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

	pinSvc, err := pin.NewRedisPINService(":6379", "pin")
	if err != nil {
		slog.Error("Failed to connect to Redis", err)
		os.Exit(1)
	}

	s := server.NewServer(pinSvc)

	pin, err := pinSvc.CurrentPIN()
	if err != nil {
		slog.Error("Failed to get initial PIN", err)
		os.Exit(1)
	}
	slog.Info("Starting", "addr", Addr, "pin", pin)

	http.ListenAndServe(Addr, s)
}
