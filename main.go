package main

import (
	"log"
	"net/http"

	"github.com/interrrp/pinned/pin"
	"github.com/interrrp/pinned/server"
)

const Addr = ":8080"

func main() {
	pinSvc := pin.NewInMemoryPINService()
	s := server.NewServer(pinSvc)

	log.Printf("Starting with initial PIN %s", pinSvc.CurrentPIN())
	log.Printf("Listening on %s\n", Addr)

	http.ListenAndServe(Addr, s)
}
