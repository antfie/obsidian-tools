package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewID(t *testing.T) {
	for i := 0; i < 1000; i++ {
		result := NewID()
		assert.Len(t, result, IDLength)
	}
}
