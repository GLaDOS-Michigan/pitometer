// Must be verified with /arith:0
include "Definitions.i.dfy"

module TimeBounds_i {
import opened Definitions_i

function RequestPacketT() : Timestamp
{
  TimeZero() + PrimaryReqT + D
}

function AckPacketT() : Timestamp
{
  RequestPacketT() + MaxQ + BackupProcessT + D
}

function ReplyPacketT() : Timestamp
{
  AckPacketT() + MaxQ + PrimaryAckT + D
}

lemma lemma_receiveThenSendQ(nodeTs:Timestamp, nodeTs':Timestamp, inMsgTs:Timestamp, X:Timestamp, outMsgTs:Timestamp)
  requires TimeLe(nodeTs, inMsgTs + MaxQ)
  requires TimeLe(nodeTs', TimeMax(nodeTs, inMsgTs) + X)
  requires TimeLe(outMsgTs, nodeTs' + D)

  ensures  TimeLe(outMsgTs, inMsgTs + MaxQ + X + D)
{
}

}
