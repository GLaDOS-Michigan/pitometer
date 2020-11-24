// Must be verified with /arith:2

include "../Types.dfy"
include "../Timestamps/TimestampedType.dfy"

module Zookeeper_Performance_Definitions {
import opened ZooKeeper_Types
import opened ZKTimestamp

ghost const D:Timestamp

ghost const SendFI:Timestamp
ghost const ProcLI:Timestamp
ghost const ProcSync:Timestamp
ghost const ProcSnap:Timestamp

ghost const ProcFI:Timestamp
ghost const ProcEpAck:Timestamp
ghost const PreSync:Timestamp
ghost const Sync:Timestamp
ghost const SyncSnap:Timestamp
ghost const ProcAck:Timestamp


function StepToTimeDelta(hstep:HostStep) : Timestamp {
    match hstep
        case F(fs) => (
            match fs
                case SendFollowerInfo => SendFI
                case ProcessLeaderInfo => ProcLI
                case ProcessSync => ProcSync
                case ProcessSnap => ProcSnap
        )
        case L(ls) => 
            match ls
                case ProcessFollowerInfo => ProcFI
                case ProcessEpochAck => ProcEpAck
                case PrepSync => PreSync
                case DoSync => Sync
                case DoSyncSNAP => SyncSnap
                case ProcessAck => ProcAck
}

function TLS_NoRecvPerfUpdate(node_pr:Timestamp, hstep:HostStep) : Timestamp
{
    var total_time := TimeAdd2(node_pr, StepToTimeDelta(hstep));
    total_time
}

function TLS_RecvPerfUpdate(node_pr:Timestamp, pkt_pr:Timestamp, hstep:HostStep) : Timestamp
{
    var deliveryTime := TimeAdd2(pkt_pr, D);    // add D to packet sent time
    var handlerStartTime := TimeMax(deliveryTime, node_pr);
    var total_time := TimeAdd2(handlerStartTime, StepToTimeDelta(hstep));
    total_time
}


predicate TimeEq(p1:Timestamp, p2:Timestamp)
{
    p1 == p2
}

predicate TimeLe(p1:Timestamp, p2:Timestamp)
{
    p1 <= p2
}
}
