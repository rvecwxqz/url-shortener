package server

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetHandler(t *testing.T) {
	type want struct {
		statusCode int
		longLink   string
	}
	tests := []struct {
		name   string
		method string
		key    string
		want   want
	}{
		{
			name:   "simple test 1",
			method: http.MethodGet,
			key:    "someVal",
			want:   want{307, "https://web.postman.co/workspace"},
		},
		{
			name:   "simple test 2",
			method: http.MethodGet,
			key:    "someVal2",
			want:   want{statusCode: 307, longLink: "http://otsphokaxju58.biz/wti20xvlw"},
		},
		{
			name:   "non exists key test",
			method: http.MethodGet,
			key:    "nonExistsKey",
			want:   want{404, ""},
		},
	}
	h := NewHandlersTest(context.Background())
	h.storage.AppendURL(context.TODO(), "someVal", "https://web.postman.co/workspace", "")
	h.storage.AppendURL(context.TODO(), "someVal2", "http://otsphokaxju58.biz/wti20xvlw", "")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			recorder := testRequest(t, nil, tt.method, "/"+tt.key, h)
			resp := recorder.Result()
			defer resp.Body.Close()
			assert.Equal(t, tt.want.statusCode, resp.StatusCode)
			assert.Equal(t, tt.want.longLink, resp.Header.Get("Location"))

		})
	}
}
