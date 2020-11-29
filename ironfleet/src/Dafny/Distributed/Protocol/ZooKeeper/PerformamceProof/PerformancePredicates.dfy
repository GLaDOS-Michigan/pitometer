include "../../../Impl/Common/SeqIsUniqueDef.i.dfy"
include "../../../Common/Framework/EnvironmentTCP.s.dfy"
include "../Timestamps/TimestampedType.dfy"
include "TimestampedLS.dfy"
include "BasicInvariants.dfy"

include "../Types.dfy"
include "../DistributedSystem.dfy"
include "../ZKEnvironment.dfy"
include "../ZKDatabase.dfy"
include "../Follower.dfy"
include "../Leader.dfy"
include "../LearnerHandler.dfy"
include "Definitions.dfy"
include "Commons.dfy"


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
import opened Zookeeper_BasicInvariants
import opened Zookeeper_Commons


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
    assert q >= 0;
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
        && tls.t_servers[ep].dts == TimeZero()
        && tls.t_servers[ep].ts == TimeZero()
    )
}

function FollowerInfo_Message_ts_Formula() : Timestamp {
    SendFI + D
}

/* Every FollowerInfo packet has ts specified by FollowerInfo_Message_Formula */
predicate FollowerInfo_Message_ts_Invariant(tls:TLS_State) {
    forall pkt | pkt in tls.t_environment.sentPackets && pkt.msg.v.FollowerInfo?
    :: pkt.msg.ts == FollowerInfo_Message_ts_Formula()
}


function ProcessFI_PreQuorum_ts_Formula(l:TQuorumPeer) : Timestamp 
    requires l.v.LeaderPeer?
{
    var connectingFollowers := l.v.leader.globals.connectingFollowers;
    l.dts + ProcFI * |connectingFollowers|
}

function ProcessFI_PreQuorum_dts_Formula() : Timestamp {
    FollowerInfo_Message_ts_Formula()
}


/* Before the leader has a complete connectingFollowers quorum, its ts and dts are 
* specified by the respective formulas */
predicate ProcessFI_PreQuorum_Invariant(tls:TLS_State) 
    requires DS_Config_Invariant(tls.config, tls)
    requires ZK_Config_Invariant(tls.config, tls)
{
    var n := |tls.config|;
    var leaderTQP := tls.t_servers[tls.config[0]];
    !IsQuorum(|tls.config|, leaderTQP.v.leader.globals.connectingFollowers)
    ==> 
    && leaderTQP.ts <= ProcessFI_PreQuorum_ts_Formula(leaderTQP)
    && leaderTQP.dts <= ProcessFI_PreQuorum_dts_Formula()
}


function LeaderInfo_Message_PreQuorum_ts_Formula(tls:TLS_State, pkt:TimestampedLPacket<EndPoint, ZKMessage>) : Timestamp 
    requires pkt.msg.v.LeaderInfo?
{
    if pkt.msg.v.serial == 0 then
        // This is the first message sent. 
        SendFI + D 
        + ProcFI * |tls.config| 
        + ProcFI + D
    else 
        // This is (#serial+1)^th message sent
        // Before this, I may have sent #serial messages and processed their responses in sequence
        SendFI + D 
        + ProcFI * |tls.config| 
        + (ProcFI + D) * (pkt.msg.v.serial + 1)
        + ProcLI
        + (ProcEpAck + D) * pkt.msg.v.serial
}


function ProcessEpAck_PreQuorum_ts_Formula(tls:TLS_State, l:TQuorumPeer) : Timestamp 
    requires l.v.LeaderPeer?
{
    var electingFollowers := l.v.leader.globals.electingFollowers;
    if |electingFollowers| <= 1 then 
        l.dts + ProcFI * (l.v.leader.globals.nextSerialLI)
    else 
        l.dts  
        + (ProcEpAck)  // Add processing time
}


function ProcessEpAck_PreQuorum_dts_Formula(tls:TLS_State, l:TQuorumPeer) : Timestamp 
    requires l.v.LeaderPeer?
    requires |l.v.leader.globals.electingFollowers| <= l.v.leader.globals.nextSerialLI + 1
    requires |l.v.leader.globals.electingFollowers| >= 1
{
    var electingFollowers := l.v.leader.globals.electingFollowers;
    var n := |tls.config|;
    assert l.v.leader.globals.nextSerialLI - |electingFollowers| + 1 >= 0;
    if |electingFollowers| <= 1 then 
        SendFI + D 
        + ProcFI * n  // Received at most n FI's
    else 
        var res := 
            SendFI + D 
            + ProcFI * n   // Receiving all of the FI
            // First sending out the excess LI, 
            // And then for all the EpochAck in electingFollowers, I sent the LI and receive 
            // the EpochAck in sequence
            + ProcFI * (l.v.leader.globals.nextSerialLI - |electingFollowers| + 1)
            + (ProcFI + D) * (|electingFollowers| - 1)
            + ProcLI
            + (ProcEpAck + D) * (|electingFollowers| - 1);
        lemma_Math_Inequality(|l.v.leader.globals.electingFollowers|, l.v.leader.globals.nextSerialLI);
        assert res >= 0;
        res
}



/* Every LeaderInfo packet has ts specified by FollowerInfo_Message_Formula.
* This upper bound comes from the worst case execution where the leader first receives
* all FI messages. Then it sends one LI and receives one EpAck. Then it sends the next 
* LI and receives one more EpAck */
predicate ProcessEpAck_PreQuorum_Invariant(tls:TLS_State) 
    requires DS_Config_Invariant(tls.config, tls)
    requires ZK_Config_Invariant(tls.config, tls)
    requires Leader_NextSerialLI_Invariant(tls)
{
    var n := |tls.config|;
    var leaderTQP := tls.t_servers[tls.config[0]];
    !IsQuorum(|tls.config|, leaderTQP.v.leader.globals.electingFollowers)
    ==> 
    && (forall pkt | pkt in tls.t_environment.sentPackets && pkt.msg.v.LeaderInfo? && pkt.msg.v.serial < tls.f
    :: pkt.msg.ts <= LeaderInfo_Message_PreQuorum_ts_Formula(tls, pkt))
    && leaderTQP.ts <= ProcessEpAck_PreQuorum_ts_Formula(tls, leaderTQP)
    && leaderTQP.dts <= ProcessEpAck_PreQuorum_dts_Formula(tls, leaderTQP)
}

}
