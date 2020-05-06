// Package _143_AbstractServiceLock__s_Compile
// Dafny module _143_AbstractServiceLock__s_Compile compiled into Go

package _143_AbstractServiceLock__s_Compile

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
_133_Lock__DistributedSystem__i_Compile "133_Lock__DistributedSystem__i_Compile_"
_138_Concrete__NodeIdentity__i_Compile "138_Concrete__NodeIdentity__i_Compile_"
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

type Dummy__ struct{}


// Definition of data type ServiceState_k
type ServiceState_k struct {
  Data_ServiceState_k_
}

func (_this ServiceState_k) Get() Data_ServiceState_k_ {
  return _this.Data_ServiceState_k_
}

type Data_ServiceState_k_ interface {
  isServiceState_k()
}

type CompanionStruct_ServiceState_k_ struct {}
var Companion_ServiceState_k_ = CompanionStruct_ServiceState_k_{}

type ServiceState_k_ServiceState_k struct {
  Hosts _dafny.Set
History _dafny.Seq
}

func (ServiceState_k_ServiceState_k) isServiceState_k() {}

func (CompanionStruct_ServiceState_k_) Create_ServiceState_k_(Hosts _dafny.Set, History _dafny.Seq) ServiceState_k {
  return ServiceState_k{ServiceState_k_ServiceState_k{Hosts,History}}
}

func (_this ServiceState_k) Is_ServiceState_k() bool {
  _, ok := _this.Get().(ServiceState_k_ServiceState_k)
return ok
}

func (_this ServiceState_k) Dtor_hosts() _dafny.Set {
  return _this.Get().(ServiceState_k_ServiceState_k).Hosts
}

func (_this ServiceState_k) Dtor_history() _dafny.Seq {
  return _this.Get().(ServiceState_k_ServiceState_k).History
}

func (_this ServiceState_k) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case ServiceState_k_ServiceState_k: {
      return "_143_AbstractServiceLock__s_Compile.ServiceState'.ServiceState'" + "(" + _dafny.String(data.Hosts) + ", " + _dafny.String(data.History) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this ServiceState_k) Equals(other ServiceState_k) bool {
  switch data1 := _this.Get().(type) {
    case ServiceState_k_ServiceState_k: {
      data2, ok := other.Get().(ServiceState_k_ServiceState_k)
return ok && data1.Hosts.Equals(data2.Hosts) && data1.History.Equals(data2.History)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this ServiceState_k) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(ServiceState_k)
return ok && _this.Equals(typed)
}
func Type_ServiceState_k_() _dafny.Type {
  return type_ServiceState_k_{}
}

type type_ServiceState_k_ struct {
}

func (_this type_ServiceState_k_) Default() interface{} {
  return ServiceState_k{ServiceState_k_ServiceState_k{_dafny.EmptySet, _dafny.EmptySeq}}
}

func (_this type_ServiceState_k_) String() string {
  return "_143_AbstractServiceLock__s_Compile.ServiceState_k"
}
// End of data type ServiceState_k





