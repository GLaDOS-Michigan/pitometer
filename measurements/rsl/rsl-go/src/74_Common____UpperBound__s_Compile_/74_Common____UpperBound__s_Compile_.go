// Package _74_Common____UpperBound__s_Compile
// Dafny module _74_Common____UpperBound__s_Compile compiled into Go

package _74_Common____UpperBound__s_Compile

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

type Dummy__ struct{}


// Definition of data type UpperBound
type UpperBound struct {
  Data_UpperBound_
}

func (_this UpperBound) Get() Data_UpperBound_ {
  return _this.Data_UpperBound_
}

type Data_UpperBound_ interface {
  isUpperBound()
}

type CompanionStruct_UpperBound_ struct {}
var Companion_UpperBound_ = CompanionStruct_UpperBound_{}

type UpperBound_UpperBoundFinite struct {
  N _dafny.Int
}

func (UpperBound_UpperBoundFinite) isUpperBound() {}

func (CompanionStruct_UpperBound_) Create_UpperBoundFinite_(N _dafny.Int) UpperBound {
  return UpperBound{UpperBound_UpperBoundFinite{N}}
}

func (_this UpperBound) Is_UpperBoundFinite() bool {
  _, ok := _this.Get().(UpperBound_UpperBoundFinite)
return ok
}

type UpperBound_UpperBoundInfinite struct {
}

func (UpperBound_UpperBoundInfinite) isUpperBound() {}

func (CompanionStruct_UpperBound_) Create_UpperBoundInfinite_() UpperBound {
  return UpperBound{UpperBound_UpperBoundInfinite{}}
}

func (_this UpperBound) Is_UpperBoundInfinite() bool {
  _, ok := _this.Get().(UpperBound_UpperBoundInfinite)
return ok
}

func (_this UpperBound) Dtor_n() _dafny.Int {
  return _this.Get().(UpperBound_UpperBoundFinite).N
}

func (_this UpperBound) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case UpperBound_UpperBoundFinite: {
      return "_74_Common____UpperBound__s_Compile.UpperBound.UpperBoundFinite" + "(" + _dafny.String(data.N) + ")"
    }
    case UpperBound_UpperBoundInfinite: {
      return "_74_Common____UpperBound__s_Compile.UpperBound.UpperBoundInfinite"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this UpperBound) Equals(other UpperBound) bool {
  switch data1 := _this.Get().(type) {
    case UpperBound_UpperBoundFinite: {
      data2, ok := other.Get().(UpperBound_UpperBoundFinite)
return ok && data1.N.Cmp(data2.N) == 0
    }
    case UpperBound_UpperBoundInfinite: {
      _, ok := other.Get().(UpperBound_UpperBoundInfinite)
return ok
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this UpperBound) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(UpperBound)
return ok && _this.Equals(typed)
}
func Type_UpperBound_() _dafny.Type {
  return type_UpperBound_{}
}

type type_UpperBound_ struct {
}

func (_this type_UpperBound_) Default() interface{} {
  return UpperBound{UpperBound_UpperBoundFinite{_dafny.Zero}}
}

func (_this type_UpperBound_) String() string {
  return "_74_Common____UpperBound__s_Compile.UpperBound"
}
// End of data type UpperBound

