// Package _111_PacketParsing__i_Compile
// Dafny module _111_PacketParsing__i_Compile compiled into Go

package _111_PacketParsing__i_Compile

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
  return "_111_PacketParsing__i_Compile.Default__"
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
  return "_111_PacketParsing__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) CMessageTransferGrammar() _107_Common____GenericMarshalling__i_Compile.G {
  return _107_Common____GenericMarshalling__i_Compile.G{_107_Common____GenericMarshalling__i_Compile.G_GUint64{}}
}
func (_this *CompanionStruct_Default___) CMessageLockedGrammar() _107_Common____GenericMarshalling__i_Compile.G {
  return _107_Common____GenericMarshalling__i_Compile.G{_107_Common____GenericMarshalling__i_Compile.G_GUint64{}}
}
func (_this *CompanionStruct_Default___) CMessageGrammar() _107_Common____GenericMarshalling__i_Compile.G {
  return _107_Common____GenericMarshalling__i_Compile.G{_107_Common____GenericMarshalling__i_Compile.G_GTaggedUnion{_dafny.SeqOf(Companion_Default___.CMessageTransferGrammar(), Companion_Default___.CMessageLockedGrammar())}}
}
func (_this *CompanionStruct_Default___) ParseCMessageTransfer(val _107_Common____GenericMarshalling__i_Compile.V) _39_Message__i_Compile.CMessage {
  return _39_Message__i_Compile.CMessage{_39_Message__i_Compile.CMessage_CTransfer{(val).Dtor_u()}}
}
func (_this *CompanionStruct_Default___) ParseCMessageLocked(val _107_Common____GenericMarshalling__i_Compile.V) _39_Message__i_Compile.CMessage {
  return _39_Message__i_Compile.CMessage{_39_Message__i_Compile.CMessage_CLocked{(val).Dtor_u()}}
}
func (_this *CompanionStruct_Default___) ParseCMessage(val _107_Common____GenericMarshalling__i_Compile.V) _39_Message__i_Compile.CMessage {
  if (((val).Dtor_c()) == (uint64(0))) {
    return Companion_Default___.ParseCMessageTransfer((val).Dtor_val())
  } else  {
    return Companion_Default___.ParseCMessageLocked((val).Dtor_val())
  }
}
func (_this *CompanionStruct_Default___) DemarshallDataMethod(data *_dafny.Array) _39_Message__i_Compile.CMessage {
  goto TAIL_CALL_START
TAIL_CALL_START:
var msg _39_Message__i_Compile.CMessage = _39_Message__i_Compile.Type_CMessage_().Default().(_39_Message__i_Compile.CMessage)
  var _ = msg
  var _1664_success bool
  var _ = _1664_success
var _1665_val _107_Common____GenericMarshalling__i_Compile.V
  var _ = _1665_val
var _out59 bool
  var _ = _out59
var _out60 _107_Common____GenericMarshalling__i_Compile.V
  var _ = _out60
_out59,_out60 = _107_Common____GenericMarshalling__i_Compile.Companion_Default___.Demarshall(data, Companion_Default___.CMessageGrammar())
_1664_success = _out59
_1665_val = _out60
  if (_1664_success) {
    msg = Companion_Default___.ParseCMessage(_1665_val)
    { }
  } else {
    msg = _39_Message__i_Compile.CMessage{_39_Message__i_Compile.CMessage_CInvalid{}}
  }
  return msg
}
func (_this *CompanionStruct_Default___) MarshallMessageTransfer(c _39_Message__i_Compile.CMessage) _107_Common____GenericMarshalling__i_Compile.V {
  goto TAIL_CALL_START
TAIL_CALL_START:
var val _107_Common____GenericMarshalling__i_Compile.V = _107_Common____GenericMarshalling__i_Compile.Type_V_().Default().(_107_Common____GenericMarshalling__i_Compile.V)
  var _ = val
  val = _107_Common____GenericMarshalling__i_Compile.V{_107_Common____GenericMarshalling__i_Compile.V_VUint64{(c).Dtor_transfer__epoch()}}
  return val
}
func (_this *CompanionStruct_Default___) MarshallMessageLocked(c _39_Message__i_Compile.CMessage) _107_Common____GenericMarshalling__i_Compile.V {
  goto TAIL_CALL_START
TAIL_CALL_START:
var val _107_Common____GenericMarshalling__i_Compile.V = _107_Common____GenericMarshalling__i_Compile.Type_V_().Default().(_107_Common____GenericMarshalling__i_Compile.V)
  var _ = val
  val = _107_Common____GenericMarshalling__i_Compile.V{_107_Common____GenericMarshalling__i_Compile.V_VUint64{(c).Dtor_locked__epoch()}}
  return val
}
func (_this *CompanionStruct_Default___) MarshallMessage(c _39_Message__i_Compile.CMessage) _107_Common____GenericMarshalling__i_Compile.V {
  var val _107_Common____GenericMarshalling__i_Compile.V = _107_Common____GenericMarshalling__i_Compile.Type_V_().Default().(_107_Common____GenericMarshalling__i_Compile.V)
  var _ = val
  if ((c).Is_CTransfer()) {
    var _1666_msg _107_Common____GenericMarshalling__i_Compile.V
    var _ = _1666_msg
var _out61 _107_Common____GenericMarshalling__i_Compile.V
    var _ = _out61
_out61 = Companion_Default___.MarshallMessageTransfer(c)
_1666_msg = _out61
    val = _107_Common____GenericMarshalling__i_Compile.V{_107_Common____GenericMarshalling__i_Compile.V_VCase{uint64(0), _1666_msg}}
  } else if ((c).Is_CLocked()) {
    var _1667_msg _107_Common____GenericMarshalling__i_Compile.V
    var _ = _1667_msg
var _out62 _107_Common____GenericMarshalling__i_Compile.V
    var _ = _out62
_out62 = Companion_Default___.MarshallMessageLocked(c)
_1667_msg = _out62
    val = _107_Common____GenericMarshalling__i_Compile.V{_107_Common____GenericMarshalling__i_Compile.V_VCase{uint64(1), _1667_msg}}
  } else { }
  return val
}
func (_this *CompanionStruct_Default___) MarshallLockMessage(msg _39_Message__i_Compile.CMessage) *_dafny.Array {
  goto TAIL_CALL_START
TAIL_CALL_START:
var data *_dafny.Array = _dafny.NewArrayWithValue(0, _dafny.IntOf(0))
  var _ = data
  var _1668_val _107_Common____GenericMarshalling__i_Compile.V
  var _ = _1668_val
var _out63 _107_Common____GenericMarshalling__i_Compile.V
  var _ = _out63
_out63 = Companion_Default___.MarshallMessage(msg)
_1668_val = _out63
  var _out64 *_dafny.Array
  var _ = _out64
_out64 = _107_Common____GenericMarshalling__i_Compile.Companion_Default___.Marshall(_1668_val, Companion_Default___.CMessageGrammar())
data = _out64
  return data
}
// End of class Default__
