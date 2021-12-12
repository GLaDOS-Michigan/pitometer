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
  // requires RslConsistency(s)
  requires CommonAssumptions(s)
{
  // TODO: add stuff about packets

  && s.t_replicas[1].v.replica.proposer.election_state.current_view == Ballot(1,1)
  && s.t_replicas[1].v.nextActionIndex == 9 // just checked for quorum of views, and got one
  && TimeLe(s.t_replicas[1].ts, FailoverTime())
}

predicate LeaderNotInNewView(s:TimestampedRslState)
  requires CommonAssumptions(s)
{
  && s.t_replicas[1].v.replica.proposer.election_state.current_view == Ballot(1,0)
}

predicate LeaderInNewViewNonFinal(s:TimestampedRslState)
  requires CommonAssumptions(s)
{
  && s.t_replicas[1].v.replica.proposer.election_state.current_view == Ballot(1,1)
  && s.t_replicas[1].v.nextActionIndex == 1 // become leader from a new-view heartbeat
  // TODO: maintain time bound until action 9?
}

////////////////////////////////////////////////////////////////////////////////
// Proof
////////////////////////////////////////////////////////////////////////////////
lemma NonSuspector1_ind_leader_trivial(s:TimestampedRslState, s':TimestampedRslState)
  requires FOAssumption2(s, s')

  requires NewViewSpec(s)
  requires NewViewSpec(s')
  requires SusHBSpec(s)
  requires SusHBSpec(s')
  requires BecomeSuspectorStepSpec(s, s')

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[1];

  requires s.t_environment.nextStep.nodeStep != RslStep(0) // on step 0, we might recv new-view HB
  requires s.t_environment.nextStep.nodeStep != RslStep(8) // might transition to new view ourselves

  requires TimestampedRslNextOneReplica(s, s', 1, s.t_environment.nextStep.ios);

  requires LeaderNotInNewView(s)
  ensures  LeaderNotInNewView(s')
{
}

lemma LeaderNewView_ind_leader_trivial(s:TimestampedRslState, s':TimestampedRslState)
  requires FOAssumption2(s, s')

  requires NewViewSpec(s)
  requires NewViewSpec(s')
  requires SusHBSpec(s)
  requires SusHBSpec(s')
  requires BecomeSuspectorStepSpec(s, s')

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[1];

  requires s.t_environment.nextStep.nodeStep != RslStep(0) // on step 0, we might recv new-view HB
  requires s.t_environment.nextStep.nodeStep != RslStep(8) // might transition to new view ourselves

  requires TimestampedRslNextOneReplica(s, s', 1, s.t_environment.nextStep.ios);

  requires LeaderNotInNewView(s)
  ensures  LeaderNotInNewView(s')
{
}

lemma LeaderNewView_ind_leader_recv(s:TimestampedRslState, s':TimestampedRslState)
  requires FOAssumption2(s, s')

  requires NewViewSpec(s)
  requires NewViewSpec(s')
  requires SusHBSpec(s)
  requires SusHBSpec(s')
  requires BecomeSuspectorStepSpec(s, s')

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[1];

  requires s.t_environment.nextStep.nodeStep == RslStep(0) // on step 0, we might recv new-view HB

  requires TimestampedRslNextOneReplica(s, s', 1, s.t_environment.nextStep.ios);

  requires LeaderNotInNewView(s)
  ensures  LeaderNotInNewView(s') || LeaderInNewViewNonFinal(s')
{
  var ios := s.t_environment.nextStep.ios;

  if ios[0].LIoOpReceive?  {
    if ios[0].r.msg.v.RslMessage_Heartbeat? {
      var m := ios[0].r.msg.v;
      if m.bal_heartbeat == Ballot(1,1) {
        assert LeaderInNewViewNonFinal(s');
      } else {
        assert m.bal_heartbeat == Ballot(1,0); // FIXME: prove or assume this
        assert LeaderNotInNewView(s');
      }
    } else {
      assert LeaderNotInNewView(s');
    }
  } else {
    assert LeaderNotInNewView(s');
  }
}

/*
   if s.t_environment.nextStep.nodeStep == RslStep(0) {
   assert LeaderNotInNewView(s');
   }
   else if s.t_environment.nextStep.nodeStep == RslStep(1) {
   assert LeaderNotInNewView(s');
   }
   else if s.t_environment.nextStep.nodeStep == RslStep(2) {
   assert LeaderNotInNewView(s');
   }
   else if s.t_environment.nextStep.nodeStep == RslStep(3) {
   assert LeaderNotInNewView(s');
   }
   else if s.t_environment.nextStep.nodeStep == RslStep(4) {
   assert LeaderNotInNewView(s');
   }
   else if s.t_environment.nextStep.nodeStep == RslStep(5) {
   assert LeaderNotInNewView(s');
   }
   else if s.t_environment.nextStep.nodeStep == RslStep(6) {
   assert LeaderNotInNewView(s');
   }
   else if s.t_environment.nextStep.nodeStep == RslStep(7) {
   assert LeaderNotInNewView(s');
   }
   else if s.t_environment.nextStep.nodeStep == RslStep(8) {
   assert LeaderNotInNewView(s');
   }
   else if s.t_environment.nextStep.nodeStep == RslStep(9) {
    assert LeaderNotInNewView(s');
  }
*/
}
