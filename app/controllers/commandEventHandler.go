package controllers

import (
	"IRCService/app/core"
	"github.com/dustin/go-coap"
	"log"
	"net"
)

//CommandHandler .
func CommandHandler(ci core.CoapInterface) core.CoapHandler {
	return func(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {

		if m.IsConfirmable() {

			cmdstr := string(m.Payload)
			log.Println(string(cmdstr))

			result, _ := ci.OnDebugCmds(cmdstr)

			res := &coap.Message{
				Type:      coap.Acknowledgement,
				Code:      coap.Content,
				MessageID: m.MessageID,
				Token:     m.Token,
				Payload:   []byte(result),
			}
			res.SetOption(coap.ContentFormat, coap.TextPlain)
			return res
		}
		return nil
	}
}
