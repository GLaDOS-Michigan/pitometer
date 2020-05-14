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
	var ok, udpClient = native.NewUDPClient(native.UDPAddrToIPEndPoint(s.LocalAddr))

	if !ok {
		fmt.Printf("Error: failed to initialize server\n")
		os.Exit(1)
	}

	// Main event loop
	for true {
		_, _, remote, packet := udpClient.Receive()
		packet.Dest = remote // Do the dest switcheroo
		udpClient.Send(packet)
	}
}
