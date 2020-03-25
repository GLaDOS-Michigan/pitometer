include "../../Impl/Lock/Host.i.dfy"
  include "../../Common/Framework/TaggedDistributedSystem.s.dfy"
  include "LockDistributedSystem.i.dfy"
  include "LockPerformance.s.dfy"

module LockTaggedDistributedSystem_i refines TaggedDistributedSystem_s {

  import opened P_s = LockPerformance_i

}
