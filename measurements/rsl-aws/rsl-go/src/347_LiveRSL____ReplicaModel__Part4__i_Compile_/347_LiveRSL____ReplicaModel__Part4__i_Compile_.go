// Package _347_LiveRSL____ReplicaModel__Part4__i_Compile
// Dafny module _347_LiveRSL____ReplicaModel__Part4__i_Compile compiled into Go

package _347_LiveRSL____ReplicaModel__Part4__i_Compile

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
	return "_347_LiveRSL____ReplicaModel__Part4__i_Compile.Default__"
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
	return "_347_LiveRSL____ReplicaModel__Part4__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) Replica__Next__Process__AppStateRequest(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, inp _217_LiveRSL____CMessage__i_Compile.CPacket, reply__cache__mutable *_9_Native____Io__s_Compile.MutableMap) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
	var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
	var _ = replica_k
	var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
	var _ = packets__sent
	var _5131_start__time uint64
	var _ = _5131_start__time
	var _out315 uint64
	var _ = _out315
	_out315 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_5131_start__time = _out315
	var _5132_newExecutor _276_LiveRSL____ExecutorState__i_Compile.ExecutorState
	var _ = _5132_newExecutor
	var _5133_packets _217_LiveRSL____CMessage__i_Compile.OutboundPackets
	var _ = _5133_packets
	var _out316 _276_LiveRSL____ExecutorState__i_Compile.ExecutorState
	var _ = _out316
	var _out317 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
	var _ = _out317
	_out316, _out317 = _321_LiveRSL____ExecutorModel__i_Compile.Companion_Default___.ExecutorProcessAppStateRequest((replica).Dtor_executor(), inp, reply__cache__mutable)
	_5132_newExecutor = _out316
	_5133_packets = _out317
	replica_k = func(_pat_let134_0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
		return func(_5134_dt__update__tmp_h0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
			return func(_pat_let135_0 _276_LiveRSL____ExecutorState__i_Compile.ExecutorState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
				return func(_5135_dt__update_hexecutor_h0 _276_LiveRSL____ExecutorState__i_Compile.ExecutorState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
					return _285_LiveRSL____ReplicaState__i_Compile.ReplicaState{_285_LiveRSL____ReplicaState__i_Compile.ReplicaState_ReplicaState{(_5134_dt__update__tmp_h0).Dtor_constants(), (_5134_dt__update__tmp_h0).Dtor_nextHeartbeatTime(), (_5134_dt__update__tmp_h0).Dtor_proposer(), (_5134_dt__update__tmp_h0).Dtor_acceptor(), (_5134_dt__update__tmp_h0).Dtor_learner(), _5135_dt__update_hexecutor_h0}}
				}(_pat_let135_0)
			}(_5132_newExecutor)
		}(_pat_let134_0)
	}(replica)
	packets__sent = _5133_packets
	var _5136_end__time uint64
	var _ = _5136_end__time
	var _out318 uint64
	var _ = _out318
	_out318 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_5136_end__time = _out318
	_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("Replica_Next_Process_AppStateRequest"), _5131_start__time, _5136_end__time)
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
	return replica_k, packets__sent
}
func (_this *CompanionStruct_Default___) Replica__Next__Process__Heartbeat(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, inp _217_LiveRSL____CMessage__i_Compile.CPacket, clock uint64, cur__req__set *_9_Native____Io__s_Compile.MutableSet, prev__req__set *_9_Native____Io__s_Compile.MutableSet) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
	var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
	var _ = replica_k
	var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
	var _ = packets__sent
	var _5137_start__time uint64
	var _ = _5137_start__time
	var _out319 uint64
	var _ = _out319
	_out319 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_5137_start__time = _out319
	packets__sent = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_Broadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast_CBroadcastNop{}}}}
	{
	}
	var _5138_newProposer _254_LiveRSL____ProposerState__i_Compile.ProposerState
	var _ = _5138_newProposer
	var _out320 _254_LiveRSL____ProposerState__i_Compile.ProposerState
	var _ = _out320
	_out320 = _308_LiveRSL____ProposerModel__i_Compile.Companion_Default___.ProposerProcessHeartbeat((replica).Dtor_proposer(), inp, clock, cur__req__set, prev__req__set)
	_5138_newProposer = _out320
	var _5139_newAcceptor _269_LiveRSL____AcceptorState__i_Compile.AcceptorState
	var _ = _5139_newAcceptor
	var _out321 _269_LiveRSL____AcceptorState__i_Compile.AcceptorState
	var _ = _out321
	_out321 = _312_LiveRSL____AcceptorModel__i_Compile.Companion_Default___.NextAcceptorState__ProcessHeartbeat((replica).Dtor_acceptor(), (inp).Dtor_msg(), (inp).Dtor_src())
	_5139_newAcceptor = _out321
	replica_k = func(_pat_let136_0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
		return func(_5140_dt__update__tmp_h0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
			return func(_pat_let137_0 _269_LiveRSL____AcceptorState__i_Compile.AcceptorState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
				return func(_5141_dt__update_hacceptor_h0 _269_LiveRSL____AcceptorState__i_Compile.AcceptorState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
					return func(_pat_let138_0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
						return func(_5142_dt__update_hproposer_h0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
							return _285_LiveRSL____ReplicaState__i_Compile.ReplicaState{_285_LiveRSL____ReplicaState__i_Compile.ReplicaState_ReplicaState{(_5140_dt__update__tmp_h0).Dtor_constants(), (_5140_dt__update__tmp_h0).Dtor_nextHeartbeatTime(), _5142_dt__update_hproposer_h0, _5141_dt__update_hacceptor_h0, (_5140_dt__update__tmp_h0).Dtor_learner(), (_5140_dt__update__tmp_h0).Dtor_executor()}}
						}(_pat_let138_0)
					}(_5138_newProposer)
				}(_pat_let137_0)
			}(_5139_newAcceptor)
		}(_pat_let136_0)
	}(replica)
	var _5143_end__time uint64
	var _ = _5143_end__time
	var _out322 uint64
	var _ = _out322
	_out322 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_5143_end__time = _out322
	_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("Replica_Next_Process_Heartbeat"), _5137_start__time, _5143_end__time)
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
	{
	}
	return replica_k, packets__sent
}
func (_this *CompanionStruct_Default___) Replica__Next__ReadClock__CheckForViewTimeout(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, clock _283_LiveRSL____CClockReading__i_Compile.CClockReading, cur__req__set *_9_Native____Io__s_Compile.MutableSet, prev__req__set *_9_Native____Io__s_Compile.MutableSet) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets, bool) {
	var noop bool
	var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
	var _ = replica_k
	var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
	var _ = packets__sent
	var _5144_start__time uint64
	var _ = _5144_start__time
	var _out323 uint64
	var _ = _out323
	_out323 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_5144_start__time = _out323
	var _5145_newProposer _254_LiveRSL____ProposerState__i_Compile.ProposerState
	var _ = _5145_newProposer
	var _out324 _254_LiveRSL____ProposerState__i_Compile.ProposerState
	var _ = _out324
	_out324, noop = _308_LiveRSL____ProposerModel__i_Compile.Companion_Default___.ProposerCheckForViewTimeout((replica).Dtor_proposer(), (clock).Dtor_t(), cur__req__set, prev__req__set)
	_5145_newProposer = _out324
	replica_k = func(_pat_let139_0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
		return func(_5146_dt__update__tmp_h0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
			return func(_pat_let140_0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
				return func(_5147_dt__update_hproposer_h0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
					return _285_LiveRSL____ReplicaState__i_Compile.ReplicaState{_285_LiveRSL____ReplicaState__i_Compile.ReplicaState_ReplicaState{(_5146_dt__update__tmp_h0).Dtor_constants(), (_5146_dt__update__tmp_h0).Dtor_nextHeartbeatTime(), _5147_dt__update_hproposer_h0, (_5146_dt__update__tmp_h0).Dtor_acceptor(), (_5146_dt__update__tmp_h0).Dtor_learner(), (_5146_dt__update__tmp_h0).Dtor_executor()}}
				}(_pat_let140_0)
			}(_5145_newProposer)
		}(_pat_let139_0)
	}(replica)
	packets__sent = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_Broadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast_CBroadcastNop{}}}}
	{
	}
	{
	}
	var _5148_end__time uint64
	var _ = _5148_end__time
	var _out325 uint64
	var _ = _out325
	_out325 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_5148_end__time = _out325
	_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("Replica_Next_ReadClock_CheckForViewTimeout"), _5144_start__time, _5148_end__time)
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
	return replica_k, packets__sent, noop
}
func (_this *CompanionStruct_Default___) Replica__Next__ReadClock__CheckForQuorumOfViewSuspicions(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, clock _283_LiveRSL____CClockReading__i_Compile.CClockReading, cur__req__set *_9_Native____Io__s_Compile.MutableSet, prev__req__set *_9_Native____Io__s_Compile.MutableSet) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets, bool) {
	var noop bool
	var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
	var _ = replica_k
	var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
	var _ = packets__sent
	var _5149_start__time uint64
	var _ = _5149_start__time
	var _out326 uint64
	var _ = _out326
	_out326 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_5149_start__time = _out326
	var _5150_newProposer _254_LiveRSL____ProposerState__i_Compile.ProposerState
	var _ = _5150_newProposer
	var _out327 _254_LiveRSL____ProposerState__i_Compile.ProposerState
	var _ = _out327
	_out327, noop = _308_LiveRSL____ProposerModel__i_Compile.Companion_Default___.ProposerCheckForQuorumOfViewSuspicions((replica).Dtor_proposer(), (clock).Dtor_t(), cur__req__set, prev__req__set)
	_5150_newProposer = _out327
	replica_k = func(_pat_let141_0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
		return func(_5151_dt__update__tmp_h0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
			return func(_pat_let142_0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
				return func(_5152_dt__update_hproposer_h0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
					return _285_LiveRSL____ReplicaState__i_Compile.ReplicaState{_285_LiveRSL____ReplicaState__i_Compile.ReplicaState_ReplicaState{(_5151_dt__update__tmp_h0).Dtor_constants(), (_5151_dt__update__tmp_h0).Dtor_nextHeartbeatTime(), _5152_dt__update_hproposer_h0, (_5151_dt__update__tmp_h0).Dtor_acceptor(), (_5151_dt__update__tmp_h0).Dtor_learner(), (_5151_dt__update__tmp_h0).Dtor_executor()}}
				}(_pat_let142_0)
			}(_5150_newProposer)
		}(_pat_let141_0)
	}(replica)
	packets__sent = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_Broadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast_CBroadcastNop{}}}}
	{
	}
	{
	}
	var _5153_end__time uint64
	var _ = _5153_end__time
	var _out328 uint64
	var _ = _out328
	_out328 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_5153_end__time = _out328
	_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("Replica_Next_ReadClock_CheckForQuorumOfViewSuspicions"), _5149_start__time, _5153_end__time)
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
	return replica_k, packets__sent, noop
}

// End of class Default__
