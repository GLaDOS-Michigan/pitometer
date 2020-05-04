// Package _71_LiveRSL____ClockReading__i_Compile
// Dafny module _71_LiveRSL____ClockReading__i_Compile compiled into Go

package _71_LiveRSL____ClockReading__i_Compile

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

type Dummy__ struct{}


// Definition of data type ClockReading
type ClockReading struct {
  Data_ClockReading_
}

func (_this ClockReading) Get() Data_ClockReading_ {
  return _this.Data_ClockReading_
}

type Data_ClockReading_ interface {
  isClockReading()
}

type CompanionStruct_ClockReading_ struct {}
var Companion_ClockReading_ = CompanionStruct_ClockReading_{}

type ClockReading_ClockReading struct {
  T _dafny.Int
}

func (ClockReading_ClockReading) isClockReading() {}

func (CompanionStruct_ClockReading_) Create_ClockReading_(T _dafny.Int) ClockReading {
  return ClockReading{ClockReading_ClockReading{T}}
}

func (_this ClockReading) Is_ClockReading() bool {
  _, ok := _this.Get().(ClockReading_ClockReading)
return ok
}

func (_this ClockReading) Dtor_t() _dafny.Int {
  return _this.Get().(ClockReading_ClockReading).T
}

func (_this ClockReading) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case ClockReading_ClockReading: {
      return "_71_LiveRSL____ClockReading__i_Compile.ClockReading.ClockReading" + "(" + _dafny.String(data.T) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this ClockReading) Equals(other ClockReading) bool {
  switch data1 := _this.Get().(type) {
    case ClockReading_ClockReading: {
      data2, ok := other.Get().(ClockReading_ClockReading)
return ok && data1.T.Cmp(data2.T) == 0
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this ClockReading) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(ClockReading)
return ok && _this.Equals(typed)
}
func Type_ClockReading_() _dafny.Type {
  return type_ClockReading_{}
}

type type_ClockReading_ struct {
}

func (_this type_ClockReading_) Default() interface{} {
  return ClockReading{ClockReading_ClockReading{_dafny.Zero}}
}

func (_this type_ClockReading_) String() string {
  return "_71_LiveRSL____ClockReading__i_Compile.ClockReading"
}
// End of data type ClockReading

