include "Phase2Proof.i.dfy"

module Rs2Phase1Proof_Tony_i {
import opened RslPhase2Proof_i


predicate PerformanceInvariant(s:TimestampedRslState)
{
    || Phase2AcceptedLeaderInvariant(s, TimeVoid(), 0, )
}


lemma PerformanceGuaranteeHolds(con:LConstants, tb:seq<TimestampedRslState>) 
    requires |tb| > 0;
    requires forall i | 0 <= i < |tb| :: RslConsistency(tb[i])
    requires forall i | 0 <= i < |tb|-1 :: TimestampedRslNext(tb[i], tb[i+1])
    requires Pre2aInvariant(tb[0], TimeVoid(), 0)
    requires forall i | 0 <= i < |tb| :: RslAssumption(tb[i])
    ensures forall i | 0 <= i < |tb| :: PerformanceGuarantee(tb[i], TimeVoid())
{
    assert PerformanceGuarantee(tb[0], TimeVoid());
    var i := 0;
    while i < |tb| 
        decreases |tb| - i
        invariant 0 <= i <= |tb|
        invariant forall k | 0 <= k < i :: PerformanceGuarantee(tb[k], TimeVoid())
    {
        assume PerformanceGuarantee(tb[i], TimeVoid());
        i := i + 1;
    }
}

}