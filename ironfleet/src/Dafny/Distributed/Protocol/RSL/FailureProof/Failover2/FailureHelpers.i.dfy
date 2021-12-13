include "../../DistributedSystem.i.dfy"

include "../../../../Common/Collections/Maps2.i.dfy"
include "../../Constants.i.dfy"
include "../../Environment.i.dfy"
include "../../Replica.i.dfy"

include "../TimestampedRslSystem.i.dfy"

include "../../CommonProof/Constants.i.dfy"

module FailureHelpers_i {
import opened TimestampedRslSystem_i
import opened CommonProof__Constants_i

function {:opaque} TBEpoch1() : Timestamp
{
  0
}

function {:opaque} TBEpoch2() : Timestamp
{
  TBEpoch1() + EpochQD(7) + StepToTimeDelta(RslStep(7)) + EpochLength
}

function {:opaque} EpochQD(nextActionIndex:int) : Timestamp
{
  if nextActionIndex == 8 then
    StepToTimeDelta(RslStep(7))
  else if nextActionIndex == 9 then
    StepToTimeDelta(RslStep(7)) + StepToTimeDelta(RslStep(8))
  else if nextActionIndex == 0 then
    StepToTimeDelta(RslStep(7)) + StepToTimeDelta(RslStep(8)) + StepToTimeDelta(RslStep(9))
  else if nextActionIndex == 1 then
    StepToTimeDelta(RslStep(7)) + StepToTimeDelta(RslStep(8)) + StepToTimeDelta(RslStep(9)) +
    Timeout() +
    StepToTimeDelta(RslStep(0))
  else if nextActionIndex == 2 then
    StepToTimeDelta(RslStep(7)) + StepToTimeDelta(RslStep(8)) + StepToTimeDelta(RslStep(9)) +
    Timeout() +
    StepToTimeDelta(RslStep(0)) + StepToTimeDelta(RslStep(1))
  else if nextActionIndex == 3 then
    StepToTimeDelta(RslStep(7)) + StepToTimeDelta(RslStep(8)) + StepToTimeDelta(RslStep(9)) +
    Timeout() +
    StepToTimeDelta(RslStep(0)) + StepToTimeDelta(RslStep(1)) + StepToTimeDelta(RslStep(2))
  else if nextActionIndex == 4 then
    StepToTimeDelta(RslStep(7)) + StepToTimeDelta(RslStep(8)) + StepToTimeDelta(RslStep(9)) +
    Timeout() +
    StepToTimeDelta(RslStep(0)) + StepToTimeDelta(RslStep(1)) + StepToTimeDelta(RslStep(2)) +
    StepToTimeDelta(RslStep(3))
  else if nextActionIndex == 5 then
    StepToTimeDelta(RslStep(7)) + StepToTimeDelta(RslStep(8)) + StepToTimeDelta(RslStep(9)) +
    Timeout() +
    StepToTimeDelta(RslStep(0)) + StepToTimeDelta(RslStep(1)) + StepToTimeDelta(RslStep(2)) +
    StepToTimeDelta(RslStep(3)) + StepToTimeDelta(RslStep(4))
  else if nextActionIndex == 6 then
    StepToTimeDelta(RslStep(7)) + StepToTimeDelta(RslStep(8)) + StepToTimeDelta(RslStep(9)) +
    Timeout() +
    StepToTimeDelta(RslStep(0)) + StepToTimeDelta(RslStep(1)) + StepToTimeDelta(RslStep(2)) +
    StepToTimeDelta(RslStep(3)) + StepToTimeDelta(RslStep(4)) + StepToTimeDelta(RslStep(5))
  else if nextActionIndex == 7 then
    StepToTimeDelta(RslStep(7)) + StepToTimeDelta(RslStep(8)) + StepToTimeDelta(RslStep(9)) +
    Timeout() +
    StepToTimeDelta(RslStep(0)) + StepToTimeDelta(RslStep(1)) + StepToTimeDelta(RslStep(2)) +
    StepToTimeDelta(RslStep(3)) + StepToTimeDelta(RslStep(4)) + StepToTimeDelta(RslStep(5)) +
    StepToTimeDelta(RslStep(6))
  else
    0
}

function {:opaque} HeartbeatQD(nextActionIndex:int) : Timestamp
{
  0
}

function {:opaque} HBPeriodEnd() : Timestamp
{
  TBEpoch2() + EpochQD(7) + StepToTimeDelta(RslStep(7)) + HBPeriod
}

function {:opaque} TBFirstSuspectingHB() : Timestamp
{
  HBPeriodEnd() + HeartbeatQD(9) + StepToTimeDelta(RslStep(9)) + D
}

function {:opaque} TBBecomeSuspector() : Timestamp
{
  0 // TODO
}

// function {:opaque} TBJustBeforeNewView() : Timestamp
// {
  // TBEpoch2() + EpochQD(7) + StepToTimeDelta(RslStep(7))
// }

function {:opaque} ActionsUpTo(nextActionIndex:int) : Timestamp
{
  if nextActionIndex == 1 then
    StepToTimeDelta(RslStep(0)) + Timeout()
  else if nextActionIndex == 2 then
    StepToTimeDelta(RslStep(0)) + Timeout() +
    StepToTimeDelta(RslStep(1))
  else if nextActionIndex == 3 then
    StepToTimeDelta(RslStep(0)) + Timeout() +
    StepToTimeDelta(RslStep(1)) + StepToTimeDelta(RslStep(2))
  else if nextActionIndex == 4 then
    StepToTimeDelta(RslStep(0)) + Timeout() +
    StepToTimeDelta(RslStep(1)) + StepToTimeDelta(RslStep(2)) + StepToTimeDelta(RslStep(3))
  else if nextActionIndex == 5 then
    StepToTimeDelta(RslStep(0)) + Timeout() +
    StepToTimeDelta(RslStep(1)) + StepToTimeDelta(RslStep(2)) + StepToTimeDelta(RslStep(3)) +
    StepToTimeDelta(RslStep(4))
  else if nextActionIndex == 6 then
    StepToTimeDelta(RslStep(0)) + Timeout() +
    StepToTimeDelta(RslStep(1)) + StepToTimeDelta(RslStep(2)) + StepToTimeDelta(RslStep(3)) +
    StepToTimeDelta(RslStep(4)) + StepToTimeDelta(RslStep(5))
  else if nextActionIndex == 7 then
    StepToTimeDelta(RslStep(0)) + Timeout() +
    StepToTimeDelta(RslStep(1)) + StepToTimeDelta(RslStep(2)) + StepToTimeDelta(RslStep(3)) +
    StepToTimeDelta(RslStep(4)) + StepToTimeDelta(RslStep(5)) + StepToTimeDelta(RslStep(6))
  else if nextActionIndex == 8 then
    StepToTimeDelta(RslStep(0)) + Timeout() +
    StepToTimeDelta(RslStep(1)) + StepToTimeDelta(RslStep(2)) + StepToTimeDelta(RslStep(3)) +
    StepToTimeDelta(RslStep(4)) + StepToTimeDelta(RslStep(5)) + StepToTimeDelta(RslStep(6)) +
    StepToTimeDelta(RslStep(7))
  else
    0 // don't care about these cases
}

function {:opaque} TBJustBeforeNewView(nextActionIndex:int) : Timestamp
{
  TBFirstSuspectingHB() + MaxQueueTime + ActionsUpTo(nextActionIndex)
}

function {:opaque} TBNewView() : Timestamp
{
  TBJustBeforeNewView(8) + StepToTimeDelta(RslStep(8))
}

lemma EpochQDHelper(t:Timestamp, t':Timestamp)
  requires TimeLe(t, TBEpoch1() + EpochQD(7));
  requires TimeLe(t', t + StepToTimeDelta(RslStep(7)))
  ensures TimeLe(t' + EpochLength, TBEpoch2());
{
  reveal_TBEpoch2();
}

lemma HeartbeatQDHelper(t:Timestamp, t':Timestamp)
  requires TimeLe(t, HBPeriodEnd() + HeartbeatQD(9));
  requires TimeLe(t', t + StepToTimeDelta(RslStep(9)))
  ensures TimeLe(t' + D, TBFirstSuspectingHB());
{
  reveal_TBFirstSuspectingHB();
}

}
