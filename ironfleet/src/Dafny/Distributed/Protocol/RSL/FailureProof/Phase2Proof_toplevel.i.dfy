include "Phase2Proof.i.dfy"
include "Phase2Proof_helper.i.dfy"

module Rs2Phase2Proof_Top {
import opened RslPhase2Proof_postFail_i
import opened Rs2Phase2Proof_Helper


/**** MAIN INVARIANT ****/
lemma PerfInvariantMaintained(s:TimestampedRslState, s':TimestampedRslState, req_time:Timestamp, opn:OperationNumber)
    requires RslAssumption(s) && RslAssumption(s')
    requires RslConsistency(s) && RslConsistency(s')
    requires TimestampedRslNext(s, s')
    requires RslPerfInvariant(s, opn)
    ensures RslPerfInvariant(s', opn)
{   
    PacketsBallotInvariant_Maintained(s, s', opn);
    AlwaysInvariant_Maintained(s, s', opn);
    
    if (!exists pkt :: pkt in s.t_environment.sentPackets && IsNew2aPacket(pkt, opn))
    && (!exists pkt :: pkt in s.t_environment.sentPackets && IsNew2bPacket(pkt, opn)) 
    {
        Before2a_to_MaybeBefore2b(s, s', opn);
    } else if (exists pkt :: pkt in s.t_environment.sentPackets && IsNew2aPacket(pkt, opn))
        && (!exists pkt :: pkt in s.t_environment.sentPackets && IsNew2bPacket(pkt, opn))  
    {
        Before2b_to_MaybeAfter2b(s, s', opn);
    }

    assert AlwaysInvariant(s');        
    assert PacketsBallotInvariant(s');
    
    assume false;
    assert PerformanceGuarantee(s', opn);   // TODO
    assert RslPerfInvariant(s', opn);
}


/*****************************************************************************************
*                                         Lemmas                                         *
*****************************************************************************************/

/****  MAIN THEOREM ****
* Proof that PacketsBallotInvariant is maintained */
lemma PacketsBallotInvariant_Maintained(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber) 
    requires RslAssumption(ts) && RslAssumption(ts')
    requires RslConsistency(ts') && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires RslPerfInvariant(ts, opn)
    ensures PacketsBallotInvariant(ts')
{
    if TimestampedRslNextEnvironment(ts, ts') {
        assert PacketsBallotInvariant(ts');
        return;
    }
    var idx, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>> :| TimestampedRslNextOneReplica(ts, ts', idx, tios);
    var nextActionIndex := ts.t_replicas[idx].v.nextActionIndex;
    if nextActionIndex == 0 {
        PacketsBallotInvariant_ReceiveStep(ts, ts', opn, idx, tios);
    } else {
        PacketsBallotInvariant_NoReceiveStep(ts, ts', opn, idx, tios);
    }
}


/* Proof that a Before_2a_Sent state transitions to a Before_2a_Sent state or 
* Before_2b_Sent state */
lemma Before2a_to_MaybeBefore2b(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber) 
    requires RslAssumption(ts) && RslConsistency(ts)
    requires RslAssumption(ts') && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires RslPerfInvariant(ts, opn)
    requires Before_2a_Sent_Invariant(ts, opn);
    ensures Before_2a_Sent_Invariant(ts', opn) || Before_2b_Sent_Invariant(ts', opn);
{
    if TimestampedRslNextEnvironment(ts, ts') {
        assert Before_2a_Sent_Invariant(ts', opn);
        return;
    }
    var idx, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>> :| TimestampedRslNextOneReplica(ts, ts', idx, tios);
    var nextActionIndex := ts.t_replicas[idx].v.nextActionIndex;
    if idx != 1 {
        Before2a_to_Before2a_NonLeaderAction(ts, ts', opn, idx, tios);
        return;
    }
    Before2a_to_Before2b(ts, ts', opn, tios);
}


/* Proof that a Before_2b_Sent state transitions to a Before_2b_Sent state or 
* After_2b_Sent state */
lemma Before2b_to_MaybeAfter2b(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber) 
    requires RslAssumption(ts) && RslConsistency(ts)
    requires RslAssumption(ts') && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires RslPerfInvariant(ts, opn)
    requires Before_2b_Sent_Invariant(ts, opn);
    ensures Before_2b_Sent_Invariant(ts', opn) || After_2b_Sent_Invariant(ts', opn);
{
    if TimestampedRslNextEnvironment(ts, ts') {
        assert Before_2b_Sent_Invariant(ts', opn);
        return;
    }
    var idx, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>> :| TimestampedRslNextOneReplica(ts, ts', idx, tios);
    var nextActionIndex := ts.t_replicas[idx].v.nextActionIndex;
    if nextActionIndex != 0 {
        Before2b_to_Before2b_NonReceive(ts, ts', opn, idx, tios);
        return;
    } 
    var s, s', ios := UntimestampRslState(ts), UntimestampRslState(ts'), UntagLIoOpSeq(tios);
    var r, r' := s.replicas[idx].replica, s'.replicas[idx].replica;
    if ios[0].LIoOpTimeoutReceive? {
        assert ts'.t_environment.sentPackets == ts.t_environment.sentPackets;
        assert Before_2b_Sent_Invariant(ts', opn);
        return;
    }
    // assert LReplicaNextProcessPacketWithoutReadingClock(r, r', ios);
    var sent_packets := ExtractSentPacketsFromIos(ios);
    if !ios[0].r.msg.RslMessage_2a? {
        Before2b_to_Before2b_Receive(ts, ts', opn, idx, tios);
        return;
    }
    // assert LReplicaNextProcess2a(r, r', ios[0].r, sent_packets);
    var m := ios[0].r.msg;
    if ios[0].r.src in r.acceptor.constants.all.config.replica_ids
       && BalLeq(r.acceptor.max_bal, m.bal_2a)
       && LeqUpperBound(m.opn_2a, r.acceptor.constants.all.params.max_integer_val)
    {
        if m.bal_2a == Ballot(1, 1) {
            Before2b_to_After2b(ts, ts', opn, idx, tios, s, s', ios);
        } else {
            assert forall p | p in sent_packets && p.msg.RslMessage_2b? :: p.msg.bal_2b != Ballot(1, 1);
            assert Before_2b_Sent_Invariant(ts', opn);
        }
    } else {
        assert sent_packets == [];
        reveal_ExtractSentPacketsFromIos();
        assert forall io | io in ios :: !io.LIoOpSend?;
        forall tio | tio in tios 
        ensures !tio.LIoOpSend? {}
        assert Before_2b_Sent_Invariant(ts', opn) || After_2b_Sent_Invariant(ts', opn);
    }
}
}