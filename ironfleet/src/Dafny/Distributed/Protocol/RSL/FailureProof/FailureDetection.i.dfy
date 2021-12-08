include "../DistributedSystem.i.dfy"

include "../../../Common/Collections/Maps2.i.dfy"
include "../Constants.i.dfy"
include "../Environment.i.dfy"
include "../Replica.i.dfy"

include "TimestampedRslSystem.i.dfy"
include "FailureHelpers.i.dfy"

include "../CommonProof/Constants.i.dfy"
// include "../CommonProof/Requests.i.dfy"
include "PureHelpers.i.dfy"

module FailureDetection_i {
import opened TimestampedRslSystem_i
import opened FailureHelpers_i
import opened PureHelpers_i
// import opened CommonProof__Requests_i

predicate RslConsistency(s:TimestampedRslState)
{
  ConstantsAllConsistentInv(UntimestampRslState(s))
    && WellFormedLConfiguration(s.constants.config)
    && 1 < |s.t_replicas|
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
  requires RslConsistency(s')
  requires s.constants.config.replica_ids == s'.constants.config.replica_ids
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

predicate NoStateTransfer(s:TimestampedRslState)
{
  s.t_environment.nextStep.LEnvStepHostIos? ==>
    (forall io :: io in s.t_environment.nextStep.ios && io.LIoOpReceive? ==>
    && !io.r.msg.v.RslMessage_AppStateRequest?
    && !io.r.msg.v.RslMessage_AppStateSupply?
    && !io.r.msg.v.RslMessage_StartingPhase2?
    )
}

predicate NoExternalSteps(s:TimestampedRslState)
{
  s.t_environment.nextStep.LEnvStepHostIos? ==>
  s.t_environment.nextStep.actor in s.constants.config.replica_ids
}

ghost const req:Request

predicate RslAssumption(s:TimestampedRslState)
{
  && RslConsistency(s)
  && BoundedQueueingAssumption(s)
  && NoStateTransfer(s)
  && NoExternalSteps(s)
  && OneAndOnlyOneRequest(s)
}

predicate RslAssumption2(s:TimestampedRslState, s':TimestampedRslState)
{
  && RslAssumption(s)
  && RslAssumption(s')
  // FIXME: shouldn't need to assume this, pretty trivial
  && s.constants.config.replica_ids == s'.constants.config.replica_ids
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

// This should be a self-contained invariant; maybe even an assumption
predicate OneAndOnlyOneRequest(s:TimestampedRslState)
  requires RslConsistency(s)
{
  && (forall pkt :: pkt in s.t_environment.sentPackets ==>
    pkt.msg.v.RslMessage_Request? ==>
    && pkt.msg.v.seqno_req == req.seqno
    && pkt.msg.v.val == req.request
    && pkt.src == req.client)
  && (forall j :: 0 <= j < |s.t_replicas| ==>
    var es := s.t_replicas[j].v.replica.proposer.election_state;
    (forall req' :: req' in es.requests_received_prev_epochs ==> req' == req)
    && (forall req' :: req' in es.requests_received_this_epoch ==> req' == req)
    )
}

// "suspecting_replicas"
predicate SuspectingReplicaInv(s:TimestampedRslState, suspecting_replicas:set<int>)
{
  && (forall j :: j in suspecting_replicas ==> 0 <= j < |s.t_replicas|)
  && (forall j :: 0 <= j < |s.t_replicas| ==>
        s.t_replicas[j].v.replica.proposer.election_state.current_view_suspectors <= suspecting_replicas
    )
}

// no one is in view 2
// This invariant becomes untrue as soon as |sr| >= MinQuorumSize

predicate InView1Packets(s:TimestampedRslState)
  requires RslConsistency(s)
{
  && (forall pkt ::
     pkt in s.t_environment.sentPackets ==>

     && (pkt.msg.v.RslMessage_Heartbeat? ==>
     pkt.msg.v.bal_heartbeat == Ballot(1, 0))
     && (pkt.msg.v.RslMessage_2a? ==>
     pkt.msg.v.bal_2a == Ballot(1, 0))
     && (pkt.msg.v.RslMessage_2b? ==>
     pkt.msg.v.bal_2b == Ballot(1, 0))
     && (pkt.msg.v.RslMessage_1a? ==>
     pkt.msg.v.bal_1a == Ballot(1, 0))
     && (pkt.msg.v.RslMessage_1b? ==>
     pkt.msg.v.bal_1b == Ballot(1, 0))
  )
}

// TODO: probably want a InV2Local without the else case
predicate InView1Local(s:TimestampedRslState, j:int, sus:bool)
  requires RslConsistency(s)
  requires 0 <= j < |s.t_replicas|
{
  if sus then
    Suspector(s, j) // Steps of j and of the leader affect this
  else
    // HB unsent and no one thinks j is a suspector
    NonSuspector(s, j) &&
    (NonSuspector0(s,j) || NonSuspector1(s,j) || NonSuspector2(s, j))
}

predicate CurrView(s:TimestampedRslState)
{
  && (
    forall j :: 0 <= j < |s.t_replicas| ==>
    s.t_replicas[j].v.replica.proposer.election_state.current_view == Ballot(1,0)
    )
}


predicate InView1(s:TimestampedRslState, suspecting_replicas:set<int>)
  requires RslConsistency(s)
{
  SuspectingReplicaInv(s, suspecting_replicas)
  && |suspecting_replicas| < LMinQuorumSize(s.constants.config)

  && InView1Packets(s)
  && CurrView(s)
  && (
    forall j :: 0 <= j < |s.t_replicas| ==> InView1Local(s, j, j in suspecting_replicas)
  )
}

predicate InView2(s:TimestampedRslState, suspecting_replicas:set<int>)
  requires RslConsistency(s)
{
  SuspectingReplicaInv(s, suspecting_replicas)
  && |suspecting_replicas| >= LMinQuorumSize(s.constants.config)
  // TODO: this should also have a majority with Suspector(j) true
}

predicate NonSuspector(s:TimestampedRslState, j:int)
  requires RslConsistency(s)
  requires 0 <= j < |s.t_replicas|
  requires 0 <= j < |s.constants.config.replica_ids|
{
  // Not a suspector ourselves, and no heartbeats sent indicating that we are one.
  var suspectors := s.t_replicas[j].v.replica.proposer.election_state.current_view_suspectors;
  s.t_replicas[j].v.replica.constants.my_index !in suspectors &&
  HBUnsent(s, j) &&
  // leader doesn't suspect this node
  s.t_replicas[j].v.replica.constants.my_index !in s.t_replicas[1].v.replica.proposer.election_state.current_view_suspectors
}

// start out in this state initially
// can also enter this after processing request
predicate NonSuspector0(s:TimestampedRslState, j:int)
  requires RslConsistency(s)
  requires 0 <= j < |s.t_replicas|
  requires 0 <= j < |s.constants.config.replica_ids|
{
  s.t_replicas[j].v.replica.proposer.election_state.requests_received_this_epoch == []
  && s.t_replicas[j].v.replica.proposer.election_state.requests_received_prev_epochs == []
}

predicate NonSuspector1(s:TimestampedRslState, j:int)
  requires RslConsistency(s)
  requires 0 <= j < |s.t_replicas|
  requires 0 <= j < |s.constants.config.replica_ids|
{
  && s.t_replicas[j].v.replica.proposer.election_state.requests_received_this_epoch == [req]
  && s.t_replicas[j].v.replica.proposer.election_state.requests_received_prev_epochs == []
  && s.t_replicas[j].v.replica.proposer.election_state.epoch_end_time >= 0
  && TimeLe(s.t_replicas[j].v.replica.proposer.election_state.epoch_end_time, TBEpoch1())
}

predicate NonSuspector2(s:TimestampedRslState, j:int)
  requires RslConsistency(s)
  requires 0 <= j < |s.t_replicas|
  requires 0 <= j < |s.constants.config.replica_ids|
{
  && s.t_replicas[j].v.replica.proposer.election_state.requests_received_prev_epochs == [req]
  && s.t_replicas[j].v.replica.proposer.election_state.epoch_end_time >= 0
  && TimeLe(s.t_replicas[j].v.replica.proposer.election_state.epoch_end_time, TBEpoch2())
}

predicate HBUnsent(s:TimestampedRslState, j:int)
  requires RslConsistency(s)
  requires 0 <= j < |s.constants.config.replica_ids|
{
  // All HBs are unsuspecting
  forall pkt ::
  pkt in s.t_environment.sentPackets ==>
  pkt.msg.v.RslMessage_Heartbeat? ==>
  pkt.src == s.constants.config.replica_ids[j] ==>
  pkt.msg.v.suspicious == false
}

predicate Suspector(s:TimestampedRslState, j:int)
  requires RslConsistency(s)
  requires 0 <= j < |s.constants.config.replica_ids|
{
  // Either there's a suspecting HB from that node, or the new leader already knows about it
  (exists pkt ::
  && pkt in s.t_environment.sentPackets
  && pkt.msg.v.RslMessage_Heartbeat?
  && pkt.src == s.constants.config.replica_ids[j]
  && pkt.dst == s.constants.config.replica_ids[1]
  && pkt.msg.v.suspicious == true
  && TimeLe(pkt.msg.ts, TBFirstSuspectingHB())
  ) ||
  true // FIXME: the leader knows us to be a suspector
}

////////////////////////////////////////////////////////////////////////////////
// PF_NONSUSP
////////////////////////////////////////////////////////////////////////////////

lemma NonSuspector1_ind_most(s:TimestampedRslState, s':TimestampedRslState, sr:set<int>, j:int)
  requires RslAssumption2(s, s')
  // requires EpochTimeoutQDInv(s)
  // requires EpochTimeoutQDInv(s')
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  // requires s.t_environment.nextStep.nodeStep != RslStep(0) // on step 0, we might just enter a new view because of HB
  requires s.t_environment.nextStep.nodeStep != RslStep(7) // on step 7, we might start new epoch
  requires s.t_environment.nextStep.nodeStep != RslStep(6) // on step 6, we might go to NS0(j)

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires InView1(s, sr);

  requires NonSuspector1(s, j);
  ensures  NonSuspector1(s', j);
  // ensures InView1(s', sr, req);
{
  if s.t_environment.nextStep.nodeStep == RslStep(0) {
    var ios := s.t_environment.nextStep.ios;

    if ios[0].LIoOpReceive?  {
      if ios[0].r.msg.v.RslMessage_Heartbeat? {
        assert NonSuspector1(s', j);
      } else if ios[0].r.msg.v.RslMessage_1a? {
        assert NonSuspector1(s', j);
      } else if ios[0].r.msg.v.RslMessage_1b? {
        assert NonSuspector1(s', j);
      } else if ios[0].r.msg.v.RslMessage_2b? {
        assert NonSuspector1(s', j);
      } else if ios[0].r.msg.v.RslMessage_2a? {
        assert NonSuspector1(s', j);
      } else if ios[0].r.msg.v.RslMessage_Request? {
        var es := s.t_replicas[j].v.replica.proposer.election_state;
        var newReq := Request(ios[0].r.src, ios[0].r.msg.v.seqno_req, ios[0].r.msg.v.val);
        assert ios[0].r in s.t_environment.sentPackets;
        assert ios[0].r.src == req.client;
        assert ios[0].r.msg.v.seqno_req == req.seqno;
        assert (req in es.requests_received_this_epoch) && RequestsMatch(req, newReq);

        assert s.t_replicas[j].v.replica.proposer.election_state.requests_received_this_epoch ==
          s'.t_replicas[j].v.replica.proposer.election_state.requests_received_this_epoch;
        assert NonSuspector1(s', j);
      } else if ios[0].r.msg.v.RslMessage_Reply? {
        assert NonSuspector1(s', j);
      } else if ios[0].r.msg.v.RslMessage_AppStateRequest? {
        assert NonSuspector1(s', j);
      } else if ios[0].r.msg.v.RslMessage_AppStateSupply? {
        assert NonSuspector1(s', j);
      } else if ios[0].r.msg.v.RslMessage_StartingPhase2? {
        assert NonSuspector1(s', j);
      } else{
        assert NonSuspector1(s', j);
      }
    } else {
        assert NonSuspector1(s', j);
    }
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(1) {
    assert NonSuspector1(s', j);
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(2) {
    assert NonSuspector1(s', j);
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(3) {
    assert NonSuspector1(s', j);
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(4) {
    assert NonSuspector1(s', j);
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(5) {
    assert NonSuspector1(s', j);
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(6) {
    assert false;
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(7) {
    assert false;
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(8) {
    var r := s.t_replicas[j].v.replica;
    // assert r.proposer.election_state.current_view_suspectors <= sr;
    SubsetCardinality(r.proposer.election_state.current_view_suspectors, sr);
    // assert |r.proposer.election_state.current_view_suspectors| <= |sr|;
    // assert |sr| < LMinQuorumSize(s.constants.config);
    // assert |r.proposer.election_state.current_view_suspectors| < LMinQuorumSize(s.constants.config);
    // assert NonSuspector1(s', j, req);
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(9) {
    assert NonSuspector1(s', j);
  }
}

// TODO: this should be used from CommonProof/Requests.i.dfy
lemma lemma_RemoveExecutedRequestBatchProducesSubsequence(s':seq<Request>, s:seq<Request>, batch:RequestBatch)
  requires s' == RemoveExecutedRequestBatch(s, batch);
  ensures  forall x :: x in s' ==> x in s;
  decreases |batch|;
{
}

lemma NonSuspector1_ind_6(s:TimestampedRslState, s':TimestampedRslState, j:int)
  requires RslAssumption2(s, s')
  // requires EpochTimeoutQDInv(s)
  // requires EpochTimeoutQDInv(s')
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep == RslStep(6) // on step 6, we might go to NS0(j)

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires NonSuspector1(s, j);
  ensures  NonSuspector1(s', j) || NonSuspector0(s', j);
{
  if s.t_replicas[j].v.replica.executor.next_op_to_execute.OutstandingOpKnown? {
    var es := s.t_replicas[j].v.replica.proposer.election_state;
    var es' := s'.t_replicas[j].v.replica.proposer.election_state;
    var batch := s.t_replicas[j].v.replica.executor.next_op_to_execute.v;
    if ElectionStateReflectExecutedRequestBatch(es, es', batch) {
      // FIXME: things can only get removed; this should be easy to prove
      // This works if we assume that reqs' <= reqs, which is true because the
      // sequence has at most one elt in it
      lemma_RemoveExecutedRequestBatchProducesSubsequence(
      es'.requests_received_this_epoch,
      es.requests_received_this_epoch,
      batch);
      lemma_RemoveExecutedRequestBatchProducesSubsequence(
      es'.requests_received_prev_epochs,
      es.requests_received_prev_epochs,
      batch);

      assume es'.requests_received_prev_epochs == [];
      assume es'.requests_received_this_epoch <= es.requests_received_this_epoch;

      if s'.t_replicas[j].v.replica.proposer.election_state.requests_received_this_epoch == [] {
        assert NonSuspector0(s', j);
      } else{
        assert NonSuspector1(s', j);
      }
    } else {
      assert NonSuspector1(s', j);
    }
  } else {
    assert NonSuspector1(s', j);
  }
}

lemma NonSuspector1_ind_7(s:TimestampedRslState, s':TimestampedRslState, j:int)
  requires RslAssumption2(s, s')
  requires EpochTimeoutQDInv(s)
  requires EpochTimeoutQDInv(s')
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep == RslStep(7)

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires NonSuspector1(s, j);
  ensures  NonSuspector1(s', j);
{
  assert false; // TODO
}

lemma InView1Local_self_ind(s:TimestampedRslState, s':TimestampedRslState, sr:set<int>, j:int) returns (sr':set<int>)
  requires RslAssumption2(s, s')
  requires EpochTimeoutQDInv(s)
  requires EpochTimeoutQDInv(s')
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires InView1(s, sr);

  ensures SuspectingReplicaInv(s', sr');
  ensures InView1Local(s', j, j in sr');
  ensures sr' == sr  || sr' == sr + {j};
{
  sr' := sr;
  var sus := j in sr;
  if sus {
    assert false;
    assert Suspector(s', j);
    return;
  }
  // else, non-sus
  if NonSuspector0(s, j) {
    // TODO: write lemmas for this?
    assert false;
  } else if NonSuspector1(s, j) {
    var step := s.t_environment.nextStep.nodeStep;
    if step == RslStep(7) {
      NonSuspector1_ind_7(s, s', j);
    } else if step == RslStep(6) {
      NonSuspector1_ind_6(s, s', j);
    } else {
      NonSuspector1_ind_most(s, s', sr, j);
    }
  } else {
    assert false;
    assert NonSuspector2(s, j);
    // NOTE: j has a change of becoming a suspector in this case
  }
}

lemma InView1Local_leader_ind(s:TimestampedRslState, s':TimestampedRslState, sr:set<int>, k:int)
  requires RslAssumption2(s, s')
  requires EpochTimeoutQDInv(s)
  requires EpochTimeoutQDInv(s')
  requires 0 <= k < |s.constants.config.replica_ids|;
  requires k != 1

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[1];

  requires TimestampedRslNextOneReplica(s, s', 1, s.t_environment.nextStep.ios);

  requires InView1(s, sr);
  ensures InView1Local(s', k, k in sr)
{
  assert false;
}

lemma InView1Local_all_ind(s:TimestampedRslState, s':TimestampedRslState, sr:set<int>, j:int) returns (sr':set<int>)
  requires RslAssumption2(s, s')
  requires EpochTimeoutQDInv(s)
  requires EpochTimeoutQDInv(s')
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires InView1(s, sr);

  ensures  SuspectingReplicaInv(s', sr')
  ensures  forall k :: 0 <= k < |s'.t_replicas| ==> InView1Local(s', k, k in sr');
{
  sr' := InView1Local_self_ind(s, s', sr, j);
  // Also use InV1L_leader_ind
  if j == 1 {
    forall k | 0 <= k < |s'.t_replicas|
      ensures InView1Local(s', k, k in sr')
    {
      if k == j { // lemma above finishes this
      } else {
        InView1Local_leader_ind(s, s', sr, k);
      }
    }
    return;
  }

  forall k | 0 <= k < |s'.t_replicas|
    ensures InView1Local(s', k, k in sr')
  {
    if k == j {
      // lemma above finishes this.
    } else { // trivial; the state of node k is unchanged, and j is not a
      // leader. So, inv for k should be maintained
      assert ReplicasDistinct(s.constants.config.replica_ids, j, k);
    }
  }
}

lemma InView1_to_Packets(s:TimestampedRslState, s':TimestampedRslState, j:int, sr:set<int>)
  requires RslAssumption2(s, s')

  requires 0 <= j < |s.constants.config.replica_ids|;
  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires InView1(s, sr);
  ensures  InView1Packets(s');
{ // trivial
}

lemma InView1_to_CurrView(s:TimestampedRslState, s':TimestampedRslState, j:int, sr:set<int>)
  requires RslAssumption2(s, s')

  requires 0 <= j < |s.constants.config.replica_ids|;
  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires InView1(s, sr);
  ensures  CurrView(s')
{
  if s.t_environment.nextStep.nodeStep == RslStep(8) {
    // have to use subset cardinality argument
    SubsetCardinality(s.t_replicas[j].v.replica.proposer.election_state.current_view_suspectors, sr);
    assert CurrView(s');
  }
}

lemma InView1_ind_hostStep(s:TimestampedRslState, s':TimestampedRslState, j:int, sr:set<int>) returns (sr':set<int>)
  requires RslAssumption2(s, s')
  requires EpochTimeoutQDInv(s)
  requires EpochTimeoutQDInv(s')

  requires 0 <= j < |s.constants.config.replica_ids|;
  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires InView1(s, sr);
  ensures  InView1(s', sr') || InView2(s', sr');
{
  sr' := InView1Local_all_ind(s, s', sr, j);

  if |sr'| < LMinQuorumSize(s.constants.config) {
    InView1_to_Packets(s, s', j, sr);
    InView1_to_CurrView(s, s', j, sr);
    assert InView1(s', sr');
  } else {
    assert InView2(s', sr'); // TODO: this will change
  }
}

lemma InView1_ind(s:TimestampedRslState, s':TimestampedRslState, sr:set<int>) returns (sr':set<int>)
  requires RslAssumption2(s, s')
  requires EpochTimeoutQDInv(s)
  requires EpochTimeoutQDInv(s')
  requires TimestampedRslNext(s, s');

  requires InView1(s, sr);
  ensures  InView1(s', sr') || InView2(s', sr');
{
  sr' := sr;
  // three cases:
  if TimestampedRslNextEnvironment(s, s') {
    assert InView1(s', sr');
  } else if (exists j, ios :: TimestampedRslNextOneReplica(s, s', j, ios)) {
    // XXX: this is where the heavy lifting happens
    var j, ios :| TimestampedRslNextOneReplica(s, s', j, ios);
    sr' := InView1_ind_hostStep(s, s', j, sr);
  } else {
    var idx, ios :| TimestampedRslNextOneExternal(s, s', idx, ios);
    assert false; // Because we assume no external steps
  }
}

/*
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
*/
}
