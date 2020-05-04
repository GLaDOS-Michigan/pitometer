// Package _214_LiveRSL____CTypes__i_Compile
// Dafny module _214_LiveRSL____CTypes__i_Compile compiled into Go

package _214_LiveRSL____CTypes__i_Compile

import (
	_0_Native____NativeTypes__s_Compile "0_Native____NativeTypes__s_Compile_"
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
	_26_Collections____Seqs__s_Compile "26_Collections____Seqs__s_Compile_"
	_2_Collections____Maps2__s_Compile "2_Collections____Maps2__s_Compile_"
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
	_5_Temporal____Temporal__s_Compile "5_Temporal____Temporal__s_Compile_"
	_61_LiveRSL____Configuration__i_Compile "61_LiveRSL____Configuration__i_Compile_"
	_65_LiveRSL____Message__i_Compile "65_LiveRSL____Message__i_Compile_"
	_68_LiveRSL____Environment__i_Compile "68_LiveRSL____Environment__i_Compile_"
	_71_LiveRSL____ClockReading__i_Compile "71_LiveRSL____ClockReading__i_Compile_"
	_74_Common____UpperBound__s_Compile "74_Common____UpperBound__s_Compile_"
	_76_LiveRSL____Parameters__i_Compile "76_LiveRSL____Parameters__i_Compile_"
	_78_LiveRSL____Constants__i_Compile "78_LiveRSL____Constants__i_Compile_"
	_7_Environment__s_Compile "7_Environment__s_Compile_"
	_85_LiveRSL____Broadcast__i_Compile "85_LiveRSL____Broadcast__i_Compile_"
	_91_Collections____CountMatches__i_Compile "91_Collections____CountMatches__i_Compile_"
	_93_LiveRSL____Acceptor__i_Compile "93_LiveRSL____Acceptor__i_Compile_"
	_99_LiveRSL____Election__i_Compile "99_LiveRSL____Election__i_Compile_"
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

type Dummy__ struct{}

// Definition of data type CBallot
type CBallot struct {
	Data_CBallot_
}

func (_this CBallot) Get() Data_CBallot_ {
	return _this.Data_CBallot_
}

type Data_CBallot_ interface {
	isCBallot()
}

type CompanionStruct_CBallot_ struct{}

var Companion_CBallot_ = CompanionStruct_CBallot_{}

type CBallot_CBallot struct {
	Seqno        uint64
	Proposer__id uint64
}

func (CBallot_CBallot) isCBallot() {}

func (CompanionStruct_CBallot_) Create_CBallot_(Seqno uint64, Proposer__id uint64) CBallot {
	return CBallot{CBallot_CBallot{Seqno, Proposer__id}}
}

func (_this CBallot) Is_CBallot() bool {
	_, ok := _this.Get().(CBallot_CBallot)
	return ok
}

func (_this CBallot) Dtor_seqno() uint64 {
	return _this.Get().(CBallot_CBallot).Seqno
}

func (_this CBallot) Dtor_proposer__id() uint64 {
	return _this.Get().(CBallot_CBallot).Proposer__id
}

func (_this CBallot) String() string {
	switch data := _this.Get().(type) {
	case nil:
		return "null"
	case CBallot_CBallot:
		{
			return "CBallot.CBallot" + "(" + _dafny.String(data.Seqno) + ", " + _dafny.String(data.Proposer__id) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this CBallot) Equals(other CBallot) bool {
	switch data1 := _this.Get().(type) {
	case CBallot_CBallot:
		{
			data2, ok := other.Get().(CBallot_CBallot)
			return ok && data1.Seqno == data2.Seqno && data1.Proposer__id == data2.Proposer__id
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this CBallot) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(CBallot)
	return ok && _this.Equals(typed)
}
func Type_CBallot_() _dafny.Type {
	return type_CBallot_{}
}

type type_CBallot_ struct {
}

func (_this type_CBallot_) Default() interface{} {
	return CBallot{CBallot_CBallot{0, 0}}
}

func (_this type_CBallot_) String() string {
	return "CBallot"
}

// End of data type CBallot

// Definition of data type COperationNumber
type COperationNumber struct {
	Data_COperationNumber_
}

func (_this COperationNumber) Get() Data_COperationNumber_ {
	return _this.Data_COperationNumber_
}

type Data_COperationNumber_ interface {
	isCOperationNumber()
}

type CompanionStruct_COperationNumber_ struct{}

var Companion_COperationNumber_ = CompanionStruct_COperationNumber_{}

type COperationNumber_COperationNumber struct {
	N uint64
}

func (COperationNumber_COperationNumber) isCOperationNumber() {}

func (CompanionStruct_COperationNumber_) Create_COperationNumber_(N uint64) COperationNumber {
	return COperationNumber{COperationNumber_COperationNumber{N}}
}

func (_this COperationNumber) Is_COperationNumber() bool {
	_, ok := _this.Get().(COperationNumber_COperationNumber)
	return ok
}

func (_this COperationNumber) Dtor_n() uint64 {
	return _this.Get().(COperationNumber_COperationNumber).N
}

func (_this COperationNumber) String() string {
	switch data := _this.Get().(type) {
	case nil:
		return "null"
	case COperationNumber_COperationNumber:
		{
			return "COperationNumber.COperationNumber" + "(" + _dafny.String(data.N) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this COperationNumber) Equals(other COperationNumber) bool {
	switch data1 := _this.Get().(type) {
	case COperationNumber_COperationNumber:
		{
			data2, ok := other.Get().(COperationNumber_COperationNumber)
			return ok && data1.N == data2.N
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this COperationNumber) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(COperationNumber)
	return ok && _this.Equals(typed)
}
func Type_COperationNumber_() _dafny.Type {
	return type_COperationNumber_{}
}

type type_COperationNumber_ struct {
}

func (_this type_COperationNumber_) Default() interface{} {
	return COperationNumber{COperationNumber_COperationNumber{0}}
}

func (_this type_COperationNumber_) String() string {
	return "COperationNumber"
}

// End of data type COperationNumber

// Definition of data type CRequest
type CRequest struct {
	Data_CRequest_
}

func (_this CRequest) Get() Data_CRequest_ {
	return _this.Data_CRequest_
}

type Data_CRequest_ interface {
	isCRequest()
}

type CompanionStruct_CRequest_ struct{}

var Companion_CRequest_ = CompanionStruct_CRequest_{}

type CRequest_CRequest struct {
	Client  _9_Native____Io__s_Compile.EndPoint
	Seqno   uint64
	Request _197_LiveRSL____AppInterface__i_Compile.CAppMessage
}

func (CRequest_CRequest) isCRequest() {}

func (CompanionStruct_CRequest_) Create_CRequest_(Client _9_Native____Io__s_Compile.EndPoint, Seqno uint64, Request _197_LiveRSL____AppInterface__i_Compile.CAppMessage) CRequest {
	return CRequest{CRequest_CRequest{Client, Seqno, Request}}
}

func (_this CRequest) Is_CRequest() bool {
	_, ok := _this.Get().(CRequest_CRequest)
	return ok
}

func (_this CRequest) Dtor_client() _9_Native____Io__s_Compile.EndPoint {
	return _this.Get().(CRequest_CRequest).Client
}

func (_this CRequest) Dtor_seqno() uint64 {
	return _this.Get().(CRequest_CRequest).Seqno
}

func (_this CRequest) Dtor_request() _197_LiveRSL____AppInterface__i_Compile.CAppMessage {
	return _this.Get().(CRequest_CRequest).Request
}

func (_this CRequest) String() string {
	switch data := _this.Get().(type) {
	case nil:
		return "null"
	case CRequest_CRequest:
		{
			return "CRequest.CRequest" + "(" + _dafny.String(data.Client) + ", " + _dafny.String(data.Seqno) + ", " + _dafny.String(data.Request) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this CRequest) Equals(other CRequest) bool {
	switch data1 := _this.Get().(type) {
	case CRequest_CRequest:
		{
			data2, ok := other.Get().(CRequest_CRequest)
			return ok && data1.Client.Equals(data2.Client) && data1.Seqno == data2.Seqno && data1.Request.Equals(data2.Request)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this CRequest) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(CRequest)
	return ok && _this.Equals(typed)
}
func Type_CRequest_() _dafny.Type {
	return type_CRequest_{}
}

type type_CRequest_ struct {
}

func (_this type_CRequest_) Default() interface{} {
	return CRequest{CRequest_CRequest{_9_Native____Io__s_Compile.Type_EndPoint_().Default().(_9_Native____Io__s_Compile.EndPoint), 0, _197_LiveRSL____AppInterface__i_Compile.Type_CAppMessage_().Default().(_197_LiveRSL____AppInterface__i_Compile.CAppMessage)}}
}

func (_this type_CRequest_) String() string {
	return "CRequest"
}

// End of data type CRequest

// Definition of data type CReply
type CReply struct {
	Data_CReply_
}

func (_this CReply) Get() Data_CReply_ {
	return _this.Data_CReply_
}

type Data_CReply_ interface {
	isCReply()
}

type CompanionStruct_CReply_ struct{}

var Companion_CReply_ = CompanionStruct_CReply_{}

type CReply_CReply struct {
	Client _9_Native____Io__s_Compile.EndPoint
	Seqno  uint64
	Reply  _197_LiveRSL____AppInterface__i_Compile.CAppMessage
}

func (CReply_CReply) isCReply() {}

func (CompanionStruct_CReply_) Create_CReply_(Client _9_Native____Io__s_Compile.EndPoint, Seqno uint64, Reply _197_LiveRSL____AppInterface__i_Compile.CAppMessage) CReply {
	return CReply{CReply_CReply{Client, Seqno, Reply}}
}

func (_this CReply) Is_CReply() bool {
	_, ok := _this.Get().(CReply_CReply)
	return ok
}

func (_this CReply) Dtor_client() _9_Native____Io__s_Compile.EndPoint {
	return _this.Get().(CReply_CReply).Client
}

func (_this CReply) Dtor_seqno() uint64 {
	return _this.Get().(CReply_CReply).Seqno
}

func (_this CReply) Dtor_reply() _197_LiveRSL____AppInterface__i_Compile.CAppMessage {
	return _this.Get().(CReply_CReply).Reply
}

func (_this CReply) String() string {
	switch data := _this.Get().(type) {
	case nil:
		return "null"
	case CReply_CReply:
		{
			return "CReply.CReply" + "(" + _dafny.String(data.Client) + ", " + _dafny.String(data.Seqno) + ", " + _dafny.String(data.Reply) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this CReply) Equals(other CReply) bool {
	switch data1 := _this.Get().(type) {
	case CReply_CReply:
		{
			data2, ok := other.Get().(CReply_CReply)
			return ok && data1.Client.Equals(data2.Client) && data1.Seqno == data2.Seqno && data1.Reply.Equals(data2.Reply)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this CReply) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(CReply)
	return ok && _this.Equals(typed)
}
func Type_CReply_() _dafny.Type {
	return type_CReply_{}
}

type type_CReply_ struct {
}

func (_this type_CReply_) Default() interface{} {
	return CReply{CReply_CReply{_9_Native____Io__s_Compile.Type_EndPoint_().Default().(_9_Native____Io__s_Compile.EndPoint), 0, _197_LiveRSL____AppInterface__i_Compile.Type_CAppMessage_().Default().(_197_LiveRSL____AppInterface__i_Compile.CAppMessage)}}
}

func (_this type_CReply_) String() string {
	return "CReply"
}

// End of data type CReply

// Definition of data type CVote
type CVote struct {
	Data_CVote_
}

func (_this CVote) Get() Data_CVote_ {
	return _this.Data_CVote_
}

type Data_CVote_ interface {
	isCVote()
}

type CompanionStruct_CVote_ struct{}

var Companion_CVote_ = CompanionStruct_CVote_{}

type CVote_CVote struct {
	Max__value__bal CBallot
	Max__val        _dafny.Seq
}

func (CVote_CVote) isCVote() {}

func (CompanionStruct_CVote_) Create_CVote_(Max__value__bal CBallot, Max__val _dafny.Seq) CVote {
	return CVote{CVote_CVote{Max__value__bal, Max__val}}
}

func (_this CVote) Is_CVote() bool {
	_, ok := _this.Get().(CVote_CVote)
	return ok
}

func (_this CVote) Dtor_max__value__bal() CBallot {
	return _this.Get().(CVote_CVote).Max__value__bal
}

func (_this CVote) Dtor_max__val() _dafny.Seq {
	return _this.Get().(CVote_CVote).Max__val
}

func (_this CVote) String() string {
	switch data := _this.Get().(type) {
	case nil:
		return "null"
	case CVote_CVote:
		{
			return "CVote.CVote" + "(" + _dafny.String(data.Max__value__bal) + ", " + _dafny.String(data.Max__val) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this CVote) Equals(other CVote) bool {
	switch data1 := _this.Get().(type) {
	case CVote_CVote:
		{
			data2, ok := other.Get().(CVote_CVote)
			return ok && data1.Max__value__bal.Equals(data2.Max__value__bal) && data1.Max__val.Equals(data2.Max__val)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this CVote) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(CVote)
	return ok && _this.Equals(typed)
}
func Type_CVote_() _dafny.Type {
	return type_CVote_{}
}

type type_CVote_ struct {
}

func (_this type_CVote_) Default() interface{} {
	return CVote{CVote_CVote{Type_CBallot_().Default().(CBallot), _dafny.EmptySeq}}
}

func (_this type_CVote_) String() string {
	return "CVote"
}

// End of data type CVote

// Definition of data type CVotes
type CVotes struct {
	Data_CVotes_
}

func (_this CVotes) Get() Data_CVotes_ {
	return _this.Data_CVotes_
}

type Data_CVotes_ interface {
	isCVotes()
}

type CompanionStruct_CVotes_ struct{}

var Companion_CVotes_ = CompanionStruct_CVotes_{}

type CVotes_CVotes struct {
	V _dafny.Map
}

func (CVotes_CVotes) isCVotes() {}

func (CompanionStruct_CVotes_) Create_CVotes_(V _dafny.Map) CVotes {
	return CVotes{CVotes_CVotes{V}}
}

func (_this CVotes) Is_CVotes() bool {
	_, ok := _this.Get().(CVotes_CVotes)
	return ok
}

func (_this CVotes) Dtor_v() _dafny.Map {
	return _this.Get().(CVotes_CVotes).V
}

func (_this CVotes) String() string {
	switch data := _this.Get().(type) {
	case nil:
		return "null"
	case CVotes_CVotes:
		{
			return "CVotes.CVotes" + "(" + _dafny.String(data.V) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this CVotes) Equals(other CVotes) bool {
	switch data1 := _this.Get().(type) {
	case CVotes_CVotes:
		{
			data2, ok := other.Get().(CVotes_CVotes)
			return ok && data1.V.Equals(data2.V)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this CVotes) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(CVotes)
	return ok && _this.Equals(typed)
}
func Type_CVotes_() _dafny.Type {
	return type_CVotes_{}
}

type type_CVotes_ struct {
}

func (_this type_CVotes_) Default() interface{} {
	return CVotes{CVotes_CVotes{_dafny.EmptyMap}}
}

func (_this type_CVotes_) String() string {
	return "CVotes"
}

// End of data type CVotes

// Definition of data type OptionOpn
type OptionOpn struct {
	Data_OptionOpn_
}

func (_this OptionOpn) Get() Data_OptionOpn_ {
	return _this.Data_OptionOpn_
}

type Data_OptionOpn_ interface {
	isOptionOpn()
}

type CompanionStruct_OptionOpn_ struct{}

var Companion_OptionOpn_ = CompanionStruct_OptionOpn_{}

type OptionOpn_ExistsOperation struct {
	Opn COperationNumber
}

func (OptionOpn_ExistsOperation) isOptionOpn() {}

func (CompanionStruct_OptionOpn_) Create_ExistsOperation_(Opn COperationNumber) OptionOpn {
	return OptionOpn{OptionOpn_ExistsOperation{Opn}}
}

func (_this OptionOpn) Is_ExistsOperation() bool {
	_, ok := _this.Get().(OptionOpn_ExistsOperation)
	return ok
}

type OptionOpn_NotExistsOperation struct {
}

func (OptionOpn_NotExistsOperation) isOptionOpn() {}

func (CompanionStruct_OptionOpn_) Create_NotExistsOperation_() OptionOpn {
	return OptionOpn{OptionOpn_NotExistsOperation{}}
}

func (_this OptionOpn) Is_NotExistsOperation() bool {
	_, ok := _this.Get().(OptionOpn_NotExistsOperation)
	return ok
}

func (_this OptionOpn) Dtor_opn() COperationNumber {
	return _this.Get().(OptionOpn_ExistsOperation).Opn
}

func (_this OptionOpn) String() string {
	switch data := _this.Get().(type) {
	case nil:
		return "null"
	case OptionOpn_ExistsOperation:
		{
			return "OptionOpn.ExistsOperation" + "(" + _dafny.String(data.Opn) + ")"
		}
	case OptionOpn_NotExistsOperation:
		{
			return "OptionOpn.NotExistsOperation"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this OptionOpn) Equals(other OptionOpn) bool {
	switch data1 := _this.Get().(type) {
	case OptionOpn_ExistsOperation:
		{
			data2, ok := other.Get().(OptionOpn_ExistsOperation)
			return ok && data1.Opn.Equals(data2.Opn)
		}
	case OptionOpn_NotExistsOperation:
		{
			_, ok := other.Get().(OptionOpn_NotExistsOperation)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this OptionOpn) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(OptionOpn)
	return ok && _this.Equals(typed)
}
func Type_OptionOpn_() _dafny.Type {
	return type_OptionOpn_{}
}

type type_OptionOpn_ struct {
}

func (_this type_OptionOpn_) Default() interface{} {
	return OptionOpn{OptionOpn_ExistsOperation{Type_COperationNumber_().Default().(COperationNumber)}}
}

func (_this type_OptionOpn_) String() string {
	return "OptionOpn"
}

// End of data type OptionOpn

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

var Companion_Default___ = CompanionStruct_Default___{}

func (_this *Default__) Equals(other *Default__) bool {
	return _this == other
}

func (_this *Default__) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*Default__)
	return ok && _this.Equals(other)
}

func (*Default__) String() string {
	return "Default__"
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
	return "Default__"
}
func (_this *CompanionStruct_Default___) BallotSize() uint64 {
	return (_138_Native____NativeTypes__i_Compile.Companion_Default___.Uint64Size()) + (_138_Native____NativeTypes__i_Compile.Companion_Default___.Uint64Size())
}
func (_this *CompanionStruct_Default___) CBallotIsLessThan(lhs CBallot, rhs CBallot) bool {
	return (((lhs).Dtor_seqno()) < ((rhs).Dtor_seqno())) || ((((lhs).Dtor_seqno()) == ((rhs).Dtor_seqno())) && (((lhs).Dtor_proposer__id()) < ((rhs).Dtor_proposer__id())))
}
func (_this *CompanionStruct_Default___) CBallotIsNotGreaterThan(lhs CBallot, rhs CBallot) bool {
	return (((lhs).Dtor_seqno()) < ((rhs).Dtor_seqno())) || ((((lhs).Dtor_seqno()) == ((rhs).Dtor_seqno())) && (((lhs).Dtor_proposer__id()) <= ((rhs).Dtor_proposer__id())))
}
func (_this *CompanionStruct_Default___) CBalLeq(ba CBallot, bb CBallot) bool {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var b bool = false
	var _ = b
	if (((ba).Dtor_seqno()) < ((bb).Dtor_seqno())) || ((((ba).Dtor_seqno()) == ((bb).Dtor_seqno())) && (((ba).Dtor_proposer__id()) <= ((bb).Dtor_proposer__id()))) {
		b = true
	} else {
		b = false
	}
	return b
}
func (_this *CompanionStruct_Default___) CBalLt(ba CBallot, bb CBallot) bool {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var b bool = false
	var _ = b
	if (((ba).Dtor_seqno()) < ((bb).Dtor_seqno())) || ((((ba).Dtor_seqno()) == ((bb).Dtor_seqno())) && (((ba).Dtor_proposer__id()) < ((bb).Dtor_proposer__id()))) {
		b = true
	} else {
		b = false
	}
	return b
}
func (_this *CompanionStruct_Default___) OpNumSize() uint64 {
	return _138_Native____NativeTypes__i_Compile.Companion_Default___.Uint64Size()
}
func (_this *CompanionStruct_Default___) ValidRequest(c CRequest) bool {
	return !((c).Is_CRequest()) || ((_180_Common____UdpClient__i_Compile.Companion_Default___.EndPointIsValidIPV4((c).Dtor_client())) && (_197_LiveRSL____AppInterface__i_Compile.Companion_Default___.ValidAppMessage((c).Dtor_request())))
}
func (_this *CompanionStruct_Default___) RequestBatchSizeLimit() _dafny.Int {
	return _dafny.IntOfInt64(100)
}
func (_this *CompanionStruct_Default___) ValidReply(c CReply) bool {
	return !((c).Is_CReply()) || ((_180_Common____UdpClient__i_Compile.Companion_Default___.EndPointIsValidIPV4((c).Dtor_client())) && (_197_LiveRSL____AppInterface__i_Compile.Companion_Default___.ValidAppMessage((c).Dtor_reply())))
}
func (_this *CompanionStruct_Default___) Max__votes__len() _dafny.Int {
	return _dafny.IntOfInt64(8)
}

// End of class Default__
