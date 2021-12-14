// Package _50_AppStateMachine__i_Compile
// Dafny module _50_AppStateMachine__i_Compile compiled into Go

package _50_AppStateMachine__i_Compile

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

type Dummy__ struct{}



// Definition of data type AppMessage_k
type AppMessage_k struct {
  Data_AppMessage_k_
}

func (_this AppMessage_k) Get() Data_AppMessage_k_ {
  return _this.Data_AppMessage_k_
}

type Data_AppMessage_k_ interface {
  isAppMessage_k()
}

type CompanionStruct_AppMessage_k_ struct {}
var Companion_AppMessage_k_ = CompanionStruct_AppMessage_k_{}

type AppMessage_k_AppIncrementRequest struct {
}

func (AppMessage_k_AppIncrementRequest) isAppMessage_k() {}

func (CompanionStruct_AppMessage_k_) Create_AppIncrementRequest_() AppMessage_k {
  return AppMessage_k{AppMessage_k_AppIncrementRequest{}}
}

func (_this AppMessage_k) Is_AppIncrementRequest() bool {
  _, ok := _this.Get().(AppMessage_k_AppIncrementRequest)
return ok
}

type AppMessage_k_AppIncrementReply struct {
  Response uint64
}

func (AppMessage_k_AppIncrementReply) isAppMessage_k() {}

func (CompanionStruct_AppMessage_k_) Create_AppIncrementReply_(Response uint64) AppMessage_k {
  return AppMessage_k{AppMessage_k_AppIncrementReply{Response}}
}

func (_this AppMessage_k) Is_AppIncrementReply() bool {
  _, ok := _this.Get().(AppMessage_k_AppIncrementReply)
return ok
}

type AppMessage_k_AppInvalidReply struct {
}

func (AppMessage_k_AppInvalidReply) isAppMessage_k() {}

func (CompanionStruct_AppMessage_k_) Create_AppInvalidReply_() AppMessage_k {
  return AppMessage_k{AppMessage_k_AppInvalidReply{}}
}

func (_this AppMessage_k) Is_AppInvalidReply() bool {
  _, ok := _this.Get().(AppMessage_k_AppInvalidReply)
return ok
}

func (_this AppMessage_k) Dtor_response() uint64 {
  return _this.Get().(AppMessage_k_AppIncrementReply).Response
}

func (_this AppMessage_k) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case AppMessage_k_AppIncrementRequest: {
      return "_50_AppStateMachine__i_Compile.AppMessage'.AppIncrementRequest"
    }
    case AppMessage_k_AppIncrementReply: {
      return "_50_AppStateMachine__i_Compile.AppMessage'.AppIncrementReply" + "(" + _dafny.String(data.Response) + ")"
    }
    case AppMessage_k_AppInvalidReply: {
      return "_50_AppStateMachine__i_Compile.AppMessage'.AppInvalidReply"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this AppMessage_k) Equals(other AppMessage_k) bool {
  switch data1 := _this.Get().(type) {
    case AppMessage_k_AppIncrementRequest: {
      _, ok := other.Get().(AppMessage_k_AppIncrementRequest)
return ok
    }
    case AppMessage_k_AppIncrementReply: {
      data2, ok := other.Get().(AppMessage_k_AppIncrementReply)
return ok && data1.Response == data2.Response
    }
    case AppMessage_k_AppInvalidReply: {
      _, ok := other.Get().(AppMessage_k_AppInvalidReply)
return ok
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this AppMessage_k) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(AppMessage_k)
return ok && _this.Equals(typed)
}
func Type_AppMessage_k_() _dafny.Type {
  return type_AppMessage_k_{}
}

type type_AppMessage_k_ struct {
}

func (_this type_AppMessage_k_) Default() interface{} {
  return AppMessage_k{AppMessage_k_AppIncrementRequest{}}
}

func (_this type_AppMessage_k_) String() string {
  return "_50_AppStateMachine__i_Compile.AppMessage_k"
}
// End of data type AppMessage_k



