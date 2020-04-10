// Package _44_Logic____Option__i_Compile
// Dafny module _44_Logic____Option__i_Compile compiled into Go

package _44_Logic____Option__i_Compile

import (
  _dafny "dafny"
_System "System_"
_0_Native____NativeTypes__s_Compile "0_Native____NativeTypes__s_Compile_"
_2_Collections____Maps2__s_Compile "2_Collections____Maps2__s_Compile_"
_5_Temporal____Temporal__s_Compile "5_Temporal____Temporal__s_Compile_"
_7_Environment__s_Compile "7_Environment__s_Compile_"
_9_Native____Io__s_Compile "9_Native____Io__s_Compile_"
_26_Collections____Seqs__s_Compile "26_Collections____Seqs__s_Compile_"
_29_Collections____Sets__i_Compile "29_Collections____Sets__i_Compile_"
_33_Types__i_Compile "33_Types__i_Compile_"
_36_Protocol__Node__i_Compile "36_Protocol__Node__i_Compile_"
_39_Message__i_Compile "39_Message__i_Compile_"
_42_Common____UdpClient__i_Compile "42_Common____UdpClient__i_Compile_"
)
var _ _dafny.Dummy__
var _ _System.Dummy__
var _ _0_Native____NativeTypes__s_Compile.Dummy__
var _ _2_Collections____Maps2__s_Compile.Dummy__
var _ _5_Temporal____Temporal__s_Compile.Dummy__
var _ _7_Environment__s_Compile.Dummy__
var _ _9_Native____Io__s_Compile.Dummy__
var _ _26_Collections____Seqs__s_Compile.Dummy__
var _ _29_Collections____Sets__i_Compile.Dummy__
var _ _33_Types__i_Compile.Dummy__
var _ _36_Protocol__Node__i_Compile.Dummy__
var _ _39_Message__i_Compile.Dummy__
var _ _42_Common____UdpClient__i_Compile.Dummy__

type Dummy__ struct{}


// Definition of data type Option
type Option struct {
  Data_Option_
}

func (_this Option) Get() Data_Option_ {
  return _this.Data_Option_
}

type Data_Option_ interface {
  isOption()
}

type CompanionStruct_Option_ struct {}
var Companion_Option_ = CompanionStruct_Option_{}

type Option_Some struct {
  V interface{}
}

func (Option_Some) isOption() {}

func (CompanionStruct_Option_) Create_Some_(V interface{}) Option {
  return Option{Option_Some{V}}
}

func (_this Option) Is_Some() bool {
  _, ok := _this.Get().(Option_Some)
return ok
}

type Option_None struct {
}

func (Option_None) isOption() {}

func (CompanionStruct_Option_) Create_None_() Option {
  return Option{Option_None{}}
}

func (_this Option) Is_None() bool {
  _, ok := _this.Get().(Option_None)
return ok
}

func (_this Option) Dtor_v() interface{} {
  return _this.Get().(Option_Some).V
}

func (_this Option) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case Option_Some: {
      return "_44_Logic____Option__i_Compile.Option.Some" + "(" + _dafny.String(data.V) + ")"
    }
    case Option_None: {
      return "_44_Logic____Option__i_Compile.Option.None"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this Option) Equals(other Option) bool {
  switch data1 := _this.Get().(type) {
    case Option_Some: {
      data2, ok := other.Get().(Option_Some)
return ok && _dafny.AreEqual(data1.V, data2.V)
    }
    case Option_None: {
      _, ok := other.Get().(Option_None)
return ok
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this Option) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(Option)
return ok && _this.Equals(typed)
}
func Type_Option_() _dafny.Type {
  return type_Option_{}
}

type type_Option_ struct {
}

func (_this type_Option_) Default() interface{} {
  return Option{Option_None{}}
}

func (_this type_Option_) String() string {
  return "_44_Logic____Option__i_Compile.Option"
}
// End of data type Option

