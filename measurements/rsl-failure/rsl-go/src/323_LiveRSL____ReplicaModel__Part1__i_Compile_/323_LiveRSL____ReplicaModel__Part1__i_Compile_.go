// Package _323_LiveRSL____ReplicaModel__Part1__i_Compile
// Dafny module _323_LiveRSL____ReplicaModel__Part1__i_Compile compiled into Go

package _323_LiveRSL____ReplicaModel__Part1__i_Compile

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
	"fmt"
	"os"
	"time"
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

var COUNTDOWN int

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
	return "_323_LiveRSL____ReplicaModel__Part1__i_Compile.Default__"
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
	return "_323_LiveRSL____ReplicaModel__Part1__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) InitReplicaState(constants _245_LiveRSL____ReplicaConstantsState__i_Compile.ReplicaConstantsState) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, *_9_Native____Io__s_Compile.MutableSet, *_9_Native____Io__s_Compile.MutableSet, *_9_Native____Io__s_Compile.MutableMap) {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
	var _ = replica
	var cur__req__set *_9_Native____Io__s_Compile.MutableSet = (*_9_Native____Io__s_Compile.MutableSet)(nil)
	var _ = cur__req__set
	var prev__req__set *_9_Native____Io__s_Compile.MutableSet = (*_9_Native____Io__s_Compile.MutableSet)(nil)
	var _ = prev__req__set
	var reply__cache__mutable *_9_Native____Io__s_Compile.MutableMap = (*_9_Native____Io__s_Compile.MutableMap)(nil)
	var _ = reply__cache__mutable
	var _5056_proposer _254_LiveRSL____ProposerState__i_Compile.ProposerState = _254_LiveRSL____ProposerState__i_Compile.Type_ProposerState_().Default().(_254_LiveRSL____ProposerState__i_Compile.ProposerState)
	var _ = _5056_proposer
	var _out246 _254_LiveRSL____ProposerState__i_Compile.ProposerState
	var _ = _out246
	var _out247 *_9_Native____Io__s_Compile.MutableSet
	var _ = _out247
	var _out248 *_9_Native____Io__s_Compile.MutableSet
	var _ = _out248
	_out246, _out247, _out248 = _308_LiveRSL____ProposerModel__i_Compile.Companion_Default___.InitProposerState(constants)
	_5056_proposer = _out246
	cur__req__set = _out247
	prev__req__set = _out248
	var _5057_acceptor _269_LiveRSL____AcceptorState__i_Compile.AcceptorState
	var _ = _5057_acceptor
	var _out249 _269_LiveRSL____AcceptorState__i_Compile.AcceptorState
	var _ = _out249
	_out249 = _312_LiveRSL____AcceptorModel__i_Compile.Companion_Default___.InitAcceptorState(constants)
	_5057_acceptor = _out249
	var _5058_learner _278_LiveRSL____LearnerState__i_Compile.CLearnerState
	var _ = _5058_learner
	var _out250 _278_LiveRSL____LearnerState__i_Compile.CLearnerState
	var _ = _out250
	_out250 = _278_LiveRSL____LearnerState__i_Compile.Companion_Default___.LearnerState__Init(constants)
	_5058_learner = _out250
	var _5059_executor _276_LiveRSL____ExecutorState__i_Compile.ExecutorState = _276_LiveRSL____ExecutorState__i_Compile.Type_ExecutorState_().Default().(_276_LiveRSL____ExecutorState__i_Compile.ExecutorState)
	var _ = _5059_executor
	var _out251 _276_LiveRSL____ExecutorState__i_Compile.ExecutorState
	var _ = _out251
	var _out252 *_9_Native____Io__s_Compile.MutableMap
	var _ = _out252
	_out251, _out252 = _321_LiveRSL____ExecutorModel__i_Compile.Companion_Default___.ExecutorInit(constants)
	_5059_executor = _out251
	reply__cache__mutable = _out252
	replica = _285_LiveRSL____ReplicaState__i_Compile.ReplicaState{_285_LiveRSL____ReplicaState__i_Compile.ReplicaState_ReplicaState{constants, uint64(0), _5056_proposer, _5057_acceptor, _5058_learner, _5059_executor}}
	{
	}
	return replica, cur__req__set, prev__req__set, reply__cache__mutable
}
func (_this *CompanionStruct_Default___) ReplicaNextProcessRequestImplCaseUncached(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, inp _217_LiveRSL____CMessage__i_Compile.CPacket, cur__req__set *_9_Native____Io__s_Compile.MutableSet, prev__req__set *_9_Native____Io__s_Compile.MutableSet, reply__cache__mutable *_9_Native____Io__s_Compile.MutableMap) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
	var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
	var _ = replica_k
	var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
	var _ = packets__sent
	{
	}
	{
	}
	{
	}
	{
	}
	var _5060_newProposer _254_LiveRSL____ProposerState__i_Compile.ProposerState
	var _ = _5060_newProposer
	var _out253 _254_LiveRSL____ProposerState__i_Compile.ProposerState
	var _ = _out253
	_out253 = _308_LiveRSL____ProposerModel__i_Compile.Companion_Default___.ProposerProcessRequest((replica).Dtor_proposer(), inp, cur__req__set, prev__req__set)
	_5060_newProposer = _out253
	replica_k = func(_pat_let111_0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
		return func(_5061_dt__update__tmp_h0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
			return func(_pat_let112_0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
				return func(_5062_dt__update_hproposer_h0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
					return _285_LiveRSL____ReplicaState__i_Compile.ReplicaState{_285_LiveRSL____ReplicaState__i_Compile.ReplicaState_ReplicaState{(_5061_dt__update__tmp_h0).Dtor_constants(), (_5061_dt__update__tmp_h0).Dtor_nextHeartbeatTime(), _5062_dt__update_hproposer_h0, (_5061_dt__update__tmp_h0).Dtor_acceptor(), (_5061_dt__update__tmp_h0).Dtor_learner(), (_5061_dt__update__tmp_h0).Dtor_executor()}}
				}(_pat_let112_0)
			}(_5060_newProposer)
		}(_pat_let111_0)
	}(replica)
	{
	}
	packets__sent = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_Broadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast_CBroadcastNop{}}}}
	{
	}
	var _5063_notCachedTime uint64
	var _ = _5063_notCachedTime
	var _out254 uint64
	var _ = _out254
	_out254 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_5063_notCachedTime = _out254
	{
	}
	{
	}
	return replica_k, packets__sent
}
func (_this *CompanionStruct_Default___) ReplicaNextProcessRequestImplCaseCachedNonReply(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, inp _217_LiveRSL____CMessage__i_Compile.CPacket, cur__req__set *_9_Native____Io__s_Compile.MutableSet, prev__req__set *_9_Native____Io__s_Compile.MutableSet, reply__cache__mutable *_9_Native____Io__s_Compile.MutableMap, cached__reply _214_LiveRSL____CTypes__i_Compile.CReply) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
	var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
	var _ = replica_k
	var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
	var _ = packets__sent
	{
	}
	{
	}
	{
	}
	{
	}
	var _5064_newProposer _254_LiveRSL____ProposerState__i_Compile.ProposerState
	var _ = _5064_newProposer
	var _out255 _254_LiveRSL____ProposerState__i_Compile.ProposerState
	var _ = _out255
	_out255 = _308_LiveRSL____ProposerModel__i_Compile.Companion_Default___.ProposerProcessRequest((replica).Dtor_proposer(), inp, cur__req__set, prev__req__set)
	_5064_newProposer = _out255
	replica_k = func(_pat_let113_0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
		return func(_5065_dt__update__tmp_h0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
			return func(_pat_let114_0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
				return func(_5066_dt__update_hproposer_h0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
					return _285_LiveRSL____ReplicaState__i_Compile.ReplicaState{_285_LiveRSL____ReplicaState__i_Compile.ReplicaState_ReplicaState{(_5065_dt__update__tmp_h0).Dtor_constants(), (_5065_dt__update__tmp_h0).Dtor_nextHeartbeatTime(), _5066_dt__update_hproposer_h0, (_5065_dt__update__tmp_h0).Dtor_acceptor(), (_5065_dt__update__tmp_h0).Dtor_learner(), (_5065_dt__update__tmp_h0).Dtor_executor()}}
				}(_pat_let114_0)
			}(_5064_newProposer)
		}(_pat_let113_0)
	}(replica)
	packets__sent = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_Broadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast_CBroadcastNop{}}}}
	{
	}
	var _5067_notReplyTime uint64
	var _ = _5067_notReplyTime
	var _out256 uint64
	var _ = _out256
	_out256 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_5067_notReplyTime = _out256
	{
	}
	return replica_k, packets__sent
}
func (_this *CompanionStruct_Default___) ReplicaNextProcessRequestImplCaseCachedOld(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, inp _217_LiveRSL____CMessage__i_Compile.CPacket, cur__req__set *_9_Native____Io__s_Compile.MutableSet, prev__req__set *_9_Native____Io__s_Compile.MutableSet, reply__cache__mutable *_9_Native____Io__s_Compile.MutableMap, cached__reply _214_LiveRSL____CTypes__i_Compile.CReply) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
	var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
	var _ = replica_k
	var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
	var _ = packets__sent
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
	var _5068_newProposer _254_LiveRSL____ProposerState__i_Compile.ProposerState
	var _ = _5068_newProposer
	var _out257 _254_LiveRSL____ProposerState__i_Compile.ProposerState
	var _ = _out257
	_out257 = _308_LiveRSL____ProposerModel__i_Compile.Companion_Default___.ProposerProcessRequest((replica).Dtor_proposer(), inp, cur__req__set, prev__req__set)
	_5068_newProposer = _out257
	replica_k = func(_pat_let115_0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
		return func(_5069_dt__update__tmp_h0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
			return func(_pat_let116_0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
				return func(_5070_dt__update_hproposer_h0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
					return _285_LiveRSL____ReplicaState__i_Compile.ReplicaState{_285_LiveRSL____ReplicaState__i_Compile.ReplicaState_ReplicaState{(_5069_dt__update__tmp_h0).Dtor_constants(), (_5069_dt__update__tmp_h0).Dtor_nextHeartbeatTime(), _5070_dt__update_hproposer_h0, (_5069_dt__update__tmp_h0).Dtor_acceptor(), (_5069_dt__update__tmp_h0).Dtor_learner(), (_5069_dt__update__tmp_h0).Dtor_executor()}}
				}(_pat_let116_0)
			}(_5068_newProposer)
		}(_pat_let115_0)
	}(replica)
	packets__sent = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_Broadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast_CBroadcastNop{}}}}
	{
	}
	var _5071_seqnoIsBeyondTime uint64
	var _ = _5071_seqnoIsBeyondTime
	var _out258 uint64
	var _ = _out258
	_out258 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_5071_seqnoIsBeyondTime = _out258
	{
	}
	{
	}
	return replica_k, packets__sent
}
func (_this *CompanionStruct_Default___) ReplicaNextProcessRequestImplCaseCachedFresh(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, inp _217_LiveRSL____CMessage__i_Compile.CPacket, reply__cache__mutable *_9_Native____Io__s_Compile.MutableMap, cached__reply _214_LiveRSL____CTypes__i_Compile.CReply) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
	var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
	var _ = replica_k
	var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
	var _ = packets__sent
	{
	}
	{
	}
	{
	}
	{
	}
	var _out259 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
	var _ = _out259
	_out259 = _321_LiveRSL____ExecutorModel__i_Compile.Companion_Default___.ExecutorProcessRequest((replica).Dtor_executor(), inp, reply__cache__mutable)
	packets__sent = _out259
	{
	}
	replica_k = replica
	var _5072_isCachedTime uint64
	var _ = _5072_isCachedTime
	var _out260 uint64
	var _ = _out260
	_out260 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_5072_isCachedTime = _out260
	{
	}
	return replica_k, packets__sent
}
func (_this *CompanionStruct_Default___) Replica__Next__Process__Request(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, inp _217_LiveRSL____CMessage__i_Compile.CPacket, cur__req__set *_9_Native____Io__s_Compile.MutableSet, prev__req__set *_9_Native____Io__s_Compile.MutableSet, reply__cache__mutable *_9_Native____Io__s_Compile.MutableMap) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
	goto TAIL_CALL_START
TAIL_CALL_START:
	fmt.Printf("Received request\n")
	// TONY: Start leader crash timer here
	if replica.Dtor_constants().Dtor_my__index() == 0 {
		f := func() {
			fmt.Printf("Detonated: %vus\n", COUNTDOWN)
			os.Exit(1)
		}
		fmt.Printf("Start countdown timer\n")
		time.AfterFunc(time.Duration(COUNTDOWN)*time.Microsecond, f)
	}

	var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
	var _ = replica_k
	var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
	var _ = packets__sent
	var _5073_cached bool
	var _ = _5073_cached
	var _5074_cached__reply _214_LiveRSL____CTypes__i_Compile.CReply
	var _ = _5074_cached__reply
	var _out261 bool
	var _ = _out261
	var _out262 interface{}
	var _ = _out262
	_out261, _out262 = (reply__cache__mutable).TryGetValue((inp).Dtor_src())
	_5073_cached = _out261

	if !(_5073_cached) {
		var _out263 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
		var _ = _out263
		var _out264 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
		var _ = _out264
		_out263, _out264 = Companion_Default___.ReplicaNextProcessRequestImplCaseUncached(replica, inp, cur__req__set, prev__req__set, reply__cache__mutable)
		replica_k = _out263
		packets__sent = _out264
	} else {
		_5074_cached__reply = _out262.(_214_LiveRSL____CTypes__i_Compile.CReply)
		if !((_5074_cached__reply).Is_CReply()) {
			var _out265 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
			var _ = _out265
			var _out266 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
			var _ = _out266
			_out265, _out266 = Companion_Default___.ReplicaNextProcessRequestImplCaseCachedNonReply(replica, inp, cur__req__set, prev__req__set, reply__cache__mutable, _5074_cached__reply)
			replica_k = _out265
			packets__sent = _out266
		} else if (((inp).Dtor_msg()).Dtor_seqno()) > ((_5074_cached__reply).Dtor_seqno()) {
			var _out267 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
			var _ = _out267
			var _out268 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
			var _ = _out268
			_out267, _out268 = Companion_Default___.ReplicaNextProcessRequestImplCaseCachedOld(replica, inp, cur__req__set, prev__req__set, reply__cache__mutable, _5074_cached__reply)
			replica_k = _out267
			packets__sent = _out268
		} else {
			var _out269 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
			var _ = _out269
			var _out270 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
			var _ = _out270
			_out269, _out270 = Companion_Default___.ReplicaNextProcessRequestImplCaseCachedFresh(replica, inp, reply__cache__mutable, _5074_cached__reply)
			replica_k = _out269
			packets__sent = _out270
		}
	}
	{
	}
	return replica_k, packets__sent
}
func (_this *CompanionStruct_Default___) Replica__Next__Process__1a(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, inp _217_LiveRSL____CMessage__i_Compile.CPacket) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
	var _ = replica_k
	var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
	var _ = packets__sent
	var _5075_newAcceptor _269_LiveRSL____AcceptorState__i_Compile.AcceptorState
	var _ = _5075_newAcceptor
	var _5076_packets _217_LiveRSL____CMessage__i_Compile.CBroadcast
	var _ = _5076_packets
	var _out271 _269_LiveRSL____AcceptorState__i_Compile.AcceptorState
	var _ = _out271
	var _out272 _217_LiveRSL____CMessage__i_Compile.CBroadcast
	var _ = _out272
	_out271, _out272 = _312_LiveRSL____AcceptorModel__i_Compile.Companion_Default___.NextAcceptorState__Phase1((replica).Dtor_acceptor(), (inp).Dtor_msg(), (inp).Dtor_src())
	_5075_newAcceptor = _out271
	_5076_packets = _out272
	replica_k = func(_pat_let117_0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
		return func(_5077_dt__update__tmp_h0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
			return func(_pat_let118_0 _269_LiveRSL____AcceptorState__i_Compile.AcceptorState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
				return func(_5078_dt__update_hacceptor_h0 _269_LiveRSL____AcceptorState__i_Compile.AcceptorState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
					return _285_LiveRSL____ReplicaState__i_Compile.ReplicaState{_285_LiveRSL____ReplicaState__i_Compile.ReplicaState_ReplicaState{(_5077_dt__update__tmp_h0).Dtor_constants(), (_5077_dt__update__tmp_h0).Dtor_nextHeartbeatTime(), (_5077_dt__update__tmp_h0).Dtor_proposer(), _5078_dt__update_hacceptor_h0, (_5077_dt__update__tmp_h0).Dtor_learner(), (_5077_dt__update__tmp_h0).Dtor_executor()}}
				}(_pat_let118_0)
			}(_5075_newAcceptor)
		}(_pat_let117_0)
	}(replica)
	{
	}
	{
	}
	{
	}
	packets__sent = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_Broadcast{_5076_packets}}
	return replica_k, packets__sent
}

// End of class Default__
