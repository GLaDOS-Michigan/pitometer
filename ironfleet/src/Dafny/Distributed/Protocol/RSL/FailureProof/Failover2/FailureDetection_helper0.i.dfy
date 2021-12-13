include "../../DistributedSystem.i.dfy"

include "../../../../Common/Collections/Maps2.i.dfy"
include "../../Constants.i.dfy"
include "../../Environment.i.dfy"
include "../../Replica.i.dfy"

include "../TimestampedRslSystem.i.dfy"
include "FailureHelpers.i.dfy"
include "FailureDetection_defns.i.dfy"

include "../../CommonProof/Constants.i.dfy"
// include "../CommonProof/Requests.i.dfy"

module FailureDetection_helper0_i {
import opened TimestampedRslSystem_i
import opened FailureHelpers_i
import opened FailureDetection_defns_i

////////////////////////////////////////////////////////////////////////////////
// Queueing delay invariants + lemmas
////////////////////////////////////////////////////////////////////////////////

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
  requires FOAssumption2(s, s')
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires EpochTimeoutQDInv(s)
  ensures EpochTimeoutQDInv(s')
{
  reveal_EpochQD();
  var r := s.t_replicas[j];
  var r' := s'.t_replicas[j];

  if s.t_environment.nextStep.nodeStep == RslStep(0) {
    assert EpochTimeoutQDInv(s');
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(1) {
    assert EpochTimeoutQDInv(s');
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(2) {
    assert EpochTimeoutQDInv(s');
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(3) {
    assert EpochTimeoutQDInv(s');
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(4) {
    assert EpochTimeoutQDInv(s');
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(5) {
    assert EpochTimeoutQDInv(s');
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(6) {
    assert EpochTimeoutQDInv(s');
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(7) {
    assert EpochTimeoutQDInv(s');
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(8) {
    assert EpochTimeoutQDInv(s');
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(9) {
    assert EpochTimeoutQDInv(s');
  }
  return;


  if s.t_environment.nextStep.nodeStep != RslStep(7) {
    // assert r'.v.replica.proposer.election_state.epoch_end_time == r.v.replica.proposer.election_state.epoch_end_time;
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

predicate EpochDelayInv(s:TimestampedRslState)
{
  && (forall idx :: && 0 <= idx < |s.t_replicas| ==>
    s.t_replicas[idx].v.replica.proposer.election_state.epoch_end_time >= 0 // so it's a valid Timestamp
    && TimeLe(s.t_replicas[idx].v.replica.proposer.election_state.epoch_end_time,
     s.t_replicas[idx].ts + EpochLength)
    )
}

// Hearbeat delay invariant; self-contained
predicate HeartbeatDelayInv(s:TimestampedRslState)
{
  && (forall idx :: && 0 <= idx < |s.t_replicas| ==>
    s.t_replicas[idx].v.replica.nextHeartbeatTime >= 0 // so it's a valid Timestamp
    && TimeLe(s.t_replicas[idx].v.replica.nextHeartbeatTime, s.t_replicas[idx].ts + HBPeriod)
    )
}

// Heartbeat queueing delay invariant; self-contained
predicate HeartbeatQDInv(s:TimestampedRslState)
{
  && (forall idx :: && 0 <= idx < |s.t_replicas| ==>
    s.t_replicas[idx].v.replica.proposer.election_state.epoch_end_time >= 0 // so it's a valid Timestamp
    && TimeLe(s.t_replicas[idx].ts, s.t_replicas[idx].v.replica.proposer.election_state.epoch_end_time + EpochQD(s.t_replicas[idx].v.nextActionIndex))
    )

  && (forall idx :: && 0 <= idx < |s.t_replicas| ==>
    s.t_replicas[idx].v.replica.nextHeartbeatTime >= 0 // so it's a valid Timestamp
    && TimeLe(s.t_replicas[idx].ts, s.t_replicas[idx].v.replica.nextHeartbeatTime + HeartbeatQD(s.t_replicas[idx].v.nextActionIndex))
    )
}

predicate DelayInvs(s:TimestampedRslState)
{
  && EpochTimeoutQDInv(s)
  && EpochDelayInv(s)
  && HeartbeatDelayInv(s)
  && HeartbeatQDInv(s)
}

// TODO: this should be used from CommonProof/Requests.i.dfy
lemma lemma_RemoveExecutedRequestBatchProducesSubsequence(s':seq<Request>, s:seq<Request>, batch:RequestBatch)
  requires s' == RemoveExecutedRequestBatch(s, batch);
  ensures  forall x :: x in s' ==> x in s;
  ensures |s'| <= |s|;
  decreases |batch|;
{
}

}
