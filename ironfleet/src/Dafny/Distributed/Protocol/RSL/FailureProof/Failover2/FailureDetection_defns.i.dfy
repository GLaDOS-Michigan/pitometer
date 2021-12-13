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
  // FIXME: rewrite without quantifiers?
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

predicate NonLeadersView1(s:TimestampedRslState)
{
  forall j :: 0 <= j < |s.t_replicas| ==>
    j != 1 ==>
    s.t_replicas[j].v.replica.proposer.election_state.current_view == Ballot(1,0)
}

predicate EpochLengthAssumption(s:TimestampedRslState)
{
  s.constants.params.baseline_view_timeout_period == EpochLength

  && forall j :: 0 <= j < |s.t_replicas| ==>
    s.t_replicas[j].v.replica.proposer.election_state.current_view == Ballot(1,0) ==>
    s.t_replicas[j].v.replica.proposer.election_state.epoch_length == EpochLength

}

predicate HBPeriodAssumption(s:TimestampedRslState)
{
  s.constants.params.heartbeat_period == HBPeriod
}

predicate ViewNoOverflow(s:TimestampedRslState)
{
  forall j :: 0 <= j < |s.t_replicas| ==>
  var es := s.t_replicas[j].v.replica.proposer.election_state;
  LtUpperBound(es.current_view.seqno, es.constants.all.params.max_integer_val)
}

predicate EpochAndHeartbeatNoOverflow(s:TimestampedRslState)
{
  s.t_environment.nextStep.LEnvStepHostIos?
  ==>
  var ios := s.t_environment.nextStep.ios;
  forall io :: io in s.t_environment.nextStep.ios && io.LIoOpReadClock? ==>
    io.t >= 0
    && UpperBoundedAddition(io.t, EpochLength, s.constants.params.max_integer_val) == io.t + EpochLength
    && UpperBoundedAddition(io.t, HBPeriod, s.constants.params.max_integer_val) == io.t + HBPeriod
}

predicate RequestTimeAssumption(s:TimestampedRslState)
{
  && (forall pkt :: pkt in s.t_environment.sentPackets ==>
    pkt.msg.v.RslMessage_Request? ==>
    TimeLe(pkt.msg.ts, RequestTime))
}

// XXX: this could be part of the model; don't want to introduce it there to
// avoid breaking things for now
predicate MonotoneTime(s:TimestampedRslState, s':TimestampedRslState)
{
  forall j :: 0 <= j < |s.t_replicas| ==> TimeLe(s.t_replicas[j].ts, s'.t_replicas[j].ts)
}

predicate FOAssumption(s:TimestampedRslState)
{
  && CommonAssumptions(s)
  && NoStateTransfer(s)
  && OneAndOnlyOneRequest(s)
  && NonLeadersView1(s)
  && EpochLengthAssumption(s)
  && HBPeriodAssumption(s)
  && ViewNoOverflow(s)
  && RequestTimeAssumption(s)
  && EpochAndHeartbeatNoOverflow(s)
}

predicate FOAssumption2(s:TimestampedRslState, s':TimestampedRslState)
{
  && FOAssumption(s)
  && FOAssumption(s')
  // XXX: shouldn't need to assume this, pretty trivial
  && s.constants.config.replica_ids == s'.constants.config.replica_ids
  && ClockAssumption(s, s')
  && MonotoneTime(s, s')
}

////////////////////////////////////////////////////////////////////////////////
// Main invariants
////////////////////////////////////////////////////////////////////////////////

predicate LeaderView0(s:TimestampedRslState)
  requires CommonAssumptions(s)
{
  s.t_replicas[1].v.replica.proposer.election_state.current_view == Ballot(1,0)
}

// "suspecting_replicas"
predicate SuspectingReplicaInv(s:TimestampedRslState, suspecting_replicas:set<int>)
  requires CommonAssumptions(s)
{
  && (forall j :: j in suspecting_replicas ==> 0 <= j < |s.t_replicas|)
  && s.t_replicas[1].v.replica.proposer.election_state.current_view_suspectors <= suspecting_replicas
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

predicate InView1Local(s:TimestampedRslState, j:int, sus:bool)
  requires RslConsistency(s)
  requires CommonAssumptions(s)
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
  requires CommonAssumptions(s)
{
  SuspectingReplicaInv(s, suspecting_replicas)

  && LeaderQuorumBound(s)
  && s.t_replicas[1].v.replica.proposer.election_state.current_view == Ballot(1,0)

  && InView1Packets(s)
  && CurrView(s)
  && (
    forall j :: 0 <= j < |s.t_replicas| ==> InView1Local(s, j, j in suspecting_replicas)
  )
}

predicate FinalStage(s:TimestampedRslState)
  requires CommonAssumptions(s)
{
  && s.t_replicas[1].v.replica.proposer.election_state.current_view == Ballot(1, 1)
  && TimeLe(s.t_replicas[1].ts, TBNewView())
  // FIXME: this should maintain that leader.current_state != 2, up until it is.
  // That should help prove a bound on the 1a packets that are sent out
  // Also should add requirement that no reply packet is sent to the client.
  // Probably gonna have to assume no executions for that.
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
  requires CommonAssumptions(s)
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
  && s.t_replicas[j].v.replica.nextHeartbeatTime >= 0
  && TimeLe(s.t_replicas[j].v.replica.nextHeartbeatTime, HBPeriodEnd())
}


predicate Suspector(s:TimestampedRslState, j:int)
  requires CommonAssumptions(s);
  requires RslConsistency(s);
  requires 0 <= j < |s.t_replicas|
{
  // If the leader is a suspector, it already knows about itself
  && (j == 1 ==> s.t_replicas[j].v.replica.constants.my_index in s.t_replicas[1].v.replica.proposer.election_state.current_view_suspectors)
  && (
  // Either there's a suspecting HB from that node, or the new leader already knows about it
  (exists pkt ::
  && pkt in s.undeliveredPackets
  && pkt.msg.v.RslMessage_Heartbeat?
  && pkt.src == s.constants.config.replica_ids[j]
  && pkt.dst == s.constants.config.replica_ids[1]
  && pkt.msg.v.suspicious == true
  && TimeLe(pkt.msg.ts, TBFirstSuspectingHB())
  )
  || (s.t_replicas[j].v.replica.constants.my_index in s.t_replicas[1].v.replica.proposer.election_state.current_view_suspectors)
  )
}

predicate LeaderQuorumBound(s:TimestampedRslState)
  requires CommonAssumptions(s)
{
  && (
    |s.t_replicas[1].v.replica.proposer.election_state.current_view_suspectors| >= LMinQuorumSize(s.constants.config) ==>
    && 1 <= s.t_replicas[1].v.nextActionIndex <= 8
    && TimeLe(s.t_replicas[1].ts, TBJustBeforeNewView(s.t_replicas[1].v.nextActionIndex))
  )
}

}
