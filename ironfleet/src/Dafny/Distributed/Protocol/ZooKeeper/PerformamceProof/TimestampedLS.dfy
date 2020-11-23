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
import opened PerformanceProof__Definitions_i

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
    && LEnvironment_Init(tls.t_environment, config)
    && forall id | id in tls.t_servers :: tls.t_servers[id].ts == TimeZero()
}


predicate TLS_NextOneServer(tls:TLS_State, tls':TLS_State, id:EndPoint, ios:seq<TZKIo>)
        requires id in tls.t_servers;
{
    && LS_NextOneServer(UntagLS_State(tls), UntagLS_State(tls'), id, UntagLIoOpSeq(ios))
    && (forall t_io | t_io in ios && t_io.LIoOpSend? :: t_io.s.msg.ts == tls'.t_servers[id].ts)
    && tls'.t_servers == tls.t_servers[id := tls'.t_servers[id]]
    && var hs := ActionToHostStep(tls, tls', id, ios);

    && (if |ios| > 0 && ios[0].LIoOpReceive? then
            tls'.t_servers[id].ts == TLS_RecvPerfUpdate(tls.t_servers[id].ts, ios[0].r.msg.ts, hs)
        else
            tls'.t_servers[id].ts == TLS_NoRecvPerfUpdate(tls.t_servers[id].ts, hs)
    )
}


predicate TLS_Next(tls:TLS_State, tls':TLS_State){
        LEnvironment_Next(tls.t_environment, tls'.t_environment)
    && if tls.t_environment.nextStep.LEnvStepHostIos? && tls.t_environment.nextStep.actor in tls.t_servers then
            TLS_NextOneServer(tls, tls', tls.t_environment.nextStep.actor, tls.t_environment.nextStep.ios)
        else
            tls'.t_servers == tls.t_servers
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
    // TODO
    L(GetEpoch)
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