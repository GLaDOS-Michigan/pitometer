// Package _439_LiveRSL____ReplicaImplNoReceiveNoClock__i_Compile
// Dafny module _439_LiveRSL____ReplicaImplNoReceiveNoClock__i_Compile compiled into Go

package _439_LiveRSL____ReplicaImplNoReceiveNoClock__i_Compile

import (
	_0_Native____NativeTypes__s_Compile "0_Native____NativeTypes__s_Compile_"
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
	_26_Collections____Seqs__s_Compile "26_Collections____Seqs__s_Compile_"
	_276_LiveRSL____ExecutorState__i_Compile "276_LiveRSL____ExecutorState__i_Compile_"
	_278_LiveRSL____LearnerState__i_Compile "278_LiveRSL____LearnerState__i_Compile_"
	_283_LiveRSL____CClockReading__i_Compile "283_LiveRSL____CClockReading__i_Compile_"
	_285_LiveRSL____ReplicaState__i_Compile "285_LiveRSL____ReplicaState__i_Compile_"
	_294_LiveRSL____MinCQuorumSize__i_Compile "294_LiveRSL____MinCQuorumSize__i_Compile_"
	_297_LiveRSL____ElectionModel__i_Compile "297_LiveRSL____ElectionModel__i_Compile_"
	_2_Collections____Maps2__s_Compile "2_Collections____Maps2__s_Compile_"
	_301_Impl____LiveRSL____Broadcast__i_Compile "301_Impl____LiveRSL____Broadcast__i_Compile_"
	_305_LiveRSL____ProposerLemmas__i_Compile "305_LiveRSL____ProposerLemmas__i_Compile_"
	_308_LiveRSL____ProposerModel__i_Compile "308_LiveRSL____ProposerModel__i_Compile_"
	_30_Collections____Seqs__i_Compile "30_Collections____Seqs__i_Compile_"
	_312_LiveRSL____AcceptorModel__i_Compile "312_LiveRSL____AcceptorModel__i_Compile_"
	_316_LiveRSL____LearnerModel__i_Compile "316_LiveRSL____LearnerModel__i_Compile_"
	_321_LiveRSL____ExecutorModel__i_Compile "321_LiveRSL____ExecutorModel__i_Compile_"
	_323_LiveRSL____ReplicaModel__Part1__i_Compile "323_LiveRSL____ReplicaModel__Part1__i_Compile_"
	_32_Math____mul__nonlinear__i_Compile "32_Math____mul__nonlinear__i_Compile_"
	_331_LiveRSL____ReplicaModel__Part2__i_Compile "331_LiveRSL____ReplicaModel__Part2__i_Compile_"
	_339_LiveRSL____ReplicaModel__Part3__i_Compile "339_LiveRSL____ReplicaModel__Part3__i_Compile_"
	_347_LiveRSL____ReplicaModel__Part4__i_Compile "347_LiveRSL____ReplicaModel__Part4__i_Compile_"
	_34_Math____mul__auto__proofs__i_Compile "34_Math____mul__auto__proofs__i_Compile_"
	_355_LiveRSL____ReplicaModel__Part5__i_Compile "355_LiveRSL____ReplicaModel__Part5__i_Compile_"
	_357_LiveRSL____ReplicaModel__i_Compile "357_LiveRSL____ReplicaModel__i_Compile_"
	_368_LiveRSL____UdpRSL__i_Compile "368_LiveRSL____UdpRSL__i_Compile_"
	_36_Math____mul__auto__i_Compile "36_Math____mul__auto__i_Compile_"
	_372_LiveRSL____QRelations__i_Compile "372_LiveRSL____QRelations__i_Compile_"
	_374_LiveRSL____ReplicaImplLemmas__i_Compile "374_LiveRSL____ReplicaImplLemmas__i_Compile_"
	_383_LiveRSL____ReplicaImplClass__i_Compile "383_LiveRSL____ReplicaImplClass__i_Compile_"
	_405_LiveRSL____ReplicaImplDelivery__i_Compile "405_LiveRSL____ReplicaImplDelivery__i_Compile_"
	_409_LiveRSL____ReplicaImplReadClock__i_Compile "409_LiveRSL____ReplicaImplReadClock__i_Compile_"
	_40_Math____mul__i_Compile "40_Math____mul__i_Compile_"
	_420_LiveRSL____ReplicaImplProcessPacketNoClock__i_Compile "420_LiveRSL____ReplicaImplProcessPacketNoClock__i_Compile_"
	_426_LiveRSL____Unsendable__i_Compile "426_LiveRSL____Unsendable__i_Compile_"
	_428_LiveRSL____ReplicaImplProcessPacketX__i_Compile "428_LiveRSL____ReplicaImplProcessPacketX__i_Compile_"
	_42_Math____div__nonlinear__i_Compile "42_Math____div__nonlinear__i_Compile_"
	_44_Math____mod__auto__proofs__i_Compile "44_Math____mod__auto__proofs__i_Compile_"
	_46_Math____mod__auto__i_Compile "46_Math____mod__auto__i_Compile_"
	_50_AppStateMachine__i_Compile "50_AppStateMachine__i_Compile_"
	_54_Concrete__NodeIdentity__i_Compile "54_Concrete__NodeIdentity__i_Compile_"
	_56_LiveRSL____Types__i_Compile "56_LiveRSL____Types__i_Compile_"
	_58_Collections____Sets__i_Compile "58_Collections____Sets__i_Compile_"
	_5_Temporal____Temporal__s_Compile "5_Temporal____Temporal__s_Compile_"
	_61_LiveRSL____Configuration__i_Compile "61_LiveRSL____Configuration__i_Compile_"
	_65_LiveRSL____Message__i_Compile "65_LiveRSL____Message__i_Compile_"
	_68_LiveRSL____Environment__i_Compile "68_LiveRSL____Environment__i_Compile_"
	_71_LiveRSL____ClockReading__i_Compile "71_LiveRSL____ClockReading__i_Compile_"
	_74_Common____UpperBound__s_Compile "74_Common____UpperBound__s_Compile_"
	_76_LiveRSL____Parameters__i_Compile "76_LiveRSL____Parameters__i_Compile_"
	_78_LiveRSL____Constants__i_Compile "78_LiveRSL____Constants__i_Compile_"
	_7_Environment__s_Compile "7_Environment__s_Compile_"
	_85_LiveRSL____Broadcast__i_Compile "85_LiveRSL____Broadcast__i_Compile_"
	_91_Collections____CountMatches__i_Compile "91_Collections____CountMatches__i_Compile_"
	_93_LiveRSL____Acceptor__i_Compile "93_LiveRSL____Acceptor__i_Compile_"
	_99_LiveRSL____Election__i_Compile "99_LiveRSL____Election__i_Compile_"
	_9_Native____Io__s_Compile "9_Native____Io__s_Compile_"
	_System "System_"
	"clock"
	_dafny "dafny"
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
var _ _420_LiveRSL____ReplicaImplProcessPacketNoClock__i_Compile.Dummy__
var _ _426_LiveRSL____Unsendable__i_Compile.Dummy__
var _ _428_LiveRSL____ReplicaImplProcessPacketX__i_Compile.Dummy__

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

var Companion_Default___ = CompanionStruct_Default___{}

func (_this *Default__) Equals(other *Default__) bool {
	return _this == other
}

func (_this *Default__) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*Default__)
	return ok && _this.Equals(other)
}

func (*Default__) String() string {
	return "_439_LiveRSL____ReplicaImplNoReceiveNoClock__i_Compile.Default__"
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
	return "_439_LiveRSL____ReplicaImplNoReceiveNoClock__i_Compile.Default__"
}

func (_this *CompanionStruct_Default___) ReplicaNoReceiveNoClockNextSpontaneousMaybeEnterNewViewAndSend1a(r *_383_LiveRSL____ReplicaImplClass__i_Compile.ReplicaImpl, LReplicaNextSpontaneousMaybeEnterNewViewAndSend1aLog *clock.Stopwatch, LReplicaNextSpontaneousMaybeEnterNewViewAndSend1aNoopLog *clock.Stopwatch) bool {
	LReplicaNextSpontaneousMaybeEnterNewViewAndSend1aLog.LogStartEvent("LReplicaNextSpontaneousMaybeEnterNewViewAndSend1a")
	LReplicaNextSpontaneousMaybeEnterNewViewAndSend1aNoopLog.LogStartEvent("LReplicaNextSpontaneousMaybeEnterNewViewAndSend1aNoop")
	var ok bool = false
	var _ = ok
	{
	}
	var _5265_sent__packets _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
	var _ = _5265_sent__packets
	var _out472 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
	var _ = _out472
	var _out473 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
	var _ = _out473
	var noop bool
	_out472, _out473, noop = _339_LiveRSL____ReplicaModel__Part3__i_Compile.Companion_Default___.Replica__Next__Spontaneous__MaybeEnterNewViewAndSend1a(r.Replica)
	(r).Replica = _out472
	_5265_sent__packets = _out473
	{
	}
	var _out474 bool
	var _ = _out474
	_out474 = _405_LiveRSL____ReplicaImplDelivery__i_Compile.Companion_Default___.DeliverOutboundPackets(r, _5265_sent__packets)
	ok = _out474
	if !(ok) {
		if noop {
			LReplicaNextSpontaneousMaybeEnterNewViewAndSend1aNoopLog.LogEndEvent("LReplicaNextSpontaneousMaybeEnterNewViewAndSend1aNoop")
			LReplicaNextSpontaneousMaybeEnterNewViewAndSend1aLog.PopStartEvent()
		} else {
			LReplicaNextSpontaneousMaybeEnterNewViewAndSend1aLog.LogEndEvent("LReplicaNextSpontaneousMaybeEnterNewViewAndSend1a")
			LReplicaNextSpontaneousMaybeEnterNewViewAndSend1aNoopLog.PopStartEvent()
		}
		return ok
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	if noop {
		LReplicaNextSpontaneousMaybeEnterNewViewAndSend1aNoopLog.LogEndEvent("LReplicaNextSpontaneousMaybeEnterNewViewAndSend1aNoop")
		LReplicaNextSpontaneousMaybeEnterNewViewAndSend1aLog.PopStartEvent()
	} else {
		LReplicaNextSpontaneousMaybeEnterNewViewAndSend1aLog.LogEndEvent("LReplicaNextSpontaneousMaybeEnterNewViewAndSend1a")
		LReplicaNextSpontaneousMaybeEnterNewViewAndSend1aNoopLog.PopStartEvent()
	}
	return ok
}

func (_this *CompanionStruct_Default___) ReplicaNoReceiveNoClockNextSpontaneousMaybeEnterPhase2(r *_383_LiveRSL____ReplicaImplClass__i_Compile.ReplicaImpl, LReplicaNextSpontaneousMaybeEnterPhase2Log *clock.Stopwatch, LReplicaNextSpontaneousMaybeEnterPhase2NoopLog *clock.Stopwatch) bool {
	LReplicaNextSpontaneousMaybeEnterPhase2Log.LogStartEvent("LReplicaNextSpontaneousMaybeEnterPhase2")
	LReplicaNextSpontaneousMaybeEnterPhase2NoopLog.LogStartEvent("LReplicaNextSpontaneousMaybeEnterPhase2Noop")
	var noop bool
	var ok bool = false
	var _ = ok
	{
	}
	var _5266_sent__packets _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
	var _ = _5266_sent__packets
	var _out475 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
	var _ = _out475
	var _out476 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
	var _ = _out476
	_out475, _out476, noop = _339_LiveRSL____ReplicaModel__Part3__i_Compile.Companion_Default___.Replica__Next__Spontaneous__MaybeEnterPhase2(r.Replica)
	(r).Replica = _out475
	_5266_sent__packets = _out476
	{
	}
	var _out477 bool
	var _ = _out477
	_out477 = _405_LiveRSL____ReplicaImplDelivery__i_Compile.Companion_Default___.DeliverOutboundPackets(r, _5266_sent__packets)
	ok = _out477
	if !(ok) {
		if noop {
			LReplicaNextSpontaneousMaybeEnterPhase2NoopLog.LogEndEvent("LReplicaNextSpontaneousMaybeEnterPhase2Noop")
			LReplicaNextSpontaneousMaybeEnterPhase2Log.PopStartEvent()
		} else {
			LReplicaNextSpontaneousMaybeEnterPhase2Log.LogEndEvent("LReplicaNextSpontaneousMaybeEnterPhase2Log")
			LReplicaNextSpontaneousMaybeEnterPhase2NoopLog.PopStartEvent()
		}
		return ok
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	if noop {
		LReplicaNextSpontaneousMaybeEnterPhase2NoopLog.LogEndEvent("LReplicaNextSpontaneousMaybeEnterPhase2Noop")
		LReplicaNextSpontaneousMaybeEnterPhase2Log.PopStartEvent()
	} else {
		LReplicaNextSpontaneousMaybeEnterPhase2Log.LogEndEvent("LReplicaNextSpontaneousMaybeEnterPhase2Log")
		LReplicaNextSpontaneousMaybeEnterPhase2NoopLog.PopStartEvent()
	}
	return ok
}

func (_this *CompanionStruct_Default___) ReplicaNoReceiveNoClockNextSpontaneousTruncateLogBasedOnCheckpoints(r *_383_LiveRSL____ReplicaImplClass__i_Compile.ReplicaImpl, LReplicaNextSpontaneousTruncateLogBasedOnCheckpointsLog *clock.Stopwatch) bool {
	LReplicaNextSpontaneousTruncateLogBasedOnCheckpointsLog.LogStartEvent("LReplicaNextSpontaneousTruncateLogBasedOnCheckpoints")
	var ok bool = false
	var _ = ok
	{
	}
	var _5267_sent__packets _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
	var _ = _5267_sent__packets
	var _out478 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
	var _ = _out478
	var _out479 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
	var _ = _out479
	_out478, _out479 = _355_LiveRSL____ReplicaModel__Part5__i_Compile.Companion_Default___.Replica__Next__Spontaneous__TruncateLogBasedOnCheckpoints(r.Replica)
	(r).Replica = _out478
	_5267_sent__packets = _out479
	{
	}
	var _out480 bool
	var _ = _out480
	_out480 = _405_LiveRSL____ReplicaImplDelivery__i_Compile.Companion_Default___.DeliverOutboundPackets(r, _5267_sent__packets)
	ok = _out480
	if !(ok) {
		LReplicaNextSpontaneousTruncateLogBasedOnCheckpointsLog.LogEndEvent("LReplicaNextSpontaneousTruncateLogBasedOnCheckpoints")
		return ok
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	LReplicaNextSpontaneousTruncateLogBasedOnCheckpointsLog.LogEndEvent("LReplicaNextSpontaneousTruncateLogBasedOnCheckpoints")
	return ok
}
func (_this *CompanionStruct_Default___) ReplicaNoReceiveNoClockNextSpontaneousMaybeMakeDecision(r *_383_LiveRSL____ReplicaImplClass__i_Compile.ReplicaImpl, LReplicaNextSpontaneousMaybeMakeDecisionLog *clock.Stopwatch) bool {
	LReplicaNextSpontaneousMaybeMakeDecisionLog.LogStartEvent("LReplicaNextSpontaneousMaybeMakeDecision")
	var ok bool = false
	var _ = ok
	{
	}
	var _5268_sent__packets _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
	var _ = _5268_sent__packets
	var _out481 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
	var _ = _out481
	var _out482 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
	var _ = _out482
	_out481, _out482 = _355_LiveRSL____ReplicaModel__Part5__i_Compile.Companion_Default___.Replica__Next__Spontaneous__MaybeMakeDecision(r.Replica)
	(r).Replica = _out481
	_5268_sent__packets = _out482
	{
	}
	var _out483 bool
	var _ = _out483
	_out483 = _405_LiveRSL____ReplicaImplDelivery__i_Compile.Companion_Default___.DeliverOutboundPackets(r, _5268_sent__packets)
	ok = _out483
	if !(ok) {
		LReplicaNextSpontaneousMaybeMakeDecisionLog.LogEndEvent("LReplicaNextSpontaneousMaybeMakeDecision")
		return ok
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	LReplicaNextSpontaneousMaybeMakeDecisionLog.LogEndEvent("LReplicaNextSpontaneousMaybeMakeDecision")
	return ok
}
func (_this *CompanionStruct_Default___) ReplicaNoReceiveNoClockNextSpontaneousMaybeExecute(r *_383_LiveRSL____ReplicaImplClass__i_Compile.ReplicaImpl, LReplicaNextSpontaneousMaybeExecuteLog *clock.Stopwatch) bool {
	LReplicaNextSpontaneousMaybeExecuteLog.LogStartEvent("LReplicaNextSpontaneousMaybeExecute")
	var ok bool = false
	var _ = ok
	{
	}
	var _5269_sent__packets _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
	var _ = _5269_sent__packets
	var _out484 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
	var _ = _out484
	var _out485 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
	var _ = _out485
	_out484, _out485 = _355_LiveRSL____ReplicaModel__Part5__i_Compile.Companion_Default___.Replica__Next__Spontaneous__MaybeExecute(r.Replica, r.Cur__req__set, r.Prev__req__set, r.Reply__cache__mutable)
	(r).Replica = _out484
	_5269_sent__packets = _out485
	{
	}
	var _out486 bool
	var _ = _out486
	_out486 = _405_LiveRSL____ReplicaImplDelivery__i_Compile.Companion_Default___.DeliverOutboundPackets(r, _5269_sent__packets)
	ok = _out486
	if !(ok) {
		LReplicaNextSpontaneousMaybeExecuteLog.LogEndEvent("LReplicaNextSpontaneousMaybeExecute")
		return ok
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	LReplicaNextSpontaneousMaybeExecuteLog.LogEndEvent("LReplicaNextSpontaneousMaybeExecute")
	return ok
}
func (_this *CompanionStruct_Default___) Replica__NoReceive__NoClock__Next(r *_383_LiveRSL____ReplicaImplClass__i_Compile.ReplicaImpl, logs map[string]*clock.Stopwatch) bool {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var ok bool = false
	var _ = ok
	if (r.NextActionIndex) == (uint64(1)) {
		var _out487 bool
		var _ = _out487
		_out487 = Companion_Default___.ReplicaNoReceiveNoClockNextSpontaneousMaybeEnterNewViewAndSend1a(r, logs["LReplicaNextSpontaneousMaybeEnterNewViewAndSend1a"], logs["LReplicaNextSpontaneousMaybeEnterNewViewAndSend1aNoop"])
		ok = _out487
	} else if (r.NextActionIndex) == (uint64(2)) {
		var _out488 bool
		var _ = _out488
		_out488 = Companion_Default___.ReplicaNoReceiveNoClockNextSpontaneousMaybeEnterPhase2(r, logs["LReplicaNextSpontaneousMaybeEnterPhase2"], logs["LReplicaNextSpontaneousMaybeEnterPhase2Noop"])
		ok = _out488
	} else if (r.NextActionIndex) == (uint64(4)) {
		var _out489 bool
		var _ = _out489
		_out489 = Companion_Default___.ReplicaNoReceiveNoClockNextSpontaneousTruncateLogBasedOnCheckpoints(r, logs["LReplicaNextSpontaneousTruncateLogBasedOnCheckpoints"])
		ok = _out489
	} else if (r.NextActionIndex) == (uint64(5)) {
		var _out490 bool
		var _ = _out490
		_out490 = Companion_Default___.ReplicaNoReceiveNoClockNextSpontaneousMaybeMakeDecision(r, logs["LReplicaNextSpontaneousMaybeMakeDecision"])
		ok = _out490
	} else if (r.NextActionIndex) == (uint64(6)) {
		var _out491 bool
		var _ = _out491
		_out491 = Companion_Default___.ReplicaNoReceiveNoClockNextSpontaneousMaybeExecute(r, logs["LReplicaNextSpontaneousMaybeExecute"])
		ok = _out491
	}
	return ok
}

// End of class Default__
