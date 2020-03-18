include "../../../Services/Lock/LockTaggedDistributedSystem.i.dfy"
  include "TaggedGLS.i.dfy"

module PerformanceProof__Definitions_i {
import opened LockTaggedDistributedSystem_i
  // import opened TaggedGLS_i

function PerfBoundLockHeld(epoch: int) : PerfReport
  requires 0 <= epoch
{
  var s : multiset<PerformanceReport> := multiset{};
  var s2 := s[GetStepRuntime(GrantStep) := epoch];
  PerfAdd(s2)
}

function PerfBoundLockInNetwork(epoch: int) : PerfReport
  requires 0 < epoch
{
  var s : multiset<PerformanceReport> := multiset{};
  var s2 := s[GetStepRuntime(GrantStep) := epoch - 1];
  PerfAdd2(PerfAdd(s2), GetStepRuntime(GrantStep))
}
  
}
