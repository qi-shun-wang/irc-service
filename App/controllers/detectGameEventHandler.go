package controllers

import (
	core "IRCService/app/core"
	"net"
	"strings"

	coap "github.com/dustin/go-coap"
)

//DetectGameEventHandler .
func DetectGameEventHandler(ci core.CoapInterface) core.CoapHandler {
	return func(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {

		if m.IsConfirmable() {
			result, _ := ci.OnDebugCmds("getevent -p")
			number := string([]rune(strings.SplitAfter(strings.SplitAfter(result, "Zinwell Gamepad F310")[0], "/dev/input/event")[1])[0])

			res := &coap.Message{
				Type:      coap.Acknowledgement,
				Code:      coap.Content,
				MessageID: m.MessageID,
				Token:     m.Token,
				Payload:   []byte(number),
			}
			res.SetOption(coap.ContentFormat, coap.TextPlain)
			return res
		}
		return nil
	}
}
