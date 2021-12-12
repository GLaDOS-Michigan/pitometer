include "Phase2Proof.i.dfy"
include "GenericLemmas.i.dfy"

module Rs2Phase2Proof_PostFail_Helper1 {
import opened RslPhase2Proof_PostFail_i
import opened RslPhase2Proof_PostFail_Generic


lemma PacketsBallotInvariant_ReceiveStep(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires P2Assumption(ts, opn) && RslConsistency(ts)
    requires P2Assumption(ts', opn) && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires Phase2Invariant(ts, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires ts.t_replicas[idx].v.nextActionIndex == 0 
    ensures PacketsBallotInvariant(ts')
{
    forall pkt | pkt in ts'.undeliveredPackets 
    ensures ExistingPacketsBallot(pkt)
    {}
}

lemma PacketsBallotInvariant_NoReceiveStep(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires P2Assumption(ts, opn) && RslConsistency(ts)
    requires P2Assumption(ts', opn) && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires Phase2Invariant(ts, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires ts.t_replicas[idx].v.nextActionIndex != 0 
    ensures PacketsBallotInvariant(ts')
{
    forall pkt | pkt in ts'.undeliveredPackets 
    ensures ExistingPacketsBallot(pkt)
    {}
}

lemma Before2a_to_Before2b(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires P2Assumption(ts, opn) && RslConsistency(ts)
    requires P2Assumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', 1, tios);
    requires Phase2Invariant(ts, opn)
    requires Before_2a_Sent_Invariant(ts, opn);
    ensures Before_2b_Sent_Invariant(ts', opn);
{
    var ios := UntagLIoOpSeq(tios);
    var ls, ls' := ts.t_replicas[1], ts'.t_replicas[1];
    var r, r' := ts.t_replicas[1].v.replica, ts'.t_replicas[1].v.replica;
    var sent_packets := ExtractSentPacketsFromIos(ios);
    var v :| 
        (&& LValIsHighestNumberedProposal(v, r.proposer.received_1b_packets, opn)
         && LBroadcastToEveryone(r.proposer.constants.all.config, r.proposer.constants.my_index, RslMessage_2a(r.proposer.max_ballot_i_sent_1a, opn, v), sent_packets)
    );

    forall p | p in ts'.t_environment.sentPackets && p.msg.v.RslMessage_2a?
    ensures && BalLeq(p.msg.v.bal_2a, Ballot(1, 1))
            && p.msg.v.opn_2a == opn
    {
        if p !in ts.t_environment.sentPackets {
            var clock := SpontaneousClock(ios);
            if !LProposerCanNominateUsingOperationNumber(ls.v.replica.proposer, ls.v.replica.acceptor.log_truncation_point, ls.v.replica.proposer.next_operation_number_to_propose) {
                reveal_ExtractSentPacketsFromIos();
                assert |sent_packets| == 0;
                assert forall io | io in tios :: !io.LIoOpSend?;
                assert p in ts.t_environment.sentPackets;
                assert false;
            } else if !LAllAcceptorsHadNoProposal(ls.v.replica.proposer.received_1b_packets, ls.v.replica.proposer.next_operation_number_to_propose){
                assert p.msg.v.bal_2a == ls.v.replica.proposer.max_ballot_i_sent_1a == Ballot(1, 1);
                assert p.msg.v.opn_2a == ls.v.replica.proposer.next_operation_number_to_propose == opn;
            } else if (exists opn' :: opn' > ls.v.replica.proposer.next_operation_number_to_propose && !LAllAcceptorsHadNoProposal(ls.v.replica.proposer.received_1b_packets, opn'))
                        || |ls.v.replica.proposer.request_queue| >= ls.v.replica.proposer.constants.all.params.max_batch_size
                        || (|ls.v.replica.proposer.request_queue| > 0 && ls.v.replica.proposer.incomplete_batch_timer.IncompleteBatchTimerOn? && clock.t >= ls.v.replica.proposer.incomplete_batch_timer.when) {
                assert p.msg.v.bal_2a == ls.v.replica.proposer.max_ballot_i_sent_1a == Ballot(1, 1);
                assert p.msg.v.opn_2a == ls.v.replica.proposer.next_operation_number_to_propose == opn;
            } else if |ls.v.replica.proposer.request_queue| > 0 && ls.v.replica.proposer.incomplete_batch_timer.IncompleteBatchTimerOff? {
                assert forall io | io in tios :: !io.LIoOpSend?;
                assert p in ts.t_environment.sentPackets;
                assert false;
            } else {
                assert forall io | io in tios :: !io.LIoOpSend?;
                assert p in ts.t_environment.sentPackets;
                assert false;
            }
        }
    }
    assert |sent_packets| > 0;
    var m := RslMessage_2a(Ballot(1, 1), opn, v);
    assert sent_packets[0] == LPacket(r.proposer.constants.all.config.replica_ids[0], r.proposer.constants.all.config.replica_ids[1], m);
    var pkt_witness := sent_packets[0];
    assert LIoOpSend(pkt_witness) in ios;
    lemma_No2bSentInNonReceiveStep(ts, ts', opn, 1, tios);
    lemma_NoRepliesSentInNonExecutionStep(ts, ts', opn, 1, tios);
}


lemma Before2a_to_Before2a_NonLeaderAction(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires P2Assumption(ts, opn) && RslConsistency(ts)
    requires P2Assumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires idx != 1;
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires Phase2Invariant(ts, opn)
    requires Before_2a_Sent_Invariant(ts, opn);
    ensures Before_2a_Sent_Invariant(ts', opn);
{
    reveal_ExtractSentPacketsFromIos();
    reveal_UntagLIoOpSeq();
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    var nextActionIndex := ls.v.nextActionIndex;
    lemma_NonLeaderDoesNotSend2a(ts, ts', opn, idx, tios);
    lemma_NoNew2bBefore2aSent(ts, ts', opn, idx, tios);
    forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2b?
    ensures BalLeq(pkt.msg.v.bal_2b, Ballot(1, 0)) && pkt.msg.v.opn_2b == opn
    {
        if pkt !in ts.t_environment.sentPackets {
            assert false;
        }
    }
    lemma_NonLeaderDoesNotSendReply(ts, ts', opn, idx, tios);
}


lemma Before2b_to_Before2b_NonReceive(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires P2Assumption(ts, opn) && RslConsistency(ts)
    requires P2Assumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires Phase2Invariant(ts, opn)
    requires Before_2b_Sent_Invariant(ts, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires ts.t_replicas[idx].v.nextActionIndex != 0
    ensures Before_2b_Sent_Invariant(ts', opn)
{   
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    var nextActionIndex, nextActionIndex' := ls.v.nextActionIndex, ls'.v.nextActionIndex;
    Before2b_to_Before2b_NonReceive_NoNewReply(ts, ts', opn, idx, tios);
    lemma_No2bSentInNonReceiveStep(ts, ts', opn, idx, tios);
    if idx != 1 {
        lemma_NonLeaderDoesNotSend2a(ts, ts', opn, idx, tios);
        assert Before_2b_Sent_Invariant(ts', opn);
    } else {
        assert ts'.t_replicas[1].v.replica.learner.unexecuted_learner_state == map[];
        assert forall opn' | opn' in ts.t_replicas[1].v.replica.learner.unexecuted_learner_state
        ::  |ts.t_replicas[1].v.replica.learner.unexecuted_learner_state[opn].received_2b_message_senders| < LMinQuorumSize(ts.t_replicas[1].v.replica.learner.constants.all.config);
        assert ts'.t_replicas[1].v.replica.executor == ts.t_replicas[1].v.replica.executor;
        assert !LProposerCanNominateUsingOperationNumber(ts.t_replicas[1].v.replica.proposer, ts.t_replicas[1].v.replica.acceptor.log_truncation_point, ts.t_replicas[1].v.replica.proposer.next_operation_number_to_propose);
        forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2a?
        ensures pkt in ts.t_environment.sentPackets
        assert Before_2b_Sent_Invariant(ts', opn);
    }
}


lemma Before2b_to_Before2b_NonReceive_NoNewReply(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires P2Assumption(ts, opn) && RslConsistency(ts)
    requires P2Assumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires Phase2Invariant(ts, opn)
    requires Before_2b_Sent_Invariant(ts, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires ts.t_replicas[idx].v.nextActionIndex != 0
    ensures forall pkt | pkt in ts'.t_environment.sentPackets :: !IsNewReplyPacket(ts', pkt) 
{   
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    var nextActionIndex, nextActionIndex' := ls.v.nextActionIndex, ls'.v.nextActionIndex;
    forall pkt | pkt in ts'.t_environment.sentPackets
    ensures !IsNewReplyPacket(ts', pkt) 
    {
        if IsNewReplyPacket(ts', pkt) {
            assert pkt !in ts.t_environment.sentPackets;
            if nextActionIndex == 6 {
                if idx == 1 {
                    assert ls.v.replica.executor.next_op_to_execute == OutstandingOpUnknown();
                    var ios := UntagLIoOpSeq(tios);
                    var sent_packets := ExtractSentPacketsFromIos(ios);
                    assert sent_packets == [];
                    assert forall io | io in tios :: !io.LIoOpSend?;
                } else {
                    lemma_NonLeaderDoesNotSendReply(ts, ts', opn, idx, tios);
                }
                assert false;
            } else {
                lemma_NoRepliesSentInNonExecutionStep(ts, ts', opn, idx, tios);
                assert false;
            }
        }
    }
}


lemma Before2b_to_Before2b_Receive(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires P2Assumption(ts, opn) && RslConsistency(ts)
    requires P2Assumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires Phase2Invariant(ts, opn)
    requires Before_2b_Sent_Invariant(ts, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires ts.t_replicas[idx].v.nextActionIndex == 0
    requires !tios[0].LIoOpTimeoutReceive?
    requires !tios[0].r.msg.v.RslMessage_2a?
    ensures Before_2b_Sent_Invariant(ts', opn)
{
    var ios := UntagLIoOpSeq(tios);
    var sent_packets := ExtractSentPacketsFromIos(ios);
    assert forall p | p in sent_packets :: !p.msg.RslMessage_2b?;
    forall pkt | pkt in ts'.t_environment.sentPackets && IsNew2bPacket(pkt, opn) 
    ensures pkt in ts.t_environment.sentPackets {
        if pkt !in ts.t_environment.sentPackets {
            assert UntagLPacket(pkt) in sent_packets;
            assert false;
        }
    }
    lemma_No2aSentInNon2aStep(ts, ts', opn, idx, tios);
    assert opn == ts'.t_replicas[1].v.replica.executor.ops_complete;   // Can be changed by LExecutorProcessAppStateSupply packet
    assert BalLt(ts'.t_replicas[1].v.replica.learner.max_ballot_seen, Ballot(1, 1)); 
    lemma_NoRepliesSentInNonExecutionStep(ts, ts', opn, idx, tios);
    lemma_No2aSentInNon2aStep(ts, ts', opn, idx, tios);
    assert Before_2b_Sent_Invariant(ts', opn);
}

}