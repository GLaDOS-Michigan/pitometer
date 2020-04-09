include "../../../Services/Lock/LockTaggedDistributedSystem.i.dfy"
  //include "TaggedGLS.i.dfy"

module PerformanceProof__Definitions_i {
 import opened LockTaggedDistributedSystem_i
  // import opened TaggedGLS_i

function PerfBoundLockHeld(epoch: int) : PerfExpr
  requires 0 < epoch
{
    var s : multiset<PerfExpr> := multiset{};
    var s2 := s[PerfStep(GrantStep) := epoch - 1][PerfStep(AcceptStep) := epoch - 1][PerfDelivery := epoch - 1];
    PerfAdd(s2)
}

function PerfBoundLockInNetwork(epoch: int) : PerfExpr
  requires 1 < epoch
{
  var s : multiset<PerfExpr> := multiset{};
  var s2 := s[PerfStep(GrantStep) := epoch - 1][PerfStep(AcceptStep) := epoch - 2][PerfDelivery := epoch - 2];
  PerfAdd(s2)
}

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

/*
lemma {:verify false} specific_axiom(a:multiset<PerfExpr>, b:multiset<PerfExpr>)
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
