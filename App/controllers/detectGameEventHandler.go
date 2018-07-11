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

//GameEventHandler .
func GameEventHandler(ci core.CoapInterface) core.CoapHandler {
	return func(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {

		number := strings.Split(string(m.Payload), ";")
		if len(number) > 2 {
			cmds := parsedGameKeySerial(number[0], number[1])
			ci.OnCmds(cmds)
		}

		return nil
	}
}

//GameBeganHandler .
func GameBeganHandler(ci core.CoapInterface) core.CoapHandler {
	return func(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {

		number := strings.Split(string(m.Payload), ";")
		if len(number) > 2 {
			cmds := parsedGameBeganKeySerial(number[0], number[1])
			ci.OnCmds(cmds)
		}

		return nil
	}
}

//GameEndHandler .
func GameEndHandler(ci core.CoapInterface) core.CoapHandler {
	return func(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {

		number := strings.Split(string(m.Payload), ";")
		if len(number) > 2 {
			cmds := parsedGameEndKeySerial(number[0], number[1])
			ci.OnCmds(cmds)
		}

		return nil
	}
}

//GameDPADHandler .
func GameDPADHandler(ci core.CoapInterface) core.CoapHandler {
	return func(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {

		number := strings.Split(string(m.Payload), ";")
		if len(number) > 3 {
			cmds := parsedGameDPadKeySerial(number[0], number[1], number[2])
			ci.OnCmds(cmds)
		}

		return nil
	}
}

//GameDPADBeganHandler .
func GameDPADBeganHandler(ci core.CoapInterface) core.CoapHandler {
	return func(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {

		number := strings.Split(string(m.Payload), ";")
		if len(number) > 3 {
			cmds := parsedGameDPadBeganKeySerial(number[0], number[1], number[2])
			ci.OnCmds(cmds)
		}

		return nil
	}
}

//GameDPADEndHandler .
func GameDPADEndHandler(ci core.CoapInterface) core.CoapHandler {
	return func(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {

		number := strings.Split(string(m.Payload), ";")
		if len(number) > 2 {
			cmds := parsedGameDPadEndKeySerial(number[0], number[1])
			ci.OnCmds(cmds)
		}

		return nil
	}
}

//GameAxisEventHandler .
func GameAxisEventHandler(ci core.CoapInterface) core.CoapHandler {
	return func(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {

		number := strings.Split(string(m.Payload), ";")
		if len(number) > 3 {
			cmds := parsedGameAxisEventKeySerial(number[0], number[1], number[2])
			ci.OnCmds(cmds)
		}

		return nil
	}
}

func parsedGameKeySerial(eventNumber string, number string) string {
	cmds := []string{}
	cmds = append(cmds, "sendevent /dev/input/event"+eventNumber+" 1 "+number+" 1")
	cmds = append(cmds, "sendevent /dev/input/event"+eventNumber+" 0 0 0")
	cmds = append(cmds, "sendevent /dev/input/event"+eventNumber+" 1 "+number+" 0")
	cmds = append(cmds, "sendevent /dev/input/event"+eventNumber+" 0 0 0")
	fullCmds := strings.Join(cmds, ";")
	return fullCmds
}

func parsedGameAxisEventKeySerial(eventNumber string, number string, value string) string {
	cmds := []string{}
	cmds = append(cmds, "sendevent /dev/input/event"+eventNumber+" 3 "+number+" "+value)
	cmds = append(cmds, "sendevent /dev/input/event"+eventNumber+" 0 0 0")
	fullCmds := strings.Join(cmds, ";")
	return fullCmds
}

func parsedGameBeganKeySerial(eventNumber string, number string) string {
	cmds := []string{}
	cmds = append(cmds, "sendevent /dev/input/event"+eventNumber+" 1 "+number+" 1")
	cmds = append(cmds, "sendevent /dev/input/event"+eventNumber+" 0 0 0")
	fullCmds := strings.Join(cmds, ";")
	return fullCmds
}

func parsedGameEndKeySerial(eventNumber string, number string) string {
	cmds := []string{}
	cmds = append(cmds, "sendevent /dev/input/event"+eventNumber+" 1 "+number+" 0")
	cmds = append(cmds, "sendevent /dev/input/event"+eventNumber+" 0 0 0")
	fullCmds := strings.Join(cmds, ";")
	return fullCmds
}

func parsedGameDPadKeySerial(eventNumber string, number string, direction string) string {
	cmds := []string{}
	cmds = append(cmds, "sendevent /dev/input/event"+eventNumber+" 3 "+number+" "+direction)
	cmds = append(cmds, "sendevent /dev/input/event"+eventNumber+" 0 0 0")
	cmds = append(cmds, "sendevent /dev/input/event"+eventNumber+" 3 "+number+" 0")
	cmds = append(cmds, "sendevent /dev/input/event"+eventNumber+" 0 0 0")
	fullCmds := strings.Join(cmds, ";")
	return fullCmds
}

func parsedGameDPadBeganKeySerial(eventNumber string, number string, direction string) string {
	cmds := []string{}
	cmds = append(cmds, "sendevent /dev/input/event"+eventNumber+" 3 "+number+" "+direction)
	cmds = append(cmds, "sendevent /dev/input/event"+eventNumber+" 0 0 0")
	fullCmds := strings.Join(cmds, ";")
	return fullCmds
}

func parsedGameDPadEndKeySerial(eventNumber string, number string) string {
	cmds := []string{}
	cmds = append(cmds, "sendevent /dev/input/event"+eventNumber+" 3 "+number+" 0")
	cmds = append(cmds, "sendevent /dev/input/event"+eventNumber+" 0 0 0")
	fullCmds := strings.Join(cmds, ";")
	return fullCmds
}
