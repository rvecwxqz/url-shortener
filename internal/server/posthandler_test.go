package server

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

func TestPostHandler(t *testing.T) {
	type want struct {
		statusCode int
	}
	tests := []struct {
		name    string
		request string
		method  string
		link    string
		want    want
	}{
		{
			name:    "simple test #1",
			request: "/",
			method:  http.MethodPost,
			link:    "https://mail.google.com/mail/u/0/",
			want:    want{statusCode: 201},
		},
		{
			name:    "incorrect link test",
			request: "/",
			method:  http.MethodPost,
			link:    "heqwea",
			want:    want{statusCode: 400},
		},
	}
	h := NewHandlersTest(context.Background())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := tt.link

			recorder := testRequest(t, strings.NewReader(body), tt.method, tt.request, h)
			res := recorder.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.want.statusCode, res.StatusCode)
			assert.NotNil(t, res.Body)
		})
	}
}
