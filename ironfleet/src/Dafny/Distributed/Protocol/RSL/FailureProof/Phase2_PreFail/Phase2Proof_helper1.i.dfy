include "Phase2Proof.i.dfy"
include "GenericLemmas.i.dfy"

module Rs2Phase2Proof_PreFail_Helper1 {
import opened RslPhase2Proof_PreFail_i
import opened Rs2Phase2Proof_PreFail_Generic

/* WARNING: this file a timeout of 60s to verify */

lemma PacketsBallotInvariant_ReceiveStep(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts, opn) && RslConsistency(ts)
    requires RslAssumption(ts', opn) && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires RslPerfInvariant(ts, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires ts.t_replicas[idx].v.nextActionIndex == 0 
    ensures PacketsBallotInvariant(ts')
{
    forall pkt | pkt in ts'.undeliveredPackets 
    ensures ExistingPacketsBallot(pkt)
    {}
}

lemma PacketsBallotInvariant_NoReceiveStep(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts, opn) && RslConsistency(ts)
    requires RslAssumption(ts', opn) && RslConsistency(ts')
    requires TimestampedRslNext(ts, ts')
    requires RslPerfInvariant(ts, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires ts.t_replicas[idx].v.nextActionIndex != 0 
    ensures PacketsBallotInvariant(ts')
{
    forall pkt | pkt in ts'.undeliveredPackets 
    ensures ExistingPacketsBallot(pkt)
    {}
}

// lemma Before2a_to_Before2b(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
//     requires RslAssumption(ts, opn) && RslConsistency(ts)
//     requires RslAssumption(ts', opn) && RslConsistency(ts')
//     requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
//     // requires AlwaysInvariant(ts', opn)
//     requires TimestampedRslNext(ts, ts')
//     requires !TimestampedRslNextEnvironment(ts, ts')
//     requires TimestampedRslNextOneReplica(ts, ts', 1, tios);
//     requires RslPerfInvariant(ts, opn)
//     requires Before_2a_Sent_Invariant(ts, opn);
//     ensures Before_2b_Sent_Invariant(ts', opn);
// {
//     var s, s', ios := UntimestampRslState(ts), UntimestampRslState(ts'), UntagLIoOpSeq(tios);
//     var r, r' := s.replicas[1].replica, s'.replicas[1].replica;
//     var sent_packets := ExtractSentPacketsFromIos(ios);
//     var v :| 
//         (&& LValIsHighestNumberedProposal(v, r.proposer.received_1b_packets, opn)
//          && LBroadcastToEveryone(r.proposer.constants.all.config, r.proposer.constants.my_index, RslMessage_2a(r.proposer.max_ballot_i_sent_1a, opn, v), sent_packets)
//     );
//     assert |sent_packets| > 0;
//     var m := RslMessage_2a(Ballot(1, 1), opn, v);
//     assert sent_packets[0] == LPacket(r.proposer.constants.all.config.replica_ids[0], r.proposer.constants.all.config.replica_ids[1], m);
//     var pkt_witness := sent_packets[0];
//     assert LIoOpSend(pkt_witness) in ios;
//     forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2b?
//     ensures pkt in ts.t_environment.sentPackets
//     {}
//     assert forall p | p in sent_packets :: !p.msg.RslMessage_Reply?;
//     forall io | io in tios && io.LIoOpSend?
//     ensures !io.s.msg.v.RslMessage_Reply? {}
// }

lemma Before2a_to_Before2a_NonLeaderAction(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts, opn) && RslConsistency(ts)
    requires RslAssumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires idx != 0;
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires RslPerfInvariant(ts, opn)
    requires Before_2a_Sent_Invariant(ts, opn);
    ensures Before_2a_Sent_Invariant(ts', opn);
{
    assert ts'.t_replicas[0] == ts.t_replicas[0];
    lemma_NonLeaderDoesNotSend2a(ts, ts', opn, idx, tios);
    lemma_No2bBefore2aSent(ts, ts', opn, idx, tios);
    lemma_NonLeaderDoesNotSendReply(ts, ts', opn, idx, tios);
}


// lemma Before2b_to_Before2b_NonReceive(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
//     requires RslAssumption(ts, opn) && RslConsistency(ts)
//     requires RslAssumption(ts', opn) && RslConsistency(ts')
//     requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
//     requires TimestampedRslNext(ts, ts')
//     requires !TimestampedRslNextEnvironment(ts, ts')
//     requires RslPerfInvariant(ts, opn)
//     requires Before_2b_Sent_Invariant(ts, opn)
//     requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
//     requires ts.t_replicas[idx].v.nextActionIndex != 0
//     ensures Before_2b_Sent_Invariant(ts', opn)
// {   
//     var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
//     var nextActionIndex, nextActionIndex' := ls.v.nextActionIndex, ls'.v.nextActionIndex;


//     Before2b_to_Before2b_NonReceive_NoNewReply(ts, ts', opn, idx, tios);
//     forall pkt | pkt in ts'.t_environment.sentPackets
//     ensures !IsNewReplyPacket(ts', pkt) 
//     {}

//     if idx != 1 {
//         forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2a?
//         ensures pkt in ts.t_environment.sentPackets
//         {}
//         forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2b?
//         ensures pkt in ts.t_environment.sentPackets
//         {
//             if pkt !in ts.t_environment.sentPackets {
//                 var ios := UntagLIoOpSeq(tios);
//                 var sent_packets := ExtractSentPacketsFromIos(ios);
//                 assert nextActionIndex != 0;
//                 assert LReplicaNoReceiveNext(ls.v.replica, nextActionIndex, ls'.v.replica, ios);
//                 assert forall p | p in sent_packets :: !p.msg.RslMessage_2b?;
//                 forall io | io in tios && io.LIoOpSend?
//                 ensures !io.s.msg.v.RslMessage_2b? {}
//                 assert false;
//             }
//         }
//         assert Before_2b_Sent_Invariant(ts', opn);
//     } else {
//         assert ts'.t_replicas[1].v.replica.learner.unexecuted_learner_state == map[];
//         assert forall opn' | opn' in ts.t_replicas[1].v.replica.learner.unexecuted_learner_state
//         ::  |ts.t_replicas[1].v.replica.learner.unexecuted_learner_state[opn].received_2b_message_senders| < LMinQuorumSize(ts.t_replicas[1].v.replica.learner.constants.all.config);
//         assert ts'.t_replicas[1].v.replica.executor == ts.t_replicas[1].v.replica.executor;

//         assert !LProposerCanNominateUsingOperationNumber(ts.t_replicas[1].v.replica.proposer, ts.t_replicas[1].v.replica.acceptor.log_truncation_point, ts.t_replicas[1].v.replica.proposer.next_operation_number_to_propose);
//         forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2a?
//         ensures pkt in ts.t_environment.sentPackets
//         forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2b?
//         ensures pkt in ts.t_environment.sentPackets
//         {
//             if pkt !in ts.t_environment.sentPackets {
//                 var ios := UntagLIoOpSeq(tios);
//                 var sent_packets := ExtractSentPacketsFromIos(ios);
//                 assert nextActionIndex != 0;
//                 assert LReplicaNoReceiveNext(ls.v.replica, nextActionIndex, ls'.v.replica, ios);
//                 assert forall p | p in sent_packets :: !p.msg.RslMessage_2b?;
//                 forall io | io in tios && io.LIoOpSend?
//                 ensures !io.s.msg.v.RslMessage_2b? {}
//                 assert false;
//             }
//         }
//         assert Before_2b_Sent_Invariant(ts', opn);
//     }
// }


// lemma Before2b_to_Before2b_NonReceive_NoNewReply(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
//     requires RslAssumption(ts, opn) && RslConsistency(ts)
//     requires RslAssumption(ts', opn) && RslConsistency(ts')
//     requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
//     requires TimestampedRslNext(ts, ts')
//     requires !TimestampedRslNextEnvironment(ts, ts')
//     requires RslPerfInvariant(ts, opn)
//     requires Before_2b_Sent_Invariant(ts, opn)
//     requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
//     requires ts.t_replicas[idx].v.nextActionIndex != 0
//     ensures forall pkt | pkt in ts'.t_environment.sentPackets :: !IsNewReplyPacket(ts', pkt) 
// {   
//     var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
//     var nextActionIndex, nextActionIndex' := ls.v.nextActionIndex, ls'.v.nextActionIndex;
//     forall pkt | pkt in ts'.t_environment.sentPackets
//     ensures !IsNewReplyPacket(ts', pkt) 
//     {
//         if IsNewReplyPacket(ts', pkt) {
//             assert pkt !in ts.t_environment.sentPackets;
//             if nextActionIndex == 6 {
//                 if idx == 1 {
//                     assert ls.v.replica.executor.next_op_to_execute == OutstandingOpUnknown();
//                     var ios := UntagLIoOpSeq(tios);
//                     var sent_packets := ExtractSentPacketsFromIos(ios);
//                     assert sent_packets == [];
//                     assert forall io | io in tios :: !io.LIoOpSend?;
//                 } else {
//                     assert ReplicasDistinct(ts.constants.config.replica_ids, 1, idx);
//                     forall io | io in tios && io.LIoOpSend?
//                     ensures io.s.src != ts'.constants.config.replica_ids[1] {}
//                 }
//                 assert false;
//             } else {
//                 var ios := UntagLIoOpSeq(tios);
//                 var sent_packets := ExtractSentPacketsFromIos(ios);
//                 assert nextActionIndex != 0;
//                 assert forall p | p in sent_packets :: !p.msg.RslMessage_Reply?;
//                 forall io | io in tios && io.LIoOpSend?
//                 ensures !io.s.msg.v.RslMessage_Reply? {}
//                 assert false;
//             }
//         }
//     }
// }






// lemma Before2b_to_Before2b_Receive(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
//     requires RslAssumption(ts, opn) && RslConsistency(ts)
//     requires RslAssumption(ts', opn) && RslConsistency(ts')
//     requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
//     requires TimestampedRslNext(ts, ts')
//     requires !TimestampedRslNextEnvironment(ts, ts')
//     requires RslPerfInvariant(ts, opn)
//     requires Before_2b_Sent_Invariant(ts, opn)
//     requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
//     requires ts.t_replicas[idx].v.nextActionIndex == 0
//     requires !tios[0].LIoOpTimeoutReceive?
//     requires !tios[0].r.msg.v.RslMessage_2a?
//     ensures Before_2b_Sent_Invariant(ts', opn)
// {
//     reveal_ExtractSentPacketsFromIos();
//     reveal_UntagLIoOpSeq();
//     var ios := UntagLIoOpSeq(tios);
//     var sent_packets := ExtractSentPacketsFromIos(ios);
//     assert forall p | p in sent_packets :: !p.msg.RslMessage_2b?;
//     forall pkt | pkt in ts'.t_environment.sentPackets && IsNew2bPacket(pkt, opn) 
//     ensures pkt in ts.t_environment.sentPackets {
//         if pkt !in ts.t_environment.sentPackets {
//             assert UntagLPacket(pkt) in sent_packets;
//             assert false;
//         }
//     }
//     forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2a?
//     ensures pkt in ts.t_environment.sentPackets
//     {}
//     assert opn == ts'.t_replicas[1].v.replica.executor.ops_complete;   // Can be changed by LExecutorProcessAppStateSupply packet
//     assert BalLt(ts'.t_replicas[1].v.replica.learner.max_ballot_seen, Ballot(1, 1)); 
//     forall pkt | pkt in ts'.t_environment.sentPackets && IsNewReplyPacket(ts', pkt) 
//     ensures pkt in ts.t_environment.sentPackets
//     {}

//     forall pkt | pkt in ts'.undeliveredPackets && IsNew2aPacket(pkt, opn)
//     ensures TimeLe(pkt.msg.ts, TimeBound2aDeliveryPost())
//     {
//         if pkt !in ts.undeliveredPackets {
//             forall io | io in tios && io.LIoOpSend?
//             ensures !io.s.msg.v.RslMessage_2a?
//             assert false;
//         }
//     }

//     assert Before_2b_Sent_Invariant(ts', opn);
// }


}