package main

import (
	"log/slog"
	"net/http"

	"github.com/interrrp/pinned/pin"
	"github.com/interrrp/pinned/server"
)

const Addr = ":8080"

func main() {
	pinSvc := pin.NewInMemoryPINService()
	s := server.NewServer(pinSvc)

	slog.Info("Starting", "addr", Addr, "pin", pinSvc.CurrentPIN())

	http.ListenAndServe(Addr, s)
}
