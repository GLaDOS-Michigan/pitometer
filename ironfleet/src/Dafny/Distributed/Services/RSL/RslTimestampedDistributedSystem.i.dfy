include "../../Impl/RSL/Host.i.dfy"
  include "../../Common/Framework/TimestampedDistributedSystem.s.dfy"

  include "RSLDistributedSystem.i.dfy"
  include "RslTimestamp.s.dfy"

module RslTimestampedDistributedSystem_i refines TimestampedDistributedSystem_s {

    import opened T_s = RslTimestamp_s
}
