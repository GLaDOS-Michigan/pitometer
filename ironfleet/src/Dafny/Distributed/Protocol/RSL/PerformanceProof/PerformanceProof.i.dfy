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
  && NoPacketDuplication(s)
  && |s.t_replicas| > 0

  // TODO: Remove this assumption
  && (var nextStep := s.t_environment.nextStep; nextStep.LEnvStepHostIos? ==>
  (forall io :: io in nextStep.ios && io.LIoOpReceive? ==> !io.r.msg.v.RslMessage_Heartbeat?))
}

predicate RslConsistency(s:TimestampedRslState)
{
  ConstantsAllConsistentInv(UntimestampRslState(s))
    && WellFormedLConfiguration(s.constants.config)
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

predicate LSchedulerTimeBound(tls:TimestampedLScheduler)
{
  0 <= tls.v.nextActionIndex <= 9
  && TimeLe(tls.ts, tls.dts + TimeActionRange(tls.v.nextActionIndex))
}

predicate InitInvariant(s:TimestampedRslState)
  requires RslConsistency(s)
{
  && (0 < |s.t_replicas|
    ==>
    var r := s.t_replicas[0].v.replica.proposer;
    (if s.t_replicas[0].v.nextActionIndex == 0 then
      TimeEq(s.t_replicas[0].ts, TimeZero())
    else
      && s.t_replicas[0].v.nextActionIndex == 1
      && TimeLe(s.t_replicas[0].ts, Timeout() + ProcessPacket)
    )
  )

  && (forall idx :: 0 < idx < |s.t_replicas|
  ==> LSchedulerTimeBound(s.t_replicas[idx])
  )

  && (0 < |s.t_replicas| ==>
  var p := s.t_replicas[0].v.replica.proposer;
    && p.current_state == 0
    && p.election_state.current_view == Ballot(1, 0)
    && p.max_ballot_i_sent_1a == Ballot(0, 0)
  )

  && (forall idx :: 0 <= idx < |s.t_replicas|
  ==> s.t_replicas[idx].v.replica.acceptor.max_bal == Ballot(0, 0)
  && s.t_replicas[idx].v.replica.proposer.current_state == 0
  && s.t_replicas[idx].v.replica.executor.next_op_to_execute == OutstandingOpUnknown()
  && s.t_replicas[idx].v.replica.learner.unexecuted_learner_state == map[]
  )

  && (forall pkt :: pkt in s.undeliveredPackets
  ==> pkt.msg.v.RslMessage_Heartbeat?
  )
}

predicate No1aToDst(undeliveredPackets:multiset<TimestampedLPacket<EndPoint, RslMessage>>, dst:NodeIdentity)
{
  && (forall pkt :: pkt in undeliveredPackets
    && pkt.msg.v.RslMessage_1a?
    ==> pkt.dst != dst)
}

predicate Phase1Invariant(s:TimestampedRslState)
  requires RslConsistency(s)
{
  && (forall idx :: 0 <= idx < |s.t_replicas|
    ==>
  && s.t_replicas[idx].v.replica.executor.next_op_to_execute == OutstandingOpUnknown()
  && s.t_replicas[idx].v.replica.learner.unexecuted_learner_state == map[]
  )

  && (forall idx :: 0 < idx < |s.t_replicas|
  ==>
  && s.t_replicas[idx].v.replica.proposer.current_state == 0
  )

  && (0 < |s.t_replicas| ==>
    && TimeLe(s.t_replicas[0].ts, TimeBound1aSent())
  )
  && (0 < |s.t_replicas| ==>
    && s.t_replicas[0].v.replica.proposer.current_state == 1
  )

  && (forall pkt :: pkt in s.undeliveredPackets
  ==>
   || pkt.msg.v.RslMessage_Heartbeat?
   || pkt.msg.v.RslMessage_1a?
   || pkt.msg.v.RslMessage_1b?
  )

  && (forall pkt :: pkt in s.undeliveredPackets
  && pkt.msg.v.RslMessage_1b?
  ==> pkt.dst == s.constants.config.replica_ids[0]
  // && TimeLe(pkt.msg.ts, TimeBound1bSent())
  )

  && (forall pkt :: pkt in s.undeliveredPackets
      && pkt.msg.v.RslMessage_1a?
   ==> TimeLe(pkt.msg.ts, TimeBound1aSent())
      && pkt.msg.v.bal_1a == Ballot(1, 0)
      && pkt.src == s.constants.config.replica_ids[0]

      && (var undeliveredPackets' := s.undeliveredPackets - multiset({pkt});
          No1aToDst(undeliveredPackets', pkt.dst)
        )
  )

  && (forall idx :: 0 < idx < |s.t_replicas|
  ==>
  if s.t_replicas[idx].v.replica.acceptor.max_bal == Ballot(0, 0) then
    LSchedulerTimeBound(s.t_replicas[idx])
  else
    (forall pkt :: pkt in s.undeliveredPackets
    && pkt.msg.v.RslMessage_1a?
    ==> pkt.dst != s.constants.config.replica_ids[idx]
    // (&& s.t_replicas[idx].v.replica.acceptor.max_bal == s.t_replicas[0].v.replica.proposer.max_ballot_i_sent_1a
    // && true // TimeLe(s.t_replicas[idx].ts, TimeBound1aReceived())
    )
  )
}

lemma lemma_RslInitImpliesInitInv(s:TimestampedRslState)
  requires RslAssumption(s);
  requires TimestampedRslInit(s.constants, s);
  ensures InitInvariant(s);
{
}

function MultisetToSet<S(!new)>(m:multiset<S>) : set<S>
  ensures forall x :: x in m <==> x in MultisetToSet(m)
{
  set x | x in m :: x
}

lemma lemma_0_1_InitGoesToPhase1(s:TimestampedRslState, s':TimestampedRslState, j:int)
  requires RslAssumption(s);
  requires ViewAlwaysZero(s)
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 == j < |s.constants.config.replica_ids|

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep == RslStep(1);

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires InitInvariant(s)
  ensures Phase1Invariant(s')
{
  // var ios := s.t_environment.nextStep.ios;
  // var newPacketsSeq := ExtractSentPacketsFromIos(ios);
  // var newPackets := set x | (x in s'.undeliveredPackets && x !in s.undeliveredPackets) :: x;

  // assume (MultisetToSet(s'.undeliveredPackets) - MultisetToSet(s.undeliveredPackets)) <= (set io | io in ios && io.LIoOpSend? :: io.s);
  // assume newPackets <= (set io | io in ios && io.LIoOpSend? :: io.s);
  // assume forall pkt :: (pkt in s'.undeliveredPackets && pkt !in s.undeliveredPackets) ==> ;
  // assert forall pkt :: pkt in newPackets ==> pkt.msg.v.RslMessage_1a?;
  // D := (A - B) + C
  // (D - A) <= C

}

lemma lemma_InitGoesToInitOrPhase1(s:TimestampedRslState, s':TimestampedRslState, j:int)
  requires RslAssumption(s);
  requires ViewAlwaysZero(s)
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires InitInvariant(s)
  ensures InitInvariant(s') || Phase1Invariant(s')
{
  if (s.t_replicas[0].v.nextActionIndex == 1) {
    if (j == 0) {
        assert Phase1Invariant(s');
    } else {
        assert InitInvariant(s');
    }
  } else {
    assert InitInvariant(s');
  }
}

lemma lemma_0_j_Phase1GoesToPhase1(s:TimestampedRslState, s':TimestampedRslState, j:int)
  requires RslAssumption(s);
  requires ViewAlwaysZero(s)
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|

  requires j > 0;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep == RslStep(0)

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Phase1Invariant(s)
  ensures Phase1Invariant(s')
{
  assert ReplicasDistinct(s.constants.config.replica_ids, 0, j);
  var ios := s.t_environment.nextStep.ios;
  if ios[0].LIoOpReceive? {
    // assert s.t_replicas[j].v.replica.acceptor.max_bal == Ballot(0, 0);
    assert ios[0].r.msg.v.RslMessage_1a?;
    BoundedLagImpliesBoundedProcessingTime(s.t_replicas[j].dts, s.t_replicas[j].ts, ios[0].r.msg.ts, s'.t_replicas[j].ts);
  } else {
    assert Phase1Invariant(s');
  }
}

lemma lemma_a_j_Phase1GoesToPhase1(s:TimestampedRslState, s':TimestampedRslState, j:int)
  requires RslAssumption(s);
  requires ViewAlwaysZero(s)
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|

  requires j > 0;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep != RslStep(0)

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Phase1Invariant(s)
  ensures Phase1Invariant(s')
{
  assert ReplicasDistinct(s.constants.config.replica_ids, 0, j);
}

}
