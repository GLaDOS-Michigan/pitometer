// Package _115_Impl__Node__i_Compile
// Dafny module _115_Impl__Node__i_Compile compiled into Go

package _115_Impl__Node__i_Compile

import (
	_0_Native____NativeTypes__s_Compile "0_Native____NativeTypes__s_Compile_"
	_101_Common____Util__i_Compile "101_Common____Util__i_Compile_"
	_105_Common____MarshallInt__i_Compile "105_Common____MarshallInt__i_Compile_"
	_107_Common____GenericMarshalling__i_Compile "107_Common____GenericMarshalling__i_Compile_"
	_111_PacketParsing__i_Compile "111_PacketParsing__i_Compile_"
	_113_Common____SeqIsUniqueDef__i_Compile "113_Common____SeqIsUniqueDef__i_Compile_"
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
	"fmt"
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

type Dummy__ struct{}

// Definition of data type CNode
type CNode struct {
	Data_CNode_
}

func (_this CNode) Get() Data_CNode_ {
	return _this.Data_CNode_
}

type Data_CNode_ interface {
	isCNode()
}

type CompanionStruct_CNode_ struct{}

var Companion_CNode_ = CompanionStruct_CNode_{}

type CNode_CNode struct {
	Held      bool
	Epoch     uint64
	My__index uint64
	Config    _dafny.Seq
}

func (CNode_CNode) isCNode() {}

func (CompanionStruct_CNode_) Create_CNode_(Held bool, Epoch uint64, My__index uint64, Config _dafny.Seq) CNode {
	return CNode{CNode_CNode{Held, Epoch, My__index, Config}}
}

func (_this CNode) Is_CNode() bool {
	_, ok := _this.Get().(CNode_CNode)
	return ok
}

func (_this CNode) Dtor_held() bool {
	return _this.Get().(CNode_CNode).Held
}

func (_this CNode) Dtor_epoch() uint64 {
	return _this.Get().(CNode_CNode).Epoch
}

func (_this CNode) Dtor_my__index() uint64 {
	return _this.Get().(CNode_CNode).My__index
}

func (_this CNode) Dtor_config() _dafny.Seq {
	return _this.Get().(CNode_CNode).Config
}

func (_this CNode) String() string {
	switch data := _this.Get().(type) {
	case nil:
		return "null"
	case CNode_CNode:
		{
			return "CNode.CNode" + "(" + _dafny.String(data.Held) + ", " + _dafny.String(data.Epoch) + ", " + _dafny.String(data.My__index) + ", " + _dafny.String(data.Config) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this CNode) Equals(other CNode) bool {
	switch data1 := _this.Get().(type) {
	case CNode_CNode:
		{
			data2, ok := other.Get().(CNode_CNode)
			return ok && data1.Held == data2.Held && data1.Epoch == data2.Epoch && data1.My__index == data2.My__index && data1.Config.Equals(data2.Config)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this CNode) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(CNode)
	return ok && _this.Equals(typed)
}
func Type_CNode_() _dafny.Type {
	return type_CNode_{}
}

type type_CNode_ struct {
}

func (_this type_CNode_) Default() interface{} {
	return CNode{CNode_CNode{false, 0, 0, _dafny.EmptySeq}}
}

func (_this type_CNode_) String() string {
	return "CNode"
}

// End of data type CNode

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
	return "_115_Impl__Node__i_Compile.Default__"
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
	return "_115_Impl__Node__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) NodeInitImpl(my__index uint64, config _dafny.Seq) CNode {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var node CNode = Type_CNode_().Default().(CNode)
	var _ = node
	node = CNode{CNode_CNode{(my__index) == (uint64(0)), (func() uint64 {
		if (my__index) == (uint64(0)) {
			return uint64(1)
		}
		return uint64(0)
	})(), my__index, config}}
	if (node).Dtor_held() {
		_dafny.Print(_dafny.SeqOfString("I start holding the lock\n"))
	}
	return node
}

// TONY: This is the method that we want to time
func (_this *CompanionStruct_Default___) NodeGrantImpl(s CNode, delay int, nodeGrantCounter *clock.Counter, nodeGrantLog *clock.Stopwatch) (CNode, _44_Logic____Option__i_Compile.Option) {
	nodeGrantLog.LogStartEvent("NodeNextGrant")
	time.Sleep(time.Duration(delay) * time.Microsecond)
	goto TAIL_CALL_START
TAIL_CALL_START:
	var s_k CNode = Type_CNode_().Default().(CNode)
	var _ = s_k
	var packet _44_Logic____Option__i_Compile.Option = _44_Logic____Option__i_Compile.Type_Option_().Default().(_44_Logic____Option__i_Compile.Option)
	var _ = packet
	if ((s).Dtor_held()) && (((s).Dtor_epoch()) < (uint64(18446744073709551615))) {
		var _1669_ssss CNode
		var _ = _1669_ssss
		_1669_ssss = CNode{CNode_CNode{false, (s).Dtor_epoch(), (s).Dtor_my__index(), (s).Dtor_config()}}
		s_k = _1669_ssss
		var _1670_dst__index uint64
		var _ = _1670_dst__index
		_1670_dst__index = (((s).Dtor_my__index()) + (uint64(1))) % (uint64(((s).Dtor_config()).CardinalityInt64()))
		packet = _44_Logic____Option__i_Compile.Option{_44_Logic____Option__i_Compile.Option_Some{_7_Environment__s_Compile.LPacket{_7_Environment__s_Compile.LPacket_LPacket{((s).Dtor_config()).IndexUint64(_1670_dst__index).(_9_Native____Io__s_Compile.EndPoint), ((s).Dtor_config()).IndexUint64((s).Dtor_my__index()).(_9_Native____Io__s_Compile.EndPoint), _39_Message__i_Compile.CMessage{_39_Message__i_Compile.CMessage_CTransfer{((s).Dtor_epoch()) + (uint64(1))}}}}}}
		{
		}
		// _dafny.Print(_dafny.SeqOfString("I grant the lock "))
		// _dafny.Print((s).Dtor_epoch())
		// _dafny.Print(_dafny.SeqOfString("\n"))
		nodeGrantCounter.Increment()
	} else {
		// TONY: This branch is observed to never be taken
		s_k = s
		{
		}
		packet = _44_Logic____Option__i_Compile.Option{_44_Logic____Option__i_Compile.Option_None{}}
	}
	nodeGrantLog.LogEndEvent("NodeNextGrant")
	return s_k, packet
}

func (_this *CompanionStruct_Default___) NodeAcceptImpl(s CNode, transfer__packet _7_Environment__s_Compile.LPacket, delay int, nodeAcceptLog *clock.Stopwatch) (CNode, _44_Logic____Option__i_Compile.Option) {
	nodeAcceptLog.LogStartEvent("NodeNextAccept")
	time.Sleep(time.Duration(delay) * time.Millisecond)
	goto TAIL_CALL_START
TAIL_CALL_START:
	var s_k CNode = Type_CNode_().Default().(CNode)
	var _ = s_k
	var locked__packet _44_Logic____Option__i_Compile.Option = _44_Logic____Option__i_Compile.Type_Option_().Default().(_44_Logic____Option__i_Compile.Option)
	var _ = locked__packet
	{
	}
	if (((!((s).Dtor_held())) && (((s).Dtor_config()).Contains((transfer__packet).Dtor_src().(_9_Native____Io__s_Compile.EndPoint)))) && (((transfer__packet).Dtor_msg().(_39_Message__i_Compile.CMessage)).Is_CTransfer())) && ((((transfer__packet).Dtor_msg().(_39_Message__i_Compile.CMessage)).Dtor_transfer__epoch()) > ((s).Dtor_epoch())) {
		var _1671_ssss CNode
		var _ = _1671_ssss
		_1671_ssss = CNode{CNode_CNode{true, ((transfer__packet).Dtor_msg().(_39_Message__i_Compile.CMessage)).Dtor_transfer__epoch(), (s).Dtor_my__index(), (s).Dtor_config()}}
		s_k = _1671_ssss
		locked__packet = _44_Logic____Option__i_Compile.Option{_44_Logic____Option__i_Compile.Option_Some{_7_Environment__s_Compile.LPacket{_7_Environment__s_Compile.LPacket_LPacket{(transfer__packet).Dtor_src().(_9_Native____Io__s_Compile.EndPoint), ((s).Dtor_config()).IndexUint64((s).Dtor_my__index()).(_9_Native____Io__s_Compile.EndPoint), _39_Message__i_Compile.CMessage{_39_Message__i_Compile.CMessage_CLocked{((transfer__packet).Dtor_msg().(_39_Message__i_Compile.CMessage)).Dtor_transfer__epoch()}}}}}}
		{
		}
		// _dafny.Print(_dafny.SeqOfString("I hold the lock!\n"))
	} else {
		// TONY: This branch should not execute
		fmt.Printf("TONY DEBUG: bad bad bad\n")
		s_k = s
		locked__packet = _44_Logic____Option__i_Compile.Option{_44_Logic____Option__i_Compile.Option_None{}}
	}
	nodeAcceptLog.LogEndEvent("NodeNextAccept")
	return s_k, locked__packet
}

// End of class Default__
