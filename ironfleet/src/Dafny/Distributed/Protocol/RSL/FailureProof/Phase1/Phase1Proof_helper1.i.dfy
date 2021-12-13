include "Phase1Proof.i.dfy"
// include "GenericLemmas.i.dfy"

module RslPhase1Proof_Helper1 {
import opened RslPhase1Proof_i
// import opened RslPhase1Proof_Generic


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

}