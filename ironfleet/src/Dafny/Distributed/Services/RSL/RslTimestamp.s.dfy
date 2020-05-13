  include "../../Common/Framework/TimestampedDistributedSystem.s.dfy"
  include "RSLDistributedSystem.i.dfy"

module RslTimestamp_s refines Timestamp_s {
  import opened DS_s = RSL_DistributedSystem_i

}
