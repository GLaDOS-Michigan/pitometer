include "../Node.i.dfy"
  include "../RefinementProof/DistributedSystem.i.dfy"
include "../../../Services/Lock/LockTaggedDistributedSystem.i.dfy"

module TaggedGLS_i {
import opened Protocol_Node_i
import opened DistributedSystem_i
import opened LockTaggedDistributedSystem_i

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
    LS_Init(UntagLS_State(tls), config)
      && LEnvironment_Init(tls.t_environment)
      && tls.t_servers[config[0]].pr == PerfZero
      && forall id :: id in tls.t_servers && id != config[0] ==> tls.t_servers[id].pr == PerfVoid
  }

  predicate TLS_NextOneServer(tls: TaggedLS_State, tls': TaggedLS_State, id:EndPoint, ios:seq<TaggedLIoOp<EndPoint, LockMessage>>, hstep:HostStep)
    requires id in tls.t_servers;
    reads *
  {
    LS_NextOneServer(UntagLS_State(tls), UntagLS_State(tls'), id, UntagLIoOpSeq(ios), hstep)
      && (var recvTime := PerfMax(GetReceivePRs(ios) + [tls.t_servers[id].pr]);
      var totalTime := PerfAdd2(recvTime, GetStepRuntime(hstep));
      tls'.t_servers[id].pr == totalTime
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
            && (forall t_io :: t_io in tls.t_environment.nextStep.ios && t_io.LIoOpSend? ==> t_io.s.msg.pr == PerformanceReport(0, 0))
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
      && TGLS_Init(tglb[0], config)
      && (forall i :: i >= 0 ==> TGLS_Next(tglb[i], tglb[i+1]))
  }
}
