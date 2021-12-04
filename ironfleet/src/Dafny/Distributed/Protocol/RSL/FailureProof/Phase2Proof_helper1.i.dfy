include "Phase2Proof.i.dfy"

module Rs2Phase2Proof_Helper_1 {
import opened RslPhase2Proof_postFail_i

/* WARNING: this file a timeout of 50s to verify */

lemma PacketsBallotInvariant_ReceiveStep(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts, opn) && RslConsistency(ts)
    requires RslAssumption(ts', opn) && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires RslPerfInvariant(ts, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires ts.t_replicas[idx].v.nextActionIndex == 0 
    ensures PacketsBallotInvariant(ts')
{}

lemma PacketsBallotInvariant_NoReceiveStep(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts, opn) && RslConsistency(ts)
    requires RslAssumption(ts', opn) && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires RslPerfInvariant(ts, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires ts.t_replicas[idx].v.nextActionIndex != 0 
    ensures PacketsBallotInvariant(ts')
{}

lemma AlwaysInvariant_Maintained(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber) 
    requires RslAssumption(ts, opn) && RslAssumption(ts', opn)
    requires RslConsistency(ts) && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires RslPerfInvariant(ts, opn)
    ensures AlwaysInvariant(ts', opn)
{
    assert ts'.t_replicas[1].v.replica.proposer.request_queue == [];
    assert forall pkt | pkt in ts'.t_replicas[1].v.replica.proposer.received_1b_packets && pkt.msg.RslMessage_1b? :: forall op | op in pkt.msg.votes :: RequestBatchSrcInClientIds(ts', pkt.msg.votes[op].max_val);
    forall pkt | pkt in ts'.undeliveredPackets && pkt.msg.v.RslMessage_2a? 
    ensures RequestBatchSrcInClientIds(ts', pkt.msg.v.val_2a)
    {}
    forall pkt | pkt in ts'.undeliveredPackets && pkt.msg.v.RslMessage_2b? 
    ensures RequestBatchSrcInClientIds(ts', pkt.msg.v.val_2b)
    {}
}


lemma Before2a_to_Before2b(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts, opn) && RslConsistency(ts)
    requires RslAssumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', 1, tios);
    requires RslPerfInvariant(ts, opn)
    requires Before_2a_Sent_Invariant(ts, opn);
    ensures Before_2b_Sent_Invariant(ts', opn);
{
    var s, s', ios := UntimestampRslState(ts), UntimestampRslState(ts'), UntagLIoOpSeq(tios);
    var r, r' := s.replicas[1].replica, s'.replicas[1].replica;
    var sent_packets := ExtractSentPacketsFromIos(ios);
    var v :| 
        (&& LValIsHighestNumberedProposal(v, r.proposer.received_1b_packets, opn)
         && LBroadcastToEveryone(r.proposer.constants.all.config, r.proposer.constants.my_index, RslMessage_2a(r.proposer.max_ballot_i_sent_1a, opn, v), sent_packets)
    );
    assert |sent_packets| > 0;
    var m := RslMessage_2a(Ballot(1, 1), opn, v);
    assert sent_packets[0] == LPacket(r.proposer.constants.all.config.replica_ids[0], r.proposer.constants.all.config.replica_ids[1], m);
    var pkt_witness := sent_packets[0];
    assert LIoOpSend(pkt_witness) in ios;
    forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2b?
    ensures pkt in ts.t_environment.sentPackets
    {}
}

lemma Before2a_to_Before2a_NonLeaderAction(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts, opn) && RslConsistency(ts)
    requires RslAssumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires idx != 1;
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires RslPerfInvariant(ts, opn)
    requires Before_2a_Sent_Invariant(ts, opn);
    ensures Before_2a_Sent_Invariant(ts', opn);
{
    reveal_ExtractSentPacketsFromIos();
    reveal_UntagLIoOpSeq();
    forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2a?
    ensures pkt in ts.t_environment.sentPackets
    {}
    forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2b?
    ensures BalLeq(pkt.msg.v.bal_2b, Ballot(1, 0)) && pkt.msg.v.opn_2b == opn
    {
        if pkt !in ts.t_environment.sentPackets {
            assert forall pkt | pkt in ts.t_environment.sentPackets && pkt.msg.v.RslMessage_2a?
                :: BalLeq(pkt.msg.v.bal_2a, Ballot(1, 0)) && pkt.msg.v.opn_2a == opn;
        }
    }
    
    forall pkt | pkt in ts'.t_environment.sentPackets
    ensures !IsNewReplyPacket(ts', pkt) 
    {
        if pkt !in ts.t_environment.sentPackets {
            assert pkt.src == ts.constants.config.replica_ids[idx];
            assert ReplicasDistinct(ts.constants.config.replica_ids, 1, idx);
        }
    }
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
    forall pkt | pkt in ts'.t_environment.sentPackets
    ensures !IsNewReplyPacket(ts', pkt) 
    {
        if pkt !in ts.t_environment.sentPackets {
            assert pkt.src == ts.constants.config.replica_ids[idx];
            assert ReplicasDistinct(ts.constants.config.replica_ids, 1, idx);
        }
    }
    if idx != 1 {
        forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2a?
        ensures pkt in ts.t_environment.sentPackets
        {}
        forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2b?
        ensures pkt in ts.t_environment.sentPackets
        {}
        assert Before_2b_Sent_Invariant(ts', opn);
    } else {
        assert forall opn' | opn' in ts.t_replicas[1].v.replica.learner.unexecuted_learner_state
        ::  |ts.t_replicas[1].v.replica.learner.unexecuted_learner_state[opn].received_2b_message_senders| < LMinQuorumSize(ts.t_replicas[1].v.replica.learner.constants.all.config);
        assert ts'.t_replicas[1].v.replica.executor == ts.t_replicas[1].v.replica.executor;

        assert !LProposerCanNominateUsingOperationNumber(ts.t_replicas[1].v.replica.proposer, ts.t_replicas[1].v.replica.acceptor.log_truncation_point, ts.t_replicas[1].v.replica.proposer.next_operation_number_to_propose);
        forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2a?
        ensures pkt in ts.t_environment.sentPackets
        forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2b?
        ensures pkt in ts.t_environment.sentPackets
        {}
    }
}

lemma Before2b_to_Before2b_Receive(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
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
    forall pkt | pkt in ts'.t_environment.sentPackets && IsNew2aPacket(pkt, opn) 
    ensures pkt in ts.t_environment.sentPackets {}
    assert opn == ts'.t_replicas[1].v.replica.executor.ops_complete;   // Can be changed by LExecutorProcessAppStateSupply packet
    assert BalLt(ts'.t_replicas[1].v.replica.learner.max_ballot_seen, Ballot(1, 1)); 
    forall pkt | pkt in ts'.t_environment.sentPackets && IsNewReplyPacket(ts', pkt) 
    ensures pkt in ts.t_environment.sentPackets
    {}
    assert Before_2b_Sent_Invariant(ts', opn);
}


}