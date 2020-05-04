// Package _420_LiveRSL____ReplicaImplProcessPacketNoClock__i_Compile
// Dafny module _420_LiveRSL____ReplicaImplProcessPacketNoClock__i_Compile compiled into Go

package _420_LiveRSL____ReplicaImplProcessPacketNoClock__i_Compile

import (
  _dafny "dafny"
_System "System_"
_0_Native____NativeTypes__s_Compile "0_Native____NativeTypes__s_Compile_"
_2_Collections____Maps2__s_Compile "2_Collections____Maps2__s_Compile_"
_5_Temporal____Temporal__s_Compile "5_Temporal____Temporal__s_Compile_"
_7_Environment__s_Compile "7_Environment__s_Compile_"
_9_Native____Io__s_Compile "9_Native____Io__s_Compile_"
_26_Collections____Seqs__s_Compile "26_Collections____Seqs__s_Compile_"
_30_Collections____Seqs__i_Compile "30_Collections____Seqs__i_Compile_"
_32_Math____mul__nonlinear__i_Compile "32_Math____mul__nonlinear__i_Compile_"
_34_Math____mul__auto__proofs__i_Compile "34_Math____mul__auto__proofs__i_Compile_"
_36_Math____mul__auto__i_Compile "36_Math____mul__auto__i_Compile_"
_40_Math____mul__i_Compile "40_Math____mul__i_Compile_"
_42_Math____div__nonlinear__i_Compile "42_Math____div__nonlinear__i_Compile_"
_44_Math____mod__auto__proofs__i_Compile "44_Math____mod__auto__proofs__i_Compile_"
_46_Math____mod__auto__i_Compile "46_Math____mod__auto__i_Compile_"
_50_AppStateMachine__i_Compile "50_AppStateMachine__i_Compile_"
_54_Concrete__NodeIdentity__i_Compile "54_Concrete__NodeIdentity__i_Compile_"
_56_LiveRSL____Types__i_Compile "56_LiveRSL____Types__i_Compile_"
_58_Collections____Sets__i_Compile "58_Collections____Sets__i_Compile_"
_61_LiveRSL____Configuration__i_Compile "61_LiveRSL____Configuration__i_Compile_"
_65_LiveRSL____Message__i_Compile "65_LiveRSL____Message__i_Compile_"
_68_LiveRSL____Environment__i_Compile "68_LiveRSL____Environment__i_Compile_"
_71_LiveRSL____ClockReading__i_Compile "71_LiveRSL____ClockReading__i_Compile_"
_74_Common____UpperBound__s_Compile "74_Common____UpperBound__s_Compile_"
_76_LiveRSL____Parameters__i_Compile "76_LiveRSL____Parameters__i_Compile_"
_78_LiveRSL____Constants__i_Compile "78_LiveRSL____Constants__i_Compile_"
_85_LiveRSL____Broadcast__i_Compile "85_LiveRSL____Broadcast__i_Compile_"
_91_Collections____CountMatches__i_Compile "91_Collections____CountMatches__i_Compile_"
_93_LiveRSL____Acceptor__i_Compile "93_LiveRSL____Acceptor__i_Compile_"
_99_LiveRSL____Election__i_Compile "99_LiveRSL____Election__i_Compile_"
_101_LiveRSL____Proposer__i_Compile "101_LiveRSL____Proposer__i_Compile_"
_115_LiveRSL____StateMachine__i_Compile "115_LiveRSL____StateMachine__i_Compile_"
_118_Collections____Maps__i_Compile "118_Collections____Maps__i_Compile_"
_120_LiveRSL____Executor__i_Compile "120_LiveRSL____Executor__i_Compile_"
_123_LiveRSL____Learner__i_Compile "123_LiveRSL____Learner__i_Compile_"
_126_LiveRSL____Replica__i_Compile "126_LiveRSL____Replica__i_Compile_"
_135_Logic____Option__i_Compile "135_Logic____Option__i_Compile_"
_138_Native____NativeTypes__i_Compile "138_Native____NativeTypes__i_Compile_"
_141_Libraries____base__s_Compile "141_Libraries____base__s_Compile_"
_143_Math____power2__s_Compile "143_Math____power2__s_Compile_"
_145_Math____power__s_Compile "145_Math____power__s_Compile_"
_149_Math____power__i_Compile "149_Math____power__i_Compile_"
_153_Math____div__def__i_Compile "153_Math____div__def__i_Compile_"
_157_Math____div__boogie__i_Compile "157_Math____div__boogie__i_Compile_"
_162_Math____div__auto__proofs__i_Compile "162_Math____div__auto__proofs__i_Compile_"
_164_Math____div__auto__i_Compile "164_Math____div__auto__i_Compile_"
_166_Math____div__i_Compile "166_Math____div__i_Compile_"
_168_Math____power2__i_Compile "168_Math____power2__i_Compile_"
_170_Common____Util__i_Compile "170_Common____Util__i_Compile_"
_174_Common____MarshallInt__i_Compile "174_Common____MarshallInt__i_Compile_"
_176_Common____GenericMarshalling__i_Compile "176_Common____GenericMarshalling__i_Compile_"
_180_Common____UdpClient__i_Compile "180_Common____UdpClient__i_Compile_"
_182_Common____SeqIsUniqueDef__i_Compile "182_Common____SeqIsUniqueDef__i_Compile_"
_185_Common____SeqIsUnique__i_Compile "185_Common____SeqIsUnique__i_Compile_"
_191_GenericRefinement__i_Compile "191_GenericRefinement__i_Compile_"
_194_Common____NodeIdentity__i_Compile "194_Common____NodeIdentity__i_Compile_"
_197_LiveRSL____AppInterface__i_Compile "197_LiveRSL____AppInterface__i_Compile_"
_214_LiveRSL____CTypes__i_Compile "214_LiveRSL____CTypes__i_Compile_"
_217_LiveRSL____CMessage__i_Compile "217_LiveRSL____CMessage__i_Compile_"
_225_LiveRSL____CMessageRefinements__i_Compile "225_LiveRSL____CMessageRefinements__i_Compile_"
_228_LiveRSL____PacketParsing__i_Compile "228_LiveRSL____PacketParsing__i_Compile_"
_234_Common____UpperBound__i_Compile "234_Common____UpperBound__i_Compile_"
_236_LiveRSL____ParametersState__i_Compile "236_LiveRSL____ParametersState__i_Compile_"
_238_LiveRSL____CPaxosConfiguration__i_Compile "238_LiveRSL____CPaxosConfiguration__i_Compile_"
_240_LiveRSL____ConstantsState__i_Compile "240_LiveRSL____ConstantsState__i_Compile_"
_243_LiveRSL____PaxosWorldState__i_Compile "243_LiveRSL____PaxosWorldState__i_Compile_"
_245_LiveRSL____ReplicaConstantsState__i_Compile "245_LiveRSL____ReplicaConstantsState__i_Compile_"
_251_LiveRSL____ElectionState__i_Compile "251_LiveRSL____ElectionState__i_Compile_"
_254_LiveRSL____ProposerState__i_Compile "254_LiveRSL____ProposerState__i_Compile_"
_265_LiveRSL____COperationNumberSort__i_Compile "265_LiveRSL____COperationNumberSort__i_Compile_"
_267_LiveRSL____CLastCheckpointedMap__i_Compile "267_LiveRSL____CLastCheckpointedMap__i_Compile_"
_269_LiveRSL____AcceptorState__i_Compile "269_LiveRSL____AcceptorState__i_Compile_"
_276_LiveRSL____ExecutorState__i_Compile "276_LiveRSL____ExecutorState__i_Compile_"
_278_LiveRSL____LearnerState__i_Compile "278_LiveRSL____LearnerState__i_Compile_"
_283_LiveRSL____CClockReading__i_Compile "283_LiveRSL____CClockReading__i_Compile_"
_285_LiveRSL____ReplicaState__i_Compile "285_LiveRSL____ReplicaState__i_Compile_"
_294_LiveRSL____MinCQuorumSize__i_Compile "294_LiveRSL____MinCQuorumSize__i_Compile_"
_297_LiveRSL____ElectionModel__i_Compile "297_LiveRSL____ElectionModel__i_Compile_"
_301_Impl____LiveRSL____Broadcast__i_Compile "301_Impl____LiveRSL____Broadcast__i_Compile_"
_305_LiveRSL____ProposerLemmas__i_Compile "305_LiveRSL____ProposerLemmas__i_Compile_"
_308_LiveRSL____ProposerModel__i_Compile "308_LiveRSL____ProposerModel__i_Compile_"
_312_LiveRSL____AcceptorModel__i_Compile "312_LiveRSL____AcceptorModel__i_Compile_"
_316_LiveRSL____LearnerModel__i_Compile "316_LiveRSL____LearnerModel__i_Compile_"
_321_LiveRSL____ExecutorModel__i_Compile "321_LiveRSL____ExecutorModel__i_Compile_"
_323_LiveRSL____ReplicaModel__Part1__i_Compile "323_LiveRSL____ReplicaModel__Part1__i_Compile_"
_331_LiveRSL____ReplicaModel__Part2__i_Compile "331_LiveRSL____ReplicaModel__Part2__i_Compile_"
_339_LiveRSL____ReplicaModel__Part3__i_Compile "339_LiveRSL____ReplicaModel__Part3__i_Compile_"
_347_LiveRSL____ReplicaModel__Part4__i_Compile "347_LiveRSL____ReplicaModel__Part4__i_Compile_"
_355_LiveRSL____ReplicaModel__Part5__i_Compile "355_LiveRSL____ReplicaModel__Part5__i_Compile_"
_357_LiveRSL____ReplicaModel__i_Compile "357_LiveRSL____ReplicaModel__i_Compile_"
_368_LiveRSL____UdpRSL__i_Compile "368_LiveRSL____UdpRSL__i_Compile_"
_372_LiveRSL____QRelations__i_Compile "372_LiveRSL____QRelations__i_Compile_"
_374_LiveRSL____ReplicaImplLemmas__i_Compile "374_LiveRSL____ReplicaImplLemmas__i_Compile_"
_383_LiveRSL____ReplicaImplClass__i_Compile "383_LiveRSL____ReplicaImplClass__i_Compile_"
_405_LiveRSL____ReplicaImplDelivery__i_Compile "405_LiveRSL____ReplicaImplDelivery__i_Compile_"
_409_LiveRSL____ReplicaImplReadClock__i_Compile "409_LiveRSL____ReplicaImplReadClock__i_Compile_"
)
var _ _dafny.Dummy__
var _ _System.Dummy__
var _ _0_Native____NativeTypes__s_Compile.Dummy__
var _ _2_Collections____Maps2__s_Compile.Dummy__
var _ _5_Temporal____Temporal__s_Compile.Dummy__
var _ _7_Environment__s_Compile.Dummy__
var _ _9_Native____Io__s_Compile.Dummy__
var _ _26_Collections____Seqs__s_Compile.Dummy__
var _ _30_Collections____Seqs__i_Compile.Dummy__
var _ _32_Math____mul__nonlinear__i_Compile.Dummy__
var _ _34_Math____mul__auto__proofs__i_Compile.Dummy__
var _ _36_Math____mul__auto__i_Compile.Dummy__
var _ _40_Math____mul__i_Compile.Dummy__
var _ _42_Math____div__nonlinear__i_Compile.Dummy__
var _ _44_Math____mod__auto__proofs__i_Compile.Dummy__
var _ _46_Math____mod__auto__i_Compile.Dummy__
var _ _50_AppStateMachine__i_Compile.Dummy__
var _ _54_Concrete__NodeIdentity__i_Compile.Dummy__
var _ _56_LiveRSL____Types__i_Compile.Dummy__
var _ _58_Collections____Sets__i_Compile.Dummy__
var _ _61_LiveRSL____Configuration__i_Compile.Dummy__
var _ _65_LiveRSL____Message__i_Compile.Dummy__
var _ _68_LiveRSL____Environment__i_Compile.Dummy__
var _ _71_LiveRSL____ClockReading__i_Compile.Dummy__
var _ _74_Common____UpperBound__s_Compile.Dummy__
var _ _76_LiveRSL____Parameters__i_Compile.Dummy__
var _ _78_LiveRSL____Constants__i_Compile.Dummy__
var _ _85_LiveRSL____Broadcast__i_Compile.Dummy__
var _ _91_Collections____CountMatches__i_Compile.Dummy__
var _ _93_LiveRSL____Acceptor__i_Compile.Dummy__
var _ _99_LiveRSL____Election__i_Compile.Dummy__
var _ _101_LiveRSL____Proposer__i_Compile.Dummy__
var _ _115_LiveRSL____StateMachine__i_Compile.Dummy__
var _ _118_Collections____Maps__i_Compile.Dummy__
var _ _120_LiveRSL____Executor__i_Compile.Dummy__
var _ _123_LiveRSL____Learner__i_Compile.Dummy__
var _ _126_LiveRSL____Replica__i_Compile.Dummy__
var _ _135_Logic____Option__i_Compile.Dummy__
var _ _138_Native____NativeTypes__i_Compile.Dummy__
var _ _141_Libraries____base__s_Compile.Dummy__
var _ _143_Math____power2__s_Compile.Dummy__
var _ _145_Math____power__s_Compile.Dummy__
var _ _149_Math____power__i_Compile.Dummy__
var _ _153_Math____div__def__i_Compile.Dummy__
var _ _157_Math____div__boogie__i_Compile.Dummy__
var _ _162_Math____div__auto__proofs__i_Compile.Dummy__
var _ _164_Math____div__auto__i_Compile.Dummy__
var _ _166_Math____div__i_Compile.Dummy__
var _ _168_Math____power2__i_Compile.Dummy__
var _ _170_Common____Util__i_Compile.Dummy__
var _ _174_Common____MarshallInt__i_Compile.Dummy__
var _ _176_Common____GenericMarshalling__i_Compile.Dummy__
var _ _180_Common____UdpClient__i_Compile.Dummy__
var _ _182_Common____SeqIsUniqueDef__i_Compile.Dummy__
var _ _185_Common____SeqIsUnique__i_Compile.Dummy__
var _ _191_GenericRefinement__i_Compile.Dummy__
var _ _194_Common____NodeIdentity__i_Compile.Dummy__
var _ _197_LiveRSL____AppInterface__i_Compile.Dummy__
var _ _214_LiveRSL____CTypes__i_Compile.Dummy__
var _ _217_LiveRSL____CMessage__i_Compile.Dummy__
var _ _225_LiveRSL____CMessageRefinements__i_Compile.Dummy__
var _ _228_LiveRSL____PacketParsing__i_Compile.Dummy__
var _ _234_Common____UpperBound__i_Compile.Dummy__
var _ _236_LiveRSL____ParametersState__i_Compile.Dummy__
var _ _238_LiveRSL____CPaxosConfiguration__i_Compile.Dummy__
var _ _240_LiveRSL____ConstantsState__i_Compile.Dummy__
var _ _243_LiveRSL____PaxosWorldState__i_Compile.Dummy__
var _ _245_LiveRSL____ReplicaConstantsState__i_Compile.Dummy__
var _ _251_LiveRSL____ElectionState__i_Compile.Dummy__
var _ _254_LiveRSL____ProposerState__i_Compile.Dummy__
var _ _265_LiveRSL____COperationNumberSort__i_Compile.Dummy__
var _ _267_LiveRSL____CLastCheckpointedMap__i_Compile.Dummy__
var _ _269_LiveRSL____AcceptorState__i_Compile.Dummy__
var _ _276_LiveRSL____ExecutorState__i_Compile.Dummy__
var _ _278_LiveRSL____LearnerState__i_Compile.Dummy__
var _ _283_LiveRSL____CClockReading__i_Compile.Dummy__
var _ _285_LiveRSL____ReplicaState__i_Compile.Dummy__
var _ _294_LiveRSL____MinCQuorumSize__i_Compile.Dummy__
var _ _297_LiveRSL____ElectionModel__i_Compile.Dummy__
var _ _301_Impl____LiveRSL____Broadcast__i_Compile.Dummy__
var _ _305_LiveRSL____ProposerLemmas__i_Compile.Dummy__
var _ _308_LiveRSL____ProposerModel__i_Compile.Dummy__
var _ _312_LiveRSL____AcceptorModel__i_Compile.Dummy__
var _ _316_LiveRSL____LearnerModel__i_Compile.Dummy__
var _ _321_LiveRSL____ExecutorModel__i_Compile.Dummy__
var _ _323_LiveRSL____ReplicaModel__Part1__i_Compile.Dummy__
var _ _331_LiveRSL____ReplicaModel__Part2__i_Compile.Dummy__
var _ _339_LiveRSL____ReplicaModel__Part3__i_Compile.Dummy__
var _ _347_LiveRSL____ReplicaModel__Part4__i_Compile.Dummy__
var _ _355_LiveRSL____ReplicaModel__Part5__i_Compile.Dummy__
var _ _357_LiveRSL____ReplicaModel__i_Compile.Dummy__
var _ _368_LiveRSL____UdpRSL__i_Compile.Dummy__
var _ _372_LiveRSL____QRelations__i_Compile.Dummy__
var _ _374_LiveRSL____ReplicaImplLemmas__i_Compile.Dummy__
var _ _383_LiveRSL____ReplicaImplClass__i_Compile.Dummy__
var _ _405_LiveRSL____ReplicaImplDelivery__i_Compile.Dummy__
var _ _409_LiveRSL____ReplicaImplReadClock__i_Compile.Dummy__

type Dummy__ struct{}











// Definition of class Default__
type Default__ struct {
  dummy byte
}

func New_Default___() *Default__ {
  _this := Default__{}

  return &_this
}

type CompanionStruct_Default___ struct {
}
var Companion_Default___ = CompanionStruct_Default___ {
}

func (_this *Default__) Equals(other *Default__) bool {
  return _this == other
}

func (_this *Default__) EqualsGeneric(x interface{}) bool {
  other, ok := x.(*Default__)
return ok && _this.Equals(other)
}

func (*Default__) String() string {
  return "_420_LiveRSL____ReplicaImplProcessPacketNoClock__i_Compile.Default__"
}

func Type_Default___() _dafny.Type {
  return type_Default___{}
}

type type_Default___ struct {
}

func (_this type_Default___) Default() interface{} {
  return (*Default__)(nil)
}

func (_this type_Default___) String() string {
  return "_420_LiveRSL____ReplicaImplProcessPacketNoClock__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) ReplicaNextProcessPacketInvalid(r *_383_LiveRSL____ReplicaImplClass__i_Compile.ReplicaImpl, cpacket _217_LiveRSL____CMessage__i_Compile.CPacket) {
  { }
  { }
  { }
  var _5251_sent__packets _217_LiveRSL____CMessage__i_Compile.OutboundPackets
  var _ = _5251_sent__packets
  _5251_sent__packets = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_OutboundPacket{_135_Logic____Option__i_Compile.Option{_135_Logic____Option__i_Compile.Option_None{}}}}
  { }
  { }
  { }
  { }
  { }
  { }
  return
}
func (_this *CompanionStruct_Default___) ReplicaNextProcessPacketRequest(r *_383_LiveRSL____ReplicaImplClass__i_Compile.ReplicaImpl, cpacket _217_LiveRSL____CMessage__i_Compile.CPacket) bool {
  var ok bool = false
  var _ = ok
  { }
  { }
  { }
  { }
  { }
  { }
  var _5252_sent__packets _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
  var _ = _5252_sent__packets
  var _out432 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
  var _ = _out432
var _out433 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
  var _ = _out433
_out432,_out433 = _323_LiveRSL____ReplicaModel__Part1__i_Compile.Companion_Default___.Replica__Next__Process__Request(r.Replica, cpacket, r.Cur__req__set, r.Prev__req__set, r.Reply__cache__mutable)
(r).Replica = _out432
_5252_sent__packets = _out433
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  var _out434 bool
  var _ = _out434
_out434 = _405_LiveRSL____ReplicaImplDelivery__i_Compile.Companion_Default___.DeliverOutboundPackets(r, _5252_sent__packets)
ok = _out434
  if (!(ok)) {
    return ok
  }
  { }
  { }
  return ok
}
func (_this *CompanionStruct_Default___) ReplicaNextProcessPacket1a(r *_383_LiveRSL____ReplicaImplClass__i_Compile.ReplicaImpl, cpacket _217_LiveRSL____CMessage__i_Compile.CPacket) bool {
  var ok bool = false
  var _ = ok
  { }
  { }
  { }
  { }
  { }
  { }
  var _5253_sent__packets _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
  var _ = _5253_sent__packets
  var _out435 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
  var _ = _out435
var _out436 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
  var _ = _out436
_out435,_out436 = _323_LiveRSL____ReplicaModel__Part1__i_Compile.Companion_Default___.Replica__Next__Process__1a(r.Replica, cpacket)
(r).Replica = _out435
_5253_sent__packets = _out436
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  var _out437 bool
  var _ = _out437
_out437 = _405_LiveRSL____ReplicaImplDelivery__i_Compile.Companion_Default___.DeliverOutboundPackets(r, _5253_sent__packets)
ok = _out437
  if (!(ok)) {
    return ok
  }
  { }
  { }
  return ok
}
func (_this *CompanionStruct_Default___) ReplicaNextProcessPacket1b(r *_383_LiveRSL____ReplicaImplClass__i_Compile.ReplicaImpl, cpacket _217_LiveRSL____CMessage__i_Compile.CPacket) bool {
  var ok bool = false
  var _ = ok
  { }
  { }
  { }
  { }
  { }
  { }
  var _5254_sent__packets _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
  var _ = _5254_sent__packets
  var _out438 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
  var _ = _out438
var _out439 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
  var _ = _out439
_out438,_out439 = _331_LiveRSL____ReplicaModel__Part2__i_Compile.Companion_Default___.Replica__Next__Process__1b(r.Replica, cpacket)
(r).Replica = _out438
_5254_sent__packets = _out439
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  var _out440 bool
  var _ = _out440
_out440 = _405_LiveRSL____ReplicaImplDelivery__i_Compile.Companion_Default___.DeliverOutboundPackets(r, _5254_sent__packets)
ok = _out440
  if (!(ok)) {
    return ok
  }
  { }
  { }
  return ok
}
func (_this *CompanionStruct_Default___) ReplicaNextProcessPacketStartingPhase2(r *_383_LiveRSL____ReplicaImplClass__i_Compile.ReplicaImpl, cpacket _217_LiveRSL____CMessage__i_Compile.CPacket) bool {
  var ok bool = false
  var _ = ok
  { }
  { }
  { }
  { }
  { }
  { }
  var _5255_sent__packets _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
  var _ = _5255_sent__packets
  var _out441 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
  var _ = _out441
var _out442 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
  var _ = _out442
_out441,_out442 = _331_LiveRSL____ReplicaModel__Part2__i_Compile.Companion_Default___.Replica__Next__Process__StartingPhase2(r.Replica, cpacket)
(r).Replica = _out441
_5255_sent__packets = _out442
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  var _out443 bool
  var _ = _out443
_out443 = _405_LiveRSL____ReplicaImplDelivery__i_Compile.Companion_Default___.DeliverOutboundPackets(r, _5255_sent__packets)
ok = _out443
  if (!(ok)) {
    return ok
  }
  { }
  { }
  return ok
}
func (_this *CompanionStruct_Default___) ReplicaNextProcessPacket2a(r *_383_LiveRSL____ReplicaImplClass__i_Compile.ReplicaImpl, cpacket _217_LiveRSL____CMessage__i_Compile.CPacket) bool {
  var ok bool = false
  var _ = ok
  { }
  { }
  { }
  { }
  { }
  { }
  var _5256_sent__packets _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
  var _ = _5256_sent__packets
  var _out444 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
  var _ = _out444
var _out445 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
  var _ = _out445
_out444,_out445 = _331_LiveRSL____ReplicaModel__Part2__i_Compile.Companion_Default___.Replica__Next__Process__2a(r.Replica, cpacket)
(r).Replica = _out444
_5256_sent__packets = _out445
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  var _out446 bool
  var _ = _out446
_out446 = _405_LiveRSL____ReplicaImplDelivery__i_Compile.Companion_Default___.DeliverOutboundPackets(r, _5256_sent__packets)
ok = _out446
  if (!(ok)) {
    return ok
  }
  { }
  { }
  return ok
}
func (_this *CompanionStruct_Default___) ReplicaNextProcessPacket2b(r *_383_LiveRSL____ReplicaImplClass__i_Compile.ReplicaImpl, cpacket _217_LiveRSL____CMessage__i_Compile.CPacket) bool {
  var ok bool = false
  var _ = ok
  { }
  { }
  { }
  { }
  { }
  { }
  var _5257_sent__packets _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
  var _ = _5257_sent__packets
  var _out447 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
  var _ = _out447
var _out448 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
  var _ = _out448
_out447,_out448 = _339_LiveRSL____ReplicaModel__Part3__i_Compile.Companion_Default___.Replica__Next__Process__2b(r.Replica, cpacket)
(r).Replica = _out447
_5257_sent__packets = _out448
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  var _out449 bool
  var _ = _out449
_out449 = _405_LiveRSL____ReplicaImplDelivery__i_Compile.Companion_Default___.DeliverOutboundPackets(r, _5257_sent__packets)
ok = _out449
  if (!(ok)) {
    return ok
  }
  { }
  { }
  return ok
}
func (_this *CompanionStruct_Default___) ReplicaNextProcessPacketReply(r *_383_LiveRSL____ReplicaImplClass__i_Compile.ReplicaImpl, cpacket _217_LiveRSL____CMessage__i_Compile.CPacket) {
  { }
  { }
  { }
  var _5258_sent__packets _217_LiveRSL____CMessage__i_Compile.OutboundPackets
  var _ = _5258_sent__packets
  _5258_sent__packets = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_Broadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast_CBroadcastNop{}}}}
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  return
}
func (_this *CompanionStruct_Default___) ReplicaNextProcessPacketAppStateRequest(r *_383_LiveRSL____ReplicaImplClass__i_Compile.ReplicaImpl, cpacket _217_LiveRSL____CMessage__i_Compile.CPacket) bool {
  var ok bool = false
  var _ = ok
  { }
  { }
  { }
  { }
  { }
  { }
  var _5259_sent__packets _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
  var _ = _5259_sent__packets
  var _out450 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
  var _ = _out450
var _out451 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
  var _ = _out451
_out450,_out451 = _347_LiveRSL____ReplicaModel__Part4__i_Compile.Companion_Default___.Replica__Next__Process__AppStateRequest(r.Replica, cpacket, r.Reply__cache__mutable)
(r).Replica = _out450
_5259_sent__packets = _out451
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  var _out452 bool
  var _ = _out452
_out452 = _405_LiveRSL____ReplicaImplDelivery__i_Compile.Companion_Default___.DeliverOutboundPackets(r, _5259_sent__packets)
ok = _out452
  if (!(ok)) {
    return ok
  }
  { }
  { }
  return ok
}
func (_this *CompanionStruct_Default___) ReplicaNextProcessPacketAppStateSupply(r *_383_LiveRSL____ReplicaImplClass__i_Compile.ReplicaImpl, cpacket _217_LiveRSL____CMessage__i_Compile.CPacket) bool {
  var ok bool = false
  var _ = ok
  { }
  { }
  { }
  { }
  { }
  { }
  var _5260_sent__packets _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
  var _ = _5260_sent__packets
var _5261_replicaChanged bool = false
  var _ = _5261_replicaChanged
var _5262_newCache *_9_Native____Io__s_Compile.MutableMap = (*_9_Native____Io__s_Compile.MutableMap)(nil)
  var _ = _5262_newCache
  var _out453 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
  var _ = _out453
var _out454 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
  var _ = _out454
var _out455 bool
  var _ = _out455
var _out456 *_9_Native____Io__s_Compile.MutableMap
  var _ = _out456
_out453,_out454,_out455,_out456 = _355_LiveRSL____ReplicaModel__Part5__i_Compile.Companion_Default___.Replica__Next__Process__AppStateSupply(r.Replica, cpacket)
(r).Replica = _out453
_5260_sent__packets = _out454
_5261_replicaChanged = _out455
_5262_newCache = _out456
  if (_5261_replicaChanged) {
    (r).Reply__cache__mutable = _5262_newCache
  }
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  var _out457 bool
  var _ = _out457
_out457 = _405_LiveRSL____ReplicaImplDelivery__i_Compile.Companion_Default___.DeliverOutboundPackets(r, _5260_sent__packets)
ok = _out457
  if (!(ok)) {
    return ok
  }
  { }
  { }
  return ok
}
func (_this *CompanionStruct_Default___) Replica__Next__ProcessPacketWithoutReadingClock__body(r *_383_LiveRSL____ReplicaImplClass__i_Compile.ReplicaImpl, cpacket _217_LiveRSL____CMessage__i_Compile.CPacket) bool {
  var ok bool = false
  var _ = ok
  if (((cpacket).Dtor_msg()).Is_CMessage__Invalid()) {
    ok = true
    Companion_Default___.ReplicaNextProcessPacketInvalid(r, cpacket)
  } else if (((cpacket).Dtor_msg()).Is_CMessage__Request()) {
    var _out458 bool
    var _ = _out458
_out458 = Companion_Default___.ReplicaNextProcessPacketRequest(r, cpacket)
ok = _out458
  } else if (((cpacket).Dtor_msg()).Is_CMessage__1a()) {
    var _out459 bool
    var _ = _out459
_out459 = Companion_Default___.ReplicaNextProcessPacket1a(r, cpacket)
ok = _out459
  } else if (((cpacket).Dtor_msg()).Is_CMessage__1b()) {
    var _out460 bool
    var _ = _out460
_out460 = Companion_Default___.ReplicaNextProcessPacket1b(r, cpacket)
ok = _out460
  } else if (((cpacket).Dtor_msg()).Is_CMessage__StartingPhase2()) {
    var _out461 bool
    var _ = _out461
_out461 = Companion_Default___.ReplicaNextProcessPacketStartingPhase2(r, cpacket)
ok = _out461
  } else if (((cpacket).Dtor_msg()).Is_CMessage__2a()) {
    var _out462 bool
    var _ = _out462
_out462 = Companion_Default___.ReplicaNextProcessPacket2a(r, cpacket)
ok = _out462
  } else if (((cpacket).Dtor_msg()).Is_CMessage__2b()) {
    var _out463 bool
    var _ = _out463
_out463 = Companion_Default___.ReplicaNextProcessPacket2b(r, cpacket)
ok = _out463
  } else if (((cpacket).Dtor_msg()).Is_CMessage__Reply()) {
    ok = true
    Companion_Default___.ReplicaNextProcessPacketReply(r, cpacket)
  } else if (((cpacket).Dtor_msg()).Is_CMessage__AppStateRequest()) {
    var _out464 bool
    var _ = _out464
_out464 = Companion_Default___.ReplicaNextProcessPacketAppStateRequest(r, cpacket)
ok = _out464
  } else if (((cpacket).Dtor_msg()).Is_CMessage__AppStateSupply()) {
    var _out465 bool
    var _ = _out465
_out465 = Companion_Default___.ReplicaNextProcessPacketAppStateSupply(r, cpacket)
ok = _out465
  } else { }
  return ok
}
// End of class Default__
