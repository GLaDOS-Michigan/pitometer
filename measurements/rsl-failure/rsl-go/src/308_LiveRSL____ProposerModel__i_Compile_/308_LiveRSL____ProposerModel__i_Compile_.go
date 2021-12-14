// Package _308_LiveRSL____ProposerModel__i_Compile
// Dafny module _308_LiveRSL____ProposerModel__i_Compile compiled into Go

package _308_LiveRSL____ProposerModel__i_Compile

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
	"fmt"
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
	return "_308_LiveRSL____ProposerModel__i_Compile.Default__"
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
	return "_308_LiveRSL____ProposerModel__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) InitProposerState(constants _245_LiveRSL____ReplicaConstantsState__i_Compile.ReplicaConstantsState) (_254_LiveRSL____ProposerState__i_Compile.ProposerState, *_9_Native____Io__s_Compile.MutableSet, *_9_Native____Io__s_Compile.MutableSet) {
	var proposer _254_LiveRSL____ProposerState__i_Compile.ProposerState = _254_LiveRSL____ProposerState__i_Compile.Type_ProposerState_().Default().(_254_LiveRSL____ProposerState__i_Compile.ProposerState)
	var _ = proposer
	var cur__req__set *_9_Native____Io__s_Compile.MutableSet = (*_9_Native____Io__s_Compile.MutableSet)(nil)
	var _ = cur__req__set
	var prev__req__set *_9_Native____Io__s_Compile.MutableSet = (*_9_Native____Io__s_Compile.MutableSet)(nil)
	var _ = prev__req__set
	var _4822_election _251_LiveRSL____ElectionState__i_Compile.CElectionState = _251_LiveRSL____ElectionState__i_Compile.Type_CElectionState_().Default().(_251_LiveRSL____ElectionState__i_Compile.CElectionState)
	var _ = _4822_election
	var _out163 _251_LiveRSL____ElectionState__i_Compile.CElectionState
	var _ = _out163
	var _out164 *_9_Native____Io__s_Compile.MutableSet
	var _ = _out164
	var _out165 *_9_Native____Io__s_Compile.MutableSet
	var _ = _out165
	_out163, _out164, _out165 = _297_LiveRSL____ElectionModel__i_Compile.Companion_Default___.InitElectionState(constants)
	_4822_election = _out163
	cur__req__set = _out164
	prev__req__set = _out165
	proposer = _254_LiveRSL____ProposerState__i_Compile.ProposerState{_254_LiveRSL____ProposerState__i_Compile.ProposerState_ProposerState{constants, uint8(0), _dafny.SeqOf(), _214_LiveRSL____CTypes__i_Compile.CBallot{_214_LiveRSL____CTypes__i_Compile.CBallot_CBallot{uint64(0), (constants).Dtor_my__index()}}, uint64(0), _dafny.SetOf(), _dafny.NewMapBuilder().ToMap(), _254_LiveRSL____ProposerState__i_Compile.CIncompleteBatchTimer{_254_LiveRSL____ProposerState__i_Compile.CIncompleteBatchTimer_CIncompleteBatchTimerOff{}}, _4822_election, _214_LiveRSL____CTypes__i_Compile.COperationNumber{_214_LiveRSL____CTypes__i_Compile.COperationNumber_COperationNumber{uint64(0)}}, _214_LiveRSL____CTypes__i_Compile.COperationNumber{_214_LiveRSL____CTypes__i_Compile.COperationNumber_COperationNumber{uint64(0)}}}}
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
	return proposer, cur__req__set, prev__req__set
}
func (_this *CompanionStruct_Default___) ProposerProcessRequest(proposer _254_LiveRSL____ProposerState__i_Compile.ProposerState, packet _217_LiveRSL____CMessage__i_Compile.CPacket, cur__req__set *_9_Native____Io__s_Compile.MutableSet, prev__req__set *_9_Native____Io__s_Compile.MutableSet) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
	var proposer_k _254_LiveRSL____ProposerState__i_Compile.ProposerState = _254_LiveRSL____ProposerState__i_Compile.Type_ProposerState_().Default().(_254_LiveRSL____ProposerState__i_Compile.ProposerState)
	var _ = proposer_k
	{
	}
	{
	}
	{
	}
	var _4823_val _214_LiveRSL____CTypes__i_Compile.CRequest
	var _ = _4823_val
	_4823_val = _214_LiveRSL____CTypes__i_Compile.CRequest{_214_LiveRSL____CTypes__i_Compile.CRequest_CRequest{(packet).Dtor_src(), ((packet).Dtor_msg()).Dtor_seqno(), ((packet).Dtor_msg()).Dtor_val()}}
	var _4824_newElectionState _251_LiveRSL____ElectionState__i_Compile.CElectionState
	var _ = _4824_newElectionState
	var _out166 _251_LiveRSL____ElectionState__i_Compile.CElectionState
	var _ = _out166
	_out166 = _297_LiveRSL____ElectionModel__i_Compile.Companion_Default___.ElectionReflectReceivedRequest((proposer).Dtor_election__state(), _4823_val, cur__req__set, prev__req__set)
	_4824_newElectionState = _out166
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
	if (((proposer).Dtor_current__state()) != (uint8(0)) /* dircomp */) && ((!((proposer).Dtor_highest__seqno__requested__by__client__this__view()).Contains((packet).Dtor_src())) || ((((packet).Dtor_msg()).Dtor_seqno()) > (((proposer).Dtor_highest__seqno__requested__by__client__this__view()).Get((packet).Dtor_src()).(uint64)))) {
		{
		}
		var _4825_new__seqno__map _dafny.Map
		var _ = _4825_new__seqno__map
		_4825_new__seqno__map = ((proposer).Dtor_highest__seqno__requested__by__client__this__view()).Update((packet).Dtor_src(), ((packet).Dtor_msg()).Dtor_seqno())
		proposer_k = func(_pat_let33_0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
			return func(_4826_dt__update__tmp_h0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
				return func(_pat_let34_0 _dafny.Map) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
					return func(_4827_dt__update_hhighest__seqno__requested__by__client__this__view_h0 _dafny.Map) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
						return func(_pat_let35_0 _dafny.Seq) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
							return func(_4828_dt__update_hrequest__queue_h0 _dafny.Seq) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
								return func(_pat_let36_0 _251_LiveRSL____ElectionState__i_Compile.CElectionState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
									return func(_4829_dt__update_helection__state_h0 _251_LiveRSL____ElectionState__i_Compile.CElectionState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
										return _254_LiveRSL____ProposerState__i_Compile.ProposerState{_254_LiveRSL____ProposerState__i_Compile.ProposerState_ProposerState{(_4826_dt__update__tmp_h0).Dtor_constants(), (_4826_dt__update__tmp_h0).Dtor_current__state(), _4828_dt__update_hrequest__queue_h0, (_4826_dt__update__tmp_h0).Dtor_max__ballot__i__sent__1a(), (_4826_dt__update__tmp_h0).Dtor_next__operation__number__to__propose(), (_4826_dt__update__tmp_h0).Dtor_received__1b__packets(), _4827_dt__update_hhighest__seqno__requested__by__client__this__view_h0, (_4826_dt__update__tmp_h0).Dtor_incomplete__batch__timer(), _4829_dt__update_helection__state_h0, (_4826_dt__update__tmp_h0).Dtor_maxOpnWithProposal(), (_4826_dt__update__tmp_h0).Dtor_maxLogTruncationPoint()}}
									}(_pat_let36_0)
								}(_4824_newElectionState)
							}(_pat_let35_0)
						}(((proposer).Dtor_request__queue()).Concat(_dafny.SeqOf(_4823_val)))
					}(_pat_let34_0)
				}(_4825_new__seqno__map)
			}(_pat_let33_0)
		}(proposer)
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
	} else {
		proposer_k = func(_pat_let37_0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
			return func(_4830_dt__update__tmp_h2 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
				return func(_pat_let38_0 _251_LiveRSL____ElectionState__i_Compile.CElectionState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
					return func(_4831_dt__update_helection__state_h2 _251_LiveRSL____ElectionState__i_Compile.CElectionState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
						return _254_LiveRSL____ProposerState__i_Compile.ProposerState{_254_LiveRSL____ProposerState__i_Compile.ProposerState_ProposerState{(_4830_dt__update__tmp_h2).Dtor_constants(), (_4830_dt__update__tmp_h2).Dtor_current__state(), (_4830_dt__update__tmp_h2).Dtor_request__queue(), (_4830_dt__update__tmp_h2).Dtor_max__ballot__i__sent__1a(), (_4830_dt__update__tmp_h2).Dtor_next__operation__number__to__propose(), (_4830_dt__update__tmp_h2).Dtor_received__1b__packets(), (_4830_dt__update__tmp_h2).Dtor_highest__seqno__requested__by__client__this__view(), (_4830_dt__update__tmp_h2).Dtor_incomplete__batch__timer(), _4831_dt__update_helection__state_h2, (_4830_dt__update__tmp_h2).Dtor_maxOpnWithProposal(), (_4830_dt__update__tmp_h2).Dtor_maxLogTruncationPoint()}}
					}(_pat_let38_0)
				}(_4824_newElectionState)
			}(_pat_let37_0)
		}(proposer)
		{
		}
		{
		}
		{
		}
		{
		}
	}
	return proposer_k
}
func (_this *CompanionStruct_Default___) ProposerMaybeEnterNewViewAndSend1a(proposer _254_LiveRSL____ProposerState__i_Compile.ProposerState) (_254_LiveRSL____ProposerState__i_Compile.ProposerState, _217_LiveRSL____CMessage__i_Compile.CBroadcast, bool) {
	var proposer_k _254_LiveRSL____ProposerState__i_Compile.ProposerState = _254_LiveRSL____ProposerState__i_Compile.Type_ProposerState_().Default().(_254_LiveRSL____ProposerState__i_Compile.ProposerState)
	var _ = proposer_k
	var sent__packets _217_LiveRSL____CMessage__i_Compile.CBroadcast = _217_LiveRSL____CMessage__i_Compile.Type_CBroadcast_().Default().(_217_LiveRSL____CMessage__i_Compile.CBroadcast)
	var _ = sent__packets
	var _4832_start__time uint64
	var _ = _4832_start__time
	var _out167 uint64
	var _ = _out167
	_out167 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_4832_start__time = _out167
	{
	}
	var _4833_lt bool
	var _ = _4833_lt
	var _out168 bool
	var _ = _out168
	_out168 = _214_LiveRSL____CTypes__i_Compile.Companion_Default___.CBalLt((proposer).Dtor_max__ballot__i__sent__1a(), ((proposer).Dtor_election__state()).Dtor_current__view())
	_4833_lt = _out168
	var noop = false
	if (((((proposer).Dtor_election__state()).Dtor_current__view()).Dtor_proposer__id()) == (((proposer).Dtor_constants()).Dtor_my__index())) && (_4833_lt) {
		{
		}
		{
		}
		{
		}
		var _4834_new__requestQueue _dafny.Seq
		var _ = _4834_new__requestQueue
		_4834_new__requestQueue = (((proposer).Dtor_election__state()).Dtor_requests__received__prev__epochs()).Concat(((proposer).Dtor_election__state()).Dtor_requests__received__this__epoch())
		{
		}
		proposer_k = func(_pat_let39_0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
			return func(_4835_dt__update__tmp_h0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
				return func(_pat_let40_0 _dafny.Seq) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
					return func(_4836_dt__update_hrequest__queue_h0 _dafny.Seq) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
						return func(_pat_let41_0 _dafny.Map) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
							return func(_4837_dt__update_hhighest__seqno__requested__by__client__this__view_h0 _dafny.Map) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
								return func(_pat_let42_0 _dafny.Set) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
									return func(_4838_dt__update_hreceived__1b__packets_h0 _dafny.Set) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
										return func(_pat_let43_0 _214_LiveRSL____CTypes__i_Compile.CBallot) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
											return func(_4839_dt__update_hmax__ballot__i__sent__1a_h0 _214_LiveRSL____CTypes__i_Compile.CBallot) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
												return func(_pat_let44_0 uint8) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
													return func(_4840_dt__update_hcurrent__state_h0 uint8) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
														return _254_LiveRSL____ProposerState__i_Compile.ProposerState{_254_LiveRSL____ProposerState__i_Compile.ProposerState_ProposerState{(_4835_dt__update__tmp_h0).Dtor_constants(), _4840_dt__update_hcurrent__state_h0, _4836_dt__update_hrequest__queue_h0, _4839_dt__update_hmax__ballot__i__sent__1a_h0, (_4835_dt__update__tmp_h0).Dtor_next__operation__number__to__propose(), _4838_dt__update_hreceived__1b__packets_h0, _4837_dt__update_hhighest__seqno__requested__by__client__this__view_h0, (_4835_dt__update__tmp_h0).Dtor_incomplete__batch__timer(), (_4835_dt__update__tmp_h0).Dtor_election__state(), (_4835_dt__update__tmp_h0).Dtor_maxOpnWithProposal(), (_4835_dt__update__tmp_h0).Dtor_maxLogTruncationPoint()}}
													}(_pat_let44_0)
												}(uint8(1))
											}(_pat_let43_0)
										}(((proposer).Dtor_election__state()).Dtor_current__view())
									}(_pat_let42_0)
								}(_dafny.SetOf())
							}(_pat_let41_0)
						}(_dafny.NewMapBuilder().ToMap())
					}(_pat_let40_0)
				}(_4834_new__requestQueue)
			}(_pat_let39_0)
		}(proposer)
		{
		}
		var _4841_msg _217_LiveRSL____CMessage__i_Compile.CMessage
		var _ = _4841_msg
		_4841_msg = _217_LiveRSL____CMessage__i_Compile.CMessage{_217_LiveRSL____CMessage__i_Compile.CMessage_CMessage__1a{((proposer).Dtor_election__state()).Dtor_current__view()}}
		{
		}
		{
		}
		var _out169 _217_LiveRSL____CMessage__i_Compile.CBroadcast
		var _ = _out169
		_out169 = _301_Impl____LiveRSL____Broadcast__i_Compile.Companion_Default___.BuildBroadcastToEveryone((((proposer).Dtor_constants()).Dtor_all()).Dtor_config(), ((proposer).Dtor_constants()).Dtor_my__index(), _4841_msg)
		sent__packets = _out169
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
		var _4842_end__time uint64
		var _ = _4842_end__time
		var _out170 uint64
		var _ = _out170
		_out170 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
		_4842_end__time = _out170
		_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("ProposerMaybeEnterNewViewAndSend1a_work"), _4832_start__time, _4842_end__time)
		noop = false
	} else {
		proposer_k = proposer
		sent__packets = _217_LiveRSL____CMessage__i_Compile.CBroadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast_CBroadcastNop{}}
		var _4843_end__time uint64
		var _ = _4843_end__time
		var _out171 uint64
		var _ = _out171
		_out171 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
		_4843_end__time = _out171
		_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("ProposerMaybeEnterNewViewAndSend1a_nada"), _4832_start__time, _4843_end__time)
		noop = true
	}
	return proposer_k, sent__packets, noop
}
func (_this *CompanionStruct_Default___) ProposerProcess1b(proposer _254_LiveRSL____ProposerState__i_Compile.ProposerState, packet _217_LiveRSL____CMessage__i_Compile.CPacket) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
	var proposer_k _254_LiveRSL____ProposerState__i_Compile.ProposerState = _254_LiveRSL____ProposerState__i_Compile.Type_ProposerState_().Default().(_254_LiveRSL____ProposerState__i_Compile.ProposerState)
	var _ = proposer_k
	proposer_k = func(_pat_let45_0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
		return func(_4844_dt__update__tmp_h0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
			return func(_pat_let46_0 _dafny.Set) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
				return func(_4845_dt__update_hreceived__1b__packets_h0 _dafny.Set) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
					return _254_LiveRSL____ProposerState__i_Compile.ProposerState{_254_LiveRSL____ProposerState__i_Compile.ProposerState_ProposerState{(_4844_dt__update__tmp_h0).Dtor_constants(), (_4844_dt__update__tmp_h0).Dtor_current__state(), (_4844_dt__update__tmp_h0).Dtor_request__queue(), (_4844_dt__update__tmp_h0).Dtor_max__ballot__i__sent__1a(), (_4844_dt__update__tmp_h0).Dtor_next__operation__number__to__propose(), _4845_dt__update_hreceived__1b__packets_h0, (_4844_dt__update__tmp_h0).Dtor_highest__seqno__requested__by__client__this__view(), (_4844_dt__update__tmp_h0).Dtor_incomplete__batch__timer(), (_4844_dt__update__tmp_h0).Dtor_election__state(), (_4844_dt__update__tmp_h0).Dtor_maxOpnWithProposal(), (_4844_dt__update__tmp_h0).Dtor_maxLogTruncationPoint()}}
				}(_pat_let46_0)
			}(((proposer).Dtor_received__1b__packets()).Union(_dafny.SetOf(packet)))
		}(_pat_let45_0)
	}(proposer)
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
	return proposer_k
}
func (_this *CompanionStruct_Default___) GetMaxOpnWithProposalFromSingleton(m _dafny.Map) _214_LiveRSL____CTypes__i_Compile.COperationNumber {
	var maxOpn _214_LiveRSL____CTypes__i_Compile.COperationNumber = _214_LiveRSL____CTypes__i_Compile.Type_COperationNumber_().Default().(_214_LiveRSL____CTypes__i_Compile.COperationNumber)
	var _ = maxOpn
	if ((m).Cardinality()).Cmp(_dafny.IntOfInt64(1)) == 0 {
		var _4846_opn _214_LiveRSL____CTypes__i_Compile.COperationNumber
		var _ = _4846_opn
		for _iter7 := _dafny.Iterate((m).Keys().Elements()); ; {
			_val7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			_assign_such_that_4 := _val7.(_214_LiveRSL____CTypes__i_Compile.COperationNumber)
			_4846_opn = _assign_such_that_4
			if (m).Contains(_4846_opn) {
				goto L_ASSIGN_SUCH_THAT_4
			}
		}
		panic("assign-such-that search produced no value (line 300)")
	L_ASSIGN_SUCH_THAT_4:
		{
		}
		{
		}
		maxOpn = _4846_opn
	} else {
		var _4847_opn _214_LiveRSL____CTypes__i_Compile.COperationNumber
		var _ = _4847_opn
		for _iter8 := _dafny.Iterate((m).Keys().Elements()); ; {
			_val8, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			_assign_such_that_5 := _val8.(_214_LiveRSL____CTypes__i_Compile.COperationNumber)
			_4847_opn = _assign_such_that_5
			if (m).Contains(_4847_opn) {
				goto L_ASSIGN_SUCH_THAT_5
			}
		}
		panic("assign-such-that search produced no value (line 305)")
	L_ASSIGN_SUCH_THAT_5:
		var _4848_rest _dafny.Map
		var _ = _4848_rest
		_4848_rest = _118_Collections____Maps__i_Compile.Companion_Default___.RemoveElt(m, _4847_opn)
		var _4849_restMax _214_LiveRSL____CTypes__i_Compile.COperationNumber = _214_LiveRSL____CTypes__i_Compile.Type_COperationNumber_().Default().(_214_LiveRSL____CTypes__i_Compile.COperationNumber)
		var _ = _4849_restMax
		var _out172 _214_LiveRSL____CTypes__i_Compile.COperationNumber
		var _ = _out172
		_out172 = Companion_Default___.GetMaxOpnWithProposalFromSingleton(_4848_rest)
		_4849_restMax = _out172
		if ((_4849_restMax).Dtor_n()) > ((_4847_opn).Dtor_n()) {
			maxOpn = _4849_restMax
		} else {
			maxOpn = _4847_opn
		}
	}
	return maxOpn
}
func (_this *CompanionStruct_Default___) GetMaxOpnWithProposalFromSet(s _dafny.Set) (_214_LiveRSL____CTypes__i_Compile.COperationNumber, bool) {
	var maxOpn _214_LiveRSL____CTypes__i_Compile.COperationNumber = _214_LiveRSL____CTypes__i_Compile.Type_COperationNumber_().Default().(_214_LiveRSL____CTypes__i_Compile.COperationNumber)
	var _ = maxOpn
	var foundNonEmpty bool = false
	var _ = foundNonEmpty
	if ((s).Cardinality()).Cmp(_dafny.IntOfInt64(1)) == 0 {
		var _4850_p _217_LiveRSL____CMessage__i_Compile.CPacket
		var _ = _4850_p
		for _iter9 := _dafny.Iterate((s).Elements()); ; {
			_val9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			_assign_such_that_6 := _val9.(_217_LiveRSL____CMessage__i_Compile.CPacket)
			_4850_p = _assign_such_that_6
			if (s).Contains(_4850_p) {
				goto L_ASSIGN_SUCH_THAT_6
			}
		}
		panic("assign-such-that search produced no value (line 327)")
	L_ASSIGN_SUCH_THAT_6:
		{
		}
		{
		}
		if (((((_4850_p).Dtor_msg()).Dtor_votes()).Dtor_v()).Cardinality()).Cmp(_dafny.Zero) > 0 {
			var _out173 _214_LiveRSL____CTypes__i_Compile.COperationNumber
			var _ = _out173
			_out173 = Companion_Default___.GetMaxOpnWithProposalFromSingleton((((_4850_p).Dtor_msg()).Dtor_votes()).Dtor_v())
			maxOpn = _out173
			foundNonEmpty = true
		} else {
			maxOpn = _214_LiveRSL____CTypes__i_Compile.COperationNumber{_214_LiveRSL____CTypes__i_Compile.COperationNumber_COperationNumber{uint64(0)}}
			foundNonEmpty = false
		}
	} else {
		var _4851_p _217_LiveRSL____CMessage__i_Compile.CPacket
		var _ = _4851_p
		for _iter10 := _dafny.Iterate((s).Elements()); ; {
			_val10, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			_assign_such_that_7 := _val10.(_217_LiveRSL____CMessage__i_Compile.CPacket)
			_4851_p = _assign_such_that_7
			if (s).Contains(_4851_p) {
				goto L_ASSIGN_SUCH_THAT_7
			}
		}
		panic("assign-such-that search produced no value (line 343)")
	L_ASSIGN_SUCH_THAT_7:
		var _4852_rest _dafny.Set
		var _ = _4852_rest
		_4852_rest = (s).Difference(_dafny.SetOf(_4851_p))
		var _4853_candidateOpn _214_LiveRSL____CTypes__i_Compile.COperationNumber = _214_LiveRSL____CTypes__i_Compile.Type_COperationNumber_().Default().(_214_LiveRSL____CTypes__i_Compile.COperationNumber)
		var _ = _4853_candidateOpn
		var _4854_foundLocal bool = false
		var _ = _4854_foundLocal
		if (((((_4851_p).Dtor_msg()).Dtor_votes()).Dtor_v()).Cardinality()).Cmp(_dafny.Zero) > 0 {
			var _out174 _214_LiveRSL____CTypes__i_Compile.COperationNumber
			var _ = _out174
			_out174 = Companion_Default___.GetMaxOpnWithProposalFromSingleton((((_4851_p).Dtor_msg()).Dtor_votes()).Dtor_v())
			_4853_candidateOpn = _out174
			_4854_foundLocal = true
			{
			}
		} else {
			_4853_candidateOpn = _214_LiveRSL____CTypes__i_Compile.COperationNumber{_214_LiveRSL____CTypes__i_Compile.COperationNumber_COperationNumber{uint64(0)}}
			_4854_foundLocal = false
		}
		{
		}
		var _4855_restMaxOpn _214_LiveRSL____CTypes__i_Compile.COperationNumber
		var _ = _4855_restMaxOpn
		var _4856_foundTemp bool
		var _ = _4856_foundTemp
		var _out175 _214_LiveRSL____CTypes__i_Compile.COperationNumber
		var _ = _out175
		var _out176 bool
		var _ = _out176
		_out175, _out176 = Companion_Default___.GetMaxOpnWithProposalFromSet(_4852_rest)
		_4855_restMaxOpn = _out175
		_4856_foundTemp = _out176
		if (_4856_foundTemp) || (_4854_foundLocal) {
			foundNonEmpty = true
		} else {
			foundNonEmpty = false
		}
		if ((_4853_candidateOpn).Dtor_n()) > ((_4855_restMaxOpn).Dtor_n()) {
			maxOpn = _4853_candidateOpn
		} else {
			maxOpn = _4855_restMaxOpn
		}
	}
	return maxOpn, foundNonEmpty
}
func (_this *CompanionStruct_Default___) GetMaxLogTruncationPoint(s _dafny.Set) _214_LiveRSL____CTypes__i_Compile.COperationNumber {
	var maxLogTruncationPoint _214_LiveRSL____CTypes__i_Compile.COperationNumber = _214_LiveRSL____CTypes__i_Compile.Type_COperationNumber_().Default().(_214_LiveRSL____CTypes__i_Compile.COperationNumber)
	var _ = maxLogTruncationPoint
	if ((s).Cardinality()).Cmp(_dafny.IntOfInt64(1)) == 0 {
		var _4857_p _217_LiveRSL____CMessage__i_Compile.CPacket
		var _ = _4857_p
		for _iter11 := _dafny.Iterate((s).Elements()); ; {
			_val11, _ok11 := _iter11()
			if !_ok11 {
				break
			}
			_assign_such_that_8 := _val11.(_217_LiveRSL____CMessage__i_Compile.CPacket)
			_4857_p = _assign_such_that_8
			if (s).Contains(_4857_p) {
				goto L_ASSIGN_SUCH_THAT_8
			}
		}
		panic("assign-such-that search produced no value (line 383)")
	L_ASSIGN_SUCH_THAT_8:
		{
		}
		{
		}
		maxLogTruncationPoint = ((_4857_p).Dtor_msg()).Dtor_log__truncation__point()
	} else {
		var _4858_p _217_LiveRSL____CMessage__i_Compile.CPacket
		var _ = _4858_p
		for _iter12 := _dafny.Iterate((s).Elements()); ; {
			_val12, _ok12 := _iter12()
			if !_ok12 {
				break
			}
			_assign_such_that_9 := _val12.(_217_LiveRSL____CMessage__i_Compile.CPacket)
			_4858_p = _assign_such_that_9
			if (s).Contains(_4858_p) {
				goto L_ASSIGN_SUCH_THAT_9
			}
		}
		panic("assign-such-that search produced no value (line 393)")
	L_ASSIGN_SUCH_THAT_9:
		var _4859_rest _dafny.Set
		var _ = _4859_rest
		_4859_rest = (s).Difference(_dafny.SetOf(_4858_p))
		var _4860_candidateOpn _214_LiveRSL____CTypes__i_Compile.COperationNumber
		var _ = _4860_candidateOpn
		_4860_candidateOpn = ((_4858_p).Dtor_msg()).Dtor_log__truncation__point()
		{
		}
		var _4861_restMaxOpn _214_LiveRSL____CTypes__i_Compile.COperationNumber
		var _ = _4861_restMaxOpn
		var _out177 _214_LiveRSL____CTypes__i_Compile.COperationNumber
		var _ = _out177
		_out177 = Companion_Default___.GetMaxLogTruncationPoint(_4859_rest)
		_4861_restMaxOpn = _out177
		if ((_4860_candidateOpn).Dtor_n()) > ((_4861_restMaxOpn).Dtor_n()) {
			maxLogTruncationPoint = _4860_candidateOpn
		} else {
			maxLogTruncationPoint = _4861_restMaxOpn
		}
	}
	return maxLogTruncationPoint
}
func (_this *CompanionStruct_Default___) ProposerMaybeEnterPhase2(proposer _254_LiveRSL____ProposerState__i_Compile.ProposerState, log__truncation__point _214_LiveRSL____CTypes__i_Compile.COperationNumber) (_254_LiveRSL____ProposerState__i_Compile.ProposerState, _217_LiveRSL____CMessage__i_Compile.CBroadcast, bool) {
	var noop bool
	var proposer_k _254_LiveRSL____ProposerState__i_Compile.ProposerState = _254_LiveRSL____ProposerState__i_Compile.Type_ProposerState_().Default().(_254_LiveRSL____ProposerState__i_Compile.ProposerState)
	var _ = proposer_k
	var sent__packets _217_LiveRSL____CMessage__i_Compile.CBroadcast = _217_LiveRSL____CMessage__i_Compile.Type_CBroadcast_().Default().(_217_LiveRSL____CMessage__i_Compile.CBroadcast)
	var _ = sent__packets
	var _4862_start__time uint64
	var _ = _4862_start__time
	var _out178 uint64
	var _ = _out178
	_out178 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_4862_start__time = _out178
	{
	}
	{
	}
	{
	}
	{
	}
	var _4863_quorum__size uint64
	var _ = _4863_quorum__size
	var _out179 uint64
	var _ = _out179
	_out179 = _294_LiveRSL____MinCQuorumSize__i_Compile.Companion_Default___.MinCQuorumSize((((proposer).Dtor_constants()).Dtor_all()).Dtor_config())
	_4863_quorum__size = _out179
	{
	}
	if ((uint64(((proposer).Dtor_received__1b__packets()).CardinalityInt())) >= (_4863_quorum__size)) && (((proposer).Dtor_current__state()) == (uint8(1))) {
		{
		}
		{
		}
		{
		}
		var _4864_maxOpn _214_LiveRSL____CTypes__i_Compile.COperationNumber
		var _ = _4864_maxOpn
		var _4865_foundNonEmpty bool
		var _ = _4865_foundNonEmpty
		var _out180 _214_LiveRSL____CTypes__i_Compile.COperationNumber
		var _ = _out180
		var _out181 bool
		var _ = _out181
		_out180, _out181 = Companion_Default___.GetMaxOpnWithProposalFromSet((proposer).Dtor_received__1b__packets())
		_4864_maxOpn = _out180
		_4865_foundNonEmpty = _out181
		if !(_4865_foundNonEmpty) {
			_4864_maxOpn = _214_LiveRSL____CTypes__i_Compile.COperationNumber{_214_LiveRSL____CTypes__i_Compile.COperationNumber_COperationNumber{uint64(0)}}
		} else if ((_4864_maxOpn).Dtor_n()) < (uint64(18446744073709551615)) {
			_4864_maxOpn = _214_LiveRSL____CTypes__i_Compile.COperationNumber{_214_LiveRSL____CTypes__i_Compile.COperationNumber_COperationNumber{((_4864_maxOpn).Dtor_n()) + (uint64(1))}}
		}
		var _4866_maxLogTP _214_LiveRSL____CTypes__i_Compile.COperationNumber
		var _ = _4866_maxLogTP
		var _out182 _214_LiveRSL____CTypes__i_Compile.COperationNumber
		var _ = _out182
		_out182 = Companion_Default___.GetMaxLogTruncationPoint((proposer).Dtor_received__1b__packets())
		_4866_maxLogTP = _out182
		proposer_k = func(_pat_let47_0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
			return func(_4867_dt__update__tmp_h0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
				return func(_pat_let48_0 _214_LiveRSL____CTypes__i_Compile.COperationNumber) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
					return func(_4868_dt__update_hmaxLogTruncationPoint_h0 _214_LiveRSL____CTypes__i_Compile.COperationNumber) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
						return func(_pat_let49_0 _214_LiveRSL____CTypes__i_Compile.COperationNumber) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
							return func(_4869_dt__update_hmaxOpnWithProposal_h0 _214_LiveRSL____CTypes__i_Compile.COperationNumber) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
								return func(_pat_let50_0 uint64) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
									return func(_4870_dt__update_hnext__operation__number__to__propose_h0 uint64) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
										return func(_pat_let51_0 uint8) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
											return func(_4871_dt__update_hcurrent__state_h0 uint8) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
												return _254_LiveRSL____ProposerState__i_Compile.ProposerState{_254_LiveRSL____ProposerState__i_Compile.ProposerState_ProposerState{(_4867_dt__update__tmp_h0).Dtor_constants(), _4871_dt__update_hcurrent__state_h0, (_4867_dt__update__tmp_h0).Dtor_request__queue(), (_4867_dt__update__tmp_h0).Dtor_max__ballot__i__sent__1a(), _4870_dt__update_hnext__operation__number__to__propose_h0, (_4867_dt__update__tmp_h0).Dtor_received__1b__packets(), (_4867_dt__update__tmp_h0).Dtor_highest__seqno__requested__by__client__this__view(), (_4867_dt__update__tmp_h0).Dtor_incomplete__batch__timer(), (_4867_dt__update__tmp_h0).Dtor_election__state(), _4869_dt__update_hmaxOpnWithProposal_h0, _4868_dt__update_hmaxLogTruncationPoint_h0}}
											}(_pat_let51_0)
										}(uint8(2))
									}(_pat_let50_0)
								}((log__truncation__point).Dtor_n())
							}(_pat_let49_0)
						}(_4864_maxOpn)
					}(_pat_let48_0)
				}(_4866_maxLogTP)
			}(_pat_let47_0)
		}(proposer)
		var _4872_msg _217_LiveRSL____CMessage__i_Compile.CMessage
		var _ = _4872_msg
		_4872_msg = _217_LiveRSL____CMessage__i_Compile.CMessage{_217_LiveRSL____CMessage__i_Compile.CMessage_CMessage__StartingPhase2{(proposer).Dtor_max__ballot__i__sent__1a(), log__truncation__point}}
		{
		}
		var _out183 _217_LiveRSL____CMessage__i_Compile.CBroadcast
		var _ = _out183
		_out183 = _301_Impl____LiveRSL____Broadcast__i_Compile.Companion_Default___.BuildBroadcastToEveryone((((proposer).Dtor_constants()).Dtor_all()).Dtor_config(), ((proposer).Dtor_constants()).Dtor_my__index(), _4872_msg)
		sent__packets = _out183
		var _4873_end__time uint64
		var _ = _4873_end__time
		var _out184 uint64
		var _ = _out184
		_out184 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
		_4873_end__time = _out184

		var proposerState = proposer.Dtor_current__state()
		fmt.Printf("Enter phase 2: proposer state %v\n", proposerState)

		_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("ProposerMaybeEnterPhase2_work"), _4862_start__time, _4873_end__time)
		noop = false
	} else {
		proposer_k = proposer
		sent__packets = _217_LiveRSL____CMessage__i_Compile.CBroadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast_CBroadcastNop{}}
		{
		}
		var _4874_end__time uint64
		var _ = _4874_end__time
		var _out185 uint64
		var _ = _out185
		_out185 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
		_4874_end__time = _out185
		_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("ProposerMaybeEnterPhase2_nada"), _4862_start__time, _4874_end__time)
		noop = true
	}
	{
	}
	{
	}
	return proposer_k, sent__packets, noop
}
func (_this *CompanionStruct_Default___) ProposerNominateNewValueAndSend2a(proposer _254_LiveRSL____ProposerState__i_Compile.ProposerState, clock uint64, log__truncation__point _214_LiveRSL____CTypes__i_Compile.COperationNumber) (_254_LiveRSL____ProposerState__i_Compile.ProposerState, _217_LiveRSL____CMessage__i_Compile.CBroadcast) {
	// fmt.Printf("TONY DEBUG: ProposerNominateNewValueAndSend2a\n")
	var proposer_k _254_LiveRSL____ProposerState__i_Compile.ProposerState = _254_LiveRSL____ProposerState__i_Compile.Type_ProposerState_().Default().(_254_LiveRSL____ProposerState__i_Compile.ProposerState)
	var _ = proposer_k
	var sent__packets _217_LiveRSL____CMessage__i_Compile.CBroadcast = _217_LiveRSL____CMessage__i_Compile.Type_CBroadcast_().Default().(_217_LiveRSL____CMessage__i_Compile.CBroadcast)
	var _ = sent__packets
	var _4875_batchSize _dafny.Int
	var _ = _4875_batchSize
	_4875_batchSize = (func() _dafny.Int {
		if ((((proposer).Dtor_request__queue()).Cardinality()).Cmp(_dafny.IntOfUint64(((((proposer).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__batch__size())) <= 0) || ((((((proposer).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__batch__size()) < (uint64(0))) {
			return ((proposer).Dtor_request__queue()).Cardinality()
		}
		return _dafny.IntOfUint64(((((proposer).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__batch__size())
	})()
	var _4876_v _dafny.Seq
	var _ = _4876_v
	_4876_v = ((proposer).Dtor_request__queue()).Subseq(_dafny.NilInt, _4875_batchSize)
	var _4877_opn uint64
	var _ = _4877_opn
	_4877_opn = (proposer).Dtor_next__operation__number__to__propose()
	var _4878_opn__op _214_LiveRSL____CTypes__i_Compile.COperationNumber
	var _ = _4878_opn__op
	_4878_opn__op = _214_LiveRSL____CTypes__i_Compile.COperationNumber{_214_LiveRSL____CTypes__i_Compile.COperationNumber_COperationNumber{(proposer).Dtor_next__operation__number__to__propose()}}
	var _4879_clock__sum uint64
	var _ = _4879_clock__sum
	var _out186 uint64
	var _ = _out186
	_out186 = _234_Common____UpperBound__i_Compile.Companion_Default___.UpperBoundedAdditionImpl(clock, ((((proposer).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__batch__delay(), ((((proposer).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__integer__val())
	_4879_clock__sum = _out186
	var _4880_newTimer _254_LiveRSL____ProposerState__i_Compile.CIncompleteBatchTimer
	var _ = _4880_newTimer
	_4880_newTimer = (func() _254_LiveRSL____ProposerState__i_Compile.CIncompleteBatchTimer {
		if (((proposer).Dtor_request__queue()).Cardinality()).Cmp(_4875_batchSize) > 0 {
			return _254_LiveRSL____ProposerState__i_Compile.CIncompleteBatchTimer{_254_LiveRSL____ProposerState__i_Compile.CIncompleteBatchTimer_CIncompleteBatchTimerOn{_4879_clock__sum}}
		}
		return _254_LiveRSL____ProposerState__i_Compile.CIncompleteBatchTimer{_254_LiveRSL____ProposerState__i_Compile.CIncompleteBatchTimer_CIncompleteBatchTimerOff{}}
	})()
	proposer_k = func(_pat_let52_0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
		return func(_4881_dt__update__tmp_h0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
			return func(_pat_let53_0 _254_LiveRSL____ProposerState__i_Compile.CIncompleteBatchTimer) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
				return func(_4882_dt__update_hincomplete__batch__timer_h0 _254_LiveRSL____ProposerState__i_Compile.CIncompleteBatchTimer) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
					return func(_pat_let54_0 uint64) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
						return func(_4883_dt__update_hnext__operation__number__to__propose_h0 uint64) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
							return func(_pat_let55_0 _dafny.Seq) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
								return func(_4884_dt__update_hrequest__queue_h0 _dafny.Seq) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
									return _254_LiveRSL____ProposerState__i_Compile.ProposerState{_254_LiveRSL____ProposerState__i_Compile.ProposerState_ProposerState{(_4881_dt__update__tmp_h0).Dtor_constants(), (_4881_dt__update__tmp_h0).Dtor_current__state(), _4884_dt__update_hrequest__queue_h0, (_4881_dt__update__tmp_h0).Dtor_max__ballot__i__sent__1a(), _4883_dt__update_hnext__operation__number__to__propose_h0, (_4881_dt__update__tmp_h0).Dtor_received__1b__packets(), (_4881_dt__update__tmp_h0).Dtor_highest__seqno__requested__by__client__this__view(), _4882_dt__update_hincomplete__batch__timer_h0, (_4881_dt__update__tmp_h0).Dtor_election__state(), (_4881_dt__update__tmp_h0).Dtor_maxOpnWithProposal(), (_4881_dt__update__tmp_h0).Dtor_maxLogTruncationPoint()}}
								}(_pat_let55_0)
							}(((proposer).Dtor_request__queue()).Subseq(_4875_batchSize, _dafny.NilInt))
						}(_pat_let54_0)
					}((_4877_opn) + (uint64(1)))
				}(_pat_let53_0)
			}(_4880_newTimer)
		}(_pat_let52_0)
	}(proposer)
	var _4885_msg _217_LiveRSL____CMessage__i_Compile.CMessage
	var _ = _4885_msg
	_4885_msg = _217_LiveRSL____CMessage__i_Compile.CMessage{_217_LiveRSL____CMessage__i_Compile.CMessage_CMessage__2a{(proposer).Dtor_max__ballot__i__sent__1a(), _4878_opn__op, _4876_v}}
	{
	}
	var _out187 _217_LiveRSL____CMessage__i_Compile.CBroadcast
	var _ = _out187
	_out187 = _301_Impl____LiveRSL____Broadcast__i_Compile.Companion_Default___.BuildBroadcastToEveryone((((proposer).Dtor_constants()).Dtor_all()).Dtor_config(), ((proposer).Dtor_constants()).Dtor_my__index(), _4885_msg)
	sent__packets = _out187
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
	return proposer_k, sent__packets
}
func (_this *CompanionStruct_Default___) FindValWithHighestNumberedProposal(received__1b__packets _dafny.Set, opn _214_LiveRSL____CTypes__i_Compile.COperationNumber) _dafny.Seq {
	var v _dafny.Seq = _dafny.EmptySeq
	var _ = v
	var _4886_packets _dafny.Set = _dafny.EmptySet
	var _ = _4886_packets
	{
	}
	_4886_packets = received__1b__packets
	var _4887_pkt _217_LiveRSL____CMessage__i_Compile.CPacket
	var _ = _4887_pkt
	for _iter13 := _dafny.Iterate((_4886_packets).Elements()); ; {
		_val13, _ok13 := _iter13()
		if !_ok13 {
			break
		}
		_assign_such_that_10 := _val13.(_217_LiveRSL____CMessage__i_Compile.CPacket)
		_4887_pkt = _assign_such_that_10
		if ((_4886_packets).Contains(_4887_pkt)) && (((((_4887_pkt).Dtor_msg()).Dtor_votes()).Dtor_v()).Contains(opn)) {
			goto L_ASSIGN_SUCH_THAT_10
		}
	}
	panic("assign-such-that search produced no value (line 733)")
L_ASSIGN_SUCH_THAT_10:
	v = (((((_4887_pkt).Dtor_msg()).Dtor_votes()).Dtor_v()).Get(opn).(_214_LiveRSL____CTypes__i_Compile.CVote)).Dtor_max__val()
	var _4888_bal _214_LiveRSL____CTypes__i_Compile.CBallot
	var _ = _4888_bal
	_4888_bal = (((((_4887_pkt).Dtor_msg()).Dtor_votes()).Dtor_v()).Get(opn).(_214_LiveRSL____CTypes__i_Compile.CVote)).Dtor_max__value__bal()
	var _4889_p__bal _217_LiveRSL____CMessage__i_Compile.CPacket
	var _ = _4889_p__bal
	_4889_p__bal = _4887_pkt
	_4886_packets = (_4886_packets).Difference(_dafny.SetOf(_4887_pkt))
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
	for !(_4886_packets).Equals(_dafny.SetOf()) {
		for _iter14 := _dafny.Iterate((_4886_packets).Elements()); ; {
			_val14, _ok14 := _iter14()
			if !_ok14 {
				break
			}
			_assign_such_that_11 := _val14.(_217_LiveRSL____CMessage__i_Compile.CPacket)
			_4887_pkt = _assign_such_that_11
			if (_4886_packets).Contains(_4887_pkt) {
				goto L_ASSIGN_SUCH_THAT_11
			}
		}
		panic("assign-such-that search produced no value (line 760)")
	L_ASSIGN_SUCH_THAT_11:
		if ((((_4887_pkt).Dtor_msg()).Dtor_votes()).Dtor_v()).Contains(opn) {
			var _4890_foundHigherBallot bool
			var _ = _4890_foundHigherBallot
			var _out188 bool
			var _ = _out188
			_out188 = _214_LiveRSL____CTypes__i_Compile.Companion_Default___.CBalLeq(_4888_bal, (((((_4887_pkt).Dtor_msg()).Dtor_votes()).Dtor_v()).Get(opn).(_214_LiveRSL____CTypes__i_Compile.CVote)).Dtor_max__value__bal())
			_4890_foundHigherBallot = _out188
			if _4890_foundHigherBallot {
				_4889_p__bal = _4887_pkt
				v = (((((_4887_pkt).Dtor_msg()).Dtor_votes()).Dtor_v()).Get(opn).(_214_LiveRSL____CTypes__i_Compile.CVote)).Dtor_max__val()
				_4888_bal = (((((_4887_pkt).Dtor_msg()).Dtor_votes()).Dtor_v()).Get(opn).(_214_LiveRSL____CTypes__i_Compile.CVote)).Dtor_max__value__bal()
			}
		}
		_4886_packets = (_4886_packets).Difference(_dafny.SetOf(_4887_pkt))
		{
		}
		{
		}
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
	return v
}
func (_this *CompanionStruct_Default___) ProposerNominateOldValueAndSend2a(proposer _254_LiveRSL____ProposerState__i_Compile.ProposerState, log__truncation__point _214_LiveRSL____CTypes__i_Compile.COperationNumber) (_254_LiveRSL____ProposerState__i_Compile.ProposerState, _217_LiveRSL____CMessage__i_Compile.CBroadcast) {
	// fmt.Printf("TONY DEBUG : ProposerNominateOldValueAndSend2a\n")
	var proposer_k _254_LiveRSL____ProposerState__i_Compile.ProposerState = _254_LiveRSL____ProposerState__i_Compile.Type_ProposerState_().Default().(_254_LiveRSL____ProposerState__i_Compile.ProposerState)
	var _ = proposer_k
	var sent__packets _217_LiveRSL____CMessage__i_Compile.CBroadcast = _217_LiveRSL____CMessage__i_Compile.Type_CBroadcast_().Default().(_217_LiveRSL____CMessage__i_Compile.CBroadcast)
	var _ = sent__packets
	{
	}
	var _4891_opn uint64
	var _ = _4891_opn
	_4891_opn = (proposer).Dtor_next__operation__number__to__propose()
	var _4892_opn__op _214_LiveRSL____CTypes__i_Compile.COperationNumber
	var _ = _4892_opn__op
	_4892_opn__op = _214_LiveRSL____CTypes__i_Compile.COperationNumber{_214_LiveRSL____CTypes__i_Compile.COperationNumber_COperationNumber{(proposer).Dtor_next__operation__number__to__propose()}}
	var _4893_val _dafny.Seq
	var _ = _4893_val
	var _out189 _dafny.Seq
	var _ = _out189
	_out189 = Companion_Default___.FindValWithHighestNumberedProposal((proposer).Dtor_received__1b__packets(), _4892_opn__op)
	_4893_val = _out189
	{
	}
	var _4894_sum uint64
	var _ = _4894_sum
	_4894_sum = (_4891_opn) + (uint64(1))
	proposer_k = func(_pat_let56_0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
		return func(_4895_dt__update__tmp_h0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
			return func(_pat_let57_0 uint64) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
				return func(_4896_dt__update_hnext__operation__number__to__propose_h0 uint64) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
					return _254_LiveRSL____ProposerState__i_Compile.ProposerState{_254_LiveRSL____ProposerState__i_Compile.ProposerState_ProposerState{(_4895_dt__update__tmp_h0).Dtor_constants(), (_4895_dt__update__tmp_h0).Dtor_current__state(), (_4895_dt__update__tmp_h0).Dtor_request__queue(), (_4895_dt__update__tmp_h0).Dtor_max__ballot__i__sent__1a(), _4896_dt__update_hnext__operation__number__to__propose_h0, (_4895_dt__update__tmp_h0).Dtor_received__1b__packets(), (_4895_dt__update__tmp_h0).Dtor_highest__seqno__requested__by__client__this__view(), (_4895_dt__update__tmp_h0).Dtor_incomplete__batch__timer(), (_4895_dt__update__tmp_h0).Dtor_election__state(), (_4895_dt__update__tmp_h0).Dtor_maxOpnWithProposal(), (_4895_dt__update__tmp_h0).Dtor_maxLogTruncationPoint()}}
				}(_pat_let57_0)
			}(_4894_sum)
		}(_pat_let56_0)
	}(proposer)
	{
	}
	{
	}
	var _4897_msg _217_LiveRSL____CMessage__i_Compile.CMessage
	var _ = _4897_msg
	_4897_msg = _217_LiveRSL____CMessage__i_Compile.CMessage{_217_LiveRSL____CMessage__i_Compile.CMessage_CMessage__2a{(proposer).Dtor_max__ballot__i__sent__1a(), _4892_opn__op, _4893_val}}
	{
	}
	var _out190 _217_LiveRSL____CMessage__i_Compile.CBroadcast
	var _ = _out190
	_out190 = _301_Impl____LiveRSL____Broadcast__i_Compile.Companion_Default___.BuildBroadcastToEveryone((((proposer).Dtor_constants()).Dtor_all()).Dtor_config(), ((proposer).Dtor_constants()).Dtor_my__index(), _4897_msg)
	sent__packets = _out190
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
	return proposer_k, sent__packets
}
func (_this *CompanionStruct_Default___) IsAfterLogTruncationPointImpl(opn _214_LiveRSL____CTypes__i_Compile.COperationNumber, received__1b__packets _dafny.Set) bool {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var b bool = false
	var _ = b
	{
	}
	{
	}
	{
	}
	b = _dafny.Quantifier((received__1b__packets).Elements(), true, func(_4898_p _217_LiveRSL____CMessage__i_Compile.CPacket) bool {
		return !(((received__1b__packets).Contains(_4898_p)) && (((_4898_p).Dtor_msg()).Is_CMessage__1b())) || (((((_4898_p).Dtor_msg()).Dtor_log__truncation__point()).Dtor_n()) <= ((opn).Dtor_n()))
	})
	return b
}
func (_this *CompanionStruct_Default___) Proposer__CanNominateUsingOperationNumberImpl(proposer _254_LiveRSL____ProposerState__i_Compile.ProposerState, log__truncation__point _214_LiveRSL____CTypes__i_Compile.COperationNumber) bool {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var b bool = false
	var _ = b
	if ((proposer).Dtor_current__state()) == (uint8(2)) {
		var _4899_opn _214_LiveRSL____CTypes__i_Compile.COperationNumber
		var _ = _4899_opn
		_4899_opn = _214_LiveRSL____CTypes__i_Compile.COperationNumber{_214_LiveRSL____CTypes__i_Compile.COperationNumber_COperationNumber{(proposer).Dtor_next__operation__number__to__propose()}}
		var _4900_quorum__size uint64
		var _ = _4900_quorum__size
		var _out191 uint64
		var _ = _out191
		_out191 = _294_LiveRSL____MinCQuorumSize__i_Compile.Companion_Default___.MinCQuorumSize((((proposer).Dtor_constants()).Dtor_all()).Dtor_config())
		_4900_quorum__size = _out191
		var _4901_after__trunk bool = false
		var _ = _4901_after__trunk
		if ((_4899_opn).Dtor_n()) >= (((proposer).Dtor_maxLogTruncationPoint()).Dtor_n()) {
			_4901_after__trunk = true
		} else {
			var _out192 bool
			var _ = _out192
			_out192 = Companion_Default___.IsAfterLogTruncationPointImpl(_4899_opn, (proposer).Dtor_received__1b__packets())
			_4901_after__trunk = _out192
		}
		var _4902_sum uint64
		var _ = _4902_sum
		var _out193 uint64
		var _ = _out193
		_out193 = _234_Common____UpperBound__i_Compile.Companion_Default___.UpperBoundedAdditionImpl((log__truncation__point).Dtor_n(), ((((proposer).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__log__length(), ((((proposer).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__integer__val())
		_4902_sum = _out193
		{
		}
		{
		}
		b = ((((((((proposer).Dtor_election__state()).Dtor_current__view()).Equals((proposer).Dtor_max__ballot__i__sent__1a())) && (((proposer).Dtor_current__state()) == (uint8(2)))) && ((uint64(((proposer).Dtor_received__1b__packets()).CardinalityInt())) >= (_4900_quorum__size))) && (_4901_after__trunk)) && (((_4899_opn).Dtor_n()) < (_4902_sum))) && (((_4899_opn).Dtor_n()) >= (uint64(0)))
		{
		}
		{
		}
	} else {
		b = false
	}
	return b
}
func (_this *CompanionStruct_Default___) AllAcceptorsHadNoProposalImpl(proposer _254_LiveRSL____ProposerState__i_Compile.ProposerState) bool {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var b bool = false
	var _ = b
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
	var _4903_start__time uint64
	var _ = _4903_start__time
	var _out194 uint64
	var _ = _out194
	_out194 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_4903_start__time = _out194
	var _4904_end__time uint64 = 0
	var _ = _4904_end__time
	if (((proposer).Dtor_next__operation__number__to__propose()) < (((proposer).Dtor_maxOpnWithProposal()).Dtor_n())) || ((((proposer).Dtor_maxOpnWithProposal()).Dtor_n()) == (uint64(18446744073709551615))) {
		var _4905_opn _214_LiveRSL____CTypes__i_Compile.COperationNumber
		var _ = _4905_opn
		_4905_opn = _214_LiveRSL____CTypes__i_Compile.COperationNumber{_214_LiveRSL____CTypes__i_Compile.COperationNumber_COperationNumber{(proposer).Dtor_next__operation__number__to__propose()}}
		b = _dafny.Quantifier(((proposer).Dtor_received__1b__packets()).Elements(), true, func(_4906_p _217_LiveRSL____CMessage__i_Compile.CPacket) bool {
			return !(((proposer).Dtor_received__1b__packets()).Contains(_4906_p)) || (!(((((_4906_p).Dtor_msg()).Dtor_votes()).Dtor_v()).Contains(_4905_opn)))
		})
		var _out195 uint64
		var _ = _out195
		_out195 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
		_4904_end__time = _out195
		_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("AllAcceptorsHadNoProposalImpl_full"), _4903_start__time, _4904_end__time)
	} else {
		b = true
		var _out196 uint64
		var _ = _out196
		_out196 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
		_4904_end__time = _out196
		_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("AllAcceptorsHadNoProposalImpl_memoized"), _4903_start__time, _4904_end__time)
	}
	{
	}
	return b
}
func (_this *CompanionStruct_Default___) DidSomeAcceptorHaveProposal(proposer _254_LiveRSL____ProposerState__i_Compile.ProposerState) bool {
	var b bool = false
	var _ = b
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
	if ((proposer).Dtor_next__operation__number__to__propose()) >= (((proposer).Dtor_maxOpnWithProposal()).Dtor_n()) {
		b = false
	} else {
		b = _dafny.Quantifier(((proposer).Dtor_received__1b__packets()).Elements(), false, func(_4907_p _217_LiveRSL____CMessage__i_Compile.CPacket) bool {
			return (((proposer).Dtor_received__1b__packets()).Contains(_4907_p)) && (_dafny.Quantifier(((((_4907_p).Dtor_msg()).Dtor_votes()).Dtor_v()).Keys().Elements(), false, func(_4908_opn _214_LiveRSL____CTypes__i_Compile.COperationNumber) bool {
				return (((((_4907_p).Dtor_msg()).Dtor_votes()).Dtor_v()).Contains(_4908_opn)) && (((_4908_opn).Dtor_n()) > ((proposer).Dtor_next__operation__number__to__propose()))
			}))
		})
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
	return b
}
func (_this *CompanionStruct_Default___) ProposerMaybeNominateValueAndSend2a(proposer _254_LiveRSL____ProposerState__i_Compile.ProposerState, clock uint64, log__truncation__point _214_LiveRSL____CTypes__i_Compile.COperationNumber) (_254_LiveRSL____ProposerState__i_Compile.ProposerState, _217_LiveRSL____CMessage__i_Compile.CBroadcast, bool) {
	// fmt.Printf("\nTONY DEBUG: ProposerMaybeNominateValueAndSend2a\n")
	var noop bool
	var proposer_k _254_LiveRSL____ProposerState__i_Compile.ProposerState = _254_LiveRSL____ProposerState__i_Compile.Type_ProposerState_().Default().(_254_LiveRSL____ProposerState__i_Compile.ProposerState)
	var _ = proposer_k
	var sent__packets _217_LiveRSL____CMessage__i_Compile.CBroadcast = _217_LiveRSL____CMessage__i_Compile.Type_CBroadcast_().Default().(_217_LiveRSL____CMessage__i_Compile.CBroadcast)
	var _ = sent__packets
	var _4909_canNominate bool
	var _ = _4909_canNominate
	var _out197 bool
	var _ = _out197
	_out197 = Companion_Default___.Proposer__CanNominateUsingOperationNumberImpl(proposer, log__truncation__point)
	_4909_canNominate = _out197
	if !(_4909_canNominate) {
		// fmt.Printf("TONY DEBUG: I can't nominate\n")
		proposer_k = proposer
		sent__packets = _217_LiveRSL____CMessage__i_Compile.CBroadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast_CBroadcastNop{}}
		noop = true
	} else {
		if ((((proposer).Dtor_next__operation__number__to__propose()) >= (((proposer).Dtor_maxOpnWithProposal()).Dtor_n())) && ((((proposer).Dtor_request__queue()).Cardinality()).Cmp(_dafny.Zero) == 0)) && ((((proposer).Dtor_maxOpnWithProposal()).Dtor_n()) < (uint64(18446744073709551615))) {
			{
			}
			{
			}
			proposer_k = proposer
			// fmt.Printf("TONY DEBUG: The happily ever after branch\n")
			sent__packets = _217_LiveRSL____CMessage__i_Compile.CBroadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast_CBroadcastNop{}}
			noop = true
		} else {
			var _4910_noProposal bool
			var _ = _4910_noProposal
			var _out198 bool
			var _ = _out198
			_out198 = Companion_Default___.AllAcceptorsHadNoProposalImpl(proposer)
			_4910_noProposal = _out198
			if !(_4910_noProposal) {
				var _out199 _254_LiveRSL____ProposerState__i_Compile.ProposerState
				var _ = _out199
				var _out200 _217_LiveRSL____CMessage__i_Compile.CBroadcast
				var _ = _out200
				_out199, _out200 = Companion_Default___.ProposerNominateOldValueAndSend2a(proposer, log__truncation__point)
				proposer_k = _out199
				sent__packets = _out200
				{
				}
				noop = false
			} else {
				var _4911_queueSize _dafny.Int
				var _ = _4911_queueSize
				_4911_queueSize = ((proposer).Dtor_request__queue()).Cardinality()
				var _4912_existsOpn bool
				var _ = _4912_existsOpn
				var _out201 bool
				var _ = _out201
				_out201 = Companion_Default___.DidSomeAcceptorHaveProposal(proposer)
				_4912_existsOpn = _out201
				{
				}
				{
				}
				{
				}
				{
				}
				// fmt.Printf("TONY DEBUG: Begining test to see if I can nominate new value\n")
				// fmt.Printf("\t\t Is the queue size non-zero : %t\n", (_4911_queueSize).Cmp(_dafny.Zero) > 0)
				// var isTimerOn = ((proposer).Dtor_incomplete__batch__timer()).Is_CIncompleteBatchTimerOn()
				// fmt.Printf("\t\t Is the timer on            : %t\n", isTimerOn)
				// if isTimerOn {
				// 	fmt.Printf("\t\t Current clock              : %v\n", clock)
				// 	fmt.Printf("\t\t Batch timer                : %v\n", ((proposer).Dtor_incomplete__batch__timer()).Dtor_when())
				// }
				if (((((_4911_queueSize).Cmp(_dafny.Zero) > 0) && (((proposer).Dtor_incomplete__batch__timer()).Is_CIncompleteBatchTimerOn())) && ((clock) >= (((proposer).Dtor_incomplete__batch__timer()).Dtor_when()))) || ((_4911_queueSize).Cmp(_dafny.IntOfUint64(((((proposer).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__batch__size())) >= 0)) || (_4912_existsOpn) {
					{
					}
					var _out202 _254_LiveRSL____ProposerState__i_Compile.ProposerState
					var _ = _out202
					var _out203 _217_LiveRSL____CMessage__i_Compile.CBroadcast
					var _ = _out203
					_out202, _out203 = Companion_Default___.ProposerNominateNewValueAndSend2a(proposer, clock, log__truncation__point)
					proposer_k = _out202
					sent__packets = _out203
					{
					}
					noop = false
				} else {
					if ((_4911_queueSize).Cmp(_dafny.Zero) > 0) && (((proposer).Dtor_incomplete__batch__timer()).Is_CIncompleteBatchTimerOff()) {
						// fmt.Printf("TONY DEBUG: Batch timer is off; turn on the batch timer\n")
						var _4913_sum uint64
						var _ = _4913_sum
						var _out204 uint64
						var _ = _out204
						_out204 = _234_Common____UpperBound__i_Compile.Companion_Default___.UpperBoundedAdditionImpl(clock, ((((proposer).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__batch__delay(), ((((proposer).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__integer__val())
						_4913_sum = _out204
						proposer_k = func(_pat_let58_0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
							return func(_4914_dt__update__tmp_h0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
								return func(_pat_let59_0 _254_LiveRSL____ProposerState__i_Compile.CIncompleteBatchTimer) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
									return func(_4915_dt__update_hincomplete__batch__timer_h0 _254_LiveRSL____ProposerState__i_Compile.CIncompleteBatchTimer) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
										return _254_LiveRSL____ProposerState__i_Compile.ProposerState{_254_LiveRSL____ProposerState__i_Compile.ProposerState_ProposerState{(_4914_dt__update__tmp_h0).Dtor_constants(), (_4914_dt__update__tmp_h0).Dtor_current__state(), (_4914_dt__update__tmp_h0).Dtor_request__queue(), (_4914_dt__update__tmp_h0).Dtor_max__ballot__i__sent__1a(), (_4914_dt__update__tmp_h0).Dtor_next__operation__number__to__propose(), (_4914_dt__update__tmp_h0).Dtor_received__1b__packets(), (_4914_dt__update__tmp_h0).Dtor_highest__seqno__requested__by__client__this__view(), _4915_dt__update_hincomplete__batch__timer_h0, (_4914_dt__update__tmp_h0).Dtor_election__state(), (_4914_dt__update__tmp_h0).Dtor_maxOpnWithProposal(), (_4914_dt__update__tmp_h0).Dtor_maxLogTruncationPoint()}}
									}(_pat_let59_0)
								}(_254_LiveRSL____ProposerState__i_Compile.CIncompleteBatchTimer{_254_LiveRSL____ProposerState__i_Compile.CIncompleteBatchTimer_CIncompleteBatchTimerOn{_4913_sum}})
							}(_pat_let58_0)
						}(proposer)
						sent__packets = _217_LiveRSL____CMessage__i_Compile.CBroadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast_CBroadcastNop{}}
						noop = false
					} else {
						// fmt.Printf("TONY DEBUG: Give up and not do anything\n")
						proposer_k = proposer
						sent__packets = _217_LiveRSL____CMessage__i_Compile.CBroadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast_CBroadcastNop{}}
						noop = true
					}
				}
			}
		}
	}
	return proposer_k, sent__packets, noop
}
func (_this *CompanionStruct_Default___) ProposerProcessHeartbeat(proposer _254_LiveRSL____ProposerState__i_Compile.ProposerState, packet _217_LiveRSL____CMessage__i_Compile.CPacket, clock uint64, cur__req__set *_9_Native____Io__s_Compile.MutableSet, prev__req__set *_9_Native____Io__s_Compile.MutableSet) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var proposer_k _254_LiveRSL____ProposerState__i_Compile.ProposerState = _254_LiveRSL____ProposerState__i_Compile.Type_ProposerState_().Default().(_254_LiveRSL____ProposerState__i_Compile.ProposerState)
	var _ = proposer_k
	var _4916_election__state_k _251_LiveRSL____ElectionState__i_Compile.CElectionState
	var _ = _4916_election__state_k
	var _out205 _251_LiveRSL____ElectionState__i_Compile.CElectionState
	var _ = _out205
	_out205 = _297_LiveRSL____ElectionModel__i_Compile.Companion_Default___.ElectionProcessHeartbeat((proposer).Dtor_election__state(), packet, clock, cur__req__set, prev__req__set)
	_4916_election__state_k = _out205
	var _4917_current__state_k uint8
	var _ = _4917_current__state_k
	_4917_current__state_k = uint8(0)
	var _4918_request__queue_k _dafny.Seq
	var _ = _4918_request__queue_k
	_4918_request__queue_k = _dafny.SeqOf()
	var _4919_lt bool
	var _ = _4919_lt
	var _out206 bool
	var _ = _out206
	_out206 = _214_LiveRSL____CTypes__i_Compile.Companion_Default___.CBalLt(((proposer).Dtor_election__state()).Dtor_current__view(), (_4916_election__state_k).Dtor_current__view())
	_4919_lt = _out206
	if _4919_lt {
		_4917_current__state_k = uint8(0)
		_4918_request__queue_k = _dafny.SeqOf()
	} else {
		_4917_current__state_k = (proposer).Dtor_current__state()
		_4918_request__queue_k = (proposer).Dtor_request__queue()
	}
	proposer_k = func(_pat_let60_0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
		return func(_4920_dt__update__tmp_h0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
			return func(_pat_let61_0 _dafny.Seq) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
				return func(_4921_dt__update_hrequest__queue_h0 _dafny.Seq) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
					return func(_pat_let62_0 uint8) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
						return func(_4922_dt__update_hcurrent__state_h0 uint8) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
							return func(_pat_let63_0 _251_LiveRSL____ElectionState__i_Compile.CElectionState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
								return func(_4923_dt__update_helection__state_h0 _251_LiveRSL____ElectionState__i_Compile.CElectionState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
									return _254_LiveRSL____ProposerState__i_Compile.ProposerState{_254_LiveRSL____ProposerState__i_Compile.ProposerState_ProposerState{(_4920_dt__update__tmp_h0).Dtor_constants(), _4922_dt__update_hcurrent__state_h0, _4921_dt__update_hrequest__queue_h0, (_4920_dt__update__tmp_h0).Dtor_max__ballot__i__sent__1a(), (_4920_dt__update__tmp_h0).Dtor_next__operation__number__to__propose(), (_4920_dt__update__tmp_h0).Dtor_received__1b__packets(), (_4920_dt__update__tmp_h0).Dtor_highest__seqno__requested__by__client__this__view(), (_4920_dt__update__tmp_h0).Dtor_incomplete__batch__timer(), _4923_dt__update_helection__state_h0, (_4920_dt__update__tmp_h0).Dtor_maxOpnWithProposal(), (_4920_dt__update__tmp_h0).Dtor_maxLogTruncationPoint()}}
								}(_pat_let63_0)
							}(_4916_election__state_k)
						}(_pat_let62_0)
					}(_4917_current__state_k)
				}(_pat_let61_0)
			}(_4918_request__queue_k)
		}(_pat_let60_0)
	}(proposer)
	return proposer_k
}
func (_this *CompanionStruct_Default___) ProposerCheckForViewTimeout(proposer _254_LiveRSL____ProposerState__i_Compile.ProposerState, clock uint64, cur__req__set *_9_Native____Io__s_Compile.MutableSet, prev__req__set *_9_Native____Io__s_Compile.MutableSet) (_254_LiveRSL____ProposerState__i_Compile.ProposerState, bool) {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var noop bool
	var proposer_k _254_LiveRSL____ProposerState__i_Compile.ProposerState = _254_LiveRSL____ProposerState__i_Compile.Type_ProposerState_().Default().(_254_LiveRSL____ProposerState__i_Compile.ProposerState)
	var _ = proposer_k
	var _4924_election__state_k _251_LiveRSL____ElectionState__i_Compile.CElectionState
	var _ = _4924_election__state_k
	var _out207 _251_LiveRSL____ElectionState__i_Compile.CElectionState
	var _ = _out207
	_out207, noop = _297_LiveRSL____ElectionModel__i_Compile.Companion_Default___.ElectionCheckForViewTimeout((proposer).Dtor_election__state(), clock, cur__req__set, prev__req__set)
	_4924_election__state_k = _out207
	proposer_k = func(_pat_let64_0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
		return func(_4925_dt__update__tmp_h0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
			return func(_pat_let65_0 _251_LiveRSL____ElectionState__i_Compile.CElectionState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
				return func(_4926_dt__update_helection__state_h0 _251_LiveRSL____ElectionState__i_Compile.CElectionState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
					return _254_LiveRSL____ProposerState__i_Compile.ProposerState{_254_LiveRSL____ProposerState__i_Compile.ProposerState_ProposerState{(_4925_dt__update__tmp_h0).Dtor_constants(), (_4925_dt__update__tmp_h0).Dtor_current__state(), (_4925_dt__update__tmp_h0).Dtor_request__queue(), (_4925_dt__update__tmp_h0).Dtor_max__ballot__i__sent__1a(), (_4925_dt__update__tmp_h0).Dtor_next__operation__number__to__propose(), (_4925_dt__update__tmp_h0).Dtor_received__1b__packets(), (_4925_dt__update__tmp_h0).Dtor_highest__seqno__requested__by__client__this__view(), (_4925_dt__update__tmp_h0).Dtor_incomplete__batch__timer(), _4926_dt__update_helection__state_h0, (_4925_dt__update__tmp_h0).Dtor_maxOpnWithProposal(), (_4925_dt__update__tmp_h0).Dtor_maxLogTruncationPoint()}}
				}(_pat_let65_0)
			}(_4924_election__state_k)
		}(_pat_let64_0)
	}(proposer)
	return proposer_k, noop
}
func (_this *CompanionStruct_Default___) ProposerCheckForQuorumOfViewSuspicions(proposer _254_LiveRSL____ProposerState__i_Compile.ProposerState, clock uint64, cur__req__set *_9_Native____Io__s_Compile.MutableSet, prev__req__set *_9_Native____Io__s_Compile.MutableSet) (_254_LiveRSL____ProposerState__i_Compile.ProposerState, bool) {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var noop bool
	var proposer_k _254_LiveRSL____ProposerState__i_Compile.ProposerState = _254_LiveRSL____ProposerState__i_Compile.Type_ProposerState_().Default().(_254_LiveRSL____ProposerState__i_Compile.ProposerState)
	var _ = proposer_k
	var _4927_start__time uint64
	var _ = _4927_start__time
	var _out208 uint64
	var _ = _out208
	_out208 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_4927_start__time = _out208
	var _4928_election__state_k _251_LiveRSL____ElectionState__i_Compile.CElectionState
	var _ = _4928_election__state_k
	var _out209 _251_LiveRSL____ElectionState__i_Compile.CElectionState
	var _ = _out209
	_out209 = _297_LiveRSL____ElectionModel__i_Compile.Companion_Default___.ElectionCheckForQuorumOfViewSuspicions((proposer).Dtor_election__state(), clock, cur__req__set, prev__req__set)
	_4928_election__state_k = _out209
	var _4929_current__state_k uint8
	var _ = _4929_current__state_k
	_4929_current__state_k = uint8(0)
	var _4930_request__queue_k _dafny.Seq
	var _ = _4930_request__queue_k
	_4930_request__queue_k = _dafny.SeqOf()
	var _4931_lt bool
	var _ = _4931_lt
	var _out210 bool
	var _ = _out210
	_out210 = _214_LiveRSL____CTypes__i_Compile.Companion_Default___.CBalLt(((proposer).Dtor_election__state()).Dtor_current__view(), (_4928_election__state_k).Dtor_current__view())
	_4931_lt = _out210
	if _4931_lt {
		_4929_current__state_k = uint8(0)
		_4930_request__queue_k = _dafny.SeqOf()
	} else {
		_4929_current__state_k = (proposer).Dtor_current__state()
		_4930_request__queue_k = (proposer).Dtor_request__queue()
	}
	proposer_k = func(_pat_let66_0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
		return func(_4932_dt__update__tmp_h0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
			return func(_pat_let67_0 _dafny.Seq) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
				return func(_4933_dt__update_hrequest__queue_h0 _dafny.Seq) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
					return func(_pat_let68_0 uint8) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
						return func(_4934_dt__update_hcurrent__state_h0 uint8) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
							return func(_pat_let69_0 _251_LiveRSL____ElectionState__i_Compile.CElectionState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
								return func(_4935_dt__update_helection__state_h0 _251_LiveRSL____ElectionState__i_Compile.CElectionState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
									return _254_LiveRSL____ProposerState__i_Compile.ProposerState{_254_LiveRSL____ProposerState__i_Compile.ProposerState_ProposerState{(_4932_dt__update__tmp_h0).Dtor_constants(), _4934_dt__update_hcurrent__state_h0, _4933_dt__update_hrequest__queue_h0, (_4932_dt__update__tmp_h0).Dtor_max__ballot__i__sent__1a(), (_4932_dt__update__tmp_h0).Dtor_next__operation__number__to__propose(), (_4932_dt__update__tmp_h0).Dtor_received__1b__packets(), (_4932_dt__update__tmp_h0).Dtor_highest__seqno__requested__by__client__this__view(), (_4932_dt__update__tmp_h0).Dtor_incomplete__batch__timer(), _4935_dt__update_helection__state_h0, (_4932_dt__update__tmp_h0).Dtor_maxOpnWithProposal(), (_4932_dt__update__tmp_h0).Dtor_maxLogTruncationPoint()}}
								}(_pat_let69_0)
							}(_4928_election__state_k)
						}(_pat_let68_0)
					}(_4929_current__state_k)
				}(_pat_let67_0)
			}(_4930_request__queue_k)
		}(_pat_let66_0)
	}(proposer)
	var _4936_end__time uint64
	var _ = _4936_end__time
	var _out211 uint64
	var _ = _out211
	_out211 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_4936_end__time = _out211
	if _4931_lt {
		_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("ProposerCheckForQuorumOfViewSuspicions_changed"), _4927_start__time, _4936_end__time)
		noop = false
	} else {
		_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("ProposerCheckForQuorumOfViewSuspicions_nada"), _4927_start__time, _4936_end__time)
		noop = true
	}
	return proposer_k, noop
}
func (_this *CompanionStruct_Default___) ProposerResetViewTimerDueToExecution(proposer _254_LiveRSL____ProposerState__i_Compile.ProposerState, val _dafny.Seq, cur__req__set *_9_Native____Io__s_Compile.MutableSet, prev__req__set *_9_Native____Io__s_Compile.MutableSet) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var proposer_k _254_LiveRSL____ProposerState__i_Compile.ProposerState = _254_LiveRSL____ProposerState__i_Compile.Type_ProposerState_().Default().(_254_LiveRSL____ProposerState__i_Compile.ProposerState)
	var _ = proposer_k
	var _4937_election__state_k _251_LiveRSL____ElectionState__i_Compile.CElectionState
	var _ = _4937_election__state_k
	var _out212 _251_LiveRSL____ElectionState__i_Compile.CElectionState
	var _ = _out212
	_out212 = _297_LiveRSL____ElectionModel__i_Compile.Companion_Default___.ElectionReflectExecutedRequestBatch((proposer).Dtor_election__state(), val, cur__req__set, prev__req__set)
	_4937_election__state_k = _out212
	proposer_k = func(_pat_let70_0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
		return func(_4938_dt__update__tmp_h0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
			return func(_pat_let71_0 _251_LiveRSL____ElectionState__i_Compile.CElectionState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
				return func(_4939_dt__update_helection__state_h0 _251_LiveRSL____ElectionState__i_Compile.CElectionState) _254_LiveRSL____ProposerState__i_Compile.ProposerState {
					return _254_LiveRSL____ProposerState__i_Compile.ProposerState{_254_LiveRSL____ProposerState__i_Compile.ProposerState_ProposerState{(_4938_dt__update__tmp_h0).Dtor_constants(), (_4938_dt__update__tmp_h0).Dtor_current__state(), (_4938_dt__update__tmp_h0).Dtor_request__queue(), (_4938_dt__update__tmp_h0).Dtor_max__ballot__i__sent__1a(), (_4938_dt__update__tmp_h0).Dtor_next__operation__number__to__propose(), (_4938_dt__update__tmp_h0).Dtor_received__1b__packets(), (_4938_dt__update__tmp_h0).Dtor_highest__seqno__requested__by__client__this__view(), (_4938_dt__update__tmp_h0).Dtor_incomplete__batch__timer(), _4939_dt__update_helection__state_h0, (_4938_dt__update__tmp_h0).Dtor_maxOpnWithProposal(), (_4938_dt__update__tmp_h0).Dtor_maxLogTruncationPoint()}}
				}(_pat_let71_0)
			}(_4937_election__state_k)
		}(_pat_let70_0)
	}(proposer)
	return proposer_k
}

// End of class Default__
