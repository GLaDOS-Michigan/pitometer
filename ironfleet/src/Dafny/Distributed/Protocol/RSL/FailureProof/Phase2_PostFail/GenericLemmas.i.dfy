include "Phase2Proof.i.dfy"

module RslPhase2Proof_PostFail_Generic {
import opened RslPhase2Proof_PostFail_i

/* WARNING: this file a timeout of 60s to verify */


/* There can be no 2b messages in the system before there are 2a's */
lemma {:timeLimitMultiplier 2} lemma_NoNew2bBefore2aSent(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires P2Assumption(ts, opn) && RslConsistency(ts)
    requires P2Assumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires Before_2a_Sent_Invariant(ts, opn)
    ensures !exists pkt :: pkt in ts'.t_environment.sentPackets && IsNew2bPacket(pkt, opn)
{
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    var nextActionIndex := ls.v.nextActionIndex;
    forall p | p in ts'.t_environment.sentPackets 
    ensures !IsNew2bPacket(p, opn) {
        if IsNew2bPacket(p, opn) {
            assert p !in ts.t_environment.sentPackets;
            if nextActionIndex == 0 {
                assert !IsNew2aPacket(tios[0].r, opn);
                forall io | io in tios && io.LIoOpSend? 
                ensures !IsNew2bPacket(io.s, opn) {}
                assert false;
            } else {
                assert forall io | io in tios && io.LIoOpSend? :: !io.s.msg.v.RslMessage_2b?;
                assert false;
            }
        }
    }
}


// /* There can be no 1b messages in the system before there are 1a's */
// lemma lemma_No1bBefore1aSent(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
//     requires P2Assumption(ts, opn) && RslConsistency(ts)
//     requires P2Assumption(ts', opn) && RslConsistency(ts')
//     requires TimestampedRslNext(ts, ts')
//     requires !TimestampedRslNextEnvironment(ts, ts')
//     requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
//     requires AlwaysInvariant(ts, opn)
//     ensures forall pkt | pkt in ts'.undeliveredPackets :: !pkt.msg.v.RslMessage_1b?
// {
//     var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
//     var nextActionIndex := ls.v.nextActionIndex;
//     forall p | p in ts'.undeliveredPackets
//     ensures !p.msg.v.RslMessage_1b? {
//         if p.msg.v.RslMessage_1b? {
//             assert p !in ts.undeliveredPackets;
//             if nextActionIndex == 0 {
//                 assert !tios[0].r.msg.v.RslMessage_1a?;
//                 assert forall io | io in tios && io.LIoOpSend? :: !io.s.msg.v.RslMessage_1b?;
//                 assert false;
//             } else {
//                 assert forall io | io in tios && io.LIoOpSend? :: !io.s.msg.v.RslMessage_1b?;
//                 assert false;
//             }
//         }
//     }
// }


// /* Non-leader replicas do not send 2a's */
lemma lemma_NonLeaderDoesNotSend2a(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires P2Assumption(ts, opn) && RslConsistency(ts)
    requires P2Assumption(ts', opn) && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires Before_2a_Sent_Invariant(ts, opn) || Before_2b_Sent_Invariant(ts, opn) || After_2b_Sent_Invariant(ts, opn)
    requires idx != 1
    ensures forall p | p in ts'.t_environment.sentPackets && p.msg.v.RslMessage_2a? :: p in ts.t_environment.sentPackets
    ensures forall p | p in ts'.undeliveredPackets && IsNew2aPacket(p, opn) :: p in ts.undeliveredPackets
{
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    var nextActionIndex := ls.v.nextActionIndex;
    forall p | p in ts'.t_environment.sentPackets && p.msg.v.RslMessage_2a?
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
    lemma_NonLeaderDoesNotSend2a_Undelivered(ts, ts', opn, idx, tios);
}

/* Non-leader replicas do not send 2a's */
lemma lemma_NonLeaderDoesNotSend2a_Undelivered(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires P2Assumption(ts, opn) && RslConsistency(ts)
    requires P2Assumption(ts', opn) && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires Before_2a_Sent_Invariant(ts, opn) || Before_2b_Sent_Invariant(ts, opn) || After_2b_Sent_Invariant(ts, opn)
    requires idx != 1
    ensures forall p | p in ts'.undeliveredPackets && IsNew2aPacket(p, opn) :: p in ts.undeliveredPackets
{
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    var nextActionIndex := ls.v.nextActionIndex;
    forall p | p in ts'.undeliveredPackets && IsNew2aPacket(p, opn)
    ensures p in ts.undeliveredPackets {
        if p !in ts.undeliveredPackets {
            if nextActionIndex == 3 {
                assert !LProposerCanNominateUsingOperationNumber(ls.v.replica.proposer, ls.v.replica.acceptor.log_truncation_point, ls.v.replica.proposer.next_operation_number_to_propose);
                assert forall io | io in tios && io.LIoOpSend? :: !io.s.msg.v.RslMessage_2a?;
                assert false;
            } else {
                lemma_No2aSentInNon2aStep(ts, ts', opn, idx, tios);
                assert forall io | io in tios && io.LIoOpSend? :: !io.s.msg.v.RslMessage_2a?;
                assert false;
            }
        }
    }
}

// /* Non-leader replicas do not send replies */
lemma lemma_NonLeaderDoesNotSendReply(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires P2Assumption(ts, opn) && RslConsistency(ts)
    requires P2Assumption(ts', opn) && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires Phase2Invariant(ts, opn)
    requires idx != 1
    ensures forall p | p in ts'.t_environment.sentPackets && IsNewReplyPacket(ts, p) :: p in ts.t_environment.sentPackets
{
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    var nextActionIndex := ls.v.nextActionIndex;
    forall p | p in ts'.t_environment.sentPackets && IsNewReplyPacket(ts, p)
    ensures p in ts.t_environment.sentPackets {
        if p !in ts.t_environment.sentPackets {
            if nextActionIndex == 6 {
                assert forall io | io in tios && io.LIoOpSend? :: io.s.src == ts'.constants.config.replica_ids[idx];
                assert p.src == ts'.constants.config.replica_ids[idx];
                assert ReplicasDistinct(ts'.constants.config.replica_ids, 1, idx);
                assert false;
            } else {
                assert forall io | io in tios && io.LIoOpSend? :: !io.s.msg.v.RslMessage_2a?;
                assert false;
            }
        }
    }
}

// /* Non-leader replicas do not send replies */
// lemma lemma_NonLeaderDoesNotSendReply_Undelivered(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
//     requires P2Assumption(ts, opn) && RslConsistency(ts)
//     requires P2Assumption(ts', opn) && RslConsistency(ts')
//     // requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
//     requires TimestampedRslNext(ts, ts')
//     requires !TimestampedRslNextEnvironment(ts, ts')
//     requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
//     requires Phase2Invariant(ts, opn)
//     requires idx != 0
//     ensures forall p | p in ts'.undeliveredPackets && IsNewReplyPacket(ts, p) :: p in ts.undeliveredPackets
// {
//     var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
//     var nextActionIndex := ls.v.nextActionIndex;
//     forall p | p in ts'.undeliveredPackets && IsNewReplyPacket(ts, p)
//     ensures p in ts.undeliveredPackets {
//         if p !in ts.undeliveredPackets {
//             if nextActionIndex == 6 {
//                 assert forall io | io in tios && io.LIoOpSend? :: io.s.src == ts'.constants.config.replica_ids[idx];
//                 assert p.src == ts'.constants.config.replica_ids[idx];
//                 assert ReplicasDistinct(ts'.constants.config.replica_ids, 0, idx);
//                 assert false;
//             } else {
//                 assert forall io | io in tios && io.LIoOpSend? :: !io.s.msg.v.RslMessage_2a?;
//                 assert false;
//             }
//         }
//     }
// }


// /* There can be no 2b messages sent in a Non-Receive step  */
lemma lemma_No2bSentInNonReceiveStep(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires P2Assumption(ts, opn) && RslConsistency(ts)
    requires P2Assumption(ts', opn) && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires Phase2Invariant(ts, opn)
    requires ts.t_replicas[idx].v.nextActionIndex != 0
    ensures forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2b? :: pkt in ts.t_environment.sentPackets
{
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    var nextActionIndex := ls.v.nextActionIndex;
    forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2b?
    ensures pkt in ts.t_environment.sentPackets 
    {
        if pkt !in ts.t_environment.sentPackets {
            var ios := UntagLIoOpSeq(tios);
            var sent_packets := ExtractSentPacketsFromIos(ios);
            assert LReplicaNoReceiveNext(ls.v.replica, nextActionIndex, ls'.v.replica, ios);
            assert forall p | p in sent_packets :: !p.msg.RslMessage_2b?;
            forall io | io in tios && io.LIoOpSend?
            ensures !io.s.msg.v.RslMessage_2b? {}
            assert false;
        }
    }
}

/* There can be no 2a messages sent in a Receive step  */
lemma {:timeLimitMultiplier 2} lemma_No2aSentInNon2aStep(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires P2Assumption(ts, opn) && RslConsistency(ts)
    requires P2Assumption(ts', opn) && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires Phase2Invariant(ts, opn)
    requires ts.t_replicas[idx].v.nextActionIndex != 3
    ensures forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2a? :: pkt in ts.t_environment.sentPackets
    ensures forall p | p in ts'.undeliveredPackets && IsNew2aPacket(p, opn) :: p in ts.undeliveredPackets
{
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    var nextActionIndex := ls.v.nextActionIndex;
    forall p | p in ts'.undeliveredPackets && IsNew2aPacket(p, opn)
    ensures p in ts.undeliveredPackets {
        if p !in ts.undeliveredPackets {
            if nextActionIndex == 3 {
                assert false;
            } else {
                forall io | io in tios && io.LIoOpSend?
                ensures !io.s.msg.v.RslMessage_2a? {}
                assert false;
            }
        }
    }
    forall io | io in tios && io.LIoOpSend?
    ensures !io.s.msg.v.RslMessage_2a? {}
}


// /* There can be no Reply messages sent in a Non-Execution step  */
lemma lemma_NoRepliesSentInNonExecutionStep(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires P2Assumption(ts, opn) && RslConsistency(ts)
    requires P2Assumption(ts', opn) && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires Phase2Invariant(ts, opn)
    requires ts.t_replicas[idx].v.nextActionIndex != 6
    ensures forall pkt | pkt in ts'.t_environment.sentPackets && IsNewReplyPacket(ts', pkt) :: pkt in ts.t_environment.sentPackets
{}

}