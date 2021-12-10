include "../DistributedSystem.i.dfy"

include "../../../Common/Collections/Maps2.i.dfy"
include "../Constants.i.dfy"
include "../Environment.i.dfy"
include "../Replica.i.dfy"

include "TimestampedRslSystem.i.dfy"

include "../CommonProof/Constants.i.dfy"
include "Phase1Helpers.i.dfy"

module RslPhase1Proof_i {
import opened TimestampedRslSystem_i
import opened RslPhase1Helpers_i

lemma lemma_RslInitImpliesInitInv(s:TimestampedRslState)
  requires RslAssumption(s);
  requires TimestampedRslInit(s.constants, s);
  ensures InitInvariant(s);
{
}

lemma lemma_leader_1_InitGoesToPhase1UnpreparedLeader(s:TimestampedRslState, s':TimestampedRslState, j:int)
  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 == j < |s.constants.config.replica_ids|

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  // FIXME: need to have this as part of InitInvariant (that the leader is in step 0 or 1)
  requires s.t_environment.nextStep.nodeStep == RslStep(1);

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires InitInvariant(s)
  ensures Phase1Invariant(s')
{
}

lemma lemma_not_receiveAndLeader_InitGoesToInit(s:TimestampedRslState, s':TimestampedRslState, j:int)
  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires !(j == 0 && s.t_environment.nextStep.nodeStep == RslStep(1));
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires InitInvariant(s)
  ensures InitInvariant(s')
{
}

lemma lemma_noreceive_Phase1UnpreparedLeaderGoesToPhase1UnpreparedLeader(s_prev:TimestampedRslState, s:TimestampedRslState, j:int)
  requires RslAssumption(s_prev);
  requires RslAssumption(s);
  requires RslConsistency(s_prev);
  requires RslConsistency(s);
  requires 0 <= j < |s_prev.constants.config.replica_ids|;
  requires 0 < j;

  requires s_prev.t_environment.nextStep.LEnvStepHostIos?;
  requires s_prev.t_environment.nextStep.actor == s_prev.constants.config.replica_ids[j];
  requires s_prev.t_environment.nextStep.nodeStep != RslStep(0)

  requires TimestampedRslNextOneReplica(s_prev, s, j, s_prev.t_environment.nextStep.ios);

  requires Phase1Invariant(s_prev)
  ensures Phase1Invariant(s)
{
}

lemma lemma_receive_Phase1UnpreparedLeaderGoesToPhase1UnpreparedLeader(s_prev:TimestampedRslState, s:TimestampedRslState, j:int)
  requires RslAssumption(s_prev);
  requires RslAssumption(s);
  requires RslConsistency(s_prev);
  requires RslConsistency(s);
  requires 0 <= j < |s_prev.constants.config.replica_ids|

  requires s_prev.t_environment.nextStep.LEnvStepHostIos?;
  requires s_prev.t_environment.nextStep.actor == s_prev.constants.config.replica_ids[j];
  requires s_prev.t_environment.nextStep.nodeStep == RslStep(0)

  requires TimestampedRslNextOneReplica(s_prev, s, j, s_prev.t_environment.nextStep.ios);

  requires Phase1Invariant(s_prev)
  ensures Phase1Invariant(s)
{
  var ios := s_prev.t_environment.nextStep.ios;
  var ts := s_prev.t_replicas[j].ts;
  var ts' := s.t_replicas[j].ts;
  assert |ios| > 0;
  if ios[0].LIoOpReceive? {
    if ios[0].r.msg.v.RslMessage_1a? {
      // show that the corresponding 1b packet has an OK timestamp
      if |ios| > 1 && ios[1].LIoOpSend? {
        var pkt := ios[1].s;
        Phase1aTimeHelper(ts, ios[0].r.msg.ts, pkt.msg.ts);
        // assert TimeLe(pkt.msg.ts, TimeBound1bDelivery());
        // assert 0 < |constants.config.replica_ids|;
        // assert pkt.dst == constants.config.replica_ids[0];
      } else {
      }
    }
  }
}

/*
lemma lemma_notleader_0_Phase1UnpreparedLeaderGoesToPhase1UnpreparedLeader(s:TimestampedRslState, s':TimestampedRslState, j:int, pkt:TimestampedRslPacket, progresses:map<NodeIdentity,Phase1Progress>)
  returns (progresses':map<NodeIdentity,Phase1Progress>)

  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|

  requires j > 0;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep == RslStep(0)

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Phase1UnpreparedLeaderInvariant(s, pkt, progresses);
  ensures Phase1UnpreparedLeaderInvariant(s', pkt, progresses');
{
  var ios := s.t_environment.nextStep.ios;
  assert ReplicasDistinct(s.constants.config.replica_ids, 0, j);
  if ios[0].LIoOpReceive? {
    assert ios[0].r.msg.v.RslMessage_1a?;
    assert ios[0].r.src == s.constants.config.replica_ids[0];
    BoundedLagImpliesBoundedProcessingTime(s.t_replicas[j].dts, s.t_replicas[j].ts, ios[0].r.msg.ts, s'.t_replicas[j].ts, TimeActionRange(0));
    progresses' := progresses[ios[0].r.dst := P1b];
  } else {
    progresses' := progresses;
  }
}

lemma lemma_leader_noreceive_Phase1UnpreparedLeaderGoesToPhase1UnpreparedLeader(s:TimestampedRslState, s':TimestampedRslState, j:int, pkt:TimestampedRslPacket, progresses:map<NodeIdentity,Phase1Progress>)
  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|

  requires j == 0;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep != RslStep(0)

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Phase1UnpreparedLeaderInvariant(s, pkt, progresses)
  ensures Phase1UnpreparedLeaderInvariant(s', pkt, progresses)
{
}

lemma lemma_leader_0_Phase1UnpreparedLeaderGoesToPhase1(s:TimestampedRslState, s':TimestampedRslState, j:int, pkt:TimestampedRslPacket, progresses:map<NodeIdentity,Phase1Progress>)
  returns (received_1b_ids:set<NodeIdentity>, progresses':map<NodeIdentity,Phase1Progress>)

  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|

  requires j == 0

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep == RslStep(0)

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Phase1UnpreparedLeaderInvariant(s, pkt, progresses);
  ensures Phase1UnpreparedLeaderInvariant(s', pkt, progresses') || Phase1PreparedLeaderInvariant(s', received_1b_ids, progresses');

{
  reveal_PacketDeliveredInOrder();
  var ios := s.t_environment.nextStep.ios;
  if ios[0].LIoOpReceive? {
    if ios[0].r.msg.v.RslMessage_1b? {
      assert ios[0].r.msg.ts <= pkt.msg.ts;
      assert pkt.msg.ts <= TimeBound1aSelfDelivery();
      assert TimeBound1aSelfDelivery() < ios[0].r.msg.ts;
      assert false;
    }
    assert ios[0].r == pkt;
    assert ios[0].r.src == s.constants.config.replica_ids[0];
    BoundedLagImpliesBoundedProcessingTime(s.t_replicas[j].dts, s.t_replicas[j].ts, ios[0].r.msg.ts, s'.t_replicas[j].ts, TimeActionRange(0));
    received_1b_ids := {};
    progresses' := progresses[s.constants.config.replica_ids[0] := P1b];
    assert Phase1PreparedLeaderInvariant(s', received_1b_ids, progresses');
  } else {
    progresses' := progresses;
    assert Phase1UnpreparedLeaderInvariant(s', pkt, progresses');
  }
}

lemma lemma_notleader_noreceive_Phase1GoesToPhase1(s:TimestampedRslState, s':TimestampedRslState, j:int, received_1b_ids:set<NodeIdentity>, progresses:map<NodeIdentity,Phase1Progress>)
  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|

  requires j > 0;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep != RslStep(0)

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Phase1PreparedLeaderInvariant(s, received_1b_ids, progresses)
  ensures Phase1PreparedLeaderInvariant(s', received_1b_ids, progresses)
{
  assert ReplicasDistinct(s.constants.config.replica_ids, 0, j);
}

lemma lemma_notleader_receive_Phase1GoesToPhase1(s:TimestampedRslState, s':TimestampedRslState, j:int, received_1b_ids:set<NodeIdentity>, progresses:map<NodeIdentity,Phase1Progress>)
  returns (progresses':map<NodeIdentity,Phase1Progress>)

  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|

  requires j > 0;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep == RslStep(0)

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Phase1PreparedLeaderInvariant(s, received_1b_ids, progresses)
  ensures Phase1PreparedLeaderInvariant(s', received_1b_ids, progresses')
{
  var ios := s.t_environment.nextStep.ios;
  assert ReplicasDistinct(s.constants.config.replica_ids, 0, j);
  progresses' := progresses;
  if ios[0].LIoOpReceive? {
    assert (ios[0].r.msg.v.RslMessage_1a?);
    BoundedLagImpliesBoundedProcessingTime(s.t_replicas[j].dts, s.t_replicas[j].ts, ios[0].r.msg.ts, s'.t_replicas[j].ts, TimeActionRange(0));
    progresses' := progresses[s.constants.config.replica_ids[j] := P1b];
  }
}

lemma lemma_leader_notreceiveOrPhase2_Phase1GoesToPhase1(s:TimestampedRslState, s':TimestampedRslState, j:int, received_1b_ids:set<NodeIdentity>, progresses:map<NodeIdentity,Phase1Progress>)
  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|

  requires j == 0;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep != RslStep(0)
  requires s.t_environment.nextStep.nodeStep != RslStep(2)

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Phase1PreparedLeaderInvariant(s, received_1b_ids, progresses)
  ensures Phase1PreparedLeaderInvariant(s', received_1b_ids, progresses)
{
}

lemma lemma_leader_mbePhase2_Phase1GoesToPhase1OrPhase2(s:TimestampedRslState, s':TimestampedRslState, j:int, received_1b_ids:set<NodeIdentity>, progresses:map<NodeIdentity,Phase1Progress>)
  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|

  requires j == 0;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep == RslStep(2)

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Phase1PreparedLeaderInvariant(s, received_1b_ids, progresses)
  ensures Phase1PreparedLeaderInvariant(s', received_1b_ids, progresses) || Phase2Invariant(s')
{
}

lemma lemma_leader_receive_Phase1GoesToPhase1(s:TimestampedRslState, s':TimestampedRslState, j:int, received_1b_ids:set<NodeIdentity>, progresses:map<NodeIdentity,Phase1Progress>)
  returns (received_1b_ids':set<NodeIdentity>, progresses':map<NodeIdentity, Phase1Progress>)

  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|

  requires j == 0;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep == RslStep(0)

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Phase1PreparedLeaderInvariant(s, received_1b_ids, progresses)
  ensures Phase1PreparedLeaderInvariant(s', received_1b_ids', progresses')
{
  var ios := s.t_environment.nextStep.ios;
  if ios[0].LIoOpReceive? {
    assert ios[0].r.msg.v.RslMessage_1b?;
    var size := |s.t_replicas[0].v.replica.proposer.received_1b_packets|;
    assert |s'.t_replicas[0].v.replica.proposer.received_1b_packets| == size + 1;
    received_1b_ids' := received_1b_ids + {ios[0].r.src};
    BoundedSizeLagImpliesBoundedProcessingTime(s.t_replicas[0].dts,
      s.t_replicas[0].ts,
      ios[0].r.msg.ts,
      s'.t_replicas[0].ts,
      size + 1
      );
    progresses' := progresses[ios[0].r.src := P1done];
  } else {
    received_1b_ids' := received_1b_ids;
    progresses' := progresses;
    LeaderTimeoutPreservesPhase1Invariant(s'.t_replicas[0].dts, |s.t_replicas[0].v.replica.proposer.received_1b_packets|, s'.t_replicas[0].v.nextActionIndex);
    // assert Phase1PreparedLeaderInvariant(s', received_1b_ids', progresses');
  }
}
*/
/*
lemma lemma_0_0_Phase1GoesToPhase1(s:TimestampedRslState, s':TimestampedRslState, j:int)
  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|

  requires j == 0;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires || s.t_environment.nextStep.nodeStep == RslStep(0);

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Phase1Invariant(s)
  ensures Phase1Invariant(s')
{
  assert forall idx :: 0 < idx < |s'.t_replicas| ==> ReplicasDistinct(s.constants.config.replica_ids, 0, idx);
  if (s.t_environment.nextStep.nodeStep == RslStep(0)) {
    var ios := s.t_environment.nextStep.ios;
    if (ios[0].LIoOpReceive?) {
      if ios[0].r.msg.v.RslMessage_1b? {
        var size := |s.t_replicas[0].v.replica.proposer.received_1b_packets|;
        var size' := |s'.t_replicas[0].v.replica.proposer.received_1b_packets|;
        assert UntagLPacket(ios[0].r) !in s.t_replicas[0].v.replica.proposer.received_1b_packets;
        assert s'.t_replicas[0].v.replica.proposer.received_1b_packets == s.t_replicas[0].v.replica.proposer.received_1b_packets + {UntagLPacket(ios[0].r)};
        assert size' == size + 1;
        BoundedSizeLagImpliesBoundedProcessingTime(s.t_replicas[j].dts, s.t_replicas[j].ts, ios[0].r.msg.ts, s'.t_replicas[j].ts, size);
      } else if (ios[0].r.msg.v.RslMessage_1a?) {
      } else{
      }
    } else {
    }
  } else if (s.t_environment.nextStep.nodeStep == RslStep(1)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(2)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(3)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(4)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(5)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(6)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(7)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(8)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(9)) {
  }
}

lemma lemma_a_0_Phase1GoesToPhase1(s:TimestampedRslState, s':TimestampedRslState, j:int)
  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|

  requires j == 0;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires || s.t_environment.nextStep.nodeStep != RslStep(2);
  requires || s.t_environment.nextStep.nodeStep != RslStep(0);

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Phase1Invariant(s)
  ensures Phase1Invariant(s')
{
  assert forall idx :: 0 < idx < |s'.t_replicas| ==> ReplicasDistinct(s.constants.config.replica_ids, 0, idx);
  if (s.t_environment.nextStep.nodeStep == RslStep(0)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(1)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(2)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(3)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(4)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(5)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(6)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(7)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(8)) {
  } else if (s.t_environment.nextStep.nodeStep == RslStep(9)) {
  }
}

lemma lemma_0_0_Phase1GoesToPhase1OrPhase2(s:TimestampedRslState, s':TimestampedRslState, j:int)
  requires RslAssumption(s);
  requires RslAssumption(s');
  requires RslConsistency(s);
  requires RslConsistency(s');
  requires 0 <= j < |s.constants.config.replica_ids|

  requires j == 0;

  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.constants.config.replica_ids[j];
  requires s.t_environment.nextStep.nodeStep == RslStep(2);

  requires TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);

  requires Phase1Invariant(s)
  ensures Phase1Invariant(s') // || Phase1Invariant(s')
{
}

*/

}
