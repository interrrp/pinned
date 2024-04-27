package pin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	pin := generateRandomPIN()

	// 4 digits
	assert.Regexp(t, `\d{4}`, pin)
}
