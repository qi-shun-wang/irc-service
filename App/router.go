package app

import (
	handlers "IRCService/app/controllers"
	core "IRCService/app/core"

	coap "github.com/dustin/go-coap"
)

func setRouters(ci core.CoapInterface) {
	mux.Handle("cmd", coap.FuncHandler(handlers.CommandHandler(ci)))
	mux.Handle("mouseEvent", coap.FuncHandler(handlers.MouseEventHandler(ci)))
	mux.Handle("keyEvent", coap.FuncHandler(handlers.KeyEventHandler(ci)))
	mux.Handle("textInput", coap.FuncHandler(handlers.TextInputHandler(ci)))
}
