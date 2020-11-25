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


predicate FollowerInit_Invariant(tls:TLS_State){
    forall ep | ep in tls.t_servers :: (
        && tls.t_servers[ep].v.FollowerPeer? 
        && tls.t_servers[ep].v.follower.state == F_HANDSHAKE_A
        ==> 
        && tls.t_servers[ep].ts == TimeZero()
    )
}

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


predicate FollowerInfo_Message_Invariant(tls:TLS_State)
{
    var perf_formula := SendFI;
    forall pkt | pkt in tls.t_environment.sentPackets && pkt.msg.v.FollowerInfo?
    :: pkt.msg.ts == perf_formula
}


lemma lemma_FollowerInfo_Message_Invariant_Induction(config:Config, tls:TLS_State, tls':TLS_State)
    requires TLS_Next(tls, tls')
    requires Basic_Invariants(config, tls) && Basic_Invariants(config, tls')
    requires General_LS_Performance_Assumption(tls)
    // Induction hypothesis
    requires FollowerInfo_Message_Invariant(tls)
    ensures FollowerInfo_Message_Invariant(tls')
{
    var actor, tios :| actor in tls.t_servers && TLS_NextOneServer(tls, tls', actor, tios);
    var ls, ls', ios := UntagLS_State(tls), UntagLS_State(tls'), UntagLIoOpSeq(tios);
    assert LS_NextOneServer(ls, ls', actor, ios);
    if ls.servers[actor].FollowerPeer? {
        // TODO
        assume FollowerInfo_Message_Invariant(tls');
    } else {
        /* No new FollowerInfo messages get sent, so nothing to worry */
        var s, s' := ls.servers[actor].leader, ls'.servers[actor].leader;
        assert LeaderNext(s, s', ios);
        if s.state == L_STARTING {
            if |ios| > 0 {
                var h, h' := s.handlers[s.nextHandlerToStep], s'.handlers[s.nextHandlerToStep];
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
        assert FollowerInfo_Message_Invariant(tls');  
    }
}

}
