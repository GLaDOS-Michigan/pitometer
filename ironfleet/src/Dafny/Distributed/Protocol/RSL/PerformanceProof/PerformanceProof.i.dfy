include "../DistributedSystem.i.dfy"

include "../../../Common/Collections/Maps2.i.dfy"
include "../Constants.i.dfy"
include "../Environment.i.dfy"
include "../Replica.i.dfy"

include "TimestampedRslSystem.i.dfy"

include "../CommonProof/Constants.i.dfy"

module RslPerformanceProof_i {
import opened TimestampedRslPerformanceProof_i
import opened CommonProof__Constants_i

predicate RslAssumption(s:TimestampedRslState)
{
  && NoPacketDuplication(s)
  && |s.t_replicas| > 0

  // TODO: Remove this assumption
  && (var nextStep := s.t_environment.nextStep; nextStep.LEnvStepHostIos? ==>
  (forall io :: io in nextStep.ios && io.LIoOpReceive? ==> !io.r.msg.v.RslMessage_Heartbeat?))
  && ViewAlwaysZero(s)
  && SelfDelivery <= D
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

predicate InitInvariant(s:TimestampedRslState)
  requires RslConsistency(s)
{
  && AlwaysInvariant(s)

  && (0 < |s.t_replicas|
    ==>
    var r := s.t_replicas[0].v.replica.proposer;
    (if s.t_replicas[0].v.nextActionIndex == 0 then
      && TimeEq(s.t_replicas[0].ts, TimeZero())
      && TimeLe(s.t_replicas[0].dts, TimeZero())
    else
      && s.t_replicas[0].v.nextActionIndex == 1
      && TimeLe(s.t_replicas[0].ts, Timeout() + ProcessPacket)
      && TimeEq(s.t_replicas[0].dts, Timeout())
    )
  )

  && (forall idx :: 0 < idx < |s.t_replicas|
  ==> LSchedulerLagTimeBound(s.t_replicas[idx])
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
  ==> pkt.msg.v.RslMessage_Heartbeat?
  )
}

predicate No1aToDst(undeliveredPackets:UndeliveredPackets, dst:NodeIdentity)
{
  && (forall pkt :: pkt in undeliveredPackets
    && pkt.msg.v.RslMessage_1a?
    ==> pkt.dst != dst)
}

//predicate TriggerPacket(pkt:TimestampedRslPacket)
//{
  //true
//}
//
//predicate TriggerIdx(idx:int)
//{
  //true
//}

predicate Phase1GenericUndeliveredPacketsInvariant(undeliveredPackets:UndeliveredPackets, constants:LConstants)
{
  && (forall pkt :: && pkt in undeliveredPackets
  ==>
   || pkt.msg.v.RslMessage_Heartbeat?
   || pkt.msg.v.RslMessage_1a?
   || pkt.msg.v.RslMessage_1b?
  )

  && (forall pkt :: && pkt in undeliveredPackets
  && pkt.msg.v.RslMessage_1b?
  && 0 < |constants.config.replica_ids|
  ==> pkt.dst == constants.config.replica_ids[0]
  && TimeLe(pkt.msg.ts, TimeBound1bDelivery())
  )

  && (forall pkt :: && pkt in undeliveredPackets
      && pkt.msg.v.RslMessage_1a?
      && 0 < |constants.config.replica_ids|
   ==> TimeLe(pkt.msg.ts, TimeBound1aDelivery())
      && pkt.msg.v.bal_1a == Ballot(1, 0)
      && pkt.src == constants.config.replica_ids[0]

      && (var undeliveredPackets' := undeliveredPackets - {pkt};
          No1aToDst(undeliveredPackets', pkt.dst)
        )
  )
}

predicate Phase1GenericInvariant(s:TimestampedRslState)
  requires RslConsistency(s)
{
  && AlwaysInvariant(s)
  && Phase1GenericUndeliveredPacketsInvariant(s.undeliveredPackets, s.constants)

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
    && s.t_replicas[0].v.replica.proposer.current_state == 1
    && s.t_replicas[0].v.replica.proposer.max_ballot_i_sent_1a == s.t_replicas[0].v.replica.proposer.election_state.current_view
  )

  && (0 < |s.t_replicas|
  ==> forall pkt :: pkt in s.undeliveredPackets
  ==> UntagLPacket(pkt) !in s.t_replicas[0].v.replica.proposer.received_1b_packets
  )

  && (forall idx :: && 0 < idx < |s.t_replicas|
  ==>
  if s.t_replicas[idx].v.replica.acceptor.max_bal == Ballot(0, 0) then
    LSchedulerLagTimeBound(s.t_replicas[idx])
  else
    (
    No1aToDst(s.undeliveredPackets, s.constants.config.replica_ids[idx])
    )
  )
}

predicate UnpreparedLeaderPhase1TimeBound(leader:TimestampedLScheduler, undeliveredPackets:UndeliveredPackets, leader_id:NodeIdentity, pkt:TimestampedRslPacket)
{
  && leader.v.replica.proposer.current_state == 1
    && leader.v.replica.acceptor.max_bal == Ballot(0, 0)
    && leader.v.replica.proposer.max_ballot_i_sent_1a == Ballot(1, 0)
    && LSchedulerLagTimeBound(leader)
    && pkt in undeliveredPackets
    && pkt.msg.v == RslMessage_1a(leader.v.replica.proposer.max_ballot_i_sent_1a)
    && pkt.src == pkt.dst == leader_id
    && TimeLe(pkt.msg.ts, TimeBound1aSelfDelivery())
    && No1aToDst(undeliveredPackets - {pkt}, leader_id)
    && |leader.v.replica.proposer.received_1b_packets| == 0
}

predicate PreparedLeaderPhase1TimeBound(leader:TimestampedLScheduler, undeliveredPackets:UndeliveredPackets, leader_id:NodeIdentity)
{
  && leader.v.replica.proposer.current_state == 1
    && leader.v.replica.acceptor.max_bal == leader.v.replica.proposer.max_ballot_i_sent_1a
    && No1aToDst(undeliveredPackets, leader_id)
    && 0 <= leader.v.nextActionIndex < 10
    && (var ell := |leader.v.replica.proposer.received_1b_packets|;
        TimeLe(leader.ts, TimeBoundPhase1Leader(leader.dts, ell, leader.v.nextActionIndex))
    )
}

predicate Phase1UnpreparedLeaderInvariant(s:TimestampedRslState, leader_1a_pkt:TimestampedRslPacket)
  requires RslConsistency(s)
{
  && (0 < |s.t_replicas| ==>
    && UnpreparedLeaderPhase1TimeBound(s.t_replicas[0], s.undeliveredPackets, s.constants.config.replica_ids[0], leader_1a_pkt)
  )
  && (forall pkt :: && pkt in s.undeliveredPackets && pkt.msg.v.RslMessage_1b?
  ==> pkt.src in s.constants.config.replica_ids[1..]
  && TimeBound1aSelfDelivery() < pkt.msg.ts
  )

  && Phase1GenericInvariant(s)
}

predicate Phase1PreparedLeaderInvariant(s:TimestampedRslState)
  requires RslConsistency(s)
{
  && (0 < |s.t_replicas| ==>
    && PreparedLeaderPhase1TimeBound(s.t_replicas[0], s.undeliveredPackets, s.constants.config.replica_ids[0])
  )
  && (forall pkt :: pkt in s.undeliveredPackets && pkt.msg.v.RslMessage_1b?
  ==> pkt.src in s.constants.config.replica_ids[0..]
  )
  && Phase1GenericInvariant(s)
}

predicate Phase2Invariant(s:TimestampedRslState)
  requires RslConsistency(s)
{
  && (0 < |s.t_replicas| ==>
    && s.t_replicas[0].v.replica.proposer.current_state == 2
  )
}

lemma lemma_RslInitImpliesInitInv(s:TimestampedRslState)
  requires RslAssumption(s);
  requires TimestampedRslInit(s.constants, s);
  ensures InitInvariant(s);
{
}

lemma lemma_0_1_InitGoesToPhase1NoLeader1b(s:TimestampedRslState, s':TimestampedRslState, j:int) returns (leader_1a_pkt:TimestampedRslPacket)
  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 == j < |s.constants.config.replica_ids|

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep == RslStep(1);

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires InitInvariant(s)
  ensures Phase1UnpreparedLeaderInvariant(s', leader_1a_pkt)
{
  var t_ios := s.t_environment.nextStep.ios;
  var ios := UntagLIoOpSeq(t_ios);
  var sent_packets := ExtractSentPacketsFromIos(ios);
  assert LIoOpSend(sent_packets[0]) in ios;
  var i :| 0 <= i < |ios| && ios[i] == LIoOpSend(sent_packets[0]);
  leader_1a_pkt := t_ios[i].s;
}

lemma lemma_InitGoesToInitOrPhase1NoLeader1b(s:TimestampedRslState, s':TimestampedRslState, j:int)
  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires !(j == 0 && s.t_environment.nextStep.nodeStep == RslStep(1));
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires InitInvariant(s)
  ensures InitInvariant(s')
{
}

lemma lemma_notleader_nonetwork_Phase1UnpreparedLeaderGoesToPhase1UnpreparedLeader(s_prev:TimestampedRslState, s:TimestampedRslState, j:int, pkt:TimestampedRslPacket)
  requires RslAssumption(s_prev);
  requires RslAssumption(s);
  requires RslConsistency(s_prev);
  requires RslConsistency(s);
  requires 0 <= j < |s_prev.constants.config.replica_ids|

  requires j > 0;

  requires s_prev.t_environment.nextStep.LEnvStepHostIos?;
  requires s_prev.t_environment.nextStep.actor == s_prev.constants.config.replica_ids[j];
  requires s_prev.t_environment.nextStep.nodeStep != RslStep(0)
  // requires s_prev.t_environment.nextStep.nodeStep != RslStep(2)
  // requires s_prev.t_environment.nextStep.nodeStep == RslStep(3)

  requires TimestampedRslNextOneReplica(s_prev, s, j, s_prev.t_environment.nextStep.ios);

  requires Phase1UnpreparedLeaderInvariant(s_prev, pkt);
  ensures Phase1UnpreparedLeaderInvariant(s, pkt);
{
  assert ReplicasDistinct(s_prev.constants.config.replica_ids, 0, j);
}

lemma lemma_notleader_0_Phase1UnpreparedLeaderGoesToPhase1UnpreparedLeader(s:TimestampedRslState, s':TimestampedRslState, j:int, pkt:TimestampedRslPacket)
  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|

  requires j > 0;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep == RslStep(0)

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Phase1UnpreparedLeaderInvariant(s, pkt);
  ensures Phase1UnpreparedLeaderInvariant(s', pkt);
{
  var ios := s.t_environment.nextStep.ios;
  assert ReplicasDistinct(s.constants.config.replica_ids, 0, j);
  if ios[0].LIoOpReceive? {
    assert ios[0].r.msg.v.RslMessage_1a?;
    assert ios[0].r.src == s.constants.config.replica_ids[0];
    BoundedLagImpliesBoundedProcessingTime(s.t_replicas[j].dts, s.t_replicas[j].ts, ios[0].r.msg.ts, s'.t_replicas[j].ts, TimeActionRange(0));
  }
}

lemma lemma_leader_0_Phase1UnpreparedLeaderGoesToPhase1(s:TimestampedRslState, s':TimestampedRslState, j:int, pkt:TimestampedRslPacket)
  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|

  requires j == 0

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep == RslStep(0)

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Phase1UnpreparedLeaderInvariant(s, pkt)
  ensures Phase1UnpreparedLeaderInvariant(s', pkt) || Phase1PreparedLeaderInvariant(s')
{
  reveal_PacketDeliveredInOrder();
  var ios := s.t_environment.nextStep.ios;
  if ios[0].LIoOpReceive? {
    // assert s.t_replicas[j].v.replica.acceptor.max_bal == Ballot(0, 0);
    if ios[0].r.msg.v.RslMessage_1b? {
      assert ios[0].r.msg.ts <= pkt.msg.ts;
      assert pkt.msg.ts <= TimeBound1aSelfDelivery();
      assert TimeBound1aSelfDelivery() < ios[0].r.msg.ts;
      assert false;
    }
    assert ios[0].r == pkt;
    assert ios[0].r.src == s.constants.config.replica_ids[0];
    BoundedLagImpliesBoundedProcessingTime(s.t_replicas[j].dts, s.t_replicas[j].ts, ios[0].r.msg.ts, s'.t_replicas[j].ts, TimeActionRange(0));
    assert Phase1PreparedLeaderInvariant(s');
  } else {
    assert Phase1UnpreparedLeaderInvariant(s', pkt);
  }
}

/*
lemma lemma_a_j_Phase1GoesToPhase1(s:TimestampedRslState, s':TimestampedRslState, j:int)
  requires RslAssumption(s);
  requires ViewAlwaysZero(s)
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|

  requires j > 0;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep != RslStep(0)

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Phase1Invariant(s)
  ensures Phase1Invariant(s')
{
  assert ReplicasDistinct(s.constants.config.replica_ids, 0, j);
}

lemma lemma_0_0_Phase1GoesToPhase1(s:TimestampedRslState, s':TimestampedRslState, j:int)
  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|

  requires j == 0;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires || s.t_environment.nextStep.nodeStep == RslStep(0);

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Phase1Invariant(s)
  ensures Phase1Invariant(s')
{
  assert forall idx :: 0 < idx < |s'.t_replicas| ==> ReplicasDistinct(s.constants.config.replica_ids, 0, idx);
  if (s.t_environment.nextStep.nodeStep == RslStep(0)) {
    var ios := s.t_environment.nextStep.ios;
    if (ios[0].LIoOpReceive?) {
      if ios[0].r.msg.v.RslMessage_1b? {
        var size := |s.t_replicas[0].v.replica.proposer.received_1b_packets|;
        var size' := |s'.t_replicas[0].v.replica.proposer.received_1b_packets|;
        assert UntagLPacket(ios[0].r) !in s.t_replicas[0].v.replica.proposer.received_1b_packets;
        assert s'.t_replicas[0].v.replica.proposer.received_1b_packets == s.t_replicas[0].v.replica.proposer.received_1b_packets + {UntagLPacket(ios[0].r)};
        assert size' == size + 1;
        BoundedSizeLagImpliesBoundedProcessingTime(s.t_replicas[j].dts, s.t_replicas[j].ts, ios[0].r.msg.ts, s'.t_replicas[j].ts, size);
      } else if (ios[0].r.msg.v.RslMessage_1a?) {
      } else{
      }
    } else {
    }
  } else if (s.t_environment.nextStep.nodeStep == RslStep(1)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(2)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(3)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(4)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(5)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(6)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(7)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(8)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(9)) {
  }
}

lemma lemma_a_0_Phase1GoesToPhase1(s:TimestampedRslState, s':TimestampedRslState, j:int)
  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|

  requires j == 0;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires || s.t_environment.nextStep.nodeStep != RslStep(2);
  requires || s.t_environment.nextStep.nodeStep != RslStep(0);

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Phase1Invariant(s)
  ensures Phase1Invariant(s')
{
  assert forall idx :: 0 < idx < |s'.t_replicas| ==> ReplicasDistinct(s.constants.config.replica_ids, 0, idx);
  if (s.t_environment.nextStep.nodeStep == RslStep(0)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(1)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(2)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(3)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(4)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(5)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(6)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(7)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(8)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(9)) {
  }
}

lemma lemma_0_0_Phase1GoesToPhase1OrPhase2(s:TimestampedRslState, s':TimestampedRslState, j:int)
  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|

  requires j == 0;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep == RslStep(2);

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Phase1Invariant(s)
  ensures Phase1Invariant(s') // || Phase1Invariant(s')
{
}

*/

}
