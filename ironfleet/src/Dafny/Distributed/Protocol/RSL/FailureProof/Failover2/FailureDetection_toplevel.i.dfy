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
  requires FOAssumptionSealed(s)
  requires CommonAssumptions(s)
  requires exists con :: TimestampedRslInit(con, s)
  ensures FailoverInv(s);
  ensures PerfGuarantee(s);
  ensures InFailover(s);
{
  reveal_FOAssumptionSealed();
  // reveal_PerfGuarantee();
  assert InView1(s, {}) && DelayInvs(s);
}

predicate FailoverInv(s:TimestampedRslState)
  requires CommonAssumptions(s)
{
  (DelayInvs(s) &&
    (exists sr :: InView1(s, sr))) || FinalStage(s)
}

predicate {:opaque} FOAssumptionSealed(s:TimestampedRslState)
{
  FOAssumption(s)
}

predicate PerfGuarantee(s:TimestampedRslState)
  requires CommonAssumptions(s)
{
  forall pkt :: pkt in s.t_environment.sentPackets ==> !IsNewReplyPacket(s, pkt)
}

lemma FailoverInvMaintained(s:TimestampedRslState, s':TimestampedRslState)
  requires CommonAssumptions(s) && CommonAssumptions(s');
  requires FOAssumption2State(s,s')
  requires InFailover(s);
  requires InFailover(s) ==> FOAssumptionSealed(s);
  requires InFailover(s') ==> FOAssumptionSealed(s');
  requires TimestampedRslNext(s, s');
  requires FailoverInv(s);
  requires PerfGuarantee(s);

  ensures PerfGuarantee(s');
  ensures InFailover(s') || InPhase1(s');
  ensures InFailover(s') ==> FailoverInv(s');
  ensures InPhase1(s') ==> FailoverFinal(s');
{
  // reveal_PerfGuarantee();
  reveal_FOAssumptionSealed();
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

lemma FailoverTopLevel(tglb:seq<TimestampedRslState>) returns (startPhase1Idx:int)
  requires |tglb| > 0;
  requires exists con :: ValidTimestampedRSLBehavior(con, tglb)
  requires forall i | 0 <= i < |tglb| :: CommonAssumptions(tglb[i]);

  requires forall i | 0 <= i < |tglb| :: InFailover(tglb[i]) ==> FOAssumptionSealed(tglb[i]);
  requires forall i :: 0 < i < |tglb| ==> FOAssumption2State(tglb[i - 1], tglb[i])

  ensures startPhase1Idx >= 0
  ensures startPhase1Idx < |tglb| ==> FailoverFinal(tglb[startPhase1Idx])
  ensures startPhase1Idx < |tglb| ==> InPhase1(tglb[startPhase1Idx])

  ensures forall j :: 0 <= j < |tglb| ==> j < startPhase1Idx ==> PerfGuarantee(tglb[j]);
{
  startPhase1Idx := |tglb|;

  var i := 1;
  InvariantsHoldInitially(tglb[0]);

  assert FailoverInv(tglb[0]);
  while i < |tglb|
    decreases |tglb| - i
    invariant 0 < i <= |tglb|
    invariant FailoverInv(tglb[i - 1]) && PerfGuarantee(tglb[i - 1])
    invariant InFailover(tglb[i - 1])
    invariant forall j :: 0 <= j < i ==> PerfGuarantee(tglb[j]);
  {
    var s := tglb[i-1];
    var s' := tglb[i];

    FailoverInvMaintained(s, s');
    forall j | 0 <= j < i+1
      ensures PerfGuarantee(tglb[j])
    {
      if j == i {
        assert PerfGuarantee(tglb[j]);
      } else {
        assert PerfGuarantee(tglb[j]);
      }
    }

    if FailoverFinal(s') {
      return i;
    }
    i := i + 1;
  }
}

}
