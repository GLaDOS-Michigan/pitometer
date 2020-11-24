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


predicate FollowerInfo_Message_Guarantee(tls:TLS_State)
{
    var perf_formula := SendFI;
    forall ep | ep in tls.t_environment.channels :: (
        forall pkt | pkt in tls.t_environment.channels[ep].channel
        :: pkt.msg.ts == perf_formula
    )
}

lemma lemma_FollowerInfo_Message_Guarantee(config:Config, tlb:seq<TLS_State>, f:int)
    requires SeqIsUnique(config);
    requires ValidTLSBehavior(config, tlb, f)
    requires Performance_Assumption_EmptyDiff(tlb)
    ensures forall i | 0 <= i < |tlb| :: FollowerInfo_Message_Guarantee(tlb[i])
{
    assume false;
}
}