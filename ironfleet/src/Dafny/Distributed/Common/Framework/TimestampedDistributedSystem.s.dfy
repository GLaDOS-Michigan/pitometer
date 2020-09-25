include "Environment.s.dfy"
include "Host.s.dfy"
include "DistributedSystem.s.dfy"
include "../Collections/Seqs.i.dfy"
include "Timestamp.s.dfy"

abstract module TimestampedDistributedSystem_s {
    import opened Environment_s
    import opened T_s : Timestamp_s
    import opened Collections__Seqs_i

datatype TimestampedType<Type> = TimestampedType(v:Type, ts:Timestamp)

type TimestampedLPacket<IdType, MessageType(==)> = LPacket<IdType, TimestampedType<MessageType>>

type TimestampedLIoOp<IdType, MessageType(==)> = LIoOp<IdType, TimestampedType<MessageType>>

type TimestampedLEnvStep<IdType, MessageType(==), NodeStepType> = LEnvStep<IdType,TimestampedType<MessageType>, NodeStepType>

type TimestampedLHostInfo<IdType, MessageType(==)> = LHostInfo<IdType,TimestampedType<MessageType>>

type TimestampedLEnvironment<IdType, MessageType(==), NodeStepType> = LEnvironment<IdType, TimestampedType<MessageType>, NodeStepType>

type TimestampedHostState = TimestampedType<HostState>

datatype TimestampedDS_State = TimestampedDS_State(
    config:ConcreteConfiguration,
    t_environment:TimestampedLEnvironment<EndPoint, seq<byte>, HostStep>,
    t_servers:map<EndPoint,TimestampedHostState>,
    clients:set<EndPoint>,
    num_steps:NumSteps
)
        
function UntagTimestampedType<T>(t_t: TimestampedType<T>) : T {
    t_t.v
}

function UntagLPacket<I,M>(pkt: TimestampedLPacket<I,M>) : LPacket<I,M> {
    LPacket(pkt.dst, pkt.src, pkt.msg.v)
}

function UntagSentPkts<I,M>(t_sentPkts : set<TimestampedLPacket<I,M>>) : set<LPacket<I,M>> {
    set pkt | pkt in t_sentPkts :: UntagLPacket(pkt)
}

function UntagLEnvStep(t_nextStep : TimestampedLEnvStep) : LEnvStep {
    match t_nextStep
    {
        case LEnvStepHostIos(actor, ios, nodeStep) => LEnvStepHostIos(actor, UntagLIoOpSeq(ios), nodeStep)
        case LEnvStepDeliverPacket(p) => LEnvStepDeliverPacket(UntagLPacket(p))
        case LEnvStepAdvanceTime() => LEnvStepAdvanceTime
        case LEnvStepStutter() => LEnvStepStutter
    }
}

function UntagHostInfo(t_hi: TimestampedLHostInfo) : LHostInfo {
    LHostInfo(MapSeqToSeq(t_hi.queue, UntagLPacket))
}

function UntagHostInfoMap<I,M>(t_hostInfo:map<I, LHostInfo<I,TimestampedType<M>>>) : map<I,LHostInfo<I,M>> {
    map id | id in t_hostInfo :: UntagHostInfo(t_hostInfo[id])
}

function UntagLEnvironment<I,M,S>(t_env: TimestampedLEnvironment<I,M,S>) : LEnvironment<I,M,S>{
    LEnvironment(t_env.time,
        UntagSentPkts(t_env.sentPackets),
        UntagHostInfoMap(t_env.hostInfo),
        UntagLEnvStep(t_env.nextStep))
}

function UntagServers(t_servers: map<EndPoint, TimestampedHostState>) : map<EndPoint, HostState> {
    map id | id in t_servers :: t_servers[id].v
}

function UntagDS_State(tds:TimestampedDS_State) : DS_State {
    DS_State(tds.config,
        UntagLEnvironment(tds.t_environment),
        UntagServers(tds.t_servers),
        tds.clients)
}

function UntagLIoOp(t_io : TimestampedLIoOp) : LIoOp {
    match t_io
    {
        case LIoOpSend(s) => LIoOpSend(UntagLPacket(s))
        case LIoOpReceive(r) => LIoOpReceive(UntagLPacket(r))
        case LIoOpTimeoutReceive() => LIoOpTimeoutReceive()
        case LIoOpReadClock(t) => LIoOpReadClock(t)
    }
}

function {:opaque} UntagLIoOpSeq(t_ios: seq<TimestampedLIoOp>) : seq<LIoOp>
    ensures |UntagLIoOpSeq(t_ios)| == |t_ios|
    ensures forall i :: 0 <= i < |UntagLIoOpSeq(t_ios)| ==> UntagLIoOpSeq(t_ios)[i] == UntagLIoOp(t_ios[i])
{
    MapSeqToSeq(t_ios, UntagLIoOp)
}

predicate TDS_NumStepsValid(tds:TimestampedDS_State){
    ValidNumSteps(tds.num_steps)
}

predicate TDS_Init(tds: TimestampedDS_State, config:ConcreteConfiguration)
    reads *
{
    && tds.num_steps == NumStepsInit()
    && TDS_NumStepsValid(tds)
    && DS_Init(UntagDS_State(tds), config)
            && forall id :: id in tds.t_servers ==> tds.t_servers[id].ts == TimeZero()
}

predicate TDS_NextOneServer(tds: TimestampedDS_State, tds': TimestampedDS_State, id:EndPoint, ios:seq<TimestampedLIoOp<EndPoint,seq<byte>>>, hstep:HostStep)
    requires TDS_NumStepsValid(tds)
    requires id in tds.t_servers;
    reads *
{
    && DS_NextOneServer(UntagDS_State(tds), UntagDS_State(tds'), id, UntagLIoOpSeq(ios))
    && if |ios| > 0 && ios[0].LIoOpReceive? then
        (tds'.t_servers[id].ts, tds'.num_steps) == RecvTimestampUpdate(tds.t_servers[id].ts, ios[0].r.msg.ts, hstep, tds.num_steps)
    else
        && (tds'.t_servers[id].ts, tds'.num_steps) == NoRecvTimestampUpdate(tds.t_servers[id].ts, hstep, tds.num_steps)
        && (forall t_io :: t_io in ios && t_io.LIoOpSend? ==> t_io.s.msg.ts == tds'.t_servers[id].ts)
}

predicate TDS_Next(tds:TimestampedDS_State, tds': TimestampedDS_State)
    reads *
{
    && TDS_NumStepsValid(tds)
    && DS_Next(UntagDS_State(tds), UntagDS_State(tds'))
    && LEnvironment_Next(tds.t_environment, tds'.t_environment)
    && if tds.t_environment.nextStep.LEnvStepHostIos? && tds.t_environment.nextStep.actor in tds.t_servers 
        then TDS_NextOneServer(tds, tds', tds.t_environment.nextStep.actor, tds.t_environment.nextStep.ios, tds.t_environment.nextStep.nodeStep)
        else tds'.t_servers == tds.t_servers
}

predicate IsValidBehavior(tdb:seq<TimestampedDS_State>, config:ConcreteConfiguration)
    reads *
{
    && |tdb| > 0
    && TDS_Init(tdb[0], config)
    && (forall i {:trigger TDS_Next(tdb[i], tdb[i+1])} :: 0 <= i < |tdb| - 1 ==> TDS_Next(tdb[i], tdb[i+1]))
}
}
