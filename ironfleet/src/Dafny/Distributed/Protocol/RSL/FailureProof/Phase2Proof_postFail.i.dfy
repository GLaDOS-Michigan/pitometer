include "../DistributedSystem.i.dfy"

include "../../../Common/Collections/Maps2.i.dfy"
include "../Constants.i.dfy"
include "../Environment.i.dfy"
include "../Replica.i.dfy"

include "TimestampedRslSystem.i.dfy"

include "../CommonProof/Constants.i.dfy"

module RslPhase2Proof_postFail_i {
import opened TimestampedRslSystem_i
import opened CommonProof__Constants_i

/*****************************************************************************************
*                                      Guarantees                                        *
*****************************************************************************************/

/* Main performance guarantee for phase 2 post-failure */
predicate PerformanceGuarantee(s:TimestampedRslState){
    && PerformanceGuarantee_Response(s)
    && PerformanceGuarantee_2b(s)
    && PerformanceGuarantee_2a(s)
}

predicate PerformanceGuarantee_Response(s:TimestampedRslState) {
    // 0 < |s.constants.config.replica_ids|
    //     && (forall pkt :: 
    //             && pkt in s.undeliveredPackets
    //             && pkt.msg.v.RslMessage_Reply?
    //             && pkt.src == s.constants.config.replica_ids[0]
    //             ==>
    //             TimeLe(pkt.msg.ts, TimeBoundReply(req_time, s.constants))
    //     )
    //TODO
    true
}

predicate PerformanceGuarantee_2b(b:TimestampedRslState) {
    // TODO
    true
}

predicate PerformanceGuarantee_2a(s:TimestampedRslState) {
    forall pkt {:trigger pkt.msg.v.RslMessage_2a?} | 
        && pkt in s.undeliveredPackets 
        && pkt.msg.v.RslMessage_2a?
        && pkt.msg.v.bal_2a == Ballot(1, 1)
    :: TimeLe(pkt.msg.ts, TimeBound2aDeliveryPost())
}

function TimeBound2aDeliveryPost() : Timestamp {
    NewLeaderInitTS + TimeActionRange(0) + TimeActionRange(4) + D
}

/*****************************************************************************************
*                                     Assumptions                                        *
*****************************************************************************************/

/* Conjunction of all assumptions */
predicate RslAssumption(s:TimestampedRslState)
{
    && RelationsAssumptions()
    && RslConsistency(s)
    && NoPacketDuplication(s)
    && |s.t_replicas| > 2  // For failure to be meaningful
    && s.constants.params.max_batch_size == 1

    && (var nextStep := s.t_environment.nextStep; nextStep.LEnvStepHostIos? ==>
    (forall io :: io in nextStep.ios && io.LIoOpReceive? ==> !io.r.msg.v.RslMessage_Heartbeat?))
    && LeaderAlwaysOne(s)
    && minD < SelfDelivery < D < 2*minD
    && ProcessPacket > 0
    && SelfDelivery + TimeActionRange(0) < D
    && NoExternalPackets(s)
    && NoExternalSteps(s)
    && BoundedQueueingAssumption(s)
}

predicate RelationsAssumptions() {
    && TimeLe(req_time, NewLeaderInitTS)
}

predicate LeaderAlwaysOne(s:TimestampedRslState) 
    requires |s.t_replicas| > 2
{
    && (forall idx | 0 <= idx < |s.t_replicas| && idx != 1
        :: s.t_replicas[idx].v.replica.proposer.current_state == 0)
    && s.t_replicas[1].v.replica.proposer.current_state == 2
    && s.t_replicas[1].v.replica.proposer.election_state.current_view == Ballot(1, 1)
}

predicate NoExternalPackets(s:TimestampedRslState) {
    forall pkt | pkt in s.undeliveredPackets ::
        && pkt in s.t_environment.sentPackets
        && pkt.src in s.constants.config.replica_ids
        && pkt.dst in s.constants.config.replica_ids
}

predicate NoExternalSteps(s:TimestampedRslState) {
    !(exists eid, ios :: s.t_environment.nextStep == LEnvStepHostIos(eid, ios, ExternalStep()))
}

// Bounded queueing assumption for simplified RSL
predicate BoundedQueueingAssumption(s:TimestampedRslState) 
    requires RslConsistency(s)
{
    forall idx, ios :: (
        && 0 <= idx < |s.constants.config.replica_ids|
        && s.t_environment.nextStep == LEnvStepHostIos(s.constants.config.replica_ids[idx], ios, RslStep(s.t_replicas[idx].v.nextActionIndex))
        ==>
        (forall io | io in s.t_environment.nextStep.ios && io.LIoOpReceive? ::
            // this means that max(replica.ts, msg.ts) <= msg.ts + MaxQueueTime
            s.t_replicas[idx].ts <= io.r.msg.ts + MaxQueueTime
        )
    )
}

predicate RslConsistency(s:TimestampedRslState)
{
    ConstantsAllConsistentInv(UntimestampRslState(s))
        && WellFormedLConfiguration(s.constants.config)
}

predicate ServersAreNotClients(s:TimestampedRslState)
{
    forall id :: id in s.constants.config.clientIds && id in s.constants.config.replica_ids
        ==> false
}

predicate RequestBatchSrcInClientIds(s:TimestampedRslState, v:RequestBatch)
{
    forall r :: r in v ==> r.client in s.constants.config.clientIds
}


/*****************************************************************************************
*                                    Boundary State                                      *
*****************************************************************************************/

/* Timestamp of initial client request */
ghost const req_time:Timestamp
/* Initial timestamp of replica 1 */
ghost const NewLeaderInitTS:Timestamp

//  Invariant of starting state of Phase2 proof
predicate BoundaryCondition(s:TimestampedRslState) 
    requires |s.t_replicas| > 2
{
    && BoundaryCondition_ExistingPacketsBallot(s)
    && BoundaryCondition_NewLeader(s)
}


predicate BoundaryCondition_NewLeader(s:TimestampedRslState) {
    // TODO
    true
}


predicate BoundaryCondition_ExistingPacketsBallot(s:TimestampedRslState) {
    forall pkt | pkt in s.undeliveredPackets :: Boundary_ExistingPacketsBallot(pkt)
}
predicate Boundary_ExistingPacketsBallot(pkt:TimestampedLPacket<EndPoint, RslMessage>) {
    match pkt.msg.v {
        // All 1a and 1b packets have Ballot (1, 1)
        case RslMessage_1a(bal_1a)              => bal_1a == Ballot(1, 1) 
        case RslMessage_1b(bal_1b,_,_)          => bal_1b == Ballot(1, 1) 
        // All 2a and 2b messages have Ballot(1, 0)
        case RslMessage_2a(bal_2a,_,_)          => bal_2a == Ballot(1, 0)
        case RslMessage_2b(bal_2b,_,_)          => bal_2b == Ballot(1, 0)

        // Cases where I don't care
        case RslMessage_Heartbeat(_,_,_)        => true
        case RslMessage_Invalid                 => true
        case RslMessage_Request(_,_)            => true
        case RslMessage_Reply(_,_)              => true
        case RslMessage_AppStateRequest(_,_)    => true
        case RslMessage_AppStateSupply(_,_,_,_) => true
        case RslMessage_StartingPhase2(_,_)     => true
    }
}





/*****************************************************************************************
*                                     Invariants                                        *
*****************************************************************************************/


// Main invariant 
predicate RslPerfInvariant(s:TimestampedRslState, req_time:Timestamp, opn:OperationNumber) 
    requires |s.t_replicas| > 2
{
    && RslConsistency(s)
    && PerformanceGuarantee(s)
    && PacketsBallotInvariant(s)
}

predicate PacketsBallotInvariant(s:TimestampedRslState) {
    forall pkt | pkt in s.undeliveredPackets :: ExistingPacketsBallot(pkt)
}

predicate ExistingPacketsBallot(pkt:TimestampedLPacket<EndPoint, RslMessage>) {
    match pkt.msg.v {
        // All 1a and 1b packets have Ballot (1, 1)
        case RslMessage_1a(bal_1a)              => bal_1a == Ballot(1, 1) 
        case RslMessage_1b(bal_1b,_,_)          => bal_1b == Ballot(1, 1) 
        // All 2a and 2b messages have Ballot(1, 0) or (1, 1)
        case RslMessage_2a(bal_2a,_,_)          => bal_2a == Ballot(1, 0) || bal_2a == Ballot(1, 1)
        case RslMessage_2b(bal_2b,_,_)          => bal_2b == Ballot(1, 0) || bal_2b == Ballot(1, 1)

        // Cases where I don't care
        case RslMessage_Heartbeat(_,_,_)        => true
        case RslMessage_Invalid                 => true
        case RslMessage_Request(_,_)            => true
        case RslMessage_Reply(_,_)              => true
        case RslMessage_AppStateRequest(_,_)    => true
        case RslMessage_AppStateSupply(_,_,_,_) => true
        case RslMessage_StartingPhase2(_,_)     => true
    }
}
}
