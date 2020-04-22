include "../../../Services/Lock/LockTaggedDistributedSystem.i.dfy"
include "TaggedGLS.i.dfy"
include "Definitions_arith.i.dfy"
include "../../../../Libraries/Math/mod_auto.i.dfy"

module MultiRoundPerformanceProof_i {

import opened LockTaggedDistributedSystem_i
  import opened TaggedGLS_i
  import opened PerformanceProof__Definitions_i
  import opened Math__mod_auto_i

predicate {:verify false} SingleGLSPerformanceAssumption(tgls:TaggedGLS_State)
{
  // The only nodes that take steps are in the ring
  && (tgls.tls.t_environment.nextStep.LEnvStepHostIos? ==> tgls.tls.t_environment.nextStep.actor in tgls.tls.t_servers)

  // The size of the history/epochs never reaches UINT64_MAX
  && (|tgls.history| < 0xFFFF_FFFF_FFFF_FFFF)

  // No timeouts
  && (var nextStep := tgls.tls.t_environment.nextStep; nextStep.LEnvStepHostIos? ==>
  (forall io :: io in nextStep.ios ==> !io.LIoOpTimeoutReceive?))
  
  // No 'Locked' packets are ever received
  && (var nextStep := tgls.tls.t_environment.nextStep; nextStep.LEnvStepHostIos? ==>
  (forall io :: io in nextStep.ios && io.LIoOpReceive? ==> io.r.msg.v.Locked? == false))

  // TODO: This should be part of the state machine definition, and not an assumption
  // The only nodes that can do NodeGrant() currently hold the lock
  && (var nextStep := tgls.tls.t_environment.nextStep; nextStep.LEnvStepHostIos? ==> (nextStep.nodeStep == GrantStep <==> tgls.tls.t_servers[nextStep.actor].v.held == true))

}

predicate {:verify false} GLSPerformanceAssumption(tglb:seq<TaggedGLS_State>)
{
  forall tgls :: tgls in tglb ==> SingleGLSPerformanceAssumption(tgls)
}

predicate {:verify false} SingleGLSPerformanceGuarantee(gls:TaggedGLS_State)
{
  |gls.tls.config| > 1 ==> 
  (forall pkt :: pkt in gls.tls.t_environment.sentPackets &&
    pkt.msg.v.Transfer? ==> pkt.msg.v.transfer_epoch > 1 && PerfLe(pkt.msg.pr, PerfBoundLockInNetwork(pkt.msg.v.transfer_epoch)))
}

predicate {:verify false} GLSPerformanceGuarantee(tglb:seq<TaggedGLS_State>)
{
  forall tgls :: tgls in tglb ==> SingleGLSPerformanceGuarantee(tgls)
}

predicate {:verify false} TGLS_Consistency(tgls: TaggedGLS_State)
{
  && (forall id :: id in tgls.tls.config <==> id in tgls.tls.t_servers)
    && (forall id :: id in tgls.tls.t_servers ==> (tgls.tls.t_servers[id].v.config == tgls.tls.config))
    && (forall i :: 0 <= i < |tgls.tls.config| ==> tgls.tls.t_servers[tgls.tls.config[i]].v.my_index == i)
}

predicate PerfInvariantAlways_NodeNotHeld(tgls:TaggedGLS_State, t_hs:TaggedType<Node>)
{
  // || (forall pkt in env.sentPackets )
    || (false)
}

predicate PerfInvariantAlways_Node(t_hs:TaggedType<Node>)
{
  && (t_hs.v.epoch >= 0)
  && (t_hs.v.held == false ==> t_hs.v.epoch > 0 && PerfLe(t_hs.pr, PerfBoundLockInNetwork(t_hs.v.epoch + 1)))
  && (t_hs.v.held == true  ==> t_hs.v.epoch > 0 && PerfLe(t_hs.pr, PerfBoundLockHeld(t_hs.v.epoch)))
}

predicate {:verify false} PerfInvariantAlways(tgls:TaggedGLS_State)
  requires TGLS_Consistency(tgls)
{
  // No irrelevant packets in sentPackets
  && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && !(pkt.src in tgls.tls.t_servers)
  ==> false)

  // No nodes with epoch higher
  &&  (forall id :: id in tgls.tls.t_servers ==> 0 <= tgls.tls.t_servers[id].v.epoch <= |tgls.history|)

  // No transfer packets with epoch higher than history size
  // All transfer packets have correct performance object, determined by its transfer_epoch
  && (forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Transfer? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
  ==> 0 < pkt.msg.v.transfer_epoch ) // <= |tgls.history|)

  && (forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Locked? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
  ==> 0 < pkt.msg.v.locked_epoch ) // <= |tgls.history|)

  && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Transfer? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
  ==> pkt.msg.v.transfer_epoch > 1 && PerfLe(pkt.msg.pr, PerfBoundLockInNetwork(pkt.msg.v.transfer_epoch)))

  && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Locked? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
  ==> pkt.msg.v.locked_epoch > 1 && PerfLe(pkt.msg.pr, PerfBoundLockHeld(pkt.msg.v.locked_epoch)))

  && (forall id :: id in tgls.tls.t_servers ==> PerfInvariantAlways_Node(tgls.tls.t_servers[id]))
}

predicate {:verify false} PerfInvariantLockHeld(tgls: TaggedGLS_State, j:int, epoch:int)
  requires 0 < epoch
  requires 0 <= j < |tgls.tls.config|
  requires TGLS_Consistency(tgls)
{
  && PerfInvariantAlways(tgls)
  && SingleGLSPerformanceGuarantee(tgls)
  && tgls.tls.t_servers[tgls.tls.config[j]].v.held == true
  && |tgls.history| == epoch
  && tgls.tls.t_servers[tgls.tls.config[j]].v.epoch == |tgls.history|
  && j == (epoch - 1) % |tgls.tls.config|

  && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Transfer? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
  ==> && pkt.msg.v.transfer_epoch > 1 && PerfLe(pkt.msg.pr, PerfBoundLockInNetwork(pkt.msg.v.transfer_epoch)))

  && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Locked? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
  ==> pkt.msg.v.locked_epoch > 1 && PerfLe(pkt.msg.pr, PerfBoundLockHeld(pkt.msg.v.locked_epoch)))

  && (forall id :: id in tgls.tls.t_servers && tgls.tls.t_servers[id].v.held == true ==> PerfLe(tgls.tls.t_servers[id].pr, PerfBoundLockHeld(tgls.tls.t_servers[id].v.epoch)) )
  // && (forall id :: id in tgls.tls.t_servers && tgls.tls.t_servers[id].v.held == false ==> PerfLe(tgls.tls.t_servers[id].pr, tgls.tls.t_servers[id].v.epoch) )


  &&  (forall id :: id in tgls.tls.t_servers && id != tgls.tls.config[j]
  ==> tgls.tls.t_servers[id].v.held == false && tgls.tls.t_servers[id].v.epoch < |tgls.history|)

  // The node with the lock has correct PerfReport
  && PerfEq(tgls.tls.t_servers[tgls.tls.config[j]].pr, PerfBoundLockHeld(epoch))
}

predicate {:verify false} PerfInvariantLockInNetwork(tgls: TaggedGLS_State, j:int, epoch:int)
  requires 1 < epoch
  requires 0 <= j < |tgls.tls.config|
  requires TGLS_Consistency(tgls)
{
  && PerfInvariantAlways(tgls)
  && SingleGLSPerformanceGuarantee(tgls)
  && |tgls.history| == epoch
  && j == (epoch - 1) % |tgls.tls.config|


  && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Transfer? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
  ==> pkt.msg.v.transfer_epoch > 1 && PerfLe(pkt.msg.pr, PerfBoundLockInNetwork(pkt.msg.v.transfer_epoch)))

  && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Locked? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
  ==> pkt.msg.v.locked_epoch > 1 && PerfLe(pkt.msg.pr, PerfBoundLockHeld(pkt.msg.v.locked_epoch)))

  && (forall id :: id in tgls.tls.t_servers && tgls.tls.t_servers[id].v.held == true ==> PerfLe(tgls.tls.t_servers[id].pr, PerfBoundLockHeld(tgls.tls.t_servers[id].v.epoch)) )
  // && (forall id :: id in tgls.tls.t_servers && tgls.tls.t_servers[id].v.held == false ==> PerfLe(tgls.tls.t_servers[id].pr, tgls.tls.t_servers[id].v.epoch) )


  // No one holds the lock; everyone's epoch is below the epoch of the newest packet.
  && (forall id :: id in tgls.tls.t_servers ==> tgls.tls.t_servers[id].v.held == false && tgls.tls.t_servers[id].v.epoch < |tgls.history|)

  && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && pkt.msg.v.Transfer? && pkt.dst in tgls.tls.t_servers && pkt.src in tgls.tls.t_servers
      && pkt.msg.v.transfer_epoch == |tgls.history|
      ==> pkt.dst == tgls.tls.config[j] && pkt.src == tgls.tls.config[(j - 1) % |tgls.tls.config|] && PerfEq(pkt.msg.pr, PerfBoundLockInNetwork(epoch)))

}

lemma Grant_InvImpliesInv(s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires TGLS_Consistency(s) && TGLS_Consistency(s')
  requires SingleGLSPerformanceAssumption(s)
  requires SingleGLSPerformanceAssumption(s')

  // A node other than node j taking a grant step
  requires s.tls.t_environment.nextStep.LEnvStepHostIos?
  requires s.tls.t_environment.nextStep.nodeStep == GrantStep
  requires s.tls.t_servers[s.tls.t_environment.nextStep.actor].v.held == true

  requires PerfInvariantAlways(s);
  ensures PerfInvariantAlways(s');
{
  lemma_mod_auto(|s.tls.config|);
  var step := s.tls.t_environment.nextStep;
  var id := step.actor;
  var pkt := step.ios[0].s;
  var epoch := s.tls.t_servers[id].v.epoch;
  var node_pr := s.tls.t_servers[id].pr;
  // assert node_pr == PerfBoundLockHeld(epoch);
  lemma_Grant(node_pr, epoch);
  assert PerfInvariantAlways_Node(s'.tls.t_servers[id]);
}

lemma Accept_InvImpliesInv(s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires TGLS_Consistency(s) && TGLS_Consistency(s')
  requires SingleGLSPerformanceAssumption(s)
  requires SingleGLSPerformanceAssumption(s')

  // A node other than node j taking a grant step
  requires s.tls.t_environment.nextStep.LEnvStepHostIos?
  requires s.tls.t_environment.nextStep.nodeStep == AcceptStep
  // requires s.tls.t_servers[s.tls.t_environment.nextStep.actor].v.held == true

  requires PerfInvariantAlways(s);
  ensures PerfInvariantAlways(s');
{
  lemma_mod_auto(|s.tls.config|);
  var step := s.tls.t_environment.nextStep;
  var id := step.actor;
  var ios := step.ios;
  var pkt := ios[0].r;
  var epoch := s.tls.t_servers[id].v.epoch;
  var node_pr := s.tls.t_servers[id].pr;
  assert pkt.msg.v.Transfer?;

  if (pkt.msg.v.transfer_epoch > s.tls.t_servers[id].v.epoch) {
    assert PerfInvariantAlways(s');
  }
}

lemma {:verify false} NotHostIos_InvLockHeldImpliesInvLockHeld(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
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

lemma {:verify false} Grant_not_j_InvLockHeldImpliesInvLockHeld(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config|
  requires j == (epoch - 1) % |s.tls.config|
  requires 0 < epoch
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

lemma {:verify false} Accept_not_j_InvLockHeldImpliesInvLockHeld(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config|
  requires 0 < epoch
  requires j == (epoch - 1) % |s.tls.config|

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
  // assert false;
  //reveal_PerfInvariantLockHeld();
}

lemma {:verify false} Grant_j_InvLockHeldImpliesInvLockInNetwork(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config|
  requires 0 < epoch
  requires j == (epoch - 1) % |s.tls.config|
  
  requires TGLS_Consistency(s) && TGLS_Consistency(s')
  requires SingleGLSPerformanceAssumption(s)
  requires SingleGLSPerformanceAssumption(s')

  // A node other than node j taking a grant step
  requires s.tls.t_environment.nextStep.LEnvStepHostIos?
  // requires s.tls.t_environment.nextStep.actor == s.tls.config[j]
  requires s.tls.t_environment.nextStep.nodeStep == GrantStep

  requires PerfInvariantLockHeld(s, j, epoch);
  ensures PerfInvariantLockInNetwork(s', (j + 1) % |s.tls.config|, epoch + 1);
{
  //reveal_PerfInvariantLockHeld();
  //reveal_PerfInvariantLockInNetwork();
  assert false;

  lemma_mod_auto(|s.tls.config|);
  //PerfProperties();
  //var p := PerfBoundLockHeld(epoch);
  //var p' := PerfBoundLockInNetwork(epoch + 1);
  //assert PerfEq(p', PerfAdd2(p, PerfStep(GrantStep)));
}

lemma {:verify false} NotHostIos_InvLockInNetworkImpliesInvLockInNetwork(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config|
  requires 1 < epoch
  requires j == (epoch - 1) % |s.tls.config|

  requires TGLS_Consistency(s) && TGLS_Consistency(s')

  // Not a HostIos step
  requires !s.tls.t_environment.nextStep.LEnvStepHostIos?

  requires PerfInvariantLockInNetwork(s, j, epoch);
  ensures PerfInvariantLockInNetwork(s', j, epoch);
{
  //reveal_PerfInvariantLockInNetwork();
}

lemma {:verify false} Accept_not_j_InvLockInNetworkImpliesInvLockInNetwork(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config|
  requires 1 < epoch
  requires j == (epoch - 1) % |s.tls.config|

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

  lemma_mod_auto(|s.tls.config|);
}

lemma {:verify false} Accept_j_InvLockInNetworkImpliesInvLockHeld(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config|
  requires 1 < epoch
  requires j == (epoch - 1) % |s.tls.config|

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
  //PerfProperties();

  lemma_mod_auto(|s.tls.config|);
}

lemma {:verify false} PerfInvariantLockInNetworkGoesToPerfInvariant(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires SingleGLSPerformanceAssumption(s) && TGLS_Consistency(s)
  requires SingleGLSPerformanceAssumption(s') && TGLS_Consistency(s')
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config|
  requires 1 < epoch
  requires j == (epoch - 1) % |s.tls.config|

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
        // assert false;
      } else {
        // Accept_not_j_InvLockInNetworkImpliesInvLockInNetwork(j, epoch, s, s');
      }
    }
  }
}

////////////////////////////////////////////////////////////////////////////////

predicate {:verify false} PerfInvariant(tgls:TaggedGLS_State)
{
  && TGLS_Consistency(tgls)
  && ( 
    || (exists epoch, j :: 0 < epoch && j == (epoch - 1) % |tgls.tls.config| && 0 <= j < |tgls.tls.config| && PerfInvariantLockHeld(tgls, j, epoch))
    || (exists epoch, j :: 1 < epoch && j == (epoch - 1) % |tgls.tls.config| && 0 <= j < |tgls.tls.config| && PerfInvariantLockInNetwork(tgls, j, epoch))
  )
}

lemma {:verify false} PerfInvariantLockHeldGoesToPerfInvariant(j:int, epoch:int, s:TaggedGLS_State, s':TaggedGLS_State)
  requires SingleGLSPerformanceAssumption(s) && TGLS_Consistency(s)
  requires SingleGLSPerformanceAssumption(s') && TGLS_Consistency(s')
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config|
  requires 0 < epoch
  requires j == (epoch - 1) % |s.tls.config|
  requires PerfInvariantLockHeld(s, j, epoch)

  ensures PerfInvariant(s')
{
  lemma_mod_auto(|s.tls.config|);
  if !s.tls.t_environment.nextStep.LEnvStepHostIos? {
    NotHostIos_InvLockHeldImpliesInvLockHeld(j, epoch, s, s');
  } else {
    if s.tls.t_environment.nextStep.actor == s.tls.config[j] {
      if s.tls.t_environment.nextStep.nodeStep == GrantStep {
        Grant_j_InvLockHeldImpliesInvLockInNetwork(j, epoch, s, s');
      } else {
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

lemma {:verify false} PerfInvariantMaintained(s:TaggedGLS_State, s':TaggedGLS_State)
  requires SingleGLSPerformanceAssumption(s) && TGLS_Consistency(s)
  requires SingleGLSPerformanceAssumption(s') && TGLS_Consistency(s')
  requires TGLS_Next(s, s')
  requires PerfInvariant(s)

  ensures PerfInvariant(s')
{
  if (exists j, epoch :: 0 < epoch <= |s.tls.config| && epoch == j + 1 && 0 <= j < |s.tls.config| && PerfInvariantLockHeld(s, j, epoch)) {
    var epoch, j :| 0 < epoch <= |s.tls.config| && epoch == j + 1 && 0 <= j < |s.tls.config| && PerfInvariantLockHeld(s, j, epoch);
    PerfInvariantLockHeldGoesToPerfInvariant(j, epoch, s, s');
  }
  else {
    var epoch, j :| 1 < epoch <= |s.tls.config| && epoch == j + 1 && 0 <= j < |s.tls.config| && PerfInvariantLockInNetwork(s, j, epoch);
    PerfInvariantLockInNetworkGoesToPerfInvariant(j, epoch, s, s');
  }
}

lemma {:verify false} Establish_TGLS_Consistency(config:Config, tglb:seq<TaggedGLS_State>, i:int)
  requires ValidTaggedGLSBehavior(tglb, config)
  requires 0 <= i < |tglb|
  ensures TGLS_Consistency(tglb[i])

  decreases i
{
  if i > 0{
    Establish_TGLS_Consistency(config, tglb, i - 1);
  }
}

lemma {:verify false} InitImpliesPerfInvariant(config:Config, s:TaggedGLS_State)
  requires TGLS_Init(s, config)
  requires TGLS_Consistency(s);
  ensures PerfInvariant(s)
{
  //reveal_PerfInvariantLockHeld();
  assert PerfInvariantLockHeld(s, 0, 1);
}

lemma {:verify false} PerfInvariantImpliesPerfGuarantee(s:TaggedGLS_State)
  requires PerfInvariant(s)
  ensures SingleGLSPerformanceGuarantee(s)
{
  //reveal_PerfInvariantLockHeld();
  //reveal_PerfInvariantLockInNetwork();
}

lemma {:verify false} PerformanceGuaranteeHolds(config:Config, tglb:seq<TaggedGLS_State>)
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
