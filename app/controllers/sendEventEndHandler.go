package controllers

import (
	"IRCService/app/core"
	"github.com/dustin/go-coap"
	"log"
	"net"
	"strings"
)

//SendEventEndHandler .
func SendEventEndHandler(ci core.CoapInterface) core.CoapHandler {
	return func(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {

		number := string(m.Payload)
		log.Println(string(number))
		cmds := parsedSendEndKeySerial(number)
		ci.OnCmds(cmds)
		return nil
	}
}

func parsedSendEndKeySerial(number string) string {
	//use regular expression will slow down response speed
	// isMatch, _ := regexp.MatchString("^[0-9]+$", "0123456789")
	// if !isMatch {
	// 	log.Println("Not supported type :" + number)
	// 	return ";"
	// }

	//use input keyevent will slow down response speed
	// cmd := "input keyevent " + number
	// log.Println(cmd)
	// return cmd
	cmds := []string{}

	cmds = append(cmds, "sendevent /dev/input/event0 1 "+number+" 0")
	cmds = append(cmds, "sendevent /dev/input/event0 0 0 0")
	log.Println("current key serial:" + number)
	log.Println()
	fullCmds := strings.Join(cmds, ";")
	log.Println(fullCmds)
	return fullCmds

}
