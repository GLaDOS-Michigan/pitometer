include "Phase1Proof.i.dfy"

module RslPhase1Proof_Top {
import opened RslPhase1Proof_i


// This is actually the first state of Phase2
predicate Phase2Begin(ts:TimestampedRslState, opn:OperationNumber)
    requires |ts.t_replicas| > 2
    requires RslConsistency(ts)
{
    var l := ts.t_replicas[1];
    var r := l.v.replica;
    && AlwaysInvariant(ts, opn)
    && ExistingPacketsBallot(ts)
    && (!exists pkt :: pkt in ts.t_environment.sentPackets && IsNew2aPacket(pkt, opn))
    && (!exists pkt :: pkt in ts.t_environment.sentPackets && IsNew2bPacket(pkt, opn))
    && (!exists pkt :: pkt in ts.t_environment.sentPackets && IsNewReplyPacket(ts, pkt))
    && TimeLe(l.ts, NewLeaderInitTS)     // leader timestamp
    && l.v.nextActionIndex == 3          // leader action index is 3
    && LeaderSet1bContainsRequest(ts)
    && r.proposer.current_state == 2
    && r.proposer.election_state.current_view == Ballot(1, 1)
    && r.proposer.max_ballot_i_sent_1a == Ballot(1, 1)
    && opn == r.proposer.next_operation_number_to_propose

    // Learner and Executor states
    && opn == r.executor.ops_complete
    && BalLt(r.learner.max_ballot_seen, Ballot(1, 1))
}



predicate ExistingPacketsBallot(ts:TimestampedRslState) {
    forall pkt | pkt in ts.undeliveredPackets ::
    match pkt.msg.v {
        // All 1a and 1b packets have Ballot (1, 1)
        case RslMessage_1a(bal_1a)              => bal_1a == Ballot(1, 1) 
        case RslMessage_1b(bal_1b,_,_)          => bal_1b == Ballot(1, 1) 
        // All 2a and 2b messages have Ballot(1, 0) or (1, 1)
        case RslMessage_2a(bal_2a,_,_)          => bal_2a == Ballot(1, 0)
        case RslMessage_2b(bal_2b,_,_)          => bal_2b == Ballot(1, 0)

        // Cases where I don't care
        case RslMessage_Heartbeat(_,_,_)        => true
        case RslMessage_Invalid                 => true
        case RslMessage_Request(_,_)            => true
        case RslMessage_Reply(_,_)              => true
        case RslMessage_AppStateRequest(_,_)    => true
        case RslMessage_AppStateSupply(_,_,_,_) => true
        case RslMessage_StartingPhase2(_,_)     => true
    }
}

predicate ExistingPacketsBallotOpn(ts:TimestampedRslState, opn:OperationNumber) {
    forall pkt | pkt in ts.undeliveredPackets ::
    match pkt.msg.v {
        // All 1a and 1b packets have Ballot (1, 1)
        case RslMessage_1a(bal_1a)              => true
        case RslMessage_1b(bal_1b,_,_)          => true // votes specified in Always Invariant
        case RslMessage_2a(bal_2a,_,_)          => pkt.msg.v.opn_2a == opn
        case RslMessage_2b(bal_2b,_,_)          => pkt.msg.v.opn_2b == opn

        // Cases where I don't care
        case RslMessage_Heartbeat(_,_,_)        => true
        case RslMessage_Invalid                 => true
        case RslMessage_Request(_,_)            => true
        case RslMessage_Reply(_,_)              => true
        case RslMessage_AppStateRequest(_,_)    => true
        case RslMessage_AppStateSupply(_,_,_,_) => true
        case RslMessage_StartingPhase2(_,_)     => true
    }
}

// Main invariant 
predicate Phase1Invariant(ts:TimestampedRslState, opn:OperationNumber) 
    requires |ts.t_replicas| > 2
{
    && AlwaysInvariant(ts, opn)
    && true // TODO
}



lemma Phase1TopLevel(tglb:seq<TimestampedRslState>, opn:OperationNumber) returns (startPhase2Idx:int)
    requires exists con :: ValidTimestampedRSLBehavior(con, tglb)
    requires forall i | 0 <= i < |tglb| :: P1Assumption(tglb[i])
    // Phase1 lasts up till right before startPhase2Idx
    // startPhase2Idx is the initial state of phase 2
    ensures forall j | 0 <= j < |tglb| && j < startPhase2Idx :: Phase1Invariant(tglb[j], opn)
    ensures startPhase2Idx >= 0
    ensures startPhase2Idx < |tglb| ==> Phase2Begin(tglb[startPhase2Idx], opn)
{
    // TODO:
    assume false;
}



}
