include "../../DistributedSystem.i.dfy"

include "../../../../Common/Collections/Maps2.i.dfy"
include "../../Constants.i.dfy"
include "../../Environment.i.dfy"
include "../../Replica.i.dfy"
include "../Common/assumptions.i.dfy"

include "../TimestampedRslSystem.i.dfy"
include "FailureHelpers.i.dfy"
include "FailureSuspectingHB.i.dfy"

include "../../CommonProof/Constants.i.dfy"
// include "../CommonProof/Requests.i.dfy"

module FailureNewView_i {
import opened TimestampedRslSystem_i
import opened FailureHelpers_i
import opened Common_Assumptions
import opened FailureSuspectingHB_i

// In this proof, we can assume bound about suspecting HB and about the step
// that makes a node a suspector.

predicate NewViewInv(s:TimestampedRslState)
{
  // && NoNewViewHBs(s)
  (forall j :: 0 <= j < |s.t_replicas| ==>
  (
   var es := s.t_replicas[j].v.replica.proposer.election_state;
   es.current_view == Ballot(1,0) ==>
   && |es.current_view_suspectors| <= LMinQuorumSize(s.constants.config)
   && (|es.current_view_suspectors| == LMinQuorumSize(s.constants.config) ==> // just about to enter new view
      && s.t_replicas[j].v.nextActionIndex == 8
      && TimeLe(s.t_replicas[j].ts, TBJustBeforeNewView()))
  ))

  && (forall j :: 0 <= j < |s.t_replicas| ==>
  (
   var es := s.t_replicas[j].v.replica.proposer.election_state;
   es.current_view == Ballot(1,1) ==>
    true
    // some time bound about the new-view heartbeatperiod
  ))
}

predicate NewViewSpec(s:TimestampedRslState)
{
  forall pkt ::
    && pkt in s.t_environment.sentPackets
    && pkt.msg.v.RslMessage_Heartbeat?
    && pkt.msg.v.bal_heartbeat == Ballot(1,1)
    ==>
  (exists pkt' ::
    && pkt' in s.t_environment.sentPackets
    && pkt'.msg.v.RslMessage_Heartbeat?
    && pkt'.dst == pkt.dst
    && pkt.msg.v.bal_heartbeat == Ballot(1,1)
    && TimeLe(pkt.msg.ts, TBFirstNewViewHB())
  )
}

}
