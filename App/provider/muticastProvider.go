package provider

import (
	"IRCService/App/provider/multicast"
	model "IRCService/app/model"
	"log"
	"net"
	"time"
)

var device = model.Device{}

//RunListenner .
func RunListenner(defaultMulticastAddress string) {
	log.Print("Muticast Provider RunListenner On Start....")
	log.Printf("Listening on %s\n", defaultMulticastAddress)
	err := multicast.Listen(defaultMulticastAddress, msgHandler)
	defer func() {
		if r := recover(); r != nil {
			log.Println("Listener  will be recovered in ", r)
			time.Sleep(1 * time.Second)
			RunListenner(defaultMulticastAddress)
		}
	}()
	if err != nil {
		panic("Wifi conn is broken...")
	}
}

func msgHandler(src *net.UDPAddr, n int, b []byte) {
	// log.Println(n, "bytes read from", src)
	// log.Println(string(hex.Dump(b[:n])))
}

//RunPinger .
func RunPinger(defaultMulticastAddress string) {
	log.Print("Muticast Provider Pinger On Start....")
	log.Printf("Broadcasting to %s\n", defaultMulticastAddress)
	err := ping(defaultMulticastAddress)
	defer func() {
		if r := recover(); r != nil {
			log.Println("Pinger will be recovered in ", r)
			time.Sleep(1 * time.Second)
			RunPinger(defaultMulticastAddress)
		}
	}()
	if err != nil {
		log.Print(err)
		panic("Wifi conn is broken...")
	}
}

func ping(addr string) error {
	conn, err := multicast.NewBroadcaster(addr)
	if err != nil {
		log.Print(err)
		return err
	}

	for {
		// log.Println("device.Prepare....")
		device = model.Prepare()
		// log.Println(device.ToJSONString())
		withInfo := device.ToJSONString()

		_, err := conn.Write([]byte(withInfo))
		if err != nil {
			log.Print(err)
			return err
		}
	}
}
