// Package _285_LiveRSL____ReplicaState__i_Compile
// Dafny module _285_LiveRSL____ReplicaState__i_Compile compiled into Go

package _285_LiveRSL____ReplicaState__i_Compile

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

type Dummy__ struct{}









// Definition of data type ReplicaState
type ReplicaState struct {
  Data_ReplicaState_
}

func (_this ReplicaState) Get() Data_ReplicaState_ {
  return _this.Data_ReplicaState_
}

type Data_ReplicaState_ interface {
  isReplicaState()
}

type CompanionStruct_ReplicaState_ struct {}
var Companion_ReplicaState_ = CompanionStruct_ReplicaState_{}

type ReplicaState_ReplicaState struct {
  Constants _245_LiveRSL____ReplicaConstantsState__i_Compile.ReplicaConstantsState
NextHeartbeatTime uint64
Proposer _254_LiveRSL____ProposerState__i_Compile.ProposerState
Acceptor _269_LiveRSL____AcceptorState__i_Compile.AcceptorState
Learner _278_LiveRSL____LearnerState__i_Compile.CLearnerState
Executor _276_LiveRSL____ExecutorState__i_Compile.ExecutorState
}

func (ReplicaState_ReplicaState) isReplicaState() {}

func (CompanionStruct_ReplicaState_) Create_ReplicaState_(Constants _245_LiveRSL____ReplicaConstantsState__i_Compile.ReplicaConstantsState, NextHeartbeatTime uint64, Proposer _254_LiveRSL____ProposerState__i_Compile.ProposerState, Acceptor _269_LiveRSL____AcceptorState__i_Compile.AcceptorState, Learner _278_LiveRSL____LearnerState__i_Compile.CLearnerState, Executor _276_LiveRSL____ExecutorState__i_Compile.ExecutorState) ReplicaState {
  return ReplicaState{ReplicaState_ReplicaState{Constants,NextHeartbeatTime,Proposer,Acceptor,Learner,Executor}}
}

func (_this ReplicaState) Is_ReplicaState() bool {
  _, ok := _this.Get().(ReplicaState_ReplicaState)
return ok
}

func (_this ReplicaState) Dtor_constants() _245_LiveRSL____ReplicaConstantsState__i_Compile.ReplicaConstantsState {
  return _this.Get().(ReplicaState_ReplicaState).Constants
}

func (_this ReplicaState) Dtor_nextHeartbeatTime() uint64 {
  return _this.Get().(ReplicaState_ReplicaState).NextHeartbeatTime
}

func (_this ReplicaState) Dtor_proposer() _254_LiveRSL____ProposerState__i_Compile.ProposerState {
  return _this.Get().(ReplicaState_ReplicaState).Proposer
}

func (_this ReplicaState) Dtor_acceptor() _269_LiveRSL____AcceptorState__i_Compile.AcceptorState {
  return _this.Get().(ReplicaState_ReplicaState).Acceptor
}

func (_this ReplicaState) Dtor_learner() _278_LiveRSL____LearnerState__i_Compile.CLearnerState {
  return _this.Get().(ReplicaState_ReplicaState).Learner
}

func (_this ReplicaState) Dtor_executor() _276_LiveRSL____ExecutorState__i_Compile.ExecutorState {
  return _this.Get().(ReplicaState_ReplicaState).Executor
}

func (_this ReplicaState) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case ReplicaState_ReplicaState: {
      return "_285_LiveRSL____ReplicaState__i_Compile.ReplicaState.ReplicaState" + "(" + _dafny.String(data.Constants) + ", " + _dafny.String(data.NextHeartbeatTime) + ", " + _dafny.String(data.Proposer) + ", " + _dafny.String(data.Acceptor) + ", " + _dafny.String(data.Learner) + ", " + _dafny.String(data.Executor) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this ReplicaState) Equals(other ReplicaState) bool {
  switch data1 := _this.Get().(type) {
    case ReplicaState_ReplicaState: {
      data2, ok := other.Get().(ReplicaState_ReplicaState)
return ok && data1.Constants.Equals(data2.Constants) && data1.NextHeartbeatTime == data2.NextHeartbeatTime && data1.Proposer.Equals(data2.Proposer) && data1.Acceptor.Equals(data2.Acceptor) && data1.Learner.Equals(data2.Learner) && data1.Executor.Equals(data2.Executor)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this ReplicaState) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(ReplicaState)
return ok && _this.Equals(typed)
}
func Type_ReplicaState_() _dafny.Type {
  return type_ReplicaState_{}
}

type type_ReplicaState_ struct {
}

func (_this type_ReplicaState_) Default() interface{} {
  return ReplicaState{ReplicaState_ReplicaState{_245_LiveRSL____ReplicaConstantsState__i_Compile.Type_ReplicaConstantsState_().Default().(_245_LiveRSL____ReplicaConstantsState__i_Compile.ReplicaConstantsState), 0, _254_LiveRSL____ProposerState__i_Compile.Type_ProposerState_().Default().(_254_LiveRSL____ProposerState__i_Compile.ProposerState), _269_LiveRSL____AcceptorState__i_Compile.Type_AcceptorState_().Default().(_269_LiveRSL____AcceptorState__i_Compile.AcceptorState), _278_LiveRSL____LearnerState__i_Compile.Type_CLearnerState_().Default().(_278_LiveRSL____LearnerState__i_Compile.CLearnerState), _276_LiveRSL____ExecutorState__i_Compile.Type_ExecutorState_().Default().(_276_LiveRSL____ExecutorState__i_Compile.ExecutorState)}}
}

func (_this type_ReplicaState_) String() string {
  return "_285_LiveRSL____ReplicaState__i_Compile.ReplicaState"
}
// End of data type ReplicaState

