package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/interrrp/pinned/pin"
	"github.com/stretchr/testify/assert"
)

func TestValidGuess(t *testing.T) {
	pinSvc := pin.NewInMemoryPINService()
	s := NewServer(pinSvc)

	correctPIN, err := pinSvc.CurrentPIN()
	assert.NoError(t, err)

	req := newGuessRequest(correctPIN)
	res := httptest.NewRecorder()
	s.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
}

func TestInvalidGuess(t *testing.T) {
	pinSvc := pin.NewInMemoryPINService()
	s := NewServer(pinSvc)

	req := newGuessRequest("Some invalid PIN")
	res := httptest.NewRecorder()
	s.ServeHTTP(res, req)

	assert.Equal(t, http.StatusForbidden, res.Result().StatusCode)
}

func newGuessRequest(pin string) *http.Request {
	req, _ := http.NewRequest(
		http.MethodPost,
		"/",
		strings.NewReader(pin),
	)
	return req
}
