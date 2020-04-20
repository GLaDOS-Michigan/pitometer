include "Environment.s.dfy"
include "Host.s.dfy"
include "DistributedSystem.s.dfy"
include "../Collections/Seqs.i.dfy"

abstract module Performance_s {
  import opened Collections__Seqs_i
  import opened DS_s : DistributedSystem_s

  datatype PerfExpr = PerfDelivery | PerfStep(hstep:HostStep)| PerfVoid | PerfZero | PerfAdd(prs:multiset<PerfExpr>) | PerfMax(prs:multiset<PerfExpr>)

  function PerfAdd2(p1:PerfExpr, p2:PerfExpr) : PerfExpr
  {
    PerfAdd(multiset{p1, p2})
  }

  predicate {:axiom} PerfEq(p1:PerfExpr, p2:PerfExpr)
    ensures p1 == p2 ==> PerfEq(p1, p2)

  predicate {:axiom} PerfLe(p1:PerfExpr, p2:PerfExpr)

  lemma {:axiom} PerfEq_IsEquivRelation()
    ensures forall p1, p2 :: PerfEq(p1, p2) == PerfEq(p2, p1);
    ensures forall p1, p2 :: p1 == p2 ==> PerfEq(p1, p2);
    ensures forall p1, p2, p3 :: PerfEq(p1, p2) && PerfEq(p2, p3) ==> PerfEq(p1, p3);

  lemma {:axiom} PerfAdd_IsAssociative()

  lemma {:axiom} PerfProperties()
    // PerfLe is a partial order
    // ensures forall p1, p2 :: PerfEq(p1, p2) ==> PerfLe(p1, p2) && PerfLe(p2, p1);
    // ensures forall p1, p2, p3 :: PerfLe(p1, p2) && PerfLe(p2, p3) ==> PerfLe(p1, p3);

    // 
    // ensures forall p1, p2, p1', p2' :: PerfEq(p1, p1') && PerfEq(p2, p2') ==> PerfLe(p1, p2) == PerfLe(p1', p2')

    // PerfMax is bigger than all prs
    // ensures forall prs {:trigger PerfMax(prs)} :: (forall pr :: pr in prs ==> PerfLe(pr, PerfMax(prs)));

    // PerfEq is a equivalence relation
    ensures forall p1, p2 :: PerfEq(p1, p2) == PerfEq(p2, p1);
    ensures forall p1, p2 :: p1 == p2 ==> PerfEq(p1, p2);
    ensures forall p1, p2, p3 :: PerfEq(p1, p2) && PerfEq(p2, p3) ==> PerfEq(p1, p3);

    // axioms involving PerfAdd
    ensures forall p1:PerfExpr, p2 :: p1.PerfAdd? && p1.prs == multiset{} && p2 == PerfZero ==> PerfEq(p1, p2)
    ensures forall p1, p2, p3 {:trigger PerfAdd2(p1, PerfAdd2(p2, p3))} :: PerfEq(PerfAdd2(p1, PerfAdd2(p2, p3)), PerfAdd(multiset{p1, p2, p3}));
    ensures forall p1, p2 :: (forall p1' :: PerfEq(p1, p1') ==> PerfEq(PerfAdd2(p1, p2), PerfAdd2(p1', p2)));
    ensures forall p1, p2, p1', p2' :: PerfEq(p1, p1') && PerfEq(p2, p2') ==> PerfEq(PerfAdd2(p1, p2), PerfAdd2(p1', p2'));

    // PerfAdd2 is associative with PerfAdd
    ensures forall p1 : PerfExpr, p2 : PerfExpr {:trigger PerfAdd2(p1, p2)} :: (p1.PerfAdd? ==> PerfEq(PerfAdd2(p1, p2), PerfAdd(p1.prs + multiset{p2})));

    // axioms involving PerfMax
    ensures forall p1, p2 :: PerfEq(p2, PerfVoid) ==> PerfEq(PerfMax(multiset{p1, p2}), PerfMax(multiset{p1}));
    ensures forall prs {:trigger PerfMax(prs)} :: PerfEq(PerfMax(prs), PerfMax(prs[PerfVoid := 0]));
    ensures forall p1 {:trigger PerfMax(multiset{p1})} :: PerfEq(PerfMax(multiset{p1}), p1)


  function NoRecvPerfUpdate(node_pr:PerfExpr, hstep:HostStep) : PerfExpr
  {
    var totalTime := PerfAdd2(node_pr, PerfStep(hstep));
    totalTime
  }

  function RecvPerfUpdate(node_pr:PerfExpr, pkt_pr:PerfExpr, hstep:HostStep) : PerfExpr
  {
    var deliveryTime := PerfAdd2(pkt_pr, PerfDelivery);
    var handlerStartTime := PerfMax(multiset{deliveryTime, node_pr});
    var totalTime := PerfAdd2(handlerStartTime, PerfStep(hstep));
    totalTime
  }
      
}
