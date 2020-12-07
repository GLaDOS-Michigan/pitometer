include "../../../Impl/Common/SeqIsUniqueDef.i.dfy"
include "../../../Common/Framework/EnvironmentTCP.s.dfy"
include "../Timestamps/TimestampedType.dfy"
include "TimestampedLS.dfy"

include "../Types.dfy"
include "../DistributedSystem.dfy"
include "../ZKEnvironment.dfy"
include "../ZKDatabase.dfy"
include "../Follower.dfy"
include "../Leader.dfy"
include "../LearnerHandler.dfy"
include "Definitions.dfy"
include "PerformancePredicates.dfy"
include "ProtocolInvariants.dfy"
include "Commons.dfy"
include "ProofEmptyDiff_Invariants.dfy"
include "EmptyDiffInvariants.dfy"


module Zookeeper_PerformanceProof {
import opened Common__SeqIsUniqueDef_i
import opened ZKTimestamp
import opened ZooKeeper_Types
import opened ZooKeeper_Environment
import opened ZooKeeper_DistributedSystem
import opened EnvironmentTCP_s
import opened ZooKeeper_ZKDatabase
import opened ZooKeeper_Follower
import opened ZooKeeper_Leader
import opened ZooKeeper_LearnerHandler
import opened Zookeeper_Performance_Definitions
import opened ZooKeeper_TimestampedDistributedSystem
import opened Zookeeper_PerformancePredicates
import opened Zookeeper_ProtocoIInvariants
import opened Zookeeper_Commons
import opened Zookeeper_PerformanceProof_Invariants
import opened Zookeeper_EmptyDiffInvariants


/* Main theorem */
lemma theorem_ZK_Performance_Guarantee(config:Config, tlb:seq<TLS_State>, f:int)
    requires SeqIsUnique(config);
    requires ValidTLSBehavior(config, tlb, f)
    requires Performance_Assumption_EmptyDiff(tlb)
    // Protocol guarantees
    ensures forall i | 0 <= i < |tlb| :: DS_Config_Invariant(config, tlb[i])
    ensures forall i | 0 <= i < |tlb| :: ZK_Config_Invariant(config, tlb[i])
    ensures forall i | 0 <= i < |tlb| :: Quorums_Size_Invariant(tlb[i])
    // Performance guarantees
    ensures Performance_Guarantee_EmptyDiff(config, tlb)
{
    
    lemma_Basic_Invariants(config, tlb, f);
    // TODO: These will later be replaced by lemmas proving these
    assume forall i | 0 <= i < |tlb| :: EmptyDiff_Invariant(tlb[i]);

    theorem_ZK_Sync_Performance_Guarantee(config, tlb, f);
    assert LS_Performance_Guarantee_EmptyDiff(tlb[0]);

    var i := 0;
    while i < |tlb| - 1 
        decreases |tlb| - i
        invariant 0 <= i < |tlb|
        invariant forall k | 0 <= k <= i :: LS_Performance_Guarantee_EmptyDiff(tlb[k])
    {
        var t, t' := tlb[i], tlb[i+1];
        theorem_ZK_Performance_Guarantee_Induction(config, t, t', f);
        i := i + 1;
    }
}


lemma theorem_ZK_Performance_Guarantee_Induction(config:Config, tls:TLS_State, tls':TLS_State, f:int)
    requires SeqIsUnique(config);
    requires TLS_Next(tls, tls');
    requires General_LS_Performance_Assumption(tls) && General_LS_Performance_Assumption(tls') 
    // Protocol guarantees
    requires Basic_Invariants(config, tls) && Basic_Invariants(config, tls')
    requires EmptyDiff_Invariant(tls) && EmptyDiff_Invariant(tls')
    // Induction hypothesis
    requires Sync_Messages_Invariant(tls) && Sync_Messages_Invariant(tls')
    requires Sync_Follower_Invariant(tls) && Sync_Follower_Invariant(tls')
    requires Sync_Leader_PreQuorum_Invariant(tls) && Sync_Leader_PreQuorum_Invariant(tls')
    requires LS_Performance_Guarantee_EmptyDiff(tls)
    // Performance guarantees
    ensures LS_Performance_Guarantee_EmptyDiff(tls')
{
    var f, m := tls.f, |config|;
    assert TLS_Next(tls, tls');
    var actor, tios:seq<TZKIo> :| actor in tls.t_servers && TLS_NextOneServer(tls, tls', actor, tios);
    if actor == config[0] {
        var st, st' := tls.t_servers[config[0]], tls'.t_servers[config[0]];
        var s, s', ios := st.v.leader, st'.v.leader, UntagLIoOpSeq(tios);
        if !IsQuorum(m, s.globals.ackSet) {
            if IsQuorum(m, s'.globals.ackSet) {
                assert StepSingleHandler_Rcv(s, s', ios);
                var fid := ios[0].r.sender_index;
                var h, h', g, g' := s.handlers[fid], s'.handlers[fid], s.globals, s'.globals;
                assert LearnerHandlerNext(h, h', g, g', ios);
                assert ZooKeeper_LearnerHandler.ProcessAck(h, h', g, g', ios);
                if !g.zkdb.isRunning {
                    lemma_Math_Mult_b();
                    lemma_Math_Inequalities_Mult();
                    var pkt := tios[0].r;
                    assert pkt.msg.v.Ack?;
                    assert pkt.msg.ts <= Ack_Message_ts_Formula(tls.f, pkt.msg.v.serial);
                    if |g.ackSet| <= 1 {
                        lemma_Math_Inequalities_Mult();
                        lemma_Math_Inequalities_CommonMult(ProcEpAck, |g.electingFollowers|-1, f);
                        lemma_Math_Inequalities_CommonMult(ProcEpAck, g.procEpCount, f);
                        lemma_Math_Inequalities_CommonMult(PreSync, g.prepCount, f);
                        lemma_Math_Inequalities_CommonMult(Sync, g.nextSerialSync, f);
                        lemma_Math_Inequalities_CommonMult(Sync, g.nextSerialNL, f);
                        lemma_Math_Inequalities_CommonMult(Sync, pkt.msg.v.serial + 1, f);
                        var lt_max := AckEpoch_Message_PreQuorum_ts_Formula(f, f-1)
                                + ProcEpAck * f
                                + ProcEpAck * f
                                + PreSync * f    // max possible PrepSyncs done before I was sent
                                + Sync * f  // max possible Syncs sent before I was sent
                                + Sync * f;
                        assert st.ts <= lt_max;
                        var pkt_max := lt_max + D + ProcSyncI + ProcSync + D;
                        assert pkt.msg.ts  <= pkt_max;
                        lemma_Math_MaxOfInequalities(st.ts, pkt.msg.ts, lt_max, pkt_max);
                        assert st'.ts <= pkt_max + ProcAck;
                        assert st'.ts <= Ack_Message_ts_Formula(f, f-1) + ProcAck * f;
                    } else {
                        // assume false;
                        assert st.ts <= Ack_Message_ts_Formula(f, f-1) + ProcAck * (|g.ackSet|-1);
                        assert st'.ts <= Ack_Message_ts_Formula(f, f-1) + ProcAck * f;
                    }
                }
            }
        } else {
            assert IsVerifiedQuorum(s.my_id, m, s.globals.ackSet);
            assert st'.ts == st.ts;
        }
    }
}


lemma theorem_ZK_Handshake_Performance_Guarantee(config:Config, tlb:seq<TLS_State>, f:int)
    requires SeqIsUnique(config);
    requires ValidTLSBehavior(config, tlb, f)
    requires Performance_Assumption_EmptyDiff(tlb)
    requires forall i | 0 <= i < |tlb| :: Basic_Invariants(config, tlb[i])
    ensures forall i | 0 <= i < |tlb| :: FollowerInit_Invariant(tlb[i])
    ensures forall i | 0 <= i < |tlb| :: Handshake_Follower_Invariant(tlb[i])
    ensures forall i | 0 <= i < |tlb| :: Handshake_Messages_Invariant(tlb[i])
    ensures forall i | 0 <= i < |tlb| :: Handshake_Leader_PreQuorum_Invariant(tlb[i])
{
    assert FollowerInit_Invariant(tlb[0]);
    assert Handshake_Follower_Invariant(tlb[0]);
    assert Handshake_Messages_Invariant(tlb[0]);
    assert Handshake_Leader_PreQuorum_Invariant(tlb[0]);
    var i := 0;
    while i < |tlb| - 1 
        decreases |tlb| - i
        invariant 0 <= i < |tlb|
        invariant forall k | 0 <= k <= i :: FollowerInit_Invariant(tlb[k])
        invariant forall k | 0 <= k <= i :: Handshake_Follower_Invariant(tlb[k])
        invariant forall k | 0 <= k <= i :: Handshake_Leader_PreQuorum_Invariant(tlb[k])
        invariant forall k | 0 <= k <= i :: Handshake_Messages_Invariant(tlb[k])
    {
        var tls, tls' := tlb[i], tlb[i+1];
        assert TLS_Next(tls, tls');
        lemma_FollowerInit_Invariant_Induction(config, tls, tls');
        lemma_Leader_ProcessEpAck_PreQuorum_Invariant_Induction(config, tls, tls');
        assert Handshake_Follower_Invariant(tls');
        assert Handshake_Messages_Invariant(tls');
        assert Handshake_Leader_PreQuorum_Invariant(tls');
        assert forall k | 0 <= k <= i+1 :: 
            && Handshake_Leader_PreQuorum_Invariant(tlb[k])
            && Handshake_Messages_Invariant(tlb[k])
            && Handshake_Follower_Invariant(tlb[k]);
        i := i + 1;
        assert tls' == tlb[i];
    }
}


lemma theorem_ZK_Sync_Performance_Guarantee(config:Config, tlb:seq<TLS_State>, f:int)
    requires SeqIsUnique(config);
    requires ValidTLSBehavior(config, tlb, f)
    requires Performance_Assumption_EmptyDiff(tlb)
    requires forall i | 0 <= i < |tlb| :: Basic_Invariants(config, tlb[i])
    requires forall i | 0 <= i < |tlb| :: EmptyDiff_Invariant(tlb[i])
    ensures forall i | 0 <= i < |tlb| :: Sync_Messages_Invariant(tlb[i])
    ensures forall i | 0 <= i < |tlb| :: Sync_Follower_Invariant(tlb[i])
    ensures forall i | 0 <= i < |tlb| :: Sync_Leader_PreQuorum_Invariant(tlb[i])
{
    theorem_ZK_Handshake_Performance_Guarantee(config, tlb, f);
    assert Sync_Messages_Invariant(tlb[0]);
    assert Sync_Follower_Invariant(tlb[0]);
    assert Sync_Leader_PreQuorum_Invariant(tlb[0]);
    var i := 0;
    while i < |tlb| - 1 
        decreases |tlb| - i
        invariant 0 <= i < |tlb|
        invariant forall k | 0 <= k <= i :: Sync_Messages_Invariant(tlb[k])
        invariant forall k | 0 <= k <= i :: Sync_Follower_Invariant(tlb[k])
        invariant forall k | 0 <= k <= i :: Sync_Leader_PreQuorum_Invariant(tlb[k])
    {
        var tls, tls' := tlb[i], tlb[i+1];
        assert TLS_Next(tls, tls');
        lemma_Leader_ProcessAck_PreQuorum_Invariant_Induction(config, tls, tls');
        i := i + 1;
    }
}

}
