include "../../../Impl/Common/SeqIsUniqueDef.i.dfy"
include "../../../Common/Framework/EnvironmentTCP.s.dfy"
include "../Timestamps/TimestampedType.dfy"

include "../Types.dfy"
include "../DistributedSystem.dfy"
include "../ZKEnvironment.dfy"
include "../ZKDatabase.dfy"
include "../Follower.dfy"
include "../Leader.dfy"
include "../LearnerHandler.dfy"
include "Definitions.dfy"


module ZooKeeper_TimestampedDistributedSystem {
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

type TQuorumPeer = TimestampedType<QuorumPeer>

type TZKEnvironment = TimestampedLEnvironment<EndPoint, ZKMessage>

type TZKIo = TimestampedLIoOp<EndPoint, ZKMessage>

datatype TLS_State = TLS_State(
    config:Config,
    t_environment: TZKEnvironment,
    initialZkdbState: seq<ZKDatabase>,
    t_servers: map<EndPoint, TQuorumPeer>
)



predicate TLS_Init(config:Config, tls:TLS_State, f: int) {
    && LS_Init(config, UntagLS_State(tls), f)
    && tls.config == config
    && LEnvironment_Init(config, tls.t_environment)
    && forall id | id in tls.t_servers :: tls.t_servers[id].ts == TimeZero()
}


predicate TLS_NextOneServer(tls:TLS_State, tls':TLS_State, id:EndPoint, ios:seq<TZKIo>)
        requires id in tls.t_servers;
{
    && LS_NextOneServer(UntagLS_State(tls), UntagLS_State(tls'), id, UntagLIoOpSeq(ios))
    && tls.t_environment.nextStep == LEnvStepHostIos(id, ios)
    && (forall t_io | t_io in ios && t_io.LIoOpSend? :: t_io.s.msg.ts == tls'.t_servers[id].ts)
    && tls'.t_servers == tls.t_servers[id := tls'.t_servers[id]]
    && var hs := ActionToHostStep(tls, tls', id, ios);

    && (if |ios| > 0 && ios[0].LIoOpReceive? then   // Note that in performal, at most one rcv in each step
            tls'.t_servers[id].ts == TLS_RecvPerfUpdate(tls.t_servers[id].ts, ios[0].r.msg.ts, hs)
        else
            tls'.t_servers[id].ts == TLS_NoRecvPerfUpdate(tls.t_servers[id].ts, hs)
    )
}


predicate TLS_Next(tls:TLS_State, tls':TLS_State){
        && LS_Next(UntagLS_State(tls), UntagLS_State(tls'))
        && tls'.config == tls.config
        && LEnvironment_Next(tls.t_environment, tls'.t_environment)
        && (exists ep, ios :: ep in tls.t_servers && TLS_NextOneServer(tls, tls', ep, ios))
}


predicate ValidTLSBehavior(config:Config, tlb:seq<TLS_State>, f: int)  {
    && |tlb| > 0
    && TLS_Init(config, tlb[0], f)
    && (forall i :: 0 <= i < |tlb| -1 ==> TLS_Next(tlb[i], tlb[i+1]))
}



/*****************************************************************************************
*                                       Hoststep                                         *
*****************************************************************************************/

function ActionToHostStep(tls:TLS_State, tls':TLS_State, id:EndPoint, ios:seq<TZKIo>) : HostStep
    requires id in tls.t_servers    
    requires LS_NextOneServer(UntagLS_State(tls), UntagLS_State(tls'), id, UntagLIoOpSeq(ios))
{
    var ls, ls', zkios := UntagLS_State(tls), UntagLS_State(tls'), UntagLIoOpSeq(ios);
    match ls.servers[id] 
    case LeaderPeer(leader) => (
        match leader.state
        case L_RUNNING => L(LStutter)
        case L_STARTING => (
            var s, s' := ls.servers[id].leader, ls'.servers[id].leader;
            if IsVerifiedQuorum(s.my_id, |s.config|, s.globals.ackSet) 
            then L(LStutter)
            else (
                var i := s.nextHandlerToStep;
                assert StepSingleHandler(s, s', zkios);
                assert LearnerHandlerNext(s.handlers[i], s'.handlers[i], s.globals, s'.globals, zkios);
                match s.handlers[i].state
                case LH_HANDSHAKE_A => L(ProcessFollowerInfo)
                case LH_HANDSHAKE_B => L(ProcessEpochAck)
                case LH_PREP_SYNC => L(PrepSync)
                case LH_SYNC => (
                    if zkios[0].s.msg.SyncSNAP?
                    then L(DoSyncSNAP)  // Sending a snapshot is a special event
                    else L(DoSync)
                )
                case LH_PROCESS_ACK => L(ProcessAck)
                case LH_RUNNING => L(LStutter)
                case LH_ERROR => L(LStutter)
            )   
        )
    )
    case FollowerPeer(follower) => (
        match follower.state 
        case F_HANDSHAKE_A => F(SendFollowerInfo)
        case F_HANDSHAKE_B => F(ProcessLeaderInfo)
        case F_SYNC => (
            var s, s' := ls.servers[id].follower, ls'.servers[id].follower;
            assert SyncWithLeader(s, s', zkios);
            if zkios[0].r.msg.SyncSNAP? 
                then F(ProcessSnap)   // Processing a snapshot is a special event
            else if zkios[0].r.msg.FollowerInfo? || zkios[0].r.msg.LeaderInfo? || zkios[0].r.msg.AckEpoch? || zkios[0].r.msg.Ack?
                then F(FStutter)
            else
                F(ProcessSync)
        )
        case F_RUNNING => F(FStutter)
        case F_ERROR => F(FStutter)
    )
}



/*****************************************************************************************
*                                       Utilities                                        *
*****************************************************************************************/

function UntagLSServers(t_servers: map<EndPoint, TQuorumPeer>) : map<EndPoint, QuorumPeer>
    ensures forall id :: id in t_servers <==> id in UntagLSServers(t_servers);
    ensures forall id | id in t_servers :: UntagLSServers(t_servers)[id] == t_servers[id].v;
{
    map id | id in t_servers :: t_servers[id].v
}

function UntagLS_State(tds:TLS_State) : LS_State
    ensures UntagLS_State(tds).servers == UntagLSServers(tds.t_servers);
    ensures UntagLS_State(tds).environment == UntagLEnvironment(tds.t_environment);
{
    LS_State(
        UntagLEnvironment(tds.t_environment),
        tds.initialZkdbState,
        UntagLSServers(tds.t_servers))
}
}