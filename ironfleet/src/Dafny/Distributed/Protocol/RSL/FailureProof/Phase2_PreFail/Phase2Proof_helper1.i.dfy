include "Phase2Proof.i.dfy"
include "GenericLemmas.i.dfy"

module Rs2Phase2Proof_PreFail_Helper1 {
import opened RslPhase2Proof_PreFail_i
import opened Rs2Phase2Proof_PreFail_Generic

/* WARNING: this file a timeout of 60s to verify */

lemma PacketsBallotInvariant_ReceiveStep(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts, opn) && RslConsistency(ts)
    requires RslAssumption(ts', opn) && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires RslPerfInvariant(ts, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires ts.t_replicas[idx].v.nextActionIndex == 0 
    ensures PacketsBallotInvariant(ts')
{
    forall pkt | pkt in ts'.undeliveredPackets 
    ensures ExistingPacketsBallot(pkt)
    {}
}

lemma PacketsBallotInvariant_NoReceiveStep(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts, opn) && RslConsistency(ts)
    requires RslAssumption(ts', opn) && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires RslPerfInvariant(ts, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires ts.t_replicas[idx].v.nextActionIndex != 0 
    ensures PacketsBallotInvariant(ts')
{
    forall pkt | pkt in ts'.undeliveredPackets 
    ensures ExistingPacketsBallot(pkt)
    {}
}

lemma Before2a_to_Before2b(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts, opn) && RslConsistency(ts)
    requires RslAssumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', 0, tios);
    requires RslPerfInvariant(ts, opn)
    requires ts.t_replicas[0].v.nextActionIndex == 3
    requires Before_2a_Sent_Invariant(ts, opn);
    ensures Before_2b_Sent_Invariant(ts', opn);
{
    var ls, ls' := ts.t_replicas[0], ts'.t_replicas[0];
    var nextActionIndex := ls.v.nextActionIndex;

    forall p | p in ts'.t_environment.sentPackets && p.msg.v.RslMessage_2a?
    ensures && p.msg.v.bal_2a == Ballot(1, 0)
            && p.msg.v.opn_2a == opn
    {
        if p !in ts.t_environment.sentPackets {
            var clock := SpontaneousClock(UntagLIoOpSeq(tios)).t;
            if !LProposerCanNominateUsingOperationNumber(ls.v.replica.proposer, ls.v.replica.acceptor.log_truncation_point, ls.v.replica.proposer.next_operation_number_to_propose) {
                assert p in ts.t_environment.sentPackets;
                assert false;
            } else if !LAllAcceptorsHadNoProposal(ls.v.replica.proposer.received_1b_packets, ls.v.replica.proposer.next_operation_number_to_propose){
                assert p.msg.v.bal_2a == ls.v.replica.proposer.max_ballot_i_sent_1a == Ballot(1, 0);
                assert p.msg.v.opn_2a == ls.v.replica.proposer.next_operation_number_to_propose == opn;
            } else if (exists opn' :: opn' > ls.v.replica.proposer.next_operation_number_to_propose && !LAllAcceptorsHadNoProposal(ls.v.replica.proposer.received_1b_packets, opn'))
                        || |ls.v.replica.proposer.request_queue| >= ls.v.replica.proposer.constants.all.params.max_batch_size
                        || (|ls.v.replica.proposer.request_queue| > 0 && ls.v.replica.proposer.incomplete_batch_timer.IncompleteBatchTimerOn? && clock >= ls.v.replica.proposer.incomplete_batch_timer.when) {
                assert p.msg.v.bal_2a == ls.v.replica.proposer.max_ballot_i_sent_1a == Ballot(1, 0);
                assert p.msg.v.opn_2a == ls.v.replica.proposer.next_operation_number_to_propose == opn;
            } else if |ls.v.replica.proposer.request_queue| > 0 && ls.v.replica.proposer.incomplete_batch_timer.IncompleteBatchTimerOff? {
                assert p in ts.t_environment.sentPackets;
                assert false;
            } else {
                assert p in ts.t_environment.sentPackets;
                assert false;
            }
        } else {
            assert false;
        }
    }

    var sent_packets := ExtractSentPacketsFromIos(UntagLIoOpSeq(tios));
    assert LProposerCanNominateUsingOperationNumber(ls.v.replica.proposer, ls.v.replica.acceptor.log_truncation_point, ls.v.replica.proposer.next_operation_number_to_propose);
    assert LAllAcceptorsHadNoProposal(ls.v.replica.proposer.received_1b_packets, ls.v.replica.proposer.next_operation_number_to_propose);
    assert (exists opn' :: opn' > ls.v.replica.proposer.next_operation_number_to_propose && !LAllAcceptorsHadNoProposal(ls.v.replica.proposer.received_1b_packets, opn'));
    assert LProposerCanNominateUsingOperationNumber(ls.v.replica.proposer, ls.v.replica.acceptor.log_truncation_point, ls.v.replica.proposer.next_operation_number_to_propose);
    assert |sent_packets| > 0;
    var v := ls.v.replica.proposer.request_queue[..1];
    var m := RslMessage_2a(Ballot(1, 0), opn, v);
    assert sent_packets[0] == LPacket(ls.v.replica.proposer.constants.all.config.replica_ids[0], ls.v.replica.proposer.constants.all.config.replica_ids[0], m);
    var pkt_witness := sent_packets[0];
    assert LIoOpSend(pkt_witness) in UntagLIoOpSeq(tios);
    assert (exists pkt :: pkt in ts'.t_environment.sentPackets && IsPreFail2aPacket(pkt, opn));
}

lemma Before2a_to_Before2a_NonLeaderAction(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts, opn) && RslConsistency(ts)
    requires RslAssumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires idx != 0;
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires RslPerfInvariant(ts, opn)
    requires Before_2a_Sent_Invariant(ts, opn);
    ensures Before_2a_Sent_Invariant(ts', opn);
{
    assert ts'.t_replicas[0] == ts.t_replicas[0];
    lemma_NonLeaderDoesNotSend2a(ts, ts', opn, idx, tios);
    lemma_No2bBefore2aSent(ts, ts', opn, idx, tios);
    lemma_NonLeaderDoesNotSendReply(ts, ts', opn, idx, tios);
}

lemma Before2a_to_Before2a_LeaderAction(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts, opn) && RslConsistency(ts)
    requires RslAssumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', 0, tios);
    requires RslPerfInvariant(ts, opn)
    requires Before_2a_Sent_Invariant(ts, opn);
    ensures Before_2a_Sent_Invariant(ts', opn);
{
    lemma_No2bBefore2aSent(ts, ts', opn, 0, tios);
}


lemma Before2b_to_Before2b_NonReceive(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts, opn) && RslConsistency(ts)
    requires RslAssumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires RslPerfInvariant(ts, opn)
    requires Before_2b_Sent_Invariant(ts, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires ts.t_replicas[idx].v.nextActionIndex != 0
    ensures Before_2b_Sent_Invariant(ts', opn)
{   
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    var nextActionIndex, nextActionIndex' := ls.v.nextActionIndex, ls'.v.nextActionIndex;

    Before2b_to_Before2b_NonReceive_NoReply(ts, ts', opn, idx, tios);

    if idx != 0 {
        lemma_NonLeaderDoesNotSend2a(ts, ts', opn, idx, tios);
        lemma_No2bSentInNonReceiveStep(ts, ts', opn, idx, tios);
        assert Before_2b_Sent_Invariant(ts', opn);
    } else {
        assert ts'.t_replicas[0].v.replica.learner.unexecuted_learner_state == map[];
        assert forall opn' | opn' in ts.t_replicas[0].v.replica.learner.unexecuted_learner_state
        ::  |ts.t_replicas[0].v.replica.learner.unexecuted_learner_state[opn].received_2b_message_senders| < LMinQuorumSize(ts.t_replicas[0].v.replica.learner.constants.all.config);
        assert ts'.t_replicas[0].v.replica.executor == ts.t_replicas[0].v.replica.executor;

        assert !LProposerCanNominateUsingOperationNumber(ts.t_replicas[0].v.replica.proposer, ts.t_replicas[0].v.replica.acceptor.log_truncation_point, ts.t_replicas[0].v.replica.proposer.next_operation_number_to_propose);
        forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2a?
        ensures pkt in ts.t_environment.sentPackets
        lemma_No2bSentInNonReceiveStep(ts, ts', opn, idx, tios);
        assert Before_2b_Sent_Invariant(ts', opn);
    }
}


lemma Before2b_to_Before2b_NonReceive_NoReply(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts, opn) && RslConsistency(ts)
    requires RslAssumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires RslPerfInvariant(ts, opn)
    requires Before_2b_Sent_Invariant(ts, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires ts.t_replicas[idx].v.nextActionIndex != 0
    ensures forall pkt | pkt in ts'.t_environment.sentPackets :: !IsPreFailReplyPacket(ts', pkt) 
{   
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    var nextActionIndex, nextActionIndex' := ls.v.nextActionIndex, ls'.v.nextActionIndex;
    forall pkt | pkt in ts'.t_environment.sentPackets
    ensures !IsPreFailReplyPacket(ts', pkt) 
    {
        if IsPreFailReplyPacket(ts', pkt) {
            assert pkt !in ts.t_environment.sentPackets;
            if nextActionIndex == 6 {
                if idx == 0 {
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
                var ios := UntagLIoOpSeq(tios);
                var sent_packets := ExtractSentPacketsFromIos(ios);
                assert nextActionIndex != 0;
                assert forall p | p in sent_packets :: !p.msg.RslMessage_Reply?;
                forall io | io in tios && io.LIoOpSend?
                ensures !io.s.msg.v.RslMessage_Reply? {}
                assert false;
            }
        }
    }
}


lemma Before2b_to_Before2b_ReceiveNot2a(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts, opn) && RslConsistency(ts)
    requires RslAssumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires RslPerfInvariant(ts, opn)
    requires Before_2b_Sent_Invariant(ts, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires ts.t_replicas[idx].v.nextActionIndex == 0
    requires !tios[0].LIoOpTimeoutReceive?
    requires !tios[0].r.msg.v.RslMessage_2a?
    ensures Before_2b_Sent_Invariant(ts', opn)
{
    reveal_ExtractSentPacketsFromIos();
    reveal_UntagLIoOpSeq();
    var ios := UntagLIoOpSeq(tios);
    var sent_packets := ExtractSentPacketsFromIos(ios);
    assert forall p | p in sent_packets :: !p.msg.RslMessage_2b?;
    forall pkt | pkt in ts'.t_environment.sentPackets && IsPreFail2bPacket(pkt, opn) 
    ensures pkt in ts.t_environment.sentPackets {
        if pkt !in ts.t_environment.sentPackets {
            assert UntagLPacket(pkt) in sent_packets;
            assert false;
        }
    }
    lemma_No2aSentInReceiveStep(ts, ts', opn, idx, tios);
    lemma_NoRepliesSentInNonExecutionStep(ts, ts', opn, idx, tios);

    assert opn == ts'.t_replicas[0].v.replica.executor.ops_complete;
    assert Before_2b_Sent_Invariant(ts', opn);
}


}