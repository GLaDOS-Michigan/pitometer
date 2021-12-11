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
    && (FO_def.FOAssumption(ts))
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
    var P1_idx := FO_top.FailoverTopLevel(tglb);

    if P1_idx >= |tglb| {
        assert forall i | 0 <= i < |tglb| :: |tglb[i].t_replicas| > 2;
        assert forall i | 0 <= i < |tglb| :: PerformanceGuarantee_Final(tglb[i]);  
    } else {
        assert forall i | 0 <= i < P1_idx :: |tglb[i].t_replicas| > 2;
        assert forall i | 0 <= i < P1_idx :: PerformanceGuarantee_Final(tglb[i]);
        
        assume false;
        var P2_idx := P1_top.Phase1TopLevel(tglb[P1_idx..], opn);
        if P2_idx >= |tglb| {
            assert forall i | 0 <= i < |tglb| :: |tglb[i].t_replicas| > 2;
            assert forall i | 0 <= i < |tglb| :: PerformanceGuarantee_Final(tglb[i]);  
        } else {
            assert forall i | 0 <= i < P2_idx :: |tglb[i].t_replicas| > 2;
            assert forall i | 0 <= i < P2_idx :: PerformanceGuarantee_Final(tglb[i]);  

            P2_top.Phase2TopLevel(tglb[P2_idx..], opn);
            assert forall i | 0 <= i < |tglb| :: PerformanceGuarantee_Final(tglb[i]);
        }
    }
}


/*****************************************************************************************
*                                  Misc Definitions                                      *
*****************************************************************************************/




predicate InFailover(ts:TimestampedRslState) 
    requires |ts.t_replicas| > 2
{
    ts.t_replicas[1].v.replica.proposer.election_state.current_view == Ballot(1,0)
}

predicate InPhase1(ts:TimestampedRslState) 
    requires |ts.t_replicas| > 2
{
    && ts.t_replicas[1].v.replica.proposer.election_state.current_view == Ballot(1,1)
    && ts.t_replicas[1].v.replica.proposer.current_state != 2
}

predicate InPhase2(ts:TimestampedRslState) 
    requires |ts.t_replicas| > 2
{
    && ts.t_replicas[1].v.replica.proposer.election_state.current_view == Ballot(1,1)
    && ts.t_replicas[1].v.replica.proposer.current_state == 2
}

}
