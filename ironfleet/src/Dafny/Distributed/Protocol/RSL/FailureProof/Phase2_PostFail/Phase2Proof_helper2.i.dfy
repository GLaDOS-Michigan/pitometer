include "Phase2Proof.i.dfy"

module Rs2Phase2Proof_PostFail_Helper2 {
import opened RslPhase2Proof_PostFail_i

/* WARNING: this file a timeout of 60s to verify */


/* Proof that a Before_2b_Sent state transitions to a Before_2b_Sent state or 
* After_2b_Sent state */
lemma Before2b_to_MaybeAfter2b_Process2a(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires P2Assumption(ts, opn) && RslConsistency(ts)
    requires P2Assumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires ts.t_replicas[idx].v.nextActionIndex == 0
    requires tios[0].r.msg.v.RslMessage_2a?
    requires Phase2Invariant(ts, opn)
    
    requires Before_2b_Sent_Invariant(ts, opn)
    ensures Before_2b_Sent_Invariant(ts', opn) || After_2b_Sent_Invariant(ts', opn)
{
    var nextActionIndex := ts.t_replicas[idx].v.nextActionIndex;

    // From this point on, nextActionIndex == 0
    var idx_s, idx_s', ios := UntimestampRslState(ts), UntimestampRslState(ts'), UntagLIoOpSeq(tios);
    var idx_r, idx_r' := idx_s.replicas[idx].replica, idx_s'.replicas[idx].replica;
    var sent_packets := ExtractSentPacketsFromIos(ios);
    var m := ios[0].r.msg;
    if ios[0].r.src in idx_r.acceptor.constants.all.config.replica_ids
       && BalLeq(idx_r.acceptor.max_bal, m.bal_2a)
       && LeqUpperBound(m.opn_2a, idx_r.acceptor.constants.all.params.max_integer_val)
    {
        if m.bal_2a == Ballot(1, 1) {
            Before2b_to_After2b(ts, ts', opn, idx, tios, idx_s, idx_s', ios);
        } else {
            assert m.bal_2a == Ballot(1, 0);
            assert forall p | p in sent_packets && p.msg.RslMessage_2b? :: p.msg.bal_2b == Ballot(1, 0);
            assert forall p | p in sent_packets && p.msg.RslMessage_2b? :: p.msg.opn_2b == opn;
            assert forall p | p in sent_packets :: !p.msg.RslMessage_2a?;
            reveal_ExtractSentPacketsFromIos();
            reveal_UntagLIoOpSeq();
            forall pkt | pkt in ts'.undeliveredPackets && pkt.msg.v.RslMessage_2a? 
            ensures pkt in ts.undeliveredPackets
            {}
            assert All2bPackets_BalLeq_Opn(ts', Ballot(1, 0), opn);
            assert Before_2b_Sent_Invariant(ts', opn);
        }
    } else {
        assert sent_packets == [];
        reveal_ExtractSentPacketsFromIos();
        forall io | io in ios 
        ensures !io.LIoOpSend? 
        {}
        forall tio | tio in tios 
        ensures !tio.LIoOpSend? {
            if tio.LIoOpSend? {
                reveal_UntagLIoOpSeq();
                assert UntagLIoOp(tio) in ios;
                assert UntagLIoOp(tio).LIoOpSend?;
            }
        }
        forall p | p in ts'.t_environment.sentPackets 
        ensures p in ts.t_environment.sentPackets 
        {}
        assert Before_2b_Sent_Invariant(ts', opn);
    }
}


lemma Before2b_to_After2b(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>,
    rs:RslState, rs':RslState, iops:seq<RslIo>
) 
    requires rs == UntimestampRslState(ts)
    requires rs' == UntimestampRslState(ts')
    requires iops == UntagLIoOpSeq(tios);
    requires P2Assumption(ts, opn) && RslConsistency(ts)
    requires P2Assumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires Phase2Invariant(ts, opn)
    requires Before_2b_Sent_Invariant(ts, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires LReplicaNextProcessPacket(rs.replicas[idx].replica, rs'.replicas[idx].replica, iops);
    requires ts.t_replicas[idx].v.nextActionIndex == 0
    requires |tios| > 0 && tios[0].LIoOpReceive? && tios[0].r.msg.v.RslMessage_2a? 
    requires iops[0].r.src in rs.replicas[idx].replica.acceptor.constants.all.config.replica_ids
    requires BalLeq(rs.replicas[idx].replica.acceptor.max_bal, iops[0].r.msg.bal_2a)
    requires LeqUpperBound(iops[0].r.msg.opn_2a, rs.replicas[idx].replica.acceptor.constants.all.params.max_integer_val)
    requires iops[0].r.msg.bal_2a == Ballot(1, 1)
    ensures Before_2b_Sent_Invariant(ts', opn) || After_2b_Sent_Invariant(ts', opn)
{
    var m := iops[0].r.msg;
    var sent_packets := ExtractSentPacketsFromIos(iops);
    assert LAcceptorProcess2a(rs.replicas[idx].replica.acceptor, rs'.replicas[idx].replica.acceptor, iops[0].r, sent_packets);
    var msg2b := RslMessage_2b(m.bal_2a, m.opn_2a, m.val_2a);
    assert LBroadcastToEveryone(rs.replicas[idx].replica.acceptor.constants.all.config, rs.replicas[idx].replica.acceptor.constants.my_index, msg2b, sent_packets);
    assert forall p | p in sent_packets :: LIoOpSend(p) in iops;
    if m.opn_2a == opn {
        var pkt_witness := sent_packets[0];
        reveal_ExtractSentPacketsFromIos();
        reveal_UntagLIoOpSeq();
        assert forall p | p in sent_packets :: !p.msg.RslMessage_Reply?;
        assert LIoOpSend(pkt_witness) in iops;
        assert rs.replicas[1].replica.learner == rs'.replicas[1].replica.learner;
        assert BalLeq(rs'.replicas[1].replica.learner.max_ballot_seen, Ballot(1, 1));
        assert After_2b_Sent_Invariant(ts', opn);          
    } else {
        assert forall p | p in sent_packets && p.msg.RslMessage_2b? :: p.msg.opn_2b != opn;
        forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2b?
        ensures BalLeq(pkt.msg.v.bal_2b, Ballot(1, 1))
        {}
        assert Before_2b_Sent_Invariant(ts', opn);    
    }
}



lemma After2b_to_After2b_NonLeaderAction(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires P2Assumption(ts, opn) && RslConsistency(ts)
    requires P2Assumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires Phase2Invariant(ts, opn)
    requires idx != 1;
    requires After_2b_Sent_Invariant(ts, opn)
    ensures After_2b_Sent_Invariant(ts', opn)
{
    assert ts'.t_replicas[1] == ts.t_replicas[1];
    After2b_to_After2b_2bBalOpn(ts, ts', opn, idx, tios);

    forall p | p in ts'.undeliveredPackets && IsNew2aPacket(p, opn)
    ensures p in ts.undeliveredPackets {
        if p !in ts.undeliveredPackets{
            reveal_ExtractSentPacketsFromIos();
            reveal_UntagLIoOpSeq();
            var sr, sr', ios := UntimestampRslState(ts), UntimestampRslState(ts'), UntagLIoOpSeq(tios);
            var sent_packets := ExtractSentPacketsFromIos(ios);
            forall p' | p' in sent_packets 
            ensures !p'.msg.RslMessage_2a?
            {}
            assert false;
        }
    }
    forall p | p in ts'.undeliveredPackets && p !in ts.undeliveredPackets
    ensures p.src == ts'.constants.config.replica_ids[idx]
    {}
    assert ReplicasDistinct(ts'.constants.config.replica_ids, 1, idx);
    forall p | p in ts'.undeliveredPackets && IsNewReplyPacket(ts', p)
    ensures p in ts.undeliveredPackets
    {}
}


lemma After2b_to_After2b_2bBalOpn(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires P2Assumption(ts, opn) && RslConsistency(ts)
    requires P2Assumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires Phase2Invariant(ts, opn)
    requires After_2b_Sent_Invariant(ts, opn)
    ensures forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2b? :: BalLeq(pkt.msg.v.bal_2b, Ballot(1, 1)) && pkt.msg.v.opn_2b == opn
{
    // Any 2b sent must have opn and ballot matching existing 2a's
    forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2b?
    ensures BalLeq(pkt.msg.v.bal_2b, Ballot(1, 1)) && pkt.msg.v.opn_2b == opn
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
                var sent_packets := ExtractSentPacketsFromIos(ios);
                forall p | p in sent_packets 
                ensures !p.msg.RslMessage_2b? {}
            }
        } else {
            var sent_packets := ExtractSentPacketsFromIos(ios);
            forall p | p in sent_packets 
            ensures !p.msg.RslMessage_2b?
            {}
            assert forall p | p in ts'.t_environment.sentPackets && p.msg.v.RslMessage_2b? 
            :: p in ts.t_environment.sentPackets;
        }
    }
}


lemma After2b_to_After2b_LeaderAction(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires P2Assumption(ts, opn) && RslConsistency(ts)
    requires P2Assumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires Phase2Invariant(ts, opn)
    requires idx == 1;
    requires After_2b_Sent_Invariant(ts, opn)
    ensures After_2b_Sent_Invariant(ts', opn)
{
    var lr, lr' := ts.t_replicas[1].v.replica, ts'.t_replicas[1].v.replica;
    var nextActionIndex, nextActionIndex' := ts.t_replicas[1].v.nextActionIndex, ts'.t_replicas[1].v.nextActionIndex;

    After2b_to_After2b_2bBalOpn(ts, ts', opn, idx, tios);
    assert All2aPackets_BalLeq_Opn(ts', Ballot(1, 1), opn);
    assert All2bPackets_BalLeq_Opn(ts', Ballot(1, 1), opn);
    assert PerformanceGuarantee_2a(ts', opn); 

    After2b_to_After2b_LeaderAction_TimeBound2b(ts, ts', opn, tios);
    forall p | p in ts'.undeliveredPackets && IsNew2bPacket(p, opn)
    ensures TimeLe(p.msg.ts, TimeBound2bDeliveryPost())
    {}

    After2b_to_After2b_LeaderAction_TimeBoundReply(ts, ts', opn, tios);
    forall p | p in ts'.undeliveredPackets && IsNewReplyPacket(ts', p)
    ensures TimeLe(p.msg.ts, TimeBoundReplyFinal())
    {}


    After2b_to_After2b_LeaderAction_PreExecution(ts, ts', opn, tios);

    After2b_to_After2b_LeaderAction_PostExecution(ts, ts', opn, tios);

    After2b_to_After2b_LeaderAction_LearnedBatchNotEmpty(ts, ts', opn, tios);

    assert lr'.proposer.next_operation_number_to_propose > opn;

    forall opn' | opn' in lr'.learner.unexecuted_learner_state 
    ensures opn' == opn && |lr'.learner.unexecuted_learner_state[opn'].received_2b_message_senders| > 0
    {}
}

lemma After2b_to_After2b_LeaderAction_TimeBoundReply(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires P2Assumption(ts, opn) && RslConsistency(ts)
    requires P2Assumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', 1, tios)
    requires Phase2Invariant(ts, opn)
    requires After_2b_Sent_Invariant(ts, opn)
    requires All2aPackets_BalLeq_Opn(ts, Ballot(1, 1), opn)
    requires All2bPackets_BalLeq_Opn(ts, Ballot(1, 1), opn)
    requires PerformanceGuarantee_2a(ts, opn)
    requires PerformanceGuarantee_2b(ts, opn)
    requires PerformanceGuarantee_Response(ts)
    ensures forall p | p in ts'.undeliveredPackets && IsNewReplyPacket(ts', p) 
            :: TimeLe(p.msg.ts, TimeBoundReplyFinal())
{}


lemma After2b_to_After2b_LeaderAction_TimeBound2b(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires P2Assumption(ts, opn) && RslConsistency(ts)
    requires P2Assumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', 1, tios)
    requires Phase2Invariant(ts, opn)
    requires After_2b_Sent_Invariant(ts, opn)
    requires All2aPackets_BalLeq_Opn(ts, Ballot(1, 1), opn)
    requires All2bPackets_BalLeq_Opn(ts, Ballot(1, 1), opn)
    requires PerformanceGuarantee_2a(ts, opn)
    requires PerformanceGuarantee_2b(ts, opn)
    requires PerformanceGuarantee_Response(ts)
    ensures forall p | p in ts'.undeliveredPackets && IsNew2bPacket(p, opn) :: TimeLe(p.msg.ts, TimeBound2bDeliveryPost())
{
    var ls, ls' := ts.t_replicas[1], ts'.t_replicas[1];
    var nextActionIndex, nextActionIndex' := ls.v.nextActionIndex, ls'.v.nextActionIndex;
    forall p | p in ts'.undeliveredPackets && IsNew2bPacket(p, opn) 
    ensures TimeLe(p.msg.ts, TimeBound2bDeliveryPost())
    {
        if p !in ts.undeliveredPackets {
            if nextActionIndex == 0 {
                assert TimeLe(p.msg.ts, TimeBound2bDeliveryPost());
            } else {
                forall io | io in tios && io.LIoOpSend?
                ensures !io.s.msg.v.RslMessage_2b? {}
                assert false;
            }
        }
    }
}


lemma After2b_to_After2b_LeaderAction_LearnedBatchNotEmpty(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires P2Assumption(ts, opn) && RslConsistency(ts)
    requires P2Assumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', 1, tios)
    requires Phase2Invariant(ts, opn)
    requires After_2b_Sent_Invariant(ts, opn)
    requires All2aPackets_BalLeq_Opn(ts, Ballot(1, 1), opn)
    requires All2bPackets_BalLeq_Opn(ts, Ballot(1, 1), opn)
    requires PerformanceGuarantee_2a(ts, opn)
    requires PerformanceGuarantee_2b(ts, opn)
    requires PerformanceGuarantee_Response(ts)
    ensures Get2bCount(ts'.t_replicas[1].v.replica, opn, Ballot(1, 1)) > 0 ==>|ts'.t_replicas[1].v.replica.learner.unexecuted_learner_state[opn].candidate_learned_value| >= 1
{
    var ls, ls' := ts.t_replicas[1], ts'.t_replicas[1];
    var nextActionIndex, nextActionIndex' := ls.v.nextActionIndex, ls'.v.nextActionIndex;
    if Get2bCount(ls.v.replica, opn, Ballot(1, 1)) == 0 {
        if nextActionIndex == 0 && tios[0].r.msg.v.RslMessage_2b?{
            var op_learnable := ls.v.replica.executor.ops_complete < opn || (ls.v.replica.executor.ops_complete == opn && ls.v.replica.executor.next_op_to_execute.OutstandingOpUnknown?);
            var m := tios[0].r.msg;
            if op_learnable {
                if tios[0].r.src !in ls.v.replica.learner.constants.all.config.replica_ids || BalLt(m.v.bal_2b, ls.v.replica.learner.max_ballot_seen) {
                    assert Get2bCount(ls'.v.replica, opn, Ballot(1, 1)) == 0;
                } else if BalLt(ls.v.replica.learner.max_ballot_seen, m.v.bal_2b) {
                    assert |m.v.val_2b| > 0;
                    assert ls'.v.replica.learner.unexecuted_learner_state[opn].candidate_learned_value == m.v.val_2b;
                } else if opn !in ls.v.replica.learner.unexecuted_learner_state {
                    assert |m.v.val_2b| > 0;
                    assert ls'.v.replica.learner.unexecuted_learner_state[opn].candidate_learned_value == m.v.val_2b;
                }
                else if tios[0].r.src in ls.v.replica.learner.unexecuted_learner_state[opn].received_2b_message_senders {
                    assert Get2bCount(ls'.v.replica, opn, Ballot(1, 1)) == 0;
                } else {
                    assert false;
                }
                assert |ls'.v.replica.learner.unexecuted_learner_state[opn].candidate_learned_value| >= 1;
            } else {
                assert Get2bCount(ls'.v.replica, opn, Ballot(1, 1)) == 0;
            }
        } else {
            assert Get2bCount(ls'.v.replica, opn, Ballot(1, 1)) == 0;
        }
    }
}


lemma After2b_to_After2b_LeaderAction_PostExecution(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires P2Assumption(ts, opn) && RslConsistency(ts)
    requires P2Assumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', 1, tios)
    requires Phase2Invariant(ts, opn)
    requires After_2b_Sent_Invariant(ts, opn)
    requires All2aPackets_BalLeq_Opn(ts, Ballot(1, 1), opn)
    requires All2bPackets_BalLeq_Opn(ts, Ballot(1, 1), opn)
    requires PerformanceGuarantee_2a(ts, opn)
    requires PerformanceGuarantee_2b(ts, opn)
    requires PerformanceGuarantee_Response(ts)
    ensures ts'.t_replicas[1].v.replica.executor.ops_complete > opn ==> After_Request_Executed(ts', ts'.t_replicas[1], opn)
{
    var ls, ls' := ts.t_replicas[1], ts'.t_replicas[1];
    var nextActionIndex, nextActionIndex' := ls.v.nextActionIndex, ls'.v.nextActionIndex;
    if ls'.v.replica.executor.ops_complete > opn {     
        if ls.v.replica.executor.ops_complete == opn {
            if nextActionIndex == 6 {
                var ios := UntagLIoOpSeq(tios);
                var sent_packets := ExtractSentPacketsFromIos(ios);
                assert ls.v.replica.executor.next_op_to_execute.OutstandingOpKnown?;
                reveal_ExtractSentPacketsFromIos();
                reveal_UntagLIoOpSeq();
                assert After_Request_Executed(ts', ts'.t_replicas[1], opn);
            } else {
                assert ls'.v.replica.executor.ops_complete == opn;
                assert false;
            }
        }
       assert  After_Request_Executed(ts', ts'.t_replicas[1], opn);
    }
}


lemma After2b_to_After2b_LeaderAction_PreExecution(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires P2Assumption(ts, opn) && RslConsistency(ts)
    requires P2Assumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', 1, tios)
    requires Phase2Invariant(ts, opn)
    requires After_2b_Sent_Invariant(ts, opn)
    requires All2aPackets_BalLeq_Opn(ts, Ballot(1, 1), opn)
    requires All2bPackets_BalLeq_Opn(ts, Ballot(1, 1), opn)
    requires PerformanceGuarantee_2a(ts, opn)
    requires PerformanceGuarantee_2b(ts, opn)
    requires PerformanceGuarantee_Response(ts)
    ensures ts'.t_replicas[1].v.replica.executor.ops_complete == opn ==> Before_Request_Executed(ts', ts'.t_replicas[1], opn)
{
    var ls, ls' := ts.t_replicas[1], ts'.t_replicas[1];
    var nextActionIndex, nextActionIndex' := ls.v.nextActionIndex, ls'.v.nextActionIndex;

    if ls'.v.replica.executor.ops_complete == opn {      
        if Get2bCount(ls.v.replica, opn, Ballot(1, 1)) < LMinQuorumSize(ts.constants.config) {
            if nextActionIndex == 0 {
                if tios[0].r.msg.v.RslMessage_2b? && tios[0].r.msg.v.opn_2b == opn {
                    var pkt := tios[0].r;
                    assert nextActionIndex == 0 && nextActionIndex' == 1;
                    assert pkt.msg.v.bal_2b == Ballot(1, 1);
                    assert pkt in ts.undeliveredPackets;
                    assert IsNew2bPacket(pkt, opn);
                    assert TimeLe(pkt.msg.ts, TimeBound2bDeliveryPost());
                    assert ts.t_environment.nextStep.nodeStep == RslStep(0);
                    assert TimeLe(TimeMax(pkt.msg.ts, ls.ts), pkt.msg.ts + MaxQueueTime);
                    assert TimeLe(ls'.ts, pkt.msg.ts + MaxQueueTime + StepToTimeDelta(RslStep(0)));
                    assert TimeLe(ls'.ts, TimeBound2bDeliveryPost() + MaxQueueTime + TimeActionRange(1));
                    assert TimeLe(ls'.ts, TimeBoundPhase2LeaderPost(nextActionIndex'));
                } else {
                    assert Get2bCount(ls'.v.replica, opn, Ballot(1, 1)) <= Get2bCount(ls.v.replica, opn, Ballot(1, 1));
                }
            } else {
                assert Get2bCount(ls'.v.replica, opn, Ballot(1, 1)) <= Get2bCount(ls.v.replica, opn, Ballot(1, 1));
            }
        } else if Get2bCount(ls.v.replica, opn, Ballot(1, 1)) == LMinQuorumSize(ts.constants.config) {
            if nextActionIndex == 0 || nextActionIndex > 6{
                assert exists pkt :: pkt in ts'.t_environment.sentPackets && IsNewReplyPacket(ts', pkt);
                assert false;
            } else {
                if |tios| > 0 && tios[0].LIoOpReceive? {
                    assert false;
                } else if |tios| > 0 && tios[0].LIoOpTimeoutReceive? {
                    assert false;
                } else {
                    assert TimestampedRslNextOneReplica(ts, ts', 1, tios);
                    var hstep := ts.t_environment.nextStep.nodeStep;
                    assert hstep == RslStep(nextActionIndex);
                    assert nextActionIndex' == nextActionIndex + 1;
                    assert TimeLe(ls'.ts, TimeBoundPhase2LeaderPost(nextActionIndex'));
                }
            }
        } 
    }   
}




}