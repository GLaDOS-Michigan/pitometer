// Must be verified with /arith:2

include "../../../Services/Lock/LockTaggedDistributedSystem.i.dfy"

module PerformanceProof__Definitions_i {
import opened LockTaggedDistributedSystem_i

ghost const Gs := PerfStep(GrantStep);
ghost const As := PerfStep(AcceptStep);
ghost const Ds := PerfDelivery();
ghost const G:PerfObject
ghost const A:PerfObject
ghost const D:PerfObject

function StepToTimeDelta(hstep:HostStep) : PerfReport
{
  if hstep == GrantStep then G else A
}

function TLS_NoRecvPerfUpdate(node_pr:PerfReport, hstep:HostStep) : PerfExpr
{
  var total_time := PerfAdd2(node_pr, StepToTimeDelta(hstep));
  total_time
}

function TLS_RecvPerfUpdate(node_pr:PerfExpr, pkt_pr:PerfExpr, hstep:HostStep) : PerfExpr
{
  var deliveryTime := PerfAdd2(pkt_pr, D);
  var handlerStartTime := PerfMax(deliveryTime, node_pr);
  var total_time := PerfAdd2(handlerStartTime, StepToTimeDelta(hstep));
  total_time
}

function PerfBoundLockHeld(epoch: int) : PerfExpr
  requires 0 < epoch
{
  (epoch - 1) * G + (epoch - 1) * A + (epoch - 1) * D
}

function PerfBoundLockInNetwork(epoch: int) : PerfExpr
  requires 1 < epoch
{
  (epoch - 1) * G + (epoch - 2) * A + (epoch - 2) * D
}

predicate PerfEq(p1:PerfObject, p2:PerfObject)
{
  p1 == p2
}

predicate PerfLe(p1:PerfObject, p2:PerfObject)
{
  p1 <= p2
}

lemma Grant_j_helper_specific(epoch:int)
  requires epoch > 0
  ensures PerfEq(PerfBoundLockInNetwork(epoch + 1), PerfAdd2(PerfBoundLockHeld(epoch), G));
{
}

lemma Grant_j_helper()
  ensures forall epoch :: epoch > 0 ==> PerfEq(PerfBoundLockInNetwork(epoch + 1), PerfAdd2(PerfBoundLockHeld(epoch), G));
{
}

lemma Accept_j_helper()
  ensures forall epoch :: epoch > 1 ==> PerfEq(PerfBoundLockHeld(epoch), TLS_RecvPerfUpdate(PerfVoid(), PerfBoundLockInNetwork(epoch), AcceptStep))
{
}

// lemma lemma_Accept_j_PR(pkt_pr:PerfReport, node_pr:PerfReport, node_pr':PerfReport, epoch:int, num_steps:NumSteps)
//   requires 1 < epoch
//   requires node_pr' == RecvPerfUpdate(node_pr, pkt_pr, AcceptStep, num_steps).0;
//   requires pkt_pr == PerfBoundLockInNetwork(epoch)
//   requires node_pr == PerfVoid()
//   ensures PerfEq(node_pr', PerfBoundLockHeld(epoch))
// {
//   PerfProperties(); 
// 
//   var s : multiset<PerfExpr> := multiset{};
//   var s2 := s[PerfStep(GrantStep) := epoch - 1];
//   var s3 := s[PerfStep(AcceptStep) := epoch - 2];
//   var s3' := s[PerfStep(AcceptStep) := epoch - 1];
//   var s4 := s[PerfDelivery() := epoch - 2];
//   var s4' := s[PerfDelivery() := epoch - 1];
//   var sNetwork := s2 + s3 + s4;
//   var sDelivery := s2 + s3 + s4';
//   var sHeld := s2 + s3' + s4';
// 
//   assert sDelivery == sNetwork + multiset{PerfDelivery()};
//   assert sHeld == sDelivery + multiset{PerfStep(AcceptStep)};
// }

/*

lemma Test(j:int)
  requires 1 < j
{
  PerfProperties();
  quotient_axioms();
  var p := quotient_map(PerfVoid);
  var p' := quotient_map(PerfBoundLockHeld(j));
  var pkt_pr := quotient_map(PerfBoundLockInNetwork(j));

  var deliveryTime := PerfAdd(pkt_pr, quotient_map(PerfDelivery));
  var handlerStartTime := PerfMax(deliveryTime, p);
  var totalTime := PerfAdd(handlerStartTime, quotient_map(PerfStep(AcceptStep)));

  assert totalTime == p';
}

lemma specific_axiom(a:multiset<PerfExpr>, b:multiset<PerfExpr>)
  ensures PerfEq(PerfAdd(b + multiset{PerfAdd(a)}), PerfAdd(b + a))
{
}

lemma Other(j:int, pkt_pr: PerfReport, nd_pr:PerfReport, final_pr:PerfReport)
  requires 1 < j;
  requires PerfEq(pkt_pr, PerfBoundLockInNetwork(j));
  // requires PerfEq(nd_pr, PerfVoid);
  requires nd_pr == PerfVoid;
  ensures PerfEq(PerfAdd2(PerfMax(multiset{PerfAdd2(pkt_pr, PerfDelivery), nd_pr}), PerfStep(AcceptStep)), PerfBoundLockHeld(j));
{
  Test(j, pkt_pr, nd_pr, final_pr);
}


lemma Test(j:int, pkt_pr: PerfReport, nd_pr:PerfReport, final_pr:PerfReport)
  requires 1 < j;
  // requires PerfEq(pkt_pr, PerfBoundLockInNetwork(j));
  requires pkt_pr == PerfBoundLockInNetwork(j);
  // requires PerfEq(nd_pr, PerfVoid);
  requires nd_pr == PerfVoid;
  ensures PerfEq(PerfAdd2(PerfMax(multiset{PerfAdd2(pkt_pr, PerfDelivery), nd_pr}), PerfStep(AcceptStep)), PerfBoundLockHeld(j));

{
  PerfProperties();
// 
  // var p := PerfVoid;
  // var p' := PerfBoundLockHeld(j);
  // // var pkt_pr := PerfBoundLockInNetwork(j);
// 
  // var deliveryTime := PerfAdd2(pkt_pr, PerfDelivery);
  // var handlerStartTime := PerfMax(multiset{deliveryTime, p});
  // var totalTime := PerfAdd2(handlerStartTime, PerfStep(AcceptStep));
// 
  // assert PerfEq(totalTime, p');
}

// if |ios| > 0 && ios[0].LIoOpReceive? then
  // var deliveryTime := PerfAdd2(ios[0].pr, PerfDeliver);
  // var handlerStartTime := PerfMax(multiset{deliveryTime, tagged_node_state.pr});
  // var totalTime := PerfAdd2(handlerStartTime, PerfStep(GrantStep));
  //
// else
  // var totalTime := PerfAdd2(ios[0].pr, PerfStep(GrantStep));

*/
  
}
