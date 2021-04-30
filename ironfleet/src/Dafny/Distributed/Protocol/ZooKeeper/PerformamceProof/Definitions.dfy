// Must be verified with /arith:2

include "../Types.dfy"
include "../Timestamps/TimestampedType.dfy"

module Zookeeper_Performance_Definitions {
import opened ZooKeeper_Types
import opened ZKTimestamp

ghost const D:Timestamp

ghost const SendFI:Timestamp
ghost const ProcLI:Timestamp
ghost const ProcSyncI:Timestamp
ghost const ProcSync:Timestamp
ghost const ProcSnap:Timestamp

ghost const ProcFI:Timestamp
ghost const ProcEpAck:Timestamp
ghost const PreSync:Timestamp // PrepSync
ghost const Sync:Timestamp  // DoSync
ghost const SyncSnap:Timestamp
ghost const ProcAck:Timestamp

ghost const NoOp:Timestamp := TimeZero()


function StepToTimeDelta(hstep:HostStep) : Timestamp {
    match hstep
        case F(fs) => (
            match fs
                case SendFollowerInfo => SendFI
                case ProcessLeaderInfo => ProcLI
                case ProcessSyncInfo => ProcSyncI
                case ProcessSync => ProcSync
                case ProcessSnap => ProcSnap
                case FStutter => NoOp
        )
        case L(ls) => 
            match ls
                case ProcessFollowerInfo => ProcFI
                case ProcessEpochAck => ProcEpAck
                case PrepSync => PreSync
                case DoSync => Sync
                case DoSyncSNAP => SyncSnap
                case ProcessAck => ProcAck
                case LStutter => NoOp
}

function TLS_NoRecvPerfUpdate(node_pr:Timestamp, hstep:HostStep) : Timestamp
    ensures TLS_NoRecvPerfUpdate(node_pr, hstep) == node_pr + StepToTimeDelta(hstep)
{
    var total_time := TimeAdd2(node_pr, StepToTimeDelta(hstep));
    total_time
}

function TLS_RecvPerfUpdate(node_pr:Timestamp, pkt_pr:Timestamp, hstep:HostStep) : Timestamp
{
    var handlerStartTime := TimeMax(pkt_pr, node_pr);
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
