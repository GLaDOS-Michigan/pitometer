include "../TimestampedRslSystem.i.dfy"
include "../../CommonProof/Constants.i.dfy"
include "../Common/assumptions.i.dfy"
include "../Common/definitions.i.dfy"
include "../Phase2_PostFail/Phase2Proof.i.dfy"

module RslPhase1Proof_i {
import opened TimestampedRslSystem_i
import opened CommonProof__Constants_i
import opened Common_Assumptions
import opened Common_Definitions
import P2 = RslPhase2Proof_PostFail_i



/*****************************************************************************************
*                                     Assumptions                                        *
*****************************************************************************************/

predicate P1Assumption(ts:TimestampedRslState, opn:OperationNumber) {
    && CommonAssumptions(ts)
    && (var nextStep := ts.t_environment.nextStep; 
        nextStep.LEnvStepHostIos? ==>
            && (forall io | io in nextStep.ios :: !io.LIoOpTimeoutReceive?)
            && (forall io | io in nextStep.ios && io.LIoOpReceive? :: !io.r.msg.v.RslMessage_Heartbeat? && !io.r.msg.v.RslMessage_Request? && !io.r.msg.v.RslMessage_AppStateSupply?)
    )
    && |ts.t_replicas| > 2 
    && NewLeaderDoesNotReceiveOld2a2b(ts)
    && NewLeaderDoesNotProposeFurtherOps(ts, opn)
    && LeaderAlwaysOne(ts)
    && QuorumOf2bPacketsImpliesValue(ts, opn)
    && minD < SelfDelivery < D < 2*minD
    && ProcessPacket > 0
    && SelfDelivery + TimeActionRange(0) < D
}

predicate LeaderAlwaysOne(ts:TimestampedRslState)
    requires |ts.t_replicas| > 2 
{
    && ts.t_replicas[1].v.replica.proposer.election_state.current_view == Ballot(1, 1)
    && (forall idx | 2 <= idx < |ts.t_replicas| 
        :: ts.t_replicas[idx].v.replica.proposer.current_state == 0)
}

predicate OldLeaderCrashed(ts:TimestampedRslState, ts':TimestampedRslState)
    requires |ts.t_replicas| > 2 
{
    forall idx, tios | TimestampedRslNextOneReplica(ts, ts', idx, tios) :: idx != 0
}


predicate QuorumOf2bPacketsImpliesValue(ts:TimestampedRslState, opn:OperationNumber)
    requires |ts.t_replicas| > 2 
{
    var r := ts.t_replicas[1].v.replica;
    LProposerCanNominateUsingOperationNumber(r.proposer, r.acceptor.log_truncation_point, r.proposer.next_operation_number_to_propose)
    ==>
    && P2.AlwaysInvariantP2_RequestSrcAndBatchSize(ts, opn)
}


/*****************************************************************************************
*                                     Invariants                                        *
*****************************************************************************************/

predicate Phase1Invariant(ts:TimestampedRslState, opn:OperationNumber) 
    requires |ts.t_replicas| > 2
    requires RslConsistency(ts)
{
    && AlwaysInvariantP1(ts, opn)
    && PacketsBallotInvariant(ts)
    && Invariant(ts, opn)
}


predicate AlwaysInvariantP1(ts:TimestampedRslState, opn:OperationNumber)
    requires |ts.t_replicas| > 2
    requires RslConsistency(ts)
{
    var l := ts.t_replicas[1];
    var r := l.v.replica;
    // && AlwaysInvariantP1_RequestSrcAndBatchSize(ts, opn)
    && 0 <= ts.t_replicas[1].v.nextActionIndex <= 9
    && ServersAreNotClients(ts)
    && LSetOfMessage1b(ts.t_replicas[1].v.replica.proposer.received_1b_packets)
    && All2aPackets_BalLeq_Opn(ts, Ballot(1, 0), opn)
    && All2bPackets_BalLeq_Opn(ts, Ballot(1, 0), opn)
    && (!exists pkt :: pkt in ts.t_environment.sentPackets && IsNew2aPacket(pkt, opn))
    && (!exists pkt :: pkt in ts.t_environment.sentPackets && IsNew2bPacket(pkt, opn))
    && (!exists pkt :: pkt in ts.t_environment.sentPackets && IsNewReplyPacket(ts, pkt))
    && (forall pkt | pkt in ts.undeliveredPackets :: pkt in ts.t_environment.sentPackets)

    // Proposer stuff
    && r.proposer.request_queue == []
    && r.proposer.election_state.current_view_suspectors == {}
    && r.proposer.current_state == 2
    && r.proposer.election_state.current_view == Ballot(1, 1)
    && r.proposer.max_ballot_i_sent_1a == Ballot(1, 1)
    && r.proposer.next_operation_number_to_propose == opn

    // Learner and Executor states, maintain for phase 2
    && opn == r.executor.ops_complete
    && r.learner.unexecuted_learner_state == map[]
    && r.executor.next_op_to_execute == OutstandingOpUnknown()
    && BalLt(r.learner.max_ballot_seen, Ballot(1, 1))
}


predicate PacketsBallotInvariant(ts:TimestampedRslState) {
    forall pkt | pkt in ts.undeliveredPackets :: 
    if pkt.msg.v.RslMessage_1a? then 
        pkt.msg.v.bal_1a == Ballot(1, 1)
    else if pkt.msg.v.RslMessage_1b? then
        pkt.msg.v.bal_1b == Ballot(1, 1)
}


/* Current identical as in P2 */
// predicate AlwaysInvariantP1_RequestSrcAndBatchSize(ts:TimestampedRslState, opn:OperationNumber)
//     requires |ts.t_replicas| > 2
//     requires LSetOfMessage1b(ts.t_replicas[1].v.replica.proposer.received_1b_packets)
// {   
//     && All2aPackets_BalLeq_Opn(ts, Ballot(1, 0), opn)
//     && All2bPackets_BalLeq_Opn(ts, Ballot(1, 0), opn)
//     && (forall pkt | IsUndelivered_2aPkt(ts, pkt) :: RequestBatchSrcInClientIds(ts, pkt.msg.v.val_2a))
//     && (forall pkt | IsUndelivered_2bPkt(ts, pkt) :: RequestBatchSrcInClientIds(ts, pkt.msg.v.val_2b))
//     && (forall pkt | pkt in ts.t_environment.sentPackets && pkt.msg.v.RslMessage_2a? :: |pkt.msg.v.val_2a| > 0)
//     && (forall pkt | pkt in ts.t_environment.sentPackets && pkt.msg.v.RslMessage_2b? :: |pkt.msg.v.val_2b| > 0)
//     && (forall pkt, opn'| pkt in ts.t_environment.sentPackets && pkt.msg.v.RslMessage_1b? && opn' in pkt.msg.v.votes
//         ::  && |pkt.msg.v.votes[opn'].max_val| > 0
//             && RequestBatchSrcInClientIds(ts, pkt.msg.v.votes[opn'].max_val))
//     && (forall v | LValIsHighestNumberedProposal(v, ts.t_replicas[1].v.replica.proposer.received_1b_packets, opn)
//         :: |v| > 0)
//     // Leader's received_1b_packets
//     && (forall pkt, op | pkt in ts.t_replicas[1].v.replica.proposer.received_1b_packets && pkt.msg.RslMessage_1b? && op in pkt.msg.votes
//         :: RequestBatchSrcInClientIds(ts, pkt.msg.votes[op].max_val))
//     // Acceptors
//     && (forall idx, opn'| 0 <= idx < |ts.t_replicas| && opn' in ts.t_replicas[idx].v.replica.acceptor.votes
//         ::  && |ts.t_replicas[idx].v.replica.acceptor.votes[opn'].max_val| > 0   
//             &&  RequestBatchSrcInClientIds(ts, ts.t_replicas[idx].v.replica.acceptor.votes[opn'].max_val)
//     )
// }


// Things that are true before 1a packets are sent out by the leader
predicate Invariant(ts:TimestampedRslState, opn:OperationNumber) 
    requires |ts.t_replicas| > 2
    requires RslConsistency(ts)
{
    var l := ts.t_replicas[1];
    var r := l.v.replica;

    // Timestamps
    && TimeLe(l.ts, TimeBoundPhase1Leader(l.v.nextActionIndex)) 

    && (forall pkt | pkt in ts.undeliveredPackets && IsNew1aPacket(pkt)
        :: TimeLe(pkt.msg.ts, TimeBound1aDelivery()))

    && (forall pkt | pkt in ts.undeliveredPackets && IsNew1bPacket(pkt)
        :: TimeLe(pkt.msg.ts, TimeBound1bDeliveryPost()))

    // Proposer


}



}
