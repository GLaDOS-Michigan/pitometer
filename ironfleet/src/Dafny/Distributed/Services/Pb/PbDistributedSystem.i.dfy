include "../../Impl/Pb/Host.i.dfy"
include "../../Common/Framework/DistributedSystem.s.dfy"

module Pb_DistributedSystem_i refines DistributedSystem_s {
    import opened H_s = Host_i
}
