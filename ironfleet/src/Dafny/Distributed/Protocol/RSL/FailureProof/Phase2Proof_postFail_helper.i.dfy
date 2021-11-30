include "Phase2Proof_postFail.i.dfy"

module Rs2Phase2Proof_Helper {
import opened RslPhase2Proof_postFail_i


/* Note: this requires a long timeout to verify */
lemma BoundaryInvariantsMaintained_ReceiveStep_ExistingPacketsBallot(ts:TimestampedRslState, ts':TimestampedRslState, req_time:Timestamp, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts) && RslConsistency(ts)
    requires RslAssumption(ts') && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires RslPerfInvariant(ts, req_time, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires ts.t_replicas[idx].v.nextActionIndex == 0 
    ensures BoundaryConditionInvariant_ExistingPacketsBallot(ts')
{
    var tsSentPackets := (set io | io in tios && io.LIoOpSend? :: io.s);
    var s, s', ios := UntimestampRslState(ts), UntimestampRslState(ts'), UntagLIoOpSeq(tios);
    var r, r' := s.replicas[idx].replica, s'.replicas[idx].replica;

    if ios[0].LIoOpTimeoutReceive? {
        return;
    }
    if ios[0].r.msg.RslMessage_Heartbeat? {
        assert false;
    } else {
        var sentPackets := ExtractSentPacketsFromIos(ios);
        match ios[0].r.msg {
            case RslMessage_Invalid => assert true;
            case RslMessage_Request(_, _) => 
                forall pkt | pkt in ts'.undeliveredPackets 
                ensures Boundary_ExistingPacketsBallot(pkt)
                {
                    if pkt !in ts.undeliveredPackets {
                        assert pkt in tsSentPackets;
                        assert pkt.msg.v.RslMessage_Reply?;
                    }
                }                
            case RslMessage_1a(_) => 
                forall pkt | pkt in ts'.undeliveredPackets 
                ensures Boundary_ExistingPacketsBallot(pkt)
                {}
            case RslMessage_1b(_, _, _) => assert true;
            case RslMessage_StartingPhase2(_, _) =>
                assert forall pkt | pkt in tsSentPackets :: pkt.msg.v.RslMessage_AppStateRequest?;
                forall pkt | pkt in ts'.undeliveredPackets 
                ensures Boundary_ExistingPacketsBallot(pkt)
                {
                    if pkt !in ts.undeliveredPackets {
                        assert pkt in tsSentPackets;
                        assert pkt.msg.v.RslMessage_AppStateRequest?;
                    }
                }
            case RslMessage_2a(_, _, _) => 
                forall pkt | pkt in ts'.undeliveredPackets 
                ensures Boundary_ExistingPacketsBallot(pkt)
                {}
            case RslMessage_2b(_, _, _) => assert true;
            case RslMessage_Reply(_, _) => assert true;
            case RslMessage_AppStateRequest(_, _) => 
                forall pkt | pkt in ts'.undeliveredPackets 
                ensures Boundary_ExistingPacketsBallot(pkt)
                {
                    if pkt !in ts.undeliveredPackets {
                        assert pkt in tsSentPackets;
                        assert pkt.msg.v.RslMessage_AppStateSupply?;
                    }
                }                
            case RslMessage_AppStateSupply(_, _, _, _) => assert true;               
        }
         
    }
}


lemma BoundaryInvariantsMaintained_ReceiveStep_ExistingPacketsTS(ts:TimestampedRslState, ts':TimestampedRslState, req_time:Timestamp, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts) && RslConsistency(ts)
    requires RslAssumption(ts') && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires RslPerfInvariant(ts, req_time, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires ts.t_replicas[idx].v.nextActionIndex == 0 
    ensures BoundaryConditionInvariant_ExistingPacketsTS(ts')
{
    // var tsSentPackets := (set io | io in tios && io.LIoOpSend? :: io.s);
    var s, s', ios := UntimestampRslState(ts), UntimestampRslState(ts'), UntagLIoOpSeq(tios);
    // var r, r' := s.replicas[idx].replica, s'.replicas[idx].replica;

    if ios[0].LIoOpTimeoutReceive? {
        return;
    }
    if ios[0].r.msg.RslMessage_Heartbeat? {
        assert false;
    }
}

lemma BoundaryInvariantsMaintained_NoReceiveStep_ExistingPacketsBallot(ts:TimestampedRslState, ts':TimestampedRslState, req_time:Timestamp, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts) && RslConsistency(ts)
    requires RslAssumption(ts') && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires RslPerfInvariant(ts, req_time, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires ts.t_replicas[idx].v.nextActionIndex != 0 
    ensures BoundaryConditionInvariant_ExistingPacketsBallot(ts')
{
    assert BoundaryConditionInvariant_ExistingPacketsBallot(ts');
}

lemma BoundaryInvariantsMaintained_NoReceiveStep_ExistingPacketsTS(ts:TimestampedRslState, ts':TimestampedRslState, req_time:Timestamp, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts) && RslConsistency(ts)
    requires RslAssumption(ts') && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires RslPerfInvariant(ts, req_time, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires ts.t_replicas[idx].v.nextActionIndex != 0 
    ensures BoundaryConditionInvariant_ExistingPacketsTS(ts')
{}

}