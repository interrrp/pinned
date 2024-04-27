package pin

import (
	"math/rand"
	"strconv"
	"strings"
)

func generateRandomPIN() string {
	var sb strings.Builder
	for range 4 {
		sb.WriteString(randomDigit())
	}
	return sb.String()
}

func randomDigit() string {
	return strconv.Itoa(rand.Intn(9))
}
