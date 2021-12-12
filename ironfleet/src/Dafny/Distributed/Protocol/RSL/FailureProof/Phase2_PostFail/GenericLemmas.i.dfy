include "Phase2Proof.i.dfy"

module RslPhase2Proof_PostFail_Generic {
import opened RslPhase2Proof_PostFail_i


/* There can be no 2b messages in the system before there are 2a's */
lemma {:timeLimitMultiplier 2} lemma_NoNew2bBefore2aSent(ts:TimestampedRslState, ts':TimestampedRslState, op:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P2Assumption(ts, op)
    requires PacketsBallotInvariant(ts)
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires Before_2a_Sent_Invariant(ts, op)
    ensures !exists pkt :: pkt in ts'.t_environment.sentPackets && IsNew2bPacket(pkt, op)
{
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    var nextActionIndex := ls.v.nextActionIndex;
    forall p | p in ts'.t_environment.sentPackets 
    ensures !IsNew2bPacket(p, op) {
        if IsNew2bPacket(p, op) {
            assert p !in ts.t_environment.sentPackets;
            if nextActionIndex == 0 {
                assert !IsNew2aPacket(tios[0].r, op);
                forall io | io in tios && io.LIoOpSend? 
                ensures !IsNew2bPacket(io.s, op) {}
                assert false;
            } else {
                assert forall io | io in tios && io.LIoOpSend? :: !io.s.msg.v.RslMessage_2b?;
                assert false;
            }
        }
    }
}


// /* There can be no 1b messages in the system before there are 1a's */
// lemma lemma_No1bBefore1aSent(ts:TimestampedRslState, ts':TimestampedRslState, op:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
//     requires P2Assumption(ts, op) && RslConsistency(ts)
//     requires P2Assumption(ts', op) && RslConsistency(ts')
//     requires TimestampedRslNext(ts, ts')
//     requires !TimestampedRslNextEnvironment(ts, ts')
//     requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
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


/* Non-leader replicas do not send 2a's */
lemma lemma_NonLeaderDoesNotSend2a(ts:TimestampedRslState, ts':TimestampedRslState, op:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P2Assumption(ts, op)
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires Before_2a_Sent_Invariant(ts, op) || Before_2b_Sent_Invariant(ts, op) || After_2b_Sent_Invariant(ts, op)
    requires idx != 1
    ensures forall p | p in ts'.t_environment.sentPackets && p.msg.v.RslMessage_2a? :: p in ts.t_environment.sentPackets
    ensures forall p | p in ts'.undeliveredPackets && p.msg.v.RslMessage_2a? :: p in ts.undeliveredPackets
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
    lemma_NonLeaderDoesNotSend2a_Undelivered(ts, ts', op, idx, tios);
}

/* Non-leader replicas do not send 2a's */
lemma lemma_NonLeaderDoesNotSend2a_Undelivered(ts:TimestampedRslState, ts':TimestampedRslState, op:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P2Assumption(ts, op)
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires Before_2a_Sent_Invariant(ts, op) || Before_2b_Sent_Invariant(ts, op) || After_2b_Sent_Invariant(ts, op)
    requires idx != 1
    ensures forall p | p in ts'.undeliveredPackets && p.msg.v.RslMessage_2a? :: p in ts.undeliveredPackets
{
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    var nextActionIndex := ls.v.nextActionIndex;
    forall p | p in ts'.undeliveredPackets && p.msg.v.RslMessage_2a?
    ensures p in ts.undeliveredPackets {
        if p !in ts.undeliveredPackets {
            if nextActionIndex == 3 {
                assert !LProposerCanNominateUsingOperationNumber(ls.v.replica.proposer, ls.v.replica.acceptor.log_truncation_point, ls.v.replica.proposer.next_operation_number_to_propose);
                assert forall io | io in tios && io.LIoOpSend? :: !io.s.msg.v.RslMessage_2a?;
                assert false;
            } else {
                lemma_No2aSentInNon2aStep(ts, ts', op, idx, tios);
                assert forall io | io in tios && io.LIoOpSend? :: !io.s.msg.v.RslMessage_2a?;
                assert false;
            }
        }
    }
}

/* Non-leader replicas do not send replies */
lemma lemma_NonLeaderDoesNotSendReply(ts:TimestampedRslState, ts':TimestampedRslState, op:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P2Assumption(ts, op)
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires Phase2Invariant(ts, op)
    requires idx != 1
    ensures forall p | p in ts'.t_environment.sentPackets && IsNewReplyPacket(ts, p) :: p in ts.t_environment.sentPackets
    ensures forall p | p in ts'.undeliveredPackets && IsNewReplyPacket(ts, p) :: p in ts.undeliveredPackets
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
    lemma_NonLeaderDoesNotSendReply_Undelivered(ts, ts', op, idx, tios);
}

/* Non-leader replicas do not send replies */
lemma lemma_NonLeaderDoesNotSendReply_Undelivered(ts:TimestampedRslState, ts':TimestampedRslState, op:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P2Assumption(ts, op)
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires Phase2Invariant(ts, op)
    requires idx != 1
    ensures forall p | p in ts'.undeliveredPackets && IsNewReplyPacket(ts, p) :: p in ts.undeliveredPackets
{
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    var nextActionIndex := ls.v.nextActionIndex;
    forall p | p in ts'.undeliveredPackets && IsNewReplyPacket(ts, p)
    ensures p in ts.undeliveredPackets {
        if p !in ts.undeliveredPackets {
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

/* There can be no 1b messages sent in a Non-Receive step  */
lemma lemma_No1bSentInNonReceiveStep(ts:TimestampedRslState, ts':TimestampedRslState, op:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P2Assumption(ts, op)
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires Phase2Invariant(ts, op)
    requires ts.t_replicas[idx].v.nextActionIndex != 0
    // ensures forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2b? :: pkt in ts.t_environment.sentPackets
    ensures forall pkt | pkt in ts'.undeliveredPackets && pkt.msg.v.RslMessage_1b? :: pkt in ts.undeliveredPackets
{
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    var nextActionIndex := ls.v.nextActionIndex;
    forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_1b?
    ensures pkt in ts.t_environment.sentPackets 
    {
        if pkt !in ts.t_environment.sentPackets {
            var ios := UntagLIoOpSeq(tios);
            var sent_packets := ExtractSentPacketsFromIos(ios);
            assert LReplicaNoReceiveNext(ls.v.replica, nextActionIndex, ls'.v.replica, ios);
            assert forall p | p in sent_packets :: !p.msg.RslMessage_1b?;
            forall io | io in tios && io.LIoOpSend?
            ensures !io.s.msg.v.RslMessage_1b? {}
            assert false;
        }
    }
    lemma_No2bSentInNonReceiveStep_Undelivered(ts, ts', op, idx, tios);
}


/* There can be no 2b messages sent in a Non-Receive step  */
lemma lemma_No2bSentInNonReceiveStep(ts:TimestampedRslState, ts':TimestampedRslState, op:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P2Assumption(ts, op)
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires Phase2Invariant(ts, op)
    requires ts.t_replicas[idx].v.nextActionIndex != 0
    ensures forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2b? :: pkt in ts.t_environment.sentPackets
    ensures forall pkt | pkt in ts'.undeliveredPackets && pkt.msg.v.RslMessage_2b? :: pkt in ts.undeliveredPackets
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
    lemma_No2bSentInNonReceiveStep_Undelivered(ts, ts', op, idx, tios);
}


lemma lemma_No2bSentInNonReceiveStep_Undelivered(ts:TimestampedRslState, ts':TimestampedRslState, op:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P2Assumption(ts, op)
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires Phase2Invariant(ts, op)
    requires ts.t_replicas[idx].v.nextActionIndex != 0
    ensures forall pkt | pkt in ts'.undeliveredPackets && pkt.msg.v.RslMessage_2b? :: pkt in ts.undeliveredPackets
{
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    var nextActionIndex := ls.v.nextActionIndex;
    forall pkt | pkt in ts'.undeliveredPackets && pkt.msg.v.RslMessage_2b?
    ensures pkt in ts.undeliveredPackets
    {
        if pkt !in ts.undeliveredPackets {
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


/* There can be no 2b messages sent in a receive step unless receiving 2a  */
lemma lemma_No2bSentInReceiveStep_NotReceive2a(ts:TimestampedRslState, ts':TimestampedRslState, op:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P2Assumption(ts, op)
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires Phase2Invariant(ts, op)
    requires ts.t_replicas[idx].v.nextActionIndex == 0
    requires !tios[0].r.msg.v.RslMessage_2a? 
    ensures forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2b? :: pkt in ts.t_environment.sentPackets
    ensures forall pkt | pkt in ts'.undeliveredPackets && pkt.msg.v.RslMessage_2b? :: pkt in ts.undeliveredPackets
{
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    forall p | p in ts'.t_environment.sentPackets && p.msg.v.RslMessage_2b? 
    ensures p in ts.t_environment.sentPackets
    {
        if p !in ts.t_environment.sentPackets {
            var sent_packets := ExtractSentPacketsFromIos(UntagLIoOpSeq(tios));
            forall p | p in sent_packets 
            ensures !p.msg.RslMessage_2b? {}
            assert false;
        }
    }
    forall p | p in ts'.undeliveredPackets && p.msg.v.RslMessage_2b?
    ensures p in ts.undeliveredPackets
    {
        if p !in ts.t_environment.sentPackets {
            var sent_packets := ExtractSentPacketsFromIos(UntagLIoOpSeq(tios));
            forall p | p in sent_packets 
            ensures !p.msg.RslMessage_2b? {}
            assert false;
        }
    }
}


lemma lemma_No1bSentInReceiveStep_NotReceive1a(ts:TimestampedRslState, ts':TimestampedRslState, op:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P2Assumption(ts, op)
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires Phase2Invariant(ts, op)
    requires ts.t_replicas[idx].v.nextActionIndex == 0
    requires !tios[0].r.msg.v.RslMessage_1a? 
    // ensures forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2b? :: pkt in ts.t_environment.sentPackets
    ensures forall pkt | pkt in ts'.undeliveredPackets && pkt.msg.v.RslMessage_1b? :: pkt in ts.undeliveredPackets
{
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    forall p | p in ts'.undeliveredPackets && p.msg.v.RslMessage_1b?
    ensures p in ts.undeliveredPackets
    {
        if p !in ts.t_environment.sentPackets {
            var sent_packets := ExtractSentPacketsFromIos(UntagLIoOpSeq(tios));
            forall p | p in sent_packets 
            ensures !p.msg.RslMessage_1b? {}
            assert false;
        }
    }
}


/* There can be no 2a messages sent in a Receive step  */
lemma {:timeLimitMultiplier 2} lemma_No2aSentInNon2aStep(ts:TimestampedRslState, ts':TimestampedRslState, op:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P2Assumption(ts, op)
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires Phase2Invariant(ts, op)
    requires ts.t_replicas[idx].v.nextActionIndex != 3
    ensures forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2a? :: pkt in ts.t_environment.sentPackets
    ensures forall p | p in ts'.undeliveredPackets && p.msg.v.RslMessage_2a? :: p in ts.undeliveredPackets
{
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    var nextActionIndex := ls.v.nextActionIndex;

    forall p | p in ts'.t_environment.sentPackets && p.msg.v.RslMessage_2a?
    ensures p in ts.t_environment.sentPackets {
        if p !in ts.t_environment.sentPackets {
            if nextActionIndex == 3 {
                assert false;
            } else {
                forall io | io in tios && io.LIoOpSend?
                ensures !io.s.msg.v.RslMessage_2a? {}
                assert false;
            }
        }
    }
    forall p | p in ts'.undeliveredPackets && p.msg.v.RslMessage_2a?
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
}


/* There can be no Reply messages sent in a Non-Execution step  */
lemma lemma_NoRepliesSentInNonExecutionStep(ts:TimestampedRslState, ts':TimestampedRslState, op:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P2Assumption(ts, op)
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires Phase2Invariant(ts, op)
    requires ts.t_replicas[idx].v.nextActionIndex != 6
    ensures forall pkt | pkt in ts'.t_environment.sentPackets && IsNewReplyPacket(ts', pkt) :: pkt in ts.t_environment.sentPackets
{}


/* In a receive step, if sent_packets is empty, no packets get added to the network */
lemma lemma_EmptySentPackets(ts:TimestampedRslState, ts':TimestampedRslState, op:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P2Assumption(ts, op)
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires Phase2Invariant(ts, op)
    requires ExtractSentPacketsFromIos(UntagLIoOpSeq(tios)) == []
    ensures forall pkt | pkt in ts'.t_environment.sentPackets :: pkt in ts.t_environment.sentPackets
{
    var ios :=  UntagLIoOpSeq(tios);
    forall tio | tio in tios 
    ensures !tio.LIoOpSend? {
        if tio.LIoOpSend? {
            assert UntagLIoOp(tio) in ios;
            assert UntagLIoOp(tio).LIoOpSend?;
            assert false;
        }
    }
    forall p | p in ts'.t_environment.sentPackets 
    ensures p in ts.t_environment.sentPackets 
    {}
}

}