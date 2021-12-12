include "../../DistributedSystem.i.dfy"

include "../../../../Common/Collections/Maps2.i.dfy"
include "../../Constants.i.dfy"
include "../../Environment.i.dfy"
include "../../Replica.i.dfy"
include "../Common/assumptions.i.dfy"

include "../TimestampedRslSystem.i.dfy"
include "FailureHelpers.i.dfy"

include "../../CommonProof/Constants.i.dfy"
// include "../CommonProof/Requests.i.dfy"

module FailureDetection_defns_i {
import opened TimestampedRslSystem_i
import opened FailureHelpers_i
import opened Common_Assumptions

////////////////////////////////////////////////////////////////////////////////
// Assumptions
////////////////////////////////////////////////////////////////////////////////

ghost const req:Request

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


// This should be a self-contained invariant; could remove this assumption.
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
    && (es.requests_received_this_epoch == [req] ||
       es.requests_received_this_epoch == [])
    && (es.requests_received_prev_epochs == [req] ||
       es.requests_received_prev_epochs == [])
    )
}


predicate FOAssumption(s:TimestampedRslState)
{
  && CommonAssumptions(s)
  && NoStateTransfer(s)
  && OneAndOnlyOneRequest(s)
}

predicate FOAssumption2(s:TimestampedRslState, s':TimestampedRslState)
{
  && FOAssumption(s)
  && FOAssumption(s')
  // XXX: shouldn't need to assume this, pretty trivial
  && s.constants.config.replica_ids == s'.constants.config.replica_ids
  && ClockAssumption(s, s')
}

////////////////////////////////////////////////////////////////////////////////
// Main invariants
////////////////////////////////////////////////////////////////////////////////

/*
// "suspecting_replicas"
predicate SuspectingReplicaInv(s:TimestampedRslState, suspecting_replicas:set<int>)
{
  && (forall j :: j in suspecting_replicas ==> 0 <= j < |s.t_replicas|)
  && (forall j :: 0 <= j < |s.t_replicas| ==>
  s.t_replicas[j].v.replica.proposer.election_state.current_view_suspectors <= suspecting_replicas + {j}
    )
    // FIXME: decide what to do about this issue
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
    NotKnownSuspector(s, j) &&
    (NonSuspector0(s,j) || NonSuspector1(s,j) || NonSuspector2(s, j) || InternalSuspector3(s, j))
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

predicate InView2Local(s:TimestampedRslState, j:int, sus:bool)
  requires RslConsistency(s)
{
  if sus then
    Suspector(s, j)
  else
    true
}

predicate InView2(s:TimestampedRslState, suspecting_replicas:set<int>)
  requires RslConsistency(s)
{
  SuspectingReplicaInv(s, suspecting_replicas)
  && |suspecting_replicas| >= LMinQuorumSize(s.constants.config)
  && (
    forall j :: 0 <= j < |s.t_replicas| ==> InView1Local(s, j, j in suspecting_replicas)
  )
  && s.t_replicas[1].v.replica.proposer.election_state.current_view == Ballot(1, 0)
  // TODO: the leader will be in its "final state" as soon as it enters a new ballot
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

predicate NotKnownSuspector(s:TimestampedRslState, j:int)
  requires RslConsistency(s)
  requires 0 <= j < |s.t_replicas|
  requires 0 <= j < |s.constants.config.replica_ids|
{
  // No heartbeats sent indicating that we are one, and leader doesn't think we are.
  HBUnsent(s, j)
  // leader doesn't know about this node being a suspector
  && (s.t_replicas[j].v.replica.constants.my_index !in s.t_replicas[1].v.replica.proposer.election_state.current_view_suspectors)
  // but the leader might be a 
}

// start out in this state initially
// can also enter this after processing request
predicate NonSuspector0(s:TimestampedRslState, j:int)
  requires RslConsistency(s)
  requires 0 <= j < |s.t_replicas|
  requires 0 <= j < |s.constants.config.replica_ids|
{
  var suspectors := s.t_replicas[j].v.replica.proposer.election_state.current_view_suspectors;
  && s.t_replicas[j].v.replica.constants.my_index !in suspectors

  && s.t_replicas[j].v.replica.proposer.election_state.requests_received_this_epoch == []
  && s.t_replicas[j].v.replica.proposer.election_state.requests_received_prev_epochs == []
}

predicate NonSuspector1(s:TimestampedRslState, j:int)
  requires RslConsistency(s)
  requires 0 <= j < |s.t_replicas|
  requires 0 <= j < |s.constants.config.replica_ids|
{

  var suspectors := s.t_replicas[j].v.replica.proposer.election_state.current_view_suspectors;
  && s.t_replicas[j].v.replica.constants.my_index !in suspectors
  && s.t_replicas[j].v.replica.proposer.election_state.requests_received_this_epoch == [req]
  && s.t_replicas[j].v.replica.proposer.election_state.requests_received_prev_epochs == []
  && s.t_replicas[j].v.replica.proposer.election_state.epoch_end_time >= 0
  && TimeLe(s.t_replicas[j].v.replica.proposer.election_state.epoch_end_time, TBEpoch1())
}

predicate NonSuspector2(s:TimestampedRslState, j:int)
  requires RslConsistency(s)
  requires 0 <= j < |s.t_replicas|
{
  var suspectors := s.t_replicas[j].v.replica.proposer.election_state.current_view_suspectors;
  && s.t_replicas[j].v.replica.constants.my_index !in suspectors
  && s.t_replicas[j].v.replica.proposer.election_state.requests_received_prev_epochs == [req]
  && s.t_replicas[j].v.replica.proposer.election_state.epoch_end_time >= 0
  && TimeLe(s.t_replicas[j].v.replica.proposer.election_state.epoch_end_time, TBEpoch2())
}

// Is actually a suspector, but not externally known as a suspector
predicate InternalSuspector3(s:TimestampedRslState, j:int)
  requires RslConsistency(s)
  requires 0 <= j < |s.t_replicas|
{
  var suspectors2 := s.t_replicas[j].v.replica.proposer.election_state.current_view_suspectors;
  && s.t_replicas[j].v.replica.constants.my_index in suspectors2
  // && s.t_replicas[j].v.replica.nextHeartbeatTime >= 0
  // && TimeLe(s.t_replicas[j].v.replica.nextHeartbeatTime, HBPeriodEnd())
}


predicate Suspector(s:TimestampedRslState, j:int)
  requires RslConsistency(s)
  requires 0 <= j < |s.t_replicas|
{
  // Either there's a suspecting HB from that node, or the new leader already knows about it
  (exists pkt ::
  && pkt in s.t_environment.sentPackets
  && pkt.msg.v.RslMessage_Heartbeat?
  && pkt.src == s.constants.config.replica_ids[j]
  && pkt.dst == s.constants.config.replica_ids[1]
  && pkt.msg.v.suspicious == true
  && TimeLe(pkt.msg.ts, TBFirstSuspectingHB())
  )
  || (s.t_replicas[j].v.replica.constants.my_index in s.t_replicas[1].v.replica.proposer.election_state.current_view_suspectors)
}
*/

}
