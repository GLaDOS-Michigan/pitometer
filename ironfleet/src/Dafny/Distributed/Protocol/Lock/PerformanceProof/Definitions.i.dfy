include "../../../Services/Lock/LockTaggedDistributedSystem.i.dfy"
  include "TaggedGLS.i.dfy"

module PerformanceProof__Definitions_i {
import opened LockTaggedDistributedSystem_i
  // import opened TaggedGLS_i

function PerfBound(epoch: int) : PerfReport
  requires 0 <= epoch
{
  if epoch == 0 then
    PerfNone()
  else
    PerfAdd2(GetStepRuntime(GrantStep), PerfBound(epoch - 1))
}

  
}
