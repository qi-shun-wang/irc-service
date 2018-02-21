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
	//use regular expression will slow down response speed
	// isMatch, _ := regexp.MatchString("^[0-9]+$", "0123456789")
	// if !isMatch {
	// 	fmt.Println("Not supported type :" + number)
	// 	return ";"
	// }

	//use input keyevent will slow down response speed
	// cmd := "input keyevent " + number
	// fmt.Println(cmd)
	// return cmd
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
