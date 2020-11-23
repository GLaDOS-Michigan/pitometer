include "../Collections/Maps2.s.dfy"
include "../Logic/Temporal/Temporal.s.dfy"

module EnvironmentTCP_s {
import opened Collections__Maps2_s
import opened Temporal__Temporal_s


/*****************************************************************************************
*                                        Types                                           *
******************************************************************************************/   

datatype LPacket<IdType, MessageType(==)> = LPacket(dst:IdType, src:IdType, msg:MessageType)

datatype LIoOp<IdType, MessageType(==)> = LIoOpSend(s:LPacket<IdType, MessageType>)
                                        | LIoOpReceive(r:LPacket<IdType, MessageType>)
                                        | LIoOpTimeoutReceive()
                                        | LIoOpReadClock(t:int)

datatype LEnvStep<IdType, MessageType(==)> = LEnvStepHostIos(actor:IdType, ios:seq<LIoOp<IdType, MessageType>>)
                                           | LEnvStepDeliverPacket(p:LPacket<IdType, MessageType>)
                                           | LEnvStepAdvanceTime()
                                           | LEnvStepStutter()

datatype LHostInfo<IdType, MessageType(==)> = LHostInfo(queue:seq<LPacket<IdType, MessageType>>)

// Channel[..i] are the delievered prefix
datatype HostChannel<IdType, MessageType> = HostChannel(index:int, channel:seq<LPacket<IdType, MessageType>>)

datatype LEnvironment<IdType, MessageType(==)> = LEnvironment(time:int,
                                                              channels:map<IdType, HostChannel>, 
                                                              hostInfo:map<IdType, LHostInfo<IdType, MessageType>>,
                                                              nextStep:LEnvStep<IdType, MessageType>)

/*****************************************************************************************
*                                     Validity Checks                                    *
******************************************************************************************/                                        
predicate IsValidLIoOp<IdType, MessageType>(io:LIoOp, actor:IdType, e:LEnvironment<IdType, MessageType>)
{
    match io
        case LIoOpSend(s) => s.src == actor
        case LIoOpReceive(r) => r.dst == actor
        case LIoOpTimeoutReceive => true
        case LIoOpReadClock(t) => true
}

predicate LIoOpOrderingOKForAction<IdType, MessageType>(
    io1:LIoOp<IdType, MessageType>,
    io2:LIoOp<IdType, MessageType>
    )
{
    io1.LIoOpReceive? || io2.LIoOpSend?
}

predicate LIoOpSeqCompatibleWithReduction<IdType, MessageType>(
    ios:seq<LIoOp<IdType, MessageType>>
    )
{
    forall i {:trigger ios[i], ios[i+1]} :: 0 <= i < |ios| - 1 ==> LIoOpOrderingOKForAction(ios[i], ios[i+1])
}

/* Returns true iff s1 is a prefix of s2 */
predicate IsPrefix<T>(s1: seq<T>, s2: seq<T>) {
    && |s1| <= |s2|
    && s2[0..|s1|] == s2
}

predicate IsValidReceiveSeq<IdType, MessageType>(rcvSeq:seq<LPacket<IdType, MessageType>>, hc:HostChannel) 
{
    && 0 <= hc.index <= |hc.channel|
    && IsPrefix(rcvSeq, hc.channel[hc.index..])
}

predicate IsValidLEnvStep<IdType, MessageType>(e:LEnvironment<IdType, MessageType>, step:LEnvStep)
{
    match step
        case LEnvStepHostIos(actor, ios) =>    
            var rcvMap := IosToRcvMap(ios);
            && (forall io :: io in ios ==> IsValidLIoOp(io, actor, e))
            && (forall receiver | receiver in rcvMap :: receiver in e.channels ==> IsValidReceiveSeq(rcvMap[receiver], e.channels[receiver]))
            && LIoOpSeqCompatibleWithReduction(ios)
        case LEnvStepDeliverPacket(p) => p.dst in e.channels && p in e.channels[p.dst].channel
        case LEnvStepAdvanceTime => true
        case LEnvStepStutter => true
}


/*****************************************************************************************
*                                       Actions                                          *
******************************************************************************************/ 


/* Transform a seq of IoOps to a map that maps each dst to a sequence of send ops to that dst*/
function IosToSendMap<IdType, MessageType>(ios:seq<LIoOp<IdType, MessageType>>) : map<IdType, seq<LPacket<IdType, MessageType>>> {
    if |ios| == 0 then map[]
    else 
        var head := ios[0];
        var sub_map := IosToSendMap(ios[1..]);
        if head.LIoOpSend? then (
            if head.s.dst in sub_map then sub_map[head.s.dst := sub_map[head.s.dst] + [head.s]]
            else sub_map[head.s.dst := [head.s]]
        ) else sub_map
}

/* Transform a seq of IoOps to a map that maps each dst to a sequence of receive ops to that dst*/
function IosToRcvMap<IdType, MessageType>(ios:seq<LIoOp<IdType, MessageType>>) : map<IdType, seq<LPacket<IdType, MessageType>>> {
    if |ios| == 0 then map[]
    else 
        var head := ios[0];
        var sub_map := IosToRcvMap(ios[1..]);
        if head.LIoOpReceive? then (
            if head.r.dst in sub_map then sub_map[head.r.dst := sub_map[head.r.dst] + [head.r]]
            else sub_map[head.r.dst := [head.r]]
        ) else sub_map
}


predicate LEnvironment_Init<IdType, MessageType>(
    e:LEnvironment<IdType, MessageType>,
    hosts: seq<IdType>)
{
    && (forall h | h in hosts :: h in e.channels && e.channels[h] == HostChannel(0, []))
    && e.time >= 0
}


/* Maps channels to its new state after a sending and receiving a set of ios in sendMap
* and rcvMap */
function PerformIos<IdType, MessageType>(
    channels:map<IdType, HostChannel>, 
    sendMap: map<IdType, seq<LPacket<IdType, MessageType>>>,
    rcvMap: map<IdType, seq<LPacket<IdType, MessageType>>>) :
    map<IdType, HostChannel>
{
    map h | h in channels :: (
        if h in rcvMap then (
            if h in sendMap 
                then HostChannel(channels[h].index + |rcvMap[h]|, channels[h].channel + sendMap[h])
                else HostChannel(channels[h].index + |rcvMap[h]|, channels[h].channel)
        ) else (
            if h in sendMap 
                then HostChannel(channels[h].index, channels[h].channel + sendMap[h])
                else channels[h]
        )
    )
}


predicate LEnvironment_PerformIos<IdType, MessageType>(
    e:LEnvironment<IdType, MessageType>,
    e':LEnvironment<IdType, MessageType>,
    actor:IdType,
    ios:seq<LIoOp<IdType, MessageType>>
    )
{
    var sendMap := IosToSendMap(ios);
    var rcvMap := IosToRcvMap(ios);
    && e'.channels == PerformIos(e.channels, sendMap, rcvMap)
    && e'.time == e.time
}

predicate LEnvironment_AdvanceTime<IdType, MessageType>(
    e:LEnvironment<IdType, MessageType>,
    e':LEnvironment<IdType, MessageType>
    )
{
       e'.time > e.time
    // UNCHANGED
    && e'.channels == e.channels
}

predicate LEnvironment_Stutter<IdType, MessageType>(
    e:LEnvironment<IdType, MessageType>,
    e':LEnvironment<IdType, MessageType>
    )
{
       e'.time == e.time
    && e'.channels == e.channels
}

predicate LEnvironment_Next<IdType, MessageType>(
    e:LEnvironment<IdType, MessageType>,
    e':LEnvironment<IdType, MessageType>
    )
{
       IsValidLEnvStep(e, e.nextStep)
    && match e.nextStep
           case LEnvStepHostIos(actor, ios) => LEnvironment_PerformIos(e, e', actor, ios)
           case LEnvStepDeliverPacket(p) => LEnvironment_Stutter(e, e') // this is only relevant for synchrony
           case LEnvStepAdvanceTime => LEnvironment_AdvanceTime(e, e')
           case LEnvStepStutter => LEnvironment_Stutter(e, e')
}


/*****************************************************************************************
*                                     Temporals                                          *
******************************************************************************************/ 


// function{:opaque} EnvironmentNextTemporal<IdType, MessageType, NodeStepType>(b:Behavior<LEnvironment<IdType, MessageType, NodeStepType>>):temporal
//     requires imaptotal(b);
//     ensures forall i {:trigger sat(i, EnvironmentNextTemporal(b))} ::
//                 sat(i, EnvironmentNextTemporal(b)) <==> LEnvironment_Next(b[i], b[i+1]);
// {
//     stepmap(imap i :: LEnvironment_Next(b[i], b[i+1]))
// }

// predicate LEnvironment_BehaviorSatisfiesSpec<IdType, MessageType, NodeStepType>(
//     b:Behavior<LEnvironment<IdType, MessageType, NodeStepType>>
//     )
// {
//        imaptotal(b)
//     && LEnvironment_Init(b[0])
//     && sat(0, always(EnvironmentNextTemporal(b)))
// }

} 
