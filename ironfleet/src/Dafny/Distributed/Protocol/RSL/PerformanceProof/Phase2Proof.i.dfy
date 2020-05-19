include "../DistributedSystem.i.dfy"

include "../../../Common/Collections/Maps2.i.dfy"
include "../Constants.i.dfy"
include "../Environment.i.dfy"
include "../Replica.i.dfy"

include "TimestampedRslSystem.i.dfy"

include "../CommonProof/Constants.i.dfy"

module RslPhase1Proof_i {
import opened TimestampedRslSystem_i
import opened CommonProof__Constants_i

predicate LeaderAlwaysZero(s:TimestampedRslState)
{
  && (forall idx:: 0 <= idx < |s.t_replicas|
    ==>
    && s.t_replicas[idx].v.replica.proposer.election_state.current_view == Ballot(1, 0)
    && (idx > 0 ==> s.t_replicas[idx].v.replica.proposer.current_state == 0)
    && s.t_replicas[idx].v.replica.proposer.election_state.current_view == s.t_replicas[0].v.replica.proposer.max_ballot_i_sent_1a
  )

    && (0 < |s.t_replicas|
    ==>
    s.t_replicas[0].v.replica.proposer.current_state == 2
    )
}

predicate RslAssumption(s:TimestampedRslState)
{
  && NoPacketDuplication(s)
  && |s.t_replicas| > 0
  && s.constants.params.max_batch_size == 1

  // TODO: Remove this assumption
  && (var nextStep := s.t_environment.nextStep; nextStep.LEnvStepHostIos? ==>
  (forall io :: io in nextStep.ios && io.LIoOpReceive? ==> !io.r.msg.v.RslMessage_Heartbeat?))
  && LeaderAlwaysZero(s)
  && SelfDelivery < D
}

predicate RslConsistency(s:TimestampedRslState)
{
  ConstantsAllConsistentInv(UntimestampRslState(s))
    && WellFormedLConfiguration(s.constants.config)
}

predicate AlwaysInvariant(s:TimestampedRslState)
{
  && (forall pkt :: pkt in s.undeliveredPackets ==>
    pkt in s.t_environment.sentPackets
    && pkt.src in s.constants.config.replica_ids
    && pkt.dst in s.constants.config.replica_ids)
}

predicate LSchedulerLagTimeBound(tls:TimestampedLScheduler)
{
  0 <= tls.v.nextActionIndex <= 9
  && TimeLe(tls.ts, tls.dts + TimeActionRange(tls.v.nextActionIndex))
}

predicate LeaderNextOperationNumberInvariant(s:LProposer, log_truncation_point:OperationNumber)
{
  && LSetOfMessage1bAboutBallot(s.received_1b_packets, s.max_ballot_i_sent_1a)

    && 0 <= s.next_operation_number_to_propose < UpperBoundedAddition(log_truncation_point, s.constants.all.params.max_log_length, s.constants.all.params.max_integer_val)

    && (forall pkt :: pkt in s.received_1b_packets
    ==>
    && pkt.msg.log_truncation_point == 0
    && pkt.msg.votes == map[]
    )

    && |s.received_1b_packets| >= LMinQuorumSize(s.constants.all.config)

    && LtUpperBound(s.next_operation_number_to_propose, s.constants.all.params.max_integer_val)
}

predicate Pre2aInvariant(s:TimestampedRslState, req_time:Timestamp)
  requires RslConsistency(s)
{
  && AlwaysInvariant(s)

  && (0 < |s.t_replicas|
    ==>
    |s.t_replicas[0].v.replica.proposer.request_queue| == 1
    // TODO: put invariants about the content of the request
    && 1 <= s.t_replicas[0].v.nextActionIndex <= 3
    && LSchedulerLagTimeBound(s.t_replicas[0])
    && TimeLe(s.t_replicas[0].dts, req_time)

    && LeaderNextOperationNumberInvariant(s.t_replicas[0].v.replica.proposer,
        s.t_replicas[0].v.replica.acceptor.log_truncation_point
    )
  )

  && (forall idx :: 0 < idx < |s.t_replicas|
  ==> LSchedulerLagTimeBound(s.t_replicas[idx])
  )

  && (forall idx :: 0 <= idx < |s.t_replicas|
  ==> 
  && s.t_replicas[idx].v.replica.executor.next_op_to_execute == OutstandingOpUnknown()
  && s.t_replicas[idx].v.replica.learner.unexecuted_learner_state == map[]
  )

  && |s.undeliveredPackets| == 0
}

datatype Phase2Progress = P2a | P2b | P2done

predicate GenericPhase2UndeliveredPacketInvariant(undeliveredPackets:UndeliveredPackets)
{
  && (
  forall pkt :: pkt in undeliveredPackets
  ==>
  pkt.msg.v.RslMessage_2a?
  )
}
  
predicate GenericPhase2Invariant(s:TimestampedRslState, req_time:Timestamp, progresses:map<NodeIdentity, Phase2Progress>)
{
  && (0 < |s.t_replicas|
    ==>
    |s.t_replicas[0].v.replica.proposer.request_queue| == 0
    && LSchedulerLagTimeBound(s.t_replicas[0])
    // && TimeLe(s.t_replicas[0].dts, req_time)
  )

  && GenericPhase2UndeliveredPacketInvariant(s.undeliveredPackets)
}

predicate Phase2UnacceptedLeaderInvariant(s:TimestampedRslState, req_time:Timestamp, progresses:map<NodeIdentity, Phase2Progress>)
{
  && GenericPhase2Invariant(s, req_time, progresses)
}

lemma lemma_leader_3_Pre2aGoesToPhase2UnacceptedLeader(s:TimestampedRslState, s':TimestampedRslState, j:int, req_time:Timestamp) returns
  (progresses':map<NodeIdentity,Phase2Progress>)
  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires j == 0;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep == RslStep(3);

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Pre2aInvariant(s, req_time)
  ensures Phase2UnacceptedLeaderInvariant(s', req_time, map[])
{
  var pr := s.t_replicas[0].v.replica.proposer;
  var pr' := s'.t_replicas[0].v.replica.proposer;
  var t_ios := s.t_environment.nextStep.ios;
  var ios := UntagLIoOpSeq(t_ios);
  var sent_packets := ExtractSentPacketsFromIos(ios);
  var clock := SpontaneousClock(ios);
  var log_truncation_point := s.t_replicas[0].v.replica.acceptor.log_truncation_point;
  
  assert LProposerMaybeNominateValueAndSend2a(pr, pr', clock.t, log_truncation_point, sent_packets);

  assert LProposerCanNominateUsingOperationNumber(pr, log_truncation_point, pr.next_operation_number_to_propose);
  assert LAllAcceptorsHadNoProposal(pr.received_1b_packets, pr.next_operation_number_to_propose);
}

}
