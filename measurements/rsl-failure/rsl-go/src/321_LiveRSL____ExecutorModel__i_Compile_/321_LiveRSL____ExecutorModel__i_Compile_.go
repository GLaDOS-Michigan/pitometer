// Package _321_LiveRSL____ExecutorModel__i_Compile
// Dafny module _321_LiveRSL____ExecutorModel__i_Compile compiled into Go

package _321_LiveRSL____ExecutorModel__i_Compile

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
	return "_321_LiveRSL____ExecutorModel__i_Compile.Default__"
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
	return "_321_LiveRSL____ExecutorModel__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) HandleRequestBatchImpl(state uint64, batch _dafny.Seq, reply__cache__mutable *_9_Native____Io__s_Compile.MutableMap) (uint64, _dafny.Seq) {
	var final__state uint64 = 0
	var _ = final__state
	var replies__seq _dafny.Seq = _dafny.EmptySeq
	var _ = replies__seq
	{
	}
	{
	}
	{
	}
	var _5009_i uint64
	var _ = _5009_i
	_5009_i = uint64(0)
	{
	}
	final__state = state
	{
	}
	var _5010_repliesArr *_dafny.Array
	var _ = _5010_repliesArr
	var _nw11 = _dafny.NewArrayWithValue(_214_LiveRSL____CTypes__i_Compile.Type_CReply_().Default().(_214_LiveRSL____CTypes__i_Compile.CReply), (batch).CardinalityInt())
	var _ = _nw11
	_5010_repliesArr = _nw11
	{
	}
	for (_5009_i) < (uint64((batch).CardinalityUint64())) {
		{
		}
		{
		}
		{
		}
		var _5011_new__state uint64
		var _ = _5011_new__state
		var _5012_reply _197_LiveRSL____AppInterface__i_Compile.CAppMessage
		var _ = _5012_reply
		var _out227 uint64
		var _ = _out227
		var _out228 _197_LiveRSL____AppInterface__i_Compile.CAppMessage
		var _ = _out228
		_out227, _out228 = _197_LiveRSL____AppInterface__i_Compile.Companion_Default___.HandleAppRequest(final__state, ((batch).IndexUint(_5009_i).(_214_LiveRSL____CTypes__i_Compile.CRequest)).Dtor_request())
		_5011_new__state = _out227
		_5012_reply = _out228
		var _5013_newReply _214_LiveRSL____CTypes__i_Compile.CReply
		var _ = _5013_newReply
		_5013_newReply = _214_LiveRSL____CTypes__i_Compile.CReply{_214_LiveRSL____CTypes__i_Compile.CReply_CReply{((batch).IndexUint(_5009_i).(_214_LiveRSL____CTypes__i_Compile.CRequest)).Dtor_client(), ((batch).IndexUint(_5009_i).(_214_LiveRSL____CTypes__i_Compile.CRequest)).Dtor_seqno(), _5012_reply}}
		{
		}
		*((_5010_repliesArr).Index(_dafny.IntOfAny((_5009_i)))) = _5013_newReply
		{
		}
		final__state = _5011_new__state
		Companion_Default___.UpdateReplyCache(reply__cache__mutable, ((batch).IndexUint(_5009_i).(_214_LiveRSL____CTypes__i_Compile.CRequest)).Dtor_client(), _5013_newReply, _5012_reply, _5009_i, batch)
		_5009_i = (_5009_i) + (uint64(1))
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
	}
	replies__seq = (_5010_repliesArr).RangeToSeq(_dafny.NilInt, _dafny.NilInt)
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
	return final__state, replies__seq
}
func (_this *CompanionStruct_Default___) UpdateReplyCache(reply__cache__mutable *_9_Native____Io__s_Compile.MutableMap, ep _9_Native____Io__s_Compile.EndPoint, newReply _214_LiveRSL____CTypes__i_Compile.CReply, reply _197_LiveRSL____AppInterface__i_Compile.CAppMessage, i uint64, batch _dafny.Seq) {
	{
	}
	{
	}
	var _5014_staleEntry _9_Native____Io__s_Compile.EndPoint = _9_Native____Io__s_Compile.Type_EndPoint_().Default().(_9_Native____Io__s_Compile.EndPoint)
	var _ = _5014_staleEntry
	var _5015_cache__size uint64
	var _ = _5015_cache__size
	var _out229 uint64
	var _ = _out229
	_out229 = (reply__cache__mutable).SizeModest()
	_5015_cache__size = _out229
	if (_5015_cache__size) == ((uint64(256)) - (func() uint64 { return (uint64(1)) })()) {
		for _iter18 := _dafny.Iterate((_9_Native____Io__s_Compile.Companion_MutableMap_.MapOf(reply__cache__mutable)).Keys().Elements()); ; {
			_val18, _ok18 := _iter18()
			if !_ok18 {
				break
			}
			_assign_such_that_12 := _val18.(_9_Native____Io__s_Compile.EndPoint)
			_5014_staleEntry = _assign_such_that_12
			if (_9_Native____Io__s_Compile.Companion_MutableMap_.MapOf(reply__cache__mutable)).Contains(_5014_staleEntry) {
				goto L_ASSIGN_SUCH_THAT_12
			}
		}
		panic("assign-such-that search produced no value (line 306)")
	L_ASSIGN_SUCH_THAT_12:
		{
		}
		(reply__cache__mutable).Remove(_5014_staleEntry)
	} else {
	}
	{
	}
	{
	}
	{
	}
	{
	}
	(reply__cache__mutable).Set(ep, newReply)
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
	{
	}
	{
	}
	{
	}
	{
	}
	return
}
func (_this *CompanionStruct_Default___) ExecutorInit(ccons _245_LiveRSL____ReplicaConstantsState__i_Compile.ReplicaConstantsState) (_276_LiveRSL____ExecutorState__i_Compile.ExecutorState, *_9_Native____Io__s_Compile.MutableMap) {
	var cs _276_LiveRSL____ExecutorState__i_Compile.ExecutorState = _276_LiveRSL____ExecutorState__i_Compile.Type_ExecutorState_().Default().(_276_LiveRSL____ExecutorState__i_Compile.ExecutorState)
	var _ = cs
	var reply__cache__mutable *_9_Native____Io__s_Compile.MutableMap = (*_9_Native____Io__s_Compile.MutableMap)(nil)
	var _ = reply__cache__mutable
	{
	}
	{
	}
	var _5016_app__state uint64
	var _ = _5016_app__state
	var _out230 uint64
	var _ = _out230
	_out230 = _197_LiveRSL____AppInterface__i_Compile.Companion_Default___.CAppState__Init()
	_5016_app__state = _out230
	cs = _276_LiveRSL____ExecutorState__i_Compile.ExecutorState{_276_LiveRSL____ExecutorState__i_Compile.ExecutorState_ExecutorState{ccons, _5016_app__state, _214_LiveRSL____CTypes__i_Compile.COperationNumber{_214_LiveRSL____CTypes__i_Compile.COperationNumber_COperationNumber{uint64(0)}}, _214_LiveRSL____CTypes__i_Compile.CBallot{_214_LiveRSL____CTypes__i_Compile.CBallot_CBallot{uint64(0), uint64(0)}}, _276_LiveRSL____ExecutorState__i_Compile.COutstandingOperation{_276_LiveRSL____ExecutorState__i_Compile.COutstandingOperation_COutstandingOpUnknown{}}}}
	var _out231 *_9_Native____Io__s_Compile.MutableMap
	var _ = _out231
	_out231 = _9_Native____Io__s_Compile.Companion_MutableMap_.EmptyMap()
	reply__cache__mutable = _out231
	{
	}
	{
	}
	return cs, reply__cache__mutable
}
func (_this *CompanionStruct_Default___) ExecutorGetDecision(cs _276_LiveRSL____ExecutorState__i_Compile.ExecutorState, cbal _214_LiveRSL____CTypes__i_Compile.CBallot, copn _214_LiveRSL____CTypes__i_Compile.COperationNumber, ca _dafny.Seq) _276_LiveRSL____ExecutorState__i_Compile.ExecutorState {
	var cs_k _276_LiveRSL____ExecutorState__i_Compile.ExecutorState = _276_LiveRSL____ExecutorState__i_Compile.Type_ExecutorState_().Default().(_276_LiveRSL____ExecutorState__i_Compile.ExecutorState)
	var _ = cs_k
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
	cs_k = func(_pat_let99_0 _276_LiveRSL____ExecutorState__i_Compile.ExecutorState) _276_LiveRSL____ExecutorState__i_Compile.ExecutorState {
		return func(_5017_dt__update__tmp_h1 _276_LiveRSL____ExecutorState__i_Compile.ExecutorState) _276_LiveRSL____ExecutorState__i_Compile.ExecutorState {
			return func(_pat_let100_0 _276_LiveRSL____ExecutorState__i_Compile.COutstandingOperation) _276_LiveRSL____ExecutorState__i_Compile.ExecutorState {
				return func(_5018_dt__update_hnext__op__to__execute_h1 _276_LiveRSL____ExecutorState__i_Compile.COutstandingOperation) _276_LiveRSL____ExecutorState__i_Compile.ExecutorState {
					return _276_LiveRSL____ExecutorState__i_Compile.ExecutorState{_276_LiveRSL____ExecutorState__i_Compile.ExecutorState_ExecutorState{(_5017_dt__update__tmp_h1).Dtor_constants(), (_5017_dt__update__tmp_h1).Dtor_app(), (_5017_dt__update__tmp_h1).Dtor_ops__complete(), (_5017_dt__update__tmp_h1).Dtor_max__bal__reflected(), _5018_dt__update_hnext__op__to__execute_h1}}
				}(_pat_let100_0)
			}(_276_LiveRSL____ExecutorState__i_Compile.COutstandingOperation{_276_LiveRSL____ExecutorState__i_Compile.COutstandingOperation_COutstandingOpKnown{ca, cbal}})
		}(_pat_let99_0)
	}(cs)
	{
	}
	return cs_k
}
func (_this *CompanionStruct_Default___) GetPacketsFromRepliesImpl(me _9_Native____Io__s_Compile.EndPoint, requests _dafny.Seq, replies _dafny.Seq) _dafny.Seq {
	var cout__seq _dafny.Seq = _dafny.EmptySeq
	var _ = cout__seq
	var _5019_i uint64
	var _ = _5019_i
	_5019_i = uint64(0)
	{
	}
	var _5020_coutArr *_dafny.Array
	var _ = _5020_coutArr
	var _nw12 = _dafny.NewArrayWithValue(_217_LiveRSL____CMessage__i_Compile.Type_CPacket_().Default().(_217_LiveRSL____CMessage__i_Compile.CPacket), (replies).CardinalityInt())
	var _ = _nw12
	_5020_coutArr = _nw12
	for (_5019_i) < (uint64((replies).CardinalityUint64())) {
		{
		}
		var _5021_cmsg _217_LiveRSL____CMessage__i_Compile.CMessage
		var _ = _5021_cmsg
		_5021_cmsg = _217_LiveRSL____CMessage__i_Compile.CMessage{_217_LiveRSL____CMessage__i_Compile.CMessage_CMessage__Reply{((requests).IndexUint(_5019_i).(_214_LiveRSL____CTypes__i_Compile.CRequest)).Dtor_seqno(), ((replies).IndexUint(_5019_i).(_214_LiveRSL____CTypes__i_Compile.CReply)).Dtor_reply()}}
		var _5022_cp _217_LiveRSL____CMessage__i_Compile.CPacket
		var _ = _5022_cp
		_5022_cp = _217_LiveRSL____CMessage__i_Compile.CPacket{_217_LiveRSL____CMessage__i_Compile.CPacket_CPacket{((requests).IndexUint(_5019_i).(_214_LiveRSL____CTypes__i_Compile.CRequest)).Dtor_client(), me, _5021_cmsg}}
		{
		}
		*((_5020_coutArr).Index(_dafny.IntOfAny((_5019_i)))) = _5022_cp
		_5019_i = (_5019_i) + (uint64(1))
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
	{
	}
	cout__seq = (_5020_coutArr).RangeToSeq(_dafny.NilInt, _dafny.NilInt)
	return cout__seq
}
func (_this *CompanionStruct_Default___) ExecutorExecute(cs _276_LiveRSL____ExecutorState__i_Compile.ExecutorState, reply__cache__mutable *_9_Native____Io__s_Compile.MutableMap) (_276_LiveRSL____ExecutorState__i_Compile.ExecutorState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
	var cs_k _276_LiveRSL____ExecutorState__i_Compile.ExecutorState = _276_LiveRSL____ExecutorState__i_Compile.Type_ExecutorState_().Default().(_276_LiveRSL____ExecutorState__i_Compile.ExecutorState)
	var _ = cs_k
	var cout _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
	var _ = cout
	var _5023_cv _dafny.Seq
	var _ = _5023_cv
	_5023_cv = ((cs).Dtor_next__op__to__execute()).Dtor_v()
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
	var _5024_final__state uint64 = 0
	var _ = _5024_final__state
	var _5025_creplies _dafny.Seq = _dafny.EmptySeq
	var _ = _5025_creplies
	{
	}
	var _5026_start__time__request__batch uint64
	var _ = _5026_start__time__request__batch
	var _out232 uint64
	var _ = _out232
	_out232 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_5026_start__time__request__batch = _out232
	var _out233 uint64
	var _ = _out233
	var _out234 _dafny.Seq
	var _ = _out234
	_out233, _out234 = Companion_Default___.HandleRequestBatchImpl((cs).Dtor_app(), _5023_cv, reply__cache__mutable)
	_5024_final__state = _out233
	_5025_creplies = _out234
	var _5027_end__time__request__batch uint64
	var _ = _5027_end__time__request__batch
	var _out235 uint64
	var _ = _out235
	_out235 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_5027_end__time__request__batch = _out235
	_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("ExecutorExecute_HandleRequestBatch"), _5026_start__time__request__batch, _5027_end__time__request__batch)
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
	var _5028_newMaxBalReflected _214_LiveRSL____CTypes__i_Compile.CBallot
	var _ = _5028_newMaxBalReflected
	_5028_newMaxBalReflected = (func() _214_LiveRSL____CTypes__i_Compile.CBallot {
		if _214_LiveRSL____CTypes__i_Compile.Companion_Default___.CBallotIsNotGreaterThan((cs).Dtor_max__bal__reflected(), ((cs).Dtor_next__op__to__execute()).Dtor_bal()) {
			return ((cs).Dtor_next__op__to__execute()).Dtor_bal()
		}
		return (cs).Dtor_max__bal__reflected()
	})()
	cs_k = func(_pat_let101_0 _276_LiveRSL____ExecutorState__i_Compile.ExecutorState) _276_LiveRSL____ExecutorState__i_Compile.ExecutorState {
		return func(_5029_dt__update__tmp_h0 _276_LiveRSL____ExecutorState__i_Compile.ExecutorState) _276_LiveRSL____ExecutorState__i_Compile.ExecutorState {
			return func(_pat_let102_0 _276_LiveRSL____ExecutorState__i_Compile.COutstandingOperation) _276_LiveRSL____ExecutorState__i_Compile.ExecutorState {
				return func(_5030_dt__update_hnext__op__to__execute_h0 _276_LiveRSL____ExecutorState__i_Compile.COutstandingOperation) _276_LiveRSL____ExecutorState__i_Compile.ExecutorState {
					return func(_pat_let103_0 _214_LiveRSL____CTypes__i_Compile.CBallot) _276_LiveRSL____ExecutorState__i_Compile.ExecutorState {
						return func(_5031_dt__update_hmax__bal__reflected_h0 _214_LiveRSL____CTypes__i_Compile.CBallot) _276_LiveRSL____ExecutorState__i_Compile.ExecutorState {
							return func(_pat_let104_0 _214_LiveRSL____CTypes__i_Compile.COperationNumber) _276_LiveRSL____ExecutorState__i_Compile.ExecutorState {
								return func(_5032_dt__update_hops__complete_h0 _214_LiveRSL____CTypes__i_Compile.COperationNumber) _276_LiveRSL____ExecutorState__i_Compile.ExecutorState {
									return func(_pat_let105_0 uint64) _276_LiveRSL____ExecutorState__i_Compile.ExecutorState {
										return func(_5033_dt__update_happ_h0 uint64) _276_LiveRSL____ExecutorState__i_Compile.ExecutorState {
											return _276_LiveRSL____ExecutorState__i_Compile.ExecutorState{_276_LiveRSL____ExecutorState__i_Compile.ExecutorState_ExecutorState{(_5029_dt__update__tmp_h0).Dtor_constants(), _5033_dt__update_happ_h0, _5032_dt__update_hops__complete_h0, _5031_dt__update_hmax__bal__reflected_h0, _5030_dt__update_hnext__op__to__execute_h0}}
										}(_pat_let105_0)
									}(_5024_final__state)
								}(_pat_let104_0)
							}(_214_LiveRSL____CTypes__i_Compile.COperationNumber{_214_LiveRSL____CTypes__i_Compile.COperationNumber_COperationNumber{(((cs).Dtor_ops__complete()).Dtor_n()) + (uint64(1))}})
						}(_pat_let103_0)
					}(_5028_newMaxBalReflected)
				}(_pat_let102_0)
			}(_276_LiveRSL____ExecutorState__i_Compile.COutstandingOperation{_276_LiveRSL____ExecutorState__i_Compile.COutstandingOperation_COutstandingOpUnknown{}})
		}(_pat_let101_0)
	}(cs)
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
	var _5034_i _dafny.Int
	var _ = _5034_i
	_5034_i = ((_5023_cv).Cardinality()).Minus(_dafny.IntOfInt64(1))
	{
	}
	{
	}
	{
	}
	var _5035_cme _9_Native____Io__s_Compile.EndPoint
	var _ = _5035_cme
	_5035_cme = (((((cs).Dtor_constants()).Dtor_all()).Dtor_config()).Dtor_replica__ids()).IndexUint(((cs).Dtor_constants()).Dtor_my__index()).(_9_Native____Io__s_Compile.EndPoint)
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
	{
	}
	var _5036_start__time__get__packets uint64
	var _ = _5036_start__time__get__packets
	var _out236 uint64
	var _ = _out236
	_out236 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_5036_start__time__get__packets = _out236
	var _5037_packets _dafny.Seq
	var _ = _5037_packets
	var _out237 _dafny.Seq
	var _ = _out237
	_out237 = Companion_Default___.GetPacketsFromRepliesImpl(_5035_cme, _5023_cv, _5025_creplies)
	_5037_packets = _out237
	cout = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_PacketSequence{_5037_packets}}
	var _5038_end__time__get__packets uint64
	var _ = _5038_end__time__get__packets
	var _out238 uint64
	var _ = _out238
	_out238 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_5038_end__time__get__packets = _out238
	_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("ExecutorExecute_GetPackets"), _5036_start__time__get__packets, _5038_end__time__get__packets)
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
	{
	}
	return cs_k, cout
}
func (_this *CompanionStruct_Default___) ExecutorProcessAppStateSupply(cs _276_LiveRSL____ExecutorState__i_Compile.ExecutorState, cinp _217_LiveRSL____CMessage__i_Compile.CPacket) (_276_LiveRSL____ExecutorState__i_Compile.ExecutorState, *_9_Native____Io__s_Compile.MutableMap) {
	var cs_k _276_LiveRSL____ExecutorState__i_Compile.ExecutorState = _276_LiveRSL____ExecutorState__i_Compile.Type_ExecutorState_().Default().(_276_LiveRSL____ExecutorState__i_Compile.ExecutorState)
	var _ = cs_k
	var reply__cache__mutable *_9_Native____Io__s_Compile.MutableMap = (*_9_Native____Io__s_Compile.MutableMap)(nil)
	var _ = reply__cache__mutable
	{
	}
	{
	}
	{
	}
	{
	}
	var _5039_cm _217_LiveRSL____CMessage__i_Compile.CMessage
	var _ = _5039_cm
	_5039_cm = (cinp).Dtor_msg()
	cs_k = func(_pat_let106_0 _276_LiveRSL____ExecutorState__i_Compile.ExecutorState) _276_LiveRSL____ExecutorState__i_Compile.ExecutorState {
		return func(_5040_dt__update__tmp_h1 _276_LiveRSL____ExecutorState__i_Compile.ExecutorState) _276_LiveRSL____ExecutorState__i_Compile.ExecutorState {
			return func(_pat_let107_0 _276_LiveRSL____ExecutorState__i_Compile.COutstandingOperation) _276_LiveRSL____ExecutorState__i_Compile.ExecutorState {
				return func(_5041_dt__update_hnext__op__to__execute_h1 _276_LiveRSL____ExecutorState__i_Compile.COutstandingOperation) _276_LiveRSL____ExecutorState__i_Compile.ExecutorState {
					return func(_pat_let108_0 _214_LiveRSL____CTypes__i_Compile.CBallot) _276_LiveRSL____ExecutorState__i_Compile.ExecutorState {
						return func(_5042_dt__update_hmax__bal__reflected_h1 _214_LiveRSL____CTypes__i_Compile.CBallot) _276_LiveRSL____ExecutorState__i_Compile.ExecutorState {
							return func(_pat_let109_0 _214_LiveRSL____CTypes__i_Compile.COperationNumber) _276_LiveRSL____ExecutorState__i_Compile.ExecutorState {
								return func(_5043_dt__update_hops__complete_h1 _214_LiveRSL____CTypes__i_Compile.COperationNumber) _276_LiveRSL____ExecutorState__i_Compile.ExecutorState {
									return func(_pat_let110_0 uint64) _276_LiveRSL____ExecutorState__i_Compile.ExecutorState {
										return func(_5044_dt__update_happ_h1 uint64) _276_LiveRSL____ExecutorState__i_Compile.ExecutorState {
											return _276_LiveRSL____ExecutorState__i_Compile.ExecutorState{_276_LiveRSL____ExecutorState__i_Compile.ExecutorState_ExecutorState{(_5040_dt__update__tmp_h1).Dtor_constants(), _5044_dt__update_happ_h1, _5043_dt__update_hops__complete_h1, _5042_dt__update_hmax__bal__reflected_h1, _5041_dt__update_hnext__op__to__execute_h1}}
										}(_pat_let110_0)
									}((_5039_cm).Dtor_app__state())
								}(_pat_let109_0)
							}((_5039_cm).Dtor_opn__state__supply())
						}(_pat_let108_0)
					}((_5039_cm).Dtor_bal__state__supply())
				}(_pat_let107_0)
			}(_276_LiveRSL____ExecutorState__i_Compile.COutstandingOperation{_276_LiveRSL____ExecutorState__i_Compile.COutstandingOperation_COutstandingOpUnknown{}})
		}(_pat_let106_0)
	}(cs)
	var _out239 *_9_Native____Io__s_Compile.MutableMap
	var _ = _out239
	_out239 = _9_Native____Io__s_Compile.Companion_MutableMap_.FromMap((_5039_cm).Dtor_reply__cache())
	reply__cache__mutable = _out239
	{
	}
	return cs_k, reply__cache__mutable
}
func (_this *CompanionStruct_Default___) ExecutorProcessAppStateRequest(cs _276_LiveRSL____ExecutorState__i_Compile.ExecutorState, cinp _217_LiveRSL____CMessage__i_Compile.CPacket, reply__cache__mutable *_9_Native____Io__s_Compile.MutableMap) (_276_LiveRSL____ExecutorState__i_Compile.ExecutorState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
	var cs_k _276_LiveRSL____ExecutorState__i_Compile.ExecutorState = _276_LiveRSL____ExecutorState__i_Compile.Type_ExecutorState_().Default().(_276_LiveRSL____ExecutorState__i_Compile.ExecutorState)
	var _ = cs_k
	var cout _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
	var _ = cout
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
	cs_k = cs
	{
	}
	{
	}
	{
	}
	if (((((((cs).Dtor_constants()).Dtor_all()).Dtor_config()).Dtor_replica__ids()).Contains((cinp).Dtor_src())) && (_214_LiveRSL____CTypes__i_Compile.Companion_Default___.CBallotIsNotGreaterThan((cs).Dtor_max__bal__reflected(), ((cinp).Dtor_msg()).Dtor_bal__state__req()))) && ((((cs).Dtor_ops__complete()).Dtor_n()) >= ((((cinp).Dtor_msg()).Dtor_opn__state__req()).Dtor_n())) {
		{
		}
		var _5045_cme _9_Native____Io__s_Compile.EndPoint
		var _ = _5045_cme
		_5045_cme = (((((cs).Dtor_constants()).Dtor_all()).Dtor_config()).Dtor_replica__ids()).IndexUint(((cs).Dtor_constants()).Dtor_my__index()).(_9_Native____Io__s_Compile.EndPoint)
		var _5046_reply__cache _dafny.Map
		var _ = _5046_reply__cache
		_5046_reply__cache = _9_Native____Io__s_Compile.Companion_MutableMap_.MapOf(reply__cache__mutable)
		{
		}
		cout = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_OutboundPacket{_135_Logic____Option__i_Compile.Option{_135_Logic____Option__i_Compile.Option_Some{_217_LiveRSL____CMessage__i_Compile.CPacket{_217_LiveRSL____CMessage__i_Compile.CPacket_CPacket{(cinp).Dtor_src(), _5045_cme, _217_LiveRSL____CMessage__i_Compile.CMessage{_217_LiveRSL____CMessage__i_Compile.CMessage_CMessage__AppStateSupply{(cs).Dtor_max__bal__reflected(), (cs).Dtor_ops__complete(), (cs).Dtor_app(), _5046_reply__cache}}}}}}}}
	} else {
		{
		}
		cout = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_OutboundPacket{_135_Logic____Option__i_Compile.Option{_135_Logic____Option__i_Compile.Option_None{}}}}
	}
	{
	}
	return cs_k, cout
}
func (_this *CompanionStruct_Default___) ExecutorProcessStartingPhase2(cs _276_LiveRSL____ExecutorState__i_Compile.ExecutorState, cinp _217_LiveRSL____CMessage__i_Compile.CPacket) (_276_LiveRSL____ExecutorState__i_Compile.ExecutorState, _217_LiveRSL____CMessage__i_Compile.CBroadcast) {
	var cs_k _276_LiveRSL____ExecutorState__i_Compile.ExecutorState = _276_LiveRSL____ExecutorState__i_Compile.Type_ExecutorState_().Default().(_276_LiveRSL____ExecutorState__i_Compile.ExecutorState)
	var _ = cs_k
	var cout _217_LiveRSL____CMessage__i_Compile.CBroadcast = _217_LiveRSL____CMessage__i_Compile.Type_CBroadcast_().Default().(_217_LiveRSL____CMessage__i_Compile.CBroadcast)
	var _ = cout
	var _5047_start__time uint64
	var _ = _5047_start__time
	var _out240 uint64
	var _ = _out240
	_out240 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_5047_start__time = _out240
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
	var _5048_copn _214_LiveRSL____CTypes__i_Compile.COperationNumber
	var _ = _5048_copn
	_5048_copn = ((cinp).Dtor_msg()).Dtor_logTruncationPoint__2()
	cs_k = cs
	{
	}
	{
	}
	{
	}
	if ((((((cs).Dtor_constants()).Dtor_all()).Dtor_config()).Dtor_replica__ids()).Contains((cinp).Dtor_src())) && (((_5048_copn).Dtor_n()) > (((cs).Dtor_ops__complete()).Dtor_n())) {
		{
		}
		var _5049_cmsg _217_LiveRSL____CMessage__i_Compile.CMessage
		var _ = _5049_cmsg
		_5049_cmsg = _217_LiveRSL____CMessage__i_Compile.CMessage{_217_LiveRSL____CMessage__i_Compile.CMessage_CMessage__AppStateRequest{((cinp).Dtor_msg()).Dtor_bal__2(), _5048_copn}}
		{
		}
		var _out241 _217_LiveRSL____CMessage__i_Compile.CBroadcast
		var _ = _out241
		_out241 = _301_Impl____LiveRSL____Broadcast__i_Compile.Companion_Default___.BuildBroadcastToEveryone((((cs).Dtor_constants()).Dtor_all()).Dtor_config(), ((cs).Dtor_constants()).Dtor_my__index(), _5049_cmsg)
		cout = _out241
		var _5050_end__time uint64
		var _ = _5050_end__time
		var _out242 uint64
		var _ = _out242
		_out242 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
		_5050_end__time = _out242
		_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("ExecutorProcessStartingPhase2_request"), _5047_start__time, _5050_end__time)
	} else {
		{
		}
		cout = _217_LiveRSL____CMessage__i_Compile.CBroadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast_CBroadcastNop{}}
		var _5051_end__time uint64
		var _ = _5051_end__time
		var _out243 uint64
		var _ = _out243
		_out243 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
		_5051_end__time = _out243
		_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("ExecutorProcessStartingPhase2_nada"), _5047_start__time, _5051_end__time)
	}
	return cs_k, cout
}
func (_this *CompanionStruct_Default___) ExecutorProcessRequest(cs _276_LiveRSL____ExecutorState__i_Compile.ExecutorState, cinp _217_LiveRSL____CMessage__i_Compile.CPacket, reply__cache__mutable *_9_Native____Io__s_Compile.MutableMap) _217_LiveRSL____CMessage__i_Compile.OutboundPackets {
	var cout _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
	var _ = cout
	{
	}
	{
	}
	{
	}
	{
	}
	var _5052_contains bool
	var _ = _5052_contains
	var _5053_cachedReply _214_LiveRSL____CTypes__i_Compile.CReply
	var _ = _5053_cachedReply
	var _out244 bool
	var _ = _out244
	var _out245 interface{}
	var _ = _out245
	_out244, _out245 = (reply__cache__mutable).TryGetValue((cinp).Dtor_src())
	_5052_contains = _out244
	_5053_cachedReply = _out245.(_214_LiveRSL____CTypes__i_Compile.CReply)
	{
	}
	if (((cinp).Dtor_msg()).Dtor_seqno()) == ((_5053_cachedReply).Dtor_seqno()) {
		var _5054_cr _214_LiveRSL____CTypes__i_Compile.CReply
		var _ = _5054_cr
		_5054_cr = _5053_cachedReply
		var _5055_msg _217_LiveRSL____CMessage__i_Compile.CMessage
		var _ = _5055_msg
		_5055_msg = _217_LiveRSL____CMessage__i_Compile.CMessage{_217_LiveRSL____CMessage__i_Compile.CMessage_CMessage__Reply{(_5054_cr).Dtor_seqno(), (_5054_cr).Dtor_reply()}}
		cout = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_OutboundPacket{_135_Logic____Option__i_Compile.Option{_135_Logic____Option__i_Compile.Option_Some{_217_LiveRSL____CMessage__i_Compile.CPacket{_217_LiveRSL____CMessage__i_Compile.CPacket_CPacket{(_5054_cr).Dtor_client(), (((((cs).Dtor_constants()).Dtor_all()).Dtor_config()).Dtor_replica__ids()).IndexUint(((cs).Dtor_constants()).Dtor_my__index()).(_9_Native____Io__s_Compile.EndPoint), _5055_msg}}}}}}
		{
		}
	} else {
		cout = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_OutboundPacket{_135_Logic____Option__i_Compile.Option{_135_Logic____Option__i_Compile.Option_None{}}}}
	}
	return cout
}

// End of class Default__
