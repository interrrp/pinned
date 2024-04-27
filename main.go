package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/interrrp/pinned/pin"
	"github.com/interrrp/pinned/server"
	"github.com/lmittmann/tint"
)

const Addr = ":8080"

func main() {
	slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, nil)))

	pinSvc := pin.NewInMemoryPINService()
	s := server.NewServer(pinSvc)

	slog.Info("Starting", "addr", Addr, "pin", pinSvc.CurrentPIN())

	http.ListenAndServe(Addr, s)
}
