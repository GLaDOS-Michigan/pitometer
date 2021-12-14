// Package _65_LiveRSL____Message__i_Compile
// Dafny module _65_LiveRSL____Message__i_Compile compiled into Go

package _65_LiveRSL____Message__i_Compile

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

type Dummy__ struct{}



// Definition of data type RslMessage
type RslMessage struct {
  Data_RslMessage_
}

func (_this RslMessage) Get() Data_RslMessage_ {
  return _this.Data_RslMessage_
}

type Data_RslMessage_ interface {
  isRslMessage()
}

type CompanionStruct_RslMessage_ struct {}
var Companion_RslMessage_ = CompanionStruct_RslMessage_{}

type RslMessage_RslMessage__Invalid struct {
}

func (RslMessage_RslMessage__Invalid) isRslMessage() {}

func (CompanionStruct_RslMessage_) Create_RslMessage__Invalid_() RslMessage {
  return RslMessage{RslMessage_RslMessage__Invalid{}}
}

func (_this RslMessage) Is_RslMessage__Invalid() bool {
  _, ok := _this.Get().(RslMessage_RslMessage__Invalid)
return ok
}

type RslMessage_RslMessage__Request struct {
  Seqno__req _dafny.Int
Val _50_AppStateMachine__i_Compile.AppMessage_k
}

func (RslMessage_RslMessage__Request) isRslMessage() {}

func (CompanionStruct_RslMessage_) Create_RslMessage__Request_(Seqno__req _dafny.Int, Val _50_AppStateMachine__i_Compile.AppMessage_k) RslMessage {
  return RslMessage{RslMessage_RslMessage__Request{Seqno__req,Val}}
}

func (_this RslMessage) Is_RslMessage__Request() bool {
  _, ok := _this.Get().(RslMessage_RslMessage__Request)
return ok
}

type RslMessage_RslMessage__1a struct {
  Bal__1a _56_LiveRSL____Types__i_Compile.Ballot
}

func (RslMessage_RslMessage__1a) isRslMessage() {}

func (CompanionStruct_RslMessage_) Create_RslMessage__1a_(Bal__1a _56_LiveRSL____Types__i_Compile.Ballot) RslMessage {
  return RslMessage{RslMessage_RslMessage__1a{Bal__1a}}
}

func (_this RslMessage) Is_RslMessage__1a() bool {
  _, ok := _this.Get().(RslMessage_RslMessage__1a)
return ok
}

type RslMessage_RslMessage__1b struct {
  Bal__1b _56_LiveRSL____Types__i_Compile.Ballot
Log__truncation__point _dafny.Int
Votes _dafny.Map
}

func (RslMessage_RslMessage__1b) isRslMessage() {}

func (CompanionStruct_RslMessage_) Create_RslMessage__1b_(Bal__1b _56_LiveRSL____Types__i_Compile.Ballot, Log__truncation__point _dafny.Int, Votes _dafny.Map) RslMessage {
  return RslMessage{RslMessage_RslMessage__1b{Bal__1b,Log__truncation__point,Votes}}
}

func (_this RslMessage) Is_RslMessage__1b() bool {
  _, ok := _this.Get().(RslMessage_RslMessage__1b)
return ok
}

type RslMessage_RslMessage__2a struct {
  Bal__2a _56_LiveRSL____Types__i_Compile.Ballot
Opn__2a _dafny.Int
Val__2a _dafny.Seq
}

func (RslMessage_RslMessage__2a) isRslMessage() {}

func (CompanionStruct_RslMessage_) Create_RslMessage__2a_(Bal__2a _56_LiveRSL____Types__i_Compile.Ballot, Opn__2a _dafny.Int, Val__2a _dafny.Seq) RslMessage {
  return RslMessage{RslMessage_RslMessage__2a{Bal__2a,Opn__2a,Val__2a}}
}

func (_this RslMessage) Is_RslMessage__2a() bool {
  _, ok := _this.Get().(RslMessage_RslMessage__2a)
return ok
}

type RslMessage_RslMessage__2b struct {
  Bal__2b _56_LiveRSL____Types__i_Compile.Ballot
Opn__2b _dafny.Int
Val__2b _dafny.Seq
}

func (RslMessage_RslMessage__2b) isRslMessage() {}

func (CompanionStruct_RslMessage_) Create_RslMessage__2b_(Bal__2b _56_LiveRSL____Types__i_Compile.Ballot, Opn__2b _dafny.Int, Val__2b _dafny.Seq) RslMessage {
  return RslMessage{RslMessage_RslMessage__2b{Bal__2b,Opn__2b,Val__2b}}
}

func (_this RslMessage) Is_RslMessage__2b() bool {
  _, ok := _this.Get().(RslMessage_RslMessage__2b)
return ok
}

type RslMessage_RslMessage__Heartbeat struct {
  Bal__heartbeat _56_LiveRSL____Types__i_Compile.Ballot
Suspicious bool
Opn__ckpt _dafny.Int
}

func (RslMessage_RslMessage__Heartbeat) isRslMessage() {}

func (CompanionStruct_RslMessage_) Create_RslMessage__Heartbeat_(Bal__heartbeat _56_LiveRSL____Types__i_Compile.Ballot, Suspicious bool, Opn__ckpt _dafny.Int) RslMessage {
  return RslMessage{RslMessage_RslMessage__Heartbeat{Bal__heartbeat,Suspicious,Opn__ckpt}}
}

func (_this RslMessage) Is_RslMessage__Heartbeat() bool {
  _, ok := _this.Get().(RslMessage_RslMessage__Heartbeat)
return ok
}

type RslMessage_RslMessage__Reply struct {
  Seqno__reply _dafny.Int
Reply _50_AppStateMachine__i_Compile.AppMessage_k
}

func (RslMessage_RslMessage__Reply) isRslMessage() {}

func (CompanionStruct_RslMessage_) Create_RslMessage__Reply_(Seqno__reply _dafny.Int, Reply _50_AppStateMachine__i_Compile.AppMessage_k) RslMessage {
  return RslMessage{RslMessage_RslMessage__Reply{Seqno__reply,Reply}}
}

func (_this RslMessage) Is_RslMessage__Reply() bool {
  _, ok := _this.Get().(RslMessage_RslMessage__Reply)
return ok
}

type RslMessage_RslMessage__AppStateRequest struct {
  Bal__state__req _56_LiveRSL____Types__i_Compile.Ballot
Opn__state__req _dafny.Int
}

func (RslMessage_RslMessage__AppStateRequest) isRslMessage() {}

func (CompanionStruct_RslMessage_) Create_RslMessage__AppStateRequest_(Bal__state__req _56_LiveRSL____Types__i_Compile.Ballot, Opn__state__req _dafny.Int) RslMessage {
  return RslMessage{RslMessage_RslMessage__AppStateRequest{Bal__state__req,Opn__state__req}}
}

func (_this RslMessage) Is_RslMessage__AppStateRequest() bool {
  _, ok := _this.Get().(RslMessage_RslMessage__AppStateRequest)
return ok
}

type RslMessage_RslMessage__AppStateSupply struct {
  Bal__state__supply _56_LiveRSL____Types__i_Compile.Ballot
Opn__state__supply _dafny.Int
App__state uint64
Reply__cache _dafny.Map
}

func (RslMessage_RslMessage__AppStateSupply) isRslMessage() {}

func (CompanionStruct_RslMessage_) Create_RslMessage__AppStateSupply_(Bal__state__supply _56_LiveRSL____Types__i_Compile.Ballot, Opn__state__supply _dafny.Int, App__state uint64, Reply__cache _dafny.Map) RslMessage {
  return RslMessage{RslMessage_RslMessage__AppStateSupply{Bal__state__supply,Opn__state__supply,App__state,Reply__cache}}
}

func (_this RslMessage) Is_RslMessage__AppStateSupply() bool {
  _, ok := _this.Get().(RslMessage_RslMessage__AppStateSupply)
return ok
}

type RslMessage_RslMessage__StartingPhase2 struct {
  Bal__2 _56_LiveRSL____Types__i_Compile.Ballot
LogTruncationPoint__2 _dafny.Int
}

func (RslMessage_RslMessage__StartingPhase2) isRslMessage() {}

func (CompanionStruct_RslMessage_) Create_RslMessage__StartingPhase2_(Bal__2 _56_LiveRSL____Types__i_Compile.Ballot, LogTruncationPoint__2 _dafny.Int) RslMessage {
  return RslMessage{RslMessage_RslMessage__StartingPhase2{Bal__2,LogTruncationPoint__2}}
}

func (_this RslMessage) Is_RslMessage__StartingPhase2() bool {
  _, ok := _this.Get().(RslMessage_RslMessage__StartingPhase2)
return ok
}

func (_this RslMessage) Dtor_seqno__req() _dafny.Int {
  return _this.Get().(RslMessage_RslMessage__Request).Seqno__req
}

func (_this RslMessage) Dtor_val() _50_AppStateMachine__i_Compile.AppMessage_k {
  return _this.Get().(RslMessage_RslMessage__Request).Val
}

func (_this RslMessage) Dtor_bal__1a() _56_LiveRSL____Types__i_Compile.Ballot {
  return _this.Get().(RslMessage_RslMessage__1a).Bal__1a
}

func (_this RslMessage) Dtor_bal__1b() _56_LiveRSL____Types__i_Compile.Ballot {
  return _this.Get().(RslMessage_RslMessage__1b).Bal__1b
}

func (_this RslMessage) Dtor_log__truncation__point() _dafny.Int {
  return _this.Get().(RslMessage_RslMessage__1b).Log__truncation__point
}

func (_this RslMessage) Dtor_votes() _dafny.Map {
  return _this.Get().(RslMessage_RslMessage__1b).Votes
}

func (_this RslMessage) Dtor_bal__2a() _56_LiveRSL____Types__i_Compile.Ballot {
  return _this.Get().(RslMessage_RslMessage__2a).Bal__2a
}

func (_this RslMessage) Dtor_opn__2a() _dafny.Int {
  return _this.Get().(RslMessage_RslMessage__2a).Opn__2a
}

func (_this RslMessage) Dtor_val__2a() _dafny.Seq {
  return _this.Get().(RslMessage_RslMessage__2a).Val__2a
}

func (_this RslMessage) Dtor_bal__2b() _56_LiveRSL____Types__i_Compile.Ballot {
  return _this.Get().(RslMessage_RslMessage__2b).Bal__2b
}

func (_this RslMessage) Dtor_opn__2b() _dafny.Int {
  return _this.Get().(RslMessage_RslMessage__2b).Opn__2b
}

func (_this RslMessage) Dtor_val__2b() _dafny.Seq {
  return _this.Get().(RslMessage_RslMessage__2b).Val__2b
}

func (_this RslMessage) Dtor_bal__heartbeat() _56_LiveRSL____Types__i_Compile.Ballot {
  return _this.Get().(RslMessage_RslMessage__Heartbeat).Bal__heartbeat
}

func (_this RslMessage) Dtor_suspicious() bool {
  return _this.Get().(RslMessage_RslMessage__Heartbeat).Suspicious
}

func (_this RslMessage) Dtor_opn__ckpt() _dafny.Int {
  return _this.Get().(RslMessage_RslMessage__Heartbeat).Opn__ckpt
}

func (_this RslMessage) Dtor_seqno__reply() _dafny.Int {
  return _this.Get().(RslMessage_RslMessage__Reply).Seqno__reply
}

func (_this RslMessage) Dtor_reply() _50_AppStateMachine__i_Compile.AppMessage_k {
  return _this.Get().(RslMessage_RslMessage__Reply).Reply
}

func (_this RslMessage) Dtor_bal__state__req() _56_LiveRSL____Types__i_Compile.Ballot {
  return _this.Get().(RslMessage_RslMessage__AppStateRequest).Bal__state__req
}

func (_this RslMessage) Dtor_opn__state__req() _dafny.Int {
  return _this.Get().(RslMessage_RslMessage__AppStateRequest).Opn__state__req
}

func (_this RslMessage) Dtor_bal__state__supply() _56_LiveRSL____Types__i_Compile.Ballot {
  return _this.Get().(RslMessage_RslMessage__AppStateSupply).Bal__state__supply
}

func (_this RslMessage) Dtor_opn__state__supply() _dafny.Int {
  return _this.Get().(RslMessage_RslMessage__AppStateSupply).Opn__state__supply
}

func (_this RslMessage) Dtor_app__state() uint64 {
  return _this.Get().(RslMessage_RslMessage__AppStateSupply).App__state
}

func (_this RslMessage) Dtor_reply__cache() _dafny.Map {
  return _this.Get().(RslMessage_RslMessage__AppStateSupply).Reply__cache
}

func (_this RslMessage) Dtor_bal__2() _56_LiveRSL____Types__i_Compile.Ballot {
  return _this.Get().(RslMessage_RslMessage__StartingPhase2).Bal__2
}

func (_this RslMessage) Dtor_logTruncationPoint__2() _dafny.Int {
  return _this.Get().(RslMessage_RslMessage__StartingPhase2).LogTruncationPoint__2
}

func (_this RslMessage) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case RslMessage_RslMessage__Invalid: {
      return "_65_LiveRSL____Message__i_Compile.RslMessage.RslMessage_Invalid"
    }
    case RslMessage_RslMessage__Request: {
      return "_65_LiveRSL____Message__i_Compile.RslMessage.RslMessage_Request" + "(" + _dafny.String(data.Seqno__req) + ", " + _dafny.String(data.Val) + ")"
    }
    case RslMessage_RslMessage__1a: {
      return "_65_LiveRSL____Message__i_Compile.RslMessage.RslMessage_1a" + "(" + _dafny.String(data.Bal__1a) + ")"
    }
    case RslMessage_RslMessage__1b: {
      return "_65_LiveRSL____Message__i_Compile.RslMessage.RslMessage_1b" + "(" + _dafny.String(data.Bal__1b) + ", " + _dafny.String(data.Log__truncation__point) + ", " + _dafny.String(data.Votes) + ")"
    }
    case RslMessage_RslMessage__2a: {
      return "_65_LiveRSL____Message__i_Compile.RslMessage.RslMessage_2a" + "(" + _dafny.String(data.Bal__2a) + ", " + _dafny.String(data.Opn__2a) + ", " + _dafny.String(data.Val__2a) + ")"
    }
    case RslMessage_RslMessage__2b: {
      return "_65_LiveRSL____Message__i_Compile.RslMessage.RslMessage_2b" + "(" + _dafny.String(data.Bal__2b) + ", " + _dafny.String(data.Opn__2b) + ", " + _dafny.String(data.Val__2b) + ")"
    }
    case RslMessage_RslMessage__Heartbeat: {
      return "_65_LiveRSL____Message__i_Compile.RslMessage.RslMessage_Heartbeat" + "(" + _dafny.String(data.Bal__heartbeat) + ", " + _dafny.String(data.Suspicious) + ", " + _dafny.String(data.Opn__ckpt) + ")"
    }
    case RslMessage_RslMessage__Reply: {
      return "_65_LiveRSL____Message__i_Compile.RslMessage.RslMessage_Reply" + "(" + _dafny.String(data.Seqno__reply) + ", " + _dafny.String(data.Reply) + ")"
    }
    case RslMessage_RslMessage__AppStateRequest: {
      return "_65_LiveRSL____Message__i_Compile.RslMessage.RslMessage_AppStateRequest" + "(" + _dafny.String(data.Bal__state__req) + ", " + _dafny.String(data.Opn__state__req) + ")"
    }
    case RslMessage_RslMessage__AppStateSupply: {
      return "_65_LiveRSL____Message__i_Compile.RslMessage.RslMessage_AppStateSupply" + "(" + _dafny.String(data.Bal__state__supply) + ", " + _dafny.String(data.Opn__state__supply) + ", " + _dafny.String(data.App__state) + ", " + _dafny.String(data.Reply__cache) + ")"
    }
    case RslMessage_RslMessage__StartingPhase2: {
      return "_65_LiveRSL____Message__i_Compile.RslMessage.RslMessage_StartingPhase2" + "(" + _dafny.String(data.Bal__2) + ", " + _dafny.String(data.LogTruncationPoint__2) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this RslMessage) Equals(other RslMessage) bool {
  switch data1 := _this.Get().(type) {
    case RslMessage_RslMessage__Invalid: {
      _, ok := other.Get().(RslMessage_RslMessage__Invalid)
return ok
    }
    case RslMessage_RslMessage__Request: {
      data2, ok := other.Get().(RslMessage_RslMessage__Request)
return ok && data1.Seqno__req.Cmp(data2.Seqno__req) == 0 && data1.Val.Equals(data2.Val)
    }
    case RslMessage_RslMessage__1a: {
      data2, ok := other.Get().(RslMessage_RslMessage__1a)
return ok && data1.Bal__1a.Equals(data2.Bal__1a)
    }
    case RslMessage_RslMessage__1b: {
      data2, ok := other.Get().(RslMessage_RslMessage__1b)
return ok && data1.Bal__1b.Equals(data2.Bal__1b) && data1.Log__truncation__point.Cmp(data2.Log__truncation__point) == 0 && data1.Votes.Equals(data2.Votes)
    }
    case RslMessage_RslMessage__2a: {
      data2, ok := other.Get().(RslMessage_RslMessage__2a)
return ok && data1.Bal__2a.Equals(data2.Bal__2a) && data1.Opn__2a.Cmp(data2.Opn__2a) == 0 && data1.Val__2a.Equals(data2.Val__2a)
    }
    case RslMessage_RslMessage__2b: {
      data2, ok := other.Get().(RslMessage_RslMessage__2b)
return ok && data1.Bal__2b.Equals(data2.Bal__2b) && data1.Opn__2b.Cmp(data2.Opn__2b) == 0 && data1.Val__2b.Equals(data2.Val__2b)
    }
    case RslMessage_RslMessage__Heartbeat: {
      data2, ok := other.Get().(RslMessage_RslMessage__Heartbeat)
return ok && data1.Bal__heartbeat.Equals(data2.Bal__heartbeat) && data1.Suspicious == data2.Suspicious && data1.Opn__ckpt.Cmp(data2.Opn__ckpt) == 0
    }
    case RslMessage_RslMessage__Reply: {
      data2, ok := other.Get().(RslMessage_RslMessage__Reply)
return ok && data1.Seqno__reply.Cmp(data2.Seqno__reply) == 0 && data1.Reply.Equals(data2.Reply)
    }
    case RslMessage_RslMessage__AppStateRequest: {
      data2, ok := other.Get().(RslMessage_RslMessage__AppStateRequest)
return ok && data1.Bal__state__req.Equals(data2.Bal__state__req) && data1.Opn__state__req.Cmp(data2.Opn__state__req) == 0
    }
    case RslMessage_RslMessage__AppStateSupply: {
      data2, ok := other.Get().(RslMessage_RslMessage__AppStateSupply)
return ok && data1.Bal__state__supply.Equals(data2.Bal__state__supply) && data1.Opn__state__supply.Cmp(data2.Opn__state__supply) == 0 && data1.App__state == data2.App__state && data1.Reply__cache.Equals(data2.Reply__cache)
    }
    case RslMessage_RslMessage__StartingPhase2: {
      data2, ok := other.Get().(RslMessage_RslMessage__StartingPhase2)
return ok && data1.Bal__2.Equals(data2.Bal__2) && data1.LogTruncationPoint__2.Cmp(data2.LogTruncationPoint__2) == 0
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this RslMessage) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(RslMessage)
return ok && _this.Equals(typed)
}
func Type_RslMessage_() _dafny.Type {
  return type_RslMessage_{}
}

type type_RslMessage_ struct {
}

func (_this type_RslMessage_) Default() interface{} {
  return RslMessage{RslMessage_RslMessage__Invalid{}}
}

func (_this type_RslMessage_) String() string {
  return "_65_LiveRSL____Message__i_Compile.RslMessage"
}
// End of data type RslMessage

