include "Environment.s.dfy"
include "Host.s.dfy"
include "DistributedSystem.s.dfy"
include "../Collections/Seqs.i.dfy"

module Performance_s {
  import opened Collections__Seqs_i

  /*
  datatype PerformanceReport = PerfVoid | PerfZero | PerformanceReport(instructions: int, pktsDelivered: int)
  type PerfReport = PerformanceReport

  predicate {:axiom} PerfLe(lhs:PerfReport, rhs:PerfReport)
    ensures lhs == PerfZero() ==> PerfLe(lhs, rhs)
    ensures lhs == PerfVoid() ==> PerfLe(lhs, rhs)

  function {:axiom} PerfMax(prs:seq<PerformanceReport>): PerformanceReport
    ensures |prs| == 1 ==> PerfMax(prs) == prs[0]
    ensures forall pr :: pr in prs ==> PerfLe(pr, PerfMax(prs))
    // ensures PerfVoid() in prs ==> PerfMax(prs) == PerfVoid()
    // ensures PerfMax(prs) == PerfMax(prs[PerfZero() := 0])
    ensures prs == []  ==> PerfMax(prs) == PerfZero()

  function {:axiom} PerfAdd(prs:multiset<PerformanceReport>) : PerformanceReport
    ensures PerfVoid() in prs ==> PerfAdd(prs) == PerfVoid()
    ensures PerfAdd(prs) == PerfAdd(prs[PerfZero() := 0])
    ensures forall s:multiset<PerformanceReport> :: var p := PerfAdd(s); p in prs ==> PerfAdd(prs) == PerfAdd(prs[p := prs[p] - 1])

  function {:axiom} PerfAdd2(pr1:PerformanceReport, pr2:PerformanceReport) : PerformanceReport
    ensures PerfAdd2(pr1, pr2) == PerfAdd(multiset{pr1, pr2})
    ensures pr1 == PerfVoid ==> PerfAdd2(pr1, pr2) == PerfVoid()
    ensures pr2 == PerfVoid ==> PerfAdd2(pr1, pr2) == PerfVoid()
    */

  datatype PerfExpr = PerfVoid | PerfZero | PerfAdd(prs:multiset<PerfExpr>) | PerfMax(prs:multiset<PerfExpr>)
  type PerformanceReport = PerfExpr
  type PerfReport = PerfExpr

  function {:axiom} PerfAdd2(p1:PerfExpr, p2:PerfExpr) : PerfExpr
    ensures p1.PerfAdd? ==> PerfAdd2(p1, p2) == PerfAdd(multiset{p2} + p1.prs)
  {
    PerfAdd(multiset{p1, p2})
  }

  predicate {:axiom} PerfEq(p1:PerfExpr, p2:PerfExpr)
    ensures PerfEq(p1, p2) == PerfEq(p2, p1)
    ensures p1 == p2 ==> PerfEq(p1, p2)
    ensures p1.PerfMax? && p1.prs == multiset{p2} ==> PerfEq(p1, p2)
    ensures p1.PerfMax? && p1.prs == multiset{} && p2 == PerfZero() ==> PerfEq(p1, p2)

    ensures p1.PerfAdd? && PerfVoid in p1.prs && p2 == PerfZero ==> PerfEq(p1, p2)
    ensures p1.PerfAdd? && p2 == PerfAdd(p1.prs[PerfZero() := 0]) ==> PerfEq(p1, p2)

    // Want to establish associativity
    // PerfAdd(multiset{PerfAdd(p1), prs'}) == PerfAdd(p1 + prs')
    ensures forall prs :: p1.PerfAdd? && PerfAdd(prs) in p1.prs && p2 == PerfAdd(p1.prs + prs - multiset{PerfAdd(prs)}) ==> PerfEq(p1, p2)
    ensures p1.PerfAdd? ==> (forall pr :: pr in p1.prs && pr.PerfAdd? && p2 == PerfAdd(p1.prs + pr.prs - multiset{PerfAdd(pr.prs)}) ==> PerfEq(p1, p2))

  lemma PerfProperties()
    // ensures forall p1, p2 :: p1.PerfAdd? ==> 
  {

  }
}
