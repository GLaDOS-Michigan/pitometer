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

predicate NoTransferPackets(pkts:multiset<TimestampedLPacket<EndPoint, LockMessage>>)
{
  (forall pkt :: pkt in pkts && pkt.msg.v.Transfer? ==> false)
}

predicate PerfInvariant_LockHeld(tgls:TimestampedGLS_State)
{
  |tgls.history| > 0
  && (forall h :: h in tgls.history ==> h in tgls.tls.t_servers)
  && tgls.tls.t_servers[tgls.history[|tgls.history| - 1]].v.held
  && tgls.tls.t_servers[tgls.history[|tgls.history| - 1]].v.epoch == |tgls.history|

  && (forall id :: id in tgls.tls.t_servers && id != tgls.history[|tgls.history| - 1]
  ==> !tgls.tls.t_servers[id].v.held
          && tgls.tls.t_servers[id].v.epoch < |tgls.history|
  )
  && NoTransferPackets(tgls.tls.undeliveredPackets)
}


predicate PerfInvariant_LockNotHeld(tgls:TimestampedGLS_State)
{
  |tgls.history| > 0
  && (forall id :: id in tgls.tls.t_servers
      ==> !tgls.tls.t_servers[id].v.held
          && tgls.tls.t_servers[id].v.epoch < |tgls.history|
  )

  && (forall pkt :: pkt in tgls.tls.undeliveredPackets && pkt.msg.v.Transfer?
      ==> pkt.dst == tgls.history[|tgls.history| - 1]
      && pkt.src in tgls.tls.t_servers
      && pkt.msg.v.transfer_epoch == |tgls.history|
      && (var undeliveredPackets' := tgls.tls.undeliveredPackets - multiset({pkt});
          NoTransferPackets(undeliveredPackets')
        )
  )
}

predicate PerfInvariantAlways_Alt(tgls:TimestampedGLS_State)
{
  |tgls.history| > 0
  && (forall h :: h in tgls.history ==> h in tgls.tls.t_servers)

  // No Invalid packets are ever sent
  && (forall pkt :: pkt in tgls.tls.undeliveredPackets
  ==>
  !pkt.msg.v.Invalid?)

  && (forall id :: id in tgls.tls.t_servers ==> var t_hs := tgls.tls.t_servers[id];
      && (t_hs.v.held == false ==> 0 <= t_hs.v.epoch < |tgls.history| && TimeLe(t_hs.ts, PerfBoundLockInNetwork(t_hs.v.epoch + 1)))
      && (t_hs.v.held == true  ==> 0 < t_hs.v.epoch == |tgls.history| && TimeLe(t_hs.ts, PerfBoundLockHeld(t_hs.v.epoch)))
  )

  && (forall pkt :: pkt in tgls.tls.undeliveredPackets && pkt.msg.v.Transfer?
  ==> 1 < pkt.msg.v.transfer_epoch && TimeLe(pkt.msg.ts, PerfBoundLockInNetwork(pkt.msg.v.transfer_epoch))
  )

  && (if tgls.tls.t_servers[tgls.history[|tgls.history| - 1]].v.held then
    PerfInvariant_LockHeld(tgls)
  else
    PerfInvariant_LockNotHeld(tgls)
    )
}

lemma NotHostIos_InvariantMaintained(s:TimestampedGLS_State, s':TimestampedGLS_State)
  requires TGLS_Next(s, s')
  requires TGLS_Consistency(s) && TGLS_Consistency(s')
  requires SingleGLSPerformanceAssumption(s)

  requires !s.tls.t_environment.nextStep.LEnvStepHostIos?

  requires PerfInvariantAlways_Alt(s);
  ensures PerfInvariantAlways_Alt(s');
{
  
}

lemma Grant_InvariantMaintained(s:TimestampedGLS_State, s':TimestampedGLS_State)
  requires TGLS_Next(s, s')
  requires TGLS_Consistency(s) && TGLS_Consistency(s')
  requires SingleGLSPerformanceAssumption(s)

  requires s.tls.t_environment.nextStep.LEnvStepHostIos?
  requires s.tls.t_environment.nextStep.nodeStep == GrantStep

  requires PerfInvariantAlways_Alt(s);
  ensures PerfInvariantAlways_Alt(s');
{
  var t_hs := s.tls.t_servers[s.tls.t_environment.nextStep.actor];
  lemma_Grant(t_hs.ts, t_hs.v.epoch);
}


lemma Accept_InvariantMaintained(s:TimestampedGLS_State, s':TimestampedGLS_State)
  requires TGLS_Next(s, s')
  requires TGLS_Consistency(s) && TGLS_Consistency(s')
  requires SingleGLSPerformanceAssumption(s)

  requires s.tls.t_environment.nextStep.LEnvStepHostIos?
  requires s.tls.t_environment.nextStep.nodeStep == AcceptStep

  requires PerfInvariantAlways_Alt(s);
  ensures PerfInvariantAlways_Alt(s');
{
  var id := s.tls.t_environment.nextStep.actor;
  var t_hs := s.tls.t_servers[id];
  var ios := s.tls.t_environment.nextStep.ios;
  assert IsValidLIoOp(ios[0], id, s.tls.t_environment);
  var pkt := ios[0].r;
  assert pkt.msg.v.Transfer?;

  // These two lines are not necessary, but they seem to speed up verification of this lemma
  var s_pkt := ios[1].s;
  assert !s_pkt.msg.v.Transfer?;

  lemma_Accept(s.tls.t_servers[id].ts, pkt.msg.ts, s.tls.t_servers[id].v.epoch, pkt.msg.v.transfer_epoch);
}

lemma Init_InvariantHolds(s:TimestampedGLS_State)
  requires TGLS_Init(s, s.tls.config)
  requires TGLS_Consistency(s)
  ensures PerfInvariantAlways_Alt(s)
{
}

}
