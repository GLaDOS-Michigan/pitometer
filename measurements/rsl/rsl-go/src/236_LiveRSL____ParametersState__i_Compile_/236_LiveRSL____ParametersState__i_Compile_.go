// Package _236_LiveRSL____ParametersState__i_Compile
// Dafny module _236_LiveRSL____ParametersState__i_Compile compiled into Go

package _236_LiveRSL____ParametersState__i_Compile

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

// BatchSize is the size of each paxos batch
var BatchSize int64 = 1 // make 1 the default

type Dummy__ struct{}

// Definition of data type ParametersState
type ParametersState struct {
	Data_ParametersState_
}

func (_this ParametersState) Get() Data_ParametersState_ {
	return _this.Data_ParametersState_
}

type Data_ParametersState_ interface {
	isParametersState()
}

type CompanionStruct_ParametersState_ struct{}

var Companion_ParametersState_ = CompanionStruct_ParametersState_{}

type ParametersState_ParametersState struct {
	Max__log__length                uint64
	Baseline__view__timeout__period uint64
	Heartbeat__period               uint64
	Max__integer__val               uint64
	Max__batch__size                uint64
	Max__batch__delay               uint64
}

func (ParametersState_ParametersState) isParametersState() {}

func (CompanionStruct_ParametersState_) Create_ParametersState_(Max__log__length uint64, Baseline__view__timeout__period uint64, Heartbeat__period uint64, Max__integer__val uint64, Max__batch__size uint64, Max__batch__delay uint64) ParametersState {
	return ParametersState{ParametersState_ParametersState{Max__log__length, Baseline__view__timeout__period, Heartbeat__period, Max__integer__val, Max__batch__size, Max__batch__delay}}
}

func (_this ParametersState) Is_ParametersState() bool {
	_, ok := _this.Get().(ParametersState_ParametersState)
	return ok
}

func (_this ParametersState) Dtor_max__log__length() uint64 {
	return _this.Get().(ParametersState_ParametersState).Max__log__length
}

func (_this ParametersState) Dtor_baseline__view__timeout__period() uint64 {
	return _this.Get().(ParametersState_ParametersState).Baseline__view__timeout__period
}

func (_this ParametersState) Dtor_heartbeat__period() uint64 {
	return _this.Get().(ParametersState_ParametersState).Heartbeat__period
}

func (_this ParametersState) Dtor_max__integer__val() uint64 {
	return _this.Get().(ParametersState_ParametersState).Max__integer__val
}

func (_this ParametersState) Dtor_max__batch__size() uint64 {
	return _this.Get().(ParametersState_ParametersState).Max__batch__size
}

func (_this ParametersState) Dtor_max__batch__delay() uint64 {
	return _this.Get().(ParametersState_ParametersState).Max__batch__delay
}

func (_this ParametersState) String() string {
	switch data := _this.Get().(type) {
	case nil:
		return "null"
	case ParametersState_ParametersState:
		{
			return "_236_LiveRSL____ParametersState__i_Compile.ParametersState.ParametersState" + "(" + _dafny.String(data.Max__log__length) + ", " + _dafny.String(data.Baseline__view__timeout__period) + ", " + _dafny.String(data.Heartbeat__period) + ", " + _dafny.String(data.Max__integer__val) + ", " + _dafny.String(data.Max__batch__size) + ", " + _dafny.String(data.Max__batch__delay) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ParametersState) Equals(other ParametersState) bool {
	switch data1 := _this.Get().(type) {
	case ParametersState_ParametersState:
		{
			data2, ok := other.Get().(ParametersState_ParametersState)
			return ok && data1.Max__log__length == data2.Max__log__length && data1.Baseline__view__timeout__period == data2.Baseline__view__timeout__period && data1.Heartbeat__period == data2.Heartbeat__period && data1.Max__integer__val == data2.Max__integer__val && data1.Max__batch__size == data2.Max__batch__size && data1.Max__batch__delay == data2.Max__batch__delay
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ParametersState) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ParametersState)
	return ok && _this.Equals(typed)
}
func Type_ParametersState_() _dafny.Type {
	return type_ParametersState_{}
}

type type_ParametersState_ struct {
}

func (_this type_ParametersState_) Default() interface{} {
	return ParametersState{ParametersState_ParametersState{0, 0, 0, 0, 0, 0}}
}

func (_this type_ParametersState_) String() string {
	return "_236_LiveRSL____ParametersState__i_Compile.ParametersState"
}

// End of data type ParametersState

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

var Companion_Default___ = CompanionStruct_Default___{}

func (_this *Default__) Equals(other *Default__) bool {
	return _this == other
}

func (_this *Default__) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*Default__)
	return ok && _this.Equals(other)
}

func (*Default__) String() string {
	return "_236_LiveRSL____ParametersState__i_Compile.Default__"
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
	return "_236_LiveRSL____ParametersState__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) StaticParams() ParametersState {
	// The values for periods and delays need to be tuned
	// return ParametersState{ParametersState_ParametersState{ // These are units of milliseconds, same as original C# implementation
	// 	Max__log__length:                uint64(7),
	// 	Baseline__view__timeout__period: uint64(1000),
	// 	Heartbeat__period:               uint64(100),
	// 	Max__integer__val:               (uint64(9223372036854775808)) - (func() uint64 { return (uint64(1)) })(),
	// 	Max__batch__size:                uint64(32),
	// 	Max__batch__delay:               uint64(10)}}
	return ParametersState{ParametersState_ParametersState{
		// These are units of milliseconds, same as original C# implementation
		// Tony's tune: make log length really long to avoid it happening during an execution
		// make heartbeats infrequent such that they don't interfere with our message queues
		Max__log__length:                uint64(10_000),
		Baseline__view__timeout__period: uint64(1000),
		Heartbeat__period:               uint64(30_000),
		Max__integer__val:               (uint64(9223372036854775808)) - (func() uint64 { return (uint64(1)) })(),
		Max__batch__size:                uint64(BatchSize),
		Max__batch__delay:               uint64(10)}}
}

// End of class Default__
