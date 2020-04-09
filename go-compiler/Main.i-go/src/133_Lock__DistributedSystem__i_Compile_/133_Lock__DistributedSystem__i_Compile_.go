// Package _133_Lock__DistributedSystem__i_Compile
// Dafny module _133_Lock__DistributedSystem__i_Compile compiled into Go

package _133_Lock__DistributedSystem__i_Compile

import (
  _dafny "dafny"
_System "System_"
_0_Native____NativeTypes__s_Compile "0_Native____NativeTypes__s_Compile_"
_2_Collections____Maps2__s_Compile "2_Collections____Maps2__s_Compile_"
_5_Temporal____Temporal__s_Compile "5_Temporal____Temporal__s_Compile_"
_7_Environment__s_Compile "7_Environment__s_Compile_"
_9_Native____Io__s_Compile "9_Native____Io__s_Compile_"
_26_Collections____Seqs__s_Compile "26_Collections____Seqs__s_Compile_"
_29_Collections____Sets__i_Compile "29_Collections____Sets__i_Compile_"
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
_61_Math____power__s_Compile "61_Math____power__s_Compile_"
_64_Math____mul__nonlinear__i_Compile "64_Math____mul__nonlinear__i_Compile_"
_67_Math____mul__auto__proofs__i_Compile "67_Math____mul__auto__proofs__i_Compile_"
_69_Math____mul__auto__i_Compile "69_Math____mul__auto__i_Compile_"
_71_Math____mul__i_Compile "71_Math____mul__i_Compile_"
_73_Math____power__i_Compile "73_Math____power__i_Compile_"
_77_Math____div__def__i_Compile "77_Math____div__def__i_Compile_"
_81_Math____div__boogie__i_Compile "81_Math____div__boogie__i_Compile_"
_83_Math____div__nonlinear__i_Compile "83_Math____div__nonlinear__i_Compile_"
_88_Math____mod__auto__proofs__i_Compile "88_Math____mod__auto__proofs__i_Compile_"
_90_Math____mod__auto__i_Compile "90_Math____mod__auto__i_Compile_"
_93_Math____div__auto__proofs__i_Compile "93_Math____div__auto__proofs__i_Compile_"
_95_Math____div__auto__i_Compile "95_Math____div__auto__i_Compile_"
_97_Math____div__i_Compile "97_Math____div__i_Compile_"
_99_Math____power2__i_Compile "99_Math____power2__i_Compile_"
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

type Dummy__ struct{}








// Definition of data type DS__State
type DS__State struct {
  Data_DS__State_
}

func (_this DS__State) Get() Data_DS__State_ {
  return _this.Data_DS__State_
}

type Data_DS__State_ interface {
  isDS__State()
}

type CompanionStruct_DS__State_ struct {}
var Companion_DS__State_ = CompanionStruct_DS__State_{}

type DS__State_DS__State struct {
  Config _dafny.Seq
Environment _7_Environment__s_Compile.LEnvironment
Servers _dafny.Map
Clients _dafny.Set
}

func (DS__State_DS__State) isDS__State() {}

func (CompanionStruct_DS__State_) Create_DS__State_(Config _dafny.Seq, Environment _7_Environment__s_Compile.LEnvironment, Servers _dafny.Map, Clients _dafny.Set) DS__State {
  return DS__State{DS__State_DS__State{Config,Environment,Servers,Clients}}
}

func (_this DS__State) Is_DS__State() bool {
  _, ok := _this.Get().(DS__State_DS__State)
return ok
}

func (_this DS__State) Dtor_config() _dafny.Seq {
  return _this.Get().(DS__State_DS__State).Config
}

func (_this DS__State) Dtor_environment() _7_Environment__s_Compile.LEnvironment {
  return _this.Get().(DS__State_DS__State).Environment
}

func (_this DS__State) Dtor_servers() _dafny.Map {
  return _this.Get().(DS__State_DS__State).Servers
}

func (_this DS__State) Dtor_clients() _dafny.Set {
  return _this.Get().(DS__State_DS__State).Clients
}

func (_this DS__State) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case DS__State_DS__State: {
      return "_133_Lock__DistributedSystem__i_Compile.DS_State.DS_State" + "(" + _dafny.String(data.Config) + ", " + _dafny.String(data.Environment) + ", " + _dafny.String(data.Servers) + ", " + _dafny.String(data.Clients) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this DS__State) Equals(other DS__State) bool {
  switch data1 := _this.Get().(type) {
    case DS__State_DS__State: {
      data2, ok := other.Get().(DS__State_DS__State)
return ok && data1.Config.Equals(data2.Config) && data1.Environment.Equals(data2.Environment) && data1.Servers.Equals(data2.Servers) && data1.Clients.Equals(data2.Clients)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this DS__State) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(DS__State)
return ok && _this.Equals(typed)
}
func Type_DS__State_() _dafny.Type {
  return type_DS__State_{}
}

type type_DS__State_ struct {
}

func (_this type_DS__State_) Default() interface{} {
  return DS__State{DS__State_DS__State{_dafny.EmptySeq, _7_Environment__s_Compile.Type_LEnvironment_().Default().(_7_Environment__s_Compile.LEnvironment), _dafny.EmptyMap, _dafny.EmptySet}}
}

func (_this type_DS__State_) String() string {
  return "_133_Lock__DistributedSystem__i_Compile.DS__State"
}
// End of data type DS__State
