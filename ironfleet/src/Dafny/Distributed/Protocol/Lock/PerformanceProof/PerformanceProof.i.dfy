include "../../../Services/Lock/LockTaggedDistributedSystem.i.dfy"

module PerformanceProof_i {
import opened LockTaggedDistributedSystem_i

predicate dummy(config:ConcreteConfiguration, tdb:seq<TaggedDS_State>)
{
  false
}
  
}
