package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/rvecwxqz/url-shortener/internal/middleware"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testRequest(t *testing.T, bodyBytes io.Reader, method, path string, h Handlers) *httptest.ResponseRecorder {
	r := chi.NewRouter()
	r.Use(middleware.GzipHandle)
	r.Use(middleware.CookieHandler)

	r.Route("/", func(r chi.Router) {
		r.Get("/{id}", h.GetHandler)
		r.Post("/", h.PostHandler)
		r.Post("/api/shorten", h.APIHandler)
		r.Get("/api/user/urls", h.URLsHandler)
	})

	req, err := http.NewRequest(method, path, bodyBytes)
	require.NoError(t, err)

	writer := httptest.NewRecorder()
	r.ServeHTTP(writer, req)

	return writer
}
