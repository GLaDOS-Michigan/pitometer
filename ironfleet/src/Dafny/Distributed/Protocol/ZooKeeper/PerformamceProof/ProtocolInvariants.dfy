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
include "Commons.dfy"


/* This module contains invariants that have nothing to do with performance */
module Zookeeper_ProtocoIInvariants {
import opened Zookeeper_Commons
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
    && Leader_Sends_All_LI_Before_Receiving_EpochAck_Invariant(tls)
    && ProcessFI_PreQuorum_Implies_No_Future_Quorum_Invariant(config, tls)
    && ProcessEA_PreQuorum_Implies_No_Future_Quorum_Invariant(config, tls)
    && ProcessFI_PreQuorum_Implies_Only_FI_Messages_Invariant(config, tls)
    && Handshake_Serial_Invariant(tls)
    && Sync_Serial_Invariant(tls)
    && Follower_Serials_In_PreSync_Invariant(tls)
    && Leader_Cannot_Receive_Ack_Before_Sending_All_Syncs_Invariant(tls)
    && Follower_Cannot_Receive_NewLeader_Before_Sync(tls)
}


lemma lemma_Basic_Invariants(config:Config, tlb:seq<TLS_State>, f:int) 
    requires SeqIsUnique(config);
    requires ValidTLSBehavior(config, tlb, f)
    requires |tlb| > 1
    ensures forall i | 0 <= i < |tlb| :: Basic_Invariants(config, tlb[i]) 
{
    lemma_DS_Config_Invariant_Proof(config, tlb, f);
    lemma_ZK_Config_Invariant_Proof(config, tlb, f);
    lemma_Leader_QueuedPackets_Invariant_Proof(config, tlb, f);
    lemma_Quorums_Size_Invariant_Proof(config, tlb, f);
    lemma_ProcessFI_PreQuorum_Implies_Only_FI_Messages_Invariant_Proof(config, tlb, f);
    lemma_ProcessFI_PreQuorum_Implies_No_Future_Quorum_Invariant_Proof(config, tlb, f);
    lemma_ProcessEA_PreQuorum_Implies_No_Future_Quorum_Invariant_Proof(config, tlb, f);
    // TODO
    assume false;
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
    && ZKDB_Always_Good_Invariant(tls)
}


lemma lemma_ZK_Config_Invariant_Proof(config:Config, tlb:seq<TLS_State>, f:int)
    requires SeqIsUnique(config);
    requires ValidTLSBehavior(config, tlb, f)
    requires forall i | 0 <= i < |tlb| :: DS_Config_Invariant(config, tlb[i])
    ensures forall i | 0 <= i < |tlb| :: ZK_Config_Invariant(config, tlb[i])
    ensures forall i | 0 <= i < |tlb| :: ZKDB_Always_Good_Invariant(tlb[i])
{
    assert ZK_Config_Invariant(config, tlb[0]);
    assert ZKDB_Always_Good_Invariant(tlb[0]);
    
    var i := 0;
    while i < |tlb| - 1 
        decreases |tlb| - i
        invariant 0 <= i < |tlb|
        invariant forall k | 0 <= k <= i :: ZK_Config_Invariant(config, tlb[k])
        invariant forall k | 0 <= k <= i :: ZKDB_Always_Good_Invariant(tlb[k])
    {
        var tls, tls' := tlb[i], tlb[i+1];
        assert ZK_Config_Invariant(config, tls');
        var actor, tios:seq<TZKIo> :| actor in tls.t_servers && TLS_NextOneServer(tls, tls', actor, tios);
        if tls.t_servers[actor].v.LeaderPeer? {
            var s, s', ios := tls.t_servers[actor].v.leader, tls'.t_servers[actor].v.leader, UntagLIoOpSeq(tios);
            assert s'.globals.zkdb.initialized;
            assert isValidZKDatabase(s'.globals.zkdb);
        } else {
            var s, s', ios := tls.t_servers[actor].v.follower, tls'.t_servers[actor].v.follower, UntagLIoOpSeq(tios);
            match s.state 
            case F_HANDSHAKE_A => assert s'.zkdb == s.zkdb;
            case F_HANDSHAKE_B => assert s'.zkdb == s.zkdb;
            case F_PRESYNC => assert isValidZKDatabase(s'.zkdb);
            case F_SYNC => assert isValidZKDatabase(s'.zkdb);
            case F_RUNNING => assert s'.zkdb == s.zkdb;
            case F_ERROR => assert s'.zkdb == s.zkdb;
            assert s'.zkdb.initialized;
            assert isValidZKDatabase(s'.zkdb);
        }
        assert ZKDB_Always_Good_Invariant(tls');
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

// Needs no proof. Dafny figured this one out automatically when lemmas requiring it are called. Good boi dafny!
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
            && IsQuorum(|config|, leader.globals.connectingFollowers)
            && !IsQuorum(|config|, leader.globals.electingFollowers)
            ==> 
            && leader.globals.ackSet == {0}
            && (forall id | id in leader.handlers :: 
                if id in leader.globals.connectingFollowers
                then leader.handlers[id].state == LH_HANDSHAKE_A || leader.handlers[id].state == LH_HANDSHAKE_B
                else leader.handlers[id].state == LH_HANDSHAKE_A
            )
        )
    )
}

lemma lemma_ProcessEA_PreQuorum_Implies_No_Future_Quorum_Invariant_Proof(config:Config, tlb:seq<TLS_State>, f:int)
    requires SeqIsUnique(config);
    requires ValidTLSBehavior(config, tlb, f)
    requires forall i | 0 <= i < |tlb| :: DS_Config_Invariant(config, tlb[i])
    requires forall i | 0 <= i < |tlb| :: ZK_Config_Invariant(config, tlb[i])
    requires forall i | 0 <= i < |tlb| :: Quorums_Size_Invariant(tlb[i])
    ensures forall i | 0 <= i < |tlb| :: ProcessEA_PreQuorum_Implies_No_Future_Quorum_Invariant(config, tlb[i]) 
{
    assert ProcessEA_PreQuorum_Implies_No_Future_Quorum_Invariant(config, tlb[0]);
    lemma_ProcessFI_PreQuorum_Implies_No_Future_Quorum_Invariant_Proof(config, tlb, f);
    var i := 0;
    while i < |tlb| - 1 
        decreases |tlb| - i
        invariant 0 <= i < |tlb|
        invariant forall k | 0 <= k <= i :: ProcessEA_PreQuorum_Implies_No_Future_Quorum_Invariant(config, tlb[k])
    {
        var tls, tls' := tlb[i], tlb[i+1];
        var actor, tios:seq<TZKIo> :| actor in tls.t_servers && TLS_NextOneServer(tls, tls', actor, tios);
        if actor == config[0] {
            var s, s', ios := tls.t_servers[actor].v.leader, tls'.t_servers[actor].v.leader, UntagLIoOpSeq(tios);
            if s.state == L_STARTING {
                if  && IsQuorum(|config|, s.globals.connectingFollowers)
                    && !IsQuorum(|config|, s.globals.electingFollowers) {
                    assert !IsVerifiedQuorum(s.my_id, |s.globals.config|, s.globals.ackSet);
                    assert StepSingleHandler(s, s', ios);
                    if StepSingleHandler_NoRcv(s, s', ios) {
                        assert ProcessEA_PreQuorum_Implies_No_Future_Quorum_Invariant(config, tls');
                    } else {
                        assert ProcessEA_PreQuorum_Implies_No_Future_Quorum_Invariant(config, tls');
                    }
                } else {
                    if !IsQuorum(|config|, s.globals.connectingFollowers) {
                        assert !IsQuorum(|config|, s'.globals.electingFollowers);
                        if IsQuorum(|config|, s'.globals.connectingFollowers) {
                            // ConnectingFollowers manage to form a quorum at this step
                            assert s'.globals.electingFollowers == s.globals.electingFollowers;
                            forall id | id in s'.handlers 
                            ensures if id in s'.globals.connectingFollowers
                                    then s'.handlers[id].state == LH_HANDSHAKE_A || s'.handlers[id].state == LH_HANDSHAKE_B
                                    else s'.handlers[id].state == LH_HANDSHAKE_A
                            {
                                var h, h' := s.handlers[id], s'.handlers[id];
                                if id in s.globals.connectingFollowers {
                                    assert h.state == LH_HANDSHAKE_A;
                                    assert h'.state == LH_HANDSHAKE_A;
                                } else {
                                    assert h'.state == LH_HANDSHAKE_A;
                                }
                            }
                            assert s'.globals.ackSet == s.globals.ackSet;
                            assert ProcessEA_PreQuorum_Implies_No_Future_Quorum_Invariant(config, tls');
                        }
                    }
                }
            }
        }
        assert ProcessEA_PreQuorum_Implies_No_Future_Quorum_Invariant(config, tls');
        i := i + 1;
    }
}


lemma lemma_ProcessFI_PreQuorum_Implies_No_Future_Quorum_Invariant_Proof(config:Config, tlb:seq<TLS_State>, f:int)
    requires SeqIsUnique(config);
    requires ValidTLSBehavior(config, tlb, f)
    requires forall i | 0 <= i < |tlb| :: DS_Config_Invariant(config, tlb[i])
    requires forall i | 0 <= i < |tlb| :: ZK_Config_Invariant(config, tlb[i])
    requires forall i | 0 <= i < |tlb| :: Quorums_Size_Invariant(tlb[i])
    ensures forall i | 0 <= i < |tlb| :: ProcessFI_PreQuorum_Implies_No_Future_Quorum_Invariant(config, tlb[i]) 
{
    assert ProcessFI_PreQuorum_Implies_No_Future_Quorum_Invariant(config, tlb[0]);
    var i := 0;
    while i < |tlb| - 1 
        decreases |tlb| - i
        invariant 0 <= i < |tlb|
        invariant forall k | 0 <= k <= i :: ProcessFI_PreQuorum_Implies_No_Future_Quorum_Invariant(config, tlb[k])
    {
        var tls, tls' := tlb[i], tlb[i+1];
        assert ProcessFI_PreQuorum_Implies_No_Future_Quorum_Invariant(config, tls');
        i := i + 1;
    }
}


// TODO: Needs Proof
predicate Leader_Sends_All_LI_Before_Receiving_EpochAck_Invariant(tls:TLS_State) {
    forall ep | ep in tls.t_servers && tls.t_servers[ep].v.LeaderPeer? :: (
        && var leader := tls.t_servers[ep].v.leader;
        && (
            && |leader.globals.electingFollowers| > 1
            ==> 
            leader.globals.nextSerialLI == tls.f
        )
    )
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
    requires forall i | 0 <= i < |tlb| :: DS_Config_Invariant(config, tlb[i])
    requires forall i | 0 <= i < |tlb| :: ZK_Config_Invariant(config, tlb[i])
    requires forall i | 0 <= i < |tlb| :: Quorums_Size_Invariant(tlb[i])
    ensures forall i | 0 <= i < |tlb| :: ProcessFI_PreQuorum_Implies_Only_FI_Messages_Invariant(config, tlb[i]) 
{
    assert ProcessFI_PreQuorum_Implies_Only_FI_Messages_Invariant(config, tlb[0]);
    lemma_ProcessFI_PreQuorum_Implies_No_Future_Quorum_Invariant_Proof(config, tlb, f);
    var i := 0;
    while i < |tlb| - 1 
        decreases |tlb| - i
        invariant 0 <= i < |tlb|
        invariant forall k | 0 <= k <= i :: ProcessFI_PreQuorum_Implies_Only_FI_Messages_Invariant(config, tlb[k])
    {
        var tls, tls' := tlb[i], tlb[i+1];
        assert ProcessFI_PreQuorum_Implies_Only_FI_Messages_Invariant(config, tls');
        i := i + 1;
    }
}


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


lemma lemma_Handshake_Serial_Invariant_Proof(config:Config, tlb:seq<TLS_State>, t:int)
    requires SeqIsUnique(config);
    requires ValidTLSBehavior(config, tlb, t)
    requires forall i | 0 <= i < |tlb| :: DS_Config_Invariant(config, tlb[i])
    requires forall i | 0 <= i < |tlb| :: ZK_Config_Invariant(config, tlb[i])
    requires forall i | 0 <= i < |tlb| :: Quorums_Size_Invariant(tlb[i])
    requires forall i | 0 <= i < |tlb| :: Leader_Sends_All_LI_Before_Receiving_EpochAck_Invariant(tlb[i])
    ensures forall i | 0 <= i < |tlb| :: Handshake_Serial_Invariant(tlb[i])
{
    lemma_nextSerialLI_Equals_NumHandlers_Past_LH_HANDSHAKE_B(config, tlb, t);
    lemma_Hander_Past_HANDSHAKE_B_Implies_In_ConnectingFollowers(config, tlb, t);
    assert Handshake_Serial_Invariant(tlb[0]);
    var i := 0;
    while i < |tlb| - 1 
        decreases |tlb| - i
        invariant 0 <= i < |tlb|
        invariant forall k | 0 <= k <= i :: Handshake_Serial_Invariant(tlb[k])
    {
        var tls, tls' := tlb[i], tlb[i+1];

        forall ep | ep in tls'.t_servers && tls'.t_servers[ep].v.FollowerPeer?
        ensures tls'.t_servers[ep].v.follower.serialLI < t;
        {}
        
        forall ep | ep in tls.t_servers && tls.t_servers[ep].v.LeaderPeer? 
        ensures && tls'.t_servers[ep].v.leader.globals.nextSerialLI <= t
        {
            var actor, tios:seq<TZKIo> :| actor in tls.t_servers && TLS_NextOneServer(tls, tls', actor, tios);
            if actor == config[0] {
                var s, s', ios := tls.t_servers[actor].v.leader, tls'.t_servers[actor].v.leader, UntagLIoOpSeq(tios);
                assert s.globals.nextSerialLI >= |Handlers_Past_HandshakeB(s)|;
                if s'.globals.nextSerialLI > t {
                    /* Proof by contradiction 
                    * nextSerialLI => num handlers in state LH_HANDSHAKE_B
                    * handler in state LH_HANDSHAKE_B ==> handler in connectingFollowers - {0}
                    * but |connectingFollowers - {0}| <= f */
                    var hsb_handlers := Handlers_Past_HandshakeB(s');
                    lemma_Size_of_Supeset_2(hsb_handlers, s'.globals.connectingFollowers - {0});
                    assert |s'.globals.connectingFollowers| > t;
                    assert false;
                }
            }
        }

        forall pkt | pkt in tls'.t_environment.sentPackets 
        ensures && (pkt.msg.v.LeaderInfo? ==> pkt.msg.v.serial < t)
                && (pkt.msg.v.AckEpoch? ==> pkt.msg.v.serial < t)
        {}

        forall ep | ep in tls.t_servers && tls.t_servers[ep].v.LeaderPeer? 
        ensures 1 <= |tls'.t_servers[ep].v.leader.globals.electingFollowers| <= tls'.t_servers[ep].v.leader.globals.nextSerialLI + 1
        {
            // True due to Leader_Sends_All_LI_Before_Receiving_EpochAck_Invariant
            // if |electingFollowers| >1, then nextSerialLI = f.
        }
        i := i + 1;
    }
}


lemma lemma_Hander_Past_HANDSHAKE_B_Implies_In_ConnectingFollowers(config:Config, tlb:seq<TLS_State>, t:int)
    requires SeqIsUnique(config);
    requires ValidTLSBehavior(config, tlb, t)
    requires forall i | 0 <= i < |tlb| :: DS_Config_Invariant(config, tlb[i])
    requires forall i | 0 <= i < |tlb| :: ZK_Config_Invariant(config, tlb[i])
    requires forall i | 0 <= i < |tlb| :: Quorums_Size_Invariant(tlb[i])
    ensures forall i | 0 <= i < |tlb| :: (
        forall h | 
            && h in tlb[i].t_servers[config[0]].v.leader.handlers 
            && Past_Handshake_B(tlb[i].t_servers[config[0]].v.leader.handlers[h])
        ::
            h in tlb[i].t_servers[config[0]].v.leader.globals.connectingFollowers - {0}
    )
{
    forall h | 
            && h in tlb[0].t_servers[config[0]].v.leader.handlers 
            && Past_Handshake_B(tlb[0].t_servers[config[0]].v.leader.handlers[h])
    ensures h in tlb[0].t_servers[config[0]].v.leader.globals.connectingFollowers - {0}
    {}
    var i := 0;
    while i < |tlb| - 1 
        decreases |tlb| - i
        invariant 0 <= i < |tlb|
        invariant forall k | 0 <= k <= i :: 
            forall h | 
            && h in tlb[k].t_servers[config[0]].v.leader.handlers 
            && Past_Handshake_B(tlb[k].t_servers[config[0]].v.leader.handlers[h])
            ::
            h in tlb[k].t_servers[config[0]].v.leader.globals.connectingFollowers - {0}
    {   
        var tls, tls' := tlb[i], tlb[i+1];
        i := i + 1;
    }
}

predicate Past_Handshake_B(h : LearnerHandler) {
    || h.state == LH_HANDSHAKE_B
    || h.state == LH_PREP_SYNC
    || h.state == LH_SYNC
    || h.state == LH_PROCESS_ACK
}


lemma lemma_nextSerialLI_Equals_NumHandlers_Past_LH_HANDSHAKE_B(config:Config, tlb:seq<TLS_State>, t:int)
    requires SeqIsUnique(config);
    requires ValidTLSBehavior(config, tlb, t)
    requires forall i | 0 <= i < |tlb| :: DS_Config_Invariant(config, tlb[i])
    requires forall i | 0 <= i < |tlb| :: ZK_Config_Invariant(config, tlb[i])
    requires forall i | 0 <= i < |tlb| :: Quorums_Size_Invariant(tlb[i])
    ensures forall i | 0 <= i < |tlb| :: 
        tlb[i].t_servers[config[0]].v.leader.globals.nextSerialLI 
        == 
        |Handlers_Past_HandshakeB(tlb[i].t_servers[config[0]].v.leader)|
{
    assert tlb[0].t_servers[config[0]].v.leader.globals.nextSerialLI == |Handlers_Past_HandshakeB(tlb[0].t_servers[config[0]].v.leader)|;
    var i := 0;
    while i < |tlb| - 1 
        decreases |tlb| - i
        invariant 0 <= i < |tlb|
        invariant forall k | 0 <= k <= i :: tlb[k].t_servers[config[0]].v.leader.globals.nextSerialLI == |Handlers_Past_HandshakeB(tlb[k].t_servers[config[0]].v.leader)|
        invariant forall k | 0 <= k <= i :: 
            && tlb[k].t_servers[config[0]].v.leader.globals.zkdb.initialized == tlb[0].t_servers[config[0]].v.leader.globals.zkdb.initialized
            && tlb[k].t_servers[config[0]].v.leader.globals.zkdb.commitLog == tlb[0].t_servers[config[0]].v.leader.globals.zkdb.commitLog
            && tlb[k].t_servers[config[0]].v.leader.globals.zkdb.minCommittedLog == tlb[0].t_servers[config[0]].v.leader.globals.zkdb.minCommittedLog
            && tlb[k].t_servers[config[0]].v.leader.globals.zkdb.maxCommittedLog == tlb[0].t_servers[config[0]].v.leader.globals.zkdb.maxCommittedLog
    {   
        var tls, tls' := tlb[i], tlb[i+1];
        assert |getInMemorySuffix(tls.t_servers[config[0]].v.leader.globals.zkdb)| == 0;
        lemma_nextSerialLI_Equals_NumHandlers_Past_LH_HANDSHAKE_B_Helper(config, tls, tls', t);
        i := i + 1;
    }
}


lemma lemma_nextSerialLI_Equals_NumHandlers_Past_LH_HANDSHAKE_B_Helper(config:Config, tls:TLS_State, tls':TLS_State, t:int)
    requires SeqIsUnique(config);
    requires TLS_Next(tls, tls')
    requires DS_Config_Invariant(config, tls) && DS_Config_Invariant(config, tls')
    requires ZK_Config_Invariant(config, tls) && ZK_Config_Invariant(config, tls')
    requires Quorums_Size_Invariant(tls) && Quorums_Size_Invariant(tls')
    requires |getInMemorySuffix(tls.t_servers[config[0]].v.leader.globals.zkdb)| == 0;
    requires tls.t_servers[config[0]].v.leader.globals.nextSerialLI 
        == 
        |Handlers_Past_HandshakeB(tls.t_servers[config[0]].v.leader)|
    ensures tls'.t_servers[config[0]].v.leader.globals.nextSerialLI 
        == 
        |Handlers_Past_HandshakeB(tls'.t_servers[config[0]].v.leader)|   
{
    var actor, tios:seq<TZKIo> :| actor in tls.t_servers && TLS_NextOneServer(tls, tls', actor, tios);
    if actor == config[0] {
        var s, s', ios := tls.t_servers[actor].v.leader, tls'.t_servers[actor].v.leader, UntagLIoOpSeq(tios);
        var sh, sh' := Handlers_Past_HandshakeB(s), Handlers_Past_HandshakeB(s');
        if s.state == L_STARTING {
            if !IsVerifiedQuorum(s.my_id, |s.globals.config|, s.globals.ackSet) {
                assert StepSingleHandler(s, s', ios);
                if StepSingleHandler_NoRcv(s, s', ios) {
                    var fid :| LHNext(s, s', fid, ios);
                    var h, h', g, g' := s.handlers[fid], s'.handlers[fid], s.globals, s'.globals;
                    if h.state == LH_HANDSHAKE_A {
                        if IsVerifiedQuorum(h.follower_id, |g.config|, g.connectingFollowers) {
                            assert s'.globals.nextSerialLI == s.globals.nextSerialLI + 1;
                            assert sh' == sh + {h.follower_id};
                        }
                    } else if h.state == LH_PREP_SYNC {
                        assert g.zkdb.initialized && isValidZKDatabase(g.zkdb);
                        assert sh' == sh;
                    } else {
                        assert sh' == sh;
                    }
                } else {
                    assert StepSingleHandler_Rcv(s, s', ios);
                    assert sh' == sh;
                }
                assert s'.globals.nextSerialLI == |sh'|; 
            } else {
                assert sh' == sh;
            }
        }
    } 
}


function Handlers_Past_HandshakeB(leader:Leader) : set<nat> {
    set h | h in leader.handlers && Past_Handshake_B(leader.handlers[h]) :: h
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


predicate Quorums_Size_Invariant(tls:TLS_State) {
    var n := |tls.config|;
    forall ep | ep in tls.t_servers && tls.t_servers[ep].v.LeaderPeer? :: (
        // && var ld := tls.t_servers[ep].v.leader;
        && {tls.t_servers[ep].v.leader.my_id} <= tls.t_servers[ep].v.leader.globals.connectingFollowers
        && {tls.t_servers[ep].v.leader.my_id} <= tls.t_servers[ep].v.leader.globals.electingFollowers
        && {tls.t_servers[ep].v.leader.my_id} <= tls.t_servers[ep].v.leader.globals.ackSet
        && 1 <= |tls.t_servers[ep].v.leader.globals.connectingFollowers| <= (n/2) + 1
        && 1 <= |tls.t_servers[ep].v.leader.globals.electingFollowers| <= (n/2) + 1
        && 1 <=|tls.t_servers[ep].v.leader.globals.ackSet| <= (n/2) + 1
    )
}


lemma lemma_Quorums_Size_Invariant_Proof(config:Config, tlb:seq<TLS_State>, t:int)
    requires SeqIsUnique(config);
    requires ValidTLSBehavior(config, tlb, t)
    requires forall i | 0 <= i < |tlb| :: DS_Config_Invariant(config, tlb[i])
    requires forall i | 0 <= i < |tlb| :: ZK_Config_Invariant(config, tlb[i])
    ensures forall i | 0 <= i < |tlb| :: Quorums_Size_Invariant(tlb[i])
{
    assert Quorums_Size_Invariant(tlb[0]);
    var i := 0;
    while i < |tlb| - 1 
        decreases |tlb| - i
        invariant 0 <= i < |tlb|
        invariant forall k | 0 <= k <= i :: Quorums_Size_Invariant(tlb[k])
    {
        var tls, tls' := tlb[i], tlb[i+1];
        var actor, tios:seq<TZKIo> :| actor in tls.t_servers && TLS_NextOneServer(tls, tls', actor, tios);
        if tls.t_servers[actor].v.LeaderPeer? {
            var s, s', ios := tls.t_servers[actor].v.leader, tls'.t_servers[actor].v.leader, UntagLIoOpSeq(tios);
            assert Quorums_Size_Invariant(tls');
        }
        i := i + 1;
    }
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
