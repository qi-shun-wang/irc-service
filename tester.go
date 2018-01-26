package main

import (
	multicast "IRCService/App/provider/multicast"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"time"
)

const (
	defaultMulticastAddress = "239.0.0.0:9999"
)

func main() {
	testListenner()
}

func testPiner() {
	log.Print("Muticast Provider Pinger On Start....")
	fmt.Printf("Broadcasting to %s\n", defaultMulticastAddress)
	ping(defaultMulticastAddress)
}

func testListenner() {
	log.Print("Muticast Provider Listener On Start....")
	fmt.Printf("Listening on %s\n", defaultMulticastAddress)
	multicast.Listen(defaultMulticastAddress, msgHandler)
}

func msgHandler(src *net.UDPAddr, n int, b []byte) {
	log.Println(n, "bytes read from", src)
	log.Println(hex.Dump(b[:n]))
}

func ping(addr string) {
	conn, err := multicast.NewBroadcaster(addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn.Write([]byte("hello, world\n"))
		time.Sleep(1 * time.Second)
	}
}
