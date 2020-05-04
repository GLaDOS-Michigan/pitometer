// Package _217_LiveRSL____CMessage__i_Compile
// Dafny module _217_LiveRSL____CMessage__i_Compile compiled into Go

package _217_LiveRSL____CMessage__i_Compile

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

type Dummy__ struct{}



// Definition of data type CMessage
type CMessage struct {
  Data_CMessage_
}

func (_this CMessage) Get() Data_CMessage_ {
  return _this.Data_CMessage_
}

type Data_CMessage_ interface {
  isCMessage()
}

type CompanionStruct_CMessage_ struct {}
var Companion_CMessage_ = CompanionStruct_CMessage_{}

type CMessage_CMessage__Invalid struct {
}

func (CMessage_CMessage__Invalid) isCMessage() {}

func (CompanionStruct_CMessage_) Create_CMessage__Invalid_() CMessage {
  return CMessage{CMessage_CMessage__Invalid{}}
}

func (_this CMessage) Is_CMessage__Invalid() bool {
  _, ok := _this.Get().(CMessage_CMessage__Invalid)
return ok
}

type CMessage_CMessage__Request struct {
  Seqno uint64
Val _197_LiveRSL____AppInterface__i_Compile.CAppMessage
}

func (CMessage_CMessage__Request) isCMessage() {}

func (CompanionStruct_CMessage_) Create_CMessage__Request_(Seqno uint64, Val _197_LiveRSL____AppInterface__i_Compile.CAppMessage) CMessage {
  return CMessage{CMessage_CMessage__Request{Seqno,Val}}
}

func (_this CMessage) Is_CMessage__Request() bool {
  _, ok := _this.Get().(CMessage_CMessage__Request)
return ok
}

type CMessage_CMessage__1a struct {
  Bal__1a _214_LiveRSL____CTypes__i_Compile.CBallot
}

func (CMessage_CMessage__1a) isCMessage() {}

func (CompanionStruct_CMessage_) Create_CMessage__1a_(Bal__1a _214_LiveRSL____CTypes__i_Compile.CBallot) CMessage {
  return CMessage{CMessage_CMessage__1a{Bal__1a}}
}

func (_this CMessage) Is_CMessage__1a() bool {
  _, ok := _this.Get().(CMessage_CMessage__1a)
return ok
}

type CMessage_CMessage__1b struct {
  Bal__1b _214_LiveRSL____CTypes__i_Compile.CBallot
Log__truncation__point _214_LiveRSL____CTypes__i_Compile.COperationNumber
Votes _214_LiveRSL____CTypes__i_Compile.CVotes
}

func (CMessage_CMessage__1b) isCMessage() {}

func (CompanionStruct_CMessage_) Create_CMessage__1b_(Bal__1b _214_LiveRSL____CTypes__i_Compile.CBallot, Log__truncation__point _214_LiveRSL____CTypes__i_Compile.COperationNumber, Votes _214_LiveRSL____CTypes__i_Compile.CVotes) CMessage {
  return CMessage{CMessage_CMessage__1b{Bal__1b,Log__truncation__point,Votes}}
}

func (_this CMessage) Is_CMessage__1b() bool {
  _, ok := _this.Get().(CMessage_CMessage__1b)
return ok
}

type CMessage_CMessage__2a struct {
  Bal__2a _214_LiveRSL____CTypes__i_Compile.CBallot
Opn__2a _214_LiveRSL____CTypes__i_Compile.COperationNumber
Val__2a _dafny.Seq
}

func (CMessage_CMessage__2a) isCMessage() {}

func (CompanionStruct_CMessage_) Create_CMessage__2a_(Bal__2a _214_LiveRSL____CTypes__i_Compile.CBallot, Opn__2a _214_LiveRSL____CTypes__i_Compile.COperationNumber, Val__2a _dafny.Seq) CMessage {
  return CMessage{CMessage_CMessage__2a{Bal__2a,Opn__2a,Val__2a}}
}

func (_this CMessage) Is_CMessage__2a() bool {
  _, ok := _this.Get().(CMessage_CMessage__2a)
return ok
}

type CMessage_CMessage__2b struct {
  Bal__2b _214_LiveRSL____CTypes__i_Compile.CBallot
Opn__2b _214_LiveRSL____CTypes__i_Compile.COperationNumber
Val__2b _dafny.Seq
}

func (CMessage_CMessage__2b) isCMessage() {}

func (CompanionStruct_CMessage_) Create_CMessage__2b_(Bal__2b _214_LiveRSL____CTypes__i_Compile.CBallot, Opn__2b _214_LiveRSL____CTypes__i_Compile.COperationNumber, Val__2b _dafny.Seq) CMessage {
  return CMessage{CMessage_CMessage__2b{Bal__2b,Opn__2b,Val__2b}}
}

func (_this CMessage) Is_CMessage__2b() bool {
  _, ok := _this.Get().(CMessage_CMessage__2b)
return ok
}

type CMessage_CMessage__Heartbeat struct {
  Bal__heartbeat _214_LiveRSL____CTypes__i_Compile.CBallot
Suspicious bool
Opn__ckpt _214_LiveRSL____CTypes__i_Compile.COperationNumber
}

func (CMessage_CMessage__Heartbeat) isCMessage() {}

func (CompanionStruct_CMessage_) Create_CMessage__Heartbeat_(Bal__heartbeat _214_LiveRSL____CTypes__i_Compile.CBallot, Suspicious bool, Opn__ckpt _214_LiveRSL____CTypes__i_Compile.COperationNumber) CMessage {
  return CMessage{CMessage_CMessage__Heartbeat{Bal__heartbeat,Suspicious,Opn__ckpt}}
}

func (_this CMessage) Is_CMessage__Heartbeat() bool {
  _, ok := _this.Get().(CMessage_CMessage__Heartbeat)
return ok
}

type CMessage_CMessage__Reply struct {
  Seqno__reply uint64
Reply _197_LiveRSL____AppInterface__i_Compile.CAppMessage
}

func (CMessage_CMessage__Reply) isCMessage() {}

func (CompanionStruct_CMessage_) Create_CMessage__Reply_(Seqno__reply uint64, Reply _197_LiveRSL____AppInterface__i_Compile.CAppMessage) CMessage {
  return CMessage{CMessage_CMessage__Reply{Seqno__reply,Reply}}
}

func (_this CMessage) Is_CMessage__Reply() bool {
  _, ok := _this.Get().(CMessage_CMessage__Reply)
return ok
}

type CMessage_CMessage__AppStateRequest struct {
  Bal__state__req _214_LiveRSL____CTypes__i_Compile.CBallot
Opn__state__req _214_LiveRSL____CTypes__i_Compile.COperationNumber
}

func (CMessage_CMessage__AppStateRequest) isCMessage() {}

func (CompanionStruct_CMessage_) Create_CMessage__AppStateRequest_(Bal__state__req _214_LiveRSL____CTypes__i_Compile.CBallot, Opn__state__req _214_LiveRSL____CTypes__i_Compile.COperationNumber) CMessage {
  return CMessage{CMessage_CMessage__AppStateRequest{Bal__state__req,Opn__state__req}}
}

func (_this CMessage) Is_CMessage__AppStateRequest() bool {
  _, ok := _this.Get().(CMessage_CMessage__AppStateRequest)
return ok
}

type CMessage_CMessage__AppStateSupply struct {
  Bal__state__supply _214_LiveRSL____CTypes__i_Compile.CBallot
Opn__state__supply _214_LiveRSL____CTypes__i_Compile.COperationNumber
App__state uint64
Reply__cache _dafny.Map
}

func (CMessage_CMessage__AppStateSupply) isCMessage() {}

func (CompanionStruct_CMessage_) Create_CMessage__AppStateSupply_(Bal__state__supply _214_LiveRSL____CTypes__i_Compile.CBallot, Opn__state__supply _214_LiveRSL____CTypes__i_Compile.COperationNumber, App__state uint64, Reply__cache _dafny.Map) CMessage {
  return CMessage{CMessage_CMessage__AppStateSupply{Bal__state__supply,Opn__state__supply,App__state,Reply__cache}}
}

func (_this CMessage) Is_CMessage__AppStateSupply() bool {
  _, ok := _this.Get().(CMessage_CMessage__AppStateSupply)
return ok
}

type CMessage_CMessage__StartingPhase2 struct {
  Bal__2 _214_LiveRSL____CTypes__i_Compile.CBallot
LogTruncationPoint__2 _214_LiveRSL____CTypes__i_Compile.COperationNumber
}

func (CMessage_CMessage__StartingPhase2) isCMessage() {}

func (CompanionStruct_CMessage_) Create_CMessage__StartingPhase2_(Bal__2 _214_LiveRSL____CTypes__i_Compile.CBallot, LogTruncationPoint__2 _214_LiveRSL____CTypes__i_Compile.COperationNumber) CMessage {
  return CMessage{CMessage_CMessage__StartingPhase2{Bal__2,LogTruncationPoint__2}}
}

func (_this CMessage) Is_CMessage__StartingPhase2() bool {
  _, ok := _this.Get().(CMessage_CMessage__StartingPhase2)
return ok
}

func (_this CMessage) Dtor_seqno() uint64 {
  return _this.Get().(CMessage_CMessage__Request).Seqno
}

func (_this CMessage) Dtor_val() _197_LiveRSL____AppInterface__i_Compile.CAppMessage {
  return _this.Get().(CMessage_CMessage__Request).Val
}

func (_this CMessage) Dtor_bal__1a() _214_LiveRSL____CTypes__i_Compile.CBallot {
  return _this.Get().(CMessage_CMessage__1a).Bal__1a
}

func (_this CMessage) Dtor_bal__1b() _214_LiveRSL____CTypes__i_Compile.CBallot {
  return _this.Get().(CMessage_CMessage__1b).Bal__1b
}

func (_this CMessage) Dtor_log__truncation__point() _214_LiveRSL____CTypes__i_Compile.COperationNumber {
  return _this.Get().(CMessage_CMessage__1b).Log__truncation__point
}

func (_this CMessage) Dtor_votes() _214_LiveRSL____CTypes__i_Compile.CVotes {
  return _this.Get().(CMessage_CMessage__1b).Votes
}

func (_this CMessage) Dtor_bal__2a() _214_LiveRSL____CTypes__i_Compile.CBallot {
  return _this.Get().(CMessage_CMessage__2a).Bal__2a
}

func (_this CMessage) Dtor_opn__2a() _214_LiveRSL____CTypes__i_Compile.COperationNumber {
  return _this.Get().(CMessage_CMessage__2a).Opn__2a
}

func (_this CMessage) Dtor_val__2a() _dafny.Seq {
  return _this.Get().(CMessage_CMessage__2a).Val__2a
}

func (_this CMessage) Dtor_bal__2b() _214_LiveRSL____CTypes__i_Compile.CBallot {
  return _this.Get().(CMessage_CMessage__2b).Bal__2b
}

func (_this CMessage) Dtor_opn__2b() _214_LiveRSL____CTypes__i_Compile.COperationNumber {
  return _this.Get().(CMessage_CMessage__2b).Opn__2b
}

func (_this CMessage) Dtor_val__2b() _dafny.Seq {
  return _this.Get().(CMessage_CMessage__2b).Val__2b
}

func (_this CMessage) Dtor_bal__heartbeat() _214_LiveRSL____CTypes__i_Compile.CBallot {
  return _this.Get().(CMessage_CMessage__Heartbeat).Bal__heartbeat
}

func (_this CMessage) Dtor_suspicious() bool {
  return _this.Get().(CMessage_CMessage__Heartbeat).Suspicious
}

func (_this CMessage) Dtor_opn__ckpt() _214_LiveRSL____CTypes__i_Compile.COperationNumber {
  return _this.Get().(CMessage_CMessage__Heartbeat).Opn__ckpt
}

func (_this CMessage) Dtor_seqno__reply() uint64 {
  return _this.Get().(CMessage_CMessage__Reply).Seqno__reply
}

func (_this CMessage) Dtor_reply() _197_LiveRSL____AppInterface__i_Compile.CAppMessage {
  return _this.Get().(CMessage_CMessage__Reply).Reply
}

func (_this CMessage) Dtor_bal__state__req() _214_LiveRSL____CTypes__i_Compile.CBallot {
  return _this.Get().(CMessage_CMessage__AppStateRequest).Bal__state__req
}

func (_this CMessage) Dtor_opn__state__req() _214_LiveRSL____CTypes__i_Compile.COperationNumber {
  return _this.Get().(CMessage_CMessage__AppStateRequest).Opn__state__req
}

func (_this CMessage) Dtor_bal__state__supply() _214_LiveRSL____CTypes__i_Compile.CBallot {
  return _this.Get().(CMessage_CMessage__AppStateSupply).Bal__state__supply
}

func (_this CMessage) Dtor_opn__state__supply() _214_LiveRSL____CTypes__i_Compile.COperationNumber {
  return _this.Get().(CMessage_CMessage__AppStateSupply).Opn__state__supply
}

func (_this CMessage) Dtor_app__state() uint64 {
  return _this.Get().(CMessage_CMessage__AppStateSupply).App__state
}

func (_this CMessage) Dtor_reply__cache() _dafny.Map {
  return _this.Get().(CMessage_CMessage__AppStateSupply).Reply__cache
}

func (_this CMessage) Dtor_bal__2() _214_LiveRSL____CTypes__i_Compile.CBallot {
  return _this.Get().(CMessage_CMessage__StartingPhase2).Bal__2
}

func (_this CMessage) Dtor_logTruncationPoint__2() _214_LiveRSL____CTypes__i_Compile.COperationNumber {
  return _this.Get().(CMessage_CMessage__StartingPhase2).LogTruncationPoint__2
}

func (_this CMessage) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case CMessage_CMessage__Invalid: {
      return "_217_LiveRSL____CMessage__i_Compile.CMessage.CMessage_Invalid"
    }
    case CMessage_CMessage__Request: {
      return "_217_LiveRSL____CMessage__i_Compile.CMessage.CMessage_Request" + "(" + _dafny.String(data.Seqno) + ", " + _dafny.String(data.Val) + ")"
    }
    case CMessage_CMessage__1a: {
      return "_217_LiveRSL____CMessage__i_Compile.CMessage.CMessage_1a" + "(" + _dafny.String(data.Bal__1a) + ")"
    }
    case CMessage_CMessage__1b: {
      return "_217_LiveRSL____CMessage__i_Compile.CMessage.CMessage_1b" + "(" + _dafny.String(data.Bal__1b) + ", " + _dafny.String(data.Log__truncation__point) + ", " + _dafny.String(data.Votes) + ")"
    }
    case CMessage_CMessage__2a: {
      return "_217_LiveRSL____CMessage__i_Compile.CMessage.CMessage_2a" + "(" + _dafny.String(data.Bal__2a) + ", " + _dafny.String(data.Opn__2a) + ", " + _dafny.String(data.Val__2a) + ")"
    }
    case CMessage_CMessage__2b: {
      return "_217_LiveRSL____CMessage__i_Compile.CMessage.CMessage_2b" + "(" + _dafny.String(data.Bal__2b) + ", " + _dafny.String(data.Opn__2b) + ", " + _dafny.String(data.Val__2b) + ")"
    }
    case CMessage_CMessage__Heartbeat: {
      return "_217_LiveRSL____CMessage__i_Compile.CMessage.CMessage_Heartbeat" + "(" + _dafny.String(data.Bal__heartbeat) + ", " + _dafny.String(data.Suspicious) + ", " + _dafny.String(data.Opn__ckpt) + ")"
    }
    case CMessage_CMessage__Reply: {
      return "_217_LiveRSL____CMessage__i_Compile.CMessage.CMessage_Reply" + "(" + _dafny.String(data.Seqno__reply) + ", " + _dafny.String(data.Reply) + ")"
    }
    case CMessage_CMessage__AppStateRequest: {
      return "_217_LiveRSL____CMessage__i_Compile.CMessage.CMessage_AppStateRequest" + "(" + _dafny.String(data.Bal__state__req) + ", " + _dafny.String(data.Opn__state__req) + ")"
    }
    case CMessage_CMessage__AppStateSupply: {
      return "_217_LiveRSL____CMessage__i_Compile.CMessage.CMessage_AppStateSupply" + "(" + _dafny.String(data.Bal__state__supply) + ", " + _dafny.String(data.Opn__state__supply) + ", " + _dafny.String(data.App__state) + ", " + _dafny.String(data.Reply__cache) + ")"
    }
    case CMessage_CMessage__StartingPhase2: {
      return "_217_LiveRSL____CMessage__i_Compile.CMessage.CMessage_StartingPhase2" + "(" + _dafny.String(data.Bal__2) + ", " + _dafny.String(data.LogTruncationPoint__2) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this CMessage) Equals(other CMessage) bool {
  switch data1 := _this.Get().(type) {
    case CMessage_CMessage__Invalid: {
      _, ok := other.Get().(CMessage_CMessage__Invalid)
return ok
    }
    case CMessage_CMessage__Request: {
      data2, ok := other.Get().(CMessage_CMessage__Request)
return ok && data1.Seqno == data2.Seqno && data1.Val.Equals(data2.Val)
    }
    case CMessage_CMessage__1a: {
      data2, ok := other.Get().(CMessage_CMessage__1a)
return ok && data1.Bal__1a.Equals(data2.Bal__1a)
    }
    case CMessage_CMessage__1b: {
      data2, ok := other.Get().(CMessage_CMessage__1b)
return ok && data1.Bal__1b.Equals(data2.Bal__1b) && data1.Log__truncation__point.Equals(data2.Log__truncation__point) && data1.Votes.Equals(data2.Votes)
    }
    case CMessage_CMessage__2a: {
      data2, ok := other.Get().(CMessage_CMessage__2a)
return ok && data1.Bal__2a.Equals(data2.Bal__2a) && data1.Opn__2a.Equals(data2.Opn__2a) && data1.Val__2a.Equals(data2.Val__2a)
    }
    case CMessage_CMessage__2b: {
      data2, ok := other.Get().(CMessage_CMessage__2b)
return ok && data1.Bal__2b.Equals(data2.Bal__2b) && data1.Opn__2b.Equals(data2.Opn__2b) && data1.Val__2b.Equals(data2.Val__2b)
    }
    case CMessage_CMessage__Heartbeat: {
      data2, ok := other.Get().(CMessage_CMessage__Heartbeat)
return ok && data1.Bal__heartbeat.Equals(data2.Bal__heartbeat) && data1.Suspicious == data2.Suspicious && data1.Opn__ckpt.Equals(data2.Opn__ckpt)
    }
    case CMessage_CMessage__Reply: {
      data2, ok := other.Get().(CMessage_CMessage__Reply)
return ok && data1.Seqno__reply == data2.Seqno__reply && data1.Reply.Equals(data2.Reply)
    }
    case CMessage_CMessage__AppStateRequest: {
      data2, ok := other.Get().(CMessage_CMessage__AppStateRequest)
return ok && data1.Bal__state__req.Equals(data2.Bal__state__req) && data1.Opn__state__req.Equals(data2.Opn__state__req)
    }
    case CMessage_CMessage__AppStateSupply: {
      data2, ok := other.Get().(CMessage_CMessage__AppStateSupply)
return ok && data1.Bal__state__supply.Equals(data2.Bal__state__supply) && data1.Opn__state__supply.Equals(data2.Opn__state__supply) && data1.App__state == data2.App__state && data1.Reply__cache.Equals(data2.Reply__cache)
    }
    case CMessage_CMessage__StartingPhase2: {
      data2, ok := other.Get().(CMessage_CMessage__StartingPhase2)
return ok && data1.Bal__2.Equals(data2.Bal__2) && data1.LogTruncationPoint__2.Equals(data2.LogTruncationPoint__2)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this CMessage) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(CMessage)
return ok && _this.Equals(typed)
}
func Type_CMessage_() _dafny.Type {
  return type_CMessage_{}
}

type type_CMessage_ struct {
}

func (_this type_CMessage_) Default() interface{} {
  return CMessage{CMessage_CMessage__Invalid{}}
}

func (_this type_CMessage_) String() string {
  return "_217_LiveRSL____CMessage__i_Compile.CMessage"
}
// End of data type CMessage

// Definition of data type CPacket
type CPacket struct {
  Data_CPacket_
}

func (_this CPacket) Get() Data_CPacket_ {
  return _this.Data_CPacket_
}

type Data_CPacket_ interface {
  isCPacket()
}

type CompanionStruct_CPacket_ struct {}
var Companion_CPacket_ = CompanionStruct_CPacket_{}

type CPacket_CPacket struct {
  Dst _9_Native____Io__s_Compile.EndPoint
Src _9_Native____Io__s_Compile.EndPoint
Msg _217_LiveRSL____CMessage__i_Compile.CMessage
}

func (CPacket_CPacket) isCPacket() {}

func (CompanionStruct_CPacket_) Create_CPacket_(Dst _9_Native____Io__s_Compile.EndPoint, Src _9_Native____Io__s_Compile.EndPoint, Msg _217_LiveRSL____CMessage__i_Compile.CMessage) CPacket {
  return CPacket{CPacket_CPacket{Dst,Src,Msg}}
}

func (_this CPacket) Is_CPacket() bool {
  _, ok := _this.Get().(CPacket_CPacket)
return ok
}

func (_this CPacket) Dtor_dst() _9_Native____Io__s_Compile.EndPoint {
  return _this.Get().(CPacket_CPacket).Dst
}

func (_this CPacket) Dtor_src() _9_Native____Io__s_Compile.EndPoint {
  return _this.Get().(CPacket_CPacket).Src
}

func (_this CPacket) Dtor_msg() _217_LiveRSL____CMessage__i_Compile.CMessage {
  return _this.Get().(CPacket_CPacket).Msg
}

func (_this CPacket) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case CPacket_CPacket: {
      return "_217_LiveRSL____CMessage__i_Compile.CPacket.CPacket" + "(" + _dafny.String(data.Dst) + ", " + _dafny.String(data.Src) + ", " + _dafny.String(data.Msg) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this CPacket) Equals(other CPacket) bool {
  switch data1 := _this.Get().(type) {
    case CPacket_CPacket: {
      data2, ok := other.Get().(CPacket_CPacket)
return ok && data1.Dst.Equals(data2.Dst) && data1.Src.Equals(data2.Src) && data1.Msg.Equals(data2.Msg)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this CPacket) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(CPacket)
return ok && _this.Equals(typed)
}
func Type_CPacket_() _dafny.Type {
  return type_CPacket_{}
}

type type_CPacket_ struct {
}

func (_this type_CPacket_) Default() interface{} {
  return CPacket{CPacket_CPacket{_9_Native____Io__s_Compile.Type_EndPoint_().Default().(_9_Native____Io__s_Compile.EndPoint), _9_Native____Io__s_Compile.Type_EndPoint_().Default().(_9_Native____Io__s_Compile.EndPoint), Type_CMessage_().Default().(_217_LiveRSL____CMessage__i_Compile.CMessage)}}
}

func (_this type_CPacket_) String() string {
  return "_217_LiveRSL____CMessage__i_Compile.CPacket"
}
// End of data type CPacket

// Definition of data type CBroadcast
type CBroadcast struct {
  Data_CBroadcast_
}

func (_this CBroadcast) Get() Data_CBroadcast_ {
  return _this.Data_CBroadcast_
}

type Data_CBroadcast_ interface {
  isCBroadcast()
}

type CompanionStruct_CBroadcast_ struct {}
var Companion_CBroadcast_ = CompanionStruct_CBroadcast_{}

type CBroadcast_CBroadcast struct {
  Src _9_Native____Io__s_Compile.EndPoint
Dsts _dafny.Seq
Msg _217_LiveRSL____CMessage__i_Compile.CMessage
}

func (CBroadcast_CBroadcast) isCBroadcast() {}

func (CompanionStruct_CBroadcast_) Create_CBroadcast_(Src _9_Native____Io__s_Compile.EndPoint, Dsts _dafny.Seq, Msg _217_LiveRSL____CMessage__i_Compile.CMessage) CBroadcast {
  return CBroadcast{CBroadcast_CBroadcast{Src,Dsts,Msg}}
}

func (_this CBroadcast) Is_CBroadcast() bool {
  _, ok := _this.Get().(CBroadcast_CBroadcast)
return ok
}

type CBroadcast_CBroadcastNop struct {
}

func (CBroadcast_CBroadcastNop) isCBroadcast() {}

func (CompanionStruct_CBroadcast_) Create_CBroadcastNop_() CBroadcast {
  return CBroadcast{CBroadcast_CBroadcastNop{}}
}

func (_this CBroadcast) Is_CBroadcastNop() bool {
  _, ok := _this.Get().(CBroadcast_CBroadcastNop)
return ok
}

func (_this CBroadcast) Dtor_src() _9_Native____Io__s_Compile.EndPoint {
  return _this.Get().(CBroadcast_CBroadcast).Src
}

func (_this CBroadcast) Dtor_dsts() _dafny.Seq {
  return _this.Get().(CBroadcast_CBroadcast).Dsts
}

func (_this CBroadcast) Dtor_msg() _217_LiveRSL____CMessage__i_Compile.CMessage {
  return _this.Get().(CBroadcast_CBroadcast).Msg
}

func (_this CBroadcast) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case CBroadcast_CBroadcast: {
      return "_217_LiveRSL____CMessage__i_Compile.CBroadcast.CBroadcast" + "(" + _dafny.String(data.Src) + ", " + _dafny.String(data.Dsts) + ", " + _dafny.String(data.Msg) + ")"
    }
    case CBroadcast_CBroadcastNop: {
      return "_217_LiveRSL____CMessage__i_Compile.CBroadcast.CBroadcastNop"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this CBroadcast) Equals(other CBroadcast) bool {
  switch data1 := _this.Get().(type) {
    case CBroadcast_CBroadcast: {
      data2, ok := other.Get().(CBroadcast_CBroadcast)
return ok && data1.Src.Equals(data2.Src) && data1.Dsts.Equals(data2.Dsts) && data1.Msg.Equals(data2.Msg)
    }
    case CBroadcast_CBroadcastNop: {
      _, ok := other.Get().(CBroadcast_CBroadcastNop)
return ok
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this CBroadcast) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(CBroadcast)
return ok && _this.Equals(typed)
}
func Type_CBroadcast_() _dafny.Type {
  return type_CBroadcast_{}
}

type type_CBroadcast_ struct {
}

func (_this type_CBroadcast_) Default() interface{} {
  return CBroadcast{CBroadcast_CBroadcast{_9_Native____Io__s_Compile.Type_EndPoint_().Default().(_9_Native____Io__s_Compile.EndPoint), _dafny.EmptySeq, Type_CMessage_().Default().(_217_LiveRSL____CMessage__i_Compile.CMessage)}}
}

func (_this type_CBroadcast_) String() string {
  return "_217_LiveRSL____CMessage__i_Compile.CBroadcast"
}
// End of data type CBroadcast

// Definition of data type OutboundPackets
type OutboundPackets struct {
  Data_OutboundPackets_
}

func (_this OutboundPackets) Get() Data_OutboundPackets_ {
  return _this.Data_OutboundPackets_
}

type Data_OutboundPackets_ interface {
  isOutboundPackets()
}

type CompanionStruct_OutboundPackets_ struct {}
var Companion_OutboundPackets_ = CompanionStruct_OutboundPackets_{}

type OutboundPackets_Broadcast struct {
  Broadcast _217_LiveRSL____CMessage__i_Compile.CBroadcast
}

func (OutboundPackets_Broadcast) isOutboundPackets() {}

func (CompanionStruct_OutboundPackets_) Create_Broadcast_(Broadcast _217_LiveRSL____CMessage__i_Compile.CBroadcast) OutboundPackets {
  return OutboundPackets{OutboundPackets_Broadcast{Broadcast}}
}

func (_this OutboundPackets) Is_Broadcast() bool {
  _, ok := _this.Get().(OutboundPackets_Broadcast)
return ok
}

type OutboundPackets_OutboundPacket struct {
  P _135_Logic____Option__i_Compile.Option
}

func (OutboundPackets_OutboundPacket) isOutboundPackets() {}

func (CompanionStruct_OutboundPackets_) Create_OutboundPacket_(P _135_Logic____Option__i_Compile.Option) OutboundPackets {
  return OutboundPackets{OutboundPackets_OutboundPacket{P}}
}

func (_this OutboundPackets) Is_OutboundPacket() bool {
  _, ok := _this.Get().(OutboundPackets_OutboundPacket)
return ok
}

type OutboundPackets_PacketSequence struct {
  S _dafny.Seq
}

func (OutboundPackets_PacketSequence) isOutboundPackets() {}

func (CompanionStruct_OutboundPackets_) Create_PacketSequence_(S _dafny.Seq) OutboundPackets {
  return OutboundPackets{OutboundPackets_PacketSequence{S}}
}

func (_this OutboundPackets) Is_PacketSequence() bool {
  _, ok := _this.Get().(OutboundPackets_PacketSequence)
return ok
}

func (_this OutboundPackets) Dtor_broadcast() _217_LiveRSL____CMessage__i_Compile.CBroadcast {
  return _this.Get().(OutboundPackets_Broadcast).Broadcast
}

func (_this OutboundPackets) Dtor_p() _135_Logic____Option__i_Compile.Option {
  return _this.Get().(OutboundPackets_OutboundPacket).P
}

func (_this OutboundPackets) Dtor_s() _dafny.Seq {
  return _this.Get().(OutboundPackets_PacketSequence).S
}

func (_this OutboundPackets) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case OutboundPackets_Broadcast: {
      return "_217_LiveRSL____CMessage__i_Compile.OutboundPackets.Broadcast" + "(" + _dafny.String(data.Broadcast) + ")"
    }
    case OutboundPackets_OutboundPacket: {
      return "_217_LiveRSL____CMessage__i_Compile.OutboundPackets.OutboundPacket" + "(" + _dafny.String(data.P) + ")"
    }
    case OutboundPackets_PacketSequence: {
      return "_217_LiveRSL____CMessage__i_Compile.OutboundPackets.PacketSequence" + "(" + _dafny.String(data.S) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this OutboundPackets) Equals(other OutboundPackets) bool {
  switch data1 := _this.Get().(type) {
    case OutboundPackets_Broadcast: {
      data2, ok := other.Get().(OutboundPackets_Broadcast)
return ok && data1.Broadcast.Equals(data2.Broadcast)
    }
    case OutboundPackets_OutboundPacket: {
      data2, ok := other.Get().(OutboundPackets_OutboundPacket)
return ok && data1.P.Equals(data2.P)
    }
    case OutboundPackets_PacketSequence: {
      data2, ok := other.Get().(OutboundPackets_PacketSequence)
return ok && data1.S.Equals(data2.S)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this OutboundPackets) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(OutboundPackets)
return ok && _this.Equals(typed)
}
func Type_OutboundPackets_() _dafny.Type {
  return type_OutboundPackets_{}
}

type type_OutboundPackets_ struct {
}

func (_this type_OutboundPackets_) Default() interface{} {
  return OutboundPackets{OutboundPackets_Broadcast{Type_CBroadcast_().Default().(_217_LiveRSL____CMessage__i_Compile.CBroadcast)}}
}

func (_this type_OutboundPackets_) String() string {
  return "_217_LiveRSL____CMessage__i_Compile.OutboundPackets"
}
// End of data type OutboundPackets

