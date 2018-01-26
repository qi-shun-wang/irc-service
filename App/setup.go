package app

import (
	core "IRCService/app/core"
	"log"

	coap "github.com/dustin/go-coap"
)

var (
	mux      = coap.NewServeMux()
	coapPort = "5683"
)

//Run app.
func Run() {

	log.Print("Kod_Coap On Start....")
	log.Fatal(coap.ListenAndServe("udp", ":"+coapPort, mux))
	// go serveMulticastUDP(srvAddr, msgHandler)
	// go coap.ListenAndServe("udp", ":5684", coap.FuncHandler(observeHandler))
}

//Setup .
func Setup() {
	c := core.Commander{}
	setRouters(&c)
}
