package agents

import (
	"fmt"
	native "native"
	"net"
	"os"
	"time"
)

// Client is an agent that sends UDP packets and times their round-trip time (RTT)
type Client struct {
	LocalAddr  *net.UDPAddr
	Target     *net.UDPAddr // remote address to send packet
	Interval   uint64       // milliseconds to sleep in between pings
	PacketSize uint64       // size of UDP payload to send
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
	var pack = &native.Packet{Dest: dest, Buffer: payload[0:]}

	// Main event loop
	for c.active {

		// Send packet
		udpClient.Send(pack)

		// Receive packet
		_, _, remote, receivedPacket := udpClient.Receive()

		// Decode and check for errors
		if !remote.GetUDPAddr().IP.Equal(c.LocalAddr.IP) {
			fmt.Printf("Error: mismatched IP\n")
			os.Exit(1)
		}
		if remote.GetUDPAddr().Port != c.LocalAddr.Port {
			fmt.Printf("Error: mismatched port\n")
			os.Exit(1)
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
	fmt.Printf("Client at %v targeting %v deactivated\n", c.LocalAddr, c.Target)
}

// StopClientLoop stops the main event loop of the Client.
func (c *Client) StopClientLoop() {
	c.active = false
}
