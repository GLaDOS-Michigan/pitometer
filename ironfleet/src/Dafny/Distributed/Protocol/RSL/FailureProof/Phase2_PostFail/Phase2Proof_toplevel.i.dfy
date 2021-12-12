include "Phase2Proof.i.dfy"
include "Phase2Proof_helper0.i.dfy"
include "Phase2Proof_helper1.i.dfy"
include "Phase2Proof_helper2.i.dfy"

module RslPhase2Proof_PostFail_Top {
import opened RslPhase2Proof_PostFail_i
import opened Rs2Phase2Proof_PostFail_Helper0
import opened Rs2Phase2Proof_PostFail_Helper1
import opened Rs2Phase2Proof_PostFail_Helper2


/**** MAIN INVARIANT THEOREM ****/
lemma PerfInvariantMaintained(s:TimestampedRslState, s':TimestampedRslState, opn:OperationNumber)
    requires CommonAssumptions(s) && CommonAssumptions(s')
    requires InPhase2(s)
    requires InPhase2(s) ==> P2Assumption(s, opn)
    requires InPhase2(s') ==> P2Assumption(s', opn)
    requires TimestampedRslNext(s, s')
    requires Phase2Invariant(s, opn)
    ensures Phase2Invariant(s', opn)
    ensures InPhase2(s')
{   
    PacketsBallotInvariant_Maintained(s, s', opn);
    AlwaysInvariantP2_Maintained(s, s', opn);
    
    if  && (!exists pkt :: pkt in s.t_environment.sentPackets && IsNew2aPacket(pkt, opn))
        && (!exists pkt :: pkt in s.t_environment.sentPackets && IsNew2bPacket(pkt, opn)) 
    {
        Before2a_to_MaybeBefore2b(s, s', opn);
    } else if (exists pkt :: pkt in s.t_environment.sentPackets && IsNew2aPacket(pkt, opn))
        && (!exists pkt :: pkt in s.t_environment.sentPackets && IsNew2bPacket(pkt, opn))  
    {
        Before2b_to_MaybeAfter2b(s, s', opn);
    } else {
        assert After_2b_Sent_Invariant(s, opn);
        After2b_to_After2b(s, s', opn);
    }
    assert Phase2Invariant(s', opn);
}


lemma Phase2TopLevel(tglb:seq<TimestampedRslState>, opn:OperationNumber)
    requires |tglb| > 0
    requires forall i | 0 <= i < |tglb| :: CommonAssumptions(tglb[i])
    requires forall i | 0 < i < |tglb| :: TimestampedRslNext(tglb[i - 1], tglb[i])
    requires InPhase2(tglb[0])
    requires Phase2Invariant(tglb[0], opn)
    requires forall i | 0 <= i < |tglb| :: InPhase2(tglb[i]) ==> P2Assumption(tglb[i], opn)
    ensures forall j | 0 <= j < |tglb| :: Phase2Invariant(tglb[j], opn)
{
    assert Phase2Invariant(tglb[0], opn);
    var i := 1;
    while i < |tglb| 
        decreases |tglb| - i
        invariant 1 <= i <= |tglb| 
        invariant forall k | 0 <= k < i :: Phase2Invariant(tglb[k], opn)
        invariant forall k | 0 <= k < i :: InPhase2(tglb[k])
    {
        PerfInvariantMaintained(tglb[i-1], tglb[i], opn);
        assert InPhase2(tglb[i]);
        assert Phase2Invariant(tglb[i], opn);
        i := i + 1;
        var k := i - 1;
        forall j | 0 <= j <= k ensures Phase2Invariant(tglb[j], opn);
        forall j | 0 <= j <= k ensures InPhase2(tglb[j]);
    }
}


/*****************************************************************************************
*                                         Lemmas                                         *
*****************************************************************************************/


/* Proof that PacketsBallotInvariant is maintained */
lemma PacketsBallotInvariant_Maintained(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires InPhase2(ts)
    requires InPhase2(ts) ==> P2Assumption(ts, opn)
    requires InPhase2(ts') ==> P2Assumption(ts', opn)
    requires TimestampedRslNext(ts, ts')
    requires Phase2Invariant(ts, opn)
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
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires InPhase2(ts)
    requires InPhase2(ts) ==> P2Assumption(ts, opn)
    requires InPhase2(ts') ==> P2Assumption(ts', opn)
    requires TimestampedRslNext(ts, ts')
    requires Phase2Invariant(ts, opn)
    requires Before_2a_Sent_Invariant(ts, opn)
    ensures Before_2a_Sent_Invariant(ts', opn) || Before_2b_Sent_Invariant(ts', opn)
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
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires InPhase2(ts)
    requires InPhase2(ts) ==> P2Assumption(ts, opn)
    requires InPhase2(ts') ==> P2Assumption(ts', opn)
    requires TimestampedRslNext(ts, ts')
    requires Phase2Invariant(ts, opn)
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
        Before2b_to_Before2b_Receive(ts, ts', opn, idx, tios);
        return;
    }
    // From this point on, replica idx is processing a 2a packet
    Before2b_to_MaybeAfter2b_Process2a(ts, ts', opn, idx, tios);
}


/* Proof that a After_2b_Sent state transitions to a After_2b_Sent state */
lemma After2b_to_After2b(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires InPhase2(ts)
    requires InPhase2(ts) ==> P2Assumption(ts, opn)
    requires InPhase2(ts') ==> P2Assumption(ts', opn)
    requires TimestampedRslNext(ts, ts')
    requires Phase2Invariant(ts, opn)
    requires After_2b_Sent_Invariant(ts, opn)
    ensures After_2b_Sent_Invariant(ts', opn)
{
    if TimestampedRslNextEnvironment(ts, ts') {
        assert After_2b_Sent_Invariant(ts', opn);
        return;
    }
    var idx, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>> :| TimestampedRslNextOneReplica(ts, ts', idx, tios);
    if idx == 1 {
        After2b_to_After2b_LeaderAction(ts, ts', opn, idx, tios);
    } else {
        After2b_to_After2b_NonLeaderAction(ts, ts', opn, idx, tios);
    }
}
}