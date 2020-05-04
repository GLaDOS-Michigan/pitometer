// Package _120_LiveRSL____Executor__i_Compile
// Dafny module _120_LiveRSL____Executor__i_Compile compiled into Go

package _120_LiveRSL____Executor__i_Compile

import (
	_0_Native____NativeTypes__s_Compile "0_Native____NativeTypes__s_Compile_"
	_101_LiveRSL____Proposer__i_Compile "101_LiveRSL____Proposer__i_Compile_"
	_115_LiveRSL____StateMachine__i_Compile "115_LiveRSL____StateMachine__i_Compile_"
	_118_Collections____Maps__i_Compile "118_Collections____Maps__i_Compile_"
	_26_Collections____Seqs__s_Compile "26_Collections____Seqs__s_Compile_"
	_2_Collections____Maps2__s_Compile "2_Collections____Maps2__s_Compile_"
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
	_5_Temporal____Temporal__s_Compile "5_Temporal____Temporal__s_Compile_"
	_61_LiveRSL____Configuration__i_Compile "61_LiveRSL____Configuration__i_Compile_"
	_65_LiveRSL____Message__i_Compile "65_LiveRSL____Message__i_Compile_"
	_68_LiveRSL____Environment__i_Compile "68_LiveRSL____Environment__i_Compile_"
	_71_LiveRSL____ClockReading__i_Compile "71_LiveRSL____ClockReading__i_Compile_"
	_74_Common____UpperBound__s_Compile "74_Common____UpperBound__s_Compile_"
	_76_LiveRSL____Parameters__i_Compile "76_LiveRSL____Parameters__i_Compile_"
	_78_LiveRSL____Constants__i_Compile "78_LiveRSL____Constants__i_Compile_"
	_7_Environment__s_Compile "7_Environment__s_Compile_"
	_85_LiveRSL____Broadcast__i_Compile "85_LiveRSL____Broadcast__i_Compile_"
	_91_Collections____CountMatches__i_Compile "91_Collections____CountMatches__i_Compile_"
	_93_LiveRSL____Acceptor__i_Compile "93_LiveRSL____Acceptor__i_Compile_"
	_99_LiveRSL____Election__i_Compile "99_LiveRSL____Election__i_Compile_"
	_9_Native____Io__s_Compile "9_Native____Io__s_Compile_"
	_System "System_"
	_dafny "dafny"
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

type Dummy__ struct{}

// Definition of data type OutstandingOperation
type OutstandingOperation struct {
	Data_OutstandingOperation_
}

func (_this OutstandingOperation) Get() Data_OutstandingOperation_ {
	return _this.Data_OutstandingOperation_
}

type Data_OutstandingOperation_ interface {
	isOutstandingOperation()
}

type CompanionStruct_OutstandingOperation_ struct{}

var Companion_OutstandingOperation_ = CompanionStruct_OutstandingOperation_{}

type OutstandingOperation_OutstandingOpKnown struct {
	V   _dafny.Seq
	Bal _56_LiveRSL____Types__i_Compile.Ballot
}

func (OutstandingOperation_OutstandingOpKnown) isOutstandingOperation() {}

func (CompanionStruct_OutstandingOperation_) Create_OutstandingOpKnown_(V _dafny.Seq, Bal _56_LiveRSL____Types__i_Compile.Ballot) OutstandingOperation {
	return OutstandingOperation{OutstandingOperation_OutstandingOpKnown{V, Bal}}
}

func (_this OutstandingOperation) Is_OutstandingOpKnown() bool {
	_, ok := _this.Get().(OutstandingOperation_OutstandingOpKnown)
	return ok
}

type OutstandingOperation_OutstandingOpUnknown struct {
}

func (OutstandingOperation_OutstandingOpUnknown) isOutstandingOperation() {}

func (CompanionStruct_OutstandingOperation_) Create_OutstandingOpUnknown_() OutstandingOperation {
	return OutstandingOperation{OutstandingOperation_OutstandingOpUnknown{}}
}

func (_this OutstandingOperation) Is_OutstandingOpUnknown() bool {
	_, ok := _this.Get().(OutstandingOperation_OutstandingOpUnknown)
	return ok
}

func (_this OutstandingOperation) Dtor_v() _dafny.Seq {
	return _this.Get().(OutstandingOperation_OutstandingOpKnown).V
}

func (_this OutstandingOperation) Dtor_bal() _56_LiveRSL____Types__i_Compile.Ballot {
	return _this.Get().(OutstandingOperation_OutstandingOpKnown).Bal
}

func (_this OutstandingOperation) String() string {
	switch data := _this.Get().(type) {
	case nil:
		return "null"
	case OutstandingOperation_OutstandingOpKnown:
		{
			return "_120_LiveRSL____Executor__i_Compile.OutstandingOperation.OutstandingOpKnown" + "(" + _dafny.String(data.V) + ", " + _dafny.String(data.Bal) + ")"
		}
	case OutstandingOperation_OutstandingOpUnknown:
		{
			return "_120_LiveRSL____Executor__i_Compile.OutstandingOperation.OutstandingOpUnknown"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this OutstandingOperation) Equals(other OutstandingOperation) bool {
	switch data1 := _this.Get().(type) {
	case OutstandingOperation_OutstandingOpKnown:
		{
			data2, ok := other.Get().(OutstandingOperation_OutstandingOpKnown)
			return ok && data1.V.Equals(data2.V) && data1.Bal.Equals(data2.Bal)
		}
	case OutstandingOperation_OutstandingOpUnknown:
		{
			_, ok := other.Get().(OutstandingOperation_OutstandingOpUnknown)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this OutstandingOperation) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(OutstandingOperation)
	return ok && _this.Equals(typed)
}
func Type_OutstandingOperation_() _dafny.Type {
	return type_OutstandingOperation_{}
}

type type_OutstandingOperation_ struct {
}

func (_this type_OutstandingOperation_) Default() interface{} {
	return OutstandingOperation{OutstandingOperation_OutstandingOpKnown{_dafny.EmptySeq, _56_LiveRSL____Types__i_Compile.Type_Ballot_().Default().(_56_LiveRSL____Types__i_Compile.Ballot)}}
}

func (_this type_OutstandingOperation_) String() string {
	return "_120_LiveRSL____Executor__i_Compile.OutstandingOperation"
}

// End of data type OutstandingOperation

// Definition of data type LExecutor
type LExecutor struct {
	Data_LExecutor_
}

func (_this LExecutor) Get() Data_LExecutor_ {
	return _this.Data_LExecutor_
}

type Data_LExecutor_ interface {
	isLExecutor()
}

type CompanionStruct_LExecutor_ struct{}

var Companion_LExecutor_ = CompanionStruct_LExecutor_{}

type LExecutor_LExecutor struct {
	Constants             _78_LiveRSL____Constants__i_Compile.LReplicaConstants
	App                   uint64
	Ops__complete         _dafny.Int
	Max__bal__reflected   _56_LiveRSL____Types__i_Compile.Ballot
	Next__op__to__execute OutstandingOperation
	Reply__cache          _dafny.Map
}

func (LExecutor_LExecutor) isLExecutor() {}

func (CompanionStruct_LExecutor_) Create_LExecutor_(Constants _78_LiveRSL____Constants__i_Compile.LReplicaConstants, App uint64, Ops__complete _dafny.Int, Max__bal__reflected _56_LiveRSL____Types__i_Compile.Ballot, Next__op__to__execute OutstandingOperation, Reply__cache _dafny.Map) LExecutor {
	return LExecutor{LExecutor_LExecutor{Constants, App, Ops__complete, Max__bal__reflected, Next__op__to__execute, Reply__cache}}
}

func (_this LExecutor) Is_LExecutor() bool {
	_, ok := _this.Get().(LExecutor_LExecutor)
	return ok
}

func (_this LExecutor) Dtor_constants() _78_LiveRSL____Constants__i_Compile.LReplicaConstants {
	return _this.Get().(LExecutor_LExecutor).Constants
}

func (_this LExecutor) Dtor_app() uint64 {
	return _this.Get().(LExecutor_LExecutor).App
}

func (_this LExecutor) Dtor_ops__complete() _dafny.Int {
	return _this.Get().(LExecutor_LExecutor).Ops__complete
}

func (_this LExecutor) Dtor_max__bal__reflected() _56_LiveRSL____Types__i_Compile.Ballot {
	return _this.Get().(LExecutor_LExecutor).Max__bal__reflected
}

func (_this LExecutor) Dtor_next__op__to__execute() OutstandingOperation {
	return _this.Get().(LExecutor_LExecutor).Next__op__to__execute
}

func (_this LExecutor) Dtor_reply__cache() _dafny.Map {
	return _this.Get().(LExecutor_LExecutor).Reply__cache
}

func (_this LExecutor) String() string {
	switch data := _this.Get().(type) {
	case nil:
		return "null"
	case LExecutor_LExecutor:
		{
			return "_120_LiveRSL____Executor__i_Compile.LExecutor.LExecutor" + "(" + _dafny.String(data.Constants) + ", " + _dafny.String(data.App) + ", " + _dafny.String(data.Ops__complete) + ", " + _dafny.String(data.Max__bal__reflected) + ", " + _dafny.String(data.Next__op__to__execute) + ", " + _dafny.String(data.Reply__cache) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this LExecutor) Equals(other LExecutor) bool {
	switch data1 := _this.Get().(type) {
	case LExecutor_LExecutor:
		{
			data2, ok := other.Get().(LExecutor_LExecutor)
			return ok && data1.Constants.Equals(data2.Constants) && data1.App == data2.App && data1.Ops__complete.Cmp(data2.Ops__complete) == 0 && data1.Max__bal__reflected.Equals(data2.Max__bal__reflected) && data1.Next__op__to__execute.Equals(data2.Next__op__to__execute) && data1.Reply__cache.Equals(data2.Reply__cache)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this LExecutor) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(LExecutor)
	return ok && _this.Equals(typed)
}
func Type_LExecutor_() _dafny.Type {
	return type_LExecutor_{}
}

type type_LExecutor_ struct {
}

func (_this type_LExecutor_) Default() interface{} {
	return LExecutor{LExecutor_LExecutor{_78_LiveRSL____Constants__i_Compile.Type_LReplicaConstants_().Default().(_78_LiveRSL____Constants__i_Compile.LReplicaConstants), 0, _dafny.Zero, _56_LiveRSL____Types__i_Compile.Type_Ballot_().Default().(_56_LiveRSL____Types__i_Compile.Ballot), Type_OutstandingOperation_().Default().(OutstandingOperation), _dafny.EmptyMap}}
}

func (_this type_LExecutor_) String() string {
	return "_120_LiveRSL____Executor__i_Compile.LExecutor"
}

// End of data type LExecutor
