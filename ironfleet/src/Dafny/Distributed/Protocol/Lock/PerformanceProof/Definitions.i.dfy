include "../../../Services/Lock/LockTaggedDistributedSystem.i.dfy"
  include "TaggedGLS.i.dfy"

module PerformanceProof__Definitions_i {
import opened LockTaggedDistributedSystem_i
  // import opened TaggedGLS_i

function PerfBoundLockHeld(epoch: int) : PerfReport
  requires 0 < epoch
{
  var s : multiset<PerformanceReport> := multiset{};
  var s2 := s[PerfStep(GrantStep) := epoch - 1][PerfStep(AcceptStep) := epoch - 1][PerfDelivery := epoch - 1];
  PerfAdd(s2)
}

function PerfBoundLockInNetwork(epoch: int) : PerfReport
  requires 1 < epoch
{
  var s : multiset<PerformanceReport> := multiset{};
  var s2 := s[PerfStep(GrantStep) := epoch - 1][PerfStep(AcceptStep) := epoch - 2][PerfDelivery := epoch - 2];
  PerfAdd(s2)
}

lemma {:verify false} specific_axiom(a:multiset<PerfExpr>, b:multiset<PerfExpr>)
  ensures PerfEq(PerfAdd(b + multiset{PerfAdd(a)}), PerfAdd(b + a))
{
}

lemma Test(ios:seq<TaggedLIoOp<EndPoint, LockMessage>>, j:int)
  requires 1 < j
{
  PerfProperties();
  // var p := PerfBoundLockHeld(j)
  // var p' := PerfBoundLockInNetwork(j + 1);

  assume |ios| == 1 && ios[0].LIoOpReceive?;
  assume PerfEq(ios[0].r.msg.pr, PerfBoundLockInNetwork(j));

  var p := PerfVoid;
  var p' := PerfBoundLockHeld(j);

  var deliveryTime := PerfAdd2(ios[0].r.msg.pr, PerfDelivery);
  var handlerStartTime := PerfMax(multiset{deliveryTime, p});
  var totalTime := PerfAdd2(handlerStartTime, PerfStep(AcceptStep));

  assert PerfEq(totalTime, p');
}

// if |ios| > 0 && ios[0].LIoOpReceive? then
  // var deliveryTime := PerfAdd2(ios[0].pr, PerfDeliver);
  // var handlerStartTime := PerfMax(multiset{deliveryTime, tagged_node_state.pr});
  // var totalTime := PerfAdd2(handlerStartTime, PerfStep(GrantStep));
  //
// else
  // var totalTime := PerfAdd2(ios[0].pr, PerfStep(GrantStep));

  
}
