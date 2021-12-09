include "../DistributedSystem.i.dfy"

include "../../../Common/Collections/Maps2.i.dfy"
include "../Constants.i.dfy"
include "../Environment.i.dfy"
include "../Replica.i.dfy"

include "TimestampedRslSystem.i.dfy"
include "FailureHelpers.i.dfy"
include "FailureDetection_defns.i.dfy"
include "FailureDetection_helper0.i.dfy"

include "../CommonProof/Constants.i.dfy"
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
  requires RslAssumption2(s, s')
  requires EpochTimeoutQDInv(s)
  requires EpochTimeoutQDInv(s')
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
        assert NonSuspector1(s', j);
        // FIXME: prove TimeLe using EpochQD
        return;
      }
    }
  }

  assert NonSuspector0(s', j); // maybe 1
}


lemma NonSuspector0_ind(s:TimestampedRslState, s':TimestampedRslState, sr:set<int>, j:int)
  requires RslAssumption2(s, s')
  requires EpochTimeoutQDInv(s)
  requires EpochTimeoutQDInv(s')
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

lemma NonSuspector2_ind_7(s:TimestampedRslState, s':TimestampedRslState, sr:set<int>, j:int)
  requires RslAssumption2(s, s')
  requires EpochTimeoutQDInv(s)
  requires EpochTimeoutQDInv(s')
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);
  requires s.t_environment.nextStep.nodeStep == RslStep(7);

  requires InView1(s, sr);
  requires j !in sr;
  requires NonSuspector2(s, j);

  ensures  NonSuspector2(s', j) || NonSuspector0(s', j) || InternalSuspector3(s', j);
  ensures NotKnownSuspector(s', j);
{
  var ios := s.t_environment.nextStep.ios;
  var clock := SpontaneousClock(UntagLIoOpSeq(ios));
  var es := s.t_replicas[j].v.replica.proposer.election_state;
  if clock.t < es.epoch_end_time {
    assert NotKnownSuspector(s', j);
    assert NonSuspector2(s', j);
  } else {
    if j == 1 {
      assert NotKnownSuspector(s', j); // might not be true
      assert InternalSuspector3(s', j);
      // FIXME: might go straight to InFinalState(), if j == 1 is the leader
    } else {
      assert NotKnownSuspector(s', j); // might not be true
      assert InternalSuspector3(s', j);
      // FIXME: prove TimeLe() using Heartbeat lemmas
    }
  }
}

lemma NonSuspector2_ind(s:TimestampedRslState, s':TimestampedRslState, sr:set<int>, j:int)
  requires RslAssumption2(s, s')
  requires EpochTimeoutQDInv(s)
  requires EpochTimeoutQDInv(s')
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires InView1(s, sr);
  requires j !in sr;
  requires NonSuspector2(s, j);

  ensures  NonSuspector2(s', j) || NonSuspector0(s', j)
    || InternalSuspector3(s', j);

  ensures NotKnownSuspector(s', j);
{
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
    NonSuspector2_ind_7(s, s', sr, j);
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

////////////////////////////////////////////////////////////////////////////////
// Suspector inductive
////////////////////////////////////////////////////////////////////////////////

// self-step
lemma Suspector_ind_self(s:TimestampedRslState, s':TimestampedRslState, sr:set<int>, j:int)
  requires RslAssumption2(s, s')
  requires EpochTimeoutQDInv(s)
  requires EpochTimeoutQDInv(s')
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  // requires InView1(s, sr);
  // requires j in sr;
  requires Suspector(s, j);

  ensures Suspector(s, j);
{
}

// leader-step
lemma Suspector_ind_leader(s:TimestampedRslState, s':TimestampedRslState, sr:set<int>, j:int)
  requires RslAssumption2(s, s')
  requires EpochTimeoutQDInv(s)
  requires EpochTimeoutQDInv(s')
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[1];

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  // requires InView1(s, sr);
  // requires j in sr;
  requires Suspector(s, j);

  ensures Suspector(s, j);
{
}

}
