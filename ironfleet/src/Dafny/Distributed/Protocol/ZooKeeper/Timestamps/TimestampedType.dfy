include "../../../Common/Framework/EnvironmentTCP.s.dfy"
include "Timestamp.dfy"

module ZKTimestamp refines Timestamp_s{
import opened EnvironmentTCP_s

datatype TimestampedType<Type> = TimestampedType(v:Type, ts:Timestamp, dts:Timestamp)

type TimestampedLPacket<IdType, MessageType(==)> = LPacket<IdType, TimestampedType<MessageType>>

type TimestampedLIoOp<IdType, MessageType(==)> = LIoOp<IdType, TimestampedType<MessageType>>

type TimestampedLEnvStep<IdType, MessageType(==)> = LEnvStep<IdType,TimestampedType<MessageType>>

type TimestampedLHostInfo<IdType, MessageType(==)> = LHostInfo<IdType,TimestampedType<MessageType>>

type TimestampedLEnvironment<IdType, MessageType(==)> = LEnvironment<IdType, TimestampedType<MessageType>>

function UntagTimestampedType<T>(t_t: TimestampedType<T>) : T 
    ensures UntagTimestampedType(t_t) == t_t.v;
{
    t_t.v
}

function UntagLPacket<I,M>(pkt: TimestampedLPacket<I,M>) : LPacket<I,M>
    ensures UntagLPacket(pkt) == LPacket(pkt.dst, pkt.src, pkt.sender_index, pkt.msg.v);
{
    LPacket(pkt.dst, pkt.src, pkt.sender_index, pkt.msg.v)
}

function UntagLPacketSeq<I,M>(pkts: seq<TimestampedLPacket<I,M>>) : seq<LPacket<I,M>>
    ensures |UntagLPacketSeq(pkts)| == |pkts|
    ensures forall i | 0 <= i < |pkts| :: UntagLPacketSeq(pkts)[i] == UntagLPacket(pkts[i])
{
    if |pkts| == 0 then [] else [UntagLPacket(pkts[0])] + UntagLPacketSeq(pkts[1..])
}

function UntagSentPkts<I,M>(t_sentPkts : set<TimestampedLPacket<I,M>>) : set<LPacket<I,M>> 
    ensures UntagSentPkts(t_sentPkts) == (set pkt | pkt in t_sentPkts :: UntagLPacket(pkt));
{
    set pkt | pkt in t_sentPkts :: UntagLPacket(pkt)
}

function UntagChannels<I,M>(t_channels : map<I, HostChannel<I, TimestampedType<M>>>) : map<I, HostChannel<I,M>>
    ensures t_channels.Keys == UntagChannels(t_channels).Keys
    ensures forall id | id in t_channels :: UntagChannels(t_channels)[id].index == t_channels[id].index && UntagChannels(t_channels)[id].channel == UntagLPacketSeq(t_channels[id].channel)
    // ensures UntagSentPkts(t_sentPkts) == (set pkt | pkt in t_sentPkts :: UntagLPacket(pkt));
{
    map id | id in t_channels :: HostChannel(t_channels[id].index, UntagLPacketSeq(t_channels[id].channel))
}

function UntagLEnvStep<I,M>(t_nextStep : TimestampedLEnvStep<I,M>) : LEnvStep<I,M> 
    ensures UntagLEnvStep(t_nextStep) ==
    match t_nextStep
    case LEnvStepHostIos(actor, ios) => LEnvStepHostIos(actor, UntagLIoOpSeq(ios))
    case LEnvStepDeliverPacket(p) => LEnvStepDeliverPacket(UntagLPacket(p))
    case LEnvStepAdvanceTime() => LEnvStepAdvanceTime
    case LEnvStepStutter() => LEnvStepStutter
{
    match t_nextStep
    case LEnvStepHostIos(actor, ios) => LEnvStepHostIos(actor, UntagLIoOpSeq(ios))
    case LEnvStepDeliverPacket(p) => LEnvStepDeliverPacket(UntagLPacket(p))
    case LEnvStepAdvanceTime() => LEnvStepAdvanceTime
    case LEnvStepStutter() => LEnvStepStutter
}

function UntagHostInfo(t_hi: TimestampedLHostInfo) : LHostInfo 
    ensures UntagHostInfo(t_hi) == LHostInfo(MapSeqToSeq(t_hi.queue, UntagLPacket));
{
    LHostInfo(MapSeqToSeq(t_hi.queue, UntagLPacket))
}

function UntagHostInfoMap<I,M>(t_hostInfo:map<I, LHostInfo<I,TimestampedType<M>>>) : map<I,LHostInfo<I,M>> 
    ensures UntagHostInfoMap(t_hostInfo) == (map id | id in t_hostInfo :: UntagHostInfo(t_hostInfo[id]));
{
    map id | id in t_hostInfo :: UntagHostInfo(t_hostInfo[id])
}

function UntagLEnvironment<I,M>(t_env: TimestampedLEnvironment<I,M>) : LEnvironment<I,M>{
    LEnvironment(t_env.time,
        t_env.config,
        UntagChannels(t_env.channels),
        UntagSentPkts(t_env.sentPackets),
        UntagHostInfoMap(t_env.hostInfo),
        UntagLEnvStep(t_env.nextStep))
}


function UntagLIoOp(t_io : TimestampedLIoOp) : LIoOp 
    ensures UntagLIoOp(t_io) == 
    match t_io
    case LIoOpSend(s) => LIoOpSend(UntagLPacket(s))
    case LIoOpReceive(r) => LIoOpReceive(UntagLPacket(r))
    case LIoOpTimeoutReceive() => LIoOpTimeoutReceive()
    case LIoOpReadClock(t) => LIoOpReadClock(t)
{
    match t_io
    case LIoOpSend(s) => LIoOpSend(UntagLPacket(s))
    case LIoOpReceive(r) => LIoOpReceive(UntagLPacket(r))
    case LIoOpTimeoutReceive() => LIoOpTimeoutReceive()
    case LIoOpReadClock(t) => LIoOpReadClock(t)
}

function UntagLIoOpSeq(t_ios: seq<TimestampedLIoOp>) : seq<LIoOp>
    ensures |UntagLIoOpSeq(t_ios)| == |t_ios|
    ensures forall i :: 0 <= i < |UntagLIoOpSeq(t_ios)| ==> UntagLIoOpSeq(t_ios)[i] == UntagLIoOp(t_ios[i])
{
    MapSeqToSeq(t_ios, UntagLIoOp)
}
}
