include "../TimestampedRslSystem.i.dfy"
include "../../CommonProof/Constants.i.dfy"
include "../Common/assumptions.i.dfy"
include "../Common/definitions.i.dfy"

module RslPhase1Proof_i {
import opened TimestampedRslSystem_i
import opened CommonProof__Constants_i
import opened Common_Assumptions
import opened Common_Definitions



/*****************************************************************************************
*                                     Assumptions                                        *
*****************************************************************************************/

predicate P1Assumption(s:TimestampedRslState) {
    && CommonAssumptions(s)
    && (var nextStep := ts.t_environment.nextStep; 
        nextStep.LEnvStepHostIos? ==>
            && (forall io | io in nextStep.ios :: !io.LIoOpTimeoutReceive?)
            && (forall io | io in nextStep.ios && io.LIoOpReceive? :: !io.r.msg.v.RslMessage_Heartbeat? && !io.r.msg.v.RslMessage_Request? && !io.r.msg.v.RslMessage_AppStateSupply?)
    )
    && |ts.t_replicas| > 2 
    && NewLeaderDoesNotReceiveOld2a2b(ts)
    && NewLeaderDoesNotProposeFurtherOps(ts, opn)
    && LeaderAlwaysOne(ts)
    && minD < SelfDelivery < D < 2*minD
    && ProcessPacket > 0
    && SelfDelivery + TimeActionRange(0) < D
}

predicate LeaderAlwaysOne(s:TimestampedRslState)
    requires |s.t_replicas| > 2 
{
    && s.t_replicas[1].v.replica.proposer.election_state.current_view == Ballot(1, 1)
    && (forall idx | 2 <= idx < |s.t_replicas| 
        :: s.t_replicas[idx].v.replica.proposer.current_state == 0)
}

predicate OldLeaderCrashed(s:TimestampedRslState, s':TimestampedRslState)
    requires |s.t_replicas| > 2 
{
    forall idx, tios | TimestampedRslNextOneReplica(ps, ps', idx, ios) :: idx != 0
}


/*****************************************************************************************
*                                     Invariants                                        *
*****************************************************************************************/




predicate AlwaysInvariantP1(ts:TimestampedRslState, opn:OperationNumber)
    requires |ts.t_replicas| > 2
{
    var l := ts.t_replicas[1];
    var r := l.v.replica;
    && 0 <= ts.t_replicas[1].v.nextActionIndex <= 9
    && ServersAreNotClients(ts)
    && LSetOfMessage1b(ts.t_replicas[1].v.replica.proposer.received_1b_packets)
    && AlwaysInvariantP1_RequestSrcAndBatchSize(ts, opn)
    && All1aPackets_BalLeq_Opn(ts, Ballot(1, 1))
    && All1bPackets_BalLeq_Opn(ts, Ballot(1, 1))
    && All2aPackets_BalLeq_Opn(ts, Ballot(1, 0), opn)
    && All2bPackets_BalLeq_Opn(ts, Ballot(1, 0), opn)
    && (!exists pkt :: pkt in ts.t_environment.sentPackets && IsNew2aPacket(pkt, opn))
    && (!exists pkt :: pkt in ts.t_environment.sentPackets && IsNew2bPacket(pkt, opn))
    && (!exists pkt :: pkt in ts.t_environment.sentPackets && IsNewReplyPacket(ts, pkt))
    && ts.t_replicas[1].v.replica.proposer.request_queue == []
    && (forall pkt | pkt in ts.undeliveredPackets :: pkt in ts.t_environment.sentPackets)
    && ts.t_replicas[1].v.replica.learner.unexecuted_learner_state == map[]
    && ts.t_replicas[1].v.replica.executor.next_op_to_execute == OutstandingOpUnknown()
    && ts.t_replicas[1].v.replica.proposer.current_view_suspectors == {}

    && r.proposer.current_state == 2
    && r.proposer.election_state.current_view == Ballot(1, 1)
    && r.proposer.max_ballot_i_sent_1a == Ballot(1, 1)
    && opn == r.proposer.next_operation_number_to_propose

    // Learner and Executor states, maintain for phase 2
    && opn == r.executor.ops_complete
    && r.learner.unexecuted_learner_state == map[]
    && r.executor.next_op_to_execute == OutstandingOpUnknown()
    && BalLt(r.learner.max_ballot_seen, Ballot(1, 1))
}


/* Current identical as in P2 */
predicate AlwaysInvariantP1_RequestSrcAndBatchSize(ts:TimestampedRslState, opn:OperationNumber)
    requires |ts.t_replicas| > 2
    requires LSetOfMessage1b(ts.t_replicas[1].v.replica.proposer.received_1b_packets)
{   
    && All2aPackets_BalLeq_Opn(ts, Ballot(1, 0), opn)
    && All2bPackets_BalLeq_Opn(ts, Ballot(1, 0), opn)
    && (forall pkt | IsUndelivered_2aPkt(ts, pkt) :: RequestBatchSrcInClientIds(ts, pkt.msg.v.val_2a))
    && (forall pkt | IsUndelivered_2bPkt(ts, pkt) :: RequestBatchSrcInClientIds(ts, pkt.msg.v.val_2b))
    && (forall pkt | pkt in ts.t_environment.sentPackets && pkt.msg.v.RslMessage_2a? :: |pkt.msg.v.val_2a| > 0)
    && (forall pkt | pkt in ts.t_environment.sentPackets && pkt.msg.v.RslMessage_2b? :: |pkt.msg.v.val_2b| > 0)
    && (forall pkt, opn'| pkt in ts.t_environment.sentPackets && pkt.msg.v.RslMessage_1b? && opn' in pkt.msg.v.votes
        ::  && |pkt.msg.v.votes[opn'].max_val| > 0
            && RequestBatchSrcInClientIds(ts, pkt.msg.v.votes[opn'].max_val))
    && (forall v | LValIsHighestNumberedProposal(v, ts.t_replicas[1].v.replica.proposer.received_1b_packets, opn)
        :: |v| > 0)
    // Leader's received_1b_packets
    && (forall pkt, op | pkt in ts.t_replicas[1].v.replica.proposer.received_1b_packets && pkt.msg.RslMessage_1b? && op in pkt.msg.votes
        :: RequestBatchSrcInClientIds(ts, pkt.msg.votes[op].max_val))
    // Acceptors
    && (forall idx, opn'| 0 <= idx < |ts.t_replicas| && opn' in ts.t_replicas[idx].v.replica.acceptor.votes
        ::  && |ts.t_replicas[idx].v.replica.acceptor.votes[opn'].max_val| > 0   
            &&  RequestBatchSrcInClientIds(ts, ts.t_replicas[idx].v.replica.acceptor.votes[opn'].max_val)
    )
}


// Things that are true before 1a packets are sent out by the leader
predicate Before_1a_Sent_Invariant(ts:TimestampedRslState, opn:OperationNumber) 
    requires |ts.t_replicas| > 2
    requires RslConsistency(ts)
{
    // Timestamp 
    && TimeLe(l.ts, NewLeaderP1_InitTS)  // TODO
    // TODO also need timestamp on 1a packets, since these are already sent

    // Network
    var l := ts.t_replicas[1];
    var r := l.v.replica;
    && All1aPackets_BalLeq_Opn(ts, Ballot(1, 1))
    && All1bPackets_BalLeq_Opn(ts, Ballot(1, 0))
    && (forall pkt :: pkt in ts.t_environment.sentPackets && pkt.msg.v.RslMessage_1a? :: IsNew1aPacket(pkt))
    && (!exists pkt :: pkt in ts.t_environment.sentPackets && IsNew1bPacket(pkt))

    // Proposer
    && |r.proposer.received_1b_packets| < LMinQuorumSize(s.constants.all.config)

    // Acceptors
    // TODO: Need quorum of acceptors accepted only opn. Or somehow assume this?
}



}
