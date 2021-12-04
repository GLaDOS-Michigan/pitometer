include "Phase2Proof.i.dfy"

module Rs2Phase2Proof_Helper_2 {
import opened RslPhase2Proof_postFail_i

/* WARNING: this file a timeout of 50s to verify */


lemma Before2b_to_After2b(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>,
    rs:RslState, rs':RslState, iops:seq<RslIo>
) 
    requires rs == UntimestampRslState(ts)
    requires rs' == UntimestampRslState(ts')
    requires iops == UntagLIoOpSeq(tios);
    requires RslAssumption(ts, opn) && RslConsistency(ts)
    requires RslAssumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires AlwaysInvariant(ts', opn)
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires RslPerfInvariant(ts, opn)
    requires Before_2b_Sent_Invariant(ts, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires LReplicaNextProcessPacket(rs.replicas[idx].replica, rs'.replicas[idx].replica, iops);
    requires ts.t_replicas[idx].v.nextActionIndex == 0
    requires |tios| > 0 && tios[0].LIoOpReceive? && tios[0].r.msg.v.RslMessage_2a? 
    requires iops[0].r.src in rs.replicas[idx].replica.acceptor.constants.all.config.replica_ids
    requires BalLeq(rs.replicas[idx].replica.acceptor.max_bal, iops[0].r.msg.bal_2a)
    requires LeqUpperBound(iops[0].r.msg.opn_2a, rs.replicas[idx].replica.acceptor.constants.all.params.max_integer_val)
    requires iops[0].r.msg.bal_2a == Ballot(1, 1)
    ensures Before_2b_Sent_Invariant(ts', opn) || After_2b_Sent_Invariant(ts', opn)
{
    var r, r' := rs.replicas[idx].replica, rs'.replicas[idx].replica;
    var m := iops[0].r.msg;
    var sent_packets := ExtractSentPacketsFromIos(iops);
    assert LAcceptorProcess2a(r.acceptor, r'.acceptor, iops[0].r, sent_packets);
    var msg2b := RslMessage_2b(m.bal_2a, m.opn_2a, m.val_2a);
    assert LBroadcastToEveryone(r.acceptor.constants.all.config, r.acceptor.constants.my_index, msg2b, sent_packets);
    assert forall p | p in sent_packets :: LIoOpSend(p) in iops;
    if m.opn_2a == opn {
        var pkt_witness := sent_packets[0];
        assert LIoOpSend(pkt_witness) in iops;
        assert After_2b_Sent_Invariant(ts', opn);          
    } else {
        assert forall p | p in sent_packets && p.msg.RslMessage_2b? :: p.msg.opn_2b != opn;
        forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2b?
        ensures BalLeq(pkt.msg.v.bal_2b, Ballot(1, 1))
        {}
        assert Before_2b_Sent_Invariant(ts', opn);    
    }
}


/* Proof that a Before_2b_Sent state transitions to a Before_2b_Sent state or 
* After_2b_Sent state */
lemma Before2b_to_MaybeAfter2b_Process2a(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts, opn) && RslConsistency(ts)
    requires RslAssumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires AlwaysInvariant(ts', opn)
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires ts.t_replicas[idx].v.nextActionIndex == 0
    requires tios[0].r.msg.v.RslMessage_2a?
    requires RslPerfInvariant(ts, opn)
    
    requires Before_2b_Sent_Invariant(ts, opn)
    ensures Before_2b_Sent_Invariant(ts', opn) || After_2b_Sent_Invariant(ts', opn)
{
    var nextActionIndex := ts.t_replicas[idx].v.nextActionIndex;

    // From this point on, nextActionIndex == 0
    var idx_s, idx_s', ios := UntimestampRslState(ts), UntimestampRslState(ts'), UntagLIoOpSeq(tios);
    var idx_r, idx_r' := idx_s.replicas[idx].replica, idx_s'.replicas[idx].replica;
    var sent_packets := ExtractSentPacketsFromIos(ios);
    var m := ios[0].r.msg;
    if ios[0].r.src in idx_r.acceptor.constants.all.config.replica_ids
       && BalLeq(idx_r.acceptor.max_bal, m.bal_2a)
       && LeqUpperBound(m.opn_2a, idx_r.acceptor.constants.all.params.max_integer_val)
    {
        if m.bal_2a == Ballot(1, 1) {
            Before2b_to_After2b(ts, ts', opn, idx, tios, idx_s, idx_s', ios);
        } else {
            assert m.bal_2a == Ballot(1, 0);
            assert forall p | p in sent_packets && p.msg.RslMessage_2b? :: p.msg.bal_2b == Ballot(1, 0);
            assert forall p | p in sent_packets && p.msg.RslMessage_2b? :: p.msg.opn_2b == opn;
            assert forall p | p in sent_packets :: !p.msg.RslMessage_2a?;
            reveal_ExtractSentPacketsFromIos();
            reveal_UntagLIoOpSeq();
            forall pkt | pkt in ts'.undeliveredPackets && pkt.msg.v.RslMessage_2a? 
            ensures pkt in ts.undeliveredPackets
            {}
            assert All2bPackets_BalLeq_Opn(ts', Ballot(1, 0), opn);
            assert Before_2b_Sent_Invariant(ts', opn);
        }
    } else {
        assert sent_packets == [];
        reveal_ExtractSentPacketsFromIos();
        forall io | io in ios 
        ensures !io.LIoOpSend? 
        {}
        forall tio | tio in tios 
        ensures !tio.LIoOpSend? {
            if tio.LIoOpSend? {
                reveal_UntagLIoOpSeq();
                assert UntagLIoOp(tio) in ios;
                assert UntagLIoOp(tio).LIoOpSend?;
            }
        }
        forall p | p in ts'.t_environment.sentPackets 
        ensures p in ts.t_environment.sentPackets 
        {}
        assert Before_2b_Sent_Invariant(ts', opn);
    }
}


lemma After2b_to_After2b_NonLeaderAction(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts, opn) && RslConsistency(ts)
    requires RslAssumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires AlwaysInvariant(ts', opn)
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires RslPerfInvariant(ts, opn)
    requires idx != 1;
    requires After_2b_Sent_Invariant(ts, opn)
    ensures After_2b_Sent_Invariant(ts', opn)
{
    assert ts'.t_replicas[1] == ts.t_replicas[1];
    After2b_to_After2b_NonLeaderAction_2bBalOpn(ts, ts', opn, idx, tios);

    forall p | p in ts'.undeliveredPackets && IsNew2aPacket(p, opn)
    ensures p in ts.undeliveredPackets {
        if p !in ts.undeliveredPackets{
            reveal_ExtractSentPacketsFromIos();
            reveal_UntagLIoOpSeq();
            var sr, sr', ios := UntimestampRslState(ts), UntimestampRslState(ts'), UntagLIoOpSeq(tios);
            var sent_packets := ExtractSentPacketsFromIos(ios);
            forall p' | p' in sent_packets 
            ensures !p'.msg.RslMessage_2a?
            {}
            assert false;
        }
    }
    forall p | p in ts'.undeliveredPackets && p !in ts.undeliveredPackets
    ensures p.src == ts'.constants.config.replica_ids[idx]
    {}
    assert ReplicasDistinct(ts'.constants.config.replica_ids, 1, idx);
    forall p | p in ts'.undeliveredPackets && IsNewReplyPacket(ts', p)
    ensures p in ts.undeliveredPackets
    {}
}


lemma After2b_to_After2b_NonLeaderAction_2bBalOpn(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires RslAssumption(ts, opn) && RslConsistency(ts)
    requires RslAssumption(ts', opn) && RslConsistency(ts')
    requires PacketsBallotInvariant(ts) && PacketsBallotInvariant(ts')
    requires AlwaysInvariant(ts', opn)
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires RslPerfInvariant(ts, opn)
    requires idx != 1;
    requires After_2b_Sent_Invariant(ts, opn)
    ensures forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2b? :: BalLeq(pkt.msg.v.bal_2b, Ballot(1, 1)) && pkt.msg.v.opn_2b == opn
{
    // Any 2b sent must have opn and ballot matching existing 2a's
    forall pkt | pkt in ts'.t_environment.sentPackets && pkt.msg.v.RslMessage_2b?
    ensures BalLeq(pkt.msg.v.bal_2b, Ballot(1, 1)) && pkt.msg.v.opn_2b == opn
    {
        var sr, sr', ios := UntimestampRslState(ts), UntimestampRslState(ts'), UntagLIoOpSeq(tios);
        var r, r' := sr.replicas[idx].replica, sr'.replicas[idx].replica;
        if sr.replicas[idx].nextActionIndex == 0 {
            var msg := tios[0].r.msg;
            if msg.v.RslMessage_2a? {
                if tios[0].r.src in r.acceptor.constants.all.config.replica_ids
                    && BalLeq(r.acceptor.max_bal, msg.v.bal_2a)
                    && LeqUpperBound(msg.v.opn_2a, r.acceptor.constants.all.params.max_integer_val)
                {
                    var sent_packets := ExtractSentPacketsFromIos(ios);
                    forall sp | sp in sent_packets 
                    ensures  && sp.msg.RslMessage_2b? 
                        && sp.msg.opn_2b == msg.v.opn_2a
                        && sp.msg.bal_2b == msg.v.bal_2a
                    {}
                } else {
                    assert forall p | p in ts'.t_environment.sentPackets && p.msg.v.RslMessage_2b? 
                    :: p in ts.t_environment.sentPackets;
                }
            } else {
                var sent_packets := ExtractSentPacketsFromIos(ios);
                forall p | p in sent_packets 
                ensures !p.msg.RslMessage_2b? {}
            }
        } else {
            var sent_packets := ExtractSentPacketsFromIos(ios);
            forall p | p in sent_packets 
            ensures !p.msg.RslMessage_2b?
            {}
            assert forall p | p in ts'.t_environment.sentPackets && p.msg.v.RslMessage_2b? 
            :: p in ts.t_environment.sentPackets;
        }
    }
}

}