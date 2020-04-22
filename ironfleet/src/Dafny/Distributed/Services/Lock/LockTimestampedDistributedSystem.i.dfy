include "../../Impl/Lock/Host.i.dfy"
  include "../../Common/Framework/TimestampedDistributedSystem.s.dfy"
  include "LockDistributedSystem.i.dfy"
  include "LockTimestamp.s.dfy"

module LockTimestampedDistributedSystem_i refines TimestampedDistributedSystem_s {

  import opened T_s = LockTimestamp_s

}
