// Package _197_LiveRSL____AppInterface__i_Compile
// Dafny module _197_LiveRSL____AppInterface__i_Compile compiled into Go

package _197_LiveRSL____AppInterface__i_Compile

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

type Dummy__ struct{}










// Definition of data type CAppMessage
type CAppMessage struct {
  Data_CAppMessage_
}

func (_this CAppMessage) Get() Data_CAppMessage_ {
  return _this.Data_CAppMessage_
}

type Data_CAppMessage_ interface {
  isCAppMessage()
}

type CompanionStruct_CAppMessage_ struct {}
var Companion_CAppMessage_ = CompanionStruct_CAppMessage_{}

type CAppMessage_CAppIncrement struct {
}

func (CAppMessage_CAppIncrement) isCAppMessage() {}

func (CompanionStruct_CAppMessage_) Create_CAppIncrement_() CAppMessage {
  return CAppMessage{CAppMessage_CAppIncrement{}}
}

func (_this CAppMessage) Is_CAppIncrement() bool {
  _, ok := _this.Get().(CAppMessage_CAppIncrement)
return ok
}

type CAppMessage_CAppReply struct {
  Response uint64
}

func (CAppMessage_CAppReply) isCAppMessage() {}

func (CompanionStruct_CAppMessage_) Create_CAppReply_(Response uint64) CAppMessage {
  return CAppMessage{CAppMessage_CAppReply{Response}}
}

func (_this CAppMessage) Is_CAppReply() bool {
  _, ok := _this.Get().(CAppMessage_CAppReply)
return ok
}

type CAppMessage_CAppInvalid struct {
}

func (CAppMessage_CAppInvalid) isCAppMessage() {}

func (CompanionStruct_CAppMessage_) Create_CAppInvalid_() CAppMessage {
  return CAppMessage{CAppMessage_CAppInvalid{}}
}

func (_this CAppMessage) Is_CAppInvalid() bool {
  _, ok := _this.Get().(CAppMessage_CAppInvalid)
return ok
}

func (_this CAppMessage) Dtor_response() uint64 {
  return _this.Get().(CAppMessage_CAppReply).Response
}

func (_this CAppMessage) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case CAppMessage_CAppIncrement: {
      return "_197_LiveRSL____AppInterface__i_Compile.CAppMessage.CAppIncrement"
    }
    case CAppMessage_CAppReply: {
      return "_197_LiveRSL____AppInterface__i_Compile.CAppMessage.CAppReply" + "(" + _dafny.String(data.Response) + ")"
    }
    case CAppMessage_CAppInvalid: {
      return "_197_LiveRSL____AppInterface__i_Compile.CAppMessage.CAppInvalid"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this CAppMessage) Equals(other CAppMessage) bool {
  switch data1 := _this.Get().(type) {
    case CAppMessage_CAppIncrement: {
      _, ok := other.Get().(CAppMessage_CAppIncrement)
return ok
    }
    case CAppMessage_CAppReply: {
      data2, ok := other.Get().(CAppMessage_CAppReply)
return ok && data1.Response == data2.Response
    }
    case CAppMessage_CAppInvalid: {
      _, ok := other.Get().(CAppMessage_CAppInvalid)
return ok
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this CAppMessage) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(CAppMessage)
return ok && _this.Equals(typed)
}
func Type_CAppMessage_() _dafny.Type {
  return type_CAppMessage_{}
}

type type_CAppMessage_ struct {
}

func (_this type_CAppMessage_) Default() interface{} {
  return CAppMessage{CAppMessage_CAppIncrement{}}
}

func (_this type_CAppMessage_) String() string {
  return "_197_LiveRSL____AppInterface__i_Compile.CAppMessage"
}
// End of data type CAppMessage

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
  return "_197_LiveRSL____AppInterface__i_Compile.Default__"
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
  return "_197_LiveRSL____AppInterface__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) CAppState__grammar() _176_Common____GenericMarshalling__i_Compile.G {
  return _176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GUint64{}}
}
func (_this *CompanionStruct_Default___) Parse__AppState(val _176_Common____GenericMarshalling__i_Compile.V) uint64 {
  return uint64((val).Dtor_u())
}
func (_this *CompanionStruct_Default___) AppStateMarshallable(msg uint64) bool {
  return true
}
func (_this *CompanionStruct_Default___) MarshallAppState(c uint64) _176_Common____GenericMarshalling__i_Compile.V {
  goto TAIL_CALL_START
TAIL_CALL_START:
var val _176_Common____GenericMarshalling__i_Compile.V = _176_Common____GenericMarshalling__i_Compile.Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = val
  val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VUint64{c}}
  return val
}
func (_this *CompanionStruct_Default___) ValidAppMessage(c _197_LiveRSL____AppInterface__i_Compile.CAppMessage) bool {
  return true
}
func (_this *CompanionStruct_Default___) CAppMessage__grammar() _176_Common____GenericMarshalling__i_Compile.G {
  return _176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GTaggedUnion{_dafny.SeqOf(_176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GTuple{_dafny.SeqOf()}}, _176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GUint64{}}, _176_Common____GenericMarshalling__i_Compile.G{_176_Common____GenericMarshalling__i_Compile.G_GTuple{_dafny.SeqOf()}})}}
}
func (_this *CompanionStruct_Default___) Parse__AppMessage(val _176_Common____GenericMarshalling__i_Compile.V) _197_LiveRSL____AppInterface__i_Compile.CAppMessage {
  if (((val).Dtor_c()) == (uint64(0))) {
    return CAppMessage{CAppMessage_CAppIncrement{}}
  } else  {
    if (((val).Dtor_c()) == (uint64(1))) {
      return CAppMessage{CAppMessage_CAppReply{((val).Dtor_val()).Dtor_u()}}
    } else  {
      return CAppMessage{CAppMessage_CAppInvalid{}}
    }
  }
}
func (_this *CompanionStruct_Default___) MarshallCAppMessage(c _197_LiveRSL____AppInterface__i_Compile.CAppMessage) _176_Common____GenericMarshalling__i_Compile.V {
  goto TAIL_CALL_START
TAIL_CALL_START:
var val _176_Common____GenericMarshalling__i_Compile.V = _176_Common____GenericMarshalling__i_Compile.Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = val
  var _source4 _197_LiveRSL____AppInterface__i_Compile.CAppMessage = c
  var _ = _source4
if (_source4.Is_CAppIncrement()) {
    val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VCase{uint64(0), _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VTuple{_dafny.SeqOf()}}}}
  } else if (_source4.Is_CAppReply()) {
    var _4638_response uint64 = _source4.Get().(_197_LiveRSL____AppInterface__i_Compile.CAppMessage_CAppReply).Response
    var _ = _4638_response
    val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VCase{uint64(1), _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VUint64{_4638_response}}}}
  } else {
    val = _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VCase{uint64(2), _176_Common____GenericMarshalling__i_Compile.V{_176_Common____GenericMarshalling__i_Compile.V_VTuple{_dafny.SeqOf()}}}}
  }
  return val
}
func (_this *CompanionStruct_Default___) Max__val__len() _dafny.Int {
  return _dafny.IntOfInt64(64)
}
func (_this *CompanionStruct_Default___) CAppState__Init() uint64 {
  goto TAIL_CALL_START
TAIL_CALL_START:
var s uint64 = 0
  var _ = s
  s = uint64(0)
  return s
}
func (_this *CompanionStruct_Default___) CappedIncrImpl(v uint64) uint64 {
  goto TAIL_CALL_START
TAIL_CALL_START:
var v_k uint64 = 0
  var _ = v_k
  if ((v) == (uint64(18446744073709551615))) {
    v_k = v
  } else {
    v_k = (v) + (uint64(1))
  }
  return v_k
}
func (_this *CompanionStruct_Default___) HandleAppRequest(appState uint64, request _197_LiveRSL____AppInterface__i_Compile.CAppMessage) (uint64, _197_LiveRSL____AppInterface__i_Compile.CAppMessage) {
  goto TAIL_CALL_START
TAIL_CALL_START:
var appState_k uint64 = 0
  var _ = appState_k
var reply _197_LiveRSL____AppInterface__i_Compile.CAppMessage = Type_CAppMessage_().Default().(_197_LiveRSL____AppInterface__i_Compile.CAppMessage)
  var _ = reply
  if ((request).Is_CAppIncrement()) {
    var _out60 uint64
    var _ = _out60
_out60 = Companion_Default___.CappedIncrImpl(appState)
appState_k = _out60
    reply = CAppMessage{CAppMessage_CAppReply{appState_k}}
  } else {
    appState_k = appState
    reply = CAppMessage{CAppMessage_CAppInvalid{}}
  }
  return appState_k,reply
}
// End of class Default__
