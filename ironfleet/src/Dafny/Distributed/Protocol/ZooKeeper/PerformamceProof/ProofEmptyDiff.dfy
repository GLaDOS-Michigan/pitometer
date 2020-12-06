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
    ensures Performance_Guarantee_EmptyDiff(tlb)
{
    // TODO
    assume false;
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
