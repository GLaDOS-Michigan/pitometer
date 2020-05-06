// Package _368_LiveRSL____UdpRSL__i_Compile
// Dafny module _368_LiveRSL____UdpRSL__i_Compile compiled into Go

package _368_LiveRSL____UdpRSL__i_Compile

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
	_214_LiveRSL____CTypes__i_Compile "214_LiveRSL____CTypes__i_Compile_"
	_217_LiveRSL____CMessage__i_Compile "217_LiveRSL____CMessage__i_Compile_"
	_225_LiveRSL____CMessageRefinements__i_Compile "225_LiveRSL____CMessageRefinements__i_Compile_"
	_228_LiveRSL____PacketParsing__i_Compile "228_LiveRSL____PacketParsing__i_Compile_"
	_234_Common____UpperBound__i_Compile "234_Common____UpperBound__i_Compile_"
	_236_LiveRSL____ParametersState__i_Compile "236_LiveRSL____ParametersState__i_Compile_"
	_238_LiveRSL____CPaxosConfiguration__i_Compile "238_LiveRSL____CPaxosConfiguration__i_Compile_"
	_240_LiveRSL____ConstantsState__i_Compile "240_LiveRSL____ConstantsState__i_Compile_"
	_243_LiveRSL____PaxosWorldState__i_Compile "243_LiveRSL____PaxosWorldState__i_Compile_"
	_245_LiveRSL____ReplicaConstantsState__i_Compile "245_LiveRSL____ReplicaConstantsState__i_Compile_"
	_251_LiveRSL____ElectionState__i_Compile "251_LiveRSL____ElectionState__i_Compile_"
	_254_LiveRSL____ProposerState__i_Compile "254_LiveRSL____ProposerState__i_Compile_"
	_265_LiveRSL____COperationNumberSort__i_Compile "265_LiveRSL____COperationNumberSort__i_Compile_"
	_267_LiveRSL____CLastCheckpointedMap__i_Compile "267_LiveRSL____CLastCheckpointedMap__i_Compile_"
	_269_LiveRSL____AcceptorState__i_Compile "269_LiveRSL____AcceptorState__i_Compile_"
	_26_Collections____Seqs__s_Compile "26_Collections____Seqs__s_Compile_"
	_276_LiveRSL____ExecutorState__i_Compile "276_LiveRSL____ExecutorState__i_Compile_"
	_278_LiveRSL____LearnerState__i_Compile "278_LiveRSL____LearnerState__i_Compile_"
	_283_LiveRSL____CClockReading__i_Compile "283_LiveRSL____CClockReading__i_Compile_"
	_285_LiveRSL____ReplicaState__i_Compile "285_LiveRSL____ReplicaState__i_Compile_"
	_294_LiveRSL____MinCQuorumSize__i_Compile "294_LiveRSL____MinCQuorumSize__i_Compile_"
	_297_LiveRSL____ElectionModel__i_Compile "297_LiveRSL____ElectionModel__i_Compile_"
	_2_Collections____Maps2__s_Compile "2_Collections____Maps2__s_Compile_"
	_301_Impl____LiveRSL____Broadcast__i_Compile "301_Impl____LiveRSL____Broadcast__i_Compile_"
	_305_LiveRSL____ProposerLemmas__i_Compile "305_LiveRSL____ProposerLemmas__i_Compile_"
	_308_LiveRSL____ProposerModel__i_Compile "308_LiveRSL____ProposerModel__i_Compile_"
	_30_Collections____Seqs__i_Compile "30_Collections____Seqs__i_Compile_"
	_312_LiveRSL____AcceptorModel__i_Compile "312_LiveRSL____AcceptorModel__i_Compile_"
	_316_LiveRSL____LearnerModel__i_Compile "316_LiveRSL____LearnerModel__i_Compile_"
	_321_LiveRSL____ExecutorModel__i_Compile "321_LiveRSL____ExecutorModel__i_Compile_"
	_323_LiveRSL____ReplicaModel__Part1__i_Compile "323_LiveRSL____ReplicaModel__Part1__i_Compile_"
	_32_Math____mul__nonlinear__i_Compile "32_Math____mul__nonlinear__i_Compile_"
	_331_LiveRSL____ReplicaModel__Part2__i_Compile "331_LiveRSL____ReplicaModel__Part2__i_Compile_"
	_339_LiveRSL____ReplicaModel__Part3__i_Compile "339_LiveRSL____ReplicaModel__Part3__i_Compile_"
	_347_LiveRSL____ReplicaModel__Part4__i_Compile "347_LiveRSL____ReplicaModel__Part4__i_Compile_"
	_34_Math____mul__auto__proofs__i_Compile "34_Math____mul__auto__proofs__i_Compile_"
	_355_LiveRSL____ReplicaModel__Part5__i_Compile "355_LiveRSL____ReplicaModel__Part5__i_Compile_"
	_357_LiveRSL____ReplicaModel__i_Compile "357_LiveRSL____ReplicaModel__i_Compile_"
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
var _ _214_LiveRSL____CTypes__i_Compile.Dummy__
var _ _217_LiveRSL____CMessage__i_Compile.Dummy__
var _ _225_LiveRSL____CMessageRefinements__i_Compile.Dummy__
var _ _228_LiveRSL____PacketParsing__i_Compile.Dummy__
var _ _234_Common____UpperBound__i_Compile.Dummy__
var _ _236_LiveRSL____ParametersState__i_Compile.Dummy__
var _ _238_LiveRSL____CPaxosConfiguration__i_Compile.Dummy__
var _ _240_LiveRSL____ConstantsState__i_Compile.Dummy__
var _ _243_LiveRSL____PaxosWorldState__i_Compile.Dummy__
var _ _245_LiveRSL____ReplicaConstantsState__i_Compile.Dummy__
var _ _251_LiveRSL____ElectionState__i_Compile.Dummy__
var _ _254_LiveRSL____ProposerState__i_Compile.Dummy__
var _ _265_LiveRSL____COperationNumberSort__i_Compile.Dummy__
var _ _267_LiveRSL____CLastCheckpointedMap__i_Compile.Dummy__
var _ _269_LiveRSL____AcceptorState__i_Compile.Dummy__
var _ _276_LiveRSL____ExecutorState__i_Compile.Dummy__
var _ _278_LiveRSL____LearnerState__i_Compile.Dummy__
var _ _283_LiveRSL____CClockReading__i_Compile.Dummy__
var _ _285_LiveRSL____ReplicaState__i_Compile.Dummy__
var _ _294_LiveRSL____MinCQuorumSize__i_Compile.Dummy__
var _ _297_LiveRSL____ElectionModel__i_Compile.Dummy__
var _ _301_Impl____LiveRSL____Broadcast__i_Compile.Dummy__
var _ _305_LiveRSL____ProposerLemmas__i_Compile.Dummy__
var _ _308_LiveRSL____ProposerModel__i_Compile.Dummy__
var _ _312_LiveRSL____AcceptorModel__i_Compile.Dummy__
var _ _316_LiveRSL____LearnerModel__i_Compile.Dummy__
var _ _321_LiveRSL____ExecutorModel__i_Compile.Dummy__
var _ _323_LiveRSL____ReplicaModel__Part1__i_Compile.Dummy__
var _ _331_LiveRSL____ReplicaModel__Part2__i_Compile.Dummy__
var _ _339_LiveRSL____ReplicaModel__Part3__i_Compile.Dummy__
var _ _347_LiveRSL____ReplicaModel__Part4__i_Compile.Dummy__
var _ _355_LiveRSL____ReplicaModel__Part5__i_Compile.Dummy__
var _ _357_LiveRSL____ReplicaModel__i_Compile.Dummy__

type Dummy__ struct{}

// Definition of data type ReceiveResult
type ReceiveResult struct {
	Data_ReceiveResult_
}

func (_this ReceiveResult) Get() Data_ReceiveResult_ {
	return _this.Data_ReceiveResult_
}

type Data_ReceiveResult_ interface {
	isReceiveResult()
}

type CompanionStruct_ReceiveResult_ struct{}

var Companion_ReceiveResult_ = CompanionStruct_ReceiveResult_{}

type ReceiveResult_RRFail struct {
}

func (ReceiveResult_RRFail) isReceiveResult() {}

func (CompanionStruct_ReceiveResult_) Create_RRFail_() ReceiveResult {
	return ReceiveResult{ReceiveResult_RRFail{}}
}

func (_this ReceiveResult) Is_RRFail() bool {
	_, ok := _this.Get().(ReceiveResult_RRFail)
	return ok
}

type ReceiveResult_RRTimeout struct {
}

func (ReceiveResult_RRTimeout) isReceiveResult() {}

func (CompanionStruct_ReceiveResult_) Create_RRTimeout_() ReceiveResult {
	return ReceiveResult{ReceiveResult_RRTimeout{}}
}

func (_this ReceiveResult) Is_RRTimeout() bool {
	_, ok := _this.Get().(ReceiveResult_RRTimeout)
	return ok
}

type ReceiveResult_RRPacket struct {
	Cpacket _217_LiveRSL____CMessage__i_Compile.CPacket
}

func (ReceiveResult_RRPacket) isReceiveResult() {}

func (CompanionStruct_ReceiveResult_) Create_RRPacket_(Cpacket _217_LiveRSL____CMessage__i_Compile.CPacket) ReceiveResult {
	return ReceiveResult{ReceiveResult_RRPacket{Cpacket}}
}

func (_this ReceiveResult) Is_RRPacket() bool {
	_, ok := _this.Get().(ReceiveResult_RRPacket)
	return ok
}

func (_this ReceiveResult) Dtor_cpacket() _217_LiveRSL____CMessage__i_Compile.CPacket {
	return _this.Get().(ReceiveResult_RRPacket).Cpacket
}

func (_this ReceiveResult) String() string {
	switch data := _this.Get().(type) {
	case nil:
		return "null"
	case ReceiveResult_RRFail:
		{
			return "ReceiveResult.RRFail"
		}
	case ReceiveResult_RRTimeout:
		{
			return "ReceiveResult.RRTimeout"
		}
	case ReceiveResult_RRPacket:
		{
			return "ReceiveResult.RRPacket" + "(" + _dafny.String(data.Cpacket) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ReceiveResult) Equals(other ReceiveResult) bool {
	switch data1 := _this.Get().(type) {
	case ReceiveResult_RRFail:
		{
			_, ok := other.Get().(ReceiveResult_RRFail)
			return ok
		}
	case ReceiveResult_RRTimeout:
		{
			_, ok := other.Get().(ReceiveResult_RRTimeout)
			return ok
		}
	case ReceiveResult_RRPacket:
		{
			data2, ok := other.Get().(ReceiveResult_RRPacket)
			return ok && data1.Cpacket.Equals(data2.Cpacket)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ReceiveResult) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ReceiveResult)
	return ok && _this.Equals(typed)
}
func Type_ReceiveResult_() _dafny.Type {
	return type_ReceiveResult_{}
}

type type_ReceiveResult_ struct {
}

func (_this type_ReceiveResult_) Default() interface{} {
	return ReceiveResult{ReceiveResult_RRFail{}}
}

func (_this type_ReceiveResult_) String() string {
	return "ReceiveResult"
}

// End of data type ReceiveResult

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
func (_this *CompanionStruct_Default___) GetEndPoint(ipe *_9_Native____Io__s_Compile.IPEndPoint) _9_Native____Io__s_Compile.EndPoint {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var ep _9_Native____Io__s_Compile.EndPoint = _9_Native____Io__s_Compile.Type_EndPoint_().Default().(_9_Native____Io__s_Compile.EndPoint)
	var _ = ep
	var _5199_addr *_dafny.Array
	var _ = _5199_addr
	var _out377 *_dafny.Array
	var _ = _out377
	_out377 = (ipe).GetAddress()
	_5199_addr = _out377
	var _5200_port uint16
	var _ = _5200_port
	_5200_port = (ipe).GetPort()
	ep = _9_Native____Io__s_Compile.EndPoint{_9_Native____Io__s_Compile.EndPoint_EndPoint{(_5199_addr).RangeToSeq(_dafny.NilInt, _dafny.NilInt), _5200_port}}
	return ep
}
func (_this *CompanionStruct_Default___) Receive(udpClient *_9_Native____Io__s_Compile.UdpClient, localAddr _9_Native____Io__s_Compile.EndPoint, config _238_LiveRSL____CPaxosConfiguration__i_Compile.CPaxosConfiguration, msg__grammar _176_Common____GenericMarshalling__i_Compile.G) ReceiveResult {
	// TONY: Check to see why every message is parsed as Invalid
	var rr ReceiveResult = Type_ReceiveResult_().Default().(ReceiveResult)
	var _ = rr
	var _5201_timeout int32
	var _ = _5201_timeout
	_5201_timeout = int32(0)
	{
	}
	var _5202_ok bool
	var _ = _5202_ok
	var _5203_timedOut bool
	var _ = _5203_timedOut
	var _5204_remote *_9_Native____Io__s_Compile.IPEndPoint
	var _ = _5204_remote
	var _5205_buffer *_dafny.Array
	var _ = _5205_buffer
	var _out378 bool
	var _ = _out378
	var _out379 bool
	var _ = _out379
	var _out380 *_9_Native____Io__s_Compile.IPEndPoint
	var _ = _out380
	var _out381 *_dafny.Array
	var _ = _out381
	_out378, _out379, _out380, _out381 = (udpClient).Receive(_5201_timeout)
	_5202_ok = _out378
	_5203_timedOut = _out379
	_5204_remote = _out380
	_5205_buffer = _out381
	if !(_5202_ok) {
		rr = ReceiveResult{ReceiveResult_RRFail{}}
		return rr
	}
	if _5203_timedOut {
		rr = ReceiveResult{ReceiveResult_RRTimeout{}}
		{
		}
		return rr
	}
	{
	}
	{
	}
	{
	}
	var _5206_start__time uint64
	var _ = _5206_start__time
	var _out382 uint64
	var _ = _out382
	_out382 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_5206_start__time = _out382
	{
	}
	var _5207_cmessage _217_LiveRSL____CMessage__i_Compile.CMessage
	var _ = _5207_cmessage
	var _out383 _217_LiveRSL____CMessage__i_Compile.CMessage
	var _ = _out383
	_out383 = _228_LiveRSL____PacketParsing__i_Compile.Companion_Default___.PaxosDemarshallDataMethod(_5205_buffer, msg__grammar)
	_5207_cmessage = _out383
	var _5208_end__time uint64
	var _ = _5208_end__time
	var _out384 uint64
	var _ = _out384
	_out384 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
	_5208_end__time = _out384
	_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("PaxosDemarshallDataMethod"), _5206_start__time, _5208_end__time)
	var _5209_srcEp _9_Native____Io__s_Compile.EndPoint
	var _ = _5209_srcEp
	var _out385 _9_Native____Io__s_Compile.EndPoint
	var _ = _out385
	_out385 = Companion_Default___.GetEndPoint(_5204_remote)
	_5209_srcEp = _out385
	var _5210_cpacket _217_LiveRSL____CMessage__i_Compile.CPacket
	var _ = _5210_cpacket
	_5210_cpacket = _217_LiveRSL____CMessage__i_Compile.CPacket{_217_LiveRSL____CMessage__i_Compile.CPacket_CPacket{localAddr, _5209_srcEp, _5207_cmessage}}
	rr = ReceiveResult{ReceiveResult_RRPacket{_5210_cpacket}}
	{
	}
	if (_5207_cmessage).Is_CMessage__Invalid() {
		_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("DemarshallMessage_Invalid"), _5206_start__time, _5208_end__time)
	} else if (_5207_cmessage).Is_CMessage__Request() {
		_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("DemarshallMessage_Request"), _5206_start__time, _5208_end__time)
	} else if (_5207_cmessage).Is_CMessage__1a() {
		_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("DemarshallMessage_1a"), _5206_start__time, _5208_end__time)
	} else if (_5207_cmessage).Is_CMessage__1b() {
		_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("DemarshallMessage_1b"), _5206_start__time, _5208_end__time)
	} else if (_5207_cmessage).Is_CMessage__2a() {
		_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("DemarshallMessage_2a"), _5206_start__time, _5208_end__time)
	} else if (_5207_cmessage).Is_CMessage__2b() {
		_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("DemarshallMessage_2b"), _5206_start__time, _5208_end__time)
	} else if (_5207_cmessage).Is_CMessage__Heartbeat() {
		_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("DemarshallMessage_Heartbeat"), _5206_start__time, _5208_end__time)
	} else if (_5207_cmessage).Is_CMessage__Reply() {
		_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("DemarshallMessage_Reply"), _5206_start__time, _5208_end__time)
	} else if (_5207_cmessage).Is_CMessage__AppStateRequest() {
		_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("DemarshallMessage_AppStateRequest"), _5206_start__time, _5208_end__time)
	} else if (_5207_cmessage).Is_CMessage__AppStateSupply() {
		_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("DemarshallMessage_AppStateSupply"), _5206_start__time, _5208_end__time)
	} else if (_5207_cmessage).Is_CMessage__StartingPhase2() {
		_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("DemarshallMessage_StartingPhase2"), _5206_start__time, _5208_end__time)
	}
	{
	}
	{
	}
	{
	}
	return rr
}
func (_this *CompanionStruct_Default___) ReadClock(udpClient *_9_Native____Io__s_Compile.UdpClient) _283_LiveRSL____CClockReading__i_Compile.CClockReading {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var clock _283_LiveRSL____CClockReading__i_Compile.CClockReading = _283_LiveRSL____CClockReading__i_Compile.Type_CClockReading_().Default().(_283_LiveRSL____CClockReading__i_Compile.CClockReading)
	var _ = clock
	var _5211_t uint64
	var _ = _5211_t
	var _out386 uint64
	var _ = _out386
	_out386 = _9_Native____Io__s_Compile.Companion_Time_.GetTime()
	_5211_t = _out386
	{
	}
	clock = _283_LiveRSL____CClockReading__i_Compile.CClockReading{_283_LiveRSL____CClockReading__i_Compile.CClockReading_CClockReading{_5211_t}}
	return clock
}
func (_this *CompanionStruct_Default___) SendBroadcast(udpClient *_9_Native____Io__s_Compile.UdpClient, broadcast _217_LiveRSL____CMessage__i_Compile.CBroadcast) bool {
	var ok bool = false
	var _ = ok
	ok = true
	{
	}
	if (broadcast).Is_CBroadcastNop() {
	} else {
		{
		}
		{
		}
		var _5212_buffer *_dafny.Array
		var _ = _5212_buffer
		var _out387 *_dafny.Array
		var _ = _out387
		_out387 = _228_LiveRSL____PacketParsing__i_Compile.Companion_Default___.PaxosMarshall((broadcast).Dtor_msg())
		_5212_buffer = _out387
		{
		}
		{
		}
		var _5213_i uint64
		var _ = _5213_i
		_5213_i = uint64(0)
		for (_5213_i) < (uint64(((broadcast).Dtor_dsts()).CardinalityUint64())) {
			{
			}
			var _5214_dstEp _9_Native____Io__s_Compile.EndPoint
			var _ = _5214_dstEp
			_5214_dstEp = ((broadcast).Dtor_dsts()).IndexUint(_5213_i).(_9_Native____Io__s_Compile.EndPoint)
			var _5215_dstAddrAry *_dafny.Array
			var _ = _5215_dstAddrAry
			var _out388 *_dafny.Array
			var _ = _out388
			_out388 = _170_Common____Util__i_Compile.Companion_Default___.SeqToArrayOpt(_0_Native____NativeTypes__s_Compile.Type_Byte_(), (_5214_dstEp).Dtor_addr())
			_5215_dstAddrAry = _out388
			var _5216_remote *_9_Native____Io__s_Compile.IPEndPoint = (*_9_Native____Io__s_Compile.IPEndPoint)(nil)
			var _ = _5216_remote
			var _out389 bool
			var _ = _out389
			var _out390 *_9_Native____Io__s_Compile.IPEndPoint
			var _ = _out390
			_out389, _out390 = _9_Native____Io__s_Compile.Companion_IPEndPoint_.Construct(_5215_dstAddrAry, (_5214_dstEp).Dtor_port())
			ok = _out389
			_5216_remote = _out390
			if !(ok) {
				return ok
			}
			var _out391 bool
			var _ = _out391
			_out391 = (udpClient).Send(_5216_remote, _5212_buffer)
			ok = _out391
			if !(ok) {
				return ok
			}
			{
			}
			{
			}
			{
			}
			_5213_i = (_5213_i) + (uint64(1))
		}
	}
	return ok
}
func (_this *CompanionStruct_Default___) SendPacket(udpClient *_9_Native____Io__s_Compile.UdpClient, packets _217_LiveRSL____CMessage__i_Compile.OutboundPackets) bool {
	var ok bool = false
	var _ = ok
	var _5217_j uint64
	var _ = _5217_j
	_5217_j = uint64(0)
	{
	}
	ok = true
	var _5218_opt__packet _135_Logic____Option__i_Compile.Option
	var _ = _5218_opt__packet
	_5218_opt__packet = (packets).Dtor_p()
	if (_5218_opt__packet).Is_None() {
	} else {
		var _5219_cpacket _217_LiveRSL____CMessage__i_Compile.CPacket
		var _ = _5219_cpacket
		_5219_cpacket = (_5218_opt__packet).Dtor_v().(_217_LiveRSL____CMessage__i_Compile.CPacket)
		{
		}
		var _5220_dstEp _9_Native____Io__s_Compile.EndPoint
		var _ = _5220_dstEp
		_5220_dstEp = (_5219_cpacket).Dtor_dst()
		var _5221_dstAddrAry *_dafny.Array
		var _ = _5221_dstAddrAry
		var _out392 *_dafny.Array
		var _ = _out392
		_out392 = _170_Common____Util__i_Compile.Companion_Default___.SeqToArrayOpt(_0_Native____NativeTypes__s_Compile.Type_Byte_(), (_5220_dstEp).Dtor_addr())
		_5221_dstAddrAry = _out392
		var _5222_remote *_9_Native____Io__s_Compile.IPEndPoint = (*_9_Native____Io__s_Compile.IPEndPoint)(nil)
		var _ = _5222_remote
		var _out393 bool
		var _ = _out393
		var _out394 *_9_Native____Io__s_Compile.IPEndPoint
		var _ = _out394
		_out393, _out394 = _9_Native____Io__s_Compile.Companion_IPEndPoint_.Construct(_5221_dstAddrAry, (_5220_dstEp).Dtor_port())
		ok = _out393
		_5222_remote = _out394
		if !(ok) {
			return ok
		}
		{
		}
		{
		}
		var _5223_marshall__start__time uint64
		var _ = _5223_marshall__start__time
		var _out395 uint64
		var _ = _out395
		_out395 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
		_5223_marshall__start__time = _out395
		var _5224_buffer *_dafny.Array
		var _ = _5224_buffer
		var _out396 *_dafny.Array
		var _ = _out396
		_out396 = _228_LiveRSL____PacketParsing__i_Compile.Companion_Default___.PaxosMarshall((_5219_cpacket).Dtor_msg())
		_5224_buffer = _out396
		var _5225_marshall__end__time uint64
		var _ = _5225_marshall__end__time
		var _out397 uint64
		var _ = _out397
		_out397 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
		_5225_marshall__end__time = _out397
		_170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("SendBatch_PaxosMarshall"), _5223_marshall__start__time, _5225_marshall__end__time)
		{
		}
		{
		}
		var _out398 bool
		var _ = _out398
		_out398 = (udpClient).Send(_5222_remote, _5224_buffer)
		ok = _out398
		if !(ok) {
			return ok
		}
		{
		}
		{
		}
		{
		}
		{
		}
		{
		}
	}
	return ok
}
func (_this *CompanionStruct_Default___) SendPacketSequence(udpClient *_9_Native____Io__s_Compile.UdpClient, packets _217_LiveRSL____CMessage__i_Compile.OutboundPackets) bool {
	var ok bool = false
	var _ = ok
	var _5226_cpackets _dafny.Seq
	var _ = _5226_cpackets
	_5226_cpackets = (packets).Dtor_s()
	var _5227_j uint64
	var _ = _5227_j
	_5227_j = uint64(0)
	{
	}
	ok = true
	{
	}
	{
	}
	var _5228_i uint64
	var _ = _5228_i
	_5228_i = uint64(0)
	for (_5228_i) < (uint64((_5226_cpackets).CardinalityUint64())) {
		var _5229_cpacket _217_LiveRSL____CMessage__i_Compile.CPacket
		var _ = _5229_cpacket
		_5229_cpacket = (_5226_cpackets).IndexUint(_5228_i).(_217_LiveRSL____CMessage__i_Compile.CPacket)
		var _5230_dstEp _9_Native____Io__s_Compile.EndPoint
		var _ = _5230_dstEp
		_5230_dstEp = (_5229_cpacket).Dtor_dst()
		{
		}
		{
		}
		var _5231_dstAddrAry *_dafny.Array
		var _ = _5231_dstAddrAry
		var _out399 *_dafny.Array
		var _ = _out399
		_out399 = _170_Common____Util__i_Compile.Companion_Default___.SeqToArrayOpt(_0_Native____NativeTypes__s_Compile.Type_Byte_(), (_5230_dstEp).Dtor_addr())
		_5231_dstAddrAry = _out399
		var _5232_remote *_9_Native____Io__s_Compile.IPEndPoint = (*_9_Native____Io__s_Compile.IPEndPoint)(nil)
		var _ = _5232_remote
		var _out400 bool
		var _ = _out400
		var _out401 *_9_Native____Io__s_Compile.IPEndPoint
		var _ = _out401
		_out400, _out401 = _9_Native____Io__s_Compile.Companion_IPEndPoint_.Construct(_5231_dstAddrAry, (_5230_dstEp).Dtor_port())
		ok = _out400
		_5232_remote = _out401
		if !(ok) {
			return ok
		}
		{
		}
		{
		}
		var _5233_buffer *_dafny.Array
		var _ = _5233_buffer
		var _out402 *_dafny.Array
		var _ = _out402
		_out402 = _228_LiveRSL____PacketParsing__i_Compile.Companion_Default___.PaxosMarshall((_5229_cpacket).Dtor_msg())
		_5233_buffer = _out402
		{
		}
		{
		}
		var _out403 bool
		var _ = _out403
		_out403 = (udpClient).Send(_5232_remote, _5233_buffer)
		ok = _out403
		if !(ok) {
			return ok
		}
		{
		}
		{
		}
		{
		}
		{
		}
		{
		}
		{
		}
		{
		}
		_5228_i = (_5228_i) + (uint64(1))
	}
	return ok
}

// End of class Default__
