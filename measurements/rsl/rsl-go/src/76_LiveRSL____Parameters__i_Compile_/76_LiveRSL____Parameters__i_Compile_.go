// Package _76_LiveRSL____Parameters__i_Compile
// Dafny module _76_LiveRSL____Parameters__i_Compile compiled into Go

package _76_LiveRSL____Parameters__i_Compile

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

type Dummy__ struct{}



// Definition of data type LParameters
type LParameters struct {
  Data_LParameters_
}

func (_this LParameters) Get() Data_LParameters_ {
  return _this.Data_LParameters_
}

type Data_LParameters_ interface {
  isLParameters()
}

type CompanionStruct_LParameters_ struct {}
var Companion_LParameters_ = CompanionStruct_LParameters_{}

type LParameters_LParameters struct {
  Max__log__length _dafny.Int
Baseline__view__timeout__period _dafny.Int
Heartbeat__period _dafny.Int
Max__integer__val _74_Common____UpperBound__s_Compile.UpperBound
Max__batch__size _dafny.Int
Max__batch__delay _dafny.Int
}

func (LParameters_LParameters) isLParameters() {}

func (CompanionStruct_LParameters_) Create_LParameters_(Max__log__length _dafny.Int, Baseline__view__timeout__period _dafny.Int, Heartbeat__period _dafny.Int, Max__integer__val _74_Common____UpperBound__s_Compile.UpperBound, Max__batch__size _dafny.Int, Max__batch__delay _dafny.Int) LParameters {
  return LParameters{LParameters_LParameters{Max__log__length,Baseline__view__timeout__period,Heartbeat__period,Max__integer__val,Max__batch__size,Max__batch__delay}}
}

func (_this LParameters) Is_LParameters() bool {
  _, ok := _this.Get().(LParameters_LParameters)
return ok
}

func (_this LParameters) Dtor_max__log__length() _dafny.Int {
  return _this.Get().(LParameters_LParameters).Max__log__length
}

func (_this LParameters) Dtor_baseline__view__timeout__period() _dafny.Int {
  return _this.Get().(LParameters_LParameters).Baseline__view__timeout__period
}

func (_this LParameters) Dtor_heartbeat__period() _dafny.Int {
  return _this.Get().(LParameters_LParameters).Heartbeat__period
}

func (_this LParameters) Dtor_max__integer__val() _74_Common____UpperBound__s_Compile.UpperBound {
  return _this.Get().(LParameters_LParameters).Max__integer__val
}

func (_this LParameters) Dtor_max__batch__size() _dafny.Int {
  return _this.Get().(LParameters_LParameters).Max__batch__size
}

func (_this LParameters) Dtor_max__batch__delay() _dafny.Int {
  return _this.Get().(LParameters_LParameters).Max__batch__delay
}

func (_this LParameters) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case LParameters_LParameters: {
      return "_76_LiveRSL____Parameters__i_Compile.LParameters.LParameters" + "(" + _dafny.String(data.Max__log__length) + ", " + _dafny.String(data.Baseline__view__timeout__period) + ", " + _dafny.String(data.Heartbeat__period) + ", " + _dafny.String(data.Max__integer__val) + ", " + _dafny.String(data.Max__batch__size) + ", " + _dafny.String(data.Max__batch__delay) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this LParameters) Equals(other LParameters) bool {
  switch data1 := _this.Get().(type) {
    case LParameters_LParameters: {
      data2, ok := other.Get().(LParameters_LParameters)
return ok && data1.Max__log__length.Cmp(data2.Max__log__length) == 0 && data1.Baseline__view__timeout__period.Cmp(data2.Baseline__view__timeout__period) == 0 && data1.Heartbeat__period.Cmp(data2.Heartbeat__period) == 0 && data1.Max__integer__val.Equals(data2.Max__integer__val) && data1.Max__batch__size.Cmp(data2.Max__batch__size) == 0 && data1.Max__batch__delay.Cmp(data2.Max__batch__delay) == 0
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this LParameters) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(LParameters)
return ok && _this.Equals(typed)
}
func Type_LParameters_() _dafny.Type {
  return type_LParameters_{}
}

type type_LParameters_ struct {
}

func (_this type_LParameters_) Default() interface{} {
  return LParameters{LParameters_LParameters{_dafny.Zero, _dafny.Zero, _dafny.Zero, _74_Common____UpperBound__s_Compile.Type_UpperBound_().Default().(_74_Common____UpperBound__s_Compile.UpperBound), _dafny.Zero, _dafny.Zero}}
}

func (_this type_LParameters_) String() string {
  return "_76_LiveRSL____Parameters__i_Compile.LParameters"
}
// End of data type LParameters

