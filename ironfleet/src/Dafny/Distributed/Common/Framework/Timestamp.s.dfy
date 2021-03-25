include "Environment.s.dfy"
include "Host.s.dfy"
include "DistributedSystem.s.dfy"
include "../Collections/Seqs.i.dfy"
include "../Collections/Maps2.s.dfy"

abstract module Timestamp_s {
  import opened Collections__Seqs_i
  import opened DS_s : DistributedSystem_s
  // import opened Host_s
  import opened Collections__Maps2_s

  type Timestamp = i : int | 0 <= i

  /* Timestamp datatype def */

  function {:axiom} TimeMax(p1:Timestamp, p2:Timestamp) : Timestamp
    // ensures TimeMax(p1, p2) >= p1 && TimeMax(p1, p2) >= p2
    // ensures TimeMax(p1, p2) == p1 || TimeMax(p1, p2) == p2
    ensures p1 >= p2 ==> TimeMax(p1, p2) == p1
    ensures p2 >= p1 ==> TimeMax(p1, p2) == p2
  {
    if p1 < p2 then
      p2
    else 
      p1
  }

  function TimeZero() : Timestamp
  {
    0
  }

  function TimeVoid() : Timestamp
  {
    0
  }

  function {:axiom} DeliveryTime() : imap<nat, Timestamp>
    ensures imaptotal(DeliveryTime())

  function {:axiom} StepTime(hstep:HostStep) : imap<nat, Timestamp>
    ensures imaptotal(StepTime(hstep))

  function {:axiom} Timeout(): Timestamp
    ensures Timeout() >= 0

  function TimeAdd2(p1:Timestamp, p2:Timestamp) : Timestamp
  {
    p1 + p2
  }

  datatype NumSteps = NumSteps(num_hsteps:imap<HostStep, nat>, num_deliveries:nat)
    
  predicate ValidNumSteps(num_steps:NumSteps)
  {
    imaptotal(num_steps.num_hsteps)
  }

  function NumStepsInit() : NumSteps
    ensures ValidNumSteps(NumStepsInit())
  {
    NumSteps(imap h:HostStep | true :: 0 as nat , 0)
  }

  function NoRecvTimestampUpdate(node_ts:Timestamp, hstep:HostStep, num_steps:NumSteps) : (Timestamp, NumSteps)
    requires ValidNumSteps(num_steps)
    ensures ValidNumSteps(num_steps)
  {
    var step_idx := num_steps.num_hsteps[hstep];
    var total_time := TimeAdd2(node_ts, StepTime(hstep)[step_idx]);
    var num_steps' := num_steps.(num_hsteps := num_steps.num_hsteps[hstep := num_steps.num_hsteps[hstep] + 1]);
    (total_time, num_steps')
  }

  function RecvTimestampUpdate(node_ts:Timestamp, pkt_ts:Timestamp, hstep:HostStep, num_steps:NumSteps) : (Timestamp, NumSteps)
    requires ValidNumSteps(num_steps)
    ensures ValidNumSteps(num_steps)
  {
    var delivery_idx := num_steps.num_deliveries;
    var deliveryTime := TimeAdd2(pkt_ts, DeliveryTime()[delivery_idx]);
    var handlerStartTime := TimeMax(deliveryTime, node_ts);
    var step_idx := num_steps.num_hsteps[hstep];
    var total_time := TimeAdd2(handlerStartTime, StepTime(hstep)[step_idx]);
    var num_steps' := num_steps.(num_hsteps := num_steps.num_hsteps[hstep := num_steps.num_hsteps[hstep] + 1], num_deliveries := num_steps.num_deliveries + 1);
    (total_time, num_steps')
  }
}
