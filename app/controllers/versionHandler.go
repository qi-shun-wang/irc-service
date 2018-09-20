package controllers

import (
	"IRCService/app/core"
	"github.com/dustin/go-coap"
	"net"
)

//VersionHandler .
func VersionHandler(ci core.CoapInterface) core.CoapHandler {
	return func(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {

		if m.IsConfirmable() {

			res := &coap.Message{
				Type:      coap.Acknowledgement,
				Code:      coap.Content,
				MessageID: m.MessageID,
				Token:     m.Token,
				Payload:   []byte("1.0.0-RC.3"),
			}
			res.SetOption(coap.ContentFormat, coap.TextPlain)
			return res
		}
		return nil
	}
}
