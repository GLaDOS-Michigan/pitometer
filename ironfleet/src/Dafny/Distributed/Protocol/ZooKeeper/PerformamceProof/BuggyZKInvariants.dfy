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
include "../LearnerHandler_Bug.dfy"
include "Definitions.dfy"
include "ProtocolInvariants.dfy"
include "PerformancePredicates_Bug.dfy"


/* This module contains invariants that have nothing to do with performance */
module Zookeeper_BuggyInvariants {
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
import opened Zookeeper_ProtocoIInvariants
import opened Zookeeper_PerformancePredicates


/*****************************************************************************************
*                                    Empty Diff facts                                    *
*****************************************************************************************/

predicate EmptyDiff_Invariant(tls:TLS_State) {
    forall pkt | pkt in tls.t_environment.sentPackets ::
        && !pkt.msg.v.SyncDIFF?
        && !pkt.msg.v.SyncTRUNC?
        && !pkt.msg.v.Commit?
}


predicate EmptyDiff_Invariant_Inductive(tls:TLS_State) 
    requires DS_Config_Invariant(tls.config, tls)
    requires ZK_Config_Invariant(tls.config, tls)
    requires |tls.initialZkdbState| == |tls.config|
{
    && EmptyDiff_Invariant(tls)
    && Dbs_Are_Identical(tls)
    && All_Messages_Have_Same_Zxid(tls)
    && Leader_Have_Bad_Zxid(tls)
    && Leader_Only_Send_SyncSnap(tls)
}

predicate Dbs_Are_Identical(tls:TLS_State) 
    requires DS_Config_Invariant(tls.config, tls)
    requires ZK_Config_Invariant(tls.config, tls)
    requires |tls.initialZkdbState| == |tls.config|
{
    var target_db := tls.initialZkdbState[0];
    var target_zxid := getLastLoggedZxid(target_db);
    tls.t_servers[tls.config[0]].v.leader.globals.zkdb.commitLog == target_db.commitLog
    && (forall ep | ep in tls.t_servers && tls.t_servers[ep].v.FollowerPeer?
        ::
        tls.t_servers[ep].v.follower.zkdb.commitLog == target_db.commitLog)
}

predicate All_Messages_Have_Same_Zxid(tls:TLS_State) 
    requires DS_Config_Invariant(tls.config, tls)
    requires ZK_Config_Invariant(tls.config, tls)
    requires |tls.initialZkdbState| == |tls.config|
{
    var target_db := tls.initialZkdbState[0];
    var target_zxid := getLastLoggedZxid(target_db);
    forall pkt | pkt in tls.t_environment.sentPackets ::
    && (pkt.msg.v.FollowerInfo? ==> pkt.msg.v.latestZxid == target_zxid) 
    && (pkt.msg.v.AckEpoch? ==> pkt.msg.v.lastLoggedZxid == target_zxid) 
}

predicate Leader_Have_Bad_Zxid(tls:TLS_State) 
    requires DS_Config_Invariant(tls.config, tls)
    requires ZK_Config_Invariant(tls.config, tls)
    requires |tls.initialZkdbState| == |tls.config|
{
    var target_db := tls.initialZkdbState[0];
    var target_zxid := getLastLoggedZxid(target_db);
    var l := tls.t_servers[tls.config[0]].v.leader;
    forall h | h in l.handlers ::
        && (l.handlers[h].peerLastZxid != NullZxid ==> l.handlers[h].peerLastZxid == target_zxid)
        && (l.handlers[h].peerLastZxid != NullZxid ==> l.globals.currZxid == Zxid(l.globals.leaderEpoch, 0))
        && (l.handlers[h].peerLastZxid != NullZxid ==> l.handlers[h].peerLastZxid != l.globals.currZxid)
}

predicate Leader_Only_Send_SyncSnap(tls:TLS_State) 
    requires DS_Config_Invariant(tls.config, tls)
    requires ZK_Config_Invariant(tls.config, tls)
    requires |tls.initialZkdbState| == |tls.config|
{
    var l := tls.t_servers[tls.config[0]].v.leader;
    forall h | h in l.handlers ::
        forall p | p in l.handlers[h].queuedPackets :: p.SyncSNAP?
}



/* Main theorem */
lemma theorem_ZK_Buggy_Guarantee(config:Config, tlb:seq<TLS_State>, f:int)
    requires |tlb| > 1
    requires SeqIsUnique(config);
    requires ValidTLSBehavior(config, tlb, f)
    requires Performance_Assumption_EmptyDiff(tlb)
    requires forall i | 0 <= i < |tlb| :: Basic_Invariants(config, tlb[i])
    ensures forall i | 0 <= i < |tlb| :: |tlb[i].initialZkdbState| == |config|
    ensures forall i | 0 <= i < |tlb| :: EmptyDiff_Invariant(tlb[i])
{
    assert |tlb[0].initialZkdbState| == |config|;
    assert EmptyDiff_Invariant_Inductive(tlb[0]);
    var i := 0;
    while i < |tlb| - 1 
        decreases |tlb| - i
        invariant 0 <= i < |tlb|
        invariant forall k | 0 <= k <= i :: |tlb[k].initialZkdbState| == |config|
        invariant forall k | 0 <= k <= i :: EmptyDiff_Invariant_Inductive(tlb[k])
    {   
        var tls, tls' := tlb[i], tlb[i+1];
        theorem_ZK_Buggy_Guarantee_Induction(config, tls, tls', f);
        assert EmptyDiff_Invariant_Inductive(tls');
        i := i + 1;
    }
}


/* Main theorem */
lemma theorem_ZK_Buggy_Guarantee_Induction(config:Config, tls:TLS_State, tls':TLS_State, f:int)
    requires TLS_Next(tls, tls')
    requires Basic_Invariants(config, tls) && Basic_Invariants(config, tls')
    requires |tls.initialZkdbState| == |config|
    requires EmptyDiff_Invariant_Inductive(tls)
    ensures |tls'.initialZkdbState| == |config|
    ensures EmptyDiff_Invariant_Inductive(tls')
{
    assert Dbs_Are_Identical(tls');
    assert All_Messages_Have_Same_Zxid(tls');
    assert Leader_Have_Bad_Zxid(tls');
    assert Leader_Only_Send_SyncSnap(tls');
    assert EmptyDiff_Invariant(tls');
}

}
