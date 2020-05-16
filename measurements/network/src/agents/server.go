package agents

import (
	"fmt"
	native "native"
	"net"
	"os"
)

// Server is an agent that receives UDP packets and bounces them back to the sender
type Server struct {
	LocalAddr *net.UDPAddr
}

// StartServerLoop is the main event loop of the Server. It receives UDP packets and
// bounces them back to the sender
func (s *Server) StartServerLoop() {
	var ok, udpClient = native.NewUDPClient(native.UDPAddrToIPEndPoint(s.LocalAddr), 1, 1)

	if !ok {
		fmt.Printf("Error: failed to initialize server\n")
		os.Exit(1)
	}

	fmt.Printf("Starting new server at %v\n", s.LocalAddr)
	// Main event loop
	for true {
		_, _, remote, packet := udpClient.Receive()
		native.Debug(fmt.Sprintf("Server %v received packet from %v, %v", s.LocalAddr, remote.GetUDPAddr(), packet))
		packet.EndPoint = remote // Do the dest switcheroo
		native.Debug(fmt.Sprintf("Server %v responding %v with %v", s.LocalAddr, remote.GetUDPAddr(), packet))
		udpClient.Send(packet)
	}
}
