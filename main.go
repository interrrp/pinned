package main

import (
	"log"
	"net/http"

	"github.com/interrrp/pinned/pin"
)

const Addr = ":8080"

func main() {
	pinSvc := pin.NewInMemoryPINService()
	s := NewServer(pinSvc)

	log.Printf("Starting with initial PIN %s", pinSvc.CurrentPIN())
	log.Printf("Listening on %s\n", Addr)

	http.ListenAndServe(Addr, s)
}
