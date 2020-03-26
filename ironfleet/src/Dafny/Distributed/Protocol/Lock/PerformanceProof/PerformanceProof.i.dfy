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
  |gls.tls.config| > 0 ==> 
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
  &&  (forall id :: id in tgls.tls.t_servers ==> 0 <= tgls.tls.t_servers[id].v.epoch <= |tgls.history|)

  // No transfer packets with epoch higher than history size
  // All transfer packets have correct performance object, determined by its transfer_epoch
  && (forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Transfer? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
  ==> 0 < pkt.msg.v.transfer_epoch <= |tgls.history|)

  && (forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Locked? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
  ==> 0 < pkt.msg.v.locked_epoch <= |tgls.history|)
}

predicate PerfInvariantLockHeld(tgls: TaggedGLS_State, j:int, epoch:int)
  requires 0 <= epoch
  requires 0 <= j < |tgls.tls.config|
  requires TGLS_Consistency(tgls)
{
    && PerfInvariantAlways(tgls)
    && SingleGLSPerformanceGuarantee(tgls)
    && tgls.tls.t_servers[tgls.tls.config[j]].v.held == true
    && |tgls.history| == epoch
    && tgls.tls.t_servers[tgls.tls.config[j]].v.epoch == |tgls.history|
    && j == (epoch % |tgls.tls.config|)

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
  requires 0 < epoch
  requires 0 < j < |tgls.tls.config|
  requires TGLS_Consistency(tgls)
{
  && PerfInvariantAlways(tgls)
  && SingleGLSPerformanceGuarantee(tgls)
  && |tgls.history| == epoch
  && j == (epoch) % |tgls.tls.config|

  // No one holds the lock; everyone's epoch is below the epoch of the newest packet.
  && (forall id :: id in tgls.tls.t_servers ==> tgls.tls.t_servers[id].v.held == false && tgls.tls.t_servers[id].v.epoch < |tgls.history|)

  && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Transfer? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
  && pkt.dst != tgls.tls.config[j % |tgls.tls.config|]
  ==> pkt.msg.v.transfer_epoch < |tgls.history| && pkt.msg.v.transfer_epoch <= tgls.tls.t_servers[pkt.dst].v.epoch)

  && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Transfer? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
      && pkt.msg.v.transfer_epoch == |tgls.history|
      ==> pkt.dst == tgls.tls.config[j] && pkt.src == tgls.tls.config[j - 1] && PerfEq(pkt.msg.pr, PerfBoundLockInNetwork(epoch)))

  // All nodes beyond j have void PerfReport
  && (forall k :: j <= k < |tgls.tls.config| ==> tgls.tls.t_servers[tgls.tls.config[k]].pr == PerfVoid())

  // The only packet sent to node j is an acceptable Transfer message
  && (forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.dst in tgls.tls.t_servers && pkt.dst == tgls.tls.config[j % |tgls.tls.config|]
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
  requires 0 <= epoch
  requires TGLS_Consistency(s) && TGLS_Consistency(s')

  // Not a HostIos step
  requires !s.tls.t_environment.nextStep.LEnvStepHostIos?

  requires PerfInvariantLockHeld(s, j, epoch);
  ensures PerfInvariantLockHeld(s', j, epoch);
{
  //revealPerfInvariantLockHeld();
}

lemma Grant_not_j_InvLockHeldImpliesInvLockHeld(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config|
  requires 0 <= epoch < |s.tls.config|
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
  //revealPerfInvariantLockHeld();
}

lemma Accept_not_j_InvLockHeldImpliesInvLockHeld(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config|
  requires 0 <= epoch < |s.tls.config|
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
  //revealPerfInvariantLockHeld();
}

lemma Grant_j_InvLockHeldImpliesInvLockInNetwork(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config| - 1
  requires 0 <= epoch < |s.tls.config| - 1
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
  //revealPerfInvariantLockHeld();
  //revealPerfInvariantLockInNetwork();

  lemma_mod_auto(|s.tls.config|);
  PerfProperties();
  var p := PerfBoundLockHeld(epoch);
  var p' := PerfBoundLockInNetwork(epoch + 1);
  assert PerfEq(p', PerfAdd2(p, PerfStep(GrantStep)));
}

lemma Grant_LastNode_InvLockHeldImpliesInvEpochHigher(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 <= j == |s.tls.config| - 1
  requires 0 <= epoch == |s.tls.config| - 1
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
  //revealPerfInvariantLockHeld();
  //revealPerfInvariantLockInNetwork();

  lemma_mod_auto(|s.tls.config|);
  PerfProperties();
  var p := PerfBoundLockHeld(epoch);
  var p' := PerfBoundLockInNetwork(epoch + 1);
  assert PerfEq(p', PerfAdd2(p, PerfStep(GrantStep)));
}

lemma NotHostIos_InvLockInNetworkImpliesInvLockInNetwork(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 < j < |s.tls.config|
  requires 0 < epoch < |s.tls.config|
  requires TGLS_Consistency(s) && TGLS_Consistency(s')

  // Not a HostIos step
  requires !s.tls.t_environment.nextStep.LEnvStepHostIos?

  requires PerfInvariantLockInNetwork(s, j, epoch);
  ensures PerfInvariantLockInNetwork(s', j, epoch);
{
  //revealPerfInvariantLockInNetwork();
}

lemma Accept_not_j_InvLockInNetworkImpliesInvLockInNetwork(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 < j < |s.tls.config|
  requires 0 < epoch < |s.tls.config|
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
  //revealPerfInvariantLockInNetwork();

  lemma_mod_auto(|s.tls.config|);
}

lemma Accept_j_InvLockInNetworkImpliesInvLockHeld(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 < j < |s.tls.config|
  requires 0 < epoch < |s.tls.config|
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
  //revealPerfInvariantLockInNetwork();
  //revealPerfInvariantLockHeld();
  PerfProperties();

  lemma_mod_auto(|s.tls.config|);
  var p2 := PerfBoundLockInNetwork(epoch);
  var p2' := PerfBoundLockHeld(epoch);

  assert PerfEq(p2', PerfAdd2(p2, PerfStep(AcceptStep)));

  var id := s.tls.t_environment.nextStep.actor;
  var tgls := s;
  var tgls' := s';
  var ios := s.tls.t_environment.nextStep.ios;
  assert ios[1].s.dst == s.tls.config[j - 1];
  assert ios[1].s.msg.v.locked_epoch == epoch;

  assert s.tls.t_servers[id].pr == PerfVoid;
  assert GetReceivePRs(ios) == [ios[0].r.msg.pr];
  assert multiset(GetReceivePRs(ios)) + multiset{s.tls.t_servers[id].pr} == multiset{ios[0].r.msg.pr, PerfVoid};
  var recvTime := PerfMax(multiset(GetReceivePRs(ios)) + multiset{s.tls.t_servers[id].pr});
  assert multiset{ios[0].r.msg.pr, PerfVoid} - multiset{PerfVoid} == multiset{ios[0].r.msg.pr};
  assert PerfEq(PerfMax(multiset{ios[0].r.msg.pr, PerfVoid}), PerfMax(multiset{ios[0].r.msg.pr}));
  assert PerfEq(recvTime, p2);

  // var totalTime := PerfAdd2(recvTime, PerfStep(hstep));
  // tls'.t_servers[id].pr == totalTime
  // assert PerfInvariantLockHeldNonOpaque(s', j);
}

lemma PerfInvariantLockInNetworkGoesToPerfInvariant(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires SingleGLSPerformanceAssumption(s) && TGLS_Consistency(s)
  requires SingleGLSPerformanceAssumption(s') && TGLS_Consistency(s')
  requires TGLS_Next(s, s')
  requires 0 < j < |s.tls.config|
  requires j == epoch
  requires 0 < epoch < |s.tls.config|
  requires PerfInvariantLockInNetwork(s, j, epoch)
  requires SingleGLSPerformanceGuarantee(s)

  ensures PerfInvariant(s')
{
  if !s.tls.t_environment.nextStep.LEnvStepHostIos? {
    NotHostIos_InvLockInNetworkImpliesInvLockInNetwork(j, epoch, s, s');
  } else {
    if s.tls.t_environment.nextStep.actor == s.tls.config[j] {
      if s.tls.t_environment.nextStep.nodeStep == GrantStep {
        //revealPerfInvariantLockInNetwork();
        assert false;
      } else {
        Accept_j_InvLockInNetworkImpliesInvLockHeld(j, epoch, s, s');
      }
    } else {
      if s.tls.t_environment.nextStep.nodeStep == GrantStep {
        //revealPerfInvariantLockInNetwork();
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
    || (exists epoch, j :: 0 <= epoch < |tgls.tls.config| && epoch == j && 0 <= j < |tgls.tls.config| && PerfInvariantLockHeld(tgls, j, epoch))
    || (exists epoch, j :: 0 < epoch < |tgls.tls.config| && epoch == j && 0 < j < |tgls.tls.config| && PerfInvariantLockInNetwork(tgls, j, epoch))
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

lemma Accept_j_InvEpochHigherGoesToInvEpochHigher(j:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config|
  requires TGLS_Consistency(s) && TGLS_Consistency(s')
  requires SingleGLSPerformanceAssumption(s)

  // A node other than node j taking an accept step
  requires s.tls.t_environment.nextStep.LEnvStepHostIos?
  requires s.tls.t_environment.nextStep.actor == s.tls.config[j]
  requires s.tls.t_environment.nextStep.nodeStep == AcceptStep
  requires PerfInvariantEpochHigherThanNumServers(s)

  ensures PerfInvariantEpochHigherThanNumServers(s')
{
  lemma_mod_auto(|s.tls.config|);
  var ios := s.tls.t_environment.nextStep.ios;
  if |ios| == 1 {
    assert PerfInvariantEpochHigherThanNumServers(s');
  } else {
    assert PerfInvariantEpochHigherThanNumServers(s');
  }
}

lemma PerfInvariantEpochHigherGoesToPerfInvariant(s:TaggedGLS_State, s':TaggedGLS_State)
  requires SingleGLSPerformanceAssumption(s) && TGLS_Consistency(s)
  requires SingleGLSPerformanceAssumption(s') && TGLS_Consistency(s')
  requires TGLS_Next(s, s')
  requires PerfInvariantEpochHigherThanNumServers(s)

  ensures PerfInvariantEpochHigherThanNumServers(s')
  ensures PerfInvariant(s')
{
  
}

lemma PerfInvariantLockHeldGoesToPerfInvariant(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires SingleGLSPerformanceAssumption(s) && TGLS_Consistency(s)
  requires SingleGLSPerformanceAssumption(s') && TGLS_Consistency(s')
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config|
  requires 0 <= epoch < |s.tls.config|
  requires j == epoch
  requires PerfInvariantLockHeld(s, j, epoch)
  requires SingleGLSPerformanceGuarantee(s)

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
        //revealPerfInvariantLockHeld();
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
  else if (exists j, epoch :: 0 <= epoch && epoch == j && 0 <= j < |s.tls.config| && PerfInvariantLockHeld(s, j, epoch)) {
    var epoch, j :| 0 <= epoch < |s.tls.config| && epoch == j && 0 <= j < |s.tls.config| && PerfInvariantLockHeld(s, j, epoch);
    PerfInvariantLockHeldGoesToPerfInvariant(j, epoch, s, s');
  }
  else {
    var epoch, j :| 0 < epoch && epoch == j && 0 < j < |s.tls.config| && PerfInvariantLockInNetwork(s, j, epoch);
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
