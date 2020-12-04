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
include "ProtocolInvariants.dfy"
include "Commons.dfy"


module Zookeeper_PerformanceProof_Invariants {
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
import opened Zookeeper_ProtocoIInvariants
import opened Zookeeper_Commons



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
    requires Basic_Invariants(config, tls) && Basic_Invariants(config, tls')
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
    var ld, ld' := tls.t_servers[actor], tls'.t_servers[actor];
    if IsQuorum(|config|, ld.v.leader.globals.connectingFollowers) {
        // Theorem is vacuosly true
        assert ProcessFI_PreQuorum_Invariant(tls');
        return;
    }

    // First prove leaderTQP.dts <= ProcessFI_PreQuorum_dts_Formula();
    if |tios| > 0 && tios[0].LIoOpReceive? {
        assert ld.dts == tios[0].r.msg.ts;
        assert tios[0].r.msg.v.FollowerInfo?;
        assert tios[0].r in tls.t_environment.sentPackets;
        assert ld'.dts <= ProcessFI_PreQuorum_dts_Formula();
    } else {
        assert ld'.dts == ld.dts;
    }

    if |tios| > 0 && tios[0].LIoOpReceive? {
        // TLS_RecvPerfUpdate case
        assert ld'.ts == TLS_RecvPerfUpdate(ld.ts, tios[0].r.msg.ts, L(ProcessFollowerInfo));
        assert tios[0].r.msg.v.FollowerInfo?;
        assert tios[0].r in tls.t_environment.sentPackets;
        assert ld'.ts <= ProcessFI_PreQuorum_dts_Formula() + ProcFI * |ld.v.leader.globals.connectingFollowers| + ProcFI;

        // Prove |ld'.v.leader.globals.connectingFollowers| == |ld.v.leader.globals.connectingFollowers| + 1;
        var follower_id := ios[0].r.sender_index;
        var h, h', g, g' := ld.v.leader.handlers[follower_id], ld'.v.leader.handlers[follower_id], ld.v.leader.globals, ld'.v.leader.globals;
        assert GetEpochToPropose(h, h', g, g', ios);
        assert !IsVerifiedQuorum(h.follower_id, |config|, g.connectingFollowers); 
        assert |g'.connectingFollowers| == |g.connectingFollowers| + 1;
        assert ld'.ts <= ProcessFI_PreQuorum_ts_Formula(ld');
    } else {
        // TLS_NoRecvPerfUpdate case.
        if ld.v.leader.state == L_RUNNING {
            assert LeaderStutter(ld.v.leader, ld'.v.leader, ios);
            assert ld'.ts == ld.ts + StepToTimeDelta(L(LStutter)) == ld.ts;
        } else {
            assert StepSingleHandler_NoRcv(ld.v.leader, ld'.v.leader, ios);
            var follower_id :| LHNext(ld.v.leader, ld'.v.leader, follower_id, ios);
            var h, h', g, g' := ld.v.leader.handlers[follower_id], ld'.v.leader.handlers[follower_id], ld.v.leader.globals, ld'.v.leader.globals;
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
    requires Basic_Invariants(config, tls) && Basic_Invariants(config, tls')
    requires FollowerInit_Invariant(tls) && FollowerInit_Invariant(tls')
    requires QuorumsMonotoneIncreasing_Property(tls, tls')
    // Induction hypothesis
    requires Handshake_Messages_Invariant(tls)
    requires Handshake_Leader_PreQuorum_Invariant(tls)
    requires Handshake_Follower_Invariant(tls)
    ensures Handshake_Messages_Invariant(tls')
    ensures Handshake_Leader_PreQuorum_Invariant(tls')
    ensures Handshake_Follower_Invariant(tls')
{
    // Invariant on followers after sending FollowerInfo messages
    forall ep | ep in tls'.t_servers && tls'.t_servers[ep].v.FollowerPeer? && tls'.t_servers[ep].v.follower.state == F_HANDSHAKE_B
    ensures tls'.t_servers[ep].dts == TimeZero() && tls'.t_servers[ep].ts == SendFI {
        var actor, tios:seq<TZKIo> :| actor in tls.t_servers && TLS_NextOneServer(tls, tls', actor, tios);
        if ep == actor {
            var f, f', ios := tls.t_servers[ep].v.follower, tls'.t_servers[ep].v.follower, UntagLIoOpSeq(tios);
            if f.state != F_HANDSHAKE_A {  // else case is true by Handshake_Messages_Invariant
                assert f'.state != F_HANDSHAKE_B; assert false;
            } 
        }
    }
    // Invariant on LeaderInfo messages 
    forall pkt | pkt in tls'.t_environment.sentPackets && pkt.msg.v.LeaderInfo? 
    ensures pkt.msg.ts <= AckEpoch_Message_PreQuorum_ts_Formula(tls'.f, pkt.msg.v.serial)
    {}
    // Invariant on EpochAck messages 
    forall pkt | pkt in tls'.t_environment.sentPackets && pkt.msg.v.AckEpoch? 
    ensures pkt.msg.ts <= AckEpoch_Message_PreQuorum_ts_Formula(tls'.f, pkt.msg.v.serial)
    {}
    // Invariants on leader at syncrhonization barrier at ProcessEpAck -- Handshake_Leader_PreQuorum_Invariant(tls')
    lemma_Leader_ProcessEpAck_PreQuorum_Invariant_Induction_Helper(config, tls, tls');
}

lemma lemma_Leader_ProcessEpAck_PreQuorum_Invariant_Induction_Helper(config:Config, tls:TLS_State, tls':TLS_State)
    requires TLS_Next(tls, tls')
    requires General_LS_Performance_Assumption(tls)
    requires Basic_Invariants(config, tls) && Basic_Invariants(config, tls')
    requires FollowerInit_Invariant(tls) && FollowerInit_Invariant(tls')
    requires QuorumsMonotoneIncreasing_Property(tls, tls')
    // Induction hypothesis
    requires Handshake_Messages_Invariant(tls)
    requires Handshake_Leader_PreQuorum_Invariant(tls)
    requires Handshake_Follower_Invariant(tls)
    requires Handshake_Follower_Invariant(tls')
    requires Handshake_Messages_Invariant(tls')
    ensures Handshake_Leader_PreQuorum_Invariant(tls')
{
    if !IsQuorum(|tls.config|, tls.t_servers[config[0]].v.leader.globals.electingFollowers) {
        var actor, tios:seq<TZKIo> :| actor in tls.t_servers && TLS_NextOneServer(tls, tls', actor, tios);
        if actor == config[0] {
            var l, l', ios := tls.t_servers[config[0]].v.leader, tls'.t_servers[config[0]].v.leader, UntagLIoOpSeq(tios);
            var lt, lt' := tls.t_servers[config[0]], tls'.t_servers[config[0]];
            if l.state == L_RUNNING {
                assert tls'.t_servers[actor].dts == tls.t_servers[actor].dts && tls'.t_servers[actor].ts == tls.t_servers[actor].ts;
            } else {
                if StepSingleHandler_Rcv(l, l', ios) {
                    var fid := ios[0].r.sender_index;
                    var h , h', g, g' := l.handlers[ios[0].r.sender_index], l'.handlers[ios[0].r.sender_index], l.globals, l'.globals;
                    var f, n := tls.f, |tls.config|;
                    if h.state == LH_HANDSHAKE_B {
                        assert ios[0].r.msg.AckEpoch?;
                        var tmsg := tios[0].r.msg;
                        if !IsQuorum(|tls'.config|, tls'.t_servers[config[0]].v.leader.globals.electingFollowers) {
                            if |g.electingFollowers| <= 1 {
                                // First deal with the dts
                                lemma_Size_One_Sets(g.electingFollowers, 0);
                                assert g'.electingFollowers == g.electingFollowers + {tmsg.v.sid};
                                assert tmsg.v.sid != 0;
                                assert |g'.electingFollowers| == 2;
                                lemma_Math_Inequalities_CommonMult(ProcFI, tmsg.v.serial + 1, f);
                                assert lt'.dts <= ProcessEpAck_PreQuorum_dts_Formula(f, n, tls'.t_servers[actor]);
                                
                                // Next is the ts
                                lemma_Math_Inequalities_CommonMult(ProcFI, |g.connectingFollowers|-1, f);
                                lemma_Math_Inequalities_CommonMult(ProcFI, g.nextSerialLI, f);
                                var handlerStartTime := TimeMax(tmsg.ts, lt.ts);
                                assert handlerStartTime <= AckEpoch_Message_PreQuorum_ts_Formula(f, f-1);
                                assert lt'.ts  == handlerStartTime + ProcEpAck;                       
                                assert lt'.ts <= ProcessEpAck_PreQuorum_ts_Formula(f, n, tls'.t_servers[actor]);
                            } else {
                                // First deal with the dts
                                lemma_Math_Inequalities_CommonMult(ProcFI, tmsg.v.serial+1, f);
                                assert AckEpoch_Message_PreQuorum_ts_Formula(f, tmsg.v.serial) <= AckEpoch_Message_PreQuorum_ts_Formula(f, f-1);
                                assert lt'.dts <= ProcessEpAck_PreQuorum_dts_Formula(f, n, tls'.t_servers[actor]);
                                
                                // Next is the ts
                                assert lt.ts <= AckEpoch_Message_PreQuorum_ts_Formula(f, f-1) + ProcEpAck * (|g.electingFollowers|-1);
                                assert tmsg.ts <= AckEpoch_Message_PreQuorum_ts_Formula(f, tmsg.v.serial);
                                assert |g.electingFollowers|-1 >= 0;
                                lemma_Math_Inequalities_Mult();
                                assert ProcEpAck * (|g.electingFollowers|-1) >= 0;
                                lemma_Math_Addition(ProcEpAck * (|g.electingFollowers|-1));
                                assert AckEpoch_Message_PreQuorum_ts_Formula(f, tmsg.v.serial) <= AckEpoch_Message_PreQuorum_ts_Formula(f, f-1) <= AckEpoch_Message_PreQuorum_ts_Formula(f, f-1) + ProcEpAck * (|g.electingFollowers|-1);
                                var handlerStartTime := TimeMax(tmsg.ts, lt.ts);
                                lemma_Math_MaxOfInequalities(tmsg.ts, lt.ts, AckEpoch_Message_PreQuorum_ts_Formula(f, tmsg.v.serial), AckEpoch_Message_PreQuorum_ts_Formula(f, f-1) + ProcEpAck * (|g.electingFollowers|-1));
                                assert handlerStartTime <= AckEpoch_Message_PreQuorum_ts_Formula(f, f-1) + ProcEpAck * (|g.electingFollowers|-1);
                                assert g'.electingFollowers == g.electingFollowers + {tmsg.v.sid};
                                assert |g'.electingFollowers| == |g.electingFollowers| + 1;
                                assert lt'.ts <= AckEpoch_Message_PreQuorum_ts_Formula(f, f-1) + ProcEpAck * (|g.electingFollowers|-1) + ProcEpAck;
                                lemma_Math_Mult_b();
                                assert ProcEpAck * (|g.electingFollowers|-1) + ProcEpAck == ProcEpAck * (|g'.electingFollowers|-1);
                                assert lt'.ts <= ProcessEpAck_PreQuorum_ts_Formula(f, n, tls'.t_servers[actor]);
                            }
                        }
                    }
                }
            }
        }
    }
}



/*****************************************************************************************
*                                  ProcessAck Invariant                                  *
*****************************************************************************************/


lemma lemma_Leader_ProcessAck_PreQuorum_Invariant_Induction(config:Config, tls:TLS_State, tls':TLS_State)
    requires TLS_Next(tls, tls')
    requires General_LS_Performance_Assumption(tls)
    requires Basic_Invariants(config, tls) && Basic_Invariants(config, tls')
    requires EmptyDiff_Invariant(tls) && EmptyDiff_Invariant(tls')
    // requires FollowerInit_Invariant(tls) && FollowerInit_Invariant(tls')
    // requires QuorumsMonotoneIncreasing_Property(tls, tls')
    // Induction hypothesis
    requires Sync_Messages_Invariant(tls)
    requires Sync_Follower_Invariant(tls)
    requires Sync_Leader_PreQuorum_Invariant(tls)
    ensures Sync_Messages_Invariant(tls')
    ensures Sync_Follower_Invariant(tls')
    ensures Sync_Leader_PreQuorum_Invariant(tls')
{
    var f, n := tls.f, |tls.config|;

    // First prove Sync_Messages_Invariant(tls');
    forall pkt | pkt in tls'.t_environment.sentPackets && (pkt.msg.v.SyncDIFF? || pkt.msg.v.SyncSNAP?)
    ensures pkt.msg.ts <= Sync_Message_ts_Formula(f, pkt.msg.v.serial)
    {
        assert pkt.msg.v.SyncDIFF?;  // By EmptyDiff_Invariant
        Sync_Messages_Invariant_Helper_A(config, tls, tls', pkt);
    }
    forall pkt | pkt in tls'.t_environment.sentPackets && pkt.msg.v.NewLeader?
    ensures pkt.msg.ts <= NewLeader_Message_ts_Formula(tls.f, pkt.msg.v.serial)
    {
        Sync_Messages_Invariant_Helper_B(config, tls, tls', pkt);
    }

    assume Sync_Messages_Invariant(tls');
    assume Sync_Follower_Invariant(tls');
    assume Sync_Leader_PreQuorum_Invariant(tls');
}


lemma Sync_Messages_Invariant_Helper_A(config:Config, tls:TLS_State, tls':TLS_State, pkt:TimestampedLPacket<EndPoint,ZKMessage>) 
    requires TLS_Next(tls, tls')
    requires General_LS_Performance_Assumption(tls)
    requires Basic_Invariants(config, tls) && Basic_Invariants(config, tls')
    // Induction hypothesis
    requires Sync_Messages_Invariant(tls)
    requires Sync_Follower_Invariant(tls)
    requires Sync_Leader_PreQuorum_Invariant(tls)
    requires pkt in tls'.t_environment.sentPackets && pkt.msg.v.SyncDIFF?
    ensures pkt.msg.ts <= Sync_Message_ts_Formula(tls.f, pkt.msg.v.serial)
{
    var f, n := tls.f, |tls.config|;
    if pkt !in tls.t_environment.sentPackets {
        var actor, tios:seq<TZKIo> :| actor in tls.t_servers && TLS_NextOneServer(tls, tls', actor, tios);
        var lt, lt' := tls.t_servers[config[0]], tls'.t_servers[config[0]];
        var l, l', ios := lt.v.leader, lt'.v.leader, UntagLIoOpSeq(tios);
        if actor == config[0] {
            var fid :| LHNext(l, l', fid, ios);
            var h , h', g, g' := l.handlers[fid], l'.handlers[fid], l.globals, l'.globals;
            assert ZooKeeper_LearnerHandler.DoSync(h, h', g, g', ios);
            lemma_Math_Mult_b();
            assert g'.prepCount ==  g.prepCount <= f;
            assert g'.nextSerialNL == g.nextSerialNL <= g.nextSerialSync == pkt.msg.v.serial;
            lemma_Math_Inequalities_CommonMult(Sync, g'.nextSerialNL, pkt.msg.v.serial);
            lemma_Math_Inequalities_CommonMult(PreSync, g'.prepCount, f);
            assert pkt.msg.ts <= Sync_Message_ts_Formula(f, pkt.msg.v.serial);
        }
    }
}

lemma Sync_Messages_Invariant_Helper_B(config:Config, tls:TLS_State, tls':TLS_State, pkt:TimestampedLPacket<EndPoint,ZKMessage>) 
    requires TLS_Next(tls, tls')
    requires General_LS_Performance_Assumption(tls)
    requires Basic_Invariants(config, tls) && Basic_Invariants(config, tls')
    // Induction hypothesis
    requires Sync_Messages_Invariant(tls)
    requires Sync_Follower_Invariant(tls)
    requires Sync_Leader_PreQuorum_Invariant(tls)
    requires pkt in tls'.t_environment.sentPackets && pkt.msg.v.NewLeader?
    ensures pkt.msg.ts <= NewLeader_Message_ts_Formula(tls.f, pkt.msg.v.serial)
{
    var f, n := tls.f, |tls.config|;
    if pkt !in tls.t_environment.sentPackets {
        var actor, tios:seq<TZKIo> :| actor in tls.t_servers && TLS_NextOneServer(tls, tls', actor, tios);
        var lt, lt' := tls.t_servers[config[0]], tls'.t_servers[config[0]];
        var l, l', ios := lt.v.leader, lt'.v.leader, UntagLIoOpSeq(tios);
        if actor == config[0] {
            var fid :| LHNext(l, l', fid, ios);
            var h , h', g, g' := l.handlers[fid], l'.handlers[fid], l.globals, l'.globals;
            assert ZooKeeper_LearnerHandler.DoSync(h, h', g, g', ios);
            lemma_Math_Mult_b();
            lemma_Math_Inequalities_CommonMult(Sync, g.nextSerialSync, f);
            lemma_Math_Inequalities_CommonMult(PreSync, g'.prepCount, f);
            assert pkt.msg.ts <= NewLeader_Message_ts_Formula(f, pkt.msg.v.serial);
        }
    }
}


}
