// Package _170_Common____Util__i_Compile
// Dafny module _170_Common____Util__i_Compile compiled into Go

package _170_Common____Util__i_Compile

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

type Dummy__ struct{}

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
	return "_170_Common____Util__i_Compile.Default__"
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
	return "_170_Common____Util__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) SeqToArray__slow(Type_A_ _dafny.Type, s _dafny.Seq) *_dafny.Array {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var a *_dafny.Array = _dafny.NewArrayWithValue(Type_A_.Default(), _dafny.IntOf(0))
	var _ = a
	var _4531_len _dafny.Int
	var _ = _4531_len
	_4531_len = (s).Cardinality()
	var _nw0 = _dafny.NewArrayWithValue(Type_A_.Default(), _4531_len)
	var _ = _nw0
	a = _nw0
	var _4532_i _dafny.Int
	var _ = _4532_i
	_4532_i = _dafny.Zero
	for (_4532_i).Cmp(_4531_len) < 0 {
		*((a).Index(_dafny.IntOfAny((_4532_i)))) = (s).Index(_4532_i).(interface{})
		_4532_i = (_4532_i).Plus(_dafny.IntOfInt64(1))
	}
	return a
}
func (_this *CompanionStruct_Default___) SeqIntoArrayOpt(s _dafny.Seq, a *_dafny.Array) {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var _4533_i uint64
	var _ = _4533_i
	_4533_i = uint64(0)
	for (_4533_i) < (uint64((s).CardinalityUint64())) {
		*((a).Index(_dafny.IntOfAny((_4533_i)))) = (s).IndexUint(_4533_i).(interface{})
		_4533_i = (_4533_i) + (uint64(1))
	}
}
func (_this *CompanionStruct_Default___) SeqToArrayOpt(Type_A_ _dafny.Type, s _dafny.Seq) *_dafny.Array {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var a *_dafny.Array = _dafny.NewArrayWithValue(Type_A_.Default(), _dafny.IntOf(0))
	var _ = a
	var _nw1 = _dafny.NewArrayWithValue(Type_A_.Default(), _dafny.IntOfAny(uint64((s).CardinalityUint64())))
	var _ = _nw1
	a = _nw1
	Companion_Default___.SeqIntoArrayOpt(s, a)
	return a
}
func (_this *CompanionStruct_Default___) SeqIntoArrayChar(s _dafny.Seq, a *_dafny.Array) {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var _4534_i uint64
	var _ = _4534_i
	_4534_i = uint64(0)
	for (_4534_i) < (uint64((s).CardinalityUint64())) {
		*((a).Index(_dafny.IntOfAny((_4534_i)))) = (s).IndexUint(_4534_i).(_dafny.Char)
		_4534_i = (_4534_i) + (uint64(1))
	}
}
func (_this *CompanionStruct_Default___) RecordTimingSeq(name _dafny.Seq, start uint64, end uint64) {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var _4535_name__array *_dafny.Array
	var _ = _4535_name__array
	var _nw2 = _dafny.NewArrayWithValue(_dafny.Char('D'), (name).Cardinality())
	var _ = _nw2
	_4535_name__array = _nw2
	Companion_Default___.SeqIntoArrayChar(name, _4535_name__array)
	var _4536_time uint64 = 0
	var _ = _4536_time
	if (start) <= (end) {
		_4536_time = (end) - (func() uint64 { return (start) })()
	} else {
		_4536_time = uint64(18446744073709551615)
	}
	// TONY: Comment this out
	// _9_Native____Io__s_Compile.Companion_Time_.RecordTiming(_4535_name__array, _4536_time)
}
func (_this *CompanionStruct_Default___) SeqByteToUint64(bs _dafny.Seq) uint64 {
	return (((((((((((uint64((bs).IndexUint(uint64(0)).(uint8))) * (uint64(256))) * (uint64(256))) * (uint64(256))) * (uint64(4294967296))) + ((((uint64((bs).IndexUint(uint64(1)).(uint8))) * (uint64(256))) * (uint64(256))) * (uint64(4294967296)))) + (((uint64((bs).IndexUint(uint64(2)).(uint8))) * (uint64(256))) * (uint64(4294967296)))) + ((uint64((bs).IndexUint(uint64(3)).(uint8))) * (uint64(4294967296)))) + ((((uint64((bs).IndexUint(uint64(4)).(uint8))) * (uint64(256))) * (uint64(256))) * (uint64(256)))) + (((uint64((bs).IndexUint(uint64(5)).(uint8))) * (uint64(256))) * (uint64(256)))) + ((uint64((bs).IndexUint(uint64(6)).(uint8))) * (uint64(256)))) + (uint64((bs).IndexUint(uint64(7)).(uint8)))
}
func (_this *CompanionStruct_Default___) Uint64ToSeqByte(u uint64) _dafny.Seq {
	var _4537_bs _dafny.Seq = _dafny.SeqOf(uint8((u)/(uint64(72057594037927936))), uint8(((u)/(uint64(281474976710656)))%(uint64(256))), uint8(((u)/(uint64(1099511627776)))%(uint64(256))), uint8(((u)/(uint64(4294967296)))%(uint64(256))), uint8(((u)/(uint64(16777216)))%(uint64(256))), uint8(((u)/(uint64(65536)))%(uint64(256))), uint8(((u)/(uint64(256)))%(uint64(256))), uint8((u)%(uint64(256))))
	var _ = _4537_bs
	var _4538_u__int _dafny.Int = _dafny.IntOfUint64(u)
	var _ = _4538_u__int
	return _4537_bs
}
func (_this *CompanionStruct_Default___) SeqByteToUint16(bs _dafny.Seq) uint16 {
	return ((uint16((bs).IndexUint(uint64(0)).(uint8))) * (uint16(256))) + (uint16((bs).IndexUint(uint64(1)).(uint8)))
}
func (_this *CompanionStruct_Default___) Uint16ToSeqByte(u uint16) _dafny.Seq {
	var _4539_s _dafny.Seq = _dafny.SeqOf(uint8(((u)/(uint16(256)))%(uint16(256))), uint8((u)%(uint16(256))))
	var _ = _4539_s
	var _4540_u__int _dafny.Int = _dafny.IntOfUint16(u)
	var _ = _4540_u__int
	return _4539_s
}

// End of class Default__
