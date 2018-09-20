// Package multicast  sends packets to all devices
// in a specified group. Membership in a group is set up when devices
// send "join" packets to an upstream router, and routers and switches
//keep track of this membership. When multicast packets arrive at a switch,
//they are only sent to devices or segments (such as WiFi) where at least one device wants them.
// Multicast can traverse the networks where it has been configured.
// Author by dmichael @ https://github.com/dmichael/go-multicast
package multicast

import (
	"log"
	"net"
)

const (
	maxDatagramSize = 8192
)

// Listen binds to the UDP address and port given and writes packets received
// from that address to a buffer which is passed to a hander
func Listen(address string, handler func(*net.UDPAddr, int, []byte)) error {
	// Parse the string address
	group, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Println("UDP address can't resolved by listener: ", err)
		return err
	}

	// Open up a connection
	conn, err := net.ListenMulticastUDP("udp", nil, group)
	if err != nil {
		log.Println("ListenMulticastUDP can't open the connection: ", err)
		return err
	}

	conn.SetReadBuffer(maxDatagramSize)

	// Loop forever reading from the socket
	for {
		buffer := make([]byte, maxDatagramSize)
		numBytes, src, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println("ReadFromUDP failed: ", err)
		}

		handler(src, numBytes, buffer)
	}
}
