package app

import (
	core "IRCService/app/core"
	multicastProvider "IRCService/app/provider"
	"fmt"
	"log"
	"time"

	coap "github.com/dustin/go-coap"
)

var (
	mux = coap.NewServeMux()
	c   = core.Commander{}
)

const (
	coapPort                = "5683"
	defaultMulticastAddress = "239.0.0.0:9999"
)

func startServcie() {
	err := coap.ListenAndServe("udp", ":"+coapPort, mux)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Coap service will be recovered in ", r)
			time.Sleep(1 * time.Second)
			c.OnCmds("kill -9 $(lsof -t -i:5683)")
			startServcie()
		}
	}()
	if err != nil {
		panic("5683 port is already in use ...")
	}
	log.Println("IRCService Started ....")
}

//Run app.
func Run() {
	go multicastProvider.RunPinger(defaultMulticastAddress)
	go multicastProvider.RunListenner(defaultMulticastAddress)
	log.Println("IRCService On Start....")
	startServcie()

}

//Setup everything before call Run for running app.
func Setup() {
	setRouters(&c)
}
