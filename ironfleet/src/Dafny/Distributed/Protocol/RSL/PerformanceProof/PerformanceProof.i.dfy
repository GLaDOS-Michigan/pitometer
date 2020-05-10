include "../DistributedSystem.i.dfy"

include "../../../Common/Collections/Maps2.i.dfy"
include "../Constants.i.dfy"
include "../Environment.i.dfy"
include "../Replica.i.dfy"

include "TimestampedRslSystem.i.dfy"

include "../CommonProof/Constants.i.dfy"

module RslPerformanceProof_i {
import opened TimestampedRslPerformanceProof_i
import opened CommonProof__Constants_i

predicate RslAssumption(s:TimestampedRslState)
{
  |s.t_replicas| > 0
}

predicate RslConsistency(s:TimestampedRslState)
{
  ConstantsAllConsistentInv(UntimestampRslState(s))
}

predicate ViewAlwaysZero(s:TimestampedRslState)
{
  && (forall idx:: 0 <= idx < |s.t_replicas|
    ==>
    s.t_replicas[idx].v.replica.proposer.election_state.current_view.proposer_id == 0
    )
}

predicate AlwaysInvariant(s:TimestampedRslState)
{
  && ViewAlwaysZero(s)
}

predicate InitInvariant(s:TimestampedRslState)
  requires RslConsistency(s)
{
  && (forall idx:: 0 <= idx < |s.t_replicas|
    ==> 
    TimeEq(s.t_replicas[idx].ts, TimeZero())
    )

  && (0 < |s.t_replicas| ==>
  var p := s.t_replicas[0].v.replica.proposer;
    && p.current_state == 0
    && BalLt(p.max_ballot_i_sent_1a, p.election_state.current_view)
  )
}

predicate Phase1Invariant(s:TimestampedRslState)
  requires RslConsistency(s)
{
  && (0 < |s.t_replicas| ==>
    && TimeEq(s.t_replicas[0].ts, L)
  )
  && (0 < |s.t_replicas| ==>
    && s.t_replicas[0].v.replica.proposer.current_state == 1
  )
}

lemma Init(s:TimestampedRslState, s':TimestampedRslState)
  requires RslAssumption(s);
  requires ViewAlwaysZero(s)
  requires RslConsistency(s);
  requires RslConsistency(s');

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[0];
  requires s.t_environment.nextStep.nodeStep == RslStep(1);

  requires TimestampedRslNextOneReplica(s, s', 0, s.t_environment.nextStep.ios);

  requires InitInvariant(s)
  ensures Phase1Invariant(s')
{
}

}
