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
  && (var nextStep := tgls.tls.t_environment.nextStep; nextStep.LEnvStepHostIos? ==> (nextStep.nodeStep == GrantStep <==> tgls.tls.t_servers[nextStep.actor].v.held == true))

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

predicate PerfInvariantAlways(tgls:TaggedGLS_State)
  requires TGLS_Consistency(tgls)
{
  // No one giving node 0 the lock
  && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
  && pkt.msg.v.Transfer? ==> !(pkt.dst == tgls.tls.config[0]) )

  // No irrelevant packets in sentPackets
  && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && !(pkt.src in tgls.tls.t_servers)
  ==> false)

  // No packets with epoch higher than history size
  && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Transfer? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
  ==> pkt.msg.v.transfer_epoch <= |tgls.history|)
}

predicate {:opaque} PerfInvariantLockHeld(tgls: TaggedGLS_State, j:int, epoch:int)
  requires 0 <= j < |tgls.tls.config|
  requires TGLS_Consistency(tgls)
{
    && PerfInvariantAlways(tgls)
    && tgls.tls.t_servers[tgls.tls.config[j]].v.held == true
    && |tgls.history| == epoch + 1
    && tgls.tls.t_servers[tgls.tls.config[j]].v.epoch == |tgls.history|

    &&  (forall id :: id in tgls.tls.t_servers && id != tgls.tls.config[j]
    ==> tgls.tls.t_servers[id].v.held == false && tgls.tls.t_servers[id].v.epoch < |tgls.history|)

    // All packets have epoch too small, so that no packets can be accepted in the next step
    && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Transfer? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
    ==> pkt.msg.v.transfer_epoch <= tgls.tls.t_servers[pkt.dst].v.epoch)

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

    // No packets sent by nodes that have my_index higher than j
    && (forall pkt, k :: j < k < |tgls.tls.config| && pkt in tgls.tls.t_environment.sentPackets && pkt.src in tgls.tls.t_servers && pkt.src == tgls.tls.config[k]
    ==> false
    )
}

predicate {:opaque} PerfInvariantLockInNetwork(tgls: TaggedGLS_State, j:int, epoch:int)
  requires 0 <= epoch
  requires 0 <= j < |tgls.tls.config|
  requires TGLS_Consistency(tgls)
{
  && PerfInvariantAlways(tgls)
  && |tgls.history| == epoch + 1
  && j == (epoch % |tgls.tls.config|)

  // No one holds the lock; everyone's epoch is below the epoch of the newest packet.
  && (forall id :: id in tgls.tls.t_servers ==> tgls.tls.t_servers[id].v.held == false && tgls.tls.t_servers[id].v.epoch < |tgls.history|)

  && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Transfer? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
  && pkt.dst != tgls.tls.config[j]
  ==> pkt.msg.v.transfer_epoch < |tgls.history|  && pkt.msg.v.transfer_epoch <= tgls.tls.t_servers[pkt.dst].v.epoch)

  && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Transfer? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
      && pkt.msg.v.transfer_epoch == |tgls.history|
      ==> pkt.dst == tgls.tls.config[j] && pkt.src == tgls.tls.config[(j - 1) % |tgls.tls.config|] && PerfEq(pkt.msg.pr, PerfBoundLockInNetwork(j)))

  // All nodes beyond j have void PerfReport
  && (forall k :: j < k < |tgls.tls.config| ==> PerfEq(tgls.tls.t_servers[tgls.tls.config[k]].pr, PerfVoid()))

  // The only packet sent to node j is an acceptable Transfer message
  && (forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.dst in tgls.tls.t_servers && pkt.dst == tgls.tls.config[j]
  ==> pkt.msg.v.Transfer? && pkt.msg.v.transfer_epoch == |tgls.history|
  )
  
  // No packets sent to nodes that have my_index higher than j
  && (forall pkt, k :: j < k < |tgls.tls.config| && pkt in tgls.tls.t_environment.sentPackets && pkt.dst in tgls.tls.t_servers && pkt.dst == tgls.tls.config[k]
  ==> false
  )

  // No packets sent by nodes that have my_index higher than j
  && (forall pkt, k :: j <= k < |tgls.tls.config| && pkt in tgls.tls.t_environment.sentPackets && pkt.src in tgls.tls.t_servers && pkt.src == tgls.tls.config[k]
  ==> false
  )
}

lemma NotHostIos_InvLockHeldImpliesInvLockHeld(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config|
  requires TGLS_Consistency(s) && TGLS_Consistency(s')

  // Not a HostIos step
  requires !s.tls.t_environment.nextStep.LEnvStepHostIos?

  requires PerfInvariantLockHeld(s, j, epoch);
  ensures PerfInvariantLockHeld(s', j, epoch);
{
  reveal_PerfInvariantLockHeld();
}

lemma Grant_not_j_InvLockHeldImpliesInvLockHeld(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config|
  requires TGLS_Consistency(s) && TGLS_Consistency(s')
  requires SingleGLSPerformanceAssumption(s)

  // A node other than node j taking a grant step
  requires s.tls.t_environment.nextStep.LEnvStepHostIos?
  requires s.tls.t_environment.nextStep.actor != s.tls.config[j]
  requires s.tls.t_environment.nextStep.nodeStep == GrantStep

  requires PerfInvariantLockHeld(s, j, epoch);
  ensures PerfInvariantLockHeld(s', j, epoch);
{
  reveal_PerfInvariantLockHeld();
}

lemma Accept_not_j_InvLockHeldImpliesInvLockHeld(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config|
  requires TGLS_Consistency(s) && TGLS_Consistency(s')
  requires SingleGLSPerformanceAssumption(s)

  // A node other than node j taking an accept step
  requires s.tls.t_environment.nextStep.LEnvStepHostIos?
  requires s.tls.t_environment.nextStep.actor != s.tls.config[j]
  requires s.tls.t_environment.nextStep.nodeStep == AcceptStep

  requires PerfInvariantLockHeld(s, j, epoch);
  ensures PerfInvariantLockHeld(s', j, epoch);
{
  reveal_PerfInvariantLockHeld();
}

lemma Grant_j_InvLockHeldImpliesInvLockInNetwork(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config| - 1
  requires TGLS_Consistency(s) && TGLS_Consistency(s')
  requires SingleGLSPerformanceAssumption(s)
  requires SingleGLSPerformanceAssumption(s')

  // A node other than node j taking a grant step
  requires s.tls.t_environment.nextStep.LEnvStepHostIos?
  requires s.tls.t_environment.nextStep.actor == s.tls.config[j]
  requires s.tls.t_environment.nextStep.nodeStep == GrantStep

  requires PerfInvariantLockHeld(s, j, epoch);
  ensures PerfInvariantLockInNetwork(s', j, epoch + 1);
{
  reveal_PerfInvariantLockHeld();
  reveal_PerfInvariantLockInNetwork();

  lemma_mod_auto(|s.tls.config|);
  PerfProperties();
  var p := PerfBoundLockHeld(j);
  var p' := PerfBoundLockInNetwork(j + 1);
  assert PerfEq(p', PerfAdd2(p, PerfStep(GrantStep)));
}

lemma NotHostIos_InvLockInNetworkImpliesInvLockInNetwork(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 < j < |s.tls.config|
  requires TGLS_Consistency(s) && TGLS_Consistency(s')

  // Not a HostIos step
  requires !s.tls.t_environment.nextStep.LEnvStepHostIos?

  requires PerfInvariantLockInNetwork(s, j, epoch);
  ensures PerfInvariantLockInNetwork(s', j, epoch);
{
  reveal_PerfInvariantLockInNetwork();
}

lemma Accept_not_j_InvLockInNetworkImpliesInvLockInNetwork(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 < j < |s.tls.config| - 1
  requires TGLS_Consistency(s) && TGLS_Consistency(s')
  requires SingleGLSPerformanceAssumption(s)
  requires SingleGLSPerformanceAssumption(s')

  // A node other than node j taking a grant step
  requires s.tls.t_environment.nextStep.LEnvStepHostIos?
  requires s.tls.t_environment.nextStep.actor != s.tls.config[j]
  requires s.tls.t_environment.nextStep.nodeStep == AcceptStep

  requires PerfInvariantLockInNetwork(s, j, epoch);
  ensures PerfInvariantLockInNetwork(s', j, epoch);
{
  reveal_PerfInvariantLockInNetwork();

  lemma_mod_auto(|s.tls.config|);
}

lemma Accept_j_InvLockInNetworkImpliesInvLockHeld(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 < j < |s.tls.config| - 1
  requires TGLS_Consistency(s) && TGLS_Consistency(s')
  requires SingleGLSPerformanceAssumption(s)
  requires SingleGLSPerformanceAssumption(s')

  // A node other than node j taking a grant step
  requires s.tls.t_environment.nextStep.LEnvStepHostIos?
  requires s.tls.t_environment.nextStep.actor == s.tls.config[j]
  requires s.tls.t_environment.nextStep.nodeStep == AcceptStep

  requires PerfInvariantLockInNetwork(s, j, epoch);
  ensures PerfInvariantLockHeld(s', j, epoch);
{
  reveal_PerfInvariantLockInNetwork();
  reveal_PerfInvariantLockHeld();
  PerfProperties();

  lemma_mod_auto(|s.tls.config|);
  var p2 := PerfBoundLockInNetwork(j);
  var p2' := PerfBoundLockHeld(j);

  assert PerfEq(p2', PerfAdd2(p2, PerfStep(AcceptStep)));

  var tgls := s;
  var tgls' := s';
  var ios := s.tls.t_environment.nextStep.ios;
  assert ios[1].s.dst == s.tls.config[j - 1];
  // assert PerfInvariantLockHeldNonOpaque(s', j);
}

predicate PerfInvariant(tgls:TaggedGLS_State)
{
  && TGLS_Consistency(tgls)
  && ( 
    || (exists epoch, j :: 0 <= epoch && 0 <= j < |tgls.tls.config| && PerfInvariantLockHeld(tgls, j, epoch))
    || (exists epoch, j :: 0 <= epoch && 0 <= j < |tgls.tls.config| && PerfInvariantLockInNetwork(tgls, j, epoch))
  )
}

lemma PerfInvariantLockHeldGoesToPerfInvariant(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires SingleGLSPerformanceAssumption(s) && TGLS_Consistency(s)
  requires SingleGLSPerformanceAssumption(s') && TGLS_Consistency(s')
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config| - 1
  requires PerfInvariantLockHeld(s, j, epoch)

  ensures PerfInvariant(s')
{
  if !s.tls.t_environment.nextStep.LEnvStepHostIos? {
    NotHostIos_InvLockHeldImpliesInvLockHeld(j, epoch, s, s');
  } else {
    if s.tls.t_environment.nextStep.actor == s.tls.config[j] {
      if s.tls.t_environment.nextStep.nodeStep == GrantStep {
        Grant_j_InvLockHeldImpliesInvLockInNetwork(j, epoch, s, s');
      } else {
        reveal_PerfInvariantLockHeld();
        assert false;
      }
    } else {
      if s.tls.t_environment.nextStep.nodeStep == GrantStep {
        Grant_not_j_InvLockHeldImpliesInvLockHeld(j, epoch, s, s');
      } else {
        Accept_not_j_InvLockHeldImpliesInvLockHeld(j, epoch, s, s');
      }
    }
  }
}

lemma PerfInvariantLockInNetworkGoesToPerfInvariant(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires SingleGLSPerformanceAssumption(s) && TGLS_Consistency(s)
  requires SingleGLSPerformanceAssumption(s') && TGLS_Consistency(s')
  requires TGLS_Next(s, s')
  requires 0 < j < |s.tls.config|
  requires PerfInvariantLockInNetwork(s, j, epoch)

  ensures PerfInvariant(s')
{
  if !s.tls.t_environment.nextStep.LEnvStepHostIos? {
    NotHostIos_InvLockInNetworkImpliesInvLockInNetwork(j, epoch, s, s');
  } else {
    if s.tls.t_environment.nextStep.actor == s.tls.config[j] {
      if s.tls.t_environment.nextStep.nodeStep == GrantStep {
        reveal_PerfInvariantLockInNetwork();
        assert false;
      } else {
        Accept_j_InvLockInNetworkImpliesInvLockHeld(j, epoch, s, s');
      }
    } else {
      if s.tls.t_environment.nextStep.nodeStep == GrantStep {
        reveal_PerfInvariantLockInNetwork();
        assert false;
      } else {
        Accept_not_j_InvLockInNetworkImpliesInvLockInNetwork(j, epoch, s, s');
      }
    }
  }
}

lemma PerfInvariantMaintained(s:TaggedGLS_State, s':TaggedGLS_State)
  requires SingleGLSPerformanceAssumption(s) && TGLS_Consistency(s)
  requires SingleGLSPerformanceAssumption(s') && TGLS_Consistency(s')
  requires TGLS_Next(s, s')
  requires PerfInvariant(s)

  ensures PerfInvariant(s')
{
  if (exists j, epoch :: 0 <= epoch && 0 <= j < |s.tls.config| && PerfInvariantLockHeld(s, j, epoch)) {
    var epoch, j :| 0 <= epoch && 0 <= j < |s.tls.config| && PerfInvariantLockHeld(s, j, epoch);
    PerfInvariantLockHeldGoesToPerfInvariant(j, epoch, s, s');
  }
  else {
    var epoch, j :| 0 <= epoch && 0 <= j < |s.tls.config| && PerfInvariantLockInNetwork(s, j, epoch);
    PerfInvariantLockInNetworkGoesToPerfInvariant(j, epoch, s, s');
  }
}

lemma PerformanceGuaranteeHolds(config:Config, tglb:seq<TaggedGLS_State>)
  requires ValidTaggedGLSBehavior(tglb, config)
  requires GLSPerformanceAssumption(tglb)
  ensures GLSPerformanceGuarantee(tglb)
{
}

}
