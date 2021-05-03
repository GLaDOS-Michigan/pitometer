include "../../Common/Native/NativeTypes.s.dfy"
include "../../Common/Framework/Environment.s.dfy"
include "Types.dfy"

module FailureDetector_Environment {
import opened Native__NativeTypes_s
import opened Environment_s
import opened FailureDetector_Types


datatype EndPoint = EndPoint(addr:seq<byte>, port:uint16)

datatype Config = Conf(nodeEp:EndPoint, detectorEp:EndPoint, clientEp:EndPoint)

datatype FDMessage = 
    | Heartbeat(time:int)
    | Alert(time:int)

type FDEnvironment = LEnvironment<EndPoint, FDMessage, FDStep>
type FDPacket = LPacket<EndPoint, FDMessage>
type FDIo = LIoOp<EndPoint, FDMessage>


predicate ConfigIsUnique(c:Config) {
    && c.nodeEp != c.detectorEp && c.nodeEp != c.clientEp
    && c.detectorEp != c.clientEp
}



// Borrowed from RSL replica
predicate SpontaneousIos(ios:seq<FDIo>, clocks:int)
    requires 0<=clocks<=1;
{
       clocks <= |ios|
    && (forall i | 0 <= i < clocks :: ios[i].LIoOpReadClock?)
    && (forall i | clocks <=i<|ios| :: ios[i].LIoOpSend?)
}

// Borrowed from RSL replica
function SpontaneousClock(ios:seq<FDIo>) : ClockReading
{
    if SpontaneousIos(ios, 1)
        then ClockReading(ios[0].t)
        else ClockReading(0)  // nonsense to avoid putting a precondition on this function
}

// Borrowed from RSL replica
function {:opaque} ExtractSentPacketsFromIos(ios:seq<FDIo>) : seq<FDPacket>
    ensures forall p :: p in ExtractSentPacketsFromIos(ios) <==> LIoOpSend(p) in ios;
{
    if |ios| == 0 then
        []
    else if ios[0].LIoOpSend? then
        [ios[0].s] + ExtractSentPacketsFromIos(ios[1..])
    else
        ExtractSentPacketsFromIos(ios[1..])
}

}
