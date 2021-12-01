include "Phase2Proof.i.dfy"
include "Phase2Proof_helper.i.dfy"

module Rs2Phase2Proof_Top {
import opened RslPhase2Proof_postFail_i
import opened Rs2Phase2Proof_Helper



lemma PerfInvariantMaintained(s:TimestampedRslState, s':TimestampedRslState, req_time:Timestamp, opn:OperationNumber)
    requires RslAssumption(s) && RslConsistency(s)
    requires RslAssumption(s') && RslConsistency(s')
    requires TimestampedRslNext(s, s')
    requires RslPerfInvariant(s, opn)
    ensures RslPerfInvariant(s', opn)
{   
    PacketsBallotInvariant_Maintained(s, s', opn);
    assert PacketsBallotInvariant(s');
    
    PerformanceGuarantee_2a_Maintained(s, s', opn);
    assert PerformanceGuarantee_2a(s', opn);
    
    assume false;
    assert PerformanceGuarantee(s', opn);
}


/*****************************************************************************************
*                                         Lemmas                                         *
*****************************************************************************************/

lemma PacketsBallotInvariant_Maintained(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber) 
    requires RslAssumption(ts) && RslConsistency(ts)
    requires RslAssumption(ts') && RslConsistency(ts')
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

lemma PerformanceGuarantee_2a_Maintained(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber) 
    requires RslAssumption(ts) && RslConsistency(ts)
    requires RslAssumption(ts') && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires RslPerfInvariant(ts, opn)
    ensures PerformanceGuarantee_2a(ts', opn);
{
    if TimestampedRslNextEnvironment(ts, ts') {
        assert PerformanceGuarantee_2a(ts', opn);
        return;
    }
    if exists pkt :: pkt in ts.t_environment.sentPackets && IsNew2aPacket(pkt, opn) {
        /* TODO: I can't say anything about leader timestamp at this point.
        Leader could send new 2a's out. And these 2a's will be unbounded
        */

        forall pkt | pkt in ts'.undeliveredPackets && IsNew2aPacket(pkt, opn)
        ensures TimeLe(pkt.msg.ts, TimeBound2aDeliveryPost())
        {}
        assert PerformanceGuarantee_2a(ts', opn);
        return;
    }
    
    var idx, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>> :| TimestampedRslNextOneReplica(ts, ts', idx, tios);
    if idx != 1 {
        assert PerformanceGuarantee_2a(ts', opn);
        return;
    }
    
    var nextActionIndex := ts.t_replicas[idx].v.nextActionIndex;
    assert nextActionIndex == 3;
    assert PerformanceGuarantee_2a(ts', opn);
}

}