// Package _42_Common____UdpClient__i_Compile
// Dafny module _42_Common____UdpClient__i_Compile compiled into Go

package _42_Common____UdpClient__i_Compile

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

type Dummy__ struct{}



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
  return "_42_Common____UdpClient__i_Compile.Default__"
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
  return "_42_Common____UdpClient__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) EndPointIsValidIPV4(endPoint _9_Native____Io__s_Compile.EndPoint) bool {
  return ((((endPoint).Dtor_addr()).Cardinality()).Cmp(_dafny.IntOfInt64(4)) == 0) && (((uint16(0)) <= ((endPoint).Dtor_port())) && (((endPoint).Dtor_port()) <= (uint16(65535))))
}
// End of class Default__
