include "Environment.s.dfy"
include "Host.s.dfy"
include "DistributedSystem.s.dfy"
include "../Collections/Seqs.i.dfy"

abstract module Performance_s {
  import opened Collections__Seqs_i
  import opened DS_s : DistributedSystem_s

  type PerfObject = i : int | 0 <= i

  function {:axiom} PerfMax(p1:PerfObject, p2:PerfObject) : PerfObject
    // ensures PerfMax(p1, p2) >= p1 && PerfMax(p1, p2) >= p2
    // ensures PerfMax(p1, p2) == p1 || PerfMax(p1, p2) == p2
    ensures p1 >= p2 ==> PerfMax(p1, p2) == p1
    ensures p2 >= p1 ==> PerfMax(p1, p2) == p2
  {
    if p1 < p2 then
      p2
    else 
      p1
  }

  function PerfZero() : PerfObject
  {
    0
  }

  function PerfVoid() : PerfObject
  {
    0
  }

  function {:axiom} PerfDelivery() : imap<nat, PerfObject>
    ensures imaptotal(PerfDelivery())

  function {:axiom} PerfStep(hstep:HostStep) : imap<nat, PerfObject>
    ensures imaptotal(PerfStep(hstep))

  function PerfAdd2(p1:PerfObject, p2:PerfObject) : PerfObject
  {
    p1 + p2
  }

  type PerfExpr = PerfObject

  datatype NumSteps = NumSteps(num_hsteps:imap<HostStep, nat>, num_deliveries:nat)
    
  predicate ValidNumSteps(num_steps:NumSteps)
  {
    imaptotal(num_steps.num_hsteps)
  }

  function NumStepsInit() : NumSteps
  {
    NumSteps(imap h:HostStep | true :: 0 as nat , 0)
  }

  function NoRecvPerfUpdate(node_pr:PerfExpr, hstep:HostStep, num_steps:NumSteps) : (PerfExpr, NumSteps)
    requires ValidNumSteps(num_steps)
    ensures ValidNumSteps(num_steps)
  {
    var step_idx := num_steps.num_hsteps[hstep];
    var total_time := PerfAdd2(node_pr, PerfStep(hstep)[step_idx]);
    var num_steps' := num_steps.(num_hsteps := num_steps.num_hsteps[hstep := num_steps.num_hsteps[hstep] + 1]);
    (total_time, num_steps')
  }

  function RecvPerfUpdate(node_pr:PerfExpr, pkt_pr:PerfExpr, hstep:HostStep, num_steps:NumSteps) : (PerfExpr, NumSteps)
    requires ValidNumSteps(num_steps)
    ensures ValidNumSteps(num_steps)
  {
    var delivery_idx := num_steps.num_deliveries;
    var deliveryTime := PerfAdd2(pkt_pr, PerfDelivery()[delivery_idx]);
    var handlerStartTime := PerfMax(deliveryTime, node_pr);
    var step_idx := num_steps.num_hsteps[hstep];
    var total_time := PerfAdd2(handlerStartTime, PerfStep(hstep)[step_idx]);
    var num_steps' := num_steps.(num_hsteps := num_steps.num_hsteps[hstep := num_steps.num_hsteps[hstep] + 1], num_deliveries := num_steps.num_deliveries + 1);
    (total_time, num_steps')
  }
}
