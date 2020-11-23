// Must be verified with /arith:2

include "../Types.dfy"
include "../Timestamps/TimestampedType.dfy"

module PerformanceProof__Definitions_i {
import opened ZooKeeper_Types
import opened ZKTimestamp

ghost const D:Timestamp

ghost const FSFI:Timestamp
ghost const FANE:Timestamp
ghost const FPS:Timestamp
ghost const FPSNAP:Timestamp

ghost const LGE:Timestamp
ghost const LWFE:Timestamp
ghost const LPS:Timestamp
ghost const LDS:Timestamp
ghost const LDSSNAP:Timestamp
ghost const LPA:Timestamp


function StepToTimeDelta(hstep:HostStep) : Timestamp
{
    // if hstep == GrantStep then G else A
    match hstep
        case F(fs) => (
            match fs
                case SendFollowerInfo => FSFI
                case AcceptNewEpoch => FANE
                case ProcessSync => FPS
                case ProcessSnap => FPSNAP
        )
        case L(ls) => 
            match ls
                case GetEpoch => LGE
                case WaitForEpoch => LWFE
                case PrepSync => LPS
                case DoSync => LDS
                case DoSyncSNAP => LDSSNAP
                case ProcessAck => LPA
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

// function PerfBoundLockHeld(epoch: int) : Timestamp
//     requires 0 < epoch
//     ensures PerfBoundLockHeld(epoch) >= 0
// {
//     (epoch - 1) * G + (epoch - 1) * A + (epoch - 1) * D
// }

// function PerfBoundLockInNetwork(epoch: int) : Timestamp
//     requires 0 < epoch
//     ensures PerfBoundLockInNetwork(epoch) >= 0
// {
//     if epoch == 1 then    // note that a valid Transfer packet never has epoch 0
//         0
//     else
//         (epoch - 1) * G + (epoch - 2) * A + (epoch - 2) * D
// }

predicate TimeEq(p1:Timestamp, p2:Timestamp)
{
    p1 == p2
}

predicate TimeLe(p1:Timestamp, p2:Timestamp)
{
    p1 <= p2
}

// lemma Grant_j_helper_specific(epoch:int)
//     requires epoch > 0
//     ensures TimeEq(PerfBoundLockInNetwork(epoch + 1), TimeAdd2(PerfBoundLockHeld(epoch), G));
// {
// }

// lemma Grant_j_helper()
//     ensures forall epoch :: epoch > 0 ==> TimeEq(PerfBoundLockInNetwork(epoch + 1), TimeAdd2(PerfBoundLockHeld(epoch), G));
// {
// }

// lemma Accept_j_helper()
//     ensures forall epoch :: epoch > 1 ==> TimeEq(PerfBoundLockHeld(epoch), TLS_RecvPerfUpdate(TimeVoid(), PerfBoundLockInNetwork(epoch), AcceptStep))
// {
// }
}
