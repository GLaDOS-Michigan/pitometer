include "../TimestampedRslSystem.i.dfy"
include "../../CommonProof/Constants.i.dfy"

module Common_Definitions {
import opened TimestampedRslSystem_i
import opened CommonProof__Constants_i


/*****************************************************************************************
*                                   Stage Predicates                                     *
*****************************************************************************************/


predicate InFailover(ts:TimestampedRslState) 
    requires |ts.t_replicas| > 2
{
    ts.t_replicas[1].v.replica.proposer.election_state.current_view == Ballot(1,0)
}

predicate InPhase1(ts:TimestampedRslState) 
    requires |ts.t_replicas| > 2
{
    && ts.t_replicas[1].v.replica.proposer.election_state.current_view == Ballot(1,1)
    && ts.t_replicas[1].v.replica.proposer.current_state != 2
}

predicate InPhase2(ts:TimestampedRslState) 
    requires |ts.t_replicas| > 2
{
    && ts.t_replicas[1].v.replica.proposer.election_state.current_view == Ballot(1,1)
    && ts.t_replicas[1].v.replica.proposer.current_state == 2
}


/*****************************************************************************************
*                                      Time Bounds                                       *
*****************************************************************************************/


function TimeBound1aDeliveryPost() : Timestamp {
    NewLeaderP1_InitTS + D
}

function TimeBound1bDeliveryPost() : Timestamp {
    TimeBound1aDeliveryPost() + ProcessPacket + MaxQueueTime + D
}

function TimeBoundPhase1LeaderPost(nextActionIndex:int) : Timestamp
    requires 0 <= nextActionIndex < 10
{
    TimeBound1bDeliveryPost() + MaxQueueTime + TimeActionRange(nextActionIndex)
}


function TimeBound2aDeliveryPost() : Timestamp {
    NewLeaderP2_InitTS + MbeP2a + D
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
*                                     Definitions                                        *
*****************************************************************************************/


/* Timestamp of initial client request */
ghost const req_time : Timestamp
/* Initial timestamp of replica 1 in Phase 1 */
ghost const NewLeaderP1_InitTS : Timestamp
/* Initial timestamp of replica 1 in Phase 2 */
ghost const NewLeaderP2_InitTS : Timestamp


predicate RslConsistency(ts:TimestampedRslState)
{
    ConstantsAllConsistentInv(UntimestampRslState(ts))
        && WellFormedLConfiguration(ts.constants.config)
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

predicate LeaderSet1bContainsRequest(ts:TimestampedRslState) 
    requires |ts.t_replicas| > 2
{
    var r := ts.t_replicas[1].v.replica;
    && LProposerCanNominateUsingOperationNumber(r.proposer, r.acceptor.log_truncation_point, r.proposer.next_operation_number_to_propose)
    && !LAllAcceptorsHadNoProposal(r.proposer.received_1b_packets, r.proposer.next_operation_number_to_propose)
}


predicate IsNew1aPacket(pkt:TimestampedRslPacket) {
    && pkt.msg.v.RslMessage_1a?
    && pkt.msg.v.bal_1a == Ballot(1, 1)
}


predicate IsNew1bPacket(pkt:TimestampedRslPacket) {
    && pkt.msg.v.RslMessage_1b?
    && pkt.msg.v.bal_1b == Ballot(1, 1)
}

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


predicate IsUndelivered_1aPkt(ts:TimestampedRslState, p:TimestampedRslPacket) {
  p in ts.undeliveredPackets && p.msg.v.RslMessage_1a?
}

predicate IsUndelivered_1bPkt(ts:TimestampedRslState, p:TimestampedRslPacket) {
  p in ts.undeliveredPackets && p.msg.v.RslMessage_1b?
}

predicate IsUndelivered_2aPkt(ts:TimestampedRslState, p:TimestampedRslPacket) {
  p in ts.undeliveredPackets && p.msg.v.RslMessage_2a?
}

predicate IsUndelivered_2bPkt(ts:TimestampedRslState, p:TimestampedRslPacket) {
  p in ts.undeliveredPackets && p.msg.v.RslMessage_2b?
}


predicate All1aPackets_BalLeq(ts:TimestampedRslState, ballot:Ballot) {
    forall pkt | pkt in ts.t_environment.sentPackets && pkt.msg.v.RslMessage_1a?
    ::  && BalLeq(pkt.msg.v.bal_1a, ballot)
}

predicate All1bPackets_BalLeq(ts:TimestampedRslState, ballot:Ballot) {
    forall pkt | pkt in ts.t_environment.sentPackets && pkt.msg.v.RslMessage_1b?
    ::  && BalLeq(pkt.msg.v.bal_1b, ballot)
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

}
