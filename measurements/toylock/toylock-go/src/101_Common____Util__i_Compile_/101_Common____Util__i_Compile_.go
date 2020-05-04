// Package _101_Common____Util__i_Compile
// Dafny module _101_Common____Util__i_Compile compiled into Go

package _101_Common____Util__i_Compile

import (
	_0_Native____NativeTypes__s_Compile "0_Native____NativeTypes__s_Compile_"
	_26_Collections____Seqs__s_Compile "26_Collections____Seqs__s_Compile_"
	_29_Collections____Sets__i_Compile "29_Collections____Sets__i_Compile_"
	_2_Collections____Maps2__s_Compile "2_Collections____Maps2__s_Compile_"
	_33_Types__i_Compile "33_Types__i_Compile_"
	_36_Protocol__Node__i_Compile "36_Protocol__Node__i_Compile_"
	_39_Message__i_Compile "39_Message__i_Compile_"
	_42_Common____UdpClient__i_Compile "42_Common____UdpClient__i_Compile_"
	_44_Logic____Option__i_Compile "44_Logic____Option__i_Compile_"
	_47_Collections____Maps__i_Compile "47_Collections____Maps__i_Compile_"
	_50_Collections____Seqs__i_Compile "50_Collections____Seqs__i_Compile_"
	_54_Native____NativeTypes__i_Compile "54_Native____NativeTypes__i_Compile_"
	_57_Libraries____base__s_Compile "57_Libraries____base__s_Compile_"
	_59_Math____power2__s_Compile "59_Math____power2__s_Compile_"
	_5_Temporal____Temporal__s_Compile "5_Temporal____Temporal__s_Compile_"
	_61_Math____power__s_Compile "61_Math____power__s_Compile_"
	_64_Math____mul__nonlinear__i_Compile "64_Math____mul__nonlinear__i_Compile_"
	_67_Math____mul__auto__proofs__i_Compile "67_Math____mul__auto__proofs__i_Compile_"
	_69_Math____mul__auto__i_Compile "69_Math____mul__auto__i_Compile_"
	_71_Math____mul__i_Compile "71_Math____mul__i_Compile_"
	_73_Math____power__i_Compile "73_Math____power__i_Compile_"
	_77_Math____div__def__i_Compile "77_Math____div__def__i_Compile_"
	_7_Environment__s_Compile "7_Environment__s_Compile_"
	_81_Math____div__boogie__i_Compile "81_Math____div__boogie__i_Compile_"
	_83_Math____div__nonlinear__i_Compile "83_Math____div__nonlinear__i_Compile_"
	_88_Math____mod__auto__proofs__i_Compile "88_Math____mod__auto__proofs__i_Compile_"
	_90_Math____mod__auto__i_Compile "90_Math____mod__auto__i_Compile_"
	_93_Math____div__auto__proofs__i_Compile "93_Math____div__auto__proofs__i_Compile_"
	_95_Math____div__auto__i_Compile "95_Math____div__auto__i_Compile_"
	_97_Math____div__i_Compile "97_Math____div__i_Compile_"
	_99_Math____power2__i_Compile "99_Math____power2__i_Compile_"
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
var _ _29_Collections____Sets__i_Compile.Dummy__
var _ _33_Types__i_Compile.Dummy__
var _ _36_Protocol__Node__i_Compile.Dummy__
var _ _39_Message__i_Compile.Dummy__
var _ _42_Common____UdpClient__i_Compile.Dummy__
var _ _44_Logic____Option__i_Compile.Dummy__
var _ _47_Collections____Maps__i_Compile.Dummy__
var _ _50_Collections____Seqs__i_Compile.Dummy__
var _ _54_Native____NativeTypes__i_Compile.Dummy__
var _ _57_Libraries____base__s_Compile.Dummy__
var _ _59_Math____power2__s_Compile.Dummy__
var _ _61_Math____power__s_Compile.Dummy__
var _ _64_Math____mul__nonlinear__i_Compile.Dummy__
var _ _67_Math____mul__auto__proofs__i_Compile.Dummy__
var _ _69_Math____mul__auto__i_Compile.Dummy__
var _ _71_Math____mul__i_Compile.Dummy__
var _ _73_Math____power__i_Compile.Dummy__
var _ _77_Math____div__def__i_Compile.Dummy__
var _ _81_Math____div__boogie__i_Compile.Dummy__
var _ _83_Math____div__nonlinear__i_Compile.Dummy__
var _ _88_Math____mod__auto__proofs__i_Compile.Dummy__
var _ _90_Math____mod__auto__i_Compile.Dummy__
var _ _93_Math____div__auto__proofs__i_Compile.Dummy__
var _ _95_Math____div__auto__i_Compile.Dummy__
var _ _97_Math____div__i_Compile.Dummy__
var _ _99_Math____power2__i_Compile.Dummy__

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
	return "_101_Common____Util__i_Compile.Default__"
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
	return "_101_Common____Util__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) SeqToArray__slow(Type_A_ _dafny.Type, s _dafny.Seq) *_dafny.Array {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var a *_dafny.Array = _dafny.NewArrayWithValue(Type_A_.Default(), _dafny.IntOf(0))
	var _ = a
	var _1568_len _dafny.Int
	var _ = _1568_len
	_1568_len = (s).Cardinality()
	var _nw0 = _dafny.NewArrayWithValue(Type_A_.Default(), _1568_len)
	var _ = _nw0
	a = _nw0
	var _1569_i _dafny.Int
	var _ = _1569_i
	_1569_i = _dafny.Zero
	for (_1569_i).Cmp(_1568_len) < 0 {
		*((a).Index(_dafny.IntOfAny((_1569_i)))) = (s).Index(_1569_i).(interface{})
		_1569_i = (_1569_i).Plus(_dafny.IntOfInt64(1))
	}
	return a
}
func (_this *CompanionStruct_Default___) SeqIntoArrayOpt(s _dafny.Seq, a *_dafny.Array) {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var _1570_i uint64
	var _ = _1570_i
	_1570_i = uint64(0)
	for (_1570_i) < ((s).CardinalityInt()).Uint64() {
		*((a).Index(_dafny.IntOfAny((_1570_i)))) = (s).Index(_dafny.IntOfUint64(_1570_i)).(interface{})
		_1570_i = (_1570_i) + (uint64(1))
	}
}
func (_this *CompanionStruct_Default___) SeqToArrayOpt(Type_A_ _dafny.Type, s _dafny.Seq) *_dafny.Array {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var a *_dafny.Array = _dafny.NewArrayWithValue(Type_A_.Default(), _dafny.IntOf(0))
	var _ = a
	var _nw1 = _dafny.NewArrayWithValue(Type_A_.Default(), (s).CardinalityInt())
	var _ = _nw1
	a = _nw1
	Companion_Default___.SeqIntoArrayOpt(s, a)
	return a
}
func (_this *CompanionStruct_Default___) SeqIntoArrayChar(s _dafny.Seq, a *_dafny.Array) {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var _1571_i uint64
	var _ = _1571_i
	_1571_i = uint64(0)
	for (_1571_i) < ((s).CardinalityInt()).Uint64() {
		*((a).Index(_dafny.IntOfAny((_1571_i)))) = (s).Index(_dafny.IntOfAny(_1571_i)).(_dafny.Char)
		_1571_i = (_1571_i) + (uint64(1))
	}
}
func (_this *CompanionStruct_Default___) RecordTimingSeq(name _dafny.Seq, start uint64, end uint64) {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var _1572_name__array *_dafny.Array
	var _ = _1572_name__array
	var _nw2 = _dafny.NewArrayWithValue(_dafny.Char('D'), (name).Cardinality())
	var _ = _nw2
	_1572_name__array = _nw2
	Companion_Default___.SeqIntoArrayChar(name, _1572_name__array)
	var _1573_time uint64 = 0
	var _ = _1573_time
	if (start) <= (end) {
		_1573_time = (end) - (func() uint64 { return (start) })()
	} else {
		_1573_time = uint64(18446744073709551615)
	}
	// TONY: REMOVED
	// _9_Native____Io__s_Compile.Companion_Time_.RecordTiming(_1572_name__array, _1573_time)
}
func (_this *CompanionStruct_Default___) SeqByteToUint64(bs _dafny.Seq) uint64 {
	return (((((((((((uint64((bs).Index(_dafny.IntOfAny(uint64(0))).(int))) * (uint64(256))) * (uint64(256))) * (uint64(256))) * (uint64(4294967296))) + ((((uint64((bs).Index(_dafny.IntOfAny(uint64(1))).(int))) * (uint64(256))) * (uint64(256))) * (uint64(4294967296)))) + (((uint64((bs).Index(_dafny.IntOfAny(uint64(2))).(int))) * (uint64(256))) * (uint64(4294967296)))) + ((uint64((bs).Index(_dafny.IntOfAny(uint64(3))).(int))) * (uint64(4294967296)))) + ((((uint64((bs).Index(_dafny.IntOfAny(uint64(4))).(int))) * (uint64(256))) * (uint64(256))) * (uint64(256)))) + (((uint64((bs).Index(_dafny.IntOfAny(uint64(5))).(int))) * (uint64(256))) * (uint64(256)))) + ((uint64((bs).Index(_dafny.IntOfAny(uint64(6))).(int))) * (uint64(256)))) + (uint64((bs).Index(_dafny.IntOfAny(uint64(7))).(int)))
}
func (_this *CompanionStruct_Default___) Uint64ToSeqByte(u uint64) _dafny.Seq {
	var _1574_bs _dafny.Seq = _dafny.SeqOf(uint8((u)/(uint64(72057594037927936))), uint8(((u)/(uint64(281474976710656)))%(uint64(256))), uint8(((u)/(uint64(1099511627776)))%(uint64(256))), uint8(((u)/(uint64(4294967296)))%(uint64(256))), uint8(((u)/(uint64(16777216)))%(uint64(256))), uint8(((u)/(uint64(65536)))%(uint64(256))), uint8(((u)/(uint64(256)))%(uint64(256))), uint8((u)%(uint64(256))))
	var _ = _1574_bs
	var _1575_u__int _dafny.Int = _dafny.IntOfUint64(u)
	var _ = _1575_u__int
	return _1574_bs
}
func (_this *CompanionStruct_Default___) SeqByteToUint16(bs _dafny.Seq) uint16 {
	return ((uint16((bs).Index(_dafny.IntOfAny(uint64(0))).(uint8))) * (uint16(256))) + (uint16((bs).Index(_dafny.IntOfAny(uint64(1))).(uint8)))
}
func (_this *CompanionStruct_Default___) Uint16ToSeqByte(u uint16) _dafny.Seq {
	var _1576_s _dafny.Seq = _dafny.SeqOf(uint8(((u)/(uint16(256)))%(uint16(256))), uint8((u)%(uint16(256))))
	var _ = _1576_s
	var _1577_u__int _dafny.Int = _dafny.IntOfUint16(u)
	var _ = _1577_u__int
	return _1576_s
}

// End of class Default__
