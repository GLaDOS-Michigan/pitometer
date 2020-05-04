// Package _56_LiveRSL____Types__i_Compile
// Dafny module _56_LiveRSL____Types__i_Compile compiled into Go

package _56_LiveRSL____Types__i_Compile

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

type Dummy__ struct{}





// Definition of data type Ballot
type Ballot struct {
  Data_Ballot_
}

func (_this Ballot) Get() Data_Ballot_ {
  return _this.Data_Ballot_
}

type Data_Ballot_ interface {
  isBallot()
}

type CompanionStruct_Ballot_ struct {}
var Companion_Ballot_ = CompanionStruct_Ballot_{}

type Ballot_Ballot struct {
  Seqno _dafny.Int
Proposer__id _dafny.Int
}

func (Ballot_Ballot) isBallot() {}

func (CompanionStruct_Ballot_) Create_Ballot_(Seqno _dafny.Int, Proposer__id _dafny.Int) Ballot {
  return Ballot{Ballot_Ballot{Seqno,Proposer__id}}
}

func (_this Ballot) Is_Ballot() bool {
  _, ok := _this.Get().(Ballot_Ballot)
return ok
}

func (_this Ballot) Dtor_seqno() _dafny.Int {
  return _this.Get().(Ballot_Ballot).Seqno
}

func (_this Ballot) Dtor_proposer__id() _dafny.Int {
  return _this.Get().(Ballot_Ballot).Proposer__id
}

func (_this Ballot) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case Ballot_Ballot: {
      return "_56_LiveRSL____Types__i_Compile.Ballot.Ballot" + "(" + _dafny.String(data.Seqno) + ", " + _dafny.String(data.Proposer__id) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this Ballot) Equals(other Ballot) bool {
  switch data1 := _this.Get().(type) {
    case Ballot_Ballot: {
      data2, ok := other.Get().(Ballot_Ballot)
return ok && data1.Seqno.Cmp(data2.Seqno) == 0 && data1.Proposer__id.Cmp(data2.Proposer__id) == 0
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this Ballot) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(Ballot)
return ok && _this.Equals(typed)
}
func Type_Ballot_() _dafny.Type {
  return type_Ballot_{}
}

type type_Ballot_ struct {
}

func (_this type_Ballot_) Default() interface{} {
  return Ballot{Ballot_Ballot{_dafny.Zero, _dafny.Zero}}
}

func (_this type_Ballot_) String() string {
  return "_56_LiveRSL____Types__i_Compile.Ballot"
}
// End of data type Ballot

// Definition of data type Request
type Request struct {
  Data_Request_
}

func (_this Request) Get() Data_Request_ {
  return _this.Data_Request_
}

type Data_Request_ interface {
  isRequest()
}

type CompanionStruct_Request_ struct {}
var Companion_Request_ = CompanionStruct_Request_{}

type Request_Request struct {
  Client _9_Native____Io__s_Compile.EndPoint
Seqno _dafny.Int
Request _50_AppStateMachine__i_Compile.AppMessage_k
}

func (Request_Request) isRequest() {}

func (CompanionStruct_Request_) Create_Request_(Client _9_Native____Io__s_Compile.EndPoint, Seqno _dafny.Int, Request _50_AppStateMachine__i_Compile.AppMessage_k) Request {
  return Request{Request_Request{Client,Seqno,Request}}
}

func (_this Request) Is_Request() bool {
  _, ok := _this.Get().(Request_Request)
return ok
}

func (_this Request) Dtor_client() _9_Native____Io__s_Compile.EndPoint {
  return _this.Get().(Request_Request).Client
}

func (_this Request) Dtor_seqno() _dafny.Int {
  return _this.Get().(Request_Request).Seqno
}

func (_this Request) Dtor_request() _50_AppStateMachine__i_Compile.AppMessage_k {
  return _this.Get().(Request_Request).Request
}

func (_this Request) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case Request_Request: {
      return "_56_LiveRSL____Types__i_Compile.Request.Request" + "(" + _dafny.String(data.Client) + ", " + _dafny.String(data.Seqno) + ", " + _dafny.String(data.Request) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this Request) Equals(other Request) bool {
  switch data1 := _this.Get().(type) {
    case Request_Request: {
      data2, ok := other.Get().(Request_Request)
return ok && data1.Client.Equals(data2.Client) && data1.Seqno.Cmp(data2.Seqno) == 0 && data1.Request.Equals(data2.Request)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this Request) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(Request)
return ok && _this.Equals(typed)
}
func Type_Request_() _dafny.Type {
  return type_Request_{}
}

type type_Request_ struct {
}

func (_this type_Request_) Default() interface{} {
  return Request{Request_Request{_9_Native____Io__s_Compile.Type_EndPoint_().Default().(_9_Native____Io__s_Compile.EndPoint), _dafny.Zero, _50_AppStateMachine__i_Compile.Type_AppMessage_k_().Default().(_50_AppStateMachine__i_Compile.AppMessage_k)}}
}

func (_this type_Request_) String() string {
  return "_56_LiveRSL____Types__i_Compile.Request"
}
// End of data type Request

// Definition of data type Reply
type Reply struct {
  Data_Reply_
}

func (_this Reply) Get() Data_Reply_ {
  return _this.Data_Reply_
}

type Data_Reply_ interface {
  isReply()
}

type CompanionStruct_Reply_ struct {}
var Companion_Reply_ = CompanionStruct_Reply_{}

type Reply_Reply struct {
  Client _9_Native____Io__s_Compile.EndPoint
Seqno _dafny.Int
Reply _50_AppStateMachine__i_Compile.AppMessage_k
}

func (Reply_Reply) isReply() {}

func (CompanionStruct_Reply_) Create_Reply_(Client _9_Native____Io__s_Compile.EndPoint, Seqno _dafny.Int, Reply _50_AppStateMachine__i_Compile.AppMessage_k) Reply {
  return Reply{Reply_Reply{Client,Seqno,Reply}}
}

func (_this Reply) Is_Reply() bool {
  _, ok := _this.Get().(Reply_Reply)
return ok
}

func (_this Reply) Dtor_client() _9_Native____Io__s_Compile.EndPoint {
  return _this.Get().(Reply_Reply).Client
}

func (_this Reply) Dtor_seqno() _dafny.Int {
  return _this.Get().(Reply_Reply).Seqno
}

func (_this Reply) Dtor_reply() _50_AppStateMachine__i_Compile.AppMessage_k {
  return _this.Get().(Reply_Reply).Reply
}

func (_this Reply) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case Reply_Reply: {
      return "_56_LiveRSL____Types__i_Compile.Reply.Reply" + "(" + _dafny.String(data.Client) + ", " + _dafny.String(data.Seqno) + ", " + _dafny.String(data.Reply) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this Reply) Equals(other Reply) bool {
  switch data1 := _this.Get().(type) {
    case Reply_Reply: {
      data2, ok := other.Get().(Reply_Reply)
return ok && data1.Client.Equals(data2.Client) && data1.Seqno.Cmp(data2.Seqno) == 0 && data1.Reply.Equals(data2.Reply)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this Reply) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(Reply)
return ok && _this.Equals(typed)
}
func Type_Reply_() _dafny.Type {
  return type_Reply_{}
}

type type_Reply_ struct {
}

func (_this type_Reply_) Default() interface{} {
  return Reply{Reply_Reply{_9_Native____Io__s_Compile.Type_EndPoint_().Default().(_9_Native____Io__s_Compile.EndPoint), _dafny.Zero, _50_AppStateMachine__i_Compile.Type_AppMessage_k_().Default().(_50_AppStateMachine__i_Compile.AppMessage_k)}}
}

func (_this type_Reply_) String() string {
  return "_56_LiveRSL____Types__i_Compile.Reply"
}
// End of data type Reply



// Definition of data type Vote
type Vote struct {
  Data_Vote_
}

func (_this Vote) Get() Data_Vote_ {
  return _this.Data_Vote_
}

type Data_Vote_ interface {
  isVote()
}

type CompanionStruct_Vote_ struct {}
var Companion_Vote_ = CompanionStruct_Vote_{}

type Vote_Vote struct {
  Max__value__bal _56_LiveRSL____Types__i_Compile.Ballot
}

func (Vote_Vote) isVote() {}

func (CompanionStruct_Vote_) Create_Vote_(Max__value__bal _56_LiveRSL____Types__i_Compile.Ballot) Vote {
  return Vote{Vote_Vote{Max__value__bal}}
}

func (_this Vote) Is_Vote() bool {
  _, ok := _this.Get().(Vote_Vote)
return ok
}

func (_this Vote) Dtor_max__value__bal() _56_LiveRSL____Types__i_Compile.Ballot {
  return _this.Get().(Vote_Vote).Max__value__bal
}

func (_this Vote) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case Vote_Vote: {
      return "_56_LiveRSL____Types__i_Compile.Vote.Vote" + "(" + _dafny.String(data.Max__value__bal) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this Vote) Equals(other Vote) bool {
  switch data1 := _this.Get().(type) {
    case Vote_Vote: {
      data2, ok := other.Get().(Vote_Vote)
return ok && data1.Max__value__bal.Equals(data2.Max__value__bal)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this Vote) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(Vote)
return ok && _this.Equals(typed)
}
func Type_Vote_() _dafny.Type {
  return type_Vote_{}
}

type type_Vote_ struct {
}

func (_this type_Vote_) Default() interface{} {
  return Vote{Vote_Vote{Type_Ballot_().Default().(_56_LiveRSL____Types__i_Compile.Ballot)}}
}

func (_this type_Vote_) String() string {
  return "_56_LiveRSL____Types__i_Compile.Vote"
}
// End of data type Vote


// Definition of data type LearnerTuple
type LearnerTuple struct {
  Data_LearnerTuple_
}

func (_this LearnerTuple) Get() Data_LearnerTuple_ {
  return _this.Data_LearnerTuple_
}

type Data_LearnerTuple_ interface {
  isLearnerTuple()
}

type CompanionStruct_LearnerTuple_ struct {}
var Companion_LearnerTuple_ = CompanionStruct_LearnerTuple_{}

type LearnerTuple_LearnerTuple struct {
  Received__2b__message__senders _dafny.Set
Candidate__learned__value _dafny.Seq
}

func (LearnerTuple_LearnerTuple) isLearnerTuple() {}

func (CompanionStruct_LearnerTuple_) Create_LearnerTuple_(Received__2b__message__senders _dafny.Set, Candidate__learned__value _dafny.Seq) LearnerTuple {
  return LearnerTuple{LearnerTuple_LearnerTuple{Received__2b__message__senders,Candidate__learned__value}}
}

func (_this LearnerTuple) Is_LearnerTuple() bool {
  _, ok := _this.Get().(LearnerTuple_LearnerTuple)
return ok
}

func (_this LearnerTuple) Dtor_received__2b__message__senders() _dafny.Set {
  return _this.Get().(LearnerTuple_LearnerTuple).Received__2b__message__senders
}

func (_this LearnerTuple) Dtor_candidate__learned__value() _dafny.Seq {
  return _this.Get().(LearnerTuple_LearnerTuple).Candidate__learned__value
}

func (_this LearnerTuple) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case LearnerTuple_LearnerTuple: {
      return "_56_LiveRSL____Types__i_Compile.LearnerTuple.LearnerTuple" + "(" + _dafny.String(data.Received__2b__message__senders) + ", " + _dafny.String(data.Candidate__learned__value) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this LearnerTuple) Equals(other LearnerTuple) bool {
  switch data1 := _this.Get().(type) {
    case LearnerTuple_LearnerTuple: {
      data2, ok := other.Get().(LearnerTuple_LearnerTuple)
return ok && data1.Received__2b__message__senders.Equals(data2.Received__2b__message__senders) && data1.Candidate__learned__value.Equals(data2.Candidate__learned__value)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this LearnerTuple) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(LearnerTuple)
return ok && _this.Equals(typed)
}
func Type_LearnerTuple_() _dafny.Type {
  return type_LearnerTuple_{}
}

type type_LearnerTuple_ struct {
}

func (_this type_LearnerTuple_) Default() interface{} {
  return LearnerTuple{LearnerTuple_LearnerTuple{_dafny.EmptySet, _dafny.EmptySeq}}
}

func (_this type_LearnerTuple_) String() string {
  return "_56_LiveRSL____Types__i_Compile.LearnerTuple"
}
// End of data type LearnerTuple


