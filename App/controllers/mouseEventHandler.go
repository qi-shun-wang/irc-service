package controllers

import (
	core "IRCService/app/core"
	"fmt"
	"log"
	"net"
	"strings"

	coap "github.com/dustin/go-coap"
)

//MouseEventHandler .
func MouseEventHandler(ci core.CoapInterface) core.CoapHandler {
	return func(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {
		number := string(m.Payload)
		log.Println(string(number))
		cmds := parsedMouseSerial(number)
		ci.OnCmds(cmds)
		return nil
	}
}

func parsedMouseSerial(number string) string {
	cmds := []string{}
	delta := "32"
	for pos, char := range number {

		fmt.Printf("%d %c", pos, char)
		switch char {
		case '1':
			cmds = append(cmds, "sendevent /dev/input/event0 2 0 0")
			cmds = append(cmds, "sendevent /dev/input/event0 2 1 -"+delta)
			cmds = append(cmds, "sendevent /dev/input/event0 0 0 0")
		case '2':
			cmds = append(cmds, "sendevent /dev/input/event0 2 0 "+delta)
			cmds = append(cmds, "sendevent /dev/input/event0 2 1 -"+delta)
			cmds = append(cmds, "sendevent /dev/input/event0 0 0 0")
		case '3':
			cmds = append(cmds, "sendevent /dev/input/event0 2 0 "+delta)
			cmds = append(cmds, "sendevent /dev/input/event0 2 1 0")
			cmds = append(cmds, "sendevent /dev/input/event0 0 0 0")
		case '4':
			cmds = append(cmds, "sendevent /dev/input/event0 2 0 "+delta)
			cmds = append(cmds, "sendevent /dev/input/event0 2 1 "+delta)
			cmds = append(cmds, "sendevent /dev/input/event0 0 0 0")
		case '5':
			cmds = append(cmds, "sendevent /dev/input/event0 2 0 0")
			cmds = append(cmds, "sendevent /dev/input/event0 2 1 "+delta)
			cmds = append(cmds, "sendevent /dev/input/event0 0 0 0")
		case '6':
			cmds = append(cmds, "sendevent /dev/input/event0 2 0 -"+delta)
			cmds = append(cmds, "sendevent /dev/input/event0 2 1 "+delta)
			cmds = append(cmds, "sendevent /dev/input/event0 0 0 0")
		case '7':
			cmds = append(cmds, "sendevent /dev/input/event0 2 0 -"+delta)
			cmds = append(cmds, "sendevent /dev/input/event0 2 1 0")
			cmds = append(cmds, "sendevent /dev/input/event0 0 0 0")
		case '8':
			cmds = append(cmds, "sendevent /dev/input/event0 2 0 -"+delta)
			cmds = append(cmds, "sendevent /dev/input/event0 2 1 -"+delta)
			cmds = append(cmds, "sendevent /dev/input/event0 0 0 0")
		}

	}
	cmds = append(cmds, "sendevent /dev/input/event0 0 0 0")
	fmt.Println()
	fullCmds := strings.Join(cmds, ";")
	fmt.Println(fullCmds)
	return fullCmds
}
