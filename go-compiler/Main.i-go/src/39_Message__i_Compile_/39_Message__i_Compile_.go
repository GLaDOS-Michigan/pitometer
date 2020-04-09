// Package _39_Message__i_Compile
// Dafny module _39_Message__i_Compile compiled into Go

package _39_Message__i_Compile

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

type Dummy__ struct{}



// Definition of data type CMessage
type CMessage struct {
  Data_CMessage_
}

func (_this CMessage) Get() Data_CMessage_ {
  return _this.Data_CMessage_
}

type Data_CMessage_ interface {
  isCMessage()
}

type CompanionStruct_CMessage_ struct {}
var Companion_CMessage_ = CompanionStruct_CMessage_{}

type CMessage_CTransfer struct {
  Transfer__epoch uint64
}

func (CMessage_CTransfer) isCMessage() {}

func (CompanionStruct_CMessage_) Create_CTransfer_(Transfer__epoch uint64) CMessage {
  return CMessage{CMessage_CTransfer{Transfer__epoch}}
}

func (_this CMessage) Is_CTransfer() bool {
  _, ok := _this.Get().(CMessage_CTransfer)
return ok
}

type CMessage_CLocked struct {
  Locked__epoch uint64
}

func (CMessage_CLocked) isCMessage() {}

func (CompanionStruct_CMessage_) Create_CLocked_(Locked__epoch uint64) CMessage {
  return CMessage{CMessage_CLocked{Locked__epoch}}
}

func (_this CMessage) Is_CLocked() bool {
  _, ok := _this.Get().(CMessage_CLocked)
return ok
}

type CMessage_CInvalid struct {
}

func (CMessage_CInvalid) isCMessage() {}

func (CompanionStruct_CMessage_) Create_CInvalid_() CMessage {
  return CMessage{CMessage_CInvalid{}}
}

func (_this CMessage) Is_CInvalid() bool {
  _, ok := _this.Get().(CMessage_CInvalid)
return ok
}

func (_this CMessage) Dtor_transfer__epoch() uint64 {
  return _this.Get().(CMessage_CTransfer).Transfer__epoch
}

func (_this CMessage) Dtor_locked__epoch() uint64 {
  return _this.Get().(CMessage_CLocked).Locked__epoch
}

func (_this CMessage) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case CMessage_CTransfer: {
      return "_39_Message__i_Compile.CMessage.CTransfer" + "(" + _dafny.String(data.Transfer__epoch) + ")"
    }
    case CMessage_CLocked: {
      return "_39_Message__i_Compile.CMessage.CLocked" + "(" + _dafny.String(data.Locked__epoch) + ")"
    }
    case CMessage_CInvalid: {
      return "_39_Message__i_Compile.CMessage.CInvalid"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this CMessage) Equals(other CMessage) bool {
  switch data1 := _this.Get().(type) {
    case CMessage_CTransfer: {
      data2, ok := other.Get().(CMessage_CTransfer)
return ok && data1.Transfer__epoch == data2.Transfer__epoch
    }
    case CMessage_CLocked: {
      data2, ok := other.Get().(CMessage_CLocked)
return ok && data1.Locked__epoch == data2.Locked__epoch
    }
    case CMessage_CInvalid: {
      _, ok := other.Get().(CMessage_CInvalid)
return ok
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this CMessage) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(CMessage)
return ok && _this.Equals(typed)
}
func Type_CMessage_() _dafny.Type {
  return type_CMessage_{}
}

type type_CMessage_ struct {
}

func (_this type_CMessage_) Default() interface{} {
  return CMessage{CMessage_CTransfer{0}}
}

func (_this type_CMessage_) String() string {
  return "_39_Message__i_Compile.CMessage"
}
// End of data type CMessage


