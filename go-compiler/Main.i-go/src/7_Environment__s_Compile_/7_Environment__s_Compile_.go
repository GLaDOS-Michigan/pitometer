// Package _7_Environment__s_Compile
// Dafny module _7_Environment__s_Compile compiled into Go

package _7_Environment__s_Compile

import (
	_0_Native____NativeTypes__s_Compile "0_Native____NativeTypes__s_Compile_"
	_2_Collections____Maps2__s_Compile "2_Collections____Maps2__s_Compile_"
	_5_Temporal____Temporal__s_Compile "5_Temporal____Temporal__s_Compile_"
	_System "System_"
	_dafny "dafny"
)

var _ _dafny.Dummy__
var _ _System.Dummy__
var _ _0_Native____NativeTypes__s_Compile.Dummy__
var _ _2_Collections____Maps2__s_Compile.Dummy__
var _ _5_Temporal____Temporal__s_Compile.Dummy__

type Dummy__ struct{}

// Definition of data type LPacket
type LPacket struct {
	Data_LPacket_
}

func (_this LPacket) Get() Data_LPacket_ {
	return _this.Data_LPacket_
}

type Data_LPacket_ interface {
	isLPacket()
}

type CompanionStruct_LPacket_ struct{}

var Companion_LPacket_ = CompanionStruct_LPacket_{}

type LPacket_LPacket struct {
	Dst interface{}
	Src interface{}
	Msg interface{}
}

func (LPacket_LPacket) isLPacket() {}

func (CompanionStruct_LPacket_) Create_LPacket_(Dst interface{}, Src interface{}, Msg interface{}) LPacket {
	return LPacket{LPacket_LPacket{Dst, Src, Msg}}
}

func (_this LPacket) Is_LPacket() bool {
	_, ok := _this.Get().(LPacket_LPacket)
	return ok
}

func (_this LPacket) Dtor_dst() interface{} {
	return _this.Get().(LPacket_LPacket).Dst
}

func (_this LPacket) Dtor_src() interface{} {
	return _this.Get().(LPacket_LPacket).Src
}

func (_this LPacket) Dtor_msg() interface{} {
	return _this.Get().(LPacket_LPacket).Msg
}

func (_this LPacket) String() string {
	switch data := _this.Get().(type) {
	case nil:
		return "null"
	case LPacket_LPacket:
		{
			return "_7_Environment__s_Compile.LPacket.LPacket" + "(" + _dafny.String(data.Dst) + ", " + _dafny.String(data.Src) + ", " + _dafny.String(data.Msg) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this LPacket) Equals(other LPacket) bool {
	switch data1 := _this.Get().(type) {
	case LPacket_LPacket:
		{
			data2, ok := other.Get().(LPacket_LPacket)
			return ok && _dafny.AreEqual(data1.Dst, data2.Dst) && _dafny.AreEqual(data1.Src, data2.Src) && _dafny.AreEqual(data1.Msg, data2.Msg)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this LPacket) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(LPacket)
	return ok && _this.Equals(typed)
}
func Type_LPacket_(Type_IdType_ _dafny.Type, Type_MessageType_ _dafny.Type) _dafny.Type {
	return type_LPacket_{Type_IdType_, Type_MessageType_}
}

type type_LPacket_ struct {
	Type_IdType_      _dafny.Type
	Type_MessageType_ _dafny.Type
}

func (_this type_LPacket_) Default() interface{} {
	Type_IdType_ := _this.Type_IdType_
	_ = Type_IdType_
	Type_MessageType_ := _this.Type_MessageType_
	_ = Type_MessageType_
	return LPacket{LPacket_LPacket{Type_IdType_.Default(), Type_IdType_.Default(), Type_MessageType_.Default()}}
}

func (_this type_LPacket_) String() string {
	return "_7_Environment__s_Compile.LPacket"
}

// End of data type LPacket

// Definition of data type LIoOp
type LIoOp struct {
	Data_LIoOp_
}

func (_this LIoOp) Get() Data_LIoOp_ {
	return _this.Data_LIoOp_
}

type Data_LIoOp_ interface {
	isLIoOp()
}

type CompanionStruct_LIoOp_ struct{}

var Companion_LIoOp_ = CompanionStruct_LIoOp_{}

type LIoOp_LIoOpSend struct {
	S LPacket
}

func (LIoOp_LIoOpSend) isLIoOp() {}

func (CompanionStruct_LIoOp_) Create_LIoOpSend_(S LPacket) LIoOp {
	return LIoOp{LIoOp_LIoOpSend{S}}
}

func (_this LIoOp) Is_LIoOpSend() bool {
	_, ok := _this.Get().(LIoOp_LIoOpSend)
	return ok
}

type LIoOp_LIoOpReceive struct {
	R LPacket
}

func (LIoOp_LIoOpReceive) isLIoOp() {}

func (CompanionStruct_LIoOp_) Create_LIoOpReceive_(R LPacket) LIoOp {
	return LIoOp{LIoOp_LIoOpReceive{R}}
}

func (_this LIoOp) Is_LIoOpReceive() bool {
	_, ok := _this.Get().(LIoOp_LIoOpReceive)
	return ok
}

type LIoOp_LIoOpTimeoutReceive struct {
}

func (LIoOp_LIoOpTimeoutReceive) isLIoOp() {}

func (CompanionStruct_LIoOp_) Create_LIoOpTimeoutReceive_() LIoOp {
	return LIoOp{LIoOp_LIoOpTimeoutReceive{}}
}

func (_this LIoOp) Is_LIoOpTimeoutReceive() bool {
	_, ok := _this.Get().(LIoOp_LIoOpTimeoutReceive)
	return ok
}

type LIoOp_LIoOpReadClock struct {
	T _dafny.Int
}

func (LIoOp_LIoOpReadClock) isLIoOp() {}

func (CompanionStruct_LIoOp_) Create_LIoOpReadClock_(T _dafny.Int) LIoOp {
	return LIoOp{LIoOp_LIoOpReadClock{T}}
}

func (_this LIoOp) Is_LIoOpReadClock() bool {
	_, ok := _this.Get().(LIoOp_LIoOpReadClock)
	return ok
}

func (_this LIoOp) Dtor_s() LPacket {
	return _this.Get().(LIoOp_LIoOpSend).S
}

func (_this LIoOp) Dtor_r() LPacket {
	return _this.Get().(LIoOp_LIoOpReceive).R
}

func (_this LIoOp) Dtor_t() _dafny.Int {
	return _this.Get().(LIoOp_LIoOpReadClock).T
}

func (_this LIoOp) String() string {
	switch data := _this.Get().(type) {
	case nil:
		return "null"
	case LIoOp_LIoOpSend:
		{
			return "_7_Environment__s_Compile.LIoOp.LIoOpSend" + "(" + _dafny.String(data.S) + ")"
		}
	case LIoOp_LIoOpReceive:
		{
			return "_7_Environment__s_Compile.LIoOp.LIoOpReceive" + "(" + _dafny.String(data.R) + ")"
		}
	case LIoOp_LIoOpTimeoutReceive:
		{
			return "_7_Environment__s_Compile.LIoOp.LIoOpTimeoutReceive"
		}
	case LIoOp_LIoOpReadClock:
		{
			return "_7_Environment__s_Compile.LIoOp.LIoOpReadClock" + "(" + _dafny.String(data.T) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this LIoOp) Equals(other LIoOp) bool {
	switch data1 := _this.Get().(type) {
	case LIoOp_LIoOpSend:
		{
			data2, ok := other.Get().(LIoOp_LIoOpSend)
			return ok && data1.S.Equals(data2.S)
		}
	case LIoOp_LIoOpReceive:
		{
			data2, ok := other.Get().(LIoOp_LIoOpReceive)
			return ok && data1.R.Equals(data2.R)
		}
	case LIoOp_LIoOpTimeoutReceive:
		{
			_, ok := other.Get().(LIoOp_LIoOpTimeoutReceive)
			return ok
		}
	case LIoOp_LIoOpReadClock:
		{
			data2, ok := other.Get().(LIoOp_LIoOpReadClock)
			return ok && data1.T.Cmp(data2.T) == 0
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this LIoOp) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(LIoOp)
	return ok && _this.Equals(typed)
}
func Type_LIoOp_() _dafny.Type {
	return type_LIoOp_{}
}

type type_LIoOp_ struct {
}

func (_this type_LIoOp_) Default() interface{} {
	return LIoOp{LIoOp_LIoOpTimeoutReceive{}}
}

func (_this type_LIoOp_) String() string {
	return "_7_Environment__s_Compile.LIoOp"
}

// End of data type LIoOp

// Definition of data type LEnvStep
type LEnvStep struct {
	Data_LEnvStep_
}

func (_this LEnvStep) Get() Data_LEnvStep_ {
	return _this.Data_LEnvStep_
}

type Data_LEnvStep_ interface {
	isLEnvStep()
}

type CompanionStruct_LEnvStep_ struct{}

var Companion_LEnvStep_ = CompanionStruct_LEnvStep_{}

type LEnvStep_LEnvStepHostIos struct {
	Actor    interface{}
	Ios      _dafny.Seq
	NodeStep interface{}
}

func (LEnvStep_LEnvStepHostIos) isLEnvStep() {}

func (CompanionStruct_LEnvStep_) Create_LEnvStepHostIos_(Actor interface{}, Ios _dafny.Seq, NodeStep interface{}) LEnvStep {
	return LEnvStep{LEnvStep_LEnvStepHostIos{Actor, Ios, NodeStep}}
}

func (_this LEnvStep) Is_LEnvStepHostIos() bool {
	_, ok := _this.Get().(LEnvStep_LEnvStepHostIos)
	return ok
}

type LEnvStep_LEnvStepDeliverPacket struct {
	P LPacket
}

func (LEnvStep_LEnvStepDeliverPacket) isLEnvStep() {}

func (CompanionStruct_LEnvStep_) Create_LEnvStepDeliverPacket_(P LPacket) LEnvStep {
	return LEnvStep{LEnvStep_LEnvStepDeliverPacket{P}}
}

func (_this LEnvStep) Is_LEnvStepDeliverPacket() bool {
	_, ok := _this.Get().(LEnvStep_LEnvStepDeliverPacket)
	return ok
}

type LEnvStep_LEnvStepAdvanceTime struct {
}

func (LEnvStep_LEnvStepAdvanceTime) isLEnvStep() {}

func (CompanionStruct_LEnvStep_) Create_LEnvStepAdvanceTime_() LEnvStep {
	return LEnvStep{LEnvStep_LEnvStepAdvanceTime{}}
}

func (_this LEnvStep) Is_LEnvStepAdvanceTime() bool {
	_, ok := _this.Get().(LEnvStep_LEnvStepAdvanceTime)
	return ok
}

type LEnvStep_LEnvStepStutter struct {
}

func (LEnvStep_LEnvStepStutter) isLEnvStep() {}

func (CompanionStruct_LEnvStep_) Create_LEnvStepStutter_() LEnvStep {
	return LEnvStep{LEnvStep_LEnvStepStutter{}}
}

func (_this LEnvStep) Is_LEnvStepStutter() bool {
	_, ok := _this.Get().(LEnvStep_LEnvStepStutter)
	return ok
}

func (_this LEnvStep) Dtor_actor() interface{} {
	return _this.Get().(LEnvStep_LEnvStepHostIos).Actor
}

func (_this LEnvStep) Dtor_ios() _dafny.Seq {
	return _this.Get().(LEnvStep_LEnvStepHostIos).Ios
}

func (_this LEnvStep) Dtor_nodeStep() interface{} {
	return _this.Get().(LEnvStep_LEnvStepHostIos).NodeStep
}

func (_this LEnvStep) Dtor_p() LPacket {
	return _this.Get().(LEnvStep_LEnvStepDeliverPacket).P
}

func (_this LEnvStep) String() string {
	switch data := _this.Get().(type) {
	case nil:
		return "null"
	case LEnvStep_LEnvStepHostIos:
		{
			return "_7_Environment__s_Compile.LEnvStep.LEnvStepHostIos" + "(" + _dafny.String(data.Actor) + ", " + _dafny.String(data.Ios) + ", " + _dafny.String(data.NodeStep) + ")"
		}
	case LEnvStep_LEnvStepDeliverPacket:
		{
			return "_7_Environment__s_Compile.LEnvStep.LEnvStepDeliverPacket" + "(" + _dafny.String(data.P) + ")"
		}
	case LEnvStep_LEnvStepAdvanceTime:
		{
			return "_7_Environment__s_Compile.LEnvStep.LEnvStepAdvanceTime"
		}
	case LEnvStep_LEnvStepStutter:
		{
			return "_7_Environment__s_Compile.LEnvStep.LEnvStepStutter"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this LEnvStep) Equals(other LEnvStep) bool {
	switch data1 := _this.Get().(type) {
	case LEnvStep_LEnvStepHostIos:
		{
			data2, ok := other.Get().(LEnvStep_LEnvStepHostIos)
			return ok && _dafny.AreEqual(data1.Actor, data2.Actor) && data1.Ios.Equals(data2.Ios) && _dafny.AreEqual(data1.NodeStep, data2.NodeStep)
		}
	case LEnvStep_LEnvStepDeliverPacket:
		{
			data2, ok := other.Get().(LEnvStep_LEnvStepDeliverPacket)
			return ok && data1.P.Equals(data2.P)
		}
	case LEnvStep_LEnvStepAdvanceTime:
		{
			_, ok := other.Get().(LEnvStep_LEnvStepAdvanceTime)
			return ok
		}
	case LEnvStep_LEnvStepStutter:
		{
			_, ok := other.Get().(LEnvStep_LEnvStepStutter)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this LEnvStep) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(LEnvStep)
	return ok && _this.Equals(typed)
}
func Type_LEnvStep_() _dafny.Type {
	return type_LEnvStep_{}
}

type type_LEnvStep_ struct {
}

func (_this type_LEnvStep_) Default() interface{} {
	return LEnvStep{LEnvStep_LEnvStepAdvanceTime{}}
}

func (_this type_LEnvStep_) String() string {
	return "_7_Environment__s_Compile.LEnvStep"
}

// End of data type LEnvStep

// Definition of data type LHostInfo
type LHostInfo struct {
	Data_LHostInfo_
}

func (_this LHostInfo) Get() Data_LHostInfo_ {
	return _this.Data_LHostInfo_
}

type Data_LHostInfo_ interface {
	isLHostInfo()
}

type CompanionStruct_LHostInfo_ struct{}

var Companion_LHostInfo_ = CompanionStruct_LHostInfo_{}

type LHostInfo_LHostInfo struct {
	Queue _dafny.Seq
}

func (LHostInfo_LHostInfo) isLHostInfo() {}

func (CompanionStruct_LHostInfo_) Create_LHostInfo_(Queue _dafny.Seq) LHostInfo {
	return LHostInfo{LHostInfo_LHostInfo{Queue}}
}

func (_this LHostInfo) Is_LHostInfo() bool {
	_, ok := _this.Get().(LHostInfo_LHostInfo)
	return ok
}

func (_this LHostInfo) Dtor_queue() _dafny.Seq {
	return _this.Get().(LHostInfo_LHostInfo).Queue
}

func (_this LHostInfo) String() string {
	switch data := _this.Get().(type) {
	case nil:
		return "null"
	case LHostInfo_LHostInfo:
		{
			return "_7_Environment__s_Compile.LHostInfo.LHostInfo" + "(" + _dafny.String(data.Queue) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this LHostInfo) Equals(other LHostInfo) bool {
	switch data1 := _this.Get().(type) {
	case LHostInfo_LHostInfo:
		{
			data2, ok := other.Get().(LHostInfo_LHostInfo)
			return ok && data1.Queue.Equals(data2.Queue)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this LHostInfo) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(LHostInfo)
	return ok && _this.Equals(typed)
}
func Type_LHostInfo_() _dafny.Type {
	return type_LHostInfo_{}
}

type type_LHostInfo_ struct {
}

func (_this type_LHostInfo_) Default() interface{} {
	return LHostInfo{LHostInfo_LHostInfo{_dafny.EmptySeq}}
}

func (_this type_LHostInfo_) String() string {
	return "_7_Environment__s_Compile.LHostInfo"
}

// End of data type LHostInfo

// Definition of data type LEnvironment
type LEnvironment struct {
	Data_LEnvironment_
}

func (_this LEnvironment) Get() Data_LEnvironment_ {
	return _this.Data_LEnvironment_
}

type Data_LEnvironment_ interface {
	isLEnvironment()
}

type CompanionStruct_LEnvironment_ struct{}

var Companion_LEnvironment_ = CompanionStruct_LEnvironment_{}

type LEnvironment_LEnvironment struct {
	Time        _dafny.Int
	SentPackets _dafny.Set
	HostInfo    _dafny.Map
	NextStep    LEnvStep
}

func (LEnvironment_LEnvironment) isLEnvironment() {}

func (CompanionStruct_LEnvironment_) Create_LEnvironment_(Time _dafny.Int, SentPackets _dafny.Set, HostInfo _dafny.Map, NextStep LEnvStep) LEnvironment {
	return LEnvironment{LEnvironment_LEnvironment{Time, SentPackets, HostInfo, NextStep}}
}

func (_this LEnvironment) Is_LEnvironment() bool {
	_, ok := _this.Get().(LEnvironment_LEnvironment)
	return ok
}

func (_this LEnvironment) Dtor_time() _dafny.Int {
	return _this.Get().(LEnvironment_LEnvironment).Time
}

func (_this LEnvironment) Dtor_sentPackets() _dafny.Set {
	return _this.Get().(LEnvironment_LEnvironment).SentPackets
}

func (_this LEnvironment) Dtor_hostInfo() _dafny.Map {
	return _this.Get().(LEnvironment_LEnvironment).HostInfo
}

func (_this LEnvironment) Dtor_nextStep() LEnvStep {
	return _this.Get().(LEnvironment_LEnvironment).NextStep
}

func (_this LEnvironment) String() string {
	switch data := _this.Get().(type) {
	case nil:
		return "null"
	case LEnvironment_LEnvironment:
		{
			return "_7_Environment__s_Compile.LEnvironment.LEnvironment" + "(" + _dafny.String(data.Time) + ", " + _dafny.String(data.SentPackets) + ", " + _dafny.String(data.HostInfo) + ", " + _dafny.String(data.NextStep) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this LEnvironment) Equals(other LEnvironment) bool {
	switch data1 := _this.Get().(type) {
	case LEnvironment_LEnvironment:
		{
			data2, ok := other.Get().(LEnvironment_LEnvironment)
			return ok && data1.Time.Cmp(data2.Time) == 0 && data1.SentPackets.Equals(data2.SentPackets) && data1.HostInfo.Equals(data2.HostInfo) && data1.NextStep.Equals(data2.NextStep)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this LEnvironment) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(LEnvironment)
	return ok && _this.Equals(typed)
}
func Type_LEnvironment_() _dafny.Type {
	return type_LEnvironment_{}
}

type type_LEnvironment_ struct {
}

func (_this type_LEnvironment_) Default() interface{} {
	return LEnvironment{LEnvironment_LEnvironment{_dafny.Zero, _dafny.EmptySet, _dafny.EmptyMap, Type_LEnvStep_().Default().(LEnvStep)}}
}

func (_this type_LEnvironment_) String() string {
	return "_7_Environment__s_Compile.LEnvironment"
}

// End of data type LEnvironment
