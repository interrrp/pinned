package server

import (
	"io"
	"net/http"
)

func readBodyToString(r *http.Request) (string, error) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
