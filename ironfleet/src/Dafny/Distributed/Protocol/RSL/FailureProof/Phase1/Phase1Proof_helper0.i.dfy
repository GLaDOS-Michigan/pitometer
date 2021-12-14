include "Phase1Proof.i.dfy"
include "GenericLemmas.i.dfy"

module RslPhase1Proof_Helper0 {
import opened RslPhase1Proof_i
import opened RslPhase1Proof_Generic


lemma AlwaysInvariantP1_Maintained(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P1Assumption(ts, opn) 
    requires InPhase1(ts') ==> P1Assumption(ts', opn)
    requires TimestampedRslNext(ts, ts')
    requires Phase1Invariant(ts, opn)
    ensures AlwaysInvariantP1(ts', opn)
{   
    assert ts'.t_replicas[1].v.replica.proposer.request_queue == [];
    assert All2aPackets_BalLeq_Opn(ts', Ballot(1, 0), opn);
    AlwaysInvariantP1_Maintained_2bBalOpn(ts, ts', opn);
    AlwaysInvariantP1_Maintained_ViewSuspectors(ts, ts', opn);
    AlwaysInvariantP1_Maintained_NoReplies(ts, ts', opn);
}


lemma AlwaysInvariantP1_Maintained_2bBalOpn(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P1Assumption(ts, opn) 
    requires InPhase1(ts') ==> P1Assumption(ts', opn)
    requires TimestampedRslNext(ts, ts')
    requires Phase1Invariant(ts, opn)
    ensures forall p | p in ts'.t_environment.sentPackets && p.msg.v.RslMessage_2b?
    ::  && BalLeq(p.msg.v.bal_2b, Ballot(1, 0))
        && p.msg.v.opn_2b == opn
{
    if TimestampedRslNextEnvironment(ts, ts') {
        return;
    }
    var idx, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>> :| TimestampedRslNextOneReplica(ts, ts', idx, tios);
    // Any 2b sent must have opn and ballot matching existing 2a's
    forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2b?
    ensures BalLeq(pkt.msg.v.bal_2b, Ballot(1, 0)) && pkt.msg.v.opn_2b == opn
    {
        var sr, sr', ios := UntimestampRslState(ts), UntimestampRslState(ts'), UntagLIoOpSeq(tios);
        if sr.replicas[idx].nextActionIndex == 0 {
            var msg := tios[0].r.msg;
            if msg.v.RslMessage_2a? {
                if tios[0].r.src in sr.replicas[idx].replica.acceptor.constants.all.config.replica_ids
                    && BalLeq(sr.replicas[idx].replica.acceptor.max_bal, msg.v.bal_2a)
                    && LeqUpperBound(msg.v.opn_2a, sr.replicas[idx].replica.acceptor.constants.all.params.max_integer_val)
                {
                    var sent_packets := ExtractSentPacketsFromIos(ios);
                    forall sp | sp in sent_packets 
                    ensures  && sp.msg.RslMessage_2b? 
                        && sp.msg.opn_2b == msg.v.opn_2a
                        && sp.msg.bal_2b == msg.v.bal_2a
                    {}
                } else {
                    assert forall p | p in ts'.t_environment.sentPackets && p.msg.v.RslMessage_2b? 
                    :: p in ts.t_environment.sentPackets;
                }
            } else {
                lemma_No2bSentInReceiveStep_NotReceive2a(ts, ts', opn, idx, tios);
            }
        } else {
            lemma_No2bSentInNonReceiveStep(ts, ts', opn, idx, tios);
        }
    }
}


lemma AlwaysInvariantP1_Maintained_NoReplies(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P1Assumption(ts, opn) 
    requires InPhase1(ts') ==> P1Assumption(ts', opn)
    requires TimestampedRslNext(ts, ts')
    requires Phase1Invariant(ts, opn)
    ensures !exists pkt :: pkt in ts'.t_environment.sentPackets && IsNewReplyPacket(ts', pkt)
{
    if TimestampedRslNextEnvironment(ts, ts') {
        return;
    }
    var idx, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>> :| TimestampedRslNextOneReplica(ts, ts', idx, tios);
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    var nextActionIndex := ls.v.nextActionIndex;
    assert idx != 0;
    if idx == 1 {
        if nextActionIndex == 6 {
            assert ls.v.replica.executor.next_op_to_execute.OutstandingOpUnknown?;
            forall io | io in tios 
            ensures !io.LIoOpSend? {}
        } else {
            lemma_NoRepliesSentInNonExecutionStep(ts, ts', opn, idx, tios);
        }
    } else {
        lemma_NonLeaderDoesNotSendReply(ts, ts', opn, idx, tios);
    }
}


lemma AlwaysInvariantP1_Maintained_ViewSuspectors(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P1Assumption(ts, opn) 
    requires InPhase1(ts') ==> P1Assumption(ts', opn)
    requires TimestampedRslNext(ts, ts')
    requires Phase1Invariant(ts, opn)
    ensures || ts'.t_replicas[1].v.replica.proposer.election_state.current_view_suspectors == {}
            || ts'.t_replicas[1].v.replica.proposer.election_state.current_view_suspectors == {ts'.t_replicas[1].v.replica.proposer.election_state.constants.my_index}
{
    if !TimestampedRslNextEnvironment(ts, ts') {
        var idx, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>> :| TimestampedRslNextOneReplica(ts, ts', idx, tios);
        var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
        var nextActionIndex := ls.v.nextActionIndex;
        if idx == 1 {
            var es, es' := ls.v.replica.proposer.election_state, ls'.v.replica.proposer.election_state;
            assert es'.constants.my_index == es.constants.my_index;
            if nextActionIndex == 0 {
                assert !tios[0].r.msg.v.RslMessage_Heartbeat?;
                assert es.current_view_suspectors == es'.current_view_suspectors;
            } else if nextActionIndex == 7 {
                if es.current_view_suspectors == {es.constants.my_index} {
                    assert es'.current_view_suspectors == {es'.constants.my_index};
                } else {
                    assert  || es'.current_view_suspectors == {} 
                            || es'.current_view_suspectors == {es'.constants.my_index};
                }
            } else {
                assert es.current_view_suspectors == es'.current_view_suspectors;
            }
        }
    }
}

}