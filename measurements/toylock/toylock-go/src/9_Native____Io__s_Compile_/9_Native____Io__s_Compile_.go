// Package _9_Native____Io__s_Compile
// Dafny module _9_Native____Io__s_Compile compiled into Go

package _9_Native____Io__s_Compile

import (
	_0_Native____NativeTypes__s_Compile "0_Native____NativeTypes__s_Compile_"
	_2_Collections____Maps2__s_Compile "2_Collections____Maps2__s_Compile_"
	_5_Temporal____Temporal__s_Compile "5_Temporal____Temporal__s_Compile_"
	_7_Environment__s_Compile "7_Environment__s_Compile_"
	_System "System_"
	_dafny "dafny"
	"encoding/json"
	"fmt"
	"goconcurrentqueue"
	"log"
	"net"
	"os"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

var _ _dafny.Dummy__
var _ _System.Dummy__
var _ _0_Native____NativeTypes__s_Compile.Dummy__
var _ _2_Collections____Maps2__s_Compile.Dummy__
var _ _5_Temporal____Temporal__s_Compile.Dummy__
var _ _7_Environment__s_Compile.Dummy__

//TONY: traceAndExit utility
func traceAndExit() {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	fmt.Printf("Error: unimplemented\n")
	fmt.Printf("%s:%d %s\n\n", frame.File, frame.Line, frame.Function)

	debug.PrintStack()
	os.Exit(1)
}

type Dummy__ struct{}

// Definition of class HostEnvironment
type HostEnvironment struct {
	dummy byte
}

func New_HostEnvironment_() *HostEnvironment {
	_this := HostEnvironment{}

	return &_this
}

type CompanionStruct_HostEnvironment_ struct {
}

var Companion_HostEnvironment_ = CompanionStruct_HostEnvironment_{}

func (_this *HostEnvironment) Equals(other *HostEnvironment) bool {
	return _this == other
}

func (_this *HostEnvironment) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*HostEnvironment)
	return ok && _this.Equals(other)
}

func (*HostEnvironment) String() string {
	return "_9_Native____Io__s_Compile.HostEnvironment"
}

func Type_HostEnvironment_() _dafny.Type {
	return type_HostEnvironment_{}
}

type type_HostEnvironment_ struct {
}

func (_this type_HostEnvironment_) Default() interface{} {
	return (*HostEnvironment)(nil)
}

func (_this type_HostEnvironment_) String() string {
	return "_9_Native____Io__s_Compile.HostEnvironment"
}

// End of class HostEnvironment

// Definition of class HostConstants
type HostConstants struct {
	dummy byte
}

func New_HostConstants_() *HostConstants {
	_this := HostConstants{}

	return &_this
}

type CompanionStruct_HostConstants_ struct {
}

var Companion_HostConstants_ = CompanionStruct_HostConstants_{}

// TONY
func (comp_hc *CompanionStruct_HostConstants_) NumCommandLineArgs() uint32 {
	// This count includes the first item which is the name of the program
	var res = uint32(len(os.Args))
	// fmt.Printf("Number of command line args is %d\n", res)
	return res
}

// TONY
func (comp_hc *CompanionStruct_HostConstants_) GetCommandLineArg(i uint64) *_dafny.Array {
	var byteArray = []byte(os.Args[i])
	var uint16Array []interface{}
	for _, value := range byteArray {
		uint16Array = append(uint16Array, interface{}(uint16(value)))
	}
	var res = _dafny.NewArrayWithValues(uint16Array...)
	// fmt.Printf("Command line arg[%d] is %s\n", i, res)
	return res
}

func (_this *HostConstants) Equals(other *HostConstants) bool {
	return _this == other
}

func (_this *HostConstants) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*HostConstants)
	return ok && _this.Equals(other)
}

func (*HostConstants) String() string {
	return "_9_Native____Io__s_Compile.HostConstants"
}

func Type_HostConstants_() _dafny.Type {
	return type_HostConstants_{}
}

type type_HostConstants_ struct {
}

func (_this type_HostConstants_) Default() interface{} {
	return (*HostConstants)(nil)
}

func (_this type_HostConstants_) String() string {
	return "_9_Native____Io__s_Compile.HostConstants"
}

// End of class HostConstants

// Definition of class OkState
type OkState struct {
	dummy byte
}

func New_OkState_() *OkState {
	_this := OkState{}

	return &_this
}

type CompanionStruct_OkState_ struct {
}

var Companion_OkState_ = CompanionStruct_OkState_{}

func (_this *OkState) Equals(other *OkState) bool {
	return _this == other
}

func (_this *OkState) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*OkState)
	return ok && _this.Equals(other)
}

func (*OkState) String() string {
	return "_9_Native____Io__s_Compile.OkState"
}

func Type_OkState_() _dafny.Type {
	return type_OkState_{}
}

type type_OkState_ struct {
}

func (_this type_OkState_) Default() interface{} {
	return (*OkState)(nil)
}

func (_this type_OkState_) String() string {
	return "_9_Native____Io__s_Compile.OkState"
}

// End of class OkState

// Definition of class NowState
type NowState struct {
	dummy byte
}

func New_NowState_() *NowState {
	_this := NowState{}

	return &_this
}

type CompanionStruct_NowState_ struct {
}

var Companion_NowState_ = CompanionStruct_NowState_{}

func (_this *NowState) Equals(other *NowState) bool {
	return _this == other
}

func (_this *NowState) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*NowState)
	return ok && _this.Equals(other)
}

func (*NowState) String() string {
	return "_9_Native____Io__s_Compile.NowState"
}

func Type_NowState_() _dafny.Type {
	return type_NowState_{}
}

type type_NowState_ struct {
}

func (_this type_NowState_) Default() interface{} {
	return (*NowState)(nil)
}

func (_this type_NowState_) String() string {
	return "_9_Native____Io__s_Compile.NowState"
}

// End of class NowState

// Definition of class Time
type Time struct {
	dummy byte
}

func New_Time_() *Time {
	_this := Time{}

	return &_this
}

type CompanionStruct_Time_ struct {
}

var Companion_Time_ = CompanionStruct_Time_{}

func (_this *Time) Equals(other *Time) bool {
	return _this == other
}

func (_this *Time) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*Time)
	return ok && _this.Equals(other)
}

func (*Time) String() string {
	return "_9_Native____Io__s_Compile.Time"
}

func Type_Time_() _dafny.Type {
	return type_Time_{}
}

type type_Time_ struct {
}

func (_this type_Time_) Default() interface{} {
	return (*Time)(nil)
}

func (_this type_Time_) String() string {
	return "_9_Native____Io__s_Compile.Time"
}

// End of class Time

// Definition of data type EndPoint
type EndPoint struct {
	Data_EndPoint_
}

func (_this EndPoint) Get() Data_EndPoint_ {
	return _this.Data_EndPoint_
}

type Data_EndPoint_ interface {
	isEndPoint()
}

type CompanionStruct_EndPoint_ struct{}

var Companion_EndPoint_ = CompanionStruct_EndPoint_{}

type EndPoint_EndPoint struct {
	Addr _dafny.Seq
	Port uint16
}

func (EndPoint_EndPoint) isEndPoint() {}

func (CompanionStruct_EndPoint_) Create_EndPoint_(Addr _dafny.Seq, Port uint16) EndPoint {
	return EndPoint{EndPoint_EndPoint{Addr, Port}}
}

func (_this EndPoint) Is_EndPoint() bool {
	_, ok := _this.Get().(EndPoint_EndPoint)
	return ok
}

func (_this EndPoint) Dtor_addr() _dafny.Seq {
	return _this.Get().(EndPoint_EndPoint).Addr
}

func (_this EndPoint) Dtor_port() uint16 {
	return _this.Get().(EndPoint_EndPoint).Port
}

func (_this EndPoint) String() string {
	switch data := _this.Get().(type) {
	case nil:
		return "null"
	case EndPoint_EndPoint:
		{
			return "_9_Native____Io__s_Compile.EndPoint.EndPoint" + "(" + _dafny.String(data.Addr) + ", " + _dafny.String(data.Port) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this EndPoint) Equals(other EndPoint) bool {
	switch data1 := _this.Get().(type) {
	case EndPoint_EndPoint:
		{
			data2, ok := other.Get().(EndPoint_EndPoint)
			return ok && data1.Addr.Equals(data2.Addr) && data1.Port == data2.Port
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this EndPoint) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(EndPoint)
	return ok && _this.Equals(typed)
}
func Type_EndPoint_() _dafny.Type {
	return type_EndPoint_{}
}

type type_EndPoint_ struct {
}

func (_this type_EndPoint_) Default() interface{} {
	return EndPoint{EndPoint_EndPoint{_dafny.EmptySeq, 0}}
}

func (_this type_EndPoint_) String() string {
	return "_9_Native____Io__s_Compile.EndPoint"
}

// End of data type EndPoint

// Definition of class UdpState
type UdpState struct {
	dummy byte
}

func New_UdpState_() *UdpState {
	_this := UdpState{}

	return &_this
}

type CompanionStruct_UdpState_ struct {
}

var Companion_UdpState_ = CompanionStruct_UdpState_{}

func (_this *UdpState) Equals(other *UdpState) bool {
	return _this == other
}

func (_this *UdpState) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*UdpState)
	return ok && _this.Equals(other)
}

func (*UdpState) String() string {
	return "_9_Native____Io__s_Compile.UdpState"
}

func Type_UdpState_() _dafny.Type {
	return type_UdpState_{}
}

type type_UdpState_ struct {
}

func (_this type_UdpState_) Default() interface{} {
	return (*UdpState)(nil)
}

func (_this type_UdpState_) String() string {
	return "_9_Native____Io__s_Compile.UdpState"
}

// End of class UdpState

// Definition of class IPEndPoint
type IPEndPoint struct {
	ip_addr *_dafny.Array
	port    uint16
}

// TONY : DONE
func UDPAddrToIPEndPoint(udpAddr *net.UDPAddr) *IPEndPoint {
	var port = uint16(udpAddr.Port)
	var byteIPArr = []byte(udpAddr.IP.To4())
	var interfaceIPArray []interface{}
	for _, value := range byteIPArr {
		interfaceIPArray = append(interfaceIPArray, interface{}(value))
	}
	var ip = _dafny.NewArrayWithValues(interfaceIPArray...)
	var res = IPEndPoint{ip, port}
	return &res
}

// GetUDPAddr returns the address of this endpoint as a net.UDPAddr data structure
// TONY : DONE
func (ep *IPEndPoint) GetUDPAddr() *net.UDPAddr {
	// First get the string of ip:port
	var ipArrStr = ep.ip_addr.String()
	var intArr []int
	err := json.Unmarshal([]byte(ipArrStr), &intArr)
	if err != nil {
		fmt.Printf("Cannot unmarshal %v\n", ipArrStr)
		log.Fatal(err)
	}
	var ip = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(intArr)), "."), "[]")
	var ipAndPortStr = ip + ":" + strconv.FormatUint(uint64(ep.port), 10)
	fmt.Printf("%v\n", ipAndPortStr)

	// Next convert to net.UDPAddr
	udpAddr, err := net.ResolveUDPAddr("udp", ipAndPortStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
	return udpAddr
}

// TONY : DONE
func (ep *IPEndPoint) GetAddress() *_dafny.Array {
	return ep.ip_addr
}

// TONY : DONE
func (ep *IPEndPoint) GetPort() uint16 {
	return ep.port
}

type CompanionStruct_IPEndPoint_ struct {
}

var Companion_IPEndPoint_ = CompanionStruct_IPEndPoint_{}

// TONY : DONE
func (comp_ep *CompanionStruct_IPEndPoint_) Construct(ip_addr *_dafny.Array, port uint16) (bool, *IPEndPoint) {
	res := &IPEndPoint{ip_addr, port}
	return true, res
}

func (_this *IPEndPoint) Equals(other *IPEndPoint) bool {
	return _this == other
}

func (_this *IPEndPoint) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*IPEndPoint)
	return ok && _this.Equals(other)
}

func (*IPEndPoint) String() string {
	return "_9_Native____Io__s_Compile.IPEndPoint"
}

func Type_IPEndPoint_() _dafny.Type {
	return type_IPEndPoint_{}
}

type type_IPEndPoint_ struct {
}

func (_this type_IPEndPoint_) Default() interface{} {
	return (*IPEndPoint)(nil)
}

func (_this type_IPEndPoint_) String() string {
	return "_9_Native____Io__s_Compile.IPEndPoint"
}

// End of class IPEndPoint

// Definition of class Packet
type Packet struct {
	ep     *IPEndPoint
	buffer []byte
}

// Definition of class UdpClient
// TONY : DONE
type UdpClient struct {
	localEndpoint *IPEndPoint
	connection    *net.UDPConn
	send_queue    goconcurrentqueue.Queue
	receive_queue goconcurrentqueue.Queue
}

type CompanionStruct_UdpClient_ struct {
}

var Companion_UdpClient_ = CompanionStruct_UdpClient_{}

// TONY : DONE
func new_UdpClient_(my_ep *IPEndPoint, conn *net.UDPConn) *UdpClient {
	// Initialize record and start send and receive loops
	_this := UdpClient{
		localEndpoint: my_ep,
		connection:    conn,
		send_queue:    goconcurrentqueue.NewFIFO(),
		receive_queue: goconcurrentqueue.NewFIFO(),
	}
	fmt.Printf("Starting new UDPClient %v\n", conn.LocalAddr())
	go _this.sendLoop()
	go _this.receiveLoop()
	return &_this
}

// TONY : DONE
func (comp_udpclient *CompanionStruct_UdpClient_) Construct(localEndpoint *IPEndPoint) (bool, *UdpClient) {
	var localEp = localEndpoint.GetUDPAddr()
	conn, err := net.ListenUDP("udp", localEp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		return false, nil
	}
	var udp = new_UdpClient_(localEndpoint, conn)
	return true, udp
}

// TONY : DONE
func (client *UdpClient) sendLoop() {
	for true {
		var packInterface, _ = client.send_queue.DequeueOrWaitForNextElement()
		var pack, ok = packInterface.(Packet)
		// fmt.Printf("TONY DEBUG: sendLoop() found a packet with dest %v and contents %v\n", pack.ep.GetUDPAddr(), pack.buffer)
		if !ok {
			fmt.Printf("Fatal error: Cannot convert %v to Packet\n", pack)
			os.Exit(1)
		}
		var _, err2 = client.connection.WriteToUDP(pack.buffer, pack.ep.GetUDPAddr())
		// fmt.Printf("TONY DEBUG: sendLoop() sent %v bytes over UDP to %v\n", n, pack.ep.GetUDPAddr())
		if err2 != nil {
			fmt.Printf("Fatal error %s", err2.Error())
			os.Exit(1)
		}
	}
}

// TONY : DONE
func (client *UdpClient) receiveLoop() {
	// Read from UDP connection, initialize packet and enqueue to receive_queue
	// fmt.Printf("TONY DEBUG: starting receiveLoop()\n")
	for true {
		var buffer [16]byte
		// TONY: There is a Golang bug on OSX where ReadFromUDP does not block, but should work fine on Linux
		var _, addr, err = client.connection.ReadFromUDP(buffer[0:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
			os.Exit(1)
		}
		if addr != nil {
			var packetEp = UDPAddrToIPEndPoint(addr)
			var packet = Packet{packetEp, buffer[0:]}
			// fmt.Printf("TONY DEBUG: receiveLoop() found a packet with source %v and contents %v \n", addr, packet.buffer)
			client.receive_queue.Enqueue(packet)
		}
	}
}

// TONY : DONE
func (client *UdpClient) Send(remote *IPEndPoint, buffer *_dafny.Array) bool {
	// Create Packet struct and enqueue to send_queue
	var bufferStr = buffer.String()
	var bufferByte []byte
	err := json.Unmarshal([]byte(bufferStr), &bufferByte)
	if err != nil {
		log.Fatal(err)
	}
	var packet = Packet{remote, bufferByte}
	client.send_queue.Enqueue(packet)
	return true
}

// TONY : DONE
func (client *UdpClient) Receive(timeLimit int32) (bool, bool, *IPEndPoint, *_dafny.Array) {
	// Note that in Toylock, this is only ever called with timeout 0
	var packet, err = client.receive_queue.Dequeue()
	if err != nil {
		// receive queue is empty
		// fmt.Printf("TONY DEBUG: receive_queue empty\n")
		if timeLimit == 0 {
			return true, true, nil, nil
		} else {
			fmt.Printf("Going to sleep unexpectedly!")
			time.Sleep(time.Duration(timeLimit) * time.Millisecond)
			return client.Receive(0)
		}
	} else {
		var pack, ok = packet.(Packet)
		if !ok {
			fmt.Fprintf(os.Stderr, "Fatal error: Cannot convert %v to Packet\n", pack)
			os.Exit(1)
		}
		// var buf = pack.buffer
		// var addr = pack.ep.GetUDPAddr()
		// fmt.Printf("TONY DEBUG: received a packet with source %v and contents %v: \n", addr, buf)
		var interfaceBuf []interface{}
		for _, value := range pack.buffer {
			interfaceBuf = append(interfaceBuf, interface{}(value))
		}
		return true, false, pack.ep, _dafny.NewArrayWithValues(interfaceBuf...)
	}
}

func (_this *UdpClient) Equals(other *UdpClient) bool {
	return _this == other
}

func (_this *UdpClient) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*UdpClient)
	return ok && _this.Equals(other)
}

func (*UdpClient) String() string {
	return "_9_Native____Io__s_Compile.UdpClient"
}

func Type_UdpClient_() _dafny.Type {
	return type_UdpClient_{}
}

type type_UdpClient_ struct {
}

func (_this type_UdpClient_) Default() interface{} {
	return (*UdpClient)(nil)
}

func (_this type_UdpClient_) String() string {
	return "_9_Native____Io__s_Compile.UdpClient"
}
func (_this *UdpClient) Ctor__() {
	goto TAIL_CALL_START
TAIL_CALL_START:
}

// End of class UdpClient

// Definition of class FileSystemState
type FileSystemState struct {
	dummy byte
}

func New_FileSystemState_() *FileSystemState {
	_this := FileSystemState{}

	return &_this
}

type CompanionStruct_FileSystemState_ struct {
}

var Companion_FileSystemState_ = CompanionStruct_FileSystemState_{}

func (_this *FileSystemState) Equals(other *FileSystemState) bool {
	return _this == other
}

func (_this *FileSystemState) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*FileSystemState)
	return ok && _this.Equals(other)
}

func (*FileSystemState) String() string {
	return "_9_Native____Io__s_Compile.FileSystemState"
}

func Type_FileSystemState_() _dafny.Type {
	return type_FileSystemState_{}
}

type type_FileSystemState_ struct {
}

func (_this type_FileSystemState_) Default() interface{} {
	return (*FileSystemState)(nil)
}

func (_this type_FileSystemState_) String() string {
	return "_9_Native____Io__s_Compile.FileSystemState"
}

// End of class FileSystemState

// Definition of class MutableSet
type MutableSet struct {
	Type_T_ _dafny.Type
	dummy   byte
}

func New_MutableSet_(Type_T_ _dafny.Type) *MutableSet {
	_this := MutableSet{}

	_this.Type_T_ = Type_T_
	return &_this
}

type CompanionStruct_MutableSet_ struct {
}

var Companion_MutableSet_ = CompanionStruct_MutableSet_{}

func (_this *MutableSet) Equals(other *MutableSet) bool {
	return _this == other
}

func (_this *MutableSet) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*MutableSet)
	return ok && _this.Equals(other)
}

func (*MutableSet) String() string {
	return "_9_Native____Io__s_Compile.MutableSet"
}

func Type_MutableSet_(Type_T_ _dafny.Type) _dafny.Type {
	return type_MutableSet_{Type_T_}
}

type type_MutableSet_ struct {
	Type_T_ _dafny.Type
}

func (_this type_MutableSet_) Default() interface{} {
	return (*MutableSet)(nil)
}

func (_this type_MutableSet_) String() string {
	return "_9_Native____Io__s_Compile.MutableSet"
}

// End of class MutableSet

// Definition of class MutableMap
type MutableMap struct {
	Type_K_ _dafny.Type
	Type_V_ _dafny.Type
	dummy   byte
}

func New_MutableMap_(Type_K_ _dafny.Type, Type_V_ _dafny.Type) *MutableMap {
	_this := MutableMap{}

	_this.Type_K_ = Type_K_
	_this.Type_V_ = Type_V_
	return &_this
}

type CompanionStruct_MutableMap_ struct {
}

var Companion_MutableMap_ = CompanionStruct_MutableMap_{}

func (_this *MutableMap) Equals(other *MutableMap) bool {
	return _this == other
}

func (_this *MutableMap) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*MutableMap)
	return ok && _this.Equals(other)
}

func (*MutableMap) String() string {
	return "_9_Native____Io__s_Compile.MutableMap"
}

func Type_MutableMap_(Type_K_ _dafny.Type, Type_V_ _dafny.Type) _dafny.Type {
	return type_MutableMap_{Type_K_, Type_V_}
}

type type_MutableMap_ struct {
	Type_K_ _dafny.Type
	Type_V_ _dafny.Type
}

func (_this type_MutableMap_) Default() interface{} {
	return (*MutableMap)(nil)
}

func (_this type_MutableMap_) String() string {
	return "_9_Native____Io__s_Compile.MutableMap"
}

// End of class MutableMap

// Definition of class Arrays
type Arrays struct {
	dummy byte
}

func New_Arrays_() *Arrays {
	_this := Arrays{}

	return &_this
}

type CompanionStruct_Arrays_ struct {
}

var Companion_Arrays_ = CompanionStruct_Arrays_{}

func (_this *Arrays) Equals(other *Arrays) bool {
	return _this == other
}

func (_this *Arrays) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*Arrays)
	return ok && _this.Equals(other)
}

func (*Arrays) String() string {
	return "_9_Native____Io__s_Compile.Arrays"
}

func Type_Arrays_() _dafny.Type {
	return type_Arrays_{}
}

type type_Arrays_ struct {
}

func (_this type_Arrays_) Default() interface{} {
	return (*Arrays)(nil)
}

func (_this type_Arrays_) String() string {
	return "_9_Native____Io__s_Compile.Arrays"
}

// End of class Arrays
