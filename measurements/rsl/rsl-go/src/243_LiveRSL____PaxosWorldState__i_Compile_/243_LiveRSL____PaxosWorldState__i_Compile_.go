// Package _243_LiveRSL____PaxosWorldState__i_Compile
// Dafny module _243_LiveRSL____PaxosWorldState__i_Compile compiled into Go

package _243_LiveRSL____PaxosWorldState__i_Compile

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
_78_LiveRSL____Constants__i_Compile "78_LiveRSL____Constants__i_Compile_"
_85_LiveRSL____Broadcast__i_Compile "85_LiveRSL____Broadcast__i_Compile_"
_91_Collections____CountMatches__i_Compile "91_Collections____CountMatches__i_Compile_"
_93_LiveRSL____Acceptor__i_Compile "93_LiveRSL____Acceptor__i_Compile_"
_99_LiveRSL____Election__i_Compile "99_LiveRSL____Election__i_Compile_"
_101_LiveRSL____Proposer__i_Compile "101_LiveRSL____Proposer__i_Compile_"
_115_LiveRSL____StateMachine__i_Compile "115_LiveRSL____StateMachine__i_Compile_"
_118_Collections____Maps__i_Compile "118_Collections____Maps__i_Compile_"
_120_LiveRSL____Executor__i_Compile "120_LiveRSL____Executor__i_Compile_"
_123_LiveRSL____Learner__i_Compile "123_LiveRSL____Learner__i_Compile_"
_126_LiveRSL____Replica__i_Compile "126_LiveRSL____Replica__i_Compile_"
_135_Logic____Option__i_Compile "135_Logic____Option__i_Compile_"
_138_Native____NativeTypes__i_Compile "138_Native____NativeTypes__i_Compile_"
_141_Libraries____base__s_Compile "141_Libraries____base__s_Compile_"
_143_Math____power2__s_Compile "143_Math____power2__s_Compile_"
_145_Math____power__s_Compile "145_Math____power__s_Compile_"
_149_Math____power__i_Compile "149_Math____power__i_Compile_"
_153_Math____div__def__i_Compile "153_Math____div__def__i_Compile_"
_157_Math____div__boogie__i_Compile "157_Math____div__boogie__i_Compile_"
_162_Math____div__auto__proofs__i_Compile "162_Math____div__auto__proofs__i_Compile_"
_164_Math____div__auto__i_Compile "164_Math____div__auto__i_Compile_"
_166_Math____div__i_Compile "166_Math____div__i_Compile_"
_168_Math____power2__i_Compile "168_Math____power2__i_Compile_"
_170_Common____Util__i_Compile "170_Common____Util__i_Compile_"
_174_Common____MarshallInt__i_Compile "174_Common____MarshallInt__i_Compile_"
_176_Common____GenericMarshalling__i_Compile "176_Common____GenericMarshalling__i_Compile_"
_180_Common____UdpClient__i_Compile "180_Common____UdpClient__i_Compile_"
_182_Common____SeqIsUniqueDef__i_Compile "182_Common____SeqIsUniqueDef__i_Compile_"
_185_Common____SeqIsUnique__i_Compile "185_Common____SeqIsUnique__i_Compile_"
_191_GenericRefinement__i_Compile "191_GenericRefinement__i_Compile_"
_194_Common____NodeIdentity__i_Compile "194_Common____NodeIdentity__i_Compile_"
_197_LiveRSL____AppInterface__i_Compile "197_LiveRSL____AppInterface__i_Compile_"
_214_LiveRSL____CTypes__i_Compile "214_LiveRSL____CTypes__i_Compile_"
_217_LiveRSL____CMessage__i_Compile "217_LiveRSL____CMessage__i_Compile_"
_225_LiveRSL____CMessageRefinements__i_Compile "225_LiveRSL____CMessageRefinements__i_Compile_"
_228_LiveRSL____PacketParsing__i_Compile "228_LiveRSL____PacketParsing__i_Compile_"
_234_Common____UpperBound__i_Compile "234_Common____UpperBound__i_Compile_"
_236_LiveRSL____ParametersState__i_Compile "236_LiveRSL____ParametersState__i_Compile_"
_238_LiveRSL____CPaxosConfiguration__i_Compile "238_LiveRSL____CPaxosConfiguration__i_Compile_"
_240_LiveRSL____ConstantsState__i_Compile "240_LiveRSL____ConstantsState__i_Compile_"
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
var _ _78_LiveRSL____Constants__i_Compile.Dummy__
var _ _85_LiveRSL____Broadcast__i_Compile.Dummy__
var _ _91_Collections____CountMatches__i_Compile.Dummy__
var _ _93_LiveRSL____Acceptor__i_Compile.Dummy__
var _ _99_LiveRSL____Election__i_Compile.Dummy__
var _ _101_LiveRSL____Proposer__i_Compile.Dummy__
var _ _115_LiveRSL____StateMachine__i_Compile.Dummy__
var _ _118_Collections____Maps__i_Compile.Dummy__
var _ _120_LiveRSL____Executor__i_Compile.Dummy__
var _ _123_LiveRSL____Learner__i_Compile.Dummy__
var _ _126_LiveRSL____Replica__i_Compile.Dummy__
var _ _135_Logic____Option__i_Compile.Dummy__
var _ _138_Native____NativeTypes__i_Compile.Dummy__
var _ _141_Libraries____base__s_Compile.Dummy__
var _ _143_Math____power2__s_Compile.Dummy__
var _ _145_Math____power__s_Compile.Dummy__
var _ _149_Math____power__i_Compile.Dummy__
var _ _153_Math____div__def__i_Compile.Dummy__
var _ _157_Math____div__boogie__i_Compile.Dummy__
var _ _162_Math____div__auto__proofs__i_Compile.Dummy__
var _ _164_Math____div__auto__i_Compile.Dummy__
var _ _166_Math____div__i_Compile.Dummy__
var _ _168_Math____power2__i_Compile.Dummy__
var _ _170_Common____Util__i_Compile.Dummy__
var _ _174_Common____MarshallInt__i_Compile.Dummy__
var _ _176_Common____GenericMarshalling__i_Compile.Dummy__
var _ _180_Common____UdpClient__i_Compile.Dummy__
var _ _182_Common____SeqIsUniqueDef__i_Compile.Dummy__
var _ _185_Common____SeqIsUnique__i_Compile.Dummy__
var _ _191_GenericRefinement__i_Compile.Dummy__
var _ _194_Common____NodeIdentity__i_Compile.Dummy__
var _ _197_LiveRSL____AppInterface__i_Compile.Dummy__
var _ _214_LiveRSL____CTypes__i_Compile.Dummy__
var _ _217_LiveRSL____CMessage__i_Compile.Dummy__
var _ _225_LiveRSL____CMessageRefinements__i_Compile.Dummy__
var _ _228_LiveRSL____PacketParsing__i_Compile.Dummy__
var _ _234_Common____UpperBound__i_Compile.Dummy__
var _ _236_LiveRSL____ParametersState__i_Compile.Dummy__
var _ _238_LiveRSL____CPaxosConfiguration__i_Compile.Dummy__
var _ _240_LiveRSL____ConstantsState__i_Compile.Dummy__

type Dummy__ struct{}



// Definition of data type ActionStatus
type ActionStatus struct {
  Data_ActionStatus_
}

func (_this ActionStatus) Get() Data_ActionStatus_ {
  return _this.Data_ActionStatus_
}

type Data_ActionStatus_ interface {
  isActionStatus()
}

type CompanionStruct_ActionStatus_ struct {}
var Companion_ActionStatus_ = CompanionStruct_ActionStatus_{}

type ActionStatus_Ok struct {
}

func (ActionStatus_Ok) isActionStatus() {}

func (CompanionStruct_ActionStatus_) Create_Ok_() ActionStatus {
  return ActionStatus{ActionStatus_Ok{}}
}

func (_this ActionStatus) Is_Ok() bool {
  _, ok := _this.Get().(ActionStatus_Ok)
return ok
}

type ActionStatus_Ignore struct {
}

func (ActionStatus_Ignore) isActionStatus() {}

func (CompanionStruct_ActionStatus_) Create_Ignore_() ActionStatus {
  return ActionStatus{ActionStatus_Ignore{}}
}

func (_this ActionStatus) Is_Ignore() bool {
  _, ok := _this.Get().(ActionStatus_Ignore)
return ok
}

type ActionStatus_Fail struct {
}

func (ActionStatus_Fail) isActionStatus() {}

func (CompanionStruct_ActionStatus_) Create_Fail_() ActionStatus {
  return ActionStatus{ActionStatus_Fail{}}
}

func (_this ActionStatus) Is_Fail() bool {
  _, ok := _this.Get().(ActionStatus_Fail)
return ok
}

func (_ CompanionStruct_ActionStatus_) AllSingletonConstructors() _dafny.Iterator {
  i := -1
return func() (interface{}, bool) {
    i++
switch i {
      case 0: return Companion_ActionStatus_.Create_Ok_(), true
case 1: return Companion_ActionStatus_.Create_Ignore_(), true
case 2: return Companion_ActionStatus_.Create_Fail_(), true
default: return ActionStatus{}, false
    }
  }
}

func (_this ActionStatus) String() string {
  switch _this.Get().(type) {
    case nil: return "null"
case ActionStatus_Ok: {
      return "_243_LiveRSL____PaxosWorldState__i_Compile.ActionStatus.Ok"
    }
    case ActionStatus_Ignore: {
      return "_243_LiveRSL____PaxosWorldState__i_Compile.ActionStatus.Ignore"
    }
    case ActionStatus_Fail: {
      return "_243_LiveRSL____PaxosWorldState__i_Compile.ActionStatus.Fail"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this ActionStatus) Equals(other ActionStatus) bool {
  switch _this.Get().(type) {
    case ActionStatus_Ok: {
      _, ok := other.Get().(ActionStatus_Ok)
return ok
    }
    case ActionStatus_Ignore: {
      _, ok := other.Get().(ActionStatus_Ignore)
return ok
    }
    case ActionStatus_Fail: {
      _, ok := other.Get().(ActionStatus_Fail)
return ok
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this ActionStatus) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(ActionStatus)
return ok && _this.Equals(typed)
}
func Type_ActionStatus_() _dafny.Type {
  return type_ActionStatus_{}
}

type type_ActionStatus_ struct {
}

func (_this type_ActionStatus_) Default() interface{} {
  return ActionStatus{ActionStatus_Ok{}}
}

func (_this type_ActionStatus_) String() string {
  return "_243_LiveRSL____PaxosWorldState__i_Compile.ActionStatus"
}
// End of data type ActionStatus

// Definition of data type PaxosWorldState
type PaxosWorldState struct {
  Data_PaxosWorldState_
}

func (_this PaxosWorldState) Get() Data_PaxosWorldState_ {
  return _this.Data_PaxosWorldState_
}

type Data_PaxosWorldState_ interface {
  isPaxosWorldState()
}

type CompanionStruct_PaxosWorldState_ struct {}
var Companion_PaxosWorldState_ = CompanionStruct_PaxosWorldState_{}

type PaxosWorldState_PaxosWorldState struct {
  Good bool
Config _238_LiveRSL____CPaxosConfiguration__i_Compile.CPaxosConfiguration
}

func (PaxosWorldState_PaxosWorldState) isPaxosWorldState() {}

func (CompanionStruct_PaxosWorldState_) Create_PaxosWorldState_(Good bool, Config _238_LiveRSL____CPaxosConfiguration__i_Compile.CPaxosConfiguration) PaxosWorldState {
  return PaxosWorldState{PaxosWorldState_PaxosWorldState{Good,Config}}
}

func (_this PaxosWorldState) Is_PaxosWorldState() bool {
  _, ok := _this.Get().(PaxosWorldState_PaxosWorldState)
return ok
}

func (_this PaxosWorldState) Dtor_good() bool {
  return _this.Get().(PaxosWorldState_PaxosWorldState).Good
}

func (_this PaxosWorldState) Dtor_config() _238_LiveRSL____CPaxosConfiguration__i_Compile.CPaxosConfiguration {
  return _this.Get().(PaxosWorldState_PaxosWorldState).Config
}

func (_this PaxosWorldState) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case PaxosWorldState_PaxosWorldState: {
      return "_243_LiveRSL____PaxosWorldState__i_Compile.PaxosWorldState.PaxosWorldState" + "(" + _dafny.String(data.Good) + ", " + _dafny.String(data.Config) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this PaxosWorldState) Equals(other PaxosWorldState) bool {
  switch data1 := _this.Get().(type) {
    case PaxosWorldState_PaxosWorldState: {
      data2, ok := other.Get().(PaxosWorldState_PaxosWorldState)
return ok && data1.Good == data2.Good && data1.Config.Equals(data2.Config)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this PaxosWorldState) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(PaxosWorldState)
return ok && _this.Equals(typed)
}
func Type_PaxosWorldState_() _dafny.Type {
  return type_PaxosWorldState_{}
}

type type_PaxosWorldState_ struct {
}

func (_this type_PaxosWorldState_) Default() interface{} {
  return PaxosWorldState{PaxosWorldState_PaxosWorldState{false, _238_LiveRSL____CPaxosConfiguration__i_Compile.Type_CPaxosConfiguration_().Default().(_238_LiveRSL____CPaxosConfiguration__i_Compile.CPaxosConfiguration)}}
}

func (_this type_PaxosWorldState_) String() string {
  return "_243_LiveRSL____PaxosWorldState__i_Compile.PaxosWorldState"
}
// End of data type PaxosWorldState

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
  return "_243_LiveRSL____PaxosWorldState__i_Compile.Default__"
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
  return "_243_LiveRSL____PaxosWorldState__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) F__max__uint64() uint64 {
  return uint64(18446744073709551615)
}
func (_this *CompanionStruct_Default___) UpdatePaxosWorld(world _243_LiveRSL____PaxosWorldState__i_Compile.PaxosWorldState, status _243_LiveRSL____PaxosWorldState__i_Compile.ActionStatus) _243_LiveRSL____PaxosWorldState__i_Compile.PaxosWorldState {
  goto TAIL_CALL_START
TAIL_CALL_START:
var world_k _243_LiveRSL____PaxosWorldState__i_Compile.PaxosWorldState = Type_PaxosWorldState_().Default().(_243_LiveRSL____PaxosWorldState__i_Compile.PaxosWorldState)
  var _ = world_k
  if ((status).Is_Fail()) {
    world_k = func (_pat_let0_0 _243_LiveRSL____PaxosWorldState__i_Compile.PaxosWorldState) _243_LiveRSL____PaxosWorldState__i_Compile.PaxosWorldState {
      return func (_4731_dt__update__tmp_h0 _243_LiveRSL____PaxosWorldState__i_Compile.PaxosWorldState) _243_LiveRSL____PaxosWorldState__i_Compile.PaxosWorldState {
        return func (_pat_let1_0 bool) _243_LiveRSL____PaxosWorldState__i_Compile.PaxosWorldState {
          return func (_4732_dt__update_hgood_h0 bool) _243_LiveRSL____PaxosWorldState__i_Compile.PaxosWorldState {
            return PaxosWorldState{PaxosWorldState_PaxosWorldState{_4732_dt__update_hgood_h0, (_4731_dt__update__tmp_h0).Dtor_config()}}
          }(_pat_let1_0)
        }(false)
      }(_pat_let0_0)
    }(world)
  } else {
    world_k = world
  }
  return world_k
}
// End of class Default__
