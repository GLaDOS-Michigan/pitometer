package agents

import (
	"clock"
	"fmt"
	native "native"
	"net"
	"os"
	"time"
)

// ClientTimeOut duration
var ClientTimeOut time.Duration

// Client is an agent that sends UDP packets and times their round-trip time (RTT)
type Client struct {
	LocalAddr    *net.UDPAddr
	Targets      []*net.UDPAddr   // remote addresses to send packet
	Interval     uint64           // milliseconds to sleep in between pings
	PacketSize   uint64           // size of UDP payload to send
	PingLog      *clock.Stopwatch // record of ping events
	TimeoutCount *clock.Counter   // count timeout events
	active       bool
}

// StartClientLoop is the main event loop of the Client. It receives UDP packets and
// bounces them back to the sender
func (c *Client) StartClientLoop() {
	var ok, udpClient = native.NewUDPClient(native.UDPAddrToIPEndPoint(c.LocalAddr), 1, 1)

	if !ok {
		fmt.Printf("Error: failed to initialize client\n")
		os.Exit(1)
	}

	// Craft packet and initialize dest
	var payload = make([]byte, c.PacketSize, c.PacketSize)
	var dests = make([]*native.IPEndPoint, 0)
	for _, dst := range c.Targets {
		dests = append(dests, native.UDPAddrToIPEndPoint(dst))
	}

	var packs = make([]*native.Packet, 0)
	for _, dst := range dests {
		packs = append(packs, &native.Packet{EndPoint: dst, Buffer: payload[0:]})
	}
	c.active = true

	fmt.Printf("Starting new client at %v targeting %v\n", c.LocalAddr, c.Targets)
	var sendNote = fmt.Sprintf("Send to targets,%v", c.Targets)
	var receiveNote = fmt.Sprintf("Receive from targets,%v", c.Targets)
	// Main event loop
	for c.active {

		// Send packet
		c.PingLog.LogStartEvent(sendNote)
		for _, p := range packs {
			udpClient.Send(p)
		}

		// Receive packet
		var remote *native.IPEndPoint = nil
		var receivedPacket *native.Packet = nil
		var timedOutChan = make(chan bool, 2)
		go func(c *Client, timedOutChan chan bool) {
			var n = len(c.Targets) // number of targets from which to wait or a response
			for n > 0 {
				_, _, remote, receivedPacket = udpClient.Receive()
				native.Debug(fmt.Sprintf("Client %v received response from %v, %v", c.LocalAddr, remote.GetUDPAddr(), receivedPacket))
				n -= 1
			}
			c.PingLog.LogEndEvent(receiveNote)
			timedOutChan <- false
		}(c, timedOutChan)
		go func(timedOutChan chan bool) {
			time.Sleep(ClientTimeOut)
			timedOutChan <- true
		}(timedOutChan)

		var timedOut = <-timedOutChan

		if timedOut {
			native.Debug("Timed out!")
			ok, err := udpClient.ReceiveQueue.DequeueOrWaitForNextElementCancel()
			if !ok {
				fmt.Printf("Error: DequeueOrWaitForNextElementCancel failed. %v\n", err)
				os.Exit(1)
			}
			c.PingLog.PopStartEvent()
			c.TimeoutCount.Increment()
			continue
		}
		if len(receivedPacket.Buffer) != int(c.PacketSize) {
			fmt.Printf("Error: got packet length %v, expected %v\n",
				len(receivedPacket.Buffer),
				c.PacketSize)
			os.Exit(1)
		}

		// Sleep
		time.Sleep(time.Duration(c.Interval) * time.Millisecond)
	}
	fmt.Printf("Client at %v targeting %v deactivated\n", c.LocalAddr, c.Targets)
}

// StopClientLoop stops the main event loop of the Client.
func (c *Client) StopClientLoop() {
	c.active = false
}
