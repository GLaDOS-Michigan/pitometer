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


////////////////////////////////////////////////////////////////////////////////
// PF_NONSUSP
////////////////////////////////////////////////////////////////////////////////

lemma NonSuspector1_ind_most(s:TimestampedRslState, s':TimestampedRslState, sr:set<int>, j:int)
  requires FOAssumption2(s, s')
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  // requires s.t_environment.nextStep.nodeStep != RslStep(0) // on step 0, we might just enter a new view because of HB
  requires s.t_environment.nextStep.nodeStep != RslStep(7) // on step 7, we might start new epoch
  requires s.t_environment.nextStep.nodeStep != RslStep(6) // on step 6, we might go to NS0(j)

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires InView1(s, sr);
  requires LeaderView0(s'); // assume that the leader doesn't enter the new view on this step

  requires j !in sr;
  requires NonSuspector1(s, j);
  ensures NonSuspector1(s', j);
  ensures NotKnownSuspector(s', j);
{
  if s.t_environment.nextStep.nodeStep == RslStep(0) {
    var ios := s.t_environment.nextStep.ios;

    if ios[0].LIoOpReceive?  {
      if ios[0].r.msg.v.RslMessage_Heartbeat? {
        assert NonSuspector1(s', j);
      } else if ios[0].r.msg.v.RslMessage_1a? {
        assert NonSuspector1(s', j);
      } else if ios[0].r.msg.v.RslMessage_1b? {
        assert NonSuspector1(s', j);
      } else if ios[0].r.msg.v.RslMessage_2b? {
        assert NonSuspector1(s', j);
      } else if ios[0].r.msg.v.RslMessage_2a? {
        assert NonSuspector1(s', j);
      } else if ios[0].r.msg.v.RslMessage_Request? {
        var es := s.t_replicas[j].v.replica.proposer.election_state;
        var newReq := Request(ios[0].r.src, ios[0].r.msg.v.seqno_req, ios[0].r.msg.v.val);
        assert ios[0].r in s.t_environment.sentPackets;
        assert ios[0].r.src == req.client;
        assert ios[0].r.msg.v.seqno_req == req.seqno;
        assert (req in es.requests_received_this_epoch) && RequestsMatch(req, newReq);

        assert s.t_replicas[j].v.replica.proposer.election_state.requests_received_this_epoch ==
          s'.t_replicas[j].v.replica.proposer.election_state.requests_received_this_epoch;
        assert NonSuspector1(s', j);
      } else if ios[0].r.msg.v.RslMessage_Reply? {
        assert NonSuspector1(s', j);
      } else if ios[0].r.msg.v.RslMessage_AppStateRequest? {
        assert NonSuspector1(s', j);
      } else if ios[0].r.msg.v.RslMessage_AppStateSupply? {
        assert NonSuspector1(s', j);
      } else if ios[0].r.msg.v.RslMessage_StartingPhase2? {
        assert NonSuspector1(s', j);
      } else{
        assert NonSuspector1(s', j);
      }
    } else {
        assert NonSuspector1(s', j);
    }
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(1) {
    assert NonSuspector1(s', j);
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(2) {
    assert NonSuspector1(s', j);
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(3) {
    assert NonSuspector1(s', j);
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(4) {
    assert NonSuspector1(s', j);
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(5) {
    assert NonSuspector1(s', j);
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(6) {
    assert false;
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(7) {
    assert false;
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(8) {
    var r := s.t_replicas[j].v.replica;
    // assert r.proposer.election_state.current_view_suspectors <= sr;
    SubsetCardinality(r.proposer.election_state.current_view_suspectors, sr);
    // if s'.t_replicas[j].v.replica.proposer.election_state.requests_received_this_epoch != [req] {
      // assert s'.t_replicas[j].v.replica.proposer.election_state.requests_received_this_epoch == [];
      // assert NonSuspector0(s', j);
    // } else {
      // assert NonSuspector1(s', j);
    // }
    // assert |r.proposer.election_state.current_view_suspectors| <= |sr|;
    // assert |sr| < LMinQuorumSize(s.constants.config);
    // assert |r.proposer.election_state.current_view_suspectors| < LMinQuorumSize(s.constants.config);
    // assert NonSuspector1(s', j, req);
  }
  else if s.t_environment.nextStep.nodeStep == RslStep(9) {
    assert NonSuspector1(s', j);
  }
}

lemma NonSuspector1_ind_6(s:TimestampedRslState, s':TimestampedRslState, j:int)
  requires FOAssumption2(s, s')
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep == RslStep(6) // on step 6, we might go to NS0(j)

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires NonSuspector1(s, j);
  requires NotKnownSuspector(s, j);
  ensures  NonSuspector1(s', j) || NonSuspector0(s', j);
  ensures NotKnownSuspector(s', j);
{
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

      // XXX: For proof stablity
      assert es'.requests_received_prev_epochs != [] ==>
        es'.requests_received_prev_epochs[0] in es'.requests_received_prev_epochs;
      assert es'.requests_received_prev_epochs == [];
      assert es'.requests_received_this_epoch <= es.requests_received_this_epoch;

      if s'.t_replicas[j].v.replica.proposer.election_state.requests_received_this_epoch == [] {
        assert NonSuspector0(s', j);
      } else{
        assert NonSuspector1(s', j);
      }
    } else {
      assert NonSuspector1(s', j);
    }
  } else {
    assert NonSuspector1(s', j);
  }
}

lemma NonSuspector1_ind_7(s:TimestampedRslState, s':TimestampedRslState, j:int)
  requires FOAssumption2(s, s')
  requires EpochTimeoutQDInv(s)
  // requires EpochTimeoutQDInv(s')
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep == RslStep(7)

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires NotKnownSuspector(s, j);
  requires NonSuspector1(s, j);
  ensures  NonSuspector1(s', j) || NonSuspector2(s', j);
  ensures NotKnownSuspector(s', j);
{
  // epoch might expire, and we might enter NonSuspector2()

  var ios := s.t_environment.nextStep.ios;
  var clock := SpontaneousClock(UntagLIoOpSeq(ios));
  var es := s.t_replicas[j].v.replica.proposer.election_state;
  if clock.t < es.epoch_end_time {
    assert NonSuspector1(s', j);
  } else {
    EpochQDHelper(s.t_replicas[j].ts, s'.t_replicas[j].ts);
    // var t' := s'.t_replicas[j].v.replica.proposer.election_state.epoch_end_time;
    // assert t' <= clock.t + EpochLength;
    assert NonSuspector2(s', j);
  }
}

lemma NonSuspector1_ind(s:TimestampedRslState, s':TimestampedRslState, sr:set<int>, j:int)
  requires FOAssumption2(s, s')
  requires EpochTimeoutQDInv(s)
  // requires EpochTimeoutQDInv(s')
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires InView1(s, sr);
  requires j !in sr;
  requires NonSuspector1(s, j);

  requires LeaderView0(s'); // assume that the leader doesn't enter the new view on this step

  ensures NonSuspector1(s', j) || NonSuspector0(s', j) || NonSuspector2(s', j);
  ensures NotKnownSuspector(s', j);
{
  var step := s.t_environment.nextStep.nodeStep;
  if step == RslStep(7) {
    NonSuspector1_ind_7(s, s', j);
  } else if step == RslStep(6) {
    NonSuspector1_ind_6(s, s', j);
  } else {
    NonSuspector1_ind_most(s, s', sr, j);
  }
}

lemma InView1Local_leader_quorum(s:TimestampedRslState, s':TimestampedRslState, sr:set<int>, j:int)
  requires FOAssumption2(s, s')
  requires DelayInvs(s)
  // requires DelayInvs(s')
  requires 0 <= j < |s.constants.config.replica_ids|;
  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);
  requires InView1(s, sr);
  requires LeaderView0(s');

  ensures LeaderQuorumBound(s')
{
  reveal_TBJustBeforeNewView();
  reveal_ActionsUpTo();
  reveal_TBFirstSuspectingHB();
  reveal_HBPeriodEnd();
  SubsetCardinality(s.t_replicas[1].v.replica.proposer.election_state.current_view_suspectors, sr);
  if j != 1 {
    assert LeaderQuorumBound(s');
    return;
  }

  // If already have quorum, then just maintain bound
  if |s.t_replicas[1].v.replica.proposer.election_state.current_view_suspectors| >= LMinQuorumSize(s'.constants.config) {
    // already have quorum, just maintain timestamp
    assert LeaderQuorumBound(s');
    return;
  }

  if s.t_environment.nextStep.nodeStep == RslStep(7) && |s'.t_replicas[1].v.replica.proposer.election_state.current_view_suspectors| >= LMinQuorumSize(s'.constants.config) {
    // Leader takes a local step that gives it a quorum.
    // This can only happen if the epoch2 timer is expired.
    // epoch_end <= Epoch2
    // t <= epoch_end + EpochQD(7)
    // --> t' <= Epoch2 + EpochQD(7) + RslStep(7)
    assert s.t_replicas[1].v.replica.constants.my_index !in s.t_replicas[1].v.replica.proposer.election_state.current_view_suspectors;
    assert s'.t_replicas[1].v.replica.constants.my_index in s'.t_replicas[1].v.replica.proposer.election_state.current_view_suspectors;
    if s.t_replicas[1].v.replica.constants.my_index in sr {
      assert s.t_replicas[1].v.replica.constants.my_index in s.t_replicas[1].v.replica.proposer.election_state.current_view_suspectors;
      assert false;
    }
    assert s.t_replicas[1].v.replica.constants.my_index !in sr;
    if NonSuspector1(s, 1) || NonSuspector0(s, 1) {
      assert s.t_replicas[1].v.replica.proposer.election_state.current_view_suspectors == s'.t_replicas[1].v.replica.proposer.election_state.current_view_suspectors;
      assert false;
      return;
    }
    assert NonSuspector2(s, 1);
    assert LeaderQuorumBound(s');
    return;
  } else if s.t_environment.nextStep.nodeStep == RslStep(0) {
    // Leader receives a packet that gives it a quorum.
    var ios := s.t_environment.nextStep.ios;
    if ios[0].LIoOpReceive?  {
      if ios[0].r.msg.v.RslMessage_Heartbeat? {
        if ios[0].r.msg.v.suspicious {
          assert ios[0].r.src in s.constants.config.replica_ids;
          var k := GetReplicaIndex(ios[0].r.src, s.constants.config);
          if k in s.t_replicas[1].v.replica.proposer.election_state.current_view_suspectors {
            return;
          }

          if k !in sr {
            assert NotKnownSuspector(s, k);
            assert false;
            return;
          }
          assert Suspector(s, k);
          var pkt :|
            && pkt in s.undeliveredPackets
            && pkt.msg.v.RslMessage_Heartbeat?
            && pkt.src == s.constants.config.replica_ids[k]
            && pkt.dst == s.constants.config.replica_ids[1]
            && pkt.msg.v.suspicious == true
            && TimeLe(pkt.msg.ts, TBFirstSuspectingHB())
            ;

          reveal_PacketDeliveredInOrder();
          assert TimeLe(ios[0].r.msg.ts, TBFirstSuspectingHB());
          reveal_TBJustBeforeNewView();
          assert LeaderQuorumBound(s');
          return;
        }
      }
    }
  }
  assert LeaderQuorumBound(s');
  return;
}

lemma InView1Local_self_ind(s:TimestampedRslState, s':TimestampedRslState, sr:set<int>, j:int) returns (sr':set<int>)
  requires FOAssumption2(s, s')
  requires DelayInvs(s)
  // requires DelayInvs(s')

  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires InView1(s, sr);
  requires LeaderView0(s');

  ensures SuspectingReplicaInv(s', sr');
  ensures InView1Local(s', j, j in sr');
  ensures sr' == sr  || sr' == sr + {j};
{
  sr' := sr;
  var sus := j in sr;
  if sus {
    Suspector_ind_self(s, s', sr, j);
    return;
  }

  // else, non-sus
  if NonSuspector0(s, j) {
    NonSuspector0_ind(s, s', sr, j);
  } else if NonSuspector1(s, j) {
    NonSuspector1_ind(s, s', sr, j);
  } else if NonSuspector2(s,j) {
    // if j == 1, it has a change of becoming a known suspector in this case
    sr' := NonSuspector2_ind(s, s', sr, j);
  } else {
    // j has a change of becoming a known suspector in this case
    sr' := InternalSuspector3_ind(s, s', sr, j);
  }
}

lemma InView1Local_leader_ind(s:TimestampedRslState, s':TimestampedRslState, sr:set<int>, k:int)
  requires FOAssumption2(s, s')
  requires EpochTimeoutQDInv(s)
  // requires EpochTimeoutQDInv(s')

  requires 0 <= k < |s.constants.config.replica_ids|;
  requires k != 1

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[1];

  requires TimestampedRslNextOneReplica(s, s', 1, s.t_environment.nextStep.ios);

  requires InView1(s, sr);
  requires LeaderView0(s');
  ensures InView1Local(s', k, k in sr)
{
  assert ReplicasDistinct(s.constants.config.replica_ids, k, 1);
  if k in sr {
    Suspector_ind_leader(s, s', sr, k);
  } else {
    assert InView1Local(s', k, k in sr);
  }
}

lemma InView1Local_all_ind(s:TimestampedRslState, s':TimestampedRslState, sr:set<int>, j:int) returns (sr':set<int>)
  requires FOAssumption2(s, s')
  requires DelayInvs(s)
  // requires DelayInvs(s')

  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires InView1(s, sr);
  requires LeaderView0(s');

  ensures  SuspectingReplicaInv(s', sr')
  ensures  forall k :: 0 <= k < |s'.t_replicas| ==> InView1Local(s', k, k in sr');

  // this needs to only be proven in the case that a leader takes a step
  ensures LeaderQuorumBound(s');
{
  sr' := InView1Local_self_ind(s, s', sr, j);
  InView1Local_leader_quorum(s, s', sr, j);
  // Also use InV1L_leader_ind
  if j == 1 {
    forall k | 0 <= k < |s'.t_replicas|
      ensures InView1Local(s', k, k in sr')
    {
      if k == j { // lemma above finishes this
      } else {
        InView1Local_leader_ind(s, s', sr, k);
      }
    }
    return;
  }

  assert ReplicasDistinct(s.constants.config.replica_ids, j, 1);
  forall k | 0 <= k < |s'.t_replicas|
    ensures InView1Local(s', k, k in sr')
  {
    if k == j {
      // lemma above finishes this.
    } else { // trivial; the state of node k is unchanged, and j is not a
      // leader. So, inv for k should be maintained
      assert ReplicasDistinct(s.constants.config.replica_ids, j, k);
    }
  }
}

lemma InView1_to_Packets(s:TimestampedRslState, s':TimestampedRslState, j:int, sr:set<int>)
  requires FOAssumption2(s, s')

  requires 0 <= j < |s.constants.config.replica_ids|;
  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires InView1(s, sr);
  ensures  InView1Packets(s');
{ // trivial
}

lemma InView1_to_CurrView(s:TimestampedRslState, s':TimestampedRslState, j:int, sr:set<int>)
  requires FOAssumption2(s, s')

  requires 0 <= j < |s.constants.config.replica_ids|;
  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires LeaderView0(s');
  requires InView1(s, sr);
  ensures  CurrView(s')
{
  SubsetCardinality(s.t_replicas[j].v.replica.proposer.election_state.current_view_suspectors, sr);
  // if s.t_environment.nextStep.nodeStep == RslStep(8) {
    // have to use subset cardinality argument
    // assert CurrView(s');
  // }
}

lemma LeaderEntersNewView(s:TimestampedRslState, s':TimestampedRslState, j:int, sr:set<int>)
  requires FOAssumption2(s, s')
  requires DelayInvs(s)
  // requires DelayInvs(s')

  requires 0 <= j < |s.constants.config.replica_ids|;
  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires InView1(s, sr);
  requires !LeaderView0(s')
  ensures FinalStage(s')
{
  reveal_TBNewView();
}

lemma InView1_ind_hostStep(s:TimestampedRslState, s':TimestampedRslState, j:int, sr:set<int>) returns (sr':set<int>)
  requires FOAssumption2(s, s')
  requires DelayInvs(s)
  // requires DelayInvs(s')

  requires 0 <= j < |s.constants.config.replica_ids|;
  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires InView1(s, sr);
  ensures  InView1(s', sr') || FinalStage(s');
{
  if LeaderView0(s') {
    sr' := InView1Local_all_ind(s, s', sr, j);
    InView1_to_Packets(s, s', j, sr);
    InView1_to_CurrView(s, s', j, sr);
    SubsetCardinality(s.t_replicas[j].v.replica.proposer.election_state.current_view_suspectors, sr);
    assert InView1(s', sr');
  } else {
    LeaderEntersNewView(s, s', j, sr);
  }
}

lemma InView1_ind(s:TimestampedRslState, s':TimestampedRslState, sr:set<int>) returns (sr':set<int>)
  requires FOAssumption2(s, s')
  requires DelayInvs(s)
  // requires DelayInvs(s')

  requires TimestampedRslNext(s, s');

  requires InView1(s, sr);
  ensures  InView1(s', sr') || FinalStage(s');
{
  sr' := sr;
  // three cases:
  if TimestampedRslNextEnvironment(s, s') {
    assert InView1(s', sr');
  } else if (exists j, ios :: TimestampedRslNextOneReplica(s, s', j, ios)) {
    // XXX: this is where the heavy lifting happens
    var j, ios :| TimestampedRslNextOneReplica(s, s', j, ios);
    sr' := InView1_ind_hostStep(s, s', j, sr);
  } else {
    var idx, ios :| TimestampedRslNextOneExternal(s, s', idx, ios);
    assert false; // Because we assume no external steps
  }
}

}
