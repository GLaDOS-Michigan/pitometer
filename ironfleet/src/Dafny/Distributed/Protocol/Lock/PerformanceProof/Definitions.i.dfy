include "../../../Services/Lock/LockTaggedDistributedSystem.i.dfy"
  include "TaggedGLS.i.dfy"

module PerformanceProof__Definitions_i {
import opened LockTaggedDistributedSystem_i
  // import opened TaggedGLS_i

function PerfBoundLockHeld(epoch: int) : PerfReport
  requires 0 < epoch
{
  var s : multiset<PerformanceReport> := multiset{};
  var s2 := s[PerfStep(GrantStep) := epoch - 1][PerfStep(AcceptStep) := epoch - 1];
  PerfAdd(s2)
}

function PerfBoundLockInNetwork(epoch: int) : PerfReport
  requires 1 < epoch
{
  var s : multiset<PerformanceReport> := multiset{};
  var s2 := s[PerfStep(GrantStep) := epoch - 1][PerfStep(AcceptStep) := epoch - 2];
  PerfAdd(s2)
}

lemma {:verify false} specific_axiom(a:multiset<PerfExpr>, b:multiset<PerfExpr>)
  ensures PerfEq(PerfAdd(b + multiset{PerfAdd(a)}), PerfAdd(b + a))
{
}

lemma Test(j:int)
  requires 0 <= j
{
  var p := PerfBoundLockHeld(j);
  var p' := PerfBoundLockInNetwork(j + 1);

  var s : multiset<PerformanceReport> := multiset{};
  var s2 := s[PerfStep(GrantStep) := j][PerfStep(AcceptStep) := j];
  var s3 := s[PerfStep(GrantStep) := j + 1][PerfStep(AcceptStep) := j];

  assert s3 == s2 + multiset{PerfStep(GrantStep)};

  assert PerfEq(p', PerfAdd2(p, PerfStep(GrantStep)));

  var p2 := PerfBoundLockHeld(j);
  var p2' := PerfBoundLockInNetwork(j + 1);

  assert p2' == PerfAdd2(p2, PerfStep(AcceptStep));
}

  
}
