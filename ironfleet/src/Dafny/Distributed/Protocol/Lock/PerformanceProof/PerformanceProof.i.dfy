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

  // No timeouts
  && (var nextStep := tgls.tls.t_environment.nextStep; nextStep.LEnvStepHostIos? ==>
  (forall io :: io in nextStep.ios ==> !io.LIoOpTimeoutReceive?))

  // TODO: This should be part of the state machine definition, and not an assumption
  // The only nodes that can do NodeGrant() currently hold the lock
  && (var nextStep := tgls.tls.t_environment.nextStep; nextStep.LEnvStepHostIos? && nextStep.nodeStep == GrantStep ==> tgls.tls.t_servers[nextStep.actor].v.held == true)

}

predicate GLSPerformanceAssumption(tglb:seq<TaggedGLS_State>)
{
  forall tgls :: tgls in tglb ==> SingleGLSPerformanceAssumption(tgls)
}

predicate SingleGLSPerformanceGuarantee(gls:TaggedGLS_State)
{
  forall pkt :: pkt in gls.tls.t_environment.sentPackets &&
    pkt.msg.v == Locked(|gls.tls.config|) ==> pkt.msg.pr == PerfZero()
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

    // No irrelevant packets in sentPackets
    && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && !(pkt.src in tgls.tls.t_servers)
    ==> false)

    // The node with the lock has correct PerfReport
    && PerfEq(tgls.tls.t_servers[tgls.tls.config[j]].pr, PerfBoundLockHeld(j))

    // TODO: This is unneeded
    // If there's a packet that grants lock to node j, it's PerfReport must be correct
    && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.dst == tgls.tls.config[j] && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
    ==> pkt.msg.pr == PerfBoundLockHeld(j) )

    // All nodes beyond j have void PerfReport
    && (forall k :: j < k < |tgls.tls.config| ==> PerfEq(tgls.tls.t_servers[tgls.tls.config[k]].pr, PerfVoid()))

    // No packets sent to nodes that have my_index higher than j
    && (forall pkt, k :: j < k < |tgls.tls.config| && pkt in tgls.tls.t_environment.sentPackets && pkt.dst in tgls.tls.t_servers && pkt.dst == tgls.tls.config[k]
    ==> false
    )
}

predicate PerfInvariantLockInNetwork(tgls: TaggedGLS_State, j:int)
    requires 0 < j < |tgls.tls.config|
  {
    // No irrelevant packets in sentPackets
    && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && !(pkt.src in tgls.tls.t_servers)
    ==> false)

    && |tgls.history| == j + 1
    && (forall id :: id in tgls.tls.t_servers ==> tgls.tls.t_servers[id].v.held == false && tgls.tls.t_servers[id].v.epoch < |tgls.history|)

    && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Transfer? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
    ==> pkt.msg.v.transfer_epoch <= |tgls.history|)

    && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Transfer? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
    && pkt.dst != tgls.tls.config[j]
    ==> pkt.msg.v.transfer_epoch < |tgls.history|  && pkt.msg.v.transfer_epoch <= tgls.tls.t_servers[pkt.dst].v.epoch)

    && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Transfer? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
        && pkt.msg.v.transfer_epoch == |tgls.history|
        ==> pkt.dst == tgls.tls.config[j] && PerfEq(pkt.msg.pr, PerfBoundLockInNetwork(j)))

    // The only packet sent to node j is an acceptable Transfer message
    && (forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.dst in tgls.tls.t_servers && pkt.dst == tgls.tls.config[j]
    ==> pkt.msg.v.Transfer? && pkt.msg.v.transfer_epoch > tgls.tls.t_servers[pkt.dst].v.epoch
    

    // No packets sent to nodes that have my_index higher than j
    && (forall pkt, k :: j < k < |tgls.tls.config| && pkt in tgls.tls.t_environment.sentPackets && pkt.dst in tgls.tls.t_servers && pkt.dst == tgls.tls.config[k]
    ==> false
    )
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

lemma Grant_j_InvLockHeldImpliesInvLockInNetwork(j:int, s:TaggedGLS_State, s':TaggedGLS_State)
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
  var p := PerfBoundLockHeld(j);
  var p' := PerfBoundLockInNetwork(j + 1);
  assert p' == PerfAdd2(p, GetStepRuntime(GrantStep));
}

// TODO: If the node holds the lock, then it cannot perform node accept. This
// needs to be part of the state machine definition
lemma Accept_j_InvLockHeldImpliesInvLockGranted(j:int, s:TaggedGLS_State, s':TaggedGLS_State)
{
}

lemma Accept_not_j_InvLockInNetworkImpliesInvLockHeld(j:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 < j < |s.tls.config| - 1
  requires TGLS_Consistency(s) && TGLS_Consistency(s')
  requires SingleGLSPerformanceAssumption(s)
  requires SingleGLSPerformanceAssumption(s')

  // A node other than node j taking a grant step
  requires s.tls.t_environment.nextStep.LEnvStepHostIos?
  requires s.tls.t_environment.nextStep.actor != s.tls.config[j]
  requires s.tls.t_environment.nextStep.nodeStep == AcceptStep

  requires PerfInvariantLockInNetwork(s, j);
  ensures PerfInvariantLockInNetwork(s', j);
{
  lemma_mod_auto(|s.tls.config|);
}

lemma Accept_j_InvLockInNetworkImpliesInvLockHeld(j:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 < j < |s.tls.config| - 1
  requires TGLS_Consistency(s) && TGLS_Consistency(s')
  requires SingleGLSPerformanceAssumption(s)
  requires SingleGLSPerformanceAssumption(s')

  // A node other than node j taking a grant step
  requires s.tls.t_environment.nextStep.LEnvStepHostIos?
  requires s.tls.t_environment.nextStep.actor == s.tls.config[j]
  requires s.tls.t_environment.nextStep.nodeStep == AcceptStep

  requires PerfInvariantLockInNetwork(s, j);
  ensures PerfInvariantLockHeld(s', j);
{
  lemma_mod_auto(|s.tls.config|);
}

lemma PerformanceGuaranteeHolds(config:Config, tglb:seq<TaggedGLS_State>)
  requires ValidTaggedGLSBehavior(tglb, config)
  requires GLSPerformanceAssumption(tglb)
  ensures GLSPerformanceGuarantee(tglb)
{
}

}
