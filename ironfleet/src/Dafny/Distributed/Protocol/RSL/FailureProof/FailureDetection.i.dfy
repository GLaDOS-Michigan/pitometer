include "../DistributedSystem.i.dfy"

include "../../../Common/Collections/Maps2.i.dfy"
include "../Constants.i.dfy"
include "../Environment.i.dfy"
include "../Replica.i.dfy"

include "TimestampedRslSystem.i.dfy"
include "FailureHelpers.i.dfy"

include "../CommonProof/Constants.i.dfy"

module FailureDetection_i {
import opened TimestampedRslSystem_i
import opened FailureHelpers_i

predicate RslConsistency(s:TimestampedRslState)
{
  ConstantsAllConsistentInv(UntimestampRslState(s))
    && WellFormedLConfiguration(s.constants.config)
}

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

predicate ClockAssumption(s:TimestampedRslState, s':TimestampedRslState)
  requires RslConsistency(s)
{
  forall idx, ios :: (
    && 0 <= idx < |s.constants.config.replica_ids|
    && s.t_environment.nextStep == LEnvStepHostIos(s.constants.config.replica_ids[idx], ios, RslStep(s.t_replicas[idx].v.nextActionIndex))
    ==>
    (forall io | io in s.t_environment.nextStep.ios && io.LIoOpReadClock? ::
    && io.t >= 0
    && TimeLe(s.t_replicas[idx].ts, io.t)
    && TimeLe(io.t, s'.t_replicas[idx].ts)
    )
  )
}

predicate RslAssumption(s:TimestampedRslState)
{
  && RslConsistency(s)
  && BoundedQueueingAssumption(s)
}

predicate RslAssumption2(s:TimestampedRslState, s':TimestampedRslState)
{
  && RslAssumption(s)
  && RslAssumption(s')
  && ClockAssumption(s, s')
}

// Queuing delay invariant for epoch timeout
// This should be inductive all on its own
predicate EpochTimeoutQDInv(s:TimestampedRslState)
{
  && (forall idx :: && 0 <= idx < |s.t_replicas| ==>
    s.t_replicas[idx].v.replica.proposer.election_state.epoch_end_time >= 0 // so it's a valid Timestamp
    && TimeLe(s.t_replicas[idx].ts, s.t_replicas[idx].v.replica.proposer.election_state.epoch_end_time + EpochQD(s.t_replicas[idx].v.nextActionIndex))
    )
}

lemma EpochTimeoutQDInductive(s:TimestampedRslState, s':TimestampedRslState, j:int)
  requires RslAssumption2(s, s')
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires EpochTimeoutQDInv(s)
  ensures EpochTimeoutQDInv(s')
{
  var r' := s'.t_replicas[j];
  if s.t_environment.nextStep.nodeStep != RslStep(7) {
    assert r'.v.replica.proposer.election_state.epoch_end_time >= 0; // so it's a valid Timestamp
    assert TimeLe(r'.ts, r'.v.replica.proposer.election_state.epoch_end_time + EpochQD(r'.v.nextActionIndex));
  } else {
    // XXX: involves reasoning about the clock
    // FIXME: prove these asserts
    assert r'.v.replica.proposer.election_state.epoch_end_time >= 0; // so it's a valid Timestamp
    // FIXME: prove these asserts
    assert
    TimeLe(r'.ts, r'.v.replica.proposer.election_state.epoch_end_time + EpochQD(r'.v.nextActionIndex));
  }
}

predicate NodeIsNotSuspector(s:TimestampedRslState, j:int)
{
  // all HBs are unsuspecting
  // Either in first epoch, or in second epoch
  // TODO: for now, assume we're in the second epoch
  var suspectors := s.t_replicas[j].v.replica.proposer.election_state.current_view_suspectors;
  s.t_replicas[j].v.replica.constants.my_index !in suspectors &&
  HBUnsent(s, j)
  // Not a suspector ourselves, and no heartbeats sent indicating that we are
  // one.
}

predicate NodeIsSuspector(s:TimestampedRslState, j:int)
{
  // Node j is a suspector
  true
}

predicate HBUnsent(s:TimestampedRslState, j:int)
{
  // All HBs are unsuspecting (and the node is a suspector)
  forall pkt ::
  pkt in s.undeliveredPackets ==>
  pkt.msg.v.RslMessage_Heartbeat? ==>
  pkt.src == s.constants.config.replica_ids[j] ==>
  pkt.msg.v.suspicious == false
}

predicate HBEnRoute(s:TimestampedRslState, j:int)
{
  |s.constants.config.replica_ids| > 1 ==>
  // There exists an HB in route with the designated TS
  exists pkt ::
  pkt in s.undeliveredPackets &&
  pkt.msg.v.RslMessage_Heartbeat? &&
  pkt.src == s.constants.config.replica_ids[j] &&
  pkt.dst == s.constants.config.replica_ids[1] &&
  pkt.msg.v.suspicious == true &&
  pkt.msg.v.bal_heartbeat == Ballot(1, 0) && // suspicious about initial ballot
  TimeLe(pkt.msg.ts, TBFirstSuspectingHB())
}

predicate NodeIsKnownSuspector(s:TimestampedRslState, j:int)
{
  1 < |s.t_replicas| ==>
    var suspectors := s.t_replicas[1].v.replica.proposer.election_state.current_view_suspectors;
    j in suspectors
}

predicate NewLeaderInvariant(s:TimestampedRslState)
{
  1 < |s.t_replicas| ==>
    (s.t_replicas[1].v.replica.proposer.election_state.current_view == Ballot(1, 0) ==>
      (var ns := |s.t_replicas[1].v.replica.proposer.election_state.current_view_suspectors|;
      ns == LMinQuorumSize(s.constants.config)) ==>
        TimeLe(s.t_replicas[1].ts, TBFirstSuspectingHB())
    )

}

predicate NodeFDInvariant(s:TimestampedRslState, idx:int)
  requires 0 <= idx && idx < |s.t_replicas|
{
  NodeIsNotSuspector(s, idx) ||
  (NodeIsSuspector(s, idx) && (
     HBUnsent(s, idx) ||
     HBEnRoute(s, idx) ||
     NodeIsKnownSuspector(s, idx)
    )
  )
}

predicate FDInvariant(s:TimestampedRslState)
  requires RslConsistency(s)
  requires EpochTimeoutQDInv(s) // keeping this here just as reminder that the EpochQD invariant should be assumed in proving that this is inductive
{
  && (forall idx :: 0 <= idx < |s.t_replicas| ==>
    // s.t_replicas[idx].
    NodeFDInvariant(s, idx)
  )
  && NewLeaderInvariant(s)
}

// if node j != idx takes a step, then NodeIsNotSuspector(s, idx) is unaffected
lemma FDInd_noninterf(s:TimestampedRslState, s':TimestampedRslState, j:int, idx:int)
  requires RslAssumption2(s, s')
  requires EpochTimeoutQDInv(s)
  requires EpochTimeoutQDInv(s')
  requires 0 <= idx < |s.t_replicas|
  requires 0 <= j < |s.constants.config.replica_ids|;
  requires j != 1;
  requires j != idx;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires NodeIsNotSuspector(s, idx)
  ensures NodeIsNotSuspector(s', idx)
{
  assert ReplicasDistinct(s.constants.config.replica_ids, j, idx);
}


lemma FDInductive(s:TimestampedRslState, s':TimestampedRslState, j:int)
  requires RslAssumption2(s, s')
  requires EpochTimeoutQDInv(s)
  requires EpochTimeoutQDInv(s')
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires FDInvariant(s)
  ensures FDInvariant(s')
{
}

}
