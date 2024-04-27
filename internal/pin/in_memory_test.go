package pin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := NewInMemoryPINService()

	pin, err := s.CurrentPIN()
	assert.NoError(t, err)

	assert.NotEmpty(t, pin)
}

func TestIsCorrect(t *testing.T) {
	s := NewInMemoryPINService()

	pin, err := s.CurrentPIN()
	assert.NoError(t, err)

	correct, err := s.IsCorrect(pin)
	assert.NoError(t, err)
	assert.True(t, correct)

	correct, err = s.IsCorrect("12345")
	assert.NoError(t, err)
	assert.False(t, correct)
}
