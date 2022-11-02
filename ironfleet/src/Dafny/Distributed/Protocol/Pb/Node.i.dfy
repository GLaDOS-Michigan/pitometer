include "Types.i.dfy"
include "../../Common/Collections/Sets.i.dfy"

module Protocol_Node_i {
import opened Types_i
import opened Native__Io_s
import opened Collections__Sets_i

type Config = seq<EndPoint>

datatype Constants = Constants(
  config:Config,
  client: EndPoint
  )

datatype Node = Node(myIndex:int, nextStep:int, gotAck:set<EndPoint>, con:Constants)

predicate NodeInit(s:Node, my_index:int, con:Constants)
{
  s.myIndex == my_index
    && 0 <= my_index < |con.config|
    && s.nextStep == 0
    && s.con == con
}

predicate PrimarySendRequest(s:Node, s':Node, ios:seq<PbIo>)
{
  s.myIndex == 0
    && s.nextStep == 0
    && s' == s.(nextStep := 1)

    // broadcast to backups
    && (|s.con.config| > 0 ==>
       |ios| == |s.con.config[1..]|
        && forall idx :: 0 <= idx < |s.con.config[1..]| ==>
          ios[idx] == LIoOpSend(LPacket(s.con.config[1..][idx], s.con.config[s.myIndex], Request()))
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
      && ios[0].r.src in s.con.config
      && ios[0].r.msg.Ack?
      && s' == s.(gotAck := s.gotAck + { ios[0].r.src } )
      && if s'.gotAck == MapSeqToSet(s.con.config, (i:EndPoint) => i) then
          && |ios| == 2
          && ios[1].LIoOpSend?
          && ios[1].s.msg.ClientReply?
          && ios[1].s.dst == s.con.client
        else
            s == s' && |ios| == 1
}

predicate BackupReceiveRequest(s:Node, s':Node, ios:seq<PbIo>)
{
  s.myIndex != 0
    && s' == s
    && |ios| >= 1
    && if ios[0].LIoOpTimeoutReceive? then
        |ios| == 1
      else
        ios[0].LIoOpReceive?
        && ios[0].r.src in s.con.config
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
