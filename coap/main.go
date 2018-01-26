package coap

import (
	"log"

	coap "github.com/dustin/go-coap"
)

func init() {
	c := commander{}

	setHandler("cmd", &c)
	log.Print("Kod_Coap On Start....")
	go serveMulticastUDP(srvAddr, msgHandler)
	go coap.ListenAndServe("udp", ":5684", coap.FuncHandler(observeHandler))
	runCoap()

}
