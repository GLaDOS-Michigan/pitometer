include "../../DistributedSystem.i.dfy"

include "../../../../Common/Collections/Maps2.i.dfy"
include "../../Constants.i.dfy"
include "../../Environment.i.dfy"
include "../../Replica.i.dfy"

include "../TimestampedRslSystem.i.dfy"

include "FailureHelpers.i.dfy"
include "FailureDetection_defns.i.dfy"
include "FailureDetection_helper0.i.dfy"
include "FailureDetection_helper1.i.dfy"

include "../../CommonProof/Constants.i.dfy"
module FailureDetection_i {
import opened TimestampedRslSystem_i
import opened FailureHelpers_i
import opened FailureDetection_defns_i
import opened FailureDetection_helper0_i
import opened FailureDetection_helper1_i

function {:opaque} FailoverTime() : Timestamp
{
  0
}

predicate FailoverFinal(s:TimestampedRslState)
{
  && (forall pkt ::
     pkt in s.t_environment.sentPackets ==>

     && (pkt.msg.v.RslMessage_Heartbeat? ==>
     pkt.msg.v.bal_heartbeat == Ballot(1, 0))
     && (pkt.msg.v.RslMessage_2a? ==>
     pkt.msg.v.bal_2a == Ballot(1, 0))
     && (pkt.msg.v.RslMessage_2b? ==>
     pkt.msg.v.bal_2b == Ballot(1, 0))
     && (pkt.msg.v.RslMessage_1a? ==>
     pkt.msg.v.bal_1a == Ballot(1, 0))
     && (pkt.msg.v.RslMessage_1b? ==>
     pkt.msg.v.bal_1b == Ballot(1, 0))
  )

  && (
    forall j :: 0 <= j < |s.t_replicas| ==>
    s.t_replicas[j].v.replica.proposer.election_state.current_view == Ballot(1,0)
    || s.t_replicas[j].v.replica.proposer.election_state.current_view == Ballot(1,1)
    )

  && 1 < |s.t_replicas| // this is actually part of an assumption of the failure proof
  && s.t_replicas[1].v.replica.proposer.election_state.current_view == Ballot(1,1)
  && s.t_replicas[1].v.nextActionIndex == 9 // just checked for quorum of views, and got one
  && TimeLe(s.t_replicas[1].ts, FailoverTime())
}

lemma FailoverTopLevel(tglb:seq<TimestampedRslState>) returns (startPhase1Idx:int)
  requires exists con :: ValidTimestampedRSLBehavior(con, tglb);
  requires forall s :: s in tglb ==> FailureDetection_defns_i.RslConsistency(s);
  // FIXME: also the 2-step assumptions have to be met

  ensures forall j :: 0 <= j < |tglb| ==> j <= startPhase1Idx ==>
    var s := tglb[j];
    && CurrView(s) // This means every node is in view (1,0); will rename to be more clear
    && InView1Packets(s); // All packets have Ballot(1,0)

  // The idx might be larger than the seq if the behavior given never has
  // failover happen, and never reaches failover final
  ensures startPhase1Idx >= 0
  ensures startPhase1Idx < |tglb| ==> FailoverFinal(tglb[startPhase1Idx])
{
  // TODO:
}
}
