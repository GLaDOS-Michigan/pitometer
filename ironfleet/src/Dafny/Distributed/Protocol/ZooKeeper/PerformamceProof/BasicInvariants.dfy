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


/* This module contains invariants that have nothing to do with performance */
module Zookeeper_BasicInvariants {
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


predicate Basic_Invariants(config:Config, tls:TLS_State) {
    && DS_Config_Invariant(config, tls)
    && ZK_Config_Invariant(config, tls)
    && Leader_QueuedPackets_Invariant(config, tls)
}


lemma lemma_Basic_Invariants(config:Config, tlb:seq<TLS_State>, f:int) 
    requires SeqIsUnique(config);
    requires ValidTLSBehavior(config, tlb, f)
    ensures forall i | 0 <= i < |tlb| :: Basic_Invariants(config, tlb[i]) 
{
    lemma_DS_Config_Invariant_Proof(config, tlb, f);
    lemma_ZK_Config_Invariant_Proof(config, tlb, f);
    lemma_Leader_QueuedPackets_Invariant_Proof(config, tlb, f);
}


/*****************************************************************************************
/                                   DSConfigInvariant                                    *
*****************************************************************************************/


/* config has same endpoints as tls.servers and each node's id is its index in config */
predicate DS_Config_Invariant(config:Config, tls:TLS_State) {
    && |config| > 0
    && (forall ep :: ep in config <==> ep in tls.t_servers)
    && (forall ep | ep in tls.t_servers :: config == if tls.t_servers[ep].v.FollowerPeer? then tls.t_servers[ep].v.follower.config else tls.t_servers[ep].v.leader.config)
    && (forall i | 0 <= i < |config| :: 
            i == if tls.t_servers[config[i]].v.FollowerPeer? then tls.t_servers[config[i]].v.follower.my_id else tls.t_servers[config[i]].v.leader.my_id
    )
    && tls.t_environment.config == config
}

lemma lemma_DS_Config_Invariant_Proof(config:Config, tlb:seq<TLS_State>, f:int)
    requires SeqIsUnique(config);
    requires ValidTLSBehavior(config, tlb, f)
    ensures forall i | 0 <= i < |tlb| :: DS_Config_Invariant(config, tlb[i])
{
    assert DS_Config_Invariant(config, tlb[0]);
    var i := 0;
    while i < |tlb| - 1 
        decreases |tlb| - i
        invariant 0 <= i < |tlb|
        invariant forall k | 0 <= k <= i :: DS_Config_Invariant(config, tlb[k])
    {
        var tls, tls' := tlb[i], tlb[i+1];
        assert TLS_Next(tls, tls');
        assert DS_Config_Invariant(config, tls');
        i := i + 1;
    }
}


/*****************************************************************************************
/                                   ZKConfigInvariant                                    *
*****************************************************************************************/

/* config[0] is the leader and everyone else are followers */
predicate ZK_Config_Invariant(config:Config, tls:TLS_State) 
    requires DS_Config_Invariant(config, tls)
{   
    // Only config[0] is the leader
    && (forall i | 0 <= i < |config| :: if i == 0 then tls.t_servers[config[i]].v.LeaderPeer? else tls.t_servers[config[i]].v.FollowerPeer?)
    // Every follower has 0 as the leader
    && (forall ep | ep in tls.t_servers && tls.t_servers[ep].v.FollowerPeer? :: tls.t_servers[ep].v.follower.leader_id == 0)
    // Each follower has a corresponding learner handler
    && |tls.t_servers[config[0]].v.leader.handlers| == |config| - 1
    && (forall follower_id | 1 <= follower_id < |config| :: tls.t_servers[config[0]].v.leader.handlers[follower_id-1].follower_id == follower_id)
}


lemma lemma_ZK_Config_Invariant_Proof(config:Config, tlb:seq<TLS_State>, f:int)
    requires SeqIsUnique(config);
    requires ValidTLSBehavior(config, tlb, f)
    requires forall i | 0 <= i < |tlb| :: DS_Config_Invariant(config, tlb[i])
    ensures forall i | 0 <= i < |tlb| :: ZK_Config_Invariant(config, tlb[i])
{
    assert ZK_Config_Invariant(config, tlb[0]);
    var i := 0;
    while i < |tlb| - 1 
        decreases |tlb| - i
        invariant 0 <= i < |tlb|
        invariant forall k | 0 <= k <= i :: ZK_Config_Invariant(config, tlb[k])
    {
        var tls, tls' := tlb[i], tlb[i+1];
        assert TLS_Next(tls, tls');
        assert ZK_Config_Invariant(config, tls');
        i := i + 1;
    }
}


/*****************************************************************************************
/                               Facts about the environment                              *
*****************************************************************************************/


/* SentPackets set is monotone increasing */
predicate SentPacketsSet_Monotone(tls:TLS_State, tls':TLS_State, ep:EndPoint, ios:seq<TZKIo>) 
    requires ep in tls.t_servers
    requires TLS_NextOneServer(tls, tls', ep, ios)
    requires LEnvironment_Next(tls.t_environment, tls'.t_environment)
{
    var new_sent := IosSeqToSentSet(ios);
    tls'.t_environment.sentPackets == tls.t_environment.sentPackets + new_sent
}


lemma lemma_SentPacketsSet_Monotone_Induction(config:Config, tls:TLS_State, tls':TLS_State)
    requires TLS_Next(tls, tls')
    requires Basic_Invariants(config, tls) && Basic_Invariants(config, tls')
    ensures forall ep, ios | ep in tls.t_servers && TLS_Next(tls, tls') && TLS_NextOneServer(tls, tls', ep, ios)
    :: SentPacketsSet_Monotone(tls, tls', ep, ios)
{
    // TODO
    assume false;
}



/*****************************************************************************************
/                                 Facts about the leader                                 *
*****************************************************************************************/


/* QueuedPackets for a LearnerHandler only contains the approriate messages */
predicate QueuedPackets_Only_Contains_LeaderMessages(q: seq<ZKMessage>) {
    forall msg | msg in q ::
        || msg.SyncTRUNC?
        || msg.SyncSNAP?
        || msg.SyncDIFF?
        || msg.Commit?
}

lemma lemma_LearnerHandler_QueuedPackets_Induction(s:LearnerHandler, s':LearnerHandler, g:LeaderGlobals, g':LeaderGlobals, ios:seq<ZKIo>) 
    requires LearnerHandlerNext(s, s', g, g', ios);
    requires QueuedPackets_Only_Contains_LeaderMessages(s.queuedPackets);
    ensures QueuedPackets_Only_Contains_LeaderMessages(s'.queuedPackets);
{}

/* QueuedPackets for a LearnerHandler only contains the approriate messages */
predicate Leader_QueuedPackets_Invariant(config:Config, tls:TLS_State) {
    forall ep | ep in tls.t_servers ::
        tls.t_servers[ep].v.LeaderPeer? ==> (
            forall lh | lh in tls.t_servers[ep].v.leader.handlers
            :: QueuedPackets_Only_Contains_LeaderMessages(lh.queuedPackets)
    )
}

lemma lemma_Leader_QueuedPackets_Invariant_Proof(config:Config, tlb:seq<TLS_State>, f:int)
    requires SeqIsUnique(config);
    requires ValidTLSBehavior(config, tlb, f)
    ensures forall i | 0 <= i < |tlb| :: Leader_QueuedPackets_Invariant(config, tlb[i]) 
{
    var handlers := tlb[0].t_servers[config[0]].v.leader.handlers;
    assert forall i | 1 <= i < |config| :: LearnerHandlerInit(handlers[i-1], 0, i, config);
    forall i | 0 <= i < |handlers| 
    ensures LearnerHandlerInit(handlers[i], 0, i+1, config)
    { 
        var k := i + 1;
        assert 1 <= k < |config|;
        assert LearnerHandlerInit(handlers[k-1], 0, k, config);
    }
    var i := 0;
    while i < |tlb| - 1 
        decreases |tlb| - i
        invariant 0 <= i < |tlb|
        invariant forall k | 0 <= k <= i :: Leader_QueuedPackets_Invariant(config, tlb[k])
    {
        var tls, tls' := tlb[i], tlb[i+1];
        assert Leader_QueuedPackets_Invariant(config, tls');
        i := i + 1;
    }
}


}
