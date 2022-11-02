include "../DistributedSystem.i.dfy"

include "../../../Common/Collections/Maps2.i.dfy"
include "../Node.i.dfy"
include "../../../Services/Pb/PbTimestampedDistributedSystem.i.dfy"

module Definitions_i {
import opened PbTimestampedDistributedSystem_i

const ZeroT:Timestamp := 0
// Generic action step time
// ghost const L:Timestamp
ghost const D:Timestamp
ghost const PrimaryReqT:Timestamp
ghost const PrimaryAckT:Timestamp
ghost const BackupProcessT:Timestamp
ghost const MaxQ:Timestamp

 //ghost const Gs := StepTime(RslStep);
 //ghost const As := StepTime(AcceptStep);
 //ghost const G:Timestamp
 //ghost const A:Timestamp

function StepToTimeDelta(hstep:PbStep) : Timestamp
{
  if hstep == PrimaryReqStep then
    PrimaryReqT
  else if hstep == PrimaryRecvStep then
    PrimaryAckT
  else
    assert (hstep == BackupRecvStep);
    BackupProcessT
}

function Pb_NoRecvPerfUpdate(node_pr:Timestamp, hstep:PbStep) : Timestamp
{
  var total_time := TimeAdd2(node_pr, StepToTimeDelta(hstep));
  total_time
}

function Pb_RecvPerfUpdate(node_ts:Timestamp, pkt_ts:Timestamp, hstep:PbStep) : Timestamp
{
  var handlerStartTime := TimeMax(pkt_ts, node_ts);
  var total_time := TimeAdd2(handlerStartTime, StepToTimeDelta(hstep));
  total_time
}

function Pb_TimeoutPerfUpdate(node_ts:Timestamp, hstep:PbStep) : Timestamp
{
  node_ts + Timeout() + StepToTimeDelta(hstep)
}

predicate TimeEq(p1:Timestamp, p2:Timestamp)
{
  p1 == p2
}

predicate TimeLe(p1:Timestamp, p2:Timestamp)
{
  p1 <= p2
}

}
