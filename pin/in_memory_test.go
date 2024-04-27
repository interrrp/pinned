package pin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := NewInMemoryPINService()
	assert.NotEmpty(t, s.CurrentPIN())
}

func TestIsCorrect(t *testing.T) {
	s := NewInMemoryPINService()

	assert.True(t, s.IsCorrect(s.CurrentPIN()))
	assert.False(t, s.IsCorrect("12345"))
}
