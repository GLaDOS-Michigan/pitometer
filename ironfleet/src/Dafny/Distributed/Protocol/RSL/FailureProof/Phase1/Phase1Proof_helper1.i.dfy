include "Phase1Proof.i.dfy"
include "GenericLemmas.i.dfy"

module RslPhase1Proof_Helper1 {
import opened RslPhase1Proof_i
import opened RslPhase1Proof_Generic


lemma PacketsBallotInvariant_ReceiveStep(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P1Assumption(ts, opn)
    requires InPhase1(ts') ==> P1Assumption(ts', opn)
    requires TimestampedRslNext(ts, ts')
    requires Phase1Invariant(ts, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires idx != 0
    requires ts.t_replicas[idx].v.nextActionIndex == 0 
    ensures PacketsBallotInvariant(ts')
{
    forall pkt | pkt in ts'.undeliveredPackets 
    ensures ExistingPacketsBallot(pkt)
    {}
}

lemma PacketsBallotInvariant_NoReceiveStep(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P1Assumption(ts, opn)
    requires InPhase1(ts') ==> P1Assumption(ts', opn)
    requires TimestampedRslNext(ts, ts')
    requires Phase1Invariant(ts, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires idx != 0
    requires ts.t_replicas[idx].v.nextActionIndex != 0 
    ensures PacketsBallotInvariant(ts')
{
    forall pkt | pkt in ts'.undeliveredPackets 
    ensures ExistingPacketsBallot(pkt)
    {}
}


/* Proof that a Phase1 state transitions to a Phase1 state when nextAction != 2 */
lemma Phase1_to_Phase1(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P1Assumption(ts, opn)
    requires InPhase1(ts') ==> P1Assumption(ts', opn)
    requires TimestampedRslNext(ts, ts')
    requires Phase1Invariant(ts, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    requires ts.t_replicas[idx].v.nextActionIndex != 2
    requires idx != 0
    ensures InPhase1(ts')
    ensures PerfInvariant(ts', opn)
{
    assert InPhase1(ts');
    Phase1_to_Phase1_LeaderTimeBound(ts, ts', opn, idx, tios);
    Phase1_to_Phase1_1aTimeboud(ts, ts', opn, idx, tios);
}

lemma Phase1_to_Phase1_1aTimeboud(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P1Assumption(ts, opn)
    requires P1Assumption(ts', opn)
    requires TimestampedRslNext(ts, ts')
    requires Phase1Invariant(ts, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    // requires ts.t_replicas[idx].v.nextActionIndex != 2
    requires idx != 0
    ensures forall pkt | pkt in ts'.undeliveredPackets && IsNew1aPacket(pkt)
        :: TimeLe(pkt.msg.ts, TimeBound1aDeliveryPost())
{   
    var ls, ls' := ts.t_replicas[idx], ts'.t_replicas[idx];
    var nextActionIndex := ls.v.nextActionIndex;

    if idx == 1 {
        if nextActionIndex == 1 {
            assert !BalLt(ls.v.replica.proposer.max_ballot_i_sent_1a, ls.v.replica.proposer.election_state.current_view);
            forall io | io in tios 
            ensures !io.LIoOpSend? {}
        } else {
            lemma_No1aSentInNon1aStep(ts, ts', opn, idx, tios);
        }
    } else {
        lemma_NonLeaderDoesNotSend1a(ts, ts', opn, idx, tios);
    }
}


lemma Phase1_to_Phase1_LeaderTimeBound(ts:TimestampedRslState, ts':TimestampedRslState, opn:OperationNumber, idx:int, tios:seq<TimestampedLIoOp<NodeIdentity, RslMessage>>) 
    requires CommonAssumptions(ts) && CommonAssumptions(ts')
    requires P1Assumption(ts, opn)
    requires P1Assumption(ts', opn)
    requires TimestampedRslNext(ts, ts')
    requires Phase1Invariant(ts, opn)
    requires TimestampedRslNextOneReplica(ts, ts', idx, tios);
    // requires ts.t_replicas[idx].v.nextActionIndex != 2
    requires idx != 0
    ensures TimeLe(ts'.t_replicas[1].ts, TimeBoundPhase1LeaderPost(ts'.t_replicas[1].v.nextActionIndex))
{  
    if idx != 1 {
        assert ts'.t_replicas[1].ts == ts.t_replicas[1].ts;
        return;
    }
    var ls, ls' := ts.t_replicas[1], ts'.t_replicas[1];
    var nextActionIndex, nextActionIndex' := ls.v.nextActionIndex, ls'.v.nextActionIndex;

    if nextActionIndex == 0 {
        var pkt := tios[0].r;
        assert ts.t_environment.nextStep.nodeStep == RslStep(0);
        if pkt.msg.v.RslMessage_1a? {
            assert IsNew1aPacket(pkt);
            assert TimeLe(pkt.msg.ts, TimeBound1aDeliveryPost());
            assert TimeLe(pkt.msg.ts, TimeBound1bDeliveryPost());
        } else if pkt.msg.v.RslMessage_1b? {
            assert IsNew1bPacket(pkt);
            assert TimeLe(pkt.msg.ts, TimeBound1bDeliveryPost());
        } else {
            assert false;
        }
        assert TimeLe(TimeMax(pkt.msg.ts, ls.ts), pkt.msg.ts + MaxQueueTime);
        assert TimeLe(ls'.ts, pkt.msg.ts + MaxQueueTime + StepToTimeDelta(RslStep(0)));
        assert TimeLe(ls'.ts, TimeBound1bDeliveryPost() + MaxQueueTime + TimeActionRange(1));
        assert TimeLe(ls'.ts, TimeBoundPhase1LeaderPost(nextActionIndex'));
    } else {
        var hstep := ts.t_environment.nextStep.nodeStep;
        assert hstep == RslStep(nextActionIndex);
    }
}

}