include "Environment.s.dfy"
include "Host.s.dfy"
include "DistributedSystem.s.dfy"
include "../Collections/Seqs.i.dfy"

module Performance_s {
  import opened Collections__Seqs_i

  datatype PerformanceReport = None | PerformanceReport(instructions: int, pktsDelivered: int)
  type PerfReport = PerformanceReport

  predicate {:axiom} PerfLe(lhs:PerfReport, rhs:PerfReport)

  function {:axiom} PerfMax(prs:seq<PerformanceReport>): PerformanceReport
    ensures forall pr :: pr in prs ==> PerfLe(pr, PerfMax(prs))

  function {:axiom} PerfAdd(pr1:PerformanceReport, pr2:PerformanceReport) : PerformanceReport
}
