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

datatype LEnvironment<IdType, MessageType(==)> = LEnvironment(time:int,
                                                              sentPackets:map<IdType, seq<LPacket<IdType, MessageType>>>,  // effectively the receive queue of each host
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

predicate IsValidLEnvStep<IdType, MessageType>(e:LEnvironment<IdType, MessageType>, step:LEnvStep)
{
    match step
        case LEnvStepHostIos(actor, ios) =>    (forall io :: io in ios ==> IsValidLIoOp(io, actor, e))
                                           && LIoOpSeqCompatibleWithReduction(ios)
        case LEnvStepDeliverPacket(p) => p.dst in e.sentPackets && p in e.sentPackets[p.dst]
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

predicate LEnvironment_Init<IdType, MessageType>(
    e:LEnvironment<IdType, MessageType>,
    hosts: set<IdType>)
{
    && (forall h | h in hosts :: h in e.sentPackets && e.sentPackets[h] == [])
    && e.time >= 0
}

predicate LEnvironment_PerformIos<IdType, MessageType>(
    e:LEnvironment<IdType, MessageType>,
    e':LEnvironment<IdType, MessageType>,
    actor:IdType,
    ios:seq<LIoOp<IdType, MessageType>>
    )
{
    var send_map := IosToSendMap(ios);
    && e'.sentPackets.Keys == e.sentPackets.Keys
    && (forall h | h in e'.sentPackets :: if h in send_map then e'.sentPackets[h] ==  e.sentPackets[h] + send_map[h] else e'.sentPackets[h] ==  e.sentPackets[h])
    && (forall io :: io in ios && io.LIoOpReceive? ==> actor in e.sentPackets && io.r in e.sentPackets[actor])
    && e'.time == e.time
}

predicate LEnvironment_AdvanceTime<IdType, MessageType>(
    e:LEnvironment<IdType, MessageType>,
    e':LEnvironment<IdType, MessageType>
    )
{
       e'.time > e.time
    // UNCHANGED
    && e'.sentPackets == e.sentPackets
}

predicate LEnvironment_Stutter<IdType, MessageType>(
    e:LEnvironment<IdType, MessageType>,
    e':LEnvironment<IdType, MessageType>
    )
{
       e'.time == e.time
    && e'.sentPackets == e.sentPackets
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
