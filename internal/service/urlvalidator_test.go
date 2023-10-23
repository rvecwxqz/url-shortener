package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsURLValid(t *testing.T) {
	type args struct {
		URL string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"simple test 1",
			args{"https://github.com/"},
			true,
		},
		{
			"simple test 2",
			args{"ewesae"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := IsURLValid(tt.args.URL)
			assert.Equal(t, tt.want, res)
		})
	}
}
