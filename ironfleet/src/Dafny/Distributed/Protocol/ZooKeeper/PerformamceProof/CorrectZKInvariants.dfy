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
include "ProtocolInvariants.dfy"
include "PerformancePredicates.dfy"


/* This module contains invariants that have nothing to do with performance */
module Zookeeper_CorrectInvariants {
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
        && !pkt.msg.v.SyncSNAP?
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
    && Leader_Have_Correct_Peer_Zxid(tls)
    && Leader_Only_Send_SyncDiff(tls)
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

predicate Leader_Have_Correct_Peer_Zxid(tls:TLS_State) 
    requires DS_Config_Invariant(tls.config, tls)
    requires ZK_Config_Invariant(tls.config, tls)
    requires |tls.initialZkdbState| == |tls.config|
{
    var target_db := tls.initialZkdbState[0];
    var target_zxid := getLastLoggedZxid(target_db);
    var l := tls.t_servers[tls.config[0]].v.leader;
    forall h | h in l.handlers ::
        l.handlers[h].peerLastZxid != NullZxid ==> l.handlers[h].peerLastZxid == target_zxid
}

predicate Leader_Only_Send_SyncDiff(tls:TLS_State) 
    requires DS_Config_Invariant(tls.config, tls)
    requires ZK_Config_Invariant(tls.config, tls)
    requires |tls.initialZkdbState| == |tls.config|
{
    var l := tls.t_servers[tls.config[0]].v.leader;
    forall h | h in l.handlers ::
        forall p | p in l.handlers[h].queuedPackets :: p.SyncDIFF?
}


/* Main theorem */
lemma theorem_ZK_Correct_Guarantee(config:Config, tlb:seq<TLS_State>, f:int)
    requires |tlb| > 1
    requires SeqIsUnique(config);
    requires ValidTLSBehavior(config, tlb, f)
    requires Performance_Assumption_EmptyDiff(tlb)
    requires forall i | 0 <= i < |tlb| :: Basic_Invariants(config, tlb[i])
    ensures forall i | 0 <= i < |tlb| :: |tlb[i].initialZkdbState| == |config|
    ensures forall i | 0 <= i < |tlb| :: EmptyDiff_Invariant_Inductive(tlb[i])
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
        assert EmptyDiff_Invariant_Inductive(tls');
        i := i + 1;
    }
}
}