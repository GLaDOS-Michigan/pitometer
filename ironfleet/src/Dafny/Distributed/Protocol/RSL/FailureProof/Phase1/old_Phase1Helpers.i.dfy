include "../DistributedSystem.i.dfy"

include "../../../Common/Collections/Maps2.i.dfy"
include "../Constants.i.dfy"
include "../Environment.i.dfy"
include "../Replica.i.dfy"

include "TimestampedRslSystem.i.dfy"

include "../CommonProof/Constants.i.dfy"
include "Phase1Invariants.i.dfy"

module RslPhase1Helpers_i {
import opened RslPhase1Invariants_i

lemma Phase1aTimeHelper(node_ts:Timestamp, pkt_ts:Timestamp, ts':Timestamp)
  requires TimeLe(pkt_ts, TimeBound1aDelivery())
  requires TimeLe(node_ts, pkt_ts + MaxQueueTime) // bounded queuing assumption
  requires TimeLe(ts', Rsl_RecvPerfUpdate(node_ts, pkt_ts, RslStep(0)) + D)
  ensures TimeLe(ts', TimeBound1bDelivery())
{
}

lemma Phase1aTimeAlways()
  ensures forall ts, ts' :: TimeLe(ts, TimeBound1aDelivery()) ==> TimeLe(ts', TimeBound1bDelivery())
{
}

}
