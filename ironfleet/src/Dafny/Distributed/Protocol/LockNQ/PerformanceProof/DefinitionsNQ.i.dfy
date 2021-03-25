// Must be verified with /arith:2

include "../../../Services/Lock/LockTimestampedDistributedSystem.i.dfy"

module PerformanceProof__Definitions_i {
import opened LockTimestampedDistributedSystem_i

ghost const Gs := StepTime(GrantStep);
ghost const As := StepTime(AcceptStep);
ghost const Ds := DeliveryTime();
ghost const G:Timestamp
ghost const A:Timestamp
ghost const D:Timestamp

function StepToTimeDelta(hstep:HostStep) : Timestamp
{
  if hstep == GrantStep then G else A
}

function TLS_NoRecvPerfUpdate(node_pr:Timestamp, hstep:HostStep) : Timestamp
{
  var total_time := TimeAdd2(node_pr, StepToTimeDelta(hstep));
  total_time
}

function TLS_RecvPerfUpdate(node_pr:Timestamp, pkt_pr:Timestamp, hstep:HostStep) : Timestamp
{
  var deliveryTime := TimeAdd2(pkt_pr, D);  // add D to packet sent time
  // var handlerStartTime := TimeMax(deliveryTime, node_pr);   
  var total_time := TimeAdd2(deliveryTime, StepToTimeDelta(hstep));
  total_time
}

function PerfBoundLockHeld(epoch: int) : Timestamp
  requires 0 < epoch
  ensures PerfBoundLockHeld(epoch) >= 0
{
  (epoch - 1) * G + (epoch - 1) * A + (epoch - 1) * D
}

function PerfBoundLockInNetwork(epoch: int) : Timestamp
  requires 0 < epoch
  ensures PerfBoundLockInNetwork(epoch) >= 0
{
  if epoch == 1 then  // note that a valid Transfer packet never has epoch 0
    0
  else
    (epoch - 1) * G + (epoch - 2) * A + (epoch - 2) * D
}

predicate TimeEq(p1:Timestamp, p2:Timestamp)
{
  p1 == p2
}

predicate TimeLe(p1:Timestamp, p2:Timestamp)
{
  p1 <= p2
}

lemma Grant_j_helper_specific(epoch:int)
  requires epoch > 0
  ensures TimeEq(PerfBoundLockInNetwork(epoch + 1), TimeAdd2(PerfBoundLockHeld(epoch), G));
{
}

lemma Grant_j_helper()
  ensures forall epoch :: epoch > 0 ==> TimeEq(PerfBoundLockInNetwork(epoch + 1), TimeAdd2(PerfBoundLockHeld(epoch), G));
{
}

lemma Accept_j_helper()
  ensures forall epoch :: epoch > 1 ==> TimeEq(PerfBoundLockHeld(epoch), TLS_RecvPerfUpdate(TimeVoid(), PerfBoundLockInNetwork(epoch), AcceptStep))
{
}
}
