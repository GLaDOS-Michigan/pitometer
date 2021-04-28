include "../DistributedSystem.i.dfy"

include "../../../Common/Collections/Maps2.i.dfy"
include "../Constants.i.dfy"
include "../Environment.i.dfy"
include "../Replica.i.dfy"

include "TimestampedRslSystem.i.dfy"

include "../CommonProof/Constants.i.dfy"

module FailureHelpers_i {
import opened TimestampedRslSystem_i
import opened CommonProof__Constants_i


// TODO: move these to a different file
function {:opaque} EpochQD(nextActionIndex:int) : Timestamp
{
  0
}

function {:opaque} FirstEpochEnd(nextActionIndex:int) : Timestamp
{
  0
}

function {:opaque} SecondEpochEnd(nextActionIndex:int) : Timestamp
{
  0
}

function {:opaque} TBFirstSuspectingHB() : Timestamp
{
  0
}

lemma EpochQDHelper()
{
}

}
