include "../../DistributedSystem.i.dfy"

// copied this here because the original one doesn't build because of random includes that it shouldn't have
module CommonProof__Requests_i {

import opened LiveRSL__DistributedSystem_i

lemma lemma_RemoveAllSatisfiedRequestsInSequenceProducesSubsequence(s':seq<Request>, s:seq<Request>, r:Request)
    requires s' == RemoveAllSatisfiedRequestsInSequence(s, r);
    decreases s, 1;
    ensures  forall x :: x in s' ==> x in s;
{
    if |s| > 0 && !RequestsMatch(s[0], r)
    {
        lemma_RemoveAllSatisfiedRequestsInSequenceProducesSubsequence(RemoveAllSatisfiedRequestsInSequence(s[1..], r), s[1..], r);
    }

}

lemma lemma_RemoveExecutedRequestBatchProducesSubsequence(s':seq<Request>, s:seq<Request>, batch:RequestBatch)
    requires s' == RemoveExecutedRequestBatch(s, batch);
    ensures  forall x :: x in s' ==> x in s;
    decreases |batch|;
{
    if |batch| > 0
    {
        var s'' := RemoveAllSatisfiedRequestsInSequence(s, batch[0]);
        lemma_RemoveAllSatisfiedRequestsInSequenceProducesSubsequence(s'', s, batch[0]);
        lemma_RemoveExecutedRequestBatchProducesSubsequence(s', s'', batch[1..]);
    }
}

}
