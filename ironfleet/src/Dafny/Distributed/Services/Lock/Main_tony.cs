// Dafny program Main.i.dfy compiled into C#
// To recompile, use 'csc' with: /r:System.Numerics.dll
// and choosing /target:exe or /target:library
// You might also want to include compiler switches like:
//     /debug /nowarn:0164 /nowarn:0219 /nowarn:1717 /nowarn:0162 /nowarn:0168
using System;
using System.Numerics;
using System.Diagnostics;
using System.Threading;
using System.Collections.Concurrent;
using System.Collections.Generic;
using FStream = System.IO.FileStream;
using UClient = System.Net.Sockets.UdpClient;
using IEndPoint = System.Net.IPEndPoint;
[assembly: DafnyAssembly.DafnySourceAttribute(@"
// Dafny 2.3.0.10506
// Command Line Options: /noCheating:1 /proverWarnings:1 /z3opt:pi.warnings=true /allowGlobals /ironDafny /warnShadowing /z3opt:nlsat.randomize=false /autoTriggers:1 /noNLarith /timeLimit:30 /spillTargetCode:1 /noVerify /compile:2 src/Dafny/Distributed/Services/Lock/Main.i.dfy
// Main.i.dfy


module Main_i refines Main_s {

  import opened DS_s = Lock_DistributedSystem_i

  import opened Environment_s = Environment_s

  import opened Concrete_NodeIdentity_i = Concrete_NodeIdentity_i

  import opened PacketParsing_i = PacketParsing_i

  import opened UdpLock_i = UdpLock_i

  import opened Host_i = Host_i

  import opened AS_s = AbstractServiceLock_s

  import opened Refinement_i = Refinement_i

  import opened RefinementProof_i = RefinementProof_i

  import opened MarshallProof_i = MarshallProof_i
  predicate IsValidBehavior(config: ConcreteConfiguration, db: seq<DS_State>)
    reads *
    decreases {}, config, db
  {
    |db| > 0 &&
    DS_Init(db[0], config) &&
    forall i: int {:trigger DS_Next(db[i], db[i + 1])} :: 
      0 <= i < |db| - 1 ==>
        DS_Next(db[i], db[i + 1])
  }

  predicate IsValidBehaviorLs(config: ConcreteConfiguration, db: seq<LS_State>)
    reads *
    decreases {}, config, db
  {
    |db| > 0 &&
    LS_Init(db[0], config) &&
    forall i: int {:trigger LS_Next(db[i], db[i + 1])} :: 
      0 <= i < |db| - 1 ==>
        LS_Next(db[i], db[i + 1])
  }

  function AbstractifyConcretePacket(p: LPacket<EndPoint, seq<byte>>): LPacket<NodeIdentity, LockMessage>
    decreases p
  {
    LPacket(p.dst, p.src, AbstractifyCMessage(DemarshallData(p.msg)))
  }

  predicate LEnvStepIsAbstractable(step: LEnvStep<EndPoint, seq<byte>, LockStep>)
    decreases step
  {
    match step {
      case LEnvStepHostIos(actor, ios, lockStep) =>
        true
      case LEnvStepDeliverPacket(p) =>
        true
      case LEnvStepAdvanceTime =>
        true
      case LEnvStepStutter =>
        true
    }
  }

  function AbstractifyConcreteEnvStep(step: LEnvStep<EndPoint, seq<byte>, LockStep>): LEnvStep<NodeIdentity, LockMessage, LockStep>
    requires LEnvStepIsAbstractable(step)
    decreases step
  {
    match step {
      case LEnvStepHostIos(actor, ios, lockStep) =>
        LEnvStepHostIos(actor, AbstractifyRawLogToIos(ios), lockStep)
      case LEnvStepDeliverPacket(p) =>
        LEnvStepDeliverPacket(AbstractifyConcretePacket(p))
      case LEnvStepAdvanceTime =>
        LEnvStepAdvanceTime()
      case LEnvStepStutter =>
        LEnvStepStutter()
    }
  }

  predicate ConcreteEnvironmentIsAbstractable(ds_env: LEnvironment<EndPoint, seq<byte>, LockStep>)
    decreases ds_env
  {
    LEnvStepIsAbstractable(ds_env.nextStep)
  }

  function AbstractifyConcreteSentPackets(sent: set<LPacket<EndPoint, seq<byte>>>): set<LPacket<NodeIdentity, LockMessage>>
    decreases sent
  {
    set p: LPacket<EndPoint, seq<byte>> {:trigger AbstractifyConcretePacket(p)} {:trigger p in sent} | p in sent :: AbstractifyConcretePacket(p)
  }

  function AbstractifyConcreteEnvironment(ds_env: LEnvironment<EndPoint, seq<byte>, LockStep>): LEnvironment<NodeIdentity, LockMessage, LockStep>
    requires ConcreteEnvironmentIsAbstractable(ds_env)
    decreases ds_env
  {
    LEnvironment(ds_env.time, AbstractifyConcreteSentPackets(ds_env.sentPackets), map[], AbstractifyConcreteEnvStep(ds_env.nextStep))
  }

  function AbstractifyConcreteReplicas(replicas: map<EndPoint, HostState>, replica_order: seq<EndPoint>): map<EndPoint, Node>
    requires forall i: int :: 0 <= i < |replica_order| ==> replica_order[i] in replicas
    requires SeqIsUnique(replica_order)
    ensures |AbstractifyConcreteReplicas(replicas, replica_order)| == |replica_order|
    ensures forall i: int :: 0 <= i < |replica_order| ==> replica_order[i] in AbstractifyConcreteReplicas(replicas, replica_order)
    ensures forall i: int :: 0 <= i < |replica_order| ==> AbstractifyConcreteReplicas(replicas, replica_order)[replica_order[i]] == replicas[replica_order[i]].node
    ensures forall e: EndPoint :: e in AbstractifyConcreteReplicas(replicas, replica_order) <==> e in replica_order
    decreases replicas, replica_order
  {
    if |replica_order| == 0 then
      map[]
    else
      lemma_UniqueSeq_SubSeqsUnique(replica_order, [replica_order[0]], replica_order[1..]); assert SeqIsUnique(replica_order[1..]); reveal_SeqIsUnique(); assert replica_order[0] !in replica_order[1..]; assert replica_order[0] !in AbstractifyConcreteReplicas(replicas, replica_order[1..]); var rest: map<EndPoint, Node> := AbstractifyConcreteReplicas(replicas, replica_order[1..]); rest[replica_order[0] := replicas[replica_order[0]].node]
  }

  function AbstractifyConcreteClients(clients: set<EndPoint>): set<NodeIdentity>
    decreases clients
  {
    set e: EndPoint {:trigger e in clients} | e in clients :: e
  }

  predicate DsStateIsAbstractable(ds: DS_State)
    decreases ds
  {
    ValidConfig(ds.config) &&
    forall r: EndPoint :: 
      r in ds.config ==>
        r in ds.servers
  }

  function AbstractifyDsState(ds: DS_State): LS_State
    requires DsStateIsAbstractable(ds)
    decreases ds
  {
    LS_State(AbstractifyConcreteEnvironment(ds.environment), AbstractifyConcreteReplicas(ds.servers, ds.config))
  }

  lemma lemma_DeduceTransitionFromDsBehavior(config: ConcreteConfiguration, db: seq<DS_State>, i: int)
    requires IsValidBehavior(config, db)
    requires 0 <= i < |db| - 1
    ensures DS_Next(db[i], db[i + 1])
    decreases config, db, i
  {
  }

  lemma lemma_DsNextOffset(db: seq<DS_State>, index: int)
    requires |db| > 0
    requires 0 < index < |db|
    requires forall i: int {:trigger DS_Next(db[i], db[i + 1])} :: 0 <= i < |db| - 1 ==> DS_Next(db[i], db[i + 1])
    ensures DS_Next(db[index - 1], db[index])
    decreases db, index
  {
    ghost var i := index - 1;
    assert DS_Next(db[i], db[i + 1]);
  }

  lemma lemma_DsConsistency(config: ConcreteConfiguration, db: seq<DS_State>, i: int)
    requires IsValidBehavior(config, db)
    requires 0 <= i < |db|
    ensures db[i].config == config
    ensures Collections__Maps2_s.mapdomain(db[i].servers) == Collections__Maps2_s.mapdomain(db[0].servers)
    decreases config, db, i
  {
    if i == 0 {
    } else {
      lemma_DsConsistency(config, db, i - 1);
      lemma_DeduceTransitionFromDsBehavior(config, db, i - 1);
      assert forall server: EndPoint :: server in db[i - 1].servers ==> server in db[i].servers;
      assert forall server: EndPoint :: server in db[i].servers ==> server in db[i - 1].servers;
      forall server: EndPoint | server in Collections__Maps2_s.mapdomain(db[i - 1].servers)
        ensures server in Collections__Maps2_s.mapdomain(db[i].servers)
      {
        assert server in db[i - 1].servers;
        assert server in db[i].servers;
      }
      forall server: EndPoint | server in Collections__Maps2_s.mapdomain(db[i].servers)
        ensures server in Collections__Maps2_s.mapdomain(db[i - 1].servers)
      {
        assert server in db[i].servers;
        assert server in db[i - 1].servers;
      }
    }
  }

  lemma lemma_IsValidEnvStep(de: LEnvironment<EndPoint, seq<byte>, LockStep>, le: LEnvironment<NodeIdentity, LockMessage, LockStep>)
    requires IsValidLEnvStep(de, de.nextStep)
    requires de.nextStep.LEnvStepHostIos?
    requires ConcreteEnvironmentIsAbstractable(de)
    requires AbstractifyConcreteEnvironment(de) == le
    ensures IsValidLEnvStep(le, le.nextStep)
    decreases de, le
  {
    ghost var id := de.nextStep.actor;
    ghost var ios := de.nextStep.ios;
    ghost var r_ios := le.nextStep.ios;
    assert LIoOpSeqCompatibleWithReduction(r_ios);
    forall io: LIoOp<NodeIdentity, LockMessage> | io in r_ios
      ensures IsValidLIoOp(io, id, le)
    {
      ghost var j :| 0 <= j < |r_ios| && r_ios[j] == io;
      assert r_ios[j] == AstractifyUdpEventToLockIo(ios[j]);
      assert IsValidLIoOp(ios[j], id, de);
    }
  }

  lemma lemma_IosRelations(ios: seq<LIoOp<EndPoint, seq<byte>>>, r_ios: seq<LIoOp<NodeIdentity, LockMessage>>)
      returns (sends: set<LPacket<EndPoint, seq<byte>>>, r_sends: set<LPacket<NodeIdentity, LockMessage>>)
    requires r_ios == AbstractifyRawLogToIos(ios)
    ensures sends == set io: LIoOp<EndPoint, seq<byte>> {:trigger io.s} {:trigger io.LIoOpSend?} {:trigger io in ios} | io in ios && io.LIoOpSend? :: io.s
    ensures r_sends == set io: LIoOp<NodeIdentity, LockMessage> {:trigger io.s} {:trigger io.LIoOpSend?} {:trigger io in r_ios} | io in r_ios && io.LIoOpSend? :: io.s
    ensures r_sends == AbstractifyConcreteSentPackets(sends)
    decreases ios, r_ios
  {
    sends := set io: LIoOp<EndPoint, seq<byte>> {:trigger io.s} {:trigger io.LIoOpSend?} {:trigger io in ios} | io in ios && io.LIoOpSend? :: io.s;
    r_sends := set io: LIoOp<NodeIdentity, LockMessage> {:trigger io.s} {:trigger io.LIoOpSend?} {:trigger io in r_ios} | io in r_ios && io.LIoOpSend? :: io.s;
    ghost var refined_sends := AbstractifyConcreteSentPackets(sends);
    forall r: LPacket<NodeIdentity, LockMessage> | r in refined_sends
      ensures r in r_sends
    {
      ghost var send :| send in sends && AbstractifyConcretePacket(send) == r;
      ghost var io :| io in ios && io.LIoOpSend? && io.s == send;
      assert AstractifyUdpEventToLockIo(io) in r_ios;
    }
    forall r: LPacket<NodeIdentity, LockMessage> | r in r_sends
      ensures r in refined_sends
    {
      ghost var r_io :| r_io in r_ios && r_io.LIoOpSend? && r_io.s == r;
      ghost var j :| 0 <= j < |r_ios| && r_ios[j] == r_io;
      assert AstractifyUdpEventToLockIo(ios[j]) == r_io;
      assert ios[j] in ios;
      assert ios[j].s in sends;
    }
  }

  lemma lemma_LEnvironmentNextHost(de: LEnvironment<EndPoint, seq<byte>, LockStep>, le: LEnvironment<NodeIdentity, LockMessage, LockStep>, de': LEnvironment<EndPoint, seq<byte>, LockStep>, le': LEnvironment<NodeIdentity, LockMessage, LockStep>)
    requires ConcreteEnvironmentIsAbstractable(de)
    requires ConcreteEnvironmentIsAbstractable(de')
    requires AbstractifyConcreteEnvironment(de) == le
    requires AbstractifyConcreteEnvironment(de') == le'
    requires de.nextStep.LEnvStepHostIos?
    requires LEnvironment_Next(de, de')
    ensures LEnvironment_Next(le, le')
    decreases de, le, de', le'
  {
    lemma_IsValidEnvStep(de, le);
    ghost var id := de.nextStep.actor;
    ghost var ios := de.nextStep.ios;
    ghost var r_ios := le.nextStep.ios;
    assert LEnvironment_PerformIos(de, de', id, ios);
    ghost var sends, r_sends := lemma_IosRelations(ios, r_ios);
    assert de.sentPackets + sends == de'.sentPackets;
    assert le.sentPackets + r_sends == le'.sentPackets;
    assert forall r_io: LIoOp<NodeIdentity, LockMessage> :: r_io in r_ios && r_io.LIoOpReceive? ==> r_io.r in le.sentPackets;
    assert LEnvironment_PerformIos(le, le', id, r_ios);
  }

  lemma {:timeLimit 60} RefinementToLSState(config: ConcreteConfiguration, db: seq<DS_State>) returns (sb: seq<LS_State>)
    requires |db| > 0
    requires DS_Init(db[0], config)
    requires forall i: int {:trigger DS_Next(db[i], db[i + 1])} :: 0 <= i < |db| - 1 ==> DS_Next(db[i], db[i + 1])
    ensures |sb| == |db|
    ensures LS_Init(sb[0], db[0].config)
    ensures forall i: int {:trigger LS_Next(sb[i], sb[i + 1])} :: 0 <= i < |sb| - 1 ==> LS_Next(sb[i], sb[i + 1])
    ensures forall i: int :: 0 <= i < |db| ==> DsStateIsAbstractable(db[i]) && sb[i] == AbstractifyDsState(db[i])
    decreases config, db
  {
    if |db| == 1 {
      ghost var ls := AbstractifyDsState(db[0]);
      sb := [ls];
      assert forall id: EndPoint :: id in db[0].servers ==> HostInit(db[0].servers[id], config, id);
      reveal_SeqIsUnique();
    } else {
      lemma_DeduceTransitionFromDsBehavior(config, db, |db| - 2);
      lemma_DsConsistency(config, db, |db| - 2);
      lemma_DsConsistency(config, db, |db| - 1);
      ghost var ls := AbstractifyDsState(db[|db| - 2]);
      ghost var ls' := AbstractifyDsState(last(db));
      ghost var rest := RefinementToLSState(config, all_but_last(db));
      sb := rest + [ls'];
      forall i: int | 0 <= i < |sb| - 1
        ensures LS_Next(sb[i], sb[i + 1])
      {
        if 0 <= i < |sb| - 2 {
          assert LS_Next(sb[i], sb[i + 1]);
        } else {
          if !db[i].environment.nextStep.LEnvStepHostIos? {
            assert LS_Next(sb[i], sb[i + 1]);
          } else {
            lemma_LEnvironmentNextHost(db[i].environment, ls.environment, db[i + 1].environment, ls'.environment);
            assert LS_Next(sb[i], sb[i + 1]);
          }
        }
      }
    }
  }

  lemma lemma_DeduceTransitionFromLsBehavior(config: ConcreteConfiguration, db: seq<LS_State>, i: int)
    requires IsValidBehaviorLs(config, db)
    requires 0 <= i < |db| - 1
    ensures LS_Next(db[i], db[i + 1])
    decreases config, db, i
  {
  }

  lemma lemma_LsConsistency(config: ConcreteConfiguration, lb: seq<LS_State>, i: int)
    requires IsValidBehaviorLs(config, lb)
    requires 0 <= i < |lb|
    ensures Collections__Maps2_s.mapdomain(lb[i].servers) == Collections__Maps2_s.mapdomain(lb[0].servers)
    ensures forall e: EndPoint :: e in lb[i].servers ==> e in lb[0].servers && lb[i].servers[e].config == lb[0].servers[e].config
    decreases config, lb, i
  {
    if i == 0 {
    } else {
      lemma_LsConsistency(config, lb, i - 1);
      lemma_DeduceTransitionFromLsBehavior(config, lb, i - 1);
      assert forall server: EndPoint :: server in lb[i - 1].servers ==> server in lb[i].servers;
      assert forall server: EndPoint :: server in lb[i].servers ==> server in lb[i - 1].servers;
      forall server: EndPoint | server in Collections__Maps2_s.mapdomain(lb[i - 1].servers)
        ensures server in Collections__Maps2_s.mapdomain(lb[i].servers)
      {
        assert server in lb[i - 1].servers;
        assert server in lb[i].servers;
      }
      forall server: EndPoint | server in Collections__Maps2_s.mapdomain(lb[i].servers)
        ensures server in Collections__Maps2_s.mapdomain(lb[i - 1].servers)
      {
        assert server in lb[i].servers;
        assert server in lb[i - 1].servers;
      }
    }
  }

  lemma {:timeLimit 60} MakeGLSBehaviorFromLS(config: ConcreteConfiguration, db: seq<LS_State>) returns (sb: seq<GLS_State>)
    requires |db| > 0
    requires LS_Init(db[0], config)
    requires forall i: int {:trigger LS_Next(db[i], db[i + 1])} :: 0 <= i < |db| - 1 ==> LS_Next(db[i], db[i + 1])
    ensures |sb| == |db|
    ensures GLS_Init(sb[0], config)
    ensures forall i: int {:trigger GLS_Next(sb[i], sb[i + 1])} :: 0 <= i < |sb| - 1 ==> GLS_Next(sb[i], sb[i + 1])
    ensures forall i: int :: 0 <= i < |db| ==> sb[i].ls == db[i]
    decreases config, db
  {
    if |db| == 1 {
      sb := [GLS_State(db[0], [config[0]])];
    } else {
      ghost var rest := MakeGLSBehaviorFromLS(config, all_but_last(db));
      ghost var last_history := last(rest).history;
      ghost var ls := db[|db| - 2];
      ghost var ls' := db[|db| - 1];
      if ls.environment.nextStep.LEnvStepHostIos? && ls.environment.nextStep.actor in ls.servers {
        ghost var id := ls.environment.nextStep.actor;
        ghost var ios := ls.environment.nextStep.ios;
        ghost var hstep := ls.environment.nextStep.nodeStep;
        lemma_DeduceTransitionFromLsBehavior(config, db, |db| - 2);
        assert LS_Next(ls, ls');
        assert LS_NextOneServer(ls, ls', id, ios, hstep);
        ghost var node := ls.servers[id];
        ghost var node' := ls'.servers[id];
        assert NodeNext(node, node', ios);
        ghost var new_history: seq<EndPoint>;
        if NodeGrant(node, node', ios) && node.held && node.epoch < 18446744073709551615 {
          new_history := last_history + [node.config[(node.my_index + 1) % |node.config|]];
        } else {
          new_history := last_history;
        }
        sb := rest + [GLS_State(db[|db| - 1], new_history)];
        assert GLS_Next(sb[|sb| - 2], sb[|sb| - 1]);
      } else {
        sb := rest + [GLS_State(db[|db| - 1], last_history)];
      }
    }
  }

  lemma {:timeLimit 60} RefinementToServiceState(config: ConcreteConfiguration, glb: seq<GLS_State>) returns (sb: seq<ServiceState>)
    requires |glb| > 0
    requires GLS_Init(glb[0], config)
    requires forall i: int {:trigger GLS_Next(glb[i], glb[i + 1])} :: 0 <= i < |glb| - 1 ==> GLS_Next(glb[i], glb[i + 1])
    ensures |sb| == |glb|
    ensures Service_Init(sb[0], MapSeqToSet(config, (x: EndPoint) => x))
    ensures forall i: int {:trigger Service_Next(sb[i], sb[i + 1])} :: 0 <= i < |sb| - 1 ==> sb[i] == sb[i + 1] || Service_Next(sb[i], sb[i + 1])
    ensures forall i: int :: 0 <= i < |glb| ==> sb[i] == AbstractifyGLS_State(glb[i])
    ensures forall i: int :: 0 <= i < |sb| ==> sb[i].hosts == sb[0].hosts
    ensures sb[|sb| - 1] == AbstractifyGLS_State(glb[|glb| - 1])
    decreases config, glb
  {
    if |glb| == 1 {
      sb := [AbstractifyGLS_State(glb[0])];
      lemma_InitRefines(glb[0], config);
      assert Service_Init(AbstractifyGLS_State(glb[0]), MapSeqToSet(config, (x: EndPoint) => x));
    } else {
      ghost var rest := RefinementToServiceState(config, all_but_last(glb));
      ghost var gls := last(all_but_last(glb));
      ghost var gls' := last(glb);
      lemma_LS_NextAbstract(glb, config, |glb| - 2);
      sb := rest + [AbstractifyGLS_State(gls')];
      if AbstractifyGLS_State(gls) == AbstractifyGLS_State(gls') {
        assert sb[|sb| - 2] == sb[|sb| - 1];
      } else {
        assert Service_Next(sb[|sb| - 2], sb[|sb| - 1]);
      }
    }
  }

  lemma lemma_LockedPacketImpliesTransferPacket(config: ConcreteConfiguration, lb: seq<LS_State>, i: int, p: LockPacket)
    requires IsValidBehaviorLs(config, lb)
    requires 0 <= i < |lb|
    requires p in lb[i].environment.sentPackets
    requires p.src in lb[i].servers
    requires p.msg.Locked?
    ensures exists q: LPacket<EndPoint, LockMessage> :: q in lb[i].environment.sentPackets && q.msg.Transfer? && q.src in lb[i].servers && q.msg.transfer_epoch == p.msg.locked_epoch && q.dst == p.src
    decreases config, lb, i, p
  {
    if i == 0 {
      return;
    }
    lemma_DeduceTransitionFromLsBehavior(config, lb, i - 1);
    lemma_LsConsistency(config, lb, i);
    assert Collections__Maps2_s.mapdomain(lb[i].servers) == Collections__Maps2_s.mapdomain(lb[0].servers);
    assert LS_Init(lb[0], config);
    if p in lb[i - 1].environment.sentPackets {
      lemma_LockedPacketImpliesTransferPacket(config, lb, i - 1, p);
    } else {
      ghost var s := lb[i - 1];
      ghost var s' := lb[i];
      assert LS_Next(lb[i - 1], lb[i]);
      if s.environment.nextStep.LEnvStepHostIos? && s.environment.nextStep.actor in s.servers {
        assert LS_NextOneServer(s, s', s.environment.nextStep.actor, s.environment.nextStep.ios, s.environment.nextStep.nodeStep);
        ghost var id := s.environment.nextStep.actor;
        ghost var node := s.servers[id];
        ghost var node' := s'.servers[id];
        ghost var ios := s.environment.nextStep.ios;
        if NodeAccept(node, node', ios) {
          ghost var packet := ios[0].r;
          assert IsValidLIoOp(ios[0], id, s.environment);
          assert packet in lb[i].environment.sentPackets && packet.msg.Transfer? && packet.msg.transfer_epoch == p.msg.locked_epoch && packet.dst == p.src && packet.src in node.config;
          assert node.config == lb[0].servers[id].config == lb[i].servers[id].config;
          assert forall e: EndPoint :: e in lb[i].servers[id].config <==> e in Collections__Maps2_s.mapdomain(lb[i].servers);
          assert packet.src in lb[i].servers;
        }
      }
    }
  }

  lemma lemma_PacketSentByServerIsDemarshallable(config: ConcreteConfiguration, db: seq<DS_State>, i: int, p: LPacket<EndPoint, seq<byte>>)
    requires IsValidBehavior(config, db)
    requires 0 <= i < |db|
    requires p.src in config
    requires p in db[i].environment.sentPackets
    ensures Demarshallable(p.msg, CMessageGrammar())
    decreases config, db, i, p
  {
    if i == 0 {
      return;
    }
    if p in db[i - 1].environment.sentPackets {
      lemma_PacketSentByServerIsDemarshallable(config, db, i - 1, p);
      return;
    }
    lemma_DeduceTransitionFromDsBehavior(config, db, i - 1);
    lemma_DsConsistency(config, db, i - 1);
  }

  lemma RefinementProof(config: ConcreteConfiguration, db: seq<DS_State>) returns (sb: seq<ServiceState>)
    requires |db| > 0
    requires DS_Init(db[0], config)
    requires forall i: int {:trigger DS_Next(db[i], db[i + 1])} :: 0 <= i < |db| - 1 ==> DS_Next(db[i], db[i + 1])
    ensures |db| == |sb|
    ensures Service_Init(sb[0], Collections__Maps2_s.mapdomain(db[0].servers))
    ensures forall i: int {:trigger Service_Next(sb[i], sb[i + 1])} :: 0 <= i < |sb| - 1 ==> sb[i] == sb[i + 1] || Service_Next(sb[i], sb[i + 1])
    ensures forall i: int :: 0 <= i < |db| ==> Service_Correspondence(db[i].environment.sentPackets, sb[i])
    decreases config, db
  {
    ghost var lsb := RefinementToLSState(config, db);
    ghost var glsb := MakeGLSBehaviorFromLS(config, lsb);
    sb := RefinementToServiceState(config, glsb);
    forall i: int | 0 <= i < |db|
      ensures Service_Correspondence(db[i].environment.sentPackets, sb[i])
    {
      ghost var ls := lsb[i];
      ghost var gls := glsb[i];
      ghost var ss := sb[i];
      ghost var history := MakeLockHistory(glsb, config, i);
      assert history == gls.history;
      forall p: LPacket<EndPoint, seq<byte>>, epoch: int | p in db[i].environment.sentPackets && p.src in ss.hosts && p.dst in ss.hosts && p.msg == MarshallLockMsg(epoch)
        ensures 2 <= epoch <= |ss.history| && p.src == ss.history[epoch - 1]
      {
        ghost var ap := AbstractifyConcretePacket(p);
        assert p.src in sb[0].hosts;
        lemma_PacketSentByServerIsDemarshallable(config, db, i, p);
        assert Demarshallable(p.msg, CMessageGrammar());
        lemma_ParseMarshallLockedAbstract(p.msg, epoch, ap.msg);
        lemma_LockedPacketImpliesTransferPacket(config, lsb, i, ap);
        ghost var q :| q in ls.environment.sentPackets && q.msg.Transfer? && q.msg.transfer_epoch == ap.msg.locked_epoch && q.dst == p.src;
        assert q in gls.ls.environment.sentPackets;
      }
    }
  }

  method Main(ghost env: HostEnvironment) returns (exitCode: int)
    requires env != null && env.Valid() && env.ok.ok()
    requires env.udp.history() == []
    requires |env.constants.CommandLineArgs()| >= 2
    modifies set x: object | true
    decreases *
  {
    ghost var ok, host_state, config, servers, clients, id := HostInitImpl(env);
    assert ok ==> HostInit(host_state, config, id);
    while ok
      invariant ok ==> HostStateInvariants(host_state, env)
      invariant ok ==> env != null && env.Valid() && env.ok.ok()
      decreases *
    {
      ghost var old_udp_history := env.udp.history();
      ghost var old_state := host_state;
      ghost var recvs, clocks, sends, ios;
      ok, host_state, recvs, clocks, sends, ios := HostNextImpl(env, host_state);
      if ok {
        assert HostNext(old_state, host_state, ios);
        assert recvs + clocks + sends == ios;
        assert env.udp.history() == old_udp_history + recvs + clocks + sends;
        assert forall e: LIoOp<EndPoint, seq<byte>> :: (e in recvs ==> e.LIoOpReceive?) && (e in clocks ==> e.LIoOpReadClock? || e.LIoOpTimeoutReceive?) && (e in sends ==> e.LIoOpSend?);
        assert |clocks| <= 1;
      }
    }
  }

  import opened Collections__Seqs_s = Collections__Seqs_s
}

abstract module Main_s {

  import opened DS_s : DistributedSystem_s

  import opened AS_s : AbstractService_s

  import opened Collections__Seqs_s = Collections__Seqs_s
  method Main(ghost env: HostEnvironment) returns (exitCode: int)
    requires env != null && env.Valid() && env.ok.ok()
    requires env.udp.history() == []
    requires |env.constants.CommandLineArgs()| >= 2
    modifies set x: object | true
    decreases *
  {
    ghost var ok, host_state, config, servers, clients, id := HostInitImpl(env);
    assert ok ==> HostInit(host_state, config, id);
    while ok
      invariant ok ==> HostStateInvariants(host_state, env)
      invariant ok ==> env != null && env.Valid() && env.ok.ok()
      decreases *
    {
      ghost var old_udp_history := env.udp.history();
      ghost var old_state := host_state;
      ghost var recvs, clocks, sends, ios;
      ok, host_state, recvs, clocks, sends, ios := HostNextImpl(env, host_state);
      if ok {
        assert HostNext(old_state, host_state, ios);
        assert recvs + clocks + sends == ios;
        assert env.udp.history() == old_udp_history + recvs + clocks + sends;
        assert forall e: LIoOp<EndPoint, seq<byte>> :: (e in recvs ==> e.LIoOpReceive?) && (e in clocks ==> e.LIoOpReadClock? || e.LIoOpTimeoutReceive?) && (e in sends ==> e.LIoOpSend?);
        assert |clocks| <= 1;
      }
    }
  }

  lemma RefinementProof(config: ConcreteConfiguration, db: seq<DS_State>) returns (sb: seq<ServiceState>)
    requires |db| > 0
    requires DS_Init(db[0], config)
    requires forall i: int {:trigger DS_Next(db[i], db[i + 1])} :: 0 <= i < |db| - 1 ==> DS_Next(db[i], db[i + 1])
    ensures |db| == |sb|
    ensures Service_Init(sb[0], Collections__Maps2_s.mapdomain(db[0].servers))
    ensures forall i: int {:trigger Service_Next(sb[i], sb[i + 1])} :: 0 <= i < |sb| - 1 ==> sb[i] == sb[i + 1] || Service_Next(sb[i], sb[i + 1])
    ensures forall i: int :: 0 <= i < |db| ==> Service_Correspondence(db[i].environment.sentPackets, sb[i])
    decreases db
}

module Lock_DistributedSystem_i refines DistributedSystem_s {

  import opened H_s = Host_i
  predicate ValidPhysicalAddress(endPoint: EndPoint)
    decreases endPoint
  {
    |endPoint.addr| == 4 &&
    0 <= endPoint.port <= 65535
  }

  predicate ValidPhysicalPacket(p: LPacket<EndPoint, seq<byte>>)
    decreases p
  {
    ValidPhysicalAddress(p.src) &&
    ValidPhysicalAddress(p.dst) &&
    |p.msg| < 18446744073709551616
  }

  predicate ValidPhysicalIo(io: LIoOp<EndPoint, seq<byte>>)
    decreases io
  {
    (io.LIoOpReceive? ==>
      ValidPhysicalPacket(io.r)) &&
    (io.LIoOpSend? ==>
      ValidPhysicalPacket(io.s))
  }

  predicate ValidPhysicalEnvironmentStep(step: LEnvStep<EndPoint, seq<byte>, HostStep>)
    decreases step
  {
    step.LEnvStepHostIos? ==>
      forall io: LIoOp<EndPoint, seq<byte>> {:trigger io in step.ios} {:trigger ValidPhysicalIo(io)} :: 
        io in step.ios ==>
          ValidPhysicalIo(io)
  }

  predicate DS_Init(s: DS_State, config: ConcreteConfiguration)
    reads *
    decreases {}, s, config
  {
    s.config == config &&
    ConcreteConfigInit(s.config, mapdomain(s.servers), s.clients) &&
    LEnvironment_Init(s.environment) &&
    forall id: EndPoint :: 
      id in s.servers ==>
        HostInit(s.servers[id], config, id)
  }

  predicate DS_NextOneServer(s: DS_State, s': DS_State, id: EndPoint, ios: seq<LIoOp<EndPoint, seq<byte>>>)
    requires id in s.servers
    reads *
    decreases {}, s, s', id, ios
  {
    id in s'.servers &&
    HostNext(s.servers[id], s'.servers[id], ios) &&
    s'.servers == s.servers[id := s'.servers[id]]
  }

  predicate DS_Next(s: DS_State, s': DS_State)
    reads *
    decreases {}, s, s'
  {
    s'.config == s.config &&
    s'.clients == s.clients &&
    LEnvironment_Next(s.environment, s'.environment) &&
    ValidPhysicalEnvironmentStep(s.environment.nextStep) &&
    if s.environment.nextStep.LEnvStepHostIos? && s.environment.nextStep.actor in s.servers then DS_NextOneServer(s, s', s.environment.nextStep.actor, s.environment.nextStep.ios) else s'.servers == s.servers
  }

  import opened Collections__Maps2_s = Collections__Maps2_s

  import opened Native__Io_s = Native__Io_s

  import opened Environment_s = Environment_s

  import opened Native__NativeTypes_s = Native__NativeTypes_s

  datatype DS_State = DS_State(config: ConcreteConfiguration, environment: LEnvironment<EndPoint, seq<byte>, HostStep>, servers: map<EndPoint, HostState>, clients: set<EndPoint>)
}

module Environment_s {

  import opened Collections__Maps2_s = Collections__Maps2_s

  import opened Temporal__Temporal_s = Temporal__Temporal_s
  datatype LPacket<IdType, MessageType(==)> = LPacket(dst: IdType, src: IdType, msg: MessageType)

  datatype LIoOp<IdType, MessageType(==)> = LIoOpSend(s: LPacket<IdType, MessageType>) | LIoOpReceive(r: LPacket<IdType, MessageType>) | LIoOpTimeoutReceive | LIoOpReadClock(t: int)

  datatype LEnvStep<IdType, MessageType(==), NodeStepType> = LEnvStepHostIos(actor: IdType, ios: seq<LIoOp<IdType, MessageType>>, nodeStep: NodeStepType) | LEnvStepDeliverPacket(p: LPacket<IdType, MessageType>) | LEnvStepAdvanceTime | LEnvStepStutter

  datatype LHostInfo<IdType, MessageType(==)> = LHostInfo(queue: seq<LPacket<IdType, MessageType>>)

  datatype LEnvironment<IdType, MessageType(==), NodeStepType> = LEnvironment(time: int, sentPackets: set<LPacket<IdType, MessageType>>, hostInfo: map<IdType, LHostInfo<IdType, MessageType>>, nextStep: LEnvStep<IdType, MessageType, NodeStepType>)

  predicate IsValidLIoOp<IdType, MessageType, NodeStepType>(io: LIoOp<IdType, MessageType>, actor: IdType, e: LEnvironment<IdType, MessageType, NodeStepType>)
    decreases io, e
  {
    match io
    case LIoOpSend(s) =>
      s.src == actor
    case LIoOpReceive(r) =>
      r.dst == actor
    case LIoOpTimeoutReceive =>
      true
    case LIoOpReadClock(t) =>
      true
  }

  predicate LIoOpOrderingOKForAction<IdType, MessageType>(io1: LIoOp<IdType, MessageType>, io2: LIoOp<IdType, MessageType>)
    decreases io1, io2
  {
    io1.LIoOpReceive? || io2.LIoOpSend?
  }

  predicate LIoOpSeqCompatibleWithReduction<IdType, MessageType>(ios: seq<LIoOp<IdType, MessageType>>)
    decreases ios
  {
    forall i: int {:trigger ios[i], ios[i + 1]} :: 
      0 <= i < |ios| - 1 ==>
        LIoOpOrderingOKForAction(ios[i], ios[i + 1])
  }

  predicate IsValidLEnvStep<IdType, MessageType, NodeStepType>(e: LEnvironment<IdType, MessageType, NodeStepType>, step: LEnvStep<IdType, MessageType, NodeStepType>)
    decreases e, step
  {
    match step
    case LEnvStepHostIos(actor, ios, nodeStep) =>
      (forall io :: 
        io in ios ==>
          IsValidLIoOp(io, actor, e)) &&
      LIoOpSeqCompatibleWithReduction(ios)
    case LEnvStepDeliverPacket(p) =>
      p in e.sentPackets
    case LEnvStepAdvanceTime =>
      true
    case LEnvStepStutter =>
      true
  }

  predicate LEnvironment_Init<IdType, MessageType, NodeStepType>(e: LEnvironment<IdType, MessageType, NodeStepType>)
    decreases e
  {
    |e.sentPackets| == 0 &&
    e.time >= 0
  }

  predicate LEnvironment_PerformIos<IdType, MessageType, NodeStepType>(e: LEnvironment<IdType, MessageType, NodeStepType>, e': LEnvironment<IdType, MessageType, NodeStepType>, actor: IdType, ios: seq<LIoOp<IdType, MessageType>>)
    decreases e, e', ios
  {
    e'.sentPackets == e.sentPackets + (set io: LIoOp<IdType, MessageType> {:trigger io.s} {:trigger io.LIoOpSend?} {:trigger io in ios} | io in ios && io.LIoOpSend? :: io.s) &&
    (forall io: LIoOp<IdType, MessageType> :: 
      io in ios &&
      io.LIoOpReceive? ==>
        io.r in e.sentPackets) &&
    e'.time == e.time
  }

  predicate LEnvironment_AdvanceTime<IdType, MessageType, NodeStepType>(e: LEnvironment<IdType, MessageType, NodeStepType>, e': LEnvironment<IdType, MessageType, NodeStepType>)
    decreases e, e'
  {
    e'.time > e.time &&
    e'.sentPackets == e.sentPackets
  }

  predicate LEnvironment_Stutter<IdType, MessageType, NodeStepType>(e: LEnvironment<IdType, MessageType, NodeStepType>, e': LEnvironment<IdType, MessageType, NodeStepType>)
    decreases e, e'
  {
    e'.time == e.time &&
    e'.sentPackets == e.sentPackets
  }

  predicate LEnvironment_Next<IdType, MessageType, NodeStepType>(e: LEnvironment<IdType, MessageType, NodeStepType>, e': LEnvironment<IdType, MessageType, NodeStepType>)
    decreases e, e'
  {
    IsValidLEnvStep(e, e.nextStep) &&
    match e.nextStep case LEnvStepHostIos(actor, ios, nodeStep) => LEnvironment_PerformIos(e, e', actor, ios) case LEnvStepDeliverPacket(p) => LEnvironment_Stutter(e, e') case LEnvStepAdvanceTime => LEnvironment_AdvanceTime(e, e') case LEnvStepStutter => LEnvironment_Stutter(e, e')
  }

  function {:opaque} {:fuel 0, 0} EnvironmentNextTemporal<IdType, MessageType, NodeStepType>(b: Behavior<LEnvironment<IdType, MessageType, NodeStepType>>): temporal
    requires imaptotal(b)
    ensures forall i: int {:trigger sat(i, EnvironmentNextTemporal(b))} :: sat(i, EnvironmentNextTemporal(b)) <==> LEnvironment_Next(b[i], b[i + 1])
  {
    stepmap(imap i: int {:trigger b[i + 1]} | true :: LEnvironment_Next(b[i], b[i + 1]))
  }

  predicate LEnvironment_BehaviorSatisfiesSpec<IdType, MessageType, NodeStepType>(b: Behavior<LEnvironment<IdType, MessageType, NodeStepType>>)
  {
    imaptotal(b) &&
    LEnvironment_Init(b[0]) &&
    sat(0, always(EnvironmentNextTemporal(b)))
  }

  lemma {:axiom} {:auto_generated} {:opaque_reveal} {:verify false} {:fuel EnvironmentNextTemporal<int, int, int>, 1, 2} reveal_EnvironmentNextTemporal()
}

module Concrete_NodeIdentity_i refines Common__NodeIdentity_s {

  import opened Native__Io_s = Native__Io_s
  type NodeIdentity = EndPoint
}

module PacketParsing_i {

  import opened Common__GenericMarshalling_i = Common__GenericMarshalling_i

  import opened Message_i = Message_i

  import opened Common__UdpClient_i = Common__UdpClient_i
  predicate UdpPacketBound(data: seq<byte>)
    decreases data
  {
    |data| < MaxPacketSize()
  }

  function method CMessageTransferGrammar(): G
  {
    GUint64
  }

  function method CMessageLockedGrammar(): G
  {
    GUint64
  }

  function method CMessageGrammar(): G
  {
    GTaggedUnion([CMessageTransferGrammar(), CMessageLockedGrammar()])
  }

  function method ParseCMessageTransfer(val: V): CMessage
    requires ValInGrammar(val, CMessageTransferGrammar())
    decreases val
  {
    CTransfer(val.u)
  }

  function method ParseCMessageLocked(val: V): CMessage
    requires ValInGrammar(val, CMessageLockedGrammar())
    decreases val
  {
    CLocked(val.u)
  }

  function method ParseCMessage(val: V): CMessage
    requires ValInGrammar(val, CMessageGrammar())
    decreases val
  {
    if val.c == 0 then
      ParseCMessageTransfer(val.val)
    else
      ParseCMessageLocked(val.val)
  }

  function DemarshallData(data: seq<byte>): CMessage
    decreases data
  {
    if Demarshallable(data, CMessageGrammar()) then
      var val: V := DemarshallFunc(data, CMessageGrammar());
      ParseCMessage(val)
    else
      CInvalid()
  }

  method DemarshallDataMethod(data: array<byte>) returns (msg: CMessage)
    requires data != null
    requires data.Length < 18446744073709551616
    ensures msg == DemarshallData(data[..])
    decreases data
  {
    var success, val := Demarshall(data, CMessageGrammar());
    if success {
      msg := ParseCMessage(val);
      assert !msg.CInvalid?;
    } else {
      msg := CInvalid();
    }
  }

  method MarshallMessageTransfer(c: CMessage) returns (val: V)
    requires c.CTransfer?
    ensures ValInGrammar(val, CMessageTransferGrammar())
    ensures ValidVal(val)
    ensures ParseCMessageTransfer(val) == c
    ensures SizeOfV(val) < MaxPacketSize()
    decreases c
  {
    val := VUint64(c.transfer_epoch);
  }

  method MarshallMessageLocked(c: CMessage) returns (val: V)
    requires c.CLocked?
    ensures ValInGrammar(val, CMessageLockedGrammar())
    ensures ValidVal(val)
    ensures ParseCMessageLocked(val) == c
    ensures SizeOfV(val) < MaxPacketSize()
    decreases c
  {
    val := VUint64(c.locked_epoch);
  }

  method MarshallMessage(c: CMessage) returns (val: V)
    requires !c.CInvalid?
    ensures ValInGrammar(val, CMessageGrammar())
    ensures ValidVal(val)
    ensures ParseCMessage(val) == c
    ensures SizeOfV(val) < MaxPacketSize()
    decreases c
  {
    if c.CTransfer? {
      var msg := MarshallMessageTransfer(c);
      val := VCase(0, msg);
    } else if c.CLocked? {
      var msg := MarshallMessageLocked(c);
      val := VCase(1, msg);
    } else {
      assert false;
    }
  }

  method MarshallLockMessage(msg: CMessage) returns (data: array<byte>)
    requires !msg.CInvalid?
    ensures data != null
    ensures fresh(data)
    ensures UdpPacketBound(data[..])
    ensures DemarshallData(data[..]) == msg
    decreases msg
  {
    var val := MarshallMessage(msg);
    data := Marshall(val, CMessageGrammar());
  }

  function AbstractifyUdpPacket(udp: UdpPacket): LockPacket
    decreases udp
  {
    LPacket(udp.dst, udp.src, AbstractifyCMessage(DemarshallData(udp.msg)))
  }

  predicate CLockPacketValid(p: CLockPacket)
    decreases p
  {
    EndPointIsValidIPV4(p.src) &&
    EndPointIsValidIPV4(p.dst) &&
    !p.msg.CInvalid?
  }

  predicate OptionCLockPacketValid(opt_packet: Option<CLockPacket>)
    decreases opt_packet
  {
    opt_packet.Some? ==>
      CLockPacketValid(opt_packet.v)
  }
}

module UdpLock_i {

  import opened Common__UdpClient_i = Common__UdpClient_i

  import opened PacketParsing_i = PacketParsing_i
  datatype ReceiveResult = RRFail | RRTimeout | RRPacket(cpacket: CLockPacket)

  function AstractifyUdpEventToLockIo(evt: UdpEvent): LockIo
    decreases evt
  {
    match evt
    case LIoOpSend(s) =>
      LIoOpSend(AbstractifyUdpPacket(s))
    case LIoOpReceive(r) =>
      LIoOpReceive(AbstractifyUdpPacket(r))
    case LIoOpTimeoutReceive =>
      LIoOpTimeoutReceive()
    case LIoOpReadClock(t) =>
      LIoOpReadClock(t as int)
  }

  function {:opaque} {:fuel 0, 0} AbstractifyRawLogToIos(rawlog: seq<UdpEvent>): seq<LockIo>
    ensures |AbstractifyRawLogToIos(rawlog)| == |rawlog|
    ensures forall i: int {:trigger AstractifyUdpEventToLockIo(rawlog[i])} {:trigger AbstractifyRawLogToIos(rawlog)[i]} :: 0 <= i < |rawlog| ==> AbstractifyRawLogToIos(rawlog)[i] == AstractifyUdpEventToLockIo(rawlog[i])
    decreases rawlog
  {
    if rawlog == [] then
      []
    else
      [AstractifyUdpEventToLockIo(rawlog[0])] + AbstractifyRawLogToIos(rawlog[1..])
  }

  lemma /*{:_induction rawlog}*/ lemma_EstablishAbstractifyRawLogToIos(rawlog: seq<UdpEvent>, ios: seq<LockIo>)
    requires |rawlog| == |ios|
    requires forall i: int :: 0 <= i < |rawlog| ==> ios[i] == AstractifyUdpEventToLockIo(rawlog[i])
    ensures AbstractifyRawLogToIos(rawlog) == ios
    decreases rawlog, ios
  {
  }

  predicate OnlySentMarshallableData(rawlog: seq<UdpEvent>)
    decreases rawlog
  {
    forall io: LIoOp<EndPoint, seq<byte>> :: 
      io in rawlog &&
      io.LIoOpSend? ==>
        UdpPacketBound(io.s.msg) &&
        Demarshallable(io.s.msg, CMessageGrammar())
  }

  method GetEndPoint(ipe: IPEndPoint) returns (ep: EndPoint)
    requires ipe != null
    ensures ep == ipe.EP()
    ensures EndPointIsValidIPV4(ep)
    decreases ipe
  {
    var addr := ipe.GetAddress();
    var port := ipe.GetPort();
    ep := EndPoint(addr[..], port);
  }

  method Receive(udpClient: UdpClient, localAddr: EndPoint)
      returns (rr: ReceiveResult, ghost udpEvent: UdpEvent)
    requires UdpClientIsValid(udpClient)
    requires udpClient.LocalEndPoint() == localAddr
    modifies UdpClientRepr(udpClient)
    ensures udpClient.env == old(udpClient.env)
    ensures udpClient.LocalEndPoint() == old(udpClient.LocalEndPoint())
    ensures UdpClientOk(udpClient) <==> !rr.RRFail?
    ensures old(UdpClientRepr(udpClient)) == UdpClientRepr(udpClient)
    ensures !rr.RRFail? ==> udpClient.IsOpen() && old(udpClient.env.udp.history()) + [udpEvent] == udpClient.env.udp.history()
    ensures rr.RRTimeout? ==> udpEvent.LIoOpTimeoutReceive?
    ensures rr.RRPacket? ==> udpEvent.LIoOpReceive? && EndPointIsValidIPV4(rr.cpacket.src) && AbstractifyCLockPacket(rr.cpacket) == AbstractifyUdpPacket(udpEvent.r) && rr.cpacket.msg == DemarshallData(udpEvent.r.msg)
    decreases udpClient, localAddr
  {
    var timeout := 0;
    ghost var old_udp_history := udpClient.env.udp.history();
    var ok, timedOut, remote, buffer := udpClient.Receive(timeout);
    if !ok {
      rr := RRFail();
      return;
    }
    if timedOut {
      rr := RRTimeout();
      udpEvent := LIoOpTimeoutReceive();
      return;
    }
    udpEvent := LIoOpReceive(LPacket(udpClient.LocalEndPoint(), remote.EP(), buffer[..]));
    var cmessage := DemarshallDataMethod(buffer);
    var srcEp := GetEndPoint(remote);
    var cpacket := LPacket(localAddr, srcEp, cmessage);
    rr := RRPacket(cpacket);
  }

  predicate SendLogEntryReflectsPacket(event: UdpEvent, cpacket: CLockPacket)
    decreases event, cpacket
  {
    event.LIoOpSend? &&
    AbstractifyCLockPacket(cpacket) == AbstractifyUdpPacket(event.s)
  }

  predicate SendLogReflectsPacket(udpEventLog: seq<UdpEvent>, packet: Option<CLockPacket>)
    decreases udpEventLog, packet
  {
    match packet {
      case Some(p) =>
        |udpEventLog| == 1 &&
        SendLogEntryReflectsPacket(udpEventLog[0], p)
      case None =>
        udpEventLog == []
    }
  }

  method SendPacket(udpClient: UdpClient, opt_packet: Option<CLockPacket>, ghost localAddr: EndPoint)
      returns (ok: bool, ghost udpEventLog: seq<UdpEvent>)
    requires UdpClientIsValid(udpClient)
    requires udpClient.LocalEndPoint() == localAddr
    requires OptionCLockPacketValid(opt_packet)
    requires opt_packet.Some? ==> opt_packet.v.src == localAddr
    modifies UdpClientRepr(udpClient)
    ensures old(UdpClientRepr(udpClient)) == UdpClientRepr(udpClient)
    ensures udpClient.env == old(udpClient.env)
    ensures udpClient.LocalEndPoint() == old(udpClient.LocalEndPoint())
    ensures UdpClientOk(udpClient) <==> ok
    ensures ok ==> UdpClientIsValid(udpClient) && udpClient.IsOpen() && old(udpClient.env.udp.history()) + udpEventLog == udpClient.env.udp.history() && OnlySentMarshallableData(udpEventLog) && SendLogReflectsPacket(udpEventLog, opt_packet)
    decreases udpClient, opt_packet, localAddr
  {
    udpEventLog := [];
    ok := true;
    if opt_packet.None? {
    } else {
      var cpacket := opt_packet.v;
      var dstEp: EndPoint := cpacket.dst;
      var dstAddrAry := seqToArrayOpt(dstEp.addr);
      var remote;
      ok, remote := IPEndPoint.Construct(dstAddrAry, dstEp.port, udpClient.env);
      if !ok {
        return;
      }
      var buffer := MarshallLockMessage(cpacket.msg);
      ok := udpClient.Send(remote, buffer);
      if !ok {
        return;
      }
      ghost var udpEvent := LIoOpSend(LPacket(remote.EP(), udpClient.LocalEndPoint(), buffer[..]));
      udpEventLog := [udpEvent];
    }
  }

  lemma {:axiom} {:auto_generated} {:opaque_reveal} {:verify false} {:fuel AbstractifyRawLogToIos, 1, 2} reveal_AbstractifyRawLogToIos()
}

module Host_i refines Host_s {

  import opened Collections__Sets_i = Collections__Sets_i

  import opened NodeImpl_i = NodeImpl_i

  import opened LockCmdLineParser_i = LockCmdLineParser_i
  datatype CScheduler = CScheduler(ghost node: Node, node_impl: NodeImpl)

  type HostState = CScheduler

  type HostStep = LockStep

  type ConcreteConfiguration = Config

  predicate ConcreteConfigurationInvariants(config: ConcreteConfiguration)
    decreases config
  {
    ValidConfig(config)
  }

  predicate HostStateInvariants(host_state: HostState, env: HostEnvironment)
    reads *
    decreases {}, host_state, env
  {
    host_state.node_impl != null &&
    host_state.node_impl.Valid() &&
    host_state.node_impl.Env() == env &&
    host_state.node == AbstractifyCNode(host_state.node_impl.node)
  }

  predicate HostInit(host_state: HostState, config: ConcreteConfiguration, id: EndPoint)
    reads *
    decreases {}, host_state, config, id
  {
    host_state.node_impl != null &&
    host_state.node_impl.Valid() &&
    host_state.node_impl.node.config == config &&
    host_state.node_impl.node.config[host_state.node_impl.node.my_index] == id &&
    NodeInit(host_state.node, host_state.node_impl.node.my_index as int, config)
  }

  predicate HostNext(host_state: HostState, host_state': HostState, ios: seq<LIoOp<EndPoint, seq<byte>>>)
    reads *
    decreases {}, host_state, host_state', ios
  {
    NodeNext(host_state.node, host_state'.node, AbstractifyRawLogToIos(ios)) &&
    OnlySentMarshallableData(ios)
  }

  predicate ConcreteConfigInit(config: ConcreteConfiguration, servers: set<EndPoint>, clients: set<EndPoint>)
    decreases config, servers, clients
  {
    ValidConfig(config) &&
    MapSeqToSet(config, (x: EndPoint) => x) == servers
  }

  function ParseCommandLineConfiguration(args: seq<seq<uint16>>): (ConcreteConfiguration, set<EndPoint>, set<EndPoint>)
    decreases args
  {
    var lock_config: seq<EndPoint> := lock_config_parsing(args);
    var endpoints_set: set<EndPoint> := set e: EndPoint {:trigger e in lock_config} | e in lock_config;
    (lock_config, endpoints_set, {})
  }

  function ParseCommandLineId(ip: seq<uint16>, port: seq<uint16>): EndPoint
    decreases ip, port
  {
    lock_parse_id(ip, port)
  }

  method HostInitImpl(ghost env: HostEnvironment)
      returns (ok: bool, host_state: HostState, config: ConcreteConfiguration, ghost servers: set<EndPoint>, ghost clients: set<EndPoint>, id: EndPoint)
    requires env != null && env.Valid()
    requires env.ok.ok()
    requires |env.constants.CommandLineArgs()| >= 2
    modifies set x: object | true
    ensures ok ==> env != null && env.Valid() && env.ok.ok()
    ensures ok ==> |env.constants.CommandLineArgs()| >= 2
    ensures ok ==> HostStateInvariants(host_state, env)
    ensures ok ==> ConcreteConfigurationInvariants(config)
    ensures ok ==> var args: seq<seq<uint16>> := env.constants.CommandLineArgs(); var (parsed_config: ConcreteConfiguration, parsed_servers: set<EndPoint>, parsed_clients: set<EndPoint>) := ParseCommandLineConfiguration(args[0 .. |args| - 2]); config == parsed_config && servers == parsed_servers && clients == parsed_clients && ConcreteConfigInit(parsed_config, parsed_servers, parsed_clients)
    ensures ok ==> var args: seq<seq<uint16>> := env.constants.CommandLineArgs(); id == ParseCommandLineId(args[|args| - 2], args[|args| - 1]) && HostInit(host_state, config, id)
    decreases env
  {
    var my_index;
    ok, config, my_index := ParseCmdLine(env);
    if !ok {
      return;
    }
    id := config[my_index];
    var node_impl := new NodeImpl();
    ok := node_impl.InitNode(config, my_index, env);
    if !ok {
      return;
    }
    host_state := CScheduler(AbstractifyCNode(node_impl.node), node_impl);
    servers := set e: EndPoint {:trigger e in config} | e in config;
    clients := {};
  }

  predicate EventsConsistent(recvs: seq<UdpEvent>, clocks: seq<UdpEvent>, sends: seq<UdpEvent>)
    decreases recvs, clocks, sends
  {
    forall e: LIoOp<EndPoint, seq<byte>> :: 
      (e in recvs ==>
        e.LIoOpReceive?) &&
      (e in clocks ==>
        e.LIoOpReadClock? || e.LIoOpTimeoutReceive?) &&
      (e in sends ==>
        e.LIoOpSend?)
  }

  ghost method RemoveRecvs(events: seq<UdpEvent>) returns (recvs: seq<UdpEvent>, rest: seq<UdpEvent>)
    ensures forall e: LIoOp<EndPoint, seq<byte>> :: e in recvs ==> e.LIoOpReceive?
    ensures events == recvs + rest
    ensures rest != [] ==> !rest[0].LIoOpReceive?
    decreases events
  {
    recvs := [];
    rest := [];
    ghost var i := 0;
    while i < |events|
      invariant 0 <= i <= |events|
      invariant forall e: LIoOp<EndPoint, seq<byte>> :: e in recvs ==> e.LIoOpReceive?
      invariant recvs == events[0 .. i]
      decreases |events| - i
    {
      if !events[i].LIoOpReceive? {
        rest := events[i..];
        return;
      }
      recvs := recvs + [events[i]];
      i := i + 1;
    }
  }

  predicate UdpEventsReductionCompatible(events: seq<UdpEvent>)
    decreases events
  {
    forall i: int :: 
      0 <= i < |events| - 1 ==>
        events[i].LIoOpReceive? || events[i + 1].LIoOpSend?
  }

  lemma RemainingEventsAreSends(events: seq<UdpEvent>)
    requires UdpEventsReductionCompatible(events)
    requires |events| > 0
    requires !events[0].LIoOpReceive?
    ensures forall e: LIoOp<EndPoint, seq<byte>> :: e in events[1..] ==> e.LIoOpSend?
    decreases events
  {
  }

  ghost method PartitionEvents(events: seq<UdpEvent>)
      returns (recvs: seq<UdpEvent>, clocks: seq<UdpEvent>, sends: seq<UdpEvent>)
    requires UdpEventsReductionCompatible(events)
    ensures events == recvs + clocks + sends
    ensures EventsConsistent(recvs, clocks, sends)
    ensures |clocks| <= 1
    decreases events
  {
    ghost var rest;
    recvs, rest := RemoveRecvs(events);
    if |rest| > 0 && (rest[0].LIoOpReadClock? || rest[0].LIoOpTimeoutReceive?) {
      clocks := [rest[0]];
      sends := rest[1..];
      RemainingEventsAreSends(rest);
    } else {
      clocks := [];
      sends := rest;
      if |rest| > 0 {
        RemainingEventsAreSends(rest);
      }
    }
  }

  lemma /*{:_induction events}*/ UdpEventsRespectReduction(s: Node, s': Node, ios: seq<LockIo>, events: seq<UdpEvent>)
    requires LIoOpSeqCompatibleWithReduction(ios)
    requires AbstractifyRawLogToIos(events) == ios
    ensures UdpEventsReductionCompatible(events)
    decreases s, s', ios, events
  {
  }

  method HostNextImpl(ghost env: HostEnvironment, host_state: HostState)
      returns (ok: bool, host_state': HostState, ghost recvs: seq<UdpEvent>, ghost clocks: seq<UdpEvent>, ghost sends: seq<UdpEvent>, ghost ios: seq<LIoOp<EndPoint, seq<byte>>>)
    requires env != null && env.Valid() && env.ok.ok()
    requires HostStateInvariants(host_state, env)
    modifies set x: object | true
    ensures ok <==> env != null && env.Valid() && env.ok.ok()
    ensures ok ==> HostStateInvariants(host_state', env)
    ensures ok ==> HostNext(host_state, host_state', ios)
    ensures ok ==> recvs + clocks + sends == ios
    ensures ok ==> env.udp.history() == old(env.udp.history()) + (recvs + clocks + sends)
    ensures forall e: LIoOp<EndPoint, seq<byte>> :: (e in recvs ==> e.LIoOpReceive?) && (e in clocks ==> e.LIoOpReadClock? || e.LIoOpTimeoutReceive?) && (e in sends ==> e.LIoOpSend?)
    ensures |clocks| <= 1
    decreases env, host_state
  {
    ghost var okay, udpEventLog, abstract_ios := host_state.node_impl.HostNextMain();
    if okay {
      UdpEventsRespectReduction(host_state.node, AbstractifyCNode(host_state.node_impl.node), abstract_ios, udpEventLog);
      recvs, clocks, sends := PartitionEvents(udpEventLog);
      ios := recvs + clocks + sends;
      assert ios == udpEventLog;
      host_state' := CScheduler(AbstractifyCNode(host_state.node_impl.node), host_state.node_impl);
    } else {
      recvs := [];
      clocks := [];
      sends := [];
    }
    ok := okay;
  }

  import opened Native__Io_s = Native__Io_s

  import opened Environment_s = Environment_s

  import opened Native__NativeTypes_s = Native__NativeTypes_s
}

module AbstractServiceLock_s refines AbstractService_s {
  datatype ServiceState' = ServiceState'(hosts: set<EndPoint>, history: seq<EndPoint>)

  type ServiceState = ServiceState'

  predicate Service_Init(s: ServiceState, serverAddresses: set<EndPoint>)
    decreases s, serverAddresses
  {
    s.hosts == serverAddresses &&
    exists e: EndPoint :: 
      e in serverAddresses &&
      s.history == [e]
  }

  predicate Service_Next(s: ServiceState, s': ServiceState)
    decreases s, s'
  {
    s'.hosts == s.hosts &&
    exists new_lock_holder: EndPoint :: 
      new_lock_holder in s.hosts &&
      s'.history == s.history + [new_lock_holder]
  }

  function Uint64ToBytes(u: uint64): seq<byte>
    decreases u
  {
    [byte(u / 72057594037927936), byte(u / 281474976710656 % 256), byte(u / 1099511627776 % 256), byte(u / 4294967296 % 256), byte(u / 16777216 % 256), byte(u / 65536 % 256), byte(u / 256 % 256), byte(u % 256)]
  }

  function MarshallLockMsg(epoch: int): seq<byte>
    decreases epoch
  {
    if 0 <= epoch < 18446744073709551616 then
      [0, 0, 0, 0, 0, 0, 0, 1] + Uint64ToBytes(uint64(epoch))
    else
      [1]
  }

  predicate Service_Correspondence(concretePkts: set<LPacket<EndPoint, seq<byte>>>, serviceState: ServiceState)
    decreases concretePkts, serviceState
  {
    forall p: LPacket<EndPoint, seq<byte>>, epoch: int :: 
      p in concretePkts &&
      p.src in serviceState.hosts &&
      p.dst in serviceState.hosts &&
      p.msg == MarshallLockMsg(epoch) ==>
        1 <= epoch <= |serviceState.history| &&
        p.src == serviceState.history[epoch - 1]
  }

  import opened Native__Io_s = Native__Io_s

  import opened Environment_s = Environment_s

  import opened Native__NativeTypes_s = Native__NativeTypes_s
}

module Refinement_i {

  import opened DistributedSystem_i = DistributedSystem_i

  import opened AbstractServiceLock_s = AbstractServiceLock_s
  function AbstractifyGLS_State(gls: GLS_State): ServiceState
    decreases gls
  {
    ServiceState'(mapdomain(gls.ls.servers), gls.history)
  }
}

module RefinementProof_i {

  import opened Refinement_i = Refinement_i

  import opened Collections__Sets_i = Collections__Sets_i

  import opened Collections__Maps_i = Collections__Maps_i

  import opened Logic__Option_i = Logic__Option_i
  lemma lemma_InitRefines(gls: GLS_State, config: Config)
    requires GLS_Init(gls, config)
    ensures Service_Init(AbstractifyGLS_State(gls), UniqueSeqToSet(config))
    decreases gls, config
  {
  }

  predicate IsValidBehavior(glb: seq<GLS_State>, config: Config)
    decreases glb, config
  {
    |glb| > 0 &&
    GLS_Init(glb[0], config) &&
    forall i: int {:trigger GLS_Next(glb[i], glb[i + 1])} :: 
      0 <= i < |glb| - 1 ==>
        GLS_Next(glb[i], glb[i + 1])
  }

  lemma lemma_LS_Next(glb: seq<GLS_State>, config: Config, i: int)
    requires IsValidBehavior(glb, config)
    requires 0 <= i < |glb| - 1
    ensures GLS_Next(glb[i], glb[i + 1])
    decreases glb, config, i
  {
  }

  lemma lemma_LSConsistent(glb: seq<GLS_State>, config: Config, i: int)
    requires IsValidBehavior(glb, config)
    requires 0 <= i < |glb|
    ensures |glb[i].ls.servers| == |config|
    ensures forall e: EndPoint :: e in config <==> e in glb[i].ls.servers
    ensures mapdomain(glb[i].ls.servers) == mapdomain(glb[0].ls.servers)
    ensures forall id: EndPoint :: id in config ==> glb[0].ls.servers[id].config == glb[i].ls.servers[id].config
    decreases glb, config, i
  {
  }

  lemma lemma_LSNodeConsistent(glb: seq<GLS_State>, config: Config, i: int, candidate: EndPoint, e: EndPoint)
    requires IsValidBehavior(glb, config)
    requires 0 <= i < |glb|
    requires e in glb[i].ls.servers
    ensures candidate in glb[i].ls.servers <==> candidate in glb[i].ls.servers[e].config
    decreases glb, config, i, candidate, e
  {
  }

  lemma lemma_HistoryIncrement(glb: seq<GLS_State>, config: Config, i: int)
    requires IsValidBehavior(glb, config)
    requires 0 <= i < |glb| - 1
    ensures |glb[i].history| + 1 == |glb[i].history| || glb[i].history == glb[i].history
    decreases glb, config, i
  {
  }

  lemma lemma_HistorySize(glb: seq<GLS_State>, config: Config, i: int)
    requires IsValidBehavior(glb, config)
    requires 0 <= i < |glb|
    ensures 1 <= |glb[i].history| <= i + 1
    decreases glb, config, i
  {
  }

  lemma lemma_HistoryMembership(glb: seq<GLS_State>, config: Config, i: int)
    requires IsValidBehavior(glb, config)
    requires 0 <= i < |glb|
    ensures 1 <= |glb[i].history| <= i + 1
    ensures last(glb[i].history) in glb[i].ls.servers
    decreases glb, config, i
  {
  }

  lemma lemma_LS_NextAbstract(glb: seq<GLS_State>, config: Config, i: int)
    requires IsValidBehavior(glb, config)
    requires 0 <= i < |glb| - 1
    ensures Service_Next(AbstractifyGLS_State(glb[i]), AbstractifyGLS_State(glb[i + 1])) || AbstractifyGLS_State(glb[i]) == AbstractifyGLS_State(glb[i + 1])
    decreases glb, config, i
  {
  }

  lemma MakeLockHistory(glb: seq<GLS_State>, config: Config, i: int)
      returns (history: seq<EndPoint>)
    requires IsValidBehavior(glb, config)
    requires 0 <= i < |glb|
    ensures |history| > 0
    ensures forall p: LPacket<EndPoint, LockMessage> :: p in glb[i].ls.environment.sentPackets && p.msg.Transfer? && p.src in glb[i].ls.servers ==> 2 <= p.msg.transfer_epoch <= |history|
    ensures forall p: LPacket<EndPoint, LockMessage> :: p in glb[i].ls.environment.sentPackets && p.msg.Transfer? && p.src in glb[i].ls.servers ==> history[p.msg.transfer_epoch - 1] == p.dst
    ensures forall h: EndPoint, j: int :: h in glb[i].ls.servers && 0 <= j < |history| - 1 && history[j] == h ==> j + 1 <= glb[i].ls.servers[h].epoch
    ensures forall h: EndPoint :: h in glb[i].ls.servers && h != last(history) ==> !glb[i].ls.servers[h].held
    ensures forall h: EndPoint :: h in glb[i].ls.servers && glb[i].ls.servers[h].held ==> glb[i].ls.servers[h].epoch == |history|
    ensures history == glb[i].history
    decreases glb, config, i
  {
  }
}

module MarshallProof_i {

  import opened AbstractServiceLock_s = AbstractServiceLock_s

  import opened Types_i = Types_i

  import opened PacketParsing_i = PacketParsing_i
  lemma lemma_ParseValCorrectVCase(data: seq<byte>, v: V, g: G)
      returns (caseId: uint64, val: V, rest: seq<byte>)
    requires ValInGrammar(v, g)
    requires |data| < 18446744073709551616
    requires ValidGrammar(g)
    requires parse_Val(data, g).0.Some?
    requires parse_Val(data, g).0.v == v
    requires g.GTaggedUnion?
    ensures parse_Uint64(data).0.Some?
    ensures caseId == parse_Uint64(data).0.v.u
    ensures 0 <= caseId as int < |g.cases|
    ensures rest == parse_Uint64(data).1
    ensures parse_Val(rest, g.cases[caseId]).0.Some?
    ensures val == parse_Val(rest, g.cases[caseId]).0.v
    ensures v == VCase(caseId, val)
    ensures ValInGrammar(val, g.cases[caseId])
    decreases data, v, g
  {
  }

  lemma lemma_ParseValCorrectVUint64(data: seq<byte>, v: V, g: G)
      returns (u: uint64, rest: seq<byte>)
    requires ValInGrammar(v, g)
    requires |data| < 18446744073709551616
    requires ValidGrammar(g)
    requires parse_Val(data, g).0.Some?
    requires parse_Val(data, g).0.v == v
    requires g.GUint64?
    ensures parse_Uint64(data).0.Some?
    ensures u == parse_Uint64(data).0.v.u
    ensures v == VUint64(u)
    ensures rest == parse_Val(data, g).1
    decreases data, v, g
  {
  }

  lemma /*{:_induction v}*/ lemma_SizeOfCMessageLocked(v: V)
    requires ValInGrammar(v, CMessageGrammar())
    requires ValInGrammar(v.val, CMessageLockedGrammar())
    ensures SizeOfV(v) == 16
    decreases v
  {
  }

  lemma lemma_ParseMarshallLockedAbstract(bytes: seq<byte>, epoch: int, msg: LockMessage)
    requires AbstractifyCMessage(DemarshallData(bytes)) == msg
    requires bytes == MarshallLockMsg(epoch)
    requires Demarshallable(bytes, CMessageGrammar())
    ensures msg.Locked?
    ensures msg.locked_epoch == epoch
    decreases bytes, epoch, msg
  {
  }
}

module Native__Io_s {

  import opened Native__NativeTypes_s = Native__NativeTypes_s

  import opened Environment_s = Environment_s
  class HostEnvironment {
    ghost var constants: HostConstants
    ghost var ok: OkState
    ghost var now: NowState
    ghost var udp: UdpState
    ghost var files: FileSystemState

    predicate Valid()
      reads this
      decreases {this}
    {
      constants != null &&
      ok != null &&
      now != null &&
      udp != null &&
      files != null
    }
  }

  class HostConstants {
    constructor {:axiom} ()
      requires false

    function {:axiom} LocalAddress(): seq<byte>
      reads this
      decreases {this}

    function {:axiom} CommandLineArgs(): seq<seq<uint16>>
      reads this
      decreases {this}

    static method {:axiom} NumCommandLineArgs(ghost env: HostEnvironment) returns (n: uint32)
      requires env != null && env.Valid()
      ensures n as int == |env.constants.CommandLineArgs()|
      decreases env

    static method {:axiom} GetCommandLineArg(i: uint64, ghost env: HostEnvironment) returns (arg: array<uint16>)
      requires env != null && env.Valid()
      requires 0 <= i as int < |env.constants.CommandLineArgs()|
      ensures arg != null
      ensures fresh(arg)
      ensures arg[..] == env.constants.CommandLineArgs()[i]
      decreases i, env
  }

  class OkState {
    constructor {:axiom} ()
      requires false

    function {:axiom} ok(): bool
      reads this
      decreases {this}
  }

  class NowState {
    constructor {:axiom} ()
      requires false

    function {:axiom} now(): int
      reads this
      decreases {this}
  }

  class Time {
    static method {:axiom} GetTime(ghost env: HostEnvironment) returns (t: uint64)
      requires env != null && env.Valid()
      modifies env.now, env.udp
      ensures t as int == env.now.now()
      ensures AdvanceTime(old(env.now.now()), env.now.now(), 0)
      ensures env.udp.history() == old(env.udp.history()) + [LIoOpReadClock(t as int)]
      decreases env

    static method {:axiom} GetDebugTimeTicks() returns (t: uint64)

    static method {:axiom} RecordTiming(name: array<char>, time: uint64)
      decreases name, time
  }

  datatype EndPoint = EndPoint(addr: seq<byte>, port: uint16)

  type UdpPacket = LPacket<EndPoint, seq<byte>>

  type UdpEvent = LIoOp<EndPoint, seq<byte>>

  class UdpState {
    constructor {:axiom} ()
      requires false

    function {:axiom} history(): seq<UdpEvent>
      reads this
      decreases {this}
  }

  class IPEndPoint {
    ghost var env: HostEnvironment

    function {:axiom} Address(): seq<byte>
      reads this
      decreases {this}

    function {:axiom} Port(): uint16
      reads this
      decreases {this}

    function EP(): EndPoint
      reads this
      decreases {this}
    {
      EndPoint(Address(), Port())
    }

    constructor {:axiom} ()
      requires false

    method {:axiom} GetAddress() returns (addr: array<byte>)
      ensures addr != null
      ensures fresh(addr)
      ensures addr[..] == Address()
      ensures addr.Length == 4

    function method {:axiom} GetPort(): uint16
      reads this
      ensures GetPort() == Port()
      decreases {this}

    static method {:axiom} Construct(ipAddress: array<byte>, port: uint16, ghost env: HostEnvironment)
        returns (ok: bool, ep: IPEndPoint)
      requires env != null && env.Valid()
      requires ipAddress != null
      modifies env.ok
      ensures env.ok.ok() == ok
      ensures ok ==> ep != null && fresh(ep) && ep.env == env && ep.Address() == ipAddress[..] && ep.Port() == port
      decreases ipAddress, port, env
  }

  class UdpClient {
    ghost var env: HostEnvironment

    function {:axiom} LocalEndPoint(): EndPoint
      reads this
      decreases {this}

    function {:axiom} IsOpen(): bool
      reads this
      decreases {this}

    constructor {:axiom} ()
    {
    }

    static method {:axiom} Construct(localEP: IPEndPoint, ghost env: HostEnvironment)
        returns (ok: bool, udp: UdpClient)
      requires env != null && env.Valid()
      requires env.ok.ok()
      requires localEP != null
      modifies env.ok
      ensures env.ok.ok() == ok
      ensures ok ==> udp != null && fresh(udp) && udp.env == env && udp.IsOpen() && udp.LocalEndPoint() == localEP.EP()
      decreases localEP, env

    method {:axiom} Close() returns (ok: bool)
      requires env != null && env.Valid()
      requires env.ok.ok()
      requires this.IsOpen()
      modifies this, env.ok
      ensures env == old(env)
      ensures env.ok.ok() == ok

    method {:axiom} Receive(timeLimit: int32)
        returns (ok: bool, timedOut: bool, remote: IPEndPoint, buffer: array<byte>)
      requires env != null && env.Valid()
      requires env.ok.ok()
      requires IsOpen()
      requires timeLimit >= 0
      requires timeLimit as int * 1000 < 2147483648
      modifies this, env.ok, env.now, env.udp
      ensures env == old(env)
      ensures env.ok.ok() == ok
      ensures AdvanceTime(old(env.now.now()), env.now.now(), timeLimit as int)
      ensures LocalEndPoint() == old(LocalEndPoint())
      ensures ok ==> IsOpen()
      ensures ok ==> timedOut ==> env.udp.history() == old(env.udp.history()) + [LIoOpTimeoutReceive()]
      ensures ok ==> !timedOut ==> remote != null && buffer != null && fresh(remote) && fresh(buffer) && env.udp.history() == old(env.udp.history()) + [LIoOpReceive(LPacket(LocalEndPoint(), remote.EP(), buffer[..]))] && buffer.Length < 18446744073709551616
      decreases timeLimit

    method {:axiom} Send(remote: IPEndPoint, buffer: array<byte>) returns (ok: bool)
      requires env != null && env.Valid()
      requires env.ok.ok()
      requires IsOpen()
      requires remote != null
      requires buffer != null
      requires buffer.Length <= MaxPacketSize()
      modifies this, env.ok, env.udp
      ensures env == old(env)
      ensures env.ok.ok() == ok
      ensures LocalEndPoint() == old(LocalEndPoint())
      ensures ok ==> IsOpen()
      ensures ok ==> env.udp.history() == old(env.udp.history()) + [LIoOpSend(LPacket(remote.EP(), LocalEndPoint(), buffer[..]))]
      decreases remote, buffer
  }

  class FileSystemState { }

  class MutableSet<T(==,0,!new)> {
    static function method {:axiom} SetOf(s: MutableSet<T>): set<T>
      reads s
      decreases {s}, s

    static method {:axiom} EmptySet() returns (s: MutableSet<T>)
      ensures SetOf(s) == {}
      ensures fresh(s)

    constructor {:axiom} ()
      requires false

    method {:axiom} Size() returns (size: int)
      ensures size == |SetOf(this)|

    method {:axiom} SizeModest() returns (size: uint64)
      requires |SetOf(this)| < 18446744073709551616
      ensures size as int == |SetOf(this)|

    method {:axiom} Contains(x: T) returns (contains: bool)
      ensures contains == (x in SetOf(this))

    method {:axiom} Add(x: T)
      modifies this
      ensures SetOf(this) == old(SetOf(this)) + {x}

    method {:axiom} AddSet(s: MutableSet<T>)
      modifies this
      ensures SetOf(this) == old(SetOf(this)) + old(SetOf(s))
      decreases s

    method {:axiom} TransferSet(s: MutableSet<T>)
      modifies this, s
      ensures SetOf(this) == old(SetOf(s))
      ensures SetOf(s) == {}
      decreases s

    method {:axiom} Remove(x: T)
      modifies this
      ensures SetOf(this) == old(SetOf(this)) - {x}

    method {:axiom} RemoveAll()
      modifies this
      ensures SetOf(this) == {}
  }

  class MutableMap<K(==), V> {
    static function method {:axiom} MapOf(m: MutableMap<K, V>): map<K, V>
      reads m
      decreases {m}, m

    static method {:axiom} EmptyMap() returns (m: MutableMap<K, V>)
      ensures MapOf(m) == map[]
      ensures fresh(m)

    static method {:axiom} FromMap(dafny_map: map<K, V>) returns (m: MutableMap<K, V>)
      ensures MapOf(m) == dafny_map
      ensures fresh(m)
      decreases dafny_map

    constructor {:axiom} ()
      requires false

    function method {:axiom} Size(): int
      reads this
      ensures this.Size() == |MapOf(this)|
      decreases {this}

    method {:axiom} SizeModest() returns (size: uint64)
      requires |MapOf(this)| < 18446744073709551616
      ensures size as int == |MapOf(this)|

    method {:axiom} Contains(key: K) returns (contains: bool)
      ensures contains == (key in MapOf(this))

    method {:axiom} TryGetValue(key: K) returns (contains: bool, val: V)
      ensures contains == (key in MapOf(this))
      ensures contains ==> val == MapOf(this)[key]

    method {:axiom} Set(key: K, val: V)
      modifies this
      ensures MapOf(this) == old(MapOf(this))[key := val]

    method {:axiom} Remove(key: K)
      modifies this
      ensures MapOf(this) == map k: K {:trigger old(MapOf(this))[k]} {:trigger k in old(MapOf(this))} | k != key && k in old(MapOf(this)) :: old(MapOf(this))[k]
  }

  class Arrays {
    static method {:axiom} CopySeqIntoArray<A>(src: seq<A>, srcIndex: uint64, dst: array<A>, dstIndex: uint64, len: uint64)
      requires dst != null
      requires srcIndex as int + len as int <= |src|
      requires dstIndex as int + len as int <= dst.Length
      modifies dst
      ensures forall i: int :: 0 <= i < dst.Length ==> dst[i] == if dstIndex as int <= i < dstIndex as int + len as int then src[i - dstIndex as int + srcIndex as int] else old(dst[..])[i]
      ensures forall i: int :: srcIndex as int <= i < srcIndex as int + len as int ==> src[i] == dst[i - srcIndex as int + dstIndex as int]
      decreases src, srcIndex, dst, dstIndex, len
  }

  function {:axiom} realTimeBound(): int

  predicate AdvanceTime(oldTime: int, newTime: int, delay: int)
    decreases oldTime, newTime, delay
  {
    oldTime <= newTime < oldTime + delay + realTimeBound()
  }

  function MaxPacketSize(): int
  {
    65507
  }
}

abstract module DistributedSystem_s {

  import opened H_s : Host_s

  import opened Collections__Maps2_s = Collections__Maps2_s

  import opened Native__Io_s = Native__Io_s

  import opened Environment_s = Environment_s

  import opened Native__NativeTypes_s = Native__NativeTypes_s
  datatype DS_State = DS_State(config: ConcreteConfiguration, environment: LEnvironment<EndPoint, seq<byte>, HostStep>, servers: map<EndPoint, HostState>, clients: set<EndPoint>)

  predicate ValidPhysicalAddress(endPoint: EndPoint)
    decreases endPoint
  {
    |endPoint.addr| == 4 &&
    0 <= endPoint.port <= 65535
  }

  predicate ValidPhysicalPacket(p: LPacket<EndPoint, seq<byte>>)
    decreases p
  {
    ValidPhysicalAddress(p.src) &&
    ValidPhysicalAddress(p.dst) &&
    |p.msg| < 18446744073709551616
  }

  predicate ValidPhysicalIo(io: LIoOp<EndPoint, seq<byte>>)
    decreases io
  {
    (io.LIoOpReceive? ==>
      ValidPhysicalPacket(io.r)) &&
    (io.LIoOpSend? ==>
      ValidPhysicalPacket(io.s))
  }

  predicate ValidPhysicalEnvironmentStep(step: LEnvStep<EndPoint, seq<byte>, HostStep>)
    decreases step
  {
    step.LEnvStepHostIos? ==>
      forall io: LIoOp<EndPoint, seq<byte>> {:trigger io in step.ios} {:trigger ValidPhysicalIo(io)} :: 
        io in step.ios ==>
          ValidPhysicalIo(io)
  }

  predicate DS_Init(s: DS_State, config: ConcreteConfiguration)
    reads *
    decreases {}, s
  {
    s.config == config &&
    ConcreteConfigInit(s.config, mapdomain(s.servers), s.clients) &&
    LEnvironment_Init(s.environment) &&
    forall id: EndPoint :: 
      id in s.servers ==>
        HostInit(s.servers[id], config, id)
  }

  predicate DS_NextOneServer(s: DS_State, s': DS_State, id: EndPoint, ios: seq<LIoOp<EndPoint, seq<byte>>>)
    requires id in s.servers
    reads *
    decreases {}, s, s', id, ios
  {
    id in s'.servers &&
    HostNext(s.servers[id], s'.servers[id], ios) &&
    s'.servers == s.servers[id := s'.servers[id]]
  }

  predicate DS_Next(s: DS_State, s': DS_State)
    reads *
    decreases {}, s, s'
  {
    s'.config == s.config &&
    s'.clients == s.clients &&
    LEnvironment_Next(s.environment, s'.environment) &&
    ValidPhysicalEnvironmentStep(s.environment.nextStep) &&
    if s.environment.nextStep.LEnvStepHostIos? && s.environment.nextStep.actor in s.servers then DS_NextOneServer(s, s', s.environment.nextStep.actor, s.environment.nextStep.ios) else s'.servers == s.servers
  }
}

abstract module AbstractService_s {

  import opened Native__Io_s = Native__Io_s

  import opened Environment_s = Environment_s

  import opened Native__NativeTypes_s = Native__NativeTypes_s
  type ServiceState

  predicate Service_Init(s: ServiceState, serverAddresses: set<EndPoint>)
    decreases serverAddresses

  predicate Service_Next(s: ServiceState, s': ServiceState)

  predicate Service_Correspondence(concretePkts: set<LPacket<EndPoint, seq<byte>>>, serviceState: ServiceState)
    decreases concretePkts
}

module Collections__Seqs_s {
  function last<T>(s: seq<T>): T
    requires |s| > 0
    decreases s
  {
    s[|s| - 1]
  }

  function all_but_last<T>(s: seq<T>): seq<T>
    requires |s| > 0
    ensures |all_but_last(s)| == |s| - 1
    decreases s
  {
    s[..|s| - 1]
  }
}

module Collections__Maps2_s {
  function mapdomain<KT, VT>(m: map<KT, VT>): set<KT>
    decreases m
  {
    set k: KT {:trigger k in m} | k in m :: k
  }

  function mapremove<KT, VT>(m: map<KT, VT>, k: KT): map<KT, VT>
    decreases m
  {
    map ki: KT {:trigger m[ki]} {:trigger ki in m} | ki in m && ki != k :: m[ki]
  }

  predicate imaptotal<KT(!new), VT>(m: imap<KT, VT>)
  {
    forall k: KT {:trigger m[k]} {:trigger k in m} :: 
      k in m
  }
}

module Temporal__Temporal_s {

  import opened Collections__Maps2_s = Collections__Maps2_s
  type temporal = imap<int, bool>

  type Behavior<S> = imap<int, S>

  function {:axiom} stepmap(f: imap<int, bool>): temporal
    ensures forall i: int :: i in f ==> sat(i, stepmap(f)) == f[i]

  predicate {:axiom} sat(s: int, t: temporal)
    decreases s
  {
    s in t &&
    t[s]
  }

  function {:opaque} {:fuel 0, 0} and(x: temporal, y: temporal): temporal
    ensures forall i: int :: sat(i, and(x, y)) == (sat(i, x) && sat(i, y))
  {
    stepmap(imap i: int {:trigger sat(i, y)} {:trigger sat(i, x)} | true :: sat(i, x) && sat(i, y))
  }

  function {:opaque} {:fuel 0, 0} or(x: temporal, y: temporal): temporal
    ensures forall i: int :: sat(i, or(x, y)) == (sat(i, x) || sat(i, y))
  {
    stepmap(imap i: int {:trigger sat(i, y)} {:trigger sat(i, x)} | true :: sat(i, x) || sat(i, y))
  }

  function {:opaque} {:fuel 0, 0} imply(x: temporal, y: temporal): temporal
    ensures forall i: int :: sat(i, imply(x, y)) == (sat(i, x) ==> sat(i, y))
  {
    stepmap(imap i: int {:trigger sat(i, y)} {:trigger sat(i, x)} | true :: sat(i, x) ==> sat(i, y))
  }

  function {:opaque} {:fuel 0, 0} equiv(x: temporal, y: temporal): temporal
    ensures forall i: int :: sat(i, equiv(x, y)) == (sat(i, x) <==> sat(i, y))
  {
    stepmap(imap i: int {:trigger sat(i, y)} {:trigger sat(i, x)} | true :: sat(i, x) <==> sat(i, y))
  }

  function {:opaque} {:fuel 0, 0} not(x: temporal): temporal
    ensures forall i: int :: sat(i, not(x)) == !sat(i, x)
  {
    stepmap(imap i: int {:trigger sat(i, x)} | true :: !sat(i, x))
  }

  function {:opaque} {:fuel 0, 0} next(x: temporal): temporal
    ensures forall i: int {:trigger sat(i, next(x))} :: sat(i, next(x)) == sat(i + 1, x)
  {
    stepmap(imap i: int {:trigger sat(i + 1, x)} | true :: sat(i + 1, x))
  }

  function {:opaque} {:fuel 0, 0} always(x: temporal): temporal
  {
    stepmap(imap i: int {:trigger sat(i, always(x))} | true :: forall j: int {:trigger sat(j, x)} :: i <= j ==> sat(j, x))
  }

  function {:opaque} {:fuel 0, 0} eventual(x: temporal): temporal
  {
    stepmap(imap i: int {:trigger sat(i, eventual(x))} | true :: exists j: int :: i <= j && sat(j, x))
  }

  lemma {:axiom} {:auto_generated} {:opaque_reveal} {:verify false} {:fuel and, 1, 2} reveal_and()

  lemma {:axiom} {:auto_generated} {:opaque_reveal} {:verify false} {:fuel or, 1, 2} reveal_or()

  lemma {:axiom} {:auto_generated} {:opaque_reveal} {:verify false} {:fuel imply, 1, 2} reveal_imply()

  lemma {:axiom} {:auto_generated} {:opaque_reveal} {:verify false} {:fuel equiv, 1, 2} reveal_equiv()

  lemma {:axiom} {:auto_generated} {:opaque_reveal} {:verify false} {:fuel not, 1, 2} reveal_not()

  lemma {:axiom} {:auto_generated} {:opaque_reveal} {:verify false} {:fuel next, 1, 2} reveal_next()

  lemma {:axiom} {:auto_generated} {:opaque_reveal} {:verify false} {:fuel always, 1, 2} reveal_always()

  lemma {:axiom} {:auto_generated} {:opaque_reveal} {:verify false} {:fuel eventual, 1, 2} reveal_eventual()
}

abstract module Common__NodeIdentity_s {
  type NodeIdentity(==)
}

module Common__GenericMarshalling_i {

  import opened Native__NativeTypes_s = Native__NativeTypes_s

  import opened Collections__Maps_i = Collections__Maps_i

  import opened Collections__Seqs_i = Collections__Seqs_i

  import opened Logic__Option_i = Logic__Option_i

  import opened Common__Util_i = Common__Util_i

  import opened Common__MarshallInt_i = Common__MarshallInt_i
  datatype G = GUint64 | GArray(elt: G) | GTuple(t: seq<G>) | GByteArray | GTaggedUnion(cases: seq<G>)

  datatype V = VUint64(u: uint64) | VArray(a: seq<V>) | VTuple(t: seq<V>) | VByteArray(b: seq<byte>) | VCase(c: uint64, val: V)

  datatype ContentsTraceStep = ContentsTraceStep(data: seq<byte>, val: Option<V>)

  predicate ValInGrammar(val: V, grammar: G)
    decreases val, grammar
  {
    match val
    case VUint64(_) =>
      grammar.GUint64?
    case VArray(a) =>
      grammar.GArray? &&
      forall v :: 
        v in a ==>
          ValInGrammar(v, grammar.elt)
    case VTuple(t) =>
      grammar.GTuple? &&
      |t| == |grammar.t| &&
      forall i :: 
        0 <= i < |t| ==>
          ValInGrammar(t[i], grammar.t[i])
    case VByteArray(b) =>
      grammar.GByteArray?
    case VCase(c, v) =>
      grammar.GTaggedUnion? &&
      c as int < |grammar.cases| &&
      ValInGrammar(v, grammar.cases[c])
  }

  predicate ValidGrammar(grammar: G)
    decreases grammar
  {
    match grammar
    case GUint64 =>
      true
    case GArray(elt) =>
      ValidGrammar(elt)
    case GTuple(t) =>
      |t| < 18446744073709551616 &&
      forall g :: 
        g in t ==>
          ValidGrammar(g)
    case GByteArray =>
      true
    case GTaggedUnion(cases) =>
      |cases| < 18446744073709551616 &&
      forall g :: 
        g in cases ==>
          ValidGrammar(g)
  }

  predicate ValidVal(val: V)
    decreases val
  {
    match val
    case VUint64(_) =>
      true
    case VArray(a) =>
      |a| < 18446744073709551616 &&
      forall v :: 
        v in a ==>
          ValidVal(v)
    case VTuple(t) =>
      |t| < 18446744073709551616 &&
      forall v :: 
        v in t ==>
          ValidVal(v)
    case VByteArray(b) =>
      |b| < 18446744073709551616
    case VCase(c, v) =>
      ValidVal(v)
  }

  function {:opaque} {:fuel 0, 0} SeqSum(t: seq<V>): int
    ensures SeqSum(t) >= 0
    decreases t
  {
    if |t| == 0 then
      0
    else
      SizeOfV(t[0]) + SeqSum(t[1..])
  }

  function SizeOfV(val: V): int
    ensures SizeOfV(val) >= 0
    decreases val
  {
    match val
    case VUint64(_) =>
      8
    case VArray(a) =>
      8 + SeqSum(a)
    case VTuple(t) =>
      SeqSum(t)
    case VByteArray(b) =>
      8 + |b|
    case VCase(c, v) =>
      8 + SizeOfV(v)
  }

  function method parse_Uint64(data: seq<byte>): (Option<V>, seq<byte>)
    requires |data| < 18446744073709551616
    decreases data
  {
    if uint64(|data|) >= Uint64Size() then
      (Some(VUint64(SeqByteToUint64(data[..Uint64Size()]))), data[Uint64Size()..])
    else
      (None, [])
  }

  method ParseUint64(data: array<byte>, index: uint64)
      returns (success: bool, v: V, rest_index: uint64)
    requires data != null
    requires index as int <= data.Length
    requires data.Length < 18446744073709551616
    ensures rest_index as int <= data.Length
    ensures var (v': Option<V>, rest': seq<byte>) := parse_Uint64(data[index..]); var v_opt: Option<V> := if success then Some(v) else None(); v_opt == v' && data[rest_index..] == rest'
    decreases data, index
  {
    lemma_2toX();
    if uint64(data.Length) >= 8 && index <= uint64(data.Length) - 8 {
      var result := uint64(data[index + uint64(0)]) * 72057594037927936 + uint64(data[index + uint64(1)]) * 281474976710656 + uint64(data[index + uint64(2)]) * 1099511627776 + uint64(data[index + uint64(3)]) * 4294967296 + uint64(data[index + uint64(4)]) * 16777216 + uint64(data[index + uint64(5)]) * 65536 + uint64(data[index + uint64(6)]) * 256 + uint64(data[index + uint64(7)]);
      success := true;
      v := VUint64(result);
      rest_index := index + Uint64Size();
    } else {
      success := false;
      rest_index := uint64(data.Length);
    }
  }

  function method {:opaque} {:fuel 0, 0} parse_Array_contents(data: seq<byte>, eltType: G, len: uint64): (Option<seq<V>>, seq<byte>)
    requires |data| < 18446744073709551616
    requires ValidGrammar(eltType)
    ensures var (opt_seq: Option<seq<V>>, rest: seq<byte>) := parse_Array_contents(data, eltType, len); |rest| <= |data| && (opt_seq.Some? ==> forall i: int :: 0 <= i < |opt_seq.v| ==> ValInGrammar(opt_seq.v[i], eltType))
    decreases eltType, 1, len
  {
    if len == 0 then
      (Some([]), data)
    else
      var (val: Option<V>, rest1: seq<byte>) := parse_Val(data, eltType); var (others: Option<seq<V>>, rest2: seq<byte>) := parse_Array_contents(rest1, eltType, len - 1); if !val.None? && !others.None? then (Some([val.v] + others.v), rest2) else (None, [])
  }

  lemma /*{:_induction data, eltType, len}*/ lemma_ArrayContents_helper(data: seq<byte>, eltType: G, len: uint64, v: seq<V>, trace: seq<ContentsTraceStep>)
    requires |data| < 18446744073709551616
    requires ValidGrammar(eltType)
    requires |trace| == len as int + 1
    requires |v| == len as int
    requires forall j: int :: 0 <= j < |trace| ==> |trace[j].data| < 18446744073709551616
    requires trace[0].data == data
    requires forall j: int :: 0 < j < len as int + 1 ==> trace[j].val == parse_Val(trace[j - 1].data, eltType).0 && trace[j].data == parse_Val(trace[j - 1].data, eltType).1
    requires forall j: int :: 0 < j < |trace| ==> trace[j].val.Some?
    requires forall j: int :: 0 < j < |trace| ==> v[j - 1] == trace[j].val.v
    ensures var (v': Option<seq<V>>, rest': seq<byte>) := parse_Array_contents(data, eltType, len); var v_opt: Option<seq<V>> := Some(v); v_opt == v' && trace[|trace| - 1].data == rest'
    decreases len
  {
  }

  lemma /*{:_induction data, eltType, len}*/ lemma_ArrayContents_helper_bailout(data: seq<byte>, eltType: G, len: uint64, trace: seq<ContentsTraceStep>)
    requires |data| < 18446744073709551616
    requires ValidGrammar(eltType)
    requires 1 < |trace| <= len as int + 1
    requires forall j: int :: 0 <= j < |trace| ==> |trace[j].data| < 18446744073709551616
    requires trace[0].data == data
    requires forall j: int :: 0 < j < |trace| ==> trace[j].val == parse_Val(trace[j - 1].data, eltType).0 && trace[j].data == parse_Val(trace[j - 1].data, eltType).1
    requires forall j: int :: 0 < j < |trace| - 1 ==> trace[j].val.Some?
    requires trace[|trace| - 1].val.None?
    ensures var (v': Option<seq<V>>, rest': seq<byte>) := parse_Array_contents(data, eltType, len); v'.None? && rest' == []
    decreases len
  {
  }

  method {:timeLimit 60} ParseArrayContents(data: array<byte>, index: uint64, eltType: G, len: uint64)
      returns (success: bool, v: seq<V>, rest_index: uint64)
    requires data != null
    requires index as int <= data.Length
    requires data.Length < 18446744073709551616
    requires ValidGrammar(eltType)
    ensures rest_index as int <= data.Length
    ensures var (v': Option<seq<V>>, rest': seq<byte>) := parse_Array_contents(data[index..], eltType, len); var v_opt: Option<seq<V>> := if success then Some(v) else None(); v_opt == v' && data[rest_index..] == rest'
    ensures success ==> ValidVal(VArray(v))
    decreases eltType, 1, len
  {
    reveal_parse_Array_contents();
    var vArr := new V[len];
    ghost var g_v := [];
    success := true;
    var i: uint64 := 0;
    var next_val_index: uint64 := index;
    ghost var trace := [ContentsTraceStep(data[index..], None())];
    while i < len
      invariant 0 <= i <= len
      invariant index <= next_val_index <= uint64(data.Length)
      invariant |trace| == i as int + 1
      invariant |g_v| == i as int
      invariant vArr[..i] == g_v
      invariant trace[0].data == data[index..]
      invariant forall j: int :: 0 <= j < i as int + 1 ==> |trace[j].data| < 18446744073709551616
      invariant trace[i].data == data[next_val_index..]
      invariant forall j: uint64 :: 0 < j <= i ==> trace[j].val.Some?
      invariant forall j: uint64 :: 0 < j <= i ==> g_v[j - 1] == trace[j].val.v
      invariant forall j: int :: 0 < j < i as int + 1 ==> trace[j].val == parse_Val(trace[j - 1].data, eltType).0 && trace[j].data == parse_Val(trace[j - 1].data, eltType).1
      invariant ValidVal(VArray(vArr[..i]))
      decreases len as int - i as int
    {
      var some1, val, rest1 := ParseVal(data, next_val_index, eltType);
      ghost var step := ContentsTraceStep(data[rest1..], if some1 then Some(val) else None());
      ghost var old_trace := trace;
      trace := trace + [step];
      if !some1 {
        success := false;
        rest_index := uint64(data.Length);
        lemma_ArrayContents_helper_bailout(data[index..], eltType, len, trace);
        return;
      }
      g_v := g_v + [val];
      vArr[i] := val;
      next_val_index := rest1;
      i := i + 1;
    }
    success := true;
    rest_index := next_val_index;
    v := vArr[..];
    lemma_ArrayContents_helper(data[index..], eltType, len, v, trace);
  }

  function method parse_Array(data: seq<byte>, eltType: G): (Option<V>, seq<byte>)
    requires ValidGrammar(eltType)
    requires |data| < 18446744073709551616
    ensures var (opt_val: Option<V>, rest: seq<byte>) := parse_Array(data, eltType); |rest| <= |data| && (opt_val.Some? ==> ValInGrammar(opt_val.v, GArray(eltType)))
    decreases eltType
  {
    var (len: Option<V>, rest: seq<byte>) := parse_Uint64(data);
    if !len.None? then
      var (contents: Option<seq<V>>, remainder: seq<byte>) := parse_Array_contents(rest, eltType, len.v.u);
      if !contents.None? then
        (Some(VArray(contents.v)), remainder)
      else
        (None, [])
    else
      (None, [])
  }

  method ParseArray(data: array<byte>, index: uint64, eltType: G)
      returns (success: bool, v: V, rest_index: uint64)
    requires data != null
    requires index as int <= data.Length
    requires data.Length < 18446744073709551616
    requires ValidGrammar(eltType)
    ensures rest_index as int <= data.Length
    ensures var (v': Option<V>, rest': seq<byte>) := parse_Array(data[index..], eltType); var v_opt: Option<V> := if success then Some(v) else None(); v_opt == v' && data[rest_index..] == rest'
    ensures success ==> ValidVal(v)
    decreases eltType
  {
    var some1, len, rest := ParseUint64(data, index);
    if some1 {
      var some2, contents, remainder := ParseArrayContents(data, rest, eltType, len.u);
      if some2 {
        success := true;
        v := VArray(contents);
        rest_index := remainder;
      } else {
        success := false;
        rest_index := uint64(data.Length);
      }
    } else {
      success := false;
      rest_index := uint64(data.Length);
    }
  }

  function method {:opaque} {:fuel 0, 0} parse_Tuple_contents(data: seq<byte>, eltTypes: seq<G>): (Option<seq<V>>, seq<byte>)
    requires |data| < 18446744073709551616
    requires |eltTypes| < 18446744073709551616
    requires forall elt: G :: elt in eltTypes ==> ValidGrammar(elt)
    ensures var (opt_val: Option<seq<V>>, rest: seq<byte>) := parse_Tuple_contents(data, eltTypes); |rest| <= |data| && (opt_val.Some? ==> |opt_val.v| == |eltTypes| && forall i: int :: 0 <= i < |opt_val.v| ==> ValInGrammar(opt_val.v[i], eltTypes[i]))
    decreases eltTypes, 0
  {
    if eltTypes == [] then
      (Some([]), data)
    else
      var (val: Option<V>, rest1: seq<byte>) := parse_Val(data, eltTypes[uint64(0)]); assert |rest1| <= |data|; var (contents: Option<seq<V>>, rest2: seq<byte>) := parse_Tuple_contents(rest1, eltTypes[uint64(1)..]); if !val.None? && !contents.None? then (Some([val.v] + contents.v), rest2) else (None, [])
  }

  lemma /*{:_induction data, eltTypes}*/ lemma_TupleContents_helper(data: seq<byte>, eltTypes: seq<G>, v: seq<V>, trace: seq<ContentsTraceStep>)
    requires |data| < 18446744073709551616
    requires |eltTypes| < 18446744073709551616
    requires forall elt: G :: elt in eltTypes ==> ValidGrammar(elt)
    requires |trace| == |eltTypes| + 1
    requires |v| == |eltTypes| as int
    requires forall j: int :: 0 <= j < |trace| ==> |trace[j].data| < 18446744073709551616
    requires trace[0].data == data
    requires forall j: int :: 0 < j < |eltTypes| as int + 1 ==> trace[j].val == parse_Val(trace[j - 1].data, eltTypes[j - 1]).0 && trace[j].data == parse_Val(trace[j - 1].data, eltTypes[j - 1]).1
    requires forall j: int :: 0 < j < |trace| ==> trace[j].val.Some?
    requires forall j: int :: 0 < j < |trace| ==> v[j - 1] == trace[j].val.v
    ensures var (v': Option<seq<V>>, rest': seq<byte>) := parse_Tuple_contents(data, eltTypes); var v_opt: Option<seq<V>> := Some(v); v_opt == v' && trace[|trace| - 1].data == rest'
    decreases |eltTypes|
  {
  }

  lemma /*{:_induction data, eltTypes}*/ lemma_TupleContents_helper_bailout(data: seq<byte>, eltTypes: seq<G>, trace: seq<ContentsTraceStep>)
    requires |data| < 18446744073709551616
    requires |eltTypes| < 18446744073709551616
    requires forall elt: G :: elt in eltTypes ==> ValidGrammar(elt)
    requires 1 < |trace| <= |eltTypes| as int + 1
    requires forall j: int :: 0 <= j < |trace| ==> |trace[j].data| < 18446744073709551616
    requires trace[0].data == data
    requires forall j: int :: 0 < j < |trace| ==> trace[j].val == parse_Val(trace[j - 1].data, eltTypes[j - 1]).0 && trace[j].data == parse_Val(trace[j - 1].data, eltTypes[j - 1]).1
    requires forall j: int :: 0 < j < |trace| - 1 ==> trace[j].val.Some?
    requires trace[|trace| - 1].val.None?
    ensures var (v': Option<seq<V>>, rest': seq<byte>) := parse_Tuple_contents(data, eltTypes); v'.None? && rest' == []
    decreases |eltTypes|
  {
  }

  method {:timeLimit 60} ParseTupleContents(data: array<byte>, index: uint64, eltTypes: seq<G>)
      returns (success: bool, v: seq<V>, rest_index: uint64)
    requires data != null
    requires index as int <= data.Length
    requires data.Length < 18446744073709551616
    requires |eltTypes| < 18446744073709551616
    requires forall elt: G :: elt in eltTypes ==> ValidGrammar(elt)
    ensures rest_index as int <= data.Length
    ensures var (v': Option<seq<V>>, rest': seq<byte>) := parse_Tuple_contents(data[index..], eltTypes); var v_opt: Option<seq<V>> := if success then Some(v) else None(); v_opt == v' && data[rest_index..] == rest'
    ensures success ==> ValidVal(VTuple(v))
    decreases eltTypes, 0
  {
    reveal_parse_Tuple_contents();
    var vArr := new V[uint64(|eltTypes|)];
    ghost var g_v := [];
    success := true;
    var i: uint64 := 0;
    var next_val_index: uint64 := index;
    ghost var trace := [ContentsTraceStep(data[index..], None())];
    while i < uint64(|eltTypes|)
      invariant 0 <= i as int <= |eltTypes|
      invariant index <= next_val_index <= uint64(data.Length)
      invariant |trace| == i as int + 1
      invariant |g_v| == i as int
      invariant vArr[..i] == g_v
      invariant trace[0].data == data[index..]
      invariant forall j: int :: 0 <= j < i as int + 1 ==> |trace[j].data| < 18446744073709551616
      invariant trace[i].data == data[next_val_index..]
      invariant forall j: uint64 :: 0 < j <= i ==> trace[j].val.Some?
      invariant forall j: uint64 :: 0 < j <= i ==> g_v[j - 1] == trace[j].val.v
      invariant forall j: int :: 0 < j < i as int + 1 ==> trace[j].val == parse_Val(trace[j - 1].data, eltTypes[j - 1]).0 && trace[j].data == parse_Val(trace[j - 1].data, eltTypes[j - 1]).1
      invariant ValidVal(VTuple(vArr[..i]))
      decreases uint64(|eltTypes|) as int - i as int
    {
      var some1, val, rest1 := ParseVal(data, next_val_index, eltTypes[i]);
      ghost var step := ContentsTraceStep(data[rest1..], if some1 then Some(val) else None());
      ghost var old_trace := trace;
      trace := trace + [step];
      if !some1 {
        success := false;
        rest_index := uint64(data.Length);
        lemma_TupleContents_helper_bailout(data[index..], eltTypes, trace);
        return;
      }
      g_v := g_v + [val];
      vArr[i] := val;
      next_val_index := rest1;
      i := i + 1;
    }
    success := true;
    rest_index := next_val_index;
    v := vArr[..];
    lemma_TupleContents_helper(data[index..], eltTypes, v, trace);
  }

  function method parse_Tuple(data: seq<byte>, eltTypes: seq<G>): (Option<V>, seq<byte>)
    requires |data| < 18446744073709551616
    requires |eltTypes| < 18446744073709551616
    requires forall elt: G :: elt in eltTypes ==> ValidGrammar(elt)
    ensures var (opt_val: Option<V>, rest: seq<byte>) := parse_Tuple(data, eltTypes); |rest| <= |data| && (opt_val.Some? ==> ValInGrammar(opt_val.v, GTuple(eltTypes)))
    decreases eltTypes, 1
  {
    var (contents: Option<seq<V>>, rest: seq<byte>) := parse_Tuple_contents(data, eltTypes);
    if !contents.None? then
      (Some(VTuple(contents.v)), rest)
    else
      (None, [])
  }

  method ParseTuple(data: array<byte>, index: uint64, eltTypes: seq<G>)
      returns (success: bool, v: V, rest_index: uint64)
    requires data != null
    requires index as int <= data.Length
    requires data.Length < 18446744073709551616
    requires |eltTypes| < 18446744073709551616
    requires forall elt: G :: elt in eltTypes ==> ValidGrammar(elt)
    ensures rest_index as int <= data.Length
    ensures var (v': Option<V>, rest': seq<byte>) := parse_Tuple(data[index..], eltTypes); var v_opt: Option<V> := if success then Some(v) else None(); v_opt == v' && data[rest_index..] == rest'
    ensures success ==> ValidVal(v)
    decreases eltTypes, 1
  {
    var some, contents, rest := ParseTupleContents(data, index, eltTypes);
    if some {
      success := true;
      v := VTuple(contents);
      rest_index := rest;
    } else {
      success := false;
      rest_index := uint64(data.Length);
    }
  }

  function method parse_ByteArray(data: seq<byte>): (Option<V>, seq<byte>)
    requires |data| < 18446744073709551616
    decreases data
  {
    var (len: Option<V>, rest: seq<byte>) := parse_Uint64(data);
    if !len.None? && len.v.u <= uint64(|rest|) then
      (Some(VByteArray(rest[uint64(0) .. len.v.u])), rest[len.v.u..])
    else
      (None, [])
  }

  method ParseByteArray(data: array<byte>, index: uint64)
      returns (success: bool, v: V, rest_index: uint64)
    requires data != null
    requires index as int <= data.Length
    requires data.Length < 18446744073709551616
    ensures rest_index as int <= data.Length
    ensures var (v': Option<V>, rest': seq<byte>) := parse_ByteArray(data[index..]); var v_opt: Option<V> := if success then Some(v) else None(); v_opt == v' && data[rest_index..] == rest'
    decreases data, index
  {
    var some, len, rest := ParseUint64(data, index);
    if some && len.u <= uint64(data.Length) - rest {
      var rest_seq := data[rest..];
      assert len.u <= uint64(|rest_seq|);
      calc {
        rest_seq[0 .. len.u];
        data[rest .. rest + len.u];
      }
      success := true;
      v := VByteArray(data[rest .. rest + len.u]);
      rest_index := rest + len.u;
    } else {
      success := false;
      rest_index := uint64(data.Length);
    }
  }

  function method parse_Case(data: seq<byte>, cases: seq<G>): (Option<V>, seq<byte>)
    requires |data| < 18446744073709551616
    requires |cases| < 18446744073709551616
    requires forall elt: G :: elt in cases ==> ValidGrammar(elt)
    ensures var (opt_val: Option<V>, rest: seq<byte>) := parse_Case(data, cases); |rest| <= |data| && (opt_val.Some? ==> ValInGrammar(opt_val.v, GTaggedUnion(cases)))
    decreases cases
  {
    var (caseID: Option<V>, rest1: seq<byte>) := parse_Uint64(data);
    if !caseID.None? && caseID.v.u < uint64(|cases|) then
      var (val: Option<V>, rest2: seq<byte>) := parse_Val(rest1, cases[caseID.v.u]);
      if !val.None? then
        (Some(VCase(caseID.v.u, val.v)), rest2)
      else
        (None, [])
    else
      (None, [])
  }

  method ParseCase(data: array<byte>, index: uint64, cases: seq<G>)
      returns (success: bool, v: V, rest_index: uint64)
    requires data != null
    requires index as int <= data.Length
    requires data.Length < 18446744073709551616
    requires |cases| < 18446744073709551616
    requires forall elt: G :: elt in cases ==> ValidGrammar(elt)
    ensures rest_index as int <= data.Length
    ensures var (v': Option<V>, rest': seq<byte>) := parse_Case(data[index..], cases); var v_opt: Option<V> := if success then Some(v) else None(); v_opt == v' && data[rest_index..] == rest'
    ensures success ==> ValidVal(v)
    decreases cases
  {
    var some1, caseID, rest1 := ParseUint64(data, index);
    if some1 && caseID.u < uint64(|cases|) {
      var some2, val, rest2 := ParseVal(data, rest1, cases[caseID.u]);
      if some2 {
        success := true;
        v := VCase(caseID.u, val);
        rest_index := rest2;
      } else {
        success := false;
        rest_index := uint64(data.Length);
      }
    } else {
      success := false;
      rest_index := uint64(data.Length);
    }
  }

  function method {:opaque} {:fuel 0, 0} parse_Val(data: seq<byte>, grammar: G): (Option<V>, seq<byte>)
    requires |data| < 18446744073709551616
    requires ValidGrammar(grammar)
    ensures var (val: Option<V>, rest: seq<byte>) := parse_Val(data, grammar); |rest| <= |data| && (!val.None? ==> ValInGrammar(val.v, grammar))
    decreases grammar, 0
  {
    match grammar
    case GUint64 =>
      parse_Uint64(data)
    case GArray(elt) =>
      parse_Array(data, elt)
    case GTuple(t) =>
      parse_Tuple(data, t)
    case GByteArray =>
      parse_ByteArray(data)
    case GTaggedUnion(cases) =>
      parse_Case(data, cases)
  }

  method ParseVal(data: array<byte>, index: uint64, grammar: G)
      returns (success: bool, v: V, rest_index: uint64)
    requires data != null
    requires index as int <= data.Length
    requires data.Length < 18446744073709551616
    requires ValidGrammar(grammar)
    ensures rest_index as int <= data.Length
    ensures var (v': Option<V>, rest': seq<byte>) := parse_Val(data[index..], grammar); var v_opt: Option<V> := if success then Some(v) else None(); v_opt == v' && data[rest_index..] == rest'
    ensures success ==> ValidVal(v)
    decreases grammar, 0
  {
    reveal_parse_Val();
    match grammar {
      case GUint64 =>
        success, v, rest_index := ParseUint64(data, index);
      case GArray(elt) =>
        success, v, rest_index := ParseArray(data, index, elt);
      case GTuple(t) =>
        success, v, rest_index := ParseTuple(data, index, t);
      case GByteArray =>
        success, v, rest_index := ParseByteArray(data, index);
      case GTaggedUnion(cases) =>
        success, v, rest_index := ParseCase(data, index, cases);
    }
  }

  predicate Demarshallable(data: seq<byte>, grammar: G)
    decreases data, grammar
  {
    |data| < 18446744073709551616 &&
    ValidGrammar(grammar) &&
    !parse_Val(data, grammar).0.None? &&
    ValidVal(parse_Val(data, grammar).0.v) &&
    parse_Val(data, grammar).1 == []
  }

  function DemarshallFunc(data: seq<byte>, grammar: G): V
    requires Demarshallable(data, grammar)
    ensures var (val: Option<V>, rest: seq<byte>) := parse_Val(data, grammar); !val.None? && ValInGrammar(val.v, grammar)
    decreases grammar, 0
  {
    parse_Val(data, grammar).0.v
  }

  method Demarshall(data: array<byte>, grammar: G)
      returns (success: bool, v: V)
    requires data != null
    requires data.Length < 18446744073709551616
    requires ValidGrammar(grammar)
    ensures success == Demarshallable(data[..], grammar)
    ensures success ==> v == DemarshallFunc(data[..], grammar)
    decreases data, grammar
  {
    var rest: uint64;
    success, v, rest := ParseVal(data, 0, grammar);
    if success && rest == uint64(data.Length) {
      assert v == parse_Val(data[..], grammar).0.v;
      assert Demarshallable(data[..], grammar);
      assert v == DemarshallFunc(data[..], grammar);
    } else {
      success := false;
      assert !Demarshallable(data[..], grammar);
    }
  }

  lemma /*{:_induction v, grammar}*/ lemma_parse_Val_view_ByteArray(data: seq<byte>, v: V, grammar: G, index: int)
    requires |data| < 18446744073709551616
    requires ValInGrammar(v, grammar)
    requires ValidGrammar(grammar)
    requires grammar.GByteArray?
    requires 0 <= index <= |data|
    requires 0 <= index + SizeOfV(v) <= |data|
    ensures forall bound: int :: Trigger(bound) ==> index + SizeOfV(v) <= bound <= |data| ==> (parse_ByteArray(data[index .. bound]).0 == Some(v) <==> parse_ByteArray(data[index .. index + SizeOfV(v)]).0 == Some(v))
    ensures forall bound: int :: index + SizeOfV(v) <= bound <= |data| ==> parse_ByteArray(data[index .. bound]).0 == Some(v) ==> parse_ByteArray(data[index .. bound]).1 == data[index + SizeOfV(v) .. bound]
    decreases data, v, grammar, index
  {
  }

  lemma /*{:_induction s, v}*/ lemma_SeqSum_prefix(s: seq<V>, v: V)
    ensures SeqSum(s + [v]) == SeqSum(s) + SizeOfV(v)
    decreases s, v
  {
  }

  lemma /*{:_induction s}*/ lemma_SeqSum_bound(s: seq<V>, bound: int)
    requires SeqSum(s) < bound
    ensures forall v: V :: v in s ==> SizeOfV(v) < bound
    decreases s, bound
  {
  }

  lemma /*{:_induction s, prefix}*/ lemma_SeqSum_bound_prefix(s: seq<V>, prefix: seq<V>, index: int)
    requires 0 <= index <= |s|
    requires prefix == s[..index]
    ensures SeqSum(prefix) <= SeqSum(s)
    decreases s, prefix, index
  {
  }

  lemma /*{:_induction data, eltType, len}*/ lemma_parse_Array_contents_len(data: seq<byte>, eltType: G, len: uint64)
    requires |data| < 18446744073709551616
    requires ValidGrammar(eltType)
    requires len >= 0
    requires !parse_Array_contents(data, eltType, len).0.None?
    ensures len as int == |parse_Array_contents(data, eltType, len).0.v|
    decreases len
  {
  }

  lemma /*{:_induction vs, grammar, len}*/ lemma_parse_Val_view_Array_contents(data: seq<byte>, vs: seq<V>, grammar: G, index: int, bound: int, len: uint64)
    requires |data| < 18446744073709551616
    requires forall v: V :: v in vs ==> ValInGrammar(v, grammar)
    requires ValidGrammar(grammar)
    requires len as int == |vs|
    requires 0 <= index <= |data|
    requires 0 <= index + SeqSum(vs) <= |data|
    requires index + SeqSum(vs) <= bound <= |data|
    ensures parse_Array_contents(data[index .. bound], grammar, len).0 == Some(vs) <==> parse_Array_contents(data[index .. index + SeqSum(vs)], grammar, len).0 == Some(vs)
    ensures parse_Array_contents(data[index .. bound], grammar, len).0 == Some(vs) ==> parse_Array_contents(data[index .. bound], grammar, len).1 == data[index + SeqSum(vs) .. bound]
    decreases grammar, 1, len
  {
  }

  lemma /*{:_induction v, grammar}*/ lemma_parse_Val_view_Array(data: seq<byte>, v: V, grammar: G, index: int, bound: int)
    requires |data| < 18446744073709551616
    requires ValInGrammar(v, grammar)
    requires ValidGrammar(grammar)
    requires grammar.GArray?
    requires 0 <= index <= |data|
    requires 0 <= index + SizeOfV(v) <= |data|
    requires index + SizeOfV(v) <= bound <= |data|
    ensures parse_Array(data[index .. bound], grammar.elt).0 == Some(v) <==> parse_Array(data[index .. index + SizeOfV(v)], grammar.elt).0 == Some(v)
    ensures parse_Array(data[index .. bound], grammar.elt).0 == Some(v) ==> parse_Array(data[index .. bound], grammar.elt).1 == data[index + SizeOfV(v) .. bound]
    decreases grammar, -1
  {
  }

  lemma /*{:_induction vs, grammar}*/ lemma_parse_Val_view_Tuple_contents(data: seq<byte>, vs: seq<V>, grammar: seq<G>, index: int, bound: int)
    requires |data| < 18446744073709551616
    requires |vs| == |grammar|
    requires forall i: int :: 0 <= i < |vs| ==> ValInGrammar(vs[i], grammar[i])
    requires |grammar| < 18446744073709551616
    requires forall g: G :: g in grammar ==> ValidGrammar(g)
    requires 0 <= index <= |data|
    requires 0 <= index + SeqSum(vs) <= |data|
    requires index + SeqSum(vs) <= bound <= |data|
    ensures parse_Tuple_contents(data[index .. bound], grammar).0 == Some(vs) <==> parse_Tuple_contents(data[index .. index + SeqSum(vs)], grammar).0 == Some(vs)
    ensures parse_Tuple_contents(data[index .. bound], grammar).0 == Some(vs) ==> parse_Tuple_contents(data[index .. bound], grammar).1 == data[index + SeqSum(vs) .. bound]
    decreases grammar, -1, vs
  {
  }

  lemma /*{:_induction v, grammar}*/ lemma_parse_Val_view_Tuple(data: seq<byte>, v: V, grammar: seq<G>, index: int, bound: int)
    requires |data| < 18446744073709551616
    requires v.VTuple?
    requires |v.t| == |grammar|
    requires forall i: int :: 0 <= i < |v.t| ==> ValInGrammar(v.t[i], grammar[i])
    requires |grammar| < 18446744073709551616
    requires forall g: G :: g in grammar ==> ValidGrammar(g)
    requires 0 <= index <= |data|
    requires 0 <= index + SizeOfV(v) <= |data|
    requires index + SizeOfV(v) <= bound <= |data|
    ensures parse_Tuple(data[index .. bound], grammar).0 == Some(v) <==> parse_Tuple(data[index .. index + SizeOfV(v)], grammar).0 == Some(v)
    ensures parse_Tuple(data[index .. bound], grammar).0 == Some(v) ==> parse_Tuple(data[index .. bound], grammar).1 == data[index + SizeOfV(v) .. bound]
    decreases grammar, -1, v
  {
  }

  lemma /*{:_induction v, grammar}*/ lemma_parse_Val_view_Union(data: seq<byte>, v: V, grammar: G, index: int, bound: int)
    requires |data| < 18446744073709551616
    requires ValInGrammar(v, grammar)
    requires ValidGrammar(grammar)
    requires grammar.GTaggedUnion?
    requires 0 <= index <= |data|
    requires 0 <= index + SizeOfV(v) <= |data|
    requires index + SizeOfV(v) <= bound <= |data|
    ensures parse_Case(data[index .. bound], grammar.cases).0 == Some(v) <==> parse_Case(data[index .. index + SizeOfV(v)], grammar.cases).0 == Some(v)
    ensures parse_Case(data[index .. bound], grammar.cases).0 == Some(v) ==> parse_Case(data[index .. bound], grammar.cases).1 == data[index + SizeOfV(v) .. bound]
    decreases grammar, -1
  {
  }

  lemma /*{:_induction v, grammar}*/ lemma_parse_Val_view(data: seq<byte>, v: V, grammar: G, index: int)
    requires |data| < 18446744073709551616
    requires ValInGrammar(v, grammar)
    requires ValidGrammar(grammar)
    requires 0 <= index <= |data|
    requires 0 <= index + SizeOfV(v) <= |data|
    ensures forall bound: int :: index + SizeOfV(v) <= bound <= |data| ==> (parse_Val(data[index .. bound], grammar).0 == Some(v) <==> parse_Val(data[index .. index + SizeOfV(v)], grammar).0 == Some(v))
    ensures forall bound: int :: index + SizeOfV(v) <= bound <= |data| ==> parse_Val(data[index .. bound], grammar).0 == Some(v) ==> parse_Val(data[index .. bound], grammar).1 == data[index + SizeOfV(v) .. bound]
    decreases grammar, 0
  {
  }

  lemma /*{:_induction v, grammar}*/ lemma_parse_Val_view_specific(data: seq<byte>, v: V, grammar: G, index: int, bound: int)
    requires |data| < 18446744073709551616
    requires ValInGrammar(v, grammar)
    requires ValidGrammar(grammar)
    requires 0 <= index <= |data|
    requires 0 <= index + SizeOfV(v) <= |data|
    requires index + SizeOfV(v) <= bound <= |data|
    requires parse_Val(data[index .. index + SizeOfV(v)], grammar).0 == Some(v)
    ensures parse_Val(data[index .. bound], grammar).0 == Some(v)
    ensures parse_Val(data[index .. bound], grammar).1 == data[index + SizeOfV(v) .. bound]
    decreases grammar, 0
  {
  }

  lemma /*{:_induction v, grammar}*/ lemma_parse_Val_view_specific_size(data: seq<byte>, v: V, grammar: G, index: int, bound: int)
    requires |data| < 18446744073709551616
    requires ValInGrammar(v, grammar)
    requires ValidGrammar(grammar)
    requires 0 <= index <= |data|
    requires 0 <= index + SizeOfV(v) <= |data|
    requires index + SizeOfV(v) <= bound <= |data|
    requires parse_Val(data[index .. bound], grammar).0 == Some(v)
    ensures parse_Val(data[index .. index + SizeOfV(v)], grammar).0 == Some(v)
    ensures parse_Val(data[index .. bound], grammar).1 == data[index + SizeOfV(v) .. bound]
    decreases grammar, 0
  {
  }

  method ComputeSeqSum(s: seq<V>) returns (size: uint64)
    requires |s| < 18446744073709551616
    requires 0 <= SeqSum(s) < 18446744073709551616
    requires forall v: V :: v in s ==> ValidVal(v)
    ensures size as int == SeqSum(s)
    decreases s
  {
    reveal_SeqSum();
    if uint64(|s|) == 0 {
      size := 0;
    } else {
      var v_size := ComputeSizeOf(s[uint64(0)]);
      var rest_size := ComputeSeqSum(s[uint64(1)..]);
      size := v_size + rest_size;
    }
  }

  method ComputeSizeOf(val: V) returns (size: uint64)
    requires 0 <= SizeOfV(val) < 18446744073709551616
    requires ValidVal(val)
    ensures size as int == SizeOfV(val)
    decreases val
  {
    match val
    case VUint64(_) =>
      size := 8;
    case VArray(a) =>
      var v := ComputeSeqSum(a);
      if v == 0 {
        size := 8;
      } else {
        size := 8 + v;
      }
    case VTuple(t) =>
      size := ComputeSeqSum(t);
    case VByteArray(b) =>
      size := 8 + uint64(|b|);
    case VCase(c, v) =>
      var vs := ComputeSizeOf(v);
      size := 8 + vs;
  }

  method MarshallUint64(n: uint64, data: array<byte>, index: uint64)
    requires data != null
    requires index as int + Uint64Size() as int <= data.Length
    requires 0 <= index as int + Uint64Size() as int < 18446744073709551616
    requires data.Length < 18446744073709551616
    modifies data
    ensures SeqByteToUint64(data[index .. index + uint64(Uint64Size())]) == n
    ensures !parse_Uint64(data[index .. index + uint64(Uint64Size())]).0.None?
    ensures !parse_Uint64(data[index..]).0.None?
    ensures var tuple: (Option<V>, seq<byte>) := parse_Uint64(data[index .. index + uint64(Uint64Size())]); tuple.0.v.u == n && tuple.1 == []
    ensures var tuple: (Option<V>, seq<byte>) := parse_Uint64(data[index..]); tuple.0.v.u == n && tuple.1 == data[index + uint64(Uint64Size())..]
    ensures data[0 .. index] == old(data[0 .. index])
    ensures data[index + uint64(Uint64Size())..] == old(data[index + uint64(Uint64Size())..])
    ensures forall i: uint64 :: 0 <= i < index ==> data[i] == old(data[i])
    ensures forall i: int :: index as int + Uint64Size() as int <= i < data.Length ==> data[i] == old(data[i])
    decreases n, data, index
  {
    var tuple := parse_Uint64(data[index..]);
    MarshallUint64_guts(n, data, index);
  }

  lemma /*{:_induction contents, eltType, marshalled_bytes, trace}*/ lemma_marshall_array_contents(contents: seq<V>, eltType: G, marshalled_bytes: seq<byte>, trace: seq<seq<byte>>)
    requires forall v: V :: v in contents ==> ValInGrammar(v, eltType)
    requires forall v: V :: v in contents ==> ValidVal(v)
    requires ValidGrammar(eltType)
    requires |marshalled_bytes| < 18446744073709551616
    requires |contents| < 18446744073709551616
    requires |contents| == |trace|
    requires |marshalled_bytes| == SeqSum(contents)
    requires marshalled_bytes == SeqCatRev(trace)
    requires forall j: int :: 0 <= j < |trace| ==> SizeOfV(contents[j]) == |trace[j]| < 18446744073709551616
    requires forall j: int :: 0 <= j < |trace| ==> var (val: Option<V>, rest: seq<byte>) := parse_Val(trace[j], eltType); val.Some? && val.v == contents[j]
    ensures parse_Array_contents(marshalled_bytes, eltType, uint64(|contents|)).0.Some?
    ensures parse_Array_contents(marshalled_bytes, eltType, uint64(|contents|)).0.v == contents
    decreases contents, eltType, marshalled_bytes, trace
  {
  }

  method {:timeLimit 120} MarshallArrayContents(contents: seq<V>, eltType: G, data: array<byte>, index: uint64)
      returns (size: uint64)
    requires data != null
    requires forall v: V :: v in contents ==> ValInGrammar(v, eltType)
    requires forall v: V :: v in contents ==> ValidVal(v)
    requires ValidGrammar(eltType)
    requires index as int + SeqSum(contents) <= data.Length
    requires 0 <= index as int + SeqSum(contents) < 18446744073709551616
    requires data.Length < 18446744073709551616
    requires |contents| < 18446744073709551616
    modifies data
    ensures parse_Array_contents(data[index .. index as int + SeqSum(contents)], eltType, uint64(|contents|)).0.Some?
    ensures parse_Array_contents(data[index .. index as int + SeqSum(contents)], eltType, uint64(|contents|)).0.v == contents
    ensures forall i: uint64 :: 0 <= i < index ==> data[i] == old(data[i])
    ensures forall i: int :: index as int + SeqSum(contents) <= i < data.Length ==> data[i] == old(data[i])
    ensures size as int == SeqSum(contents)
    decreases eltType, 1, |contents|
  {
    var i: uint64 := 0;
    var cur_index := index;
    reveal_SeqSum();
    reveal_parse_Array_contents();
    ghost var trace := [];
    ghost var marshalled_bytes := [];
    while i < uint64(|contents|)
      invariant 0 <= i as int <= |contents|
      invariant 0 <= index as int <= index as int + SeqSum(contents[..i]) <= data.Length
      invariant cur_index as int == index as int + SeqSum(contents[..i])
      invariant forall j: uint64 :: 0 <= j < index ==> data[j] == old(data[j])
      invariant forall j: int :: index as int + SeqSum(contents) <= j < data.Length ==> data[j] == old(data[j])
      invariant marshalled_bytes == data[index .. cur_index]
      invariant marshalled_bytes == SeqCatRev(trace)
      invariant |trace| == i as int
      invariant forall j: int :: 0 <= j < |trace| ==> SizeOfV(contents[j]) == |trace[j]| < 18446744073709551616
      invariant forall j: int :: 0 <= j < |trace| ==> var (val: Option<V>, rest: seq<byte>) := parse_Val(trace[j], eltType); val.Some? && val.v == contents[j]
      decreases uint64(|contents|) as int - i as int
    {
      lemma_SeqSum_bound(contents, 18446744073709551616);
      calc <= {
        cur_index as int + SizeOfV(contents[i]);
        index as int + SeqSum(contents[..i]) + SizeOfV(contents[i]);
        {
          lemma_SeqSum_prefix(contents[..i], contents[i]);
          assert contents[..i] + [contents[i]] == contents[..i + 1];
        }
        index as int + SeqSum(contents[..i + 1]);
        {
          lemma_SeqSum_bound_prefix(contents, contents[..i + 1], i as int + 1);
        }
        index as int + SeqSum(contents);
      }
      var item_size := MarshallVal(contents[i], eltType, data, cur_index);
      ghost var fresh_bytes := data[cur_index .. cur_index + item_size];
      marshalled_bytes := marshalled_bytes + fresh_bytes;
      forall
        ensures var (val: Option<V>, rest: seq<byte>) := parse_Val(fresh_bytes, eltType); val.Some? && val.v == contents[i]
      {
        assert SizeOfV(contents[i]) <= |fresh_bytes|;
        lemma_parse_Val_view(fresh_bytes, contents[i], eltType, 0);
      }
      ghost var old_trace := trace;
      trace := trace + [fresh_bytes];
      ghost var old_cur_index := cur_index;
      cur_index := cur_index + item_size;
      i := i + 1;
      calc <= {
        index as int + SeqSum(contents[..i]);
        calc {
          SeqSum(contents[..i]);
        <=
          {
            lemma_SeqSum_bound_prefix(contents, contents[..i], i as int);
          }
          SeqSum(contents);
        }
        index as int + SeqSum(contents);
        data.Length;
      }
      assert {:split_here} true;
      assert marshalled_bytes == data[index .. cur_index];
      calc {
        cur_index as int;
        old_cur_index as int + SizeOfV(contents[i - 1]);
        index as int + SeqSum(contents[..i - 1]) + SizeOfV(contents[i - 1]);
        {
          lemma_SeqSum_prefix(contents[..i - 1], contents[i - 1]);
          assert contents[..i - 1] + [contents[i - 1]] == contents[..i];
        }
        index as int + SeqSum(contents[..i]);
      }
      assert cur_index as int == index as int + SeqSum(contents[..i]);
      assert marshalled_bytes == data[index .. cur_index];
    }
    assert contents[..i] == contents;
    assert cur_index as int == index as int + SeqSum(contents);
    assert marshalled_bytes == data[index .. index as int + SeqSum(contents)];
    lemma_marshall_array_contents(contents, eltType, marshalled_bytes, trace);
    size := cur_index - index;
  }

  method MarshallArray(val: V, grammar: G, data: array<byte>, index: uint64)
      returns (size: uint64)
    requires data != null
    requires val.VArray?
    requires ValInGrammar(val, grammar)
    requires ValidGrammar(grammar)
    requires ValidVal(val)
    requires index as int + SizeOfV(val) <= data.Length
    requires 0 <= index as int + SizeOfV(val) < 18446744073709551616
    requires data.Length < 18446744073709551616
    modifies data
    ensures parse_Val(data[index .. index as int + SizeOfV(val)], grammar).0.Some? && parse_Val(data[index .. index as int + SizeOfV(val)], grammar).0.v == val
    ensures forall i: uint64 :: 0 <= i < index ==> data[i] == old(data[i])
    ensures forall i: int :: index as int + SizeOfV(val) <= i < data.Length ==> data[i] == old(data[i])
    ensures size as int == SizeOfV(val)
    decreases grammar, -1
  {
    reveal_parse_Val();
    MarshallUint64(uint64(|val.a|), data, index);
    ghost var tuple := parse_Uint64(data[index .. index as int + SizeOfV(val)]);
    ghost var len := tuple.0;
    ghost var rest := tuple.1;
    assert !len.None?;
    var contents_size := MarshallArrayContents(val.a, grammar.elt, data, index + Uint64Size());
    tuple := parse_Uint64(data[index .. index as int + SizeOfV(val)]);
    assert {:split_here} true;
    len := tuple.0;
    rest := tuple.1;
    assert !len.None?;
    ghost var contents_tuple := parse_Array_contents(rest, grammar.elt, len.v.u);
    ghost var contents := contents_tuple.0;
    ghost var remainder := contents_tuple.1;
    assert !contents.None?;
    size := 8 + contents_size;
  }

  lemma /*{:_induction contents, eltTypes, marshalled_bytes, trace}*/ lemma_marshall_tuple_contents(contents: seq<V>, eltTypes: seq<G>, marshalled_bytes: seq<byte>, trace: seq<seq<byte>>)
    requires |contents| == |eltTypes|
    requires forall i: int :: 0 <= i < |contents| ==> ValInGrammar(contents[i], eltTypes[i])
    requires forall g: G :: g in eltTypes ==> ValidGrammar(g)
    requires |eltTypes| < 18446744073709551616
    requires forall i: int :: 0 <= i < |contents| ==> ValidVal(contents[i])
    requires |marshalled_bytes| < 18446744073709551616
    requires |contents| < 18446744073709551616
    requires |contents| == |trace|
    requires |marshalled_bytes| == SeqSum(contents)
    requires marshalled_bytes == SeqCatRev(trace)
    requires forall j: int :: 0 <= j < |trace| ==> SizeOfV(contents[j]) == |trace[j]| < 18446744073709551616
    requires forall j: int :: 0 <= j < |trace| ==> var (val: Option<V>, rest: seq<byte>) := parse_Val(trace[j], eltTypes[j]); val.Some? && val.v == contents[j]
    ensures parse_Tuple_contents(marshalled_bytes, eltTypes).0.Some?
    ensures parse_Tuple_contents(marshalled_bytes, eltTypes).0.v == contents
    decreases contents, eltTypes, marshalled_bytes, trace
  {
  }

  method {:timeLimit 60} MarshallTupleContents(contents: seq<V>, eltTypes: seq<G>, data: array<byte>, index: uint64)
      returns (size: uint64)
    requires data != null
    requires |contents| == |eltTypes|
    requires forall i: int :: 0 <= i < |contents| ==> ValInGrammar(contents[i], eltTypes[i])
    requires forall g: G :: g in eltTypes ==> ValidGrammar(g)
    requires |eltTypes| < 18446744073709551616
    requires forall i: int :: 0 <= i < |contents| ==> ValidVal(contents[i])
    requires index as int + SeqSum(contents) <= data.Length
    requires 0 <= index as int + SeqSum(contents) < 18446744073709551616
    requires data.Length < 18446744073709551616
    requires |contents| < 18446744073709551616
    modifies data
    ensures parse_Tuple_contents(data[index .. index as int + SeqSum(contents)], eltTypes).0.Some?
    ensures parse_Tuple_contents(data[index .. index as int + SeqSum(contents)], eltTypes).0.v == contents
    ensures forall i: uint64 :: 0 <= i < index ==> data[i] == old(data[i])
    ensures forall i: int :: index as int + SeqSum(contents) <= i < data.Length ==> data[i] == old(data[i])
    ensures size as int == SeqSum(contents)
    decreases eltTypes, 1, |contents|
  {
    var i: uint64 := 0;
    var cur_index := index;
    reveal_SeqSum();
    reveal_parse_Tuple_contents();
    ghost var trace := [];
    ghost var marshalled_bytes := [];
    while i < uint64(|contents|)
      invariant 0 <= i as int <= |contents|
      invariant 0 <= index as int <= index as int + SeqSum(contents[..i]) <= data.Length
      invariant cur_index as int == index as int + SeqSum(contents[..i])
      invariant forall j: uint64 :: 0 <= j < index ==> data[j] == old(data[j])
      invariant forall j: int :: index as int + SeqSum(contents) <= j < data.Length ==> data[j] == old(data[j])
      invariant marshalled_bytes == data[index .. cur_index]
      invariant marshalled_bytes == SeqCatRev(trace)
      invariant |trace| == i as int
      invariant forall j: int :: 0 <= j < |trace| ==> SizeOfV(contents[j]) == |trace[j]| < 18446744073709551616
      invariant forall j: int :: 0 <= j < |trace| ==> var (val: Option<V>, rest: seq<byte>) := parse_Val(trace[j], eltTypes[j]); val.Some? && val.v == contents[j]
      decreases uint64(|contents|) as int - i as int
    {
      lemma_SeqSum_bound(contents, 18446744073709551616);
      ghost var old_marshalled_bytes := marshalled_bytes;
      ghost var old_data := data[index .. cur_index];
      assert old_marshalled_bytes == old_data;
      calc <= {
        cur_index as int + SizeOfV(contents[i]);
        index as int + SeqSum(contents[..i]) + SizeOfV(contents[i]);
        {
          lemma_SeqSum_prefix(contents[..i], contents[i]);
          assert contents[..i] + [contents[i]] == contents[..i + 1];
        }
        index as int + SeqSum(contents[..i + 1]);
        {
          lemma_SeqSum_bound_prefix(contents, contents[..i + 1], i as int + 1);
        }
        index as int + SeqSum(contents);
      }
      var item_size := MarshallVal(contents[i], eltTypes[i], data, cur_index);
      ghost var fresh_bytes := data[cur_index .. cur_index + item_size];
      marshalled_bytes := marshalled_bytes + fresh_bytes;
      forall
        ensures var (val: Option<V>, rest: seq<byte>) := parse_Val(fresh_bytes, eltTypes[i]); val.Some? && val.v == contents[i]
      {
        assert SizeOfV(contents[i]) <= |fresh_bytes|;
        lemma_parse_Val_view(fresh_bytes, contents[i], eltTypes[i], 0);
      }
      ghost var old_trace := trace;
      trace := trace + [fresh_bytes];
      ghost var old_cur_index := cur_index;
      cur_index := cur_index + item_size;
      i := i + 1;
      assert {:split_here} true;
      calc {
        marshalled_bytes;
        old_marshalled_bytes + fresh_bytes;
        old_data + fresh_bytes;
        data[index .. old_cur_index] + fresh_bytes;
        data[index .. old_cur_index] + data[old_cur_index .. cur_index];
        data[index .. cur_index];
      }
      calc <= {
        index as int + SeqSum(contents[..i]);
        calc {
          SeqSum(contents[..i]);
        <=
          {
            lemma_SeqSum_bound_prefix(contents, contents[..i], i as int);
          }
          SeqSum(contents);
        }
        index as int + SeqSum(contents);
        data.Length;
      }
      calc {
        cur_index as int;
        old_cur_index as int + SizeOfV(contents[i - 1]);
        index as int + SeqSum(contents[..i - 1]) + SizeOfV(contents[i - 1]);
        {
          lemma_SeqSum_prefix(contents[..i - 1], contents[i - 1]);
          assert contents[..i - 1] + [contents[i - 1]] == contents[..i];
        }
        index as int + SeqSum(contents[..i]);
      }
      assert cur_index as int == index as int + SeqSum(contents[..i]);
    }
    assert contents[..i] == contents;
    assert cur_index as int == index as int + SeqSum(contents);
    assert marshalled_bytes == data[index .. index as int + SeqSum(contents)];
    lemma_marshall_tuple_contents(contents, eltTypes, marshalled_bytes, trace);
    size := cur_index - index;
  }

  method MarshallTuple(val: V, grammar: G, data: array<byte>, index: uint64)
      returns (size: uint64)
    requires data != null
    requires val.VTuple?
    requires ValidVal(val)
    requires ValidGrammar(grammar)
    requires ValInGrammar(val, grammar)
    requires index as int + SizeOfV(val) <= data.Length
    requires 0 <= index as int + SizeOfV(val) < 18446744073709551616
    requires data.Length < 18446744073709551616
    modifies data
    ensures parse_Val(data[index .. index as int + SizeOfV(val)], grammar).0.Some? && parse_Val(data[index .. index as int + SizeOfV(val)], grammar).0.v == val
    ensures forall i: uint64 :: 0 <= i < index ==> data[i] == old(data[i])
    ensures forall i: int :: index as int + SizeOfV(val) <= i < data.Length ==> data[i] == old(data[i])
    ensures size as int == SizeOfV(val)
    decreases grammar, -1
  {
    size := MarshallTupleContents(val.t, grammar.t, data, index);
    calc {
      parse_Val(data[index .. index as int + SizeOfV(val)], grammar);
      {
        reveal_parse_Val();
      }
      parse_Tuple(data[index .. index as int + SizeOfV(val)], grammar.t);
    }
  }

  method MarshallBytes(bytes: seq<byte>, data: array<byte>, index: uint64)
    requires data != null
    requires index as int + |bytes| <= data.Length
    requires 0 <= index as int + |bytes| < 18446744073709551616
    requires data.Length < 18446744073709551616
    modifies data
    ensures forall i: uint64 :: 0 <= i < index ==> data[i] == old(data[i])
    ensures forall i: int :: index as int <= i < index as int + |bytes| ==> data[i] == bytes[i - index as int]
    ensures forall i: int :: index as int + |bytes| <= i < data.Length ==> data[i] == old(data[i])
    decreases bytes, data, index
  {
    Arrays.CopySeqIntoArray(bytes, 0, data, index, uint64(|bytes|));
  }

  method MarshallByteArray(val: V, grammar: G, data: array<byte>, index: uint64)
      returns (size: uint64)
    requires data != null
    requires val.VByteArray?
    requires ValidGrammar(grammar)
    requires ValInGrammar(val, grammar)
    requires ValidVal(val)
    requires index as int + SizeOfV(val) <= data.Length
    requires 0 <= index as int + SizeOfV(val) < 18446744073709551616
    requires data.Length < 18446744073709551616
    modifies data
    ensures parse_Val(data[index .. index as int + SizeOfV(val)], grammar).0.Some? && parse_Val(data[index .. index as int + SizeOfV(val)], grammar).0.v == val
    ensures forall i: uint64 :: 0 <= i < index ==> data[i] == old(data[i])
    ensures forall i: int :: index as int + SizeOfV(val) <= i < data.Length ==> data[i] == old(data[i])
    ensures size as int == SizeOfV(val)
    decreases grammar
  {
    MarshallUint64(uint64(|val.b|), data, index);
    assert SeqByteToUint64(data[index .. index + uint64(Uint64Size())]) == uint64(|val.b|);
    MarshallBytes(val.b, data, index + 8);
    calc {
      parse_Val(data[index .. index as int + SizeOfV(val)], grammar);
      {
        reveal_parse_Val();
      }
      parse_ByteArray(data[index .. index as int + SizeOfV(val)]);
    }
    ghost var data_seq := data[index .. index as int + SizeOfV(val)];
    ghost var tuple := parse_Uint64(data_seq);
    ghost var len := tuple.0;
    ghost var rest := tuple.1;
    assert {:split_here} true;
    assert rest == data[index + 8 .. index as int + SizeOfV(val)] == val.b;
    assert !len.None? && len.v.u as int <= |rest|;
    assert rest[0 .. len.v.u] == val.b;
    size := 8 + uint64(|val.b|);
  }

  method MarshallCase(val: V, grammar: G, data: array<byte>, index: uint64)
      returns (size: uint64)
    requires data != null
    requires val.VCase?
    requires ValidGrammar(grammar)
    requires ValInGrammar(val, grammar)
    requires ValidVal(val)
    requires index as int + SizeOfV(val) <= data.Length
    requires 0 <= index as int + SizeOfV(val) < 18446744073709551616
    requires data.Length < 18446744073709551616
    modifies data
    ensures parse_Val(data[index .. index as int + SizeOfV(val)], grammar).0.Some? && parse_Val(data[index .. index as int + SizeOfV(val)], grammar).0.v == val
    ensures forall i: uint64 :: 0 <= i < index ==> data[i] == old(data[i])
    ensures forall i: int :: index as int + SizeOfV(val) <= i < data.Length ==> data[i] == old(data[i])
    ensures size as int == SizeOfV(val)
    decreases grammar, -1
  {
    MarshallUint64(val.c, data, index);
    ghost var int_bytes := data[index .. index + Uint64Size()];
    ghost var tuple0 := parse_Uint64(int_bytes);
    ghost var caseID0 := tuple0.0;
    ghost var rest10 := tuple0.1;
    assert !caseID0.None?;
    assert caseID0.v.u == val.c;
    var val_size := MarshallVal(val.val, grammar.cases[val.c], data, index + 8);
    ghost var new_int_bytes := data[index .. index + Uint64Size()];
    assert forall i: uint64 {:auto_trigger} :: 0 <= i < Uint64Size() ==> int_bytes[i] == new_int_bytes[i];
    assert int_bytes == new_int_bytes;
    assert val.VCase?;
    assert grammar.GTaggedUnion?;
    assert val.c as int < |grammar.cases|;
    ghost var bytes := data[index .. index as int + SizeOfV(val)];
    assert bytes[..8] == new_int_bytes;
    calc {
      parse_Val(bytes, grammar);
      {
        reveal_parse_Val();
      }
      parse_Case(bytes, grammar.cases);
    }
    assert {:split_here} true;
    ghost var tuple1 := parse_Uint64(bytes);
    ghost var caseID := tuple1.0;
    ghost var rest1 := tuple1.1;
    assert !caseID.None?;
    assert caseID.v.u == val.c;
    assert caseID.v.u as int < |grammar.cases|;
    ghost var tuple2 := parse_Val(rest1, grammar.cases[caseID.v.u]);
    ghost var v := tuple2.0;
    ghost var rest2 := tuple2.1;
    assert !v.None?;
    size := 8 + val_size;
  }

  method MarshallVUint64(val: V, grammar: G, data: array<byte>, index: uint64)
      returns (size: uint64)
    requires data != null
    requires val.VUint64?
    requires ValidGrammar(grammar)
    requires ValInGrammar(val, grammar)
    requires index as int + SizeOfV(val) <= data.Length
    requires 0 <= index as int + SizeOfV(val) < 18446744073709551616
    requires data.Length < 18446744073709551616
    modifies data
    ensures parse_Val(data[index .. index as int + SizeOfV(val)], grammar).0.Some? && parse_Val(data[index .. index as int + SizeOfV(val)], grammar).0.v == val
    ensures forall i: uint64 :: 0 <= i < index ==> data[i] == old(data[i])
    ensures forall i: int :: index as int + SizeOfV(val) <= i < data.Length ==> data[i] == old(data[i])
    ensures size as int == SizeOfV(val)
    decreases grammar
  {
    MarshallUint64(val.u, data, index);
    calc {
      parse_Val(data[index .. index as int + SizeOfV(val)], grammar);
      {
        reveal_parse_Val();
      }
      parse_Uint64(data[index .. index as int + SizeOfV(val)]);
    }
    return 8;
  }

  method MarshallVal(val: V, grammar: G, data: array<byte>, index: uint64)
      returns (size: uint64)
    requires ValidGrammar(grammar)
    requires ValInGrammar(val, grammar)
    requires ValidVal(val)
    requires 0 <= SizeOfV(val) < 18446744073709551616
    requires data != null
    requires index as int + SizeOfV(val) <= data.Length
    requires data.Length < 18446744073709551616
    modifies data
    ensures parse_Val(data[index .. index as int + SizeOfV(val)], grammar).0.Some? && parse_Val(data[index .. index as int + SizeOfV(val)], grammar).0.v == val
    ensures forall i: uint64 :: 0 <= i < index ==> data[i] == old(data[i])
    ensures forall i: int :: index as int + SizeOfV(val) <= i < data.Length ==> data[i] == old(data[i])
    ensures size as int == SizeOfV(val)
    decreases grammar, 0
  {
    match val
    case VUint64(_) =>
      size := MarshallVUint64(val, grammar, data, index);
    case VArray(_) =>
      size := MarshallArray(val, grammar, data, index);
    case VTuple(_) =>
      size := MarshallTuple(val, grammar, data, index);
    case VByteArray(_) =>
      size := MarshallByteArray(val, grammar, data, index);
    case VCase(_, _) =>
      size := MarshallCase(val, grammar, data, index);
  }

  method Marshall(val: V, grammar: G) returns (data: array<byte>)
    requires ValidGrammar(grammar)
    requires ValInGrammar(val, grammar)
    requires ValidVal(val)
    requires 0 <= SizeOfV(val) < 18446744073709551616
    ensures data != null
    ensures fresh(data)
    ensures Demarshallable(data[..], grammar)
    ensures parse_Val(data[..], grammar).0.Some? && parse_Val(data[..], grammar).0.v == val
    ensures parse_Val(data[..], grammar).1 == []
    ensures |data[..]| == SizeOfV(val)
    decreases val, grammar
  {
    var size := ComputeSizeOf(val);
    data := new byte[size];
    var computed_size := MarshallVal(val, grammar, data, 0);
    assert data[0 .. 0 + SizeOfV(val)] == data[0 .. 0 + size] == data[..];
    lemma_parse_Val_view_specific(data[..], val, grammar, 0, size as int);
  }

  lemma {:axiom} {:auto_generated} {:opaque_reveal} {:verify false} {:fuel SeqSum, 1, 2} reveal_SeqSum()

  lemma {:axiom} {:auto_generated} {:opaque_reveal} {:verify false} {:fuel parse_Array_contents, 1, 2} reveal_parse_Array_contents()

  lemma {:axiom} {:auto_generated} {:opaque_reveal} {:verify false} {:fuel parse_Tuple_contents, 1, 2} reveal_parse_Tuple_contents()

  lemma {:axiom} {:auto_generated} {:opaque_reveal} {:verify false} {:fuel parse_Val, 1, 2} reveal_parse_Val()
}

module Common__UdpClient_i {

  import opened Native__Io_s = Native__Io_s
  function Workaround_CastHostEnvironmentToObject(env: HostEnvironment): object
    decreases env
  {
    env
  }

  function Workaround_CastOkStateToObject(okState: OkState): object
    decreases okState
  {
    okState
  }

  function Workaround_CastNowStateToObject(nowState: NowState): object
    decreases nowState
  {
    nowState
  }

  function Workaround_CastUdpStateToObject(udpState: UdpState): object
    decreases udpState
  {
    udpState
  }

  function Workaround_CastIPEndPointToObject(ip: IPEndPoint): object
    decreases ip
  {
    ip
  }

  function Workaround_CastUdpClientToObject(udpc: UdpClient): object
    decreases udpc
  {
    udpc
  }

  function HostEnvironmentDefaultFrame(env: HostEnvironment): set<object>
    reads env, if env != null then {env.now} else {}, if env != null then {env.ok} else {}, if env != null then {env.udp} else {}, if env != null then {env} else {}
    decreases (if env != null then {env.now} else {}) + (if env != null then {env.ok} else {}) + (if env != null then {env.udp} else {}) + (if env != null then {env} else {}) + {env}, env
  {
    if env != null then
      {Workaround_CastOkStateToObject(env.ok), Workaround_CastNowStateToObject(env.now), Workaround_CastUdpStateToObject(env.udp)}
    else
      {}
  }

  function UdpClientRepr(udpc: UdpClient): set<object>
    reads udpc, if udpc != null then HostEnvironmentDefaultFrame.reads(udpc.env) else {}
    decreases (if udpc != null then HostEnvironmentDefaultFrame.reads(udpc.env) else {}) + {udpc}, udpc
  {
    {Workaround_CastUdpClientToObject(udpc)} + if udpc != null then HostEnvironmentDefaultFrame(udpc.env) else {}
  }

  predicate HostEnvironmentIsValid(env: HostEnvironment)
    reads env, if env != null then env.Valid.reads() else {}, if env != null && env.ok != null then env.ok.ok.reads() else {}
    decreases (if env != null then env.Valid.reads() else {}) + (if env != null && env.ok != null then env.ok.ok.reads() else {}) + {env}, env
  {
    env != null &&
    env.Valid() &&
    env.constants != null &&
    env.now != null &&
    env.ok != null &&
    env.ok.ok() &&
    env.udp != null
  }

  predicate UdpClientOk(udpc: UdpClient)
    reads udpc, if udpc != null then HostEnvironmentDefaultFrame.reads(udpc.env) else {}
    decreases (if udpc != null then HostEnvironmentDefaultFrame.reads(udpc.env) else {}) + {udpc}, udpc
  {
    udpc != null &&
    udpc.env != null &&
    udpc.env.ok != null &&
    udpc.env.ok.ok()
  }

  function method EndPointIsValidIPV4(endPoint: EndPoint): bool
    decreases endPoint
  {
    |endPoint.addr| == 4 &&
    0 <= endPoint.port <= 65535
  }

  predicate UdpClientIsValid(udpc: UdpClient)
    reads UdpClientRepr(udpc), if udpc != null then HostEnvironmentIsValid.reads(udpc.env) else {}
    decreases UdpClientRepr(udpc) + if udpc != null then HostEnvironmentIsValid.reads(udpc.env) else {}, udpc
  {
    udpc != null &&
    udpc.env != null &&
    udpc.IsOpen() &&
    HostEnvironmentIsValid(udpc.env) &&
    EndPointIsValidIPV4(udpc.LocalEndPoint())
  }

  predicate EndPointsAreValidIPV4(eps: seq<EndPoint>)
    decreases eps
  {
    forall i: int :: 
      0 <= i < |eps| ==>
        EndPointIsValidIPV4(eps[i])
  }
}

module Message_i {

  import opened Types_i = Types_i
  datatype CMessage = CTransfer(transfer_epoch: uint64) | CLocked(locked_epoch: uint64) | CInvalid

  type CLockPacket = LPacket<EndPoint, CMessage>

  function AbstractifyCMessage(cmsg: CMessage): LockMessage
    decreases cmsg
  {
    match cmsg {
      case CTransfer(epoch) =>
        Transfer(epoch as int)
      case CLocked(epoch) =>
        Locked(epoch as int)
      case CInvalid =>
        Invalid()
    }
  }

  function AbstractifyCLockPacket(p: CLockPacket): LockPacket
    decreases p
  {
    LPacket(p.dst, p.src, AbstractifyCMessage(p.msg))
  }
}

abstract module Host_s {

  import opened Native__Io_s = Native__Io_s

  import opened Environment_s = Environment_s

  import opened Native__NativeTypes_s = Native__NativeTypes_s
  type HostState

  type HostStep(!new)

  type ConcreteConfiguration

  predicate HostInit(host_state: HostState, config: ConcreteConfiguration, id: EndPoint)
    reads *
    decreases {}, id

  predicate HostNext(host_state: HostState, host_state': HostState, ios: seq<LIoOp<EndPoint, seq<byte>>>)
    reads *
    decreases {}, ios

  predicate ConcreteConfigInit(config: ConcreteConfiguration, servers: set<EndPoint>, clients: set<EndPoint>)
    decreases servers, clients

  predicate HostStateInvariants(host_state: HostState, env: HostEnvironment)
    reads *
    decreases {}, env

  predicate ConcreteConfigurationInvariants(config: ConcreteConfiguration)

  function ParseCommandLineConfiguration(args: seq<seq<uint16>>): (ConcreteConfiguration, set<EndPoint>, set<EndPoint>)
    decreases args

  function ParseCommandLineId(ip: seq<uint16>, port: seq<uint16>): EndPoint
    decreases ip, port

  method HostInitImpl(ghost env: HostEnvironment)
      returns (ok: bool, host_state: HostState, config: ConcreteConfiguration, ghost servers: set<EndPoint>, ghost clients: set<EndPoint>, id: EndPoint)
    requires env != null && env.Valid()
    requires env.ok.ok()
    requires |env.constants.CommandLineArgs()| >= 2
    modifies set x: object | true
    ensures ok ==> env != null && env.Valid() && env.ok.ok()
    ensures ok ==> |env.constants.CommandLineArgs()| >= 2
    ensures ok ==> HostStateInvariants(host_state, env)
    ensures ok ==> ConcreteConfigurationInvariants(config)
    ensures ok ==> var args: seq<seq<uint16>> := env.constants.CommandLineArgs(); var (parsed_config: ConcreteConfiguration, parsed_servers: set<EndPoint>, parsed_clients: set<EndPoint>) := ParseCommandLineConfiguration(args[0 .. |args| - 2]); config == parsed_config && servers == parsed_servers && clients == parsed_clients && ConcreteConfigInit(parsed_config, parsed_servers, parsed_clients)
    ensures ok ==> var args: seq<seq<uint16>> := env.constants.CommandLineArgs(); id == ParseCommandLineId(args[|args| - 2], args[|args| - 1]) && HostInit(host_state, config, id)
    decreases env

  method HostNextImpl(ghost env: HostEnvironment, host_state: HostState)
      returns (ok: bool, host_state': HostState, ghost recvs: seq<UdpEvent>, ghost clocks: seq<UdpEvent>, ghost sends: seq<UdpEvent>, ghost ios: seq<LIoOp<EndPoint, seq<byte>>>)
    requires env != null && env.Valid() && env.ok.ok()
    requires HostStateInvariants(host_state, env)
    modifies set x: object | true
    ensures ok <==> env != null && env.Valid() && env.ok.ok()
    ensures ok ==> HostStateInvariants(host_state', env)
    ensures ok ==> HostNext(host_state, host_state', ios)
    ensures ok ==> recvs + clocks + sends == ios
    ensures ok ==> env.udp.history() == old(env.udp.history()) + (recvs + clocks + sends)
    ensures forall e: LIoOp<EndPoint, seq<byte>> :: (e in recvs ==> e.LIoOpReceive?) && (e in clocks ==> e.LIoOpReadClock? || e.LIoOpTimeoutReceive?) && (e in sends ==> e.LIoOpSend?)
    ensures |clocks| <= 1
    decreases env
}

module Collections__Sets_i {
  lemma ThingsIKnowAboutSubset<T>(x: set<T>, y: set<T>)
    requires x < y
    ensures |x| < |y|
    decreases x, y
  {
  }

  lemma SubsetCardinality<T>(x: set<T>, y: set<T>)
    ensures x < y ==> |x| < |y|
    ensures x <= y ==> |x| <= |y|
    decreases x, y
  {
  }

  lemma ItIsASingletonSet<T>(foo: set<T>, x: T)
    requires foo == {x}
    ensures |foo| == 1
    decreases foo
  {
  }

  lemma ThingsIKnowAboutASingletonSet<T>(foo: set<T>, x: T, y: T)
    requires |foo| == 1
    requires x in foo
    requires y in foo
    ensures x == y
    decreases foo
  {
  }

  predicate Injective<X(!new), Y>(f: X -> Y)
    requires forall x: X :: f.requires(x)
    reads f.reads
    decreases set _x0: X, _o0: object? | _o0 in f.reads(_x0) :: _o0
  {
    forall x1: X, x2: X :: 
      f(x1) == f(x2) ==>
        x1 == x2
  }

  predicate InjectiveOver<X, Y>(xs: set<X>, ys: set<Y>, f: X -> Y)
    requires forall x: X :: x in xs ==> f.requires(x)
    reads f.reads
    decreases set _x0: X, _o0: object? | _o0 in f.reads(_x0) :: _o0, xs, ys
  {
    forall x1: X, x2: X :: 
      x1 in xs &&
      x2 in xs &&
      f(x1) in ys &&
      f(x2) in ys &&
      f(x1) == f(x2) ==>
        x1 == x2
  }

  predicate InjectiveOverSeq<X, Y>(xs: seq<X>, ys: set<Y>, f: X -> Y)
    requires forall x: X :: x in xs ==> f.requires(x)
    reads f.reads
    decreases set _x0: X, _o0: object? | _o0 in f.reads(_x0) :: _o0, xs, ys
  {
    forall x1: X, x2: X :: 
      x1 in xs &&
      x2 in xs &&
      f(x1) in ys &&
      f(x2) in ys &&
      f(x1) == f(x2) ==>
        x1 == x2
  }

  lemma lemma_MapSetCardinality<X, Y>(xs: set<X>, ys: set<Y>, f: X -> Y)
    requires forall x: X :: f.requires(x)
    requires Injective(f)
    requires forall x: X :: x in xs <==> f(x) in ys
    requires forall y: Y :: y in ys ==> exists x: X :: x in xs && y == f(x)
    ensures |xs| == |ys|
    decreases xs, ys
  {
  }

  lemma lemma_MapSetCardinalityOver<X, Y>(xs: set<X>, ys: set<Y>, f: X -> Y)
    requires forall x: X :: x in xs ==> f.requires(x)
    requires InjectiveOver(xs, ys, f)
    requires forall x: X :: x in xs ==> f(x) in ys
    requires forall y: Y :: y in ys ==> exists x: X :: x in xs && y == f(x)
    ensures |xs| == |ys|
    decreases xs, ys
  {
  }

  lemma lemma_MapSubsetCardinalityOver<X, Y>(xs: set<X>, ys: set<Y>, f: X -> Y)
    requires forall x: X :: x in xs ==> f.requires(x)
    requires InjectiveOver(xs, ys, f)
    requires forall x: X :: x in xs ==> f(x) in ys
    ensures |xs| <= |ys|
    decreases xs, ys
  {
  }

  lemma lemma_MapSubseqCardinalityOver<X, Y>(xs: seq<X>, ys: set<Y>, f: X -> Y)
    requires forall x: X :: x in xs ==> f.requires(x)
    requires forall i: int, j: int :: 0 <= i < |xs| && 0 <= j < |xs| && i != j ==> xs[i] != xs[j]
    requires InjectiveOverSeq(xs, ys, f)
    requires forall x: X :: x in xs ==> f(x) in ys
    ensures |xs| <= |ys|
    decreases xs, ys
  {
  }

  function MapSetToSet<X(!new), Y>(xs: set<X>, f: X -> Y): set<Y>
    requires forall x: X :: f.requires(x)
    requires Injective(f)
    reads f.reads
    ensures forall x: X :: x in xs <==> f(x) in MapSetToSet(xs, f)
    ensures |xs| == |MapSetToSet(xs, f)|
    decreases set _x0: X, _o0: object? | _o0 in f.reads(_x0) :: _o0, xs
  {
    var ys: set<Y> := set x: X {:trigger f(x)} {:trigger x in xs} | x in xs :: f(x);
    lemma_MapSetCardinality(xs, ys, f);
    ys
  }

  function MapSetToSetOver<X, Y>(xs: set<X>, f: X -> Y): set<Y>
    requires forall x: X :: x in xs ==> f.requires(x)
    requires InjectiveOver(xs, set x: X {:trigger f(x)} {:trigger x in xs} | x in xs :: f(x), f)
    reads f.reads
    ensures forall x: X :: x in xs ==> f(x) in MapSetToSetOver(xs, f)
    ensures |xs| == |MapSetToSetOver(xs, f)|
    decreases set _x0: X, _o0: object? | _o0 in f.reads(_x0) :: _o0, xs
  {
    var ys: set<Y> := set x: X {:trigger f(x)} {:trigger x in xs} | x in xs :: f(x);
    lemma_MapSetCardinalityOver(xs, ys, f);
    ys
  }

  function MapSeqToSet<X(!new), Y>(xs: seq<X>, f: X -> Y): set<Y>
    requires forall x: X :: f.requires(x)
    requires Injective(f)
    reads f.reads
    ensures forall x: X :: x in xs <==> f(x) in MapSeqToSet(xs, f)
    decreases set _x0: X, _o0: object? | _o0 in f.reads(_x0) :: _o0, xs
  {
    set x: X {:trigger f(x)} {:trigger x in xs} | x in xs :: f(x)
  }

  lemma lemma_SubsetCardinality<X>(xs: set<X>, ys: set<X>, f: X -> bool)
    requires forall x: X :: x in xs ==> f.requires(x)
    requires forall x: X :: x in ys ==> x in xs && f(x)
    ensures |ys| <= |xs|
    decreases xs, ys
  {
  }

  function MakeSubset<X(!new)>(xs: set<X>, f: X -> bool): set<X>
    requires forall x: X :: x in xs ==> f.requires(x)
    reads f.reads
    ensures forall x: X :: x in MakeSubset(xs, f) <==> x in xs && f(x)
    ensures |MakeSubset(xs, f)| <= |xs|
    decreases set _x0: X, _o0: object? | _o0 in f.reads(_x0) :: _o0, xs
  {
    var ys: set<X> := set x: X {:trigger f(x)} {:trigger x in xs} | x in xs && f(x);
    lemma_SubsetCardinality(xs, ys, f);
    ys
  }

  lemma lemma_UnionCardinality<X>(xs: set<X>, ys: set<X>, us: set<X>)
    requires us == xs + ys
    ensures |us| >= |xs|
    decreases ys
  {
  }

  function SetOfNumbersInRightExclusiveRange(a: int, b: int): set<int>
    requires a <= b
    ensures forall opn: int :: a <= opn < b ==> opn in SetOfNumbersInRightExclusiveRange(a, b)
    ensures forall opn: int :: opn in SetOfNumbersInRightExclusiveRange(a, b) ==> a <= opn < b
    ensures |SetOfNumbersInRightExclusiveRange(a, b)| == b - a
    decreases b - a
  {
    if a == b then
      {}
    else
      {a} + SetOfNumbersInRightExclusiveRange(a + 1, b)
  }

  lemma lemma_CardinalityOfBoundedSet(s: set<int>, a: int, b: int)
    requires forall opn: int :: opn in s ==> a <= opn < b
    requires a <= b
    ensures |s| <= b - a
    decreases s, a, b
  {
  }

  function intsetmax(s: set<int>): int
    requires |s| > 0
    ensures var m: int := intsetmax(s); m in s && forall i: int :: i in s ==> m >= i
    decreases s
  {
    var x: int :| x in s;
    if |s| == 1 then
      assert |s - {x}| == 0;
      x
    else
      var sy: set<int> := s - {x}; var y: int := intsetmax(sy); assert forall i: int :: i in s ==> i in sy || i == x; if x > y then x else y
  }
}

module NodeImpl_i {

  import opened Impl_Node_i = Impl_Node_i

  import opened UdpLock_i = UdpLock_i
  class NodeImpl {
    var node: CNode
    var udpClient: UdpClient
    var localAddr: EndPoint
    ghost var Repr: set<object>

    constructor ()
    {
      udpClient := null;
    }

    predicate Valid()
      reads this, UdpClientIsValid.reads(udpClient)
      decreases UdpClientIsValid.reads(udpClient) + {this}
    {
      CNodeValid(node) &&
      UdpClientIsValid(udpClient) &&
      udpClient.LocalEndPoint() == localAddr &&
      udpClient.LocalEndPoint() == node.config[node.my_index] &&
      Repr == {this} + UdpClientRepr(udpClient)
    }

    function Env(): HostEnvironment
      reads this, UdpClientIsValid.reads(udpClient)
      decreases UdpClientIsValid.reads(udpClient) + {this}
    {
      if udpClient != null then
        udpClient.env
      else
        null
    }

    method ConstructUdpClient(me: EndPoint, ghost env_: HostEnvironment)
        returns (ok: bool, client: UdpClient)
      requires env_ != null && env_.Valid() && env_.ok.ok()
      requires EndPointIsValidIPV4(me)
      modifies env_.ok
      ensures ok ==> UdpClientIsValid(client) && client.LocalEndPoint() == me && client.env == env_
      decreases me, env_
    {
      var my_ep := me;
      var ip_byte_array := new byte[|my_ep.addr|];
      seqIntoArrayOpt(my_ep.addr, ip_byte_array);
      var ip_endpoint;
      ok, ip_endpoint := IPEndPoint.Construct(ip_byte_array, my_ep.port, env_);
      if !ok {
        return;
      }
      ok, client := UdpClient.Construct(ip_endpoint, env_);
      if ok {
        calc {
          client.LocalEndPoint();
          ip_endpoint.EP();
          my_ep;
        }
      }
    }

    method InitNode(config: Config, my_index: uint64, ghost env_: HostEnvironment)
        returns (ok: bool)
      requires env_ != null && env_.Valid() && env_.ok.ok()
      requires ValidConfig(config) && ValidConfigIndex(config, my_index)
      modifies this, udpClient, env_.ok
      ensures ok ==> Valid() && Env() == env_ && NodeInit(AbstractifyCNode(node), my_index as int, config) && node.config == config && node.my_index == my_index
      decreases config, my_index, env_
    {
      ok, udpClient := ConstructUdpClient(config[my_index], env_);
      if ok {
        node := NodeInitImpl(my_index, config);
        assert node.my_index == my_index;
        localAddr := node.config[my_index];
        Repr := {this} + UdpClientRepr(udpClient);
      }
    }

    method NodeNextGrant() returns (ok: bool, ghost udpEventLog: seq<UdpEvent>, ghost ios: seq<LockIo>)
      requires Valid()
      modifies Repr
      ensures Repr == old(Repr)
      ensures ok == UdpClientOk(udpClient)
      ensures Env() == old(Env())
      ensures ok ==> Valid() && NodeGrant(old(AbstractifyCNode(node)), AbstractifyCNode(node), ios) && AbstractifyRawLogToIos(udpEventLog) == ios && OnlySentMarshallableData(udpEventLog) && old(Env().udp.history()) + udpEventLog == Env().udp.history()
    {
      var transfer_packet;
      node, transfer_packet, ios := NodeGrantImpl(node);
      ok := true;
      if transfer_packet.Some? {
        ghost var sendEventLog;
        ok, sendEventLog := SendPacket(udpClient, transfer_packet, localAddr);
        udpEventLog := sendEventLog;
      } else {
        udpEventLog := [];
        assert AbstractifyRawLogToIos(udpEventLog) == ios;
      }
    }

    method NodeNextAccept() returns (ok: bool, ghost udpEventLog: seq<UdpEvent>, ghost ios: seq<LockIo>)
      requires Valid()
      modifies Repr
      ensures Repr == old(Repr)
      ensures ok == UdpClientOk(udpClient)
      ensures Env() == old(Env())
      ensures ok ==> Valid() && NodeAccept(old(AbstractifyCNode(node)), AbstractifyCNode(node), ios) && AbstractifyRawLogToIos(udpEventLog) == ios && OnlySentMarshallableData(udpEventLog) && old(Env().udp.history()) + udpEventLog == Env().udp.history()
    {
      var rr;
      ghost var receiveEvent;
      rr, receiveEvent := Receive(udpClient, localAddr);
      udpEventLog := [receiveEvent];
      if rr.RRFail? {
        ok := false;
        return;
      } else if rr.RRTimeout? {
        ok := true;
        ios := [LIoOpTimeoutReceive()];
        return;
      } else {
        ok := true;
        var locked_packet;
        node, locked_packet, ios := NodeAcceptImpl(node, rr.cpacket);
        if locked_packet.Some? {
          ghost var sendEventLog;
          ok, sendEventLog := SendPacket(udpClient, locked_packet, localAddr);
          udpEventLog := udpEventLog + sendEventLog;
        }
      }
    }

    method HostNextMain() returns (ok: bool, ghost udpEventLog: seq<UdpEvent>, ghost ios: seq<LockIo>)
      requires Valid()
      modifies Repr
      ensures Repr == old(Repr)
      ensures ok <==> Env() != null && Env().Valid() && Env().ok.ok()
      ensures Env() == old(Env())
      ensures ok ==> Valid() && NodeNext(old(AbstractifyCNode(node)), AbstractifyCNode(node), ios) && AbstractifyRawLogToIos(udpEventLog) == ios && OnlySentMarshallableData(udpEventLog) && old(Env().udp.history()) + udpEventLog == Env().udp.history()
    {
      if node.held {
        ok, udpEventLog, ios := NodeNextGrant();
      } else {
        ok, udpEventLog, ios := NodeNextAccept();
      }
    }
  }
}

module LockCmdLineParser_i {

  import opened CmdLineParser_i = CmdLineParser_i
  function method EndPointNull(): EndPoint
  {
    EndPoint([0, 0, 0, 0], 0)
  }

  function lock_config_parsing(args: seq<seq<uint16>>): seq<EndPoint>
    decreases args
  {
    if args != [] && |args[1..]| % 2 == 0 then
      var (ok: bool, endpoints: seq<EndPoint>) := parse_end_points(args[1..]);
      if ok && |endpoints| > 0 && |endpoints| < 18446744073709551616 then
        endpoints
      else
        []
    else
      []
  }

  function lock_parse_id(ip: seq<uint16>, port: seq<uint16>): EndPoint
    decreases ip, port
  {
    var (ok: bool, ep: EndPoint) := parse_end_point(ip, port);
    ep
  }

  function lock_cmd_line_parsing(env: HostEnvironment): (seq<EndPoint>, EndPoint)
    requires env != null && env.constants != null
    reads env, if env != null then env.constants else null
    decreases {env, if env != null then env.constants else null}, env
  {
    var args: seq<seq<uint16>> := env.constants.CommandLineArgs();
    if |args| < 2 then
      ([], EndPointNull())
    else
      var penultimate_arg: seq<uint16>, final_arg: seq<uint16> := args[|args| - 2], args[|args| - 1]; var config: seq<EndPoint> := lock_config_parsing(args[..|args| - 2]); var me: EndPoint := lock_parse_id(penultimate_arg, final_arg); (config, me)
  }

  method GetHostIndex(host: EndPoint, hosts: seq<EndPoint>)
      returns (found: bool, index: uint64)
    requires EndPointIsValidIPV4(host)
    requires SeqIsUnique(hosts)
    requires |hosts| < 18446744073709551616
    requires forall h: EndPoint :: h in hosts ==> EndPointIsValidIPV4(h)
    ensures found ==> 0 <= index as int < |hosts| && hosts[index] == host
    ensures !found ==> !(host in hosts)
    decreases host, hosts
  {
    var i: uint64 := 0;
    while i < uint64(|hosts|)
      invariant i as int <= |hosts|
      invariant forall j: uint64 :: 0 <= j < i ==> hosts[j] != host
      decreases uint64(|hosts|) as int - i as int
    {
      if host == hosts[i] {
        found := true;
        index := i;
        calc ==> {
          true;
          {
            reveal_SeqIsUnique();
          }
          forall j: int :: 
            0 <= j < |hosts| &&
            j != i as int ==>
              hosts[j] != host;
        }
        return;
      }
      if i == uint64(|hosts|) - 1 {
        found := false;
        return;
      }
      i := i + 1;
    }
    found := false;
  }

  method ParseCmdLine(ghost env: HostEnvironment)
      returns (ok: bool, host_ids: seq<EndPoint>, my_index: uint64)
    requires HostEnvironmentIsValid(env)
    ensures ok ==> |host_ids| > 0
    ensures ok ==> 0 <= my_index as int < |host_ids|
    ensures var (host_ids': seq<EndPoint>, my_ep': EndPoint) := lock_cmd_line_parsing(env); ok ==> host_ids == host_ids' && host_ids[my_index] == my_ep'
    ensures ok ==> SeqIsUnique(host_ids)
    decreases env
  {
    ok := false;
    var num_args := HostConstants.NumCommandLineArgs(env);
    if num_args < 4 || num_args % 2 != 1 {
      print ""Incorrect number of command line arguments.\n"";
      print ""Expected: ./Main.exe [IP port]+ [IP port]\n"";
      print ""  where the final argument is one of the two IP-port pairs provided earlier \n"";
      return;
    }
    var args := collect_cmd_line_args(env);
    assert args == env.constants.CommandLineArgs();
    var tuple1 := parse_end_points(args[1 .. |args| - 2]);
    ok := tuple1.0;
    var endpoints := tuple1.1;
    if !ok || |endpoints| == 0 || |endpoints| >= 18446744073709551616 {
      ok := false;
      return;
    }
    var tuple2 := parse_end_point(args[|args| - 2], args[|args| - 1]);
    ok := tuple2.0;
    if !ok {
      return;
    }
    var unique := test_unique'(endpoints);
    if !unique {
      ok := false;
      return;
    }
    ok, my_index := GetHostIndex(tuple2.1, endpoints);
    if !ok {
      return;
    }
    host_ids := endpoints;
    var me := endpoints[my_index];
    ghost var ghost_tuple := lock_cmd_line_parsing(env);
    ghost var config', my_ep' := ghost_tuple.0, ghost_tuple.1;
    assert endpoints == config';
    assert me == my_ep';
  }
}

module DistributedSystem_i {

  import opened Protocol_Node_i = Protocol_Node_i

  import opened Common__SeqIsUnique_i = Common__SeqIsUnique_i

  import opened Collections__Seqs_i = Collections__Seqs_i

  import opened Host_i = Host_i
  datatype LS_State = LS_State(environment: LockEnvironment, servers: map<EndPoint, Node>)

  datatype GLS_State = GLS_State(ls: LS_State, history: seq<EndPoint>)

  predicate LS_Init(s: LS_State, config: Config)
    decreases s, config
  {
    LEnvironment_Init(s.environment) &&
    |config| > 0 &&
    SeqIsUnique(config) &&
    (forall e: EndPoint :: 
      e in config <==> e in s.servers) &&
    forall index: int :: 
      0 <= index < |config| ==>
        NodeInit(s.servers[config[index]], index, config)
  }

  predicate LS_NextOneServer(s: LS_State, s': LS_State, id: EndPoint, ios: seq<LockIo>, lstep: LockStep)
    requires id in s.servers
    decreases s, s', id, ios, lstep
  {
    id in s'.servers &&
    NodeNext(s.servers[id], s'.servers[id], ios) &&
    var ns: Node := s.servers[id]; var ns': Node := s'.servers[id]; match lstep { case GrantStep => NodeGrant(ns, ns', ios) case AcceptStep => NodeAccept(ns, ns', ios) } && s'.servers == s.servers[id := s'.servers[id]]
  }

  predicate NodeAcquiresLock(e: EndPoint, s: LS_State, s': LS_State)
    decreases e, s, s'
  {
    e in s.servers &&
    e in s'.servers &&
    !s.servers[e].held &&
    s'.servers[e].held
  }

  predicate LS_Next(s: LS_State, s': LS_State)
    decreases s, s'
  {
    LEnvironment_Next(s.environment, s'.environment) &&
    if s.environment.nextStep.LEnvStepHostIos? && s.environment.nextStep.actor in s.servers then LS_NextOneServer(s, s', s.environment.nextStep.actor, s.environment.nextStep.ios, s.environment.nextStep.nodeStep) else s'.servers == s.servers
  }

  predicate GLS_Init(s: GLS_State, config: Config)
    decreases s, config
  {
    LS_Init(s.ls, config) &&
    s.history == [config[0]]
  }

  predicate GLS_Next(s: GLS_State, s': GLS_State)
    decreases s, s'
  {
    LS_Next(s.ls, s'.ls) &&
    if s.ls.environment.nextStep.LEnvStepHostIos? && s.ls.environment.nextStep.actor in s.ls.servers && NodeGrant(s.ls.servers[s.ls.environment.nextStep.actor], s'.ls.servers[s.ls.environment.nextStep.actor], s.ls.environment.nextStep.ios) && s.ls.servers[s.ls.environment.nextStep.actor].held && s.ls.servers[s.ls.environment.nextStep.actor].epoch < 18446744073709551615 then s'.history == s.history + [s.ls.servers[s.ls.environment.nextStep.actor].config[(s.ls.servers[s.ls.environment.nextStep.actor].my_index + 1) % |s.ls.servers[s.ls.environment.nextStep.actor].config|]] else s'.history == s.history
  }
}

module Collections__Maps_i {
  predicate eq_map<A(!new), B>(x: map<A, B>, y: map<A, B>)
    ensures eq_map(x, y) ==> x == y
    decreases x, y
  {
    (forall a: A :: 
      a in x <==> a in y) &&
    forall a: A :: 
      a in x ==>
        x[a] == y[a]
  }

  function method domain<U(!new), V>(m: map<U, V>): set<U>
    ensures forall i: U :: i in domain(m) <==> i in m
    decreases m
  {
    set s: U {:trigger s in m} | s in m
  }

  function union<U(!new), V>(m: map<U, V>, m': map<U, V>): map<U, V>
    requires m !! m'
    ensures forall i: U :: i in union(m, m') <==> i in m || i in m'
    ensures forall i: U :: i in m ==> union(m, m')[i] == m[i]
    ensures forall i: U :: i in m' ==> union(m, m')[i] == m'[i]
    decreases m, m'
  {
    map i: U {:auto_trigger} {:trigger m'[i]} {:trigger m[i]} {:trigger i in m} | i in domain(m) + domain(m') :: if i in m then m[i] else m'[i]
  }

  function method RemoveElt<U(!new), V>(m: map<U, V>, elt: U): map<U, V>
    requires elt in m
    ensures |RemoveElt(m, elt)| == |m| - 1
    ensures !(elt in RemoveElt(m, elt))
    ensures forall elt': U :: elt' in RemoveElt(m, elt) <==> elt' in m && elt' != elt
    decreases |m|
  {
    var m': map<U, V> := map elt': U {:trigger m[elt']} {:trigger elt' in m} | elt' in m && elt' != elt :: m[elt'];
    lemma_map_remove_one(m, m', elt);
    m'
  }

  lemma lemma_non_empty_map_has_elements<S, T>(m: map<S, T>)
    requires |m| > 0
    ensures exists x: S :: x in m
    decreases m
  {
  }

  lemma lemma_MapSizeIsDomainSize<S, T>(dom: set<S>, m: map<S, T>)
    requires dom == domain(m)
    ensures |m| == |dom|
    decreases dom, m
  {
  }

  lemma lemma_maps_decrease<S, T>(before: map<S, T>, after: map<S, T>, item_removed: S)
    requires item_removed in before
    requires after == map s: S {:trigger before[s]} {:trigger s in before} | s in before && s != item_removed :: before[s]
    ensures |after| < |before|
    decreases before, after
  {
  }

  lemma lemma_map_remove_one<S, T>(before: map<S, T>, after: map<S, T>, item_removed: S)
    requires item_removed in before
    requires after == map s: S {:trigger before[s]} {:trigger s in before} | s in before && s != item_removed :: before[s]
    ensures |after| + 1 == |before|
    decreases before, after
  {
  }
}

module Logic__Option_i {
  datatype Option<T> = Some(v: T) | None
}

module Types_i {

  import opened Environment_s = Environment_s

  import opened Native__Io_s = Native__Io_s
  datatype LockMessage = Transfer(transfer_epoch: int) | Locked(locked_epoch: int) | Invalid

  datatype LockStep = GrantStep | AcceptStep

  type LockEnvironment = LEnvironment<EndPoint, LockMessage, LockStep>

  type LockPacket = LPacket<EndPoint, LockMessage>

  type LockIo = LIoOp<EndPoint, LockMessage>
}

module Native__NativeTypes_s {
  newtype {:nativeType ""sbyte""} sbyte = i: int
    | -128 <= i < 128

  newtype {:nativeType ""byte""} byte = i: int
    | 0 <= i < 256

  newtype {:nativeType ""short""} int16 = i: int
    | -32768 <= i < 32768

  newtype {:nativeType ""ushort""} uint16 = i: int
    | 0 <= i < 65536

  newtype {:nativeType ""int""} int32 = i: int
    | -2147483648 <= i < 2147483648

  newtype {:nativeType ""uint""} uint32 = i: int
    | 0 <= i < 4294967296

  newtype {:nativeType ""long""} int64 = i: int
    | -9223372036854775808 <= i < 9223372036854775808

  newtype {:nativeType ""ulong""} uint64 = i: int
    | 0 <= i < 18446744073709551616

  newtype {:nativeType ""sbyte""} nat8 = i: int
    | 0 <= i < 128

  newtype {:nativeType ""short""} nat16 = i: int
    | 0 <= i < 32768

  newtype {:nativeType ""int""} nat32 = i: int
    | 0 <= i < 2147483648

  newtype {:nativeType ""long""} nat64 = i: int
    | 0 <= i < 9223372036854775808
}

module Collections__Seqs_i {

  import opened Collections__Seqs_s = Collections__Seqs_s
  lemma SeqAdditionIsAssociative<T>(a: seq<T>, b: seq<T>, c: seq<T>)
    ensures a + (b + c) == a + b + c
    decreases a, b, c
  {
  }

  predicate ItemAtPositionInSeq<T>(s: seq<T>, v: T, idx: int)
    decreases s, idx
  {
    0 <= idx < |s| &&
    s[idx] == v
  }

  lemma Lemma_ItemInSeqAtASomePosition<T>(s: seq<T>, v: T)
    requires v in s
    ensures exists idx: int :: ItemAtPositionInSeq(s, v, idx)
    decreases s
  {
  }

  function FindIndexInSeq<T>(s: seq<T>, v: T): int
    ensures var idx: int := FindIndexInSeq(s, v); if idx >= 0 then idx < |s| && s[idx] == v else v !in s
    decreases s
  {
    if v in s then
      Lemma_ItemInSeqAtASomePosition(s, v);
      var idx: int :| ItemAtPositionInSeq(s, v, idx);
      idx
    else
      -1
  }

  function {:opaque} {:fuel 0, 0} MapSeqToSeq<T, U>(s: seq<T>, f: T -> U): (s': seq<U>)
    ensures |s'| == |s|
    ensures forall i: int :: 0 <= i < |s| ==> s'[i] == f(s[i])
    decreases s
  {
    if |s| == 0 then
      []
    else
      [f(s[0])] + MapSeqToSeq(s[1..], f)
  }

  lemma Lemma_IdenticalSingletonSequencesHaveIdenticalElement<T>(x: T, y: T)
    requires [x] == [y]
    ensures x == y
  {
  }

  function SeqCat<T>(seqs: seq<seq<T>>): seq<T>
    decreases seqs
  {
    if |seqs| == 0 then
      []
    else
      seqs[0] + SeqCat(seqs[1..])
  }

  function SeqCatRev<T>(seqs: seq<seq<T>>): seq<T>
    decreases seqs
  {
    if |seqs| == 0 then
      []
    else
      SeqCatRev(all_but_last(seqs)) + last(seqs)
  }

  lemma /*{:_induction A, B}*/ lemma_SeqCat_adds<T>(A: seq<seq<T>>, B: seq<seq<T>>)
    ensures SeqCat(A + B) == SeqCat(A) + SeqCat(B)
    decreases A, B
  {
  }

  lemma /*{:_induction A, B}*/ lemma_SeqCatRev_adds<T>(A: seq<seq<T>>, B: seq<seq<T>>)
    ensures SeqCatRev(A + B) == SeqCatRev(A) + SeqCatRev(B)
    decreases A, B
  {
  }

  lemma /*{:_induction seqs}*/ lemma_SeqCat_equivalent<T>(seqs: seq<seq<T>>)
    ensures SeqCat(seqs) == SeqCatRev(seqs)
    decreases seqs
  {
  }

  lemma {:axiom} {:auto_generated} {:opaque_reveal} {:verify false} {:fuel MapSeqToSeq<int, int>, 1, 2} reveal_MapSeqToSeq()
}

module Common__Util_i {

  import opened Native__NativeTypes_i = Native__NativeTypes_i

  import opened Native__Io_s = Native__Io_s

  import opened Math__power2_i = Math__power2_i
  method seqToArray_slow<A(0)>(s: seq<A>) returns (a: array<A>)
    ensures a != null
    ensures a[..] == s
    decreases s
  {
    var len := |s|;
    a := new A[len];
    var i := 0;
    while i < len
      invariant 0 <= i <= len
      invariant a[..i] == s[..i]
      decreases len - i
    {
      a[i] := s[i];
      i := i + 1;
    }
  }

  method seqIntoArrayOpt<A>(s: seq<A>, a: array<A>)
    requires a != null
    requires |s| == a.Length
    requires |s| < 18446744073709551616
    modifies a
    ensures a[..] == s
    decreases s, a
  {
    var i: uint64 := 0;
    while i < uint64(|s|)
      invariant 0 <= i as int <= a.Length
      invariant a[..] == s[0 .. i] + old(a[i..])
      decreases uint64(|s|) as int - i as int
    {
      a[i] := s[i];
      i := i + 1;
    }
  }

  method seqToArrayOpt<A(0)>(s: seq<A>) returns (a: array<A>)
    requires |s| < 18446744073709551616
    ensures a != null
    ensures a[..] == s
    ensures fresh(a)
    decreases s
  {
    a := new A[uint64(|s|)];
    seqIntoArrayOpt(s, a);
  }

  method seqIntoArrayChar(s: seq<char>, a: array<char>)
    requires a != null
    requires |s| == a.Length
    requires |s| < 18446744073709551616
    modifies a
    ensures a[..] == s
    decreases s, a
  {
    var i: uint64 := 0;
    while i < uint64(|s|)
      invariant 0 <= i as int <= a.Length
      invariant a[..] == s[0 .. i] + old(a[i..])
      decreases uint64(|s|) as int - i as int
    {
      a[i] := s[i];
      i := i + 1;
    }
  }

  method RecordTimingSeq(name: seq<char>, start: uint64, end: uint64)
    requires 0 < |name| < 18446744073709551616
    decreases name, start, end
  {
    var name_array := new char[|name|];
    seqIntoArrayChar(name, name_array);
    var time: uint64;
    if start <= end {
      time := end - start;
    } else {
      time := 18446744073709551615;
    }
    Time.RecordTiming(name_array, time);
  }

  function BEByteSeqToInt(bytes: seq<byte>): int
    decreases |bytes|
  {
    if bytes == [] then
      0
    else
      BEByteSeqToInt(bytes[..|bytes| - 1]) * 256 + bytes[|bytes| - 1] as int
  }

  lemma /*{:_induction bytes}*/ lemma_BEByteSeqToInt_bound(bytes: seq<byte>)
    ensures 0 <= BEByteSeqToInt(bytes)
    ensures BEByteSeqToInt(bytes) < power2(8 * |bytes|)
    decreases bytes
  {
  }

  lemma /*{:_induction bs}*/ lemma_BEByteSeqToUint64_properties(bs: seq<byte>)
    requires |bs| == Uint64Size() as int
    ensures var ret: uint64 := uint64(bs[0]) * 256 * 256 * 256 * 4294967296 + uint64(bs[1]) * 256 * 256 * 4294967296 + uint64(bs[2]) * 256 * 4294967296 + uint64(bs[3]) * 4294967296 + uint64(bs[4]) * 256 * 256 * 256 + uint64(bs[5]) * 256 * 256 + uint64(bs[6]) * 256 + uint64(bs[7]); ret as int == BEByteSeqToInt(bs)
    decreases bs
  {
  }

  function method SeqByteToUint64(bs: seq<byte>): uint64
    requires |bs| == Uint64Size() as int
    ensures 0 <= BEByteSeqToInt(bs) < 18446744073709551616
    ensures SeqByteToUint64(bs) == uint64(BEByteSeqToInt(bs))
    decreases bs
  {
    lemma_2toX();
    lemma_BEByteSeqToUint64_properties(bs);
    uint64(bs[uint64(0)]) * 256 * 256 * 256 * 4294967296 + uint64(bs[uint64(1)]) * 256 * 256 * 4294967296 + uint64(bs[uint64(2)]) * 256 * 4294967296 + uint64(bs[uint64(3)]) * 4294967296 + uint64(bs[uint64(4)]) * 256 * 256 * 256 + uint64(bs[uint64(5)]) * 256 * 256 + uint64(bs[uint64(6)]) * 256 + uint64(bs[uint64(7)])
  }

  function BEUintToSeqByte(v: int, width: int): seq<byte>
    ensures width >= 0 && v >= 0 ==> |BEUintToSeqByte(v, width)| == width
    decreases v, width
  {
    if width > 0 && v >= 0 then
      BEUintToSeqByte(v / 256, width - 1) + [byte(v % 256)]
    else
      []
  }

  lemma /*{:_induction bytes, val, width}*/ lemma_BEUintToSeqByte_invertability(bytes: seq<byte>, val: int, width: nat)
    requires bytes == BEUintToSeqByte(val, width)
    requires 0 <= val < power2(8 * width)
    requires |bytes| == width
    ensures BEByteSeqToInt(bytes) == val
    decreases bytes, val, width
  {
  }

  lemma /*{:_induction bytes, val, width}*/ lemma_BEByteSeqToInt_invertability(bytes: seq<byte>, val: int, width: nat)
    requires BEByteSeqToInt(bytes) == val
    requires 0 <= val < power2(8 * width)
    requires |bytes| == width
    ensures bytes == BEUintToSeqByte(val, width)
    decreases bytes, val, width
  {
  }

  lemma lemma_BEByteSeqToInt_BEUintToSeqByte_invertability()
    ensures forall bytes: seq<byte>, width: nat :: |bytes| == width ==> bytes == BEUintToSeqByte(BEByteSeqToInt(bytes), width)
    ensures forall width: nat, val: int :: 0 <= val < power2(8 * width) ==> val == BEByteSeqToInt(BEUintToSeqByte(val, width))
  {
  }

  function method Uint64ToSeqByte(u: uint64): seq<byte>
    ensures Uint64ToSeqByte(u) == BEUintToSeqByte(u as int, 8)
    decreases u
  {
    var pv: int := 256;
    var bs: seq<byte> := [byte(u / 72057594037927936), byte(u / 281474976710656 % 256), byte(u / 1099511627776 % 256), byte(u / 4294967296 % 256), byte(u / 16777216 % 256), byte(u / 65536 % 256), byte(u / 256 % 256), byte(u % 256)];
    lemma_2toX();
    var u_int: int := u as int;
    calc {
      BEUintToSeqByte(u_int, 8);
      BEUintToSeqByte(u_int / 256, 7) + [byte(u_int % 256)];
      BEUintToSeqByte(u_int / 256 / 256, 6) + [byte(u_int / 256 % 256)] + [byte(u_int % 256)];
      {
        lemma_div_denominator(u_int as int, 256, 256);
      }
      BEUintToSeqByte(u_int / 65536, 6) + [byte(u_int / 256 % 256)] + [byte(u_int % 256)];
      {
        lemma_div_denominator(u_int as int, 65536, 256);
      }
      BEUintToSeqByte(u_int / 16777216, 5) + [byte(u_int / 65536 % 256)] + [byte(u_int / 256 % 256)] + [byte(u_int % 256)];
      {
        lemma_div_denominator(u_int as int, 16777216, 256);
      }
      BEUintToSeqByte(u_int / 4294967296, 4) + [byte(u_int / 16777216 % 256)] + [byte(u_int / 65536 % 256)] + [byte(u_int / 256 % 256)] + [byte(u_int % 256)];
      {
        lemma_div_denominator(u_int as int, 4294967296, 256);
      }
      BEUintToSeqByte(u_int / 1099511627776, 3) + [byte(u_int / 4294967296 % 256)] + [byte(u_int / 16777216 % 256)] + [byte(u_int / 65536 % 256)] + [byte(u_int / 256 % 256)] + [byte(u_int % 256)];
      {
        lemma_div_denominator(u_int as int, 1099511627776, 256);
      }
      BEUintToSeqByte(u_int / 281474976710656, 2) + [byte(u_int / 1099511627776 % 256)] + [byte(u_int / 4294967296 % 256)] + [byte(u_int / 16777216 % 256)] + [byte(u_int / 65536 % 256)] + [byte(u_int / 256 % 256)] + [byte(u_int % 256)];
      {
        lemma_div_denominator(u_int as int, 281474976710656, 256);
      }
      BEUintToSeqByte(u_int / 72057594037927936, 1) + [byte(u_int / 281474976710656 % 256)] + [byte(u_int / 1099511627776 % 256)] + [byte(u_int / 4294967296 % 256)] + [byte(u_int / 16777216 % 256)] + [byte(u_int / 65536 % 256)] + [byte(u_int / 256 % 256)] + [byte(u_int % 256)];
      {
        lemma_div_denominator(u_int as int, 72057594037927936, 256);
      }
      BEUintToSeqByte(u_int / 18446744073709551616, 0) + [byte(u_int / 72057594037927936 % 256)] + [byte(u_int / 281474976710656 % 256)] + [byte(u_int / 1099511627776 % 256)] + [byte(u_int / 4294967296 % 256)] + [byte(u_int / 16777216 % 256)] + [byte(u_int / 65536 % 256)] + [byte(u_int / 256 % 256)] + [byte(u_int % 256)];
    }
    bs
  }

  function method SeqByteToUint16(bs: seq<byte>): uint16
    requires |bs| == Uint16Size() as int
    ensures 0 <= BEByteSeqToInt(bs) < 18446744073709551616
    ensures BEByteSeqToInt(bs) < 65536
    ensures SeqByteToUint16(bs) == uint16(BEByteSeqToInt(bs))
    decreases bs
  {
    lemma_2toX();
    lemma_BEByteSeqToInt_bound(bs);
    uint16(bs[uint64(0)]) * 256 + uint16(bs[uint64(1)])
  }

  function method Uint16ToSeqByte(u: uint16): seq<byte>
    ensures Uint16ToSeqByte(u) == BEUintToSeqByte(u as int, 2)
    decreases u
  {
    var pv: int := 256;
    var s: seq<byte> := [byte(u / 256 % 256), byte(u % 256)];
    lemma_2toX();
    var u_int: int := u as int;
    calc {
      BEUintToSeqByte(u_int, 2);
      BEUintToSeqByte(u_int / 256, 1) + [byte(u_int % 256)];
      BEUintToSeqByte(u_int / 256 / 256, 0) + [byte(u_int / 256 % 256)] + [byte(u_int % 256)];
      {
        lemma_div_denominator(u_int as int, 256, 256);
      }
      BEUintToSeqByte(u_int / 65536, 0) + [byte(u_int / 256 % 256)] + [byte(u_int % 256)];
    }
    s
  }
}

module Common__MarshallInt_i {

  import opened Native__NativeTypes_s = Native__NativeTypes_s

  import opened Common__Util_i = Common__Util_i
  method MarshallUint64_guts(n: uint64, data: array<byte>, index: uint64)
    requires data != null
    requires index as int + Uint64Size() as int <= data.Length
    requires 0 <= index as int + Uint64Size() as int < 18446744073709551616
    requires data.Length < 18446744073709551616
    modifies data
    ensures SeqByteToUint64(data[index .. index + uint64(Uint64Size())]) == n
    ensures data[0 .. index] == old(data[0 .. index])
    ensures data[index + uint64(Uint64Size())..] == old(data[index + uint64(Uint64Size())..])
    decreases n, data, index
  {
    data[index] := byte(n / 72057594037927936);
    data[index + 1] := byte(n / 281474976710656 % 256);
    data[index + 2] := byte(n / 1099511627776 % 256);
    data[index + 3] := byte(n / 4294967296 % 256);
    data[index + 4] := byte(n / 16777216 % 256);
    data[index + 5] := byte(n / 65536 % 256);
    data[index + 6] := byte(n / 256 % 256);
    data[index + 7] := byte(n % 256);
    lemma_2toX();
    assert data[index .. index + uint64(Uint64Size())] == Uint64ToSeqByte(n);
    lemma_BEUintToSeqByte_invertability(data[index .. index + uint64(Uint64Size())], n as int, 8);
  }
}

module Impl_Node_i {

  import opened Protocol_Node_i = Protocol_Node_i

  import opened Message_i = Message_i

  import opened Common__UdpClient_i = Common__UdpClient_i

  import opened Logic__Option_i = Logic__Option_i

  import opened PacketParsing_i = PacketParsing_i

  import opened Common__SeqIsUniqueDef_i = Common__SeqIsUniqueDef_i
  datatype CNode = CNode(held: bool, epoch: uint64, my_index: uint64, config: Config)

  predicate ValidConfig(c: Config)
    decreases c
  {
    0 < |c| < 18446744073709551616 &&
    (forall e: EndPoint :: 
      e in c ==>
        EndPointIsValidIPV4(e)) &&
    SeqIsUnique(c)
  }

  predicate ValidConfigIndex(c: Config, index: uint64)
    decreases c, index
  {
    0 <= index as int < |c|
  }

  predicate CNodeValid(c: CNode)
    decreases c
  {
    ValidConfig(c.config) &&
    ValidConfigIndex(c.config, c.my_index)
  }

  function AbstractifyCNode(n: CNode): Node
    decreases n
  {
    Node(n.held, n.epoch as int, n.my_index as int, n.config)
  }

  method NodeInitImpl(my_index: uint64, config: Config) returns (node: CNode)
    requires 0 < |config| < 18446744073709551616
    requires 0 <= my_index as int < |config|
    requires ValidConfig(config)
    ensures CNodeValid(node)
    ensures NodeInit(AbstractifyCNode(node), my_index as int, config)
    ensures node.my_index == my_index
    ensures node.config == config
    decreases my_index, config
  {
    node := CNode(my_index == 0, if my_index == 0 then 1 else 0, my_index, config);
    if node.held {
      print ""I start holding the lock\n"";
    }
  }

  method NodeGrantImpl(s: CNode)
      returns (s': CNode, packet: Option<CLockPacket>, ghost ios: seq<LockIo>)
    requires CNodeValid(s)
    ensures NodeGrant(AbstractifyCNode(s), AbstractifyCNode(s'), ios)
    ensures s'.my_index == s.my_index && s'.config == s.config
    ensures |ios| == 0 || |ios| == 1
    ensures packet.Some? ==> |ios| == 1 && ios[0].LIoOpSend? && ios[0].s == AbstractifyCLockPacket(packet.v)
    ensures OptionCLockPacketValid(packet) && (packet.Some? ==> packet.v.src == s.config[s.my_index])
    ensures packet.None? ==> ios == [] && s' == s
    ensures CNodeValid(s')
    decreases s
  {
    if s.held && s.epoch < 18446744073709551615 {
      var ssss := CNode(false, s.epoch, s.my_index, s.config);
      s' := ssss;
      var dst_index := (s.my_index + 1) % uint64(|s.config|);
      packet := Some(LPacket(s.config[dst_index], s.config[s.my_index], CTransfer(s.epoch + 1)));
      ios := [LIoOpSend(AbstractifyCLockPacket(packet.v))];
      print ""I grant the lock "", s.epoch, ""\n"";
    } else {
      s' := s;
      ios := [];
      packet := None();
    }
  }

  method NodeAcceptImpl(s: CNode, transfer_packet: CLockPacket)
      returns (s': CNode, locked_packet: Option<CLockPacket>, ghost ios: seq<LockIo>)
    requires CNodeValid(s)
    ensures NodeAccept(AbstractifyCNode(s), AbstractifyCNode(s'), ios)
    ensures s'.my_index == s.my_index && s'.config == s.config
    ensures |ios| == 1 || |ios| == 2
    ensures locked_packet.None? ==> |ios| == 1 && ios[0].LIoOpReceive? && ios[0].r == AbstractifyCLockPacket(transfer_packet)
    ensures locked_packet.Some? ==> |ios| == 2 && ios == [LIoOpReceive(AbstractifyCLockPacket(transfer_packet)), LIoOpSend(AbstractifyCLockPacket(locked_packet.v))]
    ensures OptionCLockPacketValid(locked_packet) && (locked_packet.Some? ==> locked_packet.v.src == s.config[s.my_index])
    ensures CNodeValid(s')
    decreases s, transfer_packet
  {
    ios := [LIoOpReceive(AbstractifyCLockPacket(transfer_packet))];
    if !s.held && transfer_packet.src in s.config && transfer_packet.msg.CTransfer? && transfer_packet.msg.transfer_epoch > s.epoch {
      var ssss := CNode(true, transfer_packet.msg.transfer_epoch, s.my_index, s.config);
      s' := ssss;
      locked_packet := Some(LPacket(transfer_packet.src, s.config[s.my_index], CLocked(transfer_packet.msg.transfer_epoch)));
      ios := ios + [LIoOpSend(AbstractifyCLockPacket(locked_packet.v))];
      print ""I hold the lock!\n"";
    } else {
      s' := s;
      locked_packet := None();
    }
  }
}

module CmdLineParser_i {

  import opened Native__Io_s = Native__Io_s

  import opened Math__power_i = Math__power_i

  import opened Common__SeqIsUniqueDef_i = Common__SeqIsUniqueDef_i

  import opened Common__UdpClient_i = Common__UdpClient_i
  function method ascii_to_int(short: uint16): (bool, byte)
    ensures var tuple: (bool, byte) := ascii_to_int(short); tuple.0 ==> 0 <= tuple.1 <= 9
    decreases short
  {
    if 48 <= short <= 57 then
      (true, byte(short - 48))
    else
      (false, 0)
  }

  method power10(e: nat) returns (val: int)
    ensures val == power(10, e)
    decreases e
  {
    reveal_power();
    if e == 0 {
      return 1;
    } else {
      var tmp := power10(e - 1);
      return 10 * tmp;
    }
  }

  function method shorts_to_bytes(shorts: seq<uint16>): (bool, seq<byte>)
    decreases shorts
  {
    if |shorts| == 0 then
      (true, [])
    else
      var tuple: (bool, seq<byte>) := shorts_to_bytes(shorts[1..]); var ok: bool, rest: seq<byte> := tuple.0, tuple.1; var tuple': (bool, byte) := ascii_to_int(shorts[0]); var ok': bool, a_byte: byte := tuple'.0, tuple'.1; if ok && ok' then (true, [a_byte] + rest) else (false, [])
  }

  function method bytes_to_decimal(bytes: seq<byte>): nat
    decreases bytes
  {
    if |bytes| == 0 then
      0
    else
      bytes[|bytes| - 1] as int + 10 * bytes_to_decimal(bytes[0 .. |bytes| - 1])
  }

  function method shorts_to_nat(shorts: seq<uint16>): (bool, int)
    decreases shorts
  {
    if |shorts| == 0 then
      (false, 0)
    else
      var tuple: (bool, seq<byte>) := shorts_to_bytes(shorts); var ok: bool, bytes: seq<byte> := tuple.0, tuple.1; if !ok then (false, 0) else (true, bytes_to_decimal(bytes))
  }

  function method shorts_to_byte(shorts: seq<uint16>): (bool, byte)
    decreases shorts
  {
    var tuple: (bool, int) := shorts_to_nat(shorts);
    var ok: bool, val: int := tuple.0, tuple.1;
    if 0 <= val < 256 then
      (true, byte(val))
    else
      (false, 0)
  }

  function method shorts_to_uint16(shorts: seq<uint16>): (bool, uint16)
    decreases shorts
  {
    var tuple: (bool, int) := shorts_to_nat(shorts);
    var ok: bool, val: int := tuple.0, tuple.1;
    if 0 <= val < 65536 then
      (true, uint16(val))
    else
      (false, 0)
  }

  function method shorts_to_uint32(shorts: seq<uint16>): (bool, uint32)
    decreases shorts
  {
    var tuple: (bool, int) := shorts_to_nat(shorts);
    var ok: bool, val: int := tuple.0, tuple.1;
    if 0 <= val < 4294967296 then
      (true, uint32(val))
    else
      (false, 0)
  }

  function method is_ascii_period(short: uint16): bool
    decreases short
  {
    short == 46
  }

  function method parse_ip_addr_helper(ip_shorts: seq<uint16>, current_octet_shorts: seq<uint16>): (bool, seq<byte>)
    decreases ip_shorts, current_octet_shorts
  {
    if |ip_shorts| == 0 then
      var tuple: (bool, byte) := shorts_to_byte(current_octet_shorts);
      var okay: bool, b: byte := tuple.0, tuple.1;
      if !okay then
        (false, [])
      else
        (true, [b])
    else if is_ascii_period(ip_shorts[0]) then
      var tuple: (bool, byte) := shorts_to_byte(current_octet_shorts);
      var okay: bool, b: byte := tuple.0, tuple.1;
      if !okay then
        (false, [])
      else
        var tuple': (bool, seq<byte>) := parse_ip_addr_helper(ip_shorts[1..], []); var ok: bool, ip_bytes: seq<byte> := tuple'.0, tuple'.1; if !ok then (false, []) else (true, [b] + ip_bytes)
    else
      parse_ip_addr_helper(ip_shorts[1..], current_octet_shorts + [ip_shorts[0]])
  }

  function method parse_ip_addr(ip_shorts: seq<uint16>): (bool, seq<byte>)
    decreases ip_shorts
  {
    var tuple: (bool, seq<byte>) := parse_ip_addr_helper(ip_shorts, []);
    var ok: bool, ip_bytes: seq<byte> := tuple.0, tuple.1;
    if ok && |ip_bytes| == 4 then
      (true, ip_bytes)
    else
      (false, [])
  }

  function method {:opaque} {:fuel 0, 0} parse_end_point(ip_shorts: seq<uint16>, port_shorts: seq<uint16>): (bool, EndPoint)
    ensures var tuple: (bool, EndPoint) := parse_end_point(ip_shorts, port_shorts); var ok: bool, ep: EndPoint := tuple.0, tuple.1; ok ==> EndPointIsValidIPV4(ep)
    decreases ip_shorts, port_shorts
  {
    var tuple: (bool, seq<byte>) := parse_ip_addr(ip_shorts);
    var okay: bool, ip_bytes: seq<byte> := tuple.0, tuple.1;
    if !okay then
      (false, EndPoint([0, 0, 0, 0], 0))
    else
      var tuple': (bool, uint16) := shorts_to_uint16(port_shorts); var okay': bool, port: uint16 := tuple'.0, tuple'.1; if !okay' then (false, EndPoint([0, 0, 0, 0], 0)) else (true, EndPoint(ip_bytes, port))
  }

  method test_unique'(endpoints: seq<EndPoint>) returns (unique: bool)
    ensures unique <==> SeqIsUnique(endpoints)
    decreases endpoints
  {
    unique := true;
    var i := 0;
    while i < |endpoints|
      invariant 0 <= i <= |endpoints|
      invariant forall j: int, k: int :: 0 <= j < |endpoints| && 0 <= k < i && j != k ==> endpoints[j] != endpoints[k]
      decreases |endpoints| - i
    {
      var j := 0;
      while j < |endpoints|
        invariant 0 <= j <= |endpoints|
        invariant forall k: int :: 0 <= k < j && k != i ==> endpoints[i] != endpoints[k]
        decreases |endpoints| - j
      {
        if i != j && endpoints[i] == endpoints[j] {
          unique := false;
          reveal_SeqIsUnique();
          return;
        }
        j := j + 1;
      }
      i := i + 1;
    }
    reveal_SeqIsUnique();
  }

  function method parse_end_points(args: seq<seq<uint16>>): (bool, seq<EndPoint>)
    requires |args| % 2 == 0
    ensures var (ok: bool, endpoints: seq<EndPoint>) := parse_end_points(args); ok ==> forall e: EndPoint :: e in endpoints ==> EndPointIsValidIPV4(e)
    decreases args
  {
    if |args| == 0 then
      (true, [])
    else
      var (ok1: bool, ep: EndPoint) := parse_end_point(args[0], args[1]); var (ok2: bool, rest: seq<EndPoint>) := parse_end_points(args[2..]); if !(ok1 && ok2) then (false, []) else (true, [ep] + rest)
  }

  method collect_cmd_line_args(ghost env: HostEnvironment) returns (args: seq<seq<uint16>>)
    requires HostEnvironmentIsValid(env)
    ensures |env.constants.CommandLineArgs()| == |args|
    ensures forall i: int :: 0 <= i < |env.constants.CommandLineArgs()| ==> args[i] == env.constants.CommandLineArgs()[i]
    decreases env
  {
    var num_args := HostConstants.NumCommandLineArgs(env);
    var i := 0;
    args := [];
    while i < num_args
      invariant 0 <= i <= num_args
      invariant |env.constants.CommandLineArgs()[0 .. i]| == |args|
      invariant forall j: uint32 :: 0 <= j < i ==> args[j] == env.constants.CommandLineArgs()[j]
      decreases num_args as int - i as int
    {
      var arg := HostConstants.GetCommandLineArg(uint64(i), env);
      args := args + [arg[..]];
      i := i + 1;
    }
  }

  lemma {:axiom} {:auto_generated} {:opaque_reveal} {:verify false} {:fuel parse_end_point, 1, 2} reveal_parse_end_point()
}

module Protocol_Node_i {

  import opened Types_i = Types_i

  import opened Native__Io_s = Native__Io_s
  type Config = seq<EndPoint>

  datatype Node = Node(held: bool, epoch: int, my_index: int, config: Config)

  predicate NodeInit(s: Node, my_index: int, config: Config)
    decreases s, my_index, config
  {
    s.epoch == (if my_index == 0 then 1 else 0) &&
    0 <= my_index < |config| &&
    s.my_index == my_index &&
    s.held == (my_index == 0) &&
    s.config == config
  }

  predicate NodeGrant(s: Node, s': Node, ios: seq<LockIo>)
    decreases s, s', ios
  {
    s.my_index == s'.my_index &&
    if s.held && s.epoch < 18446744073709551615 then !s'.held && |ios| == 1 && ios[0].LIoOpSend? && |s.config| > 0 && s'.config == s.config && s'.epoch == s.epoch && var outbound_packet: LPacket<EndPoint, LockMessage> := ios[0].s; outbound_packet.msg.Transfer? && outbound_packet.msg.transfer_epoch == s.epoch + 1 && outbound_packet.dst == s.config[(s.my_index + 1) % |s.config|] else s == s' && ios == []
  }

  predicate NodeAccept(s: Node, s': Node, ios: seq<LockIo>)
    decreases s, s', ios
  {
    s.my_index == s'.my_index &&
    |ios| >= 1 &&
    if ios[0].LIoOpTimeoutReceive? then s == s' && |ios| == 1 else ios[0].LIoOpReceive? && if !s.held && ios[0].r.src in s.config && ios[0].r.msg.Transfer? && ios[0].r.msg.transfer_epoch > s.epoch then s'.held && |ios| == 2 && ios[1].LIoOpSend? && ios[1].s.msg.Locked? && ios[1].s.dst == ios[0].r.src && s'.epoch == ios[0].r.msg.transfer_epoch == ios[1].s.msg.locked_epoch && s'.config == s.config else s == s' && |ios| == 1
  }

  predicate NodeNext(s: Node, s': Node, ios: seq<LockIo>)
    decreases s, s', ios
  {
    NodeGrant(s, s', ios) || NodeAccept(s, s', ios)
  }
}

module Common__SeqIsUnique_i {

  import opened Common__SeqIsUniqueDef_i = Common__SeqIsUniqueDef_i

  import opened Native__NativeTypes_i = Native__NativeTypes_i
  function UniqueSeqToSet<X>(xs: seq<X>): set<X>
    requires SeqIsUnique(xs)
    ensures forall x: X :: x in xs ==> x in UniqueSeqToSet(xs)
    decreases xs
  {
    set x: X {:trigger x in xs} | x in xs
  }

  function {:timeLimit 90} {:opaque} {:fuel 0, 0} SetToUniqueSeq<X(!new)>(s: set<X>): seq<X>
    ensures forall x: X :: x in SetToUniqueSeq(s) <==> x in s
    ensures SeqIsUnique(SetToUniqueSeq(s))
    ensures |SetToUniqueSeq(s)| == |s|
    decreases s
  {
    if s == {} then
      var xs: seq<X> := [];
      calc ==> {
        true;
        {
          reveal_SeqIsUnique();
        }
        SeqIsUnique(xs);
      }
      xs
    else
      var x: X :| x in s; var s': set<X> := s - {x}; var xs': seq<X> := SetToUniqueSeq(s'); calc ==> {
    true;
    {
      reveal_SeqIsUnique();
    }
    SeqIsUnique(xs' + [x]);
  } xs' + [x]
  }

  function Subsequence<X(!new)>(xs: seq<X>, f: X -> bool): seq<X>
    requires forall x: X :: x in xs ==> f.requires(x)
    reads f.reads
    ensures forall x: X :: x in Subsequence(xs, f) <==> x in xs && f(x)
    decreases set _x0: X, _o0: object? | _o0 in f.reads(_x0) :: _o0, xs
  {
    var s: set<X> := set x: X {:trigger f(x)} {:trigger x in xs} | x in xs && f(x);
    SetToUniqueSeq(s)
  }

  method SeqToSetConstruct<X>(xs: seq<X>) returns (s: set<X>)
    ensures forall x: X :: x in s <==> x in xs
    ensures SeqIsUnique(xs) ==> |s| == |xs| && s == UniqueSeqToSet(xs)
    decreases xs
  {
    reveal_SeqIsUnique();
    s := {};
    var i := 0;
    while i < |xs|
      invariant 0 <= i <= |xs|
      invariant forall x: X :: x in s <==> x in xs[..i]
      invariant SeqIsUnique(xs[..i]) ==> |s| == i
      decreases |xs| - i
    {
      s := s + {xs[i]};
      i := i + 1;
    }
  }

  method {:timeLimit 150} SetToUniqueSeqConstruct<X(0)>(s: set<X>) returns (xs: seq<X>)
    ensures SeqIsUnique(xs)
    ensures UniqueSeqToSet(xs) == s
    ensures forall x: X :: x in xs <==> x in s
    ensures |xs| == |s|
    decreases s
  {
    var arr := new X[|s|];
    var s1 := s;
    ghost var s2 := {};
    ghost var i := 0;
    forall
      ensures SeqIsUnique(arr[..i])
    {
      reveal_SeqIsUnique();
    }
    while |s1| != 0
      invariant 0 <= i <= |s|
      invariant s1 + s2 == s
      invariant s1 !! s2
      invariant |s1| == |s| - i
      invariant |s2| == i
      invariant SeqIsUnique(arr[..i])
      invariant forall x: X :: x in arr[..i] <==> x in s2
      decreases if |s1| <= 0 then 0 - |s1| else |s1| - 0
    {
      reveal_SeqIsUnique();
      ghost var old_seq := arr[..i];
      var x :| x in s1;
      assert x !in old_seq;
      assert forall y: X {:trigger y in s2} {:trigger y in old_seq} :: y in s2 + {x} ==> y in old_seq + [x];
      arr[|s| - |s1|] := x;
      s1 := s1 - {x};
      s2 := s2 + {x};
      i := i + 1;
      assert arr[..i] == old_seq + [x];
    }
    xs := arr[..];
    assert xs == arr[..i];
  }

  method SubsequenceConstruct<X(==,0)>(xs: seq<X>, f: X -> bool) returns (xs': seq<X>)
    requires forall x: X :: x in xs ==> f.requires(x)
    ensures forall x: X {:trigger x in xs} {:trigger x in xs'} :: x in xs' <==> x in xs && f(x)
    ensures SeqIsUnique(xs) ==> SeqIsUnique(xs')
    decreases xs
  {
    reveal_SeqIsUnique();
    var arr := new X[|xs|];
    var i := 0;
    var j := 0;
    while i < |xs|
      invariant 0 <= i <= |xs|
      invariant 0 <= j <= i
      invariant forall x: X {:trigger x in xs[..i]} {:trigger x in arr[..j]} :: x in arr[..j] <==> x in xs[..i] && f(x)
      invariant SeqIsUnique(xs) ==> SeqIsUnique(arr[..j])
      decreases |xs| - i
    {
      ghost var old_xs := xs[..i];
      ghost var old_xs' := arr[..j];
      if f(xs[i]) {
        if SeqIsUnique(xs) {
          reveal_SeqIsUnique();
          assert forall k: int :: 0 <= k < i ==> xs[k] != xs[i];
          assert forall k: int :: 0 <= k < i ==> xs[..i][k] != xs[i];
          assert xs[i] !in arr[..j];
        }
        arr[j] := xs[i];
        j := j + 1;
        assert arr[..j] == old_xs' + [xs[i]];
      }
      i := i + 1;
      assert xs[..i] == old_xs + [xs[i - 1]];
    }
    xs' := arr[..j];
  }

  method UniqueSubsequenceConstruct<X(==,0)>(xs: seq<X>, f: X -> bool) returns (xs': seq<X>)
    requires forall x: X :: x in xs ==> f.requires(x)
    ensures forall x: X {:trigger x in xs} {:trigger x in xs'} :: x in xs' <==> x in xs && f(x)
    ensures SeqIsUnique(xs')
    decreases xs
  {
    var s := set x: X {:trigger f(x)} {:trigger x in xs} | x in xs && f(x);
    xs' := SetToUniqueSeqConstruct(s);
  }

  lemma EstablishAppendToUniqueSeq<X>(xs: seq<X>, x: X, xs': seq<X>)
    requires SeqIsUnique(xs)
    requires x !in xs
    requires xs' == xs + [x]
    ensures SeqIsUnique(xs')
    ensures x in xs'
    decreases xs, xs'
  {
  }

  function method AppendToUniqueSeq<X>(xs: seq<X>, x: X): seq<X>
    requires SeqIsUnique(xs)
    requires x !in xs
    ensures SeqIsUnique(AppendToUniqueSeq(xs, x))
    ensures x in AppendToUniqueSeq(xs, x)
    decreases xs
  {
    reveal_SeqIsUnique();
    var xs': seq<X> := xs + [x];
    EstablishAppendToUniqueSeq(xs, x, xs');
    xs'
  }

  function method AppendToUniqueSeqMaybe<X(==)>(xs: seq<X>, x: X): seq<X>
    requires SeqIsUnique(xs)
    ensures SeqIsUnique(AppendToUniqueSeqMaybe(xs, x))
    ensures x in AppendToUniqueSeqMaybe(xs, x)
    decreases xs
  {
    reveal_SeqIsUnique();
    if x in xs then
      xs
    else
      var xs': seq<X> := xs + [x]; EstablishAppendToUniqueSeq(xs, x, xs'); xs'
  }

  lemma lemma_UniqueSeq_SubSeqsUnique<X>(whole: seq<X>, left: seq<X>, right: seq<X>)
    requires SeqIsUnique(whole)
    requires whole == left + right
    requires |left| > 0
    requires |right| > 0
    requires |whole| > 0
    ensures SeqIsUnique(left)
    ensures SeqIsUnique(right)
    decreases whole, left, right
  {
  }

  lemma lemma_seqs_set_cardinality<T>(Q: seq<T>, S: set<T>)
    requires SeqIsUnique(Q)
    requires S == UniqueSeqToSet(Q)
    ensures |Q| == |S|
    decreases Q, S
  {
  }

  lemma lemma_seqs_set_membership<T>(Q: seq<T>, S: set<T>)
    requires SeqIsUnique(Q)
    requires S == UniqueSeqToSet(Q)
    ensures forall i: T :: i in Q <==> i in S
    decreases Q, S
  {
  }

  lemma {:axiom} {:auto_generated} {:opaque_reveal} {:verify false} {:fuel SetToUniqueSeq<int>, 1, 2} reveal_SetToUniqueSeq()
}

module Native__NativeTypes_i {

  import opened Native__NativeTypes_s = Native__NativeTypes_s
  function method Uint64Size(): uint64
  {
    8
  }

  function method Uint32Size(): uint64
  {
    4
  }

  function method Uint16Size(): uint64
  {
    2
  }
}

module Math__power2_i {

  import opened Math__power2_s = Math__power2_s

  import opened Math__power_s = Math__power_s

  import opened Math__power_i = Math__power_i

  import opened Math__div_i = Math__div_i
  lemma lemma_power2_is_power_2_general()
    ensures forall x: nat :: power2(x) == power(2, x)
  {
  }

  lemma /*{:_induction x}*/ lemma_power2_is_power_2(x: nat)
    ensures power2(x) == power(2, x)
    decreases x
  {
  }

  lemma lemma_power2_auto()
    ensures power2(0) == 1
    ensures power2(1) == 2
    ensures forall x: nat, y: nat {:trigger power2(x + y)} :: power2(x + y) == power2(x) * power2(y)
    ensures forall x: nat, y: nat {:trigger power2(x - y)} :: x >= y ==> power2(x - y) * power2(y) == power2(x)
    ensures forall x: nat, y: nat {:trigger x * y} :: y == 2 ==> x * y == x + x
  {
  }

  lemma /*{:_induction e1, e2}*/ lemma_power2_strictly_increases(e1: int, e2: int)
    requires 0 <= e1 < e2
    ensures power2(e1) < power2(e2)
    decreases e1, e2
  {
  }

  lemma /*{:_induction e1, e2}*/ lemma_power2_increases(e1: int, e2: int)
    requires 0 <= e1 <= e2
    ensures power2(e1) <= power2(e2)
    decreases e1, e2
  {
  }

  lemma /*{:_induction e1, e2}*/ lemma_power2_strictly_increases_converse(e1: int, e2: int)
    requires 0 <= e1
    requires 0 < e2
    requires power2(e1) < power2(e2)
    ensures e1 < e2
    decreases e1, e2
  {
  }

  lemma /*{:_induction e1, e2}*/ lemma_power2_increases_converse(e1: int, e2: int)
    requires 0 < e1
    requires 0 < e2
    requires power2(e1) <= power2(e2)
    ensures e1 <= e2
    decreases e1, e2
  {
  }

  lemma /*{:_induction e1, e2}*/ lemma_power2_adds(e1: nat, e2: nat)
    ensures power2(e1 + e2) == power2(e1) * power2(e2)
    decreases e2
  {
  }

  lemma /*{:_induction x, y}*/ lemma_power2_div_is_sub(x: int, y: int)
    requires 0 <= x <= y
    ensures power2(y - x) == power2(y) / power2(x) >= 0
    decreases x, y
  {
  }

  lemma lemma_2toX32()
    ensures power2(0) == 1
    ensures power2(1) == 2
    ensures power2(2) == 4
    ensures power2(3) == 8
    ensures power2(4) == 16
    ensures power2(5) == 32
    ensures power2(6) == 64
    ensures power2(7) == 128
    ensures power2(8) == 256
    ensures power2(9) == 512
    ensures power2(10) == 1024
    ensures power2(11) == 2048
    ensures power2(12) == 4096
    ensures power2(13) == 8192
    ensures power2(14) == 16384
    ensures power2(15) == 32768
    ensures power2(16) == 65536
    ensures power2(17) == 131072
    ensures power2(18) == 262144
    ensures power2(19) == 524288
    ensures power2(20) == 1048576
    ensures power2(21) == 2097152
    ensures power2(22) == 4194304
    ensures power2(23) == 8388608
    ensures power2(24) == 16777216
    ensures power2(25) == 33554432
    ensures power2(26) == 67108864
    ensures power2(27) == 134217728
    ensures power2(28) == 268435456
    ensures power2(29) == 536870912
    ensures power2(30) == 1073741824
    ensures power2(31) == 2147483648
    ensures power2(32) == 4294967296
  {
  }

  lemma lemma_2toX()
    ensures power2(64) == 18446744073709551616
    ensures power2(60) == 1152921504606846976
    ensures power2(32) == 4294967296
    ensures power2(24) == 16777216
    ensures power2(19) == 524288
    ensures power2(16) == 65536
    ensures power2(8) == 256
  {
  }

  lemma /*{:_induction n}*/ lemma_power2_add8(n: int)
    requires n >= 0
    ensures power2(n + 1) == 2 * power2(n)
    ensures power2(n + 2) == 4 * power2(n)
    ensures power2(n + 3) == 8 * power2(n)
    ensures power2(n + 4) == 16 * power2(n)
    ensures power2(n + 5) == 32 * power2(n)
    ensures power2(n + 6) == 64 * power2(n)
    ensures power2(n + 7) == 128 * power2(n)
    ensures power2(n + 8) == 256 * power2(n)
    decreases n
  {
  }

  lemma lemma_2to32()
    ensures power2(32) == 4294967296
    ensures power2(24) == 16777216
    ensures power2(19) == 524288
    ensures power2(16) == 65536
    ensures power2(8) == 256
    ensures power2(0) == 1
  {
  }

  lemma lemma_2to64()
    ensures power2(64) == 18446744073709551616
    ensures power2(60) == 1152921504606846976
  {
  }

  lemma lemma_power2_0_is_1()
    ensures power2(0) == 1
  {
  }

  lemma lemma_power2_1_is_2()
    ensures power2(1) == 2
  {
  }

  lemma /*{:_induction a, b}*/ lemma_bit_count_is_unique(x: int, a: int, b: int)
    requires 0 < a
    requires 0 < b
    requires power2(a - 1) <= x < power2(a)
    requires power2(b - 1) <= x < power2(b)
    ensures a == b
    decreases x, a, b
  {
  }

  lemma /*{:_induction x, y, z}*/ lemma_pull_out_powers_of_2(x: nat, y: nat, z: nat)
    ensures 0 <= x * y
    ensures 0 <= y * z
    ensures power(power2(x * y), z) == power(power2(x), y * z)
    decreases x, y, z
  {
  }

  lemma lemma_rebase_powers_of_2()
    ensures forall n: nat, e: nat {:trigger power(power2(n), e)} :: 0 <= n * e && power(power2(n), e) == power2(n * e)
  {
  }

  lemma /*{:_induction c}*/ lemma_mask_div_2(c: nat)
    requires 0 < c
    ensures (power2(c) - 1) / 2 == power2(c - 1) - 1
    decreases c
  {
  }

  lemma /*{:_induction p, s}*/ lemma_power2_division_inequality(x: nat, p: nat, s: nat)
    requires s <= p
    requires x < power2(p)
    ensures x / power2(s) < power2(p - s)
    decreases x, p, s
  {
  }

  lemma /*{:_induction a, b}*/ lemma_power2_unfolding(a: nat, b: nat)
    ensures 0 <= a * b
    ensures power(power2(a), b) == power2(a * b)
    decreases a, b
  {
  }

  function {:opaque} {:fuel 0, 0} NatNumBits(n: nat): nat
    ensures NatNumBits(n) >= 0
    decreases n
  {
    if n == 0 then
      0
    else
      1 + NatNumBits(n / 2)
  }

  lemma /*{:_induction c, n}*/ lemma_Power2BoundIsNatNumBits(c: nat, n: nat)
    ensures (c > 0 ==> power2(c - 1) <= n) && n < power2(c) <==> c == NatNumBits(n)
    decreases c, n
  {
  }

  lemma {:axiom} {:auto_generated} {:opaque_reveal} {:verify false} {:fuel NatNumBits, 1, 2} reveal_NatNumBits()
}

module Common__SeqIsUniqueDef_i {
  predicate {:opaque} {:fuel 0, 0} SeqIsUnique<X>(xs: seq<X>)
    decreases xs
  {
    forall i: int, j: int :: 
      0 <= i < |xs| &&
      0 <= j < |xs| &&
      xs[i] == xs[j] ==>
        i == j
  }

  lemma {:axiom} {:auto_generated} {:opaque_reveal} {:verify false} {:fuel SeqIsUnique<int>, 1, 2} reveal_SeqIsUnique()
}

module Math__power_i {

  import opened Math__power_s = Math__power_s

  import opened Math__mul_i = Math__mul_i
  lemma /*{:_induction b}*/ lemma_power_0(b: int)
    ensures power(b, 0) == 1
    decreases b
  {
  }

  lemma /*{:_induction b}*/ lemma_power_1(b: int)
    ensures power(b, 1) == b
    decreases b
  {
  }

  lemma /*{:_induction e}*/ lemma_0_power(e: nat)
    requires e > 0
    ensures power(0, e) == 0
    decreases e
  {
  }

  lemma /*{:_induction e}*/ lemma_1_power(e: nat)
    ensures power(1, e) == 1
    decreases e
  {
  }

  lemma /*{:_induction b, e1, e2}*/ lemma_power_adds(b: int, e1: nat, e2: nat)
    ensures power(b, e1) * power(b, e2) == power(b, e1 + e2)
    decreases e1
  {
  }

  lemma /*{:_induction a, b, c}*/ lemma_power_multiplies(a: int, b: nat, c: nat)
    ensures 0 <= b * c
    ensures power(a, b * c) == power(power(a, b), c)
    decreases c
  {
  }

  lemma /*{:_induction a, b, e}*/ lemma_power_distributes(a: int, b: int, e: nat)
    ensures power(a * b, e) == power(a, e) * power(b, e)
    decreases e
  {
  }

  lemma lemma_power_auto()
    ensures forall x: int {:trigger power(x, 0)} :: power(x, 0) == 1
    ensures forall x: int {:trigger power(x, 1)} :: power(x, 1) == x
    ensures forall x: int, y: int {:trigger power(x, y)} :: y == 0 ==> power(x, y) == 1
    ensures forall x: int, y: int {:trigger power(x, y)} :: y == 1 ==> power(x, y) == x
    ensures forall x: int, y: int {:trigger x * y} :: 0 < x && 0 < y ==> x <= x * y
    ensures forall x: int, y: int {:trigger x * y} :: 0 < x && 1 < y ==> x < x * y
    ensures forall x: int, y: nat, z: nat {:trigger power(x, y + z)} :: power(x, y + z) == power(x, y) * power(x, z)
    ensures forall x: int, y: nat, z: nat {:trigger power(x, y - z)} :: y >= z ==> power(x, y - z) * power(x, z) == power(x, y)
    ensures forall x: int, y: int, z: nat {:trigger power(x * y, z)} :: power(x * y, z) == power(x, z) * power(y, z)
  {
  }

  lemma /*{:_induction b, e}*/ lemma_power_positive(b: int, e: nat)
    requires 0 < b
    ensures (forall i: int :: i in imap u: nat {:trigger power(b, u)} | true :: 0 <= u ==> 0 < power(b, u)) ==> 0 < power(b, e)
    decreases b, e
  {
  }

  lemma /*{:_induction b, e1, e2}*/ lemma_power_increases(b: nat, e1: nat, e2: nat)
    requires 0 < b
    requires e1 <= e2
    ensures power(b, e1) <= power(b, e2)
    decreases b, e1, e2
  {
  }

  lemma /*{:_induction b, e1, e2}*/ lemma_power_strictly_increases(b: nat, e1: nat, e2: nat)
    requires 1 < b
    requires e1 < e2
    ensures power(b, e1) < power(b, e2)
    decreases b, e1, e2
  {
  }

  lemma /*{:_induction x}*/ lemma_square_is_power_2(x: nat)
    ensures power(x, 2) == x * x
    decreases x
  {
  }
}

module Math__power2_s {

  import opened Libraries__base_s = Libraries__base_s
  function {:opaque} {:fuel 0, 0} power2(exp: nat): nat
    ensures power2(exp) > 0
    decreases exp
  {
    if exp == 0 then
      1
    else
      2 * power2(exp - 1)
  }

  lemma lemma_power2_32()
    ensures power2(8) == 256
    ensures power2(16) == 65536
    ensures power2(24) == 16777216
    ensures power2(32) == 4294967296
  {
  }

  lemma {:axiom} {:auto_generated} {:opaque_reveal} {:verify false} {:fuel power2, 1, 2} reveal_power2()
}

module Math__power_s {
  function {:opaque} {:fuel 0, 0} power(b: int, e: nat): int
    decreases e
  {
    if e == 0 then
      1
    else
      b * power(b, e - 1)
  }

  lemma {:axiom} {:auto_generated} {:opaque_reveal} {:verify false} {:fuel power, 1, 2} reveal_power()
}

module Math__div_i {

  import opened Math__power_i = Math__power_i

  import opened Math__mul_i = Math__mul_i

  import opened Math__div_def_i = Math__div_def_i

  import opened Math__div_boogie_i = Math__div_boogie_i

  import opened Math__div_nonlinear_i = Math__div_nonlinear_i

  import opened Math__div_auto_i = Math__div_auto_i
  lemma lemma_div_by_one_is_identity(x: int)
    decreases x
  {
  }

  lemma lemma_div_basics(x: int)
    ensures x != 0 ==> 0 / x == 0
    ensures x / 1 == x
    ensures x != 0 ==> x / x == 1
    decreases x
  {
  }

  lemma lemma_small_div_converse()
    ensures forall x: int, d: int {:trigger x / d} :: 0 <= x && 0 < d && x / d == 0 ==> x < d
  {
  }

  lemma lemma_div_is_ordered_by_denominator(x: int, y: int, z: int)
    requires x >= 0
    requires 1 <= y <= z
    ensures x / y >= x / z
    decreases x
  {
  }

  lemma lemma_div_is_strictly_ordered_by_denominator(x: int, d: int)
    requires 0 < x
    requires 1 < d
    ensures x / d < x
    decreases x
  {
  }

  lemma lemma_dividing_sums(a: int, b: int, d: int, R: int)
    requires 0 < d
    requires R == a % d + b % d - (a + b) % d
    ensures d * (a + b) / d - R == d * a / d + d * b / d
    decreases a, b, d, R
  {
  }

  lemma lemma_div_pos_is_pos(x: int, divisor: int)
    requires 0 <= x
    requires 0 < divisor
    ensures x / divisor >= 0
    decreases x, divisor
  {
  }

  lemma lemma_div_basics_forall()
    ensures forall x: int {:trigger 0 / x} :: x != 0 ==> 0 / x == 0
    ensures forall x: int {:trigger x / 1} :: x / 1 == x
    ensures forall x: int, y: int {:trigger x / y} :: x >= 0 && y > 0 ==> x / y >= 0
    ensures forall x: int, y: int {:trigger x / y} :: x >= 0 && y > 0 ==> x / y <= x
  {
  }

  lemma lemma_div_neg_neg(x: int, d: int)
    requires d > 0
    ensures x / d == -((-x + d - 1) / d)
    decreases x, d
  {
  }

  lemma lemma_mod_2(x: int)
    decreases x
  {
  }

  lemma lemma_mod2_plus(x: int)
    decreases x
  {
  }

  lemma lemma_mod2_plus2(x: int)
    decreases x
  {
  }

  lemma lemma_mod32(x: int)
    decreases x
  {
  }

  lemma lemma_mod_remainder_neg_specific(x: int, m: int)
    decreases x, m
  {
  }

  lemma lemma_mod_remainder_neg()
  {
  }

  lemma lemma_mod_remainder_pos_specific(x: int, m: int)
    decreases x, m
  {
  }

  lemma lemma_mod_remainder_pos()
  {
  }

  lemma lemma_mod_remainder_specific(x: int, m: int)
    decreases x, m
  {
  }

  lemma lemma_mod_remainder()
  {
  }

  lemma lemma_mod_basics()
    ensures forall m: int {:trigger m % m} :: m > 0 ==> m % m == 0
    ensures forall x: int, m: int {:trigger x % m % m} :: m > 0 ==> x % m % m == x % m
  {
  }

  lemma lemma_mod_properties()
    ensures forall m: int {:trigger m % m} :: m > 0 ==> m % m == 0
    ensures forall x: int, m: int {:trigger x % m % m} :: m > 0 ==> x % m % m == x % m
    ensures forall x: int, m: int {:trigger x % m} :: m > 0 ==> 0 <= x % m < m
  {
  }

  lemma lemma_mod_decreases(x: nat, d: nat)
    requires 0 < d
    ensures x % d <= x
    decreases x, d
  {
  }

  lemma lemma_mod_add_multiples_vanish(b: int, m: int)
    requires 0 < m
    ensures (m + b) % m == b % m
    decreases b, m
  {
  }

  lemma lemma_mod_sub_multiples_vanish(b: int, m: int)
    requires 0 < m
    ensures (-m + b) % m == b % m
    decreases b, m
  {
  }

  lemma lemma_mod_multiples_vanish(a: int, b: int, m: int)
    requires 0 < m
    ensures (m * a + b) % m == b % m
    decreases if a > 0 then a else -a
  {
  }

  lemma lemma_add_mod_noop(x: int, y: int, m: int)
    requires 0 < m
    ensures (x % m + y % m) % m == (x + y) % m
    decreases x, y, m
  {
  }

  lemma lemma_add_mod_noop_right(x: int, y: int, m: int)
    requires 0 < m
    ensures (x + y % m) % m == (x + y) % m
    decreases x, y, m
  {
  }

  lemma lemma_mod_equivalence(x: int, y: int, m: int)
    requires 0 < m
    ensures x % m == y % m <==> (x - y) % m == 0
    decreases x, y, m
  {
  }

  lemma lemma_sub_mod_noop(x: int, y: int, m: int)
    requires 0 < m
    ensures (x % m - y % m) % m == (x - y) % m
    decreases x, y, m
  {
  }

  lemma lemma_sub_mod_noop_right(x: int, y: int, m: int)
    requires 0 < m
    ensures (x - y % m) % m == (x - y) % m
    decreases x, y, m
  {
  }

  lemma lemma_mod_adds(a: int, b: int, d: int)
    requires 0 < d
    ensures a % d + b % d == (a + b) % d + d * (a % d + b % d) / d
    ensures a % d + b % d < d ==> a % d + b % d == (a + b) % d
    decreases a, b, d
  {
  }

  lemma {:timeLimit 60} lemma_mod_neg_neg(x: int, d: int)
    requires d > 0
    ensures x % d == x * (1 - d) % d
    decreases x, d
  {
  }

  lemma lemma_fundamental_div_mod_converse(x: int, d: int, q: int, r: int)
    requires d != 0
    requires 0 <= r < d
    requires x == q * d + r
    ensures q == x / d
    ensures r == x % d
    decreases x, d, q, r
  {
  }

  lemma lemma_mod_pos_bound(x: int, m: int)
    requires 0 <= x
    requires 0 < m
    ensures 0 <= x % m < m
    decreases x
  {
  }

  lemma lemma_mul_mod_noop_left(x: int, y: int, m: int)
    requires 0 < m
    ensures x % m * y % m == x * y % m
    decreases x, y, m
  {
  }

  lemma lemma_mul_mod_noop_right(x: int, y: int, m: int)
    requires 0 < m
    ensures x * y % m % m == x * y % m
    decreases x, y, m
  {
  }

  lemma lemma_mul_mod_noop_general(x: int, y: int, m: int)
    requires 0 < m
    ensures x % m * y % m == x * y % m
    ensures x * y % m % m == x * y % m
    ensures x % m * y % m % m == x * y % m
    decreases x, y, m
  {
  }

  lemma lemma_mul_mod_noop(x: int, y: int, m: int)
    requires 0 < m
    ensures x % m * y % m % m == x * y % m
    decreases x, y, m
  {
  }

  lemma /*{:_induction b, e, m}*/ lemma_power_mod_noop(b: int, e: nat, m: int)
    requires 0 < m
    ensures power(b % m, e) % m == power(b, e) % m
    decreases e
  {
  }

  lemma lemma_mod_subtraction(x: nat, s: nat, d: nat)
    requires 0 < d
    requires 0 <= s <= x % d
    ensures x % d - s % d == (x - s) % d
    decreases x, s, d
  {
  }

  lemma lemma_mod_ordering(x: nat, k: nat, d: nat)
    requires 1 < d
    requires 0 < k
    ensures 0 < d * k
    ensures x % d <= x % (d * k)
    decreases x, k, d
  {
  }

  lemma lemma_mod_multiples_basic(x: int, m: int)
    requires m > 0
    ensures x * m % m == 0
    decreases x, m
  {
  }

  lemma lemma_div_plus_one(x: int, d: int)
    requires d > 0
    ensures 1 + x / d == (d + x) / d
    decreases x, d
  {
  }

  lemma lemma_div_minus_one(x: int, d: int)
    requires d > 0
    ensures -1 + x / d == (-d + x) / d
    decreases x, d
  {
  }

  lemma lemma_mod_mod(x: int, a: int, b: int)
    requires 0 < a
    requires 0 < b
    ensures 0 < a * b
    ensures x % (a * b) % a == x % a
    decreases x, a, b
  {
  }

  lemma lemma_div_is_div_recursive(x: int, d: int)
    requires d > 0
    ensures my_div_recursive(x, d) == x / d
    decreases x, d
  {
  }

  lemma lemma_div_is_div_recursive_forall()
    ensures forall x: int, d: int :: d > 0 ==> my_div_recursive(x, d) == x / d
  {
  }

  lemma /*{:_induction x, m}*/ lemma_mod_is_mod_recursive(x: int, m: int)
    requires m > 0
    ensures my_mod_recursive(x, m) == x % m
    decreases if x < 0 then -x + m else x
  {
  }

  lemma lemma_mod_is_mod_recursive_forall()
    ensures forall x: int, d: int :: d > 0 ==> my_mod_recursive(x, d) == x % d
  {
  }

  lemma lemma_basic_div(d: int)
    requires d > 0
    ensures forall x: int {:trigger x / d} :: 0 <= x < d ==> x / d == 0
    decreases d
  {
  }

  lemma lemma_div_is_ordered(x: int, y: int, z: int)
    requires x <= y
    requires z > 0
    ensures x / z <= y / z
    decreases x, y, z
  {
  }

  lemma lemma_div_decreases(x: int, d: int)
    requires 0 < x
    requires 1 < d
    ensures x / d < x
    decreases x, d
  {
  }

  lemma lemma_div_nonincreasing(x: int, d: int)
    requires 0 <= x
    requires 0 < d
    ensures x / d <= x
    decreases x, d
  {
  }

  lemma lemma_breakdown(a: int, b: int, c: int)
    requires 0 <= a
    requires 0 < b
    requires 0 < c
    ensures 0 < b * c
    ensures a % (b * c) == b * a / b % c + a % b
    decreases a, b, c
  {
  }

  lemma lemma_remainder_upper(x: int, divisor: int)
    requires 0 <= x
    requires 0 < divisor
    ensures x - divisor < x / divisor * divisor
    decreases x, divisor
  {
  }

  lemma lemma_remainder_lower(x: int, divisor: int)
    requires 0 <= x
    requires 0 < divisor
    ensures x >= x / divisor * divisor
    decreases x, divisor
  {
  }

  lemma lemma_remainder(x: int, divisor: int)
    requires 0 <= x
    requires 0 < divisor
    ensures 0 <= x - x / divisor * divisor < divisor
    decreases x, divisor
  {
  }

  lemma lemma_div_denominator(x: int, c: nat, d: nat)
    requires 0 <= x
    requires 0 < c
    requires 0 < d
    ensures c * d != 0
    ensures x / c / d == x / (c * d)
    decreases x, c, d
  {
  }

  lemma lemma_mul_hoist_inequality(x: int, y: int, z: int)
    requires 0 <= x
    requires 0 < z
    ensures x * y / z <= x * y / z
    decreases x, y, z
  {
  }

  lemma lemma_indistinguishable_quotients(a: int, b: int, d: int)
    requires 0 < d
    requires 0 <= a - a % d <= b < a + d - a % d
    ensures a / d == b / d
    decreases a, b, d
  {
  }

  lemma lemma_truncate_middle(x: int, b: int, c: int)
    requires 0 <= x
    requires 0 < b
    requires 0 < c
    ensures 0 < b * c
    ensures b * x % (b * c) == b * x % c
    decreases x, b, c
  {
  }

  lemma lemma_div_multiples_vanish_quotient(x: int, a: int, d: int)
    requires 0 < x
    requires 0 <= a
    requires 0 < d
    ensures 0 < x * d
    ensures a / d == x * a / (x * d)
    decreases x, a, d
  {
  }

  lemma lemma_round_down(a: int, r: int, d: int)
    requires 0 < d
    requires a % d == 0
    requires 0 <= r < d
    ensures a == d * (a + r) / d
    decreases a, r, d
  {
  }

  lemma lemma_div_multiples_vanish_fancy(x: int, b: int, d: int)
    requires 0 < d
    requires 0 <= b < d
    ensures (d * x + b) / d == x
    decreases x, b, d
  {
  }

  lemma lemma_div_multiples_vanish(x: int, d: int)
    requires 0 < d
    ensures d * x / d == x
    decreases x, d
  {
  }

  lemma lemma_div_by_multiple(b: int, d: int)
    requires 0 <= b
    requires 0 < d
    ensures b * d / d == b
    decreases b, d
  {
  }

  lemma lemma_div_by_multiple_is_strongly_ordered(x: int, y: int, m: int, z: int)
    requires x < y
    requires y == m * z
    requires z > 0
    ensures x / z < y / z
    decreases x, y, m, z
  {
  }

  lemma lemma_multiply_divide_le(a: int, b: int, c: int)
    requires 0 < b
    requires a <= b * c
    ensures a / b <= c
    decreases a, b, c
  {
  }

  lemma lemma_multiply_divide_lt(a: int, b: int, c: int)
    requires 0 < b
    requires a < b * c
    ensures a / b < c
    decreases a, b, c
  {
  }

  lemma lemma_hoist_over_denominator(x: int, j: int, d: nat)
    requires 0 < d
    ensures x / d + j == (x + j * d) / d
    decreases x, j, d
  {
  }

  lemma lemma_part_bound1(a: int, b: int, c: int)
    requires 0 <= a
    requires 0 < b
    requires 0 < c
    ensures 0 < b * c
    ensures b * a / b % (b * c) <= b * (c - 1)
    decreases a, b, c
  {
  }

  lemma lemma_part_bound2(a: int, b: int, c: int)
    requires 0 <= a
    requires 0 < b
    requires 0 < c
    ensures 0 < b * c
    ensures a % b % (b * c) < b
    decreases a, b, c
  {
  }

  lemma lemma_mod_breakdown(a: int, b: int, c: int)
    requires 0 <= a
    requires 0 < b
    requires 0 < c
    ensures 0 < b * c
    ensures a % (b * c) == b * a / b % c + a % b
    decreases a, b, c
  {
  }

  lemma lemma_div_denominator_forall()
    ensures forall c: nat, d: nat {:trigger c * d} :: 0 < c && 0 < d ==> c * d != 0
    ensures forall x: int, c: nat, d: nat {:trigger x / c / d} :: 0 <= x && 0 < c && 0 < d ==> x / c / d == x / (c * d)
  {
  }
}

module Math__mul_i {

  import opened Math__mul_nonlinear_i = Math__mul_nonlinear_i

  import opened Math__mul_auto_i = Math__mul_auto_i
  function mul(x: int, y: int): int
    decreases x, y
  {
    x * y
  }

  function mul_recursive(x: int, y: int): int
    decreases x, y
  {
    if x >= 0 then
      mul_pos(x, y)
    else
      -1 * mul_pos(-1 * x, y)
  }

  function {:opaque} {:fuel 0, 0} mul_pos(x: int, y: int): int
    requires x >= 0
    decreases x, y
  {
    if x == 0 then
      0
    else
      y + mul_pos(x - 1, y)
  }

  lemma lemma_mul_is_mul_recursive(x: int, y: int)
    ensures x * y == mul_recursive(x, y)
    decreases x, y
  {
  }

  lemma /*{:_induction x, y}*/ lemma_mul_is_mul_pos(x: int, y: int)
    requires x >= 0
    ensures x * y == mul_pos(x, y)
    decreases x, y
  {
  }

  lemma lemma_mul_basics(x: int)
    ensures 0 * x == 0
    ensures x * 0 == 0
    ensures 1 * x == x
    ensures x * 1 == x
    decreases x
  {
  }

  lemma lemma_mul_is_commutative(x: int, y: int)
    ensures x * y == y * x
    decreases x, y
  {
  }

  lemma lemma_mul_ordering_general()
    ensures forall x: int, y: int {:trigger x * y} :: 0 < x && 0 < y && 0 <= x * y ==> x <= x * y && y <= x * y
  {
  }

  lemma lemma_mul_is_mul_boogie(x: int, y: int)
    decreases x, y
  {
  }

  lemma lemma_mul_inequality(x: int, y: int, z: int)
    requires x <= y
    requires z >= 0
    ensures x * z <= y * z
    decreases x, y, z
  {
  }

  lemma lemma_mul_upper_bound(x: int, x_bound: int, y: int, y_bound: int)
    requires x <= x_bound
    requires y <= y_bound
    requires 0 <= x
    requires 0 <= y
    ensures x * y <= x_bound * y_bound
    decreases x, x_bound, y, y_bound
  {
  }

  lemma lemma_mul_strict_upper_bound(x: int, x_bound: int, y: int, y_bound: int)
    requires x < x_bound
    requires y < y_bound
    requires 0 <= x
    requires 0 <= y
    ensures x * y < x_bound * y_bound
    decreases x, x_bound, y, y_bound
  {
  }

  lemma lemma_mul_left_inequality(x: int, y: int, z: int)
    requires x > 0
    ensures y <= z ==> x * y <= x * z
    ensures y < z ==> x * y < x * z
    decreases x, y, z
  {
  }

  lemma lemma_mul_strict_inequality_converse(x: int, y: int, z: int)
    requires x * z < y * z
    requires z >= 0
    ensures x < y
    decreases x, y, z
  {
  }

  lemma lemma_mul_inequality_converse(x: int, y: int, z: int)
    requires x * z <= y * z
    requires z > 0
    ensures x <= y
    decreases x, y, z
  {
  }

  lemma lemma_mul_equality_converse(x: int, y: int, z: int)
    requires x * z == y * z
    requires 0 < z
    ensures x == y
    decreases x, y, z
  {
  }

  lemma lemma_mul_is_distributive_add_other_way(x: int, y: int, z: int)
    ensures (y + z) * x == y * x + z * x
    decreases x, y, z
  {
  }

  lemma lemma_mul_is_distributive_sub(x: int, y: int, z: int)
    ensures x * (y - z) == x * y - x * z
    decreases x, y, z
  {
  }

  lemma lemma_mul_is_distributive(x: int, y: int, z: int)
    ensures x * (y + z) == x * y + x * z
    ensures x * (y - z) == x * y - x * z
    ensures (y + z) * x == y * x + z * x
    ensures (y - z) * x == y * x - z * x
    ensures x * (y + z) == (y + z) * x
    ensures x * (y - z) == (y - z) * x
    ensures x * y == y * x
    ensures x * z == z * x
    decreases x, y, z
  {
  }

  lemma lemma_mul_strictly_increases(x: int, y: int)
    requires 1 < x
    requires 0 < y
    ensures y < x * y
    decreases x, y
  {
  }

  lemma lemma_mul_increases(x: int, y: int)
    requires 0 < x
    requires 0 < y
    ensures y <= x * y
    decreases x, y
  {
  }

  lemma lemma_mul_nonnegative(x: int, y: int)
    requires 0 <= x
    requires 0 <= y
    ensures 0 <= x * y
    decreases x, y
  {
  }

  lemma lemma_mul_unary_negation(x: int, y: int)
    ensures -x * y == -(x * y) == x * -y
    decreases x, y
  {
  }

  lemma lemma_mul_one_to_one_pos(m: int, x: int, y: int)
    requires 0 < m
    requires m * x == m * y
    ensures x == y
    decreases m, x, y
  {
  }

  lemma lemma_mul_one_to_one(m: int, x: int, y: int)
    requires m != 0
    requires m * x == m * y
    ensures x == y
    decreases m, x, y
  {
  }

  lemma lemma_mul_is_mul_recursive_forall()
    ensures forall x: int, y: int :: x * y == mul_recursive(x, y)
  {
  }

  lemma lemma_mul_basics_forall()
    ensures forall x: int {:trigger 0 * x} :: 0 * x == 0
    ensures forall x: int {:trigger x * 0} :: x * 0 == 0
    ensures forall x: int {:trigger 1 * x} :: 1 * x == x
    ensures forall x: int {:trigger x * 1} :: x * 1 == x
  {
  }

  lemma lemma_mul_is_commutative_forall()
    ensures forall x: int, y: int {:trigger x * y} :: x * y == y * x
  {
  }

  lemma lemma_mul_ordering_forall()
    ensures forall x: int, y: int {:trigger x * y} :: 0 < x && 0 < y && 0 <= x * y ==> x <= x * y && y <= x * y
  {
  }

  lemma lemma_mul_strict_inequality_forall()
    ensures forall x: int, y: int, z: int {:trigger x * z, y * z} :: x < y && z > 0 ==> x * z < y * z
  {
  }

  lemma lemma_mul_inequality_forall()
    ensures forall x: int, y: int, z: int {:trigger x * z, y * z} :: x <= y && z >= 0 ==> x * z <= y * z
  {
  }

  lemma lemma_mul_strict_inequality_converse_forall()
    ensures forall x: int, y: int, z: int {:trigger x * z, y * z} :: x * z < y * z && z >= 0 ==> x < y
  {
  }

  lemma lemma_mul_inequality_converse_forall()
    ensures forall x: int, y: int, z: int {:trigger x * z, y * z} :: x * z <= y * z && z > 0 ==> x <= y
  {
  }

  lemma lemma_mul_is_distributive_add_forall()
    ensures forall x: int, y: int, z: int {:trigger x * (y + z)} :: x * (y + z) == x * y + x * z
  {
  }

  lemma lemma_mul_is_distributive_sub_forall()
    ensures forall x: int, y: int, z: int {:trigger x * (y - z)} :: x * (y - z) == x * y - x * z
  {
  }

  lemma lemma_mul_is_distributive_forall()
    ensures forall x: int, y: int, z: int {:trigger x * (y + z)} :: x * (y + z) == x * y + x * z
    ensures forall x: int, y: int, z: int {:trigger x * (y - z)} :: x * (y - z) == x * y - x * z
    ensures forall x: int, y: int, z: int {:trigger (y + z) * x} :: (y + z) * x == y * x + z * x
    ensures forall x: int, y: int, z: int {:trigger (y - z) * x} :: (y - z) * x == y * x - z * x
  {
  }

  lemma lemma_mul_is_associative_forall()
    ensures forall x: int, y: int, z: int {:trigger x * y * z} {:trigger x * y * z} :: x * y * z == x * y * z
  {
  }

  lemma lemma_mul_nonzero_forall()
    ensures forall x: int, y: int {:trigger x * y} :: x * y != 0 <==> x != 0 && y != 0
  {
  }

  lemma lemma_mul_nonnegative_forall()
    ensures forall x: int, y: int {:trigger x * y} :: 0 <= x && 0 <= y ==> 0 <= x * y
  {
  }

  lemma lemma_mul_unary_negation_forall()
    ensures forall x: int, y: int {:trigger -x * y} {:trigger x * -y} :: -x * y == -(x * y) == x * -y
  {
  }

  lemma lemma_mul_strictly_increases_forall()
    ensures forall x: int, y: int {:trigger x * y} :: 1 < x && 0 < y ==> y < x * y
  {
  }

  lemma lemma_mul_increases_forall()
    ensures forall x: int, y: int {:trigger x * y} :: 0 < x && 0 < y ==> y <= x * y
  {
  }

  lemma lemma_mul_strictly_positive_forall()
    ensures forall x: int, y: int {:trigger x * y} :: 0 < x && 0 < y ==> 0 < x * y
  {
  }

  lemma lemma_mul_one_to_one_forall()
    ensures forall m: int, x: int, y: int {:trigger m * x, m * y} :: m != 0 && m * x == m * y ==> x == y
  {
  }

  lemma lemma_mul_properties()
    ensures forall x: int, y: int {:trigger x * y} :: x * y == y * x
    ensures forall x: int {:trigger x * 0} {:trigger 0 * x} :: x * 0 == 0 * x == 0
    ensures forall x: int {:trigger x * 1} {:trigger 1 * x} :: x * 1 == 1 * x == x
    ensures forall x: int, y: int, z: int {:trigger x * z, y * z} :: x < y && z > 0 ==> x * z < y * z
    ensures forall x: int, y: int, z: int {:trigger x * z, y * z} :: x <= y && z >= 0 ==> x * z <= y * z
    ensures forall x: int, y: int, z: int {:trigger x * (y + z)} :: x * (y + z) == x * y + x * z
    ensures forall x: int, y: int, z: int {:trigger x * (y - z)} :: x * (y - z) == x * y - x * z
    ensures forall x: int, y: int, z: int {:trigger (y + z) * x} :: (y + z) * x == y * x + z * x
    ensures forall x: int, y: int, z: int {:trigger (y - z) * x} :: (y - z) * x == y * x - z * x
    ensures forall x: int, y: int, z: int {:trigger x * y * z} {:trigger x * y * z} :: x * y * z == x * y * z
    ensures forall x: int, y: int {:trigger x * y} :: x * y != 0 <==> x != 0 && y != 0
    ensures forall x: int, y: int {:trigger x * y} :: 0 <= x && 0 <= y ==> 0 <= x * y
    ensures forall x: int, y: int {:trigger x * y} :: 0 < x && 0 < y && 0 <= x * y ==> x <= x * y && y <= x * y
    ensures forall x: int, y: int {:trigger x * y} :: 1 < x && 0 < y ==> y < x * y
    ensures forall x: int, y: int {:trigger x * y} :: 0 < x && 0 < y ==> y <= x * y
    ensures forall x: int, y: int {:trigger x * y} :: 0 < x && 0 < y ==> 0 < x * y
  {
  }

  lemma lemma_mul_cancels_negatives(a: int, b: int)
    ensures a * b == -a * -b
    decreases a, b
  {
  }

  function INTERNAL_mul_recursive(x: int, y: int): int
    decreases x, y
  {
    mul_recursive(x, y)
  }

  lemma {:axiom} {:auto_generated} {:opaque_reveal} {:verify false} {:fuel mul_pos, 1, 2} reveal_mul_pos()
}

module Libraries__base_s {
  function {:imported} unroll(i: int): bool
    decreases i
  {
    true
  }

  function Trigger(i: int): bool
    decreases i
  {
    true
  }

  function sizeof<A>(a: A): int
  {
    1
  }
}

module Math__div_def_i {
  function div(x: int, d: int): int
    requires d != 0
    decreases x, d
  {
    x / d
  }

  function mod(x: int, d: int): int
    requires d != 0
    decreases x, d
  {
    x % d
  }

  function div_recursive(x: int, d: int): int
    requires d != 0
    decreases x, d
  {
    INTERNAL_div_recursive(x, d)
  }

  function mod_recursive(x: int, d: int): int
    requires d > 0
    decreases x, d
  {
    INTERNAL_mod_recursive(x, d)
  }

  function mod_boogie(x: int, y: int): int
    requires y != 0
    decreases x, y
  {
    x % y
  }

  function div_boogie(x: int, y: int): int
    requires y != 0
    decreases x, y
  {
    x / y
  }

  function my_div_recursive(x: int, d: int): int
    requires d != 0
    decreases x, d
  {
    if d > 0 then
      my_div_pos(x, d)
    else
      -1 * my_div_pos(x, -1 * d)
  }

  function my_div_pos(x: int, d: int): int
    requires d > 0
    decreases if x < 0 then d - x else x
  {
    if x < 0 then
      -1 + my_div_pos(x + d, d)
    else if x < d then
      0
    else
      1 + my_div_pos(x - d, d)
  }

  function my_mod_recursive(x: int, m: int): int
    requires m > 0
    decreases if x < 0 then m - x else x
  {
    if x < 0 then
      my_mod_recursive(m + x, m)
    else if x < m then
      x
    else
      my_mod_recursive(x - m, m)
  }

  function INTERNAL_mod_recursive(x: int, m: int): int
    requires m > 0
    decreases x, m
  {
    my_mod_recursive(x, m)
  }

  function INTERNAL_div_recursive(x: int, d: int): int
    requires d != 0
    decreases x, d
  {
    my_div_recursive(x, d)
  }
}

module Math__div_boogie_i {

  import opened Math__div_def_i = Math__div_def_i

  import opened Math__mul_i = Math__mul_i
  lemma lemma_div_is_div_boogie(x: int, d: int)
    requires d != 0
    decreases x, d
  {
  }

  lemma lemma_mod_is_mod_boogie(x: int, d: int)
    requires d > 0
    decreases x, d
  {
  }
}

module Math__div_nonlinear_i {
  lemma lemma_div_of_0(d: int)
    requires d != 0
    ensures 0 / d == 0
    decreases d
  {
  }

  lemma lemma_div_by_self(d: int)
    requires d != 0
    ensures d / d == 1
    decreases d
  {
  }

  lemma lemma_small_div()
    ensures forall d: int, x: int {:trigger x / d} :: 0 <= x < d && d > 0 ==> x / d == 0
  {
  }

  lemma lemma_mod_of_zero_is_zero(m: int)
    requires 0 < m
    ensures 0 % m == 0
    decreases m
  {
  }

  lemma lemma_fundamental_div_mod(x: int, d: int)
    requires d != 0
    ensures x == d * x / d + x % d
    decreases x, d
  {
  }

  lemma lemma_0_mod_anything()
    ensures forall m: int {:trigger 0 % m} :: m > 0 ==> 0 % m == 0
  {
  }

  lemma lemma_small_mod(x: nat, m: nat)
    requires x < m
    requires 0 < m
    ensures x % m == x
    decreases x, m
  {
  }

  lemma lemma_mod_range(x: int, m: int)
    requires m > 0
    ensures 0 <= x % m < m
    decreases x, m
  {
  }

  lemma lemma_real_div_gt(x: real, y: real)
    requires x > y
    requires x >= 0.0
    requires y > 0.0
    ensures x / y > 1 as real
    decreases x, y
  {
  }
}

module Math__div_auto_i {

  import opened Math__mod_auto_i = Math__mod_auto_i

  import opened Math__div_auto_proofs_i = Math__div_auto_proofs_i
  predicate DivAuto(n: int)
    requires n > 0
    decreases n
  {
    ModAuto(n) &&
    n / n == -(-n / n) == 1 &&
    (forall x: int {:trigger x / n} :: 
      0 <= x < n <==> x / n == 0) &&
    (forall x: int, y: int {:trigger (x + y) / n} :: 
      var z: int := x % n + y % n; (0 <= z < n && (x + y) / n == x / n + y / n) || (n <= z < n + n && (x + y) / n == x / n + y / n + 1)) &&
    forall x: int, y: int {:trigger (x - y) / n} :: 
      var z: int := x % n - y % n; (0 <= z < n && (x - y) / n == x / n - y / n) || (-n <= z < 0 && (x - y) / n == x / n - y / n - 1)
  }

  lemma lemma_div_auto(n: int)
    requires n > 0
    ensures DivAuto(n)
    decreases n
  {
  }

  predicate TDivAutoLe(x: int, y: int)
    decreases x, y
  {
    x <= y
  }

  lemma lemma_div_auto_induction(n: int, x: int, f: imap<int, bool>)
    requires n > 0
    requires forall i: int :: i in f
    requires DivAuto(n) ==> (forall i: int {:trigger TDivAutoLe(0, i)} :: TDivAutoLe(0, i) && i < n ==> f[i]) && (forall i: int {:trigger TDivAutoLe(0, i)} :: TDivAutoLe(0, i) && f[i] ==> f[i + n]) && forall i: int {:trigger TDivAutoLe(i + 1, n)} :: TDivAutoLe(i + 1, n) && f[i] ==> f[i - n]
    ensures DivAuto(n)
    ensures f[x]
    decreases n, x
  {
  }

  lemma lemma_div_auto_induction_forall(n: int, f: imap<int, bool>)
    requires n > 0
    requires forall i: int :: i in f
    requires DivAuto(n) ==> (forall i: int {:trigger TDivAutoLe(0, i)} :: TDivAutoLe(0, i) && i < n ==> f[i]) && (forall i: int {:trigger TDivAutoLe(0, i)} :: TDivAutoLe(0, i) && f[i] ==> f[i + n]) && forall i: int {:trigger TDivAutoLe(i + 1, n)} :: TDivAutoLe(i + 1, n) && f[i] ==> f[i - n]
    ensures DivAuto(n)
    ensures forall i: int {:trigger f[i]} :: f[i]
    decreases n
  {
  }
}

module Math__mul_nonlinear_i {
  lemma lemma_mul_strictly_positive(x: int, y: int)
    ensures 0 < x && 0 < y ==> 0 < x * y
    decreases x, y
  {
  }

  lemma lemma_mul_nonzero(x: int, y: int)
    ensures x * y != 0 <==> x != 0 && y != 0
    decreases x, y
  {
  }

  lemma lemma_mul_is_associative(x: int, y: int, z: int)
    ensures x * y * z == x * y * z
    decreases x, y, z
  {
  }

  lemma lemma_mul_is_distributive_add(x: int, y: int, z: int)
    ensures x * (y + z) == x * y + x * z
    decreases x, y, z
  {
  }

  lemma lemma_mul_ordering(x: int, y: int)
    requires 0 < x
    requires 0 < y
    requires 0 <= x * y
    ensures x <= x * y && y <= x * y
    decreases x, y
  {
  }

  lemma lemma_mul_strict_inequality(x: int, y: int, z: int)
    requires x < y
    requires z > 0
    ensures x * z < y * z
    decreases x, y, z
  {
  }
}

module Math__mul_auto_i {

  import opened Math__mul_auto_proofs_i = Math__mul_auto_proofs_i
  predicate MulAuto()
  {
    (forall x: int, y: int {:trigger x * y} :: 
      x * y == y * x) &&
    (forall x: int, y: int, z: int {:trigger (x + y) * z} :: 
      (x + y) * z == x * z + y * z) &&
    forall x: int, y: int, z: int {:trigger (x - y) * z} :: 
      (x - y) * z == x * z - y * z
  }

  lemma lemma_mul_auto()
    ensures MulAuto()
  {
  }

  predicate TMulAutoLe(x: int, y: int)
    decreases x, y
  {
    x <= y
  }

  lemma lemma_mul_auto_induction(x: int, f: imap<int, bool>)
    requires forall i: int :: i in f
    requires MulAuto() ==> f[0] && (forall i: int {:trigger TMulAutoLe(0, i)} :: TMulAutoLe(0, i) && f[i] ==> f[i + 1]) && forall i: int {:trigger TMulAutoLe(i, 0)} :: TMulAutoLe(i, 0) && f[i] ==> f[i - 1]
    ensures MulAuto()
    ensures f[x]
    decreases x
  {
  }

  lemma lemma_mul_auto_induction_forall(f: imap<int, bool>)
    requires forall i: int :: i in f
    requires MulAuto() ==> f[0] && (forall i: int {:trigger TMulAutoLe(0, i)} :: TMulAutoLe(0, i) && f[i] ==> f[i + 1]) && forall i: int {:trigger TMulAutoLe(i, 0)} :: TMulAutoLe(i, 0) && f[i] ==> f[i - 1]
    ensures MulAuto()
    ensures forall i: int {:trigger f[i]} :: f[i]
  {
  }
}

module Math__mod_auto_i {

  import opened Math__mod_auto_proofs_i = Math__mod_auto_proofs_i
  predicate eq_mod(x: int, y: int, n: int)
    requires n > 0
    decreases x, y, n
  {
    (x - y) % n == 0
  }

  predicate ModAuto(n: int)
    requires n > 0
    decreases n
  {
    n % n == -n % n == 0 &&
    (forall x: int {:trigger x % n % n} :: 
      x % n % n == x % n) &&
    (forall x: int {:trigger x % n} :: 
      0 <= x < n <==> x % n == x) &&
    (forall x: int, y: int {:trigger (x + y) % n} :: 
      var z: int := x % n + y % n; (0 <= z < n && (x + y) % n == z) || (n <= z < n + n && (x + y) % n == z - n)) &&
    forall x: int, y: int {:trigger (x - y) % n} :: 
      var z: int := x % n - y % n; (0 <= z < n && (x - y) % n == z) || (-n <= z < 0 && (x - y) % n == z + n)
  }

  lemma lemma_mod_auto(n: int)
    requires n > 0
    ensures ModAuto(n)
    decreases n
  {
  }

  predicate TModAutoLe(x: int, y: int)
    decreases x, y
  {
    x <= y
  }

  lemma lemma_mod_auto_induction(n: int, x: int, f: imap<int, bool>)
    requires n > 0
    requires forall i: int :: i in f
    requires ModAuto(n) ==> (forall i: int {:trigger TModAutoLe(0, i)} :: TModAutoLe(0, i) && i < n ==> f[i]) && (forall i: int {:trigger TModAutoLe(0, i)} :: TModAutoLe(0, i) && f[i] ==> f[i + n]) && forall i: int {:trigger TModAutoLe(i + 1, n)} :: TModAutoLe(i + 1, n) && f[i] ==> f[i - n]
    ensures ModAuto(n)
    ensures f[x]
    decreases n, x
  {
  }

  lemma lemma_mod_auto_induction_forall(n: int, f: imap<int, bool>)
    requires n > 0
    requires forall i: int :: i in f
    requires ModAuto(n) ==> (forall i: int {:trigger TModAutoLe(0, i)} :: TModAutoLe(0, i) && i < n ==> f[i]) && (forall i: int {:trigger TModAutoLe(0, i)} :: TModAutoLe(0, i) && f[i] ==> f[i + n]) && forall i: int {:trigger TModAutoLe(i + 1, n)} :: TModAutoLe(i + 1, n) && f[i] ==> f[i - n]
    ensures ModAuto(n)
    ensures forall i: int {:trigger f[i]} :: f[i]
    decreases n
  {
  }
}

module Math__div_auto_proofs_i {

  import opened Math__mod_auto_i = Math__mod_auto_i
  lemma lemma_div_auto_basics(n: int)
    requires n > 0
    ensures n / n == -(-n / n) == 1
    ensures forall x: int {:trigger x / n} :: 0 <= x < n <==> x / n == 0
    ensures forall x: int {:trigger (x + n) / n} :: (x + n) / n == x / n + 1
    ensures forall x: int {:trigger (x - n) / n} :: (x - n) / n == x / n - 1
    decreases n
  {
  }
}

module Math__mul_auto_proofs_i {

  import opened Math__mul_nonlinear_i = Math__mul_nonlinear_i
  lemma lemma_mul_induction_helper(f: imap<int, bool>, x: int)
    requires forall i: int :: i in f
    requires f[0]
    requires forall i: int {:trigger f[i], f[i + 1]} :: i >= 0 && f[i] ==> f[i + 1]
    requires forall i: int {:trigger f[i], f[i - 1]} :: i <= 0 && f[i] ==> f[i - 1]
    ensures f[x]
    decreases if x >= 0 then x else -x
  {
  }

  lemma lemma_mul_induction_forall(f: imap<int, bool>)
    requires forall i: int :: i in f
    requires f[0]
    requires forall i: int {:trigger f[i], f[i + 1]} :: i >= 0 && f[i] ==> f[i + 1]
    requires forall i: int {:trigger f[i], f[i - 1]} :: i <= 0 && f[i] ==> f[i - 1]
    ensures forall i: int :: f[i]
  {
  }

  lemma lemma_mul_auto_commutes()
    ensures forall x: int, y: int {:trigger x * y} :: x * y == y * x
  {
  }

  lemma lemma_mul_auto_succ()
    ensures forall x: int, y: int {:trigger (x + 1) * y} :: (x + 1) * y == x * y + y
    ensures forall x: int, y: int {:trigger (x - 1) * y} :: (x - 1) * y == x * y - y
  {
  }

  lemma lemma_mul_auto_distributes()
    ensures forall x: int, y: int, z: int {:trigger (x + y) * z} :: (x + y) * z == x * z + y * z
    ensures forall x: int, y: int, z: int {:trigger (x - y) * z} :: (x - y) * z == x * z - y * z
  {
  }
}

module Math__mod_auto_proofs_i {

  import opened Math__mul_auto_i = Math__mul_auto_i

  import opened Math__mul_i = Math__mul_i

  import opened Math__div_nonlinear_i = Math__div_nonlinear_i
  lemma lemma_mod_induction_helper(n: int, f: imap<int, bool>, x: int)
    requires n > 0
    requires forall i: int :: i in f
    requires forall i: int :: 0 <= i < n ==> f[i]
    requires forall i: int {:trigger f[i], f[i + n]} :: i >= 0 && f[i] ==> f[i + n]
    requires forall i: int {:trigger f[i], f[i - n]} :: i < n && f[i] ==> f[i - n]
    ensures f[x]
    decreases if x >= n then x else -x
  {
  }

  lemma lemma_mod_induction_forall(n: int, f: imap<int, bool>)
    requires n > 0
    requires forall i: int :: i in f
    requires forall i: int :: 0 <= i < n ==> f[i]
    requires forall i: int {:trigger f[i], f[i + n]} :: i >= 0 && f[i] ==> f[i + n]
    requires forall i: int {:trigger f[i], f[i - n]} :: i < n && f[i] ==> f[i - n]
    ensures forall i: int :: f[i]
    decreases n
  {
  }

  lemma lemma_mod_induction_forall2(n: int, f: imap<(int, int), bool>)
    requires n > 0
    requires forall i: int, j: int :: (i, j) in f
    requires forall i: int, j: int :: 0 <= i < n && 0 <= j < n ==> f[(i, j)]
    requires forall i: int, j: int {:trigger f[(i, j)], f[(i + n, j)]} :: i >= 0 && f[(i, j)] ==> f[(i + n, j)]
    requires forall i: int, j: int {:trigger f[(i, j)], f[(i, j + n)]} :: j >= 0 && f[(i, j)] ==> f[(i, j + n)]
    requires forall i: int, j: int {:trigger f[(i, j)], f[(i - n, j)]} :: i < n && f[(i, j)] ==> f[(i - n, j)]
    requires forall i: int, j: int {:trigger f[(i, j)], f[(i, j - n)]} :: j < n && f[(i, j)] ==> f[(i, j - n)]
    ensures forall i: int, j: int :: f[(i, j)]
    decreases n
  {
  }

  lemma lemma_mod_auto_basics(n: int)
    requires n > 0
    ensures forall x: int {:trigger (x + n) % n} :: (x + n) % n == x % n
    ensures forall x: int {:trigger (x - n) % n} :: (x - n) % n == x % n
    ensures forall x: int {:trigger (x + n) / n} :: (x + n) / n == x / n + 1
    ensures forall x: int {:trigger (x - n) / n} :: (x - n) / n == x / n - 1
    ensures forall x: int {:trigger x % n} :: 0 <= x < n <==> x % n == x
    decreases n
  {
  }
}
")]

#if ISDAFNYRUNTIMELIB
using System; // for Func
using System.Numerics;
#endif

namespace DafnyAssembly {
  [AttributeUsage(AttributeTargets.Assembly)]
  public class DafnySourceAttribute : Attribute {
    public readonly string dafnySourceText;
    public DafnySourceAttribute(string txt) { dafnySourceText = txt; }
  }
}

namespace Dafny
{
  using System.Collections.Generic;
  // set this option if you want to use System.Collections.Immutable and if you know what you're doing.
#if DAFNY_USE_SYSTEM_COLLECTIONS_IMMUTABLE
  using System.Collections.Immutable;
  using System.Linq;
#endif

  public class Set<T>
  {
#if DAFNY_USE_SYSTEM_COLLECTIONS_IMMUTABLE
    readonly ImmutableHashSet<T> setImpl;
    readonly bool containsNull;
    Set(ImmutableHashSet<T> d, bool containsNull) {
      this.setImpl = d;
      this.containsNull = containsNull;
    }
    public static readonly Set<T> Empty = new Set<T>(ImmutableHashSet<T>.Empty, false);
    public static Set<T> FromElements(params T[] values) {
      return FromCollection(values);
    }
    public static Set<T> FromCollection(IEnumerable<T> values) {
      var d = ImmutableHashSet<T>.Empty.ToBuilder();
      var containsNull = false;
      foreach (T t in values) {
        if (t == null) {
          containsNull = true;
        } else {
          d.Add(t);
        }
      }
      return new Set<T>(d.ToImmutable(), containsNull);
    }
    public static Set<T> FromCollectionPlusOne(IEnumerable<T> values, T oneMoreValue) {
      var d = ImmutableHashSet<T>.Empty.ToBuilder();
      var containsNull = false;
      if (oneMoreValue == null) {
        containsNull = true;
      } else {
        d.Add(oneMoreValue);
      }
      foreach (T t in values) {
        if (t == null) {
          containsNull = true;
        } else {
          d.Add(t);
        }
      }
      return new Set<T>(d.ToImmutable(), containsNull);
    }
    public int Count {
      get { return this.setImpl.Count + (containsNull ? 1 : 0); }
    }
    public long LongCount {
      get { return this.setImpl.Count + (containsNull ? 1 : 0); }
    }
    public IEnumerable<T> Elements {
      get {
        if (containsNull) {
          yield return default(T);
        }
        foreach (var t in this.setImpl) {
          yield return t;
        }
      }
    }
#else
    readonly HashSet<T> setImpl;
    Set(HashSet<T> s) {
      this.setImpl = s;
    }
    public static readonly Set<T> Empty = new Set<T>(new HashSet<T>());
    public static Set<T> FromElements(params T[] values) {
      return FromCollection(values);
    }
    public static Set<T> FromCollection(IEnumerable<T> values) {
      var s = new HashSet<T>(values);
      return new Set<T>(s);
    }
    public static Set<T> FromCollectionPlusOne(IEnumerable<T> values, T oneMoreValue) {
      var s = new HashSet<T>(values);
      s.Add(oneMoreValue);
      return new Set<T>(s);
    }
    public int Count {
      get { return this.setImpl.Count; }
    }
    public long LongCount {
      get { return this.setImpl.Count; }
    }
    public IEnumerable<T> Elements {
      get {
        return this.setImpl;
      }
    }
#endif

    public static Set<T> _DafnyDefaultValue() {
      return Empty;
    }

    /// <summary>
    /// This is an inefficient iterator for producing all subsets of "this".
    /// </summary>
    public IEnumerable<Set<T>> AllSubsets {
      get {
        // Start by putting all set elements into a list, but don't include null
        var elmts = new List<T>();
        elmts.AddRange(this.setImpl);
        var n = elmts.Count;
        var which = new bool[n];
#if DAFNY_USE_SYSTEM_COLLECTIONS_IMMUTABLE
        var s = ImmutableHashSet<T>.Empty.ToBuilder();
#else
        var s = new HashSet<T>();
#endif
        while (true) {
#if DAFNY_USE_SYSTEM_COLLECTIONS_IMMUTABLE
          // yield both the subset without null and, if null is in the original set, the subset with null included
          var ihs = s.ToImmutable();
          yield return new Set<T>(ihs, false);
          if (containsNull) {
            yield return new Set<T>(ihs, true);
          }
#else
          yield return new Set<T>(new HashSet<T>(s));
#endif
          // "add 1" to "which", as if doing a carry chain.  For every digit changed, change the membership of the corresponding element in "s".
          int i = 0;
          for (; i < n && which[i]; i++) {
            which[i] = false;
            s.Remove(elmts[i]);
          }
          if (i == n) {
            // we have cycled through all the subsets
            break;
          }
          which[i] = true;
          s.Add(elmts[i]);
        }
      }
    }
    public bool Equals(Set<T> other) {
#if DAFNY_USE_SYSTEM_COLLECTIONS_IMMUTABLE
      return containsNull == other.containsNull && this.setImpl.SetEquals(other.setImpl);
#else
      return this.setImpl.Count == other.setImpl.Count && IsSubsetOf(other);
#endif
    }
    public override bool Equals(object other) {
      return other is Set<T> && Equals((Set<T>)other);
    }
    public override int GetHashCode() {
      var hashCode = 1;
#if DAFNY_USE_SYSTEM_COLLECTIONS_IMMUTABLE
      if (containsNull) {
        hashCode = hashCode * (Dafny.Helpers.GetHashCode(default(T)) + 3);
      }
#endif
      foreach (var t in this.setImpl) {
        hashCode = hashCode * (Dafny.Helpers.GetHashCode(t)+3);
      }
      return hashCode;
    }
    public override string ToString() {
      var s = "{";
      var sep = "";
#if DAFNY_USE_SYSTEM_COLLECTIONS_IMMUTABLE
      if (containsNull) {
        s += sep + Dafny.Helpers.ToString(default(T));
        sep = ", ";
      }
#endif
      foreach (var t in this.setImpl) {
        s += sep + Dafny.Helpers.ToString(t);
        sep = ", ";
      }
      return s + "}";
    }
    public bool IsProperSubsetOf(Set<T> other) {
      return this.Count < other.Count && IsSubsetOf(other);
    }
    public bool IsSubsetOf(Set<T> other) {
#if DAFNY_USE_SYSTEM_COLLECTIONS_IMMUTABLE
      if (this.containsNull && !other.containsNull) {
        return false;
      }
#endif
      if (other.setImpl.Count < this.setImpl.Count)
        return false;
      foreach (T t in this.setImpl) {
        if (!other.setImpl.Contains(t))
          return false;
      }
      return true;
    }
    public bool IsSupersetOf(Set<T> other) {
      return other.IsSubsetOf(this);
    }
    public bool IsProperSupersetOf(Set<T> other) {
      return other.IsProperSubsetOf(this);
    }
    public bool IsDisjointFrom(Set<T> other) {
#if DAFNY_USE_SYSTEM_COLLECTIONS_IMMUTABLE
      if (this.containsNull && other.containsNull) {
        return false;
      }
      ImmutableHashSet<T> a, b;
#else
      HashSet<T> a, b;
#endif
      if (this.setImpl.Count < other.setImpl.Count) {
        a = this.setImpl; b = other.setImpl;
      } else {
        a = other.setImpl; b = this.setImpl;
      }
      foreach (T t in a) {
        if (b.Contains(t))
          return false;
      }
      return true;
    }
    public bool Contains<G>(G t) {
#if DAFNY_USE_SYSTEM_COLLECTIONS_IMMUTABLE
      return t == null ? containsNull : t is T && this.setImpl.Contains((T)(object)t);
#else
      return (t == null || t is T) && this.setImpl.Contains((T)(object)t);
#endif
    }
#if DAFNY_USE_SYSTEM_COLLECTIONS_IMMUTABLE
    public Set<T> Union(Set<T> other) {
      return new Set<T>(this.setImpl.Union(other.setImpl), containsNull || other.containsNull);
    }
    public Set<T> Intersect(Set<T> other) {
      return new Set<T>(this.setImpl.Intersect(other.setImpl), containsNull && other.containsNull);
    }
    public Set<T> Difference(Set<T> other) {
        return new Set<T>(this.setImpl.Except(other.setImpl), containsNull && !other.containsNull);
    }
#else
    public Set<T> Union(Set<T> other) {
      if (this.setImpl.Count == 0)
        return other;
      else if (other.setImpl.Count == 0)
        return this;
      HashSet<T> a, b;
      if (this.setImpl.Count < other.setImpl.Count) {
        a = this.setImpl; b = other.setImpl;
      } else {
        a = other.setImpl; b = this.setImpl;
      }
      var r = new HashSet<T>();
      foreach (T t in b)
        r.Add(t);
      foreach (T t in a)
        r.Add(t);
      return new Set<T>(r);
    }
    public Set<T> Intersect(Set<T> other) {
      if (this.setImpl.Count == 0)
        return this;
      else if (other.setImpl.Count == 0)
        return other;
      HashSet<T> a, b;
      if (this.setImpl.Count < other.setImpl.Count) {
        a = this.setImpl; b = other.setImpl;
      } else {
        a = other.setImpl; b = this.setImpl;
      }
      var r = new HashSet<T>();
      foreach (T t in a) {
        if (b.Contains(t))
          r.Add(t);
      }
      return new Set<T>(r);
    }
    public Set<T> Difference(Set<T> other) {
      if (this.setImpl.Count == 0)
        return this;
      else if (other.setImpl.Count == 0)
        return this;
      var r = new HashSet<T>();
      foreach (T t in this.setImpl) {
        if (!other.setImpl.Contains(t))
          r.Add(t);
      }
      return new Set<T>(r);
    }
#endif
  }

  public class MultiSet<T>
  {
#if DAFNY_USE_SYSTEM_COLLECTIONS_IMMUTABLE
    readonly ImmutableDictionary<T, int> dict;
#else
    readonly Dictionary<T, int> dict;
#endif
    readonly BigInteger occurrencesOfNull;  // stupidly, a Dictionary in .NET cannot use "null" as a key
#if DAFNY_USE_SYSTEM_COLLECTIONS_IMMUTABLE
    MultiSet(ImmutableDictionary<T, int>.Builder d, BigInteger occurrencesOfNull) {
      dict = d.ToImmutable();
      this.occurrencesOfNull = occurrencesOfNull;
    }
    public static readonly MultiSet<T> Empty = new MultiSet<T>(ImmutableDictionary<T, int>.Empty.ToBuilder(), BigInteger.Zero);
#else
    MultiSet(Dictionary<T, int> d, BigInteger occurrencesOfNull) {
      this.dict = d;
      this.occurrencesOfNull = occurrencesOfNull;
    }
    public static MultiSet<T> Empty = new MultiSet<T>(new Dictionary<T, int>(0), BigInteger.Zero);
#endif
    public static MultiSet<T> FromElements(params T[] values) {
#if DAFNY_USE_SYSTEM_COLLECTIONS_IMMUTABLE
      var d = ImmutableDictionary<T, int>.Empty.ToBuilder();
#else
      var d = new Dictionary<T, int>(values.Length);
#endif
      var occurrencesOfNull = BigInteger.Zero;
      foreach (T t in values) {
        if (t == null) {
          occurrencesOfNull++;
        } else {
          var i = 0;
          if (!d.TryGetValue(t, out i)) {
            i = 0;
          }
          d[t] = i + 1;
        }
      }
      return new MultiSet<T>(d, occurrencesOfNull);
    }
    public static MultiSet<T> FromCollection(ICollection<T> values) {
#if DAFNY_USE_SYSTEM_COLLECTIONS_IMMUTABLE
      var d = ImmutableDictionary<T, int>.Empty.ToBuilder();
#else
      var d = new Dictionary<T, int>();
#endif
      var occurrencesOfNull = BigInteger.Zero;
      foreach (T t in values) {
        if (t == null) {
          occurrencesOfNull++;
        } else {
          var i = 0;
          if (!d.TryGetValue(t, out i)) {
            i = 0;
          }
          d[t] = i + 1;
        }
      }
      return new MultiSet<T>(d, occurrencesOfNull);
    }
    public static MultiSet<T> FromSeq(Sequence<T> values) {
#if DAFNY_USE_SYSTEM_COLLECTIONS_IMMUTABLE
      var d = ImmutableDictionary<T, int>.Empty.ToBuilder();
#else
      var d = new Dictionary<T, int>();
#endif
      var occurrencesOfNull = BigInteger.Zero;
      foreach (T t in values.Elements) {
        if (t == null) {
          occurrencesOfNull++;
        } else {
          var i = 0;
          if (!d.TryGetValue(t, out i)) {
            i = 0;
          }
          d[t] = i + 1;
        }
      }
      return new MultiSet<T>(d, occurrencesOfNull);
    }
    public static MultiSet<T> FromSet(Set<T> values) {
#if DAFNY_USE_SYSTEM_COLLECTIONS_IMMUTABLE
      var d = ImmutableDictionary<T, int>.Empty.ToBuilder();
#else
      var d = new Dictionary<T, int>();
#endif
      var containsNull = false;
      foreach (T t in values.Elements) {
        if (t == null) {
          containsNull = true;
        } else {
          d[t] = 1;
        }
      }
      return new MultiSet<T>(d, containsNull ? BigInteger.One : BigInteger.Zero);
    }

    public static MultiSet<T> _DafnyDefaultValue() {
      return Empty;
    }

    public bool Equals(MultiSet<T> other) {
      return other.IsSubsetOf(this) && this.IsSubsetOf(other);
    }
    public override bool Equals(object other) {
      return other is MultiSet<T> && Equals((MultiSet<T>)other);
    }
    public override int GetHashCode() {
      var hashCode = 1;
      if (occurrencesOfNull > 0) {
        var key = Dafny.Helpers.GetHashCode(default(T));
        key = (key << 3) | (key >> 29) ^ occurrencesOfNull.GetHashCode();
        hashCode = hashCode * (key + 3);
      }
      foreach (var kv in dict) {
        var key = Dafny.Helpers.GetHashCode(kv.Key);
        key = (key << 3) | (key >> 29) ^ kv.Value.GetHashCode();
        hashCode = hashCode * (key + 3);
      }
      return hashCode;
    }
    public override string ToString() {
      var s = "multiset{";
      var sep = "";
      for (var i = BigInteger.Zero; i < occurrencesOfNull; i++) {
        s += sep + Dafny.Helpers.ToString(default(T));
        sep = ", ";
      }
      foreach (var kv in dict) {
        var t = Dafny.Helpers.ToString(kv.Key);
        for (int i = 0; i < kv.Value; i++) {
          s += sep + t;
          sep = ", ";
        }
      }
      return s + "}";
    }
    public bool IsProperSubsetOf(MultiSet<T> other) {
      return !Equals(other) && IsSubsetOf(other);
    }
    public bool IsSubsetOf(MultiSet<T> other) {
      if (other.occurrencesOfNull < this.occurrencesOfNull) {
        return false;
      }
      foreach (T t in dict.Keys) {
        if (!other.dict.ContainsKey(t) || other.dict[t] < dict[t])
          return false;
      }
      return true;
    }
    public bool IsSupersetOf(MultiSet<T> other) {
      return other.IsSubsetOf(this);
    }
    public bool IsProperSupersetOf(MultiSet<T> other) {
      return other.IsProperSubsetOf(this);
    }
    public bool IsDisjointFrom(MultiSet<T> other) {
      if (occurrencesOfNull > 0 && other.occurrencesOfNull > 0) {
        return false;
      }
      foreach (T t in dict.Keys) {
        if (other.dict.ContainsKey(t))
          return false;
      }
      foreach (T t in other.dict.Keys) {
        if (dict.ContainsKey(t))
          return false;
      }
      return true;
    }

    public bool Contains<G>(G t) {
      return t == null ? occurrencesOfNull > 0 : t is T && dict.ContainsKey((T)(object)t);
    }
    public BigInteger Select<G>(G t) {
      if (t == null) {
        return occurrencesOfNull;
      } else if (t is T && dict.ContainsKey((T)(object)t)) {
        return dict[(T)(object)t];
      } else {
        return BigInteger.Zero;
      }
    }
    public MultiSet<T> Update<G>(G t, BigInteger i) {
      if (Select(t) == i) {
        return this;
      } else if (t == null) {
#if DAFNY_USE_SYSTEM_COLLECTIONS_IMMUTABLE
        var r = dict.ToBuilder();
#else
        var r = dict;
#endif
        return new MultiSet<T>(r, i);
      } else {
#if DAFNY_USE_SYSTEM_COLLECTIONS_IMMUTABLE
        var r = dict.ToBuilder();
#else
        var r = new Dictionary<T, int>(dict);
#endif
        r[(T)(object)t] = (int)i;
        return new MultiSet<T>(r, occurrencesOfNull);
      }
    }
    public MultiSet<T> Union(MultiSet<T> other) {
      if (dict.Count + occurrencesOfNull == 0)
        return other;
      else if (other.dict.Count + other.occurrencesOfNull == 0)
        return this;
#if DAFNY_USE_SYSTEM_COLLECTIONS_IMMUTABLE
      var r = ImmutableDictionary<T, int>.Empty.ToBuilder();
#else
      var r = new Dictionary<T, int>();
#endif
      foreach (T t in dict.Keys) {
        var i = 0;
        if (!r.TryGetValue(t, out i)) {
          i = 0;
        }
        r[t] = i + dict[t];
      }
      foreach (T t in other.dict.Keys) {
        var i = 0;
        if (!r.TryGetValue(t, out i)) {
          i = 0;
        }
        r[t] = i + other.dict[t];
      }
      return new MultiSet<T>(r, occurrencesOfNull + other.occurrencesOfNull);
    }
    public MultiSet<T> Intersect(MultiSet<T> other) {
      if (dict.Count == 0 && occurrencesOfNull == 0)
        return this;
      else if (other.dict.Count == 0 && other.occurrencesOfNull == 0)
        return other;
#if DAFNY_USE_SYSTEM_COLLECTIONS_IMMUTABLE
      var r = ImmutableDictionary<T, int>.Empty.ToBuilder();
#else
      var r = new Dictionary<T, int>();
#endif
      foreach (T t in dict.Keys) {
        if (other.dict.ContainsKey(t)) {
          r.Add(t, other.dict[t] < dict[t] ? other.dict[t] : dict[t]);
        }
      }
      return new MultiSet<T>(r, other.occurrencesOfNull < occurrencesOfNull ? other.occurrencesOfNull : occurrencesOfNull);
    }
    public MultiSet<T> Difference(MultiSet<T> other) { // \result == this - other
      if (dict.Count == 0 && occurrencesOfNull == 0)
        return this;
      else if (other.dict.Count == 0 && other.occurrencesOfNull == 0)
        return this;
#if DAFNY_USE_SYSTEM_COLLECTIONS_IMMUTABLE
      var r = ImmutableDictionary<T, int>.Empty.ToBuilder();
#else
      var r = new Dictionary<T, int>();
#endif
      foreach (T t in dict.Keys) {
        if (!other.dict.ContainsKey(t)) {
          r.Add(t, dict[t]);
        } else if (other.dict[t] < dict[t]) {
          r.Add(t, dict[t] - other.dict[t]);
        }
      }
      return new MultiSet<T>(r, other.occurrencesOfNull < occurrencesOfNull ? occurrencesOfNull - other.occurrencesOfNull : BigInteger.Zero);
    }

    public int Count {
      get { return (int)ElementCount(); }
    }
    public long LongCount {
      get { return (long)ElementCount(); }
    }
    private BigInteger ElementCount() {
      // This is inefficient
      var c = occurrencesOfNull;
      foreach (var item in dict) {
        c += item.Value;
      }
      return c;
    }

    public IEnumerable<T> Elements {
      get {
        for (var i = BigInteger.Zero; i < occurrencesOfNull; i++) {
          yield return default(T);
        }
        foreach (var item in dict) {
          for (int i = 0; i < item.Value; i++) {
            yield return item.Key;
          }
        }
      }
    }

    public IEnumerable<T> UniqueElements {
      get {
        if (!occurrencesOfNull.IsZero) {
          yield return default(T);
        }
        foreach (var item in dict) {
          yield return item.Key;
        }
      }
    }
  }

  public class Map<U, V>
  {
#if DAFNY_USE_SYSTEM_COLLECTIONS_IMMUTABLE
    readonly ImmutableDictionary<U, V> dict;
#else
    readonly Dictionary<U, V> dict;
#endif
    readonly bool hasNullValue;  // true when "null" is a key of the Map
    readonly V nullValue;  // if "hasNullValue", the value that "null" maps to

#if DAFNY_USE_SYSTEM_COLLECTIONS_IMMUTABLE
    Map(ImmutableDictionary<U, V>.Builder d, bool hasNullValue, V nullValue) {
      dict = d.ToImmutable();
      this.hasNullValue = hasNullValue;
      this.nullValue = nullValue;
    }
    public static readonly Map<U, V> Empty = new Map<U, V>(ImmutableDictionary<U, V>.Empty.ToBuilder(), false, default(V));
#else
    Map(Dictionary<U, V> d, bool hasNullValue, V nullValue) {
      this.dict = d;
      this.hasNullValue = hasNullValue;
      this.nullValue = nullValue;
    }
    public static readonly Map<U, V> Empty = new Map<U, V>(new Dictionary<U, V>(), false, default(V));
#endif

    public static Map<U, V> FromElements(params Pair<U, V>[] values) {
#if DAFNY_USE_SYSTEM_COLLECTIONS_IMMUTABLE
      var d = ImmutableDictionary<U, V>.Empty.ToBuilder();
#else
      var d = new Dictionary<U, V>(values.Length);
#endif
      var hasNullValue = false;
      var nullValue = default(V);
      foreach (Pair<U, V> p in values) {
        if (p.Car == null) {
          hasNullValue = true;
          nullValue = p.Cdr;
        } else {
          d[p.Car] = p.Cdr;
        }
      }
      return new Map<U, V>(d, hasNullValue, nullValue);
    }
    public static Map<U, V> FromCollection(List<Pair<U, V>> values) {
#if DAFNY_USE_SYSTEM_COLLECTIONS_IMMUTABLE
      var d = ImmutableDictionary<U, V>.Empty.ToBuilder();
#else
      var d = new Dictionary<U, V>(values.Count);
#endif
      var hasNullValue = false;
      var nullValue = default(V);
      foreach (Pair<U, V> p in values) {
        if (p.Car == null) {
          hasNullValue = true;
          nullValue = p.Cdr;
        } else {
          d[p.Car] = p.Cdr;
        }
      }
      return new Map<U, V>(d, hasNullValue, nullValue);
    }
    public int Count {
      get { return dict.Count + (hasNullValue ? 1 : 0); }
    }
    public long LongCount {
      get { return dict.Count + (hasNullValue ? 1 : 0); }
    }
    public static Map<U, V> _DafnyDefaultValue() {
      return Empty;
    }

    public bool Equals(Map<U, V> other) {
      if (hasNullValue != other.hasNullValue || dict.Count != other.dict.Count) {
        return false;
      } else if (hasNullValue && !Dafny.Helpers.AreEqual(nullValue, other.nullValue)) {
        return false;
      }
      foreach (U u in dict.Keys) {
        V v1 = dict[u];
        V v2;
        if (!other.dict.TryGetValue(u, out v2)) {
          return false; // other dictionary does not contain this element
        }
        if (!Dafny.Helpers.AreEqual(v1, v2)) {
          return false;
        }
      }
      return true;
    }
    public override bool Equals(object other) {
      return other is Map<U, V> && Equals((Map<U, V>)other);
    }
    public override int GetHashCode() {
      var hashCode = 1;
      if (hasNullValue) {
        var key = Dafny.Helpers.GetHashCode(default(U));
        key = (key << 3) | (key >> 29) ^ Dafny.Helpers.GetHashCode(nullValue);
        hashCode = hashCode * (key + 3);
      }
      foreach (var kv in dict) {
        var key = Dafny.Helpers.GetHashCode(kv.Key);
        key = (key << 3) | (key >> 29) ^ Dafny.Helpers.GetHashCode(kv.Value);
        hashCode = hashCode * (key + 3);
      }
      return hashCode;
    }
    public override string ToString() {
      var s = "map[";
      var sep = "";
      if (hasNullValue) {
        s += sep + Dafny.Helpers.ToString(default(U)) + " := " + Dafny.Helpers.ToString(nullValue);
        sep = ", ";
      }
      foreach (var kv in dict) {
        s += sep + Dafny.Helpers.ToString(kv.Key) + " := " + Dafny.Helpers.ToString(kv.Value);
        sep = ", ";
      }
      return s + "]";
    }
    public bool IsDisjointFrom(Map<U, V> other) {
      if (hasNullValue && other.hasNullValue) {
        return false;
      }
      foreach (U u in dict.Keys) {
        if (other.dict.ContainsKey(u))
          return false;
      }
      foreach (U u in other.dict.Keys) {
        if (dict.ContainsKey(u))
          return false;
      }
      return true;
    }
    public bool Contains<G>(G u) {
      return u == null ? hasNullValue : u is U && dict.ContainsKey((U)(object)u);
    }
    public V Select(U index) {
      // evidently, the following will throw some exception if "index" in not a key of the map
      return index == null && hasNullValue ? nullValue : dict[index];
    }
#if DAFNY_USE_SYSTEM_COLLECTIONS_IMMUTABLE
    public Map<U, V> Update(U index, V val) {
      var d = dict.ToBuilder();
      if (index == null) {
        return new Map<U, V>(d, true, val);
      } else {
        d[index] = val;
        return new Map<U, V>(d, hasNullValue, nullValue);
      }
    }
#else
    public Map<U, V> Update(U index, V val) {
      if (index == null) {
        return new Map<U, V>(dict, true, val);
      } else {
        var d = new Dictionary<U, V>(dict);
        d[index] = val;
        return new Map<U, V>(d, hasNullValue, nullValue);
      }
    }
#endif
    public Set<U> Keys {
      get {
        if (hasNullValue) {
          return Dafny.Set<U>.FromCollectionPlusOne(dict.Keys, default(U));
        } else {
          return Dafny.Set<U>.FromCollection(dict.Keys);
        }
      }
    }
    public Set<V> Values {
      get {
        if (hasNullValue) {
          return Dafny.Set<V>.FromCollectionPlusOne(dict.Values, nullValue);
        } else {
          return Dafny.Set<V>.FromCollection(dict.Values);
        }
      }
    }
    public Set<_System.Tuple2<U, V>> Items {
      get {
        HashSet<_System.Tuple2<U, V>> result = new HashSet<_System.Tuple2<U, V>>();
        if (hasNullValue) {
          result.Add(_System.Tuple2<U, V>.create(default(U), nullValue));
        }
        foreach (KeyValuePair<U, V> kvp in dict) {
          result.Add(_System.Tuple2<U, V>.create(kvp.Key, kvp.Value));
        }
        return Dafny.Set<_System.Tuple2<U, V>>.FromCollection(result);
      }
    }
  }

  public class Sequence<T>
  {
    readonly T[] elmts;
    public Sequence(T[] ee) {
      elmts = ee;
    }
    public static Sequence<T> Empty {
      get {
        return new Sequence<T>(new T[0]);
      }
    }
    public static Sequence<T> FromElements(params T[] values) {
      return new Sequence<T>(values);
    }
    public static Sequence<char> FromString(string s) {
      return new Sequence<char>(s.ToCharArray());
    }
    public static Sequence<T> _DafnyDefaultValue() {
      return Empty;
    }
    public int Count {
      get { return elmts.Length; }
    }
    public long LongCount {
      get { return elmts.LongLength; }
    }
    public T[] Elements {
      get {
        return elmts;
      }
    }
    public IEnumerable<T> UniqueElements {
      get {
        var st = Set<T>.FromElements(elmts);
        return st.Elements;
      }
    }
    public T Select(ulong index) {
      return elmts[index];
    }
    public T Select(long index) {
      return elmts[index];
    }
    public T Select(uint index) {
      return elmts[index];
    }
    public T Select(int index) {
      return elmts[index];
    }
    public T Select(BigInteger index) {
      return elmts[(int)index];
    }
    public Sequence<T> Update(long index, T t) {
      T[] a = (T[])elmts.Clone();
      a[index] = t;
      return new Sequence<T>(a);
    }
    public Sequence<T> Update(ulong index, T t) {
      return Update((long)index, t);
    }
    public Sequence<T> Update(BigInteger index, T t) {
      return Update((long)index, t);
    }
    public bool Equals(Sequence<T> other) {
      int n = elmts.Length;
      return n == other.elmts.Length && EqualUntil(other, n);
    }
    public override bool Equals(object other) {
      return other is Sequence<T> && Equals((Sequence<T>)other);
    }
    public override int GetHashCode() {
      if (elmts == null || elmts.Length == 0)
        return 0;
      var hashCode = 0;
      for (var i = 0; i < elmts.Length; i++) {
        hashCode = (hashCode << 3) | (hashCode >> 29) ^ Dafny.Helpers.GetHashCode(elmts[i]);
      }
      return hashCode;
    }
    public override string ToString() {
      if (elmts is char[]) {
        var s = "";
        foreach (var t in elmts) {
          s += t.ToString();
        }
        return s;
      } else {
        var s = "[";
        var sep = "";
        foreach (var t in elmts) {
          s += sep + Dafny.Helpers.ToString(t);
          sep = ", ";
        }
        return s + "]";
      }
    }
    bool EqualUntil(Sequence<T> other, int n) {
      for (int i = 0; i < n; i++) {
        if (!Dafny.Helpers.AreEqual(elmts[i], other.elmts[i]))
          return false;
      }
      return true;
    }
    public bool IsProperPrefixOf(Sequence<T> other) {
      int n = elmts.Length;
      return n < other.elmts.Length && EqualUntil(other, n);
    }
    public bool IsPrefixOf(Sequence<T> other) {
      int n = elmts.Length;
      return n <= other.elmts.Length && EqualUntil(other, n);
    }
    public Sequence<T> Concat(Sequence<T> other) {
      if (elmts.Length == 0)
        return other;
      else if (other.elmts.Length == 0)
        return this;
      T[] a = new T[elmts.Length + other.elmts.Length];
      System.Array.Copy(elmts, 0, a, 0, elmts.Length);
      System.Array.Copy(other.elmts, 0, a, elmts.Length, other.elmts.Length);
      return new Sequence<T>(a);
    }
    public bool Contains<G>(G g) {
      if (g == null || g is T) {
        var t = (T)(object)g;
        int n = elmts.Length;
        for (int i = 0; i < n; i++) {
          if (Dafny.Helpers.AreEqual(t, elmts[i]))
            return true;
        }
      }
      return false;
    }
    public Sequence<T> Take(long m) {
      if (elmts.LongLength == m)
        return this;
      T[] a = new T[m];
      System.Array.Copy(elmts, a, m);
      return new Sequence<T>(a);
    }
    public Sequence<T> Take(ulong n) {
      return Take((long)n);
    }
    public Sequence<T> Take(BigInteger n) {
      return Take((long)n);
    }
    public Sequence<T> Drop(long m) {
      if (m == 0)
        return this;
      T[] a = new T[elmts.Length - m];
      System.Array.Copy(elmts, m, a, 0, elmts.Length - m);
      return new Sequence<T>(a);
    }
    public Sequence<T> Drop(ulong n) {
      return Drop((long)n);
    }
    public Sequence<T> Drop(BigInteger n) {
      if (n.IsZero)
        return this;
      return Drop((long)n);
    }
  }
  public struct Pair<A, B>
  {
    public readonly A Car;
    public readonly B Cdr;
    public Pair(A a, B b) {
      this.Car = a;
      this.Cdr = b;
    }
  }
  public partial class Helpers {
    public static bool AreEqual<G>(G a, G b) {
      return a == null ? b == null : a.Equals(b);
    }
    public static int GetHashCode<G>(G g) {
      return g == null ? 1001 : g.GetHashCode();
    }
    public static string ToString<G>(G g) {
      if (g == null) {
        return "null";
      } else if (g is bool) {
        return (bool)(object)g ? "true" : "false";  // capitalize boolean literals like in Dafny
      } else {
        return g.ToString();
      }
    }
    public static void Print<G>(G g) {
      System.Console.Write(ToString(g));
    }
    public static G Default<G>() {
      System.Type ty = typeof(G);
      System.Reflection.MethodInfo mInfo = ty.GetMethod("_DafnyDefaultValue");
      if (mInfo != null) {
        G g = (G)mInfo.Invoke(null, null);
        return g;
      } else {
        return default(G);
      }
    }
    // Computing forall/exists quantifiers
    public static bool Quantifier<T>(IEnumerable<T> vals, bool frall, System.Predicate<T> pred) {
      foreach (var u in vals) {
        if (pred(u) != frall) { return !frall; }
      }
      return frall;
    }
    // Enumerating other collections
    public static IEnumerable<bool> AllBooleans() {
      yield return false;
      yield return true;
    }
    public static IEnumerable<char> AllChars() {
      for (int i = 0; i < 0x10000; i++) {
        yield return (char)i;
      }
    }
    public static IEnumerable<BigInteger> AllIntegers() {
      yield return new BigInteger(0);
      for (var j = new BigInteger(1);; j++) {
        yield return j;
        yield return -j;
      }
    }
    public static IEnumerable<BigInteger> IntegerRange(Nullable<BigInteger> lo, Nullable<BigInteger> hi) {
      if (lo == null) {
        for (var j = (BigInteger)hi; true; ) {
          j--;
          yield return j;
        }
      } else if (hi == null) {
        for (var j = (BigInteger)lo; true; j++) {
          yield return j;
        }
      } else {
        for (var j = (BigInteger)lo; j < hi; j++) {
          yield return j;
        }
      }
    }
    public static IEnumerable<T> SingleValue<T>(T e) {
      yield return e;
    }
    // pre: b != 0
    // post: result == a/b, as defined by Euclidean Division (http://en.wikipedia.org/wiki/Modulo_operation)
    public static sbyte EuclideanDivision_sbyte(sbyte a, sbyte b) {
      return (sbyte)EuclideanDivision_int(a, b);
    }
    public static short EuclideanDivision_short(short a, short b) {
      return (short)EuclideanDivision_int(a, b);
    }
    public static int EuclideanDivision_int(int a, int b) {
      if (0 <= a) {
        if (0 <= b) {
          // +a +b: a/b
          return (int)(((uint)(a)) / ((uint)(b)));
        } else {
          // +a -b: -(a/(-b))
          return -((int)(((uint)(a)) / ((uint)(unchecked(-b)))));
        }
      } else {
        if (0 <= b) {
          // -a +b: -((-a-1)/b) - 1
          return -((int)(((uint)(-(a + 1))) / ((uint)(b)))) - 1;
        } else {
          // -a -b: ((-a-1)/(-b)) + 1
          return ((int)(((uint)(-(a + 1))) / ((uint)(unchecked(-b))))) + 1;
        }
      }
    }
    public static long EuclideanDivision_long(long a, long b) {
      if (0 <= a) {
        if (0 <= b) {
          // +a +b: a/b
          return (long)(((ulong)(a)) / ((ulong)(b)));
        } else {
          // +a -b: -(a/(-b))
          return -((long)(((ulong)(a)) / ((ulong)(unchecked(-b)))));
        }
      } else {
        if (0 <= b) {
          // -a +b: -((-a-1)/b) - 1
          return -((long)(((ulong)(-(a + 1))) / ((ulong)(b)))) - 1;
        } else {
          // -a -b: ((-a-1)/(-b)) + 1
          return ((long)(((ulong)(-(a + 1))) / ((ulong)(unchecked(-b))))) + 1;
        }
      }
    }
    public static BigInteger EuclideanDivision(BigInteger a, BigInteger b) {
      if (0 <= a.Sign) {
        if (0 <= b.Sign) {
          // +a +b: a/b
          return BigInteger.Divide(a, b);
        } else {
          // +a -b: -(a/(-b))
          return BigInteger.Negate(BigInteger.Divide(a, BigInteger.Negate(b)));
        }
      } else {
        if (0 <= b.Sign) {
          // -a +b: -((-a-1)/b) - 1
          return BigInteger.Negate(BigInteger.Divide(BigInteger.Negate(a) - 1, b)) - 1;
        } else {
          // -a -b: ((-a-1)/(-b)) + 1
          return BigInteger.Divide(BigInteger.Negate(a) - 1, BigInteger.Negate(b)) + 1;
        }
      }
    }
    // pre: b != 0
    // post: result == a%b, as defined by Euclidean Division (http://en.wikipedia.org/wiki/Modulo_operation)
    public static sbyte EuclideanModulus_sbyte(sbyte a, sbyte b) {
      return (sbyte)EuclideanModulus_int(a, b);
    }
    public static short EuclideanModulus_short(short a, short b) {
      return (short)EuclideanModulus_int(a, b);
    }
    public static int EuclideanModulus_int(int a, int b) {
      uint bp = (0 <= b) ? (uint)b : (uint)(unchecked(-b));
      if (0 <= a) {
        // +a: a % b'
        return (int)(((uint)a) % bp);
      } else {
        // c = ((-a) % b')
        // -a: b' - c if c > 0
        // -a: 0 if c == 0
        uint c = ((uint)(unchecked(-a))) % bp;
        return (int)(c == 0 ? c : bp - c);
      }
    }
    public static long EuclideanModulus_long(long a, long b) {
      ulong bp = (0 <= b) ? (ulong)b : (ulong)(unchecked(-b));
      if (0 <= a) {
        // +a: a % b'
        return (long)(((ulong)a) % bp);
      } else {
        // c = ((-a) % b')
        // -a: b' - c if c > 0
        // -a: 0 if c == 0
        ulong c = ((ulong)(unchecked(-a))) % bp;
        return (long)(c == 0 ? c : bp - c);
      }
    }
    public static BigInteger EuclideanModulus(BigInteger a, BigInteger b) {
      var bp = BigInteger.Abs(b);
      if (0 <= a.Sign) {
        // +a: a % b'
        return BigInteger.Remainder(a, bp);
      } else {
        // c = ((-a) % b')
        // -a: b' - c if c > 0
        // -a: 0 if c == 0
        var c = BigInteger.Remainder(BigInteger.Negate(a), bp);
        return c.IsZero ? c : BigInteger.Subtract(bp, c);
      }
    }
    public static Sequence<T> SeqFromArray<T>(T[] array) {
      return new Sequence<T>((T[])array.Clone());
    }
    // In .NET version 4.5, it it possible to mark a method with "AggressiveInlining", which says to inline the
    // method if possible.  Method "ExpressionSequence" would be a good candidate for it:
    // [System.Runtime.CompilerServices.MethodImpl(System.Runtime.CompilerServices.MethodImplOptions.AggressiveInlining)]
    public static U ExpressionSequence<T, U>(T t, U u)
    {
      return u;
    }

    public static U Let<T, U>(T t, Func<T,U> f) {
      return f(t);
    }

    public static A Id<A>(A a) {
      return a;
    }
  }

  public class BigOrdinal {
    public static bool IsLimit(BigInteger ord) {
      return ord == 0;
    }
    public static bool IsSucc(BigInteger ord) {
      return 0 < ord;
    }
    public static BigInteger Offset(BigInteger ord) {
      return ord;
    }
    public static bool IsNat(BigInteger ord) {
      return true;  // at run time, every ORDINAL is a natural number
    }
  }

  public struct BigRational
  {
    public static readonly BigRational ZERO = new BigRational(0);

    // We need to deal with the special case "num == 0 && den == 0", because
    // that's what C#'s default struct constructor will produce for BigRational. :(
    // To deal with it, we ignore "den" when "num" is 0.
    BigInteger num, den;  // invariant 1 <= den || (num == 0 && den == 0)
    public override string ToString() {
      int log10;
      if (num.IsZero || den.IsOne) {
        return string.Format("{0}.0", num);
      } else if (IsPowerOf10(den, out log10)) {
        string sign;
        string digits;
        if (num.Sign < 0) {
          sign = "-"; digits = (-num).ToString();
        } else {
          sign = ""; digits = num.ToString();
        }
        if (log10 < digits.Length) {
          var n = digits.Length - log10;
          return string.Format("{0}{1}.{2}", sign, digits.Substring(0, n), digits.Substring(n));
        } else {
          return string.Format("{0}0.{1}{2}", sign, new string('0', log10 - digits.Length), digits);
        }
      } else {
        return string.Format("({0}.0 / {1}.0)", num, den);
      }
    }
    public bool IsPowerOf10(BigInteger x, out int log10) {
      log10 = 0;
      if (x.IsZero) {
        return false;
      }
      while (true) {  // invariant: x != 0 && x * 10^log10 == old(x)
        if (x.IsOne) {
          return true;
        } else if (x % 10 == 0) {
          log10++;
          x /= 10;
        } else {
          return false;
        }
      }
    }
    public BigRational(int n) {
      num = new BigInteger(n);
      den = BigInteger.One;
    }
    public BigRational(BigInteger n, BigInteger d) {
      // requires 1 <= d
      num = n;
      den = d;
    }
    public BigInteger ToBigInteger() {
      if (num.IsZero || den.IsOne) {
        return num;
      } else if (0 < num.Sign) {
        return num / den;
      } else {
        return (num - den + 1) / den;
      }
    }
    /// <summary>
    /// Returns values such that aa/dd == a and bb/dd == b.
    /// </summary>
    private static void Normalize(BigRational a, BigRational b, out BigInteger aa, out BigInteger bb, out BigInteger dd) {
      if (a.num.IsZero) {
        aa = a.num;
        bb = b.num;
        dd = b.den;
      } else if (b.num.IsZero) {
        aa = a.num;
        dd = a.den;
        bb = b.num;
      } else {
        var gcd = BigInteger.GreatestCommonDivisor(a.den, b.den);
        var xx = a.den / gcd;
        var yy = b.den / gcd;
        // We now have a == a.num / (xx * gcd) and b == b.num / (yy * gcd).
        aa = a.num * yy;
        bb = b.num * xx;
        dd = a.den * yy;
      }
    }
    public int CompareTo(BigRational that) {
      // simple things first
      int asign = this.num.Sign;
      int bsign = that.num.Sign;
      if (asign < 0 && 0 <= bsign) {
        return -1;
      } else if (asign <= 0 && 0 < bsign) {
        return -1;
      } else if (bsign < 0 && 0 <= asign) {
        return 1;
      } else if (bsign <= 0 && 0 < asign) {
        return 1;
      }
      BigInteger aa, bb, dd;
      Normalize(this, that, out aa, out bb, out dd);
      return aa.CompareTo(bb);
    }
    public override int GetHashCode() {
      return num.GetHashCode() + 29 * den.GetHashCode();
    }
    public override bool Equals(object obj) {
      if (obj is BigRational) {
        return this == (BigRational)obj;
      } else {
        return false;
      }
    }
    public static bool operator ==(BigRational a, BigRational b) {
      return a.CompareTo(b) == 0;
    }
    public static bool operator !=(BigRational a, BigRational b) {
      return a.CompareTo(b) != 0;
    }
    public static bool operator >(BigRational a, BigRational b) {
      return a.CompareTo(b) > 0;
    }
    public static bool operator >=(BigRational a, BigRational b) {
      return a.CompareTo(b) >= 0;
    }
    public static bool operator <(BigRational a, BigRational b) {
      return a.CompareTo(b) < 0;
    }
    public static bool operator <=(BigRational a, BigRational b) {
      return a.CompareTo(b) <= 0;
    }
    public static BigRational operator +(BigRational a, BigRational b) {
      BigInteger aa, bb, dd;
      Normalize(a, b, out aa, out bb, out dd);
      return new BigRational(aa + bb, dd);
    }
    public static BigRational operator -(BigRational a, BigRational b) {
      BigInteger aa, bb, dd;
      Normalize(a, b, out aa, out bb, out dd);
      return new BigRational(aa - bb, dd);
    }
    public static BigRational operator -(BigRational a) {
      return new BigRational(-a.num, a.den);
    }
    public static BigRational operator *(BigRational a, BigRational b) {
      return new BigRational(a.num * b.num, a.den * b.den);
    }
    public static BigRational operator /(BigRational a, BigRational b) {
      // Compute the reciprocal of b
      BigRational bReciprocal;
      if (0 < b.num.Sign) {
        bReciprocal = new BigRational(b.den, b.num);
      } else {
        // this is the case b.num < 0
        bReciprocal = new BigRational(-b.den, -b.num);
      }
      return a * bReciprocal;
    }
  }
}

namespace @_System
{
  public class Tuple2<T0,T1> {
    public readonly T0 _0;
    public readonly T1 _1;
    public Tuple2(T0 _0, T1 _1) {
      this._0 = _0;
      this._1 = _1;
    }
    public override bool Equals(object other) {
      var oth = other as _System.@Tuple2<T0,T1>;
      return oth != null && Dafny.Helpers.AreEqual(this._0, oth._0) && Dafny.Helpers.AreEqual(this._1, oth._1);
    }
    public override int GetHashCode() {
      ulong hash = 5381;
      hash = ((hash << 5) + hash) + 0;
      hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this._0));
      hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this._1));
      return (int) hash;
    }
    public override string ToString() {
      string s = "";
      s += "(";
      s += Dafny.Helpers.ToString(this._0);
      s += ", ";
      s += Dafny.Helpers.ToString(this._1);
      s += ")";
      return s;
    }
    static Tuple2<T0,T1> theDefault;
    public static Tuple2<T0,T1> Default {
      get {
        if (theDefault == null) {
          theDefault = new _System.@Tuple2<T0,T1>(Dafny.Helpers.Default<T0>(), Dafny.Helpers.Default<T1>());
        }
        return theDefault;
      }
    }
    public static Tuple2<T0,T1> _DafnyDefaultValue() { return Default; }
    public static Tuple2<T0,T1> create(T0 _0, T1 _1) {
      return new Tuple2<T0,T1>(_0, _1);
    }
    public bool is____hMake3 { get { return true; } }
    public T0 dtor__0 {
      get {
        return this._0;
      }
    }
    public T1 dtor__1 {
      get {
        return this._1;
      }
    }
  }

} // end of namespace _System
namespace Dafny {
  internal class ArrayHelpers {
    public static T[] InitNewArray1<T>(T z, BigInteger size0) {
      int s0 = (int)size0;
T[] a = new T[s0];
for (int i0 = 0; i0 < s0; i0++) {
        a[i0] = z;
      }
      return a;
    }
  }
} // end of namespace Dafny
namespace _System {


  public partial class nat {
  }

  public class Tuple3<T0,T1,T2> {
    public readonly T0 _0;
public readonly T1 _1;
public readonly T2 _2;
public Tuple3(T0 _0, T1 _1, T2 _2) {
      this._0 = _0;
this._1 = _1;
this._2 = _2;
    }
    public override bool Equals(object other) {
      var oth = other as _System.Tuple3<T0,T1,T2>;
return oth != null && Dafny.Helpers.AreEqual(this._0, oth._0) && Dafny.Helpers.AreEqual(this._1, oth._1) && Dafny.Helpers.AreEqual(this._2, oth._2);
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 0;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this._0));
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this._1));
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this._2));
return (int) hash;
    }
    public override string ToString() {
      string s = "";
s += "(";
s += Dafny.Helpers.ToString(this._0);
s += ", ";
s += Dafny.Helpers.ToString(this._1);
s += ", ";
s += Dafny.Helpers.ToString(this._2);
s += ")";
return s;
    }
    static Tuple3<T0,T1,T2> theDefault;
public static Tuple3<T0,T1,T2> Default {
      get {
        if (theDefault == null) {
          theDefault = new _System.Tuple3<T0,T1,T2>(Dafny.Helpers.Default<T0>(), Dafny.Helpers.Default<T1>(), Dafny.Helpers.Default<T2>());
        }
        return theDefault;
      }
    }
    public static Tuple3<T0,T1,T2> _DafnyDefaultValue() { return Default; }
public static Tuple3<T0,T1,T2> create(T0 _0, T1 _1, T2 _2) {
      return new Tuple3<T0,T1,T2>(_0, _1, _2);
    }
    public bool is____hMake3 { get { return true; } }
public T0 dtor__0 {
      get {
        return this._0;
      }
    }
    public T1 dtor__1 {
      get {
        return this._1;
      }
    }
    public T2 dtor__2 {
      get {
        return this._2;
      }
    }
  }








  public class Tuple0 {
    public Tuple0() {
    }
    public override bool Equals(object other) {
      var oth = other as _System.Tuple0;
return oth != null;
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 0;
return (int) hash;
    }
    public override string ToString() {
      return "()";
    }
    static Tuple0 theDefault;
public static Tuple0 Default {
      get {
        if (theDefault == null) {
          theDefault = new _System.Tuple0();
        }
        return theDefault;
      }
    }
    public static Tuple0 _DafnyDefaultValue() { return Default; }
public static Tuple0 create() {
      return new Tuple0();
    }
    public bool is____hMake0 { get { return true; } }
public static System.Collections.Generic.IEnumerable<Tuple0> AllSingletonConstructors {
      get {
        yield return Tuple0.create();
      }
    }
  }
} // end of namespace _System
namespace _0_Native____NativeTypes__s_Compile {

  public partial class @sbyte {
    public static System.Collections.Generic.IEnumerable<sbyte> IntegerRange(BigInteger lo, BigInteger hi) {
      for (var j = lo; j < hi; j++) { yield return (sbyte)j; }
    }
  }

  public partial class @byte {
    public static System.Collections.Generic.IEnumerable<byte> IntegerRange(BigInteger lo, BigInteger hi) {
      for (var j = lo; j < hi; j++) { yield return (byte)j; }
    }
  }

  public partial class int16 {
    public static System.Collections.Generic.IEnumerable<short> IntegerRange(BigInteger lo, BigInteger hi) {
      for (var j = lo; j < hi; j++) { yield return (short)j; }
    }
  }

  public partial class uint16 {
    public static System.Collections.Generic.IEnumerable<ushort> IntegerRange(BigInteger lo, BigInteger hi) {
      for (var j = lo; j < hi; j++) { yield return (ushort)j; }
    }
  }

  public partial class int32 {
    public static System.Collections.Generic.IEnumerable<int> IntegerRange(BigInteger lo, BigInteger hi) {
      for (var j = lo; j < hi; j++) { yield return (int)j; }
    }
  }

  public partial class uint32 {
    public static System.Collections.Generic.IEnumerable<uint> IntegerRange(BigInteger lo, BigInteger hi) {
      for (var j = lo; j < hi; j++) { yield return (uint)j; }
    }
  }

  public partial class int64 {
    public static System.Collections.Generic.IEnumerable<long> IntegerRange(BigInteger lo, BigInteger hi) {
      for (var j = lo; j < hi; j++) { yield return (long)j; }
    }
  }

  public partial class uint64 {
    public static System.Collections.Generic.IEnumerable<ulong> IntegerRange(BigInteger lo, BigInteger hi) {
      for (var j = lo; j < hi; j++) { yield return (ulong)j; }
    }
  }

  public partial class nat8 {
    public static System.Collections.Generic.IEnumerable<sbyte> IntegerRange(BigInteger lo, BigInteger hi) {
      for (var j = lo; j < hi; j++) { yield return (sbyte)j; }
    }
  }

  public partial class nat16 {
    public static System.Collections.Generic.IEnumerable<short> IntegerRange(BigInteger lo, BigInteger hi) {
      for (var j = lo; j < hi; j++) { yield return (short)j; }
    }
  }

  public partial class nat32 {
    public static System.Collections.Generic.IEnumerable<int> IntegerRange(BigInteger lo, BigInteger hi) {
      for (var j = lo; j < hi; j++) { yield return (int)j; }
    }
  }

  public partial class nat64 {
    public static System.Collections.Generic.IEnumerable<long> IntegerRange(BigInteger lo, BigInteger hi) {
      for (var j = lo; j < hi; j++) { yield return (long)j; }
    }
  }

} // end of namespace _0_Native____NativeTypes__s_Compile
namespace _2_Collections____Maps2__s_Compile {

} // end of namespace _2_Collections____Maps2__s_Compile
namespace _5_Temporal____Temporal__s_Compile {




} // end of namespace _5_Temporal____Temporal__s_Compile
namespace _7_Environment__s_Compile {



  public class LPacket<IdType,MessageType> {
    public readonly IdType dst;
public readonly IdType src;
public readonly MessageType msg;
public LPacket(IdType dst, IdType src, MessageType msg) {
      this.dst = dst;
this.src = src;
this.msg = msg;
    }
    public override bool Equals(object other) {
      var oth = other as _7_Environment__s_Compile.LPacket<IdType,MessageType>;
return oth != null && Dafny.Helpers.AreEqual(this.dst, oth.dst) && Dafny.Helpers.AreEqual(this.src, oth.src) && Dafny.Helpers.AreEqual(this.msg, oth.msg);
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 0;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.dst));
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.src));
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.msg));
return (int) hash;
    }
    public override string ToString() {
      string s = "_7_Environment__s_Compile.LPacket.LPacket";
s += "(";
s += Dafny.Helpers.ToString(this.dst);
s += ", ";
s += Dafny.Helpers.ToString(this.src);
s += ", ";
s += Dafny.Helpers.ToString(this.msg);
s += ")";
return s;
    }
    static LPacket<IdType,MessageType> theDefault;
public static LPacket<IdType,MessageType> Default {
      get {
        if (theDefault == null) {
          theDefault = new _7_Environment__s_Compile.LPacket<IdType,MessageType>(Dafny.Helpers.Default<IdType>(), Dafny.Helpers.Default<IdType>(), Dafny.Helpers.Default<MessageType>());
        }
        return theDefault;
      }
    }
    public static LPacket<IdType,MessageType> _DafnyDefaultValue() { return Default; }
public static LPacket<IdType,MessageType> create(IdType dst, IdType src, MessageType msg) {
      return new LPacket<IdType,MessageType>(dst, src, msg);
    }
    public bool is_LPacket { get { return true; } }
public IdType dtor_dst {
      get {
        return this.dst;
      }
    }
    public IdType dtor_src {
      get {
        return this.src;
      }
    }
    public MessageType dtor_msg {
      get {
        return this.msg;
      }
    }
  }

  public abstract class LIoOp<IdType,MessageType> {
    public LIoOp() { }
static LIoOp<IdType,MessageType> theDefault;
public static LIoOp<IdType,MessageType> Default {
      get {
        if (theDefault == null) {
          theDefault = new _7_Environment__s_Compile.LIoOp_LIoOpTimeoutReceive<IdType,MessageType>();
        }
        return theDefault;
      }
    }
    public static LIoOp<IdType,MessageType> _DafnyDefaultValue() { return Default; }
public static LIoOp<IdType,MessageType> create_LIoOpSend(_7_Environment__s_Compile.LPacket<IdType,MessageType> s) {
      return new LIoOp_LIoOpSend<IdType,MessageType>(s);
    }
    public static LIoOp<IdType,MessageType> create_LIoOpReceive(_7_Environment__s_Compile.LPacket<IdType,MessageType> r) {
      return new LIoOp_LIoOpReceive<IdType,MessageType>(r);
    }
    public static LIoOp<IdType,MessageType> create_LIoOpTimeoutReceive() {
      return new LIoOp_LIoOpTimeoutReceive<IdType,MessageType>();
    }
    public static LIoOp<IdType,MessageType> create_LIoOpReadClock(BigInteger t) {
      return new LIoOp_LIoOpReadClock<IdType,MessageType>(t);
    }
    public bool is_LIoOpSend { get { return this is LIoOp_LIoOpSend<IdType,MessageType>; } }
public bool is_LIoOpReceive { get { return this is LIoOp_LIoOpReceive<IdType,MessageType>; } }
public bool is_LIoOpTimeoutReceive { get { return this is LIoOp_LIoOpTimeoutReceive<IdType,MessageType>; } }
public bool is_LIoOpReadClock { get { return this is LIoOp_LIoOpReadClock<IdType,MessageType>; } }
public _7_Environment__s_Compile.LPacket<IdType,MessageType> dtor_s {
      get {
        var d = this;
return ((LIoOp_LIoOpSend<IdType,MessageType>)d).s; 
      }
    }
    public _7_Environment__s_Compile.LPacket<IdType,MessageType> dtor_r {
      get {
        var d = this;
return ((LIoOp_LIoOpReceive<IdType,MessageType>)d).r; 
      }
    }
    public BigInteger dtor_t {
      get {
        var d = this;
return ((LIoOp_LIoOpReadClock<IdType,MessageType>)d).t; 
      }
    }
  }
  public class LIoOp_LIoOpSend<IdType,MessageType> : LIoOp<IdType,MessageType> {
    public readonly _7_Environment__s_Compile.LPacket<IdType,MessageType> s;
public LIoOp_LIoOpSend(_7_Environment__s_Compile.LPacket<IdType,MessageType> s) {
      this.s = s;
    }
    public override bool Equals(object other) {
      var oth = other as _7_Environment__s_Compile.LIoOp_LIoOpSend<IdType,MessageType>;
return oth != null && Dafny.Helpers.AreEqual(this.s, oth.s);
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 0;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.s));
return (int) hash;
    }
    public override string ToString() {
      string ss = "_7_Environment__s_Compile.LIoOp.LIoOpSend";
ss += "(";
ss += Dafny.Helpers.ToString(this.s);
ss += ")";
return ss;
    }
  }
  public class LIoOp_LIoOpReceive<IdType,MessageType> : LIoOp<IdType,MessageType> {
    public readonly _7_Environment__s_Compile.LPacket<IdType,MessageType> r;
public LIoOp_LIoOpReceive(_7_Environment__s_Compile.LPacket<IdType,MessageType> r) {
      this.r = r;
    }
    public override bool Equals(object other) {
      var oth = other as _7_Environment__s_Compile.LIoOp_LIoOpReceive<IdType,MessageType>;
return oth != null && Dafny.Helpers.AreEqual(this.r, oth.r);
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 1;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.r));
return (int) hash;
    }
    public override string ToString() {
      string s = "_7_Environment__s_Compile.LIoOp.LIoOpReceive";
s += "(";
s += Dafny.Helpers.ToString(this.r);
s += ")";
return s;
    }
  }
  public class LIoOp_LIoOpTimeoutReceive<IdType,MessageType> : LIoOp<IdType,MessageType> {
    public LIoOp_LIoOpTimeoutReceive() {
    }
    public override bool Equals(object other) {
      var oth = other as _7_Environment__s_Compile.LIoOp_LIoOpTimeoutReceive<IdType,MessageType>;
return oth != null;
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 2;
return (int) hash;
    }
    public override string ToString() {
      string s = "_7_Environment__s_Compile.LIoOp.LIoOpTimeoutReceive";
return s;
    }
  }
  public class LIoOp_LIoOpReadClock<IdType,MessageType> : LIoOp<IdType,MessageType> {
    public readonly BigInteger t;
public LIoOp_LIoOpReadClock(BigInteger t) {
      this.t = t;
    }
    public override bool Equals(object other) {
      var oth = other as _7_Environment__s_Compile.LIoOp_LIoOpReadClock<IdType,MessageType>;
return oth != null && this.t == oth.t;
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 3;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.t));
return (int) hash;
    }
    public override string ToString() {
      string s = "_7_Environment__s_Compile.LIoOp.LIoOpReadClock";
s += "(";
s += Dafny.Helpers.ToString(this.t);
s += ")";
return s;
    }
  }

  public abstract class LEnvStep<IdType,MessageType,NodeStepType> {
    public LEnvStep() { }
static LEnvStep<IdType,MessageType,NodeStepType> theDefault;
public static LEnvStep<IdType,MessageType,NodeStepType> Default {
      get {
        if (theDefault == null) {
          theDefault = new _7_Environment__s_Compile.LEnvStep_LEnvStepAdvanceTime<IdType,MessageType,NodeStepType>();
        }
        return theDefault;
      }
    }
    public static LEnvStep<IdType,MessageType,NodeStepType> _DafnyDefaultValue() { return Default; }
public static LEnvStep<IdType,MessageType,NodeStepType> create_LEnvStepHostIos(IdType actor, Dafny.Sequence<_7_Environment__s_Compile.LIoOp<IdType,MessageType>> ios, NodeStepType nodeStep) {
      return new LEnvStep_LEnvStepHostIos<IdType,MessageType,NodeStepType>(actor, ios, nodeStep);
    }
    public static LEnvStep<IdType,MessageType,NodeStepType> create_LEnvStepDeliverPacket(_7_Environment__s_Compile.LPacket<IdType,MessageType> p) {
      return new LEnvStep_LEnvStepDeliverPacket<IdType,MessageType,NodeStepType>(p);
    }
    public static LEnvStep<IdType,MessageType,NodeStepType> create_LEnvStepAdvanceTime() {
      return new LEnvStep_LEnvStepAdvanceTime<IdType,MessageType,NodeStepType>();
    }
    public static LEnvStep<IdType,MessageType,NodeStepType> create_LEnvStepStutter() {
      return new LEnvStep_LEnvStepStutter<IdType,MessageType,NodeStepType>();
    }
    public bool is_LEnvStepHostIos { get { return this is LEnvStep_LEnvStepHostIos<IdType,MessageType,NodeStepType>; } }
public bool is_LEnvStepDeliverPacket { get { return this is LEnvStep_LEnvStepDeliverPacket<IdType,MessageType,NodeStepType>; } }
public bool is_LEnvStepAdvanceTime { get { return this is LEnvStep_LEnvStepAdvanceTime<IdType,MessageType,NodeStepType>; } }
public bool is_LEnvStepStutter { get { return this is LEnvStep_LEnvStepStutter<IdType,MessageType,NodeStepType>; } }
public IdType dtor_actor {
      get {
        var d = this;
return ((LEnvStep_LEnvStepHostIos<IdType,MessageType,NodeStepType>)d).actor; 
      }
    }
    public Dafny.Sequence<_7_Environment__s_Compile.LIoOp<IdType,MessageType>> dtor_ios {
      get {
        var d = this;
return ((LEnvStep_LEnvStepHostIos<IdType,MessageType,NodeStepType>)d).ios; 
      }
    }
    public NodeStepType dtor_nodeStep {
      get {
        var d = this;
return ((LEnvStep_LEnvStepHostIos<IdType,MessageType,NodeStepType>)d).nodeStep; 
      }
    }
    public _7_Environment__s_Compile.LPacket<IdType,MessageType> dtor_p {
      get {
        var d = this;
return ((LEnvStep_LEnvStepDeliverPacket<IdType,MessageType,NodeStepType>)d).p; 
      }
    }
  }
  public class LEnvStep_LEnvStepHostIos<IdType,MessageType,NodeStepType> : LEnvStep<IdType,MessageType,NodeStepType> {
    public readonly IdType actor;
public readonly Dafny.Sequence<_7_Environment__s_Compile.LIoOp<IdType,MessageType>> ios;
public readonly NodeStepType nodeStep;
public LEnvStep_LEnvStepHostIos(IdType actor, Dafny.Sequence<_7_Environment__s_Compile.LIoOp<IdType,MessageType>> ios, NodeStepType nodeStep) {
      this.actor = actor;
this.ios = ios;
this.nodeStep = nodeStep;
    }
    public override bool Equals(object other) {
      var oth = other as _7_Environment__s_Compile.LEnvStep_LEnvStepHostIos<IdType,MessageType,NodeStepType>;
return oth != null && Dafny.Helpers.AreEqual(this.actor, oth.actor) && Dafny.Helpers.AreEqual(this.ios, oth.ios) && Dafny.Helpers.AreEqual(this.nodeStep, oth.nodeStep);
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 0;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.actor));
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.ios));
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.nodeStep));
return (int) hash;
    }
    public override string ToString() {
      string s = "_7_Environment__s_Compile.LEnvStep.LEnvStepHostIos";
s += "(";
s += Dafny.Helpers.ToString(this.actor);
s += ", ";
s += Dafny.Helpers.ToString(this.ios);
s += ", ";
s += Dafny.Helpers.ToString(this.nodeStep);
s += ")";
return s;
    }
  }
  public class LEnvStep_LEnvStepDeliverPacket<IdType,MessageType,NodeStepType> : LEnvStep<IdType,MessageType,NodeStepType> {
    public readonly _7_Environment__s_Compile.LPacket<IdType,MessageType> p;
public LEnvStep_LEnvStepDeliverPacket(_7_Environment__s_Compile.LPacket<IdType,MessageType> p) {
      this.p = p;
    }
    public override bool Equals(object other) {
      var oth = other as _7_Environment__s_Compile.LEnvStep_LEnvStepDeliverPacket<IdType,MessageType,NodeStepType>;
return oth != null && Dafny.Helpers.AreEqual(this.p, oth.p);
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 1;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.p));
return (int) hash;
    }
    public override string ToString() {
      string s = "_7_Environment__s_Compile.LEnvStep.LEnvStepDeliverPacket";
s += "(";
s += Dafny.Helpers.ToString(this.p);
s += ")";
return s;
    }
  }
  public class LEnvStep_LEnvStepAdvanceTime<IdType,MessageType,NodeStepType> : LEnvStep<IdType,MessageType,NodeStepType> {
    public LEnvStep_LEnvStepAdvanceTime() {
    }
    public override bool Equals(object other) {
      var oth = other as _7_Environment__s_Compile.LEnvStep_LEnvStepAdvanceTime<IdType,MessageType,NodeStepType>;
return oth != null;
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 2;
return (int) hash;
    }
    public override string ToString() {
      string s = "_7_Environment__s_Compile.LEnvStep.LEnvStepAdvanceTime";
return s;
    }
  }
  public class LEnvStep_LEnvStepStutter<IdType,MessageType,NodeStepType> : LEnvStep<IdType,MessageType,NodeStepType> {
    public LEnvStep_LEnvStepStutter() {
    }
    public override bool Equals(object other) {
      var oth = other as _7_Environment__s_Compile.LEnvStep_LEnvStepStutter<IdType,MessageType,NodeStepType>;
return oth != null;
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 3;
return (int) hash;
    }
    public override string ToString() {
      string s = "_7_Environment__s_Compile.LEnvStep.LEnvStepStutter";
return s;
    }
  }

  public class LHostInfo<IdType,MessageType> {
    public readonly Dafny.Sequence<_7_Environment__s_Compile.LPacket<IdType,MessageType>> queue;
public LHostInfo(Dafny.Sequence<_7_Environment__s_Compile.LPacket<IdType,MessageType>> queue) {
      this.queue = queue;
    }
    public override bool Equals(object other) {
      var oth = other as _7_Environment__s_Compile.LHostInfo<IdType,MessageType>;
return oth != null && Dafny.Helpers.AreEqual(this.queue, oth.queue);
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 0;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.queue));
return (int) hash;
    }
    public override string ToString() {
      string s = "_7_Environment__s_Compile.LHostInfo.LHostInfo";
s += "(";
s += Dafny.Helpers.ToString(this.queue);
s += ")";
return s;
    }
    static LHostInfo<IdType,MessageType> theDefault;
public static LHostInfo<IdType,MessageType> Default {
      get {
        if (theDefault == null) {
          theDefault = new _7_Environment__s_Compile.LHostInfo<IdType,MessageType>(Dafny.Sequence<_7_Environment__s_Compile.LPacket<IdType,MessageType>>.Empty);
        }
        return theDefault;
      }
    }
    public static LHostInfo<IdType,MessageType> _DafnyDefaultValue() { return Default; }
public static LHostInfo<IdType,MessageType> create(Dafny.Sequence<_7_Environment__s_Compile.LPacket<IdType,MessageType>> queue) {
      return new LHostInfo<IdType,MessageType>(queue);
    }
    public bool is_LHostInfo { get { return true; } }
public Dafny.Sequence<_7_Environment__s_Compile.LPacket<IdType,MessageType>> dtor_queue {
      get {
        return this.queue;
      }
    }
  }

  public class LEnvironment<IdType,MessageType,NodeStepType> {
    public readonly BigInteger time;
public readonly Dafny.Set<_7_Environment__s_Compile.LPacket<IdType,MessageType>> sentPackets;
public readonly Dafny.Map<IdType,_7_Environment__s_Compile.LHostInfo<IdType,MessageType>> hostInfo;
public readonly _7_Environment__s_Compile.LEnvStep<IdType,MessageType,NodeStepType> nextStep;
public LEnvironment(BigInteger time, Dafny.Set<_7_Environment__s_Compile.LPacket<IdType,MessageType>> sentPackets, Dafny.Map<IdType,_7_Environment__s_Compile.LHostInfo<IdType,MessageType>> hostInfo, _7_Environment__s_Compile.LEnvStep<IdType,MessageType,NodeStepType> nextStep) {
      this.time = time;
this.sentPackets = sentPackets;
this.hostInfo = hostInfo;
this.nextStep = nextStep;
    }
    public override bool Equals(object other) {
      var oth = other as _7_Environment__s_Compile.LEnvironment<IdType,MessageType,NodeStepType>;
return oth != null && this.time == oth.time && Dafny.Helpers.AreEqual(this.sentPackets, oth.sentPackets) && Dafny.Helpers.AreEqual(this.hostInfo, oth.hostInfo) && Dafny.Helpers.AreEqual(this.nextStep, oth.nextStep);
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 0;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.time));
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.sentPackets));
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.hostInfo));
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.nextStep));
return (int) hash;
    }
    public override string ToString() {
      string s = "_7_Environment__s_Compile.LEnvironment.LEnvironment";
s += "(";
s += Dafny.Helpers.ToString(this.time);
s += ", ";
s += Dafny.Helpers.ToString(this.sentPackets);
s += ", ";
s += Dafny.Helpers.ToString(this.hostInfo);
s += ", ";
s += Dafny.Helpers.ToString(this.nextStep);
s += ")";
return s;
    }
    static LEnvironment<IdType,MessageType,NodeStepType> theDefault;
public static LEnvironment<IdType,MessageType,NodeStepType> Default {
      get {
        if (theDefault == null) {
          theDefault = new _7_Environment__s_Compile.LEnvironment<IdType,MessageType,NodeStepType>(BigInteger.Zero, Dafny.Set<_7_Environment__s_Compile.LPacket<IdType,MessageType>>.Empty, Dafny.Map<IdType,_7_Environment__s_Compile.LHostInfo<IdType,MessageType>>.Empty, @_7_Environment__s_Compile.LEnvStep<IdType,MessageType,NodeStepType>.Default);
        }
        return theDefault;
      }
    }
    public static LEnvironment<IdType,MessageType,NodeStepType> _DafnyDefaultValue() { return Default; }
public static LEnvironment<IdType,MessageType,NodeStepType> create(BigInteger time, Dafny.Set<_7_Environment__s_Compile.LPacket<IdType,MessageType>> sentPackets, Dafny.Map<IdType,_7_Environment__s_Compile.LHostInfo<IdType,MessageType>> hostInfo, _7_Environment__s_Compile.LEnvStep<IdType,MessageType,NodeStepType> nextStep) {
      return new LEnvironment<IdType,MessageType,NodeStepType>(time, sentPackets, hostInfo, nextStep);
    }
    public bool is_LEnvironment { get { return true; } }
public BigInteger dtor_time {
      get {
        return this.time;
      }
    }
    public Dafny.Set<_7_Environment__s_Compile.LPacket<IdType,MessageType>> dtor_sentPackets {
      get {
        return this.sentPackets;
      }
    }
    public Dafny.Map<IdType,_7_Environment__s_Compile.LHostInfo<IdType,MessageType>> dtor_hostInfo {
      get {
        return this.hostInfo;
      }
    }
    public _7_Environment__s_Compile.LEnvStep<IdType,MessageType,NodeStepType> dtor_nextStep {
      get {
        return this.nextStep;
      }
    }
  }

} // end of namespace _7_Environment__s_Compile

// TONY: Put IoNative.cs contents here
namespace _9_Native____Io__s_Compile {

  public struct Packet {
    public IEndPoint ep;
    public byte[] buffer;
  }

  public partial class HostEnvironment {
  }

  public partial class HostConstants
  {
      public static void NumCommandLineArgs(out uint n)
      {
          n = (uint)System.Environment.GetCommandLineArgs().Length;
      }

      public static void GetCommandLineArg(ulong i, out ushort[] arg)
      {
          arg = Array.ConvertAll(System.Environment.GetCommandLineArgs()[i].ToCharArray(), c => (ushort)c);
      }
  }

  public partial class OkState {
  }

  public partial class NowState {
  }

  public partial class Time
{
    static Stopwatch watch;

    public static void Initialize()
    {
        watch = new Stopwatch();
        watch.Start();
    }

    public static void GetTime(out ulong time)
    {
        time = (ulong) DateTime.Now.Ticks / 10000;
    }
    
    public static void GetDebugTimeTicks(out ulong time)
    {
        time = (ulong) watch.ElapsedTicks;
    }
    
    public static void RecordTiming(char[] name, ulong time)
    {
        var str = new string(name);
        // Common.Profiler.Record(str, (long)time);
    }
}

  public class EndPoint {
    public readonly Dafny.Sequence<byte> addr;
public readonly ushort port;
public EndPoint(Dafny.Sequence<byte> addr, ushort port) {
      this.addr = addr;
this.port = port;
    }
    public override bool Equals(object other) {
      var oth = other as _9_Native____Io__s_Compile.EndPoint;
return oth != null && Dafny.Helpers.AreEqual(this.addr, oth.addr) && this.port == oth.port;
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 0;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.addr));
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.port));
return (int) hash;
    }
    public override string ToString() {
      string s = "_9_Native____Io__s_Compile.EndPoint.EndPoint";
s += "(";
s += Dafny.Helpers.ToString(this.addr);
s += ", ";
s += Dafny.Helpers.ToString(this.port);
s += ")";
return s;
    }
    static EndPoint theDefault;
public static EndPoint Default {
      get {
        if (theDefault == null) {
          theDefault = new _9_Native____Io__s_Compile.EndPoint(Dafny.Sequence<byte>.Empty, 0);
        }
        return theDefault;
      }
    }
    public static EndPoint _DafnyDefaultValue() { return Default; }
public static EndPoint create(Dafny.Sequence<byte> addr, ushort port) {
      return new EndPoint(addr, port);
    }
    public bool is_EndPoint { get { return true; } }
public Dafny.Sequence<byte> dtor_addr {
      get {
        return this.addr;
      }
    }
    public ushort dtor_port {
      get {
        return this.port;
      }
    }
  }



  public partial class UdpState {
  }

  public partial class IPEndPoint
  {
      internal IEndPoint endpoint;
      internal IPEndPoint(IEndPoint endpoint) { this.endpoint = endpoint; }

      public void GetAddress(out byte[] addr)
      {
          // no exceptions thrown:
          addr = (byte[])(endpoint.Address.GetAddressBytes().Clone());
      }

      public ushort GetPort()
      {
          // no exceptions thrown:
          return (ushort)endpoint.Port;
      }

      public static void Construct(byte[] ipAddress, ushort port, out bool ok, out IPEndPoint endpoint)
      {
          try
          {
              ipAddress = (byte[])(ipAddress.Clone());
              endpoint = new IPEndPoint(new IEndPoint(new System.Net.IPAddress(ipAddress), port));
              ok = true;
          }
          catch (Exception e)
          {
              System.Console.Error.WriteLine(e);
              endpoint = null;
              ok = false;
          }
      }
  }

  public partial class UdpClient {
    internal UClient client;
    internal Thread sender;
    internal Thread receiver;
    internal ConcurrentQueue<Packet> send_queue;
    internal ConcurrentQueue<Packet> receive_queue;

    internal UdpClient(UClient client) { 
      this.client = client;
      this.send_queue = new ConcurrentQueue<Packet>();
      this.receive_queue = new ConcurrentQueue<Packet>();
      this.sender = new Thread(SendLoop);
      this.sender.Start();
      this.receiver = new Thread(ReceiveLoop);
      this.receiver.Start();
    }

    // TODO: remove this
    public static void ConstructDeprecated(bool useIPv6, ushort port, out bool ok, out UdpClient udp)
    {
        try
        {
            var family = useIPv6 ? System.Net.Sockets.AddressFamily.InterNetworkV6 : System.Net.Sockets.AddressFamily.InterNetwork;
            udp = new UdpClient(new UClient(port, family));
            uint SIO_UDP_CONNRESET = 0x9800000C; // suppress UDP "connection" closed exceptions, since UDP is connectionless
            udp.client.Client.IOControl((System.Net.Sockets.IOControlCode)SIO_UDP_CONNRESET, new byte[] { 0 }, new byte[0]);
            ok = true;
        }
        catch (Exception e)
        {
            System.Console.Error.WriteLine(e);
            udp = null;
            ok = false;
        }
    }

    public static void Construct(IPEndPoint localEP, out bool ok, out UdpClient udp)
    {
        try
        {
            udp = new UdpClient(new UClient(localEP.endpoint));
            uint SIO_UDP_CONNRESET = 0x9800000C; // suppress UDP "connection" closed exceptions, since UDP is connectionless
            udp.client.Client.IOControl((System.Net.Sockets.IOControlCode)SIO_UDP_CONNRESET, new byte[] { 0 }, new byte[0]);
            udp.client.Client.ReceiveBufferSize = 8192 * 100;
            ok = true;
        }
        catch (Exception e)
        {
            System.Console.Error.WriteLine(e);
            udp = null;
            ok = false;
        }
    }

    public void Close(out bool ok)
    {
        try
        {
            client.Close();
            ok = true;
        }
        catch (Exception e)
        {
            System.Console.Error.WriteLine(e);
            ok = false;
        }
    }

    public void Receive(int timeLimit, out bool ok, out bool timedOut, out IPEndPoint remote, out byte[] buffer)
    {
        buffer = null;
        remote = null;
        try
        {
            Packet packet;
            bool dequeued = this.receive_queue.TryDequeue(out packet);
            if (!dequeued) {
                if (timeLimit == 0) {
                    ok = true;
                    timedOut = true;
                    return;
                } else {
                    System.Console.Out.WriteLine("Going to sleep unexpectedly!");
                    Thread.Sleep(timeLimit);  // REVIEW: This is very conservative, but shouldn't matter, since we don't use this path
                    Receive(0, out ok, out timedOut, out remote, out buffer);
                }
            } else {
                //System.Console.Out.WriteLine("Dequeued a packet from: " + packet.ep.Address);
                timedOut = false;
                remote = new IPEndPoint(packet.ep);
                buffer = new byte[packet.buffer.Length];
                Array.Copy(packet.buffer, buffer, packet.buffer.Length);
                ok = true;
            }     
        }
        catch (Exception e)
        {
            System.Console.Error.WriteLine(e);
            timedOut = false;
            ok = false;
        }
    }

    public void ReceiveLoop() {
        while (true) {
            try {
                Packet packet = new Packet();
                packet.buffer = client.Receive(ref packet.ep);
                this.receive_queue.Enqueue(packet);
                //System.Console.Out.WriteLine("Enqueued a packet from: " + packet.ep.Address);
            } catch (Exception e) {
                System.Console.Error.WriteLine(e);
            }
        }
    }

    public void SendLoop() {
        while (true) {
            try {
                Packet packet;
                bool dequeued = this.send_queue.TryDequeue(out packet);
                if (dequeued) {                
                      int nSent = client.Send(packet.buffer, packet.buffer.Length, packet.ep);
                      if (nSent != packet.buffer.Length) {
                          //throw new Exception("only sent " + nSent + " of " + packet.buffer.Length + " bytes");
                          System.Console.Error.Write("only sent " + nSent + " of " + packet.buffer.Length + " bytes");
                      }                
                }
            } catch (Exception e) {
              System.Console.Error.WriteLine(e);
            }
        }
    }

    public void Send(IPEndPoint remote, byte[] buffer, out bool ok)
    {
        Packet p = new Packet();
        p.ep = remote.endpoint;
        p.buffer = new byte[buffer.Length];
        Array.Copy(buffer, p.buffer, buffer.Length);
        this.send_queue.Enqueue(p);
        ok = true;
    }

    public void __ctor()
    {
      var _this = this;
    TAIL_CALL_START: ;
    }
  }

  public partial class FileSystemState {
  }

  public partial class MutableSet<T>
  {
      private HashSet<T> setImpl;
      public MutableSet() {
          this.setImpl = new HashSet<T>();
      }

      public static Dafny.Set<T> SetOf(MutableSet<T> s) { return Dafny.Set<T>.FromCollection(s.setImpl); }

      public static void EmptySet(out MutableSet<T> s) { s = new MutableSet<T>(); }

      public BigInteger Size() { return new BigInteger(this.setImpl.Count); }
      
      public void SizeModest(out ulong size) { size = (ulong)this.setImpl.Count; }

      public void Contains(T x, out bool b) { b = this.setImpl.Contains(x); }

      public void Add(T x) { this.setImpl.Add(x); }
             
      public void AddSet(MutableSet<T> s) { this.setImpl.UnionWith(s.setImpl); }

      public void TransferSet(MutableSet<T> s) { this.setImpl = s.setImpl; s.setImpl = new HashSet<T>(); }
             
      public void Remove(T x) { this.setImpl.Remove(x); }

      public void RemoveAll() { this.setImpl.Clear(); }
  }

  public partial class MutableMap<K,V>
  {
    private Dictionary<K,V> mapImpl;

    public MutableMap() {
        this.mapImpl = new Dictionary<K, V>();
    }

    // TODO: This is pretty inefficient.  Should change Dafny's interface to allow us to 
    // pass in an enumerable or an ImmutableDictionary
    public static Dafny.Map<K,V> MapOf(MutableMap<K,V> s) {
      List<Dafny.Pair<K, V>> pairs = new List<Dafny.Pair<K, V>>();
      foreach (var pair in s.mapImpl) {
        pairs.Add(new Dafny.Pair<K, V>(pair.Key, pair.Value));
      }
      return Dafny.Map<K,V>.FromCollection(pairs); 
    }

    public static void EmptyMap(out MutableMap<K,V> m) { m = new MutableMap<K,V>(); }

    // TONY: manually commented this out.
    // public static void FromMap(Dafny.Map<K, V> m, out MutableMap<K, V> new_m) {
    //   new_m = new MutableMap<K,V>();
    //   foreach (var key in m.Domain) {
    //     new_m.mapImpl.Add(key, m.Select(key));
    //   }
    // }

    public BigInteger Size() { return new BigInteger(this.mapImpl.Count); }

    public void SizeModest(out ulong size) { size = (ulong)this.mapImpl.Count; }

    public bool Contains(K key) { return this.mapImpl.ContainsKey(key); }

    public void TryGetValue(K key, out bool contains, out V val) {
      contains = this.mapImpl.TryGetValue(key, out val);
    }

    public void Set(K key, V val) { this.mapImpl[key] = val; }
           
    //public void AddMap(MutableMap<K,V> s) { this.mapImpl.}

    public void Remove(K key) { this.mapImpl.Remove(key); }
  }

  public partial class @Arrays
  {
    public static void @CopySeqIntoArray<A>(Dafny.Sequence<A> src, ulong srcIndex, A[] dst, ulong dstIndex, ulong len) {
        System.Array.Copy(src.Elements, (long)srcIndex, dst, (long)dstIndex, (long)len);
    }
  }

} // end of namespace _9_Native____Io__s_Compile
//TONY: END

namespace _26_Collections____Seqs__s_Compile {

} // end of namespace _26_Collections____Seqs__s_Compile
namespace _29_Collections____Sets__i_Compile {

} // end of namespace _29_Collections____Sets__i_Compile
namespace _33_Types__i_Compile {



  public abstract class LockMessage {
    public LockMessage() { }
static LockMessage theDefault;
public static LockMessage Default {
      get {
        if (theDefault == null) {
          theDefault = new _33_Types__i_Compile.LockMessage_Transfer(BigInteger.Zero);
        }
        return theDefault;
      }
    }
    public static LockMessage _DafnyDefaultValue() { return Default; }
public static LockMessage create_Transfer(BigInteger transfer__epoch) {
      return new LockMessage_Transfer(transfer__epoch);
    }
    public static LockMessage create_Locked(BigInteger locked__epoch) {
      return new LockMessage_Locked(locked__epoch);
    }
    public static LockMessage create_Invalid() {
      return new LockMessage_Invalid();
    }
    public bool is_Transfer { get { return this is LockMessage_Transfer; } }
public bool is_Locked { get { return this is LockMessage_Locked; } }
public bool is_Invalid { get { return this is LockMessage_Invalid; } }
public BigInteger dtor_transfer__epoch {
      get {
        var d = this;
return ((LockMessage_Transfer)d).transfer__epoch; 
      }
    }
    public BigInteger dtor_locked__epoch {
      get {
        var d = this;
return ((LockMessage_Locked)d).locked__epoch; 
      }
    }
  }
  public class LockMessage_Transfer : LockMessage {
    public readonly BigInteger transfer__epoch;
public LockMessage_Transfer(BigInteger transfer__epoch) {
      this.transfer__epoch = transfer__epoch;
    }
    public override bool Equals(object other) {
      var oth = other as _33_Types__i_Compile.LockMessage_Transfer;
return oth != null && this.transfer__epoch == oth.transfer__epoch;
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 0;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.transfer__epoch));
return (int) hash;
    }
    public override string ToString() {
      string s = "_33_Types__i_Compile.LockMessage.Transfer";
s += "(";
s += Dafny.Helpers.ToString(this.transfer__epoch);
s += ")";
return s;
    }
  }
  public class LockMessage_Locked : LockMessage {
    public readonly BigInteger locked__epoch;
public LockMessage_Locked(BigInteger locked__epoch) {
      this.locked__epoch = locked__epoch;
    }
    public override bool Equals(object other) {
      var oth = other as _33_Types__i_Compile.LockMessage_Locked;
return oth != null && this.locked__epoch == oth.locked__epoch;
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 1;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.locked__epoch));
return (int) hash;
    }
    public override string ToString() {
      string s = "_33_Types__i_Compile.LockMessage.Locked";
s += "(";
s += Dafny.Helpers.ToString(this.locked__epoch);
s += ")";
return s;
    }
  }
  public class LockMessage_Invalid : LockMessage {
    public LockMessage_Invalid() {
    }
    public override bool Equals(object other) {
      var oth = other as _33_Types__i_Compile.LockMessage_Invalid;
return oth != null;
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 2;
return (int) hash;
    }
    public override string ToString() {
      string s = "_33_Types__i_Compile.LockMessage.Invalid";
return s;
    }
  }

  public abstract class LockStep {
    public LockStep() { }
static LockStep theDefault;
public static LockStep Default {
      get {
        if (theDefault == null) {
          theDefault = new _33_Types__i_Compile.LockStep_GrantStep();
        }
        return theDefault;
      }
    }
    public static LockStep _DafnyDefaultValue() { return Default; }
public static LockStep create_GrantStep() {
      return new LockStep_GrantStep();
    }
    public static LockStep create_AcceptStep() {
      return new LockStep_AcceptStep();
    }
    public bool is_GrantStep { get { return this is LockStep_GrantStep; } }
public bool is_AcceptStep { get { return this is LockStep_AcceptStep; } }
public static System.Collections.Generic.IEnumerable<LockStep> AllSingletonConstructors {
      get {
        yield return LockStep.create_GrantStep();
yield return LockStep.create_AcceptStep();
      }
    }
  }
  public class LockStep_GrantStep : LockStep {
    public LockStep_GrantStep() {
    }
    public override bool Equals(object other) {
      var oth = other as _33_Types__i_Compile.LockStep_GrantStep;
return oth != null;
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 0;
return (int) hash;
    }
    public override string ToString() {
      string s = "_33_Types__i_Compile.LockStep.GrantStep";
return s;
    }
  }
  public class LockStep_AcceptStep : LockStep {
    public LockStep_AcceptStep() {
    }
    public override bool Equals(object other) {
      var oth = other as _33_Types__i_Compile.LockStep_AcceptStep;
return oth != null;
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 1;
return (int) hash;
    }
    public override string ToString() {
      string s = "_33_Types__i_Compile.LockStep.AcceptStep";
return s;
    }
  }




} // end of namespace _33_Types__i_Compile
namespace _36_Protocol__Node__i_Compile {




  public class Node {
    public readonly bool held;
public readonly BigInteger epoch;
public readonly BigInteger my__index;
public readonly Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> config;
public Node(bool held, BigInteger epoch, BigInteger my__index, Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> config) {
      this.held = held;
this.epoch = epoch;
this.my__index = my__index;
this.config = config;
    }
    public override bool Equals(object other) {
      var oth = other as _36_Protocol__Node__i_Compile.Node;
return oth != null && this.held == oth.held && this.epoch == oth.epoch && this.my__index == oth.my__index && Dafny.Helpers.AreEqual(this.config, oth.config);
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 0;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.held));
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.epoch));
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.my__index));
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.config));
return (int) hash;
    }
    public override string ToString() {
      string s = "_36_Protocol__Node__i_Compile.Node.Node";
s += "(";
s += Dafny.Helpers.ToString(this.held);
s += ", ";
s += Dafny.Helpers.ToString(this.epoch);
s += ", ";
s += Dafny.Helpers.ToString(this.my__index);
s += ", ";
s += Dafny.Helpers.ToString(this.config);
s += ")";
return s;
    }
    static Node theDefault;
public static Node Default {
      get {
        if (theDefault == null) {
          theDefault = new _36_Protocol__Node__i_Compile.Node(false, BigInteger.Zero, BigInteger.Zero, Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint>.Empty);
        }
        return theDefault;
      }
    }
    public static Node _DafnyDefaultValue() { return Default; }
public static Node create(bool held, BigInteger epoch, BigInteger my__index, Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> config) {
      return new Node(held, epoch, my__index, config);
    }
    public bool is_Node { get { return true; } }
public bool dtor_held {
      get {
        return this.held;
      }
    }
    public BigInteger dtor_epoch {
      get {
        return this.epoch;
      }
    }
    public BigInteger dtor_my__index {
      get {
        return this.my__index;
      }
    }
    public Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> dtor_config {
      get {
        return this.config;
      }
    }
  }

} // end of namespace _36_Protocol__Node__i_Compile
namespace _39_Message__i_Compile {


  public abstract class CMessage {
    public CMessage() { }
static CMessage theDefault;
public static CMessage Default {
      get {
        if (theDefault == null) {
          theDefault = new _39_Message__i_Compile.CMessage_CTransfer(0);
        }
        return theDefault;
      }
    }
    public static CMessage _DafnyDefaultValue() { return Default; }
public static CMessage create_CTransfer(ulong transfer__epoch) {
      return new CMessage_CTransfer(transfer__epoch);
    }
    public static CMessage create_CLocked(ulong locked__epoch) {
      return new CMessage_CLocked(locked__epoch);
    }
    public static CMessage create_CInvalid() {
      return new CMessage_CInvalid();
    }
    public bool is_CTransfer { get { return this is CMessage_CTransfer; } }
public bool is_CLocked { get { return this is CMessage_CLocked; } }
public bool is_CInvalid { get { return this is CMessage_CInvalid; } }
public ulong dtor_transfer__epoch {
      get {
        var d = this;
return ((CMessage_CTransfer)d).transfer__epoch; 
      }
    }
    public ulong dtor_locked__epoch {
      get {
        var d = this;
return ((CMessage_CLocked)d).locked__epoch; 
      }
    }
  }
  public class CMessage_CTransfer : CMessage {
    public readonly ulong transfer__epoch;
public CMessage_CTransfer(ulong transfer__epoch) {
      this.transfer__epoch = transfer__epoch;
    }
    public override bool Equals(object other) {
      var oth = other as _39_Message__i_Compile.CMessage_CTransfer;
return oth != null && this.transfer__epoch == oth.transfer__epoch;
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 0;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.transfer__epoch));
return (int) hash;
    }
    public override string ToString() {
      string s = "_39_Message__i_Compile.CMessage.CTransfer";
s += "(";
s += Dafny.Helpers.ToString(this.transfer__epoch);
s += ")";
return s;
    }
  }
  public class CMessage_CLocked : CMessage {
    public readonly ulong locked__epoch;
public CMessage_CLocked(ulong locked__epoch) {
      this.locked__epoch = locked__epoch;
    }
    public override bool Equals(object other) {
      var oth = other as _39_Message__i_Compile.CMessage_CLocked;
return oth != null && this.locked__epoch == oth.locked__epoch;
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 1;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.locked__epoch));
return (int) hash;
    }
    public override string ToString() {
      string s = "_39_Message__i_Compile.CMessage.CLocked";
s += "(";
s += Dafny.Helpers.ToString(this.locked__epoch);
s += ")";
return s;
    }
  }
  public class CMessage_CInvalid : CMessage {
    public CMessage_CInvalid() {
    }
    public override bool Equals(object other) {
      var oth = other as _39_Message__i_Compile.CMessage_CInvalid;
return oth != null;
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 2;
return (int) hash;
    }
    public override string ToString() {
      string s = "_39_Message__i_Compile.CMessage.CInvalid";
return s;
    }
  }


} // end of namespace _39_Message__i_Compile
namespace _42_Common____UdpClient__i_Compile {


  public partial class __default {
    public static bool EndPointIsValidIPV4(_9_Native____Io__s_Compile.EndPoint endPoint) {
      return ((new BigInteger(((endPoint).dtor_addr).Count)) == (new BigInteger(4))) && (((0) <= ((endPoint).dtor_port)) && (((endPoint).dtor_port) <= (65535)));
    }
  }
} // end of namespace _42_Common____UdpClient__i_Compile
namespace _44_Logic____Option__i_Compile {

  public abstract class Option<T> {
    public Option() { }
static Option<T> theDefault;
public static Option<T> Default {
      get {
        if (theDefault == null) {
          theDefault = new _44_Logic____Option__i_Compile.Option_None<T>();
        }
        return theDefault;
      }
    }
    public static Option<T> _DafnyDefaultValue() { return Default; }
public static Option<T> create_Some(T v) {
      return new Option_Some<T>(v);
    }
    public static Option<T> create_None() {
      return new Option_None<T>();
    }
    public bool is_Some { get { return this is Option_Some<T>; } }
public bool is_None { get { return this is Option_None<T>; } }
public T dtor_v {
      get {
        var d = this;
return ((Option_Some<T>)d).v; 
      }
    }
  }
  public class Option_Some<T> : Option<T> {
    public readonly T v;
public Option_Some(T v) {
      this.v = v;
    }
    public override bool Equals(object other) {
      var oth = other as _44_Logic____Option__i_Compile.Option_Some<T>;
return oth != null && Dafny.Helpers.AreEqual(this.v, oth.v);
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 0;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.v));
return (int) hash;
    }
    public override string ToString() {
      string s = "_44_Logic____Option__i_Compile.Option.Some";
s += "(";
s += Dafny.Helpers.ToString(this.v);
s += ")";
return s;
    }
  }
  public class Option_None<T> : Option<T> {
    public Option_None() {
    }
    public override bool Equals(object other) {
      var oth = other as _44_Logic____Option__i_Compile.Option_None<T>;
return oth != null;
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 1;
return (int) hash;
    }
    public override string ToString() {
      string s = "_44_Logic____Option__i_Compile.Option.None";
return s;
    }
  }

} // end of namespace _44_Logic____Option__i_Compile
namespace _47_Collections____Maps__i_Compile {

  public partial class __default {
    public static Dafny.Set<U> domain<U,V>(Dafny.Map<U,V> m) {
      return ((System.Func<Dafny.Set<U>>)(() => {
        var _coll0 = new System.Collections.Generic.List<U>();
foreach (var _compr_0 in (m).Keys.Elements) { U _1565_s = (U)_compr_0;
          if ((m).Contains(_1565_s)) {
            _coll0.Add(_1565_s);
          }
        }
        return Dafny.Set<U>.FromCollection(_coll0);
      }))();
    }
    public static Dafny.Map<U,V> RemoveElt<U,V>(Dafny.Map<U,V> m, U elt)
    {
      Dafny.Map<U,V> _1566_m_k = ((System.Func<Dafny.Map<U,V>>)(() => {
        var _coll1 = new System.Collections.Generic.List<Dafny.Pair<U,V>>();
foreach (var _1567_elt_k in (m).Keys.Elements) {
          if (((m).Contains(_1567_elt_k)) && (!(_1567_elt_k).Equals(elt))) {
            _coll1.Add(new Dafny.Pair<U,V>(_1567_elt_k,(m).Select(_1567_elt_k)));
          }
        }
        return Dafny.Map<U,V>.FromCollection(_coll1);
      }))();
return _1566_m_k;
    }
  }
} // end of namespace _47_Collections____Maps__i_Compile
namespace _50_Collections____Seqs__i_Compile {


} // end of namespace _50_Collections____Seqs__i_Compile
namespace _54_Native____NativeTypes__i_Compile {


  public partial class __default {
    public static ulong Uint64Size() {
      return 8UL;
    }
    public static ulong Uint32Size() {
      return 4UL;
    }
    public static ulong Uint16Size() {
      return 2UL;
    }
  }
} // end of namespace _54_Native____NativeTypes__i_Compile
namespace _57_Libraries____base__s_Compile {

} // end of namespace _57_Libraries____base__s_Compile
namespace _59_Math____power2__s_Compile {


} // end of namespace _59_Math____power2__s_Compile
namespace _61_Math____power__s_Compile {

} // end of namespace _61_Math____power__s_Compile
namespace _64_Math____mul__nonlinear__i_Compile {

} // end of namespace _64_Math____mul__nonlinear__i_Compile
namespace _67_Math____mul__auto__proofs__i_Compile {


} // end of namespace _67_Math____mul__auto__proofs__i_Compile
namespace _69_Math____mul__auto__i_Compile {


} // end of namespace _69_Math____mul__auto__i_Compile
namespace _71_Math____mul__i_Compile {



} // end of namespace _71_Math____mul__i_Compile
namespace _73_Math____power__i_Compile {



} // end of namespace _73_Math____power__i_Compile
namespace _77_Math____div__def__i_Compile {

} // end of namespace _77_Math____div__def__i_Compile
namespace _81_Math____div__boogie__i_Compile {



} // end of namespace _81_Math____div__boogie__i_Compile
namespace _83_Math____div__nonlinear__i_Compile {

} // end of namespace _83_Math____div__nonlinear__i_Compile
namespace _88_Math____mod__auto__proofs__i_Compile {




} // end of namespace _88_Math____mod__auto__proofs__i_Compile
namespace _90_Math____mod__auto__i_Compile {


} // end of namespace _90_Math____mod__auto__i_Compile
namespace _93_Math____div__auto__proofs__i_Compile {


} // end of namespace _93_Math____div__auto__proofs__i_Compile
namespace _95_Math____div__auto__i_Compile {



} // end of namespace _95_Math____div__auto__i_Compile
namespace _97_Math____div__i_Compile {







} // end of namespace _97_Math____div__i_Compile
namespace _99_Math____power2__i_Compile {





} // end of namespace _99_Math____power2__i_Compile
namespace _101_Common____Util__i_Compile {




  public partial class __default {
    public static void seqToArray__slow<A>(Dafny.Sequence<A> s, out A[] a)
    {
    TAIL_CALL_START: ;
a = new A[0];
      BigInteger _1568_len;
      _1568_len = new BigInteger((s).Count);
      var _nw0 = Dafny.ArrayHelpers.InitNewArray1<A>(Dafny.Helpers.Default<A>(), (_1568_len));
      a = _nw0;
      BigInteger _1569_i;
      _1569_i = new BigInteger(0);
      while ((_1569_i) < (_1568_len)) {
        (a)[(int)((_1569_i))] = (s).Select(_1569_i);
        _1569_i = (_1569_i) + (new BigInteger(1));
      }
    }
    public static void seqIntoArrayOpt<A>(Dafny.Sequence<A> s, A[] a)
    {
    TAIL_CALL_START: ;
      ulong _1570_i;
      _1570_i = 0UL;
      while ((_1570_i) < ((ulong)(s).LongCount)) {
        (a)[(int)((_1570_i))] = (s).Select(_1570_i);
        _1570_i = (_1570_i) + (1UL);
      }
    }
    public static void seqToArrayOpt<A>(Dafny.Sequence<A> s, out A[] a)
    {
    TAIL_CALL_START: ;
a = new A[0];
      var _nw1 = Dafny.ArrayHelpers.InitNewArray1<A>(Dafny.Helpers.Default<A>(), ((ulong)(s).LongCount));
      a = _nw1;
      _101_Common____Util__i_Compile.__default.seqIntoArrayOpt<A>(s, a);
    }
    public static void seqIntoArrayChar(Dafny.Sequence<char> s, char[] a)
    {
    TAIL_CALL_START: ;
      ulong _1571_i;
      _1571_i = 0UL;
      while ((_1571_i) < ((ulong)(s).LongCount)) {
        (a)[(int)((_1571_i))] = (s).Select(_1571_i);
        _1571_i = (_1571_i) + (1UL);
      }
    }
    public static void RecordTimingSeq(Dafny.Sequence<char> name, ulong start, ulong end)
    {
    TAIL_CALL_START: ;
      char[] _1572_name__array;
      var _nw2 = new char[(int)(new BigInteger((name).Count))];
      _1572_name__array = _nw2;
      _101_Common____Util__i_Compile.__default.seqIntoArrayChar(name, _1572_name__array);
      ulong _1573_time = 0;
      if ((start) <= (end)) {
        _1573_time = (end) - (start);
      } else {
        _1573_time = 18446744073709551615UL;
      }
      _9_Native____Io__s_Compile.Time.RecordTiming(_1572_name__array, _1573_time);
    }
    public static ulong SeqByteToUint64(Dafny.Sequence<byte> bs) {
      return ((((((((((((ulong)((bs).Select((ulong)(0UL)))) * (256UL)) * (256UL)) * (256UL)) * (4294967296UL)) + (((((ulong)((bs).Select((ulong)(1UL)))) * (256UL)) * (256UL)) * (4294967296UL))) + ((((ulong)((bs).Select((ulong)(2UL)))) * (256UL)) * (4294967296UL))) + (((ulong)((bs).Select((ulong)(3UL)))) * (4294967296UL))) + (((((ulong)((bs).Select((ulong)(4UL)))) * (256UL)) * (256UL)) * (256UL))) + ((((ulong)((bs).Select((ulong)(5UL)))) * (256UL)) * (256UL))) + (((ulong)((bs).Select((ulong)(6UL)))) * (256UL))) + ((ulong)((bs).Select((ulong)(7UL))));
    }
    public static Dafny.Sequence<byte> Uint64ToSeqByte(ulong u) {
      Dafny.Sequence<byte> _1574_bs = Dafny.Sequence<byte>.FromElements((byte)((u) / (72057594037927936UL)), (byte)(((u) / (281474976710656UL)) % (256UL)), (byte)(((u) / (1099511627776UL)) % (256UL)), (byte)(((u) / (4294967296UL)) % (256UL)), (byte)(((u) / (16777216UL)) % (256UL)), (byte)(((u) / (65536UL)) % (256UL)), (byte)(((u) / (256UL)) % (256UL)), (byte)((u) % (256UL)));
BigInteger _1575_u__int = new BigInteger(u);
return _1574_bs;
    }
    public static ushort SeqByteToUint16(Dafny.Sequence<byte> bs) {
      return (ushort)(((ushort)(((ushort)((bs).Select((ulong)(0UL)))) * (256))) + ((ushort)((bs).Select((ulong)(1UL)))));
    }
    public static Dafny.Sequence<byte> Uint16ToSeqByte(ushort u) {
      Dafny.Sequence<byte> _1576_s = Dafny.Sequence<byte>.FromElements((byte)((ushort)(((ushort)((u) / (256))) % (256))), (byte)((ushort)((u) % (256))));
BigInteger _1577_u__int = new BigInteger(u);
return _1576_s;
    }
  }
} // end of namespace _101_Common____Util__i_Compile
namespace _105_Common____MarshallInt__i_Compile {



  public partial class __default {
    public static void MarshallUint64__guts(ulong n, byte[] data, ulong index)
    {
    TAIL_CALL_START: ;
      (data)[(int)((index))] = (byte)((n) / (72057594037927936UL));
      var _index0 = (index) + (1UL);
      (data)[(int)(_index0)] = (byte)(((n) / (281474976710656UL)) % (256UL));
      var _index1 = (index) + (2UL);
      (data)[(int)(_index1)] = (byte)(((n) / (1099511627776UL)) % (256UL));
      var _index2 = (index) + (3UL);
      (data)[(int)(_index2)] = (byte)(((n) / (4294967296UL)) % (256UL));
      var _index3 = (index) + (4UL);
      (data)[(int)(_index3)] = (byte)(((n) / (16777216UL)) % (256UL));
      var _index4 = (index) + (5UL);
      (data)[(int)(_index4)] = (byte)(((n) / (65536UL)) % (256UL));
      var _index5 = (index) + (6UL);
      (data)[(int)(_index5)] = (byte)(((n) / (256UL)) % (256UL));
      var _index6 = (index) + (7UL);
      (data)[(int)(_index6)] = (byte)((n) % (256UL));
      { }
      { }
      { }
    }
  }
} // end of namespace _105_Common____MarshallInt__i_Compile
namespace _107_Common____GenericMarshalling__i_Compile {







  public abstract class G {
    public G() { }
static G theDefault;
public static G Default {
      get {
        if (theDefault == null) {
          theDefault = new _107_Common____GenericMarshalling__i_Compile.G_GUint64();
        }
        return theDefault;
      }
    }
    public static G _DafnyDefaultValue() { return Default; }
public static G create_GUint64() {
      return new G_GUint64();
    }
    public static G create_GArray(_107_Common____GenericMarshalling__i_Compile.G elt) {
      return new G_GArray(elt);
    }
    public static G create_GTuple(Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.G> t) {
      return new G_GTuple(t);
    }
    public static G create_GByteArray() {
      return new G_GByteArray();
    }
    public static G create_GTaggedUnion(Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.G> cases) {
      return new G_GTaggedUnion(cases);
    }
    public bool is_GUint64 { get { return this is G_GUint64; } }
public bool is_GArray { get { return this is G_GArray; } }
public bool is_GTuple { get { return this is G_GTuple; } }
public bool is_GByteArray { get { return this is G_GByteArray; } }
public bool is_GTaggedUnion { get { return this is G_GTaggedUnion; } }
public _107_Common____GenericMarshalling__i_Compile.G dtor_elt {
      get {
        var d = this;
return ((G_GArray)d).elt; 
      }
    }
    public Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.G> dtor_t {
      get {
        var d = this;
return ((G_GTuple)d).t; 
      }
    }
    public Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.G> dtor_cases {
      get {
        var d = this;
return ((G_GTaggedUnion)d).cases; 
      }
    }
  }
  public class G_GUint64 : G {
    public G_GUint64() {
    }
    public override bool Equals(object other) {
      var oth = other as _107_Common____GenericMarshalling__i_Compile.G_GUint64;
return oth != null;
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 0;
return (int) hash;
    }
    public override string ToString() {
      string s = "_107_Common____GenericMarshalling__i_Compile.G.GUint64";
return s;
    }
  }
  public class G_GArray : G {
    public readonly _107_Common____GenericMarshalling__i_Compile.G elt;
public G_GArray(_107_Common____GenericMarshalling__i_Compile.G elt) {
      this.elt = elt;
    }
    public override bool Equals(object other) {
      var oth = other as _107_Common____GenericMarshalling__i_Compile.G_GArray;
return oth != null && Dafny.Helpers.AreEqual(this.elt, oth.elt);
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 1;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.elt));
return (int) hash;
    }
    public override string ToString() {
      string s = "_107_Common____GenericMarshalling__i_Compile.G.GArray";
s += "(";
s += Dafny.Helpers.ToString(this.elt);
s += ")";
return s;
    }
  }
  public class G_GTuple : G {
    public readonly Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.G> t;
public G_GTuple(Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.G> t) {
      this.t = t;
    }
    public override bool Equals(object other) {
      var oth = other as _107_Common____GenericMarshalling__i_Compile.G_GTuple;
return oth != null && Dafny.Helpers.AreEqual(this.t, oth.t);
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 2;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.t));
return (int) hash;
    }
    public override string ToString() {
      string s = "_107_Common____GenericMarshalling__i_Compile.G.GTuple";
s += "(";
s += Dafny.Helpers.ToString(this.t);
s += ")";
return s;
    }
  }
  public class G_GByteArray : G {
    public G_GByteArray() {
    }
    public override bool Equals(object other) {
      var oth = other as _107_Common____GenericMarshalling__i_Compile.G_GByteArray;
return oth != null;
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 3;
return (int) hash;
    }
    public override string ToString() {
      string s = "_107_Common____GenericMarshalling__i_Compile.G.GByteArray";
return s;
    }
  }
  public class G_GTaggedUnion : G {
    public readonly Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.G> cases;
public G_GTaggedUnion(Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.G> cases) {
      this.cases = cases;
    }
    public override bool Equals(object other) {
      var oth = other as _107_Common____GenericMarshalling__i_Compile.G_GTaggedUnion;
return oth != null && Dafny.Helpers.AreEqual(this.cases, oth.cases);
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 4;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.cases));
return (int) hash;
    }
    public override string ToString() {
      string s = "_107_Common____GenericMarshalling__i_Compile.G.GTaggedUnion";
s += "(";
s += Dafny.Helpers.ToString(this.cases);
s += ")";
return s;
    }
  }

  public abstract class V {
    public V() { }
static V theDefault;
public static V Default {
      get {
        if (theDefault == null) {
          theDefault = new _107_Common____GenericMarshalling__i_Compile.V_VUint64(0);
        }
        return theDefault;
      }
    }
    public static V _DafnyDefaultValue() { return Default; }
public static V create_VUint64(ulong u) {
      return new V_VUint64(u);
    }
    public static V create_VArray(Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V> a) {
      return new V_VArray(a);
    }
    public static V create_VTuple(Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V> t) {
      return new V_VTuple(t);
    }
    public static V create_VByteArray(Dafny.Sequence<byte> b) {
      return new V_VByteArray(b);
    }
    public static V create_VCase(ulong c, _107_Common____GenericMarshalling__i_Compile.V val) {
      return new V_VCase(c, val);
    }
    public bool is_VUint64 { get { return this is V_VUint64; } }
public bool is_VArray { get { return this is V_VArray; } }
public bool is_VTuple { get { return this is V_VTuple; } }
public bool is_VByteArray { get { return this is V_VByteArray; } }
public bool is_VCase { get { return this is V_VCase; } }
public ulong dtor_u {
      get {
        var d = this;
return ((V_VUint64)d).u; 
      }
    }
    public Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V> dtor_a {
      get {
        var d = this;
return ((V_VArray)d).a; 
      }
    }
    public Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V> dtor_t {
      get {
        var d = this;
return ((V_VTuple)d).t; 
      }
    }
    public Dafny.Sequence<byte> dtor_b {
      get {
        var d = this;
return ((V_VByteArray)d).b; 
      }
    }
    public ulong dtor_c {
      get {
        var d = this;
return ((V_VCase)d).c; 
      }
    }
    public _107_Common____GenericMarshalling__i_Compile.V dtor_val {
      get {
        var d = this;
return ((V_VCase)d).val; 
      }
    }
  }
  public class V_VUint64 : V {
    public readonly ulong u;
public V_VUint64(ulong u) {
      this.u = u;
    }
    public override bool Equals(object other) {
      var oth = other as _107_Common____GenericMarshalling__i_Compile.V_VUint64;
return oth != null && this.u == oth.u;
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 0;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.u));
return (int) hash;
    }
    public override string ToString() {
      string s = "_107_Common____GenericMarshalling__i_Compile.V.VUint64";
s += "(";
s += Dafny.Helpers.ToString(this.u);
s += ")";
return s;
    }
  }
  public class V_VArray : V {
    public readonly Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V> a;
public V_VArray(Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V> a) {
      this.a = a;
    }
    public override bool Equals(object other) {
      var oth = other as _107_Common____GenericMarshalling__i_Compile.V_VArray;
return oth != null && Dafny.Helpers.AreEqual(this.a, oth.a);
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 1;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.a));
return (int) hash;
    }
    public override string ToString() {
      string s = "_107_Common____GenericMarshalling__i_Compile.V.VArray";
s += "(";
s += Dafny.Helpers.ToString(this.a);
s += ")";
return s;
    }
  }
  public class V_VTuple : V {
    public readonly Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V> t;
public V_VTuple(Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V> t) {
      this.t = t;
    }
    public override bool Equals(object other) {
      var oth = other as _107_Common____GenericMarshalling__i_Compile.V_VTuple;
return oth != null && Dafny.Helpers.AreEqual(this.t, oth.t);
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 2;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.t));
return (int) hash;
    }
    public override string ToString() {
      string s = "_107_Common____GenericMarshalling__i_Compile.V.VTuple";
s += "(";
s += Dafny.Helpers.ToString(this.t);
s += ")";
return s;
    }
  }
  public class V_VByteArray : V {
    public readonly Dafny.Sequence<byte> b;
public V_VByteArray(Dafny.Sequence<byte> b) {
      this.b = b;
    }
    public override bool Equals(object other) {
      var oth = other as _107_Common____GenericMarshalling__i_Compile.V_VByteArray;
return oth != null && Dafny.Helpers.AreEqual(this.b, oth.b);
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 3;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.b));
return (int) hash;
    }
    public override string ToString() {
      string s = "_107_Common____GenericMarshalling__i_Compile.V.VByteArray";
s += "(";
s += Dafny.Helpers.ToString(this.b);
s += ")";
return s;
    }
  }
  public class V_VCase : V {
    public readonly ulong c;
public readonly _107_Common____GenericMarshalling__i_Compile.V val;
public V_VCase(ulong c, _107_Common____GenericMarshalling__i_Compile.V val) {
      this.c = c;
this.val = val;
    }
    public override bool Equals(object other) {
      var oth = other as _107_Common____GenericMarshalling__i_Compile.V_VCase;
return oth != null && this.c == oth.c && Dafny.Helpers.AreEqual(this.val, oth.val);
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 4;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.c));
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.val));
return (int) hash;
    }
    public override string ToString() {
      string s = "_107_Common____GenericMarshalling__i_Compile.V.VCase";
s += "(";
s += Dafny.Helpers.ToString(this.c);
s += ", ";
s += Dafny.Helpers.ToString(this.val);
s += ")";
return s;
    }
  }

  public class ContentsTraceStep {
    public readonly Dafny.Sequence<byte> data;
public readonly _44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V> val;
public ContentsTraceStep(Dafny.Sequence<byte> data, _44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V> val) {
      this.data = data;
this.val = val;
    }
    public override bool Equals(object other) {
      var oth = other as _107_Common____GenericMarshalling__i_Compile.ContentsTraceStep;
return oth != null && Dafny.Helpers.AreEqual(this.data, oth.data) && Dafny.Helpers.AreEqual(this.val, oth.val);
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 0;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.data));
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.val));
return (int) hash;
    }
    public override string ToString() {
      string s = "_107_Common____GenericMarshalling__i_Compile.ContentsTraceStep.ContentsTraceStep";
s += "(";
s += Dafny.Helpers.ToString(this.data);
s += ", ";
s += Dafny.Helpers.ToString(this.val);
s += ")";
return s;
    }
    static ContentsTraceStep theDefault;
public static ContentsTraceStep Default {
      get {
        if (theDefault == null) {
          theDefault = new _107_Common____GenericMarshalling__i_Compile.ContentsTraceStep(Dafny.Sequence<byte>.Empty, @_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>.Default);
        }
        return theDefault;
      }
    }
    public static ContentsTraceStep _DafnyDefaultValue() { return Default; }
public static ContentsTraceStep create(Dafny.Sequence<byte> data, _44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V> val) {
      return new ContentsTraceStep(data, val);
    }
    public bool is_ContentsTraceStep { get { return true; } }
public Dafny.Sequence<byte> dtor_data {
      get {
        return this.data;
      }
    }
    public _44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V> dtor_val {
      get {
        return this.val;
      }
    }
  }

  public partial class __default {
    public static _System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>> parse__Uint64(Dafny.Sequence<byte> data) {
      if (((ulong)(data).LongCount) >= (_54_Native____NativeTypes__i_Compile.__default.Uint64Size())) {
        return @_System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>>.create(@_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>.create_Some(@_107_Common____GenericMarshalling__i_Compile.V.create_VUint64(_101_Common____Util__i_Compile.__default.SeqByteToUint64((data).Take(_54_Native____NativeTypes__i_Compile.__default.Uint64Size())))), (data).Drop(_54_Native____NativeTypes__i_Compile.__default.Uint64Size()));
      } else  {
        return @_System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>>.create(@_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>.create_None(), Dafny.Sequence<byte>.FromElements());
      }
    }
    public static void ParseUint64(byte[] data, ulong index, out bool success, out _107_Common____GenericMarshalling__i_Compile.V v, out ulong rest__index)
    {
    TAIL_CALL_START: ;
success = false;
v = @_107_Common____GenericMarshalling__i_Compile.V.Default;
rest__index = 0;
      { }
      if ((((ulong)(data).LongLength) >= (8UL)) && ((index) <= (((ulong)(data).LongLength) - (8UL)))) {
        ulong _1578_result;
        _1578_result = (((((((((ulong)((data)[(int)((index) + ((ulong)(0UL)))])) * (72057594037927936UL)) + (((ulong)((data)[(int)((index) + ((ulong)(1UL)))])) * (281474976710656UL))) + (((ulong)((data)[(int)((index) + ((ulong)(2UL)))])) * (1099511627776UL))) + (((ulong)((data)[(int)((index) + ((ulong)(3UL)))])) * (4294967296UL))) + (((ulong)((data)[(int)((index) + ((ulong)(4UL)))])) * (16777216UL))) + (((ulong)((data)[(int)((index) + ((ulong)(5UL)))])) * (65536UL))) + (((ulong)((data)[(int)((index) + ((ulong)(6UL)))])) * (256UL))) + ((ulong)((data)[(int)((index) + ((ulong)(7UL)))]));
        success = true;
        v = @_107_Common____GenericMarshalling__i_Compile.V.create_VUint64(_1578_result);
        rest__index = (index) + (_54_Native____NativeTypes__i_Compile.__default.Uint64Size());
      } else {
        success = false;
        rest__index = (ulong)(data).LongLength;
      }
    }
    public static _System.Tuple2<_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>>,Dafny.Sequence<byte>> parse__Array__contents(Dafny.Sequence<byte> data, _107_Common____GenericMarshalling__i_Compile.G eltType, ulong len)
    {
      if ((len) == (0UL)) {
        return @_System.Tuple2<_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>>,Dafny.Sequence<byte>>.create(@_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>>.create_Some(Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>.FromElements()), data);
      } else  {
        _System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>> _let_tmp_rhs0 = _107_Common____GenericMarshalling__i_Compile.__default.parse__Val(data, eltType);
_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V> _1579_val = ((_System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>>)_let_tmp_rhs0)._0;
Dafny.Sequence<byte> _1580_rest1 = ((_System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>>)_let_tmp_rhs0)._1;
_System.Tuple2<_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>>,Dafny.Sequence<byte>> _let_tmp_rhs1 = _107_Common____GenericMarshalling__i_Compile.__default.parse__Array__contents(_1580_rest1, eltType, (len) - (1UL));
_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>> _1581_others = ((_System.Tuple2<_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>>,Dafny.Sequence<byte>>)_let_tmp_rhs1)._0;
Dafny.Sequence<byte> _1582_rest2 = ((_System.Tuple2<_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>>,Dafny.Sequence<byte>>)_let_tmp_rhs1)._1;
if ((!((_1579_val).is_None)) && (!((_1581_others).is_None))) {
          return @_System.Tuple2<_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>>,Dafny.Sequence<byte>>.create(@_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>>.create_Some((Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>.FromElements((_1579_val).dtor_v)).Concat((_1581_others).dtor_v)), _1582_rest2);
        } else  {
          return @_System.Tuple2<_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>>,Dafny.Sequence<byte>>.create(@_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>>.create_None(), Dafny.Sequence<byte>.FromElements());
        }
      }
    }
    public static void ParseArrayContents(byte[] data, ulong index, _107_Common____GenericMarshalling__i_Compile.G eltType, ulong len, out bool success, out Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V> v, out ulong rest__index)
    {
      success = false;
v = Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>.Empty;
rest__index = 0;
      { }
      _107_Common____GenericMarshalling__i_Compile.V[] _1583_vArr;
      var _nw3 = Dafny.ArrayHelpers.InitNewArray1<_107_Common____GenericMarshalling__i_Compile.V>(@_107_Common____GenericMarshalling__i_Compile.V.Default, (len));
      _1583_vArr = _nw3;
      { }
      success = true;
      ulong _1584_i;
      _1584_i = 0UL;
      ulong _1585_next__val__index;
      _1585_next__val__index = index;
      { }
      while ((_1584_i) < (len)) {
        bool _1586_some1;
_107_Common____GenericMarshalling__i_Compile.V _1587_val;
ulong _1588_rest1;
bool _out0;
_107_Common____GenericMarshalling__i_Compile.V _out1;
ulong _out2;
_107_Common____GenericMarshalling__i_Compile.__default.ParseVal(data, _1585_next__val__index, eltType, out _out0, out _out1, out _out2);
_1586_some1 = _out0;
_1587_val = _out1;
_1588_rest1 = _out2;
        { }
        { }
        { }
        if (!(_1586_some1)) {
          success = false;
          rest__index = (ulong)(data).LongLength;
          { }
          return;
        }
        { }
        (_1583_vArr)[(int)((_1584_i))] = _1587_val;
        _1585_next__val__index = _1588_rest1;
        _1584_i = (_1584_i) + (1UL);
      }
      success = true;
      rest__index = _1585_next__val__index;
      v = Dafny.Helpers.SeqFromArray(_1583_vArr);
      { }
    }
    public static _System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>> parse__Array(Dafny.Sequence<byte> data, _107_Common____GenericMarshalling__i_Compile.G eltType)
    {
      _System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>> _let_tmp_rhs2 = _107_Common____GenericMarshalling__i_Compile.__default.parse__Uint64(data);
_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V> _1589_len = ((_System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>>)_let_tmp_rhs2)._0;
Dafny.Sequence<byte> _1590_rest = ((_System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>>)_let_tmp_rhs2)._1;
if (!((_1589_len).is_None)) {
        _System.Tuple2<_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>>,Dafny.Sequence<byte>> _let_tmp_rhs3 = _107_Common____GenericMarshalling__i_Compile.__default.parse__Array__contents(_1590_rest, eltType, ((_1589_len).dtor_v).dtor_u);
_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>> _1591_contents = ((_System.Tuple2<_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>>,Dafny.Sequence<byte>>)_let_tmp_rhs3)._0;
Dafny.Sequence<byte> _1592_remainder = ((_System.Tuple2<_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>>,Dafny.Sequence<byte>>)_let_tmp_rhs3)._1;
if (!((_1591_contents).is_None)) {
          return @_System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>>.create(@_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>.create_Some(@_107_Common____GenericMarshalling__i_Compile.V.create_VArray((_1591_contents).dtor_v)), _1592_remainder);
        } else  {
          return @_System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>>.create(@_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>.create_None(), Dafny.Sequence<byte>.FromElements());
        }
      } else  {
        return @_System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>>.create(@_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>.create_None(), Dafny.Sequence<byte>.FromElements());
      }
    }
    public static void ParseArray(byte[] data, ulong index, _107_Common____GenericMarshalling__i_Compile.G eltType, out bool success, out _107_Common____GenericMarshalling__i_Compile.V v, out ulong rest__index)
    {
      success = false;
v = @_107_Common____GenericMarshalling__i_Compile.V.Default;
rest__index = 0;
      bool _1593_some1;
_107_Common____GenericMarshalling__i_Compile.V _1594_len;
ulong _1595_rest;
bool _out3;
_107_Common____GenericMarshalling__i_Compile.V _out4;
ulong _out5;
_107_Common____GenericMarshalling__i_Compile.__default.ParseUint64(data, index, out _out3, out _out4, out _out5);
_1593_some1 = _out3;
_1594_len = _out4;
_1595_rest = _out5;
      if (_1593_some1) {
        bool _1596_some2;
Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V> _1597_contents;
ulong _1598_remainder;
bool _out6;
Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V> _out7;
ulong _out8;
_107_Common____GenericMarshalling__i_Compile.__default.ParseArrayContents(data, _1595_rest, eltType, (_1594_len).dtor_u, out _out6, out _out7, out _out8);
_1596_some2 = _out6;
_1597_contents = _out7;
_1598_remainder = _out8;
        if (_1596_some2) {
          success = true;
          v = @_107_Common____GenericMarshalling__i_Compile.V.create_VArray(_1597_contents);
          rest__index = _1598_remainder;
        } else {
          success = false;
          rest__index = (ulong)(data).LongLength;
        }
      } else {
        success = false;
        rest__index = (ulong)(data).LongLength;
      }
    }
    public static _System.Tuple2<_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>>,Dafny.Sequence<byte>> parse__Tuple__contents(Dafny.Sequence<byte> data, Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.G> eltTypes)
    {
      if ((eltTypes).Equals(Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.G>.FromElements())) {
        return @_System.Tuple2<_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>>,Dafny.Sequence<byte>>.create(@_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>>.create_Some(Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>.FromElements()), data);
      } else  {
        _System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>> _let_tmp_rhs4 = _107_Common____GenericMarshalling__i_Compile.__default.parse__Val(data, (eltTypes).Select((ulong)(0UL)));
_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V> _1599_val = ((_System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>>)_let_tmp_rhs4)._0;
Dafny.Sequence<byte> _1600_rest1 = ((_System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>>)_let_tmp_rhs4)._1;
_System.Tuple2<_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>>,Dafny.Sequence<byte>> _let_tmp_rhs5 = _107_Common____GenericMarshalling__i_Compile.__default.parse__Tuple__contents(_1600_rest1, (eltTypes).Drop((ulong)(1UL)));
_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>> _1601_contents = ((_System.Tuple2<_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>>,Dafny.Sequence<byte>>)_let_tmp_rhs5)._0;
Dafny.Sequence<byte> _1602_rest2 = ((_System.Tuple2<_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>>,Dafny.Sequence<byte>>)_let_tmp_rhs5)._1;
if ((!((_1599_val).is_None)) && (!((_1601_contents).is_None))) {
          return @_System.Tuple2<_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>>,Dafny.Sequence<byte>>.create(@_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>>.create_Some((Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>.FromElements((_1599_val).dtor_v)).Concat((_1601_contents).dtor_v)), _1602_rest2);
        } else  {
          return @_System.Tuple2<_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>>,Dafny.Sequence<byte>>.create(@_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>>.create_None(), Dafny.Sequence<byte>.FromElements());
        }
      }
    }
    public static void ParseTupleContents(byte[] data, ulong index, Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.G> eltTypes, out bool success, out Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V> v, out ulong rest__index)
    {
      success = false;
v = Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>.Empty;
rest__index = 0;
      { }
      _107_Common____GenericMarshalling__i_Compile.V[] _1603_vArr;
      var _nw4 = Dafny.ArrayHelpers.InitNewArray1<_107_Common____GenericMarshalling__i_Compile.V>(@_107_Common____GenericMarshalling__i_Compile.V.Default, ((ulong)(eltTypes).LongCount));
      _1603_vArr = _nw4;
      { }
      success = true;
      ulong _1604_i;
      _1604_i = 0UL;
      ulong _1605_next__val__index;
      _1605_next__val__index = index;
      { }
      while ((_1604_i) < ((ulong)(eltTypes).LongCount)) {
        bool _1606_some1;
_107_Common____GenericMarshalling__i_Compile.V _1607_val;
ulong _1608_rest1;
bool _out9;
_107_Common____GenericMarshalling__i_Compile.V _out10;
ulong _out11;
_107_Common____GenericMarshalling__i_Compile.__default.ParseVal(data, _1605_next__val__index, (eltTypes).Select(_1604_i), out _out9, out _out10, out _out11);
_1606_some1 = _out9;
_1607_val = _out10;
_1608_rest1 = _out11;
        { }
        { }
        { }
        if (!(_1606_some1)) {
          success = false;
          rest__index = (ulong)(data).LongLength;
          { }
          return;
        }
        { }
        (_1603_vArr)[(int)((_1604_i))] = _1607_val;
        _1605_next__val__index = _1608_rest1;
        _1604_i = (_1604_i) + (1UL);
      }
      success = true;
      rest__index = _1605_next__val__index;
      v = Dafny.Helpers.SeqFromArray(_1603_vArr);
      { }
    }
    public static _System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>> parse__Tuple(Dafny.Sequence<byte> data, Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.G> eltTypes)
    {
      _System.Tuple2<_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>>,Dafny.Sequence<byte>> _let_tmp_rhs6 = _107_Common____GenericMarshalling__i_Compile.__default.parse__Tuple__contents(data, eltTypes);
_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>> _1609_contents = ((_System.Tuple2<_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>>,Dafny.Sequence<byte>>)_let_tmp_rhs6)._0;
Dafny.Sequence<byte> _1610_rest = ((_System.Tuple2<_44_Logic____Option__i_Compile.Option<Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V>>,Dafny.Sequence<byte>>)_let_tmp_rhs6)._1;
if (!((_1609_contents).is_None)) {
        return @_System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>>.create(@_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>.create_Some(@_107_Common____GenericMarshalling__i_Compile.V.create_VTuple((_1609_contents).dtor_v)), _1610_rest);
      } else  {
        return @_System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>>.create(@_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>.create_None(), Dafny.Sequence<byte>.FromElements());
      }
    }
    public static void ParseTuple(byte[] data, ulong index, Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.G> eltTypes, out bool success, out _107_Common____GenericMarshalling__i_Compile.V v, out ulong rest__index)
    {
      success = false;
v = @_107_Common____GenericMarshalling__i_Compile.V.Default;
rest__index = 0;
      bool _1611_some;
Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V> _1612_contents;
ulong _1613_rest;
bool _out12;
Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V> _out13;
ulong _out14;
_107_Common____GenericMarshalling__i_Compile.__default.ParseTupleContents(data, index, eltTypes, out _out12, out _out13, out _out14);
_1611_some = _out12;
_1612_contents = _out13;
_1613_rest = _out14;
      if (_1611_some) {
        success = true;
        v = @_107_Common____GenericMarshalling__i_Compile.V.create_VTuple(_1612_contents);
        rest__index = _1613_rest;
      } else {
        success = false;
        rest__index = (ulong)(data).LongLength;
      }
    }
    public static _System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>> parse__ByteArray(Dafny.Sequence<byte> data) {
      _System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>> _let_tmp_rhs7 = _107_Common____GenericMarshalling__i_Compile.__default.parse__Uint64(data);
_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V> _1614_len = ((_System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>>)_let_tmp_rhs7)._0;
Dafny.Sequence<byte> _1615_rest = ((_System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>>)_let_tmp_rhs7)._1;
if ((!((_1614_len).is_None)) && ((((_1614_len).dtor_v).dtor_u) <= ((ulong)(_1615_rest).LongCount))) {
        return @_System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>>.create(@_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>.create_Some(@_107_Common____GenericMarshalling__i_Compile.V.create_VByteArray((_1615_rest).Take(((_1614_len).dtor_v).dtor_u).Drop((ulong)(0UL)))), (_1615_rest).Drop(((_1614_len).dtor_v).dtor_u));
      } else  {
        return @_System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>>.create(@_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>.create_None(), Dafny.Sequence<byte>.FromElements());
      }
    }
    public static void ParseByteArray(byte[] data, ulong index, out bool success, out _107_Common____GenericMarshalling__i_Compile.V v, out ulong rest__index)
    {
    TAIL_CALL_START: ;
success = false;
v = @_107_Common____GenericMarshalling__i_Compile.V.Default;
rest__index = 0;
      bool _1616_some;
_107_Common____GenericMarshalling__i_Compile.V _1617_len;
ulong _1618_rest;
bool _out15;
_107_Common____GenericMarshalling__i_Compile.V _out16;
ulong _out17;
_107_Common____GenericMarshalling__i_Compile.__default.ParseUint64(data, index, out _out15, out _out16, out _out17);
_1616_some = _out15;
_1617_len = _out16;
_1618_rest = _out17;
      if ((_1616_some) && (((_1617_len).dtor_u) <= (((ulong)(data).LongLength) - (_1618_rest)))) {
        Dafny.Sequence<byte> _1619_rest__seq;
        _1619_rest__seq = Dafny.Helpers.SeqFromArray(data).Drop(_1618_rest);
        { }
        { }
        success = true;
        v = @_107_Common____GenericMarshalling__i_Compile.V.create_VByteArray(Dafny.Helpers.SeqFromArray(data).Take((_1618_rest) + ((_1617_len).dtor_u)).Drop(_1618_rest));
        rest__index = (_1618_rest) + ((_1617_len).dtor_u);
      } else {
        success = false;
        rest__index = (ulong)(data).LongLength;
      }
    }
    public static _System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>> parse__Case(Dafny.Sequence<byte> data, Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.G> cases)
    {
      _System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>> _let_tmp_rhs8 = _107_Common____GenericMarshalling__i_Compile.__default.parse__Uint64(data);
_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V> _1620_caseID = ((_System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>>)_let_tmp_rhs8)._0;
Dafny.Sequence<byte> _1621_rest1 = ((_System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>>)_let_tmp_rhs8)._1;
if ((!((_1620_caseID).is_None)) && ((((_1620_caseID).dtor_v).dtor_u) < ((ulong)(cases).LongCount))) {
        _System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>> _let_tmp_rhs9 = _107_Common____GenericMarshalling__i_Compile.__default.parse__Val(_1621_rest1, (cases).Select(((_1620_caseID).dtor_v).dtor_u));
_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V> _1622_val = ((_System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>>)_let_tmp_rhs9)._0;
Dafny.Sequence<byte> _1623_rest2 = ((_System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>>)_let_tmp_rhs9)._1;
if (!((_1622_val).is_None)) {
          return @_System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>>.create(@_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>.create_Some(@_107_Common____GenericMarshalling__i_Compile.V.create_VCase(((_1620_caseID).dtor_v).dtor_u, (_1622_val).dtor_v)), _1623_rest2);
        } else  {
          return @_System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>>.create(@_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>.create_None(), Dafny.Sequence<byte>.FromElements());
        }
      } else  {
        return @_System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>>.create(@_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>.create_None(), Dafny.Sequence<byte>.FromElements());
      }
    }
    public static void ParseCase(byte[] data, ulong index, Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.G> cases, out bool success, out _107_Common____GenericMarshalling__i_Compile.V v, out ulong rest__index)
    {
      success = false;
v = @_107_Common____GenericMarshalling__i_Compile.V.Default;
rest__index = 0;
      bool _1624_some1;
_107_Common____GenericMarshalling__i_Compile.V _1625_caseID;
ulong _1626_rest1;
bool _out18;
_107_Common____GenericMarshalling__i_Compile.V _out19;
ulong _out20;
_107_Common____GenericMarshalling__i_Compile.__default.ParseUint64(data, index, out _out18, out _out19, out _out20);
_1624_some1 = _out18;
_1625_caseID = _out19;
_1626_rest1 = _out20;
      if ((_1624_some1) && (((_1625_caseID).dtor_u) < ((ulong)(cases).LongCount))) {
        bool _1627_some2;
_107_Common____GenericMarshalling__i_Compile.V _1628_val;
ulong _1629_rest2;
bool _out21;
_107_Common____GenericMarshalling__i_Compile.V _out22;
ulong _out23;
_107_Common____GenericMarshalling__i_Compile.__default.ParseVal(data, _1626_rest1, (cases).Select((_1625_caseID).dtor_u), out _out21, out _out22, out _out23);
_1627_some2 = _out21;
_1628_val = _out22;
_1629_rest2 = _out23;
        if (_1627_some2) {
          success = true;
          v = @_107_Common____GenericMarshalling__i_Compile.V.create_VCase((_1625_caseID).dtor_u, _1628_val);
          rest__index = _1629_rest2;
        } else {
          success = false;
          rest__index = (ulong)(data).LongLength;
        }
      } else {
        success = false;
        rest__index = (ulong)(data).LongLength;
      }
    }
    public static _System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>> parse__Val(Dafny.Sequence<byte> data, _107_Common____GenericMarshalling__i_Compile.G grammar)
    {
      _107_Common____GenericMarshalling__i_Compile.G _source0 = grammar;
if (_source0.is_GUint64) {
        return _107_Common____GenericMarshalling__i_Compile.__default.parse__Uint64(data);
      } else if (_source0.is_GArray) {
        _107_Common____GenericMarshalling__i_Compile.G _1630_elt = ((_107_Common____GenericMarshalling__i_Compile.G_GArray)_source0).elt;
return _107_Common____GenericMarshalling__i_Compile.__default.parse__Array(data, _1630_elt);
      } else if (_source0.is_GTuple) {
        Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.G> _1631_t = ((_107_Common____GenericMarshalling__i_Compile.G_GTuple)_source0).t;
return _107_Common____GenericMarshalling__i_Compile.__default.parse__Tuple(data, _1631_t);
      } else if (_source0.is_GByteArray) {
        return _107_Common____GenericMarshalling__i_Compile.__default.parse__ByteArray(data);
      } else {
        Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.G> _1632_cases = ((_107_Common____GenericMarshalling__i_Compile.G_GTaggedUnion)_source0).cases;
return _107_Common____GenericMarshalling__i_Compile.__default.parse__Case(data, _1632_cases);
      }
    }
    public static void ParseVal(byte[] data, ulong index, _107_Common____GenericMarshalling__i_Compile.G grammar, out bool success, out _107_Common____GenericMarshalling__i_Compile.V v, out ulong rest__index)
    {
      success = false;
v = @_107_Common____GenericMarshalling__i_Compile.V.Default;
rest__index = 0;
      { }
      _107_Common____GenericMarshalling__i_Compile.G _source1 = grammar;
if (_source1.is_GUint64) {
        bool _out24;
_107_Common____GenericMarshalling__i_Compile.V _out25;
ulong _out26;
_107_Common____GenericMarshalling__i_Compile.__default.ParseUint64(data, index, out _out24, out _out25, out _out26);
success = _out24;
v = _out25;
rest__index = _out26;
      } else if (_source1.is_GArray) {
        _107_Common____GenericMarshalling__i_Compile.G _1633_elt = ((_107_Common____GenericMarshalling__i_Compile.G_GArray)_source1).elt;
        bool _out27;
_107_Common____GenericMarshalling__i_Compile.V _out28;
ulong _out29;
_107_Common____GenericMarshalling__i_Compile.__default.ParseArray(data, index, _1633_elt, out _out27, out _out28, out _out29);
success = _out27;
v = _out28;
rest__index = _out29;
      } else if (_source1.is_GTuple) {
        Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.G> _1634_t = ((_107_Common____GenericMarshalling__i_Compile.G_GTuple)_source1).t;
        bool _out30;
_107_Common____GenericMarshalling__i_Compile.V _out31;
ulong _out32;
_107_Common____GenericMarshalling__i_Compile.__default.ParseTuple(data, index, _1634_t, out _out30, out _out31, out _out32);
success = _out30;
v = _out31;
rest__index = _out32;
      } else if (_source1.is_GByteArray) {
        bool _out33;
_107_Common____GenericMarshalling__i_Compile.V _out34;
ulong _out35;
_107_Common____GenericMarshalling__i_Compile.__default.ParseByteArray(data, index, out _out33, out _out34, out _out35);
success = _out33;
v = _out34;
rest__index = _out35;
      } else {
        Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.G> _1635_cases = ((_107_Common____GenericMarshalling__i_Compile.G_GTaggedUnion)_source1).cases;
        bool _out36;
_107_Common____GenericMarshalling__i_Compile.V _out37;
ulong _out38;
_107_Common____GenericMarshalling__i_Compile.__default.ParseCase(data, index, _1635_cases, out _out36, out _out37, out _out38);
success = _out36;
v = _out37;
rest__index = _out38;
      }
    }
    public static void Demarshall(byte[] data, _107_Common____GenericMarshalling__i_Compile.G grammar, out bool success, out _107_Common____GenericMarshalling__i_Compile.V v)
    {
      success = false;
v = @_107_Common____GenericMarshalling__i_Compile.V.Default;
      ulong _1636_rest = 0;
      bool _out39;
_107_Common____GenericMarshalling__i_Compile.V _out40;
ulong _out41;
_107_Common____GenericMarshalling__i_Compile.__default.ParseVal(data, 0UL, grammar, out _out39, out _out40, out _out41);
success = _out39;
v = _out40;
_1636_rest = _out41;
      if ((success) && ((_1636_rest) == ((ulong)(data).LongLength))) {
        { }
        { }
        { }
      } else {
        success = false;
        { }
      }
    }
    public static void ComputeSeqSum(Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V> s, out ulong size)
    {
      size = 0;
      { }
      if (((ulong)(s).LongCount) == (0UL)) {
        size = 0UL;
      } else {
        ulong _1637_v__size;
ulong _out42;
_107_Common____GenericMarshalling__i_Compile.__default.ComputeSizeOf((s).Select((ulong)(0UL)), out _out42);
_1637_v__size = _out42;
        ulong _1638_rest__size;
ulong _out43;
_107_Common____GenericMarshalling__i_Compile.__default.ComputeSeqSum((s).Drop((ulong)(1UL)), out _out43);
_1638_rest__size = _out43;
        size = (_1637_v__size) + (_1638_rest__size);
      }
    }
    public static void ComputeSizeOf(_107_Common____GenericMarshalling__i_Compile.V val, out ulong size)
    {
      size = 0;
      _107_Common____GenericMarshalling__i_Compile.V _source2 = val;
if (_source2.is_VUint64) {
        ulong _1639___v3 = ((_107_Common____GenericMarshalling__i_Compile.V_VUint64)_source2).u;
        size = 8UL;
      } else if (_source2.is_VArray) {
        Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V> _1640_a = ((_107_Common____GenericMarshalling__i_Compile.V_VArray)_source2).a;
        ulong _1641_v;
ulong _out44;
_107_Common____GenericMarshalling__i_Compile.__default.ComputeSeqSum(_1640_a, out _out44);
_1641_v = _out44;
        if ((_1641_v) == (0UL)) {
          size = 8UL;
        } else {
          size = (8UL) + (_1641_v);
        }
      } else if (_source2.is_VTuple) {
        Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V> _1642_t = ((_107_Common____GenericMarshalling__i_Compile.V_VTuple)_source2).t;
        ulong _out45;
_107_Common____GenericMarshalling__i_Compile.__default.ComputeSeqSum(_1642_t, out _out45);
size = _out45;
      } else if (_source2.is_VByteArray) {
        Dafny.Sequence<byte> _1643_b = ((_107_Common____GenericMarshalling__i_Compile.V_VByteArray)_source2).b;
        size = (8UL) + ((ulong)(_1643_b).LongCount);
      } else {
        ulong _1644_c = ((_107_Common____GenericMarshalling__i_Compile.V_VCase)_source2).c;
_107_Common____GenericMarshalling__i_Compile.V _1645_v = ((_107_Common____GenericMarshalling__i_Compile.V_VCase)_source2).val;
        ulong _1646_vs;
ulong _out46;
_107_Common____GenericMarshalling__i_Compile.__default.ComputeSizeOf(_1645_v, out _out46);
_1646_vs = _out46;
        size = (8UL) + (_1646_vs);
      }
    }
    public static void MarshallUint64(ulong n, byte[] data, ulong index)
    {
    TAIL_CALL_START: ;
      _System.Tuple2<_44_Logic____Option__i_Compile.Option<_107_Common____GenericMarshalling__i_Compile.V>,Dafny.Sequence<byte>> _1647_tuple;
      _1647_tuple = _107_Common____GenericMarshalling__i_Compile.__default.parse__Uint64(Dafny.Helpers.SeqFromArray(data).Drop(index));
      _105_Common____MarshallInt__i_Compile.__default.MarshallUint64__guts(n, data, index);
    }
    public static void MarshallArrayContents(Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V> contents, _107_Common____GenericMarshalling__i_Compile.G eltType, byte[] data, ulong index, out ulong size)
    {
      size = 0;
      ulong _1648_i;
      _1648_i = 0UL;
      ulong _1649_cur__index;
      _1649_cur__index = index;
      { }
      { }
      { }
      { }
      while ((_1648_i) < ((ulong)(contents).LongCount)) {
        { }
        { }
        ulong _1650_item__size;
ulong _out47;
_107_Common____GenericMarshalling__i_Compile.__default.MarshallVal((contents).Select(_1648_i), eltType, data, _1649_cur__index, out _out47);
_1650_item__size = _out47;
        { }
        { }
        { }
        { }
        { }
        { }
        _1649_cur__index = (_1649_cur__index) + (_1650_item__size);
        _1648_i = (_1648_i) + (1UL);
        { }
        { }
        { }
        { }
        { }
        { }
      }
      { }
      { }
      { }
      { }
      size = (_1649_cur__index) - (index);
    }
    public static void MarshallArray(_107_Common____GenericMarshalling__i_Compile.V val, _107_Common____GenericMarshalling__i_Compile.G grammar, byte[] data, ulong index, out ulong size)
    {
      size = 0;
      { }
      _107_Common____GenericMarshalling__i_Compile.__default.MarshallUint64((ulong)((val).dtor_a).LongCount, data, index);
      { }
      { }
      { }
      { }
      ulong _1651_contents__size;
ulong _out48;
_107_Common____GenericMarshalling__i_Compile.__default.MarshallArrayContents((val).dtor_a, (grammar).dtor_elt, data, (index) + (_54_Native____NativeTypes__i_Compile.__default.Uint64Size()), out _out48);
_1651_contents__size = _out48;
      { }
      { }
      { }
      { }
      { }
      { }
      { }
      { }
      { }
      size = (8UL) + (_1651_contents__size);
    }
    public static void MarshallTupleContents(Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V> contents, Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.G> eltTypes, byte[] data, ulong index, out ulong size)
    {
      size = 0;
      ulong _1652_i;
      _1652_i = 0UL;
      ulong _1653_cur__index;
      _1653_cur__index = index;
      { }
      { }
      { }
      { }
      while ((_1652_i) < ((ulong)(contents).LongCount)) {
        { }
        { }
        { }
        { }
        { }
        ulong _1654_item__size;
ulong _out49;
_107_Common____GenericMarshalling__i_Compile.__default.MarshallVal((contents).Select(_1652_i), (eltTypes).Select(_1652_i), data, _1653_cur__index, out _out49);
_1654_item__size = _out49;
        { }
        { }
        { }
        { }
        { }
        { }
        _1653_cur__index = (_1653_cur__index) + (_1654_item__size);
        _1652_i = (_1652_i) + (1UL);
        { }
        { }
        { }
        { }
        { }
      }
      { }
      { }
      { }
      { }
      size = (_1653_cur__index) - (index);
    }
    public static void MarshallTuple(_107_Common____GenericMarshalling__i_Compile.V val, _107_Common____GenericMarshalling__i_Compile.G grammar, byte[] data, ulong index, out ulong size)
    {
      size = 0;
      ulong _out50;
_107_Common____GenericMarshalling__i_Compile.__default.MarshallTupleContents((val).dtor_t, (grammar).dtor_t, data, index, out _out50);
size = _out50;
      { }
    }
    public static void MarshallBytes(Dafny.Sequence<byte> bytes, byte[] data, ulong index)
    {
    TAIL_CALL_START: ;
      _9_Native____Io__s_Compile.Arrays.CopySeqIntoArray<byte>(bytes, 0UL, data, index, (ulong)(bytes).LongCount);
    }
    public static void MarshallByteArray(_107_Common____GenericMarshalling__i_Compile.V val, _107_Common____GenericMarshalling__i_Compile.G grammar, byte[] data, ulong index, out ulong size)
    {
      size = 0;
      _107_Common____GenericMarshalling__i_Compile.__default.MarshallUint64((ulong)((val).dtor_b).LongCount, data, index);
      { }
      _107_Common____GenericMarshalling__i_Compile.__default.MarshallBytes((val).dtor_b, data, (index) + (8UL));
      { }
      { }
      { }
      { }
      { }
      { }
      { }
      { }
      { }
      size = (8UL) + ((ulong)((val).dtor_b).LongCount);
    }
    public static void MarshallCase(_107_Common____GenericMarshalling__i_Compile.V val, _107_Common____GenericMarshalling__i_Compile.G grammar, byte[] data, ulong index, out ulong size)
    {
      size = 0;
      _107_Common____GenericMarshalling__i_Compile.__default.MarshallUint64((val).dtor_c, data, index);
      { }
      { }
      { }
      { }
      { }
      { }
      ulong _1655_val__size;
ulong _out51;
_107_Common____GenericMarshalling__i_Compile.__default.MarshallVal((val).dtor_val, ((grammar).dtor_cases).Select((val).dtor_c), data, (index) + (8UL), out _out51);
_1655_val__size = _out51;
      { }
      { }
      { }
      { }
      { }
      { }
      { }
      { }
      { }
      { }
      { }
      { }
      { }
      { }
      { }
      { }
      { }
      { }
      { }
      { }
      size = (8UL) + (_1655_val__size);
    }
    public static void MarshallVUint64(_107_Common____GenericMarshalling__i_Compile.V val, _107_Common____GenericMarshalling__i_Compile.G grammar, byte[] data, ulong index, out ulong size)
    {
    TAIL_CALL_START: ;
size = 0;
      _107_Common____GenericMarshalling__i_Compile.__default.MarshallUint64((val).dtor_u, data, index);
      { }
      size = 8UL;
return;
    }
    public static void MarshallVal(_107_Common____GenericMarshalling__i_Compile.V val, _107_Common____GenericMarshalling__i_Compile.G grammar, byte[] data, ulong index, out ulong size)
    {
      size = 0;
      _107_Common____GenericMarshalling__i_Compile.V _source3 = val;
if (_source3.is_VUint64) {
        ulong _1656___v4 = ((_107_Common____GenericMarshalling__i_Compile.V_VUint64)_source3).u;
        ulong _out52;
_107_Common____GenericMarshalling__i_Compile.__default.MarshallVUint64(val, grammar, data, index, out _out52);
size = _out52;
      } else if (_source3.is_VArray) {
        Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V> _1657___v5 = ((_107_Common____GenericMarshalling__i_Compile.V_VArray)_source3).a;
        ulong _out53;
_107_Common____GenericMarshalling__i_Compile.__default.MarshallArray(val, grammar, data, index, out _out53);
size = _out53;
      } else if (_source3.is_VTuple) {
        Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.V> _1658___v6 = ((_107_Common____GenericMarshalling__i_Compile.V_VTuple)_source3).t;
        ulong _out54;
_107_Common____GenericMarshalling__i_Compile.__default.MarshallTuple(val, grammar, data, index, out _out54);
size = _out54;
      } else if (_source3.is_VByteArray) {
        Dafny.Sequence<byte> _1659___v7 = ((_107_Common____GenericMarshalling__i_Compile.V_VByteArray)_source3).b;
        ulong _out55;
_107_Common____GenericMarshalling__i_Compile.__default.MarshallByteArray(val, grammar, data, index, out _out55);
size = _out55;
      } else {
        ulong _1660___v8 = ((_107_Common____GenericMarshalling__i_Compile.V_VCase)_source3).c;
_107_Common____GenericMarshalling__i_Compile.V _1661___v9 = ((_107_Common____GenericMarshalling__i_Compile.V_VCase)_source3).val;
        ulong _out56;
_107_Common____GenericMarshalling__i_Compile.__default.MarshallCase(val, grammar, data, index, out _out56);
size = _out56;
      }
    }
    public static void Marshall(_107_Common____GenericMarshalling__i_Compile.V val, _107_Common____GenericMarshalling__i_Compile.G grammar, out byte[] data)
    {
    TAIL_CALL_START: ;
data = new byte[0];
      ulong _1662_size;
ulong _out57;
_107_Common____GenericMarshalling__i_Compile.__default.ComputeSizeOf(val, out _out57);
_1662_size = _out57;
      var _nw5 = new byte[(int)(_1662_size)];
      data = _nw5;
      ulong _1663_computed__size;
ulong _out58;
_107_Common____GenericMarshalling__i_Compile.__default.MarshallVal(val, grammar, data, 0UL, out _out58);
_1663_computed__size = _out58;
      { }
      { }
    }
  }
} // end of namespace _107_Common____GenericMarshalling__i_Compile
namespace _111_PacketParsing__i_Compile {




  public partial class __default {
    public static _107_Common____GenericMarshalling__i_Compile.G CMessageTransferGrammar() {
      return @_107_Common____GenericMarshalling__i_Compile.G.create_GUint64();
    }
    public static _107_Common____GenericMarshalling__i_Compile.G CMessageLockedGrammar() {
      return @_107_Common____GenericMarshalling__i_Compile.G.create_GUint64();
    }
    public static _107_Common____GenericMarshalling__i_Compile.G CMessageGrammar() {
      return @_107_Common____GenericMarshalling__i_Compile.G.create_GTaggedUnion(Dafny.Sequence<_107_Common____GenericMarshalling__i_Compile.G>.FromElements(_111_PacketParsing__i_Compile.__default.CMessageTransferGrammar(), _111_PacketParsing__i_Compile.__default.CMessageLockedGrammar()));
    }
    public static _39_Message__i_Compile.CMessage ParseCMessageTransfer(_107_Common____GenericMarshalling__i_Compile.V val) {
      return @_39_Message__i_Compile.CMessage.create_CTransfer((val).dtor_u);
    }
    public static _39_Message__i_Compile.CMessage ParseCMessageLocked(_107_Common____GenericMarshalling__i_Compile.V val) {
      return @_39_Message__i_Compile.CMessage.create_CLocked((val).dtor_u);
    }
    public static _39_Message__i_Compile.CMessage ParseCMessage(_107_Common____GenericMarshalling__i_Compile.V val) {
      if (((val).dtor_c) == (0UL)) {
        return _111_PacketParsing__i_Compile.__default.ParseCMessageTransfer((val).dtor_val);
      } else  {
        return _111_PacketParsing__i_Compile.__default.ParseCMessageLocked((val).dtor_val);
      }
    }
    public static void DemarshallDataMethod(byte[] data, out _39_Message__i_Compile.CMessage msg)
    {
    TAIL_CALL_START: ;
msg = @_39_Message__i_Compile.CMessage.Default;
      bool _1664_success;
_107_Common____GenericMarshalling__i_Compile.V _1665_val;
bool _out59;
_107_Common____GenericMarshalling__i_Compile.V _out60;
_107_Common____GenericMarshalling__i_Compile.__default.Demarshall(data, _111_PacketParsing__i_Compile.__default.CMessageGrammar(), out _out59, out _out60);
_1664_success = _out59;
_1665_val = _out60;
      if (_1664_success) {
        msg = _111_PacketParsing__i_Compile.__default.ParseCMessage(_1665_val);
        { }
      } else {
        msg = @_39_Message__i_Compile.CMessage.create_CInvalid();
      }
    }
    public static void MarshallMessageTransfer(_39_Message__i_Compile.CMessage c, out _107_Common____GenericMarshalling__i_Compile.V val)
    {
    TAIL_CALL_START: ;
val = @_107_Common____GenericMarshalling__i_Compile.V.Default;
      val = @_107_Common____GenericMarshalling__i_Compile.V.create_VUint64((c).dtor_transfer__epoch);
    }
    public static void MarshallMessageLocked(_39_Message__i_Compile.CMessage c, out _107_Common____GenericMarshalling__i_Compile.V val)
    {
    TAIL_CALL_START: ;
val = @_107_Common____GenericMarshalling__i_Compile.V.Default;
      val = @_107_Common____GenericMarshalling__i_Compile.V.create_VUint64((c).dtor_locked__epoch);
    }
    public static void MarshallMessage(_39_Message__i_Compile.CMessage c, out _107_Common____GenericMarshalling__i_Compile.V val)
    {
      val = @_107_Common____GenericMarshalling__i_Compile.V.Default;
      if ((c).is_CTransfer) {
        _107_Common____GenericMarshalling__i_Compile.V _1666_msg;
_107_Common____GenericMarshalling__i_Compile.V _out61;
_111_PacketParsing__i_Compile.__default.MarshallMessageTransfer(c, out _out61);
_1666_msg = _out61;
        val = @_107_Common____GenericMarshalling__i_Compile.V.create_VCase(0UL, _1666_msg);
      } else if ((c).is_CLocked) {
        _107_Common____GenericMarshalling__i_Compile.V _1667_msg;
_107_Common____GenericMarshalling__i_Compile.V _out62;
_111_PacketParsing__i_Compile.__default.MarshallMessageLocked(c, out _out62);
_1667_msg = _out62;
        val = @_107_Common____GenericMarshalling__i_Compile.V.create_VCase(1UL, _1667_msg);
      } else { }
    }
    public static void MarshallLockMessage(_39_Message__i_Compile.CMessage msg, out byte[] data)
    {
    TAIL_CALL_START: ;
data = new byte[0];
      _107_Common____GenericMarshalling__i_Compile.V _1668_val;
_107_Common____GenericMarshalling__i_Compile.V _out63;
_111_PacketParsing__i_Compile.__default.MarshallMessage(msg, out _out63);
_1668_val = _out63;
      byte[] _out64;
_107_Common____GenericMarshalling__i_Compile.__default.Marshall(_1668_val, _111_PacketParsing__i_Compile.__default.CMessageGrammar(), out _out64);
data = _out64;
    }
  }
} // end of namespace _111_PacketParsing__i_Compile
namespace _113_Common____SeqIsUniqueDef__i_Compile {

} // end of namespace _113_Common____SeqIsUniqueDef__i_Compile
namespace _115_Impl__Node__i_Compile {







  public class CNode {
    public readonly bool held;
public readonly ulong epoch;
public readonly ulong my__index;
public readonly Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> config;
public CNode(bool held, ulong epoch, ulong my__index, Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> config) {
      this.held = held;
this.epoch = epoch;
this.my__index = my__index;
this.config = config;
    }
    public override bool Equals(object other) {
      var oth = other as _115_Impl__Node__i_Compile.CNode;
return oth != null && this.held == oth.held && this.epoch == oth.epoch && this.my__index == oth.my__index && Dafny.Helpers.AreEqual(this.config, oth.config);
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 0;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.held));
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.epoch));
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.my__index));
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.config));
return (int) hash;
    }
    public override string ToString() {
      string s = "_115_Impl__Node__i_Compile.CNode.CNode";
s += "(";
s += Dafny.Helpers.ToString(this.held);
s += ", ";
s += Dafny.Helpers.ToString(this.epoch);
s += ", ";
s += Dafny.Helpers.ToString(this.my__index);
s += ", ";
s += Dafny.Helpers.ToString(this.config);
s += ")";
return s;
    }
    static CNode theDefault;
public static CNode Default {
      get {
        if (theDefault == null) {
          theDefault = new _115_Impl__Node__i_Compile.CNode(false, 0, 0, Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint>.Empty);
        }
        return theDefault;
      }
    }
    public static CNode _DafnyDefaultValue() { return Default; }
public static CNode create(bool held, ulong epoch, ulong my__index, Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> config) {
      return new CNode(held, epoch, my__index, config);
    }
    public bool is_CNode { get { return true; } }
public bool dtor_held {
      get {
        return this.held;
      }
    }
    public ulong dtor_epoch {
      get {
        return this.epoch;
      }
    }
    public ulong dtor_my__index {
      get {
        return this.my__index;
      }
    }
    public Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> dtor_config {
      get {
        return this.config;
      }
    }
  }

  public partial class __default {
    public static void NodeInitImpl(ulong my__index, Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> config, out _115_Impl__Node__i_Compile.CNode node)
    {
    TAIL_CALL_START: ;
node = @_115_Impl__Node__i_Compile.CNode.Default;
      node = @_115_Impl__Node__i_Compile.CNode.create((my__index) == (0UL), ((my__index) == (0UL)) ? (1UL) : (0UL), my__index, config);
      if ((node).dtor_held) {
        Dafny.Helpers.Print(Dafny.Sequence<char>.FromString("I start holding the lock\n"));
      }
    }
    public static void NodeGrantImpl(_115_Impl__Node__i_Compile.CNode s, out _115_Impl__Node__i_Compile.CNode s_k, out _44_Logic____Option__i_Compile.Option<_7_Environment__s_Compile.LPacket<_9_Native____Io__s_Compile.EndPoint,_39_Message__i_Compile.CMessage>> packet)
    {
    TAIL_CALL_START: ;
s_k = @_115_Impl__Node__i_Compile.CNode.Default;
packet = @_44_Logic____Option__i_Compile.Option<_7_Environment__s_Compile.LPacket<_9_Native____Io__s_Compile.EndPoint,_39_Message__i_Compile.CMessage>>.Default;
      if (((s).dtor_held) && (((s).dtor_epoch) < (18446744073709551615UL))) {
        _115_Impl__Node__i_Compile.CNode _1669_ssss;
        _1669_ssss = @_115_Impl__Node__i_Compile.CNode.create(false, (s).dtor_epoch, (s).dtor_my__index, (s).dtor_config);
        s_k = _1669_ssss;
        ulong _1670_dst__index;
        _1670_dst__index = (((s).dtor_my__index) + (1UL)) % ((ulong)((s).dtor_config).LongCount);
        packet = @_44_Logic____Option__i_Compile.Option<_7_Environment__s_Compile.LPacket<_9_Native____Io__s_Compile.EndPoint,_39_Message__i_Compile.CMessage>>.create_Some(@_7_Environment__s_Compile.LPacket<_9_Native____Io__s_Compile.EndPoint,_39_Message__i_Compile.CMessage>.create(((s).dtor_config).Select(_1670_dst__index), ((s).dtor_config).Select((s).dtor_my__index), @_39_Message__i_Compile.CMessage.create_CTransfer(((s).dtor_epoch) + (1UL))));
        { }
        Dafny.Helpers.Print(Dafny.Sequence<char>.FromString("I grant the lock "));
Dafny.Helpers.Print((s).dtor_epoch);
Dafny.Helpers.Print(Dafny.Sequence<char>.FromString("\n"));
      } else {
        s_k = s;
        { }
        packet = @_44_Logic____Option__i_Compile.Option<_7_Environment__s_Compile.LPacket<_9_Native____Io__s_Compile.EndPoint,_39_Message__i_Compile.CMessage>>.create_None();
      }
    }
    public static void NodeAcceptImpl(_115_Impl__Node__i_Compile.CNode s, _7_Environment__s_Compile.LPacket<_9_Native____Io__s_Compile.EndPoint,_39_Message__i_Compile.CMessage> transfer__packet, out _115_Impl__Node__i_Compile.CNode s_k, out _44_Logic____Option__i_Compile.Option<_7_Environment__s_Compile.LPacket<_9_Native____Io__s_Compile.EndPoint,_39_Message__i_Compile.CMessage>> locked__packet)
    {
    TAIL_CALL_START: ;
s_k = @_115_Impl__Node__i_Compile.CNode.Default;
locked__packet = @_44_Logic____Option__i_Compile.Option<_7_Environment__s_Compile.LPacket<_9_Native____Io__s_Compile.EndPoint,_39_Message__i_Compile.CMessage>>.Default;
      { }
      if ((((!((s).dtor_held)) && (((s).dtor_config).Contains((transfer__packet).dtor_src))) && (((transfer__packet).dtor_msg).is_CTransfer)) && ((((transfer__packet).dtor_msg).dtor_transfer__epoch) > ((s).dtor_epoch))) {
        _115_Impl__Node__i_Compile.CNode _1671_ssss;
        _1671_ssss = @_115_Impl__Node__i_Compile.CNode.create(true, ((transfer__packet).dtor_msg).dtor_transfer__epoch, (s).dtor_my__index, (s).dtor_config);
        s_k = _1671_ssss;
        locked__packet = @_44_Logic____Option__i_Compile.Option<_7_Environment__s_Compile.LPacket<_9_Native____Io__s_Compile.EndPoint,_39_Message__i_Compile.CMessage>>.create_Some(@_7_Environment__s_Compile.LPacket<_9_Native____Io__s_Compile.EndPoint,_39_Message__i_Compile.CMessage>.create((transfer__packet).dtor_src, ((s).dtor_config).Select((s).dtor_my__index), @_39_Message__i_Compile.CMessage.create_CLocked(((transfer__packet).dtor_msg).dtor_transfer__epoch)));
        { }
        Dafny.Helpers.Print(Dafny.Sequence<char>.FromString("I hold the lock!\n"));
      } else {
        s_k = s;
        locked__packet = @_44_Logic____Option__i_Compile.Option<_7_Environment__s_Compile.LPacket<_9_Native____Io__s_Compile.EndPoint,_39_Message__i_Compile.CMessage>>.create_None();
      }
    }
  }
} // end of namespace _115_Impl__Node__i_Compile
namespace _119_UdpLock__i_Compile {



  public abstract class ReceiveResult {
    public ReceiveResult() { }
static ReceiveResult theDefault;
public static ReceiveResult Default {
      get {
        if (theDefault == null) {
          theDefault = new _119_UdpLock__i_Compile.ReceiveResult_RRFail();
        }
        return theDefault;
      }
    }
    public static ReceiveResult _DafnyDefaultValue() { return Default; }
public static ReceiveResult create_RRFail() {
      return new ReceiveResult_RRFail();
    }
    public static ReceiveResult create_RRTimeout() {
      return new ReceiveResult_RRTimeout();
    }
    public static ReceiveResult create_RRPacket(_7_Environment__s_Compile.LPacket<_9_Native____Io__s_Compile.EndPoint,_39_Message__i_Compile.CMessage> cpacket) {
      return new ReceiveResult_RRPacket(cpacket);
    }
    public bool is_RRFail { get { return this is ReceiveResult_RRFail; } }
public bool is_RRTimeout { get { return this is ReceiveResult_RRTimeout; } }
public bool is_RRPacket { get { return this is ReceiveResult_RRPacket; } }
public _7_Environment__s_Compile.LPacket<_9_Native____Io__s_Compile.EndPoint,_39_Message__i_Compile.CMessage> dtor_cpacket {
      get {
        var d = this;
return ((ReceiveResult_RRPacket)d).cpacket; 
      }
    }
  }
  public class ReceiveResult_RRFail : ReceiveResult {
    public ReceiveResult_RRFail() {
    }
    public override bool Equals(object other) {
      var oth = other as _119_UdpLock__i_Compile.ReceiveResult_RRFail;
return oth != null;
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 0;
return (int) hash;
    }
    public override string ToString() {
      string s = "_119_UdpLock__i_Compile.ReceiveResult.RRFail";
return s;
    }
  }
  public class ReceiveResult_RRTimeout : ReceiveResult {
    public ReceiveResult_RRTimeout() {
    }
    public override bool Equals(object other) {
      var oth = other as _119_UdpLock__i_Compile.ReceiveResult_RRTimeout;
return oth != null;
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 1;
return (int) hash;
    }
    public override string ToString() {
      string s = "_119_UdpLock__i_Compile.ReceiveResult.RRTimeout";
return s;
    }
  }
  public class ReceiveResult_RRPacket : ReceiveResult {
    public readonly _7_Environment__s_Compile.LPacket<_9_Native____Io__s_Compile.EndPoint,_39_Message__i_Compile.CMessage> cpacket;
public ReceiveResult_RRPacket(_7_Environment__s_Compile.LPacket<_9_Native____Io__s_Compile.EndPoint,_39_Message__i_Compile.CMessage> cpacket) {
      this.cpacket = cpacket;
    }
    public override bool Equals(object other) {
      var oth = other as _119_UdpLock__i_Compile.ReceiveResult_RRPacket;
return oth != null && Dafny.Helpers.AreEqual(this.cpacket, oth.cpacket);
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 2;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.cpacket));
return (int) hash;
    }
    public override string ToString() {
      string s = "_119_UdpLock__i_Compile.ReceiveResult.RRPacket";
s += "(";
s += Dafny.Helpers.ToString(this.cpacket);
s += ")";
return s;
    }
  }

  public partial class __default {
    public static void GetEndPoint(_9_Native____Io__s_Compile.IPEndPoint ipe, out _9_Native____Io__s_Compile.EndPoint ep)
    {
    TAIL_CALL_START: ;
ep = @_9_Native____Io__s_Compile.EndPoint.Default;
      byte[] _1672_addr;
byte[] _out65;
(ipe).GetAddress(out _out65);
_1672_addr = _out65;
      ushort _1673_port;
      _1673_port = (ipe).GetPort();
      ep = @_9_Native____Io__s_Compile.EndPoint.create(Dafny.Helpers.SeqFromArray(_1672_addr), _1673_port);
    }
    public static void Receive(_9_Native____Io__s_Compile.UdpClient udpClient, _9_Native____Io__s_Compile.EndPoint localAddr, out _119_UdpLock__i_Compile.ReceiveResult rr)
    {
      rr = @_119_UdpLock__i_Compile.ReceiveResult.Default;
      int _1674_timeout;
      _1674_timeout = 0;
      { }
      bool _1675_ok;
bool _1676_timedOut;
_9_Native____Io__s_Compile.IPEndPoint _1677_remote;
byte[] _1678_buffer;
bool _out66;
bool _out67;
_9_Native____Io__s_Compile.IPEndPoint _out68;
byte[] _out69;
(udpClient).Receive(_1674_timeout, out _out66, out _out67, out _out68, out _out69);
_1675_ok = _out66;
_1676_timedOut = _out67;
_1677_remote = _out68;
_1678_buffer = _out69;
      if (!(_1675_ok)) {
        rr = @_119_UdpLock__i_Compile.ReceiveResult.create_RRFail();
        return;
      }
      if (_1676_timedOut) {
        rr = @_119_UdpLock__i_Compile.ReceiveResult.create_RRTimeout();
        { }
        return;
      }
      { }
      _39_Message__i_Compile.CMessage _1679_cmessage;
_39_Message__i_Compile.CMessage _out70;
_111_PacketParsing__i_Compile.__default.DemarshallDataMethod(_1678_buffer, out _out70);
_1679_cmessage = _out70;
      _9_Native____Io__s_Compile.EndPoint _1680_srcEp;
_9_Native____Io__s_Compile.EndPoint _out71;
_119_UdpLock__i_Compile.__default.GetEndPoint(_1677_remote, out _out71);
_1680_srcEp = _out71;
      _7_Environment__s_Compile.LPacket<_9_Native____Io__s_Compile.EndPoint,_39_Message__i_Compile.CMessage> _1681_cpacket;
      _1681_cpacket = @_7_Environment__s_Compile.LPacket<_9_Native____Io__s_Compile.EndPoint,_39_Message__i_Compile.CMessage>.create(localAddr, _1680_srcEp, _1679_cmessage);
      rr = @_119_UdpLock__i_Compile.ReceiveResult.create_RRPacket(_1681_cpacket);
    }
    public static void SendPacket(_9_Native____Io__s_Compile.UdpClient udpClient, _44_Logic____Option__i_Compile.Option<_7_Environment__s_Compile.LPacket<_9_Native____Io__s_Compile.EndPoint,_39_Message__i_Compile.CMessage>> opt__packet, out bool ok)
    {
      ok = false;
      { }
      ok = true;
      if ((opt__packet).is_None) {
      } else {
        _7_Environment__s_Compile.LPacket<_9_Native____Io__s_Compile.EndPoint,_39_Message__i_Compile.CMessage> _1682_cpacket;
        _1682_cpacket = (opt__packet).dtor_v;
        _9_Native____Io__s_Compile.EndPoint _1683_dstEp;
        _1683_dstEp = (_1682_cpacket).dtor_dst;
        byte[] _1684_dstAddrAry;
byte[] _out72;
_101_Common____Util__i_Compile.__default.seqToArrayOpt<byte>((_1683_dstEp).dtor_addr, out _out72);
_1684_dstAddrAry = _out72;
        _9_Native____Io__s_Compile.IPEndPoint _1685_remote = default(_9_Native____Io__s_Compile.IPEndPoint);
        bool _out73;
_9_Native____Io__s_Compile.IPEndPoint _out74;
_9_Native____Io__s_Compile.IPEndPoint.Construct(_1684_dstAddrAry, (_1683_dstEp).dtor_port, out _out73, out _out74);
ok = _out73;
_1685_remote = _out74;
        if (!(ok)) {
          return;
        }
        byte[] _1686_buffer;
byte[] _out75;
_111_PacketParsing__i_Compile.__default.MarshallLockMessage((_1682_cpacket).dtor_msg, out _out75);
_1686_buffer = _out75;
        bool _out76;
(udpClient).Send(_1685_remote, _1686_buffer, out _out76);
ok = _out76;
        if (!(ok)) {
          return;
        }
        { }
        { }
      }
    }
  }
} // end of namespace _119_UdpLock__i_Compile
namespace _121_NodeImpl__i_Compile {



  public partial class NodeImpl {
    public _115_Impl__Node__i_Compile.CNode node = @_115_Impl__Node__i_Compile.CNode.Default;
public _9_Native____Io__s_Compile.UdpClient udpClient = default(_9_Native____Io__s_Compile.UdpClient);
public _9_Native____Io__s_Compile.EndPoint localAddr = @_9_Native____Io__s_Compile.EndPoint.Default;
public void __ctor()
    {
      var _this = this;
    TAIL_CALL_START: ;
      (_this).udpClient = (_9_Native____Io__s_Compile.UdpClient)null;
    }
    public void ConstructUdpClient(_9_Native____Io__s_Compile.EndPoint me, out bool ok, out _9_Native____Io__s_Compile.UdpClient client)
    {
      var _this = this;
    TAIL_CALL_START: ;
ok = false;
client = default(_9_Native____Io__s_Compile.UdpClient);
      _9_Native____Io__s_Compile.EndPoint _1687_my__ep;
      _1687_my__ep = me;
      byte[] _1688_ip__byte__array;
      var _nw6 = new byte[(int)(new BigInteger(((_1687_my__ep).dtor_addr).Count))];
      _1688_ip__byte__array = _nw6;
      _101_Common____Util__i_Compile.__default.seqIntoArrayOpt<byte>((_1687_my__ep).dtor_addr, _1688_ip__byte__array);
      _9_Native____Io__s_Compile.IPEndPoint _1689_ip__endpoint = default(_9_Native____Io__s_Compile.IPEndPoint);
      bool _out77;
_9_Native____Io__s_Compile.IPEndPoint _out78;
_9_Native____Io__s_Compile.IPEndPoint.Construct(_1688_ip__byte__array, (_1687_my__ep).dtor_port, out _out77, out _out78);
ok = _out77;
_1689_ip__endpoint = _out78;
      if (!(ok)) {
        return;
      }
      bool _out79;
_9_Native____Io__s_Compile.UdpClient _out80;
_9_Native____Io__s_Compile.UdpClient.Construct(_1689_ip__endpoint, out _out79, out _out80);
ok = _out79;
client = _out80;
      { }
    }
    public void InitNode(Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> config, ulong my__index, out bool ok)
    {
      var _this = this;
    TAIL_CALL_START: ;
ok = false;
      bool _out81;
_9_Native____Io__s_Compile.UdpClient _out82;
(_this).ConstructUdpClient((config).Select(my__index), out _out81, out _out82);
ok = _out81;
(_this).udpClient = _out82;
      if (ok) {
        _115_Impl__Node__i_Compile.CNode _out83;
_115_Impl__Node__i_Compile.__default.NodeInitImpl(my__index, config, out _out83);
(_this).node = _out83;
        { }
        (_this).localAddr = ((_this.node).dtor_config).Select(my__index);
        { }
      }
    }
    public void NodeNextGrant(out bool ok)
    {
      ok = false;
      _44_Logic____Option__i_Compile.Option<_7_Environment__s_Compile.LPacket<_9_Native____Io__s_Compile.EndPoint,_39_Message__i_Compile.CMessage>> _1690_transfer__packet = @_44_Logic____Option__i_Compile.Option<_7_Environment__s_Compile.LPacket<_9_Native____Io__s_Compile.EndPoint,_39_Message__i_Compile.CMessage>>.Default;
      _115_Impl__Node__i_Compile.CNode _out84;
_44_Logic____Option__i_Compile.Option<_7_Environment__s_Compile.LPacket<_9_Native____Io__s_Compile.EndPoint,_39_Message__i_Compile.CMessage>> _out85;
_115_Impl__Node__i_Compile.__default.NodeGrantImpl(this.node, out _out84, out _out85);
(this).node = _out84;
_1690_transfer__packet = _out85;
      ok = true;
      if ((_1690_transfer__packet).is_Some) {
        { }
        bool _out86;
_119_UdpLock__i_Compile.__default.SendPacket(this.udpClient, _1690_transfer__packet, out _out86);
ok = _out86;
        { }
      } else { }
    }
    public void NodeNextAccept(out bool ok)
    {
      var _this = this;
    TAIL_CALL_START: ;
ok = false;
      _119_UdpLock__i_Compile.ReceiveResult _1691_rr = @_119_UdpLock__i_Compile.ReceiveResult.Default;
      { }
      _119_UdpLock__i_Compile.ReceiveResult _out87;
_119_UdpLock__i_Compile.__default.Receive(_this.udpClient, _this.localAddr, out _out87);
_1691_rr = _out87;
      { }
      if ((_1691_rr).is_RRFail) {
        ok = false;
        return;
      } else if ((_1691_rr).is_RRTimeout) {
        ok = true;
        { }
        return;
      } else {
        ok = true;
        _44_Logic____Option__i_Compile.Option<_7_Environment__s_Compile.LPacket<_9_Native____Io__s_Compile.EndPoint,_39_Message__i_Compile.CMessage>> _1692_locked__packet = @_44_Logic____Option__i_Compile.Option<_7_Environment__s_Compile.LPacket<_9_Native____Io__s_Compile.EndPoint,_39_Message__i_Compile.CMessage>>.Default;
        _115_Impl__Node__i_Compile.CNode _out88;
_44_Logic____Option__i_Compile.Option<_7_Environment__s_Compile.LPacket<_9_Native____Io__s_Compile.EndPoint,_39_Message__i_Compile.CMessage>> _out89;
_115_Impl__Node__i_Compile.__default.NodeAcceptImpl(_this.node, (_1691_rr).dtor_cpacket, out _out88, out _out89);
(_this).node = _out88;
_1692_locked__packet = _out89;
        if ((_1692_locked__packet).is_Some) {
          { }
          bool _out90;
_119_UdpLock__i_Compile.__default.SendPacket(_this.udpClient, _1692_locked__packet, out _out90);
ok = _out90;
          { }
        }
      }
    }
    public void HostNextMain(out bool ok)
    {
      var _this = this;
    TAIL_CALL_START: ;
ok = false;
      if ((_this.node).dtor_held) {
        bool _out91;
(_this).NodeNextGrant(out _out91);
ok = _out91;
      } else {
        bool _out92;
(_this).NodeNextAccept(out _out92);
ok = _out92;
      }
    }
  }

} // end of namespace _121_NodeImpl__i_Compile
namespace _127_CmdLineParser__i_Compile {





  public partial class __default {
    public static _System.Tuple2<bool,byte> ascii__to__int(ushort @short) {
      if (((48) <= (@short)) && ((@short) <= (57))) {
        return @_System.Tuple2<bool,byte>.create(true, (byte)((ushort)((@short) - (48))));
      } else  {
        return @_System.Tuple2<bool,byte>.create(false, 0);
      }
    }
    public static void power10(BigInteger e, out BigInteger val)
    {
      val = BigInteger.Zero;
      { }
      if ((e) == (new BigInteger(0))) {
        val = new BigInteger(1);
return;
      } else {
        BigInteger _1693_tmp;
BigInteger _out93;
_127_CmdLineParser__i_Compile.__default.power10((e) - (new BigInteger(1)), out _out93);
_1693_tmp = _out93;
        val = (new BigInteger(10)) * (_1693_tmp);
return;
      }
    }
    public static _System.Tuple2<bool,Dafny.Sequence<byte>> shorts__to__bytes(Dafny.Sequence<ushort> shorts) {
      if ((new BigInteger((shorts).Count)) == (new BigInteger(0))) {
        return @_System.Tuple2<bool,Dafny.Sequence<byte>>.create(true, Dafny.Sequence<byte>.FromElements());
      } else  {
        _System.Tuple2<bool,Dafny.Sequence<byte>> _1694_tuple = _127_CmdLineParser__i_Compile.__default.shorts__to__bytes((shorts).Drop(new BigInteger(1)));
bool _1695_ok = (_1694_tuple).dtor__0;
Dafny.Sequence<byte> _1696_rest = (_1694_tuple).dtor__1;
_System.Tuple2<bool,byte> _1697_tuple_k = _127_CmdLineParser__i_Compile.__default.ascii__to__int((shorts).Select(new BigInteger(0)));
bool _1698_ok_k = (_1697_tuple_k).dtor__0;
byte _1699_a__byte = (_1697_tuple_k).dtor__1;
if ((_1695_ok) && (_1698_ok_k)) {
          return @_System.Tuple2<bool,Dafny.Sequence<byte>>.create(true, (Dafny.Sequence<byte>.FromElements(_1699_a__byte)).Concat(_1696_rest));
        } else  {
          return @_System.Tuple2<bool,Dafny.Sequence<byte>>.create(false, Dafny.Sequence<byte>.FromElements());
        }
      }
    }
    public static BigInteger bytes__to__decimal(Dafny.Sequence<byte> bytes) {
      if ((new BigInteger((bytes).Count)) == (new BigInteger(0))) {
        return new BigInteger(0);
      } else  {
        return (new BigInteger((bytes).Select((new BigInteger((bytes).Count)) - (new BigInteger(1))))) + ((new BigInteger(10)) * (_127_CmdLineParser__i_Compile.__default.bytes__to__decimal((bytes).Take((new BigInteger((bytes).Count)) - (new BigInteger(1))).Drop(new BigInteger(0)))));
      }
    }
    public static _System.Tuple2<bool,BigInteger> shorts__to__nat(Dafny.Sequence<ushort> shorts) {
      if ((new BigInteger((shorts).Count)) == (new BigInteger(0))) {
        return @_System.Tuple2<bool,BigInteger>.create(false, new BigInteger(0));
      } else  {
        _System.Tuple2<bool,Dafny.Sequence<byte>> _1700_tuple = _127_CmdLineParser__i_Compile.__default.shorts__to__bytes(shorts);
bool _1701_ok = (_1700_tuple).dtor__0;
Dafny.Sequence<byte> _1702_bytes = (_1700_tuple).dtor__1;
if (!(_1701_ok)) {
          return @_System.Tuple2<bool,BigInteger>.create(false, new BigInteger(0));
        } else  {
          return @_System.Tuple2<bool,BigInteger>.create(true, _127_CmdLineParser__i_Compile.__default.bytes__to__decimal(_1702_bytes));
        }
      }
    }
    public static _System.Tuple2<bool,byte> shorts__to__byte(Dafny.Sequence<ushort> shorts) {
      _System.Tuple2<bool,BigInteger> _1703_tuple = _127_CmdLineParser__i_Compile.__default.shorts__to__nat(shorts);
bool _1704_ok = (_1703_tuple).dtor__0;
BigInteger _1705_val = (_1703_tuple).dtor__1;
if (((new BigInteger(0)) <= (_1705_val)) && ((_1705_val) < (new BigInteger(256)))) {
        return @_System.Tuple2<bool,byte>.create(true, (byte)(_1705_val));
      } else  {
        return @_System.Tuple2<bool,byte>.create(false, 0);
      }
    }
    public static _System.Tuple2<bool,ushort> shorts__to__uint16(Dafny.Sequence<ushort> shorts) {
      _System.Tuple2<bool,BigInteger> _1706_tuple = _127_CmdLineParser__i_Compile.__default.shorts__to__nat(shorts);
bool _1707_ok = (_1706_tuple).dtor__0;
BigInteger _1708_val = (_1706_tuple).dtor__1;
if (((new BigInteger(0)) <= (_1708_val)) && ((_1708_val) < (new BigInteger(65536)))) {
        return @_System.Tuple2<bool,ushort>.create(true, (ushort)(_1708_val));
      } else  {
        return @_System.Tuple2<bool,ushort>.create(false, 0);
      }
    }
    public static _System.Tuple2<bool,uint> shorts__to__uint32(Dafny.Sequence<ushort> shorts) {
      _System.Tuple2<bool,BigInteger> _1709_tuple = _127_CmdLineParser__i_Compile.__default.shorts__to__nat(shorts);
bool _1710_ok = (_1709_tuple).dtor__0;
BigInteger _1711_val = (_1709_tuple).dtor__1;
if (((new BigInteger(0)) <= (_1711_val)) && ((_1711_val) < (BigInteger.Parse("4294967296")))) {
        return @_System.Tuple2<bool,uint>.create(true, (uint)(_1711_val));
      } else  {
        return @_System.Tuple2<bool,uint>.create(false, 0U);
      }
    }
    public static bool is__ascii__period(ushort @short) {
      return (@short) == (46);
    }
    public static _System.Tuple2<bool,Dafny.Sequence<byte>> parse__ip__addr__helper(Dafny.Sequence<ushort> ip__shorts, Dafny.Sequence<ushort> current__octet__shorts)
    {
      if ((new BigInteger((ip__shorts).Count)) == (new BigInteger(0))) {
        _System.Tuple2<bool,byte> _1712_tuple = _127_CmdLineParser__i_Compile.__default.shorts__to__byte(current__octet__shorts);
bool _1713_okay = (_1712_tuple).dtor__0;
byte _1714_b = (_1712_tuple).dtor__1;
if (!(_1713_okay)) {
          return @_System.Tuple2<bool,Dafny.Sequence<byte>>.create(false, Dafny.Sequence<byte>.FromElements());
        } else  {
          return @_System.Tuple2<bool,Dafny.Sequence<byte>>.create(true, Dafny.Sequence<byte>.FromElements(_1714_b));
        }
      } else  {
        if (_127_CmdLineParser__i_Compile.__default.is__ascii__period((ip__shorts).Select(new BigInteger(0)))) {
          _System.Tuple2<bool,byte> _1715_tuple = _127_CmdLineParser__i_Compile.__default.shorts__to__byte(current__octet__shorts);
bool _1716_okay = (_1715_tuple).dtor__0;
byte _1717_b = (_1715_tuple).dtor__1;
if (!(_1716_okay)) {
            return @_System.Tuple2<bool,Dafny.Sequence<byte>>.create(false, Dafny.Sequence<byte>.FromElements());
          } else  {
            _System.Tuple2<bool,Dafny.Sequence<byte>> _1718_tuple_k = _127_CmdLineParser__i_Compile.__default.parse__ip__addr__helper((ip__shorts).Drop(new BigInteger(1)), Dafny.Sequence<ushort>.FromElements());
bool _1719_ok = (_1718_tuple_k).dtor__0;
Dafny.Sequence<byte> _1720_ip__bytes = (_1718_tuple_k).dtor__1;
if (!(_1719_ok)) {
              return @_System.Tuple2<bool,Dafny.Sequence<byte>>.create(false, Dafny.Sequence<byte>.FromElements());
            } else  {
              return @_System.Tuple2<bool,Dafny.Sequence<byte>>.create(true, (Dafny.Sequence<byte>.FromElements(_1717_b)).Concat(_1720_ip__bytes));
            }
          }
        } else  {
          return _127_CmdLineParser__i_Compile.__default.parse__ip__addr__helper((ip__shorts).Drop(new BigInteger(1)), (current__octet__shorts).Concat(Dafny.Sequence<ushort>.FromElements((ip__shorts).Select(new BigInteger(0)))));
        }
      }
    }
    public static _System.Tuple2<bool,Dafny.Sequence<byte>> parse__ip__addr(Dafny.Sequence<ushort> ip__shorts) {
      _System.Tuple2<bool,Dafny.Sequence<byte>> _1721_tuple = _127_CmdLineParser__i_Compile.__default.parse__ip__addr__helper(ip__shorts, Dafny.Sequence<ushort>.FromElements());
bool _1722_ok = (_1721_tuple).dtor__0;
Dafny.Sequence<byte> _1723_ip__bytes = (_1721_tuple).dtor__1;
if ((_1722_ok) && ((new BigInteger((_1723_ip__bytes).Count)) == (new BigInteger(4)))) {
        return @_System.Tuple2<bool,Dafny.Sequence<byte>>.create(true, _1723_ip__bytes);
      } else  {
        return @_System.Tuple2<bool,Dafny.Sequence<byte>>.create(false, Dafny.Sequence<byte>.FromElements());
      }
    }
    public static _System.Tuple2<bool,_9_Native____Io__s_Compile.EndPoint> parse__end__point(Dafny.Sequence<ushort> ip__shorts, Dafny.Sequence<ushort> port__shorts)
    {
      _System.Tuple2<bool,Dafny.Sequence<byte>> _1724_tuple = _127_CmdLineParser__i_Compile.__default.parse__ip__addr(ip__shorts);
bool _1725_okay = (_1724_tuple).dtor__0;
Dafny.Sequence<byte> _1726_ip__bytes = (_1724_tuple).dtor__1;
if (!(_1725_okay)) {
        return @_System.Tuple2<bool,_9_Native____Io__s_Compile.EndPoint>.create(false, @_9_Native____Io__s_Compile.EndPoint.create(Dafny.Sequence<byte>.FromElements(0, 0, 0, 0), 0));
      } else  {
        _System.Tuple2<bool,ushort> _1727_tuple_k = _127_CmdLineParser__i_Compile.__default.shorts__to__uint16(port__shorts);
bool _1728_okay_k = (_1727_tuple_k).dtor__0;
ushort _1729_port = (_1727_tuple_k).dtor__1;
if (!(_1728_okay_k)) {
          return @_System.Tuple2<bool,_9_Native____Io__s_Compile.EndPoint>.create(false, @_9_Native____Io__s_Compile.EndPoint.create(Dafny.Sequence<byte>.FromElements(0, 0, 0, 0), 0));
        } else  {
          return @_System.Tuple2<bool,_9_Native____Io__s_Compile.EndPoint>.create(true, @_9_Native____Io__s_Compile.EndPoint.create(_1726_ip__bytes, _1729_port));
        }
      }
    }
    public static void test__unique_k(Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> endpoints, out bool unique)
    {
    TAIL_CALL_START: ;
unique = false;
      unique = true;
      BigInteger _1730_i;
      _1730_i = new BigInteger(0);
      while ((_1730_i) < (new BigInteger((endpoints).Count))) {
        BigInteger _1731_j;
        _1731_j = new BigInteger(0);
        while ((_1731_j) < (new BigInteger((endpoints).Count))) {
          if (((_1730_i) != (_1731_j)) && (((endpoints).Select(_1730_i)).Equals((endpoints).Select(_1731_j)))) {
            unique = false;
            { }
            return;
          }
          _1731_j = (_1731_j) + (new BigInteger(1));
        }
        _1730_i = (_1730_i) + (new BigInteger(1));
      }
      { }
    }
    public static _System.Tuple2<bool,Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint>> parse__end__points(Dafny.Sequence<Dafny.Sequence<ushort>> args) {
      if ((new BigInteger((args).Count)) == (new BigInteger(0))) {
        return @_System.Tuple2<bool,Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint>>.create(true, Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint>.FromElements());
      } else  {
        _System.Tuple2<bool,_9_Native____Io__s_Compile.EndPoint> _let_tmp_rhs10 = _127_CmdLineParser__i_Compile.__default.parse__end__point((args).Select(new BigInteger(0)), (args).Select(new BigInteger(1)));
bool _1732_ok1 = ((_System.Tuple2<bool,_9_Native____Io__s_Compile.EndPoint>)_let_tmp_rhs10)._0;
_9_Native____Io__s_Compile.EndPoint _1733_ep = ((_System.Tuple2<bool,_9_Native____Io__s_Compile.EndPoint>)_let_tmp_rhs10)._1;
_System.Tuple2<bool,Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint>> _let_tmp_rhs11 = _127_CmdLineParser__i_Compile.__default.parse__end__points((args).Drop(new BigInteger(2)));
bool _1734_ok2 = ((_System.Tuple2<bool,Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint>>)_let_tmp_rhs11)._0;
Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> _1735_rest = ((_System.Tuple2<bool,Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint>>)_let_tmp_rhs11)._1;
if (!((_1732_ok1) && (_1734_ok2))) {
          return @_System.Tuple2<bool,Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint>>.create(false, Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint>.FromElements());
        } else  {
          return @_System.Tuple2<bool,Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint>>.create(true, (Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint>.FromElements(_1733_ep)).Concat(_1735_rest));
        }
      }
    }
    public static void collect__cmd__line__args(out Dafny.Sequence<Dafny.Sequence<ushort>> args)
    {
    TAIL_CALL_START: ;
args = Dafny.Sequence<Dafny.Sequence<ushort>>.Empty;
      uint _1736_num__args;
uint _out94;
_9_Native____Io__s_Compile.HostConstants.NumCommandLineArgs(out _out94);
_1736_num__args = _out94;
      uint _1737_i;
      _1737_i = 0U;
      args = Dafny.Sequence<Dafny.Sequence<ushort>>.FromElements();
      while ((_1737_i) < (_1736_num__args)) {
        ushort[] _1738_arg;
ushort[] _out95;
_9_Native____Io__s_Compile.HostConstants.GetCommandLineArg((ulong)(_1737_i), out _out95);
_1738_arg = _out95;
        args = (args).Concat(Dafny.Sequence<Dafny.Sequence<ushort>>.FromElements(Dafny.Helpers.SeqFromArray(_1738_arg)));
        _1737_i = (_1737_i) + (1U);
      }
    }
  }
} // end of namespace _127_CmdLineParser__i_Compile
namespace _129_LockCmdLineParser__i_Compile {


  public partial class __default {
    public static _9_Native____Io__s_Compile.EndPoint EndPointNull() {
      return @_9_Native____Io__s_Compile.EndPoint.create(Dafny.Sequence<byte>.FromElements(0, 0, 0, 0), 0);
    }
    public static void GetHostIndex(_9_Native____Io__s_Compile.EndPoint host, Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> hosts, out bool found, out ulong index)
    {
    TAIL_CALL_START: ;
found = false;
index = 0;
      ulong _1739_i;
      _1739_i = 0UL;
      while ((_1739_i) < ((ulong)(hosts).LongCount)) {
        if ((host).Equals((hosts).Select(_1739_i))) {
          found = true;
          index = _1739_i;
          { }
          return;
        }
        if ((_1739_i) == (((ulong)(hosts).LongCount) - (1UL))) {
          found = false;
          return;
        }
        _1739_i = (_1739_i) + (1UL);
      }
      found = false;
    }
    public static void ParseCmdLine(out bool ok, out Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> host__ids, out ulong my__index)
    {
      ok = false;
host__ids = Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint>.Empty;
my__index = 0;
      ok = false;
      uint _1740_num__args;
uint _out96;
_9_Native____Io__s_Compile.HostConstants.NumCommandLineArgs(out _out96);
_1740_num__args = _out96;
      if (((_1740_num__args) < (4U)) || (((_1740_num__args) % (2U)) != (1U))) {
        Dafny.Helpers.Print(Dafny.Sequence<char>.FromString("Incorrect number of command line arguments.\n"));
        Dafny.Helpers.Print(Dafny.Sequence<char>.FromString("Expected: ./Main.exe [IP port]+ [IP port]\n"));
        Dafny.Helpers.Print(Dafny.Sequence<char>.FromString("  where the final argument is one of the two IP-port pairs provided earlier \n"));
        return;
      }
      Dafny.Sequence<Dafny.Sequence<ushort>> _1741_args;
Dafny.Sequence<Dafny.Sequence<ushort>> _out97;
_127_CmdLineParser__i_Compile.__default.collect__cmd__line__args(out _out97);
_1741_args = _out97;
      { }
      _System.Tuple2<bool,Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint>> _1742_tuple1;
      _1742_tuple1 = _127_CmdLineParser__i_Compile.__default.parse__end__points((_1741_args).Take((new BigInteger((_1741_args).Count)) - (new BigInteger(2))).Drop(new BigInteger(1)));
      ok = (_1742_tuple1).dtor__0;
      Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> _1743_endpoints;
      _1743_endpoints = (_1742_tuple1).dtor__1;
      if (((!(ok)) || ((new BigInteger((_1743_endpoints).Count)) == (new BigInteger(0)))) || ((new BigInteger((_1743_endpoints).Count)) >= (BigInteger.Parse("18446744073709551616")))) {
        ok = false;
        return;
      }
      _System.Tuple2<bool,_9_Native____Io__s_Compile.EndPoint> _1744_tuple2;
      _1744_tuple2 = _127_CmdLineParser__i_Compile.__default.parse__end__point((_1741_args).Select((new BigInteger((_1741_args).Count)) - (new BigInteger(2))), (_1741_args).Select((new BigInteger((_1741_args).Count)) - (new BigInteger(1))));
      ok = (_1744_tuple2).dtor__0;
      if (!(ok)) {
        return;
      }
      bool _1745_unique;
bool _out98;
_127_CmdLineParser__i_Compile.__default.test__unique_k(_1743_endpoints, out _out98);
_1745_unique = _out98;
      if (!(_1745_unique)) {
        ok = false;
        return;
      }
      bool _out99;
ulong _out100;
_129_LockCmdLineParser__i_Compile.__default.GetHostIndex((_1744_tuple2).dtor__1, _1743_endpoints, out _out99, out _out100);
ok = _out99;
my__index = _out100;
      if (!(ok)) {
        return;
      }
      host__ids = _1743_endpoints;
      _9_Native____Io__s_Compile.EndPoint _1746_me;
      _1746_me = (_1743_endpoints).Select(my__index);
      { }
      { }
      { }
      { }
    }
  }
} // end of namespace _129_LockCmdLineParser__i_Compile
namespace _131_Host__i_Compile {




  public class CScheduler {
    public readonly _121_NodeImpl__i_Compile.NodeImpl node__impl;
public CScheduler(_121_NodeImpl__i_Compile.NodeImpl node__impl) {
      this.node__impl = node__impl;
    }
    public override bool Equals(object other) {
      var oth = other as _131_Host__i_Compile.CScheduler;
return oth != null && this.node__impl == oth.node__impl;
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 0;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.node__impl));
return (int) hash;
    }
    public override string ToString() {
      string s = "_131_Host__i_Compile.CScheduler.CScheduler";
s += "(";
s += Dafny.Helpers.ToString(this.node__impl);
s += ")";
return s;
    }
    static CScheduler theDefault;
public static CScheduler Default {
      get {
        if (theDefault == null) {
          theDefault = new _131_Host__i_Compile.CScheduler(default(_121_NodeImpl__i_Compile.NodeImpl));
        }
        return theDefault;
      }
    }
    public static CScheduler _DafnyDefaultValue() { return Default; }
public static CScheduler create(_121_NodeImpl__i_Compile.NodeImpl node__impl) {
      return new CScheduler(node__impl);
    }
    public bool is_CScheduler { get { return true; } }
public _121_NodeImpl__i_Compile.NodeImpl dtor_node__impl {
      get {
        return this.node__impl;
      }
    }
  }




  public partial class __default {
    public static void HostInitImpl(out bool ok, out _131_Host__i_Compile.CScheduler host__state, out Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> config, out _9_Native____Io__s_Compile.EndPoint id)
    {
    TAIL_CALL_START: ;
ok = false;
host__state = @_131_Host__i_Compile.CScheduler.Default;
config = Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint>.Empty;
id = @_9_Native____Io__s_Compile.EndPoint.Default;
      ulong _1747_my__index = 0;
      bool _out101;
Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> _out102;
ulong _out103;
_129_LockCmdLineParser__i_Compile.__default.ParseCmdLine(out _out101, out _out102, out _out103);
ok = _out101;
config = _out102;
_1747_my__index = _out103;
      if (!(ok)) {
        return;
      }
      id = (config).Select(_1747_my__index);
      _121_NodeImpl__i_Compile.NodeImpl _1748_node__impl;
      var _nw7 = new _121_NodeImpl__i_Compile.NodeImpl();
_nw7.__ctor();
      _1748_node__impl = _nw7;
      bool _out104;
(_1748_node__impl).InitNode(config, _1747_my__index, out _out104);
ok = _out104;
      if (!(ok)) {
        return;
      }
      host__state = @_131_Host__i_Compile.CScheduler.create(_1748_node__impl);
      { }
      { }
    }
    public static void HostNextImpl(_131_Host__i_Compile.CScheduler host__state, out bool ok, out _131_Host__i_Compile.CScheduler host__state_k)
    {
      ok = false;
host__state_k = @_131_Host__i_Compile.CScheduler.Default;
      bool _1749_okay;
bool _out105;
((host__state).dtor_node__impl).HostNextMain(out _out105);
_1749_okay = _out105;
      if (_1749_okay) {
        { }
        { }
        { }
        { }
        host__state_k = @_131_Host__i_Compile.CScheduler.create((host__state).dtor_node__impl);
      } else { }
      ok = _1749_okay;
    }
  }



} // end of namespace _131_Host__i_Compile
namespace _133_Lock__DistributedSystem__i_Compile {







  public class DS__State {
    public readonly Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> config;
public readonly _7_Environment__s_Compile.LEnvironment<_9_Native____Io__s_Compile.EndPoint,Dafny.Sequence<byte>,_33_Types__i_Compile.LockStep> environment;
public readonly Dafny.Map<_9_Native____Io__s_Compile.EndPoint,_131_Host__i_Compile.CScheduler> servers;
public readonly Dafny.Set<_9_Native____Io__s_Compile.EndPoint> clients;
public DS__State(Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> config, _7_Environment__s_Compile.LEnvironment<_9_Native____Io__s_Compile.EndPoint,Dafny.Sequence<byte>,_33_Types__i_Compile.LockStep> environment, Dafny.Map<_9_Native____Io__s_Compile.EndPoint,_131_Host__i_Compile.CScheduler> servers, Dafny.Set<_9_Native____Io__s_Compile.EndPoint> clients) {
      this.config = config;
this.environment = environment;
this.servers = servers;
this.clients = clients;
    }
    public override bool Equals(object other) {
      var oth = other as _133_Lock__DistributedSystem__i_Compile.DS__State;
return oth != null && Dafny.Helpers.AreEqual(this.config, oth.config) && Dafny.Helpers.AreEqual(this.environment, oth.environment) && Dafny.Helpers.AreEqual(this.servers, oth.servers) && Dafny.Helpers.AreEqual(this.clients, oth.clients);
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 0;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.config));
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.environment));
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.servers));
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.clients));
return (int) hash;
    }
    public override string ToString() {
      string s = "_133_Lock__DistributedSystem__i_Compile.DS_State.DS_State";
s += "(";
s += Dafny.Helpers.ToString(this.config);
s += ", ";
s += Dafny.Helpers.ToString(this.environment);
s += ", ";
s += Dafny.Helpers.ToString(this.servers);
s += ", ";
s += Dafny.Helpers.ToString(this.clients);
s += ")";
return s;
    }
    static DS__State theDefault;
public static DS__State Default {
      get {
        if (theDefault == null) {
          theDefault = new _133_Lock__DistributedSystem__i_Compile.DS__State(Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint>.Empty, @_7_Environment__s_Compile.LEnvironment<_9_Native____Io__s_Compile.EndPoint,Dafny.Sequence<byte>,_33_Types__i_Compile.LockStep>.Default, Dafny.Map<_9_Native____Io__s_Compile.EndPoint,_131_Host__i_Compile.CScheduler>.Empty, Dafny.Set<_9_Native____Io__s_Compile.EndPoint>.Empty);
        }
        return theDefault;
      }
    }
    public static DS__State _DafnyDefaultValue() { return Default; }
public static DS__State create(Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> config, _7_Environment__s_Compile.LEnvironment<_9_Native____Io__s_Compile.EndPoint,Dafny.Sequence<byte>,_33_Types__i_Compile.LockStep> environment, Dafny.Map<_9_Native____Io__s_Compile.EndPoint,_131_Host__i_Compile.CScheduler> servers, Dafny.Set<_9_Native____Io__s_Compile.EndPoint> clients) {
      return new DS__State(config, environment, servers, clients);
    }
    public bool is_DS__State { get { return true; } }
public Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> dtor_config {
      get {
        return this.config;
      }
    }
    public _7_Environment__s_Compile.LEnvironment<_9_Native____Io__s_Compile.EndPoint,Dafny.Sequence<byte>,_33_Types__i_Compile.LockStep> dtor_environment {
      get {
        return this.environment;
      }
    }
    public Dafny.Map<_9_Native____Io__s_Compile.EndPoint,_131_Host__i_Compile.CScheduler> dtor_servers {
      get {
        return this.servers;
      }
    }
    public Dafny.Set<_9_Native____Io__s_Compile.EndPoint> dtor_clients {
      get {
        return this.clients;
      }
    }
  }
} // end of namespace _133_Lock__DistributedSystem__i_Compile
namespace _138_Concrete__NodeIdentity__i_Compile {



} // end of namespace _138_Concrete__NodeIdentity__i_Compile
namespace _143_AbstractServiceLock__s_Compile {

  public class ServiceState_k {
    public readonly Dafny.Set<_9_Native____Io__s_Compile.EndPoint> hosts;
public readonly Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> history;
public ServiceState_k(Dafny.Set<_9_Native____Io__s_Compile.EndPoint> hosts, Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> history) {
      this.hosts = hosts;
this.history = history;
    }
    public override bool Equals(object other) {
      var oth = other as _143_AbstractServiceLock__s_Compile.ServiceState_k;
return oth != null && Dafny.Helpers.AreEqual(this.hosts, oth.hosts) && Dafny.Helpers.AreEqual(this.history, oth.history);
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 0;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.hosts));
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.history));
return (int) hash;
    }
    public override string ToString() {
      string s = "_143_AbstractServiceLock__s_Compile.ServiceState'.ServiceState'";
s += "(";
s += Dafny.Helpers.ToString(this.hosts);
s += ", ";
s += Dafny.Helpers.ToString(this.history);
s += ")";
return s;
    }
    static ServiceState_k theDefault;
public static ServiceState_k Default {
      get {
        if (theDefault == null) {
          theDefault = new _143_AbstractServiceLock__s_Compile.ServiceState_k(Dafny.Set<_9_Native____Io__s_Compile.EndPoint>.Empty, Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint>.Empty);
        }
        return theDefault;
      }
    }
    public static ServiceState_k _DafnyDefaultValue() { return Default; }
public static ServiceState_k create(Dafny.Set<_9_Native____Io__s_Compile.EndPoint> hosts, Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> history) {
      return new ServiceState_k(hosts, history);
    }
    public bool is_ServiceState_k { get { return true; } }
public Dafny.Set<_9_Native____Io__s_Compile.EndPoint> dtor_hosts {
      get {
        return this.hosts;
      }
    }
    public Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> dtor_history {
      get {
        return this.history;
      }
    }
  }





} // end of namespace _143_AbstractServiceLock__s_Compile
namespace _148_Common____SeqIsUnique__i_Compile {



  public partial class __default {
    public static void SeqToSetConstruct<X>(Dafny.Sequence<X> xs, out Dafny.Set<X> s)
    {
    TAIL_CALL_START: ;
s = Dafny.Set<X>.Empty;
      { }
      s = Dafny.Set<X>.FromElements();
      BigInteger _1750_i;
      _1750_i = new BigInteger(0);
      while ((_1750_i) < (new BigInteger((xs).Count))) {
        s = (s).Union(Dafny.Set<X>.FromElements((xs).Select(_1750_i)));
        _1750_i = (_1750_i) + (new BigInteger(1));
      }
    }
    public static void SetToUniqueSeqConstruct<X>(Dafny.Set<X> s, out Dafny.Sequence<X> xs)
    {
      xs = Dafny.Sequence<X>.Empty;
      X[] _1751_arr;
      var _nw8 = Dafny.ArrayHelpers.InitNewArray1<X>(Dafny.Helpers.Default<X>(), (new BigInteger((s).Count)));
      _1751_arr = _nw8;
      Dafny.Set<X> _1752_s1;
      _1752_s1 = s;
      { }
      { }
      { }
      while ((new BigInteger((_1752_s1).Count)) != (new BigInteger(0))) {
        { }
        { }
        X _1753_x;
        foreach (var _assign_such_that_0 in (_1752_s1).Elements) { _1753_x = _assign_such_that_0;
          if ((_1752_s1).Contains(_1753_x)) {
            goto after__ASSIGN_SUCH_THAT_0;
          }
        }
        throw new System.Exception("assign-such-that search produced no value (line 85)");
      after__ASSIGN_SUCH_THAT_0: ;
        { }
        { }
        var _index7 = (new BigInteger((s).Count)) - (new BigInteger((_1752_s1).Count));
        (_1751_arr)[(int)(_index7)] = _1753_x;
        _1752_s1 = (_1752_s1).Difference(Dafny.Set<X>.FromElements(_1753_x));
        { }
        { }
        { }
      }
      xs = Dafny.Helpers.SeqFromArray(_1751_arr);
      { }
    }
    public static void SubsequenceConstruct<X>(Dafny.Sequence<X> xs, Func<X,bool> f, out Dafny.Sequence<X> xs_k)
    {
      xs_k = Dafny.Sequence<X>.Empty;
      { }
      X[] _1754_arr;
      var _nw9 = Dafny.ArrayHelpers.InitNewArray1<X>(Dafny.Helpers.Default<X>(), (new BigInteger((xs).Count)));
      _1754_arr = _nw9;
      BigInteger _1755_i;
      _1755_i = new BigInteger(0);
      BigInteger _1756_j;
      _1756_j = new BigInteger(0);
      while ((_1755_i) < (new BigInteger((xs).Count))) {
        { }
        { }
        if (Dafny.Helpers.Id<Func<X,bool>>(f)((xs).Select(_1755_i))) {
          { }
          (_1754_arr)[(int)((_1756_j))] = (xs).Select(_1755_i);
          _1756_j = (_1756_j) + (new BigInteger(1));
          { }
        }
        _1755_i = (_1755_i) + (new BigInteger(1));
        { }
      }
      xs_k = Dafny.Helpers.SeqFromArray(_1754_arr).Take(_1756_j);
    }
    public static void UniqueSubsequenceConstruct<X>(Dafny.Sequence<X> xs, Func<X,bool> f, out Dafny.Sequence<X> xs_k)
    {
    TAIL_CALL_START: ;
xs_k = Dafny.Sequence<X>.Empty;
      Dafny.Set<X> _1757_s;
      _1757_s = ((System.Func<Dafny.Set<X>>)(() => {
        var _coll2 = new System.Collections.Generic.List<X>();
foreach (var _compr_1 in (xs).Elements) { X _1758_x = (X)_compr_1;
          if (((xs).Contains(_1758_x)) && (Dafny.Helpers.Id<Func<X,bool>>(f)(_1758_x))) {
            _coll2.Add(_1758_x);
          }
        }
        return Dafny.Set<X>.FromCollection(_coll2);
      }))();
      Dafny.Sequence<X> _out106;
_148_Common____SeqIsUnique__i_Compile.__default.SetToUniqueSeqConstruct<X>(_1757_s, out _out106);
xs_k = _out106;
    }
    public static Dafny.Sequence<X> AppendToUniqueSeq<X>(Dafny.Sequence<X> xs, X x)
    {
      Dafny.Sequence<X> _1759_xs_k = (xs).Concat(Dafny.Sequence<X>.FromElements(x));
return _1759_xs_k;
    }
    public static Dafny.Sequence<X> AppendToUniqueSeqMaybe<X>(Dafny.Sequence<X> xs, X x)
    {
      if ((xs).Contains(x)) {
        return xs;
      } else  {
        Dafny.Sequence<X> _1760_xs_k = (xs).Concat(Dafny.Sequence<X>.FromElements(x));
return _1760_xs_k;
      }
    }
  }
} // end of namespace _148_Common____SeqIsUnique__i_Compile
namespace _152_DistributedSystem__i_Compile {





  public class LS__State {
    public readonly _7_Environment__s_Compile.LEnvironment<_9_Native____Io__s_Compile.EndPoint,_33_Types__i_Compile.LockMessage,_33_Types__i_Compile.LockStep> environment;
public readonly Dafny.Map<_9_Native____Io__s_Compile.EndPoint,_36_Protocol__Node__i_Compile.Node> servers;
public LS__State(_7_Environment__s_Compile.LEnvironment<_9_Native____Io__s_Compile.EndPoint,_33_Types__i_Compile.LockMessage,_33_Types__i_Compile.LockStep> environment, Dafny.Map<_9_Native____Io__s_Compile.EndPoint,_36_Protocol__Node__i_Compile.Node> servers) {
      this.environment = environment;
this.servers = servers;
    }
    public override bool Equals(object other) {
      var oth = other as _152_DistributedSystem__i_Compile.LS__State;
return oth != null && Dafny.Helpers.AreEqual(this.environment, oth.environment) && Dafny.Helpers.AreEqual(this.servers, oth.servers);
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 0;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.environment));
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.servers));
return (int) hash;
    }
    public override string ToString() {
      string s = "_152_DistributedSystem__i_Compile.LS_State.LS_State";
s += "(";
s += Dafny.Helpers.ToString(this.environment);
s += ", ";
s += Dafny.Helpers.ToString(this.servers);
s += ")";
return s;
    }
    static LS__State theDefault;
public static LS__State Default {
      get {
        if (theDefault == null) {
          theDefault = new _152_DistributedSystem__i_Compile.LS__State(@_7_Environment__s_Compile.LEnvironment<_9_Native____Io__s_Compile.EndPoint,_33_Types__i_Compile.LockMessage,_33_Types__i_Compile.LockStep>.Default, Dafny.Map<_9_Native____Io__s_Compile.EndPoint,_36_Protocol__Node__i_Compile.Node>.Empty);
        }
        return theDefault;
      }
    }
    public static LS__State _DafnyDefaultValue() { return Default; }
public static LS__State create(_7_Environment__s_Compile.LEnvironment<_9_Native____Io__s_Compile.EndPoint,_33_Types__i_Compile.LockMessage,_33_Types__i_Compile.LockStep> environment, Dafny.Map<_9_Native____Io__s_Compile.EndPoint,_36_Protocol__Node__i_Compile.Node> servers) {
      return new LS__State(environment, servers);
    }
    public bool is_LS__State { get { return true; } }
public _7_Environment__s_Compile.LEnvironment<_9_Native____Io__s_Compile.EndPoint,_33_Types__i_Compile.LockMessage,_33_Types__i_Compile.LockStep> dtor_environment {
      get {
        return this.environment;
      }
    }
    public Dafny.Map<_9_Native____Io__s_Compile.EndPoint,_36_Protocol__Node__i_Compile.Node> dtor_servers {
      get {
        return this.servers;
      }
    }
  }

  public class GLS__State {
    public readonly _152_DistributedSystem__i_Compile.LS__State ls;
public readonly Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> history;
public GLS__State(_152_DistributedSystem__i_Compile.LS__State ls, Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> history) {
      this.ls = ls;
this.history = history;
    }
    public override bool Equals(object other) {
      var oth = other as _152_DistributedSystem__i_Compile.GLS__State;
return oth != null && Dafny.Helpers.AreEqual(this.ls, oth.ls) && Dafny.Helpers.AreEqual(this.history, oth.history);
    }
    public override int GetHashCode() {
      ulong hash = 5381;
hash = ((hash << 5) + hash) + 0;
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.ls));
hash = ((hash << 5) + hash) + ((ulong)Dafny.Helpers.GetHashCode(this.history));
return (int) hash;
    }
    public override string ToString() {
      string s = "_152_DistributedSystem__i_Compile.GLS_State.GLS_State";
s += "(";
s += Dafny.Helpers.ToString(this.ls);
s += ", ";
s += Dafny.Helpers.ToString(this.history);
s += ")";
return s;
    }
    static GLS__State theDefault;
public static GLS__State Default {
      get {
        if (theDefault == null) {
          theDefault = new _152_DistributedSystem__i_Compile.GLS__State(@_152_DistributedSystem__i_Compile.LS__State.Default, Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint>.Empty);
        }
        return theDefault;
      }
    }
    public static GLS__State _DafnyDefaultValue() { return Default; }
public static GLS__State create(_152_DistributedSystem__i_Compile.LS__State ls, Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> history) {
      return new GLS__State(ls, history);
    }
    public bool is_GLS__State { get { return true; } }
public _152_DistributedSystem__i_Compile.LS__State dtor_ls {
      get {
        return this.ls;
      }
    }
    public Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> dtor_history {
      get {
        return this.history;
      }
    }
  }

} // end of namespace _152_DistributedSystem__i_Compile
namespace _155_Refinement__i_Compile {



} // end of namespace _155_Refinement__i_Compile
namespace _161_RefinementProof__i_Compile {





} // end of namespace _161_RefinementProof__i_Compile
namespace _166_MarshallProof__i_Compile {




} // end of namespace _166_MarshallProof__i_Compile
namespace _168_Main__i_Compile {











  public partial class __default {

    public static void Main() {
      BigInteger v = 0;
      _default_Main(out v);
    }

    public static void _default_Main(out BigInteger exitCode)
    {
      exitCode = BigInteger.Zero;
      bool _1761_ok;
_131_Host__i_Compile.CScheduler _1762_host__state;
Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> _1763_config;
_9_Native____Io__s_Compile.EndPoint _1764_id;
bool _out107;
_131_Host__i_Compile.CScheduler _out108;
Dafny.Sequence<_9_Native____Io__s_Compile.EndPoint> _out109;
_9_Native____Io__s_Compile.EndPoint _out110;
_131_Host__i_Compile.__default.HostInitImpl(out _out107, out _out108, out _out109, out _out110);
_1761_ok = _out107;
_1762_host__state = _out108;
_1763_config = _out109;
_1764_id = _out110;
      { }
      while (_1761_ok) {
        { }
        { }
        { }
        bool _out111;
_131_Host__i_Compile.CScheduler _out112;
_131_Host__i_Compile.__default.HostNextImpl(_1762_host__state, out _out111, out _out112);
_1761_ok = _out111;
_1762_host__state = _out112;
        { }
      }
    }
  }

} // end of namespace _168_Main__i_Compile
namespace _module {


























































} // end of namespace _module
