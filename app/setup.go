package app

import (
	"IRCService/app/core"
	"github.com/dustin/go-coap"
	multicastProvider "IRCService/app/provider"
	"log"
	"time"
)

var (
	mux = coap.NewServeMux()
	c   = core.Commander{}
)

const (
	coapPort                = "5683"
	coapPort2               = "5684"
	defaultMulticastAddress = "239.0.0.0:9999"
	broadcastAddress        = "192.168.34.255:9999"
)

func startServcie(port string) {
	err := coap.ListenAndServe("udp", ":"+port, mux)
	defer func() {
		if r := recover(); r != nil {
			log.Println("Coap service will be recovered in ", r)
			time.Sleep(1 * time.Second)
			c.OnCmds("kill -9 $(lsof -t  /system/bin/IRCService)")
			startServcie(port)
		}
	}()
	if err != nil {
		panic(port + " port is already in use ...")
	}
	log.Println("IRCService Started ....")
}

//Run app.
func Run() {
	log.Println("IRCService On Start....")
	go multicastProvider.RunPinger(broadcastAddress)
	go multicastProvider.RunPinger(defaultMulticastAddress)

	// go multicastProvider.RunListenner(defaultMulticastAddress)
	startServcie(coapPort)
}

//Setup everything before call Run for running app.
func Setup() {
	setRouters(&c)
}
