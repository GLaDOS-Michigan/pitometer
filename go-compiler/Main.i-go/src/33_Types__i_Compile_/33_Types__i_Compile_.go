// Package _33_Types__i_Compile
// Dafny module _33_Types__i_Compile compiled into Go

package _33_Types__i_Compile

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

type Dummy__ struct{}




// Definition of data type LockMessage
type LockMessage struct {
  Data_LockMessage_
}

func (_this LockMessage) Get() Data_LockMessage_ {
  return _this.Data_LockMessage_
}

type Data_LockMessage_ interface {
  isLockMessage()
}

type CompanionStruct_LockMessage_ struct {}
var Companion_LockMessage_ = CompanionStruct_LockMessage_{}

type LockMessage_Transfer struct {
  Transfer__epoch _dafny.Int
}

func (LockMessage_Transfer) isLockMessage() {}

func (CompanionStruct_LockMessage_) Create_Transfer_(Transfer__epoch _dafny.Int) LockMessage {
  return LockMessage{LockMessage_Transfer{Transfer__epoch}}
}

func (_this LockMessage) Is_Transfer() bool {
  _, ok := _this.Get().(LockMessage_Transfer)
return ok
}

type LockMessage_Locked struct {
  Locked__epoch _dafny.Int
}

func (LockMessage_Locked) isLockMessage() {}

func (CompanionStruct_LockMessage_) Create_Locked_(Locked__epoch _dafny.Int) LockMessage {
  return LockMessage{LockMessage_Locked{Locked__epoch}}
}

func (_this LockMessage) Is_Locked() bool {
  _, ok := _this.Get().(LockMessage_Locked)
return ok
}

type LockMessage_Invalid struct {
}

func (LockMessage_Invalid) isLockMessage() {}

func (CompanionStruct_LockMessage_) Create_Invalid_() LockMessage {
  return LockMessage{LockMessage_Invalid{}}
}

func (_this LockMessage) Is_Invalid() bool {
  _, ok := _this.Get().(LockMessage_Invalid)
return ok
}

func (_this LockMessage) Dtor_transfer__epoch() _dafny.Int {
  return _this.Get().(LockMessage_Transfer).Transfer__epoch
}

func (_this LockMessage) Dtor_locked__epoch() _dafny.Int {
  return _this.Get().(LockMessage_Locked).Locked__epoch
}

func (_this LockMessage) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case LockMessage_Transfer: {
      return "_33_Types__i_Compile.LockMessage.Transfer" + "(" + _dafny.String(data.Transfer__epoch) + ")"
    }
    case LockMessage_Locked: {
      return "_33_Types__i_Compile.LockMessage.Locked" + "(" + _dafny.String(data.Locked__epoch) + ")"
    }
    case LockMessage_Invalid: {
      return "_33_Types__i_Compile.LockMessage.Invalid"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this LockMessage) Equals(other LockMessage) bool {
  switch data1 := _this.Get().(type) {
    case LockMessage_Transfer: {
      data2, ok := other.Get().(LockMessage_Transfer)
return ok && data1.Transfer__epoch.Cmp(data2.Transfer__epoch) == 0
    }
    case LockMessage_Locked: {
      data2, ok := other.Get().(LockMessage_Locked)
return ok && data1.Locked__epoch.Cmp(data2.Locked__epoch) == 0
    }
    case LockMessage_Invalid: {
      _, ok := other.Get().(LockMessage_Invalid)
return ok
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this LockMessage) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(LockMessage)
return ok && _this.Equals(typed)
}
func Type_LockMessage_() _dafny.Type {
  return type_LockMessage_{}
}

type type_LockMessage_ struct {
}

func (_this type_LockMessage_) Default() interface{} {
  return LockMessage{LockMessage_Transfer{_dafny.Zero}}
}

func (_this type_LockMessage_) String() string {
  return "_33_Types__i_Compile.LockMessage"
}
// End of data type LockMessage

// Definition of data type LockStep
type LockStep struct {
  Data_LockStep_
}

func (_this LockStep) Get() Data_LockStep_ {
  return _this.Data_LockStep_
}

type Data_LockStep_ interface {
  isLockStep()
}

type CompanionStruct_LockStep_ struct {}
var Companion_LockStep_ = CompanionStruct_LockStep_{}

type LockStep_GrantStep struct {
}

func (LockStep_GrantStep) isLockStep() {}

func (CompanionStruct_LockStep_) Create_GrantStep_() LockStep {
  return LockStep{LockStep_GrantStep{}}
}

func (_this LockStep) Is_GrantStep() bool {
  _, ok := _this.Get().(LockStep_GrantStep)
return ok
}

type LockStep_AcceptStep struct {
}

func (LockStep_AcceptStep) isLockStep() {}

func (CompanionStruct_LockStep_) Create_AcceptStep_() LockStep {
  return LockStep{LockStep_AcceptStep{}}
}

func (_this LockStep) Is_AcceptStep() bool {
  _, ok := _this.Get().(LockStep_AcceptStep)
return ok
}

func (_ CompanionStruct_LockStep_) AllSingletonConstructors() _dafny.Iterator {
  i := -1
return func() (interface{}, bool) {
    i++
switch i {
      case 0: return Companion_LockStep_.Create_GrantStep_(), true
case 1: return Companion_LockStep_.Create_AcceptStep_(), true
default: return LockStep{}, false
    }
  }
}

func (_this LockStep) String() string {
  switch _this.Get().(type) {
    case nil: return "null"
case LockStep_GrantStep: {
      return "_33_Types__i_Compile.LockStep.GrantStep"
    }
    case LockStep_AcceptStep: {
      return "_33_Types__i_Compile.LockStep.AcceptStep"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this LockStep) Equals(other LockStep) bool {
  switch _this.Get().(type) {
    case LockStep_GrantStep: {
      _, ok := other.Get().(LockStep_GrantStep)
return ok
    }
    case LockStep_AcceptStep: {
      _, ok := other.Get().(LockStep_AcceptStep)
return ok
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this LockStep) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(LockStep)
return ok && _this.Equals(typed)
}
func Type_LockStep_() _dafny.Type {
  return type_LockStep_{}
}

type type_LockStep_ struct {
}

func (_this type_LockStep_) Default() interface{} {
  return LockStep{LockStep_GrantStep{}}
}

func (_this type_LockStep_) String() string {
  return "_33_Types__i_Compile.LockStep"
}
// End of data type LockStep




