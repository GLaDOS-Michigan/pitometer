include "Phase2Proof.i.dfy"
include "Phase2Proof_helper.i.dfy"

module Rs2Phase2Proof_Top {
import opened RslPhase2Proof_postFail_i
import opened Rs2Phase2Proof_Helper



lemma PerfInvariantMaintained(s:TimestampedRslState, s':TimestampedRslState, req_time:Timestamp, opn:OperationNumber)
    requires RslAssumption(s) && RslAssumption(s')
    requires RslConsistency(s) && RslConsistency(s')
    requires TimestampedRslNext(s, s')
    requires RslPerfInvariant(s, opn)
    ensures RslPerfInvariant(s', opn)
{   
    PacketsBallotInvariant_Maintained(s, s', opn);
    AlwaysInvariant_Maintained(s, s', opn);
    
    if !exists pkt :: pkt in s.t_environment.sentPackets && IsNew2aPacket(pkt, opn) {
        Before2a_to_BeforeOrAfter2a(s, s', opn);
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


lemma Before2a_to_BeforeOrAfter2a(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber) 
    requires RslAssumption(ts) && RslConsistency(ts)
    requires RslAssumption(ts') && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires RslPerfInvariant(ts, opn)
    ensures Before_2a_Sent_Invariant(ts', opn) || After_2a_Sent_Invariant(ts', opn);
{
    assume false;
    if TimestampedRslNextEnvironment(ts, ts') {
        assert PerformanceGuarantee_2a(ts', opn);
        return;
    }

    var idx, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>> :| TimestampedRslNextOneReplica(ts, ts', idx, tios);
    var nextActionIndex := ts.t_replicas[idx].v.nextActionIndex;
    if idx != 1 {
        assert PerformanceGuarantee_2a(ts', opn);
        return;
    }
    // Actor is the leader now
    if exists pkt :: pkt in ts.t_environment.sentPackets && IsNew2aPacket(pkt, opn) {
        /* TODO: If I enter this clause, then any new 2a packets I sent out can't be of opn */

        if nextActionIndex == 3 {
            assert PerformanceGuarantee_2a(ts', opn);
        } else {
            forall pkt | pkt in ts'.undeliveredPackets && IsNew2aPacket(pkt, opn)
            ensures pkt in ts.undeliveredPackets
            {}
            assert PerformanceGuarantee_2a(ts', opn);
        }   



        // forall pkt | pkt in ts'.undeliveredPackets && IsNew2aPacket(pkt, opn)
        // ensures TimeLe(pkt.msg.ts, TimeBound2aDeliveryPost())
        // {}
        // assert PerformanceGuarantee_2a(ts', opn);
        return;
    }
    
    
    
    assert nextActionIndex == 3;
    assert PerformanceGuarantee_2a(ts', opn);
}

}