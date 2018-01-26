package app

import (
	core "IRCService/app/core"
	model "IRCService/app/model"
	multicastProvider "IRCService/app/provider"
	"log"

	coap "github.com/dustin/go-coap"
)

var (
	mux    = coap.NewServeMux()
	device = model.Device{}
)

const (
	coapPort                = "5683"
	defaultMulticastAddress = "239.0.0.0:9999"
)

//Run app.
func Run() {

	log.Print("Muticast Provider Pinger On Start....")
	go multicastProvider.RunPinger(defaultMulticastAddress, device.ToJSONString())
	log.Print("Kod_Coap On Start....")
	log.Fatal(coap.ListenAndServe("udp", ":"+coapPort, mux))

}

//Setup everything before call Run for running app.
func Setup() {
	c := core.Commander{}
	device = model.Prepare()
	setRouters(&c)
}
