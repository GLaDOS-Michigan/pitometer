include "DistributedSystem.dfy"
include "Timestamps/TimestampedType.dfy"

module FailureDetector_TimestampedDistributedSystem {
import opened FDTimestampedEnv_s
import opened FailureDetector_DistributedSystem

ghost const Delay:Timestamp
ghost const Q:Timestamp

ghost const TryRecv:Timestamp
ghost const MbeTimeout:Timestamp
ghost const MbeSend:Timestamp
ghost const Unknown:Timestamp

predicate TimeLe(lhs:Timestamp, rhs:Timestamp)
{
  lhs <= rhs
}

function StepToTimeDelta(hstep:FDStep) : Timestamp
{
  if hstep == NodeStep then
    MbeSend
  else if hstep == DetectorStep(0) then
    TryRecv
  else if hstep == DetectorStep(1) then
    MbeTimeout
  else
    Unknown
}

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

  function FD_NoRecvPerfUpdate(node_pr:Timestamp, hstep:FDStep) : Timestamp
  {
    var total_time := TimeAdd2(node_pr, StepToTimeDelta(hstep));
    total_time
  }

  function FD_RecvPerfUpdate(node_ts:Timestamp, pkt_ts:Timestamp, hstep:FDStep) : Timestamp
  {
    var handlerStartTime := TimeMax(pkt_ts, node_ts);
    var total_time := TimeAdd2(handlerStartTime, StepToTimeDelta(hstep));
    total_time
  }

  function FD_TimeoutPerfUpdate(node_ts:Timestamp, hstep:FDStep) : Timestamp
  {
    node_ts + Timeout() + StepToTimeDelta(hstep)
  }

  predicate TFD_NextOneServer(s:TFD_State, s':TFD_State, actor:EndPoint, ios:seq<TimestampedLIoOp<EndPoint, FDMessage>>)
  {
    && s.config == s'.config
    && actor in s.t_servers
    && FD_NextOneServer(UntagFDState(s), UntagFDState(s'), actor, UntagLIoOpSeq(ios))
    && LEnvironment_Next(s.t_environment, s'.t_environment)
    && s'.t_servers == s.t_servers[actor := s'.t_servers[actor]]

    && var hstep := s.t_environment.nextStep.nodeStep;
    (if |ios| > 0 && ios[0].LIoOpReceive? then
      && ios[0] in ios
      && TimeLe(s'.t_servers[actor].ts, FD_RecvPerfUpdate(s.t_servers[actor].ts, ios[0].r.msg.ts, hstep))
      && (ios[0].r.msg.ts <= s.t_servers[actor].ts + Timeout())
    else if |ios| > 0 && ios[0].LIoOpTimeoutReceive? then
      && TimeLe(s'.t_servers[actor].ts, FD_TimeoutPerfUpdate(s.t_servers[actor].ts, hstep))
    else
      && TimeLe(s'.t_servers[actor].ts, FD_NoRecvPerfUpdate(s.t_servers[actor].ts, hstep))
      )

      && (forall io :: io in ios && io.LIoOpSend? ==>
          TimeLe(io.s.msg.ts, s'.t_servers[actor].ts + Delay)
      )

      && (forall io :: io in ios && io.LIoOpReadClock? ==>
          && TimeLe(io.t, s'.t_servers[actor].ts)
          && TimeLe(s.t_servers[actor].ts, io.t)
      )
  }
}
