// Package _101_LiveRSL____Proposer__i_Compile
// Dafny module _101_LiveRSL____Proposer__i_Compile compiled into Go

package _101_LiveRSL____Proposer__i_Compile

import (
	_0_Native____NativeTypes__s_Compile "0_Native____NativeTypes__s_Compile_"
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

type Dummy__ struct{}

// Definition of data type IncompleteBatchTimer
type IncompleteBatchTimer struct {
	Data_IncompleteBatchTimer_
}

func (_this IncompleteBatchTimer) Get() Data_IncompleteBatchTimer_ {
	return _this.Data_IncompleteBatchTimer_
}

type Data_IncompleteBatchTimer_ interface {
	isIncompleteBatchTimer()
}

type CompanionStruct_IncompleteBatchTimer_ struct{}

var Companion_IncompleteBatchTimer_ = CompanionStruct_IncompleteBatchTimer_{}

type IncompleteBatchTimer_IncompleteBatchTimerOn struct {
	When _dafny.Int
}

func (IncompleteBatchTimer_IncompleteBatchTimerOn) isIncompleteBatchTimer() {}

func (CompanionStruct_IncompleteBatchTimer_) Create_IncompleteBatchTimerOn_(When _dafny.Int) IncompleteBatchTimer {
	return IncompleteBatchTimer{IncompleteBatchTimer_IncompleteBatchTimerOn{When}}
}

func (_this IncompleteBatchTimer) Is_IncompleteBatchTimerOn() bool {
	_, ok := _this.Get().(IncompleteBatchTimer_IncompleteBatchTimerOn)
	return ok
}

type IncompleteBatchTimer_IncompleteBatchTimerOff struct {
}

func (IncompleteBatchTimer_IncompleteBatchTimerOff) isIncompleteBatchTimer() {}

func (CompanionStruct_IncompleteBatchTimer_) Create_IncompleteBatchTimerOff_() IncompleteBatchTimer {
	return IncompleteBatchTimer{IncompleteBatchTimer_IncompleteBatchTimerOff{}}
}

func (_this IncompleteBatchTimer) Is_IncompleteBatchTimerOff() bool {
	_, ok := _this.Get().(IncompleteBatchTimer_IncompleteBatchTimerOff)
	return ok
}

func (_this IncompleteBatchTimer) Dtor_when() _dafny.Int {
	return _this.Get().(IncompleteBatchTimer_IncompleteBatchTimerOn).When
}

func (_this IncompleteBatchTimer) String() string {
	switch data := _this.Get().(type) {
	case nil:
		return "null"
	case IncompleteBatchTimer_IncompleteBatchTimerOn:
		{
			return "_101_LiveRSL____Proposer__i_Compile.IncompleteBatchTimer.IncompleteBatchTimerOn" + "(" + _dafny.String(data.When) + ")"
		}
	case IncompleteBatchTimer_IncompleteBatchTimerOff:
		{
			return "_101_LiveRSL____Proposer__i_Compile.IncompleteBatchTimer.IncompleteBatchTimerOff"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this IncompleteBatchTimer) Equals(other IncompleteBatchTimer) bool {
	switch data1 := _this.Get().(type) {
	case IncompleteBatchTimer_IncompleteBatchTimerOn:
		{
			data2, ok := other.Get().(IncompleteBatchTimer_IncompleteBatchTimerOn)
			return ok && data1.When.Cmp(data2.When) == 0
		}
	case IncompleteBatchTimer_IncompleteBatchTimerOff:
		{
			_, ok := other.Get().(IncompleteBatchTimer_IncompleteBatchTimerOff)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this IncompleteBatchTimer) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(IncompleteBatchTimer)
	return ok && _this.Equals(typed)
}
func Type_IncompleteBatchTimer_() _dafny.Type {
	return type_IncompleteBatchTimer_{}
}

type type_IncompleteBatchTimer_ struct {
}

func (_this type_IncompleteBatchTimer_) Default() interface{} {
	return IncompleteBatchTimer{IncompleteBatchTimer_IncompleteBatchTimerOn{_dafny.Zero}}
}

func (_this type_IncompleteBatchTimer_) String() string {
	return "_101_LiveRSL____Proposer__i_Compile.IncompleteBatchTimer"
}

// End of data type IncompleteBatchTimer

// Definition of data type LProposer
type LProposer struct {
	Data_LProposer_
}

func (_this LProposer) Get() Data_LProposer_ {
	return _this.Data_LProposer_
}

type Data_LProposer_ interface {
	isLProposer()
}

type CompanionStruct_LProposer_ struct{}

var Companion_LProposer_ = CompanionStruct_LProposer_{}

type LProposer_LProposer struct {
	Constants                                         _78_LiveRSL____Constants__i_Compile.LReplicaConstants
	Current__state                                    _dafny.Int
	Request__queue                                    _dafny.Seq
	Max__ballot__i__sent__1a                          _56_LiveRSL____Types__i_Compile.Ballot
	Next__operation__number__to__propose              _dafny.Int
	Received__1b__packets                             _dafny.Set
	Highest__seqno__requested__by__client__this__view _dafny.Map
	Incomplete__batch__timer                          IncompleteBatchTimer
	Election__state                                   _99_LiveRSL____Election__i_Compile.ElectionState
}

func (LProposer_LProposer) isLProposer() {}

func (CompanionStruct_LProposer_) Create_LProposer_(Constants _78_LiveRSL____Constants__i_Compile.LReplicaConstants, Current__state _dafny.Int, Request__queue _dafny.Seq, Max__ballot__i__sent__1a _56_LiveRSL____Types__i_Compile.Ballot, Next__operation__number__to__propose _dafny.Int, Received__1b__packets _dafny.Set, Highest__seqno__requested__by__client__this__view _dafny.Map, Incomplete__batch__timer IncompleteBatchTimer, Election__state _99_LiveRSL____Election__i_Compile.ElectionState) LProposer {
	return LProposer{LProposer_LProposer{Constants, Current__state, Request__queue, Max__ballot__i__sent__1a, Next__operation__number__to__propose, Received__1b__packets, Highest__seqno__requested__by__client__this__view, Incomplete__batch__timer, Election__state}}
}

func (_this LProposer) Is_LProposer() bool {
	_, ok := _this.Get().(LProposer_LProposer)
	return ok
}

func (_this LProposer) Dtor_constants() _78_LiveRSL____Constants__i_Compile.LReplicaConstants {
	return _this.Get().(LProposer_LProposer).Constants
}

func (_this LProposer) Dtor_current__state() _dafny.Int {
	return _this.Get().(LProposer_LProposer).Current__state
}

func (_this LProposer) Dtor_request__queue() _dafny.Seq {
	return _this.Get().(LProposer_LProposer).Request__queue
}

func (_this LProposer) Dtor_max__ballot__i__sent__1a() _56_LiveRSL____Types__i_Compile.Ballot {
	return _this.Get().(LProposer_LProposer).Max__ballot__i__sent__1a
}

func (_this LProposer) Dtor_next__operation__number__to__propose() _dafny.Int {
	return _this.Get().(LProposer_LProposer).Next__operation__number__to__propose
}

func (_this LProposer) Dtor_received__1b__packets() _dafny.Set {
	return _this.Get().(LProposer_LProposer).Received__1b__packets
}

func (_this LProposer) Dtor_highest__seqno__requested__by__client__this__view() _dafny.Map {
	return _this.Get().(LProposer_LProposer).Highest__seqno__requested__by__client__this__view
}

func (_this LProposer) Dtor_incomplete__batch__timer() IncompleteBatchTimer {
	return _this.Get().(LProposer_LProposer).Incomplete__batch__timer
}

func (_this LProposer) Dtor_election__state() _99_LiveRSL____Election__i_Compile.ElectionState {
	return _this.Get().(LProposer_LProposer).Election__state
}

func (_this LProposer) String() string {
	switch data := _this.Get().(type) {
	case nil:
		return "null"
	case LProposer_LProposer:
		{
			return "_101_LiveRSL____Proposer__i_Compile.LProposer.LProposer" + "(" + _dafny.String(data.Constants) + ", " + _dafny.String(data.Current__state) + ", " + _dafny.String(data.Request__queue) + ", " + _dafny.String(data.Max__ballot__i__sent__1a) + ", " + _dafny.String(data.Next__operation__number__to__propose) + ", " + _dafny.String(data.Received__1b__packets) + ", " + _dafny.String(data.Highest__seqno__requested__by__client__this__view) + ", " + _dafny.String(data.Incomplete__batch__timer) + ", " + _dafny.String(data.Election__state) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this LProposer) Equals(other LProposer) bool {
	switch data1 := _this.Get().(type) {
	case LProposer_LProposer:
		{
			data2, ok := other.Get().(LProposer_LProposer)
			return ok && data1.Constants.Equals(data2.Constants) && data1.Current__state.Cmp(data2.Current__state) == 0 && data1.Request__queue.Equals(data2.Request__queue) && data1.Max__ballot__i__sent__1a.Equals(data2.Max__ballot__i__sent__1a) && data1.Next__operation__number__to__propose.Cmp(data2.Next__operation__number__to__propose) == 0 && data1.Received__1b__packets.Equals(data2.Received__1b__packets) && data1.Highest__seqno__requested__by__client__this__view.Equals(data2.Highest__seqno__requested__by__client__this__view) && data1.Incomplete__batch__timer.Equals(data2.Incomplete__batch__timer) && data1.Election__state.Equals(data2.Election__state)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this LProposer) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(LProposer)
	return ok && _this.Equals(typed)
}
func Type_LProposer_() _dafny.Type {
	return type_LProposer_{}
}

type type_LProposer_ struct {
}

func (_this type_LProposer_) Default() interface{} {
	return LProposer{LProposer_LProposer{_78_LiveRSL____Constants__i_Compile.Type_LReplicaConstants_().Default().(_78_LiveRSL____Constants__i_Compile.LReplicaConstants), _dafny.Zero, _dafny.EmptySeq, _56_LiveRSL____Types__i_Compile.Type_Ballot_().Default().(_56_LiveRSL____Types__i_Compile.Ballot), _dafny.Zero, _dafny.EmptySet, _dafny.EmptyMap, Type_IncompleteBatchTimer_().Default().(IncompleteBatchTimer), _99_LiveRSL____Election__i_Compile.Type_ElectionState_().Default().(_99_LiveRSL____Election__i_Compile.ElectionState)}}
}

func (_this type_LProposer_) String() string {
	return "_101_LiveRSL____Proposer__i_Compile.LProposer"
}

// End of data type LProposer
