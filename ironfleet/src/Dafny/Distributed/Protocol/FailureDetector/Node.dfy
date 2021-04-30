include "Types.dfy"
include "FDEnvironment.dfy"


module FailureDetector_Node {
import opened FailureDetector_Types
import opened FailureDetector_Environment


datatype NodeState = RUNNING | FAILED

datatype Node = Node(
    state:NodeState,
    heartbeatInterval:int,
    nextHeartbeatTime:int,
    config:Config
)



/*****************************************************************************************
*                                      Actions                                           *
******************************************************************************************/ 

predicate NodeInit(s:Node, heartbeatInterval:int, config:Config) {
    && s.state == RUNNING
    && s.heartbeatInterval == heartbeatInterval
    && s.nextHeartbeatTime == s.heartbeatInterval  // assuming global clock starts at 0
    && s.config == config
}

predicate NodeNext(s:Node, s':Node, ios:seq<FDIo>) {
    match s.state
        case RUNNING => 
            || NodeNext_ReadClockAndMaybeSendHeartbeat(s, s', ios)
            || NodeNext_Fail(s, s', ios)
        case FAILED => NodeNext_Stutter(s, s', ios)
}

predicate NodeNext_Stutter(s:Node, s':Node, ios:seq<FDIo>) 
    requires s.state == FAILED
{
    s' == s && |ios| == 0
}

/* Enter failed state */
predicate NodeNext_Fail(s:Node, s':Node, ios:seq<FDIo>) 
    requires s.state == RUNNING
{
    && s' == s.(state := FAILED)
    && |ios| == 0
}


/* Check clock and maybe send heartbeat */
predicate NodeNext_ReadClockAndMaybeSendHeartbeat(s:Node, s':Node, ios:seq<FDIo>) 
    requires s.state == RUNNING
{
    var clock := SpontaneousClock(ios);
    var sentPackets := ExtractSentPacketsFromIos(ios);
    && SpontaneousIos(ios, 1)
    && if clock.t < s.nextHeartbeatTime then
        s' == s && sentPackets == []
    else
        && s' == s.(nextHeartbeatTime := clock.t + s.heartbeatInterval)
        && |sentPackets| == 1
        && sentPackets[0] == LPacket(s.config.detectorEp, s.config.nodeEp, Heartbeat(clock.t))
}
}
