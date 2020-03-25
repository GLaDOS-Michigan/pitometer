include "Environment.s.dfy"
include "Host.s.dfy"
include "DistributedSystem.s.dfy"
include "../Collections/Seqs.i.dfy"
include "Performance.s.dfy"

abstract module TaggedDistributedSystem_s {
  import opened Environment_s
  import opened P_s : Performance_s
  import opened Collections__Seqs_i

  type PerformanceReport = PerfExpr
  type PerfReport = PerfExpr

  datatype TaggedType<Type> = TaggedType(v:Type, pr:PerformanceReport)

  type TaggedLPacket<IdType, MessageType(==)> = LPacket<IdType, TaggedType<MessageType>>

  type TaggedLIoOp<IdType, MessageType(==)> = LIoOp<IdType, TaggedType<MessageType>>

  type TaggedLEnvStep<IdType, MessageType(==), NodeStepType> = LEnvStep<IdType,TaggedType<MessageType>, NodeStepType>

  type TaggedLHostInfo<IdType, MessageType(==)> = LHostInfo<IdType,TaggedType<MessageType>>

  type TaggedLEnvironment<IdType, MessageType(==), NodeStepType> = LEnvironment<IdType, TaggedType<MessageType>, NodeStepType>

  type TaggedHostState = TaggedType<HostState>

  datatype TaggedDS_State = TaggedDS_State(
    config:ConcreteConfiguration,
    t_environment:TaggedLEnvironment<EndPoint, seq<byte>, HostStep>,
    t_servers:map<EndPoint,TaggedHostState>,
    clients:set<EndPoint>
    )
      
  function UntagTaggedType<T>(t_t: TaggedType<T>) : T
  {
    t_t.v
  }

  function UntagLPacket<I,M>(pkt: TaggedLPacket<I,M>) : LPacket<I,M>
  {
    LPacket(pkt.dst, pkt.src, pkt.msg.v)
  }

  function UntagSentPkts<I,M>(t_sentPkts : set<TaggedLPacket<I,M>>) : set<LPacket<I,M>>
  {
    set pkt | pkt in t_sentPkts :: UntagLPacket(pkt)
  }

  function UntagLEnvStep(t_nextStep : TaggedLEnvStep) : LEnvStep
  {
    match t_nextStep
    {
      case LEnvStepHostIos(actor, ios, nodeStep) => LEnvStepHostIos(actor, UntagLIoOpSeq(ios), nodeStep)
      case LEnvStepDeliverPacket(p) => LEnvStepDeliverPacket(UntagLPacket(p))
      case LEnvStepAdvanceTime() => LEnvStepAdvanceTime
      case LEnvStepStutter() => LEnvStepStutter
    }
  }

  function UntagHostInfo(t_hi: TaggedLHostInfo) : LHostInfo
  {
    LHostInfo(MapSeqToSeq(t_hi.queue, UntagLPacket))
  }

  function UntagHostInfoMap<I,M>(t_hostInfo:map<I, LHostInfo<I,TaggedType<M>>>) : map<I,LHostInfo<I,M>>
  {
    map id | id in t_hostInfo :: UntagHostInfo(t_hostInfo[id])
  }

  function UntagLEnvironment<I,M,S>(t_env: TaggedLEnvironment<I,M,S>) : LEnvironment<I,M,S>
  {
    LEnvironment(t_env.time,
      UntagSentPkts(t_env.sentPackets),
      UntagHostInfoMap(t_env.hostInfo),
      UntagLEnvStep(t_env.nextStep))
  }

  function UntagServers(t_servers: map<EndPoint, TaggedHostState>) : map<EndPoint, HostState>
  {
    map id | id in t_servers :: t_servers[id].v
  }
  
  function UntagDS_State(tds:TaggedDS_State) : DS_State
  {
    DS_State(tds.config,
      UntagLEnvironment(tds.t_environment),
      UntagServers(tds.t_servers),
      tds.clients)
  }
  
  function UntagLIoOp(t_io : TaggedLIoOp) : LIoOp
  {
    match t_io
    {
      case LIoOpSend(s) => LIoOpSend(UntagLPacket(s))
      case LIoOpReceive(r) => LIoOpReceive(UntagLPacket(r))
      case LIoOpTimeoutReceive() => LIoOpTimeoutReceive()
      case LIoOpReadClock(t) => LIoOpReadClock(t)
    }
  }

  function {:opaque} UntagLIoOpSeq(t_ios: seq<TaggedLIoOp>) : seq<LIoOp>
    ensures |UntagLIoOpSeq(t_ios)| == |t_ios|
    ensures forall i :: 0 <= i < |UntagLIoOpSeq(t_ios)| ==> UntagLIoOpSeq(t_ios)[i] == UntagLIoOp(t_ios[i])
  {
    MapSeqToSeq(t_ios, UntagLIoOp)
  }

  function GetReceivePRs(ios:seq<TaggedLIoOp>) : seq<PerfReport>
    decreases |ios|
    ensures (forall io :: io in ios ==> !io.LIoOpReceive?) ==> GetReceivePRs(ios) == []
  {
    if |ios| == 0 then
      []
    else if ios[0].LIoOpReceive? then
      [ios[0].r.msg.pr] + GetReceivePRs(ios[1..])
    else
      GetReceivePRs(ios[1..])
  }

  predicate TDS_Init(tds: TaggedDS_State, config:ConcreteConfiguration)
    reads *
  {
    DS_Init(UntagDS_State(tds), config)
      && forall id :: id in tds.t_servers ==> tds.t_servers[id].pr == PerfZero()
  }

  predicate TDS_NextOneServer(tds: TaggedDS_State, tds': TaggedDS_State, id:EndPoint, ios:seq<TaggedLIoOp<EndPoint,seq<byte>>>, hstep:HostStep)
    requires id in tds.t_servers;
    reads *
  {
    DS_NextOneServer(UntagDS_State(tds), UntagDS_State(tds'), id, UntagLIoOpSeq(ios))
      && (var recvTime := PerfMax(multiset(GetReceivePRs(ios)) + multiset{tds.t_servers[id].pr});
      var totalTime := PerfAdd2(recvTime, PerfStep(hstep));
      tds'.t_servers[id].pr == totalTime
      )
      && (forall t_io :: t_io in ios && t_io.LIoOpSend? ==> t_io.s.msg.pr == tds'.t_servers[id].pr)
  }

  predicate TDS_Next(tds:TaggedDS_State, tds': TaggedDS_State)
    reads *
  {
    DS_Next(UntagDS_State(tds), UntagDS_State(tds'))
      && LEnvironment_Next(tds.t_environment, tds'.t_environment)
      && if tds.t_environment.nextStep.LEnvStepHostIos? && tds.t_environment.nextStep.actor in tds.t_servers then
      TDS_NextOneServer(tds, tds', tds.t_environment.nextStep.actor, tds.t_environment.nextStep.ios, tds.t_environment.nextStep.nodeStep)
    else
      tds'.t_servers == tds.t_servers
  }

  predicate IsValidBehavior(tdb:seq<TaggedDS_State>, config:ConcreteConfiguration)
    reads *
  {
    |tdb| > 0
      && TDS_Init(tdb[0], config)
      && (forall i {:trigger TDS_Next(tdb[i], tdb[i+1])} :: 0 <= i < |tdb| - 1 ==> TDS_Next(tdb[i], tdb[i+1]))
  }

}
