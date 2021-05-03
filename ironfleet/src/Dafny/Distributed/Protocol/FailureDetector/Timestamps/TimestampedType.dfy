include "../../../Common/Framework/Environment.s.dfy"
include "Timestamp.dfy"

module FDTimestampedEnv_s {
  import opened Environment_s
  import opened FDTimestamp_s
  import opened Collections__Seqs_i

  datatype TimestampedType<Type> = TimestampedType(v:Type, ts:Timestamp)

  type TimestampedLPacket<IdType, MessageType(==)> = LPacket<IdType, TimestampedType<MessageType>>

  type TimestampedLIoOp<IdType, MessageType(==)> = LIoOp<IdType, TimestampedType<MessageType>>

  type TimestampedLEnvStep<IdType, MessageType(==), NodeStepType> = LEnvStep<IdType,TimestampedType<MessageType>, NodeStepType>

  type TimestampedLHostInfo<IdType, MessageType(==)> = LHostInfo<IdType,TimestampedType<MessageType>>

  type TimestampedLEnvironment<IdType, MessageType(==), NodeStepType> = LEnvironment<IdType, TimestampedType<MessageType>, NodeStepType>

  function UntagTimestampedType<T>(t_t: TimestampedType<T>) : T
  {
    t_t.v
  }

  function UntagLPacket<I,M>(pkt: TimestampedLPacket<I,M>) : LPacket<I,M>
  {
    LPacket(pkt.dst, pkt.src, pkt.msg.v)
  }

  function UntagSentPkts<I,M>(t_sentPkts : set<TimestampedLPacket<I,M>>) : set<LPacket<I,M>>
  {
    set pkt | pkt in t_sentPkts :: UntagLPacket(pkt)
  }

  function UntagLEnvStep(t_nextStep : TimestampedLEnvStep) : LEnvStep
  {
    match t_nextStep
    {
      case LEnvStepHostIos(actor, ios, nodeStep) => LEnvStepHostIos(actor, UntagLIoOpSeq(ios), nodeStep)
      case LEnvStepDeliverPacket(p) => LEnvStepDeliverPacket(UntagLPacket(p))
      case LEnvStepAdvanceTime() => LEnvStepAdvanceTime
      case LEnvStepStutter() => LEnvStepStutter
    }
  }

  function UntagHostInfo(t_hi: TimestampedLHostInfo) : LHostInfo
  {
    LHostInfo(MapSeqToSeq(t_hi.queue, UntagLPacket))
  }

  function UntagHostInfoMap<I,M>(t_hostInfo:map<I, LHostInfo<I,TimestampedType<M>>>) : map<I,LHostInfo<I,M>>
  {
    map id | id in t_hostInfo :: UntagHostInfo(t_hostInfo[id])
  }

  function UntagLEnvironment<I,M,S>(t_env: TimestampedLEnvironment<I,M,S>) : LEnvironment<I,M,S>
  {
    LEnvironment(t_env.time,
      UntagSentPkts(t_env.sentPackets),
      UntagHostInfoMap(t_env.hostInfo),
      UntagLEnvStep(t_env.nextStep))
  }

  function UntagLIoOp(t_io : TimestampedLIoOp) : LIoOp
  {
    match t_io
    {
      case LIoOpSend(s) => LIoOpSend(UntagLPacket(s))
      case LIoOpReceive(r) => LIoOpReceive(UntagLPacket(r))
      case LIoOpTimeoutReceive() => LIoOpTimeoutReceive()
      case LIoOpReadClock(t) => LIoOpReadClock(t)
    }
  }

  function {:opaque} UntagLIoOpSeq(t_ios: seq<TimestampedLIoOp>) : seq<LIoOp>
    ensures |UntagLIoOpSeq(t_ios)| == |t_ios|
    ensures forall i :: 0 <= i < |UntagLIoOpSeq(t_ios)| ==> UntagLIoOpSeq(t_ios)[i] == UntagLIoOp(t_ios[i])
  {
    MapSeqToSeq(t_ios, UntagLIoOp)
  }
}
