package controllers

import (
	core "IRCService/app/core"
	"fmt"
	"log"
	"net"
	"strings"

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
	cmds := []string{}

	cmds = append(cmds, "sendevent /dev/input/event0 1 "+number+" 1")
	cmds = append(cmds, "sendevent /dev/input/event0 0 0 0")
	cmds = append(cmds, "sendevent /dev/input/event0 1 "+number+" 0")
	cmds = append(cmds, "sendevent /dev/input/event0 0 0 0")
	fmt.Println("current key serial:" + number)
	fmt.Println()
	fullCmds := strings.Join(cmds, ";")
	fmt.Println(fullCmds)
	return fullCmds
}
