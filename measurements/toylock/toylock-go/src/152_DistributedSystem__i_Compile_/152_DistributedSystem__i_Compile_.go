// Package _152_DistributedSystem__i_Compile
// Dafny module _152_DistributedSystem__i_Compile compiled into Go

package _152_DistributedSystem__i_Compile

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

type Dummy__ struct{}

// Definition of data type LS__State
type LS__State struct {
	Data_LS__State_
}

func (_this LS__State) Get() Data_LS__State_ {
	return _this.Data_LS__State_
}

type Data_LS__State_ interface {
	isLS__State()
}

type CompanionStruct_LS__State_ struct{}

var Companion_LS__State_ = CompanionStruct_LS__State_{}

type LS__State_LS__State struct {
	Environment _7_Environment__s_Compile.LEnvironment
	Servers     _dafny.Map
}

func (LS__State_LS__State) isLS__State() {}

func (CompanionStruct_LS__State_) Create_LS__State_(Environment _7_Environment__s_Compile.LEnvironment, Servers _dafny.Map) LS__State {
	return LS__State{LS__State_LS__State{Environment, Servers}}
}

func (_this LS__State) Is_LS__State() bool {
	_, ok := _this.Get().(LS__State_LS__State)
	return ok
}

func (_this LS__State) Dtor_environment() _7_Environment__s_Compile.LEnvironment {
	return _this.Get().(LS__State_LS__State).Environment
}

func (_this LS__State) Dtor_servers() _dafny.Map {
	return _this.Get().(LS__State_LS__State).Servers
}

func (_this LS__State) String() string {
	switch data := _this.Get().(type) {
	case nil:
		return "null"
	case LS__State_LS__State:
		{
			return "_152_DistributedSystem__i_Compile.LS_State.LS_State" + "(" + _dafny.String(data.Environment) + ", " + _dafny.String(data.Servers) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this LS__State) Equals(other LS__State) bool {
	switch data1 := _this.Get().(type) {
	case LS__State_LS__State:
		{
			data2, ok := other.Get().(LS__State_LS__State)
			return ok && data1.Environment.Equals(data2.Environment) && data1.Servers.Equals(data2.Servers)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this LS__State) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(LS__State)
	return ok && _this.Equals(typed)
}
func Type_LS__State_() _dafny.Type {
	return type_LS__State_{}
}

type type_LS__State_ struct {
}

func (_this type_LS__State_) Default() interface{} {
	return LS__State{LS__State_LS__State{_7_Environment__s_Compile.Type_LEnvironment_().Default().(_7_Environment__s_Compile.LEnvironment), _dafny.EmptyMap}}
}

func (_this type_LS__State_) String() string {
	return "LS__State"
}

// End of data type LS__State

// Definition of data type GLS__State
type GLS__State struct {
	Data_GLS__State_
}

func (_this GLS__State) Get() Data_GLS__State_ {
	return _this.Data_GLS__State_
}

type Data_GLS__State_ interface {
	isGLS__State()
}

type CompanionStruct_GLS__State_ struct{}

var Companion_GLS__State_ = CompanionStruct_GLS__State_{}

type GLS__State_GLS__State struct {
	Ls      LS__State
	History _dafny.Seq
}

func (GLS__State_GLS__State) isGLS__State() {}

func (CompanionStruct_GLS__State_) Create_GLS__State_(Ls LS__State, History _dafny.Seq) GLS__State {
	return GLS__State{GLS__State_GLS__State{Ls, History}}
}

func (_this GLS__State) Is_GLS__State() bool {
	_, ok := _this.Get().(GLS__State_GLS__State)
	return ok
}

func (_this GLS__State) Dtor_ls() LS__State {
	return _this.Get().(GLS__State_GLS__State).Ls
}

func (_this GLS__State) Dtor_history() _dafny.Seq {
	return _this.Get().(GLS__State_GLS__State).History
}

func (_this GLS__State) String() string {
	switch data := _this.Get().(type) {
	case nil:
		return "null"
	case GLS__State_GLS__State:
		{
			return "_152_DistributedSystem__i_Compile.GLS_State.GLS_State" + "(" + _dafny.String(data.Ls) + ", " + _dafny.String(data.History) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this GLS__State) Equals(other GLS__State) bool {
	switch data1 := _this.Get().(type) {
	case GLS__State_GLS__State:
		{
			data2, ok := other.Get().(GLS__State_GLS__State)
			return ok && data1.Ls.Equals(data2.Ls) && data1.History.Equals(data2.History)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this GLS__State) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(GLS__State)
	return ok && _this.Equals(typed)
}
func Type_GLS__State_() _dafny.Type {
	return type_GLS__State_{}
}

type type_GLS__State_ struct {
}

func (_this type_GLS__State_) Default() interface{} {
	return GLS__State{GLS__State_GLS__State{Type_LS__State_().Default().(LS__State), _dafny.EmptySeq}}
}

func (_this type_GLS__State_) String() string {
	return "_152_DistributedSystem__i_Compile.GLS__State"
}

// End of data type GLS__State
