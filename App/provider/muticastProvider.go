package provider

import (
	"IRCService/App/provider/multicast"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"time"
)

//RunListenner .
func RunListenner(defaultMulticastAddress string) {
	fmt.Printf("Listening on %s\n", defaultMulticastAddress)
	multicast.Listen(defaultMulticastAddress, msgHandler)
}

func msgHandler(src *net.UDPAddr, n int, b []byte) {
	log.Println(n, "bytes read from", src)
	log.Println(hex.Dump(b[:n]))
}

//RunPinger .
func RunPinger(defaultMulticastAddress string, withInfo string) {
	fmt.Printf("Broadcasting to %s\n", defaultMulticastAddress)
	ping(defaultMulticastAddress, withInfo)
}

func ping(addr string, withInfo string) {
	conn, err := multicast.NewBroadcaster(addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn.Write([]byte(withInfo))
		time.Sleep(1 * time.Second)
	}
}
