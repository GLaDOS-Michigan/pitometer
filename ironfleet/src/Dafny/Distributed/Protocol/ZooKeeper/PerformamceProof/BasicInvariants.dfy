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
}
