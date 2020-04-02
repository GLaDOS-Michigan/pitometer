include "Environment.s.dfy"
include "Host.s.dfy"
include "DistributedSystem.s.dfy"
include "../Collections/Seqs.i.dfy"

abstract module Performance_s {
  import opened Collections__Seqs_i
  import opened DS_s : DistributedSystem_s

  datatype PerfExpr = PerfDelivery | PerfStep(hstep:HostStep)| PerfVoid | PerfZero | PerfAdd(prs:multiset<PerfExpr>) | PerfMax(prs:multiset<PerfExpr>)

  function PerfAdd2(p1:PerfExpr, p2:PerfExpr) : PerfExpr
    ensures p1.PerfAdd? ==> PerfEq(PerfAdd2(p1, p2), PerfAdd(multiset{p2} + p1.prs))
  {
    PerfAdd(multiset{p1, p2})
  }

  predicate {:axiom} PerfEq(p1:PerfExpr, p2:PerfExpr)
    ensures PerfEq(p1, p2) == PerfEq(p2, p1)
    ensures p1 == p2 ==> PerfEq(p1, p2)
    ensures p1.PerfMax? && p1.prs == multiset{p2} ==> PerfEq(p1, p2)
    ensures p1.PerfMax? && p1.prs == multiset{} && p2 == PerfZero() ==> PerfEq(p1, p2)

    ensures p1.PerfMax? && p2.PerfMax? && p2.prs == p1.prs - multiset{PerfVoid} ==> PerfEq(p1, p2)

    ensures p1.PerfAdd? && PerfVoid in p1.prs && p2 == PerfZero ==> PerfEq(p1, p2)
    ensures p1.PerfAdd? && p2 == PerfAdd(p1.prs[PerfZero() := 0]) ==> PerfEq(p1, p2)
    ensures p1.PerfAdd? && p1.prs == multiset{} && p2 == PerfZero ==> PerfEq(p1, p2)
    ensures p2.PerfAdd? && p2.prs == multiset{} && p1 == PerfZero ==> PerfEq(p1, p2)

    // Want to establish associativity
    // PerfAdd(multiset{PerfAdd(p1), prs'}) == PerfAdd(p1 + prs')
    // ensures forall prs :: p1.PerfAdd? && PerfAdd(prs) in p1.prs && p2 == PerfAdd(p1.prs + prs - multiset{PerfAdd(prs)}) ==> PerfEq(p1, p2)
    // ensures p1.PerfAdd? ==> (forall pr :: pr in p1.prs && pr.PerfAdd? && p2 == PerfAdd(p1.prs + pr.prs - multiset{PerfAdd(pr.prs)}) ==> PerfEq(p1, p2))

    predicate {:axiom} PerfLe(p1: PerfExpr, p2:PerfExpr)

  lemma {:axiom} PerfProperties()
    // ensures forall p1, p2 :: p1.PerfAdd? ==> 
    ensures forall p1, p2, p1', p2' :: PerfEq(p1, p1') && PerfEq(p2, p2') ==> PerfEq(PerfAdd2(p1, p2), PerfAdd2(p1', p2'));
    ensures forall p1, p2 :: (forall p1' :: PerfEq(p1, p1') ==> PerfEq(PerfAdd2(p1, p2), PerfAdd2(p1', p2)));

    ensures forall p1, p2, p3 :: PerfEq(p1, p2) && PerfEq(p2, p3) ==> PerfEq(p1, p3);
    ensures forall p1, p2 :: PerfEq(p2, PerfVoid) ==> PerfEq(PerfMax(multiset{p1, p2}), PerfMax(multiset{p1}));
    ensures forall p1, p2, p3 :: PerfEq(PerfAdd2(p1, PerfAdd2(p2, p3)), PerfAdd(multiset{p1, p2, p3}));
    ensures forall p1, p2, p3 :: PerfEq(p1, p2) ==> (PerfEq(p1, p3) <==> PerfEq(p2, p3));

    ensures forall p1 : PerfExpr, p2 : PerfExpr {:trigger PerfAdd2(p1, p2)} :: (p1.PerfAdd? ==> PerfEq(PerfAdd2(p1, p2), PerfAdd(p1.prs + multiset{p2})));

    ensures forall prs {:trigger PerfMax(prs)} :: PerfEq(PerfMax(prs), PerfMax(prs - multiset{PerfVoid}));
    ensures forall p1 {:trigger PerfMax(multiset{p1})} :: PerfEq(PerfMax(multiset{p1}), p1)
}
