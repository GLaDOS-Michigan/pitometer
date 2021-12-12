include "../TimestampedRslSystem.i.dfy"
include "../../CommonProof/Constants.i.dfy"

include "../Failover/FailureDetection_toplevel.i.dfy"
include "../Phase1/Phase1Proof_toplevel.i.dfy"
include "../Phase2_PostFail/Phase2Proof_toplevel.i.dfy"
include "../Common/assumptions.i.dfy"


module RslFailureProof_i {
import opened TimestampedRslSystem_i
import opened CommonProof__Constants_i
import opened Common_Assumptions

import FO_top = FailureDetection_i
import FO_def = FailureDetection_defns_i
import P1_top = RslPhase1Proof_Top  
import P1_def = RslPhase1Proof_i  
import P2_top = RslPhase2Proof_PostFail_Top
import P2_def = RslPhase2Proof_PostFail_i

/*****************************************************************************************
*                                      Guarantees                                        *
*****************************************************************************************/

/* Main performance guarantee */
predicate PerformanceGuarantee_Final(ts:TimestampedRslState) 
    requires |ts.t_replicas| > 2
    requires RslConsistency(ts)
{
    forall pkt | 
        && pkt in ts.undeliveredPackets 
        && IsNewReplyPacket(ts, pkt)
    :: TimeLe(pkt.msg.ts, P2_def.TimeBoundReplyFinal())
}


/*****************************************************************************************
*                                     Assumptions                                        *
*****************************************************************************************/

/* Conjunction of all assumptions */
predicate RslAssumption(ts:TimestampedRslState, opn:OperationNumber)
{
    && CommonAssumptions(ts)
    && (InFailover(ts) ==> FO_def.FOAssumption(ts))
    && (InPhase1(ts) ==> P1_def.P1Assumption(ts))
    && (InPhase2(ts) ==> P2_def.P2Assumption(ts, opn))
    
}

predicate RslAssumption2(ts:TimestampedRslState, ts':TimestampedRslState)
    requires |ts.t_replicas| > 2 && |ts'.t_replicas| > 2
{
    (InFailover(ts) && InFailover(ts') ==> FO_def.FOAssumption2(ts, ts'))
}



/*****************************************************************************************
*                                Main Theorem Proof                                      *
*****************************************************************************************/


lemma RSLTopLevel(tglb:seq<TimestampedRslState>, opn:OperationNumber)
    requires |tglb| > 0
    requires exists con :: ValidTimestampedRSLBehavior(con, tglb)
    requires forall i | 0 <= i < |tglb| :: RslAssumption(tglb[i], opn)
    ensures forall j | 0 <= j < |tglb| :: PerformanceGuarantee_Final(tglb[j])
{
    assert InFailover(tglb[0]);
    var P1_idx := FO_top.FailoverTopLevel_Prototype(tglb);
    if P1_idx >= |tglb| {
        assert forall i | 0 <= i < |tglb| :: |tglb[i].t_replicas| > 2;
        assert forall i | 0 <= i < |tglb| :: PerformanceGuarantee_Final(tglb[i]);  
        return;         // We are done
    } else {
        assert forall i | 0 <= i < P1_idx :: |tglb[i].t_replicas| > 2;
        assert forall i | 0 <= i < P1_idx :: PerformanceGuarantee_Final(tglb[i]);

        assert InPhase1(tglb[P1_idx]);
        forall i | 0 < i < |tglb[P1_idx..]| 
        ensures TimestampedRslNext(tglb[P1_idx..][i - 1], tglb[P1_idx..][i]){
            assert tglb[P1_idx..][i - 1] == tglb[P1_idx + i - 1];
            assert tglb[P1_idx..][i] == tglb[P1_idx + i];
        }
        var P2_idx := P1_top.Phase1TopLevel_Prototype(tglb[P1_idx..], opn);
        P2_idx := P2_idx + P1_idx;

        if P2_idx >= |tglb| {
            assert forall i | 0 <= i < |tglb| :: |tglb[i].t_replicas| > 2;
            assert forall i | 0 <= i < |tglb| :: PerformanceGuarantee_Final(tglb[i]);  
            return;     // We are done
        } else {
            assert forall i | 0 <= i < P2_idx :: |tglb[i].t_replicas| > 2;
            assert forall i | 0 <= i < P2_idx :: PerformanceGuarantee_Final(tglb[i]);  

            forall i | 0 < i < |tglb[P2_idx..]| 
            ensures TimestampedRslNext(tglb[P2_idx..][i - 1], tglb[P2_idx..][i]){
                assert tglb[P2_idx..][i - 1] == tglb[P2_idx + i - 1];
                assert tglb[P2_idx..][i] == tglb[P2_idx + i];
            }
            assert InPhase2(tglb[P2_idx]);
            assert P2_def.Phase2Invariant(tglb[P2_idx], opn);
            P2_top.Phase2TopLevel_Prototype(tglb[P2_idx..], opn);
            assert forall i | 0 <= i < |tglb| :: PerformanceGuarantee_Final(tglb[i]);
            // We are done
        }
    }
}


}
