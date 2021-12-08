include "Phase2Proof.i.dfy"
include "Phase2Proof_helper0.i.dfy"
include "Phase2Proof_helper1.i.dfy"
// include "Phase2Proof_helper2.i.dfy"

module Rs2Phase2Proof_PreFail_Top {
import opened RslPhase2Proof_PreFail_i
import opened Rs2Phase2Proof_PreFail_Helper0
import opened Rs2Phase2Proof_PreFail_Helper1
// import opened Rs2Phase2Proof_PostFail_Helper2


/**** MAIN INVARIANT THEOREM ****/
lemma PerfInvariantMaintained(s:TimestampedRslState, s':TimestampedRslState, req_time:Timestamp, opn:OperationNumber)
    requires RslAssumption(s, opn) && RslAssumption(s', opn)
    requires RslConsistency(s) && RslConsistency(s')
    requires TimestampedRslNext(s, s')
    requires RslPerfInvariant(s, opn)
    ensures RslPerfInvariant(s', opn)
{   
    PacketsBallotInvariant_Maintained(s, s', opn);
    AlwaysInvariant_Maintained(s, s', opn);
    
    if  && (!exists pkt :: pkt in s.t_environment.sentPackets && IsPreFail2aPacket(pkt, opn))
        && (!exists pkt :: pkt in s.t_environment.sentPackets && IsPreFail2bPacket(pkt, opn)) 
    {
        Before2a_to_MaybeBefore2b(s, s', opn);
    } else if (exists pkt :: pkt in s.t_environment.sentPackets && IsPreFail2aPacket(pkt, opn))
        && (!exists pkt :: pkt in s.t_environment.sentPackets && IsPreFail2bPacket(pkt, opn))  
    {
        Before2b_to_MaybeAfter2b(s, s', opn);
    } else {
        // TODO
        assume false;
        // assert After_2b_Sent_Invariant(s, opn);
        // After2b_to_After2b(s, s', opn);
    }
    assert RslPerfInvariant(s', opn);
}


/*****************************************************************************************
*                                         Lemmas                                         *
*****************************************************************************************/


/* Proof that PacketsBallotInvariant is maintained */
lemma PacketsBallotInvariant_Maintained(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber) 
    requires RslAssumption(ts, opn) && RslAssumption(ts', opn)
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
    requires RslAssumption(ts, opn) && RslConsistency(ts)
    requires RslAssumption(ts', opn) && RslConsistency(ts')
    requires AlwaysInvariant(ts', opn)
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires RslPerfInvariant(ts, opn)
    requires Before_2a_Sent_Invariant(ts, opn)
    ensures Before_2a_Sent_Invariant(ts', opn) || Before_2b_Sent_Invariant(ts', opn)
{
    if TimestampedRslNextEnvironment(ts, ts') {
        assert Before_2a_Sent_Invariant(ts', opn);
        return;
    }
    var idx, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>> :| TimestampedRslNextOneReplica(ts, ts', idx, tios);
    var nextActionIndex := ts.t_replicas[idx].v.nextActionIndex;
    if idx != 0 {
        Before2a_to_Before2a_NonLeaderAction(ts, ts', opn, idx, tios);
    } else {
        var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
        var nextActionIndex := ls.v.nextActionIndex;
        if nextActionIndex == 3 {
            Before2a_to_Before2b(ts, ts', opn, tios);
        } else {
            Before2a_to_Before2a_LeaderAction(ts, ts', opn, tios);
        }
    }
}


/* Proof that a Before_2b_Sent state transitions to a Before_2b_Sent state or 
* After_2b_Sent state */
lemma Before2b_to_MaybeAfter2b(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber) 
    requires RslAssumption(ts, opn) && RslConsistency(ts)
    requires RslAssumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires AlwaysInvariant(ts', opn)
    requires TimestampedRslNext(ts, ts')
    requires RslPerfInvariant(ts, opn)
    requires Before_2b_Sent_Invariant(ts, opn)
    ensures Before_2b_Sent_Invariant(ts', opn) || After_2b_Sent_Invariant(ts', opn)
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
    // From this point on, nextActionIndex == 0
    var s, s', ios := UntimestampRslState(ts), UntimestampRslState(ts'), UntagLIoOpSeq(tios);
    var r, r' := s.replicas[idx].replica, s'.replicas[idx].replica;
    if ios[0].LIoOpTimeoutReceive? {
        assert ts'.t_environment.sentPackets == ts.t_environment.sentPackets;
        assert Before_2b_Sent_Invariant(ts', opn);
        return;
    }
    var sent_packets := ExtractSentPacketsFromIos(ios);
    if !ios[0].r.msg.RslMessage_2a? {

        // TODO
        assume false;
        // Before2b_to_Before2b_Receive(ts, ts', opn, idx, tios);
        return;
    }

    // TODO
    assume false;
    // From this point on, replica idx is processing a 2a packet
    // Before2b_to_MaybeAfter2b_Process2a(ts, ts', opn, idx, tios);
}


// /* Proof that a After_2b_Sent state transitions to a After_2b_Sent state */
// lemma After2b_to_After2b(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber) 
//     requires RslAssumption(ts, opn) && RslConsistency(ts)
//     requires RslAssumption(ts', opn) && RslConsistency(ts')
//     requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
//     requires AlwaysInvariant(ts', opn)
//     requires TimestampedRslNext(ts, ts')
//     requires RslPerfInvariant(ts, opn)
//     requires After_2b_Sent_Invariant(ts, opn)
//     ensures After_2b_Sent_Invariant(ts', opn)
// {
//     if TimestampedRslNextEnvironment(ts, ts') {
//         assert After_2b_Sent_Invariant(ts', opn);
//         return;
//     }
//     var idx, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>> :| TimestampedRslNextOneReplica(ts, ts', idx, tios);
//     if idx == 1 {
//         After2b_to_After2b_LeaderAction(ts, ts', opn, idx, tios);
//     } else {
//         After2b_to_After2b_NonLeaderAction(ts, ts', opn, idx, tios);
//     }
// }
}