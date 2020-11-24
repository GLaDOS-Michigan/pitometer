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
// include "PerformancePredicates.dfy"


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
// import opened Zookeeper_PerformancePredicates


/* config has same endpoints as tls.servers and each node's id is its index in config */
predicate ConfigInvariant(config:Config, tls:TLS_State) {
    && (forall ep :: ep in config <==> ep in tls.t_servers)
    && (forall ep | ep in tls.t_servers :: config == if tls.t_servers[ep].v.FollowerPeer? then tls.t_servers[ep].v.follower.config else tls.t_servers[ep].v.leader.config)
    && (forall i | 0 <= i < |config| :: 
            i == if tls.t_servers[config[i]].v.FollowerPeer? then tls.t_servers[config[i]].v.follower.my_id else tls.t_servers[config[i]].v.leader.my_id
    )
}


lemma lemma_Config_Invariant_Proof(config:Config, tlb:seq<TLS_State>, f:int)
    requires SeqIsUnique(config);
    requires ValidTLSBehavior(config, tlb, f)
    ensures forall i | 0 <= i < |tlb| :: ConfigInvariant(config, tlb[i])
{
    assert ConfigInvariant(config, tlb[0]);
    var i := 0;
    while i < |tlb| - 1 
        decreases |tlb| - i
        invariant 0 <= i < |tlb|
        invariant forall k | 0 <= k <= i :: ConfigInvariant(config, tlb[k])
    {
        var tls, tls' := tlb[i], tlb[i+1];
        assert TLS_Next(tls, tls');
        assert ConfigInvariant(config, tls');
        i := i + 1;
    }
}


}