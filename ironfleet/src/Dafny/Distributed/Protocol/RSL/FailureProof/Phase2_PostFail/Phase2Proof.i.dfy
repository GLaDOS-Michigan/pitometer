include "../TimestampedRslSystem.i.dfy"
include "../../CommonProof/Constants.i.dfy"
include "../Common/assumptions.i.dfy"

module RslPhase2Proof_PostFail_i {
import opened TimestampedRslSystem_i
import opened CommonProof__Constants_i
import opened Common_Assumptions

/*****************************************************************************************
*                                      Guarantees                                        *
*****************************************************************************************/

/* Main performance guarantee for phase 2 post-failure */
predicate PerformanceGuarantee(ts:TimestampedRslState, opn:OperationNumber)
    requires |ts.t_replicas| > 2
    requires RslConsistency(ts)
{
    && PerformanceGuarantee_Response(ts)
    && PerformanceGuarantee_2b(ts, opn)
    && PerformanceGuarantee_2a(ts, opn)
}

predicate PerformanceGuarantee_Response(ts:TimestampedRslState) 
    requires |ts.t_replicas| > 2
    requires RslConsistency(ts)
{
    forall pkt | 
        && pkt in ts.undeliveredPackets 
        && IsNewReplyPacket(ts, pkt)
    :: TimeLe(pkt.msg.ts, TimeBoundReplyFinal())
}

predicate PerformanceGuarantee_2b(ts:TimestampedRslState, opn:OperationNumber) {
    forall pkt {:trigger pkt.msg.v.RslMessage_2b?} | 
        && pkt in ts.undeliveredPackets 
        && IsNew2bPacket(pkt, opn)
    :: TimeLe(pkt.msg.ts, TimeBound2bDeliveryPost())
}

predicate PerformanceGuarantee_2a(ts:TimestampedRslState, opn:OperationNumber) {
    forall pkt {:trigger pkt.msg.v.RslMessage_2a?} | 
        && pkt in ts.undeliveredPackets 
        && IsNew2aPacket(pkt, opn)
    :: TimeLe(pkt.msg.ts, TimeBound2aDeliveryPost())
}

function TimeBound2aDeliveryPost() : Timestamp {
    NewLeaderInitTS + MbeP2a + D
}

function TimeBound2bDeliveryPost() : Timestamp {
    TimeBound2aDeliveryPost() + ProcessPacket + MaxQueueTime + D
}

function TimeBoundPhase2LeaderPost(nextActionIndex:int) : Timestamp
    requires 0 <= nextActionIndex < 10
{
  TimeBound2bDeliveryPost() + MaxQueueTime + TimeActionRange(nextActionIndex)
}

function TimeBoundReplyFinal() : Timestamp {
    TimeBoundPhase2LeaderPost(7) + D
}

/*****************************************************************************************
*                                     Assumptions                                        *
*****************************************************************************************/

/* Conjunction of all assumptions */
predicate RslAssumption(ts:TimestampedRslState, opn:OperationNumber)
{
    && CommonAssumptions(ts)
    && (var nextStep := ts.t_environment.nextStep; 
        && nextStep.LEnvStepHostIos? ==>
            && (forall io | io in nextStep.ios :: !io.LIoOpTimeoutReceive?)
            && (forall io | io in nextStep.ios && io.LIoOpReceive? :: !io.r.msg.v.RslMessage_Heartbeat? && !io.r.msg.v.RslMessage_Request? && !io.r.msg.v.RslMessage_AppStateSupply?)
    )
    && NewLeaderDoesNotReceiveOld2a2b(ts)
    && NewLeaderDoesNotProposeFurtherOps(ts, opn)
    && LeaderAlwaysOne(ts)
    && minD < SelfDelivery < D < 2*minD
    && ProcessPacket > 0
    && SelfDelivery + TimeActionRange(0) < D
}

predicate NewLeaderDoesNotProposeFurtherOps(ts:TimestampedRslState, opn:OperationNumber) 
    requires |ts.t_replicas| > 2 
{
    ts.t_replicas[1].v.replica.proposer.constants.all.params.max_integer_val == UpperBoundFinite(opn + 1)
}

/* Assume that the leader does not receive leftover 2b packets from before leader election */
predicate NewLeaderDoesNotReceiveOld2a2b(ts:TimestampedRslState) 
    requires |ts.t_replicas| > 2 
{
    var nextStep := ts.t_environment.nextStep;
    nextStep.LEnvStepHostIos? ==>
    && (forall io | io in nextStep.ios && io.LIoOpReceive? && io.r.msg.v.RslMessage_2b? :: io.r.msg.v.bal_2b != Ballot(1, 0))
    && (forall io | io in nextStep.ios && io.LIoOpReceive? && io.r.msg.v.RslMessage_2a? :: io.r.msg.v.bal_2a != Ballot(1, 0))
}

predicate LeaderAlwaysOne(ts:TimestampedRslState) 
    requires |ts.t_replicas| > 2
{
    && (forall idx | 0 <= idx < |ts.t_replicas| && idx != 1
        :: ts.t_replicas[idx].v.replica.proposer.current_state == 0)
    && ts.t_replicas[1].v.replica.proposer.current_state == 2
    && ts.t_replicas[1].v.replica.proposer.election_state.current_view == Ballot(1, 1)
}


/*****************************************************************************************
*                                    Boundary State                                      *
*****************************************************************************************/

/* Timestamp of initial client request */
ghost const req_time:Timestamp
/* Initial timestamp of replica 1 */
ghost const NewLeaderInitTS:Timestamp

predicate LeaderSet1bContainsRequest(ts:TimestampedRslState) 
    requires |ts.t_replicas| > 2
{
    var r := ts.t_replicas[1].v.replica;
    && LProposerCanNominateUsingOperationNumber(r.proposer, r.acceptor.log_truncation_point, r.proposer.next_operation_number_to_propose)
    && !LAllAcceptorsHadNoProposal(r.proposer.received_1b_packets, r.proposer.next_operation_number_to_propose)
}


/*****************************************************************************************
*                                     Invariants                                        *
*****************************************************************************************/


// Main invariant 
predicate RslPerfInvariant(ts:TimestampedRslState, opn:OperationNumber) 
    requires |ts.t_replicas| > 2
{
    && RslConsistency(ts)
    && AlwaysInvariant(ts, opn)
    && PerformanceGuarantee(ts, opn)
    && PacketsBallotInvariant(ts)
    && (|| Before_2a_Sent_Invariant(ts, opn)
        || Before_2b_Sent_Invariant(ts, opn)
        || After_2b_Sent_Invariant(ts, opn)
    )
}


predicate ServersAreNotClients(ts:TimestampedRslState)
{
  forall id :: id in ts.constants.config.clientIds && id in ts.constants.config.replica_ids
    ==> false
}

predicate RequestBatchSrcInClientIds(ts:TimestampedRslState, v:RequestBatch)
{
  forall r :: r in v ==> r.client in ts.constants.config.clientIds
}

predicate AlwaysInvariant(ts:TimestampedRslState, opn:OperationNumber)
    requires |ts.t_replicas| > 2
{
    && ServersAreNotClients(ts)
    && ts.t_replicas[1].v.replica.proposer.request_queue == []
    && (forall pkt | pkt in ts.undeliveredPackets :: pkt in ts.t_environment.sentPackets)
    && (forall pkt | pkt in ts.undeliveredPackets && pkt.msg.v.RslMessage_2a? :: RequestBatchSrcInClientIds(ts, pkt.msg.v.val_2a))
    && (forall pkt | pkt in ts.undeliveredPackets && pkt.msg.v.RslMessage_2b? :: RequestBatchSrcInClientIds(ts, pkt.msg.v.val_2b))

    && (forall pkt, opn'| pkt in ts.t_environment.sentPackets && pkt.msg.v.RslMessage_1b? && opn' in pkt.msg.v.votes
        ::  && |pkt.msg.v.votes[opn'].max_val| > 0
            && RequestBatchSrcInClientIds(ts, pkt.msg.v.votes[opn'].max_val))
    && (forall pkt | pkt in ts.t_environment.sentPackets && pkt.msg.v.RslMessage_2a? :: |pkt.msg.v.val_2a| > 0)
    && (forall pkt | pkt in ts.t_environment.sentPackets && pkt.msg.v.RslMessage_2b? :: |pkt.msg.v.val_2b| > 0)
    && (forall idx, opn'| 0 <= idx < |ts.t_replicas| && opn' in ts.t_replicas[idx].v.replica.acceptor.votes
        ::  && |ts.t_replicas[idx].v.replica.acceptor.votes[opn'].max_val| > 0   
            &&  RequestBatchSrcInClientIds(ts, ts.t_replicas[idx].v.replica.acceptor.votes[opn'].max_val)
    )

    && LSetOfMessage1b(ts.t_replicas[1].v.replica.proposer.received_1b_packets)
    && (forall pkt, op | pkt in ts.t_replicas[1].v.replica.proposer.received_1b_packets && pkt.msg.RslMessage_1b? && op in pkt.msg.votes
        :: RequestBatchSrcInClientIds(ts, pkt.msg.votes[op].max_val))
    
    && (forall v | LValIsHighestNumberedProposal(v, ts.t_replicas[1].v.replica.proposer.received_1b_packets, opn)
        :: |v| > 0)
    
    && var uls := ts.t_replicas[1].v.replica.learner.unexecuted_learner_state;
    && (forall opn | opn in uls :: RequestBatchSrcInClientIds(ts, uls[opn].candidate_learned_value))
    && (ts.t_replicas[1].v.replica.executor.next_op_to_execute.OutstandingOpKnown?
        ==> RequestBatchSrcInClientIds(ts, ts.t_replicas[1].v.replica.executor.next_op_to_execute.v))
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


predicate All2aPackets_BalLeq_Opn(ts:TimestampedRslState, ballot:Ballot, opn:OperationNumber) {
    forall pkt | pkt in ts.t_environment.sentPackets && pkt.msg.v.RslMessage_2a?
    ::  && BalLeq(pkt.msg.v.bal_2a, ballot)
        && pkt.msg.v.opn_2a == opn
}

predicate All2bPackets_BalLeq_Opn(ts:TimestampedRslState, ballot:Ballot, opn:OperationNumber) {
    forall pkt | pkt in ts.t_environment.sentPackets && pkt.msg.v.RslMessage_2b?
    ::  && BalLeq(pkt.msg.v.bal_2b, ballot)
        && pkt.msg.v.opn_2b == opn
}

// Things that are true before 2a packets are sent out by the leader
predicate Before_2a_Sent_Invariant(ts:TimestampedRslState, opn:OperationNumber) 
    requires |ts.t_replicas| > 2
    requires RslConsistency(ts)
{
    var l := ts.t_replicas[1];
    var r := l.v.replica;
    && All2aPackets_BalLeq_Opn(ts, Ballot(1, 0), opn)
    && All2bPackets_BalLeq_Opn(ts, Ballot(1, 0), opn)
    && (!exists pkt :: pkt in ts.t_environment.sentPackets && IsNew2aPacket(pkt, opn))
    && (!exists pkt :: pkt in ts.t_environment.sentPackets && IsNew2bPacket(pkt, opn))
    && (!exists pkt :: pkt in ts.t_environment.sentPackets && IsNewReplyPacket(ts, pkt))
    && PerformanceGuarantee_Response(ts)
    && TimeLe(l.ts, NewLeaderInitTS)     // leader timestamp
    && l.v.nextActionIndex == 3          // leader action index is 3
    && LeaderSet1bContainsRequest(ts)
    && ts.t_replicas[1].v.nextActionIndex == 3
    && r.proposer.current_state == 2
    && r.proposer.election_state.current_view == Ballot(1, 1)
    && r.proposer.max_ballot_i_sent_1a == Ballot(1, 1)
    && opn == r.proposer.next_operation_number_to_propose
    
    // Learner and Executor states
    && opn == r.executor.ops_complete
    && r.learner.unexecuted_learner_state == map[]
    && r.executor.next_op_to_execute == OutstandingOpUnknown()
    && BalLt(r.learner.max_ballot_seen, Ballot(1, 1))
}

// Things that are true after 2a packets are sent out by the leader
predicate Before_2b_Sent_Invariant(ts:TimestampedRslState, opn:OperationNumber) 
    requires |ts.t_replicas| > 2
    requires RslConsistency(ts)
{
    var l := ts.t_replicas[1];
    var r := l.v.replica;
    && All2aPackets_BalLeq_Opn(ts, Ballot(1, 1), opn)
    && All2bPackets_BalLeq_Opn(ts, Ballot(1, 0), opn)
    && (exists pkt :: pkt in ts.t_environment.sentPackets && IsNew2aPacket(pkt, opn))
    && (!exists pkt :: pkt in ts.t_environment.sentPackets && IsNew2bPacket(pkt, opn))
    && (!exists pkt :: pkt in ts.t_environment.sentPackets && IsNewReplyPacket(ts, pkt))
    && PerformanceGuarantee_2a(ts, opn)
    && PerformanceGuarantee_Response(ts)
    && 0 <= ts.t_replicas[1].v.nextActionIndex <= 9
    && r.proposer.current_state == 2
    && r.proposer.next_operation_number_to_propose > opn

    // Learner and Executor states
    && r.learner.unexecuted_learner_state == map[]
    && (forall opn' | opn' in r.learner.unexecuted_learner_state :: opn' == opn)
    && opn == r.executor.ops_complete
    && r.executor.next_op_to_execute == OutstandingOpUnknown()
    && BalLt(r.learner.max_ballot_seen, Ballot(1, 1))
}

predicate After_2b_Sent_Invariant(ts:TimestampedRslState, opn:OperationNumber) 
    requires |ts.t_replicas| > 2
    requires RslConsistency(ts)
{
    var l := ts.t_replicas[1];
    var r := l.v.replica;
    && All2aPackets_BalLeq_Opn(ts, Ballot(1, 1), opn)
    && All2bPackets_BalLeq_Opn(ts, Ballot(1, 1), opn)
    && (exists pkt :: pkt in ts.t_environment.sentPackets && IsNew2bPacket(pkt, opn))
    && PerformanceGuarantee_2a(ts, opn)
    && PerformanceGuarantee_2b(ts, opn)
    && PerformanceGuarantee_Response(ts)
    && 0 <= ts.t_replicas[1].v.nextActionIndex <= 9
    && r.proposer.current_state == 2
    && r.proposer.next_operation_number_to_propose > opn

    // Learner and Executor states
    && BalLeq(r.learner.max_ballot_seen, Ballot(1, 1))
    && r.executor.ops_complete >= opn


    && (r.executor.ops_complete == opn
        ==> Before_Request_Executed(ts, l, opn)
    )

    && (r.executor.ops_complete > opn
        ==> After_Request_Executed(ts, l, opn)
    )

    && ((exists pkt :: pkt in ts.t_environment.sentPackets && IsNewReplyPacket(ts, pkt))
    ==> && r.executor.ops_complete > opn)

    && (Get2bCount(r, opn, Ballot(1, 1)) > 0
        ==>
        |r.learner.unexecuted_learner_state[opn].candidate_learned_value| >= 1
    )

    && (forall opn' | opn' in r.learner.unexecuted_learner_state 
        :: opn' == opn && |r.learner.unexecuted_learner_state[opn'].received_2b_message_senders| > 0 )

    && (opn in r.learner.unexecuted_learner_state 
        ==>
        && r.learner.max_ballot_seen == Ballot(1, 1)
        && (forall id :: id in r.learner.unexecuted_learner_state[opn].received_2b_message_senders ==> id in ts.constants.config.replica_ids)
    )
}

predicate Before_Request_Executed(ts:TimestampedRslState, l:TimestampedLScheduler, opn:OperationNumber)
    requires |ts.t_replicas| > 2
    requires RslConsistency(ts)
    requires 0 <= l.v.nextActionIndex <= 9
{
    var r := l.v.replica;
    && (!exists pkt :: pkt in ts.t_environment.sentPackets && IsNewReplyPacket(ts, pkt))
    && (Get2bCount(r, opn, Ballot(1, 1)) <= LMinQuorumSize(ts.constants.config))
    && (Get2bCount(r, opn, Ballot(1, 1)) < LMinQuorumSize(ts.constants.config)
        ==> Count2b_Lt_Quorum(ts, r, opn))
    && (Get2bCount(r, opn, Ballot(1, 1)) == LMinQuorumSize(ts.constants.config)
        ==> TimeLe(l.ts, TimeBoundPhase2LeaderPost(l.v.nextActionIndex)))
    && (Get2bCount(r, opn, Ballot(1, 1)) == LMinQuorumSize(ts.constants.config)
            ==> Count2b_Eq_Quorum(ts, l, r, opn))
}

predicate After_Request_Executed(ts:TimestampedRslState, l:TimestampedLScheduler, opn:OperationNumber)
    requires |ts.t_replicas| > 2
    requires RslConsistency(ts)
    requires 0 <= l.v.nextActionIndex <= 9
{
    var r := l.v.replica;
    && (exists pkt :: pkt in ts.t_environment.sentPackets && IsNewReplyPacket(ts, pkt))
}


predicate Count2b_Lt_Quorum(ts:TimestampedRslState, r:LReplica, opn:OperationNumber) 
    requires |ts.t_replicas| > 2
    requires RslConsistency(ts)
{
    r.executor.next_op_to_execute.OutstandingOpUnknown?
}

predicate Count2b_Eq_Quorum(ts:TimestampedRslState, l:TimestampedLScheduler, r:LReplica, opn:OperationNumber) 
    requires r == l.v.replica
    requires |ts.t_replicas| > 2
    requires RslConsistency(ts)
{
    || (&& 1 <= l.v.nextActionIndex < 6 
        && r.executor.next_op_to_execute.OutstandingOpUnknown?  
    )
    || (&& l.v.nextActionIndex == 6 
        && r.executor.next_op_to_execute.OutstandingOpKnown?
        && |r.executor.next_op_to_execute.v| > 0
    )
}

/*****************************************************************************************
*                                  Misc Definitions                                      *
*****************************************************************************************/

predicate IsNew2aPacket(pkt:TimestampedRslPacket, opn:OperationNumber) {
    && pkt.msg.v.RslMessage_2a?
    && pkt.msg.v.bal_2a == Ballot(1, 1)
    && pkt.msg.v.opn_2a == opn
}

predicate IsNew2bPacket(pkt:TimestampedRslPacket, opn:OperationNumber) {
    && pkt.msg.v.RslMessage_2b?
    && pkt.msg.v.bal_2b == Ballot(1, 1)
    && pkt.msg.v.opn_2b == opn
}

predicate IsNewReplyPacket(ts:TimestampedRslState, pkt:TimestampedRslPacket) 
    requires |ts.t_replicas| > 2
    requires RslConsistency(ts)
{
    && pkt.msg.v.RslMessage_Reply?
    && pkt.src == ts.constants.config.replica_ids[1]
}

function Get2bCount(s:LReplica, opn:OperationNumber, ballot:Ballot) : int
{
  if opn !in s.learner.unexecuted_learner_state then
    0
  else
    |s.learner.unexecuted_learner_state[opn].received_2b_message_senders|
}

}
