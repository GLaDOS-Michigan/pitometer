// Package _228_LiveRSL____PacketParsing__i_Compile
// Dafny module _228_LiveRSL____PacketParsing__i_Compile compiled into Go

package _228_LiveRSL____PacketParsing__i_Compile

import (
  _dafny "dafny"
_System "System_"
_0_Native____NativeTypes__s_Compile "0_Native____NativeTypes__s_Compile_"
_2_Collections____Maps2__s_Compile "2_Collections____Maps2__s_Compile_"
_5_Temporal____Temporal__s_Compile "5_Temporal____Temporal__s_Compile_"
_7_Environment__s_Compile "7_Environment__s_Compile_"
_9_Native____Io__s_Compile "9_Native____Io__s_Compile_"
_26_Collections____Seqs__s_Compile "26_Collections____Seqs__s_Compile_"
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
_61_LiveRSL____Configuration__i_Compile "61_LiveRSL____Configuration__i_Compile_"
_65_LiveRSL____Message__i_Compile "65_LiveRSL____Message__i_Compile_"
_68_LiveRSL____Environment__i_Compile "68_LiveRSL____Environment__i_Compile_"
_71_LiveRSL____ClockReading__i_Compile "71_LiveRSL____ClockReading__i_Compile_"
_74_Common____UpperBound__s_Compile "74_Common____UpperBound__s_Compile_"
_76_LiveRSL____Parameters__i_Compile "76_LiveRSL____Parameters__i_Compile_"
_78_LiveRSL____Constants__i_Compile "78_LiveRSL____Constants__i_Compile_"
_85_LiveRSL____Broadcast__i_Compile "85_LiveRSL____Broadcast__i_Compile_"
_91_Collections____CountMatches__i_Compile "91_Collections____CountMatches__i_Compile_"
_93_LiveRSL____Acceptor__i_Compile "93_LiveRSL____Acceptor__i_Compile_"
_99_LiveRSL____Election__i_Compile "99_LiveRSL____Election__i_Compile_"
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
  return "_228_LiveRSL____PacketParsing__i_Compile.Default__"
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
  return "_228_LiveRSL____PacketParsing__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) EndPoint__grammar() _176_Common____GenericMarshalling__i_Compile.G {
  return _176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GUint64{}}
}
func (_this *CompanionStruct_Default___) CRequest__grammar() _176_Common____GenericMarshalling__i_Compile.G {
  return _176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GTuple{_dafny.SeqOf(Companion_Default___.EndPoint__grammar(), _176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GUint64{}}, _197_LiveRSL____AppInterface__i_Compile.Companion_Default___.CAppMessage__grammar())}}
}
func (_this *CompanionStruct_Default___) CRequestBatch__grammar() _176_Common____GenericMarshalling__i_Compile.G {
  return _176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GArray{Companion_Default___.CRequest__grammar()}}
}
func (_this *CompanionStruct_Default___) CReply__grammar() _176_Common____GenericMarshalling__i_Compile.G {
  return _176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GTuple{_dafny.SeqOf(Companion_Default___.EndPoint__grammar(), _176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GUint64{}}, _197_LiveRSL____AppInterface__i_Compile.Companion_Default___.CAppMessage__grammar())}}
}
func (_this *CompanionStruct_Default___) CBallot__grammar() _176_Common____GenericMarshalling__i_Compile.G {
  return _176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GTuple{_dafny.SeqOf(_176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GUint64{}}, _176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GUint64{}})}}
}
func (_this *CompanionStruct_Default___) COperationNumber__grammar() _176_Common____GenericMarshalling__i_Compile.G {
  return _176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GUint64{}}
}
func (_this *CompanionStruct_Default___) CVote__grammar() _176_Common____GenericMarshalling__i_Compile.G {
  return _176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GTuple{_dafny.SeqOf(Companion_Default___.CBallot__grammar(), Companion_Default___.CRequestBatch__grammar())}}
}
func (_this *CompanionStruct_Default___) CMap__grammar(key _176_Common____GenericMarshalling__i_Compile.G, val _176_Common____GenericMarshalling__i_Compile.G) _176_Common____GenericMarshalling__i_Compile.G {
  return _176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GArray{_176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GTuple{_dafny.SeqOf(key, val)}}}}
}
func (_this *CompanionStruct_Default___) CVotes__grammar() _176_Common____GenericMarshalling__i_Compile.G {
  return _176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GArray{_176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GTuple{_dafny.SeqOf(Companion_Default___.COperationNumber__grammar(), Companion_Default___.CVote__grammar())}}}}
}
func (_this *CompanionStruct_Default___) CReplyCache__grammar() _176_Common____GenericMarshalling__i_Compile.G {
  return _176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GArray{_176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GTuple{_dafny.SeqOf(Companion_Default___.EndPoint__grammar(), Companion_Default___.CReply__grammar())}}}}
}
func (_this *CompanionStruct_Default___) CMessage__Request__grammar() _176_Common____GenericMarshalling__i_Compile.G {
  return _176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GTuple{_dafny.SeqOf(_176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GUint64{}}, _197_LiveRSL____AppInterface__i_Compile.Companion_Default___.CAppMessage__grammar())}}
}
func (_this *CompanionStruct_Default___) CMessage__1a__grammar() _176_Common____GenericMarshalling__i_Compile.G {
  return Companion_Default___.CBallot__grammar()
}
func (_this *CompanionStruct_Default___) CMessage__1b__grammar() _176_Common____GenericMarshalling__i_Compile.G {
  return _176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GTuple{_dafny.SeqOf(Companion_Default___.CBallot__grammar(), Companion_Default___.COperationNumber__grammar(), Companion_Default___.CVotes__grammar())}}
}
func (_this *CompanionStruct_Default___) CMessage__2a__grammar() _176_Common____GenericMarshalling__i_Compile.G {
  return _176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GTuple{_dafny.SeqOf(Companion_Default___.CBallot__grammar(), Companion_Default___.COperationNumber__grammar(), Companion_Default___.CRequestBatch__grammar())}}
}
func (_this *CompanionStruct_Default___) CMessage__2b__grammar() _176_Common____GenericMarshalling__i_Compile.G {
  return _176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GTuple{_dafny.SeqOf(Companion_Default___.CBallot__grammar(), Companion_Default___.COperationNumber__grammar(), Companion_Default___.CRequestBatch__grammar())}}
}
func (_this *CompanionStruct_Default___) CMessage__Heartbeat__grammar() _176_Common____GenericMarshalling__i_Compile.G {
  return _176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GTuple{_dafny.SeqOf(Companion_Default___.CBallot__grammar(), _176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GUint64{}}, Companion_Default___.COperationNumber__grammar())}}
}
func (_this *CompanionStruct_Default___) CMessage__Reply__grammar() _176_Common____GenericMarshalling__i_Compile.G {
  return _176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GTuple{_dafny.SeqOf(_176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GUint64{}}, _197_LiveRSL____AppInterface__i_Compile.Companion_Default___.CAppMessage__grammar())}}
}
func (_this *CompanionStruct_Default___) CMessage__AppStateRequest__grammar() _176_Common____GenericMarshalling__i_Compile.G {
  return _176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GTuple{_dafny.SeqOf(Companion_Default___.CBallot__grammar(), Companion_Default___.COperationNumber__grammar())}}
}
func (_this *CompanionStruct_Default___) CMessage__AppStateSupply__grammar() _176_Common____GenericMarshalling__i_Compile.G {
  return _176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GTuple{_dafny.SeqOf(Companion_Default___.CBallot__grammar(), Companion_Default___.COperationNumber__grammar(), _197_LiveRSL____AppInterface__i_Compile.Companion_Default___.CAppState__grammar(), Companion_Default___.CReplyCache__grammar())}}
}
func (_this *CompanionStruct_Default___) CMessage__StartingPhase2__grammar() _176_Common____GenericMarshalling__i_Compile.G {
  return _176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GTuple{_dafny.SeqOf(Companion_Default___.CBallot__grammar(), Companion_Default___.COperationNumber__grammar())}}
}
func (_this *CompanionStruct_Default___) CMessage__grammar() _176_Common____GenericMarshalling__i_Compile.G {
  return _176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GTaggedUnion{_dafny.SeqOf(Companion_Default___.CMessage__Request__grammar(), Companion_Default___.CMessage__1a__grammar(), Companion_Default___.CMessage__1b__grammar(), Companion_Default___.CMessage__2a__grammar(), Companion_Default___.CMessage__2b__grammar(), Companion_Default___.CMessage__Heartbeat__grammar(), Companion_Default___.CMessage__Reply__grammar(), Companion_Default___.CMessage__AppStateRequest__grammar(), Companion_Default___.CMessage__AppStateSupply__grammar(), Companion_Default___.CMessage__StartingPhase2__grammar())}}
}
func (_this *CompanionStruct_Default___) Parse__EndPoint(val _176_Common____GenericMarshalling__i_Compile.V) _9_Native____Io__s_Compile.EndPoint {
  if (((val).Dtor_u()) <= (uint64(281474976710655))) {
    return _194_Common____NodeIdentity__i_Compile.Companion_Default___.ConvertUint64ToEndPoint((val).Dtor_u())
  } else  {
    return _9_Native____Io__s_Compile.EndPoint{_9_Native____Io__s_Compile.EndPoint_EndPoint{_dafny.SeqOf(uint8(0), uint8(0), uint8(0), uint8(0)), uint16(0)}}
  }
}
func (_this *CompanionStruct_Default___) Parse__Request(val _176_Common____GenericMarshalling__i_Compile.V) _214_LiveRSL____CTypes__i_Compile.CRequest {
  var _4639_ep _9_Native____Io__s_Compile.EndPoint = Companion_Default___.Parse__EndPoint(((val).Dtor_t()).Index(_dafny.Zero).(_176_Common____GenericMarshalling__i_Compile.V))
  var _ = _4639_ep
return _214_LiveRSL____CTypes__i_Compile.CRequest{_214_LiveRSL____CTypes__i_Compile.CRequest_CRequest{_4639_ep, (((val).Dtor_t()).Index(_dafny.IntOfInt64(1)).(_176_Common____GenericMarshalling__i_Compile.V)).Dtor_u(), _197_LiveRSL____AppInterface__i_Compile.Companion_Default___.Parse__AppMessage(((val).Dtor_t()).Index(_dafny.IntOfInt64(2)).(_176_Common____GenericMarshalling__i_Compile.V))}}
}
func (_this *CompanionStruct_Default___) Parse__RequestBatch(val _176_Common____GenericMarshalling__i_Compile.V) _dafny.Seq {
  goto TAIL_CALL_START
TAIL_CALL_START:
var batch _dafny.Seq = _dafny.EmptySeq
  var _ = batch
  var _4640_batchArr *_dafny.Array
  var _ = _4640_batchArr
  var _nw8 = _dafny.NewArrayWithValue(_214_LiveRSL____CTypes__i_Compile.Type_CRequest_().Default().(_214_LiveRSL____CTypes__i_Compile.CRequest), uint64(((val).Dtor_a()).CardinalityInt()))
  var _ = _nw8
  _4640_batchArr = _nw8
  var _4641_i uint64
  var _ = _4641_i
  _4641_i = uint64(0)
  for (_4641_i) < (uint64(((val).Dtor_a()).CardinalityInt())) {
    var _4642_req _214_LiveRSL____CTypes__i_Compile.CRequest
    var _ = _4642_req
    _4642_req = Companion_Default___.Parse__Request(((val).Dtor_a()).Index(_4641_i).(_176_Common____GenericMarshalling__i_Compile.V))
    *((_4640_batchArr).Index(_dafny.IntOfAny((_4641_i)))) = _4642_req
    _4641_i = (_4641_i) + (uint64(1))
  }
  batch = (_4640_batchArr).RangeToSeq(_dafny.NilInt, _dafny.NilInt)
  return batch
}
func (_this *CompanionStruct_Default___) Parse__Reply(val _176_Common____GenericMarshalling__i_Compile.V) _214_LiveRSL____CTypes__i_Compile.CReply {
  var _4643_ep _9_Native____Io__s_Compile.EndPoint = Companion_Default___.Parse__EndPoint(((val).Dtor_t()).Index(_dafny.Zero).(_176_Common____GenericMarshalling__i_Compile.V))
  var _ = _4643_ep
return _214_LiveRSL____CTypes__i_Compile.CReply{_214_LiveRSL____CTypes__i_Compile.CReply_CReply{_4643_ep, (((val).Dtor_t()).Index(_dafny.IntOfInt64(1)).(_176_Common____GenericMarshalling__i_Compile.V)).Dtor_u(), _197_LiveRSL____AppInterface__i_Compile.Companion_Default___.Parse__AppMessage(((val).Dtor_t()).Index(_dafny.IntOfInt64(2)).(_176_Common____GenericMarshalling__i_Compile.V))}}
}
func (_this *CompanionStruct_Default___) Parse__Ballot(val _176_Common____GenericMarshalling__i_Compile.V) _214_LiveRSL____CTypes__i_Compile.CBallot {
  return _214_LiveRSL____CTypes__i_Compile.CBallot{_214_LiveRSL____CTypes__i_Compile.CBallot_CBallot{(((val).Dtor_t()).Index(uint64(0)).(_176_Common____GenericMarshalling__i_Compile.V)).Dtor_u(), (((val).Dtor_t()).Index(uint64(1)).(_176_Common____GenericMarshalling__i_Compile.V)).Dtor_u()}}
}
func (_this *CompanionStruct_Default___) Parse__OperationNumber(val _176_Common____GenericMarshalling__i_Compile.V) _214_LiveRSL____CTypes__i_Compile.COperationNumber {
  return _214_LiveRSL____CTypes__i_Compile.COperationNumber{_214_LiveRSL____CTypes__i_Compile.COperationNumber_COperationNumber{(val).Dtor_u()}}
}
func (_this *CompanionStruct_Default___) Parse__Vote(val _176_Common____GenericMarshalling__i_Compile.V) _214_LiveRSL____CTypes__i_Compile.CVote {
  goto TAIL_CALL_START
TAIL_CALL_START:
var vote _214_LiveRSL____CTypes__i_Compile.CVote = _214_LiveRSL____CTypes__i_Compile.Type_CVote_().Default().(_214_LiveRSL____CTypes__i_Compile.CVote)
  var _ = vote
  var _4644_batch _dafny.Seq
  var _ = _4644_batch
var _out61 _dafny.Seq
  var _ = _out61
_out61 = Companion_Default___.Parse__RequestBatch(((val).Dtor_t()).Index(uint64(1)).(_176_Common____GenericMarshalling__i_Compile.V))
_4644_batch = _out61
  vote = _214_LiveRSL____CTypes__i_Compile.CVote{_214_LiveRSL____CTypes__i_Compile.CVote_CVote{Companion_Default___.Parse__Ballot(((val).Dtor_t()).Index(uint64(0)).(_176_Common____GenericMarshalling__i_Compile.V)), _4644_batch}}
  return vote
}
func (_this *CompanionStruct_Default___) Parse__ReplyCache(val _176_Common____GenericMarshalling__i_Compile.V) _dafny.Map {
  if ((((val).Dtor_a()).Cardinality()).Cmp(_dafny.Zero) == 0) {
    return _dafny.NewMapBuilder().ToMap()
  } else  {
    var _4645_tuple _176_Common____GenericMarshalling__i_Compile.V = ((val).Dtor_a()).Index(_dafny.Zero).(_176_Common____GenericMarshalling__i_Compile.V)
    var _ = _4645_tuple
var _4646_e _9_Native____Io__s_Compile.EndPoint = Companion_Default___.Parse__EndPoint(((_4645_tuple).Dtor_t()).Index(_dafny.Zero).(_176_Common____GenericMarshalling__i_Compile.V))
    var _ = _4646_e
var _4647_reply _214_LiveRSL____CTypes__i_Compile.CReply = Companion_Default___.Parse__Reply(((_4645_tuple).Dtor_t()).Index(_dafny.IntOfInt64(1)).(_176_Common____GenericMarshalling__i_Compile.V))
    var _ = _4647_reply
var _4648_others _dafny.Map = Companion_Default___.Parse__ReplyCache(_176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VArray{((val).Dtor_a()).Subseq(_dafny.IntOfInt64(1), _dafny.NilInt)}})
    var _ = _4648_others
var _4649_m _dafny.Map = (_4648_others).Update(_4646_e, _4647_reply)
    var _ = _4649_m
return _4649_m
  }
}
func (_this *CompanionStruct_Default___) Parse__Votes(val _176_Common____GenericMarshalling__i_Compile.V) _214_LiveRSL____CTypes__i_Compile.CVotes {
  var votes _214_LiveRSL____CTypes__i_Compile.CVotes = _214_LiveRSL____CTypes__i_Compile.Type_CVotes_().Default().(_214_LiveRSL____CTypes__i_Compile.CVotes)
  var _ = votes
  if ((uint64(((val).Dtor_a()).CardinalityInt())) == (uint64(0))) {
    votes = _214_LiveRSL____CTypes__i_Compile.CVotes{_214_LiveRSL____CTypes__i_Compile.CVotes_CVotes{_dafny.NewMapBuilder().ToMap()}}
  } else {
    var _4650_tuple _176_Common____GenericMarshalling__i_Compile.V
    var _ = _4650_tuple
    _4650_tuple = ((val).Dtor_a()).Index(uint64(0)).(_176_Common____GenericMarshalling__i_Compile.V)
    { }
    { }
    { }
    { }
    var _4651_op _214_LiveRSL____CTypes__i_Compile.COperationNumber
    var _ = _4651_op
    _4651_op = Companion_Default___.Parse__OperationNumber(((_4650_tuple).Dtor_t()).Index(uint64(0)).(_176_Common____GenericMarshalling__i_Compile.V))
    var _4652_vote _214_LiveRSL____CTypes__i_Compile.CVote
    var _ = _4652_vote
var _out62 _214_LiveRSL____CTypes__i_Compile.CVote
    var _ = _out62
_out62 = Companion_Default___.Parse__Vote(((_4650_tuple).Dtor_t()).Index(uint64(1)).(_176_Common____GenericMarshalling__i_Compile.V))
_4652_vote = _out62
    var _4653_others _214_LiveRSL____CTypes__i_Compile.CVotes
    var _ = _4653_others
var _out63 _214_LiveRSL____CTypes__i_Compile.CVotes
    var _ = _out63
_out63 = Companion_Default___.Parse__Votes(_176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VArray{((val).Dtor_a()).Subseq(uint64(1), _dafny.NilInt)}})
_4653_others = _out63
    var _4654_m _dafny.Map
    var _ = _4654_m
    _4654_m = ((_4653_others).Dtor_v()).Update(_4651_op, _4652_vote)
    votes = _214_LiveRSL____CTypes__i_Compile.CVotes{_214_LiveRSL____CTypes__i_Compile.CVotes_CVotes{_4654_m}}
  }
  return votes
}
func (_this *CompanionStruct_Default___) Parse__Message__Request(val _176_Common____GenericMarshalling__i_Compile.V) _217_LiveRSL____CMessage__i_Compile.CMessage {
  return _217_LiveRSL____CMessage__i_Compile.CMessage{_217_LiveRSL____CMessage__i_Compile.CMessage_CMessage__Request{(((val).Dtor_t()).Index(uint64(0)).(_176_Common____GenericMarshalling__i_Compile.V)).Dtor_u(), _197_LiveRSL____AppInterface__i_Compile.Companion_Default___.Parse__AppMessage(((val).Dtor_t()).Index(uint64(1)).(_176_Common____GenericMarshalling__i_Compile.V))}}
}
func (_this *CompanionStruct_Default___) Parse__Message__1a(val _176_Common____GenericMarshalling__i_Compile.V) _217_LiveRSL____CMessage__i_Compile.CMessage {
  return _217_LiveRSL____CMessage__i_Compile.CMessage{_217_LiveRSL____CMessage__i_Compile.CMessage_CMessage__1a{Companion_Default___.Parse__Ballot(val)}}
}
func (_this *CompanionStruct_Default___) Parse__Message__1b(val _176_Common____GenericMarshalling__i_Compile.V) _217_LiveRSL____CMessage__i_Compile.CMessage {
  goto TAIL_CALL_START
TAIL_CALL_START:
var msg _217_LiveRSL____CMessage__i_Compile.CMessage = _217_LiveRSL____CMessage__i_Compile.Type_CMessage_().Default().(_217_LiveRSL____CMessage__i_Compile.CMessage)
  var _ = msg
  var _4655_votes _214_LiveRSL____CTypes__i_Compile.CVotes
  var _ = _4655_votes
var _out64 _214_LiveRSL____CTypes__i_Compile.CVotes
  var _ = _out64
_out64 = Companion_Default___.Parse__Votes(((val).Dtor_t()).Index(uint64(2)).(_176_Common____GenericMarshalling__i_Compile.V))
_4655_votes = _out64
  msg = _217_LiveRSL____CMessage__i_Compile.CMessage{_217_LiveRSL____CMessage__i_Compile.CMessage_CMessage__1b{Companion_Default___.Parse__Ballot(((val).Dtor_t()).Index(uint64(0)).(_176_Common____GenericMarshalling__i_Compile.V)), Companion_Default___.Parse__OperationNumber(((val).Dtor_t()).Index(uint64(1)).(_176_Common____GenericMarshalling__i_Compile.V)), _4655_votes}}
  return msg
}
func (_this *CompanionStruct_Default___) Parse__Message__2a(val _176_Common____GenericMarshalling__i_Compile.V) _217_LiveRSL____CMessage__i_Compile.CMessage {
  goto TAIL_CALL_START
TAIL_CALL_START:
var msg _217_LiveRSL____CMessage__i_Compile.CMessage = _217_LiveRSL____CMessage__i_Compile.Type_CMessage_().Default().(_217_LiveRSL____CMessage__i_Compile.CMessage)
  var _ = msg
  var _4656_batch _dafny.Seq
  var _ = _4656_batch
var _out65 _dafny.Seq
  var _ = _out65
_out65 = Companion_Default___.Parse__RequestBatch(((val).Dtor_t()).Index(uint64(2)).(_176_Common____GenericMarshalling__i_Compile.V))
_4656_batch = _out65
  msg = _217_LiveRSL____CMessage__i_Compile.CMessage{_217_LiveRSL____CMessage__i_Compile.CMessage_CMessage__2a{Companion_Default___.Parse__Ballot(((val).Dtor_t()).Index(uint64(0)).(_176_Common____GenericMarshalling__i_Compile.V)), Companion_Default___.Parse__OperationNumber(((val).Dtor_t()).Index(uint64(1)).(_176_Common____GenericMarshalling__i_Compile.V)), _4656_batch}}
  return msg
}
func (_this *CompanionStruct_Default___) Parse__Message__2b(val _176_Common____GenericMarshalling__i_Compile.V) _217_LiveRSL____CMessage__i_Compile.CMessage {
  goto TAIL_CALL_START
TAIL_CALL_START:
var msg _217_LiveRSL____CMessage__i_Compile.CMessage = _217_LiveRSL____CMessage__i_Compile.Type_CMessage_().Default().(_217_LiveRSL____CMessage__i_Compile.CMessage)
  var _ = msg
  var _4657_batch _dafny.Seq
  var _ = _4657_batch
var _out66 _dafny.Seq
  var _ = _out66
_out66 = Companion_Default___.Parse__RequestBatch(((val).Dtor_t()).Index(uint64(2)).(_176_Common____GenericMarshalling__i_Compile.V))
_4657_batch = _out66
  msg = _217_LiveRSL____CMessage__i_Compile.CMessage{_217_LiveRSL____CMessage__i_Compile.CMessage_CMessage__2b{Companion_Default___.Parse__Ballot(((val).Dtor_t()).Index(uint64(0)).(_176_Common____GenericMarshalling__i_Compile.V)), Companion_Default___.Parse__OperationNumber(((val).Dtor_t()).Index(uint64(1)).(_176_Common____GenericMarshalling__i_Compile.V)), _4657_batch}}
  return msg
}
func (_this *CompanionStruct_Default___) Parse__Message__Heartbeat(val _176_Common____GenericMarshalling__i_Compile.V) _217_LiveRSL____CMessage__i_Compile.CMessage {
  return _217_LiveRSL____CMessage__i_Compile.CMessage{_217_LiveRSL____CMessage__i_Compile.CMessage_CMessage__Heartbeat{Companion_Default___.Parse__Ballot(((val).Dtor_t()).Index(uint64(0)).(_176_Common____GenericMarshalling__i_Compile.V)), ((((val).Dtor_t()).Index(uint64(1)).(_176_Common____GenericMarshalling__i_Compile.V)).Dtor_u()) != (uint64(0))/* dircomp */, Companion_Default___.Parse__OperationNumber(((val).Dtor_t()).Index(uint64(2)).(_176_Common____GenericMarshalling__i_Compile.V))}}
}
func (_this *CompanionStruct_Default___) Parse__Message__Reply(val _176_Common____GenericMarshalling__i_Compile.V) _217_LiveRSL____CMessage__i_Compile.CMessage {
  return _217_LiveRSL____CMessage__i_Compile.CMessage{_217_LiveRSL____CMessage__i_Compile.CMessage_CMessage__Reply{(((val).Dtor_t()).Index(_dafny.Zero).(_176_Common____GenericMarshalling__i_Compile.V)).Dtor_u(), _197_LiveRSL____AppInterface__i_Compile.Companion_Default___.Parse__AppMessage(((val).Dtor_t()).Index(_dafny.IntOfInt64(1)).(_176_Common____GenericMarshalling__i_Compile.V))}}
}
func (_this *CompanionStruct_Default___) Parse__Message__AppStateRequest(val _176_Common____GenericMarshalling__i_Compile.V) _217_LiveRSL____CMessage__i_Compile.CMessage {
  return _217_LiveRSL____CMessage__i_Compile.CMessage{_217_LiveRSL____CMessage__i_Compile.CMessage_CMessage__AppStateRequest{Companion_Default___.Parse__Ballot(((val).Dtor_t()).Index(uint64(0)).(_176_Common____GenericMarshalling__i_Compile.V)), Companion_Default___.Parse__OperationNumber(((val).Dtor_t()).Index(uint64(1)).(_176_Common____GenericMarshalling__i_Compile.V))}}
}
func (_this *CompanionStruct_Default___) Parse__Message__AppStateSupply(val _176_Common____GenericMarshalling__i_Compile.V) _217_LiveRSL____CMessage__i_Compile.CMessage {
  return _217_LiveRSL____CMessage__i_Compile.CMessage{_217_LiveRSL____CMessage__i_Compile.CMessage_CMessage__AppStateSupply{Companion_Default___.Parse__Ballot(((val).Dtor_t()).Index(uint64(0)).(_176_Common____GenericMarshalling__i_Compile.V)), Companion_Default___.Parse__OperationNumber(((val).Dtor_t()).Index(uint64(1)).(_176_Common____GenericMarshalling__i_Compile.V)), _197_LiveRSL____AppInterface__i_Compile.Companion_Default___.Parse__AppState(((val).Dtor_t()).Index(uint64(2)).(_176_Common____GenericMarshalling__i_Compile.V)), Companion_Default___.Parse__ReplyCache(((val).Dtor_t()).Index(uint64(3)).(_176_Common____GenericMarshalling__i_Compile.V))}}
}
func (_this *CompanionStruct_Default___) Parse__Message__StartingPhase2(val _176_Common____GenericMarshalling__i_Compile.V) _217_LiveRSL____CMessage__i_Compile.CMessage {
  return _217_LiveRSL____CMessage__i_Compile.CMessage{_217_LiveRSL____CMessage__i_Compile.CMessage_CMessage__StartingPhase2{Companion_Default___.Parse__Ballot(((val).Dtor_t()).Index(uint64(0)).(_176_Common____GenericMarshalling__i_Compile.V)), Companion_Default___.Parse__OperationNumber(((val).Dtor_t()).Index(uint64(1)).(_176_Common____GenericMarshalling__i_Compile.V))}}
}
func (_this *CompanionStruct_Default___) Parse__Message(val _176_Common____GenericMarshalling__i_Compile.V) _217_LiveRSL____CMessage__i_Compile.CMessage {
  goto TAIL_CALL_START
TAIL_CALL_START:
var msg _217_LiveRSL____CMessage__i_Compile.CMessage = _217_LiveRSL____CMessage__i_Compile.Type_CMessage_().Default().(_217_LiveRSL____CMessage__i_Compile.CMessage)
  var _ = msg
  if (((val).Dtor_c()) == (uint64(0))) {
    msg = Companion_Default___.Parse__Message__Request((val).Dtor_val())
  } else if (((val).Dtor_c()) == (uint64(1))) {
    msg = Companion_Default___.Parse__Message__1a((val).Dtor_val())
  } else if (((val).Dtor_c()) == (uint64(2))) {
    var _out67 _217_LiveRSL____CMessage__i_Compile.CMessage
    var _ = _out67
_out67 = Companion_Default___.Parse__Message__1b((val).Dtor_val())
msg = _out67
  } else if (((val).Dtor_c()) == (uint64(3))) {
    var _out68 _217_LiveRSL____CMessage__i_Compile.CMessage
    var _ = _out68
_out68 = Companion_Default___.Parse__Message__2a((val).Dtor_val())
msg = _out68
  } else if (((val).Dtor_c()) == (uint64(4))) {
    var _out69 _217_LiveRSL____CMessage__i_Compile.CMessage
    var _ = _out69
_out69 = Companion_Default___.Parse__Message__2b((val).Dtor_val())
msg = _out69
  } else if (((val).Dtor_c()) == (uint64(5))) {
    msg = Companion_Default___.Parse__Message__Heartbeat((val).Dtor_val())
  } else if (((val).Dtor_c()) == (uint64(6))) {
    msg = Companion_Default___.Parse__Message__Reply((val).Dtor_val())
  } else if (((val).Dtor_c()) == (uint64(7))) {
    msg = Companion_Default___.Parse__Message__AppStateRequest((val).Dtor_val())
  } else if (((val).Dtor_c()) == (uint64(8))) {
    msg = Companion_Default___.Parse__Message__AppStateSupply((val).Dtor_val())
  } else if (((val).Dtor_c()) == (uint64(9))) {
    msg = Companion_Default___.Parse__Message__StartingPhase2((val).Dtor_val())
  } else {
    { }
    msg = Companion_Default___.Parse__Message__Request(val)
  }
  return msg
}
func (_this *CompanionStruct_Default___) PaxosDemarshallDataMethod(data *_dafny.Array, msg__grammar _176_Common____GenericMarshalling__i_Compile.G) _217_LiveRSL____CMessage__i_Compile.CMessage {
  goto TAIL_CALL_START
TAIL_CALL_START:
var msg _217_LiveRSL____CMessage__i_Compile.CMessage = _217_LiveRSL____CMessage__i_Compile.Type_CMessage_().Default().(_217_LiveRSL____CMessage__i_Compile.CMessage)
  var _ = msg
  var _4658_success bool
  var _ = _4658_success
var _4659_val _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4659_val
var _out70 bool
  var _ = _out70
var _out71 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out71
_out70,_out71 = _176_Common____GenericMarshalling__i_Compile.Companion_Default___.Demarshall(data, msg__grammar)
_4658_success = _out70
_4659_val = _out71
  if (_4658_success) {
    { }
    var _out72 _217_LiveRSL____CMessage__i_Compile.CMessage
    var _ = _out72
_out72 = Companion_Default___.Parse__Message(_4659_val)
msg = _out72
    { }
  } else {
    msg = _217_LiveRSL____CMessage__i_Compile.CMessage{_217_LiveRSL____CMessage__i_Compile.CMessage_CMessage__Invalid{}}
  }
  return msg
}
func (_this *CompanionStruct_Default___) DetermineIfValidVote(vote _214_LiveRSL____CTypes__i_Compile.CVote) bool {
  goto TAIL_CALL_START
TAIL_CALL_START:
var b bool = false
  var _ = b
  b = (uint64(((vote).Dtor_max__val()).CardinalityInt())) <= (uint64(100))
  return b
}
func (_this *CompanionStruct_Default___) DetermineIfValidVotes(votes _214_LiveRSL____CTypes__i_Compile.CVotes) bool {
  goto TAIL_CALL_START
TAIL_CALL_START:
var b bool = false
  var _ = b
  b = (uint64(((votes).Dtor_v()).CardinalityInt())) < (uint64(8))
  if (!(b)) {
    return b
  }
  var _4660_keys _dafny.Set
  var _ = _4660_keys
  _4660_keys = _118_Collections____Maps__i_Compile.Companion_Default___.Domain((votes).Dtor_v())
  { }
  for (uint64((_4660_keys).CardinalityInt())) > (uint64(0)) {
    var _4661_opn _214_LiveRSL____CTypes__i_Compile.COperationNumber
    var _ = _4661_opn
    for _iter4 := _dafny.Iterate((_4660_keys).Elements());; {
      _val4, _ok4 := _iter4()
if !_ok4 { break }
_assign_such_that_1 := _val4.(_214_LiveRSL____CTypes__i_Compile.COperationNumber)
_4661_opn = _assign_such_that_1
if ((_4660_keys).Contains(_4661_opn)) {
        goto L_ASSIGN_SUCH_THAT_1
      }
    }
    panic("assign-such-that search produced no value (line 624)")
  L_ASSIGN_SUCH_THAT_1:
    _4660_keys = (_4660_keys).Difference(_dafny.SetOf(_4661_opn))
    var _out73 bool
    var _ = _out73
_out73 = Companion_Default___.DetermineIfValidVote(((votes).Dtor_v()).Get(_4661_opn).(_214_LiveRSL____CTypes__i_Compile.CVote))
b = _out73
    if (!(b)) {
      return b
    }
  }
  return b
}
func (_this *CompanionStruct_Default___) DetermineIfValidReplyCache(m _dafny.Map) bool {
  goto TAIL_CALL_START
TAIL_CALL_START:
var b bool = false
  var _ = b
  b = (uint64((m).CardinalityInt())) < (uint64(256))
  { }
  { }
  return b
}
func (_this *CompanionStruct_Default___) DetermineIfMessageMarshallable(msg _217_LiveRSL____CMessage__i_Compile.CMessage) bool {
  var b bool = false
  var _ = b
  if ((msg).Is_CMessage__Invalid()) {
    b = false
  } else if ((msg).Is_CMessage__Request()) {
    b = _197_LiveRSL____AppInterface__i_Compile.Companion_Default___.ValidAppMessage((msg).Dtor_val())
  } else if ((msg).Is_CMessage__1a()) {
    b = true
  } else if ((msg).Is_CMessage__1b()) {
    var _out74 bool
    var _ = _out74
_out74 = Companion_Default___.DetermineIfValidVotes((msg).Dtor_votes())
b = _out74
  } else if ((msg).Is_CMessage__2a()) {
    b = (uint64(((msg).Dtor_val__2a()).CardinalityInt())) <= (uint64(100))
  } else if ((msg).Is_CMessage__2b()) {
    b = (uint64(((msg).Dtor_val__2b()).CardinalityInt())) <= (uint64(100))
  } else if ((msg).Is_CMessage__Heartbeat()) {
    b = true
  } else if ((msg).Is_CMessage__Reply()) {
    b = _197_LiveRSL____AppInterface__i_Compile.Companion_Default___.ValidAppMessage((msg).Dtor_reply())
  } else if ((msg).Is_CMessage__AppStateRequest()) {
    b = true
  } else if ((msg).Is_CMessage__AppStateSupply()) {
    var _4662_b1 bool
    var _ = _4662_b1
    _4662_b1 = _197_LiveRSL____AppInterface__i_Compile.Companion_Default___.AppStateMarshallable((msg).Dtor_app__state())
    var _4663_b2 bool
    var _ = _4663_b2
var _out75 bool
    var _ = _out75
_out75 = Companion_Default___.DetermineIfValidReplyCache((msg).Dtor_reply__cache())
_4663_b2 = _out75
    b = (_4662_b1) && (_4663_b2)
  } else if ((msg).Is_CMessage__StartingPhase2()) {
    b = true
  } else { }
  return b
}
func (_this *CompanionStruct_Default___) MarshallEndPoint(c _9_Native____Io__s_Compile.EndPoint) _176_Common____GenericMarshalling__i_Compile.V {
  goto TAIL_CALL_START
TAIL_CALL_START:
var val _176_Common____GenericMarshalling__i_Compile.V = _176_Common____GenericMarshalling__i_Compile.Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = val
  val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VUint64{_194_Common____NodeIdentity__i_Compile.Companion_Default___.ConvertEndPointToUint64(c)}}
  { }
  return val
}
func (_this *CompanionStruct_Default___) MarshallRequest(c _214_LiveRSL____CTypes__i_Compile.CRequest) _176_Common____GenericMarshalling__i_Compile.V {
  goto TAIL_CALL_START
TAIL_CALL_START:
var val _176_Common____GenericMarshalling__i_Compile.V = _176_Common____GenericMarshalling__i_Compile.Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = val
  var _4664_marshalled__app__message _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4664_marshalled__app__message
var _out76 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out76
_out76 = _197_LiveRSL____AppInterface__i_Compile.Companion_Default___.MarshallCAppMessage((c).Dtor_request())
_4664_marshalled__app__message = _out76
  var _4665_marshalled__ep _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4665_marshalled__ep
var _out77 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out77
_out77 = Companion_Default___.MarshallEndPoint((c).Dtor_client())
_4665_marshalled__ep = _out77
  val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VTuple{_dafny.SeqOf(_4665_marshalled__ep, _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VUint64{(c).Dtor_seqno()}}, _4664_marshalled__app__message)}}
  { }
  { }
  { }
  { }
  return val
}
func (_this *CompanionStruct_Default___) MarshallRequestBatch(c _dafny.Seq) _176_Common____GenericMarshalling__i_Compile.V {
  goto TAIL_CALL_START
TAIL_CALL_START:
var val _176_Common____GenericMarshalling__i_Compile.V = _176_Common____GenericMarshalling__i_Compile.Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = val
  var _4666_reqs *_dafny.Array
  var _ = _4666_reqs
  var _nw9 = _dafny.NewArrayWithValue(_176_Common____GenericMarshalling__i_Compile.Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V), uint64((c).CardinalityInt()))
  var _ = _nw9
  _4666_reqs = _nw9
  var _4667_i uint64
  var _ = _4667_i
  _4667_i = uint64(0)
  for (_4667_i) < (uint64((c).CardinalityInt())) {
    var _4668_single _176_Common____GenericMarshalling__i_Compile.V
    var _ = _4668_single
var _out78 _176_Common____GenericMarshalling__i_Compile.V
    var _ = _out78
_out78 = Companion_Default___.MarshallRequest((c).Index(_4667_i).(_214_LiveRSL____CTypes__i_Compile.CRequest))
_4668_single = _out78
    { }
    *((_4666_reqs).Index(_dafny.IntOfAny((_4667_i)))) = _4668_single
    _4667_i = (_4667_i) + (uint64(1))
  }
  val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VArray{(_4666_reqs).RangeToSeq(_dafny.NilInt, _dafny.NilInt)}}
  return val
}
func (_this *CompanionStruct_Default___) MarshallReply(c _214_LiveRSL____CTypes__i_Compile.CReply) _176_Common____GenericMarshalling__i_Compile.V {
  goto TAIL_CALL_START
TAIL_CALL_START:
var val _176_Common____GenericMarshalling__i_Compile.V = _176_Common____GenericMarshalling__i_Compile.Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = val
  var _4669_marshalled__app__message _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4669_marshalled__app__message
var _out79 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out79
_out79 = _197_LiveRSL____AppInterface__i_Compile.Companion_Default___.MarshallCAppMessage((c).Dtor_reply())
_4669_marshalled__app__message = _out79
  var _4670_marshalled__ep _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4670_marshalled__ep
var _out80 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out80
_out80 = Companion_Default___.MarshallEndPoint((c).Dtor_client())
_4670_marshalled__ep = _out80
  val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VTuple{_dafny.SeqOf(_4670_marshalled__ep, _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VUint64{(c).Dtor_seqno()}}, _4669_marshalled__app__message)}}
  { }
  { }
  { }
  { }
  return val
}
func (_this *CompanionStruct_Default___) MarshallBallot(c _214_LiveRSL____CTypes__i_Compile.CBallot) _176_Common____GenericMarshalling__i_Compile.V {
  goto TAIL_CALL_START
TAIL_CALL_START:
var val _176_Common____GenericMarshalling__i_Compile.V = _176_Common____GenericMarshalling__i_Compile.Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = val
  val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VTuple{_dafny.SeqOf(_176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VUint64{(c).Dtor_seqno()}}, _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VUint64{(c).Dtor_proposer__id()}})}}
  return val
}
func (_this *CompanionStruct_Default___) MarshallOperationNumber(c _214_LiveRSL____CTypes__i_Compile.COperationNumber) _176_Common____GenericMarshalling__i_Compile.V {
  goto TAIL_CALL_START
TAIL_CALL_START:
var val _176_Common____GenericMarshalling__i_Compile.V = _176_Common____GenericMarshalling__i_Compile.Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = val
  val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VUint64{(c).Dtor_n()}}
  return val
}
func (_this *CompanionStruct_Default___) MarshallVote(c _214_LiveRSL____CTypes__i_Compile.CVote) _176_Common____GenericMarshalling__i_Compile.V {
  goto TAIL_CALL_START
TAIL_CALL_START:
var val _176_Common____GenericMarshalling__i_Compile.V = _176_Common____GenericMarshalling__i_Compile.Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = val
  var _4671_bal _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4671_bal
var _out81 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out81
_out81 = Companion_Default___.MarshallBallot((c).Dtor_max__value__bal())
_4671_bal = _out81
  var _4672_v _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4672_v
var _out82 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out82
_out82 = Companion_Default___.MarshallRequestBatch((c).Dtor_max__val())
_4672_v = _out82
  val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VTuple{_dafny.SeqOf(_4671_bal, _4672_v)}}
  return val
}
func (_this *CompanionStruct_Default___) MarshallReplyCache(c _dafny.Map) _176_Common____GenericMarshalling__i_Compile.V {
  var val _176_Common____GenericMarshalling__i_Compile.V = _176_Common____GenericMarshalling__i_Compile.Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = val
  if ((uint64((c).CardinalityInt())) == (uint64(0))) {
    val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VArray{_dafny.SeqOf()}}
    { }
  } else {
    var _4673_ep _9_Native____Io__s_Compile.EndPoint
    var _ = _4673_ep
    for _iter5 := _dafny.Iterate((c).Keys().Elements());; {
      _val5, _ok5 := _iter5()
if !_ok5 { break }
_assign_such_that_2 := _val5.(_9_Native____Io__s_Compile.EndPoint)
_4673_ep = _assign_such_that_2
if ((c).Contains(_4673_ep)) {
        goto L_ASSIGN_SUCH_THAT_2
      }
    }
    panic("assign-such-that search produced no value (line 810)")
  L_ASSIGN_SUCH_THAT_2:
    var _4674_marshalled__ep _176_Common____GenericMarshalling__i_Compile.V
    var _ = _4674_marshalled__ep
var _out83 _176_Common____GenericMarshalling__i_Compile.V
    var _ = _out83
_out83 = Companion_Default___.MarshallEndPoint(_4673_ep)
_4674_marshalled__ep = _out83
    var _4675_marshalled__reply _176_Common____GenericMarshalling__i_Compile.V
    var _ = _4675_marshalled__reply
var _out84 _176_Common____GenericMarshalling__i_Compile.V
    var _ = _out84
_out84 = Companion_Default___.MarshallReply((c).Get(_4673_ep).(_214_LiveRSL____CTypes__i_Compile.CReply))
_4675_marshalled__reply = _out84
    var _4676_remainder _dafny.Map
    var _ = _4676_remainder
    _4676_remainder = _118_Collections____Maps__i_Compile.Companion_Default___.RemoveElt(c, _4673_ep)
    { }
    var _4677_marshalled__remainder _176_Common____GenericMarshalling__i_Compile.V
    var _ = _4677_marshalled__remainder
var _out85 _176_Common____GenericMarshalling__i_Compile.V
    var _ = _out85
_out85 = Companion_Default___.MarshallReplyCache(_4676_remainder)
_4677_marshalled__remainder = _out85
    { }
    val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VArray{(_dafny.SeqOf(_176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VTuple{_dafny.SeqOf(_4674_marshalled__ep, _4675_marshalled__reply)}})).Concat((_4677_marshalled__remainder).Dtor_a())}}
    { }
    { }
    { }
    { }
    { }
    { }
    { }
    { }
    { }
    { }
    { }
    { }
    { }
    { }
  }
  return val
}
func (_this *CompanionStruct_Default___) MarshallVotes(c _214_LiveRSL____CTypes__i_Compile.CVotes) _176_Common____GenericMarshalling__i_Compile.V {
  var val _176_Common____GenericMarshalling__i_Compile.V = _176_Common____GenericMarshalling__i_Compile.Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = val
  if ((uint64(((c).Dtor_v()).CardinalityInt())) == (uint64(0))) {
    val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VArray{_dafny.SeqOf()}}
    { }
  } else {
    { }
    var _4678_op _214_LiveRSL____CTypes__i_Compile.COperationNumber
    var _ = _4678_op
    for _iter6 := _dafny.Iterate(((c).Dtor_v()).Keys().Elements());; {
      _val6, _ok6 := _iter6()
if !_ok6 { break }
_assign_such_that_3 := _val6.(_214_LiveRSL____CTypes__i_Compile.COperationNumber)
_4678_op = _assign_such_that_3
if (((c).Dtor_v()).Contains(_4678_op)) {
        goto L_ASSIGN_SUCH_THAT_3
      }
    }
    panic("assign-such-that search produced no value (line 869)")
  L_ASSIGN_SUCH_THAT_3:
    var _4679_marshalled__op _176_Common____GenericMarshalling__i_Compile.V
    var _ = _4679_marshalled__op
var _out86 _176_Common____GenericMarshalling__i_Compile.V
    var _ = _out86
_out86 = Companion_Default___.MarshallOperationNumber(_4678_op)
_4679_marshalled__op = _out86
    var _4680_marshalled__vote _176_Common____GenericMarshalling__i_Compile.V
    var _ = _4680_marshalled__vote
var _out87 _176_Common____GenericMarshalling__i_Compile.V
    var _ = _out87
_out87 = Companion_Default___.MarshallVote(((c).Dtor_v()).Get(_4678_op).(_214_LiveRSL____CTypes__i_Compile.CVote))
_4680_marshalled__vote = _out87
    var _4681_remainder _dafny.Map
    var _ = _4681_remainder
    _4681_remainder = _118_Collections____Maps__i_Compile.Companion_Default___.RemoveElt((c).Dtor_v(), _4678_op)
    var _4682_marshalled__remainder _176_Common____GenericMarshalling__i_Compile.V
    var _ = _4682_marshalled__remainder
var _out88 _176_Common____GenericMarshalling__i_Compile.V
    var _ = _out88
_out88 = Companion_Default___.MarshallVotes(_214_LiveRSL____CTypes__i_Compile.CVotes{_214_LiveRSL____CTypes__i_Compile.CVotes_CVotes{_4681_remainder}})
_4682_marshalled__remainder = _out88
    { }
    val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VArray{(_dafny.SeqOf(_176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VTuple{_dafny.SeqOf(_4679_marshalled__op, _4680_marshalled__vote)}})).Concat((_4682_marshalled__remainder).Dtor_a())}}
    { }
    { }
    { }
    { }
    { }
    { }
    { }
    { }
    { }
    { }
    { }
    { }
    { }
    { }
  }
  return val
}
func (_this *CompanionStruct_Default___) MarshallMessage__Request(c _217_LiveRSL____CMessage__i_Compile.CMessage) _176_Common____GenericMarshalling__i_Compile.V {
  goto TAIL_CALL_START
TAIL_CALL_START:
var val _176_Common____GenericMarshalling__i_Compile.V = _176_Common____GenericMarshalling__i_Compile.Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = val
  var _4683_v _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4683_v
var _out89 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out89
_out89 = _197_LiveRSL____AppInterface__i_Compile.Companion_Default___.MarshallCAppMessage((c).Dtor_val())
_4683_v = _out89
  val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VTuple{_dafny.SeqOf(_176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VUint64{(c).Dtor_seqno()}}, _4683_v)}}
  return val
}
func (_this *CompanionStruct_Default___) MarshallMessage__1a(c _217_LiveRSL____CMessage__i_Compile.CMessage) _176_Common____GenericMarshalling__i_Compile.V {
  goto TAIL_CALL_START
TAIL_CALL_START:
var val _176_Common____GenericMarshalling__i_Compile.V = _176_Common____GenericMarshalling__i_Compile.Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = val
  var _out90 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out90
_out90 = Companion_Default___.MarshallBallot((c).Dtor_bal__1a())
val = _out90
  return val
}
func (_this *CompanionStruct_Default___) MarshallMessage__1b(c _217_LiveRSL____CMessage__i_Compile.CMessage) _176_Common____GenericMarshalling__i_Compile.V {
  goto TAIL_CALL_START
TAIL_CALL_START:
var val _176_Common____GenericMarshalling__i_Compile.V = _176_Common____GenericMarshalling__i_Compile.Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = val
  var _4684_bal _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4684_bal
var _out91 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out91
_out91 = Companion_Default___.MarshallBallot((c).Dtor_bal__1b())
_4684_bal = _out91
  var _4685_log__truncation__point _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4685_log__truncation__point
var _out92 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out92
_out92 = Companion_Default___.MarshallOperationNumber((c).Dtor_log__truncation__point())
_4685_log__truncation__point = _out92
  var _4686_votes _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4686_votes
var _out93 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out93
_out93 = Companion_Default___.MarshallVotes((c).Dtor_votes())
_4686_votes = _out93
  val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VTuple{_dafny.SeqOf(_4684_bal, _4685_log__truncation__point, _4686_votes)}}
  { }
  { }
  { }
  { }
  { }
  { }
  return val
}
func (_this *CompanionStruct_Default___) MarshallMessage__2a(c _217_LiveRSL____CMessage__i_Compile.CMessage) _176_Common____GenericMarshalling__i_Compile.V {
  goto TAIL_CALL_START
TAIL_CALL_START:
var val _176_Common____GenericMarshalling__i_Compile.V = _176_Common____GenericMarshalling__i_Compile.Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = val
  var _4687_bal _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4687_bal
var _out94 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out94
_out94 = Companion_Default___.MarshallBallot((c).Dtor_bal__2a())
_4687_bal = _out94
  var _4688_op _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4688_op
var _out95 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out95
_out95 = Companion_Default___.MarshallOperationNumber((c).Dtor_opn__2a())
_4688_op = _out95
  var _4689_v _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4689_v
var _out96 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out96
_out96 = Companion_Default___.MarshallRequestBatch((c).Dtor_val__2a())
_4689_v = _out96
  val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VTuple{_dafny.SeqOf(_4687_bal, _4688_op, _4689_v)}}
  { }
  { }
  { }
  { }
  { }
  { }
  return val
}
func (_this *CompanionStruct_Default___) MarshallMessage__2b(c _217_LiveRSL____CMessage__i_Compile.CMessage) _176_Common____GenericMarshalling__i_Compile.V {
  goto TAIL_CALL_START
TAIL_CALL_START:
var val _176_Common____GenericMarshalling__i_Compile.V = _176_Common____GenericMarshalling__i_Compile.Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = val
  var _4690_bal _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4690_bal
var _out97 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out97
_out97 = Companion_Default___.MarshallBallot((c).Dtor_bal__2b())
_4690_bal = _out97
  var _4691_op _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4691_op
var _out98 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out98
_out98 = Companion_Default___.MarshallOperationNumber((c).Dtor_opn__2b())
_4691_op = _out98
  var _4692_v _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4692_v
var _out99 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out99
_out99 = Companion_Default___.MarshallRequestBatch((c).Dtor_val__2b())
_4692_v = _out99
  val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VTuple{_dafny.SeqOf(_4690_bal, _4691_op, _4692_v)}}
  { }
  { }
  { }
  { }
  return val
}
func (_this *CompanionStruct_Default___) MarshallMessage__Heartbeat(c _217_LiveRSL____CMessage__i_Compile.CMessage) _176_Common____GenericMarshalling__i_Compile.V {
  goto TAIL_CALL_START
TAIL_CALL_START:
var val _176_Common____GenericMarshalling__i_Compile.V = _176_Common____GenericMarshalling__i_Compile.Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = val
  var _4693_ballot _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4693_ballot
var _out100 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out100
_out100 = Companion_Default___.MarshallBallot((c).Dtor_bal__heartbeat())
_4693_ballot = _out100
  var _4694_op _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4694_op
var _out101 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out101
_out101 = Companion_Default___.MarshallOperationNumber((c).Dtor_opn__ckpt())
_4694_op = _out101
  val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VTuple{_dafny.SeqOf(_4693_ballot, _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VUint64{(func () uint64 { if (c).Dtor_suspicious() { return uint64(1) }; return uint64(0) })() }}, _4694_op)}}
  return val
}
func (_this *CompanionStruct_Default___) MarshallMessage__Reply(c _217_LiveRSL____CMessage__i_Compile.CMessage) _176_Common____GenericMarshalling__i_Compile.V {
  goto TAIL_CALL_START
TAIL_CALL_START:
var val _176_Common____GenericMarshalling__i_Compile.V = _176_Common____GenericMarshalling__i_Compile.Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = val
  var _4695_app__val _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4695_app__val
var _out102 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out102
_out102 = _197_LiveRSL____AppInterface__i_Compile.Companion_Default___.MarshallCAppMessage((c).Dtor_reply())
_4695_app__val = _out102
  val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VTuple{_dafny.SeqOf(_176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VUint64{(c).Dtor_seqno__reply()}}, _4695_app__val)}}
  return val
}
func (_this *CompanionStruct_Default___) MarshallMessage__AppStateRequest(c _217_LiveRSL____CMessage__i_Compile.CMessage) _176_Common____GenericMarshalling__i_Compile.V {
  goto TAIL_CALL_START
TAIL_CALL_START:
var val _176_Common____GenericMarshalling__i_Compile.V = _176_Common____GenericMarshalling__i_Compile.Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = val
  var _4696_ballot _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4696_ballot
var _out103 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out103
_out103 = Companion_Default___.MarshallBallot((c).Dtor_bal__state__req())
_4696_ballot = _out103
  var _4697_opn _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4697_opn
var _out104 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out104
_out104 = Companion_Default___.MarshallOperationNumber((c).Dtor_opn__state__req())
_4697_opn = _out104
  val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VTuple{_dafny.SeqOf(_4696_ballot, _4697_opn)}}
  return val
}
func (_this *CompanionStruct_Default___) MarshallMessage__AppStateSupply(c _217_LiveRSL____CMessage__i_Compile.CMessage) _176_Common____GenericMarshalling__i_Compile.V {
  goto TAIL_CALL_START
TAIL_CALL_START:
var val _176_Common____GenericMarshalling__i_Compile.V = _176_Common____GenericMarshalling__i_Compile.Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = val
  var _4698_ballot _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4698_ballot
var _out105 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out105
_out105 = Companion_Default___.MarshallBallot((c).Dtor_bal__state__supply())
_4698_ballot = _out105
  var _4699_opn__state__supply _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4699_opn__state__supply
var _out106 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out106
_out106 = Companion_Default___.MarshallOperationNumber((c).Dtor_opn__state__supply())
_4699_opn__state__supply = _out106
  var _4700_app__state _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4700_app__state
var _out107 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out107
_out107 = _197_LiveRSL____AppInterface__i_Compile.Companion_Default___.MarshallAppState((c).Dtor_app__state())
_4700_app__state = _out107
  var _4701_reply__cache _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4701_reply__cache
var _out108 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out108
_out108 = Companion_Default___.MarshallReplyCache((c).Dtor_reply__cache())
_4701_reply__cache = _out108
  val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VTuple{_dafny.SeqOf(_4698_ballot, _4699_opn__state__supply, _4700_app__state, _4701_reply__cache)}}
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  return val
}
func (_this *CompanionStruct_Default___) MarshallMessage__StartingPhase2(c _217_LiveRSL____CMessage__i_Compile.CMessage) _176_Common____GenericMarshalling__i_Compile.V {
  goto TAIL_CALL_START
TAIL_CALL_START:
var val _176_Common____GenericMarshalling__i_Compile.V = _176_Common____GenericMarshalling__i_Compile.Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = val
  var _4702_bal _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4702_bal
var _out109 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out109
_out109 = Companion_Default___.MarshallBallot((c).Dtor_bal__2())
_4702_bal = _out109
  var _4703_op _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4703_op
var _out110 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out110
_out110 = Companion_Default___.MarshallOperationNumber((c).Dtor_logTruncationPoint__2())
_4703_op = _out110
  val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VTuple{_dafny.SeqOf(_4702_bal, _4703_op)}}
  return val
}
func (_this *CompanionStruct_Default___) MarshallMessage(c _217_LiveRSL____CMessage__i_Compile.CMessage) _176_Common____GenericMarshalling__i_Compile.V {
  goto TAIL_CALL_START
TAIL_CALL_START:
var val _176_Common____GenericMarshalling__i_Compile.V = _176_Common____GenericMarshalling__i_Compile.Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = val
  var _4704_start__time uint64
  var _ = _4704_start__time
var _out111 uint64
  var _ = _out111
_out111 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_4704_start__time = _out111
  { }
  if ((c).Is_CMessage__Request()) {
    var _4705_msg _176_Common____GenericMarshalling__i_Compile.V
    var _ = _4705_msg
var _out112 _176_Common____GenericMarshalling__i_Compile.V
    var _ = _out112
_out112 = Companion_Default___.MarshallMessage__Request(c)
_4705_msg = _out112
    val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VCase{uint64(0), _4705_msg}}
    var _4706_end__time uint64
    var _ = _4706_end__time
var _out113 uint64
    var _ = _out113
_out113 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_4706_end__time = _out113
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("MarshallMessage_Request"), _4704_start__time, _4706_end__time)
  } else if ((c).Is_CMessage__1a()) {
    var _4707_msg _176_Common____GenericMarshalling__i_Compile.V
    var _ = _4707_msg
var _out114 _176_Common____GenericMarshalling__i_Compile.V
    var _ = _out114
_out114 = Companion_Default___.MarshallMessage__1a(c)
_4707_msg = _out114
    val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VCase{uint64(1), _4707_msg}}
    var _4708_end__time uint64
    var _ = _4708_end__time
var _out115 uint64
    var _ = _out115
_out115 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_4708_end__time = _out115
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("MarshallMessage_1a"), _4704_start__time, _4708_end__time)
  } else if ((c).Is_CMessage__1b()) {
    var _4709_msg _176_Common____GenericMarshalling__i_Compile.V
    var _ = _4709_msg
var _out116 _176_Common____GenericMarshalling__i_Compile.V
    var _ = _out116
_out116 = Companion_Default___.MarshallMessage__1b(c)
_4709_msg = _out116
    val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VCase{uint64(2), _4709_msg}}
    var _4710_end__time uint64
    var _ = _4710_end__time
var _out117 uint64
    var _ = _out117
_out117 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_4710_end__time = _out117
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("MarshallMessage_1b"), _4704_start__time, _4710_end__time)
  } else if ((c).Is_CMessage__2a()) {
    var _4711_msg _176_Common____GenericMarshalling__i_Compile.V
    var _ = _4711_msg
var _out118 _176_Common____GenericMarshalling__i_Compile.V
    var _ = _out118
_out118 = Companion_Default___.MarshallMessage__2a(c)
_4711_msg = _out118
    val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VCase{uint64(3), _4711_msg}}
    var _4712_end__time uint64
    var _ = _4712_end__time
var _out119 uint64
    var _ = _out119
_out119 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_4712_end__time = _out119
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("MarshallMessage_2a"), _4704_start__time, _4712_end__time)
  } else if ((c).Is_CMessage__2b()) {
    var _4713_msg _176_Common____GenericMarshalling__i_Compile.V
    var _ = _4713_msg
var _out120 _176_Common____GenericMarshalling__i_Compile.V
    var _ = _out120
_out120 = Companion_Default___.MarshallMessage__2b(c)
_4713_msg = _out120
    val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VCase{uint64(4), _4713_msg}}
    var _4714_end__time uint64
    var _ = _4714_end__time
var _out121 uint64
    var _ = _out121
_out121 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_4714_end__time = _out121
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("MarshallMessage_2b"), _4704_start__time, _4714_end__time)
  } else if ((c).Is_CMessage__Heartbeat()) {
    var _4715_msg _176_Common____GenericMarshalling__i_Compile.V
    var _ = _4715_msg
var _out122 _176_Common____GenericMarshalling__i_Compile.V
    var _ = _out122
_out122 = Companion_Default___.MarshallMessage__Heartbeat(c)
_4715_msg = _out122
    val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VCase{uint64(5), _4715_msg}}
    var _4716_end__time uint64
    var _ = _4716_end__time
var _out123 uint64
    var _ = _out123
_out123 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_4716_end__time = _out123
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("MarshallMessage_Heartbeat"), _4704_start__time, _4716_end__time)
  } else if ((c).Is_CMessage__Reply()) {
    var _4717_msg _176_Common____GenericMarshalling__i_Compile.V
    var _ = _4717_msg
var _out124 _176_Common____GenericMarshalling__i_Compile.V
    var _ = _out124
_out124 = Companion_Default___.MarshallMessage__Reply(c)
_4717_msg = _out124
    val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VCase{uint64(6), _4717_msg}}
    { }
    var _4718_end__time uint64
    var _ = _4718_end__time
var _out125 uint64
    var _ = _out125
_out125 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_4718_end__time = _out125
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("MarshallMessage_Reply"), _4704_start__time, _4718_end__time)
  } else if ((c).Is_CMessage__AppStateRequest()) {
    var _4719_msg _176_Common____GenericMarshalling__i_Compile.V
    var _ = _4719_msg
var _out126 _176_Common____GenericMarshalling__i_Compile.V
    var _ = _out126
_out126 = Companion_Default___.MarshallMessage__AppStateRequest(c)
_4719_msg = _out126
    val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VCase{uint64(7), _4719_msg}}
    var _4720_end__time uint64
    var _ = _4720_end__time
var _out127 uint64
    var _ = _out127
_out127 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_4720_end__time = _out127
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("MarshallMessage_AppStateRequest"), _4704_start__time, _4720_end__time)
  } else if ((c).Is_CMessage__AppStateSupply()) {
    var _4721_msg _176_Common____GenericMarshalling__i_Compile.V
    var _ = _4721_msg
var _out128 _176_Common____GenericMarshalling__i_Compile.V
    var _ = _out128
_out128 = Companion_Default___.MarshallMessage__AppStateSupply(c)
_4721_msg = _out128
    val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VCase{uint64(8), _4721_msg}}
    var _4722_end__time uint64
    var _ = _4722_end__time
var _out129 uint64
    var _ = _out129
_out129 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_4722_end__time = _out129
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("MarshallMessage_AppStateSupply"), _4704_start__time, _4722_end__time)
  } else if ((c).Is_CMessage__StartingPhase2()) {
    var _4723_msg _176_Common____GenericMarshalling__i_Compile.V
    var _ = _4723_msg
var _out130 _176_Common____GenericMarshalling__i_Compile.V
    var _ = _out130
_out130 = Companion_Default___.MarshallMessage__StartingPhase2(c)
_4723_msg = _out130
    val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VCase{uint64(9), _4723_msg}}
    var _4724_end__time uint64
    var _ = _4724_end__time
var _out131 uint64
    var _ = _out131
_out131 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_4724_end__time = _out131
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("MarshallMessage_StartingPhase2"), _4704_start__time, _4724_end__time)
  }
  return val
}
func (_this *CompanionStruct_Default___) PaxosMarshall(msg _217_LiveRSL____CMessage__i_Compile.CMessage) *_dafny.Array {
  goto TAIL_CALL_START
TAIL_CALL_START:
var data *_dafny.Array = _dafny.NewArrayWithValue(0, _dafny.IntOf(0))
  var _ = data
  var _4725_marshall__start__time uint64
  var _ = _4725_marshall__start__time
var _out132 uint64
  var _ = _out132
_out132 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_4725_marshall__start__time = _out132
  var _4726_val _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4726_val
var _out133 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out133
_out133 = Companion_Default___.MarshallMessage(msg)
_4726_val = _out133
  var _4727_marshall__end__time uint64
  var _ = _4727_marshall__end__time
var _out134 uint64
  var _ = _out134
_out134 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_4727_marshall__end__time = _out134
  _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("PaxosMarshall_MarshallMessage"), _4725_marshall__start__time, _4727_marshall__end__time)
  { }
  { }
  var _4728_generic__marshall__start__time uint64
  var _ = _4728_generic__marshall__start__time
var _out135 uint64
  var _ = _out135
_out135 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_4728_generic__marshall__start__time = _out135
  var _out136 *_dafny.Array
  var _ = _out136
_out136 = _176_Common____GenericMarshalling__i_Compile.Companion_Default___.Marshall(_4726_val, Companion_Default___.CMessage__grammar())
data = _out136
  var _4729_generic__marshall__end__time uint64
  var _ = _4729_generic__marshall__end__time
var _out137 uint64
  var _ = _out137
_out137 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_4729_generic__marshall__end__time = _out137
  { }
  if ((msg).Is_CMessage__Request()) {
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("GenericMarshallMessage_Request"), _4728_generic__marshall__start__time, _4729_generic__marshall__end__time)
  } else if ((msg).Is_CMessage__1a()) {
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("GenericMarshallMessage_1a"), _4728_generic__marshall__start__time, _4729_generic__marshall__end__time)
  } else if ((msg).Is_CMessage__1b()) {
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("GenericMarshallMessage_1b"), _4728_generic__marshall__start__time, _4729_generic__marshall__end__time)
  } else if ((msg).Is_CMessage__2a()) {
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("GenericMarshallMessage_2a"), _4728_generic__marshall__start__time, _4729_generic__marshall__end__time)
  } else if ((msg).Is_CMessage__2b()) {
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("GenericMarshallMessage_2b"), _4728_generic__marshall__start__time, _4729_generic__marshall__end__time)
  } else if ((msg).Is_CMessage__Heartbeat()) {
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("GenericMarshallMessage_Heartbeat"), _4728_generic__marshall__start__time, _4729_generic__marshall__end__time)
  } else if ((msg).Is_CMessage__Reply()) {
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("GenericMarshallMessage_Reply"), _4728_generic__marshall__start__time, _4729_generic__marshall__end__time)
  } else if ((msg).Is_CMessage__AppStateRequest()) {
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("GenericMarshallMessage_AppStateRequest"), _4728_generic__marshall__start__time, _4729_generic__marshall__end__time)
  } else if ((msg).Is_CMessage__AppStateSupply()) {
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("GenericMarshallMessage_AppStateSupply"), _4728_generic__marshall__start__time, _4729_generic__marshall__end__time)
  } else if ((msg).Is_CMessage__StartingPhase2()) {
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("GenericMarshallMessage_StartingPhase2"), _4728_generic__marshall__start__time, _4729_generic__marshall__end__time)
  }
  { }
  return data
}
// End of class Default__
