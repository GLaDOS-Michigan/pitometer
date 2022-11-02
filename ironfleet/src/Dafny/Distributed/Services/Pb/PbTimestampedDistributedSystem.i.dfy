include "../../Common/Framework/TimestampedDistributedSystem.s.dfy"

  include "PbDistributedSystem.i.dfy"
  include "PbTimestamp.s.dfy"

module PbTimestampedDistributedSystem_i refines TimestampedDistributedSystem_s {
    import opened T_s = PbTimestamp_s
}
