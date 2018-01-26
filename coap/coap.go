package coap

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	coap "github.com/dustin/go-coap"
)

var (
	mux      = coap.NewServeMux()
	coapPort = "5683"
)

func setHandler(uri string, ci coapInterface) {
	mux.Handle(uri, coap.FuncHandler(commandHandler(ci)))
	mux.Handle("mouseEvent", coap.FuncHandler(mouseEventHandler(ci)))
}

func runCoap() {

	log.Fatal(coap.ListenAndServe("udp", ":"+coapPort, mux))
}

type coapHandler func(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message

type coapInterface interface {
	OnCmds(cmds string) error
}

func commandHandler(ci coapInterface) coapHandler {
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

func mouseEventHandler(ci coapInterface) coapHandler {
	return func(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {

		if m.IsConfirmable() {

			mouseSerialNumber := string(m.Payload)
			log.Println(string(mouseSerialNumber))
			cmds := parsing(mouseSerialNumber)
			ci.OnCmds(cmds)

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

func parsing(mouseSerialNumber string) string {
	cmds := []string{}
	delta := "32"
	for pos, char := range mouseSerialNumber {

		fmt.Printf("%d %c", pos, char)
		switch char {
		case '1':
			cmds = append(cmds, "sendevent /dev/input/event0 2 0 0")
			cmds = append(cmds, "sendevent /dev/input/event0 2 1 -"+delta)
			cmds = append(cmds, "sendevent /dev/input/event0 0 0 0")
		case '2':
			cmds = append(cmds, "sendevent /dev/input/event0 2 0 "+delta)
			cmds = append(cmds, "sendevent /dev/input/event0 2 1 -"+delta)
			cmds = append(cmds, "sendevent /dev/input/event0 0 0 0")
		case '3':
			cmds = append(cmds, "sendevent /dev/input/event0 2 0 "+delta)
			cmds = append(cmds, "sendevent /dev/input/event0 2 1 0")
			cmds = append(cmds, "sendevent /dev/input/event0 0 0 0")
		case '4':
			cmds = append(cmds, "sendevent /dev/input/event0 2 0 "+delta)
			cmds = append(cmds, "sendevent /dev/input/event0 2 1 "+delta)
			cmds = append(cmds, "sendevent /dev/input/event0 0 0 0")
		case '5':
			cmds = append(cmds, "sendevent /dev/input/event0 2 0 0")
			cmds = append(cmds, "sendevent /dev/input/event0 2 1 "+delta)
			cmds = append(cmds, "sendevent /dev/input/event0 0 0 0")
		case '6':
			cmds = append(cmds, "sendevent /dev/input/event0 2 0 -"+delta)
			cmds = append(cmds, "sendevent /dev/input/event0 2 1 "+delta)
			cmds = append(cmds, "sendevent /dev/input/event0 0 0 0")
		case '7':
			cmds = append(cmds, "sendevent /dev/input/event0 2 0 -1")
			cmds = append(cmds, "sendevent /dev/input/event0 2 1 0")
			cmds = append(cmds, "sendevent /dev/input/event0 0 0 0")
		case '8':
			cmds = append(cmds, "sendevent /dev/input/event0 2 0 -1")
			cmds = append(cmds, "sendevent /dev/input/event0 2 1 -1")
			cmds = append(cmds, "sendevent /dev/input/event0 0 0 0")
		}

	}
	cmds = append(cmds, "sendevent /dev/input/event0 0 0 0")
	fmt.Println()
	fullCmds := strings.Join(cmds, ";")
	fmt.Println(fullCmds)
	return fullCmds
}
func periodicTransmitter(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) {
	subded := time.Now()

	for {
		msg := coap.Message{
			Type:      coap.Acknowledgement,
			Code:      coap.Content,
			MessageID: m.MessageID,
			Payload:   []byte(fmt.Sprintf("Been running for %v", time.Since(subded))),
		}

		msg.SetOption(coap.ContentFormat, coap.TextPlain)
		msg.SetOption(coap.LocationPath, m.Path())

		log.Printf("Transmitting %v", msg)
		err := coap.Transmit(l, a, msg)
		if err != nil {
			log.Printf("Error on transmitter, stopping: %v", err)
			return
		}

		time.Sleep(time.Second)
	}
}

func observeHandler(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {
	log.Printf("Got message path=%q: %#v from %v", m.Path(), m, a)
	if m.Code == coap.GET && m.Option(coap.Observe) != nil {
		if value, ok := m.Option(coap.Observe).([]uint8); ok &&
			len(value) >= 1 && value[0] == 1 {
			go periodicTransmitter(l, a, m)
		}
	}
	return nil
}
