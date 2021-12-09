include "../TimestampedRslSystem.i.dfy"
include "../../CommonProof/Constants.i.dfy"

include "../Failover/FailureDetection_toplevel.i.dfy"
include "../Phase2_PostFail/Phase2Proof_toplevel.i.dfy"
include "../Common/assumptions.i.dfy"


module RslFailureProof_i {
import opened TimestampedRslSystem_i
import opened CommonProof__Constants_i
import opened Common_Assumptions

import FOStg_Def = FailureDetection_defns_i
import P2Stg_Pf = Rs2Phase2Proof_PostFail_Top  
import P2Stg_Def = RslPhase2Proof_PostFail_i

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
        && P2Stg_Def.IsNewReplyPacket(ts, pkt)
    :: TimeLe(pkt.msg.ts, P2Stg_Def.TimeBoundReplyFinal())
}


/*****************************************************************************************
*                                     Assumptions                                        *
*****************************************************************************************/

/* Conjunction of all assumptions */
predicate RslAssumption(ts:TimestampedRslState, opn:OperationNumber)
{
    && CommonAssumptions(ts)
    && (InPhase2Stage(ts) ==> P2Stg_Def.RslAssumption(ts, opn))
    && (InFailoverStage(ts) ==> FOStg_Def.RslAssumption(ts))
}

predicate RslAssumption2(ts:TimestampedRslState, ts':TimestampedRslState)
{
    (InFailoverStage(ts) && InFailoverStage(ts') ==> FOStg_Def.RslAssumption2(ts, ts'))
}



/*****************************************************************************************
*                                     Invariants                                        *
*****************************************************************************************/


// // Main invariant 
// predicate RslPerfInvariant(ts:TimestampedRslState, opn:OperationNumber) 
//     requires |ts.t_replicas| > 2
// {
//     && RslConsistency(ts)
//     && AlwaysInvariant(ts, opn)
//     && PerformanceGuarantee(ts, opn)
//     && PacketsBallotInvariant(ts)
//     && (|| Before_2a_Sent_Invariant(ts, opn)
//         || Before_2b_Sent_Invariant(ts, opn)
//         || After_2b_Sent_Invariant(ts, opn)
//     )
// }


/*****************************************************************************************
*                                  Misc Definitions                                      *
*****************************************************************************************/




predicate InFailoverStage(ts:TimestampedRslState) 

predicate InPhase1Stage(ts:TimestampedRslState) 

predicate InPhase2Stage(ts:TimestampedRslState) 

}
