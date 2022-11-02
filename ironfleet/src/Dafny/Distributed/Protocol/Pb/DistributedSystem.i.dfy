include "../../Common/Collections/Maps2.i.dfy"
include "Node.i.dfy"

module Protocol_DistributedSystem_i {
import opened Collections__Maps2_i
import opened Protocol_Node_i

type HostStep = PbStep

datatype Constants = Constants(
  config:Config,
  client: EndPoint
  )

datatype PbState = PbState(
    constants:Constants,
    environment:LEnvironment<EndPoint, PbMessage, PbStep>,
    nodes:seq<Node>
    )

predicate PbMapsComplete(ps:PbState)
{
    |ps.nodes| == |ps.constants.config|
}

predicate PbInit(constants:Constants, ps:PbState)
{
  ps.constants == constants
    && LEnvironment_Init(ps.environment)
    && PbMapsComplete(ps)
    && |ps.nodes| == |constants.config|
    && (forall i :: 0 <= i < |constants.config| ==> NodeInit(ps.nodes[i], i, constants.config, constants.client))
}

predicate PbNextCommon(ps:PbState, ps':PbState)
{
  |ps.nodes| == |ps'.nodes|
    && LEnvironment_Next(ps.environment, ps'.environment)
}

predicate PbNextOneReplica(ps:PbState, ps':PbState, idx:int, ios:seq<PbIo>)
{
  PbNextCommon(ps, ps')
    && PbMapsComplete(ps)
    && 0 <= idx < |ps.nodes|
    && NodeNext(ps.nodes[idx], ps'.nodes[idx], ios)
    && ps.environment.nextStep ==
      LEnvStepHostIos(ps.constants.config[idx], ios,
      (if idx == 0 then
        if ps.nodes[idx].nextStep == 0 then
          PrimaryReqStep()
        else
          PrimaryRecvStep()
      else
        BackupRecvStep))
    && ps'.nodes == ps.nodes[idx := ps'.nodes[idx]]
}

predicate PbNextEnvironment(ps:PbState, ps':PbState)
{
       PbNextCommon(ps, ps')
    && !ps.environment.nextStep.LEnvStepHostIos?
    && ps'.nodes == ps.nodes
}

predicate PbNext(ps:PbState, ps':PbState)
{
  (exists idx, ios :: PbNextOneReplica(ps, ps', idx, ios))
    || PbNextEnvironment(ps, ps')
}

}
