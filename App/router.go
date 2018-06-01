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
	mux.Handle("sendEvent", coap.FuncHandler(handlers.SendEventHandler(ci)))
	mux.Handle("sendLongPressedEvent", coap.FuncHandler(handlers.SendLongPressedEvent(ci)))
	mux.Handle("sendEventBegan", coap.FuncHandler(handlers.SendEventBeganHandler(ci)))
	mux.Handle("sendEventEnd", coap.FuncHandler(handlers.SendEventEndHandler(ci)))
	mux.Handle("textInput", coap.FuncHandler(handlers.TextInputHandler(ci)))
	mux.Handle("wireCheck", coap.FuncHandler(handlers.WireCheckHandler(ci)))
	mux.Handle("ping", coap.FuncHandler(handlers.PingHandler(ci)))
}
