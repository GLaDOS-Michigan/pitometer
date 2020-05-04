// Package _297_LiveRSL____ElectionModel__i_Compile
// Dafny module _297_LiveRSL____ElectionModel__i_Compile compiled into Go

package _297_LiveRSL____ElectionModel__i_Compile

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
  return "_297_LiveRSL____ElectionModel__i_Compile.Default__"
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
  return "_297_LiveRSL____ElectionModel__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) CComputeSuccessorView(cb _214_LiveRSL____CTypes__i_Compile.CBallot, constants _240_LiveRSL____ConstantsState__i_Compile.ConstantsState) _214_LiveRSL____CTypes__i_Compile.CBallot {
  var cb_k _214_LiveRSL____CTypes__i_Compile.CBallot = _214_LiveRSL____CTypes__i_Compile.Type_CBallot_().Default().(_214_LiveRSL____CTypes__i_Compile.CBallot)
  var _ = cb_k
  { }
  if (((cb).Dtor_proposer__id()) < ((uint64((((constants).Dtor_config()).Dtor_replica__ids()).CardinalityInt())) - (func () uint64 { return  (uint64(1)) })())) {
    cb_k = _214_LiveRSL____CTypes__i_Compile.CBallot{_214_LiveRSL____CTypes__i_Compile.CBallot_CBallot{(cb).Dtor_seqno(), ((cb).Dtor_proposer__id()) + (uint64(1))}}
  } else {
    cb_k = _214_LiveRSL____CTypes__i_Compile.CBallot{_214_LiveRSL____CTypes__i_Compile.CBallot_CBallot{((cb).Dtor_seqno()) + (uint64(1)), uint64(0)}}
  }
  return cb_k
}
func (_this *CompanionStruct_Default___) BoundCRequestSequence(s _dafny.Seq, lengthBound uint64) _System.Tuple2 {
  if (((uint64(0)) <= (lengthBound)) && ((lengthBound) < (uint64((s).CardinalityInt())))) {
    return _dafny.TupleOf(true, (s).Subseq(_dafny.NilInt, lengthBound))
  } else  {
    return _dafny.TupleOf(false, s)
  }
}
func (_this *CompanionStruct_Default___) BoundCRequestHeaders(s _dafny.Seq, lengthBound uint64, cur__req__set *_9_Native____Io__s_Compile.MutableSet) {
  goto TAIL_CALL_START
TAIL_CALL_START:
  var _4743_i _dafny.Int
  var _ = _4743_i
  _4743_i = _dafny.Zero
  { }
  (cur__req__set).RemoveAll()
  for (_4743_i).Cmp(_dafny.IntOfUint64(lengthBound)) < 0 {
    var _4744_new__header _251_LiveRSL____ElectionState__i_Compile.CRequestHeader
    var _ = _4744_new__header
    _4744_new__header = _251_LiveRSL____ElectionState__i_Compile.CRequestHeader{_251_LiveRSL____ElectionState__i_Compile.CRequestHeader_CRequestHeader{((s).Index(_4743_i).(_214_LiveRSL____CTypes__i_Compile.CRequest)).Dtor_client(), ((s).Index(_4743_i).(_214_LiveRSL____CTypes__i_Compile.CRequest)).Dtor_seqno()}}
    { }
    { }
    (cur__req__set).Add(_4744_new__header)
    _4743_i = (_4743_i).Plus(_dafny.IntOfInt64(1))
  }
  return
}
func (_this *CompanionStruct_Default___) RemoveAllSatisfiedCRequestsInSequenceIter(requests _dafny.Seq, cur__req__set *_9_Native____Io__s_Compile.MutableSet, r _214_LiveRSL____CTypes__i_Compile.CRequest) _dafny.Seq {
  var requests_k _dafny.Seq = _dafny.EmptySeq
  var _ = requests_k
  var _4745_i uint64
  var _ = _4745_i
  _4745_i = uint64(0)
  var _4746_len uint64
  var _ = _4746_len
  _4746_len = uint64((requests).CardinalityInt())
  { }
  { }
  requests_k = _dafny.SeqOf()
  _4745_i = uint64(0)
  { }
  for (_4745_i) < (_4746_len) {
    { }
    { }
    { }
    var _4747_h _251_LiveRSL____ElectionState__i_Compile.CRequestHeader
    var _ = _4747_h
    _4747_h = _251_LiveRSL____ElectionState__i_Compile.CRequestHeader{_251_LiveRSL____ElectionState__i_Compile.CRequestHeader_CRequestHeader{((requests).Index(_4745_i).(_214_LiveRSL____CTypes__i_Compile.CRequest)).Dtor_client(), ((requests).Index(_4745_i).(_214_LiveRSL____CTypes__i_Compile.CRequest)).Dtor_seqno()}}
    { }
    if (_251_LiveRSL____ElectionState__i_Compile.Companion_Default___.CRequestSatisfiedBy((requests).Index(_4745_i).(_214_LiveRSL____CTypes__i_Compile.CRequest), r)) {
      (cur__req__set).Remove(_4747_h)
      { }
    } else {
      { }
      requests_k = (requests_k).Concat(_dafny.SeqOf((requests).Index(_4745_i).(_214_LiveRSL____CTypes__i_Compile.CRequest)))
      { }
    }
    { }
    _4745_i = (_4745_i) + (uint64(1))
  }
  { }
  { }
  return requests_k
}
func (_this *CompanionStruct_Default___) InitElectionState(constants _245_LiveRSL____ReplicaConstantsState__i_Compile.ReplicaConstantsState) (_251_LiveRSL____ElectionState__i_Compile.CElectionState, *_9_Native____Io__s_Compile.MutableSet, *_9_Native____Io__s_Compile.MutableSet) {
  var election _251_LiveRSL____ElectionState__i_Compile.CElectionState = _251_LiveRSL____ElectionState__i_Compile.Type_CElectionState_().Default().(_251_LiveRSL____ElectionState__i_Compile.CElectionState)
  var _ = election
var cur__req__set *_9_Native____Io__s_Compile.MutableSet = (*_9_Native____Io__s_Compile.MutableSet)(nil)
  var _ = cur__req__set
var prev__req__set *_9_Native____Io__s_Compile.MutableSet = (*_9_Native____Io__s_Compile.MutableSet)(nil)
  var _ = prev__req__set
  election = _251_LiveRSL____ElectionState__i_Compile.CElectionState{_251_LiveRSL____ElectionState__i_Compile.CElectionState_CElectionState{constants, _214_LiveRSL____CTypes__i_Compile.CBallot{_214_LiveRSL____CTypes__i_Compile.CBallot_CBallot{uint64(1), uint64(0)}}, _dafny.SeqOf(), uint64(0), (((constants).Dtor_all()).Dtor_params()).Dtor_baseline__view__timeout__period(), _dafny.SeqOf(), _dafny.SeqOf()}}
  var _out141 *_9_Native____Io__s_Compile.MutableSet
  var _ = _out141
_out141 = _9_Native____Io__s_Compile.Companion_MutableSet_.EmptySet()
cur__req__set = _out141
  var _out142 *_9_Native____Io__s_Compile.MutableSet
  var _ = _out142
_out142 = _9_Native____Io__s_Compile.Companion_MutableSet_.EmptySet()
prev__req__set = _out142
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  return election,cur__req__set,prev__req__set
}
func (_this *CompanionStruct_Default___) ElectionProcessHeartbeat(ces _251_LiveRSL____ElectionState__i_Compile.CElectionState, cp _217_LiveRSL____CMessage__i_Compile.CPacket, clock uint64, cur__req__set *_9_Native____Io__s_Compile.MutableSet, prev__req__set *_9_Native____Io__s_Compile.MutableSet) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
  var ces_k _251_LiveRSL____ElectionState__i_Compile.CElectionState = _251_LiveRSL____ElectionState__i_Compile.Type_CElectionState_().Default().(_251_LiveRSL____ElectionState__i_Compile.CElectionState)
  var _ = ces_k
  var _4748_src__ep _9_Native____Io__s_Compile.EndPoint
  var _ = _4748_src__ep
  _4748_src__ep = (cp).Dtor_src()
  var _4749_found bool
  var _ = _4749_found
var _4750_index uint64
  var _ = _4750_index
var _out143 bool
  var _ = _out143
var _out144 uint64
  var _ = _out144
_out143,_out144 = _238_LiveRSL____CPaxosConfiguration__i_Compile.Companion_Default___.CGetReplicaIndex(_4748_src__ep, (((ces).Dtor_constants()).Dtor_all()).Dtor_config())
_4749_found = _out143
_4750_index = _out144
  { }
  { }
  { }
  { }
  if (!(_4749_found)) {
    { }
    ces_k = ces
  } else {
    { }
    { }
    if (((((cp).Dtor_msg()).Dtor_bal__heartbeat()).Equals((ces).Dtor_current__view())) && (((cp).Dtor_msg()).Dtor_suspicious())) {
      { }
      ces_k = func (_pat_let2_0 _251_LiveRSL____ElectionState__i_Compile.CElectionState) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
        return func (_4751_dt__update__tmp_h1 _251_LiveRSL____ElectionState__i_Compile.CElectionState) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
          return func (_pat_let3_0 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
            return func (_4752_dt__update_hcurrent__view__suspectors_h1 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
              return _251_LiveRSL____ElectionState__i_Compile.CElectionState{_251_LiveRSL____ElectionState__i_Compile.CElectionState_CElectionState{(_4751_dt__update__tmp_h1).Dtor_constants(), (_4751_dt__update__tmp_h1).Dtor_current__view(), _4752_dt__update_hcurrent__view__suspectors_h1, (_4751_dt__update__tmp_h1).Dtor_epoch__end__time(), (_4751_dt__update__tmp_h1).Dtor_epoch__length(), (_4751_dt__update__tmp_h1).Dtor_requests__received__this__epoch(), (_4751_dt__update__tmp_h1).Dtor_requests__received__prev__epochs()}}
            }(_pat_let3_0)
          }(_185_Common____SeqIsUnique__i_Compile.Companion_Default___.AppendToUniqueSeqMaybe((ces).Dtor_current__view__suspectors(), _4750_index))
        }(_pat_let2_0)
      }(ces)
      { }
      { }
    } else {
      var _4753_cmp bool
      var _ = _4753_cmp
var _out145 bool
      var _ = _out145
_out145 = _214_LiveRSL____CTypes__i_Compile.Companion_Default___.CBalLt((ces).Dtor_current__view(), ((cp).Dtor_msg()).Dtor_bal__heartbeat())
_4753_cmp = _out145
      if (_4753_cmp) {
        { }
        { }
        var _4754_cnewEpochLength uint64
        var _ = _4754_cnewEpochLength
var _out146 uint64
        var _ = _out146
_out146 = _234_Common____UpperBound__i_Compile.Companion_Default___.UpperBoundedAdditionImpl((ces).Dtor_epoch__length(), (ces).Dtor_epoch__length(), ((((ces).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__integer__val())
_4754_cnewEpochLength = _out146
        var _4755_cnewEpochEndTime uint64
        var _ = _4755_cnewEpochEndTime
var _out147 uint64
        var _ = _out147
_out147 = _234_Common____UpperBound__i_Compile.Companion_Default___.UpperBoundedAdditionImpl(clock, _4754_cnewEpochLength, ((((ces).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__integer__val())
_4755_cnewEpochEndTime = _out147
        var _4756_new__seq _dafny.Seq
        var _ = _4756_new__seq
        _4756_new__seq = ((ces).Dtor_requests__received__prev__epochs()).Concat((ces).Dtor_requests__received__this__epoch())
        { }
        (prev__req__set).AddSet(cur__req__set)
        (cur__req__set).RemoveAll()
        { }
        var _4757_tuple _System.Tuple2
        var _ = _4757_tuple
        _4757_tuple = Companion_Default___.BoundCRequestSequence(_4756_new__seq, ((((ces).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__integer__val())
        var _4758_bounded bool
        var _ = _4758_bounded
var _4759_bounded__seq _dafny.Seq
        var _ = _4759_bounded__seq
        _4758_bounded, _4759_bounded__seq = (*((_4757_tuple)).IndexInt(0)).(bool), (*((_4757_tuple)).IndexInt(1)).(_dafny.Seq)
        if (_4758_bounded) {
          Companion_Default___.BoundCRequestHeaders(_4756_new__seq, ((((ces).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__integer__val(), prev__req__set)
        }
        ces_k = func (_pat_let4_0 _251_LiveRSL____ElectionState__i_Compile.CElectionState) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
          return func (_4760_dt__update__tmp_h3 _251_LiveRSL____ElectionState__i_Compile.CElectionState) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
            return func (_pat_let5_0 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
              return func (_4761_dt__update_hrequests__received__this__epoch_h1 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                return func (_pat_let6_0 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                  return func (_4762_dt__update_hrequests__received__prev__epochs_h1 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                    return func (_pat_let7_0 uint64) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                      return func (_4763_dt__update_hepoch__end__time_h1 uint64) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                        return func (_pat_let8_0 uint64) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                          return func (_4764_dt__update_hepoch__length_h1 uint64) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                            return func (_pat_let9_0 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                              return func (_4765_dt__update_hcurrent__view__suspectors_h3 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                                return func (_pat_let10_0 _214_LiveRSL____CTypes__i_Compile.CBallot) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                                  return func (_4766_dt__update_hcurrent__view_h1 _214_LiveRSL____CTypes__i_Compile.CBallot) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                                    return _251_LiveRSL____ElectionState__i_Compile.CElectionState{_251_LiveRSL____ElectionState__i_Compile.CElectionState_CElectionState{(_4760_dt__update__tmp_h3).Dtor_constants(), _4766_dt__update_hcurrent__view_h1, _4765_dt__update_hcurrent__view__suspectors_h3, _4763_dt__update_hepoch__end__time_h1, _4764_dt__update_hepoch__length_h1, _4761_dt__update_hrequests__received__this__epoch_h1, _4762_dt__update_hrequests__received__prev__epochs_h1}}
                                  }(_pat_let10_0)
                                }(((cp).Dtor_msg()).Dtor_bal__heartbeat())
                              }(_pat_let9_0)
                            }((func () _dafny.Seq { if ((cp).Dtor_msg()).Dtor_suspicious() { return _dafny.SeqOf(_4750_index) }; return _dafny.SeqOf() })() )
                          }(_pat_let8_0)
                        }(_4754_cnewEpochLength)
                      }(_pat_let7_0)
                    }(_4755_cnewEpochEndTime)
                  }(_pat_let6_0)
                }(_4759_bounded__seq)
              }(_pat_let5_0)
            }(_dafny.SeqOf())
          }(_pat_let4_0)
        }(ces)
        { }
        { }
        { }
        { }
        { }
      } else {
        { }
        ces_k = ces
      }
    }
    { }
  }
  { }
  return ces_k
}
func (_this *CompanionStruct_Default___) ElectionCheckForViewTimeout(ces _251_LiveRSL____ElectionState__i_Compile.CElectionState, clock uint64, cur__req__set *_9_Native____Io__s_Compile.MutableSet, prev__req__set *_9_Native____Io__s_Compile.MutableSet) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
  var ces_k _251_LiveRSL____ElectionState__i_Compile.CElectionState = _251_LiveRSL____ElectionState__i_Compile.Type_CElectionState_().Default().(_251_LiveRSL____ElectionState__i_Compile.CElectionState)
  var _ = ces_k
  var _4767_start__time uint64
  var _ = _4767_start__time
var _out148 uint64
  var _ = _out148
_out148 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_4767_start__time = _out148
  { }
  { }
  { }
  if ((clock) < ((ces).Dtor_epoch__end__time())) {
    { }
    ces_k = ces
    var _4768_end__time uint64
    var _ = _4768_end__time
var _out149 uint64
    var _ = _out149
_out149 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_4768_end__time = _out149
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("ElectionCheckForViewTimeout_nada"), _4767_start__time, _4768_end__time)
    { }
  } else if ((((ces).Dtor_requests__received__prev__epochs()).Cardinality()).Cmp(_dafny.Zero) == 0) {
    { }
    { }
    var _4769_cnewEpochLength uint64
    var _ = _4769_cnewEpochLength
    _4769_cnewEpochLength = ((((ces).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_baseline__view__timeout__period()
    var _4770_cnewEpochEndTime uint64
    var _ = _4770_cnewEpochEndTime
var _out150 uint64
    var _ = _out150
_out150 = _234_Common____UpperBound__i_Compile.Companion_Default___.UpperBoundedAdditionImpl(clock, _4769_cnewEpochLength, ((((ces).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__integer__val())
_4770_cnewEpochEndTime = _out150
    ces_k = func (_pat_let11_0 _251_LiveRSL____ElectionState__i_Compile.CElectionState) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
      return func (_4771_dt__update__tmp_h1 _251_LiveRSL____ElectionState__i_Compile.CElectionState) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
        return func (_pat_let12_0 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
          return func (_4772_dt__update_hrequests__received__this__epoch_h1 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
            return func (_pat_let13_0 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
              return func (_4773_dt__update_hrequests__received__prev__epochs_h1 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                return func (_pat_let14_0 uint64) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                  return func (_4774_dt__update_hepoch__end__time_h1 uint64) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                    return func (_pat_let15_0 uint64) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                      return func (_4775_dt__update_hepoch__length_h1 uint64) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                        return _251_LiveRSL____ElectionState__i_Compile.CElectionState{_251_LiveRSL____ElectionState__i_Compile.CElectionState_CElectionState{(_4771_dt__update__tmp_h1).Dtor_constants(), (_4771_dt__update__tmp_h1).Dtor_current__view(), (_4771_dt__update__tmp_h1).Dtor_current__view__suspectors(), _4774_dt__update_hepoch__end__time_h1, _4775_dt__update_hepoch__length_h1, _4772_dt__update_hrequests__received__this__epoch_h1, _4773_dt__update_hrequests__received__prev__epochs_h1}}
                      }(_pat_let15_0)
                    }(_4769_cnewEpochLength)
                  }(_pat_let14_0)
                }(_4770_cnewEpochEndTime)
              }(_pat_let13_0)
            }((ces).Dtor_requests__received__this__epoch())
          }(_pat_let12_0)
        }(_dafny.SeqOf())
      }(_pat_let11_0)
    }(ces)
    (prev__req__set).TransferSet(cur__req__set)
    { }
    var _4776_end__time uint64
    var _ = _4776_end__time
var _out151 uint64
    var _ = _out151
_out151 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_4776_end__time = _out151
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("ElectionCheckForViewTimeout_noprev"), _4767_start__time, _4776_end__time)
  } else {
    { }
    var _4777_cnewEpochEndTime uint64
    var _ = _4777_cnewEpochEndTime
var _out152 uint64
    var _ = _out152
_out152 = _234_Common____UpperBound__i_Compile.Companion_Default___.UpperBoundedAdditionImpl(clock, (ces).Dtor_epoch__length(), ((((ces).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__integer__val())
_4777_cnewEpochEndTime = _out152
    var _4778_new__seq _dafny.Seq
    var _ = _4778_new__seq
    _4778_new__seq = ((ces).Dtor_requests__received__prev__epochs()).Concat((ces).Dtor_requests__received__this__epoch())
    { }
    (prev__req__set).AddSet(cur__req__set)
    (cur__req__set).RemoveAll()
    { }
    var _4779_tuple _System.Tuple2
    var _ = _4779_tuple
    _4779_tuple = Companion_Default___.BoundCRequestSequence(_4778_new__seq, ((((ces).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__integer__val())
    var _4780_bounded bool
    var _ = _4780_bounded
var _4781_bounded__seq _dafny.Seq
    var _ = _4781_bounded__seq
    _4780_bounded, _4781_bounded__seq = (*((_4779_tuple)).IndexInt(0)).(bool), (*((_4779_tuple)).IndexInt(1)).(_dafny.Seq)
    if (_4780_bounded) {
      Companion_Default___.BoundCRequestHeaders(_4778_new__seq, ((((ces).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__integer__val(), prev__req__set)
    }
    ces_k = func (_pat_let16_0 _251_LiveRSL____ElectionState__i_Compile.CElectionState) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
      return func (_4782_dt__update__tmp_h3 _251_LiveRSL____ElectionState__i_Compile.CElectionState) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
        return func (_pat_let17_0 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
          return func (_4783_dt__update_hrequests__received__this__epoch_h3 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
            return func (_pat_let18_0 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
              return func (_4784_dt__update_hrequests__received__prev__epochs_h3 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                return func (_pat_let19_0 uint64) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                  return func (_4785_dt__update_hepoch__end__time_h3 uint64) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                    return func (_pat_let20_0 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                      return func (_4786_dt__update_hcurrent__view__suspectors_h1 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                        return _251_LiveRSL____ElectionState__i_Compile.CElectionState{_251_LiveRSL____ElectionState__i_Compile.CElectionState_CElectionState{(_4782_dt__update__tmp_h3).Dtor_constants(), (_4782_dt__update__tmp_h3).Dtor_current__view(), _4786_dt__update_hcurrent__view__suspectors_h1, _4785_dt__update_hepoch__end__time_h3, (_4782_dt__update__tmp_h3).Dtor_epoch__length(), _4783_dt__update_hrequests__received__this__epoch_h3, _4784_dt__update_hrequests__received__prev__epochs_h3}}
                      }(_pat_let20_0)
                    }(_185_Common____SeqIsUnique__i_Compile.Companion_Default___.AppendToUniqueSeqMaybe((ces).Dtor_current__view__suspectors(), ((ces).Dtor_constants()).Dtor_my__index()))
                  }(_pat_let19_0)
                }(_4777_cnewEpochEndTime)
              }(_pat_let18_0)
            }(_4781_bounded__seq)
          }(_pat_let17_0)
        }(_dafny.SeqOf())
      }(_pat_let16_0)
    }(ces)
    { }
    { }
    { }
    { }
    var _4787_end__time uint64
    var _ = _4787_end__time
var _out153 uint64
    var _ = _out153
_out153 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_4787_end__time = _out153
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("ElectionCheckForViewTimeout_timeout"), _4767_start__time, _4787_end__time)
  }
  { }
  return ces_k
}
func (_this *CompanionStruct_Default___) ElectionCheckForQuorumOfViewSuspicions(ces _251_LiveRSL____ElectionState__i_Compile.CElectionState, clock uint64, cur__req__set *_9_Native____Io__s_Compile.MutableSet, prev__req__set *_9_Native____Io__s_Compile.MutableSet) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
  var ces_k _251_LiveRSL____ElectionState__i_Compile.CElectionState = _251_LiveRSL____ElectionState__i_Compile.Type_CElectionState_().Default().(_251_LiveRSL____ElectionState__i_Compile.CElectionState)
  var _ = ces_k
  { }
  { }
  { }
  { }
  { }
  var _4788_minq uint64
  var _ = _4788_minq
var _out154 uint64
  var _ = _out154
_out154 = _294_LiveRSL____MinCQuorumSize__i_Compile.Companion_Default___.MinCQuorumSize((((ces).Dtor_constants()).Dtor_all()).Dtor_config())
_4788_minq = _out154
  if (((((ces).Dtor_current__view__suspectors()).Cardinality()).Cmp(_dafny.IntOfUint64(_4788_minq)) < 0) || ((((ces).Dtor_current__view()).Dtor_seqno()) >= (((((ces).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__integer__val()))) {
    { }
    ces_k = ces
  } else {
    { }
    { }
    var _4789_cview _214_LiveRSL____CTypes__i_Compile.CBallot
    var _ = _4789_cview
var _out155 _214_LiveRSL____CTypes__i_Compile.CBallot
    var _ = _out155
_out155 = Companion_Default___.CComputeSuccessorView((ces).Dtor_current__view(), ((ces).Dtor_constants()).Dtor_all())
_4789_cview = _out155
    var _4790_cnewEpochLength uint64
    var _ = _4790_cnewEpochLength
var _out156 uint64
    var _ = _out156
_out156 = _234_Common____UpperBound__i_Compile.Companion_Default___.UpperBoundedAdditionImpl((ces).Dtor_epoch__length(), (ces).Dtor_epoch__length(), ((((ces).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__integer__val())
_4790_cnewEpochLength = _out156
    var _4791_cnewEpochEndTime uint64
    var _ = _4791_cnewEpochEndTime
var _out157 uint64
    var _ = _out157
_out157 = _234_Common____UpperBound__i_Compile.Companion_Default___.UpperBoundedAdditionImpl(clock, _4790_cnewEpochLength, ((((ces).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__integer__val())
_4791_cnewEpochEndTime = _out157
    var _4792_new__seq _dafny.Seq
    var _ = _4792_new__seq
    _4792_new__seq = ((ces).Dtor_requests__received__prev__epochs()).Concat((ces).Dtor_requests__received__this__epoch())
    { }
    (prev__req__set).AddSet(cur__req__set)
    (cur__req__set).RemoveAll()
    { }
    var _4793_tuple _System.Tuple2
    var _ = _4793_tuple
    _4793_tuple = Companion_Default___.BoundCRequestSequence(_4792_new__seq, ((((ces).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__integer__val())
    var _4794_bounded bool
    var _ = _4794_bounded
var _4795_bounded__seq _dafny.Seq
    var _ = _4795_bounded__seq
    _4794_bounded, _4795_bounded__seq = (*((_4793_tuple)).IndexInt(0)).(bool), (*((_4793_tuple)).IndexInt(1)).(_dafny.Seq)
    if (_4794_bounded) {
      Companion_Default___.BoundCRequestHeaders(_4792_new__seq, ((((ces).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__integer__val(), prev__req__set)
    }
    ces_k = func (_pat_let21_0 _251_LiveRSL____ElectionState__i_Compile.CElectionState) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
      return func (_4796_dt__update__tmp_h1 _251_LiveRSL____ElectionState__i_Compile.CElectionState) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
        return func (_pat_let22_0 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
          return func (_4797_dt__update_hrequests__received__this__epoch_h1 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
            return func (_pat_let23_0 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
              return func (_4798_dt__update_hrequests__received__prev__epochs_h1 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                return func (_pat_let24_0 uint64) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                  return func (_4799_dt__update_hepoch__end__time_h1 uint64) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                    return func (_pat_let25_0 uint64) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                      return func (_4800_dt__update_hepoch__length_h1 uint64) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                        return func (_pat_let26_0 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                          return func (_4801_dt__update_hcurrent__view__suspectors_h1 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                            return func (_pat_let27_0 _214_LiveRSL____CTypes__i_Compile.CBallot) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                              return func (_4802_dt__update_hcurrent__view_h1 _214_LiveRSL____CTypes__i_Compile.CBallot) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                                return _251_LiveRSL____ElectionState__i_Compile.CElectionState{_251_LiveRSL____ElectionState__i_Compile.CElectionState_CElectionState{(_4796_dt__update__tmp_h1).Dtor_constants(), _4802_dt__update_hcurrent__view_h1, _4801_dt__update_hcurrent__view__suspectors_h1, _4799_dt__update_hepoch__end__time_h1, _4800_dt__update_hepoch__length_h1, _4797_dt__update_hrequests__received__this__epoch_h1, _4798_dt__update_hrequests__received__prev__epochs_h1}}
                              }(_pat_let27_0)
                            }(_4789_cview)
                          }(_pat_let26_0)
                        }(_dafny.SeqOf())
                      }(_pat_let25_0)
                    }(_4790_cnewEpochLength)
                  }(_pat_let24_0)
                }(_4791_cnewEpochEndTime)
              }(_pat_let23_0)
            }(_4795_bounded__seq)
          }(_pat_let22_0)
        }(_dafny.SeqOf())
      }(_pat_let21_0)
    }(ces)
    { }
    { }
  }
  { }
  { }
  { }
  return ces_k
}
func (_this *CompanionStruct_Default___) FindEarlierRequestSets(ces _251_LiveRSL____ElectionState__i_Compile.CElectionState, target _214_LiveRSL____CTypes__i_Compile.CRequest, cur__req__set *_9_Native____Io__s_Compile.MutableSet, prev__req__set *_9_Native____Io__s_Compile.MutableSet) bool {
  var b bool = false
  var _ = b
  var _4803_header _251_LiveRSL____ElectionState__i_Compile.CRequestHeader
  var _ = _4803_header
  _4803_header = _251_LiveRSL____ElectionState__i_Compile.CRequestHeader{_251_LiveRSL____ElectionState__i_Compile.CRequestHeader_CRequestHeader{(target).Dtor_client(), (target).Dtor_seqno()}}
  var _4804_b1 bool
  var _ = _4804_b1
var _out158 bool
  var _ = _out158
_out158 = (cur__req__set).Contains(_4803_header)
_4804_b1 = _out158
  if (_4804_b1) {
    b = true
  } else {
    var _4805_b2 bool
    var _ = _4805_b2
var _out159 bool
    var _ = _out159
_out159 = (prev__req__set).Contains(_4803_header)
_4805_b2 = _out159
    b = _4805_b2
  }
  { }
  { }
  if (b) {
    { }
    { }
    { }
  }
  return b
}
func (_this *CompanionStruct_Default___) ElectionReflectReceivedRequest(ces _251_LiveRSL____ElectionState__i_Compile.CElectionState, creq _214_LiveRSL____CTypes__i_Compile.CRequest, cur__req__set *_9_Native____Io__s_Compile.MutableSet, prev__req__set *_9_Native____Io__s_Compile.MutableSet) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
  var ces_k _251_LiveRSL____ElectionState__i_Compile.CElectionState = _251_LiveRSL____ElectionState__i_Compile.Type_CElectionState_().Default().(_251_LiveRSL____ElectionState__i_Compile.CElectionState)
  var _ = ces_k
  { }
  { }
  { }
  var _4806_earlier bool
  var _ = _4806_earlier
var _out160 bool
  var _ = _out160
_out160 = Companion_Default___.FindEarlierRequestSets(ces, creq, cur__req__set, prev__req__set)
_4806_earlier = _out160
  if (_4806_earlier) {
    { }
    ces_k = ces
  } else {
    { }
    var _4807_new__seq _dafny.Seq
    var _ = _4807_new__seq
    _4807_new__seq = ((ces).Dtor_requests__received__this__epoch()).Concat(_dafny.SeqOf(creq))
    var _4808_header _251_LiveRSL____ElectionState__i_Compile.CRequestHeader
    var _ = _4808_header
    _4808_header = _251_LiveRSL____ElectionState__i_Compile.CRequestHeader{_251_LiveRSL____ElectionState__i_Compile.CRequestHeader_CRequestHeader{(creq).Dtor_client(), (creq).Dtor_seqno()}}
    { }
    (cur__req__set).Add(_4808_header)
    { }
    { }
    var _4809_tuple _System.Tuple2
    var _ = _4809_tuple
    _4809_tuple = Companion_Default___.BoundCRequestSequence(_4807_new__seq, ((((ces).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__integer__val())
    var _4810_bounded bool
    var _ = _4810_bounded
var _4811_bounded__seq _dafny.Seq
    var _ = _4811_bounded__seq
    _4810_bounded, _4811_bounded__seq = (*((_4809_tuple)).IndexInt(0)).(bool), (*((_4809_tuple)).IndexInt(1)).(_dafny.Seq)
    if (_4810_bounded) {
      Companion_Default___.BoundCRequestHeaders(_4807_new__seq, ((((ces).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__integer__val(), cur__req__set)
    }
    ces_k = func (_pat_let28_0 _251_LiveRSL____ElectionState__i_Compile.CElectionState) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
      return func (_4812_dt__update__tmp_h1 _251_LiveRSL____ElectionState__i_Compile.CElectionState) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
        return func (_pat_let29_0 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
          return func (_4813_dt__update_hrequests__received__this__epoch_h1 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
            return _251_LiveRSL____ElectionState__i_Compile.CElectionState{_251_LiveRSL____ElectionState__i_Compile.CElectionState_CElectionState{(_4812_dt__update__tmp_h1).Dtor_constants(), (_4812_dt__update__tmp_h1).Dtor_current__view(), (_4812_dt__update__tmp_h1).Dtor_current__view__suspectors(), (_4812_dt__update__tmp_h1).Dtor_epoch__end__time(), (_4812_dt__update__tmp_h1).Dtor_epoch__length(), _4813_dt__update_hrequests__received__this__epoch_h1, (_4812_dt__update__tmp_h1).Dtor_requests__received__prev__epochs()}}
          }(_pat_let29_0)
        }(_4811_bounded__seq)
      }(_pat_let28_0)
    }(ces)
    { }
    { }
    { }
    { }
    { }
  }
  return ces_k
}
func (_this *CompanionStruct_Default___) ElectionReflectExecutedRequestBatch(ces _251_LiveRSL____ElectionState__i_Compile.CElectionState, creqb _dafny.Seq, cur__req__set *_9_Native____Io__s_Compile.MutableSet, prev__req__set *_9_Native____Io__s_Compile.MutableSet) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
  var ces_k _251_LiveRSL____ElectionState__i_Compile.CElectionState = _251_LiveRSL____ElectionState__i_Compile.Type_CElectionState_().Default().(_251_LiveRSL____ElectionState__i_Compile.CElectionState)
  var _ = ces_k
  { }
  var _4814_i uint64
  var _ = _4814_i
  _4814_i = uint64(0)
  { }
  var _4815_tempces_k _251_LiveRSL____ElectionState__i_Compile.CElectionState
  var _ = _4815_tempces_k
  _4815_tempces_k = ces
  for (_4814_i) < (uint64((creqb).CardinalityInt())) {
    var _4816_creq _214_LiveRSL____CTypes__i_Compile.CRequest
    var _ = _4816_creq
    _4816_creq = (creqb).Index(_4814_i).(_214_LiveRSL____CTypes__i_Compile.CRequest)
    { }
    { }
    { }
    { }
    { }
    { }
    var _4817_prevEpoch _dafny.Seq = _dafny.EmptySeq
    var _ = _4817_prevEpoch
    { }
    var _out161 _dafny.Seq
    var _ = _out161
_out161 = Companion_Default___.RemoveAllSatisfiedCRequestsInSequenceIter((_4815_tempces_k).Dtor_requests__received__prev__epochs(), prev__req__set, _4816_creq)
_4817_prevEpoch = _out161
    var _4818_thisEpoch _dafny.Seq = _dafny.EmptySeq
    var _ = _4818_thisEpoch
    { }
    var _out162 _dafny.Seq
    var _ = _out162
_out162 = Companion_Default___.RemoveAllSatisfiedCRequestsInSequenceIter((_4815_tempces_k).Dtor_requests__received__this__epoch(), cur__req__set, _4816_creq)
_4818_thisEpoch = _out162
    { }
    _4815_tempces_k = func (_pat_let30_0 _251_LiveRSL____ElectionState__i_Compile.CElectionState) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
      return func (_4819_dt__update__tmp_h1 _251_LiveRSL____ElectionState__i_Compile.CElectionState) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
        return func (_pat_let31_0 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
          return func (_4820_dt__update_hrequests__received__this__epoch_h1 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
            return func (_pat_let32_0 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
              return func (_4821_dt__update_hrequests__received__prev__epochs_h1 _dafny.Seq) _251_LiveRSL____ElectionState__i_Compile.CElectionState {
                return _251_LiveRSL____ElectionState__i_Compile.CElectionState{_251_LiveRSL____ElectionState__i_Compile.CElectionState_CElectionState{(_4819_dt__update__tmp_h1).Dtor_constants(), (_4819_dt__update__tmp_h1).Dtor_current__view(), (_4819_dt__update__tmp_h1).Dtor_current__view__suspectors(), (_4819_dt__update__tmp_h1).Dtor_epoch__end__time(), (_4819_dt__update__tmp_h1).Dtor_epoch__length(), _4820_dt__update_hrequests__received__this__epoch_h1, _4821_dt__update_hrequests__received__prev__epochs_h1}}
              }(_pat_let32_0)
            }(_4817_prevEpoch)
          }(_pat_let31_0)
        }(_4818_thisEpoch)
      }(_pat_let30_0)
    }(_4815_tempces_k)
    { }
    { }
    { }
    { }
    { }
    _4814_i = (_4814_i) + (uint64(1))
  }
  { }
  ces_k = _4815_tempces_k
  return ces_k
}
// End of class Default__
