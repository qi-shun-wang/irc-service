package app

import (
	core "IRCService/app/core"
	multicastProvider "IRCService/app/provider"
	"log"

	coap "github.com/dustin/go-coap"
)

var (
	mux = coap.NewServeMux()
)

const (
	coapPort                = "5683"
	defaultMulticastAddress = "239.0.0.0:9999"
)

//Run app.
func Run() {

	log.Print("Muticast Provider Pinger On Start....")
	go multicastProvider.RunPinger(defaultMulticastAddress)
	go multicastProvider.RunListenner(defaultMulticastAddress)
	log.Print("Kod_Coap On Start....")
	log.Print(coap.ListenAndServe("udp", ":"+coapPort, mux))

}

//Setup everything before call Run for running app.
func Setup() {
	c := core.Commander{}
	setRouters(&c)
}
