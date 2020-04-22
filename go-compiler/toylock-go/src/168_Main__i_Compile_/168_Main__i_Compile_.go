// Package _168_Main__i_Compile
// Dafny module _168_Main__i_Compile compiled into Go

package _168_Main__i_Compile

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
	_129_LockCmdLineParser__i_Compile "129_LockCmdLineParser__i_Compile_"
	_131_Host__i_Compile "131_Host__i_Compile_"
	_133_Lock__DistributedSystem__i_Compile "133_Lock__DistributedSystem__i_Compile_"
	_138_Concrete__NodeIdentity__i_Compile "138_Concrete__NodeIdentity__i_Compile_"
	_143_AbstractServiceLock__s_Compile "143_AbstractServiceLock__s_Compile_"
	_148_Common____SeqIsUnique__i_Compile "148_Common____SeqIsUnique__i_Compile_"
	_152_DistributedSystem__i_Compile "152_DistributedSystem__i_Compile_"
	_155_Refinement__i_Compile "155_Refinement__i_Compile_"
	_161_RefinementProof__i_Compile "161_RefinementProof__i_Compile_"
	_166_MarshallProof__i_Compile "166_MarshallProof__i_Compile_"
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
var _ _129_LockCmdLineParser__i_Compile.Dummy__
var _ _131_Host__i_Compile.Dummy__
var _ _133_Lock__DistributedSystem__i_Compile.Dummy__
var _ _138_Concrete__NodeIdentity__i_Compile.Dummy__
var _ _143_AbstractServiceLock__s_Compile.Dummy__
var _ _148_Common____SeqIsUnique__i_Compile.Dummy__
var _ _152_DistributedSystem__i_Compile.Dummy__
var _ _155_Refinement__i_Compile.Dummy__
var _ _161_RefinementProof__i_Compile.Dummy__
var _ _166_MarshallProof__i_Compile.Dummy__

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
	return "_168_Main__i_Compile.Default__"
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
	return "_168_Main__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) Default_Main_(numRounds int) _dafny.Int {
	var exitCode _dafny.Int = _dafny.Zero
	var _ = exitCode
	var _1761_ok bool
	var _ = _1761_ok
	var _1762_host__state _131_Host__i_Compile.CScheduler
	var _ = _1762_host__state
	var _1763_config _dafny.Seq
	var _ = _1763_config
	var _1764_id _9_Native____Io__s_Compile.EndPoint
	var _ = _1764_id
	var _out107 bool
	var _ = _out107
	var _out108 _131_Host__i_Compile.CScheduler
	var _ = _out108
	var _out109 _dafny.Seq
	var _ = _out109
	var _out110 _9_Native____Io__s_Compile.EndPoint
	var _ = _out110
	_out107, _out108, _out109, _out110 = _131_Host__i_Compile.Companion_Default___.HostInitImpl()
	_1761_ok = _out107
	_1762_host__state = _out108
	_1763_config = _out109
	_1764_id = _out110
	{
	}
	var counter = 0
	for _1761_ok {
		{
		}
		{
		}
		{
		}
		var _out111 bool
		var _ = _out111
		var _out112 _131_Host__i_Compile.CScheduler
		var _ = _out112
		_out111, _out112 = _131_Host__i_Compile.Companion_Default___.HostNextImpl(_1762_host__state)
		counter++
		_1761_ok = _out111
		_1762_host__state = _out112
		{
		}
		if counter == numRounds {
			return exitCode
		}
	}
	return exitCode
}

// End of class Default__
