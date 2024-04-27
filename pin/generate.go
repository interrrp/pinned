package pin

import (
	"strconv"
	"strings"
)

func generateRandomPIN() string {
	var sb strings.Builder
	for i := range 4 {
		sb.WriteString(strconv.Itoa(i))
	}
	return sb.String()
}
