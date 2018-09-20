package controllers

import (
	"IRCService/app/core"
	"github.com/dustin/go-coap"
	"log"
	"net"
	"strings"
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
	// 	log.Println("Not supported type :" + number)
	// 	return ";"
	// }
	//use input keyevent will slow down response speed
	cmds := []string{}
	cmds = append(cmds, "input keyevent "+number)
	log.Println("current keyevent serial:" + number)
	log.Println()
	fullCmds := strings.Join(cmds, ";")
	log.Println(fullCmds)
	return fullCmds

}
