include "../DistributedSystem.i.dfy"

include "../../../Common/Collections/Maps2.i.dfy"
include "../Constants.i.dfy"
include "../Environment.i.dfy"
include "../Replica.i.dfy"
include "../../../Services/RSL/RslTimestampedDistributedSystem.i.dfy"
include "Definitions.i.dfy"

module TimestampedRslPerformanceProof_i {

import opened RslTimestampedDistributedSystem_i
  import opened LiveRSL__DistributedSystem_i
  import opened Definitions_i

import opened Collections__Maps2_i
import opened LiveRSL__Constants_i
import opened LiveRSL__Environment_i
import opened LiveRSL__Replica_i

type TimestampedLScheduler = TimestampedType<LScheduler>

datatype TimestampedRslState = TimestampedRslState(
    constants:LConstants,
    t_environment:TimestampedLEnvironment<NodeIdentity, RslMessage, RslStep>,
    t_replicas:seq<TimestampedLScheduler>,
    clients:set<NodeIdentity>
    )

function UntimestampRslReplicas(t_replicas: seq<TimestampedLScheduler>) : seq<LScheduler>
{
  MapSeqToSeq(t_replicas, (tls:TimestampedLScheduler) => tls.v)
}

function UntimestampRslState(tps:TimestampedRslState) : RslState
{
  RslState(
    tps.constants,
    UntagLEnvironment(tps.t_environment),
    UntimestampRslReplicas(tps.t_replicas),
    tps.clients
    )
}


predicate TimestampedRslMapsComplete(ps:TimestampedRslState)
{
    |ps.t_replicas| == |ps.constants.config.replica_ids|
}

predicate TimestampedRslConstantsUnchanged(ps:TimestampedRslState, ps':TimestampedRslState)
{
       |ps'.t_replicas| == |ps.t_replicas|
    && ps'.clients == ps.clients
    && ps'.constants == ps.constants
}

predicate TimestampedRslInit(con:LConstants, ps:TimestampedRslState)
{
  RslInit(con, UntimestampRslState(ps))
    && (forall i :: 0 <= i < |con.config.replica_ids| ==> ps.t_replicas[i].ts == TimeZero())

    && LEnvironment_Init(ps.t_environment)
    && TimestampedRslMapsComplete(ps)
}

predicate TimestampedRslNextCommon(ps:TimestampedRslState, ps':TimestampedRslState)
{
       TimestampedRslMapsComplete(ps)
    && TimestampedRslConstantsUnchanged(ps, ps')
    && LEnvironment_Next(ps.t_environment, ps'.t_environment)
}

predicate TimestampedRslNextOneReplica(ps:TimestampedRslState, ps':TimestampedRslState, idx:int, ios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>)
{
  RslNextOneReplica(UntimestampRslState(ps), UntimestampRslState(ps'), idx, UntagLIoOpSeq(ios))

    // LS_NextOneServer(UntagLS_State(tls), UntagLS_State(tls'), id, UntagLIoOpSeq(ios), hstep)
// 
      // && (forall t_io :: t_io in ios && t_io.LIoOpSend? ==> t_io.s.msg.ts == tls'.t_servers[id].ts)
      // && tls'.t_servers == tls.t_servers[id := tls'.t_servers[id]]

    && TimestampedRslNextCommon(ps, ps')
    && 0 <= idx < |ps.constants.config.replica_ids|
    && ps.t_environment.nextStep == LEnvStepHostIos(ps.constants.config.replica_ids[idx], ios, RslStep(ps.t_replicas[idx].v.nextActionIndex))
    && ps'.t_replicas == ps.t_replicas[idx := ps'.t_replicas[idx]]

    && var hstep := ps.t_environment.nextStep.nodeStep; (if |ios| > 0 && ios[0].LIoOpReceive? then
      ps'.t_replicas[idx].ts == Rsl_RecvPerfUpdate(ps.t_replicas[idx].ts, ios[0].r.msg.ts, hstep)
    else
      ps'.t_replicas[idx].ts == Rsl_NoRecvPerfUpdate(ps.t_replicas[idx].ts, hstep)
      )
}

predicate TimestampedRslNextEnvironment(ps:TimestampedRslState, ps':TimestampedRslState)
{
       TimestampedRslNextCommon(ps, ps')
    && !ps.t_environment.nextStep.LEnvStepHostIos?
    && ps'.t_replicas == ps.t_replicas
}

predicate TimestampedRslNextOneExternal(ps:TimestampedRslState, ps':TimestampedRslState, eid:NodeIdentity, ios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>)

{
       TimestampedRslNextCommon(ps, ps')
    && eid !in ps.constants.config.replica_ids
    && ps.t_environment.nextStep == LEnvStepHostIos(eid, ios, ExternalStep())
    && ps'.t_replicas == ps.t_replicas
}

predicate TimestampedRslNext(ps:TimestampedRslState, ps':TimestampedRslState)
{
       (exists idx, ios :: TimestampedRslNextOneReplica(ps, ps', idx, ios))
    || (exists eid, ios :: TimestampedRslNextOneExternal(ps, ps', eid, ios))
    || TimestampedRslNextEnvironment(ps, ps')
}
}
