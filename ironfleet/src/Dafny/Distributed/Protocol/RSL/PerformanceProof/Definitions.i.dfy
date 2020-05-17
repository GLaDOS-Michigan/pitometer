include "../DistributedSystem.i.dfy"

include "../../../Common/Collections/Maps2.i.dfy"
include "../Constants.i.dfy"
include "../Environment.i.dfy"
include "../Replica.i.dfy"

include "TimestampedRslSystem.i.dfy"

module Definitions_i {
import opened RslTimestampedDistributedSystem_i

// Generic action step time
ghost const {:verify false} L:Timestamp

ghost const {:verify false} D:Timestamp
ghost const {:verify false} SelfDelivery:Timestamp
ghost const {:verify false} ProcessPacket:Timestamp
ghost const {:verify false} MbeNewView:Timestamp
ghost const {:verify false} MbeP2:Timestamp
ghost const {:verify false} MbeP2a:Timestamp
ghost const {:verify false} TruncLog:Timestamp
ghost const {:verify false} MbeDecide:Timestamp
ghost const {:verify false} MbeExec:Timestamp
ghost const {:verify false} CheckViewTimeout:Timestamp
ghost const {:verify false} CheckQuorumOfViewSuspicion:Timestamp
ghost const {:verify false} MbeSendHeartbeat:Timestamp

 //ghost const Gs := StepTime(RslStep);
 //ghost const As := StepTime(AcceptStep);
 //ghost const G:Timestamp
 //ghost const A:Timestamp

function {:verify false} StepToTimeDelta(hstep:RslStep) : Timestamp
{
  if hstep == RslStep(0) then
    ProcessPacket
  else if hstep == RslStep(1) then
    MbeNewView
  else if hstep == RslStep(2) then
    MbeP2
  else if hstep == RslStep(3) then
    MbeP2a
  else if hstep == RslStep(4) then
    TruncLog
  else if hstep == RslStep(5) then
    MbeDecide
  else if hstep == RslStep(6) then
    MbeExec
  else if hstep == RslStep(7) then
    CheckViewTimeout
  else if hstep == RslStep(8) then
    CheckQuorumOfViewSuspicion
  else if hstep == RslStep(9) then
    MbeSendHeartbeat
  else
    L
}

function {:verify false} Rsl_NoRecvPerfUpdate(node_pr:Timestamp, hstep:RslStep) : Timestamp
{
  var total_time := TimeAdd2(node_pr, StepToTimeDelta(hstep));
  total_time
}

function {:verify false} Rsl_RecvPerfUpdate(node_ts:Timestamp, pkt_ts:Timestamp, hstep:RslStep) : Timestamp
{
  var handlerStartTime := TimeMax(pkt_ts, node_ts);
  var total_time := TimeAdd2(handlerStartTime, StepToTimeDelta(hstep));
  total_time
}

function {:verify false} Rsl_TimeoutPerfUpdate(node_ts:Timestamp, hstep:RslStep) : Timestamp
{
  node_ts + Timeout() + StepToTimeDelta(hstep)
}

predicate {:verify false} TimeEq(p1:Timestamp, p2:Timestamp)
{
  p1 == p2
}

predicate {:verify false} TimeLe(p1:Timestamp, p2:Timestamp)
{
  p1 <= p2
}

function {:verify false} TimeBound1aSent() : Timestamp
{
  Timeout() + ProcessPacket + MbeNewView
}

function {:verify false} TimeBound1aDelivery() : Timestamp
{
  TimeBound1aSent() + D
}

function {:verify false} TimeBound1aSelfDelivery() : Timestamp
{
  TimeBound1aSent() + SelfDelivery
}

function TimeBound1bDelivery() : Timestamp
{
  TimeBound1aDelivery() + ProcessPacket + TimeActionRange(0) + D
}

function {:verify false} {:fuel 2} TimeActionRangeInclusive(prevActionIndex:int) : Timestamp
  requires prevActionIndex >= 0
{
  if prevActionIndex == 0 then
    StepToTimeDelta(RslStep(prevActionIndex))
  else 
    TimeActionRangeInclusive(prevActionIndex - 1) + StepToTimeDelta(RslStep(prevActionIndex))
}

function {:verify false} TimeActionRange(nextActionIndex:int) : Timestamp
  requires 0 <= nextActionIndex < 10
{
  if nextActionIndex == 0 then
    TimeActionRangeInclusive(9)
  else
    TimeActionRangeInclusive(nextActionIndex - 1)
}

function TimeBoundPhase1Leader(dts:Timestamp, ell:int, nextActionIndex:int) : Timestamp
  requires ell >= 0
  requires 0 <= nextActionIndex < 10
{
  dts + (ell + 1) * TimeActionRange(0) + TimeActionRange(nextActionIndex)
}

lemma LeaderTimeoutPreservesPhase1Invariant(dts:Timestamp, ell:int, nextActionIndex:int)
  requires ell >= 0
  ensures dts + ProcessPacket <= TimeBoundPhase1Leader(dts, ell, nextActionIndex)
{
  
}

lemma {:verify false} TimeActionRangeHelper_NoRecv(dts:Timestamp, node_ts:Timestamp, nextActionIndex:int)
  requires 0 < nextActionIndex < 10
  requires node_ts <= dts + TimeActionRange(nextActionIndex)
  ensures node_ts + StepToTimeDelta(RslStep(nextActionIndex)) <= dts + TimeActionRange((nextActionIndex+1) % LReplicaNumActions())
{
}

lemma TimeActionRangeHelper_Recv(dts:Timestamp, node_ts:Timestamp, dts':Timestamp, node_ts':Timestamp)
  requires node_ts <= dts + TimeActionRange(0);
  requires dts' >= node_ts;
  requires node_ts' == TimeMax(dts', node_ts) + StepToTimeDelta(RslStep(0))
  ensures node_ts' <= dts' + TimeActionRange(1);
{
}

lemma BoundedLagImpliesBoundedProcessingTime(dts:Timestamp, node_ts:Timestamp, pkt_ts:Timestamp, node_ts':Timestamp, lag:Timestamp)
  requires node_ts <= dts + lag
  requires node_ts' == Rsl_RecvPerfUpdate(node_ts, pkt_ts, RslStep(0));
  requires pkt_ts >= dts;

  ensures node_ts' <= (pkt_ts + lag + ProcessPacket)
{
}

lemma BoundedSizeLagImpliesBoundedProcessingTime(dts:Timestamp, node_ts:Timestamp, pkt_ts:Timestamp, node_ts':Timestamp, size:int)
  requires node_ts <= dts + size * TimeActionRange(0) + TimeActionRange(0)
  requires 0 <= size
  requires node_ts' == Rsl_RecvPerfUpdate(node_ts, pkt_ts, RslStep(0));
  requires pkt_ts >= dts;

  ensures node_ts' <= (pkt_ts + (size + 1) * TimeActionRange(0) + ProcessPacket)
{
}

}
