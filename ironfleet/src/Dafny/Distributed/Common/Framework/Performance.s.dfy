include "Environment.s.dfy"
include "Host.s.dfy"
include "DistributedSystem.s.dfy"
include "../Collections/Seqs.i.dfy"

module Performance_s {
  import opened Collections__Seqs_i

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

  function {:axiom} PerfAdd2(pr1:PerformanceReport, pr2:PerformanceReport) : PerformanceReport
    ensures PerfAdd2(pr1, pr2) == PerfAdd(multiset{pr1, pr2})
    ensures pr1 == PerfVoid ==> PerfAdd2(pr1, pr2) == PerfVoid()
    ensures pr2 == PerfVoid ==> PerfAdd2(pr1, pr2) == PerfVoid()
}
