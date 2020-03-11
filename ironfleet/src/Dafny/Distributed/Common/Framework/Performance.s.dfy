include "Environment.s.dfy"
include "Host.s.dfy"
include "DistributedSystem.s.dfy"
include "../Collections/Seqs.i.dfy"

module Performance_s {
  import opened Collections__Seqs_i

  datatype PerformanceReport = PerfNone | PerformanceReport(instructions: int, pktsDelivered: int)
  type PerfReport = PerformanceReport

  predicate {:axiom} PerfLe(lhs:PerfReport, rhs:PerfReport)

  function {:axiom} PerfMax(prs:seq<PerformanceReport>): PerformanceReport
    ensures forall pr :: pr in prs ==> PerfLe(pr, PerfMax(prs))
    ensures PerfNone() in prs ==> PerfMax(prs) == PerfNone()
    ensures prs == []  ==> PerfMax(prs) == PerfNone()

  function {:axiom} PerfAdd(prs:multiset<PerformanceReport>) : PerformanceReport
    ensures PerfNone() in prs ==> PerfAdd(prs) == PerfNone()

  function {:axiom} PerfAdd2(pr1:PerformanceReport, pr2:PerformanceReport) : PerformanceReport
    ensures PerfAdd2(pr1, pr2) == PerfAdd(multiset{pr1, pr2})
    ensures pr1 == PerfNone ==> PerfAdd2(pr1, pr2) == PerfNone()
    ensures pr2 == PerfNone ==> PerfAdd2(pr1, pr2) == PerfNone()
}
