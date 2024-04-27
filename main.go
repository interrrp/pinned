package main

import (
	"log"
	"net/http"

	"github.com/interrrp/pinned/pin"
)

const Addr = ":8080"

func main() {
	s := NewServer(pin.NewInMemoryPINService())
	log.Printf("Listening on %s\n", Addr)
	http.ListenAndServe(Addr, s)
}
