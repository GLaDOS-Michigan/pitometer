include "TimestampedSystem.i.dfy"
include "TimeBounds.i.dfy"

module Performance_Proof {
import opened TimestampedPbSystem_i
import opened TimeBounds_i

predicate Consistency(ps:TimestampedPbState)
{
  |ps.t_nodes| == |ps.constants.config|
    && |ps.t_nodes| > 0
    && (forall i :: 0 <= i < |ps.t_nodes| ==> ps.t_nodes[i].v.con == ps.constants)
    && (forall i :: 0 <= i < |ps.t_nodes| ==> ps.t_nodes[i].v.myIndex == i)
}

type InvType = (TimestampedPbState --> bool)

predicate isInv(inv:InvType)
{
  && (forall ps :: Consistency(ps) ==> inv.requires(ps))
  && (forall ps, con :: TimestampedPbInit(con, ps) ==> Consistency(ps) ==> inv(ps))
  && (forall ps, ps' :: TimestampedPbNext(ps, ps') ==>
      Consistency(ps) ==> inv(ps) ==> Consistency(ps') ==> inv(ps'))
}

predicate invImpliesInv(A:InvType, B:InvType)
{
  && (forall ps :: Consistency(ps) ==> A.requires(ps))
  && (forall ps :: Consistency(ps) ==> B.requires(ps))

  && (forall ps, con :: TimestampedPbInit(con, ps) ==> Consistency(ps) ==> A(ps) ==> B(ps))
  && (forall ps, ps' :: TimestampedPbNext(ps, ps') ==>
      Consistency(ps) ==> Consistency(ps') ==>
      B(ps) ==> A(ps) ==> A(ps') ==> B(ps'))
}

function andInv(A:InvType, B:InvType) : InvType
{
  (ps:TimestampedPbState) requires (A.requires(ps) && B.requires(ps)) =>
    (A(ps) && B(ps))
}

////////////////////////////////////////////////////////////////////////////////

predicate Inv_Primary(ps:TimestampedPbState)
  requires Consistency(ps)
{
  ps.t_nodes[0].v.nextStep == 0 ==>
    ps.t_nodes[0].ts == TimeZero()
}

lemma lemma_Inv_Primary()
  ensures isInv(Inv_Primary)
{
  forall ps, con |
    TimestampedPbInit(con, ps) && Consistency(ps)
    ensures Inv_Primary(ps)
  {}

  forall ps, ps' |
    TimestampedPbNext(ps, ps') && Consistency(ps) && Consistency(ps') && Inv_Primary(ps)
    ensures Inv_Primary(ps')
  {}
}

////////////////////////////////////////////////////////////////////////////////

predicate requestPacketInv(con:Constants, pkt:TimestampedPbPacket)
  requires |con.config| > 0
  requires pkt.msg.v.Request?
{
    TimeEq(pkt.msg.ts, RequestPacketT())
      && pkt.src == con.config[0]
}

predicate Inv_RequestPacket(ps:TimestampedPbState)
  requires Consistency(ps)
{
  forall pkt :: pkt in ps.t_environment.sentPackets ==> pkt.msg.v.Request? ==>
    requestPacketInv(ps.constants, pkt)
}

lemma lemma_Inv_RequestPacket()
  ensures invImpliesInv(Inv_Primary, Inv_RequestPacket)
{
  forall ps, con |
    TimestampedPbInit(con, ps) && Consistency(ps)
    && Inv_Primary(ps)
    ensures Inv_RequestPacket(ps)
  {}

  forall ps, ps' |
    TimestampedPbNext(ps, ps') && Consistency(ps) && Consistency(ps') && Inv_RequestPacket(ps)
    && Inv_Primary(ps) && Inv_Primary(ps')
    ensures Inv_RequestPacket(ps')
  {
    /*
    forall pkt |
      pkt in ps'.t_environment.sentPackets && pkt.msg.v.Request?
      ensures requestPacketInv(ps.constants, pkt)
    {
      if pkt in ps.t_environment.sentPackets {
        assert requestPacketInv(ps.constants, pkt);
      } else {
        var ios := ps.t_environment.nextStep.ios;

        assert !ios[0].LIoOpReceive?;
        assert !ios[0].LIoOpTimeoutReceive?;

        assert pkt in (set io | io in ios && io.LIoOpSend? :: io.s);
        var io :| io in ios && io.LIoOpSend? && io.s == pkt;
        assert pkt.msg.ts == ps'.t_nodes[0].ts + D;
        assert ps'.t_nodes[0].ts == TimeZero() + PrimaryReqT;

        assert pkt.src == ps'.constants.config[0];

        // assert pkt.ts == ps'.t_nodes[idx].ts + D;
        assert requestPacketInv(ps.constants, pkt);
      }
    }
    */
  }
}

////////////////////////////////////////////////////////////////////////////////

// Bounded queueing assumption
predicate QAssumption(s:TimestampedPbState)
  requires Consistency(s)
{
  forall idx, ios, hstep :: 0 <= idx < |s.constants.config|
    ==> s.t_environment.nextStep == LEnvStepHostIos(s.constants.config[idx], ios, hstep)
    ==> (forall io :: io in s.t_environment.nextStep.ios && io.LIoOpReceive? ==>
      // this means that max(replica.ts, msg.ts) <= msg.ts + MaxQueueTime
      TimeLe(s.t_nodes[idx].ts, io.r.msg.ts + MaxQ)
    )
}

predicate ackPacketInv(con:Constants, pkt:TimestampedPbPacket)
  requires |con.config| > 0
  requires pkt.msg.v.Ack?
{
    TimeLe(pkt.msg.ts, AckPacketT())
      && pkt.dst == con.config[0]
}

predicate Inv_AckPacket(ps:TimestampedPbState)
  requires Consistency(ps)
{
  forall pkt :: pkt in ps.t_environment.sentPackets ==> pkt.msg.v.Ack? ==>
    ackPacketInv(ps.constants, pkt)
}

lemma lemma_Inv_AckPacket()
  ensures invImpliesInv(andInv(QAssumption, Inv_RequestPacket), Inv_AckPacket)
{
  forall ps, con |
    TimestampedPbInit(con, ps) && Consistency(ps)
    && QAssumption(ps) && Inv_RequestPacket(ps)
    ensures Inv_AckPacket(ps)
  {}

  forall ps, ps' |
    TimestampedPbNext(ps, ps') && Consistency(ps) && Consistency(ps') && Inv_AckPacket(ps)
    && (QAssumption(ps) && Inv_RequestPacket(ps)) && (QAssumption(ps') && Inv_RequestPacket(ps'))
    ensures Inv_AckPacket(ps')
  {
    forall pkt |
      pkt in ps'.t_environment.sentPackets && pkt.msg.v.Ack?
      ensures ackPacketInv(ps.constants, pkt)
    {
      if pkt in ps.t_environment.sentPackets {
        assert ackPacketInv(ps.constants, pkt);
      } else {
        var i, ios2 :| TimestampedPbNextOneReplica(ps, ps', i, ios2);
        var ios := ps.t_environment.nextStep.ios;
        // assert ios[0].r.msg
        assert ios[0].LIoOpReceive?;
        lemma_receiveThenSendQ(ps.t_nodes[i].ts, ps'.t_nodes[i].ts,
            ios[0].r.msg.ts, BackupProcessT, ios[1].s.msg.ts);
        assert ackPacketInv(ps.constants, pkt);
      }
    }
  }
}

////////////////////////////////////////////////////////////////////////////////

predicate replyPacketInv(con:Constants, pkt:TimestampedPbPacket)
  requires |con.config| > 0
  requires pkt.msg.v.ClientReply?
{
    TimeLe(pkt.msg.ts, ReplyPacketT())
      && pkt.src == con.config[0]
}

predicate Inv_ReplyPacket(ps:TimestampedPbState)
  requires Consistency(ps)
{
  forall pkt :: pkt in ps.t_environment.sentPackets ==> pkt.msg.v.ClientReply? ==>
    replyPacketInv(ps.constants, pkt)
}

lemma lemma_Inv_ReplyPacket()
  ensures invImpliesInv(andInv(QAssumption, Inv_AckPacket), Inv_ReplyPacket)
{
  forall ps, con |
    TimestampedPbInit(con, ps) && Consistency(ps)
    && QAssumption(ps) && Inv_AckPacket(ps)
    ensures Inv_ReplyPacket(ps)
  {}

  forall ps, ps' |
    TimestampedPbNext(ps, ps') && Consistency(ps) && Consistency(ps') && Inv_ReplyPacket(ps)
    && (QAssumption(ps) && Inv_AckPacket(ps)) && (QAssumption(ps') && Inv_AckPacket(ps'))
    ensures Inv_ReplyPacket(ps')
  {
    forall pkt |
      pkt in ps'.t_environment.sentPackets && pkt.msg.v.ClientReply?
      ensures replyPacketInv(ps.constants, pkt)
    {
      if pkt in ps.t_environment.sentPackets {
        assert replyPacketInv(ps.constants, pkt);
      } else {
        var i, ios2 :| TimestampedPbNextOneReplica(ps, ps', i, ios2);
        var ios := ps.t_environment.nextStep.ios;
        // assert ios[0].r.msg
        assert ios[0].LIoOpReceive?;

        assert i == 0;
        assert ps.t_nodes[0].v.nextStep == 1;

        lemma_receiveThenSendQ(ps.t_nodes[i].ts, ps'.t_nodes[i].ts,
            ios[0].r.msg.ts, PrimaryAckT, ios[1].s.msg.ts);
        assert replyPacketInv(ps.constants, pkt);
      }
    }
  }
}


////////////////////////////////////////////////////////////////////////////////

predicate performanceGuarantee(ps:TimestampedPbState)
{
  forall pkt :: pkt in ps.t_environment.sentPackets ==> pkt.msg.v.ClientReply? ==> TimeLe(pkt.msg.ts, ReplyPacketT())
}

predicate Inv(ps:TimestampedPbState)
{
  Consistency(ps)
  && Inv_RequestPacket(ps)
}

lemma topLevel()
{

}

lemma Init_Inv(con:Constants, ps:TimestampedPbState)
  requires TimestampedPbInit(con, ps)
  ensures Inv(ps)
{
}

predicate invTrue(inv:InvType)
{
  forall b ::
  && (forall ps :: Consistency(ps) ==> inv.requires(ps))
  && (forall ps, con :: TimestampedPbInit(con, ps) ==> Consistency(ps) ==> inv(ps))
  && (forall ps, ps' :: TimestampedPbNext(ps, ps') ==>
      Consistency(ps) ==> inv(ps) ==> Consistency(ps') ==> inv(ps'))
}

// is(A) && (A -> B) ==>
// is(A && B)

lemma Next_Inv (ps:TimestampedPbState, ps':TimestampedPbState)
  requires Inv(ps)
  requires TimestampedPbNext(ps, ps')
  ensures Inv(ps')
{
  lemma_Inv_Primary();
  // isInv(Inv_Primary)
  lemma_Inv_RequestPacket();
  // invImpliesInv(Inv_Primary, Inv_RequestPacket)
  lemma_Inv_AckPacket();
  // invImpliesInv(andInv(QAssumption, Inv_RequestPacket), Inv_AckPacket)
  lemma_Inv_ReplyPacket();
  // invImpliesInv(andInv(QAssumption, Inv_AckPacket), Inv_ReplyPacket)

}

}
