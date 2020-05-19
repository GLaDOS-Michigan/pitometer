include "../DistributedSystem.i.dfy"

include "../../../Common/Collections/Maps2.i.dfy"
include "../Constants.i.dfy"
include "../Environment.i.dfy"
include "../Replica.i.dfy"
include "../../../Services/RSL/RslTimestampedDistributedSystem.i.dfy"
include "Definitions.i.dfy"

module TimestampedRslSystem_i {

import opened RslTimestampedDistributedSystem_i
  import opened LiveRSL__DistributedSystem_i
  import opened Definitions_i

import opened Collections__Maps2_i
import opened LiveRSL__Constants_i
import opened LiveRSL__Environment_i
import opened LiveRSL__Replica_i

// type TimestampedLScheduler = TimestampedType<LScheduler>
datatype TimestampedLScheduler = TimestampedLScheduler(v:LScheduler, ts:Timestamp, dts:Timestamp)
type TimestampedRslPacket = TimestampedLPacket<EndPoint, RslMessage>
type TimestampedRslEnvironment = TimestampedLEnvironment<NodeIdentity, RslMessage, RslStep>


// type UndeliveredPackets = multiset<TimestampedLPacket<EndPoint, RslMessage>>
type UndeliveredPackets = set<TimestampedLPacket<EndPoint, RslMessage>>
function UndeliveredPacketsEmpty() : UndeliveredPackets { {} } 

datatype TimestampedRslState = TimestampedRslState(
    constants:LConstants,
    t_environment:TimestampedLEnvironment<NodeIdentity, RslMessage, RslStep>,
    t_replicas:seq<TimestampedLScheduler>,
    clients:set<NodeIdentity>,
    undeliveredPackets:UndeliveredPackets
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
    && (forall i :: 0 <= i < |con.config.replica_ids| ==> ps.t_replicas[i].ts == TimeZero()
    && ps.t_replicas[i].dts == TimeZero())

    && LEnvironment_Init(ps.t_environment)
    && TimestampedRslMapsComplete(ps)
    && ps.undeliveredPackets == UndeliveredPacketsEmpty()
}

predicate TimestampedRslNextCommon(ps:TimestampedRslState, ps':TimestampedRslState)
{
       TimestampedRslMapsComplete(ps)
    && TimestampedRslConstantsUnchanged(ps, ps')
    && LEnvironment_Next(ps.t_environment, ps'.t_environment)
}

predicate NoPacketDuplication(s:TimestampedRslState)
{
  var nextStep := s.t_environment.nextStep;
  nextStep.LEnvStepHostIos? ==> (
  (forall io :: io in nextStep.ios && io.LIoOpReceive? ==> io.r in s.undeliveredPackets)
  )
}

predicate {:opaque} PacketDeliveredInOrder(pkt:TimestampedRslPacket, undeliveredPackets:UndeliveredPackets)
{
  && pkt in undeliveredPackets
  && (forall other_pkt :: other_pkt in undeliveredPackets ==> TimeLe(pkt.msg.ts, other_pkt.msg.ts)
  )
}

predicate UndeliveredPackets_Next(s:TimestampedRslState, s':TimestampedRslState)
  requires s.t_environment.nextStep.LEnvStepHostIos?
{
  var nextStep := s.t_environment.nextStep;
  var ios := nextStep.ios;
  (
    s'.undeliveredPackets == s.undeliveredPackets -
    (set io | io in ios && io.LIoOpReceive? :: io.r) +
    (set io | io in ios && io.LIoOpSend? :: io.s)

    && (ios[0].LIoOpReceive? ==> 
        && PacketDeliveredInOrder(ios[0].r, s.undeliveredPackets)
    )
  )
}

predicate TimestampedRslNextOneReplica(ps:TimestampedRslState, ps':TimestampedRslState, idx:int, ios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>)
{
  RslNextOneReplica(UntimestampRslState(ps), UntimestampRslState(ps'), idx, UntagLIoOpSeq(ios))
    && TimestampedRslNextCommon(ps, ps')
    && 0 <= idx < |ps.constants.config.replica_ids|
    && ps.t_environment.nextStep == LEnvStepHostIos(ps.constants.config.replica_ids[idx], ios, RslStep(ps.t_replicas[idx].v.nextActionIndex))
    && UndeliveredPackets_Next(ps, ps')
    && ps'.t_replicas == ps.t_replicas[idx := ps'.t_replicas[idx]]

    && var hstep := ps.t_environment.nextStep.nodeStep;
    (if |ios| > 0 && ios[0].LIoOpReceive? then
      && ios[0] in ios
      && ps'.t_replicas[idx].ts == Rsl_RecvPerfUpdate(ps.t_replicas[idx].ts, ios[0].r.msg.ts, hstep)
      && (ps.t_replicas[idx].dts <= ios[0].r.msg.ts <= ps.t_replicas[idx].ts + Timeout())
      && ps'.t_replicas[idx].dts == ios[0].r.msg.ts
    else if |ios| > 0 && ios[0].LIoOpTimeoutReceive? then
      && ps'.t_replicas[idx].ts == Rsl_TimeoutPerfUpdate(ps.t_replicas[idx].ts, hstep)
      && ps'.t_replicas[idx].dts == ps.t_replicas[idx].ts + Timeout()
    else
      && ps'.t_replicas[idx].ts == Rsl_NoRecvPerfUpdate(ps.t_replicas[idx].ts, hstep)
      && ps'.t_replicas[idx].dts == ps.t_replicas[idx].dts
      )

      && (forall io :: io in ios && io.LIoOpSend? ==>
          io.s.msg.ts  ==
          (if io.s.src == io.s.dst then ps'.t_replicas[idx].ts + SelfDelivery else ps'.t_replicas[idx].ts + D)
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
