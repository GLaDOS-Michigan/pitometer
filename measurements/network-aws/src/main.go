package main

import (
	"agents"
	"clock"
	"fmt"
	"math/rand"
	"native"
	"net"
	"os"
	"strconv"
	"time"
)

// BaseClientPort is the first client port to use. The i^th client agent uses port BaseClientPort + i
const BaseClientPort uint64 = 5000

// Debug mode
const Debug bool = false

// TimeOut is the period after which a client retries if failed to receive response
// In all likelihood, the packet is dropped by the network
var TimeOut time.Duration = 20 * time.Second

func main() {
	// This program takes the following positional arguments
	// Each agent only needs to know a priori the target servers' addresses
	// The local client's port is dynamic based on how many client agents I spawn, let's reserve 5000-5999 for now
	// <target1 ip:s_port> <target2 ip:s_port> ... <local ip:s_port> <interval> <payload_sz>

	native.DebugMode = Debug
	agents.ClientTimeOut = TimeOut

	// Pop the next argument from os.Args -- that is the payload size
	var payloadSz, err2 = strconv.ParseInt(os.Args[len(os.Args)-1], 10, 64)
	if int(payloadSz) < 0 || err2 != nil {
		fmt.Printf("Error: Invalid payloadSz %v\n", os.Args[len(os.Args)-1])
		os.Exit(1)
	}
	os.Args = os.Args[:len(os.Args)-1]

	// Pop the next argument from os.Args -- that is the ping interval
	var interval, err3 = strconv.ParseInt(os.Args[len(os.Args)-1], 10, 64)
	if int(interval) < 0 || err3 != nil {
		fmt.Printf("Error: Invalid interval %v\n", os.Args[len(os.Args)-1])
		os.Exit(1)
	}
	os.Args = os.Args[:len(os.Args)-1]

	// Pop the next argument from os.Args -- that is my server agent's address
	var localServerAddr, err4 = net.ResolveUDPAddr("udp", os.Args[len(os.Args)-1])
	if err4 != nil {
		fmt.Printf("Error: Failed to resolve address %v\n", os.Args[len(os.Args)-1])
		os.Exit(1)
	}
	if 5000 <= localServerAddr.Port && localServerAddr.Port < 6000 {
		fmt.Printf("Error: Ports 5000-5999 are reserved and cannot be used as local server port\n")
		os.Exit(1)
	}
	os.Args = os.Args[:len(os.Args)-1]

	// The remaining arguments are now the target server addresses
	var targetServerAddresses = make([]*net.UDPAddr, 0)
	for _, targetServerStr := range os.Args[1:] {
		var targetServerAddr, err = net.ResolveUDPAddr("udp", targetServerStr)
		if err != nil {
			fmt.Printf("Error: Failed to resolve target server address %v\n", targetServerStr)
			os.Exit(1)
		}
		if 5000 <= targetServerAddr.Port && targetServerAddr.Port < 6000 {
			fmt.Printf("Error: Ports 5000-5999 are reserved and cannot be server ports\n")
			os.Exit(1)
		}
		targetServerAddresses = append(targetServerAddresses, targetServerAddr)
	}

	fmt.Printf("Initializing Network Agent with the following parameters:\n")
	fmt.Printf("    payload size  = %v bytes\n", payloadSz)
	fmt.Printf("    ping interval = %v milliseconds\n", interval)
	fmt.Printf("    local server  = %v\n", localServerAddr)
	fmt.Printf("    targets       = %v\n", targetServerAddresses)

	// Start local server
	var localServerAgent = agents.Server{LocalAddr: localServerAddr}
	go localServerAgent.StartServerLoop()

	// Wait some time for all clients to start
	var waitTime = time.Duration(2+len(targetServerAddresses)) * time.Second
	native.Debug(fmt.Sprintf("Wait for %v seconds\n", waitTime))
	time.Sleep(waitTime)

	// Start all clients
	var localIP = localServerAddr.IP
	var clientsMap = make(map[uint64]*agents.Client) // map from local port used, to the client agents

	for i, targetAddr := range targetServerAddresses {
		rand.Seed(time.Now().UnixNano())
		var clientPort uint64
		if localIP.IsLoopback() {
			// If running a local experiment, pick random client port so that they are
			// unlikely to clash. This is really janky, but ok since this is not meant to
			// be run locally anyways
			clientPort = BaseClientPort + uint64(i) + uint64(rand.Intn(990))
		} else {
			clientPort = BaseClientPort + uint64(i)
		}
		var clientAddr = &net.UDPAddr{IP: localIP, Port: int(clientPort)}
		var localClientAgent = &agents.Client{
			LocalAddr:    clientAddr,
			Target:       targetAddr,       // remote address to send packet
			Interval:     uint64(interval), // milliseconds to sleep in between pings
			PacketSize:   uint64(payloadSz),
			PingLog:      clock.NewStopwatch(fmt.Sprintf("Ping Stopwatch from %v to %v", clientAddr.IP, targetAddr.IP)),
			TimeoutCount: clock.NewCounter("Timeouts")}
		clientsMap[clientPort] = localClientAgent
	}

	// Start all local clients
	for _, clientAgent := range clientsMap {
		go clientAgent.StartClientLoop()
	}

	// Start experiment timer
	fmt.Printf("Starting experiment at %v\n", time.Now())
	for true {
		time.Sleep(30 * time.Second)
	}
}