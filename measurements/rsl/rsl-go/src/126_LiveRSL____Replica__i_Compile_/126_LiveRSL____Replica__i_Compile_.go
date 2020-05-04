// Package _126_LiveRSL____Replica__i_Compile
// Dafny module _126_LiveRSL____Replica__i_Compile compiled into Go

package _126_LiveRSL____Replica__i_Compile

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

type Dummy__ struct{}











// Definition of data type LReplica
type LReplica struct {
  Data_LReplica_
}

func (_this LReplica) Get() Data_LReplica_ {
  return _this.Data_LReplica_
}

type Data_LReplica_ interface {
  isLReplica()
}

type CompanionStruct_LReplica_ struct {}
var Companion_LReplica_ = CompanionStruct_LReplica_{}

type LReplica_LReplica struct {
  Constants _78_LiveRSL____Constants__i_Compile.LReplicaConstants
NextHeartbeatTime _dafny.Int
Proposer _101_LiveRSL____Proposer__i_Compile.LProposer
Acceptor _93_LiveRSL____Acceptor__i_Compile.LAcceptor
Learner _123_LiveRSL____Learner__i_Compile.LLearner
Executor _120_LiveRSL____Executor__i_Compile.LExecutor
}

func (LReplica_LReplica) isLReplica() {}

func (CompanionStruct_LReplica_) Create_LReplica_(Constants _78_LiveRSL____Constants__i_Compile.LReplicaConstants, NextHeartbeatTime _dafny.Int, Proposer _101_LiveRSL____Proposer__i_Compile.LProposer, Acceptor _93_LiveRSL____Acceptor__i_Compile.LAcceptor, Learner _123_LiveRSL____Learner__i_Compile.LLearner, Executor _120_LiveRSL____Executor__i_Compile.LExecutor) LReplica {
  return LReplica{LReplica_LReplica{Constants,NextHeartbeatTime,Proposer,Acceptor,Learner,Executor}}
}

func (_this LReplica) Is_LReplica() bool {
  _, ok := _this.Get().(LReplica_LReplica)
return ok
}

func (_this LReplica) Dtor_constants() _78_LiveRSL____Constants__i_Compile.LReplicaConstants {
  return _this.Get().(LReplica_LReplica).Constants
}

func (_this LReplica) Dtor_nextHeartbeatTime() _dafny.Int {
  return _this.Get().(LReplica_LReplica).NextHeartbeatTime
}

func (_this LReplica) Dtor_proposer() _101_LiveRSL____Proposer__i_Compile.LProposer {
  return _this.Get().(LReplica_LReplica).Proposer
}

func (_this LReplica) Dtor_acceptor() _93_LiveRSL____Acceptor__i_Compile.LAcceptor {
  return _this.Get().(LReplica_LReplica).Acceptor
}

func (_this LReplica) Dtor_learner() _123_LiveRSL____Learner__i_Compile.LLearner {
  return _this.Get().(LReplica_LReplica).Learner
}

func (_this LReplica) Dtor_executor() _120_LiveRSL____Executor__i_Compile.LExecutor {
  return _this.Get().(LReplica_LReplica).Executor
}

func (_this LReplica) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case LReplica_LReplica: {
      return "_126_LiveRSL____Replica__i_Compile.LReplica.LReplica" + "(" + _dafny.String(data.Constants) + ", " + _dafny.String(data.NextHeartbeatTime) + ", " + _dafny.String(data.Proposer) + ", " + _dafny.String(data.Acceptor) + ", " + _dafny.String(data.Learner) + ", " + _dafny.String(data.Executor) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this LReplica) Equals(other LReplica) bool {
  switch data1 := _this.Get().(type) {
    case LReplica_LReplica: {
      data2, ok := other.Get().(LReplica_LReplica)
return ok && data1.Constants.Equals(data2.Constants) && data1.NextHeartbeatTime.Cmp(data2.NextHeartbeatTime) == 0 && data1.Proposer.Equals(data2.Proposer) && data1.Acceptor.Equals(data2.Acceptor) && data1.Learner.Equals(data2.Learner) && data1.Executor.Equals(data2.Executor)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this LReplica) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(LReplica)
return ok && _this.Equals(typed)
}
func Type_LReplica_() _dafny.Type {
  return type_LReplica_{}
}

type type_LReplica_ struct {
}

func (_this type_LReplica_) Default() interface{} {
  return LReplica{LReplica_LReplica{_78_LiveRSL____Constants__i_Compile.Type_LReplicaConstants_().Default().(_78_LiveRSL____Constants__i_Compile.LReplicaConstants), _dafny.Zero, _101_LiveRSL____Proposer__i_Compile.Type_LProposer_().Default().(_101_LiveRSL____Proposer__i_Compile.LProposer), _93_LiveRSL____Acceptor__i_Compile.Type_LAcceptor_().Default().(_93_LiveRSL____Acceptor__i_Compile.LAcceptor), _123_LiveRSL____Learner__i_Compile.Type_LLearner_().Default().(_123_LiveRSL____Learner__i_Compile.LLearner), _120_LiveRSL____Executor__i_Compile.Type_LExecutor_().Default().(_120_LiveRSL____Executor__i_Compile.LExecutor)}}
}

func (_this type_LReplica_) String() string {
  return "_126_LiveRSL____Replica__i_Compile.LReplica"
}
// End of data type LReplica

// Definition of data type LScheduler
type LScheduler struct {
  Data_LScheduler_
}

func (_this LScheduler) Get() Data_LScheduler_ {
  return _this.Data_LScheduler_
}

type Data_LScheduler_ interface {
  isLScheduler()
}

type CompanionStruct_LScheduler_ struct {}
var Companion_LScheduler_ = CompanionStruct_LScheduler_{}

type LScheduler_LScheduler struct {
  Replica _126_LiveRSL____Replica__i_Compile.LReplica
NextActionIndex _dafny.Int
}

func (LScheduler_LScheduler) isLScheduler() {}

func (CompanionStruct_LScheduler_) Create_LScheduler_(Replica _126_LiveRSL____Replica__i_Compile.LReplica, NextActionIndex _dafny.Int) LScheduler {
  return LScheduler{LScheduler_LScheduler{Replica,NextActionIndex}}
}

func (_this LScheduler) Is_LScheduler() bool {
  _, ok := _this.Get().(LScheduler_LScheduler)
return ok
}

func (_this LScheduler) Dtor_replica() _126_LiveRSL____Replica__i_Compile.LReplica {
  return _this.Get().(LScheduler_LScheduler).Replica
}

func (_this LScheduler) Dtor_nextActionIndex() _dafny.Int {
  return _this.Get().(LScheduler_LScheduler).NextActionIndex
}

func (_this LScheduler) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case LScheduler_LScheduler: {
      return "_126_LiveRSL____Replica__i_Compile.LScheduler.LScheduler" + "(" + _dafny.String(data.Replica) + ", " + _dafny.String(data.NextActionIndex) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this LScheduler) Equals(other LScheduler) bool {
  switch data1 := _this.Get().(type) {
    case LScheduler_LScheduler: {
      data2, ok := other.Get().(LScheduler_LScheduler)
return ok && data1.Replica.Equals(data2.Replica) && data1.NextActionIndex.Cmp(data2.NextActionIndex) == 0
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this LScheduler) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(LScheduler)
return ok && _this.Equals(typed)
}
func Type_LScheduler_() _dafny.Type {
  return type_LScheduler_{}
}

type type_LScheduler_ struct {
}

func (_this type_LScheduler_) Default() interface{} {
  return LScheduler{LScheduler_LScheduler{Type_LReplica_().Default().(_126_LiveRSL____Replica__i_Compile.LReplica), _dafny.Zero}}
}

func (_this type_LScheduler_) String() string {
  return "_126_LiveRSL____Replica__i_Compile.LScheduler"
}
// End of data type LScheduler

