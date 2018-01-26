package coap

import (
	"encoding/hex"
	"log"
	"net"
)

const (
	srvAddr         = "224.0.0.1:9999"
	maxDatagramSize = 8192
)

func msgHandler(src *net.UDPAddr, n int, b []byte) {
	log.Println(n, "bytes read from", src)
	log.Println(hex.Dump(b[:n]))
}

func serveMulticastUDP(address string, handler func(*net.UDPAddr, int, []byte)) {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatal(err)
	}
	connection, err := net.ListenMulticastUDP("udp", nil, addr)
	connection.SetReadBuffer(maxDatagramSize)
	for {
		bytes := make([]byte, maxDatagramSize)
		n, src, err := connection.ReadFromUDP(bytes)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		}
		handler(src, n, bytes)
	}
}
