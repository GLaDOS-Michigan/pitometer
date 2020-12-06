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
module Zookeeper_ProtocoIInvariants {
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
    && ZKDB_Always_Good_Invariant(tls)
    && Leader_QueuedPackets_Invariant(config, tls)
    && Quorums_Size_Invariant(tls)
    && ProcessFI_PreQuorum_Implies_No_Future_Quorum_Invariant(config, tls)
    && ProcessEA_PreQuorum_Implies_No_Future_Quorum_Invariant(config, tls)
    && ProcessFI_PreQuorum_Implies_Only_FI_Messages_Invariant(config, tls)
    && Handshake_Serial_Invariant(tls)
    && Sync_Serial_Invariant(tls)
    && Follower_Serials_In_PreSync_Invariant(tls)
    && Leader_Cannot_Receive_Ack_Before_Sending_All_Syncs_Invariant(tls)
    && Follower_Cannot_Receive_NewLeader_Before_Sync(tls)
    && Leader_Only_Sends_Leader_Msgs(tls)
}


lemma lemma_Basic_Invariants(config:Config, tlb:seq<TLS_State>, f:int) 
    requires SeqIsUnique(config);
    requires ValidTLSBehavior(config, tlb, f)
    ensures forall i | 0 <= i < |tlb| :: Basic_Invariants(config, tlb[i]) 
{
    // TODO
    assume false;
    lemma_DS_Config_Invariant_Proof(config, tlb, f);
    lemma_ZK_Config_Invariant_Proof(config, tlb, f);
    lemma_Leader_QueuedPackets_Invariant_Proof(config, tlb, f);
}


/*****************************************************************************************
/                                   DSConfigInvariant                                    *
*****************************************************************************************/


/* config has same endpoints as tls.servers and each node's id is its index in config */
predicate DS_Config_Invariant(config:Config, tls:TLS_State) {
    && tls.f >= 1
    && tls.config == config
    && |config| == tls.f*2+1
    && (forall ep :: ep in config <==> ep in tls.t_servers)
    && (forall ep | ep in tls.t_servers :: config == if tls.t_servers[ep].v.FollowerPeer? then tls.t_servers[ep].v.follower.config else tls.t_servers[ep].v.leader.globals.config)
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
    assert tlb[0].f >= 1;
    assert |config| == tlb[0].f*2+1;
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
*                                   ZKConfigInvariant                                    *
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
    && (forall i :: 1 <= i < |config| <==> i in tls.t_servers[config[0]].v.leader.handlers)
    && (forall follower_id | 1 <= follower_id < |config| :: tls.t_servers[config[0]].v.leader.handlers[follower_id].follower_id == follower_id)
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


predicate ZKDB_Always_Good_Invariant(tls:TLS_State) {
    forall ep | ep in tls.t_servers:: 
    && (tls.t_servers[ep].v.LeaderPeer? ==> tls.t_servers[ep].v.leader.globals.zkdb.initialized && isValidZKDatabase(tls.t_servers[ep].v.leader.globals.zkdb))
    && (tls.t_servers[ep].v.FollowerPeer? ==> tls.t_servers[ep].v.follower.zkdb.initialized && isValidZKDatabase(tls.t_servers[ep].v.follower.zkdb))
}


/*****************************************************************************************
*                                   Transition Invariants                                *
*****************************************************************************************/

predicate SentPacketsSet_Property(tls:TLS_State, tls':TLS_State, id:EndPoint, tios:seq<TZKIo>)
    requires id in tls.t_servers
    requires TLS_Next(tls, tls')
    requires TLS_NextOneServer(tls, tls', id, tios)
{
    tls'.t_environment.sentPackets == 
    tls.t_environment.sentPackets + (set tio : TimestampedLIoOp | tio in tios && tio.LIoOpSend? :: tio.s)
}


lemma lemma_SentPacketsSet_Property(tls:TLS_State, tls':TLS_State, id:EndPoint, tios:seq<TZKIo>)
    requires id in tls.t_servers
    requires TLS_Next(tls, tls')
    requires TLS_Next(tls, tls')
    requires TLS_NextOneServer(tls, tls', id, tios)
    ensures SentPacketsSet_Property(tls, tls', id, tios)
{}

// TODO: Needs Proof
predicate QuorumsMonotoneIncreasing_Property(tls:TLS_State, tls':TLS_State)
    requires TLS_Next(tls, tls')
    requires Basic_Invariants(tls.config, tls) && Basic_Invariants(tls'.config, tls')
{
    && |tls.t_servers[tls.config[0]].v.leader.globals.connectingFollowers| <= |tls'.t_servers[tls'.config[0]].v.leader.globals.connectingFollowers|
    && |tls.t_servers[tls.config[0]].v.leader.globals.electingFollowers| <= |tls'.t_servers[tls'.config[0]].v.leader.globals.electingFollowers|
}


/*****************************************************************************************
*                                 Facts about the leader                                 *
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

/* QueuedPackets for all LearnerHandlers in all LeaderPeers only contains the approriate messages */
predicate Leader_QueuedPackets_Invariant(config:Config, tls:TLS_State) {
    forall ep | ep in tls.t_servers ::
        tls.t_servers[ep].v.LeaderPeer? ==> (
            forall lh | lh in tls.t_servers[ep].v.leader.handlers.Values
            :: && QueuedPackets_Only_Contains_LeaderMessages(lh.queuedPackets)
               && |lh.queuedPackets| <= 1
    )
}

lemma lemma_Leader_QueuedPackets_Invariant_Proof(config:Config, tlb:seq<TLS_State>, f:int)
    requires SeqIsUnique(config);
    requires ValidTLSBehavior(config, tlb, f)
    ensures forall i | 0 <= i < |tlb| :: Leader_QueuedPackets_Invariant(config, tlb[i]) 
{
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


/* If connectingFollowers is not a full quorum, then all future quorums are empty */
// TODO: Needs Proof
predicate ProcessFI_PreQuorum_Implies_No_Future_Quorum_Invariant(config:Config, tls:TLS_State){
    forall ep | ep in tls.t_servers && tls.t_servers[ep].v.LeaderPeer? :: (
        && var leader := tls.t_servers[ep].v.leader;
        && (
            && !IsQuorum(|config|, leader.globals.connectingFollowers)
            ==> 
            && leader.globals.electingFollowers == {0}
            && leader.globals.ackSet == {0}
            && leader.globals.nextSerialLI == 0
            && (forall id | id in leader.handlers :: leader.handlers[id].state == LH_HANDSHAKE_A)
        )
    )
}

/* If electingFollowers is not a full quorum, then all future quorums are empty */
predicate ProcessEA_PreQuorum_Implies_No_Future_Quorum_Invariant(config:Config, tls:TLS_State){
    forall ep | ep in tls.t_servers && tls.t_servers[ep].v.LeaderPeer? :: (
        && var leader := tls.t_servers[ep].v.leader;
        && (
            && !IsQuorum(|config|, leader.globals.electingFollowers)
            ==> 
            && leader.globals.ackSet == {0}
            && (forall id | id in leader.handlers :: 
                if id in leader.globals.connectingFollowers
                then leader.handlers[id].state == LH_HANDSHAKE_B
                else leader.handlers[id].state == LH_HANDSHAKE_A
            )
        )
    )
}


lemma lemma_ProcessFI_PreQuorum_Implies_No_Future_Quorum_Invariant_Proof(config:Config, tlb:seq<TLS_State>, f:int)
    requires SeqIsUnique(config);
    requires ValidTLSBehavior(config, tlb, f)
    ensures forall i | 0 <= i < |tlb| :: ProcessFI_PreQuorum_Implies_No_Future_Quorum_Invariant(config, tlb[i]) 
{
    // TODO
    assume false;
}


/* If connectingFollowers is not a full quorum, the only messages in the network are FollowerInfo messages */
predicate ProcessFI_PreQuorum_Implies_Only_FI_Messages_Invariant(config:Config, tls:TLS_State){
    forall ep | ep in tls.t_servers && tls.t_servers[ep].v.LeaderPeer? :: (
        && var leader := tls.t_servers[ep].v.leader;
        && (
            && !IsQuorum(|config|, leader.globals.connectingFollowers)
            ==> 
            forall pkt | pkt in tls.t_environment.sentPackets ::
                pkt.msg.v.FollowerInfo?
        )
    )
}

lemma lemma_ProcessFI_PreQuorum_Implies_Only_FI_Messages_Invariant_Proof(config:Config, tlb:seq<TLS_State>, f:int)
    requires SeqIsUnique(config);
    requires ValidTLSBehavior(config, tlb, f)
    ensures forall i | 0 <= i < |tlb| :: ProcessFI_PreQuorum_Implies_Only_FI_Messages_Invariant(config, tlb[i]) 
{
    // TODO
    assume false;
}

// TODO: Needs Proof
predicate Handshake_Serial_Invariant(tls:TLS_State) {
    && (forall ep | ep in tls.t_servers && tls.t_servers[ep].v.LeaderPeer? :: (
        && var l := tls.t_servers[ep].v.leader;
        // Size of electing followers is bound by the # LeaderInfo's sent
        && 1 <= |l.globals.electingFollowers| <= l.globals.nextSerialLI + 1
        // Don't send more than f LeaderInfo's
        && l.globals.nextSerialLI <= tls.f
    ))
    && (forall pkt | pkt in tls.t_environment.sentPackets :: (
        // LeaderInfo and AckEpoch serial bound by f
        && (pkt.msg.v.LeaderInfo? ==> pkt.msg.v.serial < tls.f)  // strictly <f bc I send out at most f messages. 1st msg is #0. f-th msg is #(f-1)
        && (pkt.msg.v.AckEpoch? ==> pkt.msg.v.serial < tls.f)
    )) // Follower serial bound by f
    && (forall ep | ep in tls.t_servers && tls.t_servers[ep].v.FollowerPeer? :: (
        && var f := tls.t_servers[ep].v.follower;
        f.serialLI < tls.f
    ))
}

// TODO: Needs Proof
predicate Sync_Serial_Invariant(tls:TLS_State) {
    && (forall ep | ep in tls.t_servers && tls.t_servers[ep].v.LeaderPeer? :: (
        && var l := tls.t_servers[ep].v.leader;
        // Size of ackSet is bound by the # NL's sent
        && 1 <= |l.globals.ackSet| <= l.globals.nextSerialNL + 1
        // Don't send more than f Sync's and NL's
        && l.globals.nextSerialSync <= tls.f
        && l.globals.nextSerialNL <= tls.f
        // Can only send NL after sending a Sync to someone
        && l.globals.nextSerialNL <= l.globals.nextSerialSync 
        // At most f of such state transitions
        && l.globals.procEpCount <= tls.f
        && l.globals.prepCount <= tls.f
    ))
    && (forall pkt | pkt in tls.t_environment.sentPackets :: (
        // LeaderInfo and AckEpoch serial bound by f
        && (pkt.msg.v.SyncDIFF? ==> pkt.msg.v.serial < tls.f)  // strictly <f bc I send out at most f messages. 1st msg is #0. f-th msg is #(f-1)
        && (pkt.msg.v.SyncSNAP? ==> pkt.msg.v.serial < tls.f)
        && (pkt.msg.v.SyncTRUNC? ==> pkt.msg.v.serial < tls.f)
        && (pkt.msg.v.NewLeader? ==> pkt.msg.v.serial < tls.f)
        && (pkt.msg.v.Ack? ==> pkt.msg.v.serial < tls.f)
        && (pkt.msg.v.Commit? ==> pkt.msg.v.serial < tls.f)
    )) // Follower serial bound by f
    && (forall ep | ep in tls.t_servers && tls.t_servers[ep].v.FollowerPeer? :: (
        && var f := tls.t_servers[ep].v.follower;
        && f.serialSync < tls.f
        && f.serialNL < tls.f
    ))
}

// TODO: Needs Proof
predicate Quorums_Size_Invariant(tls:TLS_State) {
    var n := |tls.config|;
    forall ep | ep in tls.t_servers && tls.t_servers[ep].v.LeaderPeer? :: (
        && var l := tls.t_servers[ep].v.leader;
        && {l.my_id} <= l.globals.connectingFollowers
        && {l.my_id} <= l.globals.electingFollowers
        && {l.my_id} <= l.globals.ackSet
        && 1 <= |l.globals.connectingFollowers| <= (n/2) + 1
        && 1 <= |l.globals.electingFollowers| <= (n/2) + 1
        && 1 <=|l.globals.ackSet| <= (n/2) + 1
    )
}


predicate Leader_Only_Sends_Leader_Msgs(tls:TLS_State) {
    && tls.t_environment.nextStep.LEnvStepHostIos?
    && tls.t_environment.nextStep.actor in tls.t_servers
    && tls.t_servers[tls.t_environment.nextStep.actor].v.LeaderPeer?
    ==> 
    forall tio | tio in tls.t_environment.nextStep.ios && tio.LIoOpSend? :: 
        && !tio.s.msg.v.FollowerInfo? 
        && !tio.s.msg.v.AckEpoch? 
        && !tio.s.msg.v.Ack? 
}



// TODO: Needs Proof
predicate Follower_Serials_In_PreSync_Invariant(tls:TLS_State) 
    requires |tls.config| > 0
    requires tls.config[0] in tls.t_servers
{
    forall ep | 
    && ep in tls.t_servers 
    && tls.t_servers[ep].v.FollowerPeer? 
    && tls.t_servers[ep].v.follower.state == F_PRESYNC
    :: 
    && var f := tls.t_servers[ep].v.follower;
    && f.serialLI >= 0 
    && f.serialSync < 0 
    && f.serialNL < 0
}


// TODO: Needs Proof
predicate Leader_Cannot_Receive_Ack_Before_Sending_All_Syncs_Invariant(tls:TLS_State) 
    requires |tls.config| > 0
    requires tls.config[0] in tls.t_servers
    requires tls.t_servers[tls.config[0]].v.LeaderPeer?
{
    var l := tls.t_servers[tls.config[0]].v.leader;
    (exists fid :: 
        && fid in l.handlers 
        && (l.handlers[fid].state == LH_PREP_SYNC || l.handlers[fid].state == LH_SYNC)
    )
    ==> 
    |l.globals.ackSet| == 1
}

// TODO: Needs Proof
predicate Follower_Cannot_Receive_NewLeader_Before_Sync(tls:TLS_State) {
    forall ep | ep in tls.t_servers && tls.t_servers[ep].v.FollowerPeer? 
    :: 
    && var f := tls.t_servers[ep].v.follower;
    && f.serialNL >= 0 ==> f.serialSync >= 0 && f.serialLI >= 0
}
}
