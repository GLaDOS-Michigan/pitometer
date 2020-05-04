// Package _312_LiveRSL____AcceptorModel__i_Compile
// Dafny module _312_LiveRSL____AcceptorModel__i_Compile compiled into Go

package _312_LiveRSL____AcceptorModel__i_Compile

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
	return "_312_LiveRSL____AcceptorModel__i_Compile.Default__"
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
	return "_312_LiveRSL____AcceptorModel__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) CreateSeq(len_ _dafny.Int, init__int uint64) _dafny.Seq {
	var s _dafny.Seq = _dafny.EmptySeq
	var _ = s
	if (len_).Cmp(_dafny.Zero) == 0 {
		s = _dafny.SeqOf()
	} else {
		var _4940_rest _dafny.Seq
		var _ = _4940_rest
		var _out213 _dafny.Seq
		var _ = _out213
		_out213 = Companion_Default___.CreateSeq((len_).Minus(_dafny.IntOfInt64(1)), init__int)
		_4940_rest = _out213
		s = (_dafny.SeqOf(_214_LiveRSL____CTypes__i_Compile.COperationNumber{_214_LiveRSL____CTypes__i_Compile.COperationNumber_COperationNumber{init__int}})).Concat(_4940_rest)
	}
	return s
}
func (_this *CompanionStruct_Default___) DummyInitLastCheckpointedOperation(config _238_LiveRSL____CPaxosConfiguration__i_Compile.CPaxosConfiguration) _dafny.Seq {
	var ilco _dafny.Seq = _dafny.EmptySeq
	var _ = ilco
	var _out214 _dafny.Seq
	var _ = _out214
	_out214 = Companion_Default___.CreateSeq(((config).Dtor_replica__ids()).Cardinality(), uint64(0))
	ilco = _out214
	{
	}
	{
	}
	return ilco
}
func (_this *CompanionStruct_Default___) InitAcceptorState(rcs _245_LiveRSL____ReplicaConstantsState__i_Compile.ReplicaConstantsState) _269_LiveRSL____AcceptorState__i_Compile.AcceptorState {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var acceptor _269_LiveRSL____AcceptorState__i_Compile.AcceptorState = _269_LiveRSL____AcceptorState__i_Compile.Type_AcceptorState_().Default().(_269_LiveRSL____AcceptorState__i_Compile.AcceptorState)
	var _ = acceptor
	{
	}
	var _4941_max__ballot _214_LiveRSL____CTypes__i_Compile.CBallot
	var _ = _4941_max__ballot
	_4941_max__ballot = _214_LiveRSL____CTypes__i_Compile.CBallot{_214_LiveRSL____CTypes__i_Compile.CBallot_CBallot{uint64(0), uint64(0)}}
	var _4942_votes _214_LiveRSL____CTypes__i_Compile.CVotes
	var _ = _4942_votes
	_4942_votes = _214_LiveRSL____CTypes__i_Compile.CVotes{_214_LiveRSL____CTypes__i_Compile.CVotes_CVotes{_dafny.NewMapBuilder().ToMap()}}
	var _4943_last__checkpointed__operation _dafny.Seq
	var _ = _4943_last__checkpointed__operation
	var _out215 _dafny.Seq
	var _ = _out215
	_out215 = Companion_Default___.DummyInitLastCheckpointedOperation(((rcs).Dtor_all()).Dtor_config())
	_4943_last__checkpointed__operation = _out215
	var _4944_log__truncation__point _214_LiveRSL____CTypes__i_Compile.COperationNumber
	var _ = _4944_log__truncation__point
	_4944_log__truncation__point = _214_LiveRSL____CTypes__i_Compile.COperationNumber{_214_LiveRSL____CTypes__i_Compile.COperationNumber_COperationNumber{uint64(0)}}
	var _4945_min__voted__opn _214_LiveRSL____CTypes__i_Compile.COperationNumber
	var _ = _4945_min__voted__opn
	_4945_min__voted__opn = _214_LiveRSL____CTypes__i_Compile.COperationNumber{_214_LiveRSL____CTypes__i_Compile.COperationNumber_COperationNumber{uint64(0)}}
	acceptor = _269_LiveRSL____AcceptorState__i_Compile.AcceptorState{_269_LiveRSL____AcceptorState__i_Compile.AcceptorState_AcceptorState{rcs, _4941_max__ballot, _4942_votes, _4943_last__checkpointed__operation, _4944_log__truncation__point, _4945_min__voted__opn}}
	{
	}
	return acceptor
}
func (_this *CompanionStruct_Default___) NextAcceptorState__Phase1(acceptor _269_LiveRSL____AcceptorState__i_Compile.AcceptorState, in__msg _217_LiveRSL____CMessage__i_Compile.CMessage, sender _9_Native____Io__s_Compile.EndPoint) (_269_LiveRSL____AcceptorState__i_Compile.AcceptorState, _217_LiveRSL____CMessage__i_Compile.CBroadcast) {
	var acceptor_k _269_LiveRSL____AcceptorState__i_Compile.AcceptorState = _269_LiveRSL____AcceptorState__i_Compile.Type_AcceptorState_().Default().(_269_LiveRSL____AcceptorState__i_Compile.AcceptorState)
	var _ = acceptor_k
	var packets__sent _217_LiveRSL____CMessage__i_Compile.CBroadcast = _217_LiveRSL____CMessage__i_Compile.Type_CBroadcast_().Default().(_217_LiveRSL____CMessage__i_Compile.CBroadcast)
	var _ = packets__sent
	packets__sent = _217_LiveRSL____CMessage__i_Compile.CBroadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast_CBroadcastNop{}}
	var _4946_ballot _214_LiveRSL____CTypes__i_Compile.CBallot
	var _ = _4946_ballot
	_4946_ballot = (in__msg).Dtor_bal__1a()
	{
	}
	if !(_214_LiveRSL____CTypes__i_Compile.Companion_Default___.CBallotIsLessThan((acceptor).Dtor_maxBallot(), _4946_ballot)) {
		acceptor_k = acceptor
		return acceptor_k, packets__sent
	}
	if !(((((acceptor).Dtor_constants()).Dtor_all()).Dtor_config()).Dtor_replica__ids()).Contains(sender) {
		acceptor_k = acceptor
		return acceptor_k, packets__sent
	}
	var _4947_outMsg _217_LiveRSL____CMessage__i_Compile.CMessage
	var _ = _4947_outMsg
	_4947_outMsg = _217_LiveRSL____CMessage__i_Compile.CMessage{_217_LiveRSL____CMessage__i_Compile.CMessage_CMessage__1b{_4946_ballot, (acceptor).Dtor_log__truncation__point(), (acceptor).Dtor_votes()}}
	packets__sent = _217_LiveRSL____CMessage__i_Compile.CBroadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast_CBroadcast{(((((acceptor).Dtor_constants()).Dtor_all()).Dtor_config()).Dtor_replica__ids()).IndexUint(((acceptor).Dtor_constants()).Dtor_my__index()).(_9_Native____Io__s_Compile.EndPoint), _dafny.SeqOf(sender), _4947_outMsg}}
	acceptor_k = func(_pat_let72_0 _269_LiveRSL____AcceptorState__i_Compile.AcceptorState) _269_LiveRSL____AcceptorState__i_Compile.AcceptorState {
		return func(_4948_dt__update__tmp_h0 _269_LiveRSL____AcceptorState__i_Compile.AcceptorState) _269_LiveRSL____AcceptorState__i_Compile.AcceptorState {
			return func(_pat_let73_0 _214_LiveRSL____CTypes__i_Compile.CBallot) _269_LiveRSL____AcceptorState__i_Compile.AcceptorState {
				return func(_4949_dt__update_hmaxBallot_h0 _214_LiveRSL____CTypes__i_Compile.CBallot) _269_LiveRSL____AcceptorState__i_Compile.AcceptorState {
					return _269_LiveRSL____AcceptorState__i_Compile.AcceptorState{_269_LiveRSL____AcceptorState__i_Compile.AcceptorState_AcceptorState{(_4948_dt__update__tmp_h0).Dtor_constants(), _4949_dt__update_hmaxBallot_h0, (_4948_dt__update__tmp_h0).Dtor_votes(), (_4948_dt__update__tmp_h0).Dtor_last__checkpointed__operation(), (_4948_dt__update__tmp_h0).Dtor_log__truncation__point(), (_4948_dt__update__tmp_h0).Dtor_minVotedOpn()}}
				}(_pat_let73_0)
			}(_4946_ballot)
		}(_pat_let72_0)
	}(acceptor)
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
	return acceptor_k, packets__sent
}
func (_this *CompanionStruct_Default___) AddVoteAndRemoveOldOnesImpl(acceptor _269_LiveRSL____AcceptorState__i_Compile.AcceptorState, votes _214_LiveRSL____CTypes__i_Compile.CVotes, new__opn _214_LiveRSL____CTypes__i_Compile.COperationNumber, new__vote _214_LiveRSL____CTypes__i_Compile.CVote, newLogTruncationPoint _214_LiveRSL____CTypes__i_Compile.COperationNumber, minVotedOpn _214_LiveRSL____CTypes__i_Compile.COperationNumber) (_214_LiveRSL____CTypes__i_Compile.CVotes, _214_LiveRSL____CTypes__i_Compile.COperationNumber) {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var votes_k _214_LiveRSL____CTypes__i_Compile.CVotes = _214_LiveRSL____CTypes__i_Compile.Type_CVotes_().Default().(_214_LiveRSL____CTypes__i_Compile.CVotes)
	var _ = votes_k
	var newMinVotedOpn _214_LiveRSL____CTypes__i_Compile.COperationNumber = _214_LiveRSL____CTypes__i_Compile.Type_COperationNumber_().Default().(_214_LiveRSL____CTypes__i_Compile.COperationNumber)
	var _ = newMinVotedOpn
	{
	}
	var _4950_updated__votes _dafny.Map
	var _ = _4950_updated__votes
	_4950_updated__votes = ((votes).Dtor_v()).Update(new__opn, new__vote)
	var _4951_new__votes _dafny.Map = _dafny.EmptyMap
	var _ = _4951_new__votes
	if ((newLogTruncationPoint).Dtor_n()) > ((minVotedOpn).Dtor_n()) {
		_4951_new__votes = func() _dafny.Map {
			var _coll3 = _dafny.NewMapBuilder()
			var _ = _coll3
			for _iter15 := _dafny.Iterate((_4950_updated__votes).Keys().Elements()); ; {
				_val15, _ok15 := _iter15()
				if !_ok15 {
					break
				}
				_4952_op := _val15.(_214_LiveRSL____CTypes__i_Compile.COperationNumber)
				if ((_4950_updated__votes).Contains(_4952_op)) && (((_4952_op).Dtor_n()) >= ((newLogTruncationPoint).Dtor_n())) {
					_coll3.Add(_4952_op, (_4950_updated__votes).Get(_4952_op).(_214_LiveRSL____CTypes__i_Compile.CVote))
				}
			}
			return _coll3.ToMap()
		}()
		newMinVotedOpn = newLogTruncationPoint
	} else {
		_4951_new__votes = _4950_updated__votes
		if ((new__opn).Dtor_n()) < ((minVotedOpn).Dtor_n()) {
			newMinVotedOpn = new__opn
		} else {
			newMinVotedOpn = minVotedOpn
		}
	}
	votes_k = _214_LiveRSL____CTypes__i_Compile.CVotes{_214_LiveRSL____CTypes__i_Compile.CVotes_CVotes{_4951_new__votes}}
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
	return votes_k, newMinVotedOpn
}
func (_this *CompanionStruct_Default___) NextAcceptorState__Phase2(acceptor _269_LiveRSL____AcceptorState__i_Compile.AcceptorState, in__msg _217_LiveRSL____CMessage__i_Compile.CMessage, sender _9_Native____Io__s_Compile.EndPoint) (_269_LiveRSL____AcceptorState__i_Compile.AcceptorState, _217_LiveRSL____CMessage__i_Compile.CBroadcast) {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var acceptor_k _269_LiveRSL____AcceptorState__i_Compile.AcceptorState = _269_LiveRSL____AcceptorState__i_Compile.Type_AcceptorState_().Default().(_269_LiveRSL____AcceptorState__i_Compile.AcceptorState)
	var _ = acceptor_k
	var packets__sent _217_LiveRSL____CMessage__i_Compile.CBroadcast = _217_LiveRSL____CMessage__i_Compile.Type_CBroadcast_().Default().(_217_LiveRSL____CMessage__i_Compile.CBroadcast)
	var _ = packets__sent
	var _4953_start__time uint64
	var _ = _4953_start__time
	var _out216 uint64
	var _ = _out216
	_out216 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_4953_start__time = _out216
	{
	}
	packets__sent = _217_LiveRSL____CMessage__i_Compile.CBroadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast_CBroadcastNop{}}
	acceptor_k = acceptor
	var _4954_ballot _214_LiveRSL____CTypes__i_Compile.CBallot
	var _ = _4954_ballot
	_4954_ballot = (in__msg).Dtor_bal__2a()
	{
	}
	var _4955_maxLogLengthMinus1 uint64
	var _ = _4955_maxLogLengthMinus1
	_4955_maxLogLengthMinus1 = (((((acceptor).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__log__length()) - (func() uint64 { return (uint64(1)) })()
	var _4956_newLogTruncationPoint _214_LiveRSL____CTypes__i_Compile.COperationNumber
	var _ = _4956_newLogTruncationPoint
	_4956_newLogTruncationPoint = (acceptor).Dtor_log__truncation__point()
	if (((in__msg).Dtor_opn__2a()).Dtor_n()) >= (_4955_maxLogLengthMinus1) {
		var _4957_potentialNewTruncationPoint uint64
		var _ = _4957_potentialNewTruncationPoint
		_4957_potentialNewTruncationPoint = (((in__msg).Dtor_opn__2a()).Dtor_n()) - (func() uint64 { return (_4955_maxLogLengthMinus1) })()
		if (_4957_potentialNewTruncationPoint) > (((acceptor).Dtor_log__truncation__point()).Dtor_n()) {
			_4956_newLogTruncationPoint = _214_LiveRSL____CTypes__i_Compile.COperationNumber{_214_LiveRSL____CTypes__i_Compile.COperationNumber_COperationNumber{_4957_potentialNewTruncationPoint}}
		}
	}
	{
	}
	var _4958_addition uint64
	var _ = _4958_addition
	var _out217 uint64
	var _ = _out217
	_out217 = _234_Common____UpperBound__i_Compile.Companion_Default___.UpperBoundedAdditionImpl(((acceptor).Dtor_log__truncation__point()).Dtor_n(), ((((acceptor).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__log__length(), ((((acceptor).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__integer__val())
	_4958_addition = _out217
	{
	}
	var _4959_opn _214_LiveRSL____CTypes__i_Compile.COperationNumber
	var _ = _4959_opn
	_4959_opn = (in__msg).Dtor_opn__2a()
	var _4960_value _dafny.Seq
	var _ = _4960_value
	_4960_value = (in__msg).Dtor_val__2a()
	var _4961_newVote _214_LiveRSL____CTypes__i_Compile.CVote
	var _ = _4961_newVote
	_4961_newVote = _214_LiveRSL____CTypes__i_Compile.CVote{_214_LiveRSL____CTypes__i_Compile.CVote_CVote{_4954_ballot, _4960_value}}
	var _4962_votes_k _214_LiveRSL____CTypes__i_Compile.CVotes = _214_LiveRSL____CTypes__i_Compile.Type_CVotes_().Default().(_214_LiveRSL____CTypes__i_Compile.CVotes)
	var _ = _4962_votes_k
	var _4963_newMinVotedOpn _214_LiveRSL____CTypes__i_Compile.COperationNumber = _214_LiveRSL____CTypes__i_Compile.Type_COperationNumber_().Default().(_214_LiveRSL____CTypes__i_Compile.COperationNumber)
	var _ = _4963_newMinVotedOpn
	if (((acceptor).Dtor_log__truncation__point()).Dtor_n()) <= (((in__msg).Dtor_opn__2a()).Dtor_n()) {
		var _out218 _214_LiveRSL____CTypes__i_Compile.CVotes
		var _ = _out218
		var _out219 _214_LiveRSL____CTypes__i_Compile.COperationNumber
		var _ = _out219
		_out218, _out219 = Companion_Default___.AddVoteAndRemoveOldOnesImpl(acceptor, (acceptor).Dtor_votes(), _4959_opn, _4961_newVote, _4956_newLogTruncationPoint, (acceptor).Dtor_minVotedOpn())
		_4962_votes_k = _out218
		_4963_newMinVotedOpn = _out219
	} else {
		_4962_votes_k = (acceptor).Dtor_votes()
		_4963_newMinVotedOpn = (acceptor).Dtor_minVotedOpn()
	}
	var _4964_outMsg _217_LiveRSL____CMessage__i_Compile.CMessage
	var _ = _4964_outMsg
	_4964_outMsg = _217_LiveRSL____CMessage__i_Compile.CMessage{_217_LiveRSL____CMessage__i_Compile.CMessage_CMessage__2b{_4954_ballot, _4959_opn, _4960_value}}
	acceptor_k = func(_pat_let74_0 _269_LiveRSL____AcceptorState__i_Compile.AcceptorState) _269_LiveRSL____AcceptorState__i_Compile.AcceptorState {
		return func(_4965_dt__update__tmp_h0 _269_LiveRSL____AcceptorState__i_Compile.AcceptorState) _269_LiveRSL____AcceptorState__i_Compile.AcceptorState {
			return func(_pat_let75_0 _214_LiveRSL____CTypes__i_Compile.COperationNumber) _269_LiveRSL____AcceptorState__i_Compile.AcceptorState {
				return func(_4966_dt__update_hminVotedOpn_h0 _214_LiveRSL____CTypes__i_Compile.COperationNumber) _269_LiveRSL____AcceptorState__i_Compile.AcceptorState {
					return func(_pat_let76_0 _214_LiveRSL____CTypes__i_Compile.COperationNumber) _269_LiveRSL____AcceptorState__i_Compile.AcceptorState {
						return func(_4967_dt__update_hlog__truncation__point_h0 _214_LiveRSL____CTypes__i_Compile.COperationNumber) _269_LiveRSL____AcceptorState__i_Compile.AcceptorState {
							return func(_pat_let77_0 _214_LiveRSL____CTypes__i_Compile.CBallot) _269_LiveRSL____AcceptorState__i_Compile.AcceptorState {
								return func(_4968_dt__update_hmaxBallot_h0 _214_LiveRSL____CTypes__i_Compile.CBallot) _269_LiveRSL____AcceptorState__i_Compile.AcceptorState {
									return func(_pat_let78_0 _214_LiveRSL____CTypes__i_Compile.CVotes) _269_LiveRSL____AcceptorState__i_Compile.AcceptorState {
										return func(_4969_dt__update_hvotes_h0 _214_LiveRSL____CTypes__i_Compile.CVotes) _269_LiveRSL____AcceptorState__i_Compile.AcceptorState {
											return _269_LiveRSL____AcceptorState__i_Compile.AcceptorState{_269_LiveRSL____AcceptorState__i_Compile.AcceptorState_AcceptorState{(_4965_dt__update__tmp_h0).Dtor_constants(), _4968_dt__update_hmaxBallot_h0, _4969_dt__update_hvotes_h0, (_4965_dt__update__tmp_h0).Dtor_last__checkpointed__operation(), _4967_dt__update_hlog__truncation__point_h0, _4966_dt__update_hminVotedOpn_h0}}
										}(_pat_let78_0)
									}(_4962_votes_k)
								}(_pat_let77_0)
							}(_4954_ballot)
						}(_pat_let76_0)
					}(_4956_newLogTruncationPoint)
				}(_pat_let75_0)
			}(_4963_newMinVotedOpn)
		}(_pat_let74_0)
	}(acceptor)
	var _out220 _217_LiveRSL____CMessage__i_Compile.CBroadcast
	var _ = _out220
	_out220 = _301_Impl____LiveRSL____Broadcast__i_Compile.Companion_Default___.BuildBroadcastToEveryone((((acceptor).Dtor_constants()).Dtor_all()).Dtor_config(), ((acceptor).Dtor_constants()).Dtor_my__index(), _4964_outMsg)
	packets__sent = _out220
	{
	}
	var _4970_end__time uint64
	var _ = _4970_end__time
	var _out221 uint64
	var _ = _out221
	_out221 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_4970_end__time = _out221
	_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("NextAcceptorState_Phase2"), _4953_start__time, _4970_end__time)
	return acceptor_k, packets__sent
}
func (_this *CompanionStruct_Default___) NextAcceptorState__ProcessHeartbeat(acceptor _269_LiveRSL____AcceptorState__i_Compile.AcceptorState, in__msg _217_LiveRSL____CMessage__i_Compile.CMessage, sender _9_Native____Io__s_Compile.EndPoint) _269_LiveRSL____AcceptorState__i_Compile.AcceptorState {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var acceptor_k _269_LiveRSL____AcceptorState__i_Compile.AcceptorState = _269_LiveRSL____AcceptorState__i_Compile.Type_AcceptorState_().Default().(_269_LiveRSL____AcceptorState__i_Compile.AcceptorState)
	var _ = acceptor_k
	acceptor_k = acceptor
	{
	}
	{
	}
	{
	}
	var _4971_found bool
	var _ = _4971_found
	var _4972_index uint64
	var _ = _4972_index
	var _out222 bool
	var _ = _out222
	var _out223 uint64
	var _ = _out223
	_out222, _out223 = _238_LiveRSL____CPaxosConfiguration__i_Compile.Companion_Default___.CGetReplicaIndex(sender, (((acceptor).Dtor_constants()).Dtor_all()).Dtor_config())
	_4971_found = _out222
	_4972_index = _out223
	if !(_4971_found) {
		return acceptor_k
	}
	{
	}
	if (((in__msg).Dtor_opn__ckpt()).Dtor_n()) <= ((((acceptor).Dtor_last__checkpointed__operation()).IndexUint(_4972_index).(_214_LiveRSL____CTypes__i_Compile.COperationNumber)).Dtor_n()) {
		return acceptor_k
	}
	var _4973_LCO _dafny.Seq
	var _ = _4973_LCO
	_4973_LCO = (acceptor).Dtor_last__checkpointed__operation()
	var _4974_newLCO _dafny.Seq
	var _ = _4974_newLCO
	_4974_newLCO = (((acceptor).Dtor_last__checkpointed__operation()).Update(_dafny.IntOfUint64(_4972_index), (in__msg).Dtor_opn__ckpt())).(_dafny.Seq)
	acceptor_k = func(_pat_let79_0 _269_LiveRSL____AcceptorState__i_Compile.AcceptorState) _269_LiveRSL____AcceptorState__i_Compile.AcceptorState {
		return func(_4975_dt__update__tmp_h0 _269_LiveRSL____AcceptorState__i_Compile.AcceptorState) _269_LiveRSL____AcceptorState__i_Compile.AcceptorState {
			return func(_pat_let80_0 _dafny.Seq) _269_LiveRSL____AcceptorState__i_Compile.AcceptorState {
				return func(_4976_dt__update_hlast__checkpointed__operation_h0 _dafny.Seq) _269_LiveRSL____AcceptorState__i_Compile.AcceptorState {
					return _269_LiveRSL____AcceptorState__i_Compile.AcceptorState{_269_LiveRSL____AcceptorState__i_Compile.AcceptorState_AcceptorState{(_4975_dt__update__tmp_h0).Dtor_constants(), (_4975_dt__update__tmp_h0).Dtor_maxBallot(), (_4975_dt__update__tmp_h0).Dtor_votes(), _4976_dt__update_hlast__checkpointed__operation_h0, (_4975_dt__update__tmp_h0).Dtor_log__truncation__point(), (_4975_dt__update__tmp_h0).Dtor_minVotedOpn()}}
				}(_pat_let80_0)
			}(_4974_newLCO)
		}(_pat_let79_0)
	}(acceptor)
	return acceptor_k
}
func (_this *CompanionStruct_Default___) NextAcceptorState__TruncateLog(acceptor _269_LiveRSL____AcceptorState__i_Compile.AcceptorState, opn _214_LiveRSL____CTypes__i_Compile.COperationNumber) _269_LiveRSL____AcceptorState__i_Compile.AcceptorState {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var acceptor_k _269_LiveRSL____AcceptorState__i_Compile.AcceptorState = _269_LiveRSL____AcceptorState__i_Compile.Type_AcceptorState_().Default().(_269_LiveRSL____AcceptorState__i_Compile.AcceptorState)
	var _ = acceptor_k
	if ((opn).Dtor_n()) > (((acceptor).Dtor_log__truncation__point()).Dtor_n()) {
		{
		}
		{
		}
		var _4977_truncatedVotes _dafny.Map
		var _ = _4977_truncatedVotes
		_4977_truncatedVotes = func() _dafny.Map {
			var _coll4 = _dafny.NewMapBuilder()
			var _ = _coll4
			for _iter16 := _dafny.Iterate((((acceptor).Dtor_votes()).Dtor_v()).Keys().Elements()); ; {
				_val16, _ok16 := _iter16()
				if !_ok16 {
					break
				}
				_4978_op := _val16.(_214_LiveRSL____CTypes__i_Compile.COperationNumber)
				if ((((acceptor).Dtor_votes()).Dtor_v()).Contains(_4978_op)) && (((_4978_op).Dtor_n()) >= ((opn).Dtor_n())) {
					_coll4.Add(_4978_op, (((acceptor).Dtor_votes()).Dtor_v()).Get(_4978_op).(_214_LiveRSL____CTypes__i_Compile.CVote))
				}
			}
			return _coll4.ToMap()
		}()
		acceptor_k = func(_pat_let81_0 _269_LiveRSL____AcceptorState__i_Compile.AcceptorState) _269_LiveRSL____AcceptorState__i_Compile.AcceptorState {
			return func(_4979_dt__update__tmp_h0 _269_LiveRSL____AcceptorState__i_Compile.AcceptorState) _269_LiveRSL____AcceptorState__i_Compile.AcceptorState {
				return func(_pat_let82_0 _214_LiveRSL____CTypes__i_Compile.CVotes) _269_LiveRSL____AcceptorState__i_Compile.AcceptorState {
					return func(_4982_dt__update_hvotes_h0 _214_LiveRSL____CTypes__i_Compile.CVotes) _269_LiveRSL____AcceptorState__i_Compile.AcceptorState {
						return func(_pat_let85_0 _214_LiveRSL____CTypes__i_Compile.COperationNumber) _269_LiveRSL____AcceptorState__i_Compile.AcceptorState {
							return func(_4983_dt__update_hlog__truncation__point_h0 _214_LiveRSL____CTypes__i_Compile.COperationNumber) _269_LiveRSL____AcceptorState__i_Compile.AcceptorState {
								return _269_LiveRSL____AcceptorState__i_Compile.AcceptorState{_269_LiveRSL____AcceptorState__i_Compile.AcceptorState_AcceptorState{(_4979_dt__update__tmp_h0).Dtor_constants(), (_4979_dt__update__tmp_h0).Dtor_maxBallot(), _4982_dt__update_hvotes_h0, (_4979_dt__update__tmp_h0).Dtor_last__checkpointed__operation(), _4983_dt__update_hlog__truncation__point_h0, (_4979_dt__update__tmp_h0).Dtor_minVotedOpn()}}
							}(_pat_let85_0)
						}(opn)
					}(_pat_let82_0)
				}(func(_pat_let83_0 _214_LiveRSL____CTypes__i_Compile.CVotes) _214_LiveRSL____CTypes__i_Compile.CVotes {
					return func(_4980_dt__update__tmp_h1 _214_LiveRSL____CTypes__i_Compile.CVotes) _214_LiveRSL____CTypes__i_Compile.CVotes {
						return func(_pat_let84_0 _dafny.Map) _214_LiveRSL____CTypes__i_Compile.CVotes {
							return func(_4981_dt__update_hv_h0 _dafny.Map) _214_LiveRSL____CTypes__i_Compile.CVotes {
								return _214_LiveRSL____CTypes__i_Compile.CVotes{_214_LiveRSL____CTypes__i_Compile.CVotes_CVotes{_4981_dt__update_hv_h0}}
							}(_pat_let84_0)
						}(_4977_truncatedVotes)
					}(_pat_let83_0)
				}((acceptor).Dtor_votes()))
			}(_pat_let81_0)
		}(acceptor)
		{
		}
		{
		}
		{
		}
	} else {
		acceptor_k = acceptor
	}
	return acceptor_k
}
func (_this *CompanionStruct_Default___) AcceptorModel__GetNthHighestValueAmongReportedCheckpoints(acceptor _269_LiveRSL____AcceptorState__i_Compile.AcceptorState, minQuorumSize uint64) _214_LiveRSL____CTypes__i_Compile.COperationNumber {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var opn _214_LiveRSL____CTypes__i_Compile.COperationNumber = _214_LiveRSL____CTypes__i_Compile.Type_COperationNumber_().Default().(_214_LiveRSL____CTypes__i_Compile.COperationNumber)
	var _ = opn
	var _out224 _214_LiveRSL____CTypes__i_Compile.COperationNumber
	var _ = _out224
	_out224 = _267_LiveRSL____CLastCheckpointedMap__i_Compile.Companion_Default___.ComputeNthHighestValue((acceptor).Dtor_last__checkpointed__operation(), minQuorumSize)
	opn = _out224
	return opn
}

// End of class Default__
