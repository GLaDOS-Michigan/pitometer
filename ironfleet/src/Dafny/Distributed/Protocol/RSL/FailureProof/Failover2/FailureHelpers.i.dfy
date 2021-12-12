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

// TODO: move these to a different file
function {:opaque} EpochQD(nextActionIndex:int) : Timestamp
{
  0
}

function {:opaque} HeartbeatQD(nextActionIndex:int) : Timestamp
{
  0
}

function {:opaque} HBPeriodEnd() : Timestamp
{
  TBEpoch2() // + HBPeriod()
}

function {:opaque} TBFirstSuspectingHB() : Timestamp
{
  HBPeriodEnd() + HeartbeatQD(9) + StepToTimeDelta(RslStep(9)) + D
}

function {:opaque} TBBecomeSuspector() : Timestamp
{
  0 // TODO
}

function {:opaque} TBJustBeforeNewView() : Timestamp
{
  0 // TODO
}

function {:opaque} TBFirstNewViewHB() : Timestamp
{
  0 // TODO
}

function {:opaque} FailoverTime() : Timestamp
{
  0
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
