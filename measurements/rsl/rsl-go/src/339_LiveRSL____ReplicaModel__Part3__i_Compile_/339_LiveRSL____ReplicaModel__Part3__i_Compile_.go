// Package _339_LiveRSL____ReplicaModel__Part3__i_Compile
// Dafny module _339_LiveRSL____ReplicaModel__Part3__i_Compile compiled into Go

package _339_LiveRSL____ReplicaModel__Part3__i_Compile

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
	return "_339_LiveRSL____ReplicaModel__Part3__i_Compile.Default__"
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
	return "_339_LiveRSL____ReplicaModel__Part3__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) Replica__Next__Process__2b(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, inp _217_LiveRSL____CMessage__i_Compile.CPacket) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
	var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
	var _ = replica_k
	var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
	var _ = packets__sent
	var _5106_start__time uint64
	var _ = _5106_start__time
	var _out300 uint64
	var _ = _out300
	_out300 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_5106_start__time = _out300
	{
	}
	{
	}
	{
	}
	var _5107_copn _214_LiveRSL____CTypes__i_Compile.COperationNumber
	var _ = _5107_copn
	_5107_copn = ((inp).Dtor_msg()).Dtor_opn__2b()
	var _5108_cop__learnable bool
	var _ = _5108_cop__learnable
	_5108_cop__learnable = (((((replica).Dtor_executor()).Dtor_ops__complete()).Dtor_n()) < ((_5107_copn).Dtor_n())) || ((((((replica).Dtor_executor()).Dtor_ops__complete()).Dtor_n()) == ((_5107_copn).Dtor_n())) && ((((replica).Dtor_executor()).Dtor_next__op__to__execute()).Is_COutstandingOpUnknown()))
	{
	}
	if _5108_cop__learnable {
		var _5109_newLearner _278_LiveRSL____LearnerState__i_Compile.CLearnerState
		var _ = _5109_newLearner
		var _out301 _278_LiveRSL____LearnerState__i_Compile.CLearnerState
		var _ = _out301
		_out301 = _316_LiveRSL____LearnerModel__i_Compile.Companion_Default___.LearnerModel__Process2b((replica).Dtor_learner(), (replica).Dtor_executor(), inp)
		_5109_newLearner = _out301
		replica_k = func(_pat_let126_0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
			return func(_5110_dt__update__tmp_h0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
				return func(_pat_let127_0 _278_LiveRSL____LearnerState__i_Compile.CLearnerState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
					return func(_5111_dt__update_hlearner_h0 _278_LiveRSL____LearnerState__i_Compile.CLearnerState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
						return _285_LiveRSL____ReplicaState__i_Compile.ReplicaState{_285_LiveRSL____ReplicaState__i_Compile.ReplicaState_ReplicaState{(_5110_dt__update__tmp_h0).Dtor_constants(), (_5110_dt__update__tmp_h0).Dtor_nextHeartbeatTime(), (_5110_dt__update__tmp_h0).Dtor_proposer(), (_5110_dt__update__tmp_h0).Dtor_acceptor(), _5111_dt__update_hlearner_h0, (_5110_dt__update__tmp_h0).Dtor_executor()}}
					}(_pat_let127_0)
				}(_5109_newLearner)
			}(_pat_let126_0)
		}(replica)
	} else {
		replica_k = replica
	}
	packets__sent = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_Broadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast_CBroadcastNop{}}}}
	var _5112_end__time uint64
	var _ = _5112_end__time
	var _out302 uint64
	var _ = _out302
	_out302 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_5112_end__time = _out302
	_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("Replica_Next_Process_2b"), _5106_start__time, _5112_end__time)
	return replica_k, packets__sent
}
func (_this *CompanionStruct_Default___) Replica__Next__Spontaneous__MaybeEnterNewViewAndSend1a(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets, bool) {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
	var _ = replica_k
	var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
	var _ = packets__sent
	var _5113_start__time uint64
	var _ = _5113_start__time
	var _out303 uint64
	var _ = _out303
	_out303 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_5113_start__time = _out303
	var _5114_newProposer _254_LiveRSL____ProposerState__i_Compile.ProposerState
	var _ = _5114_newProposer
	var _5115_packets _217_LiveRSL____CMessage__i_Compile.CBroadcast
	var _ = _5115_packets
	var _out304 _254_LiveRSL____ProposerState__i_Compile.ProposerState
	var _ = _out304
	var _out305 _217_LiveRSL____CMessage__i_Compile.CBroadcast
	var _ = _out305
	var noop bool
	_out304, _out305, noop = _308_LiveRSL____ProposerModel__i_Compile.Companion_Default___.ProposerMaybeEnterNewViewAndSend1a((replica).Dtor_proposer())
	_5114_newProposer = _out304
	_5115_packets = _out305
	replica_k = func(_pat_let128_0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
		return func(_5116_dt__update__tmp_h0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
			return func(_pat_let129_0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
				return func(_5117_dt__update_hproposer_h0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
					return _285_LiveRSL____ReplicaState__i_Compile.ReplicaState{_285_LiveRSL____ReplicaState__i_Compile.ReplicaState_ReplicaState{(_5116_dt__update__tmp_h0).Dtor_constants(), (_5116_dt__update__tmp_h0).Dtor_nextHeartbeatTime(), _5117_dt__update_hproposer_h0, (_5116_dt__update__tmp_h0).Dtor_acceptor(), (_5116_dt__update__tmp_h0).Dtor_learner(), (_5116_dt__update__tmp_h0).Dtor_executor()}}
				}(_pat_let129_0)
			}(_5114_newProposer)
		}(_pat_let128_0)
	}(replica)
	packets__sent = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_Broadcast{_5115_packets}}
	var _5118_end__time uint64
	var _ = _5118_end__time
	var _out306 uint64
	var _ = _out306
	_out306 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_5118_end__time = _out306
	_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("Replica_Next_Spontaneous_MaybeEnterNewViewAndSend1a"), _5113_start__time, _5118_end__time)
	return replica_k, packets__sent, noop
}
func (_this *CompanionStruct_Default___) Replica__Next__Spontaneous__MaybeEnterPhase2(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
	var _ = replica_k
	var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
	var _ = packets__sent
	var _5119_start__time uint64
	var _ = _5119_start__time
	var _out307 uint64
	var _ = _out307
	_out307 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_5119_start__time = _out307
	var _5120_newProposer _254_LiveRSL____ProposerState__i_Compile.ProposerState
	var _ = _5120_newProposer
	var _5121_packets _217_LiveRSL____CMessage__i_Compile.CBroadcast
	var _ = _5121_packets
	var _out308 _254_LiveRSL____ProposerState__i_Compile.ProposerState
	var _ = _out308
	var _out309 _217_LiveRSL____CMessage__i_Compile.CBroadcast
	var _ = _out309
	_out308, _out309 = _308_LiveRSL____ProposerModel__i_Compile.Companion_Default___.ProposerMaybeEnterPhase2((replica).Dtor_proposer(), ((replica).Dtor_acceptor()).Dtor_log__truncation__point())
	_5120_newProposer = _out308
	_5121_packets = _out309
	replica_k = func(_pat_let130_0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
		return func(_5122_dt__update__tmp_h0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
			return func(_pat_let131_0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
				return func(_5123_dt__update_hproposer_h0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
					return _285_LiveRSL____ReplicaState__i_Compile.ReplicaState{_285_LiveRSL____ReplicaState__i_Compile.ReplicaState_ReplicaState{(_5122_dt__update__tmp_h0).Dtor_constants(), (_5122_dt__update__tmp_h0).Dtor_nextHeartbeatTime(), _5123_dt__update_hproposer_h0, (_5122_dt__update__tmp_h0).Dtor_acceptor(), (_5122_dt__update__tmp_h0).Dtor_learner(), (_5122_dt__update__tmp_h0).Dtor_executor()}}
				}(_pat_let131_0)
			}(_5120_newProposer)
		}(_pat_let130_0)
	}(replica)
	packets__sent = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_Broadcast{_5121_packets}}
	var _5124_end__time uint64
	var _ = _5124_end__time
	var _out310 uint64
	var _ = _out310
	_out310 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_5124_end__time = _out310
	_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("Replica_Next_Spontaneous_MaybeEnterPhase2"), _5119_start__time, _5124_end__time)
	return replica_k, packets__sent
}
func (_this *CompanionStruct_Default___) Replica__Next__Spontaneous__MaybeNominateValueAndSend2a(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, clock _283_LiveRSL____CClockReading__i_Compile.CClockReading) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
	var _ = replica_k
	var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
	var _ = packets__sent
	var _5125_start__time uint64
	var _ = _5125_start__time
	var _out311 uint64
	var _ = _out311
	_out311 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_5125_start__time = _out311
	var _5126_newProposer _254_LiveRSL____ProposerState__i_Compile.ProposerState
	var _ = _5126_newProposer
	var _5127_packets _217_LiveRSL____CMessage__i_Compile.CBroadcast
	var _ = _5127_packets
	var _out312 _254_LiveRSL____ProposerState__i_Compile.ProposerState
	var _ = _out312
	var _out313 _217_LiveRSL____CMessage__i_Compile.CBroadcast
	var _ = _out313
	_out312, _out313 = _308_LiveRSL____ProposerModel__i_Compile.Companion_Default___.ProposerMaybeNominateValueAndSend2a((replica).Dtor_proposer(), (clock).Dtor_t(), ((replica).Dtor_acceptor()).Dtor_log__truncation__point())
	_5126_newProposer = _out312
	_5127_packets = _out313
	replica_k = func(_pat_let132_0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
		return func(_5128_dt__update__tmp_h0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
			return func(_pat_let133_0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
				return func(_5129_dt__update_hproposer_h0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
					return _285_LiveRSL____ReplicaState__i_Compile.ReplicaState{_285_LiveRSL____ReplicaState__i_Compile.ReplicaState_ReplicaState{(_5128_dt__update__tmp_h0).Dtor_constants(), (_5128_dt__update__tmp_h0).Dtor_nextHeartbeatTime(), _5129_dt__update_hproposer_h0, (_5128_dt__update__tmp_h0).Dtor_acceptor(), (_5128_dt__update__tmp_h0).Dtor_learner(), (_5128_dt__update__tmp_h0).Dtor_executor()}}
				}(_pat_let133_0)
			}(_5126_newProposer)
		}(_pat_let132_0)
	}(replica)
	packets__sent = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_Broadcast{_5127_packets}}
	var _5130_end__time uint64
	var _ = _5130_end__time
	var _out314 uint64
	var _ = _out314
	_out314 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_5130_end__time = _out314
	_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("Replica_Next_Spontaneous_MaybeNominateValueAndSend2a"), _5125_start__time, _5130_end__time)
	return replica_k, packets__sent
}

// End of class Default__
