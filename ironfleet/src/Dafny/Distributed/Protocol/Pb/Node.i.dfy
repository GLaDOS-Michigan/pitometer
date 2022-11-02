include "Types.i.dfy"
include "../../Common/Collections/Sets.i.dfy"

module Protocol_Node_i {
import opened Types_i
import opened Native__Io_s
import opened Collections__Sets_i

type Config = seq<EndPoint>

datatype Node = Node(myIndex:int, nextStep:int, gotAck:set<EndPoint>, clientEndPoint:EndPoint, config:Config)

predicate NodeInit(s:Node, my_index:int, config:Config, clientEndPoint:EndPoint)
{
  s.myIndex == my_index
    && 0 <= my_index < |config|
    && s.nextStep == 0
    && s.config == config
    && s.clientEndPoint == clientEndPoint
}

predicate PrimarySendRequest(s:Node, s':Node, ios:seq<PbIo>)
{
  s.myIndex == 0
    && s.nextStep == 0
    && s' == s.(nextStep := 1)

    // broadcast to backups
    && forall hd, backups :: s.config == [hd] + backups ==>
          (|ios| == |backups|
            && forall idx :: 0 <= idx < |backups| ==>
              ios[idx] == LIoOpSend(LPacket(backups[idx], s.config[s.myIndex], Request()))
          )
}

predicate PrimaryReceiveAck(s:Node, s':Node, ios:seq<PbIo>)
{
  s.myIndex == 0
    && s.nextStep == 1
    && |ios| >= 1
    && if ios[0].LIoOpTimeoutReceive? then
      s == s' && |ios| == 1
    else
      ios[0].LIoOpReceive?
      && ios[0].r.src in s.config
      && ios[0].r.msg.Ack?
      && s' == s.(gotAck := s.gotAck + { ios[0].r.src } )
      && if s'.gotAck == MapSeqToSet(s.config, (i:EndPoint) => i) then
          && |ios| == 2
          && ios[1].LIoOpSend?
          && ios[1].s.msg.ClientReply?
          && ios[1].s.dst == s.clientEndPoint
        else
            s == s' && |ios| == 1
}

predicate BackupReceiveRequest(s:Node, s':Node, ios:seq<PbIo>)
{
  s.myIndex != 0
    && s' == s
    && |ios| >= 1
    && if ios[0].LIoOpTimeoutReceive? then
      s == s' && |ios| == 1
    else
      ios[0].LIoOpReceive?
      && ios[0].r.src in s.config
      && ios[0].r.msg.Request?
      && |ios| == 2
      && ios[1].LIoOpSend?
      && ios[1].s.msg.Ack?
      && ios[1].s.dst == ios[0].r.src
}

predicate NodeNext(s:Node, s':Node, ios:seq<PbIo>)
{
    PrimarySendRequest(s, s', ios)
      || PrimaryReceiveAck(s, s', ios)
      || BackupReceiveRequest(s, s', ios)
}

}
