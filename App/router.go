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
	mux.Handle("detectEvent", coap.FuncHandler(handlers.DetectGameEventHandler(ci)))
	mux.Handle("textInput", coap.FuncHandler(handlers.TextInputHandler(ci)))
	mux.Handle("wireCheck", coap.FuncHandler(handlers.WireCheckHandler(ci)))
	mux.Handle("ping", coap.FuncHandler(handlers.PingHandler(ci)))
	mux.Handle("gameEvent", coap.FuncHandler(handlers.GameEventHandler(ci)))
	mux.Handle("gameEventBegan", coap.FuncHandler(handlers.GameBeganHandler(ci)))
	mux.Handle("gameEventEnd", coap.FuncHandler(handlers.GameEndHandler(ci)))
	mux.Handle("gameDPadEvent", coap.FuncHandler(handlers.GameDPADHandler(ci)))
	mux.Handle("gameDPadBegan", coap.FuncHandler(handlers.GameDPADBeganHandler(ci)))
	mux.Handle("gameDPadEnd", coap.FuncHandler(handlers.GameDPADEndHandler(ci)))
	mux.Handle("gameAxisEvent", coap.FuncHandler(handlers.GameAxisEventHandler(ci)))
	mux.Handle("version", coap.FuncHandler(handlers.VersionHandler(ci)))

}
