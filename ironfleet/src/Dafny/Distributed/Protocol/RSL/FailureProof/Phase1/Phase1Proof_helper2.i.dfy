include "Phase1Proof.i.dfy"
include "GenericLemmas.i.dfy"

module RslPhase1Proof_Helper2 {
import opened RslPhase1Proof_i
import opened RslPhase1Proof_Generic


/* Proof that a Phase1 state maybe transitions to a Phase2 state when nextAction == 2 */
lemma Phase1_to_MaybePhase2(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P1Assumption(ts, opn)
    requires InPhase1(ts') ==> P1Assumption(ts', opn)
    requires TimestampedRslNext(ts, ts')
    requires Phase1Invariant(ts, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires ts.t_replicas[idx].v.nextActionIndex == 2
    requires idx != 0
    ensures InPhase1(ts') || InPhase2(ts')
    ensures InPhase1(ts') ==> Phase1Invariant(ts', opn)
    ensures InPhase2(ts') ==> P2.Phase2Invariant(ts', opn)
{
    if idx == 1 {
        var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
        var goToPhase2 :=  && |ls.v.replica.proposer.received_1b_packets| >= LMinQuorumSize(ls.v.replica.proposer.constants.all.config)
                           && LSetOfMessage1bAboutBallot(ls.v.replica.proposer.received_1b_packets, ls.v.replica.proposer.max_ballot_i_sent_1a);
        if goToPhase2 {
            Phase1_to_Phase2(ts, ts', opn, tios);
        } else {
            // TODO
            assume false;
        }
    } else {
        assert InPhase1(ts');
        assert Phase1Invariant(ts', opn);
    }
}


lemma Phase1_to_Phase2(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P1Assumption(ts, opn)
    requires InPhase1(ts') ==> P1Assumption(ts', opn)
    requires TimestampedRslNext(ts, ts')
    requires Phase1Invariant(ts, opn)
    requires TimestampedRslNextOneReplica(ts, ts', 1, tios);
    requires ts.t_replicas[1].v.nextActionIndex == 2
    requires && |ts.t_replicas[1].v.replica.proposer.received_1b_packets| >= LMinQuorumSize(ts.t_replicas[1].v.replica.proposer.constants.all.config)
             && LSetOfMessage1bAboutBallot(ts.t_replicas[1].v.replica.proposer.received_1b_packets, ts.t_replicas[1].v.replica.proposer.max_ballot_i_sent_1a);
    ensures InPhase2(ts')
    ensures P2.Phase2Invariant(ts', opn)
{
    var ls, ls' := ts.t_replicas[1], ts'.t_replicas[1];
    assert InPhase2(ts');

    lemma_No2aSentInNon2aStep(ts, ts', opn, 1, tios);
    lemma_No2bSentInNonReceiveStep(ts, ts', opn, 1, tios);
    lemma_NoRepliesSentInNonExecutionStep(ts, ts', opn, 1, tios);
    Phase1_to_Phase2_RequestSrcAndBatchSize(ts, ts', opn, tios);

    assert TimeLe(ls'.ts, TimeBoundPhase1LeaderPost(3));
    assert TimeLe(ls'.ts, NewLeaderP2_InitTS); 
    assert P2.Before_2a_Sent_Invariant(ts', opn);
}


lemma Phase1_to_Phase2_RequestSrcAndBatchSize(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P1Assumption(ts, opn)
    requires InPhase1(ts') ==> P1Assumption(ts', opn)
    requires TimestampedRslNext(ts, ts')
    requires Phase1Invariant(ts, opn)
    requires TimestampedRslNextOneReplica(ts, ts', 1, tios);
    requires ts.t_replicas[1].v.nextActionIndex == 2
    requires InPhase2(ts')
    requires P2.AlwaysInvariantP2_RequestSrcAndBatchSize(ts, opn)
    ensures P2.AlwaysInvariantP2_RequestSrcAndBatchSize(ts', opn)
{
    var ls, ls' := ts.t_replicas[1], ts'.t_replicas[1];

    // RequestSrcAndBatchSize -- 1b
    forall p, op| IsUndelivered_1bPkt(ts', p) && op in p.msg.v.votes   
    ensures && RequestBatchSrcInClientIds(ts', p.msg.v.votes[op].max_val)
            && |p.msg.v.votes[op].max_val| > 0
    {
        lemma_No1bSentInNonReceiveStep(ts, ts', opn, 1, tios);
        assert IsUndelivered_1bPkt(ts, p);
        forall r | r in p.msg.v.votes[op].max_val
        ensures r.client in ts'.constants.config.clientIds {}
    }

    // RequestSrcAndBatchSize -- 2a
    forall p | IsUndelivered_2aPkt(ts', p)   
    ensures && RequestBatchSrcInClientIds(ts', p.msg.v.val_2a)
            && |p.msg.v.val_2a| > 0
    {
        lemma_No2aSentInNon2aStep(ts, ts', opn, 1, tios);
        assert IsUndelivered_2aPkt(ts, p);
        forall r | r in p.msg.v.val_2a
        ensures r.client in ts'.constants.config.clientIds
        {}
    }

    // RequestSrcAndBatchSize -- 2b
    lemma_No2bSentInNonReceiveStep(ts, ts', opn, 1, tios);

    // RequestSrcAndBatchSize -- leader's received_1b_packets set
    forall p, op | p in ls'.v.replica.proposer.received_1b_packets && p.msg.RslMessage_1b? && op in p.msg.votes
    ensures RequestBatchSrcInClientIds(ts', p.msg.votes[op].max_val)
    {
        assert p in ls.v.replica.proposer.received_1b_packets;
        forall r | r in p.msg.votes[op].max_val
        ensures r.client in ts'.constants.config.clientIds {}
    }

    // RequestSrcAndBatchSize -- Acceptors
    forall i, op| 0 <= i < |ts'.t_replicas| && op in ts'.t_replicas[i].v.replica.acceptor.votes
    ensures && |ts'.t_replicas[i].v.replica.acceptor.votes[op].max_val| > 0   
            &&  RequestBatchSrcInClientIds(ts', ts'.t_replicas[i].v.replica.acceptor.votes[op].max_val)
    {
        assert ls'.v.replica.acceptor == ls.v.replica.acceptor;
        forall r | r in ts'.t_replicas[i].v.replica.acceptor.votes[op].max_val
        ensures r.client in ts'.constants.config.clientIds {}
    }   
}


}