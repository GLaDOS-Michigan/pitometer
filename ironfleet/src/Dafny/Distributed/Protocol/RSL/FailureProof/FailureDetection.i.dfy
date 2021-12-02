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

// Hearbeat delay invariant; self-contained
predicate HeartbeatDelayInv(s:TimestampedRslState)
{
  && (forall idx :: && 0 <= idx < |s.t_replicas| ==>
    s.t_replicas[idx].v.replica.nextHeartbeatTime >= 0 // so it's a valid Timestamp
    && TimeLe(s.t_replicas[idx].v.replica.nextHeartbeatTime, s.t_replicas[idx].ts + HBPeriod)
    )
}

////////////////////////////////////////////////////////////////////////////////
// Main invariants
////////////////////////////////////////////////////////////////////////////////

// can enter this after processing request
predicate NodeFinishedRequest(s:TimestampedRslState, j:int)
  requires RslConsistency(s)
  requires j < |s.t_replicas|
  requires j < |s.constants.config.replica_ids|
{
  s.t_replicas[j].v.replica.proposer.election_state.requests_received_this_epoch == []
  s.t_replicas[j].v.replica.proposer.election_state.requests_received_prev_epochs == []
}

predicate NodeIsNotSuspectorInFirstEpoch(s:TimestampedRslState, j:int)
  requires RslConsistency(s)
  requires j < |s.t_replicas|
  requires j < |s.constants.config.replica_ids|
{
  // all HBs are unsuspecting
  var suspectors := s.t_replicas[j].v.replica.proposer.election_state.current_view_suspectors;
  s.t_replicas[j].v.replica.constants.my_index !in suspectors &&
  HBUnsent(s, j)
  // Not a suspector ourselves, and no heartbeats sent indicating that we are one.
  && |s.t_replicas[j].v.replica.proposer.election_state.requests_received_this_epoch| == 1
  && |s.t_replicas[j].v.replica.proposer.election_state.requests_received_prev_epochs| == 0
  && true
  // TODO: can't have exceeded FailTime + EpochLength + AllActions()
}

predicate NodeIsNotSuspectorInSecondEpoch(s:TimestampedRslState, j:int)
  requires RslConsistency(s)
  requires j < |s.t_replicas|
  requires j < |s.constants.config.replica_ids|
{
  // all HBs are unsuspecting
  var suspectors := s.t_replicas[j].v.replica.proposer.election_state.current_view_suspectors;
  s.t_replicas[j].v.replica.constants.my_index !in suspectors &&
  HBUnsent(s, j)
  && s.t_replicas[j].v.replica.proposer.election_state.requests_received_prev_epochs != []
  // Not a suspector ourselves, and no heartbeats sent indicating that we are
  // one.
  // TODO: can't have exceeded (FailTime + EpochLength + AllActions()) + EpochLength + AllActions()
}

predicate NodeIsNotSuspector(s:TimestampedRslState, j:int)
  requires RslConsistency(s)
  requires j < |s.t_replicas|
  requires j < |s.constants.config.replica_ids|
{
  NodeIsNotSuspectorInFirstEpoch(s, j) ||
  NodeIsNotSuspectorInSecondEpoch(s, j)
}

predicate NodeIsSuspector(s:TimestampedRslState, j:int)
  requires RslConsistency(s)
  requires j < |s.constants.config.replica_ids|
{
  // Node j is a suspector
  var suspectors := s.t_replicas[j].v.replica.proposer.election_state.current_view_suspectors;
  s.t_replicas[j].v.replica.constants.my_index in suspectors
}

predicate HBUnsent(s:TimestampedRslState, j:int)
  requires RslConsistency(s)
  requires j < |s.constants.config.replica_ids|
{
  // All HBs are unsuspecting
  forall pkt ::
  pkt in s.undeliveredPackets ==>
  pkt.msg.v.RslMessage_Heartbeat? ==>
  pkt.src == s.constants.config.replica_ids[j] ==>
  pkt.msg.v.suspicious == false
}

predicate HBEnRoute(s:TimestampedRslState, j:int)
  requires RslConsistency(s)
  requires j < |s.constants.config.replica_ids|
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
  requires RslConsistency(s)
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
      // TODO: add nextAction inv here
      ns == LMinQuorumSize(s.constants.config)) ==>
        TimeLe(s.t_replicas[1].ts, TBFirstSuspectingHB())
    )
}

predicate HBtimeInv(s:TimestampedRslState, idx:int)
  requires RslConsistency(s)
  requires 0 <= idx && idx < |s.t_replicas|
{
  s.t_replicas[idx].v.replica.nextHeartbeatTime >= 0
  && TimeLe(s.t_replicas[idx].v.replica.nextHeartbeatTime, 0)
  // FIXME: put real upper bound here
}

predicate NodeFDInvariant(s:TimestampedRslState, idx:int)
  requires RslConsistency(s)
  requires 0 <= idx && idx < |s.t_replicas|
{
  NodeIsNotSuspector(s, idx) ||
  (NodeIsSuspector(s, idx) && (
    (HBUnsent(s, idx) && HBtimeInv(s,idx))
    || HBEnRoute(s, idx)
    || NodeIsKnownSuspector(s, idx)
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

////////////////////////////////////////////////////////////////////////////////
// PF_NONSUSP
////////////////////////////////////////////////////////////////////////////////


lemma NonSuspectorFirst_ind(s:TimestampedRslState, s':TimestampedRslState, j:int, idx:int)
  requires RslAssumption2(s, s')
  requires EpochTimeoutQDInv(s)
  requires EpochTimeoutQDInv(s')
  requires 0 <= idx < |s.t_replicas|
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep != RslStep(7) // on step 7, we might start new epoch

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires NodeIsNotSuspectorInFirstEpoch(s, j);
  ensures  NodeIsNotSuspectorInFirstEpoch(s', j);
{
  if s.t_environment.nextStep.nodeStep == RslStep(0) {
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(1) {
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(2) {
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(3) {
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(4) {
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(5) {
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(6) {
    if s.t_replicas[j].v.replica.proposer.election_state.requests_received_this_epoch == [] {
      assert NodeFinishedRequest(s, j);
    }
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(7) {
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(8) {
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(9) {
  }
}

lemma NonSuspector_ind(s:TimestampedRslState, s':TimestampedRslState, j:int, idx:int)
  requires RslAssumption2(s, s')
  requires EpochTimeoutQDInv(s)
  requires EpochTimeoutQDInv(s')
  requires 0 <= idx < |s.t_replicas|
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep != RslStep(7) // on step 7, we might become a suspector

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires NodeIsNotSuspector(s, j);
  ensures  NodeIsNotSuspector(s', j);
{
  // var suspectors := s.t_replicas[j].v.replica.proposer.election_state.current_view_suspectors;
  // s.t_replicas[j].v.replica.constants.my_index !in suspectors &&
  // HBUnsent(s, j)
  assert ReplicasDistinct(s.constants.config.replica_ids, j, 1);

  if s.t_environment.nextStep.nodeStep == RslStep(0) {
    var suspectors := s.t_replicas[j].v.replica.proposer.election_state.current_view_suspectors;
    var suspectors' := s'.t_replicas[j].v.replica.proposer.election_state.current_view_suspectors;
    assert suspectors' == suspectors || suspectors' == {}; // FIXME: this should follow from ElectionStateProcessHeartbeat
    assert suspectors' <= suspectors;
    assert suspectors' <= suspectors;
    assert j == s'.t_replicas[j].v.replica.constants.my_index;
    assert j !in suspectors;
    assert j !in suspectors';
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(1) {
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(2) {
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(3) {
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(4) {
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(5) {
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(6) {
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(7) {
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(8) {
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(9) {
  }
}

lemma NonSuspector_ind_7(s:TimestampedRslState, s':TimestampedRslState, j:int, idx:int)
  requires RslAssumption2(s, s')
  requires EpochTimeoutQDInv(s)
  requires EpochTimeoutQDInv(s')
  requires 0 <= idx < |s.t_replicas|
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep == RslStep(7) // on step 7, we might become a suspector

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires NodeIsNotSuspector(s, j);
  ensures  NodeIsNotSuspector(s', j) || (NodeIsSuspector(s', j) && HBUnsent(s', j));
{
  var ios := s.t_environment.nextStep.ios;
  var es := s.t_replicas[j].v.replica.proposer.election_state;
  var clock := SpontaneousClock(UntagLIoOpSeq(ios));
  if clock.t < es.epoch_end_time {
    assert NodeIsNotSuspector(s', j);
  } else {
    assert |es.requests_received_prev_epochs| > 0; // FIXME: put inside of invariant
    assert HBUnsent(s', j);
    assert NodeIsSuspector(s', j);
  }
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


  ensures NodeIsNotSuspector(s, idx) ==> NodeIsNotSuspector(s', idx)
  ensures NodeIsSuspector(s, idx) ==> NodeIsSuspector(s', idx)
  ensures HBUnsent(s, idx) ==> HBUnsent(s', idx)
  ensures HBEnRoute(s, idx) ==> HBEnRoute(s', idx)
  ensures NodeIsKnownSuspector(s, idx) ==> NodeIsKnownSuspector(s', idx)
{
  assert ReplicasDistinct(s.constants.config.replica_ids, j, idx);
  assert ReplicasDistinct(s.constants.config.replica_ids, j, 1);
}

lemma FDInd_noninterf_full(s:TimestampedRslState, s':TimestampedRslState, j:int, idx:int)
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

  requires NodeFDInvariant(s, idx)
  ensures NodeFDInvariant(s', idx)
{
  FDInd_noninterf(s, s', j, idx);
}

lemma Suspector_ind(s:TimestampedRslState, s':TimestampedRslState, j:int, idx:int)
  requires RslAssumption2(s, s')
  requires EpochTimeoutQDInv(s)
  requires EpochTimeoutQDInv(s')
  requires 0 <= idx < |s.t_replicas|
  requires 0 <= j < |s.constants.config.replica_ids|;
  requires j != 1;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires HBEnRoute(s, idx);
  ensures HBEnRoute(s', idx); // || NodeIsKnownSuspector(s, j);
{
  assert ReplicasDistinct(s.constants.config.replica_ids, j, 1);
}

lemma Suspector_ind_leader(s:TimestampedRslState, s':TimestampedRslState, j:int, idx:int)
  requires RslAssumption2(s, s')
  requires EpochTimeoutQDInv(s)
  requires EpochTimeoutQDInv(s')
  requires 0 <= idx < |s.t_replicas|
  requires 0 <= j < |s.constants.config.replica_ids|;
  requires j == 1;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires HBEnRoute(s, idx);
  ensures HBEnRoute(s', idx) || NodeIsKnownSuspector(s, j);
{
  if s.t_environment.nextStep.nodeStep == RslStep(0) {
    var ios := s.t_environment.nextStep.ios;
    assert |ios| > 0;
    if ios[0].LIoOpReceive? && ios[0].r.msg.v.RslMessage_Heartbeat? {
      var p := ios[0].r;
      if p.src in s.constants.config.replica_ids {
        var sender_index := GetReplicaIndex(p.src, s.constants.config);
        // lemma_GetReplicaIndexIsUnique
        if sender_index == idx {
          assume false;
        } else {
          assert HBEnRoute(s', idx);
        }
      } else{
        assert HBEnRoute(s', idx);
      }
    }
  } else{
    assert HBEnRoute(s', idx);
  }
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
