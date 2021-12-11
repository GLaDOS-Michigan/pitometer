include "Phase2Proof.i.dfy"
include "GenericLemmas.i.dfy"

module Rs2Phase2Proof_PostFail_Helper0 {
import opened RslPhase2Proof_PostFail_i
import opened RslPhase2Proof_PostFail_Generic


lemma AlwaysInvariant_Maintained(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber) 
    requires P2Assumption(ts, opn) && P2Assumption(ts', opn)
    requires RslConsistency(ts) && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires Phase2Invariant(ts, opn)
    ensures AlwaysInvariant(ts', opn)
{   
    assert ts'.t_replicas[1].v.replica.proposer.request_queue == [];
    assert forall pkt | pkt in ts'.t_replicas[1].v.replica.proposer.received_1b_packets && pkt.msg.RslMessage_1b? :: forall op | op in pkt.msg.votes :: RequestBatchSrcInClientIds(ts', pkt.msg.votes[op].max_val);
    
    AlwaysInvariant_Maintained_BatchSize2a(ts, ts', opn);
    forall pkt | IsUndelivered_2aPkt(ts', pkt)
    ensures |pkt.msg.v.val_2a| > 0 
    {}
    
    AlwaysInvariant_Maintained_BatchSize2b(ts, ts', opn);
    forall pkt | IsUndelivered_2bPkt(ts', pkt) 
    ensures |pkt.msg.v.val_2b| > 0
    {}

    AlwaysInvariant_Maintained_ClientSrc2a(ts, ts', opn);
    forall pkt | IsUndelivered_2aPkt(ts', pkt) 
    ensures RequestBatchSrcInClientIds(ts', pkt.msg.v.val_2a)
    {}

    AlwaysInvariant_Maintained_ClientSrc2b(ts, ts', opn);
    forall pkt | IsUndelivered_2bPkt(ts', pkt) 
    ensures RequestBatchSrcInClientIds(ts', pkt.msg.v.val_2b)
    {}

    AlwaysInvariant_Maintained_ClientSrc_BatchSize1b(ts, ts', opn);
    forall pkt, opn' | IsUndelivered_1bPkt(ts', pkt)  && opn' in pkt.msg.v.votes
    ensures  && |pkt.msg.v.votes[opn'].max_val| > 0
             && RequestBatchSrcInClientIds(ts', pkt.msg.v.votes[opn'].max_val)
    {}

    AlwaysInvariant_Maintained_ClientSrc_BatchSizeAcceptor(ts, ts', opn);
    forall idx, opn'| 0 <= idx < |ts'.t_replicas| && opn' in ts'.t_replicas[idx].v.replica.acceptor.votes
    ensures && |ts'.t_replicas[idx].v.replica.acceptor.votes[opn'].max_val| > 0   
            &&  RequestBatchSrcInClientIds(ts', ts'.t_replicas[idx].v.replica.acceptor.votes[opn'].max_val)
    {}

    var uls := ts'.t_replicas[1].v.replica.learner.unexecuted_learner_state;
    forall opn | opn in uls ensures 
    RequestBatchSrcInClientIds(ts', uls[opn].candidate_learned_value)
    {}

    assert ts'.t_replicas[1].v.replica.executor.next_op_to_execute.OutstandingOpKnown?
        ==> RequestBatchSrcInClientIds(ts', ts'.t_replicas[1].v.replica.executor.next_op_to_execute.v);
}


lemma AlwaysInvariant_Maintained_ClientSrc_BatchSizeAcceptor(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber) 
    requires P2Assumption(ts, opn) && P2Assumption(ts', opn)
    requires RslConsistency(ts) && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires Phase2Invariant(ts, opn)
    ensures forall idx, opn'| 0 <= idx < |ts'.t_replicas| && opn' in ts'.t_replicas[idx].v.replica.acceptor.votes
        ::  && |ts'.t_replicas[idx].v.replica.acceptor.votes[opn'].max_val| > 0   
            &&  RequestBatchSrcInClientIds(ts', ts'.t_replicas[idx].v.replica.acceptor.votes[opn'].max_val)
{
    forall idx, opn' | 0 <= idx < |ts'.t_replicas| && opn' in ts'.t_replicas[idx].v.replica.acceptor.votes
    ensures && |ts'.t_replicas[idx].v.replica.acceptor.votes[opn'].max_val| > 0   
            &&  RequestBatchSrcInClientIds(ts', ts'.t_replicas[idx].v.replica.acceptor.votes[opn'].max_val)
    {
        if !TimestampedRslNextEnvironment(ts, ts') {
            var x, tios :| TimestampedRslNextOneReplica(ts, ts', x, tios);
            if x == idx {
                AlwaysInvariant_Maintained_ClientSrc_BatchSizeAcceptor_OneReplica(ts, ts', opn, opn', idx, tios);
            } else {
                assert ts'.t_replicas[idx] ==  ts.t_replicas[idx];
            }
        } else {
            assert ts'.t_replicas[idx] ==  ts.t_replicas[idx];
        }
        assert RequestBatchSrcInClientIds(ts', ts'.t_replicas[idx].v.replica.acceptor.votes[opn'].max_val);
    }
}

lemma AlwaysInvariant_Maintained_ClientSrc_BatchSizeAcceptor_OneReplica(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, opn':OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires P2Assumption(ts, opn) && P2Assumption(ts', opn)
    requires RslConsistency(ts) && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires Phase2Invariant(ts, opn)
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires opn' in ts'.t_replicas[idx].v.replica.acceptor.votes
    ensures |ts'.t_replicas[idx].v.replica.acceptor.votes[opn'].max_val| > 0   
    ensures RequestBatchSrcInClientIds(ts', ts'.t_replicas[idx].v.replica.acceptor.votes[opn'].max_val)
{
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    var nextActionIndex := ls.v.nextActionIndex;
    if nextActionIndex == 0 {
        if tios[0].r.msg.v.RslMessage_2a? {
            assert |ts'.t_replicas[idx].v.replica.acceptor.votes[opn'].max_val| > 0;
            assert RequestBatchSrcInClientIds(ts', ts'.t_replicas[idx].v.replica.acceptor.votes[opn'].max_val);
        } else {
            assert ts'.t_replicas[idx].v.replica.acceptor.votes == ts.t_replicas[idx].v.replica.acceptor.votes;
        }
    } else if nextActionIndex == 4 {
        assert |ts'.t_replicas[idx].v.replica.acceptor.votes[opn'].max_val| > 0;   
        assert RequestBatchSrcInClientIds(ts', ts'.t_replicas[idx].v.replica.acceptor.votes[opn'].max_val);
    } else {
        assert ts'.t_replicas[idx].v.replica.acceptor == ts.t_replicas[idx].v.replica.acceptor;
    }
    assert RequestBatchSrcInClientIds(ts', ts'.t_replicas[idx].v.replica.acceptor.votes[opn'].max_val);
}


lemma AlwaysInvariant_Maintained_ClientSrc_BatchSize1b(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber) 
    requires P2Assumption(ts, opn) && P2Assumption(ts', opn)
    requires RslConsistency(ts) && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires Phase2Invariant(ts, opn)
    ensures forall pkt, opn'| IsUndelivered_1bPkt(ts', pkt) && opn' in pkt.msg.v.votes
        ::  && |pkt.msg.v.votes[opn'].max_val| > 0
            && RequestBatchSrcInClientIds(ts', pkt.msg.v.votes[opn'].max_val)
{   
    forall pkt, opn' | IsUndelivered_1bPkt(ts', pkt) && opn' in pkt.msg.v.votes
    ensures && |pkt.msg.v.votes[opn'].max_val| > 0
            && RequestBatchSrcInClientIds(ts', pkt.msg.v.votes[opn'].max_val)
    {
        if pkt !in ts.undeliveredPackets {
            assert !TimestampedRslNextEnvironment(ts, ts');
            var idx, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>> :| TimestampedRslNextOneReplica(ts, ts', idx, tios);
            var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
            var nextActionIndex, nextActionIndex' := ls.v.nextActionIndex, ls'.v.nextActionIndex;
            if nextActionIndex == 0 {
                if tios[0].r.msg.v.RslMessage_1a? {
                    assert |pkt.msg.v.votes[opn'].max_val| > 0;
                    assert RequestBatchSrcInClientIds(ts', pkt.msg.v.votes[opn'].max_val);
                } else {
                    lemma_No1bSentInReceiveStep_NotReceive1a(ts, ts', opn, idx, tios);
                    assert false;
                }
            } else {
                assert P2Assumption(ts', opn);
                lemma_No1bSentInNonReceiveStep(ts, ts', opn, idx, tios);
                assert false;
            }
        }
        assert RequestBatchSrcInClientIds(ts', pkt.msg.v.votes[opn'].max_val);
    }
}



lemma {:timeLimitMultiplier 2} AlwaysInvariant_Maintained_BatchSize2a(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber) 
    requires P2Assumption(ts, opn)
    requires P2Assumption(ts', opn)
    requires RslConsistency(ts) && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires Phase2Invariant(ts, opn)
    ensures forall pkt | IsUndelivered_2aPkt(ts', pkt) :: |pkt.msg.v.val_2a| > 0 
{   
    forall pkt | IsUndelivered_2aPkt(ts', pkt) 
    ensures |pkt.msg.v.val_2a| > 0 
    {
        if pkt !in ts.undeliveredPackets {
            assert !TimestampedRslNextEnvironment(ts, ts');
            var idx, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>> :| TimestampedRslNextOneReplica(ts, ts', idx, tios);
            var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
            var nextActionIndex, nextActionIndex' := ls.v.nextActionIndex, ls'.v.nextActionIndex;
            if idx == 1 {
                if nextActionIndex == 3 {
                    var clock := SpontaneousClock(UntagLIoOpSeq(tios)).t;
                    if !LProposerCanNominateUsingOperationNumber(ls.v.replica.proposer, ls.v.replica.acceptor.log_truncation_point, ls.v.replica.proposer.next_operation_number_to_propose) {
                        assert pkt in ts.undeliveredPackets;
                        assert false;
                    } else if !LAllAcceptorsHadNoProposal(ls.v.replica.proposer.received_1b_packets, ls.v.replica.proposer.next_operation_number_to_propose){
                        assert |pkt.msg.v.val_2a| > 0;
                    } else if (exists opn' :: opn' > ls.v.replica.proposer.next_operation_number_to_propose && !LAllAcceptorsHadNoProposal(ls.v.replica.proposer.received_1b_packets, opn'))
                                || |ls.v.replica.proposer.request_queue| >= ls.v.replica.proposer.constants.all.params.max_batch_size
                                || (|ls.v.replica.proposer.request_queue| > 0 && ls.v.replica.proposer.incomplete_batch_timer.IncompleteBatchTimerOn? && clock >= ls.v.replica.proposer.incomplete_batch_timer.when) {
                        assert |pkt.msg.v.val_2a| > 0;
                    } else if |ls.v.replica.proposer.request_queue| > 0 && ls.v.replica.proposer.incomplete_batch_timer.IncompleteBatchTimerOff? {
                        assert pkt in ts.undeliveredPackets;
                        assert false;
                    } else {
                        assert pkt in ts.undeliveredPackets;
                        assert false;
                    }
                } else {
                    lemma_No2aSentInNon2aStep(ts, ts', opn, 1, tios);
                    assert false;
                }
            } else {
                lemma_NonLeaderDoesNotSend2a(ts, ts', opn, idx, tios);
                assert false;
            }
        }
    }
}


lemma AlwaysInvariant_Maintained_ClientSrc2b(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber) 
    requires P2Assumption(ts, opn) && P2Assumption(ts', opn)
    requires RslConsistency(ts) && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires Phase2Invariant(ts, opn)
    ensures forall pkt | IsUndelivered_2bPkt(ts', pkt) :: RequestBatchSrcInClientIds(ts', pkt.msg.v.val_2b)
{   
    forall pkt | IsUndelivered_2bPkt(ts', pkt) 
    ensures RequestBatchSrcInClientIds(ts', pkt.msg.v.val_2b)
    {
        if pkt !in ts.undeliveredPackets {
            assert !TimestampedRslNextEnvironment(ts, ts');
            var idx, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>> :| TimestampedRslNextOneReplica(ts, ts', idx, tios);
            var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
            var nextActionIndex, nextActionIndex' := ls.v.nextActionIndex, ls'.v.nextActionIndex;
            if nextActionIndex == 0 {
                if tios[0].r.msg.v.RslMessage_2a? {
                    assert pkt.msg.v.val_2b == tios[0].r.msg.v.val_2a;
                    assert |pkt.msg.v.val_2b| > 0;
                } else {
                    lemma_No2bSentInReceiveStep_NotReceive2a(ts, ts', opn, idx, tios);
                    assert false;
                }
            } else {
                assert P2Assumption(ts', opn);
                lemma_No2bSentInNonReceiveStep(ts, ts', opn, idx, tios);
                assert false;
            }
        }
    }
}

lemma AlwaysInvariant_Maintained_BatchSize2b(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber) 
    requires P2Assumption(ts, opn) && P2Assumption(ts', opn)
    requires RslConsistency(ts) && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires Phase2Invariant(ts, opn)
    ensures forall pkt | IsUndelivered_2bPkt(ts', pkt) :: |pkt.msg.v.val_2b| > 0 
{   
    forall pkt | IsUndelivered_2bPkt(ts', pkt) 
    ensures |pkt.msg.v.val_2b| > 0 
    {
        if pkt !in ts.undeliveredPackets {
            assert !TimestampedRslNextEnvironment(ts, ts');
            var idx, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>> :| TimestampedRslNextOneReplica(ts, ts', idx, tios);
            var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
            var nextActionIndex, nextActionIndex' := ls.v.nextActionIndex, ls'.v.nextActionIndex;
            if nextActionIndex == 0 && tios[0].r.msg.v.RslMessage_2a? {
                assert pkt.msg.v.val_2b == tios[0].r.msg.v.val_2a;
                assert |pkt.msg.v.val_2b| > 0;
            } else {
                lemma_No2bSentInNonReceiveStep(ts, ts', opn, idx, tios);
                lemma_No2bSentInReceiveStep_NotReceive2a(ts, ts', opn, idx, tios);
                assert false;
            }
        }
    }
}

lemma {:timeLimitMultiplier 2} AlwaysInvariant_Maintained_ClientSrc2a(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber) 
    requires P2Assumption(ts, opn) && P2Assumption(ts', opn)
    requires RslConsistency(ts) && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires Phase2Invariant(ts, opn)
    ensures forall pkt | IsUndelivered_2aPkt(ts', pkt) :: RequestBatchSrcInClientIds(ts', pkt.msg.v.val_2a)
{   
    forall pkt | IsUndelivered_2aPkt(ts', pkt) 
    ensures RequestBatchSrcInClientIds(ts', pkt.msg.v.val_2a)
    {
        if pkt !in ts.undeliveredPackets {
            assert !TimestampedRslNextEnvironment(ts, ts');
            var idx, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>> :| TimestampedRslNextOneReplica(ts, ts', idx, tios);
            var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
            var nextActionIndex, nextActionIndex' := ls.v.nextActionIndex, ls'.v.nextActionIndex;
            if idx == 1 {
                if nextActionIndex == 3 {
                    var clock := SpontaneousClock(UntagLIoOpSeq(tios)).t;
                    if !LProposerCanNominateUsingOperationNumber(ls.v.replica.proposer, ls.v.replica.acceptor.log_truncation_point, ls.v.replica.proposer.next_operation_number_to_propose) {
                        assert pkt in ts.undeliveredPackets;
                        assert false;
                    } else if !LAllAcceptorsHadNoProposal(ls.v.replica.proposer.received_1b_packets, ls.v.replica.proposer.next_operation_number_to_propose){
                        assert RequestBatchSrcInClientIds(ts', pkt.msg.v.val_2a);
                    } else if (exists opn' :: opn' > ls.v.replica.proposer.next_operation_number_to_propose && !LAllAcceptorsHadNoProposal(ls.v.replica.proposer.received_1b_packets, opn'))
                                || |ls.v.replica.proposer.request_queue| >= ls.v.replica.proposer.constants.all.params.max_batch_size
                                || (|ls.v.replica.proposer.request_queue| > 0 && ls.v.replica.proposer.incomplete_batch_timer.IncompleteBatchTimerOn? && clock >= ls.v.replica.proposer.incomplete_batch_timer.when) {
                        assert RequestBatchSrcInClientIds(ts', pkt.msg.v.val_2a);
                    } else if |ls.v.replica.proposer.request_queue| > 0 && ls.v.replica.proposer.incomplete_batch_timer.IncompleteBatchTimerOff? {
                        assert pkt in ts.undeliveredPackets;
                        assert false;
                    } else {
                        assert pkt in ts.undeliveredPackets;
                        assert false;
                    }
                } else {
                    lemma_No2aSentInNon2aStep(ts, ts', opn, 1, tios); 
                    assert false;
                }
            } else {
                lemma_NonLeaderDoesNotSend2a(ts, ts', opn, idx, tios);
                assert false;
            }
        }
    }
}

}