package controllers

import (
	"IRCService/app/core"
	"github.com/dustin/go-coap"
	"log"
	"net"
)

//WireCheckHandler .
func WireCheckHandler(ci core.CoapInterface) core.CoapHandler {
	return func(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {

		if m.IsConfirmable() {
			ip := "192.168.34.1"
			var deviceName string
			log.Println(string(ip))
			cName := make(chan string)
			go func() {
				core.GetDeviceName(cName)
			}()
			for deviceName == "" {
				select {
				case name := <-cName:
					deviceName = name
				}
			}
			json := "{\"Address\":\"" + ip + "\",\"BackupAddress\":\"" + ip + "\"" + ",\"Name\":\"" + deviceName + "\",\"Settings\":\"WIRE\"}"
			res := &coap.Message{
				Type:      coap.Acknowledgement,
				Code:      coap.Content,
				MessageID: m.MessageID,
				Token:     m.Token,
				Payload:   []byte(json),
			}
			res.SetOption(coap.ContentFormat, coap.TextPlain)
			return res
		}
		return nil
	}
}
