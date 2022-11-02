include "../../Common/Framework/Host.s.dfy"
include "../../Protocol/Pb/Types.i.dfy"

module Host_i refines Host_s {
  import opened Types_i
    type HostStep = PbStep
    // type ConcreteConfiguration = ConstantsState
}
