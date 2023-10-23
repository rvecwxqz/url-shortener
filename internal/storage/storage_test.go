package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStorage(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"simple test 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStorage("", "")
			assert.IsType(t, &s, new(Storage), "")
		})
	}
}
