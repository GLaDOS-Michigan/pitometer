// Package _119_UdpLock__i_Compile
// Dafny module _119_UdpLock__i_Compile compiled into Go

package _119_UdpLock__i_Compile

import (
	_0_Native____NativeTypes__s_Compile "0_Native____NativeTypes__s_Compile_"
	_101_Common____Util__i_Compile "101_Common____Util__i_Compile_"
	_105_Common____MarshallInt__i_Compile "105_Common____MarshallInt__i_Compile_"
	_107_Common____GenericMarshalling__i_Compile "107_Common____GenericMarshalling__i_Compile_"
	_111_PacketParsing__i_Compile "111_PacketParsing__i_Compile_"
	_113_Common____SeqIsUniqueDef__i_Compile "113_Common____SeqIsUniqueDef__i_Compile_"
	_115_Impl__Node__i_Compile "115_Impl__Node__i_Compile_"
	_26_Collections____Seqs__s_Compile "26_Collections____Seqs__s_Compile_"
	_29_Collections____Sets__i_Compile "29_Collections____Sets__i_Compile_"
	_2_Collections____Maps2__s_Compile "2_Collections____Maps2__s_Compile_"
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
	_5_Temporal____Temporal__s_Compile "5_Temporal____Temporal__s_Compile_"
	_61_Math____power__s_Compile "61_Math____power__s_Compile_"
	_64_Math____mul__nonlinear__i_Compile "64_Math____mul__nonlinear__i_Compile_"
	_67_Math____mul__auto__proofs__i_Compile "67_Math____mul__auto__proofs__i_Compile_"
	_69_Math____mul__auto__i_Compile "69_Math____mul__auto__i_Compile_"
	_71_Math____mul__i_Compile "71_Math____mul__i_Compile_"
	_73_Math____power__i_Compile "73_Math____power__i_Compile_"
	_77_Math____div__def__i_Compile "77_Math____div__def__i_Compile_"
	_7_Environment__s_Compile "7_Environment__s_Compile_"
	_81_Math____div__boogie__i_Compile "81_Math____div__boogie__i_Compile_"
	_83_Math____div__nonlinear__i_Compile "83_Math____div__nonlinear__i_Compile_"
	_88_Math____mod__auto__proofs__i_Compile "88_Math____mod__auto__proofs__i_Compile_"
	_90_Math____mod__auto__i_Compile "90_Math____mod__auto__i_Compile_"
	_93_Math____div__auto__proofs__i_Compile "93_Math____div__auto__proofs__i_Compile_"
	_95_Math____div__auto__i_Compile "95_Math____div__auto__i_Compile_"
	_97_Math____div__i_Compile "97_Math____div__i_Compile_"
	_99_Math____power2__i_Compile "99_Math____power2__i_Compile_"
	_9_Native____Io__s_Compile "9_Native____Io__s_Compile_"
	_System "System_"
	_dafny "dafny"
	"fmt"
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
	Cpacket _7_Environment__s_Compile.LPacket
}

func (ReceiveResult_RRPacket) isReceiveResult() {}

func (CompanionStruct_ReceiveResult_) Create_RRPacket_(Cpacket _7_Environment__s_Compile.LPacket) ReceiveResult {
	return ReceiveResult{ReceiveResult_RRPacket{Cpacket}}
}

func (_this ReceiveResult) Is_RRPacket() bool {
	_, ok := _this.Get().(ReceiveResult_RRPacket)
	return ok
}

func (_this ReceiveResult) Dtor_cpacket() _7_Environment__s_Compile.LPacket {
	return _this.Get().(ReceiveResult_RRPacket).Cpacket
}

func (_this ReceiveResult) String() string {
	switch data := _this.Get().(type) {
	case nil:
		return "null"
	case ReceiveResult_RRFail:
		{
			return "_119_UdpLock__i_Compile.ReceiveResult.RRFail"
		}
	case ReceiveResult_RRTimeout:
		{
			return "_119_UdpLock__i_Compile.ReceiveResult.RRTimeout"
		}
	case ReceiveResult_RRPacket:
		{
			return "_119_UdpLock__i_Compile.ReceiveResult.RRPacket" + "(" + _dafny.String(data.Cpacket) + ")"
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
	return "_119_UdpLock__i_Compile.ReceiveResult"
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
	return "_119_UdpLock__i_Compile.Default__"
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
	return "_119_UdpLock__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) GetEndPoint(ipe *_9_Native____Io__s_Compile.IPEndPoint) _9_Native____Io__s_Compile.EndPoint {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var ep _9_Native____Io__s_Compile.EndPoint = _9_Native____Io__s_Compile.Type_EndPoint_().Default().(_9_Native____Io__s_Compile.EndPoint)
	var _ = ep
	var _1672_addr *_dafny.Array
	var _ = _1672_addr
	var _out65 *_dafny.Array
	var _ = _out65
	_out65 = (ipe).GetAddress()
	_1672_addr = _out65
	var _1673_port uint16
	var _ = _1673_port
	_1673_port = (ipe).GetPort()
	ep = _9_Native____Io__s_Compile.EndPoint{_9_Native____Io__s_Compile.EndPoint_EndPoint{(_1672_addr).RangeToSeq(_dafny.NilInt, _dafny.NilInt), _1673_port}}
	return ep
}
func (_this *CompanionStruct_Default___) Receive(udpClient *_9_Native____Io__s_Compile.UdpClient, localAddr _9_Native____Io__s_Compile.EndPoint) ReceiveResult {
	var rr ReceiveResult = Type_ReceiveResult_().Default().(ReceiveResult)
	var _ = rr
	var _1674_timeout int32
	var _ = _1674_timeout
	_1674_timeout = int32(0)
	{
	}
	var _1675_ok bool
	var _ = _1675_ok
	var _1676_timedOut bool
	var _ = _1676_timedOut
	var _1677_remote *_9_Native____Io__s_Compile.IPEndPoint
	var _ = _1677_remote
	var _1678_buffer *_dafny.Array
	var _ = _1678_buffer
	var _out66 bool
	var _ = _out66
	var _out67 bool
	var _ = _out67
	var _out68 *_9_Native____Io__s_Compile.IPEndPoint
	var _ = _out68
	var _out69 *_dafny.Array
	var _ = _out69
	_out66, _out67, _out68, _out69 = (udpClient).Receive(_1674_timeout)
	_1675_ok = _out66
	_1676_timedOut = _out67
	_1677_remote = _out68
	_1678_buffer = _out69
	if !(_1675_ok) {
		rr = ReceiveResult{ReceiveResult_RRFail{}}
		return rr
	}
	if _1676_timedOut {
		rr = ReceiveResult{ReceiveResult_RRTimeout{}}
		{
		}
		return rr
	}
	{
	}
	var _1679_cmessage _39_Message__i_Compile.CMessage
	var _ = _1679_cmessage
	var _out70 _39_Message__i_Compile.CMessage
	var _ = _out70
	_out70 = _111_PacketParsing__i_Compile.Companion_Default___.DemarshallDataMethod(_1678_buffer)
	_1679_cmessage = _out70
	var _1680_srcEp _9_Native____Io__s_Compile.EndPoint
	var _ = _1680_srcEp
	var _out71 _9_Native____Io__s_Compile.EndPoint
	var _ = _out71
	_out71 = Companion_Default___.GetEndPoint(_1677_remote)
	_1680_srcEp = _out71
	var _1681_cpacket _7_Environment__s_Compile.LPacket
	var _ = _1681_cpacket
	_1681_cpacket = _7_Environment__s_Compile.LPacket{_7_Environment__s_Compile.LPacket_LPacket{localAddr, _1680_srcEp, _1679_cmessage}}
	rr = ReceiveResult{ReceiveResult_RRPacket{_1681_cpacket}}
	return rr
}
func (_this *CompanionStruct_Default___) SendPacket(udpClient *_9_Native____Io__s_Compile.UdpClient, opt__packet _44_Logic____Option__i_Compile.Option) bool {
	var ok bool = false
	var _ = ok
	{
	}
	ok = true
	if (opt__packet).Is_None() {
	} else {
		var _1682_cpacket _7_Environment__s_Compile.LPacket
		var _ = _1682_cpacket
		_1682_cpacket = (opt__packet).Dtor_v().(_7_Environment__s_Compile.LPacket)
		var _1683_dstEp _9_Native____Io__s_Compile.EndPoint
		var _ = _1683_dstEp
		_1683_dstEp = (_1682_cpacket).Dtor_dst().(_9_Native____Io__s_Compile.EndPoint)
		var _1684_dstAddrAry *_dafny.Array
		var _ = _1684_dstAddrAry
		var _out72 *_dafny.Array
		var _ = _out72
		_out72 = _101_Common____Util__i_Compile.Companion_Default___.SeqToArrayOpt(_0_Native____NativeTypes__s_Compile.Type_Byte_(), (_1683_dstEp).Dtor_addr())
		_1684_dstAddrAry = _out72
		var _1685_remote *_9_Native____Io__s_Compile.IPEndPoint = (*_9_Native____Io__s_Compile.IPEndPoint)(nil)
		var _ = _1685_remote
		var _out73 bool
		var _ = _out73
		var _out74 *_9_Native____Io__s_Compile.IPEndPoint
		var _ = _out74
		_out73, _out74 = _9_Native____Io__s_Compile.Companion_IPEndPoint_.Construct(_1684_dstAddrAry, (_1683_dstEp).Dtor_port())
		ok = _out73
		_1685_remote = _out74
		if !(ok) {
			return ok
		}
		var _1686_buffer *_dafny.Array
		var _ = _1686_buffer
		var _out75 *_dafny.Array
		var _ = _out75
		_out75 = _111_PacketParsing__i_Compile.Companion_Default___.MarshallLockMessage((_1682_cpacket).Dtor_msg().(_39_Message__i_Compile.CMessage))
		_1686_buffer = _out75
		var _out76 bool
		var _ = _out76
		fmt.Printf("TONY DEBUG: buffer = %s\n", _1686_buffer.String())
		_out76 = (udpClient).Send(_1685_remote, _1686_buffer)
		ok = _out76
		if !(ok) {
			return ok
		}
		{
		}
		{
		}
	}
	return ok
}

// End of class Default__
