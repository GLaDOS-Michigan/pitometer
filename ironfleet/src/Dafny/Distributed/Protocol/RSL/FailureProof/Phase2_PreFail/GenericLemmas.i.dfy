include "Phase2Proof.i.dfy"

module Rs2Phase2Proof_PreFail_Generic {
import opened RslPhase2Proof_PreFail_i


/* There can be no 2b messages in the system before there are 2a's */
lemma lemma_No2bBefore2aSent(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts, opn) && RslConsistency(ts)
    requires RslAssumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires Before_2a_Sent_Invariant(ts, opn);
    ensures !exists pkt :: pkt in ts'.t_environment.sentPackets && IsPreFail2bPacket(pkt, opn)
{
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    var nextActionIndex := ls.v.nextActionIndex;
    forall p | p in ts'.t_environment.sentPackets 
    ensures !IsPreFail2bPacket(p, opn) {
        if IsPreFail2bPacket(p, opn) {
            assert p !in ts.t_environment.sentPackets;
            if nextActionIndex == 0 {
                assert !tios[0].r.msg.v.RslMessage_2a?;
                assert forall io | io in tios && io.LIoOpSend? :: !io.s.msg.v.RslMessage_2b?;
                assert false;
            } else {
                assert forall io | io in tios && io.LIoOpSend? :: !io.s.msg.v.RslMessage_2b?;
                assert false;
            }
        }
    }
}


/* Non-leader replicas do not send 2a's */
lemma lemma_NonLeaderDoesNotSend2a(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts, opn) && RslConsistency(ts)
    requires RslAssumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires Before_2a_Sent_Invariant(ts, opn) || Before_2b_Sent_Invariant(ts, opn) || After_2b_Sent_Invariant(ts, opn)
    requires idx != 0
    ensures forall p | p in ts'.t_environment.sentPackets && IsPreFail2aPacket(p, opn) :: p in ts.t_environment.sentPackets
{
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    var nextActionIndex := ls.v.nextActionIndex;
    forall p | p in ts'.t_environment.sentPackets && IsPreFail2aPacket(p, opn) 
    ensures p in ts.t_environment.sentPackets {
        if p !in ts.t_environment.sentPackets {
            if nextActionIndex == 3 {
                assert !LProposerCanNominateUsingOperationNumber(ls.v.replica.proposer, ls.v.replica.acceptor.log_truncation_point, ls.v.replica.proposer.next_operation_number_to_propose);
                assert forall io | io in tios && io.LIoOpSend? :: !io.s.msg.v.RslMessage_2a?;
                assert false;
            } else {
                assert forall io | io in tios && io.LIoOpSend? :: !io.s.msg.v.RslMessage_2a?;
                assert false;
            }
        }
    }
}

/* Non-leader replicas do not send replies */
lemma lemma_NonLeaderDoesNotSendReply(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts, opn) && RslConsistency(ts)
    requires RslAssumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires Before_2a_Sent_Invariant(ts, opn) || Before_2b_Sent_Invariant(ts, opn) || After_2b_Sent_Invariant(ts, opn)
    requires idx != 0
    ensures forall p | p in ts'.t_environment.sentPackets && IsPreFailReplyPacket(ts, p) :: p in ts.t_environment.sentPackets
{
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    var nextActionIndex := ls.v.nextActionIndex;
    forall p | p in ts'.t_environment.sentPackets && IsPreFailReplyPacket(ts, p)
    ensures p in ts.t_environment.sentPackets {
        if p !in ts.t_environment.sentPackets {
            if nextActionIndex == 6 {
                assert forall io | io in tios && io.LIoOpSend? :: io.s.src == ts'.constants.config.replica_ids[idx];
                assert p.src == ts'.constants.config.replica_ids[idx];
                assert ReplicasDistinct(ts'.constants.config.replica_ids, 0, idx);
                assert false;
            } else {
                assert forall io | io in tios && io.LIoOpSend? :: !io.s.msg.v.RslMessage_2a?;
                assert false;
            }
        }
    }
}

}