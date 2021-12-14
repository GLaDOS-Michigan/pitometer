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
    && NodeZeroCrashed(ts)
    && (var nextStep := ts.t_environment.nextStep; 
        nextStep.LEnvStepHostIos? ==>
            && (forall io | io in nextStep.ios :: !io.LIoOpTimeoutReceive?)
            && (forall io | io in nextStep.ios && io.LIoOpReceive? :: 
                && !io.r.msg.v.RslMessage_Heartbeat? 
                && !io.r.msg.v.RslMessage_Invalid? 
                && !io.r.msg.v.RslMessage_Request? 
                && !io.r.msg.v.RslMessage_Reply? 
                && !io.r.msg.v.RslMessage_AppStateRequest?
                && !io.r.msg.v.RslMessage_AppStateSupply?
                && !io.r.msg.v.RslMessage_StartingPhase2?
            )
    )
    && |ts.t_replicas| > 2 
    && NewLeaderDoesNotReceiveOld2a2b(ts)
    && NewLeaderDoesNotProposeFurtherOps(ts, opn)
    && LeaderAlwaysOne(ts)
    && NonLeaderViews(ts)
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

predicate NonLeaderViews(ts:TimestampedRslState) 
    requires |ts.t_replicas| > 2
{
    forall idx | 2 <= idx < |ts.t_replicas|
    :: || ts.t_replicas[idx].v.replica.proposer.election_state.current_view == Ballot(1, 0)
       || ts.t_replicas[idx].v.replica.proposer.election_state.current_view == Ballot(1, 1)
}

predicate NodeZeroCrashed(ts:TimestampedRslState)
    requires |ts.t_replicas| > 2 
    requires RslConsistency(ts)
{
    ts.t_environment.nextStep.LEnvStepHostIos?
    ==> ts.t_environment.nextStep.actor != ts.constants.config.replica_ids[0]
}


predicate QuorumOf2bPacketsImpliesValue(ts:TimestampedRslState, opn:OperationNumber)
    requires |ts.t_replicas| > 2 
{
    var r := ts.t_replicas[1].v.replica;
    var s := r.proposer;
    && |r.proposer.received_1b_packets| >= LMinQuorumSize(r.proposer.constants.all.config)
    && LSetOfMessage1bAboutBallot(r.proposer.received_1b_packets, r.proposer.max_ballot_i_sent_1a)
    ==>
    && r.acceptor.log_truncation_point == opn 
    && P2.AlwaysInvariantP2_RequestSrcAndBatchSize(ts, opn)
    && LIsAfterLogTruncationPoint(opn, s.received_1b_packets)
    && opn < UpperBoundedAddition(r.acceptor.log_truncation_point, s.constants.all.params.max_log_length, s.constants.all.params.max_integer_val)
    && opn >= 0
    && LtUpperBound(opn, s.constants.all.params.max_integer_val)
    && !LAllAcceptorsHadNoProposal(r.proposer.received_1b_packets, r.proposer.next_operation_number_to_propose)
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
    && PerfInvariant(ts, opn)
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
    && r.proposer.max_ballot_i_sent_1a == Ballot(1, 1)
    && r.proposer.next_operation_number_to_propose == opn
    && (var es := r.proposer.election_state;
         || es.current_view_suspectors == {}
         || es.current_view_suspectors == {es.constants.my_index}
    )

    // Learner and Executor states, maintain for phase 2
    && opn == r.executor.ops_complete
    && r.learner.unexecuted_learner_state == map[]
    && r.executor.next_op_to_execute == OutstandingOpUnknown()
    && BalLt(r.learner.max_ballot_seen, Ballot(1, 1))
}


predicate PacketsBallotInvariant(ts:TimestampedRslState) {
    forall pkt | pkt in ts.undeliveredPackets :: ExistingPacketsBallot(pkt)
}

predicate ExistingPacketsBallot(pkt:TimestampedLPacket<EndPoint, RslMessage>) {
    match pkt.msg.v {
        // All 1a and 1b packets have Ballot (1, 1)
        case RslMessage_1a(bal_1a)              => bal_1a == Ballot(1, 1) 
        case RslMessage_1b(bal_1b,_,_)          => bal_1b == Ballot(1, 1) 
        // All 2a and 2b messages have Ballot(1, 0) or (1, 1)
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
predicate PerfInvariant(ts:TimestampedRslState, opn:OperationNumber) 
    requires |ts.t_replicas| > 2
    requires RslConsistency(ts)
{
    var l := ts.t_replicas[1];
    var r := l.v.replica;
    && 0 <= ts.t_replicas[1].v.nextActionIndex <= 9

    // Timestamps
    && TimeLe(l.ts, TimeBoundPhase1LeaderPost(l.v.nextActionIndex)) 

    && (forall pkt | pkt in ts.undeliveredPackets && IsNew1aPacket(pkt)
        :: TimeLe(pkt.msg.ts, TimeBound1aDeliveryPost()))

    && (forall pkt | pkt in ts.undeliveredPackets && IsNew1bPacket(pkt)
        :: TimeLe(pkt.msg.ts, TimeBound1bDeliveryPost()))

    // Proposer
    && r.proposer.current_state == 1
}



}
