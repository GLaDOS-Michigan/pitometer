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
include "CommonProof_Requests.i.dfy"

module FailureDetection_helper0_i {
import opened TimestampedRslSystem_i
import opened FailureHelpers_i
import opened FailureDetection_defns_i
import opened CommonProof__Requests_i

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

  requires exists sr :: InView1(s, sr) // XXX: Need this to know about the epoch length from FOAssumption
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
    var ios := s.t_environment.nextStep.ios;
    var clock := SpontaneousClock(UntagLIoOpSeq(ios));
    if clock.t < r.v.replica.proposer.election_state.epoch_end_time {
      assert EpochTimeoutQDInv(s');
    } else {
      // r'.ts <= r.ts + BLAH
      // NOTE: had to assume that the UpperBoundedAddition for new epoch doesn't fail
      assert r'.v.replica.proposer.election_state.epoch_end_time >= clock.t;
      assert clock.t >= r.ts;
      assert r'.v.replica.proposer.election_state.epoch_end_time >= r.ts;
      assert r'.ts <= r'.v.replica.proposer.election_state.epoch_end_time + StepToTimeDelta(RslStep(7));
      assert EpochTimeoutQDInv(s');
    }
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(8) {
    assert EpochTimeoutQDInv(s');
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(9) {
    assert EpochTimeoutQDInv(s');
  }
  return;
}

predicate EpochDelayInv(s:TimestampedRslState)
{
  && (forall idx :: && 0 <= idx < |s.t_replicas| ==>
    s.t_replicas[idx].v.replica.proposer.election_state.current_view == Ballot(1,0) ==>
    s.t_replicas[idx].v.replica.proposer.election_state.epoch_end_time >= 0 // so it's a valid Timestamp
    && TimeLe(s.t_replicas[idx].v.replica.proposer.election_state.epoch_end_time,
     s.t_replicas[idx].ts + EpochLength)
    )
}

lemma EpochDelayInductive(s:TimestampedRslState, s':TimestampedRslState, j:int)
  requires FOAssumption2(s, s')
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  // requires exists sr :: InView1(s, sr) // XXX: Need this to know about the epoch length from FOAssumption
  requires EpochDelayInv(s)
  ensures  EpochDelayInv(s')
{
  if s.t_environment.nextStep.nodeStep == RslStep(0) {
    assert EpochDelayInv(s');
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(1) {
    assert EpochDelayInv(s');
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(2) {
    assert s.t_replicas[j].v.replica.proposer.election_state.epoch_end_time ==
    s'.t_replicas[j].v.replica.proposer.election_state.epoch_end_time;
    assert EpochDelayInv(s');
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(3) {
    assert EpochDelayInv(s');
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(4) {
    assert EpochDelayInv(s');
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(5) {
    assert EpochDelayInv(s');
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(6) {
    assert EpochDelayInv(s');
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(7) {
    assert EpochDelayInv(s');
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(8) {
    assert EpochDelayInv(s');
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(9) {
    assert EpochDelayInv(s');
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

lemma HeartbeatDelayInductive(s:TimestampedRslState, s':TimestampedRslState, j:int)
  requires FOAssumption2(s, s')
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires HeartbeatDelayInv(s)
  ensures  HeartbeatDelayInv(s')
{
}

// Heartbeat queueing delay invariant; self-contained
predicate HeartbeatQDInv(s:TimestampedRslState)
{
  && (forall idx :: && 0 <= idx < |s.t_replicas| ==>
    s.t_replicas[idx].v.replica.nextHeartbeatTime >= 0 // so it's a valid Timestamp
    && TimeLe(s.t_replicas[idx].ts, s.t_replicas[idx].v.replica.nextHeartbeatTime + HeartbeatQD(s.t_replicas[idx].v.nextActionIndex))
    )
}

lemma HeartbeatQDInductive(s:TimestampedRslState, s':TimestampedRslState, j:int)
  requires FOAssumption2(s, s')
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires HeartbeatQDInv(s)
  ensures  HeartbeatQDInv(s')
{
  reveal_HeartbeatQD();
  if s.t_environment.nextStep.nodeStep == RslStep(0) {
    assert HeartbeatQDInv(s');
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(1) {
    assert HeartbeatQDInv(s');
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(2) {
    assert HeartbeatQDInv(s');
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(3) {
    assert HeartbeatQDInv(s');
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(4) {
    assert HeartbeatQDInv(s');
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(5) {
    assert HeartbeatQDInv(s');
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(6) {
    assert HeartbeatQDInv(s');
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(7) {
    assert HeartbeatQDInv(s');
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(8) {
    assert HeartbeatQDInv(s');
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(9) {
    assert HeartbeatQDInv(s');
  }
}

lemma DelayInv_ind(s:TimestampedRslState, s':TimestampedRslState)
  requires FOAssumption2(s, s')
  requires DelayInvs(s)
  requires TimestampedRslNext(s, s');

  requires (exists sr :: InView1(s, sr));
  ensures  DelayInvs(s')
{
  if TimestampedRslNextEnvironment(s, s') {
    assert DelayInvs(s'); // trivial
  } else if (exists j, ios :: TimestampedRslNextOneReplica(s, s', j, ios)) {
    // XXX: this is where the heavy lifting happens
    var j, ios :| TimestampedRslNextOneReplica(s, s', j, ios);
    HeartbeatQDInductive(s, s', j);
    HeartbeatDelayInductive(s, s', j);
    EpochTimeoutQDInductive(s, s', j);
    EpochDelayInductive(s, s', j);
  } else {
    var idx, ios :| TimestampedRslNextOneExternal(s, s', idx, ios);
    assert false; // Because we assume no external steps
  }
}

predicate DelayInvs(s:TimestampedRslState)
{
  && EpochTimeoutQDInv(s)
  && EpochDelayInv(s)
  && HeartbeatDelayInv(s)
  && HeartbeatQDInv(s)
}

}
