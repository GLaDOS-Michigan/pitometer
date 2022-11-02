include "../DistributedSystem.i.dfy"

include "../../../Services/Pb/PbTimestampedDistributedSystem.i.dfy"
include "Definitions.i.dfy"

module TimestampedPbSystem_i {

import opened PbTimestampedDistributedSystem_i
import opened Pb_DistributedSystem_i
import opened Protocol_DistributedSystem_i
import opened Definitions_i

// type TimestampedNode = TimestampedType<LScheduler>
datatype TimestampedNode = TimestampedNode(v:Node, ts:Timestamp, dts:Timestamp)
type TimestampedPbPacket = TimestampedLPacket<EndPoint, PbMessage>
type TimestampedPbEnvironment = TimestampedLEnvironment<EndPoint, PbMessage, PbStep>


// type UndeliveredPackets = multiset<TimestampedLPacket<EndPoint, PbMessage>>
type UndeliveredPackets = set<TimestampedLPacket<EndPoint, PbMessage>>
function UndeliveredPacketsEmpty() : UndeliveredPackets { {} }

datatype TimestampedPbState = TimestampedPbState(
    constants:Constants,
    t_environment:TimestampedLEnvironment<EndPoint, PbMessage, PbStep>,
    t_nodes:seq<TimestampedNode>
    )

function UntimestampPbNodes(t_nodes: seq<TimestampedNode>) : seq<Node>
{
  MapSeqToSeq(t_nodes, (tls:TimestampedNode) => tls.v)
}

function UntimestampPbState(tps:TimestampedPbState) : PbState
{
  PbState(
    tps.constants,
    UntagLEnvironment(tps.t_environment),
    UntimestampPbNodes(tps.t_nodes)
    )
}


predicate TimestampedPbMapsComplete(ps:TimestampedPbState)
{
    |ps.t_nodes| == |ps.constants.config|
}

predicate TimestampedPbConstantsUnchanged(ps:TimestampedPbState, ps':TimestampedPbState)
{
  |ps'.t_nodes| == |ps.t_nodes|
    && ps'.constants == ps.constants
}

predicate TimestampedPbInit(con:Constants, ps:TimestampedPbState)
{
  PbInit(con, UntimestampPbState(ps))
    && (forall i :: 0 <= i < |con.config| ==> ps.t_nodes[i].ts == TimeZero()
    && ps.t_nodes[i].dts == TimeZero())

    && LEnvironment_Init(ps.t_environment)
    && TimestampedPbMapsComplete(ps)
}

predicate TimestampedPbNextCommon(ps:TimestampedPbState, ps':TimestampedPbState)
{
       TimestampedPbMapsComplete(ps)
    && TimestampedPbConstantsUnchanged(ps, ps')
    && LEnvironment_Next(ps.t_environment, ps'.t_environment)
}

predicate TimestampedPbNextOneReplica(ps:TimestampedPbState, ps':TimestampedPbState, idx:int, ios:seq<TimestampedLIoOp<EndPoint, PbMessage>>)
{
  PbNextOneReplica(UntimestampPbState(ps), UntimestampPbState(ps'), idx, UntagLIoOpSeq(ios))
    && TimestampedPbNextCommon(ps, ps')
    && 0 <= idx < |ps.constants.config|
    && ps.t_environment.nextStep.LEnvStepHostIos?
    && ps.t_environment.nextStep.actor == ps.constants.config[idx]
    && ps.t_environment.nextStep.ios == ios
    // && ps.t_environment.nextStep == LEnvStepHostIos(ps.constants.config[idx], ios, PbStep(ps.t_nodes[idx].v.nextActionIndex))
    && ps'.t_nodes == ps.t_nodes[idx := ps'.t_nodes[idx]]

    && var hstep := ps.t_environment.nextStep.nodeStep;
    (if |ios| > 0 && ios[0].LIoOpReceive? then
      && ios[0] in ios
      && ps'.t_nodes[idx].ts == Pb_RecvPerfUpdate(ps.t_nodes[idx].ts, ios[0].r.msg.ts, hstep)
      && (ps.t_nodes[idx].dts <= ios[0].r.msg.ts <= ps.t_nodes[idx].ts + Timeout())
      && ps'.t_nodes[idx].dts == ios[0].r.msg.ts
    else if |ios| > 0 && ios[0].LIoOpTimeoutReceive? then
      && ps'.t_nodes[idx].ts == Pb_TimeoutPerfUpdate(ps.t_nodes[idx].ts, hstep)
      && ps'.t_nodes[idx].dts == ps.t_nodes[idx].ts + Timeout()
    else
      && ps'.t_nodes[idx].ts == Pb_NoRecvPerfUpdate(ps.t_nodes[idx].ts, hstep)
      && ps'.t_nodes[idx].dts == ps.t_nodes[idx].dts
      )
      && (forall io :: io in ios && io.LIoOpSend? ==>
          io.s.msg.ts  == ps'.t_nodes[idx].ts + D)
}

predicate TimestampedPbNextEnvironment(ps:TimestampedPbState, ps':TimestampedPbState)
{
       TimestampedPbNextCommon(ps, ps')
    && !ps.t_environment.nextStep.LEnvStepHostIos?
    && ps'.t_nodes == ps.t_nodes
}

predicate TimestampedPbNext(ps:TimestampedPbState, ps':TimestampedPbState)
{
       (exists idx, ios :: TimestampedPbNextOneReplica(ps, ps', idx, ios))
    || TimestampedPbNextEnvironment(ps, ps')
}

}
