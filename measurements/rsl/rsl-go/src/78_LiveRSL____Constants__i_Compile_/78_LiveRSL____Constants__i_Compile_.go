// Package _78_LiveRSL____Constants__i_Compile
// Dafny module _78_LiveRSL____Constants__i_Compile compiled into Go

package _78_LiveRSL____Constants__i_Compile

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

type Dummy__ struct{}




// Definition of data type LConstants
type LConstants struct {
  Data_LConstants_
}

func (_this LConstants) Get() Data_LConstants_ {
  return _this.Data_LConstants_
}

type Data_LConstants_ interface {
  isLConstants()
}

type CompanionStruct_LConstants_ struct {}
var Companion_LConstants_ = CompanionStruct_LConstants_{}

type LConstants_LConstants struct {
  Config _61_LiveRSL____Configuration__i_Compile.LConfiguration
Params _76_LiveRSL____Parameters__i_Compile.LParameters
}

func (LConstants_LConstants) isLConstants() {}

func (CompanionStruct_LConstants_) Create_LConstants_(Config _61_LiveRSL____Configuration__i_Compile.LConfiguration, Params _76_LiveRSL____Parameters__i_Compile.LParameters) LConstants {
  return LConstants{LConstants_LConstants{Config,Params}}
}

func (_this LConstants) Is_LConstants() bool {
  _, ok := _this.Get().(LConstants_LConstants)
return ok
}

func (_this LConstants) Dtor_config() _61_LiveRSL____Configuration__i_Compile.LConfiguration {
  return _this.Get().(LConstants_LConstants).Config
}

func (_this LConstants) Dtor_params() _76_LiveRSL____Parameters__i_Compile.LParameters {
  return _this.Get().(LConstants_LConstants).Params
}

func (_this LConstants) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case LConstants_LConstants: {
      return "_78_LiveRSL____Constants__i_Compile.LConstants.LConstants" + "(" + _dafny.String(data.Config) + ", " + _dafny.String(data.Params) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this LConstants) Equals(other LConstants) bool {
  switch data1 := _this.Get().(type) {
    case LConstants_LConstants: {
      data2, ok := other.Get().(LConstants_LConstants)
return ok && data1.Config.Equals(data2.Config) && data1.Params.Equals(data2.Params)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this LConstants) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(LConstants)
return ok && _this.Equals(typed)
}
func Type_LConstants_() _dafny.Type {
  return type_LConstants_{}
}

type type_LConstants_ struct {
}

func (_this type_LConstants_) Default() interface{} {
  return LConstants{LConstants_LConstants{_61_LiveRSL____Configuration__i_Compile.Type_LConfiguration_().Default().(_61_LiveRSL____Configuration__i_Compile.LConfiguration), _76_LiveRSL____Parameters__i_Compile.Type_LParameters_().Default().(_76_LiveRSL____Parameters__i_Compile.LParameters)}}
}

func (_this type_LConstants_) String() string {
  return "_78_LiveRSL____Constants__i_Compile.LConstants"
}
// End of data type LConstants

// Definition of data type LReplicaConstants
type LReplicaConstants struct {
  Data_LReplicaConstants_
}

func (_this LReplicaConstants) Get() Data_LReplicaConstants_ {
  return _this.Data_LReplicaConstants_
}

type Data_LReplicaConstants_ interface {
  isLReplicaConstants()
}

type CompanionStruct_LReplicaConstants_ struct {}
var Companion_LReplicaConstants_ = CompanionStruct_LReplicaConstants_{}

type LReplicaConstants_LReplicaConstants struct {
  My__index _dafny.Int
All _78_LiveRSL____Constants__i_Compile.LConstants
}

func (LReplicaConstants_LReplicaConstants) isLReplicaConstants() {}

func (CompanionStruct_LReplicaConstants_) Create_LReplicaConstants_(My__index _dafny.Int, All _78_LiveRSL____Constants__i_Compile.LConstants) LReplicaConstants {
  return LReplicaConstants{LReplicaConstants_LReplicaConstants{My__index,All}}
}

func (_this LReplicaConstants) Is_LReplicaConstants() bool {
  _, ok := _this.Get().(LReplicaConstants_LReplicaConstants)
return ok
}

func (_this LReplicaConstants) Dtor_my__index() _dafny.Int {
  return _this.Get().(LReplicaConstants_LReplicaConstants).My__index
}

func (_this LReplicaConstants) Dtor_all() _78_LiveRSL____Constants__i_Compile.LConstants {
  return _this.Get().(LReplicaConstants_LReplicaConstants).All
}

func (_this LReplicaConstants) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case LReplicaConstants_LReplicaConstants: {
      return "_78_LiveRSL____Constants__i_Compile.LReplicaConstants.LReplicaConstants" + "(" + _dafny.String(data.My__index) + ", " + _dafny.String(data.All) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this LReplicaConstants) Equals(other LReplicaConstants) bool {
  switch data1 := _this.Get().(type) {
    case LReplicaConstants_LReplicaConstants: {
      data2, ok := other.Get().(LReplicaConstants_LReplicaConstants)
return ok && data1.My__index.Cmp(data2.My__index) == 0 && data1.All.Equals(data2.All)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this LReplicaConstants) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(LReplicaConstants)
return ok && _this.Equals(typed)
}
func Type_LReplicaConstants_() _dafny.Type {
  return type_LReplicaConstants_{}
}

type type_LReplicaConstants_ struct {
}

func (_this type_LReplicaConstants_) Default() interface{} {
  return LReplicaConstants{LReplicaConstants_LReplicaConstants{_dafny.Zero, Type_LConstants_().Default().(_78_LiveRSL____Constants__i_Compile.LConstants)}}
}

func (_this type_LReplicaConstants_) String() string {
  return "_78_LiveRSL____Constants__i_Compile.LReplicaConstants"
}
// End of data type LReplicaConstants

