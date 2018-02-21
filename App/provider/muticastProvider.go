package provider

import (
	"IRCService/App/provider/multicast"
	model "IRCService/app/model"
	"encoding/hex"
	"log"
	"net"
	"time"
)

//RunListenner .
func RunListenner(defaultMulticastAddress string) {
	log.Printf("Listening on %s\n", defaultMulticastAddress)
	multicast.Listen(defaultMulticastAddress, msgHandler)
}

func msgHandler(src *net.UDPAddr, n int, b []byte) {
	log.Println(n, "bytes read from", src)
	log.Println(hex.Dump(b[:n]))
}

//RunPinger .
func RunPinger(defaultMulticastAddress string) {
	log.Printf("Broadcasting to %s\n", defaultMulticastAddress)
	ping(defaultMulticastAddress)
}

func ping(addr string) {
	conn, err := multicast.NewBroadcaster(addr)
	if err != nil {
		log.Print(err)
	}

	for {
		withInfo := model.Prepare().ToJSONString()
		conn.Write([]byte(withInfo))
		time.Sleep(1 * time.Second)
	}
}
