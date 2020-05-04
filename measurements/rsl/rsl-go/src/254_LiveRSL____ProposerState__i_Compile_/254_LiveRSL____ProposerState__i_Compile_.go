// Package _254_LiveRSL____ProposerState__i_Compile
// Dafny module _254_LiveRSL____ProposerState__i_Compile compiled into Go

package _254_LiveRSL____ProposerState__i_Compile

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
_251_LiveRSL____ElectionState__i_Compile "251_LiveRSL____ElectionState__i_Compile_"
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
var _ _251_LiveRSL____ElectionState__i_Compile.Dummy__

type Dummy__ struct{}





// Definition of data type CIncompleteBatchTimer
type CIncompleteBatchTimer struct {
  Data_CIncompleteBatchTimer_
}

func (_this CIncompleteBatchTimer) Get() Data_CIncompleteBatchTimer_ {
  return _this.Data_CIncompleteBatchTimer_
}

type Data_CIncompleteBatchTimer_ interface {
  isCIncompleteBatchTimer()
}

type CompanionStruct_CIncompleteBatchTimer_ struct {}
var Companion_CIncompleteBatchTimer_ = CompanionStruct_CIncompleteBatchTimer_{}

type CIncompleteBatchTimer_CIncompleteBatchTimerOn struct {
  When uint64
}

func (CIncompleteBatchTimer_CIncompleteBatchTimerOn) isCIncompleteBatchTimer() {}

func (CompanionStruct_CIncompleteBatchTimer_) Create_CIncompleteBatchTimerOn_(When uint64) CIncompleteBatchTimer {
  return CIncompleteBatchTimer{CIncompleteBatchTimer_CIncompleteBatchTimerOn{When}}
}

func (_this CIncompleteBatchTimer) Is_CIncompleteBatchTimerOn() bool {
  _, ok := _this.Get().(CIncompleteBatchTimer_CIncompleteBatchTimerOn)
return ok
}

type CIncompleteBatchTimer_CIncompleteBatchTimerOff struct {
}

func (CIncompleteBatchTimer_CIncompleteBatchTimerOff) isCIncompleteBatchTimer() {}

func (CompanionStruct_CIncompleteBatchTimer_) Create_CIncompleteBatchTimerOff_() CIncompleteBatchTimer {
  return CIncompleteBatchTimer{CIncompleteBatchTimer_CIncompleteBatchTimerOff{}}
}

func (_this CIncompleteBatchTimer) Is_CIncompleteBatchTimerOff() bool {
  _, ok := _this.Get().(CIncompleteBatchTimer_CIncompleteBatchTimerOff)
return ok
}

func (_this CIncompleteBatchTimer) Dtor_when() uint64 {
  return _this.Get().(CIncompleteBatchTimer_CIncompleteBatchTimerOn).When
}

func (_this CIncompleteBatchTimer) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case CIncompleteBatchTimer_CIncompleteBatchTimerOn: {
      return "_254_LiveRSL____ProposerState__i_Compile.CIncompleteBatchTimer.CIncompleteBatchTimerOn" + "(" + _dafny.String(data.When) + ")"
    }
    case CIncompleteBatchTimer_CIncompleteBatchTimerOff: {
      return "_254_LiveRSL____ProposerState__i_Compile.CIncompleteBatchTimer.CIncompleteBatchTimerOff"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this CIncompleteBatchTimer) Equals(other CIncompleteBatchTimer) bool {
  switch data1 := _this.Get().(type) {
    case CIncompleteBatchTimer_CIncompleteBatchTimerOn: {
      data2, ok := other.Get().(CIncompleteBatchTimer_CIncompleteBatchTimerOn)
return ok && data1.When == data2.When
    }
    case CIncompleteBatchTimer_CIncompleteBatchTimerOff: {
      _, ok := other.Get().(CIncompleteBatchTimer_CIncompleteBatchTimerOff)
return ok
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this CIncompleteBatchTimer) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(CIncompleteBatchTimer)
return ok && _this.Equals(typed)
}
func Type_CIncompleteBatchTimer_() _dafny.Type {
  return type_CIncompleteBatchTimer_{}
}

type type_CIncompleteBatchTimer_ struct {
}

func (_this type_CIncompleteBatchTimer_) Default() interface{} {
  return CIncompleteBatchTimer{CIncompleteBatchTimer_CIncompleteBatchTimerOn{0}}
}

func (_this type_CIncompleteBatchTimer_) String() string {
  return "_254_LiveRSL____ProposerState__i_Compile.CIncompleteBatchTimer"
}
// End of data type CIncompleteBatchTimer

// Definition of data type ProposerState
type ProposerState struct {
  Data_ProposerState_
}

func (_this ProposerState) Get() Data_ProposerState_ {
  return _this.Data_ProposerState_
}

type Data_ProposerState_ interface {
  isProposerState()
}

type CompanionStruct_ProposerState_ struct {}
var Companion_ProposerState_ = CompanionStruct_ProposerState_{}

type ProposerState_ProposerState struct {
  Constants _245_LiveRSL____ReplicaConstantsState__i_Compile.ReplicaConstantsState
Current__state uint8
Request__queue _dafny.Seq
Max__ballot__i__sent__1a _214_LiveRSL____CTypes__i_Compile.CBallot
Next__operation__number__to__propose uint64
Received__1b__packets _dafny.Set
Highest__seqno__requested__by__client__this__view _dafny.Map
Incomplete__batch__timer _254_LiveRSL____ProposerState__i_Compile.CIncompleteBatchTimer
Election__state _251_LiveRSL____ElectionState__i_Compile.CElectionState
MaxOpnWithProposal _214_LiveRSL____CTypes__i_Compile.COperationNumber
MaxLogTruncationPoint _214_LiveRSL____CTypes__i_Compile.COperationNumber
}

func (ProposerState_ProposerState) isProposerState() {}

func (CompanionStruct_ProposerState_) Create_ProposerState_(Constants _245_LiveRSL____ReplicaConstantsState__i_Compile.ReplicaConstantsState, Current__state uint8, Request__queue _dafny.Seq, Max__ballot__i__sent__1a _214_LiveRSL____CTypes__i_Compile.CBallot, Next__operation__number__to__propose uint64, Received__1b__packets _dafny.Set, Highest__seqno__requested__by__client__this__view _dafny.Map, Incomplete__batch__timer _254_LiveRSL____ProposerState__i_Compile.CIncompleteBatchTimer, Election__state _251_LiveRSL____ElectionState__i_Compile.CElectionState, MaxOpnWithProposal _214_LiveRSL____CTypes__i_Compile.COperationNumber, MaxLogTruncationPoint _214_LiveRSL____CTypes__i_Compile.COperationNumber) ProposerState {
  return ProposerState{ProposerState_ProposerState{Constants,Current__state,Request__queue,Max__ballot__i__sent__1a,Next__operation__number__to__propose,Received__1b__packets,Highest__seqno__requested__by__client__this__view,Incomplete__batch__timer,Election__state,MaxOpnWithProposal,MaxLogTruncationPoint}}
}

func (_this ProposerState) Is_ProposerState() bool {
  _, ok := _this.Get().(ProposerState_ProposerState)
return ok
}

func (_this ProposerState) Dtor_constants() _245_LiveRSL____ReplicaConstantsState__i_Compile.ReplicaConstantsState {
  return _this.Get().(ProposerState_ProposerState).Constants
}

func (_this ProposerState) Dtor_current__state() uint8 {
  return _this.Get().(ProposerState_ProposerState).Current__state
}

func (_this ProposerState) Dtor_request__queue() _dafny.Seq {
  return _this.Get().(ProposerState_ProposerState).Request__queue
}

func (_this ProposerState) Dtor_max__ballot__i__sent__1a() _214_LiveRSL____CTypes__i_Compile.CBallot {
  return _this.Get().(ProposerState_ProposerState).Max__ballot__i__sent__1a
}

func (_this ProposerState) Dtor_next__operation__number__to__propose() uint64 {
  return _this.Get().(ProposerState_ProposerState).Next__operation__number__to__propose
}

func (_this ProposerState) Dtor_received__1b__packets() _dafny.Set {
  return _this.Get().(ProposerState_ProposerState).Received__1b__packets
}

func (_this ProposerState) Dtor_highest__seqno__requested__by__client__this__view() _dafny.Map {
  return _this.Get().(ProposerState_ProposerState).Highest__seqno__requested__by__client__this__view
}

func (_this ProposerState) Dtor_incomplete__batch__timer() _254_LiveRSL____ProposerState__i_Compile.CIncompleteBatchTimer {
  return _this.Get().(ProposerState_ProposerState).Incomplete__batch__timer
}

func (_this ProposerState) Dtor_election__state() _251_LiveRSL____ElectionState__i_Compile.CElectionState {
  return _this.Get().(ProposerState_ProposerState).Election__state
}

func (_this ProposerState) Dtor_maxOpnWithProposal() _214_LiveRSL____CTypes__i_Compile.COperationNumber {
  return _this.Get().(ProposerState_ProposerState).MaxOpnWithProposal
}

func (_this ProposerState) Dtor_maxLogTruncationPoint() _214_LiveRSL____CTypes__i_Compile.COperationNumber {
  return _this.Get().(ProposerState_ProposerState).MaxLogTruncationPoint
}

func (_this ProposerState) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case ProposerState_ProposerState: {
      return "_254_LiveRSL____ProposerState__i_Compile.ProposerState.ProposerState" + "(" + _dafny.String(data.Constants) + ", " + _dafny.String(data.Current__state) + ", " + _dafny.String(data.Request__queue) + ", " + _dafny.String(data.Max__ballot__i__sent__1a) + ", " + _dafny.String(data.Next__operation__number__to__propose) + ", " + _dafny.String(data.Received__1b__packets) + ", " + _dafny.String(data.Highest__seqno__requested__by__client__this__view) + ", " + _dafny.String(data.Incomplete__batch__timer) + ", " + _dafny.String(data.Election__state) + ", " + _dafny.String(data.MaxOpnWithProposal) + ", " + _dafny.String(data.MaxLogTruncationPoint) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this ProposerState) Equals(other ProposerState) bool {
  switch data1 := _this.Get().(type) {
    case ProposerState_ProposerState: {
      data2, ok := other.Get().(ProposerState_ProposerState)
return ok && data1.Constants.Equals(data2.Constants) && data1.Current__state == data2.Current__state && data1.Request__queue.Equals(data2.Request__queue) && data1.Max__ballot__i__sent__1a.Equals(data2.Max__ballot__i__sent__1a) && data1.Next__operation__number__to__propose == data2.Next__operation__number__to__propose && data1.Received__1b__packets.Equals(data2.Received__1b__packets) && data1.Highest__seqno__requested__by__client__this__view.Equals(data2.Highest__seqno__requested__by__client__this__view) && data1.Incomplete__batch__timer.Equals(data2.Incomplete__batch__timer) && data1.Election__state.Equals(data2.Election__state) && data1.MaxOpnWithProposal.Equals(data2.MaxOpnWithProposal) && data1.MaxLogTruncationPoint.Equals(data2.MaxLogTruncationPoint)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this ProposerState) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(ProposerState)
return ok && _this.Equals(typed)
}
func Type_ProposerState_() _dafny.Type {
  return type_ProposerState_{}
}

type type_ProposerState_ struct {
}

func (_this type_ProposerState_) Default() interface{} {
  return ProposerState{ProposerState_ProposerState{_245_LiveRSL____ReplicaConstantsState__i_Compile.Type_ReplicaConstantsState_().Default().(_245_LiveRSL____ReplicaConstantsState__i_Compile.ReplicaConstantsState), 0, _dafny.EmptySeq, _214_LiveRSL____CTypes__i_Compile.Type_CBallot_().Default().(_214_LiveRSL____CTypes__i_Compile.CBallot), 0, _dafny.EmptySet, _dafny.EmptyMap, Type_CIncompleteBatchTimer_().Default().(_254_LiveRSL____ProposerState__i_Compile.CIncompleteBatchTimer), _251_LiveRSL____ElectionState__i_Compile.Type_CElectionState_().Default().(_251_LiveRSL____ElectionState__i_Compile.CElectionState), _214_LiveRSL____CTypes__i_Compile.Type_COperationNumber_().Default().(_214_LiveRSL____CTypes__i_Compile.COperationNumber), _214_LiveRSL____CTypes__i_Compile.Type_COperationNumber_().Default().(_214_LiveRSL____CTypes__i_Compile.COperationNumber)}}
}

func (_this type_ProposerState_) String() string {
  return "_254_LiveRSL____ProposerState__i_Compile.ProposerState"
}
// End of data type ProposerState

