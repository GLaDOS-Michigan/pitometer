// Package _278_LiveRSL____LearnerState__i_Compile
// Dafny module _278_LiveRSL____LearnerState__i_Compile compiled into Go

package _278_LiveRSL____LearnerState__i_Compile

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

type Dummy__ struct{}





// Definition of data type CLearnerTuple
type CLearnerTuple struct {
  Data_CLearnerTuple_
}

func (_this CLearnerTuple) Get() Data_CLearnerTuple_ {
  return _this.Data_CLearnerTuple_
}

type Data_CLearnerTuple_ interface {
  isCLearnerTuple()
}

type CompanionStruct_CLearnerTuple_ struct {}
var Companion_CLearnerTuple_ = CompanionStruct_CLearnerTuple_{}

type CLearnerTuple_CLearnerTuple struct {
  Received__2b__message__senders _dafny.Seq
Candidate__learned__value _dafny.Seq
}

func (CLearnerTuple_CLearnerTuple) isCLearnerTuple() {}

func (CompanionStruct_CLearnerTuple_) Create_CLearnerTuple_(Received__2b__message__senders _dafny.Seq, Candidate__learned__value _dafny.Seq) CLearnerTuple {
  return CLearnerTuple{CLearnerTuple_CLearnerTuple{Received__2b__message__senders,Candidate__learned__value}}
}

func (_this CLearnerTuple) Is_CLearnerTuple() bool {
  _, ok := _this.Get().(CLearnerTuple_CLearnerTuple)
return ok
}

func (_this CLearnerTuple) Dtor_received__2b__message__senders() _dafny.Seq {
  return _this.Get().(CLearnerTuple_CLearnerTuple).Received__2b__message__senders
}

func (_this CLearnerTuple) Dtor_candidate__learned__value() _dafny.Seq {
  return _this.Get().(CLearnerTuple_CLearnerTuple).Candidate__learned__value
}

func (_this CLearnerTuple) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case CLearnerTuple_CLearnerTuple: {
      return "_278_LiveRSL____LearnerState__i_Compile.CLearnerTuple.CLearnerTuple" + "(" + _dafny.String(data.Received__2b__message__senders) + ", " + _dafny.String(data.Candidate__learned__value) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this CLearnerTuple) Equals(other CLearnerTuple) bool {
  switch data1 := _this.Get().(type) {
    case CLearnerTuple_CLearnerTuple: {
      data2, ok := other.Get().(CLearnerTuple_CLearnerTuple)
return ok && data1.Received__2b__message__senders.Equals(data2.Received__2b__message__senders) && data1.Candidate__learned__value.Equals(data2.Candidate__learned__value)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this CLearnerTuple) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(CLearnerTuple)
return ok && _this.Equals(typed)
}
func Type_CLearnerTuple_() _dafny.Type {
  return type_CLearnerTuple_{}
}

type type_CLearnerTuple_ struct {
}

func (_this type_CLearnerTuple_) Default() interface{} {
  return CLearnerTuple{CLearnerTuple_CLearnerTuple{_dafny.EmptySeq, _dafny.EmptySeq}}
}

func (_this type_CLearnerTuple_) String() string {
  return "_278_LiveRSL____LearnerState__i_Compile.CLearnerTuple"
}
// End of data type CLearnerTuple

// Definition of data type CLearnerState
type CLearnerState struct {
  Data_CLearnerState_
}

func (_this CLearnerState) Get() Data_CLearnerState_ {
  return _this.Data_CLearnerState_
}

type Data_CLearnerState_ interface {
  isCLearnerState()
}

type CompanionStruct_CLearnerState_ struct {}
var Companion_CLearnerState_ = CompanionStruct_CLearnerState_{}

type CLearnerState_CLearnerState struct {
  Rcs _245_LiveRSL____ReplicaConstantsState__i_Compile.ReplicaConstantsState
Max__ballot__seen _214_LiveRSL____CTypes__i_Compile.CBallot
Unexecuted__ops _dafny.Map
SendDecision bool
Opn _214_LiveRSL____CTypes__i_Compile.COperationNumber
Recv2b bool
RecvCp _217_LiveRSL____CMessage__i_Compile.CPacket
}

func (CLearnerState_CLearnerState) isCLearnerState() {}

func (CompanionStruct_CLearnerState_) Create_CLearnerState_(Rcs _245_LiveRSL____ReplicaConstantsState__i_Compile.ReplicaConstantsState, Max__ballot__seen _214_LiveRSL____CTypes__i_Compile.CBallot, Unexecuted__ops _dafny.Map, SendDecision bool, Opn _214_LiveRSL____CTypes__i_Compile.COperationNumber, Recv2b bool, RecvCp _217_LiveRSL____CMessage__i_Compile.CPacket) CLearnerState {
  return CLearnerState{CLearnerState_CLearnerState{Rcs,Max__ballot__seen,Unexecuted__ops,SendDecision,Opn,Recv2b,RecvCp}}
}

func (_this CLearnerState) Is_CLearnerState() bool {
  _, ok := _this.Get().(CLearnerState_CLearnerState)
return ok
}

func (_this CLearnerState) Dtor_rcs() _245_LiveRSL____ReplicaConstantsState__i_Compile.ReplicaConstantsState {
  return _this.Get().(CLearnerState_CLearnerState).Rcs
}

func (_this CLearnerState) Dtor_max__ballot__seen() _214_LiveRSL____CTypes__i_Compile.CBallot {
  return _this.Get().(CLearnerState_CLearnerState).Max__ballot__seen
}

func (_this CLearnerState) Dtor_unexecuted__ops() _dafny.Map {
  return _this.Get().(CLearnerState_CLearnerState).Unexecuted__ops
}

func (_this CLearnerState) Dtor_sendDecision() bool {
  return _this.Get().(CLearnerState_CLearnerState).SendDecision
}

func (_this CLearnerState) Dtor_opn() _214_LiveRSL____CTypes__i_Compile.COperationNumber {
  return _this.Get().(CLearnerState_CLearnerState).Opn
}

func (_this CLearnerState) Dtor_recv2b() bool {
  return _this.Get().(CLearnerState_CLearnerState).Recv2b
}

func (_this CLearnerState) Dtor_recvCp() _217_LiveRSL____CMessage__i_Compile.CPacket {
  return _this.Get().(CLearnerState_CLearnerState).RecvCp
}

func (_this CLearnerState) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case CLearnerState_CLearnerState: {
      return "_278_LiveRSL____LearnerState__i_Compile.CLearnerState.CLearnerState" + "(" + _dafny.String(data.Rcs) + ", " + _dafny.String(data.Max__ballot__seen) + ", " + _dafny.String(data.Unexecuted__ops) + ", " + _dafny.String(data.SendDecision) + ", " + _dafny.String(data.Opn) + ", " + _dafny.String(data.Recv2b) + ", " + _dafny.String(data.RecvCp) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this CLearnerState) Equals(other CLearnerState) bool {
  switch data1 := _this.Get().(type) {
    case CLearnerState_CLearnerState: {
      data2, ok := other.Get().(CLearnerState_CLearnerState)
return ok && data1.Rcs.Equals(data2.Rcs) && data1.Max__ballot__seen.Equals(data2.Max__ballot__seen) && data1.Unexecuted__ops.Equals(data2.Unexecuted__ops) && data1.SendDecision == data2.SendDecision && data1.Opn.Equals(data2.Opn) && data1.Recv2b == data2.Recv2b && data1.RecvCp.Equals(data2.RecvCp)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this CLearnerState) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(CLearnerState)
return ok && _this.Equals(typed)
}
func Type_CLearnerState_() _dafny.Type {
  return type_CLearnerState_{}
}

type type_CLearnerState_ struct {
}

func (_this type_CLearnerState_) Default() interface{} {
  return CLearnerState{CLearnerState_CLearnerState{_245_LiveRSL____ReplicaConstantsState__i_Compile.Type_ReplicaConstantsState_().Default().(_245_LiveRSL____ReplicaConstantsState__i_Compile.ReplicaConstantsState), _214_LiveRSL____CTypes__i_Compile.Type_CBallot_().Default().(_214_LiveRSL____CTypes__i_Compile.CBallot), _dafny.EmptyMap, false, _214_LiveRSL____CTypes__i_Compile.Type_COperationNumber_().Default().(_214_LiveRSL____CTypes__i_Compile.COperationNumber), false, _217_LiveRSL____CMessage__i_Compile.Type_CPacket_().Default().(_217_LiveRSL____CMessage__i_Compile.CPacket)}}
}

func (_this type_CLearnerState_) String() string {
  return "_278_LiveRSL____LearnerState__i_Compile.CLearnerState"
}
// End of data type CLearnerState

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
  return "_278_LiveRSL____LearnerState__i_Compile.Default__"
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
  return "_278_LiveRSL____LearnerState__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) LearnerState__Init(rcs _245_LiveRSL____ReplicaConstantsState__i_Compile.ReplicaConstantsState) _278_LiveRSL____LearnerState__i_Compile.CLearnerState {
  goto TAIL_CALL_START
TAIL_CALL_START:
var learner _278_LiveRSL____LearnerState__i_Compile.CLearnerState = Type_CLearnerState_().Default().(_278_LiveRSL____LearnerState__i_Compile.CLearnerState)
  var _ = learner
  var _4741_endPoint _9_Native____Io__s_Compile.EndPoint
  var _ = _4741_endPoint
  _4741_endPoint = ((((rcs).Dtor_all()).Dtor_config()).Dtor_replica__ids()).Index((rcs).Dtor_my__index()).(_9_Native____Io__s_Compile.EndPoint)
  var _4742_unknown _197_LiveRSL____AppInterface__i_Compile.CAppMessage = _197_LiveRSL____AppInterface__i_Compile.Type_CAppMessage_().Default().(_197_LiveRSL____AppInterface__i_Compile.CAppMessage)
  var _ = _4742_unknown
  learner = CLearnerState{CLearnerState_CLearnerState{rcs, _214_LiveRSL____CTypes__i_Compile.CBallot{_214_LiveRSL____CTypes__i_Compile.CBallot_CBallot{uint64(0), uint64(0)}}, _dafny.NewMapBuilder().ToMap(), false, _214_LiveRSL____CTypes__i_Compile.COperationNumber{_214_LiveRSL____CTypes__i_Compile.COperationNumber_COperationNumber{uint64(0)}}, false, _217_LiveRSL____CMessage__i_Compile.CPacket{_217_LiveRSL____CMessage__i_Compile.CPacket_CPacket{_4741_endPoint, _4741_endPoint, _217_LiveRSL____CMessage__i_Compile.CMessage{_217_LiveRSL____CMessage__i_Compile.CMessage_CMessage__2b{_214_LiveRSL____CTypes__i_Compile.CBallot{_214_LiveRSL____CTypes__i_Compile.CBallot_CBallot{uint64(0), uint64(0)}}, _214_LiveRSL____CTypes__i_Compile.COperationNumber{_214_LiveRSL____CTypes__i_Compile.COperationNumber_COperationNumber{uint64(0)}}, _dafny.SeqOf()}}}}}}
  { }
  { }
  return learner
}
// End of class Default__
