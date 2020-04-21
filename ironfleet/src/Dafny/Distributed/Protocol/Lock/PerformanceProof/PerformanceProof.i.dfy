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


  // The size of the history/epochs never reaches UINT64_MAX
  && (|tgls.history| < 0xFFFF_FFFF_FFFF_FFFF)

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
  |gls.tls.config| > 1 ==> 
  (forall pkt :: pkt in gls.tls.t_environment.sentPackets &&
    pkt.msg.v == Transfer(|gls.tls.config|) ==> PerfEq(pkt.msg.pr, PerfBoundLockInNetwork(|gls.tls.config|)))
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
  // No irrelevant packets in sentPackets
  && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && !(pkt.src in tgls.tls.t_servers)
  ==> false)

  // No nodes with epoch higher
  // All transfer packets have correct performance object, determined by its transfer_epoch
  && (forall id :: id in tgls.tls.t_servers ==> 0 <= tgls.tls.t_servers[id].v.epoch <= |tgls.history|)

  // No transfer packets with epoch higher than history size
  // All transfer packets have correct performance object, determined by its transfer_epoch
  && (forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Transfer? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
  ==> 0 < pkt.msg.v.transfer_epoch <= |tgls.history|)

  && (forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Locked? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
  ==> 0 < pkt.msg.v.locked_epoch <= |tgls.history|)
}

predicate PerfInvariantLockHeld(tgls: TaggedGLS_State, j:int, epoch:int)
  requires 0 < epoch
  requires 0 <= j < |tgls.tls.config|
  requires TGLS_Consistency(tgls)
{
    && PerfInvariantAlways(tgls)
    && SingleGLSPerformanceGuarantee(tgls)
    && tgls.tls.t_servers[tgls.tls.config[j]].v.held == true
    && |tgls.history| == epoch
    && tgls.tls.t_servers[tgls.tls.config[j]].v.epoch == |tgls.history|
    // && j == (epoch % |tgls.tls.config|)

    &&  (forall id :: id in tgls.tls.t_servers && id != tgls.tls.config[j]
    ==> tgls.tls.t_servers[id].v.held == false && tgls.tls.t_servers[id].v.epoch < |tgls.history|)

    // All packets have epoch too small, so that no packets can be accepted in the next step
    && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Transfer? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
    ==> pkt.msg.v.transfer_epoch <= tgls.tls.t_servers[pkt.dst].v.epoch)

    // The node with the lock has correct PerfReport
    && PerfEq(tgls.tls.t_servers[tgls.tls.config[j]].pr, PerfBoundLockHeld(epoch))

    // All nodes beyond j have void PerfReport
    && (forall k :: j < k < |tgls.tls.config| ==> tgls.tls.t_servers[tgls.tls.config[k]].pr == PerfVoid())

  // No packets sent to nodes that have my_index higher than j, including node zero
  && (forall pkt, k :: j < k < |tgls.tls.config| && pkt in tgls.tls.t_environment.sentPackets && pkt.dst in tgls.tls.t_servers && pkt.dst == tgls.tls.config[k]
  ==> false
  )

  // No packets sent by nodes that have my_index higher than j
  && (forall pkt, k :: j < k < |tgls.tls.config| && pkt in tgls.tls.t_environment.sentPackets && pkt.src in tgls.tls.t_servers && pkt.src == tgls.tls.config[k]
  ==> false
  )
}

predicate PerfInvariantLockInNetwork(tgls: TaggedGLS_State, j:int, epoch:int)
  requires 1 < epoch
  requires 0 < j < |tgls.tls.config|
  requires TGLS_Consistency(tgls)
{
  && PerfInvariantAlways(tgls)
  && SingleGLSPerformanceGuarantee(tgls)
  && |tgls.history| == epoch
  // && j == (epoch) % |tgls.tls.config|

  // No one holds the lock; everyone's epoch is below the epoch of the newest packet.
  && (forall id :: id in tgls.tls.t_servers ==> tgls.tls.t_servers[id].v.held == false && tgls.tls.t_servers[id].v.epoch < |tgls.history|)

  && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Transfer? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
  && pkt.dst != tgls.tls.config[j % |tgls.tls.config|]
  ==> pkt.msg.v.transfer_epoch < |tgls.history| && pkt.msg.v.transfer_epoch <= tgls.tls.t_servers[pkt.dst].v.epoch)

  && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Transfer? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
      && pkt.msg.v.transfer_epoch == |tgls.history|
      ==> pkt.dst == tgls.tls.config[j] && pkt.src == tgls.tls.config[j - 1] && PerfEq(pkt.msg.pr, PerfBoundLockInNetwork(epoch)))

  // All nodes with index j or higher have void PerfReport
  && (forall k :: j <= k < |tgls.tls.config| ==> tgls.tls.t_servers[tgls.tls.config[k]].pr == PerfVoid())

  // The only packet sent to node j is an acceptable Transfer message
  && (forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.dst in tgls.tls.t_servers && pkt.dst == tgls.tls.config[j]
  ==> pkt.msg.v.Transfer? && pkt.msg.v.transfer_epoch == |tgls.history|
  )
  
  // No packets sent to nodes that have my_index higher than j, including node zero
  && (forall pkt, k :: j < k < |tgls.tls.config| && pkt in tgls.tls.t_environment.sentPackets && pkt.dst in tgls.tls.t_servers && pkt.dst == tgls.tls.config[k]
  ==> false
  )

  // No packets sent by nodes that have my_index higher than j
  && (forall pkt, k :: j < k < |tgls.tls.config| && pkt in tgls.tls.t_environment.sentPackets && pkt.src in tgls.tls.t_servers && pkt.src == tgls.tls.config[k]
  ==> false
  )
}

predicate PerfInvariantEpochHigherThanNumServers(tgls:TaggedGLS_State)
  requires TGLS_Consistency(tgls)
{
  && PerfInvariantAlways(tgls)
  && SingleGLSPerformanceGuarantee(tgls)
  && |tgls.history| > |tgls.tls.config| - 1

  // If a node has epoch that's too low, it does not hold the lock
  &&  (forall id :: id in tgls.tls.t_servers && tgls.tls.t_servers[id].v.epoch < |tgls.tls.config| ==> tgls.tls.t_servers[id].v.held == false)

  // All packets with epoch too low are unacceptable
  && (forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Transfer? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
      && pkt.msg.v.transfer_epoch <= |tgls.tls.config| - 1 ==> pkt.msg.v.transfer_epoch <= tgls.tls.t_servers[pkt.dst].v.epoch)
}

// TODO: Change this invariant
predicate PerfInvariantNodeZeroGrantedLock(tgls:TaggedGLS_State)
  requires 0 < |tgls.tls.config|
  requires TGLS_Consistency(tgls)
{
  && PerfInvariantAlways(tgls)
    && tgls.tls.t_servers[tgls.tls.config[0]].v.held == true
    && |tgls.history| == |tgls.tls.config|
    && tgls.tls.t_servers[tgls.tls.config[0]].v.epoch == |tgls.history|
}

predicate PerfInvariantNodeZeroReceivedLock(tgls:TaggedGLS_State)
  requires TGLS_Consistency(tgls)
{
  && 0 < |tgls.tls.config|
  && PerfInvariantAlways(tgls)
    && tgls.tls.t_servers[tgls.tls.config[0]].v.held == true
    && |tgls.history| == |tgls.tls.config|
    && tgls.tls.t_servers[tgls.tls.config[0]].v.epoch == |tgls.history|
}

lemma NotHostIos_InvLockHeldImpliesInvLockHeld(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config|
  requires 0 < epoch
  requires TGLS_Consistency(s) && TGLS_Consistency(s')

  // Not a HostIos step
  requires !s.tls.t_environment.nextStep.LEnvStepHostIos?

  requires PerfInvariantLockHeld(s, j, epoch);
  ensures PerfInvariantLockHeld(s', j, epoch);
{
  //reveal_PerfInvariantLockHeld();
}

lemma Grant_not_j_InvLockHeldImpliesInvLockHeld(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config|
  requires 0 < epoch <= |s.tls.config|
  requires TGLS_Consistency(s) && TGLS_Consistency(s')
  requires SingleGLSPerformanceAssumption(s)

  // A node other than node j taking a grant step
  requires s.tls.t_environment.nextStep.LEnvStepHostIos?
  requires s.tls.t_environment.nextStep.actor != s.tls.config[j]
  requires s.tls.t_environment.nextStep.nodeStep == GrantStep

  requires PerfInvariantLockHeld(s, j, epoch);
  ensures PerfInvariantLockHeld(s', j, epoch);
{
  lemma_mod_auto(|s.tls.config|);
  //reveal_PerfInvariantLockHeld();
}

lemma Accept_not_j_InvLockHeldImpliesInvLockHeld(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config|
  requires 0 < epoch <= |s.tls.config|
  requires TGLS_Consistency(s) && TGLS_Consistency(s')
  requires SingleGLSPerformanceAssumption(s)

  // A node other than node j taking an accept step
  requires s.tls.t_environment.nextStep.LEnvStepHostIos?
  requires s.tls.t_environment.nextStep.actor != s.tls.config[j]
  requires s.tls.t_environment.nextStep.nodeStep == AcceptStep

  requires PerfInvariantLockHeld(s, j, epoch);
  ensures PerfInvariantLockHeld(s', j, epoch);
{
  lemma_mod_auto(|s.tls.config|);
  var ios := s.tls.t_environment.nextStep.ios;
  assert IsValidLIoOp(ios[0], s.tls.t_environment.nextStep.actor, s.tls.t_environment);
}

lemma Grant_j_InvLockHeldImpliesInvLockInNetwork(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config| - 1
  requires 0 < epoch < |s.tls.config|
  requires TGLS_Consistency(s) && TGLS_Consistency(s')
  requires SingleGLSPerformanceAssumption(s)
  requires SingleGLSPerformanceAssumption(s')

  // A node other than node j taking a grant step
  requires s.tls.t_environment.nextStep.LEnvStepHostIos?
  requires s.tls.t_environment.nextStep.actor == s.tls.config[j]
  requires s.tls.t_environment.nextStep.nodeStep == GrantStep

  requires PerfInvariantLockHeld(s, j, epoch);
  ensures PerfInvariantLockInNetwork(s', j + 1, epoch + 1);
{
  //reveal_PerfInvariantLockHeld();
  //reveal_PerfInvariantLockInNetwork();

  lemma_mod_auto(|s.tls.config|);
  // PerfProperties();
  //var p := PerfBoundLockHeld(epoch);
  //var p' := PerfBoundLockInNetwork(epoch + 1);
  //assert PerfEq(p', PerfAdd2(p, PerfStep(GrantStep)));
  Grant_j_helper_specific(epoch);
}

lemma Grant_LastNode_InvLockHeldImpliesInvEpochHigher(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 <= j == |s.tls.config| - 1
  requires 0 <= epoch == |s.tls.config|
  requires TGLS_Consistency(s) && TGLS_Consistency(s')
  requires SingleGLSPerformanceAssumption(s)
  requires SingleGLSPerformanceAssumption(s')

  // A node other than node j taking a grant step
  requires s.tls.t_environment.nextStep.LEnvStepHostIos?
  requires s.tls.t_environment.nextStep.actor == s.tls.config[j]
  requires s.tls.t_environment.nextStep.nodeStep == GrantStep

  requires PerfInvariantLockHeld(s, j, epoch);
  ensures PerfInvariantEpochHigherThanNumServers(s')
{
  //reveal_PerfInvariantLockHeld();
  //reveal_PerfInvariantLockInNetwork();

  lemma_mod_auto(|s.tls.config|);
  // PerfProperties();
  var p := PerfBoundLockHeld(epoch);
  var p' := PerfBoundLockInNetwork(epoch + 1);
  // assert PerfEq(p', PerfAdd2(p, PerfStep(GrantStep)));
}

lemma NotHostIos_InvLockInNetworkImpliesInvLockInNetwork(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 < j < |s.tls.config|
  requires 1 < epoch <= |s.tls.config|
  requires TGLS_Consistency(s) && TGLS_Consistency(s')

  // Not a HostIos step
  requires !s.tls.t_environment.nextStep.LEnvStepHostIos?

  requires PerfInvariantLockInNetwork(s, j, epoch);
  ensures PerfInvariantLockInNetwork(s', j, epoch);
{
  //reveal_PerfInvariantLockInNetwork();
}

lemma Accept_not_j_InvLockInNetworkImpliesInvLockInNetwork(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 < j < |s.tls.config|
  requires 1 < epoch <= |s.tls.config|
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
  //reveal_PerfInvariantLockInNetwork();
  assert s.tls.t_environment.nextStep.ios[0] in s.tls.t_environment.nextStep.ios;
  lemma_mod_auto(|s.tls.config|);
}

lemma Accept_j_InvLockInNetworkImpliesInvLockHeld(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 < j < |s.tls.config|
  requires 1 < epoch <= |s.tls.config|
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
  //reveal_PerfInvariantLockInNetwork();
  //reveal_PerfInvariantLockHeld();
  // PerfProperties();
  lemma_mod_auto(|s.tls.config|);
  assert s.tls.t_environment.nextStep.ios[0] in s.tls.t_environment.nextStep.ios;
  Accept_j_helper();
  // lemma_Accept_j_PR(s.tls.t_environment.nextStep.ios[0].r.msg.pr, s.tls.t_servers[s.tls.config[j]].pr, s'.tls.t_servers[s.tls.config[j]].pr, epoch);
}

lemma PerfInvariantLockInNetworkGoesToPerfInvariant(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires SingleGLSPerformanceAssumption(s) && TGLS_Consistency(s)
  requires SingleGLSPerformanceAssumption(s') && TGLS_Consistency(s')
  requires TGLS_Next(s, s')
  requires 0 < j < |s.tls.config|
  requires j + 1 == epoch
  requires 1 < epoch <= |s.tls.config|
  requires PerfInvariantLockInNetwork(s, j, epoch)

  ensures PerfInvariant(s')
{
  if !s.tls.t_environment.nextStep.LEnvStepHostIos? {
    NotHostIos_InvLockInNetworkImpliesInvLockInNetwork(j, epoch, s, s');
  } else {
    if s.tls.t_environment.nextStep.actor == s.tls.config[j] {
      if s.tls.t_environment.nextStep.nodeStep == GrantStep {
        //reveal_PerfInvariantLockInNetwork();
        assert false;
      } else {
        Accept_j_InvLockInNetworkImpliesInvLockHeld(j, epoch, s, s');
      }
    } else {
      if s.tls.t_environment.nextStep.nodeStep == GrantStep {
        //reveal_PerfInvariantLockInNetwork();
        assert false;
      } else {
        Accept_not_j_InvLockInNetworkImpliesInvLockInNetwork(j, epoch, s, s');
      }
    }
  }
}

////////////////////////////////////////////////////////////////////////////////

predicate PerfInvariant(tgls:TaggedGLS_State)
{
  && TGLS_Consistency(tgls)
  && ( 
    || (exists epoch, j :: 0 < epoch <= |tgls.tls.config| && epoch == j + 1 && 0 <= j < |tgls.tls.config| && PerfInvariantLockHeld(tgls, j, epoch))
    || (exists epoch, j :: 1 < epoch <= |tgls.tls.config| && epoch == j + 1 && 0 < j < |tgls.tls.config| && PerfInvariantLockInNetwork(tgls, j, epoch))
    || (PerfInvariantEpochHigherThanNumServers(tgls))
  )
}

lemma NotHostIos_InvEpochHigherGoesToInvEpochHigher(j:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config|
  requires TGLS_Consistency(s) && TGLS_Consistency(s')
  requires SingleGLSPerformanceAssumption(s)

  // A node other than node j taking an accept step
  requires !s.tls.t_environment.nextStep.LEnvStepHostIos?
  requires PerfInvariantEpochHigherThanNumServers(s)

  ensures PerfInvariantEpochHigherThanNumServers(s')
{
  lemma_mod_auto(|s.tls.config|);
}

lemma Grant_j_InvEpochHigherGoesToInvEpochHigher(j:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config|
  requires TGLS_Consistency(s) && TGLS_Consistency(s')
  requires SingleGLSPerformanceAssumption(s)

  // A node other than node j taking an accept step
  requires s.tls.t_environment.nextStep.LEnvStepHostIos?
  requires s.tls.t_environment.nextStep.actor == s.tls.config[j]
  requires s.tls.t_environment.nextStep.nodeStep == GrantStep
  requires PerfInvariantEpochHigherThanNumServers(s)

  ensures PerfInvariantEpochHigherThanNumServers(s')
{
  lemma_mod_auto(|s.tls.config|);
}

lemma PerfInvariantEpochHigherGoesToPerfInvariant(s:TaggedGLS_State, s':TaggedGLS_State)
  requires SingleGLSPerformanceAssumption(s) && TGLS_Consistency(s)
  requires SingleGLSPerformanceAssumption(s') && TGLS_Consistency(s')
  requires TGLS_Next(s, s')
  requires PerfInvariantEpochHigherThanNumServers(s)

  ensures PerfInvariantEpochHigherThanNumServers(s')
  ensures PerfInvariant(s')
{


  if s.tls.t_environment.nextStep.LEnvStepHostIos? &&
    s.tls.t_environment.nextStep.nodeStep == AcceptStep
  {
    lemma_mod_auto(|s.tls.config|);
    var id := s.tls.t_environment.nextStep.actor;
    var ios := s.tls.t_environment.nextStep.ios;
    assert IsValidLIoOp(ios[0], id, s.tls.t_environment);
  }
}

lemma PerfInvariantLockHeldGoesToPerfInvariant(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires SingleGLSPerformanceAssumption(s) && TGLS_Consistency(s)
  requires SingleGLSPerformanceAssumption(s') && TGLS_Consistency(s')
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config|
  requires 0 <= epoch <= |s.tls.config|
  requires j + 1 == epoch
  requires PerfInvariantLockHeld(s, j, epoch)

  ensures PerfInvariant(s')
{
  if !s.tls.t_environment.nextStep.LEnvStepHostIos? {
    NotHostIos_InvLockHeldImpliesInvLockHeld(j, epoch, s, s');
  } else {
    if s.tls.t_environment.nextStep.actor == s.tls.config[j] {
      if s.tls.t_environment.nextStep.nodeStep == GrantStep {
        if j == |s.tls.config| - 1 {
            Grant_LastNode_InvLockHeldImpliesInvEpochHigher(j, epoch, s, s');
        } else {
            Grant_j_InvLockHeldImpliesInvLockInNetwork(j, epoch, s, s');
        }
      } else {
        //reveal_PerfInvariantLockHeld();
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

lemma PerfInvariantMaintained(s:TaggedGLS_State, s':TaggedGLS_State)
  requires SingleGLSPerformanceAssumption(s) && TGLS_Consistency(s)
  requires SingleGLSPerformanceAssumption(s') && TGLS_Consistency(s')
  requires TGLS_Next(s, s')
  requires PerfInvariant(s)

  ensures PerfInvariant(s')
{
  if (PerfInvariantEpochHigherThanNumServers(s)) {
    PerfInvariantEpochHigherGoesToPerfInvariant(s, s');
  }
  else if (exists j, epoch :: 0 < epoch <= |s.tls.config| && epoch == j + 1 && 0 <= j < |s.tls.config| && PerfInvariantLockHeld(s, j, epoch)) {
    var epoch, j :| 0 < epoch <= |s.tls.config| && epoch == j + 1 && 0 <= j < |s.tls.config| && PerfInvariantLockHeld(s, j, epoch);
    PerfInvariantLockHeldGoesToPerfInvariant(j, epoch, s, s');
  }
  else {
    var epoch, j :| 1 < epoch <= |s.tls.config| && epoch == j + 1 && 0 < j < |s.tls.config| && PerfInvariantLockInNetwork(s, j, epoch);
    PerfInvariantLockInNetworkGoesToPerfInvariant(j, epoch, s, s');
  }
}

lemma Establish_TGLS_Consistency(config:Config, tglb:seq<TaggedGLS_State>, i:int)
  requires ValidTaggedGLSBehavior(tglb, config)
  requires 0 <= i < |tglb|
  ensures TGLS_Consistency(tglb[i])

  decreases i
{
  if i > 0{
    Establish_TGLS_Consistency(config, tglb, i - 1);
  }
}

lemma InitImpliesPerfInvariant(config:Config, s:TaggedGLS_State)
  requires TGLS_Init(s, config)
  requires TGLS_Consistency(s);
  ensures PerfInvariant(s)
{
  //reveal_PerfInvariantLockHeld();
  // PerfProperties();
  assert PerfInvariantLockHeld(s, 0, 1);
}

lemma PerfInvariantImpliesPerfGuarantee(s:TaggedGLS_State)
  requires PerfInvariant(s)
  ensures SingleGLSPerformanceGuarantee(s)
{
  //reveal_PerfInvariantLockHeld();
  //reveal_PerfInvariantLockInNetwork();
}

lemma PerformanceGuaranteeHolds(config:Config, tglb:seq<TaggedGLS_State>)
  requires ValidTaggedGLSBehavior(tglb, config)
  requires GLSPerformanceAssumption(tglb)
  ensures GLSPerformanceGuarantee(tglb)
{
  var i := 0;
  Establish_TGLS_Consistency(config, tglb, 0);
  InitImpliesPerfInvariant(config, tglb[0]);

  while i < |tglb| - 1
    invariant 0 <= i < |tglb|
    invariant GLSPerformanceGuarantee(tglb[..i+1])
    invariant PerfInvariant(tglb[i])
  {
    Establish_TGLS_Consistency(config, tglb, i);
    Establish_TGLS_Consistency(config, tglb, i + 1);
    PerfInvariantMaintained(tglb[i], tglb[i + 1]);
    PerfInvariantImpliesPerfGuarantee(tglb[i+1]);
    i := i + 1;
  }
}

}
