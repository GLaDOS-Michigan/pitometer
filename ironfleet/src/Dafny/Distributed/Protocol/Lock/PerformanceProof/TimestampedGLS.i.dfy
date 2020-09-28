include "../Node.i.dfy"
include "../RefinementProof/DistributedSystem.i.dfy"
include "../../../Services/Lock/LockTimestampedDistributedSystem.i.dfy"
include "Definitions.i.dfy"

module TimestampedGLS_i {
import opened Protocol_Node_i
import opened DistributedSystem_i
import opened LockTimestampedDistributedSystem_i
import opened PerformanceProof__Definitions_i

type TimestampedNode = TimestampedType<Node>

    
/*****************************************************************************************
/                                    LS_State + Timestaps                                *
*****************************************************************************************/

datatype TimestampedLS_State = TimestampedLS_State(
    config:ConcreteConfiguration,
    t_environment:TimestampedLEnvironment<EndPoint, LockMessage, HostStep>,
    t_servers:map<EndPoint,TimestampedNode>
    )

function UntagLSServers(t_servers: map<EndPoint, TimestampedNode>) : map<EndPoint, Node>
    ensures forall id :: id in t_servers <==> id in UntagLSServers(t_servers);
    ensures forall id | id in t_servers :: UntagLSServers(t_servers)[id] == t_servers[id].v;
{
    map id | id in t_servers :: t_servers[id].v
}

function UntagLS_State(tds:TimestampedLS_State) : LS_State
    ensures UntagLS_State(tds).servers == UntagLSServers(tds.t_servers);
    ensures UntagLS_State(tds).environment == UntagLEnvironment(tds.t_environment);
{
    LS_State(
        UntagLEnvironment(tds.t_environment),
        UntagLSServers(tds.t_servers))
}

predicate TLS_Init(tls: TimestampedLS_State, config:Config)
    reads *
{
    && LS_Init(UntagLS_State(tls), config)
        && tls.config == config
        && LEnvironment_Init(tls.t_environment)
        && tls.t_servers[config[0]].ts == TimeZero()
        && forall id :: id in tls.t_servers && id != config[0] ==> tls.t_servers[id].ts == TimeZero()
}

predicate TLS_NextOneServer(tls: TimestampedLS_State, tls': TimestampedLS_State, id:EndPoint, ios:seq<TimestampedLIoOp<EndPoint, LockMessage>>, hstep:HostStep)
    requires id in tls.t_servers;
    reads *
{
    LS_NextOneServer(UntagLS_State(tls), UntagLS_State(tls'), id, UntagLIoOpSeq(ios), hstep)
        && (if |ios| > 0 && ios[0].LIoOpReceive? then
            tls'.t_servers[id].ts == TLS_RecvPerfUpdate(tls.t_servers[id].ts, ios[0].r.msg.ts, hstep)
        else
            tls'.t_servers[id].ts == TLS_NoRecvPerfUpdate(tls.t_servers[id].ts, hstep)
            )

        && (forall t_io :: t_io in ios && t_io.LIoOpSend? ==> t_io.s.msg.ts == tls'.t_servers[id].ts)
        && tls'.t_servers == tls.t_servers[id := tls'.t_servers[id]]
}

predicate TLS_Next(tls:TimestampedLS_State, tls': TimestampedLS_State)
    reads *
{
    tls.config == tls'.config
    && LS_Next(UntagLS_State(tls), UntagLS_State(tls'))
    && LEnvironment_Next(tls.t_environment, tls'.t_environment)
    && if tls.t_environment.nextStep.LEnvStepHostIos? && tls.t_environment.nextStep.actor in tls.t_servers then
                    TLS_NextOneServer(tls, tls', tls.t_environment.nextStep.actor, tls.t_environment.nextStep.ios, tls.t_environment.nextStep.nodeStep)
            else
            tls'.t_servers == tls.t_servers

            && (if tls.t_environment.nextStep.LEnvStepHostIos? then
                // Any irrelevant packets are given a default Zero.
                    && (forall t_io :: t_io in tls.t_environment.nextStep.ios && t_io.LIoOpSend? ==> t_io.s.msg.ts == TimeZero())
                    else
                    true)
    }


/*****************************************************************************************
/                                GLS_State + Timestaps                                   *
*****************************************************************************************/

datatype TimestampedGLS_State = TimestampedGLS_State(
    tls:TimestampedLS_State,
    history:seq<EndPoint>
    )


function UntagGLS_State(tgls:TimestampedGLS_State) : GLS_State
    ensures UntagGLS_State(tgls).ls == UntagLS_State(tgls.tls);
    ensures UntagGLS_State(tgls).history == tgls.history;
{
    GLS_State(
        UntagLS_State(tgls.tls),
        tgls.history
    )
}

predicate TGLS_Init(tgls:TimestampedGLS_State, config:Config)
    reads *
{
    && TLS_Init(tgls.tls, config)
    && tgls.history == [config[0]]
}

predicate TGLS_Next(tgls:TimestampedGLS_State, tgls':TimestampedGLS_State)
    reads *
{
    && TLS_Next(tgls.tls, tgls'.tls)
    && (if  && tgls.tls.t_environment.nextStep.LEnvStepHostIos? 
            && tgls.tls.t_environment.nextStep.actor in tgls.tls.t_servers
            && NodeGrant(tgls.tls.t_servers[tgls.tls.t_environment.nextStep.actor].v, tgls'.tls.t_servers[tgls.tls.t_environment.nextStep.actor].v, UntagLIoOpSeq(tgls.tls.t_environment.nextStep.ios))
            && tgls.tls.t_servers[tgls.tls.t_environment.nextStep.actor].v.held && tgls.tls.t_servers[tgls.tls.t_environment.nextStep.actor].v.epoch < 0xFFFF_FFFF_FFFF_FFFF 
        then
            tgls'.history == tgls.history + [tgls.tls.t_servers[tgls.tls.t_environment.nextStep.actor].v.config[(tgls.tls.t_servers[tgls.tls.t_environment.nextStep.actor].v.my_index + 1) % |tgls.tls.t_servers[tgls.tls.t_environment.nextStep.actor].v.config|]]
        else
            tgls'.history == tgls.history
        )
}

predicate ValidTimestampedGLSBehavior(tglb:seq<TimestampedGLS_State>, config:Config)
    reads *
{
    && |tglb| > 0
    && TGLS_Init(tglb[0], config)
    && (forall i :: 0 < i < |tglb| ==> TGLS_Next(tglb[i - 1], tglb[i]))
}
}
