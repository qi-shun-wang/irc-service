package controllers

import (
	"IRCService/app/core"
	"github.com/dustin/go-coap"
	"log"
	"net"
	"strings"
)

//SlideEventHandler .
func SlideEventHandler(ci core.CoapInterface) core.CoapHandler {
	return func(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {

		number := strings.Split(string(m.Payload), ";")
		if len(number) > 2 {
			cmds := parsedSlideSerial(number[0], number[1])
			ci.OnCmds(cmds)
		}
		return nil
	}
}

func parsedSlideSerial(shiftX string, shiftY string) string {
	//use regular expression will slow down response speed
	// isMatch, _ := regexp.MatchString("^[0-9]+$", "0123456789")
	// if !isMatch {
	// 	log.Println("Not supported type :" + number)
	// 	return ";"
	// }
	//use input keyevent will slow down response speed

	cmds := []string{}
	s := `'{"CommandType":"WEB_OPERATION","requestId":"0","WebOperation":"CURSOR_MOVE","shiftX":` + shiftX + `,"shiftY":` + shiftY + `}'`
	cmds = append(cmds, "am broadcast -a com.ising99.action.WEB_OPERATION --es WebViewerCommand "+s)
	log.Println("current keyevent serial:" + s)
	log.Println()
	fullCmds := strings.Join(cmds, ";")
	log.Println(fullCmds)
	return fullCmds

}
