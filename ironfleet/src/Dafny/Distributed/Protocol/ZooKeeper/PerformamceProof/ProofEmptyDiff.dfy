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
include "BasicInvariants.dfy"
include "Commons.dfy"


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
import opened Zookeeper_BasicInvariants
import opened Zookeeper_Commons


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



/*****************************************************************************************
*                                FollowerInit_Invariant                                  *
*****************************************************************************************/


lemma lemma_FollowerInit_Invariant_Induction(config:Config, tls:TLS_State, tls':TLS_State)
    requires TLS_Next(tls, tls')
    requires Basic_Invariants(config, tls) && Basic_Invariants(config, tls')
    requires General_LS_Performance_Assumption(tls)
    // Induction hypothesis
    requires FollowerInit_Invariant(tls)
    ensures FollowerInit_Invariant(tls')
{}



/*****************************************************************************************
*                            FollowerInfo_Message_Invariant                              *
*****************************************************************************************/


lemma lemma_FollowerInfo_Message_Invariant_Induction(config:Config, tls:TLS_State, tls':TLS_State)
    requires TLS_Next(tls, tls')
    requires General_LS_Performance_Assumption(tls)
    // Pre-established Invariants
    requires Basic_Invariants(config, tls) && Basic_Invariants(config, tls')
    requires FollowerInit_Invariant(tls);
    // Induction hypothesis
    requires FollowerInfo_Message_ts_Invariant(tls)
    ensures FollowerInfo_Message_ts_Invariant(tls')
{
    var actor, tios :| actor in tls.t_servers && TLS_NextOneServer(tls, tls', actor, tios);
    var ls, ls', ios := UntagLS_State(tls), UntagLS_State(tls'), UntagLIoOpSeq(tios);
    assert LS_NextOneServer(ls, ls', actor, ios);
    if ls.servers[actor].FollowerPeer? {
        /* Here we deal with a follower taking a step */
        var s, s' := ls.servers[actor].follower, ls'.servers[actor].follower;
        assert FollowerNext(s, s', ios);
        match s.state 
        case F_HANDSHAKE_A =>  
            assert SendMyInfo(s, s', ios);
            assert tls.t_servers[actor].ts == TimeZero();  // by FollowerInit_Invariant
            assert |ios| == |tios| == 1;
            var tio : TimestampedLIoOp := tios[0];
            assert tio.LIoOpSend?;
            var hs := ActionToHostStep(tls, tls', actor, tios);
            assert tls'.t_servers[actor].ts == TLS_NoRecvPerfUpdate(tls.t_servers[actor].ts, hs);
            assert tio.s.msg.ts == SendFI + D;
            lemma_SentPacketsSet_Property(tls, tls', actor, tios);
        case F_HANDSHAKE_B => 
            assert AcceptNewEpoch(s, s', ios);
            assert ios[0].LIoOpReceive?;
            if ios[0].r.msg.newZxid.epoch < s.accepted_epoch {
                assert |ios| == 1;
                assert forall io | io in ios :: !io.LIoOpSend?;
            } else { 
                assert |ios| == 2;
                assert ios[1].LIoOpSend?;
                assert ios[1].s.msg.AckEpoch?;
                assert !ios[1].s.msg.FollowerInfo?;
                forall i | 0 <= i < |ios| 
                ensures ios[i].LIoOpSend? ==> ios[i].s.msg.AckEpoch? {
                    if i == 0 {
                        assert ios[i].LIoOpReceive?;
                    } else {
                        assert i == 1;
                        assert ios[i].LIoOpSend? && ios[i].s.msg.AckEpoch?;
                    }
                }
            }
        case F_PRESYNC => 
            assert PreSyncWithLeader(s, s', ios);
        case F_SYNC => 
            assert SyncWithLeader(s, s', ios);
            assert forall io | io in ios && io.LIoOpSend? :: !io.s.msg.FollowerInfo?;
        case F_RUNNING => assert |ios| == 0;       
        case F_ERROR => assert |ios| == 0;
    } else {
        /* No new FollowerInfo messages get sent, so nothing to worry */
        var s, s' := ls.servers[actor].leader, ls'.servers[actor].leader;
        assert LeaderNext(s, s', ios);
        if s.state == L_STARTING {
            if |ios| > 0 {
                assert StepSingleHandler(s, s', ios);
                var handlerToStep :| LHNext(s, s', handlerToStep, ios);
                var h, h' := s.handlers[handlerToStep], s'.handlers[handlerToStep];
                var g, g' := s.globals, s'.globals;
                assert LearnerHandlerNext(h, h', g, g', ios);
                match h.state 
                case LH_HANDSHAKE_A => 
                    assert GetEpochToPropose(h, h', g, g', ios);
                    assert forall io | io in ios && io.LIoOpSend? :: !io.s.msg.FollowerInfo?;
                case LH_HANDSHAKE_B => 
                    assert WaitForEpochAck(h, h', g, g', ios);
                    assert forall io | io in ios && io.LIoOpSend? :: !io.s.msg.FollowerInfo?;
                case LH_PREP_SYNC => 
                    assert PrepareSync(h, h', g, g', ios);
                    assert |ios| == 0;
                case LH_SYNC => 
                    /* Note that this case uses Leader_QueuedPackets_Invariant to make sure
                     that there are no FollowerInfo sneaking inside queued packets */
                    assert ZooKeeper_LearnerHandler.DoSync(h, h', g, g', ios);
                    assert forall io | io in ios && io.LIoOpSend? :: !io.s.msg.FollowerInfo?;
                case LH_PROCESS_ACK => 
                    assert ZooKeeper_LearnerHandler.ProcessAck(h, h', g, g', ios);
                    assert forall io | io in ios && io.LIoOpSend? :: !io.s.msg.FollowerInfo?;
                case LH_RUNNING => assert |ios| == 0;
                case LH_ERROR => assert |ios| == 0;
            }
        }
        assert FollowerInfo_Message_ts_Invariant(tls');  
    }
}


/*****************************************************************************************
*                                 ProcessFI Invariant                                    *
*****************************************************************************************/

lemma lemma_Leader_ProcessFI_PreQuorum_Invariant_Induction(config:Config, tls:TLS_State, tls':TLS_State)
    requires TLS_Next(tls, tls')
    requires General_LS_Performance_Assumption(tls)
    // Pre-established Invariants
    requires Basic_Invariants(config, tls) && Basic_Invariants(config, tls')
    requires FollowerInfo_Message_ts_Invariant(tls) && FollowerInfo_Message_ts_Invariant(tls')
    requires FollowerInit_Invariant(tls);
    requires ProcessFI_PreQuorum_Implies_No_Future_Quorum_Invariant(config, tls)
    requires ProcessFI_PreQuorum_Implies_Only_FI_Messages_Invariant(config, tls)
    // Induction hypothesis
    requires ProcessFI_PreQuorum_Invariant(tls)
    ensures ProcessFI_PreQuorum_Invariant(tls')
{
    var actor, tios:seq<TZKIo> :| actor in tls.t_servers && TLS_NextOneServer(tls, tls', actor, tios);
    var ls, ls', ios := UntagLS_State(tls), UntagLS_State(tls'), UntagLIoOpSeq(tios);
    assert LS_NextOneServer(ls, ls', actor, ios);
    if ls.servers[actor].FollowerPeer? {
        // Theorem is vacuosly true
        assert tls'.t_servers[tls.config[0]] == tls'.t_servers[tls.config[0]];
        return;
    } 
    assert actor == config[0];
    var l, l' := tls.t_servers[actor], tls'.t_servers[actor];
    if IsQuorum(|config|, l.v.leader.globals.connectingFollowers) {
        // Theorem is vacuosly true
        assert ProcessFI_PreQuorum_Invariant(tls');
        return;
    }

    // First prove leaderTQP.dts <= ProcessFI_PreQuorum_dts_Formula();
    if |tios| > 0 && tios[0].LIoOpReceive? {
        assert l'.dts == tios[0].r.msg.ts;
        assert tios[0].r.msg.v.FollowerInfo?;
        assert tios[0].r in tls.t_environment.sentPackets;
        assert l'.dts <= ProcessFI_PreQuorum_dts_Formula();
    } else {
        assert l'.dts == l.dts;
    }

    if |tios| > 0 && tios[0].LIoOpReceive? {
        // TLS_RecvPerfUpdate case
        assert l'.ts == TLS_RecvPerfUpdate(l.ts, tios[0].r.msg.ts, L(ProcessFollowerInfo));
        assert tios[0].r.msg.v.FollowerInfo?;
        assert tios[0].r in tls.t_environment.sentPackets;
        assert l'.ts <= ProcessFI_PreQuorum_dts_Formula() + ProcFI * |l.v.leader.globals.connectingFollowers| + ProcFI;

        // Prove |l'.v.leader.globals.connectingFollowers| == |l.v.leader.globals.connectingFollowers| + 1;
        var follower_id := ios[0].r.sender_index;
        var h, h', g, g' := l.v.leader.handlers[follower_id], l'.v.leader.handlers[follower_id], l.v.leader.globals, l'.v.leader.globals;
        assert GetEpochToPropose(h, h', g, g', ios);
        assert !IsVerifiedQuorum(h.follower_id, |config|, g.connectingFollowers); 
        assert |g'.connectingFollowers| == |g.connectingFollowers| + 1;
        assert l'.ts <= ProcessFI_PreQuorum_ts_Formula(l');
    } else {
        // TLS_NoRecvPerfUpdate case.
        if l.v.leader.state == L_RUNNING {
            assert LeaderStutter(l.v.leader, l'.v.leader, ios);
            assert l'.ts == l.ts + StepToTimeDelta(L(LStutter)) == l.ts;
        } else {
            assert StepSingleHandler_NoRcv(l.v.leader, l'.v.leader, ios);
            var follower_id :| LHNext(l.v.leader, l'.v.leader, follower_id, ios);
            var h, h', g, g' := l.v.leader.handlers[follower_id], l'.v.leader.handlers[follower_id], l.v.leader.globals, l'.v.leader.globals;
            assert LearnerHandlerNext(h, h', g, g', ios);
            assert h.state == LH_HANDSHAKE_A;
            assert false;
        }
    }
}


/*****************************************************************************************
*                                ProcessEpAck Invariant                                  *
*****************************************************************************************/

lemma lemma_Leader_ProcessEpAck_PreQuorum_Invariant_Induction(config:Config, tls:TLS_State, tls':TLS_State)
    requires TLS_Next(tls, tls')
    requires General_LS_Performance_Assumption(tls)
    // Pre-established Invariants
    requires Basic_Invariants(config, tls) && Basic_Invariants(config, tls')
    requires Leader_NextSerialLI_Invariant(tls) && Leader_NextSerialLI_Invariant(tls')
    requires FollowerInit_Invariant(tls);
    // requires FollowerInfo_Message_ts_Invariant(tls) && FollowerInfo_Message_ts_Invariant(tls')
    // requires ProcessFI_PreQuorum_Implies_No_Future_Quorum_Invariant(config, tls)
    // requires ProcessFI_PreQuorum_Implies_Only_FI_Messages_Invariant(config, tls)
    
    // Induction hypothesis
    requires HandShake_Messages_Invariant(tls)
    requires Handshake_Leader_PreQuorum_Invariant(tls)
    requires Handshake_Follower_Invariant(tls)
    ensures HandShake_Messages_Invariant(tls')
    ensures Handshake_Leader_PreQuorum_Invariant(tls')
    ensures Handshake_Follower_Invariant(tls')
{

    // Invariant on followers after SendFI
    forall ep | ep in tls'.t_servers && tls'.t_servers[ep].v.FollowerPeer? && tls'.t_servers[ep].v.follower.state == F_HANDSHAKE_B
    ensures tls'.t_servers[ep].dts == TimeZero() && tls'.t_servers[ep].ts == SendFI {
        var actor, tios:seq<TZKIo> :| actor in tls.t_servers && TLS_NextOneServer(tls, tls', actor, tios);
        if ep == actor {
            var f, f', ios := tls.t_servers[ep].v.follower, tls'.t_servers[ep].v.follower, UntagLIoOpSeq(tios);
            if f.state != F_HANDSHAKE_A {  // else case is true by HandShake_Messages_Invariant
                assert f'.state != F_HANDSHAKE_B; assert false;
            } 
        }
    }

    // TODO: Invariants on leader after ProcessFI
    assume Handshake_Leader_PreQuorum_Invariant(tls');

    // Invariant on LeaderInfo messages
    forall pkt | pkt in tls'.t_environment.sentPackets && pkt.msg.v.LeaderInfo? && pkt.msg.v.serial < tls'.f
    ensures pkt.msg.ts <= LeaderInfo_Message_PreQuorum_ts_Formula(tls', pkt.msg.v.serial)
    {
        lemma_LeaderInfo_Message_PreQuorum_ts_Formula_Helper(config, tls, tls', pkt);
    }

    // Invariant on followers after ProcLI
    forall ep | ep in tls'.t_servers && tls'.t_servers[ep].v.FollowerPeer? && tls'.t_servers[ep].v.follower.state == F_PRESYNC && 0 <= tls'.t_servers[ep].v.follower.serialLI < tls'.f
    ensures tls'.t_servers[ep].dts <= LeaderInfo_Message_PreQuorum_ts_Formula(tls', tls'.t_servers[ep].v.follower.serialLI)
            && tls'.t_servers[ep].ts <= LeaderInfo_Message_PreQuorum_ts_Formula(tls', tls'.t_servers[ep].v.follower.serialLI) + ProcLI
    {
        var actor, tios:seq<TZKIo> :| actor in tls.t_servers && TLS_NextOneServer(tls, tls', actor, tios);
        if ep == actor {
            var f, f', ios := tls.t_servers[ep].v.follower, tls'.t_servers[ep].v.follower, UntagLIoOpSeq(tios);
            if f.state != F_HANDSHAKE_B { 
                assert f'.state != F_PRESYNC; assert false;
            } 
        }
    }

    // Invariant on AckEpoch messages
    assume (forall pkt | pkt in tls'.t_environment.sentPackets && pkt.msg.v.AckEpoch? && pkt.msg.v.serial < tls'.f
    :: pkt.msg.ts <= AckEpoch_Message_PreQuorum_ts_Formula(tls', pkt.msg.v.serial));
}


lemma lemma_LeaderInfo_Message_PreQuorum_ts_Formula_Helper(config:Config, tls:TLS_State, tls':TLS_State, pkt:LPacket<EndPoint,TimestampedType<ZKMessage>>)
    requires TLS_Next(tls, tls')
    requires General_LS_Performance_Assumption(tls)
    // Pre-established Invariants
    requires Basic_Invariants(config, tls) && Basic_Invariants(config, tls')
    requires Leader_NextSerialLI_Invariant(tls) && Leader_NextSerialLI_Invariant(tls')
    // Local conditions
    requires pkt in tls'.t_environment.sentPackets && pkt.msg.v.LeaderInfo? && pkt.msg.v.serial < tls'.f
    // Induction hypothesis
    requires HandShake_Messages_Invariant(tls)
    requires Handshake_Leader_PreQuorum_Invariant(tls)
    requires Handshake_Follower_Invariant(tls)
    ensures pkt.msg.ts <= LeaderInfo_Message_PreQuorum_ts_Formula(tls', pkt.msg.v.serial)
{
    var ps := pkt.msg.v.serial;
    if pkt in tls.t_environment.sentPackets {
        assert LeaderInfo_Message_PreQuorum_ts_Formula(tls', ps) == LeaderInfo_Message_PreQuorum_ts_Formula(tls, ps);
        // assert  pkt.msg.ts <= LeaderInfo_Message_PreQuorum_ts_Formula(tls', ps);
    } else {
        var actor, tios:seq<TZKIo> :| actor in tls.t_servers && TLS_NextOneServer(tls, tls', actor, tios);
        assert tls.t_servers[actor].v.LeaderPeer?;
        assert actor == config[0];
        var l, l', ios := tls.t_servers[actor].v.leader, tls'.t_servers[actor].v.leader, UntagLIoOpSeq(tios);
        if StepSingleHandler_Rcv(l, l', ios) {
            assert false;  // Not possible to send a new LI packet in this step
        } else {
            assert StepSingleHandler_NoRcv(l, l', ios);
            var fid :| LHNext(l, l', fid, ios);
            var h , h', g, g' := l.handlers[fid], l'.handlers[fid], l.globals, l'.globals;
            assert LearnerHandlerNext(h, h', g, g', ios);
            if h.state == LH_HANDSHAKE_A {
                var lts, ldts, n := tls.t_servers[actor].ts, tls.t_servers[actor].dts, |config|;
                assert lts <= ProcessEpAck_PreQuorum_ts_Formula(tls, tls.t_servers[actor]);
                if |g.electingFollowers| <= 1 {
                    if pkt.msg.v.serial == 0 {
                        // Case pkt.msg.v.serial == 0, so this is the first LI sent
                        assert lts <= SendFI + D + ProcFI * n;
                        assert pkt.msg.ts <= SendFI + D + ProcFI * n + ProcFI + D;
                        assert pkt.msg.ts <= LeaderInfo_Message_PreQuorum_ts_Formula(tls', ps);
                    } else {
                        // Case pkt.msg.v.serial > 0, so this is NOT the first LI sent
                        assert lts <= SendFI + D + ProcFI * n + ProcFI * g.nextSerialLI;
                        assert pkt.msg.ts 
                            <= lts + ProcFI + D
                            <=  SendFI + D 
                                + ProcFI * n 
                                + ProcFI * (g.nextSerialLI + 1) + D;
                        assert (ProcFI + D) * (g.nextSerialLI  + 1) == ProcFI * (g.nextSerialLI + 1) + D * (g.nextSerialLI + 1);
                        assert pkt.msg.ts <= LeaderInfo_Message_PreQuorum_ts_Formula(tls', ps);
                    }
                } else { 
                    assert ProcEpAck * (|g.electingFollowers|-2) + ProcEpAck == ProcEpAck * (|g.electingFollowers|-1);
                    assert pkt.msg.v.serial == g.nextSerialLI;
                    assert lts <= SendFI + D 
                                    + ProcFI * n
                                    + (ProcFI + D) * g.nextSerialLI
                                    + (ProcLI + D) * g.nextSerialLI
                                    + ProcEpAck * (|g.electingFollowers|-1);
                    assert pkt.msg.ts == lts + ProcFI + D;
                    assert |g.electingFollowers| - 1 <= g.nextSerialLI;
                    lemma_Math_Mult();
                    assert ProcEpAck * (|g.electingFollowers|-1) <= ProcEpAck * g.nextSerialLI;
                    assert pkt.msg.ts <= LeaderInfo_Message_PreQuorum_ts_Formula(tls', ps);
                }
            } else {
                // h.state is not HANDSHAKE_A
                assert false; // Not possible to send a new LI packet in this step
            }
        }
    }
}

}
