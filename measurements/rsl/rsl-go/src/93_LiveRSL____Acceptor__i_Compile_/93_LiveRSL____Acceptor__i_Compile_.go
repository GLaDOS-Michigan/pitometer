// Package _93_LiveRSL____Acceptor__i_Compile
// Dafny module _93_LiveRSL____Acceptor__i_Compile compiled into Go

package _93_LiveRSL____Acceptor__i_Compile

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

type Dummy__ struct{}







// Definition of data type LAcceptor
type LAcceptor struct {
  Data_LAcceptor_
}

func (_this LAcceptor) Get() Data_LAcceptor_ {
  return _this.Data_LAcceptor_
}

type Data_LAcceptor_ interface {
  isLAcceptor()
}

type CompanionStruct_LAcceptor_ struct {}
var Companion_LAcceptor_ = CompanionStruct_LAcceptor_{}

type LAcceptor_LAcceptor struct {
  Constants _78_LiveRSL____Constants__i_Compile.LReplicaConstants
Max__bal _56_LiveRSL____Types__i_Compile.Ballot
Votes _dafny.Map
Last__checkpointed__operation _dafny.Seq
Log__truncation__point _dafny.Int
}

func (LAcceptor_LAcceptor) isLAcceptor() {}

func (CompanionStruct_LAcceptor_) Create_LAcceptor_(Constants _78_LiveRSL____Constants__i_Compile.LReplicaConstants, Max__bal _56_LiveRSL____Types__i_Compile.Ballot, Votes _dafny.Map, Last__checkpointed__operation _dafny.Seq, Log__truncation__point _dafny.Int) LAcceptor {
  return LAcceptor{LAcceptor_LAcceptor{Constants,Max__bal,Votes,Last__checkpointed__operation,Log__truncation__point}}
}

func (_this LAcceptor) Is_LAcceptor() bool {
  _, ok := _this.Get().(LAcceptor_LAcceptor)
return ok
}

func (_this LAcceptor) Dtor_constants() _78_LiveRSL____Constants__i_Compile.LReplicaConstants {
  return _this.Get().(LAcceptor_LAcceptor).Constants
}

func (_this LAcceptor) Dtor_max__bal() _56_LiveRSL____Types__i_Compile.Ballot {
  return _this.Get().(LAcceptor_LAcceptor).Max__bal
}

func (_this LAcceptor) Dtor_votes() _dafny.Map {
  return _this.Get().(LAcceptor_LAcceptor).Votes
}

func (_this LAcceptor) Dtor_last__checkpointed__operation() _dafny.Seq {
  return _this.Get().(LAcceptor_LAcceptor).Last__checkpointed__operation
}

func (_this LAcceptor) Dtor_log__truncation__point() _dafny.Int {
  return _this.Get().(LAcceptor_LAcceptor).Log__truncation__point
}

func (_this LAcceptor) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case LAcceptor_LAcceptor: {
      return "_93_LiveRSL____Acceptor__i_Compile.LAcceptor.LAcceptor" + "(" + _dafny.String(data.Constants) + ", " + _dafny.String(data.Max__bal) + ", " + _dafny.String(data.Votes) + ", " + _dafny.String(data.Last__checkpointed__operation) + ", " + _dafny.String(data.Log__truncation__point) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this LAcceptor) Equals(other LAcceptor) bool {
  switch data1 := _this.Get().(type) {
    case LAcceptor_LAcceptor: {
      data2, ok := other.Get().(LAcceptor_LAcceptor)
return ok && data1.Constants.Equals(data2.Constants) && data1.Max__bal.Equals(data2.Max__bal) && data1.Votes.Equals(data2.Votes) && data1.Last__checkpointed__operation.Equals(data2.Last__checkpointed__operation) && data1.Log__truncation__point.Cmp(data2.Log__truncation__point) == 0
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this LAcceptor) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(LAcceptor)
return ok && _this.Equals(typed)
}
func Type_LAcceptor_() _dafny.Type {
  return type_LAcceptor_{}
}

type type_LAcceptor_ struct {
}

func (_this type_LAcceptor_) Default() interface{} {
  return LAcceptor{LAcceptor_LAcceptor{_78_LiveRSL____Constants__i_Compile.Type_LReplicaConstants_().Default().(_78_LiveRSL____Constants__i_Compile.LReplicaConstants), _56_LiveRSL____Types__i_Compile.Type_Ballot_().Default().(_56_LiveRSL____Types__i_Compile.Ballot), _dafny.EmptyMap, _dafny.EmptySeq, _dafny.Zero}}
}

func (_this type_LAcceptor_) String() string {
  return "_93_LiveRSL____Acceptor__i_Compile.LAcceptor"
}
// End of data type LAcceptor

