include "../DistributedSystem.i.dfy"

include "../../../Common/Collections/Maps2.i.dfy"
include "../Constants.i.dfy"
include "../Environment.i.dfy"
include "../Replica.i.dfy"

include "TimestampedRslSystem.i.dfy"

include "../CommonProof/Constants.i.dfy"

module RslPhase2Proof_i {
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


// Main guarantee
predicate PerformanceGuarantee(s:TimestampedRslState, req_time:Timestamp)
{
  0 < |s.constants.config.replica_ids|
    && (forall pkt :: 
        && pkt in s.undeliveredPackets
        && pkt.msg.v.RslMessage_Reply?
        && pkt.src == s.constants.config.replica_ids[0]
        ==>
        TimeLe(pkt.msg.ts, TimeBoundReply(req_time, s.constants))
    )
}

predicate RslAssumption(s:TimestampedRslState)
{
  && NoPacketDuplication(s)
  && |s.t_replicas| > 0
  && s.constants.params.max_batch_size == 1

  && (var nextStep := s.t_environment.nextStep; nextStep.LEnvStepHostIos? ==>
  (forall io :: io in nextStep.ios && io.LIoOpReceive? ==> !io.r.msg.v.RslMessage_Heartbeat?))
  && LeaderAlwaysZero(s)
  && minD < SelfDelivery < D < 2*minD
  && ProcessPacket > 0
  && SelfDelivery + TimeActionRange(0) < D
}

predicate RslConsistency(s:TimestampedRslState)
{
  ConstantsAllConsistentInv(UntimestampRslState(s))
    && WellFormedLConfiguration(s.constants.config)
}

predicate ServersAreNotClients(s:TimestampedRslState)
{
  forall id :: id in s.constants.config.clientIds && id in s.constants.config.replica_ids
    ==> false
}

predicate RequestBatchSrcInClientIds(s:TimestampedRslState, v:RequestBatch)
{
  forall r :: r in v ==> r.client in s.constants.config.clientIds
}

predicate AlwaysInvariant(s:TimestampedRslState)
{
  && (forall pkt :: pkt in s.undeliveredPackets ==>
    && pkt in s.t_environment.sentPackets
    && pkt.src in s.constants.config.replica_ids
    // && pkt.dst in s.constants.config.replica_ids
  )

  && ServersAreNotClients(s)

  && (forall pkt :: pkt in s.undeliveredPackets ==>
    && (pkt.msg.v.RslMessage_2a? ==> RequestBatchSrcInClientIds(s, pkt.msg.v.val_2a))
    && (pkt.msg.v.RslMessage_2b? ==> RequestBatchSrcInClientIds(s, pkt.msg.v.val_2b))
  )

  && (forall idx :: 0 <= idx < |s.t_replicas|
  ==>
  && (var uls := s.t_replicas[idx].v.replica.learner.unexecuted_learner_state;
   forall opn :: opn in uls ==> RequestBatchSrcInClientIds(s, uls[opn].candidate_learned_value)
  )

  && (s.t_replicas[idx].v.replica.executor.next_op_to_execute.OutstandingOpKnown?
  ==> RequestBatchSrcInClientIds(s, s.t_replicas[idx].v.replica.executor.next_op_to_execute.v)
  )

  )
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

predicate ProposedLeaderNextOperationNumberInvariant(s:LProposer, log_truncation_point:OperationNumber)
{
  && LSetOfMessage1bAboutBallot(s.received_1b_packets, s.max_ballot_i_sent_1a)

    && 0 < s.next_operation_number_to_propose <= UpperBoundedAddition(log_truncation_point, s.constants.all.params.max_log_length, s.constants.all.params.max_integer_val)

    && (forall pkt :: pkt in s.received_1b_packets
    ==>
    && pkt.msg.log_truncation_point == 0
    && pkt.msg.votes == map[]
    )
    && |s.received_1b_packets| >= LMinQuorumSize(s.constants.all.config)
}

predicate Pre2aInvariant(s:TimestampedRslState, req_time:Timestamp, opn:OperationNumber)
  requires RslConsistency(s)
{
  && AlwaysInvariant(s)

  && (0 < |s.t_replicas|
    ==>
    |s.t_replicas[0].v.replica.proposer.request_queue| == 1
    && s.t_replicas[0].v.replica.proposer.request_queue[0].client in s.constants.config.clientIds
    // TODO: put invariants about the content of the request
    && 1 <= s.t_replicas[0].v.nextActionIndex <= 3
    && TimeLe(s.t_replicas[0].ts, s.t_replicas[0].dts + TimeActionRange(0) + TimeActionRange(s.t_replicas[0].v.nextActionIndex))
    && TimeLe(s.t_replicas[0].dts, req_time)
    && s.t_replicas[0].v.replica.learner.unexecuted_learner_state == map[]

    && LeaderNextOperationNumberInvariant(s.t_replicas[0].v.replica.proposer,
        s.t_replicas[0].v.replica.acceptor.log_truncation_point
    )

    && BalLeq(s.t_replicas[0].v.replica.learner.max_ballot_seen, Ballot(1, 0))
  )


  && (forall idx :: 0 < idx < |s.t_replicas|
  ==> LSchedulerLagTimeBound(s.t_replicas[idx])
  )

  && (forall idx :: 0 <= idx < |s.t_replicas|
  ==> 
  && s.t_replicas[idx].v.replica.executor.next_op_to_execute == OutstandingOpUnknown()
  && s.t_replicas[idx].v.replica.executor.ops_complete == opn
  && s.t_replicas[idx].v.replica.proposer.next_operation_number_to_propose == opn
  && s.t_replicas[idx].v.replica.learner.unexecuted_learner_state == map[]
  )

  && |s.undeliveredPackets| == 0
}

predicate No2aToDst(undeliveredPackets:UndeliveredPackets, dst:NodeIdentity)
{
  && (forall pkt {:trigger pkt.msg.v.RslMessage_2a?} :: pkt in undeliveredPackets
    && pkt.msg.v.RslMessage_2a?
    ==> pkt.dst != dst)
}

predicate No2bToLeaderFromSrc(undeliveredPackets:UndeliveredPackets, ldr:NodeIdentity, src:NodeIdentity)
{
  && (forall pkt {:trigger pkt.msg.v.RslMessage_2b?} :: pkt in undeliveredPackets
    && pkt.msg.v.RslMessage_2b?
    && pkt.dst == ldr
    ==> pkt.src != src)
}

datatype Phase2Progress = P2a(pkt:TimestampedRslPacket) | P2b | P2done

predicate GenericPhase2UndeliveredPacketInvariant(undeliveredPackets:UndeliveredPackets, constants:LConstants, req_time:Timestamp, opn:OperationNumber, t2a:Timestamp, progresses:map<NodeIdentity, Phase2Progress>)
  requires 0 < |constants.config.replica_ids|
{
  && (forall pkt :: pkt in undeliveredPackets && pkt.dst in constants.config.replica_ids
  ==>
  pkt.msg.v.RslMessage_2a?
  || pkt.msg.v.RslMessage_2b?
  )

  && (forall pkt {:trigger pkt.msg.v.RslMessage_2a?} :: pkt in undeliveredPackets && pkt.msg.v.RslMessage_2a?
      ==> pkt.dst in progresses && progresses[pkt.dst].P2a?
      && pkt.src == constants.config.replica_ids[0]
      && pkt.msg.v.opn_2a == opn
      && pkt.msg.v.bal_2a == Ballot(1, 0)
      && TimeLe(pkt.msg.ts, TimeBound2aDelivery(req_time))
      && (pkt.dst == pkt.src ==> TimeLe(pkt.msg.ts, TimeBound2aSelfDelivery(req_time)))

      && t2a + minD < pkt.msg.ts <= t2a + D
      && (var undeliveredPackets' := undeliveredPackets - {pkt};
      No2aToDst(undeliveredPackets', pkt.dst)
      )
  )

  && (forall pkt {:trigger pkt.msg.v.RslMessage_2b?} :: pkt in undeliveredPackets && pkt.msg.v.RslMessage_2b?
      ==>
      (pkt.dst == constants.config.replica_ids[0] ==>
      && pkt.src in progresses && progresses[pkt.src] == P2b
      && TimeLe(pkt.msg.ts, TimeBound2bDelivery(req_time))

      && (var undeliveredPackets' := undeliveredPackets - {pkt};
      No2bToLeaderFromSrc(undeliveredPackets', constants.config.replica_ids[0], pkt.src)
      )
      )

      && pkt.msg.v.opn_2b == opn
      && pkt.msg.v.bal_2b == Ballot(1, 0)
      && pkt.msg.ts > t2a + 2*minD

  )

  && (forall id {:trigger id in progresses} ::
    Progress2aProperty(progresses, id, undeliveredPackets) 
  )
}

predicate {:opaque} Progress2aProperty(progresses:map<NodeIdentity,Phase2Progress>, id:NodeIdentity, undeliveredPackets:UndeliveredPackets)
{
  id in progresses && progresses[id].P2a?
    ==>
    progresses[id].pkt in undeliveredPackets
    && progresses[id].pkt.msg.v.RslMessage_2a?
    && progresses[id].pkt.dst == id
}

lemma SelectiveProgress2aProperty(progresses:map<NodeIdentity,Phase2Progress>, id:NodeIdentity, undeliveredPackets:UndeliveredPackets)
  ensures Progress2aProperty(progresses, id, undeliveredPackets) <==> (id in progresses && progresses[id].P2a? ==> progresses[id].pkt in undeliveredPackets && progresses[id].pkt.msg.v.RslMessage_2a? && progresses[id].pkt.dst == id)
{
  reveal_Progress2aProperty(); 
}

// Invariant that is common between Phase2UnacceptedLeaderInvariant and Phase2AcceptedLeaderInvariant
predicate GenericPhase2Invariant(s:TimestampedRslState, req_time:Timestamp, opn:OperationNumber, t2a:Timestamp, progresses:map<NodeIdentity, Phase2Progress>)
  requires RslConsistency(s)
{
  && AlwaysInvariant(s)
  && (forall id :: id in progresses.Keys <==> id in s.constants.config.replica_ids)
  && (0 < |s.t_replicas|
    ==>
    |s.t_replicas[0].v.replica.proposer.request_queue| == 0

    && GenericPhase2UndeliveredPacketInvariant(s.undeliveredPackets, s.constants, req_time, opn, t2a, progresses)
  )

  && (forall idx :: && 0 < idx < |s.t_replicas|
  ==>
  // && s.t_replicas[idx].v.replica.executor.next_op_to_execute == OutstandingOpUnknown()
  && (progresses[s.constants.config.replica_ids[idx]].P2a?
      ==> LSchedulerLagTimeBound(s.t_replicas[idx])
        && s.t_replicas[idx].v.replica.learner.unexecuted_learner_state == map[]
   )
  )
}

// Main invariant part 1
// Final invariant is (Phase2UnacceptedLeaderInvariant or Phase2AcceptedLeaderInvariant)
predicate Phase2UnacceptedLeaderInvariant(s:TimestampedRslState, req_time:Timestamp, opn:OperationNumber, t2a:Timestamp, progresses:map<NodeIdentity, Phase2Progress>)
  requires RslConsistency(s)
{
  && PerformanceGuarantee(s, req_time)
  && GenericPhase2Invariant(s, req_time, opn, t2a, progresses)

  && (0 < |s.t_replicas|
    ==>
    |s.t_replicas[0].v.replica.proposer.request_queue| == 0
    && 0 <= s.t_replicas[0].v.nextActionIndex < 10

    && opn == s.t_replicas[0].v.replica.proposer.next_operation_number_to_propose - 1
    && opn == s.t_replicas[0].v.replica.executor.ops_complete
    && s.t_replicas[0].v.replica.executor.next_op_to_execute.OutstandingOpUnknown?
    && BalLeq(s.t_replicas[0].v.replica.learner.max_ballot_seen, Ballot(1, 0))

    && TimeLe(s.t_replicas[0].ts, s.t_replicas[0].dts + TimeActionRange(0) + TimeActionRange(s.t_replicas[0].v.nextActionIndex))
    && progresses[s.constants.config.replica_ids[0]].P2a?
    && s.t_replicas[0].v.replica.learner.unexecuted_learner_state == map[]
    && s.t_replicas[0].v.replica.executor.next_op_to_execute == OutstandingOpUnknown()
    && ProposedLeaderNextOperationNumberInvariant(s.t_replicas[0].v.replica.proposer, s.t_replicas[0].v.replica.acceptor.log_truncation_point)
  )
}

function Get2bCount(s:LReplica, opn:OperationNumber) : int
{
  if opn !in s.learner.unexecuted_learner_state then
    0
  else
    |s.learner.unexecuted_learner_state[opn].received_2b_message_senders|
}


// Main invariant part 2
// Final invariant is (Phase2UnacceptedLeaderInvariant or Phase2AcceptedLeaderInvariant)
predicate Phase2AcceptedLeaderInvariant(s:TimestampedRslState, req_time:Timestamp, opn:OperationNumber, t2a:Timestamp, progresses:map<NodeIdentity, Phase2Progress>)
  requires RslConsistency(s)
{
  && PerformanceGuarantee(s, req_time)
  && GenericPhase2Invariant(s, req_time, opn, t2a, progresses)

  && (0 < |s.t_replicas|
    ==>
    |s.t_replicas[0].v.replica.proposer.request_queue| == 0
    && 0 <= s.t_replicas[0].v.nextActionIndex < 10

    && opn == s.t_replicas[0].v.replica.proposer.next_operation_number_to_propose - 1
    && opn == s.t_replicas[0].v.replica.executor.ops_complete
    && Get2bCount(s.t_replicas[0].v.replica, opn) <= LMinQuorumSize(s.constants.config)

    && (Get2bCount(s.t_replicas[0].v.replica, opn) < LMinQuorumSize(s.constants.config)
    ==> s.t_replicas[0].v.replica.executor.next_op_to_execute.OutstandingOpUnknown?)

    && (s.t_replicas[0].v.nextActionIndex >= 7 || s.t_replicas[0].v.nextActionIndex == 0
    ==> Get2bCount(s.t_replicas[0].v.replica, opn) < LMinQuorumSize(s.constants.config)
    )

    && (Get2bCount(s.t_replicas[0].v.replica, opn) == LMinQuorumSize(s.constants.config)
    ==> TimeLe(s.t_replicas[0].dts, TimeBound2bDelivery(req_time)))

    && (s.t_replicas[0].v.nextActionIndex == 6
        && Get2bCount(s.t_replicas[0].v.replica, opn) == LMinQuorumSize(s.constants.config)
    ==> 
    s.t_replicas[0].v.replica.executor.next_op_to_execute.OutstandingOpKnown?
    )

    && (s.t_replicas[0].v.replica.executor.next_op_to_execute.OutstandingOpKnown?
    ==> Get2bCount(s.t_replicas[0].v.replica, opn) == LMinQuorumSize(s.constants.config)
    )

    && ProposedLeaderNextOperationNumberInvariant(s.t_replicas[0].v.replica.proposer, s.t_replicas[0].v.replica.acceptor.log_truncation_point)

    && (var num_2bs := Get2bCount(s.t_replicas[0].v.replica, opn);
     TimeLe(s.t_replicas[0].ts, TimeBoundPhase2Leader(s.t_replicas[0].dts, num_2bs + 2, s.t_replicas[0].v.nextActionIndex))
    )
    && !progresses[s.constants.config.replica_ids[0]].P2a?

    && (forall opn' :: opn' in s.t_replicas[0].v.replica.learner.unexecuted_learner_state ==>
    opn' == opn
    )

    &&(BalLeq(s.t_replicas[0].v.replica.learner.max_ballot_seen, Ballot(1, 0))
    )

    && (opn in s.t_replicas[0].v.replica.learner.unexecuted_learner_state ==>
    && s.t_replicas[0].v.replica.learner.max_ballot_seen == Ballot(1, 0)
    && (forall id :: id in s.t_replicas[0].v.replica.learner.unexecuted_learner_state[opn].received_2b_message_senders ==> id in s.constants.config.replica_ids)
      && (forall id :: id in s.t_replicas[0].v.replica.learner.unexecuted_learner_state[opn].received_2b_message_senders ==> progresses[id] == P2done)
    )

  )
}

predicate Phase2CompletedInvariant(s:TimestampedRslState, req_time:Timestamp, opn:OperationNumber)
  requires RslConsistency(s)
{
  && PerformanceGuarantee(s, req_time)
    && (0 < |s.t_replicas|
    ==>
    // s.t_replicas[0].v.replica.executor.ops_complete > opn
    s.t_replicas[0].ts <= TimeBoundPhase2Leader(TimeBound2bDelivery(req_time), LMinQuorumSize(s.constants.config) + 2, 7)
    )
}

lemma {:verify false} lemma_leader_3_Pre2aGoesToPhase2UnacceptedLeader(s:TimestampedRslState, s':TimestampedRslState, j:int, req_time:Timestamp, opn:OperationNumber) returns
  (progresses':map<NodeIdentity,Phase2Progress>, t2a:Timestamp)
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

  requires Pre2aInvariant(s, req_time, opn)
  ensures Phase2UnacceptedLeaderInvariant(s', req_time, opn, t2a, progresses')
{
  reveal_Progress2aProperty();
  var t_ios := s.t_environment.nextStep.ios;
  var ios := UntagLIoOpSeq(t_ios);
  var sent_packets := ExtractSentPacketsFromIos(ios);

  // assert exists pkt :: LIoOpSend(pkt) in t_ios;
  // assert exists pkt:TimestampedRslPacket :: pkt.msg.v.RslMessage_2a? && pkt.dst == s.constants.config.replica_ids[j] && LIoOpSend(pkt) in t_ios;

  var i := 0;
  progresses' := map[];
  
  while i < |s.constants.config.replica_ids|
    invariant 0 <= i <= |s.constants.config.replica_ids|;
    invariant (forall id :: id in progresses'.Keys <==> id in s.constants.config.replica_ids[..i]);
    invariant (forall id :: id in progresses'
      ==>
      && progresses'[id].P2a? && progresses'[id].pkt in s'.undeliveredPackets
      && progresses'[id].pkt.msg.v.RslMessage_2a? && progresses'[id].pkt.dst == id
    );
  {
    var untagged_pkt := sent_packets[i];
    var k:int :| 0 <= k < |ios| && ios[k] == LIoOpSend(sent_packets[i]);
    var pkt := t_ios[k].s;

    progresses' := progresses'[s.constants.config.replica_ids[i] := P2a(pkt)];
    i := i + 1;
  }
  t2a := s'.t_replicas[0].ts;
}

lemma lemma_notleader_notreceive_P2UnacceptedGoesToP2Unaccepted(s:TimestampedRslState, s':TimestampedRslState, j:int, req_time:Timestamp, opn:OperationNumber, t2a:Timestamp, progresses:map<NodeIdentity, Phase2Progress>)
  returns (progresses':map<NodeIdentity,Phase2Progress>)

  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires j > 0;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep != RslStep(0);
  requires s.t_environment.nextStep.nodeStep == RslStep(6);

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Phase2UnacceptedLeaderInvariant(s, req_time, opn, t2a, progresses)
  ensures Phase2UnacceptedLeaderInvariant(s', req_time, opn, t2a, progresses')
{
  reveal_Progress2aProperty();
  progresses' := progresses;
  var t_ios := s.t_environment.nextStep.ios;
  var sent_packets := ExtractSentPacketsFromIos(UntagLIoOpSeq(t_ios));
  assert ReplicasDistinct(s.constants.config.replica_ids, 0, j);


  if (LExecutorExecute.requires(s.t_replicas[j].v.replica.executor, s'.t_replicas[j].v.replica.executor, sent_packets)
  && LExecutorExecute(s.t_replicas[j].v.replica.executor, s'.t_replicas[j].v.replica.executor, sent_packets)) {
    var e := s.t_replicas[j].v.replica.executor;
    var e' := s'.t_replicas[j].v.replica.executor;
    var batch := e.next_op_to_execute.v;
    var temp := HandleRequestBatch(e.app, batch);
    var new_state := temp.0[|temp.0|-1];
    var replies := temp.1;

    assert sent_packets == GetPacketsFromReplies(e.constants.all.config.replica_ids[e.constants.my_index], batch, replies);

    forall i | 0 <= i < |sent_packets|
      ensures sent_packets[i].dst in s.constants.config.clientIds
      ensures sent_packets[i].dst !in s.constants.config.replica_ids
    {
      lemma_SizeOfGetPacketsFromReplies(s.constants.config.replica_ids[j], batch, replies, sent_packets);
      lemma_SpecificPacketInGetPacketsFromReplies(s.constants.config.replica_ids[j], batch, replies, sent_packets, i);
      assert sent_packets[i] == LPacket(batch[i].client, s.constants.config.replica_ids[j], RslMessage_Reply(batch[i].seqno, replies[i].reply));
    }
  } else {
  }
}

lemma lemma_notleader_receive_P2UnacceptedGoesToP2Unaccepted(s:TimestampedRslState, s':TimestampedRslState, j:int, req_time:Timestamp, opn:OperationNumber, t2a:Timestamp, progresses:map<NodeIdentity, Phase2Progress>)
  returns (progresses':map<NodeIdentity,Phase2Progress>)

  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires j > 0;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep == RslStep(0);

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Phase2UnacceptedLeaderInvariant(s, req_time, opn, t2a, progresses)
  ensures Phase2UnacceptedLeaderInvariant(s', req_time, opn, t2a, progresses')
{
  reveal_PacketDeliveredInOrder();
  reveal_Progress2aProperty();

  progresses' := progresses;

  var ios := s.t_environment.nextStep.ios;
  assert ReplicasDistinct(s.constants.config.replica_ids, 0, j);
  progresses' := progresses;
  SelectiveProgress2aProperty(progresses, s.constants.config.replica_ids[j], s.undeliveredPackets);
  if !progresses[s.constants.config.replica_ids[j]].P2a? {
    progresses' := progresses;
  }
  else {
    if ios[0].LIoOpReceive? {
      var pkt_2a := progresses[s.constants.config.replica_ids[j]].pkt;

      if ios[0].r.msg.v.RslMessage_2b? {
        assert ios[0].r.msg.ts <= pkt_2a.msg.ts;
        assert pkt_2a.msg.ts < t2a + D;
        assert t2a + 2*minD < ios[0].r.msg.ts;
        assert false;
      }

      assert (ios[0].r.msg.v.RslMessage_2a?);
      BoundedLagImpliesBoundedProcessingTime(s.t_replicas[j].dts, s.t_replicas[j].ts, ios[0].r.msg.ts, s'.t_replicas[j].ts, TimeActionRange(0));
      progresses' := progresses[s.constants.config.replica_ids[j] := P2b];

     // assert (forall io :: io in ios && io.LIoOpSend? && io.s.msg.v.RslMessage_2b?
     //   ==>
     //   var pkt := io.s;
     //   && pkt.msg.ts >= ios[0].r.msg.ts + minD > t2a + 2*minD
     //   );

      // forall id | true
        // ensures Progress2aProperty(progresses', id, s'.undeliveredPackets)
      // {
        // SelectiveProgress2aProperty(progresses', id, s'.undeliveredPackets);
        // SelectiveProgress2aProperty(progresses, id, s.undeliveredPackets);
      // }
      assert Phase2UnacceptedLeaderInvariant(s', req_time, opn, t2a, progresses');
    }
  }
}

lemma lemma_leader_notreceive_P2UnacceptedGoesToP2Unaccepted(s:TimestampedRslState, s':TimestampedRslState, j:int, req_time:Timestamp, opn:OperationNumber, t2a:Timestamp, progresses:map<NodeIdentity, Phase2Progress>)
  returns (progresses':map<NodeIdentity,Phase2Progress>)

  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires j == 0;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep != RslStep(0);

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Phase2UnacceptedLeaderInvariant(s, req_time, opn, t2a, progresses)
  ensures Phase2UnacceptedLeaderInvariant(s', req_time, opn, t2a, progresses')
{
  reveal_Progress2aProperty();
  progresses' := progresses;
}

lemma lemma_leader_receive_P2UnacceptedGoesToP2(s:TimestampedRslState, s':TimestampedRslState, j:int, req_time:Timestamp, opn:OperationNumber, t2a:Timestamp, progresses:map<NodeIdentity, Phase2Progress>)
  returns (progresses':map<NodeIdentity,Phase2Progress>)

  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires j == 0;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep == RslStep(0);

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Phase2UnacceptedLeaderInvariant(s, req_time, opn, t2a, progresses)
  ensures Phase2UnacceptedLeaderInvariant(s', req_time, opn, t2a, progresses') || Phase2AcceptedLeaderInvariant(s', req_time, opn, t2a, progresses')
{
  reveal_Progress2aProperty();
  progresses' := progresses;

  reveal_PacketDeliveredInOrder();
  var ios := s.t_environment.nextStep.ios;
  var pkt_2a := progresses[s.constants.config.replica_ids[0]].pkt;
  if ios[0].LIoOpReceive? {
    if ios[0].r.msg.v.RslMessage_2b? {
      assert ios[0].r.msg.ts <= pkt_2a.msg.ts;
      assert pkt_2a.msg.ts <= t2a + D;
      assert t2a + 2*minD < ios[0].r.msg.ts;
      assert false;
    }
    assert ios[0].r == pkt_2a;
    assert ios[0].r.src == s.constants.config.replica_ids[0];
    BoundedLagImpliesBoundedProcessingTime(s.t_replicas[j].dts, s.t_replicas[j].ts, ios[0].r.msg.ts, s'.t_replicas[j].ts, TimeActionRange(0) + TimeActionRange(0));
    progresses' := progresses[s.constants.config.replica_ids[0] := P2b];
    // assert s.t_replicas[0].ts <= s.t_replicas[0].dts + TimeActionRange(0) + TimeActionRange(s.t_replicas[0].v.nextActionIndex);
    assert s'.t_replicas[0].ts <= s'.t_replicas[0].dts + TimeActionRange(0) + TimeActionRange(0) + ProcessPacket;
    assert Phase2AcceptedLeaderInvariant(s', req_time, opn, t2a, progresses');

  } else {
    progresses' := progresses;
    assert Phase2UnacceptedLeaderInvariant(s', req_time, opn, t2a, progresses');
  }
}

lemma lemma_notleader_notreceive_P2GoesToP2(s:TimestampedRslState, s':TimestampedRslState, j:int, req_time:Timestamp, opn:OperationNumber, t2a:Timestamp, progresses:map<NodeIdentity, Phase2Progress>)
  returns (progresses':map<NodeIdentity,Phase2Progress>)

  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires j > 0;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep != RslStep(0);

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Phase2AcceptedLeaderInvariant(s, req_time, opn, t2a, progresses)
  ensures Phase2AcceptedLeaderInvariant(s', req_time, opn, t2a, progresses')
{
  reveal_Progress2aProperty();
  progresses' := progresses;
  var t_ios := s.t_environment.nextStep.ios;
  var sent_packets := ExtractSentPacketsFromIos(UntagLIoOpSeq(t_ios));
  assert ReplicasDistinct(s.constants.config.replica_ids, 0, j);

  if (LExecutorExecute.requires(s.t_replicas[j].v.replica.executor, s'.t_replicas[j].v.replica.executor, sent_packets)
  && LExecutorExecute(s.t_replicas[j].v.replica.executor, s'.t_replicas[j].v.replica.executor, sent_packets)) {
    var e := s.t_replicas[j].v.replica.executor;
    var e' := s'.t_replicas[j].v.replica.executor;
    var batch := e.next_op_to_execute.v;
    var temp := HandleRequestBatch(e.app, batch);
    var new_state := temp.0[|temp.0|-1];
    var replies := temp.1;

    assert sent_packets == GetPacketsFromReplies(e.constants.all.config.replica_ids[e.constants.my_index], batch, replies);

    forall i |
      0 <= i < |sent_packets|

      ensures sent_packets[i].dst in s.constants.config.clientIds
      ensures sent_packets[i].dst !in s.constants.config.replica_ids
    {
      lemma_SizeOfGetPacketsFromReplies(s.constants.config.replica_ids[j], batch, replies, sent_packets);
      lemma_SpecificPacketInGetPacketsFromReplies(s.constants.config.replica_ids[j], batch, replies, sent_packets, i);
      assert sent_packets[i].dst == batch[i].client;
      assert batch[i].client in s.constants.config.clientIds;
    }
  } else {
  }
}

lemma lemma_notleader_receive_P2GoesToP2(s:TimestampedRslState, s':TimestampedRslState, j:int, req_time:Timestamp, opn:OperationNumber, t2a:Timestamp, progresses:map<NodeIdentity, Phase2Progress>)
  returns (progresses':map<NodeIdentity,Phase2Progress>)

  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires j > 0;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep == RslStep(0);

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Phase2AcceptedLeaderInvariant(s, req_time, opn, t2a, progresses)
  ensures Phase2AcceptedLeaderInvariant(s', req_time, opn, t2a, progresses')
{
  reveal_PacketDeliveredInOrder();
  reveal_Progress2aProperty();

  progresses' := progresses;

  var ios := s.t_environment.nextStep.ios;
  assert ReplicasDistinct(s.constants.config.replica_ids, 0, j);
  progresses' := progresses;
  SelectiveProgress2aProperty(progresses, s.constants.config.replica_ids[j], s.undeliveredPackets);
  if !progresses[s.constants.config.replica_ids[j]].P2a? {
    progresses' := progresses;
  }
  else {
    if ios[0].LIoOpReceive? {
      var pkt_2a := progresses[s.constants.config.replica_ids[j]].pkt;

      if ios[0].r.msg.v.RslMessage_2b? {
        assert ios[0].r.msg.ts <= pkt_2a.msg.ts;
        assert pkt_2a.msg.ts < t2a + D;
        assert t2a + 2*minD < ios[0].r.msg.ts;
        assert false;
      }

      assert (ios[0].r.msg.v.RslMessage_2a?);
      BoundedLagImpliesBoundedProcessingTime(s.t_replicas[j].dts, s.t_replicas[j].ts, ios[0].r.msg.ts, s'.t_replicas[j].ts, TimeActionRange(0));
      progresses' := progresses[s.constants.config.replica_ids[j] := P2b];
      assert Phase2AcceptedLeaderInvariant(s', req_time, opn, t2a, progresses');
    }
  }
}

lemma lemma_leader_notreceiveOrExecute_P2GoesToP2(s:TimestampedRslState, s':TimestampedRslState, j:int, req_time:Timestamp, opn:OperationNumber, t2a:Timestamp, progresses:map<NodeIdentity, Phase2Progress>)
  returns (progresses':map<NodeIdentity,Phase2Progress>)

  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires j == 0;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep != RslStep(0);
  requires s.t_environment.nextStep.nodeStep != RslStep(6);

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Phase2AcceptedLeaderInvariant(s, req_time, opn, t2a, progresses)
  ensures Phase2AcceptedLeaderInvariant(s', req_time, opn, t2a, progresses')
{
  reveal_Progress2aProperty();
  progresses' := progresses;
}

lemma lemma_leader_execute_P2GoesToP2(s:TimestampedRslState, s':TimestampedRslState, j:int, req_time:Timestamp, opn:OperationNumber, t2a:Timestamp, progresses:map<NodeIdentity, Phase2Progress>)
  returns (progresses':map<NodeIdentity,Phase2Progress>)

  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires j == 0;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep == RslStep(6);

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Phase2AcceptedLeaderInvariant(s, req_time, opn, t2a, progresses)
  ensures Phase2AcceptedLeaderInvariant(s', req_time, opn, t2a, progresses') || Phase2CompletedInvariant(s', req_time, opn)
{
  reveal_Progress2aProperty();
  progresses' := progresses;

  if (Get2bCount(s.t_replicas[0].v.replica, opn) >= LMinQuorumSize(s.constants.config)) {

    assert Get2bCount(s.t_replicas[0].v.replica, opn) == LMinQuorumSize(s.constants.config);
    assert TimeLe(s.t_replicas[0].ts, TimeBoundPhase2Leader(s.t_replicas[0].dts, Get2bCount(s.t_replicas[0].v.replica, opn) + 2, s.t_replicas[0].v.nextActionIndex));
    assert TimeLe(s'.t_replicas[0].ts, TimeBoundPhase2Leader(s'.t_replicas[0].dts, LMinQuorumSize(s.constants.config) + 2, s'.t_replicas[0].v.nextActionIndex));
    assert Phase2CompletedInvariant(s', req_time, opn);
  } else {
    assert Phase2AcceptedLeaderInvariant(s', req_time, opn, t2a, progresses');
  }
}


lemma lemma_leader_receive_P2GoesToP2(s:TimestampedRslState, s':TimestampedRslState, j:int, req_time:Timestamp, opn:OperationNumber, t2a:Timestamp, progresses:map<NodeIdentity, Phase2Progress>)
  returns (progresses':map<NodeIdentity,Phase2Progress>)

  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires j == 0;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep == RslStep(0);

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Phase2AcceptedLeaderInvariant(s, req_time, opn, t2a, progresses)
  ensures Phase2AcceptedLeaderInvariant(s', req_time, opn, t2a, progresses')
{
  reveal_Progress2aProperty();
  progresses' := progresses;

  var ios := s.t_environment.nextStep.ios;
  if ios[0].LIoOpReceive? {
    assert ios[0].r.msg.v.RslMessage_2b?;
    var size := Get2bCount(s.t_replicas[0].v.replica, opn);
    var size' := Get2bCount(s'.t_replicas[0].v.replica, opn);
    assert size' == size + 1;
    
    BoundedSizeLagImpliesBoundedProcessingTime(s.t_replicas[0].dts,
      s.t_replicas[0].ts,
      ios[0].r.msg.ts,
      s'.t_replicas[0].ts,
      size + 2
      );
    progresses' := progresses[ios[0].r.src := P2done];
  } else {
    progresses' := progresses;
    var size := Get2bCount(s.t_replicas[0].v.replica, opn);
    LeaderTimeoutPreservesPhase2Invariant(s'.t_replicas[0].dts, size + 2, s'.t_replicas[0].v.nextActionIndex);
  }
}

}
