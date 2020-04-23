// Package _121_NodeImpl__i_Compile
// Dafny module _121_NodeImpl__i_Compile compiled into Go

package _121_NodeImpl__i_Compile

import (
	_0_Native____NativeTypes__s_Compile "0_Native____NativeTypes__s_Compile_"
	_101_Common____Util__i_Compile "101_Common____Util__i_Compile_"
	_105_Common____MarshallInt__i_Compile "105_Common____MarshallInt__i_Compile_"
	_107_Common____GenericMarshalling__i_Compile "107_Common____GenericMarshalling__i_Compile_"
	_111_PacketParsing__i_Compile "111_PacketParsing__i_Compile_"
	_113_Common____SeqIsUniqueDef__i_Compile "113_Common____SeqIsUniqueDef__i_Compile_"
	_115_Impl__Node__i_Compile "115_Impl__Node__i_Compile_"
	_119_UdpLock__i_Compile "119_UdpLock__i_Compile_"
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
	"clock"
	_dafny "dafny"
	"time"
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

type Dummy__ struct{}

// Definition of class NodeImpl
type NodeImpl struct {
	Node      _115_Impl__Node__i_Compile.CNode
	UdpClient *_9_Native____Io__s_Compile.UdpClient
	LocalAddr _9_Native____Io__s_Compile.EndPoint
}

func New_NodeImpl_() *NodeImpl {
	_this := NodeImpl{}

	_this.Node = _115_Impl__Node__i_Compile.Type_CNode_().Default().(_115_Impl__Node__i_Compile.CNode)
	_this.UdpClient = (*_9_Native____Io__s_Compile.UdpClient)(nil)
	_this.LocalAddr = _9_Native____Io__s_Compile.Type_EndPoint_().Default().(_9_Native____Io__s_Compile.EndPoint)
	return &_this
}

type CompanionStruct_NodeImpl_ struct {
}

var Companion_NodeImpl_ = CompanionStruct_NodeImpl_{}

func (_this *NodeImpl) Equals(other *NodeImpl) bool {
	return _this == other
}

func (_this *NodeImpl) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*NodeImpl)
	return ok && _this.Equals(other)
}

func (*NodeImpl) String() string {
	return "_121_NodeImpl__i_Compile.NodeImpl"
}

func Type_NodeImpl_() _dafny.Type {
	return type_NodeImpl_{}
}

type type_NodeImpl_ struct {
}

func (_this type_NodeImpl_) Default() interface{} {
	return (*NodeImpl)(nil)
}

func (_this type_NodeImpl_) String() string {
	return "_121_NodeImpl__i_Compile.NodeImpl"
}
func (_this *NodeImpl) Ctor__() {
	goto TAIL_CALL_START
TAIL_CALL_START:
	(_this).UdpClient = (*_9_Native____Io__s_Compile.UdpClient)(nil)
}
func (_this *NodeImpl) ConstructUdpClient(me _9_Native____Io__s_Compile.EndPoint) (bool, *_9_Native____Io__s_Compile.UdpClient) {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var ok bool = false
	var _ = ok
	var client *_9_Native____Io__s_Compile.UdpClient = (*_9_Native____Io__s_Compile.UdpClient)(nil)
	var _ = client
	var _1687_my__ep _9_Native____Io__s_Compile.EndPoint
	var _ = _1687_my__ep
	_1687_my__ep = me
	var _1688_ip__byte__array *_dafny.Array
	var _ = _1688_ip__byte__array
	var _nw6 = _dafny.NewArrayWithValue(0, ((_1687_my__ep).Dtor_addr()).Cardinality())
	var _ = _nw6
	_1688_ip__byte__array = _nw6
	_101_Common____Util__i_Compile.Companion_Default___.SeqIntoArrayOpt((_1687_my__ep).Dtor_addr(), _1688_ip__byte__array)
	var _1689_ip__endpoint *_9_Native____Io__s_Compile.IPEndPoint = (*_9_Native____Io__s_Compile.IPEndPoint)(nil)
	var _ = _1689_ip__endpoint
	var _out77 bool
	var _ = _out77
	var _out78 *_9_Native____Io__s_Compile.IPEndPoint
	var _ = _out78
	_out77, _out78 = _9_Native____Io__s_Compile.Companion_IPEndPoint_.Construct(_1688_ip__byte__array, (_1687_my__ep).Dtor_port())
	ok = _out77
	_1689_ip__endpoint = _out78
	if !(ok) {
		return ok, client
	}
	var _out79 bool
	var _ = _out79
	var _out80 *_9_Native____Io__s_Compile.UdpClient
	var _ = _out80
	_out79, _out80 = _9_Native____Io__s_Compile.Companion_UdpClient_.Construct(_1689_ip__endpoint)
	ok = _out79
	client = _out80
	{
	}
	return ok, client
}
func (_this *NodeImpl) InitNode(config _dafny.Seq, my__index uint64) bool {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var ok bool = false
	var _ = ok
	var _out81 bool
	var _ = _out81
	var _out82 *_9_Native____Io__s_Compile.UdpClient
	var _ = _out82
	_out81, _out82 = (_this).ConstructUdpClient((config).IndexUint64(my__index).(_9_Native____Io__s_Compile.EndPoint))
	ok = _out81
	(_this).UdpClient = _out82
	if ok {
		var _out83 _115_Impl__Node__i_Compile.CNode
		var _ = _out83
		_out83 = _115_Impl__Node__i_Compile.Companion_Default___.NodeInitImpl(my__index, config)
		(_this).Node = _out83
		{
		}
		(_this).LocalAddr = ((_this.Node).Dtor_config()).IndexUint64(my__index).(_9_Native____Io__s_Compile.EndPoint)
		{
		}
	}
	return ok
}

// TONY: Measure this
func (_this *NodeImpl) NodeNextGrant(nodeGrantCounter *clock.Counter, nodeGrantLog *clock.Stopwatch) bool {
	nodeGrantLog.LogStartEvent("NodeNextGrant")
	time.Sleep(10 * time.Millisecond)
	var ok bool = false
	var _ = ok
	var _1690_transfer__packet _44_Logic____Option__i_Compile.Option = _44_Logic____Option__i_Compile.Type_Option_().Default().(_44_Logic____Option__i_Compile.Option)
	var _ = _1690_transfer__packet
	var _out84 _115_Impl__Node__i_Compile.CNode
	var _ = _out84
	var _out85 _44_Logic____Option__i_Compile.Option
	var _ = _out85
	_out84, _out85 = _115_Impl__Node__i_Compile.Companion_Default___.NodeGrantImpl(_this.Node)
	(_this).Node = _out84
	_1690_transfer__packet = _out85 // This is the packet that is sent
	ok = true
	if (_1690_transfer__packet).Is_Some() {
		{
		}
		var _out86 bool
		var _ = _out86
		_out86 = _119_UdpLock__i_Compile.Companion_Default___.SendPacket(_this.UdpClient, _1690_transfer__packet)
		ok = _out86
		{
		}
	} else {
	}
	nodeGrantLog.LogEndEvent("NodeNextGrant")
	nodeGrantCounter.Increment()
	return ok
}
func (_this *NodeImpl) NodeNextAccept() bool {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var ok bool = false
	var _ = ok
	var _1691_rr _119_UdpLock__i_Compile.ReceiveResult = _119_UdpLock__i_Compile.Type_ReceiveResult_().Default().(_119_UdpLock__i_Compile.ReceiveResult)
	var _ = _1691_rr
	{
	}
	var _out87 _119_UdpLock__i_Compile.ReceiveResult
	var _ = _out87
	_out87 = _119_UdpLock__i_Compile.Companion_Default___.Receive(_this.UdpClient, _this.LocalAddr)
	_1691_rr = _out87
	{
	}
	if (_1691_rr).Is_RRFail() {
		ok = false
		return ok
	} else if (_1691_rr).Is_RRTimeout() {
		ok = true
		{
		}
		return ok
	} else {
		ok = true
		var _1692_locked__packet _44_Logic____Option__i_Compile.Option = _44_Logic____Option__i_Compile.Type_Option_().Default().(_44_Logic____Option__i_Compile.Option)
		var _ = _1692_locked__packet
		var _out88 _115_Impl__Node__i_Compile.CNode
		var _ = _out88
		var _out89 _44_Logic____Option__i_Compile.Option
		var _ = _out89
		_out88, _out89 = _115_Impl__Node__i_Compile.Companion_Default___.NodeAcceptImpl(_this.Node, (_1691_rr).Dtor_cpacket())
		(_this).Node = _out88
		_1692_locked__packet = _out89
		if (_1692_locked__packet).Is_Some() {
			{
			}
			var _out90 bool
			var _ = _out90
			_out90 = _119_UdpLock__i_Compile.Companion_Default___.SendPacket(_this.UdpClient, _1692_locked__packet)
			ok = _out90
			{
			}
		}
	}
	return ok
}
func (_this *NodeImpl) HostNextMain(nodeGrantCounter *clock.Counter, nodeGrantLog *clock.Stopwatch) bool {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var ok bool = false
	var _ = ok
	if (_this.Node).Dtor_held() {
		var _out91 bool
		var _ = _out91
		_out91 = (_this).NodeNextGrant(nodeGrantCounter, nodeGrantLog)
		ok = _out91
	} else {
		var _out92 bool
		var _ = _out92
		_out92 = (_this).NodeNextAccept()
		ok = _out92
	}
	return ok
}

// End of class NodeImpl
