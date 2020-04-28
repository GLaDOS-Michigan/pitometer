include "../../../Services/Lock/LockTimestampedDistributedSystem.i.dfy"
include "TimestampedGLS_noduplication.i.dfy"
include "Definitions.i.dfy"
include "../../../../Libraries/Math/mod_auto.i.dfy"
include "MultiRoundTimeHelpers.i.dfy"
include "../../../Common/Collections/Multisets.s.dfy"

module MultiRoundPerformanceProof_i {

import opened LockTimestampedDistributedSystem_i
  import opened TimestampedGLS_noduplication_i
  import opened PerformanceProof__Definitions_i
  import opened Math__mod_auto_i
  import opened PerformanceProof__MultiRoundTimeHelpers_i
  import opened Collections__Multisets_s

predicate NoPacketDuplication(tls:TimestampedLS_State)
{
  var nextStep := tls.t_environment.nextStep;
  nextStep.LEnvStepHostIos? ==> (
  (forall io :: io in nextStep.ios && io.LIoOpReceive? ==> io.r in tls.undeliveredPackets)
  )
}

predicate LockedPacketNeverDelivered(tls:TimestampedLS_State)
{
  var nextStep := tls.t_environment.nextStep;
  nextStep.LEnvStepHostIos? ==> (forall io :: io in nextStep.ios && io.LIoOpReceive? ==> !io.r.msg.v.Locked?)
}

predicate SingleGLSPerformanceAssumption(tgls:TimestampedGLS_State)
{
  && NoPacketDuplication(tgls.tls)
  && LockedPacketNeverDelivered(tgls.tls)

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

predicate GLSPerformanceAssumption(tglb:seq<TimestampedGLS_State>)
{
  forall tgls :: tgls in tglb ==> SingleGLSPerformanceAssumption(tgls)
}

predicate SingleGLSPerformanceGuarantee(gls:TimestampedGLS_State)
{
  |gls.tls.config| > 1 ==> 
  (forall pkt :: pkt in gls.tls.t_environment.sentPackets &&
    pkt.msg.v.Transfer? ==> pkt.msg.v.transfer_epoch > 1 && TimeLe(pkt.msg.ts, PerfBoundLockInNetwork(pkt.msg.v.transfer_epoch)))
}

predicate GLSPerformanceGuarantee(tglb:seq<TimestampedGLS_State>)
{
  forall tgls :: tgls in tglb ==> SingleGLSPerformanceGuarantee(tgls)
}

predicate TGLS_Consistency(tgls: TimestampedGLS_State)
{
  && (forall id :: id in tgls.tls.config <==> id in tgls.tls.t_servers)
    && (forall id :: id in tgls.tls.t_servers ==> (tgls.tls.t_servers[id].v.config == tgls.tls.config))
    && (forall i :: 0 <= i < |tgls.tls.config| ==> tgls.tls.t_servers[tgls.tls.config[i]].v.my_index == i)
}

predicate NoTransferPacketsToId(config:ConcreteConfiguration, pkts:multiset<TimestampedLPacket<EndPoint, LockMessage>>, id:EndPoint)
  // ensures forall pkt :: !pkt.Transfer? ==>
{
  (forall pkt :: pkt in pkts && pkt.msg.v.Transfer? && pkt.dst == id && pkt.src in config ==> false)
}

predicate PerfInvariantAlways_Node(tgls:TimestampedGLS_State, id:EndPoint)
  requires id in tgls.tls.t_servers
{
  var t_hs := tgls.tls.t_servers[id];
  (
  && (t_hs.v.epoch >= 0)
  && (t_hs.v.held == false ==> t_hs.v.epoch > 0 && TimeLe(t_hs.ts, PerfBoundLockInNetwork(t_hs.v.epoch + 1)))
  && (t_hs.v.held == true  ==> t_hs.v.epoch > 0 && TimeLe(t_hs.ts, PerfBoundLockHeld(t_hs.v.epoch)))
  )

  && (
  if t_hs.v.held then
    NoTransferPacketsToId(tgls.tls.config, tgls.tls.undeliveredPackets, id)
  else
     (forall pkt :: pkt in tgls.tls.undeliveredPackets && pkt.msg.v.Transfer? && pkt.dst == id && pkt.src in tgls.tls.t_servers ==>
     (var undeliveredPackets' := tgls.tls.undeliveredPackets - multiset({pkt});
        NoTransferPacketsToId(tgls.tls.config, undeliveredPackets', id)
        )

    )
  )
}

predicate PerfInvariantAlways(tgls:TimestampedGLS_State)
  requires TGLS_Consistency(tgls)
{
  // No irrelevant packets in sentPackets
  && ( forall pkt :: pkt in tgls.tls.t_environment.sentPackets && !(pkt.src in tgls.tls.t_servers)
  ==> false)

  // No nodes with epoch higher
  &&  (forall id :: id in tgls.tls.t_servers ==> 0 <= tgls.tls.t_servers[id].v.epoch <= |tgls.history|)

  // No Invalid packets are ever sent
  && (forall pkt :: pkt in tgls.tls.undeliveredPackets
  ==>
  !pkt.msg.v.Invalid?)

  // No transfer packets with epoch higher than history size
  // All transfer packets have correct performance object, determined by its transfer_epoch
  // All undelivered transfer packets are acceptable
  && (forall pkt :: pkt in tgls.tls.undeliveredPackets && pkt.msg.v.Transfer? && pkt.dst in tgls.tls.t_servers
  ==> && pkt.src in tgls.tls.t_servers
  && 1 < pkt.msg.v.transfer_epoch <= |tgls.history| && TimeLe(pkt.msg.ts, PerfBoundLockInNetwork(pkt.msg.v.transfer_epoch))
  && pkt.msg.v.transfer_epoch > tgls.tls.t_servers[pkt.dst].v.epoch
  )

  && (forall id :: id in tgls.tls.t_servers ==> PerfInvariantAlways_Node(tgls, id))
}


lemma Grant_j_InvLockHeldImpliesInvLockHeld(j:int, s:TimestampedGLS_State, s':TimestampedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config|
  requires TGLS_Consistency(s) && TGLS_Consistency(s')
  requires SingleGLSPerformanceAssumption(s)

  // A node other than node j taking a grant step
  requires s.tls.t_environment.nextStep.LEnvStepHostIos?
  requires s.tls.t_environment.nextStep.actor == s.tls.config[j]
  requires s.tls.t_environment.nextStep.nodeStep == GrantStep

  requires PerfInvariantAlways(s);
  ensures PerfInvariantAlways(s');
{
  var t_hs := s.tls.t_servers[s.tls.config[j]];
  lemma_Grant(t_hs.ts, t_hs.v.epoch);
  assert PerfInvariantAlways_Node(s', s.tls.config[j]);
  // assert forall id :: id in s.tls.t_servers ==> PerfInvariant
}

lemma Accept_j_InvLockHeldImpliesInvLockHeld(j:int, s:TimestampedGLS_State, s':TimestampedGLS_State)
  requires TGLS_Next(s, s')
  requires 0 <= j < |s.tls.config|
  requires TGLS_Consistency(s) && TGLS_Consistency(s')
  requires SingleGLSPerformanceAssumption(s)

  // A node other than node j taking a grant step
  requires s.tls.t_environment.nextStep.LEnvStepHostIos?
  requires s.tls.t_environment.nextStep.actor == s.tls.config[j]
  requires s.tls.t_environment.nextStep.nodeStep == AcceptStep

  requires PerfInvariantAlways(s);
  ensures PerfInvariantAlways(s');
{
  var t_hs := s.tls.t_servers[s.tls.config[j]];
  // lemma_Grant(t_hs.ts, t_hs.v.epoch);
  var ios := s.tls.t_environment.nextStep.ios;
  var id := s.tls.t_environment.nextStep.actor;
  assert IsValidLIoOp(ios[0], id, s.tls.t_environment);
  var pkt := ios[0].r;
  assert pkt in s.tls.undeliveredPackets;
  match pkt.msg.v {
    case Locked(_) =>
      assert false;
    case Transfer(_) =>
      assert pkt.msg.v.Transfer?;
  }
  assert !(pkt.msg.v.Locked?);
  assert pkt.msg.v.Transfer?;
  assert s'.tls.t_servers[s.tls.config[j]].v.held == true;
  var undeliveredPackets' := s.tls.undeliveredPackets - multiset{pkt};
  var s_pkt := ios[1].s;
  assert !s_pkt.msg.v.Transfer?;
  assert NoTransferPacketsToId(s.tls.config, undeliveredPackets', id);
  assert s'.tls.undeliveredPackets == undeliveredPackets' + multiset{s_pkt};
  assert NoTransferPacketsToId(s.tls.config, s'.tls.undeliveredPackets, id);
  lemma_Accept(s.tls.t_servers[id].ts, pkt.msg.ts, s.tls.t_servers[id].v.epoch, pkt.msg.v.transfer_epoch);

  assert PerfInvariantAlways_Node(s', s.tls.config[j]);
}


}
