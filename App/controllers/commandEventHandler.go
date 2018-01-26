package controllers

import (
	core "IRCService/app/core"
	"log"
	"net"

	coap "github.com/dustin/go-coap"
)

//CommandHandler .
func CommandHandler(ci core.CoapInterface) core.CoapHandler {
	return func(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {

		if m.IsConfirmable() {

			cmdstr := string(m.Payload)
			log.Println(string(cmdstr))
			ci.OnCmds(cmdstr)

			res := &coap.Message{
				Type:      coap.Acknowledgement,
				Code:      coap.Content,
				MessageID: m.MessageID,
				Token:     m.Token,
				Payload:   []byte("success"),
			}
			res.SetOption(coap.ContentFormat, coap.TextPlain)
			return res
		}
		return nil
	}
}
