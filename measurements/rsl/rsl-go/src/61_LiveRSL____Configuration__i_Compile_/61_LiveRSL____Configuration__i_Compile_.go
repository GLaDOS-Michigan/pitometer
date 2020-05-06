// Package _61_LiveRSL____Configuration__i_Compile
// Dafny module _61_LiveRSL____Configuration__i_Compile compiled into Go

package _61_LiveRSL____Configuration__i_Compile

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

type Dummy__ struct{}





// Definition of data type LConfiguration
type LConfiguration struct {
  Data_LConfiguration_
}

func (_this LConfiguration) Get() Data_LConfiguration_ {
  return _this.Data_LConfiguration_
}

type Data_LConfiguration_ interface {
  isLConfiguration()
}

type CompanionStruct_LConfiguration_ struct {}
var Companion_LConfiguration_ = CompanionStruct_LConfiguration_{}

type LConfiguration_LConfiguration struct {
  ClientIds _dafny.Set
Replica__ids _dafny.Seq
}

func (LConfiguration_LConfiguration) isLConfiguration() {}

func (CompanionStruct_LConfiguration_) Create_LConfiguration_(ClientIds _dafny.Set, Replica__ids _dafny.Seq) LConfiguration {
  return LConfiguration{LConfiguration_LConfiguration{ClientIds,Replica__ids}}
}

func (_this LConfiguration) Is_LConfiguration() bool {
  _, ok := _this.Get().(LConfiguration_LConfiguration)
return ok
}

func (_this LConfiguration) Dtor_clientIds() _dafny.Set {
  return _this.Get().(LConfiguration_LConfiguration).ClientIds
}

func (_this LConfiguration) Dtor_replica__ids() _dafny.Seq {
  return _this.Get().(LConfiguration_LConfiguration).Replica__ids
}

func (_this LConfiguration) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case LConfiguration_LConfiguration: {
      return "_61_LiveRSL____Configuration__i_Compile.LConfiguration.LConfiguration" + "(" + _dafny.String(data.ClientIds) + ", " + _dafny.String(data.Replica__ids) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this LConfiguration) Equals(other LConfiguration) bool {
  switch data1 := _this.Get().(type) {
    case LConfiguration_LConfiguration: {
      data2, ok := other.Get().(LConfiguration_LConfiguration)
return ok && data1.ClientIds.Equals(data2.ClientIds) && data1.Replica__ids.Equals(data2.Replica__ids)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this LConfiguration) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(LConfiguration)
return ok && _this.Equals(typed)
}
func Type_LConfiguration_() _dafny.Type {
  return type_LConfiguration_{}
}

type type_LConfiguration_ struct {
}

func (_this type_LConfiguration_) Default() interface{} {
  return LConfiguration{LConfiguration_LConfiguration{_dafny.EmptySet, _dafny.EmptySeq}}
}

func (_this type_LConfiguration_) String() string {
  return "_61_LiveRSL____Configuration__i_Compile.LConfiguration"
}
// End of data type LConfiguration

