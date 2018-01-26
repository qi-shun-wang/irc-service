package controllers

import (
	core "IRCService/app/core"
	"log"
	"net"

	coap "github.com/dustin/go-coap"
)

//KeyEventHandler .
func KeyEventHandler(ci core.CoapInterface) core.CoapHandler {
	return func(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {

		number := string(m.Payload)
		log.Println(string(number))
		cmds := parsedKeySerial(number)
		ci.OnCmds(cmds)
		return nil
	}
}

func parsedKeySerial(number string) string {
	cmds := ""
	return cmds
}
