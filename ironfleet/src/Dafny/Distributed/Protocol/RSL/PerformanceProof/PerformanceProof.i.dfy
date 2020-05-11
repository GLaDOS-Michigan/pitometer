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

predicate {:verify false} RslAssumption(s:TimestampedRslState)
{
  |s.t_replicas| > 0
}

predicate {:verify false} RslConsistency(s:TimestampedRslState)
{
  ConstantsAllConsistentInv(UntimestampRslState(s))
}

predicate {:verify false} ViewAlwaysZero(s:TimestampedRslState)
{
  && (forall idx:: 0 <= idx < |s.t_replicas|
    ==>
    s.t_replicas[idx].v.replica.proposer.election_state.current_view.proposer_id == 0
    )
}

predicate {:verify false} AlwaysInvariant(s:TimestampedRslState)
{
  && ViewAlwaysZero(s)
}

predicate {:verify false} InitInvariant(s:TimestampedRslState)
  requires RslConsistency(s)
{
  && (0 < |s.t_replicas|
    ==>
    var r := s.t_replicas[0].v.replica.proposer;
    if s.t_replicas[0].v.nextActionIndex == 0 then
      TimeEq(s.t_replicas[0].ts, TimeZero())
    else
      && s.t_replicas[0].v.nextActionIndex == 1
      && TimeLe(s.t_replicas[0].ts, Timeout() + L)
    )

  && (0 < |s.t_replicas| ==>
  var p := s.t_replicas[0].v.replica.proposer;
    && p.current_state == 0
    && BalLt(p.max_ballot_i_sent_1a, p.election_state.current_view)
  )

  && (forall idx :: 0 <= idx < |s.t_replicas|
  ==> s.t_replicas[idx].v.replica.acceptor.max_bal == Ballot(0, 0)
  && s.t_replicas[idx].v.replica.proposer.current_state == 0
  && s.t_replicas[idx].v.replica.executor.next_op_to_execute == OutstandingOpUnknown()
  && s.t_replicas[idx].v.replica.learner.unexecuted_learner_state == map[]
  )

  && (forall pkt :: pkt in s.t_environment.sentPackets
  ==> pkt.msg.v.RslMessage_Heartbeat?
  )
}

predicate {:verify false} Phase1aInvariant(s:TimestampedRslState)
  requires RslConsistency(s)
{
  && (forall idx :: 0 <= idx < |s.t_replicas|
    ==>
  && s.t_replicas[idx].v.replica.executor.next_op_to_execute == OutstandingOpUnknown()
  && s.t_replicas[idx].v.replica.learner.unexecuted_learner_state == map[]
  )

  && (0 < |s.t_replicas| ==>
    && TimeLe(s.t_replicas[0].ts, TimeBound1aSent())
  )
  && (0 < |s.t_replicas| ==>
    && s.t_replicas[0].v.replica.proposer.current_state == 1
  )

  && (forall pkt :: pkt in s.t_environment.sentPackets
  ==>
   || pkt.msg.v.RslMessage_Heartbeat?
   || pkt.msg.v.RslMessage_1a?
  )

  && (forall pkt :: pkt in s.t_environment.sentPackets
      && pkt.msg.v.RslMessage_1a?
  ==> TimeLe(pkt.msg.ts, TimeBound1aSent())
  )

  && (forall idx :: 0 < idx < |s.t_replicas|
  ==>
  if s.t_replicas[idx].v.replica.acceptor.max_bal == Ballot(0, 0) then
    // TODO: s.t_replicas[idx].ts == TimeZero() 
    true
  else
    (&& s.t_replicas[idx].v.replica.acceptor.max_bal == s.t_replicas[0].v.replica.proposer.max_ballot_i_sent_1a
    && s.t_replicas[idx].ts == TimeBound1aReceived())
  )
}

lemma {:verify false} NextProcessPacket_0_InitGoesToInit(s:TimestampedRslState, s':TimestampedRslState)
  requires RslAssumption(s);
  requires ViewAlwaysZero(s)
  requires RslConsistency(s);
  requires RslConsistency(s');

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[0];
  requires s.t_environment.nextStep.nodeStep == RslStep(0);

  requires TimestampedRslNextOneReplica(s, s', 0, s.t_environment.nextStep.ios);

  requires InitInvariant(s)
  ensures InitInvariant(s')
{
}

lemma {:verify false} NextProcessPacket_j_InitGoesToInit(s:TimestampedRslState, s':TimestampedRslState, j:int)
  requires RslAssumption(s);
  requires ViewAlwaysZero(s)
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|

  requires j > 0;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires InitInvariant(s)
  ensures InitInvariant(s')
{
}

lemma MaybeEnterNewView_1_InitGoesToPhase1a(s:TimestampedRslState, s':TimestampedRslState)
  requires RslAssumption(s);
  requires ViewAlwaysZero(s)
  requires RslConsistency(s);
  requires RslConsistency(s');

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[0];
  requires s.t_environment.nextStep.nodeStep == RslStep(1);

  requires TimestampedRslNextOneReplica(s, s', 0, s.t_environment.nextStep.ios);

  requires InitInvariant(s)
  ensures Phase1aInvariant(s')
{
}

}
