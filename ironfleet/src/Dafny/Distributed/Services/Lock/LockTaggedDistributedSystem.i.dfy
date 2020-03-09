include "../../Impl/Lock/Host.i.dfy"
  include "../../Common/Framework/TaggedDistributedSystem.s.dfy"
  include "LockDistributedSystem.i.dfy"

module LockTaggedDistributedSystem_i refines TaggedDistributedSystem_s {

  import opened DS_s = Lock_DistributedSystem_i

}
