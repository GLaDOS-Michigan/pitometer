include "Phase1Proof.i.dfy"
include "../Phase2_PostFail/Phase2Proof.i.dfy"
include "Phase1Proof_helper0.i.dfy"
include "Phase1Proof_helper1.i.dfy"
include "Phase1Proof_helper2.i.dfy"

module RslPhase1Proof_Top {
import opened RslPhase1Proof_i
import P2 = RslPhase2Proof_PostFail_i

import opened RslPhase1Proof_Helper0
import opened RslPhase1Proof_Helper1
import opened RslPhase1Proof_Helper2



/**** MAIN INVARIANT THEOREM ****/
lemma PerfInvariantMaintained(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber)
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires InPhase1(ts)
    requires InPhase1(ts) ==> P1Assumption(ts, opn)
    requires InPhase1(ts') ==> P1Assumption(ts', opn)
    requires TimestampedRslNext(ts, ts')
    requires Phase1Invariant(ts, opn)
    ensures InPhase1(ts') || InPhase2(ts')
    ensures InPhase1(ts') ==> Phase1Invariant(ts', opn)
    ensures InPhase2(ts') ==> P2.Phase2Invariant(ts', opn)
{   
    PacketsBallotInvariant_Maintained(ts, ts', opn);
    AlwaysInvariantP1_Maintained(ts, ts', opn);
    if TimestampedRslNextEnvironment(ts, ts') {
        assert PacketsBallotInvariant(ts');
        return;
    }

    var idx, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>> :| TimestampedRslNextOneReplica(ts, ts', idx, tios);
    var nextActionIndex := ts.t_replicas[idx].v.nextActionIndex;
    if nextActionIndex == 2 {
        Phase1_to_MaybePhase2(ts, ts', opn, idx, tios);
    } else {
        Phase1_to_Phase1(ts, ts', opn, idx, tios);
    }
}

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


/* Phase1 lasts up till right before P2idx.
* Return value P2idx is the initial state of phase 2 */
lemma Phase1TopLevel(tglb:seq<TimestampedRslState>, opn:OperationNumber) returns (P2idx:int)
    requires |tglb| > 0;
    requires forall i | 0 <= i < |tglb| :: CommonAssumptions(tglb[i])
    requires forall i | 0 < i < |tglb| :: TimestampedRslNext(tglb[i - 1], tglb[i])
    requires InPhase1(tglb[0])
    requires Phase1Invariant(tglb[0], opn)
    requires forall i | 0 <= i < |tglb| :: InPhase1(tglb[i]) ==> P1Assumption(tglb[i], opn)
    ensures forall j | 0 <= j < |tglb| && j < P2idx :: Phase1Invariant(tglb[j], opn)
    ensures P2idx >= 0
    ensures P2idx < |tglb| ==> P2.Phase2Invariant(tglb[P2idx], opn);
    ensures P2idx < |tglb| ==> InPhase2(tglb[P2idx])
{
    P2idx := 1;
    while P2idx < |tglb|
        decreases |tglb| - P2idx
        invariant 0 <= P2idx <= |tglb| 
        invariant forall k | 0 <= k < P2idx :: Phase1Invariant(tglb[k], opn)
        invariant forall k | 0 <= k < P2idx :: InPhase1(tglb[k])
    {
        PerfInvariantMaintained(tglb[P2idx - 1], tglb[P2idx], opn);
        if InPhase2(tglb[P2idx]) {
            return P2idx; 
        } else {
            assert InPhase1(tglb[P2idx]);
            P2idx := P2idx + 1;
            var k := P2idx - 1;
            forall j | 0 <= j <= k ensures Phase1Invariant(tglb[j], opn);
            forall j | 0 <= j <= k ensures InPhase1(tglb[j]);
        }
    }
    return P2idx;
}

}
