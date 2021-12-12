include "../../DistributedSystem.i.dfy"

include "../../../../Common/Collections/Maps2.i.dfy"
include "../../Constants.i.dfy"
include "../../Environment.i.dfy"
include "../../Replica.i.dfy"
include "../Common/assumptions.i.dfy"

include "../TimestampedRslSystem.i.dfy"
include "FailureHelpers.i.dfy"

include "../../CommonProof/Constants.i.dfy"
// include "../CommonProof/Requests.i.dfy"

module FailureDetection_defns_i {
import opened TimestampedRslSystem_i
import opened FailureHelpers_i
import opened Common_Assumptions

/* Starting from a B(1,1)-HB-free state, this takes a step into either a
   B(1,1)-HB-free state or a state with B(1,1) HBs broadcast.
*/

predicate Suspector(s:TimestampedRslState, j:int)
  requires RslConsistency(s)
  requires 0 <= j < |s.t_replicas|
{
  forall k :: 0 <= k < |s.t_replicas| ==>
  // There's either a suspecting packet to node k, or node k already knows.
  // At this point, if there's one such packet, then if a different one is
  // received first, we can use the bound of this one to bound that one.
  (exists pkt ::
  && pkt in s.t_environment.sentPackets
  && pkt.msg.v.RslMessage_Heartbeat?
  && pkt.src == s.constants.config.replica_ids[j]
  && pkt.dst == s.constants.config.replica_ids[k]
  && pkt.msg.v.suspicious == true
  && TimeLe(pkt.msg.ts, TBFirstSuspectingHB())
  )
  || (s.t_replicas[j].v.replica.constants.my_index in s.t_replicas[k].v.replica.proposer.election_state.current_view_suspectors)
}

predicate Suspector(s:TimestampedRslState, j:int)
  requires RslConsistency(s)
  requires 0 <= j < |s.t_replicas|
{
  var suspectors := s.t_replicas[k].v.replica.proposer.election_state.current_view_suspectors;
  |suspectors|
    // q :=
}

predicate NodeLocal(s:TimestampedRslState, j:int)
  requires RslConsistency(s)
  requires 0 <= j < |s.t_replicas|
{
  NonSuspector0(s,j) || NonSuspector1(s,j) || NonSuspector2(s, j) || InternalSuspector2(s, j)
  || Suspector(s, j)
}

predicate HBFree(s:TimestampedRslState)
  requires RslConsistency(s)
{
  && InView1Packets(s)
  && (
    forall j :: 0 <= j < |s.t_replicas| ==> InView1Local(s, j)
  )
}

predicate HBSent(s:TimestampedRslState)
{
  && InView1Packets(s)
  && (
    forall j :: 0 <= j < |s.t_replicas| ==> InView1Local(s, j)
  )
}

}
