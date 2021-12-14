// Package _251_LiveRSL____ElectionState__i_Compile
// Dafny module _251_LiveRSL____ElectionState__i_Compile compiled into Go

package _251_LiveRSL____ElectionState__i_Compile

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
_243_LiveRSL____PaxosWorldState__i_Compile "243_LiveRSL____PaxosWorldState__i_Compile_"
_245_LiveRSL____ReplicaConstantsState__i_Compile "245_LiveRSL____ReplicaConstantsState__i_Compile_"
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
var _ _243_LiveRSL____PaxosWorldState__i_Compile.Dummy__
var _ _245_LiveRSL____ReplicaConstantsState__i_Compile.Dummy__

type Dummy__ struct{}





// Definition of data type CRequestHeader
type CRequestHeader struct {
  Data_CRequestHeader_
}

func (_this CRequestHeader) Get() Data_CRequestHeader_ {
  return _this.Data_CRequestHeader_
}

type Data_CRequestHeader_ interface {
  isCRequestHeader()
}

type CompanionStruct_CRequestHeader_ struct {}
var Companion_CRequestHeader_ = CompanionStruct_CRequestHeader_{}

type CRequestHeader_CRequestHeader struct {
  Client _9_Native____Io__s_Compile.EndPoint
Seqno uint64
}

func (CRequestHeader_CRequestHeader) isCRequestHeader() {}

func (CompanionStruct_CRequestHeader_) Create_CRequestHeader_(Client _9_Native____Io__s_Compile.EndPoint, Seqno uint64) CRequestHeader {
  return CRequestHeader{CRequestHeader_CRequestHeader{Client,Seqno}}
}

func (_this CRequestHeader) Is_CRequestHeader() bool {
  _, ok := _this.Get().(CRequestHeader_CRequestHeader)
return ok
}

func (_this CRequestHeader) Dtor_client() _9_Native____Io__s_Compile.EndPoint {
  return _this.Get().(CRequestHeader_CRequestHeader).Client
}

func (_this CRequestHeader) Dtor_seqno() uint64 {
  return _this.Get().(CRequestHeader_CRequestHeader).Seqno
}

func (_this CRequestHeader) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case CRequestHeader_CRequestHeader: {
      return "_251_LiveRSL____ElectionState__i_Compile.CRequestHeader.CRequestHeader" + "(" + _dafny.String(data.Client) + ", " + _dafny.String(data.Seqno) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this CRequestHeader) Equals(other CRequestHeader) bool {
  switch data1 := _this.Get().(type) {
    case CRequestHeader_CRequestHeader: {
      data2, ok := other.Get().(CRequestHeader_CRequestHeader)
return ok && data1.Client.Equals(data2.Client) && data1.Seqno == data2.Seqno
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this CRequestHeader) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(CRequestHeader)
return ok && _this.Equals(typed)
}
func Type_CRequestHeader_() _dafny.Type {
  return type_CRequestHeader_{}
}

type type_CRequestHeader_ struct {
}

func (_this type_CRequestHeader_) Default() interface{} {
  return CRequestHeader{CRequestHeader_CRequestHeader{_9_Native____Io__s_Compile.Type_EndPoint_().Default().(_9_Native____Io__s_Compile.EndPoint), 0}}
}

func (_this type_CRequestHeader_) String() string {
  return "_251_LiveRSL____ElectionState__i_Compile.CRequestHeader"
}
// End of data type CRequestHeader

// Definition of data type CElectionState
type CElectionState struct {
  Data_CElectionState_
}

func (_this CElectionState) Get() Data_CElectionState_ {
  return _this.Data_CElectionState_
}

type Data_CElectionState_ interface {
  isCElectionState()
}

type CompanionStruct_CElectionState_ struct {}
var Companion_CElectionState_ = CompanionStruct_CElectionState_{}

type CElectionState_CElectionState struct {
  Constants _245_LiveRSL____ReplicaConstantsState__i_Compile.ReplicaConstantsState
Current__view _214_LiveRSL____CTypes__i_Compile.CBallot
Current__view__suspectors _dafny.Seq
Epoch__end__time uint64
Epoch__length uint64
Requests__received__this__epoch _dafny.Seq
Requests__received__prev__epochs _dafny.Seq
}

func (CElectionState_CElectionState) isCElectionState() {}

func (CompanionStruct_CElectionState_) Create_CElectionState_(Constants _245_LiveRSL____ReplicaConstantsState__i_Compile.ReplicaConstantsState, Current__view _214_LiveRSL____CTypes__i_Compile.CBallot, Current__view__suspectors _dafny.Seq, Epoch__end__time uint64, Epoch__length uint64, Requests__received__this__epoch _dafny.Seq, Requests__received__prev__epochs _dafny.Seq) CElectionState {
  return CElectionState{CElectionState_CElectionState{Constants,Current__view,Current__view__suspectors,Epoch__end__time,Epoch__length,Requests__received__this__epoch,Requests__received__prev__epochs}}
}

func (_this CElectionState) Is_CElectionState() bool {
  _, ok := _this.Get().(CElectionState_CElectionState)
return ok
}

func (_this CElectionState) Dtor_constants() _245_LiveRSL____ReplicaConstantsState__i_Compile.ReplicaConstantsState {
  return _this.Get().(CElectionState_CElectionState).Constants
}

func (_this CElectionState) Dtor_current__view() _214_LiveRSL____CTypes__i_Compile.CBallot {
  return _this.Get().(CElectionState_CElectionState).Current__view
}

func (_this CElectionState) Dtor_current__view__suspectors() _dafny.Seq {
  return _this.Get().(CElectionState_CElectionState).Current__view__suspectors
}

func (_this CElectionState) Dtor_epoch__end__time() uint64 {
  return _this.Get().(CElectionState_CElectionState).Epoch__end__time
}

func (_this CElectionState) Dtor_epoch__length() uint64 {
  return _this.Get().(CElectionState_CElectionState).Epoch__length
}

func (_this CElectionState) Dtor_requests__received__this__epoch() _dafny.Seq {
  return _this.Get().(CElectionState_CElectionState).Requests__received__this__epoch
}

func (_this CElectionState) Dtor_requests__received__prev__epochs() _dafny.Seq {
  return _this.Get().(CElectionState_CElectionState).Requests__received__prev__epochs
}

func (_this CElectionState) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case CElectionState_CElectionState: {
      return "_251_LiveRSL____ElectionState__i_Compile.CElectionState.CElectionState" + "(" + _dafny.String(data.Constants) + ", " + _dafny.String(data.Current__view) + ", " + _dafny.String(data.Current__view__suspectors) + ", " + _dafny.String(data.Epoch__end__time) + ", " + _dafny.String(data.Epoch__length) + ", " + _dafny.String(data.Requests__received__this__epoch) + ", " + _dafny.String(data.Requests__received__prev__epochs) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this CElectionState) Equals(other CElectionState) bool {
  switch data1 := _this.Get().(type) {
    case CElectionState_CElectionState: {
      data2, ok := other.Get().(CElectionState_CElectionState)
return ok && data1.Constants.Equals(data2.Constants) && data1.Current__view.Equals(data2.Current__view) && data1.Current__view__suspectors.Equals(data2.Current__view__suspectors) && data1.Epoch__end__time == data2.Epoch__end__time && data1.Epoch__length == data2.Epoch__length && data1.Requests__received__this__epoch.Equals(data2.Requests__received__this__epoch) && data1.Requests__received__prev__epochs.Equals(data2.Requests__received__prev__epochs)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this CElectionState) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(CElectionState)
return ok && _this.Equals(typed)
}
func Type_CElectionState_() _dafny.Type {
  return type_CElectionState_{}
}

type type_CElectionState_ struct {
}

func (_this type_CElectionState_) Default() interface{} {
  return CElectionState{CElectionState_CElectionState{_245_LiveRSL____ReplicaConstantsState__i_Compile.Type_ReplicaConstantsState_().Default().(_245_LiveRSL____ReplicaConstantsState__i_Compile.ReplicaConstantsState), _214_LiveRSL____CTypes__i_Compile.Type_CBallot_().Default().(_214_LiveRSL____CTypes__i_Compile.CBallot), _dafny.EmptySeq, 0, 0, _dafny.EmptySeq, _dafny.EmptySeq}}
}

func (_this type_CElectionState_) String() string {
  return "_251_LiveRSL____ElectionState__i_Compile.CElectionState"
}
// End of data type CElectionState

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
  return "_251_LiveRSL____ElectionState__i_Compile.Default__"
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
  return "_251_LiveRSL____ElectionState__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) CRequestsMatch(r1 _214_LiveRSL____CTypes__i_Compile.CRequest, r2 _214_LiveRSL____CTypes__i_Compile.CRequest) bool {
  return (((r1).Dtor_client()).Equals((r2).Dtor_client())) && (((r1).Dtor_seqno()) == ((r2).Dtor_seqno()))
}
func (_this *CompanionStruct_Default___) CRequestSatisfiedBy(r1 _214_LiveRSL____CTypes__i_Compile.CRequest, r2 _214_LiveRSL____CTypes__i_Compile.CRequest) bool {
  return (((r1).Dtor_client()).Equals((r2).Dtor_client())) && (((r1).Dtor_seqno()) <= ((r2).Dtor_seqno()))
}
// End of class Default__
