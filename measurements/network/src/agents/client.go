package agents

import (
	"clock"
	"fmt"
	native "native"
	"net"
	"os"
	"time"
)

// Client is an agent that sends UDP packets and times their round-trip time (RTT)
type Client struct {
	LocalAddr  *net.UDPAddr
	Target     *net.UDPAddr     // remote address to send packet
	Interval   uint64           // milliseconds to sleep in between pings
	PacketSize uint64           // size of UDP payload to send
	PingLog    *clock.Stopwatch // record of ping events
	active     bool
}

// StartClientLoop is the main event loop of the Client. It receives UDP packets and
// bounces them back to the sender
func (c *Client) StartClientLoop() {
	var ok, udpClient = native.NewUDPClient(native.UDPAddrToIPEndPoint(c.LocalAddr))

	if !ok {
		fmt.Printf("Error: failed to initialize client\n")
		os.Exit(1)
	}

	// Craft packet and initialize dest
	var payload = make([]byte, c.PacketSize, c.PacketSize)
	var dest = native.UDPAddrToIPEndPoint(c.Target)
	var pack = &native.Packet{EndPoint: dest, Buffer: payload[0:]}
	c.active = true

	fmt.Printf("Starting new client at %v targeting %v\n", c.LocalAddr, c.Target)
	// Main event loop
	for c.active {

		// Send packet
		native.Debug(fmt.Sprintf("Client %v sending %v", c.LocalAddr, pack))
		udpClient.Send(pack)
		c.PingLog.LogStartEvent("Send ping")

		// Receive packet
		_, _, remote, receivedPacket := udpClient.Receive()
		native.Debug(fmt.Sprintf("Client %v received response from %v, %v", c.LocalAddr, remote.GetUDPAddr(), receivedPacket))
		c.PingLog.LogStartEvent("Receive reply")

		if len(receivedPacket.Buffer) != int(c.PacketSize) {
			fmt.Printf("Error: got packet length %v, expected %v\n",
				len(receivedPacket.Buffer),
				c.PacketSize)
			os.Exit(1)
		}

		// Sleep
		time.Sleep(time.Duration(c.Interval) * time.Millisecond)
	}
	fmt.Printf("Client at %v targeting %v deactivated\n", c.LocalAddr, c.Target)
}

// StopClientLoop stops the main event loop of the Client.
func (c *Client) StopClientLoop() {
	c.active = false
}
