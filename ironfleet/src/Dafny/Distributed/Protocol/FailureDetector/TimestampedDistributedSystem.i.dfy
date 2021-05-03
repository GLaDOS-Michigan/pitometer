include "DistributedSystem.dfy"
include "Timestamps/TimestampedType.dfy"

module FailureDetector_TimestampedDistributedSystem {
import opened FDTimestampedEnv_s
import opened FailureDetector_DistributedSystem


type TimestampedAgent = TimestampedType<Agent>

datatype TFD_State = TFD_State(
  config:Config,
  t_environment:TimestampedLEnvironment<EndPoint, FDMessage, FDStep>,
  t_servers:map<EndPoint,TimestampedAgent>
  )

  function UntagFDState(s:TFD_State) : FD_State
  {
    FD_State(
      UntagLEnvironment(s.t_environment),
      (map id | id in s.t_servers :: s.t_servers[id].v)
      )
  }

  predicate TFD_Init(config:Config, s:TFD_State, heartbeatInterval:int, timeoutInterval:int)
  {
    && FD_Init(config, UntagFDState(s), heartbeatInterval, timeoutInterval)
      && s.config == config
      && LEnvironment_Init(s.t_environment)
      && forall id :: id in s.t_servers ==> s.t_servers[id].ts == TimeZero()
  }

  predicate TFD_NextOneServer(s:TFD_State, s':TFD_State, actor:EndPoint, ios:seq<TimestampedLIoOp<EndPoint, FDMessage>>, hstep:FDStep)
  {
    && actor in s.t_servers
    && FD_NextOneServer(UntagFDState(s), UntagFDState(s'), actor, UntagLIoOpSeq(ios))
      && s.t_environment.nextStep.nodeStep == hstep
  }
}
