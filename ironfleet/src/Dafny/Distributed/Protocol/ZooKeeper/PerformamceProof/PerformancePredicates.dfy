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


module Zookeeper_PerformancePredicates {
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
*                                      Assumptions                                       *
*****************************************************************************************/


predicate General_LS_Performance_Assumption(tls:TLS_State) {
    // The only nodes that take steps are in the ring
    && (tls.t_environment.nextStep.LEnvStepHostIos? 
        ==> 
       tls.t_environment.nextStep.actor in tls.t_servers)
    // No timeouts.
    && (tls.t_environment.nextStep.LEnvStepHostIos? 
        ==>
        (forall io | io in tls.t_environment.nextStep.ios :: !io.LIoOpTimeoutReceive?))
}


/* Performance assumption Empty Diff */
predicate Performance_Assumption_EmptyDiff(tlb:seq<TLS_State>) {
    && |tlb| > 0
    && (forall i | 0 <= i < |tlb| :: General_LS_Performance_Assumption(tlb[i]))
    && InitialZkdbState_EmptyDiff(tlb[0].initialZkdbState)
}


/*****************************************************************************************
*                                     Main Guarantee                                     *
*****************************************************************************************/


function Performance_Formula_EmptyDiff(config: Config) : Timestamp {
    var q := |config| / 2 + 1;
    SendFI +
    D + ProcFI * q +
    D + ProcLI +
    D + ProcEpAck * q +
    PreSync +
    Sync * 2 +
    D + ProcSync +
    D + ProcAck * q
}


/* Performance guarantee Empty Diff 
* Every leader in the RUNNING state has the specified performance formula */
predicate LS_Performance_Guarantee_EmptyDiff(tls:TLS_State) {   
    forall ep | 
        && ep in tls.t_servers 
        && tls.t_servers[ep].v.LeaderPeer? 
        && tls.t_servers[ep].v.leader.state == L_RUNNING 
    ::
        tls.t_servers[ep].ts >= Performance_Formula_EmptyDiff(tls.config)
}


predicate Performance_Guarantee_EmptyDiff(tlb:seq<TLS_State>){
    forall i | 0 <= i < |tlb| :: LS_Performance_Guarantee_EmptyDiff(tlb[i])
}


/*****************************************************************************************
*                             Handshake Phase Guarantees                                 *
*****************************************************************************************/


predicate FollowerInit_Invariant(tls:TLS_State){
    forall ep | ep in tls.t_servers :: (
        && tls.t_servers[ep].v.FollowerPeer? 
        && tls.t_servers[ep].v.follower.state == F_HANDSHAKE_A
        ==> 
        && tls.t_servers[ep].ts == TimeZero()
    )
}

function FollowerInfo_Message_Formula() : Timestamp {
    SendFI
}

/* Every FollowerInfo packet has ts specified by FollowerInfo_Message_Formula */
predicate FollowerInfo_Message_Invariant(tls:TLS_State) {
    forall pkt | pkt in tls.t_environment.sentPackets && pkt.msg.v.FollowerInfo?
    :: pkt.msg.ts == FollowerInfo_Message_Formula()
}


function ProcessFI_PreQuorum_Formula(connectingFollowers:set<nat>) : Timestamp {
    FollowerInfo_Message_Formula() + D + ProcLI * |connectingFollowers|
}


predicate ProcessFI_PreQuorum_Invariant(tls:TLS_State) {
    var n := |tls.config|;
    forall ep | ep in tls.t_servers :: (
        && tls.t_servers[ep].v.LeaderPeer? 
        && |tls.t_servers[ep].v.leader.globals.connectingFollowers| < (n/2) + 1
        ==> 
        && tls.t_servers[ep].ts == ProcessFI_PreQuorum_Formula(tls.t_servers[ep].v.leader.globals.connectingFollowers)
    )
}
}
