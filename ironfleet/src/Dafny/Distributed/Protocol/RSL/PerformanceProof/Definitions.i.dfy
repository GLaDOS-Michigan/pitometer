include "../DistributedSystem.i.dfy"

include "../../../Common/Collections/Maps2.i.dfy"
include "../Constants.i.dfy"
include "../Environment.i.dfy"
include "../Replica.i.dfy"

include "TimestampedRslSystem.i.dfy"

module Definitions_i {
import opened RslTimestampedDistributedSystem_i

// Generic action step time
ghost const L:Timestamp

ghost const Ds := DeliveryTime();
ghost const D:Timestamp

 //ghost const Gs := StepTime(RslStep);
 //ghost const As := StepTime(AcceptStep);
 //ghost const G:Timestamp
 //ghost const A:Timestamp

function StepToTimeDelta(hstep:RslStep) : Timestamp
{
  L
}

function Rsl_NoRecvPerfUpdate(node_pr:Timestamp, hstep:RslStep) : Timestamp
{
  var total_time := TimeAdd2(node_pr, StepToTimeDelta(hstep));
  total_time
}

function Rsl_RecvPerfUpdate(node_pr:Timestamp, pkt_pr:Timestamp, hstep:RslStep) : Timestamp
{
  var deliveryTime := TimeAdd2(pkt_pr, D);
  var handlerStartTime := TimeMax(deliveryTime, node_pr);
  var total_time := TimeAdd2(handlerStartTime, StepToTimeDelta(hstep));
  total_time
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
