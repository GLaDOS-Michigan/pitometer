include "Phase1Proof.i.dfy"

module RslPhase1Proof_Generic {
import opened RslPhase1Proof_i

/* There can be no 1a messages sent in a
* non-LReplicaNextSpontaneousMaybeEnterNewViewAndSend1a step  */
lemma lemma_No1aSentInNon1aStep(ts:TimestampedRslState, ts':TimestampedRslState, op:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P1Assumption(ts, op)
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires Phase1Invariant(ts, op)
    requires idx != 0
    requires ts.t_replicas[idx].v.nextActionIndex != 1
    ensures forall p | p in ts'.t_environment.sentPackets && p.msg.v.RslMessage_1a? :: p in ts.t_environment.sentPackets
    ensures forall p | p in ts'.undeliveredPackets && p.msg.v.RslMessage_1a? :: p in ts.undeliveredPackets
{
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    var nextActionIndex := ls.v.nextActionIndex;

    forall p | p in ts'.t_environment.sentPackets && p.msg.v.RslMessage_1a?
    ensures p in ts.t_environment.sentPackets {
        if p !in ts.t_environment.sentPackets {
            if nextActionIndex == 2 {
                assert false;
            } else {
                forall io | io in tios && io.LIoOpSend?
                ensures !io.s.msg.v.RslMessage_1a? {}
                assert false;
            }
        }
    }
    lemma_No1aSentInNon1aStep_Undelivered(ts, ts', op, idx, tios);
}

lemma lemma_No1aSentInNon1aStep_Undelivered(ts:TimestampedRslState, ts':TimestampedRslState, op:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P1Assumption(ts, op)
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires Phase1Invariant(ts, op)
    requires idx != 0
    requires ts.t_replicas[idx].v.nextActionIndex != 1
    ensures forall p | p in ts'.undeliveredPackets && p.msg.v.RslMessage_1a? :: p in ts.undeliveredPackets
{
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    var nextActionIndex := ls.v.nextActionIndex;

    forall p | p in ts'.undeliveredPackets && p.msg.v.RslMessage_1a?
    ensures p in ts.undeliveredPackets {
        if p !in ts.undeliveredPackets {
            if nextActionIndex == 2 {
                assert false;
            } else {
                forall io | io in tios && io.LIoOpSend?
                ensures !io.s.msg.v.RslMessage_1a? {}
                assert false;
            }
        }
    }
}


/* Non-leader replicas do not send 1a's */
lemma lemma_NonLeaderDoesNotSend1a(ts:TimestampedRslState, ts':TimestampedRslState, op:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P1Assumption(ts, op)
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires Phase1Invariant(ts, op)
    requires idx != 0
    requires idx != 1
    ensures forall p | p in ts'.t_environment.sentPackets && p.msg.v.RslMessage_1a? :: p in ts.t_environment.sentPackets
    ensures forall p | p in ts'.undeliveredPackets && p.msg.v.RslMessage_1a? :: p in ts.undeliveredPackets
{
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    var nextActionIndex := ls.v.nextActionIndex;
    forall p | p in ts'.t_environment.sentPackets && p.msg.v.RslMessage_1a?
    ensures p in ts.t_environment.sentPackets {
        if p !in ts.t_environment.sentPackets {
            if nextActionIndex == 1 {
                assert ReplicasDistinct(ts'.constants.config.replica_ids, ls.v.replica.proposer.election_state.current_view.proposer_id, idx);
                assert ls.v.replica.proposer.election_state.current_view.proposer_id != ls.v.replica.proposer.constants.my_index;
                assert forall io | io in tios && io.LIoOpSend? :: !io.s.msg.v.RslMessage_1a?;
                assert false;
            } else {
                lemma_No1aSentInNon1aStep(ts, ts', op, idx, tios);
                assert false;
            }
        }
    }
    lemma_NonLeaderDoesNotSend1a_Undelivered(ts, ts', op, idx, tios);
}

/* Non-leader replicas do not send 1a's */
lemma lemma_NonLeaderDoesNotSend1a_Undelivered(ts:TimestampedRslState, ts':TimestampedRslState, op:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P1Assumption(ts, op)
    requires TimestampedRslNext(ts, ts')
    requires !TimestampedRslNextEnvironment(ts, ts')
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios)
    requires Phase1Invariant(ts, op)
    requires idx != 0
    requires idx != 1
    ensures forall p | p in ts'.undeliveredPackets && p.msg.v.RslMessage_1a? :: p in ts.undeliveredPackets
{
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    var nextActionIndex := ls.v.nextActionIndex;
    forall p | p in ts'.undeliveredPackets && p.msg.v.RslMessage_1a?
    ensures p in ts.undeliveredPackets {
        if p !in ts.undeliveredPackets {
            if nextActionIndex == 1 {
                assert ReplicasDistinct(ts'.constants.config.replica_ids, ls.v.replica.proposer.election_state.current_view.proposer_id, idx);
                assert ls.v.replica.proposer.election_state.current_view.proposer_id != ls.v.replica.proposer.constants.my_index;
                assert forall io | io in tios && io.LIoOpSend? :: !io.s.msg.v.RslMessage_1a?;
                assert false;
            } else {
                lemma_No1aSentInNon1aStep(ts, ts', op, idx, tios);
                assert false;
            }
        }
    }
}

}