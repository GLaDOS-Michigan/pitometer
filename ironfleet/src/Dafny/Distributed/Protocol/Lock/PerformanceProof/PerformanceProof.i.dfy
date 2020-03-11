include "../../../Services/Lock/LockTaggedDistributedSystem.i.dfy"
include "TaggedGLS.i.dfy"
include "Definitions.i.dfy"
include "../../../../Libraries/Math/mod_auto.i.dfy"

module PerformanceProof_i {

import opened LockTaggedDistributedSystem_i
  import opened TaggedGLS_i
  import opened PerformanceProof__Definitions_i
  import opened Math__mod_auto_i

predicate SingleGLSPerformanceAssumption(tgls:TaggedGLS_State)
{
  // The only nodes that take steps are in the ring
  && (tgls.tls.t_environment.nextStep.LEnvStepHostIos? ==> tgls.tls.t_environment.nextStep.actor in tgls.tls.t_servers)

  // There are fewer nodes than the uint64 upper bound
  && (|tgls.tls.config| < 0xFFFF_FFFF_FFFF_FFFF)
}

predicate GLSPerformanceAssumption(tglb:seq<TaggedGLS_State>)
{
  forall tgls :: tgls in tglb ==> SingleGLSPerformanceAssumption(tgls)
}

predicate SingleGLSPerformanceGuarantee(gls:TaggedGLS_State)
{
  forall pkt :: pkt in gls.tls.t_environment.sentPackets &&
    pkt.msg.v == Locked(|gls.tls.config|) ==> pkt.msg.pr == PerformanceReport(0, 0)
}

predicate GLSPerformanceGuarantee(tglb:seq<TaggedGLS_State>)
{
  forall tgls :: tgls in tglb ==> SingleGLSPerformanceGuarantee(tgls)
}

predicate TGLS_Consistency(tgls: TaggedGLS_State)
{
  && (forall id :: id in tgls.tls.config <==> id in tgls.tls.t_servers)
    && (forall id :: id in tgls.tls.t_servers ==> (tgls.tls.t_servers[id].v.config == tgls.tls.config))
    && (forall i :: 0 <= i < |tgls.tls.config| ==> tgls.tls.t_servers[tgls.tls.config[i]].v.my_index == i)
}

predicate PerfInvariantLockHeld(tgls: TaggedGLS_State, j:int)
    requires 0 <= j < |tgls.tls.config|
    requires TGLS_Consistency(tgls)
{
  tgls.tls.t_servers[tgls.tls.config[j]].v.held == true
    && |tgls.history| == j + 1
    && tgls.tls.t_servers[tgls.tls.config[j]].v.epoch == |tgls.history|

    &&  (forall id :: id in tgls.tls.t_servers && id != tgls.tls.config[j]
    ==> tgls.tls.t_servers[id].v.held == false && tgls.tls.t_servers[id].v.epoch < |tgls.history|)
    
    && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
    && pkt.msg.v.Transfer? ==> !(pkt.dst == tgls.tls.config[0]) )

    // All packets have epoch too small, so that no packets can be accepted in the next step
    && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Transfer? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
    ==> pkt.msg.v.transfer_epoch <= tgls.tls.t_servers[pkt.dst].v.epoch)

    // All irrelevant packets have zero PerfReport
    && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && !(pkt.src in tgls.tls.t_servers)
    ==> pkt.msg.pr == PerfNone())

    // The node with the lock has correct PerfReport
    && tgls.tls.t_servers[tgls.tls.config[j]].pr == PerfBound(j)

    // If there's a packet that grants lock to node j, it's PerfReport must be correct
    && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.dst == tgls.tls.config[j] && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
    ==> pkt.msg.pr == PerfBound(j) )

    // All nodes beyond j have zero PerfReport
    && (forall k :: j < k < |tgls.tls.config| ==> tgls.tls.t_servers[tgls.tls.config[k]].pr == PerfNone() )

    // No packets sent to nodes that have my_index higher than j
    && (forall pkt, k :: j < k < |tgls.tls.config| && pkt in tgls.tls.t_environment.sentPackets && pkt.dst in tgls.tls.t_servers && pkt.dst != tgls.tls.config[k]
    ==> !(pkt.src in tgls.tls.t_servers)
    )
}

predicate PerfInvariantLockInNetwork(tgls: TaggedGLS_State, j:int)
    requires 0 < j < |tgls.tls.config|
  {
    && |tgls.history| == j + 1
    && (forall id :: id in tgls.tls.t_servers ==> tgls.tls.t_servers[id].v.held == false && tgls.tls.t_servers[id].v.epoch < |tgls.history|)

    && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Transfer? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
    ==> pkt.msg.v.transfer_epoch <= |tgls.history|)

    && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Transfer? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
    && pkt.dst != tgls.tls.config[j]
    ==> pkt.msg.v.transfer_epoch < |tgls.history| )

    /*
    && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Transfer? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
        && pkt.msg.v.transfer_epoch == |tgls.history|
        ==> pkt.dst == tgls.tls.config[j] /*&& pkt.msg.pr ==  PerfBound(j - 1)*/)
       */
  }

lemma NotHostIos_InvLockHeldImpliesInvLockHeld(j:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config|
  requires TGLS_Consistency(s) && TGLS_Consistency(s')

  // Not a HostIos step
  requires !s.tls.t_environment.nextStep.LEnvStepHostIos?

  requires PerfInvariantLockHeld(s, j);
  ensures PerfInvariantLockHeld(s', j);
{
}

lemma Grant_not_j_InvLockHeldImpliesInvLockHeld(j:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config|
  requires TGLS_Consistency(s) && TGLS_Consistency(s')
  requires SingleGLSPerformanceAssumption(s)

  // A node other than node j taking a grant step
  requires s.tls.t_environment.nextStep.LEnvStepHostIos?
  requires s.tls.t_environment.nextStep.actor != s.tls.config[j]
  requires s.tls.t_environment.nextStep.nodeStep == GrantStep

  requires PerfInvariantLockHeld(s, j);
  ensures PerfInvariantLockHeld(s', j);
{
}

lemma Accept_not_j_InvLockHeldImpliesInvLockHeld(j:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config|
  requires TGLS_Consistency(s) && TGLS_Consistency(s')
  requires SingleGLSPerformanceAssumption(s)

  // A node other than node j taking an accept step
  requires s.tls.t_environment.nextStep.LEnvStepHostIos?
  requires s.tls.t_environment.nextStep.actor != s.tls.config[j]
  requires s.tls.t_environment.nextStep.nodeStep == AcceptStep

  requires PerfInvariantLockHeld(s, j);
  ensures PerfInvariantLockHeld(s', j);
{
}

lemma Grant_j_InvLockHeldImpliesInvLockHeld(j:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config| - 1
  requires TGLS_Consistency(s) && TGLS_Consistency(s')
  requires SingleGLSPerformanceAssumption(s)
  requires SingleGLSPerformanceAssumption(s')

  // A node other than node j taking a grant step
  requires s.tls.t_environment.nextStep.LEnvStepHostIos?
  requires s.tls.t_environment.nextStep.actor == s.tls.config[j]
  requires s.tls.t_environment.nextStep.nodeStep == GrantStep

  requires PerfInvariantLockHeld(s, j);
  ensures PerfInvariantLockInNetwork(s', j + 1);
{
  lemma_mod_auto(|s.tls.config|);
}

// TODO: If the node holds the lock, then it cannot perform node accept. This will be
// part of the state machine definition
lemma Accept_j_InvLockHeldImpliesInvLockGranted(j:int, s:TaggedGLS_State, s':TaggedGLS_State)
{
}

lemma PerformanceGuaranteeHolds(config:Config, tglb:seq<TaggedGLS_State>)
  requires ValidTaggedGLSBehavior(tglb, config)
  requires GLSPerformanceAssumption(tglb)
  ensures GLSPerformanceGuarantee(tglb)
{
}

}
