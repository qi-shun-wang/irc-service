package controllers

import (
	core "IRCService/app/core"
	"log"
	"net"
	"strings"

	coap "github.com/dustin/go-coap"
)

//ScrollEventHandler .
func ScrollEventHandler(ci core.CoapInterface) core.CoapHandler {
	return func(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {

		number := strings.Split(string(m.Payload), ";")
		if len(number) > 2 {
			cmds := parsedScrollSerial(number[0], number[1])
			ci.OnCmds(cmds)
		}
		return nil
	}
}

func parsedScrollSerial(scrollX string, scrollY string) string {
	//use regular expression will slow down response speed
	// isMatch, _ := regexp.MatchString("^[0-9]+$", "0123456789")
	// if !isMatch {
	// 	log.Println("Not supported type :" + number)
	// 	return ";"
	// }
	//use input keyevent will slow down response speed

	cmds := []string{}
	s := `'{"CommandType":"WEB_OPERATION","requestId":"0","WebOperation":"PAGE_SCROLL","scrollX":` + scrollX + `,"scrollY":` + scrollY + `}'`
	cmds = append(cmds, "am broadcast -a com.ising99.action.WEB_OPERATION --es WebViewerCommand "+s)
	log.Println("current keyevent serial:" + s)
	log.Println()
	fullCmds := strings.Join(cmds, ";")
	log.Println(fullCmds)
	return fullCmds

}
