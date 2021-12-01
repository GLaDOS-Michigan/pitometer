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


lemma Before2a_to_After2a(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts) && RslConsistency(ts)
    requires RslAssumption(ts') && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', 1, tios);
    requires RslPerfInvariant(ts, opn)
    requires Before_2a_Sent_Invariant(ts, opn);
    ensures After_2a_Sent_Invariant(ts', opn);
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
    assert After_2a_Sent_Invariant(ts', opn);
}
}