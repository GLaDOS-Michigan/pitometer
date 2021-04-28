include "../DistributedSystem.i.dfy"

include "../../../Common/Collections/Maps2.i.dfy"
include "../Constants.i.dfy"
include "../Environment.i.dfy"
include "../Replica.i.dfy"

include "TimestampedRslSystem.i.dfy"

include "../CommonProof/Constants.i.dfy"

module RslPhase1Invariants_i {
import opened CommonProof__Constants_i
import opened TimestampedRslSystem_i

predicate BoundedQueueingAssumption(s:TimestampedRslState)
  requires RslConsistency(s)
{
  forall idx, ios :: (
    && 0 <= idx < |s.constants.config.replica_ids|
    && s.t_environment.nextStep == LEnvStepHostIos(s.constants.config.replica_ids[idx], ios, RslStep(s.t_replicas[idx].v.nextActionIndex))
    ==>
    (forall io | io in s.t_environment.nextStep.ios && io.LIoOpReceive? ::
      // this means that max(replica.ts, msg.ts) <= msg.ts + MaxQueueTime
      s.t_replicas[idx].ts <= io.r.msg.ts + MaxQueueTime
    )
  )
}

predicate RslAssumption(s:TimestampedRslState)
{
  && NoPacketDuplication(s)
  && |s.t_replicas| > 0

  && (var nextStep := s.t_environment.nextStep; nextStep.LEnvStepHostIos? ==>
  (forall io :: io in nextStep.ios && io.LIoOpReceive? ==> !io.r.msg.v.RslMessage_Heartbeat?))
  && ViewAlwaysZero(s)
  && BoundedQueueingAssumption(s)
}

predicate RslConsistency(s:TimestampedRslState)
{
  ConstantsAllConsistentInv(UntimestampRslState(s))
    && WellFormedLConfiguration(s.constants.config)
}

predicate ViewAlwaysZero(s:TimestampedRslState)
{
  && (forall idx:: 0 <= idx < |s.t_replicas|
    ==>
    s.t_replicas[idx].v.replica.proposer.election_state.current_view == Ballot(1, 0)
    )
}

// FIXME: this should be part of assumptions
predicate AlwaysInvariant(s:TimestampedRslState)
{
  && (forall pkt :: pkt in s.undeliveredPackets ==>
    pkt in s.t_environment.sentPackets
    && pkt.src in s.constants.config.replica_ids
    && pkt.dst in s.constants.config.replica_ids)
}

predicate InitInvariant(s:TimestampedRslState)
  requires RslConsistency(s)
{
  && AlwaysInvariant(s)

  && (0 < |s.t_replicas|
    ==>
    var r := s.t_replicas[0].v.replica.proposer;
    (if s.t_replicas[0].v.nextActionIndex == 0 then
      && TimeEq(s.t_replicas[0].ts, TimeZero())
    else
      && s.t_replicas[0].v.nextActionIndex == 1
      && TimeLe(s.t_replicas[0].ts, Timeout() + ProcessPacket)
    )
  )

  && (0 < |s.t_replicas| ==>
  var p := s.t_replicas[0].v.replica.proposer;
    && p.current_state == 0
    && p.election_state.current_view == Ballot(1, 0)
    && p.max_ballot_i_sent_1a == Ballot(0, 0)
    && |p.received_1b_packets| == 0
  )

  && (forall idx :: 0 <= idx < |s.t_replicas|
  ==> s.t_replicas[idx].v.replica.acceptor.max_bal == Ballot(0, 0)
  && s.t_replicas[idx].v.replica.proposer.current_state == 0
  && s.t_replicas[idx].v.replica.executor.next_op_to_execute == OutstandingOpUnknown()
  && s.t_replicas[idx].v.replica.learner.unexecuted_learner_state == map[]
  )

  && (forall pkt :: pkt in s.undeliveredPackets
  ==> false
  )
}

predicate TriggerPacket(pkt:TimestampedRslPacket)
{
  true
}

predicate Phase1PacketInvariant(undeliveredPackets:UndeliveredPackets, constants:LConstants)
{
  && (forall pkt :: && pkt in undeliveredPackets
  ==>
   || pkt.msg.v.RslMessage_1a?
   || pkt.msg.v.RslMessage_1b?
  )

  && (forall pkt {:trigger pkt.msg.v.RslMessage_1b?} :: && pkt in undeliveredPackets
      && pkt.msg.v.RslMessage_1b?
  ==>
     && 0 < |constants.config.replica_ids|
     && pkt.dst == constants.config.replica_ids[0]
     && TimeLe(pkt.msg.ts, TimeBound1bDelivery())
     && pkt.msg.v.bal_1b == Ballot(1, 0)
  )

  && (forall pkt {:trigger pkt.msg.v.RslMessage_1a?} :: && pkt in undeliveredPackets
      && pkt.msg.v.RslMessage_1a?
   ==> TimeLe(pkt.msg.ts, TimeBound1aDelivery())
      && 0 < |constants.config.replica_ids|
      && pkt.msg.v.bal_1a == Ballot(1, 0)
      && pkt.src == constants.config.replica_ids[0]
  )
}

predicate Phase1GenericInvariant(s:TimestampedRslState)
  requires RslConsistency(s)
{
  && AlwaysInvariant(s)
  && Phase1PacketInvariant(s.undeliveredPackets, s.constants)

  && (forall pkt {:trigger pkt in s.t_replicas[0].v.replica.proposer.received_1b_packets}:: 0 < |s.t_replicas| && pkt in s.t_replicas[0].v.replica.proposer.received_1b_packets
      ==> pkt.src in s.constants.config.replica_ids
  )

  && (forall idx :: && 0 <= idx < |s.t_replicas|
    ==>
  && s.t_replicas[idx].v.replica.executor.next_op_to_execute == OutstandingOpUnknown()
  && s.t_replicas[idx].v.replica.learner.unexecuted_learner_state == map[]
  )

  && (forall idx :: && 0 < idx < |s.t_replicas|
  ==>
  && s.t_replicas[idx].v.replica.proposer.current_state == 0
  )

  && (0 < |s.t_replicas| ==>
    // && s.t_replicas[0].v.replica.proposer.current_state == 1
    && s.t_replicas[0].v.replica.proposer.max_ballot_i_sent_1a == s.t_replicas[0].v.replica.proposer.election_state.current_view == Ballot(1, 0)
  )

  && (0 < |s.t_replicas| ==>
    // && s.t_replicas[0].v.replica.proposer.current_state == 1
    var num_2bs := |s.t_replicas[0].v.replica.proposer.received_1b_packets|;
    num_2bs == LMinQuoromSize(s.constants.config) ==>
    // By time time we've entered phase 2, we won't be in this invariant anymore
    TimeLe(s.t_replicas[0].ts, TimeBoundPhase1Leader())
  )
}

predicate PreparedLeaderPhase1TimeBound(leader:TimestampedLScheduler, undeliveredPackets:UndeliveredPackets, leader_id:NodeIdentity)
{
  && leader.v.replica.proposer.current_state == 1
    && leader.v.replica.acceptor.max_bal == leader.v.replica.proposer.max_ballot_i_sent_1a
    && 0 <= leader.v.nextActionIndex < 10
}

predicate Phase1Invariant(s:TimestampedRslState)
  requires RslConsistency(s)
{
  && Phase1GenericInvariant(s)
}

predicate Phase2Invariant(s:TimestampedRslState)
  requires RslConsistency(s)
{
  && (0 < |s.t_replicas| ==>
    && s.t_replicas[0].v.replica.proposer.current_state == 2
  )
}


}
