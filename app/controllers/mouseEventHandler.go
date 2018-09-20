package controllers

import (
	"IRCService/app/core"
	"github.com/dustin/go-coap"
	"net"
	"strings"
	"sync"
	"strconv"
)

var mouseLock sync.Mutex

//MouseEventHandler .
func MouseEventHandler(ci core.CoapInterface) core.CoapHandler {
	return func(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {
		number := strings.Split(string(m.Payload), ":")
		eventArray := parsedMouseMove(number[0], number[1])
		core.SendEvent(eventArray)
		return nil
	}
}

//MouseTapEventHandler .
func MouseTapEventHandler(ci core.CoapInterface) core.CoapHandler {
	return func(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {
		eventArray := parsedMouseTap()
		core.SendEvent(eventArray)
		return nil
	}
}

func MouseScrollEventHandler(ci core.CoapInterface) core.CoapHandler {
	return func(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {
		eventArray := parsedMouseScroll(string(m.Payload))
		core.SendEvent(eventArray)
		return nil
	}
}

func parsedMouseMove(dx string, dy string) []core.Event {
	mouseLock.Lock()
	defer mouseLock.Unlock()
	var eventArray []core.Event
	x, _ := strconv.Atoi(dx)
	y, _ := strconv.Atoi(dy)
	eventArray = append(eventArray, core.Event{Type:  2, Code:  0, Value: x,})
	eventArray = append(eventArray, core.Event{Type:  2, Code:  1, Value: y,})
	eventArray = append(eventArray, core.Event{Type:  0, Code:  0, Value: 0,})
	return eventArray
}

func parsedMouseTap() []core.Event {
	var eventArray []core.Event
	eventArray = append(eventArray, core.Event{Type: 1, Code: 272, Value: 1,})
	eventArray = append(eventArray, core.Event{Type: 0, Code: 0, Value: 0,})
	eventArray = append(eventArray, core.Event{Type: 1, Code: 272, Value: 0,})
	eventArray = append(eventArray, core.Event{Type: 0, Code: 0, Value: 0,})
	return eventArray
}

func parsedMouseScroll(direction string) []core.Event {
	mouseLock.Lock()
	defer mouseLock.Unlock()
	var eventArray []core.Event
	scrollDirection, _ := strconv.Atoi(direction)
	if scrollDirection > 0 {
		eventArray = append(eventArray, core.Event{Type:  1, Code:  108, Value: 1,})
		eventArray = append(eventArray, core.Event{Type:  1, Code:  108, Value: 0,})
		eventArray = append(eventArray, core.Event{Type:  0, Code:  0, Value: 0,})
	}else if scrollDirection < 0 {
		eventArray = append(eventArray, core.Event{Type:  1, Code:  103, Value: 1,})
		eventArray = append(eventArray, core.Event{Type:  1, Code:  103, Value: 0,})
		eventArray = append(eventArray, core.Event{Type:  0, Code:  0, Value: 0,})
	}
	return eventArray
}