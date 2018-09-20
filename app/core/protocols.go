package core

import (
	"net"

	coap "github.com/dustin/go-coap"
)

//CoapHandler .
type CoapHandler func(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message

//CoapInterface .
type CoapInterface interface {
	OnCmds(cmds string) (string, error)
	OnDebugCmds(cmds string) (string, error)
}
