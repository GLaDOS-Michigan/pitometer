// Package _123_LiveRSL____Learner__i_Compile
// Dafny module _123_LiveRSL____Learner__i_Compile compiled into Go

package _123_LiveRSL____Learner__i_Compile

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

type Dummy__ struct{}








// Definition of data type LLearner
type LLearner struct {
  Data_LLearner_
}

func (_this LLearner) Get() Data_LLearner_ {
  return _this.Data_LLearner_
}

type Data_LLearner_ interface {
  isLLearner()
}

type CompanionStruct_LLearner_ struct {}
var Companion_LLearner_ = CompanionStruct_LLearner_{}

type LLearner_LLearner struct {
  Constants _78_LiveRSL____Constants__i_Compile.LReplicaConstants
Max__ballot__seen _56_LiveRSL____Types__i_Compile.Ballot
Unexecuted__learner__state _dafny.Map
}

func (LLearner_LLearner) isLLearner() {}

func (CompanionStruct_LLearner_) Create_LLearner_(Constants _78_LiveRSL____Constants__i_Compile.LReplicaConstants, Max__ballot__seen _56_LiveRSL____Types__i_Compile.Ballot, Unexecuted__learner__state _dafny.Map) LLearner {
  return LLearner{LLearner_LLearner{Constants,Max__ballot__seen,Unexecuted__learner__state}}
}

func (_this LLearner) Is_LLearner() bool {
  _, ok := _this.Get().(LLearner_LLearner)
return ok
}

func (_this LLearner) Dtor_constants() _78_LiveRSL____Constants__i_Compile.LReplicaConstants {
  return _this.Get().(LLearner_LLearner).Constants
}

func (_this LLearner) Dtor_max__ballot__seen() _56_LiveRSL____Types__i_Compile.Ballot {
  return _this.Get().(LLearner_LLearner).Max__ballot__seen
}

func (_this LLearner) Dtor_unexecuted__learner__state() _dafny.Map {
  return _this.Get().(LLearner_LLearner).Unexecuted__learner__state
}

func (_this LLearner) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case LLearner_LLearner: {
      return "_123_LiveRSL____Learner__i_Compile.LLearner.LLearner" + "(" + _dafny.String(data.Constants) + ", " + _dafny.String(data.Max__ballot__seen) + ", " + _dafny.String(data.Unexecuted__learner__state) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this LLearner) Equals(other LLearner) bool {
  switch data1 := _this.Get().(type) {
    case LLearner_LLearner: {
      data2, ok := other.Get().(LLearner_LLearner)
return ok && data1.Constants.Equals(data2.Constants) && data1.Max__ballot__seen.Equals(data2.Max__ballot__seen) && data1.Unexecuted__learner__state.Equals(data2.Unexecuted__learner__state)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this LLearner) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(LLearner)
return ok && _this.Equals(typed)
}
func Type_LLearner_() _dafny.Type {
  return type_LLearner_{}
}

type type_LLearner_ struct {
}

func (_this type_LLearner_) Default() interface{} {
  return LLearner{LLearner_LLearner{_78_LiveRSL____Constants__i_Compile.Type_LReplicaConstants_().Default().(_78_LiveRSL____Constants__i_Compile.LReplicaConstants), _56_LiveRSL____Types__i_Compile.Type_Ballot_().Default().(_56_LiveRSL____Types__i_Compile.Ballot), _dafny.EmptyMap}}
}

func (_this type_LLearner_) String() string {
  return "_123_LiveRSL____Learner__i_Compile.LLearner"
}
// End of data type LLearner

