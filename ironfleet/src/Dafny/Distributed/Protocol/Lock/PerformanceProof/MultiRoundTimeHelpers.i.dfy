// Must be verified with /arith:0. Make sure /noNLarith is not set
include "../../../Services/Lock/LockTimestampedDistributedSystem.i.dfy"
include "Definitions.i.dfy"

module PerformanceProof__MultiRoundTimeHelpers_i {
import opened LockTimestampedDistributedSystem_i
import opened PerformanceProof__Definitions_i

lemma lemma_Grant(node_pr:Timestamp, epoch:int)
  requires 0 < epoch
  requires TimeLe(node_pr, PerfBoundLockHeld(epoch))
  ensures TimeLe(TLS_NoRecvPerfUpdate(node_pr, GrantStep), PerfBoundLockInNetwork(epoch + 1))
{
  // TLS_NoRecvPerfUpdateIsMonotonic(node_pr, GrantStep, PerfBoundLockHeld(epoch));
}

lemma lemma_Accept(node_ts:Timestamp, pkt_ts:Timestamp, nd_epoch:int, pkt_epoch:int)
  requires 1 < pkt_epoch
  requires 0 <= nd_epoch
  requires nd_epoch + 1 <= pkt_epoch
  requires TimeLe(node_ts, PerfBoundLockInNetwork(nd_epoch + 1))
  requires TimeLe(pkt_ts, PerfBoundLockInNetwork(pkt_epoch))
  ensures TimeLe(TLS_RecvPerfUpdate(node_ts, pkt_ts, AcceptStep), PerfBoundLockHeld(pkt_epoch))
{
  if (pkt_epoch == 1) {
    
  }
  if nd_epoch > 0 {
    assert TimeLe(TLS_RecvPerfUpdate(node_ts, pkt_ts, AcceptStep), PerfBoundLockHeld(pkt_epoch));
  }
  else {
    
  }
}

}