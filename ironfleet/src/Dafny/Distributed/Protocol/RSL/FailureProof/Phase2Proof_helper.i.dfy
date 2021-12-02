include "Phase2Proof.i.dfy"

module Rs2Phase2Proof_Helper {
import opened RslPhase2Proof_postFail_i


/* Note: this requires a long timeout to verify */
lemma PacketsBallotInvariant_ReceiveStep(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts) && RslConsistency(ts)
    requires RslAssumption(ts') && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires RslPerfInvariant(ts, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires ts.t_replicas[idx].v.nextActionIndex == 0 
    ensures PacketsBallotInvariant(ts')
{}

lemma PacketsBallotInvariant_NoReceiveStep(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts) && RslConsistency(ts)
    requires RslAssumption(ts') && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires RslPerfInvariant(ts, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires ts.t_replicas[idx].v.nextActionIndex != 0 
    ensures PacketsBallotInvariant(ts')
{}

lemma AlwaysInvariant_Maintained(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber) 
    requires RslAssumption(ts) && RslAssumption(ts')
    requires RslConsistency(ts) && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires RslPerfInvariant(ts, opn)
    ensures AlwaysInvariant(ts')
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
    requires RslAssumption(ts) && RslConsistency(ts)
    requires RslAssumption(ts') && RslConsistency(ts')
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
}

lemma Before2a_to_Before2a_NonLeaderAction(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts) && RslConsistency(ts)
    requires RslAssumption(ts') && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires idx != 1;
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires RslPerfInvariant(ts, opn)
    requires Before_2a_Sent_Invariant(ts, opn);
    ensures Before_2a_Sent_Invariant(ts', opn);
{
    forall pkt | pkt in ts'.t_environment.sentPackets && IsNew2aPacket(pkt, opn) 
    ensures pkt in ts.t_environment.sentPackets
    {}
    forall pkt | pkt in ts'.t_environment.sentPackets && IsNew2bPacket(pkt, opn) 
    ensures pkt in ts.t_environment.sentPackets
    {}
}


lemma Before2b_to_After2b(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>,
    rs:RslState, rs':RslState, iops:seq<RslIo>
) 
    requires rs == UntimestampRslState(ts)
    requires rs' == UntimestampRslState(ts')
    requires iops == UntagLIoOpSeq(tios);
    requires RslAssumption(ts) && RslConsistency(ts)
    requires RslAssumption(ts') && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires RslPerfInvariant(ts, opn)
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
    var r, r' := rs.replicas[idx].replica, rs'.replicas[idx].replica;
    var m := iops[0].r.msg;
    var sent_packets := ExtractSentPacketsFromIos(iops);
    assert LAcceptorProcess2a(r.acceptor, r'.acceptor, iops[0].r, sent_packets);
    var msg2b := RslMessage_2b(m.bal_2a, m.opn_2a, m.val_2a);
    assert LBroadcastToEveryone(r.acceptor.constants.all.config, r.acceptor.constants.my_index, msg2b, sent_packets);
    assert forall p | p in sent_packets :: LIoOpSend(p) in iops;
    if m.opn_2a == opn {
        var pkt_witness := sent_packets[0];
        assert LIoOpSend(pkt_witness) in iops;
        assert After_2b_Sent_Invariant(ts', opn);
    } else {
        assert forall p | p in sent_packets && p.msg.RslMessage_2b? :: p.msg.opn_2b != opn;
        assert Before_2b_Sent_Invariant(ts', opn);
    }
}


lemma Before2b_to_Before2b_NonReceive(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts) && RslConsistency(ts)
    requires RslAssumption(ts') && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires RslPerfInvariant(ts, opn)
    requires Before_2b_Sent_Invariant(ts, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires ts.t_replicas[idx].v.nextActionIndex != 0
    ensures Before_2b_Sent_Invariant(ts', opn)
{}

lemma Before2b_to_Before2b_Receive(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts) && RslConsistency(ts)
    requires RslAssumption(ts') && RslConsistency(ts')
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
    var s, s', ios := UntimestampRslState(ts), UntimestampRslState(ts'), UntagLIoOpSeq(tios);
    var r, r' := s.replicas[idx].replica, s'.replicas[idx].replica;
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
    assert Before_2b_Sent_Invariant(ts', opn);
}


}