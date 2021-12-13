include "../../DistributedSystem.i.dfy"

include "../../../../Common/Collections/Maps2.i.dfy"
include "../../Constants.i.dfy"
include "../../Environment.i.dfy"
include "../../Replica.i.dfy"

include "../TimestampedRslSystem.i.dfy"
include "FailureHelpers.i.dfy"
include "FailureDetection_defns.i.dfy"
include "FailureDetection_helper0.i.dfy"

include "../../CommonProof/Constants.i.dfy"
// include "../CommonProof/Requests.i.dfy"

module FailureDetection_helper1_i {
import opened TimestampedRslSystem_i
import opened FailureHelpers_i
import opened FailureDetection_defns_i
import opened FailureDetection_helper0_i

////////////////////////////////////////////////////////////////////////////////
// NonSuspector0 inductive
////////////////////////////////////////////////////////////////////////////////

lemma NonSuspector0_ind_recv(s:TimestampedRslState, s':TimestampedRslState, sr:set<int>, j:int)
  requires FOAssumption2(s, s')
  requires EpochDelayInv(s)
  requires EpochDelayInv(s')
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);
  requires s.t_environment.nextStep.nodeStep == RslStep(0);

  requires InView1(s, sr);
  requires j !in sr;
  requires NonSuspector0(s, j);

  ensures  NonSuspector0(s', j) || NonSuspector1(s', j);
  ensures NotKnownSuspector(s', j);
{
  var ios := s.t_environment.nextStep.ios;

  if ios[0].LIoOpReceive?  {
    if ios[0].r.msg.v.RslMessage_Request? {
      var es := s.t_replicas[j].v.replica.proposer.election_state;
      var newReq := Request(ios[0].r.src, ios[0].r.msg.v.seqno_req, ios[0].r.msg.v.val);
      assert ios[0].r in s.t_environment.sentPackets;
      assert ios[0].r.src == req.client;
      assert ios[0].r.msg.v.seqno_req == req.seqno;
      // assert (req in es.requests_received_this_epoch) &&
      assert RequestsMatch(req, newReq);

      if s'.t_replicas[j].v.replica.proposer.election_state.requests_received_this_epoch == [req] {
        reveal_TBEpoch1();
        assert NonSuspector1(s', j);
        return;
      }
    }
  }

  assert NonSuspector0(s', j); // maybe 1
}


lemma NonSuspector0_ind(s:TimestampedRslState, s':TimestampedRslState, sr:set<int>, j:int)
  requires FOAssumption2(s, s')
  requires EpochDelayInv(s)
  requires EpochDelayInv(s')
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires InView1(s, sr);
  requires j !in sr;
  requires NonSuspector0(s, j);

  ensures  NonSuspector0(s', j) || NonSuspector1(s', j);
  ensures NotKnownSuspector(s', j);
{
  var step := s.t_environment.nextStep.nodeStep;
  if step == RslStep(0) {
    NonSuspector0_ind_recv(s, s', sr, j);
  } else if step == RslStep(6) {
    if s.t_replicas[j].v.replica.executor.next_op_to_execute.OutstandingOpKnown? {
      var es := s.t_replicas[j].v.replica.proposer.election_state;
      var es' := s'.t_replicas[j].v.replica.proposer.election_state;
      var batch := s.t_replicas[j].v.replica.executor.next_op_to_execute.v;
      if ElectionStateReflectExecutedRequestBatch(es, es', batch) {
        lemma_RemoveExecutedRequestBatchProducesSubsequence(
          es'.requests_received_this_epoch,
          es.requests_received_this_epoch,
          batch);
        lemma_RemoveExecutedRequestBatchProducesSubsequence(
          es'.requests_received_prev_epochs,
          es.requests_received_prev_epochs,
          batch);

        assert es'.requests_received_prev_epochs == [];
        assert es'.requests_received_this_epoch == [];
        assert NonSuspector0(s', j);
      }
    }
    assert NonSuspector0(s', j);
  } else {
    assert NonSuspector0(s', j);
  }
}

////////////////////////////////////////////////////////////////////////////////
// NonSuspector2 inductive
////////////////////////////////////////////////////////////////////////////////

lemma NonSuspector2_ind_7(s:TimestampedRslState, s':TimestampedRslState, sr:set<int>, j:int) returns (sr':set<int>)
  requires FOAssumption2(s, s')

  requires EpochTimeoutQDInv(s)
  requires EpochTimeoutQDInv(s')

  requires HeartbeatDelayInv(s)
  requires HeartbeatDelayInv(s')

  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);
  requires s.t_environment.nextStep.nodeStep == RslStep(7);

  requires InView1(s, sr);
  requires j !in sr;
  requires NonSuspector2(s, j);

  ensures
    (sr' == sr && NotKnownSuspector(s', j) &&
       (NonSuspector2(s', j) || NonSuspector0(s', j) || InternalSuspector3(s', j)))
    ||
    (sr' == sr + {j} && Suspector(s', j));
{
  sr' := sr;
  var ios := s.t_environment.nextStep.ios;
  var clock := SpontaneousClock(UntagLIoOpSeq(ios));
  var es := s.t_replicas[j].v.replica.proposer.election_state;
  if clock.t < es.epoch_end_time {
    assert NotKnownSuspector(s', j);
    assert NonSuspector2(s', j);
  } else {
    if j == 1 {
      // assert NotKnownSuspector(s', j); // might not be true
      // assert InternalSuspector3(s', j);
      assert Suspector(s', j);
      sr' := sr + {j};
    } else {
      assert NotKnownSuspector(s', j);

      reveal_HBPeriodEnd();
      assert InternalSuspector3(s', j);
    }
  }
}

lemma NonSuspector2_ind(s:TimestampedRslState, s':TimestampedRslState, sr:set<int>, j:int) returns (sr':set<int>)
  requires FOAssumption2(s, s')
  requires EpochTimeoutQDInv(s)
  requires EpochTimeoutQDInv(s')

  requires HeartbeatDelayInv(s)
  requires HeartbeatDelayInv(s')

  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires InView1(s, sr);
  requires j !in sr;
  requires NonSuspector2(s, j);
  requires LeaderView0(s')

  // If j is added to sr, then j must be the leader.
  ensures
    (sr' == sr && NotKnownSuspector(s', j) && (NonSuspector2(s', j) || NonSuspector0(s', j)
     || InternalSuspector3(s', j)))
    || (sr' == sr + {j} && Suspector(s', j))
    ;

  ensures SuspectingReplicaInv(s', sr');
{
  sr' := sr;

  var step := s.t_environment.nextStep.nodeStep;
  if step == RslStep(6) {
    assert NotKnownSuspector(s', j);
    if s.t_replicas[j].v.replica.executor.next_op_to_execute.OutstandingOpKnown? {
      var es := s.t_replicas[j].v.replica.proposer.election_state;
      var es' := s'.t_replicas[j].v.replica.proposer.election_state;
      var batch := s.t_replicas[j].v.replica.executor.next_op_to_execute.v;
      if ElectionStateReflectExecutedRequestBatch(es, es', batch) {
        lemma_RemoveExecutedRequestBatchProducesSubsequence(
          es'.requests_received_this_epoch,
          es.requests_received_this_epoch,
          batch);
        lemma_RemoveExecutedRequestBatchProducesSubsequence(
          es'.requests_received_prev_epochs,
          es.requests_received_prev_epochs,
          batch);

        if es'.requests_received_prev_epochs == [] {
          assert es'.requests_received_this_epoch == [];
          assert NonSuspector0(s', j);
          return;
        } else {
          assert NonSuspector2(s', j);
        }
      }
    }
    assert NonSuspector2(s', j);
  } else if step == RslStep(7) {
    sr' := NonSuspector2_ind_7(s, s', sr, j);
  } else if step == RslStep(8) {
    var r := s.t_replicas[j].v.replica;
    SubsetCardinality(r.proposer.election_state.current_view_suspectors, sr);
    assert NonSuspector2(s', j);
    assert NotKnownSuspector(s', j);
  } else {
    assert NonSuspector2(s', j);
    assert NotKnownSuspector(s', j);
  }
}

////////////////////////////////////////////////////////////////////////////////
// InternalSuspector3 inductive
////////////////////////////////////////////////////////////////////////////////

lemma InternalSuspector3_ind(s:TimestampedRslState, s':TimestampedRslState, sr:set<int>, j:int) returns (sr':set<int>)
  requires FOAssumption2(s, s')

  requires EpochTimeoutQDInv(s)
  requires EpochTimeoutQDInv(s')

  requires HeartbeatQDInv(s)
  requires HeartbeatQDInv(s')

  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires InView1(s, sr);
  requires j !in sr;
  requires InternalSuspector3(s, j);

  // might become a suspector here
  ensures
    (sr' == sr && NotKnownSuspector(s', j) && InternalSuspector3(s', j))
    || (sr' == sr + {j} && Suspector(s', j))
    ;
  ensures SuspectingReplicaInv(s', sr');
{
  sr' := sr;
  if s.t_environment.nextStep.nodeStep != RslStep(9) {
    assert InternalSuspector3(s', j);
    assert SuspectingReplicaInv(s', sr');
  } else {
    var ios := s.t_environment.nextStep.ios;
    var clock := SpontaneousClock(UntagLIoOpSeq(ios));
    if clock.t < s.t_replicas[j].v.replica.nextHeartbeatTime {
      assert InternalSuspector3(s', j);
      assert SuspectingReplicaInv(s', sr');
    } else {
      // sentPackets' := s'.t_environment.sentPackets;
      var sent_packets := ExtractSentPacketsFromIos(UntagLIoOpSeq(ios));

      assert exists pkt :: pkt in sent_packets
        && pkt.msg.RslMessage_Heartbeat?
        && pkt.src == s.constants.config.replica_ids[j]
        && pkt.dst == s.constants.config.replica_ids[1]
        && pkt.msg.suspicious == true;

      assert exists pkt :: LIoOpSend(pkt) in ios
        && pkt.msg.v.RslMessage_Heartbeat?
        && pkt.src == s.constants.config.replica_ids[j]
        && pkt.dst == s.constants.config.replica_ids[1]
        && pkt.msg.v.suspicious == true;

      // assert forall t' :: TimeLe(t', s'.t_replicas[j].ts + D) ==> TimeLe(t', TBFirstSuspectingHB());
      HeartbeatQDHelper(s.t_replicas[j].ts, s'.t_replicas[j].ts);
      // s.t_replicas[j].v.replica.nextHeartbeatTime);

      sr' := sr + {j};
      assert Suspector(s', j);
      // all sent packets are bounded by TBFirstSuspectingHB
    }
  }
}


////////////////////////////////////////////////////////////////////////////////
// Suspector inductive
////////////////////////////////////////////////////////////////////////////////

// self-step
lemma Suspector_ind_self(s:TimestampedRslState, s':TimestampedRslState, sr:set<int>, j:int)
  requires FOAssumption2(s, s')
  requires EpochTimeoutQDInv(s)
  requires EpochTimeoutQDInv(s')
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires InView1(s, sr);
  requires j in sr;
  requires Suspector(s, j);
  requires LeaderView0(s');

  ensures Suspector(s', j);
{
  if j != 1 {
    assert ReplicasDistinct(s.constants.config.replica_ids, j, 1);
  }
  SubsetCardinality(s.t_replicas[j].v.replica.proposer.election_state.current_view_suspectors, sr);
}

// leader-step
lemma Suspector_ind_leader(s:TimestampedRslState, s':TimestampedRslState, sr:set<int>, j:int)
  requires FOAssumption2(s, s')
  requires EpochTimeoutQDInv(s)
  requires EpochTimeoutQDInv(s')
  requires 0 <= j < |s.constants.config.replica_ids|;
  requires j != 1;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[1];

  requires TimestampedRslNextOneReplica(s, s', 1, s.t_environment.nextStep.ios);

  requires InView1(s, sr);
  requires j in sr;
  requires Suspector(s, j);

  requires LeaderView0(s');

  ensures Suspector(s', j) || FinalStage(s);
{
  assert ReplicasDistinct(s.constants.config.replica_ids, j, 1);
  if s.t_environment.nextStep.nodeStep == RslStep(0) {
    if (s.t_replicas[j].v.replica.constants.my_index in s.t_replicas[1].v.replica.proposer.election_state.current_view_suspectors) {
      assert Suspector(s', j);
    } else {
      var pkt :|
        && pkt in s.undeliveredPackets
        && pkt.msg.v.RslMessage_Heartbeat?
        && pkt.src == s.constants.config.replica_ids[j]
        && pkt.dst == s.constants.config.replica_ids[1]
        && pkt.msg.v.suspicious == true
        && TimeLe(pkt.msg.ts, TBFirstSuspectingHB());

      var t_ios := s.t_environment.nextStep.ios;
      var ios := UntagLIoOpSeq(t_ios);

      assert forall k :: 0 < k < |ios| ==> !ios[k].LIoOpReceive?;

      if t_ios[0].LIoOpReceive? && t_ios[0].r == pkt {
        lemma_GetReplicaIndexIsUnique(s.constants.config, j);
        assert s'.t_replicas[1].v.replica.proposer.election_state.current_view_suspectors ==
          s.t_replicas[1].v.replica.proposer.election_state.current_view_suspectors + {j};
        assert s'.t_replicas[j].v.replica.constants.my_index == j;
        assert (s'.t_replicas[j].v.replica.constants.my_index in s'.t_replicas[1].v.replica.proposer.election_state.current_view_suspectors);
        return;
      }
      assert pkt !in (set io | io in t_ios && io.LIoOpReceive? :: io.r);
      assert pkt in s'.undeliveredPackets;
      // might receive packet and remove it from undeliveredPackets
      assert Suspector(s', j);
    }
  } else if s.t_environment.nextStep.nodeStep == RslStep(8) {
    if s'.t_replicas[1].v.replica.proposer.election_state.current_view == Ballot(1,1) {
      SubsetCardinality(s.t_replicas[1].v.replica.proposer.election_state.current_view_suspectors, sr);
      assert Suspector(s', j);
    } else {
      assert Suspector(s', j);
    }
  } else {
    assert Suspector(s', j);
  }
}

}
