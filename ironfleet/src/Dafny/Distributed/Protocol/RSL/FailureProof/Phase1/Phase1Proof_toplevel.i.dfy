include "Phase1Proof.i.dfy"
include "../Phase2_PostFail/Phase2Proof.i.dfy"
include "Phase1Proof_helper0.i.dfy"
include "Phase1Proof_helper1.i.dfy"

module RslPhase1Proof_Top {
import opened RslPhase1Proof_i
import P2 = RslPhase2Proof_PostFail_i

import opened RslPhase1Proof_Helper0
import opened RslPhase1Proof_Helper1



/**** MAIN INVARIANT THEOREM ****/
lemma PerfInvariantMaintained(s:TimestampedRslState, s':TimestampedRslState, opn:OperationNumber)
    requires CommonAssumptions(s) && CommonAssumptions(s')
    requires InPhase1(s)
    requires InPhase1(s) ==> P1Assumption(s, opn)
    requires InPhase1(s') ==> P1Assumption(s', opn)
    requires TimestampedRslNext(s, s')
    requires Phase1Invariant(s, opn)
    ensures InPhase1(s') || InPhase2(s')
    ensures InPhase1(s') ==> Phase1Invariant(s', opn)
    ensures InPhase2(s') ==> P2.Phase2Invariant(s', opn)
{   
    PacketsBallotInvariant_Maintained(s, s', opn);

    assume false;
    // AlwaysInvariantP2_Maintained(s, s', opn);
    
    // if  && (!exists pkt :: pkt in s.t_environment.sentPackets && IsNew2aPacket(pkt, opn))
    //     && (!exists pkt :: pkt in s.t_environment.sentPackets && IsNew2bPacket(pkt, opn)) 
    // {
    //     Before2a_to_MaybeBefore2b(s, s', opn);
    // } else if (exists pkt :: pkt in s.t_environment.sentPackets && IsNew2aPacket(pkt, opn))
    //     && (!exists pkt :: pkt in s.t_environment.sentPackets && IsNew2bPacket(pkt, opn))  
    // {
    //     Before2b_to_MaybeAfter2b(s, s', opn);
    // } else {
    //     assert After_2b_Sent_Invariant(s, opn);
    //     After2b_to_After2b(s, s', opn);
    // }
    // assert Phase2Invariant(s', opn);
}


/*
lemma Phase1TopLevel(tglb:seq<TimestampedRslState>, opn:OperationNumber) returns (startPhase2Idx:int)
    requires forall i | 0 <= i < |tglb| :: P1Assumption(tglb[i], opn)
    requires forall i | 0 < i < |tglb| :: TimestampedRslNext(tglb[i - 1], tglb[i])
    // Phase1 lasts up till right before startPhase2Idx
    // startPhase2Idx is the initial state of phase 2
    ensures forall j | 0 <= j < |tglb| && j < startPhase2Idx :: Phase1Invariant(tglb[j], opn)
    ensures startPhase2Idx >= 0
    ensures startPhase2Idx < |tglb| ==> Phase2Begin(tglb[startPhase2Idx], opn)
{
    // TODO:
    assume false;
}

lemma Phase1TopLevel_Prototype(tglb:seq<TimestampedRslState>, opn:OperationNumber) returns (startPhase2Idx:int)
    requires InPhase1(tglb[0])
    requires forall i | 0 <= i < |tglb| :: InPhase1(tglb[i]) ==> P1Assumption(tglb[i])
    requires forall i | 0 < i < |tglb| :: TimestampedRslNext(tglb[i - 1], tglb[i])
    // Phase1 lasts up till right before startPhase2Idx
    // startPhase2Idx is the initial state of phase 2
    ensures forall j | 0 <= j < |tglb| && j < startPhase2Idx :: Phase1Invariant(tglb[j], opn)
    ensures startPhase2Idx >= 0
    ensures startPhase2Idx < |tglb| ==> Phase2Invariant(tglb[startPhase2Idx], opn);
    ensures startPhase2Idx < |tglb| ==> InPhase2(tglb[startPhase2Idx])
{
    // TODO:
    assume false;
}
*/


/*****************************************************************************************
*                                         Lemmas                                         *
*****************************************************************************************/


/* Proof that PacketsBallotInvariant is maintained */
lemma PacketsBallotInvariant_Maintained(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires InPhase1(ts)
    requires InPhase1(ts) ==> P1Assumption(ts, opn)
    requires InPhase1(ts') ==> P1Assumption(ts', opn)
    requires TimestampedRslNext(ts, ts')
    requires Phase1Invariant(ts, opn)
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



}
