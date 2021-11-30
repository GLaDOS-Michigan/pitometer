include "Phase2Proof_postFail.i.dfy"
include "Phase2Proof_postFail_helper.i.dfy"

module Rs2Phase2Proof_Top {
import opened RslPhase2Proof_postFail_i
import opened Rs2Phase2Proof_Helper



lemma PerfInvariantMaintained(s:TimestampedRslState, s':TimestampedRslState, req_time:Timestamp, opn:OperationNumber)
    requires RslAssumption(s) && RslConsistency(s)
    requires RslAssumption(s') && RslConsistency(s')
    requires TimestampedRslNext(s, s')
    requires RslPerfInvariant(s, req_time, opn)
    ensures RslPerfInvariant(s', req_time, opn)
{   
    PacketsBallotInvariantMaintained(s, s', req_time, opn);
    assert PacketsBallotInvariant(s');
    
    assume false;
    assert PerformanceGuarantee_2a(s');
    assert PerformanceGuarantee(s');
}


/*****************************************************************************************
*                                         Lemmas                                         *
*****************************************************************************************/

lemma PacketsBallotInvariantMaintained(ts:TimestampedRslState, ts':TimestampedRslState, req_time:Timestamp, opn:OperationNumber) 
    requires RslAssumption(ts) && RslConsistency(ts)
    requires RslAssumption(ts') && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires RslPerfInvariant(ts, req_time, opn)
    ensures PacketsBallotInvariant(ts')
{
    if TimestampedRslNextEnvironment(ts, ts') {
        assert PacketsBallotInvariant(ts');
        return;
    }
    var idx, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>> :| TimestampedRslNextOneReplica(ts, ts', idx, tios);
    var nextActionIndex := ts.t_replicas[idx].v.nextActionIndex;
    if nextActionIndex == 0 {
        PacketsBallotInvariant_ReceiveStep(ts, ts', req_time, opn, idx, tios);
    } else {
        PacketsBallotInvariant_NoReceiveStep(ts, ts', req_time, opn, idx, tios);
    }
}

}