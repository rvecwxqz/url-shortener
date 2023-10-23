package server

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestAPIHandler(t *testing.T) {
	type want struct {
		statusCode  int
		contentType string
	}
	tests := []struct {
		name   string
		method string
		body   LongURL
		want   want
	}{
		{
			name:   "simple test 1",
			method: http.MethodPost,
			body:   LongURL{"https://cloud.google.com/storage"},
			want:   want{201, "application/json"},
		},
	}
	h := NewHandlersTest(context.Background())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBody, err := json.Marshal(tt.body)
			assert.NoError(t, err)
			recorder := testRequest(t, bytes.NewBuffer(jsonBody), tt.method, "/api/shorten", h)
			resp := recorder.Result()
			defer resp.Body.Close()
			assert.Equal(t, tt.want.statusCode, resp.StatusCode)
			assert.Equal(t, tt.want.contentType, resp.Header.Get("Content-Type"))
		})
	}
}
