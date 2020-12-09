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
include "CorrectZKInvariants.dfy"


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
import opened Zookeeper_CorrectInvariants


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
    forall pkt | pkt in tls'.t_environment.sentPackets && pkt.msg.v.FollowerInfo?
    ensures pkt.msg.ts == FollowerInfo_Message_ts_Formula()
    {}
}


/*****************************************************************************************
*                                 ProcessFI Invariant                                    *
*****************************************************************************************/

lemma lemma_Leader_ProcessFI_PreQuorum_Invariant_Induction(config:Config, tls:TLS_State, tls':TLS_State)
    requires TLS_Next(tls, tls')
    requires General_LS_Performance_Assumption(tls)
    requires Basic_Invariants(config, tls) && Basic_Invariants(config, tls')
    // Induction hypothesis
    requires FollowerInfo_Message_ts_Invariant(tls);
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
        assert ld'.dts == tios[0].r.msg.ts;
        assert tios[0].r.msg.v.FollowerInfo?;
        assert tios[0].r in tls.t_environment.sentPackets;
        assert tios[0].r.msg.ts == FollowerInfo_Message_ts_Formula();
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
        lemma_Math_Mult_b();
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
    requires FollowerInfo_Message_ts_Invariant(tls) && FollowerInfo_Message_ts_Invariant(tls')
    requires ProcessFI_PreQuorum_Invariant(tls) && ProcessFI_PreQuorum_Invariant(tls'); 
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

    assert Follower_HandshakeB_Invariant(tls');
    // Invariant on LeaderInfo messages 
    forall pkt | pkt in tls'.t_environment.sentPackets && pkt.msg.v.LeaderInfo? 
    ensures pkt.msg.ts <= LeaderInfo_Message_PreQuorum_ts_Formula(tls'.f, pkt.msg.v.serial)
    {
        var f := tls'.f;
        if pkt !in tls.t_environment.sentPackets {
            var actor, tios:seq<TZKIo> :| actor in tls.t_servers && TLS_NextOneServer(tls, tls', actor, tios);
            var lt, lt' := tls.t_servers[config[0]], tls'.t_servers[config[0]];
            var l, l', ios := lt.v.leader, lt'.v.leader, UntagLIoOpSeq(tios);
            if actor == config[0] {
                if |l.globals.electingFollowers| <= 1 {
                    assert lt.ts 
                        <=  SendFI + D
                            + ProcFI * (|l.globals.connectingFollowers|-1)
                            + ProcFI * l.globals.nextSerialLI;
                    lemma_Math_Mult_b();
                    lemma_Math_Inequalities_CommonMult(ProcFI, |l.globals.connectingFollowers| - 1, f);
                    lemma_Math_Inequalities_CommonMult(ProcFI, l.globals.nextSerialLI, pkt.msg.v.serial + 1);
                    assert l'.globals.nextSerialLI == l.globals.nextSerialLI + 1;
                    assert pkt.msg.ts
                        <=  SendFI + D 
                            + ProcFI * f 
                            + ProcFI * (pkt.msg.v.serial + 1)
                            + D;
                    assert pkt.msg.ts <= LeaderInfo_Message_PreQuorum_ts_Formula(tls'.f, pkt.msg.v.serial);
                } else {
                    // This case isn't possible, by Leader_Sends_All_LI_Before_Receiving_EpochAck_Invariant
                    assert l.globals.nextSerialLI == f;
                    assert false;
                }
            }
        }
    }

    // Invariant on followers after sending AckEpoch messages
    forall ep | ep in tls'.t_servers && tls'.t_servers[ep].v.FollowerPeer? && tls'.t_servers[ep].v.follower.state == F_PRESYNC && 0 <= tls'.t_servers[ep].v.follower.serialLI < tls.f
    ensures && tls'.t_servers[ep].dts <= LeaderInfo_Message_PreQuorum_ts_Formula(tls'.f, tls'.t_servers[ep].v.follower.serialLI)
            && tls'.t_servers[ep].ts <= LeaderInfo_Message_PreQuorum_ts_Formula(tls'.f, tls'.t_servers[ep].v.follower.serialLI) + ProcLI
    {
        var actor, tios:seq<TZKIo> :| actor in tls.t_servers && TLS_NextOneServer(tls, tls', actor, tios);
        if ep == actor {
            var f, f', ios := tls.t_servers[ep].v.follower, tls'.t_servers[ep].v.follower, UntagLIoOpSeq(tios);
            var ft, ft' := tls.t_servers[ep], tls'.t_servers[ep];
            if f.state != F_HANDSHAKE_B {  // else case is true by Handshake_Messages_Invariant
                assert f'.state != F_PRESYNC; assert false;
            } else {
                lemma_Math_Addition();
                assert f.state == F_HANDSHAKE_B;
                assert ft.ts == SendFI;
                var pkt := tios[0].r;
                assert pkt.msg.v.LeaderInfo?;
                assert pkt.msg.ts <= SendFI + D 
                                    + ProcFI * tls'.f
                                    + ProcFI * (pkt.msg.v.serial + 1) 
                                    + D;
                lemma_Math_Addition_2(SendFI, D + ProcFI * tls'.f + ProcFI * (pkt.msg.v.serial + 1) + D);
                lemma_Math_MaxOfInequalities(ft.ts, pkt.msg.ts, SendFI, SendFI + D + ProcFI * tls'.f + ProcFI * (pkt.msg.v.serial + 1) + D);
                assert TimeMax(ft.ts, pkt.msg.ts) <= SendFI + D 
                                    + ProcFI * tls'.f 
                                    + ProcFI * (pkt.msg.v.serial + 1) 
                                    + D;
                assert ft'.ts <= SendFI + D + ProcFI * tls'.f + ProcFI * (pkt.msg.v.serial + 1) + D + ProcLI;
                assert ft'.ts <= LeaderInfo_Message_PreQuorum_ts_Formula(tls'.f, f'.serialLI) + ProcLI;
            }
        } else {
            assert tls'.t_servers[ep] == tls.t_servers[ep];
        }
    }
    assert Follower_PreSync_Invariant(tls');

    // Invariant on EpochAck messages 
    forall pkt | pkt in tls'.t_environment.sentPackets && pkt.msg.v.AckEpoch? 
    ensures pkt.msg.ts <= AckEpoch_Message_PreQuorum_ts_Formula(tls'.f, pkt.msg.v.serial)
    {
        if pkt !in tls.t_environment.sentPackets {
            var actor, tios:seq<TZKIo> :| actor in tls.t_servers && TLS_NextOneServer(tls, tls', actor, tios);
            var st, st' := tls.t_servers[actor], tls'.t_servers[actor];
            assert exists tio :: tio in tios && tio.LIoOpSend? && tio.s == pkt;
            if st.v.FollowerPeer? {
                var s, s', ios := st.v.follower, st'.v.follower, UntagLIoOpSeq(tios);
                assert AcceptNewEpoch(s, s', ios);
                assert pkt.msg.ts <= AckEpoch_Message_PreQuorum_ts_Formula(tls'.f, pkt.msg.v.serial);
            } else {
                forall tio | tio in tios && tio.LIoOpSend? 
                ensures !tio.s.msg.v.AckEpoch?
                {}
                assert false;
            }
        }
    }

    // Invariants on leader at syncrhonization barrier at ProcessEpAck -- Handshake_Leader_PreQuorum_Invariant(tls')
    assert Handshake_Messages_Invariant(tls');
    lemma_Leader_ProcessEpAck_PreQuorum_Invariant_Induction_Helper(config, tls, tls');
}

lemma lemma_Leader_ProcessEpAck_PreQuorum_Invariant_Induction_Helper(config:Config, tls:TLS_State, tls':TLS_State)
    requires TLS_Next(tls, tls')
    requires General_LS_Performance_Assumption(tls)
    requires Basic_Invariants(config, tls) && Basic_Invariants(config, tls')
    requires FollowerInit_Invariant(tls) && FollowerInit_Invariant(tls')
    requires FollowerInfo_Message_ts_Invariant(tls) && FollowerInfo_Message_ts_Invariant(tls')
    requires ProcessFI_PreQuorum_Invariant(tls) && ProcessFI_PreQuorum_Invariant(tls'); 
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
        if !IsQuorum(|tls.config|, tls.t_servers[config[0]].v.leader.globals.connectingFollowers) {
            assert Handshake_Leader_PreQuorum_Invariant(tls');
        } else {
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
                                    lemma_Math_Addition();
                                    // lemma_Math_Addition(ProcEpAck * (|g.electingFollowers|-1));
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
}



/*****************************************************************************************
*                                  ProcessAck Invariant                                  *
*****************************************************************************************/


lemma lemma_Leader_ProcessAck_PreQuorum_Invariant_Induction(config:Config, tls:TLS_State, tls':TLS_State)
    requires TLS_Next(tls, tls')
    requires General_LS_Performance_Assumption(tls)
    requires Basic_Invariants(config, tls) && Basic_Invariants(config, tls')
    requires EmptyDiff_Invariant(tls) && EmptyDiff_Invariant(tls')
    requires FollowerInfo_Message_ts_Invariant(tls) && FollowerInfo_Message_ts_Invariant(tls'); 
    requires ProcessFI_PreQuorum_Invariant(tls) && ProcessFI_PreQuorum_Invariant(tls'); 
    requires Handshake_Messages_Invariant(tls) && Handshake_Messages_Invariant(tls')
    requires Handshake_Follower_Invariant(tls) && Handshake_Follower_Invariant(tls')
    requires Handshake_Leader_PreQuorum_Invariant(tls) && Handshake_Leader_PreQuorum_Invariant(tls')
    // Induction hypothesis
    requires Sync_Messages_Invariant(tls)
    requires Sync_Follower_Invariant(tls)
    requires Sync_Leader_PreQuorum_Invariant(tls)
    ensures Sync_Messages_Invariant(tls')
    ensures Sync_Follower_Invariant(tls')
    ensures Sync_Leader_PreQuorum_Invariant(tls')
{
    var f, n := tls.f, |tls.config|;

    // Prove Sync_Follower_Invariant(tls')
    forall ep | 
        && ep in tls'.t_servers
        && tls'.t_servers[ep].v.FollowerPeer? 
        && tls'.t_servers[ep].v.follower.state == F_SYNC
        && 0 <= tls'.t_servers[ep].v.follower.serialSync < tls'.f
    ensures && tls'.t_servers[ep].dts <= Follower_F_SYNC_dts_Formula(tls'.f, tls'.t_servers[ep])
            && tls'.t_servers[ep].ts <= Follower_F_SYNC_ts_Formula(tls'.f, tls'.t_servers[ep])
    {
        Sync_Follower_Invariant_Helper(config, tls, tls', ep);
    }
    assert Sync_Follower_Invariant(tls');

    // Prove Sync_Messages_Invariant(tls');
    forall pkt | pkt in tls'.t_environment.sentPackets && (pkt.msg.v.SyncDIFF? || pkt.msg.v.SyncSNAP?)
    ensures pkt.msg.ts <= Sync_Message_ts_Formula(f, pkt.msg.v.serial) {
        assert pkt.msg.v.SyncDIFF?;  // By EmptyDiff_Invariant
        Sync_Messages_Invariant_Helper_A(config, tls, tls', pkt);
    }
    forall pkt | pkt in tls'.t_environment.sentPackets && pkt.msg.v.NewLeader?
    ensures pkt.msg.ts <= NewLeader_Message_ts_Formula(tls.f, pkt.msg.v.serial) {
        Sync_Messages_Invariant_Helper_B(config, tls, tls', pkt);
    }

    forall pkt | pkt in tls'.t_environment.sentPackets && pkt.msg.v.Ack?
    ensures pkt.msg.ts <= Ack_Message_ts_Formula(tls.f, pkt.msg.v.serial) {
        Sync_Messages_Invariant_Helper_C(config, tls, tls', pkt);
    }
    assert Sync_Messages_Invariant(tls');

    // Prove Sync_Leader_PreQuorum_Invariant(tls')
    if !IsQuorum(n, tls'.t_servers[config[0]].v.leader.globals.ackSet) {
        Sync_Leader_PreQuorum_Invariant_Helper(config, tls, tls');
    }
    assert Sync_Leader_PreQuorum_Invariant(tls');
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
            lemma_Math_Inequalities_CommonMult(ProcEpAck, g'.procEpCount, f);
            lemma_Math_Inequalities_CommonMult(Sync, g'.nextSerialNL, pkt.msg.v.serial);
            lemma_Math_Inequalities_CommonMult(Sync, g.nextSerialSync + 1, pkt.msg.v.serial + 1);
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
            lemma_Math_Inequalities_CommonMult(ProcEpAck, g'.procEpCount, f);
            lemma_Math_Inequalities_CommonMult(Sync, g.nextSerialSync, f);
            lemma_Math_Inequalities_CommonMult(PreSync, g'.prepCount, f);
            assert pkt.msg.ts <= NewLeader_Message_ts_Formula(f, pkt.msg.v.serial);
        }
    }
}

lemma Sync_Messages_Invariant_Helper_C(config:Config, tls:TLS_State, tls':TLS_State, pkt:TimestampedLPacket<EndPoint,ZKMessage>) 
    requires TLS_Next(tls, tls')
    requires General_LS_Performance_Assumption(tls)
    requires Basic_Invariants(config, tls) && Basic_Invariants(config, tls')
    // Induction hypothesis
    requires Sync_Messages_Invariant(tls)
    requires Sync_Follower_Invariant(tls) && Sync_Follower_Invariant(tls')
    requires Sync_Leader_PreQuorum_Invariant(tls)
    requires pkt in tls'.t_environment.sentPackets && pkt.msg.v.Ack?
    ensures pkt.msg.ts <= Ack_Message_ts_Formula(tls.f, pkt.msg.v.serial)
{
    var f, n := tls.f, |tls.config|;
    if pkt !in tls.t_environment.sentPackets {
        var actor, tios:seq<TZKIo> :| actor in tls.t_servers && TLS_NextOneServer(tls, tls', actor, tios);
        assert tls'.t_servers[actor].v.follower.serialSync >= 0;
        assert pkt.msg.ts <= Ack_Message_ts_Formula(f, pkt.msg.v.serial);
    }
}

lemma Sync_Follower_Invariant_Helper(config:Config, tls:TLS_State, tls':TLS_State, ep:EndPoint) 
    requires TLS_Next(tls, tls')
    requires General_LS_Performance_Assumption(tls)
    requires Basic_Invariants(config, tls) && Basic_Invariants(config, tls')
    requires EmptyDiff_Invariant(tls) && EmptyDiff_Invariant(tls')
    requires Handshake_Follower_Invariant(tls) && Handshake_Follower_Invariant(tls')
    // Induction hypothesis
    requires Sync_Messages_Invariant(tls)
    requires Sync_Follower_Invariant(tls)
    requires Sync_Leader_PreQuorum_Invariant(tls)
    // Antecedent
    requires ep in tls'.t_servers
    requires tls'.t_servers[ep].v.FollowerPeer? 
    requires tls'.t_servers[ep].v.follower.state == F_SYNC
    requires 0 <= tls'.t_servers[ep].v.follower.serialSync < tls'.f
    // Conclusion
    ensures tls'.t_servers[ep].dts <= Follower_F_SYNC_dts_Formula(tls'.f, tls'.t_servers[ep])
    ensures tls'.t_servers[ep].ts <= Follower_F_SYNC_ts_Formula(tls'.f, tls'.t_servers[ep])
{
    var f, n := tls.f, |tls.config|;
    var actor, tios:seq<TZKIo> :| actor in tls.t_servers && TLS_NextOneServer(tls, tls', actor, tios);
    var ls, ls', ios := UntagLS_State(tls), UntagLS_State(tls'), UntagLIoOpSeq(tios);
    if actor != ep {return;}
    var s, s' := ls.servers[actor].follower, ls'.servers[actor].follower;
    var st, st' := tls.t_servers[actor], tls'.t_servers[actor];
    var pkt := tios[0].r;
    
    // s can be in two possible states. 1) F_PRESYNC, 2) F_SYNC
    if s.state == F_PRESYNC {
        assert pkt.msg.v.SyncDIFF?;
        assert st'.dts <= Follower_F_SYNC_dts_Formula(f, st');
        assert st.ts <= LeaderInfo_Message_PreQuorum_ts_Formula(f, s.serialLI) + ProcLI;
        assert pkt.msg.ts <= Sync_Message_ts_Formula(f, pkt.msg.v.serial);
        assert s'.state == F_SYNC;
        var handlerStartTime := TimeMax(st.ts, pkt.msg.ts);
        lemma_Math_Inequalities_CommonMult(ProcFI, s.serialLI+1, f);
        
        assert LeaderInfo_Message_PreQuorum_ts_Formula(f, s.serialLI) + ProcLI
                <= Sync_Message_ts_Formula(f, pkt.msg.v.serial);
        lemma_Math_MaxOfInequalities(st.ts, pkt.msg.ts, 
        LeaderInfo_Message_PreQuorum_ts_Formula(f, s.serialLI) + ProcLI, 
        Sync_Message_ts_Formula(f, pkt.msg.v.serial));
        assert st'.ts <= Follower_F_SYNC_ts_Formula(f, st');
    } else {
        assert s.state == F_SYNC;
        assert st'.dts <= Follower_F_SYNC_dts_Formula(f, st');
        assert pkt.msg.v.NewLeader?;
        lemma_Math_Inequalities_CommonMult(Sync, s.serialSync, f);
        lemma_Math_Inequalities_CommonMult(Sync, s.serialSync + 1, f);
        lemma_Math_Inequalities_CommonMult(Sync, s.serialSync, f);
        lemma_Math_Inequalities_CommonMult(Sync, s.serialSync + 1, f);
        lemma_Math_Inequalities_CommonMult(Sync, pkt.msg.v.serial + 1, f);
        var handlerStartTime := TimeMax(st.ts, pkt.msg.ts);
        lemma_Math_MaxOfInequalities(pkt.msg.ts, st.ts, 
        NewLeader_Message_ts_Formula(f, pkt.msg.v.serial), 
        AckEpoch_Message_PreQuorum_ts_Formula(f, f-1)
            + ProcEpAck * f
            + ProcEpAck * f
            + PreSync * f    
            + Sync * f  
            + Sync * f  
            + D
            + ProcSyncI);      
        assert st'.ts <= Follower_F_SYNC_ts_Formula(f, st');
    }
}


lemma Sync_Leader_PreQuorum_Invariant_Helper(config:Config, tls:TLS_State, tls':TLS_State) 
    requires TLS_Next(tls, tls')
    requires General_LS_Performance_Assumption(tls)
    requires Basic_Invariants(config, tls) && Basic_Invariants(config, tls')
    requires EmptyDiff_Invariant(tls) && EmptyDiff_Invariant(tls')
    requires FollowerInfo_Message_ts_Invariant(tls) && FollowerInfo_Message_ts_Invariant(tls'); 
    requires ProcessFI_PreQuorum_Invariant(tls) && ProcessFI_PreQuorum_Invariant(tls'); 
    requires Handshake_Messages_Invariant(tls) && Handshake_Messages_Invariant(tls')
    requires Handshake_Leader_PreQuorum_Invariant(tls) && Handshake_Leader_PreQuorum_Invariant(tls')
    requires Sync_Messages_Invariant(tls) && Sync_Messages_Invariant(tls')
    requires Sync_Follower_Invariant(tls) && Sync_Follower_Invariant(tls')
    // Induction hypothesis
    requires Sync_Leader_PreQuorum_Invariant(tls)
    // Antecedent
    requires !IsQuorum(|tls'.config|, tls'.t_servers[config[0]].v.leader.globals.ackSet)
    // Conclusion
    ensures tls'.t_servers[config[0]].dts <= ProcessAck_PreQuorum_dts_Formula(tls'.f, |tls'.config|, tls'.t_servers[config[0]])
    ensures tls'.t_servers[config[0]].ts <= ProcessAck_PreQuorum_ts_Formula(tls'.f, |tls'.config|, tls'.t_servers[config[0]])
{
    var f, n := tls'.f, |tls'.config|;
    var actor, tios:seq<TZKIo> :| actor in tls.t_servers && TLS_NextOneServer(tls, tls', actor, tios);
    var lt, lt' := tls.t_servers[config[0]], tls'.t_servers[config[0]];
    var l, l', ios := lt.v.leader, lt'.v.leader, UntagLIoOpSeq(tios);
    if actor == config[0] {
        if l.state == L_RUNNING {
            assert LeaderStutter(l, l', ios);
            assert lt'.ts == lt.ts && lt'.dts == lt.dts;
        } else {
            assert LeaderStartStep(l, l', ios);
            if StepSingleHandler_NoRcv(l, l', ios) {
                Sync_Leader_PreQuorum_Invariant_Helper_NoRcv(config, tls, tls', config[0], tios);
            } else {
                Sync_Leader_PreQuorum_Invariant_Helper_Rcv(config, tls, tls', config[0], tios);
            }
        }
    }
}


lemma Sync_Leader_PreQuorum_Invariant_Helper_NoRcv(config:Config, tls:TLS_State, tls':TLS_State, actor:EndPoint, tios:seq<TZKIo>) 
    requires TLS_Next(tls, tls')
    requires General_LS_Performance_Assumption(tls)
    requires Basic_Invariants(config, tls) && Basic_Invariants(config, tls')
    requires EmptyDiff_Invariant(tls) && EmptyDiff_Invariant(tls')
    requires FollowerInfo_Message_ts_Invariant(tls) && FollowerInfo_Message_ts_Invariant(tls'); 
    requires ProcessFI_PreQuorum_Invariant(tls) && ProcessFI_PreQuorum_Invariant(tls'); 
    requires Handshake_Messages_Invariant(tls) && Handshake_Messages_Invariant(tls')
    requires Handshake_Leader_PreQuorum_Invariant(tls) && Handshake_Leader_PreQuorum_Invariant(tls')
    requires Sync_Messages_Invariant(tls) && Sync_Messages_Invariant(tls')
    requires Sync_Follower_Invariant(tls) && Sync_Follower_Invariant(tls')
    // Induction hypothesis
    requires Sync_Leader_PreQuorum_Invariant(tls)
    // Antecedent
    requires !IsQuorum(|tls'.config|, tls'.t_servers[config[0]].v.leader.globals.ackSet)
    // Local predicates
    requires actor in tls.t_servers && TLS_NextOneServer(tls, tls', actor, tios);
    requires actor == config[0]
    requires tls.t_servers[config[0]].v.leader.state == L_STARTING
    requires LeaderStartStep(tls.t_servers[config[0]].v.leader, tls'.t_servers[config[0]].v.leader, UntagLIoOpSeq(tios))
    requires StepSingleHandler_NoRcv(tls.t_servers[actor].v.leader, tls'.t_servers[actor].v.leader, UntagLIoOpSeq(tios))
    // Conclusion
    ensures tls'.t_servers[config[0]].dts <= ProcessAck_PreQuorum_dts_Formula(tls'.f, |tls'.config|, tls'.t_servers[config[0]])
    ensures tls'.t_servers[config[0]].ts <= ProcessAck_PreQuorum_ts_Formula(tls'.f, |tls'.config|, tls'.t_servers[config[0]])
{
    var f, n := tls'.f, |tls'.config|;
    var lt, lt' := tls.t_servers[config[0]], tls'.t_servers[config[0]];
    var l, l', ios := lt.v.leader, lt'.v.leader, UntagLIoOpSeq(tios);
    lemma_Math_Mult_b();
    var fid :| LHNext(l, l', fid, ios);
    var h , h', g, g' := l.handlers[fid], l'.handlers[fid], l.globals, l'.globals;
    match h.state 
    case LH_HANDSHAKE_A => {
        assert IsVerifiedQuorum(h.follower_id, n, g.connectingFollowers);
        assert ios[0].s.msg.LeaderInfo?;
        var pkt := tios[0].s;
        assert pkt.msg.ts <= SendFI + D 
                            + ProcFI * f
                            + ProcFI * (pkt.msg.v.serial + 1)
                            + D;
        assert pkt.msg.ts == lt'.ts + D;
        lemma_Math_Addition();
        lemma_Math_Inequalities_CommonMult(ProcFI, pkt.msg.v.serial + 1, f);
        assert lt'.ts <= SendFI + D 
                        + ProcFI * f
                        + ProcFI * f
                      <= AckEpoch_Message_PreQuorum_ts_Formula(f, f-1);
        var a := AckEpoch_Message_PreQuorum_ts_Formula(f, f-1);
        lemma_Math_Inequalities_Mult();
        if |g'.ackSet| <= 1 {
            var b := ProcEpAck * (|g'.electingFollowers|-1)
                + ProcEpAck * g'.procEpCount 
                + PreSync * g'.prepCount    
                + Sync * g'.nextSerialSync  
                + Sync * g'.nextSerialNL;
            lemma_Math_Addition_2(a, b);
            assert ProcessAck_PreQuorum_ts_Formula(f, n, lt') == a + b;
            assert lt'.ts <= ProcessAck_PreQuorum_ts_Formula(f, n, lt');  
        } else {
            assert lt'.ts <= ProcessAck_PreQuorum_ts_Formula(f, n, lt');  
        }
    }
    case LH_HANDSHAKE_B => {
        assert IsVerifiedQuorum(h.follower_id, |g.config|, g.electingFollowers);
        assert |g'.ackSet| == |g.ackSet| == 1;
    }
    case LH_PREP_SYNC => {
        assert lt'.ts == lt.ts + PreSync;
        assert lt'.ts <= ProcessAck_PreQuorum_ts_Formula(f, n, lt) + PreSync;
        assert g'.prepCount == g.prepCount + 1;
        assert ProcessAck_PreQuorum_ts_Formula(f, n, lt) + PreSync == ProcessAck_PreQuorum_ts_Formula(f, n, lt');
    }
    case LH_SYNC => {
        assert ZooKeeper_LearnerHandler.DoSync(h, h', g, g', ios);
    }
    case LH_PROCESS_ACK => {
        assert ZooKeeper_LearnerHandler.ProcessAck(h, h', g, g', ios);
        if g.zkdb.isRunning {
            assert !StepSingleHandler(l, l', ios);
            assert false;
        } else {
            assert ios[0].LIoOpReceive?;
            assert false;
        }
    }
    case LH_RUNNING => {assert lt'.dts == lt.dts && lt'.ts == lt.ts;}
    case LH_ERROR => {assert lt'.dts == lt.dts && lt'.ts == lt.ts;}  
    // Conclusion 
    assert lt'.dts <= ProcessAck_PreQuorum_dts_Formula(f, n, lt');
    assert lt'.ts <= ProcessAck_PreQuorum_ts_Formula(f, n, lt');              
}


lemma Sync_Leader_PreQuorum_Invariant_Helper_Rcv(config:Config, tls:TLS_State, tls':TLS_State, actor:EndPoint, tios:seq<TZKIo>) 
    requires TLS_Next(tls, tls')
    requires General_LS_Performance_Assumption(tls)
    requires Basic_Invariants(config, tls) && Basic_Invariants(config, tls')
    requires EmptyDiff_Invariant(tls) && EmptyDiff_Invariant(tls')
    requires FollowerInfo_Message_ts_Invariant(tls) && FollowerInfo_Message_ts_Invariant(tls'); 
    requires ProcessFI_PreQuorum_Invariant(tls) && ProcessFI_PreQuorum_Invariant(tls'); 
    requires Handshake_Messages_Invariant(tls) && Handshake_Messages_Invariant(tls')
    requires Handshake_Leader_PreQuorum_Invariant(tls) && Handshake_Leader_PreQuorum_Invariant(tls')
    requires Sync_Messages_Invariant(tls) && Sync_Messages_Invariant(tls')
    requires Sync_Follower_Invariant(tls) && Sync_Follower_Invariant(tls')
    // Induction hypothesis
    requires Sync_Leader_PreQuorum_Invariant(tls)
    // Antecedent
    requires !IsQuorum(|tls'.config|, tls'.t_servers[config[0]].v.leader.globals.ackSet)
    // Local predicates
    requires actor in tls.t_servers && TLS_NextOneServer(tls, tls', actor, tios);
    requires actor == config[0]
    requires tls.t_servers[config[0]].v.leader.state == L_STARTING
    requires LeaderStartStep(tls.t_servers[config[0]].v.leader, tls'.t_servers[config[0]].v.leader, UntagLIoOpSeq(tios))
    requires StepSingleHandler_Rcv(tls.t_servers[actor].v.leader, tls'.t_servers[actor].v.leader, UntagLIoOpSeq(tios))
    // Conclusion
    ensures tls'.t_servers[config[0]].dts <= ProcessAck_PreQuorum_dts_Formula(tls'.f, |tls'.config|, tls'.t_servers[config[0]])
    ensures tls'.t_servers[config[0]].ts <= ProcessAck_PreQuorum_ts_Formula(tls'.f, |tls'.config|, tls'.t_servers[config[0]])
{
    var f, n := tls'.f, |tls'.config|;
    var lt, lt' := tls.t_servers[config[0]], tls'.t_servers[config[0]];
    var l, l', ios := lt.v.leader, lt'.v.leader, UntagLIoOpSeq(tios);
    lemma_Math_Mult_b();
    var fid := ios[0].r.sender_index;
    var h , h', g, g' := l.handlers[fid], l'.handlers[fid], l.globals, l'.globals;
    match h.state 
    case LH_HANDSHAKE_A => {
        if IsQuorum(n, l.globals.electingFollowers){
            assert IsQuorum(n, l.globals.connectingFollowers);
            assert false;
        } else {
            if !IsQuorum(n, l.globals.connectingFollowers) {
                assert lt'.dts <= ProcessAck_PreQuorum_dts_Formula(f, n, lt');
                assert lt.ts <= SendFI + D + ProcFI * (|l.globals.connectingFollowers|-1);
                var pkt := tios[0].r;
                assert pkt.msg.ts <= SendFI + D;
                lemma_Math_Addition_2(SendFI + D, ProcFI * (|l.globals.connectingFollowers|-1));
                lemma_Math_MaxOfInequalities(pkt.msg.ts, lt.ts, SendFI + D, SendFI + D + ProcFI * (|l.globals.connectingFollowers|-1));
                assert TimeMax(lt.ts, pkt.msg.ts) <= SendFI + D + ProcFI * (|l.globals.connectingFollowers|-1);
                assert lt'.ts <= SendFI + D + ProcFI * (|l.globals.connectingFollowers|-1) + ProcFI;
                lemma_Math_Mult_b();
                assert ProcFI * (|l.globals.connectingFollowers|-1) + ProcFI == ProcFI * (|l'.globals.connectingFollowers|-1);
                lemma_Math_Inequalities_CommonMult(ProcFI, |l'.globals.connectingFollowers|-1, f);
                assert lt'.ts <= SendFI + D + ProcFI * f;
                var a := SendFI + D + ProcFI * f;
                var b := ProcFI * f
                        + D
                        + ProcLI + D
                        + ProcEpAck * (|g'.electingFollowers|-1)
                        + ProcEpAck * g'.procEpCount 
                        + PreSync * g'.prepCount   
                        + Sync * g'.nextSerialSync 
                        + Sync * g'.nextSerialNL;
                lemma_Math_Inequalities_Mult();
                assert ProcessAck_PreQuorum_ts_Formula(f, n, lt') == a + b;
                lemma_Math_Addition_2(a, b);
                assert a <= a+b;
                assert lt'.ts <= ProcessAck_PreQuorum_ts_Formula(f, n, lt');
            } else {
                assert lt'.dts <= ProcessAck_PreQuorum_dts_Formula(f, n, lt');
                assert lt'.ts <= ProcessAck_PreQuorum_ts_Formula(f, n, lt');
            }
        }
        // Conclusion
        assert lt'.dts <= ProcessAck_PreQuorum_dts_Formula(f, n, lt');
        assert lt'.ts <= ProcessAck_PreQuorum_ts_Formula(f, n, lt');
    }
    case LH_HANDSHAKE_B => {
        if IsVerifiedQuorum(h.follower_id, |g.config|, g.electingFollowers) {
            assert false;
        } else {
            assert !IsQuorum(n, g.electingFollowers);
            if !IsQuorum(n, g'.electingFollowers){
                assert lt'.dts <= ProcessEpAck_PreQuorum_dts_Formula(f, n, lt');
                assert lt'.ts <= ProcessEpAck_PreQuorum_ts_Formula(f, n, lt');
                assert |g'.ackSet| <= 1;
                assert |g'.electingFollowers| > 1;
                lemma_Math_Addition();
                lemma_Math_Inequalities_Mult();
                var a:Timestamp := AckEpoch_Message_PreQuorum_ts_Formula(f, f-1)
                        + ProcEpAck * (|g'.electingFollowers|-1);
                var b:Timestamp := ProcEpAck * g'.procEpCount 
                        + PreSync * g'.prepCount    
                        + Sync * g'.nextSerialSync 
                        + Sync * g'.nextSerialNL;
                assert ProcessEpAck_PreQuorum_ts_Formula(f, n, lt') == a;
                assert ProcessAck_PreQuorum_ts_Formula(f, n, lt') == a + b;
                lemma_Math_Addition_2(a, b);
                assert a <= a + b;
                assert ProcessEpAck_PreQuorum_ts_Formula(f, n, lt')
                    <= ProcessAck_PreQuorum_ts_Formula(f, n, lt');
            } else {
                var pkt := tios[0].r;
                assert pkt.msg.v.AckEpoch?;
                lemma_Math_Inequalities_CommonMult(ProcFI, pkt.msg.v.serial+1, f);
                assert g'.ackSet == g.ackSet;
                assert |g'.ackSet| <= 1;
                assert |g'.electingFollowers| == |g.electingFollowers| + 1;
                var lt_max_ts := AckEpoch_Message_PreQuorum_ts_Formula(f, f-1)
                                + ProcEpAck * (|g.electingFollowers|-1);
                var pkt_max_ts := AckEpoch_Message_PreQuorum_ts_Formula(f, pkt.msg.v.serial);
                lemma_Math_Inequalities_Mult();
                lemma_Math_Addition();
                lemma_Math_MaxOfInequalities(pkt.msg.ts, lt.ts, pkt_max_ts, lt_max_ts);
                assert lt'.ts <= lt_max_ts + ProcEpAck;
                var lt'_max_ts := AckEpoch_Message_PreQuorum_ts_Formula(f, f-1)
                                + ProcEpAck * (|g'.electingFollowers|-1);
                assert lt'.ts <= lt'_max_ts;
                lemma_Math_Inequalities_Mult();
                var b:Timestamp := ProcEpAck * g'.procEpCount 
                        + PreSync * g'.prepCount    
                        + Sync * g'.nextSerialSync 
                        + Sync * g'.nextSerialNL;
                lemma_Math_Addition_2(lt'_max_ts, b);
                assert lt'_max_ts <= lt'_max_ts + b;
            }
        }
    }
    case LH_PREP_SYNC => assert false;  // No receives in this state
    case LH_SYNC => assert false;       // No receives in this state
    case LH_PROCESS_ACK => {
        if !g.zkdb.isRunning {
            var pkt := tios[0].r;
            assert pkt.msg.v.Ack?;
            assert pkt.msg.ts <= Ack_Message_ts_Formula(tls.f, pkt.msg.v.serial);
            assert |g'.ackSet| > 1;
            assert lt'.dts <= ProcessAck_PreQuorum_dts_Formula(f, n, lt');
            if |g.ackSet| == 1 {
                lemma_Math_Inequalities_Mult();
                lemma_Math_Inequalities_CommonMult(ProcEpAck, |g.electingFollowers|-1, f);
                lemma_Math_Inequalities_CommonMult(ProcEpAck, g.procEpCount, f);
                lemma_Math_Inequalities_CommonMult(PreSync, g.prepCount, f);
                lemma_Math_Inequalities_CommonMult(Sync, g.nextSerialSync, f);
                lemma_Math_Inequalities_CommonMult(Sync, g.nextSerialNL, f);
                assert lt'.ts <= ProcessAck_PreQuorum_ts_Formula(f, n, lt');
            } 
        }
    }
    case LH_RUNNING => {assert lt'.dts == lt.dts && lt'.ts == lt.ts;}
    case LH_ERROR => {assert lt'.dts == lt.dts && lt'.ts == lt.ts;}                           
}


}
