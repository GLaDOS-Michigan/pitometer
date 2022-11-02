include "../../Common/Framework/TimestampedDistributedSystem.s.dfy"
include "PbDistributedSystem.i.dfy"

module PbTimestamp_s refines Timestamp_s {
  import opened DS_s = Pb_DistributedSystem_i
}
