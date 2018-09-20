package controllers

import (
	"IRCService/app/core"
	"github.com/dustin/go-coap"
	"log"
	"net"
	"strings"
)

//TextInputHandler .
func TextInputHandler(ci core.CoapInterface) core.CoapHandler {
	return func(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {
		text := string(m.Payload)
		log.Println(string(text))
		cmds := parsedInput(text)
		ci.OnCmds(cmds)
		return nil
	}
}

func parsedInput(text string) string {
	cmds := []string{}

	cmds = append(cmds, "input text "+text+";")

	log.Println("input text " + text + ";")
	log.Println()
	fullCmds := strings.Join(cmds, ";")
	log.Println(fullCmds)
	return fullCmds

}
