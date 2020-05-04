// Package _129_LockCmdLineParser__i_Compile
// Dafny module _129_LockCmdLineParser__i_Compile compiled into Go

package _129_LockCmdLineParser__i_Compile

import (
	_0_Native____NativeTypes__s_Compile "0_Native____NativeTypes__s_Compile_"
	_101_Common____Util__i_Compile "101_Common____Util__i_Compile_"
	_105_Common____MarshallInt__i_Compile "105_Common____MarshallInt__i_Compile_"
	_107_Common____GenericMarshalling__i_Compile "107_Common____GenericMarshalling__i_Compile_"
	_111_PacketParsing__i_Compile "111_PacketParsing__i_Compile_"
	_113_Common____SeqIsUniqueDef__i_Compile "113_Common____SeqIsUniqueDef__i_Compile_"
	_115_Impl__Node__i_Compile "115_Impl__Node__i_Compile_"
	_119_UdpLock__i_Compile "119_UdpLock__i_Compile_"
	_121_NodeImpl__i_Compile "121_NodeImpl__i_Compile_"
	_127_CmdLineParser__i_Compile "127_CmdLineParser__i_Compile_"
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
var _ _101_Common____Util__i_Compile.Dummy__
var _ _105_Common____MarshallInt__i_Compile.Dummy__
var _ _107_Common____GenericMarshalling__i_Compile.Dummy__
var _ _111_PacketParsing__i_Compile.Dummy__
var _ _113_Common____SeqIsUniqueDef__i_Compile.Dummy__
var _ _115_Impl__Node__i_Compile.Dummy__
var _ _119_UdpLock__i_Compile.Dummy__
var _ _121_NodeImpl__i_Compile.Dummy__
var _ _127_CmdLineParser__i_Compile.Dummy__

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
	return "_129_LockCmdLineParser__i_Compile.Default__"
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
	return "_129_LockCmdLineParser__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) EndPointNull() _9_Native____Io__s_Compile.EndPoint {
	return _9_Native____Io__s_Compile.EndPoint{_9_Native____Io__s_Compile.EndPoint_EndPoint{_dafny.SeqOf(uint8(0), uint8(0), uint8(0), uint8(0)), uint16(0)}}
}
func (_this *CompanionStruct_Default___) GetHostIndex(host _9_Native____Io__s_Compile.EndPoint, hosts _dafny.Seq) (bool, uint64) {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var found bool = false
	var _ = found
	var index uint64 = 0
	var _ = index
	var _1739_i uint64
	var _ = _1739_i
	_1739_i = uint64(0)
	for (_1739_i) < (uint64((hosts).CardinalityInt64())) {
		if (host).Equals((hosts).IndexUint64(_1739_i).(_9_Native____Io__s_Compile.EndPoint)) {
			found = true
			index = _1739_i
			{
			}
			return found, index
		}
		if (_1739_i) == ((uint64((hosts).CardinalityInt64())) - (func() uint64 { return (uint64(1)) })()) {
			found = false
			return found, index
		}
		_1739_i = (_1739_i) + (uint64(1))
	}
	found = false
	return found, index
}
func (_this *CompanionStruct_Default___) ParseCmdLine() (bool, _dafny.Seq, uint64) {
	var ok bool = false
	var _ = ok
	var host__ids _dafny.Seq = _dafny.EmptySeq
	var _ = host__ids
	var my__index uint64 = 0
	var _ = my__index
	ok = false
	var _1740_num__args uint32
	var _ = _1740_num__args
	var _out96 uint32
	var _ = _out96
	_out96 = _9_Native____Io__s_Compile.Companion_HostConstants_.NumCommandLineArgs()
	_1740_num__args = _out96
	if ((_1740_num__args) < (uint32(4))) || (((_1740_num__args) % (uint32(2))) != (uint32(1)) /* dircomp */) {
		_dafny.Print(_dafny.SeqOfString("Error: Incorrect number of command line arguments. " +
			"Expected: ./Main.exe [IP port]+ [IP port]. " +
			"where the final argument is one of the two IP-port pairs provided earlier"))
		return ok, host__ids, my__index
	}
	var _1741_args _dafny.Seq
	var _ = _1741_args
	var _out97 _dafny.Seq
	var _ = _out97
	_out97 = _127_CmdLineParser__i_Compile.Companion_Default___.Collect__cmd__line__args()
	_1741_args = _out97
	{
	}
	var _1742_tuple1 _System.Tuple2
	var _ = _1742_tuple1
	_1742_tuple1 = _127_CmdLineParser__i_Compile.Companion_Default___.Parse__end__points((_1741_args).Subseq(_dafny.IntOfInt64(1), ((_1741_args).Cardinality()).Minus(_dafny.IntOfInt64(2))))
	ok = (*(_1742_tuple1).IndexInt(0)).(bool)
	var _1743_endpoints _dafny.Seq
	var _ = _1743_endpoints
	_1743_endpoints = (*(_1742_tuple1).IndexInt(1)).(_dafny.Seq)
	if ((!(ok)) || (((_1743_endpoints).Cardinality()).Cmp(_dafny.Zero) == 0)) || (((_1743_endpoints).Cardinality()).Cmp(_dafny.IntOfString("18446744073709551616")) >= 0) {
		ok = false
		return ok, host__ids, my__index
	}
	var _1744_tuple2 _System.Tuple2
	var _ = _1744_tuple2
	_1744_tuple2 = _127_CmdLineParser__i_Compile.Companion_Default___.Parse__end__point((_1741_args).Index(((_1741_args).Cardinality()).Minus(_dafny.IntOfInt64(2))).(_dafny.Seq), (_1741_args).Index(((_1741_args).Cardinality()).Minus(_dafny.IntOfInt64(1))).(_dafny.Seq))
	ok = (*(_1744_tuple2).IndexInt(0)).(bool)
	if !(ok) {
		return ok, host__ids, my__index
	}
	var _1745_unique bool
	var _ = _1745_unique
	var _out98 bool
	var _ = _out98
	_out98 = _127_CmdLineParser__i_Compile.Companion_Default___.Test__unique_k(_1743_endpoints)
	_1745_unique = _out98
	if !(_1745_unique) {
		ok = false
		return ok, host__ids, my__index
	}
	var _out99 bool
	var _ = _out99
	var _out100 uint64
	var _ = _out100
	_out99, _out100 = Companion_Default___.GetHostIndex((*(_1744_tuple2).IndexInt(1)).(_9_Native____Io__s_Compile.EndPoint), _1743_endpoints)
	ok = _out99
	my__index = _out100
	if !(ok) {
		return ok, host__ids, my__index
	}
	host__ids = _1743_endpoints
	var _1746_me _9_Native____Io__s_Compile.EndPoint
	var _ = _1746_me
	_1746_me = (_1743_endpoints).IndexUint64(my__index).(_9_Native____Io__s_Compile.EndPoint)
	{
	}
	{
	}
	{
	}
	{
	}
	return ok, host__ids, my__index
}

// End of class Default__
