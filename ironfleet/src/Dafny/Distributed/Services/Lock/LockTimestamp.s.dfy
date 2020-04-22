include "../../Impl/Lock/Host.i.dfy"
  include "../../Common/Framework/TimestampedDistributedSystem.s.dfy"
  include "LockDistributedSystem.i.dfy"

module LockTimestamp_s refines Timestamp_s {
  import opened DS_s = Lock_DistributedSystem_i

}
