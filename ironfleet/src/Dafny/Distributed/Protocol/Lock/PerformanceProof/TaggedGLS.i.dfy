include "../Node.i.dfy"
  include "../RefinementProof/DistributedSystem.i.dfy"
include "../../../Services/Lock/LockTaggedDistributedSystem.i.dfy"
include "Definitions.i.dfy"

module TaggedGLS_i {
import opened Protocol_Node_i
import opened DistributedSystem_i
import opened LockTaggedDistributedSystem_i
import opened PerformanceProof__Definitions_i

type TaggedNode = TaggedType<Node>
  
datatype TaggedLS_State = TaggedLS_State(
  config:ConcreteConfiguration,
  t_environment:TaggedLEnvironment<EndPoint, LockMessage, HostStep>,
  t_servers:map<EndPoint,TaggedNode>
  )

datatype TaggedGLS_State = TaggedGLS_State(
  tls:TaggedLS_State,
  history:seq<EndPoint>
  )

  function UntagLSServers(t_servers: map<EndPoint, TaggedNode>) : map<EndPoint, Node>
  {
    map id | id in t_servers :: t_servers[id].v
  }

  function UntagLS_State(tds:TaggedLS_State) : LS_State
  {
    LS_State(
      UntagLEnvironment(tds.t_environment),
      UntagLSServers(tds.t_servers))
  }

  predicate TLS_Init(tls: TaggedLS_State, config:Config)
    reads *
  {
    && LS_Init(UntagLS_State(tls), config)
      && tls.config == config
      && LEnvironment_Init(tls.t_environment)
      && tls.t_servers[config[0]].pr == PerfZero()
      && forall id :: id in tls.t_servers && id != config[0] ==> tls.t_servers[id].pr == PerfVoid()
  }

  predicate TLS_NextOneServer(tls: TaggedLS_State, tls': TaggedLS_State, id:EndPoint, ios:seq<TaggedLIoOp<EndPoint, LockMessage>>, hstep:HostStep)
    requires id in tls.t_servers;
    reads *
  {
    LS_NextOneServer(UntagLS_State(tls), UntagLS_State(tls'), id, UntagLIoOpSeq(ios), hstep)
      && (if |ios| > 0 && ios[0].LIoOpReceive? then
        tls'.t_servers[id].pr == TLS_RecvPerfUpdate(tls.t_servers[id].pr, ios[0].r.msg.pr, hstep)
      else
        tls'.t_servers[id].pr == TLS_NoRecvPerfUpdate(tls.t_servers[id].pr, hstep)
        )

      && (forall t_io :: t_io in ios && t_io.LIoOpSend? ==> t_io.s.msg.pr == tls'.t_servers[id].pr)
      && tls'.t_servers == tls.t_servers[id := tls'.t_servers[id]]
  }

  predicate TLS_Next(tls:TaggedLS_State, tls': TaggedLS_State)
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
            && (forall t_io :: t_io in tls.t_environment.nextStep.ios && t_io.LIoOpSend? ==> t_io.s.msg.pr == PerfZero())
            else
            true)
  }

  predicate TGLS_Init(tgls:TaggedGLS_State, config:Config)
    reads *
  {
    TLS_Init(tgls.tls, config)
      && tgls.history == [config[0]]
  }

  predicate TGLS_Next(tgls:TaggedGLS_State, tgls':TaggedGLS_State)
    reads *
  {
    TLS_Next(tgls.tls, tgls'.tls)
      && (if tgls.tls.t_environment.nextStep.LEnvStepHostIos? && tgls.tls.t_environment.nextStep.actor in tgls.tls.t_servers
      && NodeGrant(tgls.tls.t_servers[tgls.tls.t_environment.nextStep.actor].v, tgls'.tls.t_servers[tgls.tls.t_environment.nextStep.actor].v, UntagLIoOpSeq(tgls.tls.t_environment.nextStep.ios))
      && tgls.tls.t_servers[tgls.tls.t_environment.nextStep.actor].v.held && tgls.tls.t_servers[tgls.tls.t_environment.nextStep.actor].v.epoch < 0xFFFF_FFFF_FFFF_FFFF then
      tgls'.history == tgls.history + [tgls.tls.t_servers[tgls.tls.t_environment.nextStep.actor].v.config[(tgls.tls.t_servers[tgls.tls.t_environment.nextStep.actor].v.my_index + 1) % |tgls.tls.t_servers[tgls.tls.t_environment.nextStep.actor].v.config|]]
      else
        tgls'.history == tgls.history
        )
  }

  predicate ValidTaggedGLSBehavior(tglb:seq<TaggedGLS_State>, config:Config)
    reads *
  {
    && |tglb| > 0
      && TGLS_Init(tglb[0], config)
      && (forall i :: 0 < i < |tglb| ==> TGLS_Next(tglb[i - 1], tglb[i]))
  }
}
