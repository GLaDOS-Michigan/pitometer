include "../../../Impl/Common/SeqIsUniqueDef.i.dfy"
include "../../../Common/Framework/EnvironmentTCP.s.dfy"
include "../Timestamps/TimestampedType.dfy"
include "TimestampedLS.dfy"
include "ProtocolInvariants.dfy"

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
import opened Zookeeper_ProtocoIInvariants
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
    Sync +
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
    requires |l.v.leader.globals.connectingFollowers| >= 1
{
    var connectingFollowers := l.v.leader.globals.connectingFollowers;
    l.dts + ProcFI * (|connectingFollowers|-1)
}

function ProcessFI_PreQuorum_dts_Formula() : Timestamp {
    FollowerInfo_Message_ts_Formula()
}


/* Before the leader has a complete connectingFollowers quorum, its ts and dts are 
* specified by the respective formulas */
predicate ProcessFI_PreQuorum_Invariant(tls:TLS_State) 
    requires DS_Config_Invariant(tls.config, tls)
    requires ZK_Config_Invariant(tls.config, tls)
    requires Quorums_Size_Invariant(tls)
{
    var n := |tls.config|;
    var leaderTQP := tls.t_servers[tls.config[0]];
    !IsQuorum(|tls.config|, leaderTQP.v.leader.globals.connectingFollowers)
    ==> 
    && leaderTQP.ts <= ProcessFI_PreQuorum_ts_Formula(leaderTQP)
    && leaderTQP.dts <= ProcessFI_PreQuorum_dts_Formula()
}


//----------------------------------------------------------------------------------------

/* For all followers after sending FollowerInfo messages */
predicate Follower_HandshakeB_Invariant(tls:TLS_State){
    forall ep | 
        && ep in tls.t_servers 
        && tls.t_servers[ep].v.FollowerPeer? 
        && tls.t_servers[ep].v.follower.state == F_HANDSHAKE_B
    ::  && tls.t_servers[ep].dts == TimeZero()
        && tls.t_servers[ep].ts == SendFI
}

/* For all LeaderInfo messages sent BEFORE electingFollowers reaches a quorum */
function LeaderInfo_Message_PreQuorum_ts_Formula(f:int, serial:nat) : Timestamp 
    requires f >= 1;
{
    SendFI + D 
    + ProcFI * f // Possibly receive at most f FI's
    + ProcFI * (serial + 1)  // I am the (serial + 1)-th LI to be sent out
    + D
}


/* Invariant for followers in F_PRESYNC mode */
predicate Follower_PreSync_Invariant(tls:TLS_State){
    forall ep | 
        && ep in tls.t_servers
        && tls.t_servers[ep].v.FollowerPeer? 
        && tls.t_servers[ep].v.follower.state == F_PRESYNC
        && 0 <= tls.t_servers[ep].v.follower.serialLI < tls.f
    ::  && tls.t_servers[ep].dts <= LeaderInfo_Message_PreQuorum_ts_Formula(tls.f, tls.t_servers[ep].v.follower.serialLI)
        && tls.t_servers[ep].ts <= LeaderInfo_Message_PreQuorum_ts_Formula(tls.f, tls.t_servers[ep].v.follower.serialLI) + ProcLI
}

/* For all AckEpoch messages sent BEFORE electingFollowers reaches a quorum */
function AckEpoch_Message_PreQuorum_ts_Formula(f:int, serial:nat) : Timestamp 
    requires f >= 1;
{
    LeaderInfo_Message_PreQuorum_ts_Formula(f, serial) + ProcLI + D
}


function ProcessEpAck_PreQuorum_ts_Formula(f:int, n:int, l:TQuorumPeer) : Timestamp 
    requires l.v.LeaderPeer?
    requires f >= 1
    requires |l.v.leader.globals.electingFollowers| >= 1;
    requires |l.v.leader.globals.connectingFollowers| >= 1;
    requires !IsQuorum(n, l.v.leader.globals.electingFollowers)
{
    var connectingFollowers, electingFollowers := l.v.leader.globals.connectingFollowers, l.v.leader.globals.electingFollowers;
    if |electingFollowers| <= 1 then 
        // Have not received any AckEpochs. Dts is from last FollowerInfo received
        // ts is from receiving |cf| FI's, and sending out #serial of them
        FollowerInfo_Message_ts_Formula()
        + ProcFI * (|connectingFollowers|-1)
        + ProcFI * l.v.leader.globals.nextSerialLI
    else 
        AckEpoch_Message_PreQuorum_ts_Formula(f, f-1)
        + ProcEpAck * (|electingFollowers|-1) // Add processing time
}


function ProcessEpAck_PreQuorum_dts_Formula(f:int, n:int, l:TQuorumPeer) : Timestamp 
    requires l.v.LeaderPeer?
    requires f >= 1
    requires !IsQuorum(n, l.v.leader.globals.electingFollowers)
{
    var electingFollowers := l.v.leader.globals.electingFollowers;
    if |electingFollowers| <= 1 then 
        // Have not received any AckEpochs. Dts is from last FollowerInfo received
        FollowerInfo_Message_ts_Formula()
    else 
        // From the last AckEpoch I received
        // The largest AckEpoch serial I can get is f
        AckEpoch_Message_PreQuorum_ts_Formula(f, f-1)
}

//----------------------------------------------------------------------------------------

/* Summary of handshake B phase messages */
predicate Handshake_Messages_Invariant(tls:TLS_State) 
    requires DS_Config_Invariant(tls.config, tls)
    requires ZK_Config_Invariant(tls.config, tls)
{   
    && (forall pkt | pkt in tls.t_environment.sentPackets && pkt.msg.v.LeaderInfo? // && pkt.msg.v.serial < tls.f
    :: pkt.msg.ts <= LeaderInfo_Message_PreQuorum_ts_Formula(tls.f, pkt.msg.v.serial))
    && (forall pkt | pkt in tls.t_environment.sentPackets && pkt.msg.v.AckEpoch?  // && pkt.msg.v.serial < tls.f
    :: pkt.msg.ts <= AckEpoch_Message_PreQuorum_ts_Formula(tls.f, pkt.msg.v.serial))
}

/* Summary of handshake B phase leader */
predicate Handshake_Leader_PreQuorum_Invariant(tls:TLS_State) 
    requires DS_Config_Invariant(tls.config, tls)
    requires ZK_Config_Invariant(tls.config, tls)
    requires Handshake_Serial_Invariant(tls)
    requires Quorums_Size_Invariant(tls)
{
    var n := |tls.config|;
    var leaderTQP := tls.t_servers[tls.config[0]];
    !IsQuorum(n, leaderTQP.v.leader.globals.electingFollowers)
    ==> 
    && leaderTQP.ts <= ProcessEpAck_PreQuorum_ts_Formula(tls.f, n, leaderTQP)
    && leaderTQP.dts <= ProcessEpAck_PreQuorum_dts_Formula(tls.f, n, leaderTQP)
}

/* Summary of handshake B phase follower */
predicate Handshake_Follower_Invariant(tls:TLS_State) 
    requires DS_Config_Invariant(tls.config, tls)
    requires ZK_Config_Invariant(tls.config, tls)
{
    && Follower_HandshakeB_Invariant(tls)
    && Follower_PreSync_Invariant(tls)
}


/*****************************************************************************************
*                                 Sync Phase Guarantees                                  *
*****************************************************************************************/


/* For all SyncDIFF | SyncSNAP messages sent */
function Sync_Message_ts_Formula(f:int, serial:nat) : Timestamp 
    requires f >= 1;
{
    AckEpoch_Message_PreQuorum_ts_Formula(f, f-1)
    + ProcEpAck * f
    + PreSync * f    // max possible PrepSyncs done before I was sent
    + Sync * serial  // max possible NewLeader sent before I was sent
    + Sync * (serial + 1)  // I am the (serial + 1)-th sync message sent
    + D
}

/* For all NewLeader messages sent */
function NewLeader_Message_ts_Formula(f:int, serial:nat) : Timestamp 
    requires f >= 1;
{
    AckEpoch_Message_PreQuorum_ts_Formula(f, f-1)
    + ProcEpAck * f
    + PreSync * f    // max possible PrepSyncs done before I was sent
    + Sync * f  // max possible Syncs sent before I was sent
    + Sync * (serial + 1)  // I am the (serial + 1)-th NL message sent
    + D
}

/* Follower in F_SYNC state. May or may not have received Newleader at this state. 
* Just received a SyncDiff | SyncSnap | NewLeader */
function Follower_F_SYNC_dts_Formula(f:int, s:TQuorumPeer) : Timestamp 
    requires s.v.FollowerPeer?    
    requires f >= 1
    requires s.v.follower.serialSync >= 0
{
    if s.v.follower.serialNL < 0 
    then // yet to receive NL. Just processed SyncDiff | SyncSnap
        Sync_Message_ts_Formula(f, s.v.follower.serialSync)
    else // just received NL. Waititng to receive final UpToDate to start running
        NewLeader_Message_ts_Formula(f, s.v.follower.serialNL)
}

/* Follower in F_SYNC state. May or may not have received Newleader at this state. 
* Just received a SyncDiff | SyncSnap | NewLeader */
function Follower_F_SYNC_ts_Formula(f:int, s:TQuorumPeer) : Timestamp 
    requires s.v.FollowerPeer?    
    requires f >= 1
    requires s.v.follower.serialSync >= 0
{
    if s.v.follower.serialNL < 0 
    then // yet to receive NL. Just processed SyncDiff | SyncSnap
        Follower_F_SYNC_dts_Formula(f, s) + ProcSyncI
    else // just received NL. Waititng to receive final UpToDate to start running
        NewLeader_Message_ts_Formula(f, f-1) + ProcSyncI + ProcSync
}


function Ack_Message_ts_Formula(f:int, serial:nat) : Timestamp
    requires f >= 1;
{
    NewLeader_Message_ts_Formula(f, f-1) + ProcSyncI + ProcSync + D
}


function ProcessAck_PreQuorum_ts_Formula(f:int, n:int, l:TQuorumPeer) : Timestamp 
    requires l.v.LeaderPeer?
    requires f >= 1
    requires !IsQuorum(n, l.v.leader.globals.ackSet)
{
    var ackSet := l.v.leader.globals.ackSet;
    if |ackSet| <= 1 then 
        AckEpoch_Message_PreQuorum_ts_Formula(f, f-1)
        + ProcEpAck * l.v.leader.globals.procEpCount 
        + PreSync * l.v.leader.globals.prepCount    // num of PrepSyncs done
        + Sync * l.v.leader.globals.nextSerialSync   // Only sync, never syncSnap
        + Sync * l.v.leader.globals.nextSerialNL   // Only sync, never syncSnap
    else 
        Ack_Message_ts_Formula(f, f-1) + ProcAck * (|ackSet|-1)
}


function ProcessAck_PreQuorum_dts_Formula(f:int, n:int, l:TQuorumPeer) : Timestamp 
    requires l.v.LeaderPeer?
    requires f >= 1
    requires !IsQuorum(n, l.v.leader.globals.ackSet)
{
    var ackSet := l.v.leader.globals.ackSet;
    if |ackSet| <= 1 then 
        // Have not received any Acks. Dts is from last AckEpoch received
        // The largest AckEpoch serial I can get is f
        AckEpoch_Message_PreQuorum_ts_Formula(f, f-1)
    else 
        // From the last Ack I received
        // The largest Ack serial I can get is f
        Ack_Message_ts_Formula(f, f-1)
}


//----------------------------------------------------------------------------------------
/* Summary of Sync phase messages */
predicate Sync_Messages_Invariant(tls:TLS_State) 
    requires DS_Config_Invariant(tls.config, tls)
    requires ZK_Config_Invariant(tls.config, tls)
{   
    && (forall pkt | pkt in tls.t_environment.sentPackets && (pkt.msg.v.SyncDIFF? || pkt.msg.v.SyncSNAP?) 
    :: pkt.msg.ts <= Sync_Message_ts_Formula(tls.f, pkt.msg.v.serial))
    && (forall pkt | pkt in tls.t_environment.sentPackets && pkt.msg.v.NewLeader? 
    :: pkt.msg.ts <= NewLeader_Message_ts_Formula(tls.f, pkt.msg.v.serial))
    && (forall pkt | pkt in tls.t_environment.sentPackets && pkt.msg.v.Ack? 
    :: pkt.msg.ts <= Ack_Message_ts_Formula(tls.f, pkt.msg.v.serial))
}


/* Summary of Sync phase followers */
predicate Sync_Follower_Invariant(tls:TLS_State){
    forall ep | 
        && ep in tls.t_servers
        && tls.t_servers[ep].v.FollowerPeer? 
        && tls.t_servers[ep].v.follower.state == F_SYNC
        && 0 <= tls.t_servers[ep].v.follower.serialSync < tls.f
    ::  && tls.t_servers[ep].dts <= Follower_F_SYNC_dts_Formula(tls.f, tls.t_servers[ep])
        && tls.t_servers[ep].ts <= Follower_F_SYNC_ts_Formula(tls.f, tls.t_servers[ep])
}


/* Summary of Sync phase leader before AckSet quorum */
predicate Sync_Leader_PreQuorum_Invariant(tls:TLS_State) 
    requires DS_Config_Invariant(tls.config, tls)
    requires ZK_Config_Invariant(tls.config, tls)
    // requires Quorums_Size_Invariant(tls)
{
    var n := |tls.config|;
    var leaderTQP := tls.t_servers[tls.config[0]];
    && !IsQuorum(n, leaderTQP.v.leader.globals.ackSet)
    ==> 
    && leaderTQP.ts <= ProcessAck_PreQuorum_ts_Formula(tls.f, n, leaderTQP)
    && leaderTQP.dts <= ProcessAck_PreQuorum_dts_Formula(tls.f, n, leaderTQP)
}

}
