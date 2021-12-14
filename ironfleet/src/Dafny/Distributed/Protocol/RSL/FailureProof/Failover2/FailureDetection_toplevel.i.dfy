include "../TimestampedRslSystem.i.dfy"

include "FailureHelpers.i.dfy"
include "FailureDetection_defns.i.dfy"
include "FailureDetection_helper0.i.dfy"
include "FailureDetection_helper1.i.dfy"
include "../Common/assumptions.i.dfy"

include "../../CommonProof/Constants.i.dfy"
include "FailureDetection.i.dfy"
include "FailureDetection_final.i.dfy"

module FailureDetection_toplevel_i {
import opened TimestampedRslSystem_i
import opened FailureHelpers_i
import opened FailureDetection_defns_i
import opened FailureDetection_helper0_i
import opened FailureDetection_helper1_i
import opened FailureDetection_i
import opened Common_Assumptions
import opened FailureDetection_final_i

lemma InvariantsHoldInitially(s:TimestampedRslState)
  requires FOAssumption(s)
  requires exists con :: TimestampedRslInit(con, s)
  ensures InView1(s, {})
  ensures DelayInvs(s)
{

}

predicate FailoverInv(s:TimestampedRslState)
  requires CommonAssumptions(s)
{
  (DelayInvs(s) &&
    (exists sr :: InView1(s, sr))) || FinalStage(s)
}

lemma FailoverInvMaintained(s:TimestampedRslState, s':TimestampedRslState)
  requires CommonAssumptions(s) && CommonAssumptions(s');
  requires FOAssumption2State(s,s')
  requires InFailover(s);
  requires InFailover(s) ==> FOAssumption(s);
  requires InFailover(s') ==> FOAssumption(s');
  requires TimestampedRslNext(s, s');
  requires FailoverInv(s);

  ensures InFailover(s') || InPhase1(s');
  ensures InFailover(s') ==> FailoverInv(s');
  ensures InPhase1(s') ==> FailoverFinal(s');
{
  if FinalStage(s) {
    FinalStageInd(s, s');
    if FailoverFinal(s') {
      assert InPhase1(s');
      assert FailoverFinal(s');
      return;
    } else {
      assert InFailover(s');
      assert FinalStage(s');
      return;
    }
  } else {
    var sr :| InView1(s, sr);
    assert InFailover(s');
    assert FOAssumption(s');
    var sr' := InView1_ind(s, s', sr);
    DelayInv_ind(s, s');
    assert FailoverInv(s');
  }
}
/*
lemma FailoverTopLevel_Prototype(tglb:seq<TimestampedRslState>) returns (startPhase1Idx:int)
  requires exists con :: ValidTimestampedRSLBehavior(con, tglb)
  requires InFailover(tglb[0])
  requires forall i :: 0 <= i < |tglb| ==> InFailover(tglb[i]) ==> FOAssumption(tglb[i])
  requires forall i :: 0 <= i < |tglb| ==> |tglb[i].t_replicas| > 2

  // ensures startPhase1Idx >= 0
  // ensures startPhase1Idx < |tglb| ==> FailoverFinal(tglb[startPhase1Idx])
  // ensures startPhase1Idx < |tglb| ==> InPhase1(tglb[startPhase1Idx])
  // ensures forall i | 0 <= i < |tglb| ::
    // forall pkt | pkt in tglb[i].t_environment.sentPackets :: !IsNewReplyPacket(tglb[i], pkt)
{
  startPhase1Idx := 0;

  var i := 1;
  InvariantsHoldInitially(tglb[0]);

  assert FailoverInv(tglb[0]);
  while i < |tglb|
    invariant i > 0
    invariant i <= |tglb|
    invariant FailoverInv(tglb[i - 1])
    invariant InFailover(tglb[i - 1])
  {
    var s := tglb[i-1];
    var s' := tglb[i];
    var sr :| InView1(s, sr);
    var sr' := InView1_ind(s, s', sr);
    if FinalStage(s') {
      return i;
    } else {
      i := i + 1;
    }
  }
}
*/
}
