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
    && LeaderAlwaysOne(s)
    && (var nextStep := s.t_environment.nextStep; nextStep.LEnvStepHostIos? ==>
        (forall io :: io in nextStep.ios && io.LIoOpReceive? ==> !io.r.msg.v.RslMessage_Heartbeat?))
    
}


predicate LeaderAlwaysOne(s:TimestampedRslState)
    requires |s.t_replicas| > 2 
{
    && s.t_replicas[1].v.replica.proposer.election_state.current_view == Ballot(1, 1)
    && (forall idx | 2 <= idx < |s.t_replicas| 
        :: s.t_replicas[idx].v.replica.proposer.current_state == 0)
}


/*****************************************************************************************
*                                     Invariants                                        *
*****************************************************************************************/




predicate AlwaysInvariant(ts:TimestampedRslState, opn:OperationNumber)
    requires |ts.t_replicas| > 2
{
    && ServersAreNotClients(ts)
    && LSetOfMessage1b(ts.t_replicas[1].v.replica.proposer.received_1b_packets)
    && AlwaysInvariant_RequestSrcAndBatchSize(ts, opn)
    && ts.t_replicas[1].v.replica.proposer.request_queue == []
    && (forall pkt | pkt in ts.undeliveredPackets :: pkt in ts.t_environment.sentPackets)
    && ts.t_replicas[1].v.replica.learner.unexecuted_learner_state == map[]
    && ts.t_replicas[1].v.replica.executor.next_op_to_execute == OutstandingOpUnknown()
}


predicate AlwaysInvariant_RequestSrcAndBatchSize(ts:TimestampedRslState, opn:OperationNumber)
    requires |ts.t_replicas| > 2
    requires LSetOfMessage1b(ts.t_replicas[1].v.replica.proposer.received_1b_packets)
{
    && (forall pkt | pkt in ts.undeliveredPackets && pkt.msg.v.RslMessage_2a? :: RequestBatchSrcInClientIds(ts, pkt.msg.v.val_2a))
    && (forall pkt | pkt in ts.undeliveredPackets && pkt.msg.v.RslMessage_2b? :: RequestBatchSrcInClientIds(ts, pkt.msg.v.val_2b))
    && (forall pkt | pkt in ts.t_environment.sentPackets && pkt.msg.v.RslMessage_2a? :: |pkt.msg.v.val_2a| > 0)
    && (forall pkt | pkt in ts.t_environment.sentPackets && pkt.msg.v.RslMessage_2b? :: |pkt.msg.v.val_2b| > 0)
    && (forall pkt, opn'| pkt in ts.t_environment.sentPackets && pkt.msg.v.RslMessage_1b? && opn' in pkt.msg.v.votes
        ::  && |pkt.msg.v.votes[opn'].max_val| > 0
            && RequestBatchSrcInClientIds(ts, pkt.msg.v.votes[opn'].max_val))
    && (forall v | LValIsHighestNumberedProposal(v, ts.t_replicas[1].v.replica.proposer.received_1b_packets, opn)
        :: |v| > 0)
    && (forall pkt, op | pkt in ts.t_replicas[1].v.replica.proposer.received_1b_packets && pkt.msg.RslMessage_1b? && op in pkt.msg.votes
        :: RequestBatchSrcInClientIds(ts, pkt.msg.votes[op].max_val))
    && (forall idx, opn'| 0 <= idx < |ts.t_replicas| && opn' in ts.t_replicas[idx].v.replica.acceptor.votes
        ::  && |ts.t_replicas[idx].v.replica.acceptor.votes[opn'].max_val| > 0   
            &&  RequestBatchSrcInClientIds(ts, ts.t_replicas[idx].v.replica.acceptor.votes[opn'].max_val)
    )
}



}
