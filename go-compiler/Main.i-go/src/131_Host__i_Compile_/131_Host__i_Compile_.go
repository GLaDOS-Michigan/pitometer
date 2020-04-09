// Package _131_Host__i_Compile
// Dafny module _131_Host__i_Compile compiled into Go

package _131_Host__i_Compile

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

type Dummy__ struct{}





// Definition of data type CScheduler
type CScheduler struct {
  Data_CScheduler_
}

func (_this CScheduler) Get() Data_CScheduler_ {
  return _this.Data_CScheduler_
}

type Data_CScheduler_ interface {
  isCScheduler()
}

type CompanionStruct_CScheduler_ struct {}
var Companion_CScheduler_ = CompanionStruct_CScheduler_{}

type CScheduler_CScheduler struct {
  Node__impl *_121_NodeImpl__i_Compile.NodeImpl
}

func (CScheduler_CScheduler) isCScheduler() {}

func (CompanionStruct_CScheduler_) Create_CScheduler_(Node__impl *_121_NodeImpl__i_Compile.NodeImpl) CScheduler {
  return CScheduler{CScheduler_CScheduler{Node__impl}}
}

func (_this CScheduler) Is_CScheduler() bool {
  _, ok := _this.Get().(CScheduler_CScheduler)
return ok
}

func (_this CScheduler) Dtor_node__impl() *_121_NodeImpl__i_Compile.NodeImpl {
  return _this.Get().(CScheduler_CScheduler).Node__impl
}

func (_this CScheduler) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case CScheduler_CScheduler: {
      return "_131_Host__i_Compile.CScheduler.CScheduler" + "(" + _dafny.String(data.Node__impl) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this CScheduler) Equals(other CScheduler) bool {
  switch data1 := _this.Get().(type) {
    case CScheduler_CScheduler: {
      data2, ok := other.Get().(CScheduler_CScheduler)
return ok && data1.Node__impl == data2.Node__impl
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this CScheduler) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(CScheduler)
return ok && _this.Equals(typed)
}
func Type_CScheduler_() _dafny.Type {
  return type_CScheduler_{}
}

type type_CScheduler_ struct {
}

func (_this type_CScheduler_) Default() interface{} {
  return CScheduler{CScheduler_CScheduler{(*_121_NodeImpl__i_Compile.NodeImpl)(nil)}}
}

func (_this type_CScheduler_) String() string {
  return "_131_Host__i_Compile.CScheduler"
}
// End of data type CScheduler




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
var Companion_Default___ = CompanionStruct_Default___ {
}

func (_this *Default__) Equals(other *Default__) bool {
  return _this == other
}

func (_this *Default__) EqualsGeneric(x interface{}) bool {
  other, ok := x.(*Default__)
return ok && _this.Equals(other)
}

func (*Default__) String() string {
  return "_131_Host__i_Compile.Default__"
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
  return "_131_Host__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) HostInitImpl() (bool, _131_Host__i_Compile.CScheduler, _dafny.Seq, _9_Native____Io__s_Compile.EndPoint) {
  goto TAIL_CALL_START
TAIL_CALL_START:
var ok bool = false
  var _ = ok
var host__state _131_Host__i_Compile.CScheduler = Type_CScheduler_().Default().(_131_Host__i_Compile.CScheduler)
  var _ = host__state
var config _dafny.Seq = _dafny.EmptySeq
  var _ = config
var id _9_Native____Io__s_Compile.EndPoint = _9_Native____Io__s_Compile.Type_EndPoint_().Default().(_9_Native____Io__s_Compile.EndPoint)
  var _ = id
  var _1747_my__index uint64 = 0
  var _ = _1747_my__index
  var _out101 bool
  var _ = _out101
var _out102 _dafny.Seq
  var _ = _out102
var _out103 uint64
  var _ = _out103
_out101,_out102,_out103 = _129_LockCmdLineParser__i_Compile.Companion_Default___.ParseCmdLine()
ok = _out101
config = _out102
_1747_my__index = _out103
  if (!(ok)) {
    return ok,host__state,config,id
  }
  id = (config).Index(_1747_my__index).(_9_Native____Io__s_Compile.EndPoint)
  var _1748_node__impl *_121_NodeImpl__i_Compile.NodeImpl
  var _ = _1748_node__impl
  var _nw7 = _121_NodeImpl__i_Compile.New_NodeImpl_()
  var _ = _nw7
_nw7.Ctor__()
  _1748_node__impl = _nw7
  var _out104 bool
  var _ = _out104
_out104 = (_1748_node__impl).InitNode(config, _1747_my__index)
ok = _out104
  if (!(ok)) {
    return ok,host__state,config,id
  }
  host__state = CScheduler{CScheduler_CScheduler{_1748_node__impl}}
  { }
  { }
  return ok,host__state,config,id
}
func (_this *CompanionStruct_Default___) HostNextImpl(host__state _131_Host__i_Compile.CScheduler) (bool, _131_Host__i_Compile.CScheduler) {
  var ok bool = false
  var _ = ok
var host__state_k _131_Host__i_Compile.CScheduler = Type_CScheduler_().Default().(_131_Host__i_Compile.CScheduler)
  var _ = host__state_k
  var _1749_okay bool
  var _ = _1749_okay
var _out105 bool
  var _ = _out105
_out105 = ((host__state).Dtor_node__impl()).HostNextMain()
_1749_okay = _out105
  if (_1749_okay) {
    { }
    { }
    { }
    { }
    host__state_k = CScheduler{CScheduler_CScheduler{(host__state).Dtor_node__impl()}}
  } else { }
  ok = _1749_okay
  return ok,host__state_k
}
// End of class Default__



