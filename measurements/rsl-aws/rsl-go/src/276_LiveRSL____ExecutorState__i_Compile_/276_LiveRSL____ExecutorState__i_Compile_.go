// Package _276_LiveRSL____ExecutorState__i_Compile
// Dafny module _276_LiveRSL____ExecutorState__i_Compile compiled into Go

package _276_LiveRSL____ExecutorState__i_Compile

import (
	_0_Native____NativeTypes__s_Compile "0_Native____NativeTypes__s_Compile_"
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
	_254_LiveRSL____ProposerState__i_Compile "254_LiveRSL____ProposerState__i_Compile_"
	_265_LiveRSL____COperationNumberSort__i_Compile "265_LiveRSL____COperationNumberSort__i_Compile_"
	_267_LiveRSL____CLastCheckpointedMap__i_Compile "267_LiveRSL____CLastCheckpointedMap__i_Compile_"
	_269_LiveRSL____AcceptorState__i_Compile "269_LiveRSL____AcceptorState__i_Compile_"
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
var _ _254_LiveRSL____ProposerState__i_Compile.Dummy__
var _ _265_LiveRSL____COperationNumberSort__i_Compile.Dummy__
var _ _267_LiveRSL____CLastCheckpointedMap__i_Compile.Dummy__
var _ _269_LiveRSL____AcceptorState__i_Compile.Dummy__

type Dummy__ struct{}

// Definition of data type COutstandingOperation
type COutstandingOperation struct {
	Data_COutstandingOperation_
}

func (_this COutstandingOperation) Get() Data_COutstandingOperation_ {
	return _this.Data_COutstandingOperation_
}

type Data_COutstandingOperation_ interface {
	isCOutstandingOperation()
}

type CompanionStruct_COutstandingOperation_ struct{}

var Companion_COutstandingOperation_ = CompanionStruct_COutstandingOperation_{}

type COutstandingOperation_COutstandingOpKnown struct {
	V   _dafny.Seq
	Bal _214_LiveRSL____CTypes__i_Compile.CBallot
}

func (COutstandingOperation_COutstandingOpKnown) isCOutstandingOperation() {}

func (CompanionStruct_COutstandingOperation_) Create_COutstandingOpKnown_(V _dafny.Seq, Bal _214_LiveRSL____CTypes__i_Compile.CBallot) COutstandingOperation {
	return COutstandingOperation{COutstandingOperation_COutstandingOpKnown{V, Bal}}
}

func (_this COutstandingOperation) Is_COutstandingOpKnown() bool {
	_, ok := _this.Get().(COutstandingOperation_COutstandingOpKnown)
	return ok
}

type COutstandingOperation_COutstandingOpUnknown struct {
}

func (COutstandingOperation_COutstandingOpUnknown) isCOutstandingOperation() {}

func (CompanionStruct_COutstandingOperation_) Create_COutstandingOpUnknown_() COutstandingOperation {
	return COutstandingOperation{COutstandingOperation_COutstandingOpUnknown{}}
}

func (_this COutstandingOperation) Is_COutstandingOpUnknown() bool {
	_, ok := _this.Get().(COutstandingOperation_COutstandingOpUnknown)
	return ok
}

func (_this COutstandingOperation) Dtor_v() _dafny.Seq {
	return _this.Get().(COutstandingOperation_COutstandingOpKnown).V
}

func (_this COutstandingOperation) Dtor_bal() _214_LiveRSL____CTypes__i_Compile.CBallot {
	return _this.Get().(COutstandingOperation_COutstandingOpKnown).Bal
}

func (_this COutstandingOperation) String() string {
	switch data := _this.Get().(type) {
	case nil:
		return "null"
	case COutstandingOperation_COutstandingOpKnown:
		{
			return "COutstandingOperation.COutstandingOpKnown" + "(" + _dafny.String(data.V) + ", " + _dafny.String(data.Bal) + ")"
		}
	case COutstandingOperation_COutstandingOpUnknown:
		{
			return "COutstandingOperation.COutstandingOpUnknown"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this COutstandingOperation) Equals(other COutstandingOperation) bool {
	switch data1 := _this.Get().(type) {
	case COutstandingOperation_COutstandingOpKnown:
		{
			data2, ok := other.Get().(COutstandingOperation_COutstandingOpKnown)
			return ok && data1.V.Equals(data2.V) && data1.Bal.Equals(data2.Bal)
		}
	case COutstandingOperation_COutstandingOpUnknown:
		{
			_, ok := other.Get().(COutstandingOperation_COutstandingOpUnknown)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this COutstandingOperation) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(COutstandingOperation)
	return ok && _this.Equals(typed)
}
func Type_COutstandingOperation_() _dafny.Type {
	return type_COutstandingOperation_{}
}

type type_COutstandingOperation_ struct {
}

func (_this type_COutstandingOperation_) Default() interface{} {
	return COutstandingOperation{COutstandingOperation_COutstandingOpKnown{_dafny.EmptySeq, _214_LiveRSL____CTypes__i_Compile.Type_CBallot_().Default().(_214_LiveRSL____CTypes__i_Compile.CBallot)}}
}

func (_this type_COutstandingOperation_) String() string {
	return "COutstandingOperation"
}

// End of data type COutstandingOperation

// Definition of data type ExecutorState
type ExecutorState struct {
	Data_ExecutorState_
}

func (_this ExecutorState) Get() Data_ExecutorState_ {
	return _this.Data_ExecutorState_
}

type Data_ExecutorState_ interface {
	isExecutorState()
}

type CompanionStruct_ExecutorState_ struct{}

var Companion_ExecutorState_ = CompanionStruct_ExecutorState_{}

type ExecutorState_ExecutorState struct {
	Constants             _245_LiveRSL____ReplicaConstantsState__i_Compile.ReplicaConstantsState
	App                   uint64
	Ops__complete         _214_LiveRSL____CTypes__i_Compile.COperationNumber
	Max__bal__reflected   _214_LiveRSL____CTypes__i_Compile.CBallot
	Next__op__to__execute COutstandingOperation
}

func (ExecutorState_ExecutorState) isExecutorState() {}

func (CompanionStruct_ExecutorState_) Create_ExecutorState_(Constants _245_LiveRSL____ReplicaConstantsState__i_Compile.ReplicaConstantsState, App uint64, Ops__complete _214_LiveRSL____CTypes__i_Compile.COperationNumber, Max__bal__reflected _214_LiveRSL____CTypes__i_Compile.CBallot, Next__op__to__execute COutstandingOperation) ExecutorState {
	return ExecutorState{ExecutorState_ExecutorState{Constants, App, Ops__complete, Max__bal__reflected, Next__op__to__execute}}
}

func (_this ExecutorState) Is_ExecutorState() bool {
	_, ok := _this.Get().(ExecutorState_ExecutorState)
	return ok
}

func (_this ExecutorState) Dtor_constants() _245_LiveRSL____ReplicaConstantsState__i_Compile.ReplicaConstantsState {
	return _this.Get().(ExecutorState_ExecutorState).Constants
}

func (_this ExecutorState) Dtor_app() uint64 {
	return _this.Get().(ExecutorState_ExecutorState).App
}

func (_this ExecutorState) Dtor_ops__complete() _214_LiveRSL____CTypes__i_Compile.COperationNumber {
	return _this.Get().(ExecutorState_ExecutorState).Ops__complete
}

func (_this ExecutorState) Dtor_max__bal__reflected() _214_LiveRSL____CTypes__i_Compile.CBallot {
	return _this.Get().(ExecutorState_ExecutorState).Max__bal__reflected
}

func (_this ExecutorState) Dtor_next__op__to__execute() COutstandingOperation {
	return _this.Get().(ExecutorState_ExecutorState).Next__op__to__execute
}

func (_this ExecutorState) String() string {
	switch data := _this.Get().(type) {
	case nil:
		return "null"
	case ExecutorState_ExecutorState:
		{
			return "ExecutorState.ExecutorState" + "(" + _dafny.String(data.Constants) + ", " + _dafny.String(data.App) + ", " + _dafny.String(data.Ops__complete) + ", " + _dafny.String(data.Max__bal__reflected) + ", " + _dafny.String(data.Next__op__to__execute) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ExecutorState) Equals(other ExecutorState) bool {
	switch data1 := _this.Get().(type) {
	case ExecutorState_ExecutorState:
		{
			data2, ok := other.Get().(ExecutorState_ExecutorState)
			return ok && data1.Constants.Equals(data2.Constants) && data1.App == data2.App && data1.Ops__complete.Equals(data2.Ops__complete) && data1.Max__bal__reflected.Equals(data2.Max__bal__reflected) && data1.Next__op__to__execute.Equals(data2.Next__op__to__execute)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ExecutorState) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ExecutorState)
	return ok && _this.Equals(typed)
}
func Type_ExecutorState_() _dafny.Type {
	return type_ExecutorState_{}
}

type type_ExecutorState_ struct {
}

func (_this type_ExecutorState_) Default() interface{} {
	return ExecutorState{ExecutorState_ExecutorState{_245_LiveRSL____ReplicaConstantsState__i_Compile.Type_ReplicaConstantsState_().Default().(_245_LiveRSL____ReplicaConstantsState__i_Compile.ReplicaConstantsState), 0, _214_LiveRSL____CTypes__i_Compile.Type_COperationNumber_().Default().(_214_LiveRSL____CTypes__i_Compile.COperationNumber), _214_LiveRSL____CTypes__i_Compile.Type_CBallot_().Default().(_214_LiveRSL____CTypes__i_Compile.CBallot), Type_COutstandingOperation_().Default().(COutstandingOperation)}}
}

func (_this type_ExecutorState_) String() string {
	return "ExecutorState"
}

// End of data type ExecutorState
