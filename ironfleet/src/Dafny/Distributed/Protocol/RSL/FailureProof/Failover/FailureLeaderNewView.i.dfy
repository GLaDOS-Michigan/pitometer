include "../../DistributedSystem.i.dfy"

include "../../../../Common/Collections/Maps2.i.dfy"
include "../../Constants.i.dfy"
include "../../Environment.i.dfy"
include "../../Replica.i.dfy"
include "../Common/assumptions.i.dfy"

include "../TimestampedRslSystem.i.dfy"
include "FailureHelpers.i.dfy"
include "FailureNewView.i.dfy"
include "FailureDetection_defns.i.dfy"

include "../../CommonProof/Constants.i.dfy"
// include "../CommonProof/Requests.i.dfy"

module FailureLeaderNewView_i {
import opened TimestampedRslSystem_i
import opened FailureHelpers_i
import opened FailureNewView_i
import opened Common_Assumptions
import opened FailureDetection_defns_i

// can assume NewViewSpec() in this proof

////////////////////////////////////////////////////////////////////////////////
// Invariants
////////////////////////////////////////////////////////////////////////////////

predicate FailoverFinal(s:TimestampedRslState)
{
  // TODO: add stuff about packets

  && s.t_replicas[1].v.replica.proposer.election_state.current_view == Ballot(1,1)
  && s.t_replicas[1].v.nextActionIndex == 9 // just checked for quorum of views, and got one
  && TimeLe(s.t_replicas[1].ts, FailoverTime())
}

predicate LeaderNotInNewView(s:TimestampedRslState)
{
  && s.t_replicas[1].v.replica.proposer.election_state.current_view == Ballot(1,0)
}

predicate LeaderInNewViewNonFinal(s:TimestampedRslState)
{
  && s.t_replicas[1].v.replica.proposer.election_state.current_view == Ballot(1,1)
  && s.t_replicas[1].v.nextActionIndex == 1 // become leader from a new-view heartbeat
  // TODO: maintain time bound until action 9?
}

////////////////////////////////////////////////////////////////////////////////
// Proof
////////////////////////////////////////////////////////////////////////////////
lemma NonSuspector1_ind_leader(s:TimestampedRslState, s':TimestampedRslState)
  requires FOAssumption2(s, s')

  requires NewViewSpec(s)
  requires NewViewSpec(s')
  requires SusHBSpec(s)
  requires SusHBSpec(s')
  requires BecomeSuspectorStepSpec(s, s')

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[1];

  // requires s.t_environment.nextStep.nodeStep != RslStep(7) // on step 7, we might start new epoch
  // requires s.t_environment.nextStep.nodeStep != RslStep(6) // on step 6, we might go to NS0(j)

  requires TimestampedRslNextOneReplica(s, s', 1, s.t_environment.nextStep.ios);

  requires LeaderNotInNewView(s)
  ensures  LeaderNotInNewView(s) || LeaderInNewViewNonFinal(s) || FailoverFinal(s);
{
  if s.t_environment.nextStep.nodeStep == RslStep(0) {
    assert LeaderNotInNewView(s);
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(1) {
    assert LeaderNotInNewView(s);
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(2) {
    assert LeaderNotInNewView(s);
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(3) {
    assert LeaderNotInNewView(s);
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(4) {
    assert LeaderNotInNewView(s);
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(5) {
    assert LeaderNotInNewView(s);
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(6) {
    assert LeaderNotInNewView(s);
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(7) {
    assert LeaderNotInNewView(s);
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(8) {
    assert LeaderNotInNewView(s);
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(9) {
    assert LeaderNotInNewView(s);
  }
}

}
