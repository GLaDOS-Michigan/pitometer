include "../TimestampedRslSystem.i.dfy"

include "FailureHelpers.i.dfy"
include "FailureDetection_defns.i.dfy"
include "FailureDetection_helper0.i.dfy"
include "FailureDetection_helper1.i.dfy"
include "../Common/assumptions.i.dfy"

include "../../CommonProof/Constants.i.dfy"
include "FailureDetection.i.dfy"

module FailureDetection_final_i {
import opened TimestampedRslSystem_i
import opened FailureDetection_i
import opened Common_Assumptions

predicate FailoverFinal(s:TimestampedRslState)
  requires CommonAssumptions(s)
{
  // && (forall pkt ::
     // pkt in s.t_environment.sentPackets ==>
//
     // && (pkt.msg.v.RslMessage_Heartbeat? ==>
     // pkt.msg.v.bal_heartbeat == Ballot(1, 0))
     // && (pkt.msg.v.RslMessage_2a? ==>
     // pkt.msg.v.bal_2a == Ballot(1, 0))
     // && (pkt.msg.v.RslMessage_2b? ==>
     // pkt.msg.v.bal_2b == Ballot(1, 0))
     // && (pkt.msg.v.RslMessage_1a? ==>
     // pkt.msg.v.bal_1a == Ballot(1, 0))
     // && (pkt.msg.v.RslMessage_1b? ==>
     // pkt.msg.v.bal_1b == Ballot(1, 0))
  // )

  && s.t_replicas[1].v.replica.proposer.election_state.current_view == Ballot(1,1)
  && s.t_replicas[1].v.replica.proposer.current_state == 1
  && s.t_replicas[1].v.nextActionIndex == 2 // just checked for quorum of views, and got one
  // && TimeLe(s.t_replicas[1].ts, Fali())
}

// SLOW LEMMA
lemma FinalStageInd(s:TimestampedRslState, s':TimestampedRslState)
  requires CommonAssumptions(s) && CommonAssumptions(s');
  requires FOAssumption2State(s,s')
  requires FOAssumption(s);
  requires TimestampedRslNext(s, s');

  requires FinalStage(s)
  ensures FinalStage(s') || FailoverFinal(s')
{
  if TimestampedRslNextEnvironment(s, s') {
    assert FinalStage(s');
  } else if (exists j, ios :: TimestampedRslNextOneReplica(s, s', j, ios)) {
    // XXX: this is where the heavy lifting happens
    var j, ios :| TimestampedRslNextOneReplica(s, s', j, ios);
    if j == 1 {
      if s.t_replicas[1].v.nextActionIndex == 0 || s.t_replicas[1].v.nextActionIndex == 9 {
        assert FinalStage(s');
      } else {
        assert FailoverFinal(s');
      }
    }
  } else {
    var idx, ios :| TimestampedRslNextOneExternal(s, s', idx, ios);
    assert false; // Because we assume no external steps
  }
}

}
