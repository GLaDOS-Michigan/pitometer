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
  0
}

// TODO: move these to a different file
function {:opaque} EpochQD(nextActionIndex:int) : Timestamp
{
  0
}

function {:opaque} FirstEpochEnd() : Timestamp
{
  0
}

function {:opaque} SecondEpochEnd() : Timestamp
{
  FirstEpochEnd() // + EpochLength() + TimeActionRange(0)
}

function {:opaque} HBPeriodEnd() : Timestamp
{
  SecondEpochEnd() // + HBPeriod()
}

function {:opaque} TBFirstSuspectingHB() : Timestamp
{
  SecondEpochEnd() // + HBPeriod() + TimeActionRange(0) + D
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

lemma EpochQDHelper()
{
}

}
