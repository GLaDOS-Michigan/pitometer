// Package _99_LiveRSL____Election__i_Compile
// Dafny module _99_LiveRSL____Election__i_Compile compiled into Go

package _99_LiveRSL____Election__i_Compile

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

type Dummy__ struct{}






// Definition of data type ElectionState
type ElectionState struct {
  Data_ElectionState_
}

func (_this ElectionState) Get() Data_ElectionState_ {
  return _this.Data_ElectionState_
}

type Data_ElectionState_ interface {
  isElectionState()
}

type CompanionStruct_ElectionState_ struct {}
var Companion_ElectionState_ = CompanionStruct_ElectionState_{}

type ElectionState_ElectionState struct {
  Constants _78_LiveRSL____Constants__i_Compile.LReplicaConstants
Current__view _56_LiveRSL____Types__i_Compile.Ballot
Current__view__suspectors _dafny.Set
Epoch__end__time _dafny.Int
Epoch__length _dafny.Int
Requests__received__this__epoch _dafny.Seq
Requests__received__prev__epochs _dafny.Seq
}

func (ElectionState_ElectionState) isElectionState() {}

func (CompanionStruct_ElectionState_) Create_ElectionState_(Constants _78_LiveRSL____Constants__i_Compile.LReplicaConstants, Current__view _56_LiveRSL____Types__i_Compile.Ballot, Current__view__suspectors _dafny.Set, Epoch__end__time _dafny.Int, Epoch__length _dafny.Int, Requests__received__this__epoch _dafny.Seq, Requests__received__prev__epochs _dafny.Seq) ElectionState {
  return ElectionState{ElectionState_ElectionState{Constants,Current__view,Current__view__suspectors,Epoch__end__time,Epoch__length,Requests__received__this__epoch,Requests__received__prev__epochs}}
}

func (_this ElectionState) Is_ElectionState() bool {
  _, ok := _this.Get().(ElectionState_ElectionState)
return ok
}

func (_this ElectionState) Dtor_constants() _78_LiveRSL____Constants__i_Compile.LReplicaConstants {
  return _this.Get().(ElectionState_ElectionState).Constants
}

func (_this ElectionState) Dtor_current__view() _56_LiveRSL____Types__i_Compile.Ballot {
  return _this.Get().(ElectionState_ElectionState).Current__view
}

func (_this ElectionState) Dtor_current__view__suspectors() _dafny.Set {
  return _this.Get().(ElectionState_ElectionState).Current__view__suspectors
}

func (_this ElectionState) Dtor_epoch__end__time() _dafny.Int {
  return _this.Get().(ElectionState_ElectionState).Epoch__end__time
}

func (_this ElectionState) Dtor_epoch__length() _dafny.Int {
  return _this.Get().(ElectionState_ElectionState).Epoch__length
}

func (_this ElectionState) Dtor_requests__received__this__epoch() _dafny.Seq {
  return _this.Get().(ElectionState_ElectionState).Requests__received__this__epoch
}

func (_this ElectionState) Dtor_requests__received__prev__epochs() _dafny.Seq {
  return _this.Get().(ElectionState_ElectionState).Requests__received__prev__epochs
}

func (_this ElectionState) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case ElectionState_ElectionState: {
      return "_99_LiveRSL____Election__i_Compile.ElectionState.ElectionState" + "(" + _dafny.String(data.Constants) + ", " + _dafny.String(data.Current__view) + ", " + _dafny.String(data.Current__view__suspectors) + ", " + _dafny.String(data.Epoch__end__time) + ", " + _dafny.String(data.Epoch__length) + ", " + _dafny.String(data.Requests__received__this__epoch) + ", " + _dafny.String(data.Requests__received__prev__epochs) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this ElectionState) Equals(other ElectionState) bool {
  switch data1 := _this.Get().(type) {
    case ElectionState_ElectionState: {
      data2, ok := other.Get().(ElectionState_ElectionState)
return ok && data1.Constants.Equals(data2.Constants) && data1.Current__view.Equals(data2.Current__view) && data1.Current__view__suspectors.Equals(data2.Current__view__suspectors) && data1.Epoch__end__time.Cmp(data2.Epoch__end__time) == 0 && data1.Epoch__length.Cmp(data2.Epoch__length) == 0 && data1.Requests__received__this__epoch.Equals(data2.Requests__received__this__epoch) && data1.Requests__received__prev__epochs.Equals(data2.Requests__received__prev__epochs)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this ElectionState) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(ElectionState)
return ok && _this.Equals(typed)
}
func Type_ElectionState_() _dafny.Type {
  return type_ElectionState_{}
}

type type_ElectionState_ struct {
}

func (_this type_ElectionState_) Default() interface{} {
  return ElectionState{ElectionState_ElectionState{_78_LiveRSL____Constants__i_Compile.Type_LReplicaConstants_().Default().(_78_LiveRSL____Constants__i_Compile.LReplicaConstants), _56_LiveRSL____Types__i_Compile.Type_Ballot_().Default().(_56_LiveRSL____Types__i_Compile.Ballot), _dafny.EmptySet, _dafny.Zero, _dafny.Zero, _dafny.EmptySeq, _dafny.EmptySeq}}
}

func (_this type_ElectionState_) String() string {
  return "_99_LiveRSL____Election__i_Compile.ElectionState"
}
// End of data type ElectionState

