include "Types.dfy"
include "FDEnvironment.dfy"


module FailureDetector_Detector {
import opened FailureDetector_Types
import opened FailureDetector_Environment


datatype Detector = Detector(
    timeoutInterval:int,
    lastHeartbeatTime:int,
    nextActionIndex:int,
    config:Config
)

function DetectorNumActions() : int {
    2
}

/*****************************************************************************************
*                                      Actions                                           *
******************************************************************************************/ 

predicate DetectorInit(s:Detector, timeoutInterval:int, config:Config) {
    && s.timeoutInterval == timeoutInterval
    && s.lastHeartbeatTime == 0  // assuming global clock starts at 0
    && s.config == config
}

predicate DetectorNext(s:Detector, s':Detector, ios:seq<FDIo>) {
    && s'.timeoutInterval == s.timeoutInterval
    && s'.config == s.config
    && s'.nextActionIndex == (s.nextActionIndex + 1) % DetectorNumActions()
    &&  if s.nextActionIndex == 0 then
            DetectorNext_TryReceiveHeartbeat(s, s', ios)
        else
            DetectorNext_ReadClockAndMaybeTimeout(s, s', ios)
}

predicate DetectorNext_TryReceiveHeartbeat(s:Detector, s':Detector, ios:seq<FDIo>) {
    && |ios| >= 1
    &&  if ios[0].LIoOpTimeoutReceive? then
           && s'.lastHeartbeatTime == s.lastHeartbeatTime
           && |ios| == 1
        else (
            && ios[0].LIoOpReceive?
            && if ios[0].r.msg.Heartbeat? then
                    // Process the heartbeat
                    DetectorNext_ReadClockAndProcessHeartbeat(s, s', ios)
              else
                    // This can never happen
                    && |ios| == 1
                    && s'.lastHeartbeatTime == s.lastHeartbeatTime
        )
}

predicate DetectorNext_ReadClockAndProcessHeartbeat(s:Detector, s':Detector, ios:seq<FDIo>)
    requires |ios| >= 1;
    requires ios[0].LIoOpReceive?;
    requires ios[0].r.msg.Heartbeat?;
{
    && |ios| == 2   // first item is the received heartbeat, second item is reading the clock
    && ios[1].LIoOpReadClock?
    && s'.lastHeartbeatTime == ios[1].t
}

predicate DetectorNext_ReadClockAndMaybeTimeout(s:Detector, s':Detector, ios:seq<FDIo>){ 
    var clock := SpontaneousClock(ios);
    var sentPackets := ExtractSentPacketsFromIos(ios);
    && s'.lastHeartbeatTime == s.lastHeartbeatTime
    && SpontaneousIos(ios, 1)
    &&  if clock.t < s.lastHeartbeatTime + s.timeoutInterval then
            sentPackets == []
        else
            && |sentPackets| == 1
            && sentPackets[0] == LPacket(s.config.clientEp, s.config.detectorEp, Alert(clock.t))
}
}
