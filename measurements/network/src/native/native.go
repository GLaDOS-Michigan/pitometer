package native

import (
	"dafny"
	"encoding/json"
	"fmt"
	"goconcurrentqueue"
	"log"
	"net"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
)

/*****************************************************************************************
*                                         Debug                                          *
*****************************************************************************************/

// DebugMode toggle
var DebugMode bool

// Debug prints
func Debug(msg string) {
	if DebugMode {
		fmt.Println(msg)
	}
}

/*****************************************************************************************
*                                    Class Packet                                        *
*****************************************************************************************/

// Packet to be sent on the network
type Packet struct {
	// When receiving a packet, EndPoint the SOURCE
	// When sending a packet, EndPoint the DEST
	EndPoint *IPEndPoint
	Buffer   []byte
}

// String formats Packet into a string
func (p *Packet) String() string {
	return fmt.Sprintf("Packet{ dest %v, size %v }", p.EndPoint.GetUDPAddr(), len(p.Buffer))
}

/*****************************************************************************************
*                                  Class IPEndPoint                                      *
*****************************************************************************************/

// IPEndPoint represents an endpoint address
type IPEndPoint struct {
	ipAddr *dafny.Array
	port   uint16
}

// UDPAddrToIPEndPoint doc
func UDPAddrToIPEndPoint(udpAddr *net.UDPAddr) *IPEndPoint {
	var port = uint16(udpAddr.Port)
	var byteIPArr = []byte(udpAddr.IP.To4())
	var interfaceIPArray []interface{}
	for _, value := range byteIPArr {
		interfaceIPArray = append(interfaceIPArray, interface{}(value))
	}
	var ip = dafny.NewArrayWithValues(interfaceIPArray...)
	var res = IPEndPoint{ip, port}
	return &res
}

// GetUDPAddr returns the address of this endpoint as a net.UDPAddr data structure
func (ep *IPEndPoint) GetUDPAddr() *net.UDPAddr {
	// First get the string of ip:port
	var ipArrStr = ep.ipAddr.String()
	var intArr []int
	err := json.Unmarshal([]byte(ipArrStr), &intArr)
	if err != nil {
		fmt.Printf("Cannot unmarshal %v\n", ipArrStr)
		log.Fatal(err)
	}
	var ip = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(intArr)), "."), "[]")
	var ipAndPortStr = ip + ":" + strconv.FormatUint(uint64(ep.port), 10)

	// Next convert to net.UDPAddr
	udpAddr, err := net.ResolveUDPAddr("udp", ipAndPortStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		debug.PrintStack()
		os.Exit(1)
	}
	return udpAddr
}

// GetAddress doc
func (ep *IPEndPoint) GetAddress() *dafny.Array {
	return ep.ipAddr
}

// GetPort doc
func (ep *IPEndPoint) GetPort() uint16 {
	return ep.port
}

// NewIPEndPoint is the constructor for IPEndPoint
func NewIPEndPoint(ipAddr *dafny.Array, port uint16) (bool, *IPEndPoint) {
	res := &IPEndPoint{ipAddr, port}
	return true, res
}

/*****************************************************************************************
*                                   Class UdpClient                                      *
*****************************************************************************************/

// UDPClient represents a client
type UDPClient struct {
	localEndpoint *IPEndPoint
	connection    *net.UDPConn
	sendQueue     goconcurrentqueue.Queue
	receiveQueue  goconcurrentqueue.Queue
}

// TONY : DONE
func newUDPClient(myEP *IPEndPoint, conn *net.UDPConn) *UDPClient {
	// Initialize record and start send and receive loops
	_this := UDPClient{
		localEndpoint: myEP,
		connection:    conn,
		sendQueue:     goconcurrentqueue.NewFIFO(),
		receiveQueue:  goconcurrentqueue.NewFIFO(),
	}
	// fmt.Printf("Starting new UDPClient %v\n", conn.LocalAddr())
	go _this.sendLoop()
	go _this.receiveLoop()
	return &_this
}

// NewUDPClient starts a UDPClient listening at the localEndPoint
func NewUDPClient(localEndpoint *IPEndPoint) (bool, *UDPClient) {
	var localEp = localEndpoint.GetUDPAddr()
	conn, err := net.ListenUDP("udp", localEp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		debug.PrintStack()
		return false, nil
	}
	var udp = newUDPClient(localEndpoint, conn)
	return true, udp
}

func (client *UDPClient) sendLoop() {
	for true {
		var packInterface, _ = client.sendQueue.DequeueOrWaitForNextElement()
		var pack, ok = packInterface.(Packet)
		if !ok {
			fmt.Fprintf(os.Stderr, "Fatal error: Cannot convert %v to Packet\n", pack)
			debug.PrintStack()
			os.Exit(1)
		}
		var _, err2 = client.connection.WriteToUDP(pack.Buffer, pack.EndPoint.GetUDPAddr())
		if err2 != nil {
			fmt.Fprintf(os.Stderr, "Fatal error %s", err2.Error())
			os.Exit(1)
		}
	}
}

func (client *UDPClient) receiveLoop() {
	// Read from UDP connection, initialize packet and enqueue to receive_queue
	// fmt.Printf("TONY DEBUG: starting receiveLoop()\n")
	for true {
		var buffer [4096]byte // max UDP Payload is 65,507 bytes
		// TONY: There is a Golang bug on OSX where ReadFromUDP does not block, but should work fine on Linux
		var n, addr, err = client.connection.ReadFromUDP(buffer[0:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
			os.Exit(1)
		}
		if addr != nil {
			var packetEp = UDPAddrToIPEndPoint(addr)
			var packet = Packet{packetEp, buffer[0:n]}
			client.receiveQueue.Enqueue(packet)
		}
	}
}

// Send a packet to the remote ep
func (client *UDPClient) Send(packet *Packet) bool {
	// Create Packet struct and enqueue to send_queue
	client.sendQueue.Enqueue(*packet)
	return true
}

// Receive blocks until a packet is received
// returns <ok> <timedOut> <remoteEp> <packet>
func (client *UDPClient) Receive() (bool, bool, *IPEndPoint, *Packet) {
	// Note that in Toylock, this is only ever called with timeout 0
	var packet, err = client.receiveQueue.DequeueOrWaitForNextElement()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	var pack, ok = packet.(Packet)
	if !ok {
		fmt.Fprintf(os.Stderr, "Fatal error: Cannot convert %v to Packet\n", pack)
		os.Exit(1)
	}
	return true, false, pack.EndPoint, &pack
}
