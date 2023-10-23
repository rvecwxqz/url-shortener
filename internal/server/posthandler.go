package server

import (
	"errors"
	"github.com/rvecwxqz/url-shortener/internal/middleware"
	"github.com/rvecwxqz/url-shortener/internal/service"
	"github.com/rvecwxqz/url-shortener/internal/storage"
	"io"
	"net/http"
	"net/url"
)

func (h Handlers) PostHandler(w http.ResponseWriter, r *http.Request) {
	uIDKey := middleware.ContextKey("user_id")
	ctx := r.Context()
	id := ctx.Value(uIDKey).(string)

	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	longURL := string(b)
	longURL, _ = url.QueryUnescape(longURL)
	// проверяем url на валидность
	if !service.IsURLValid(longURL) {
		http.Error(w, "url not valid", http.StatusBadRequest)
		return
	}
	var e *storage.NotUniqueError
	shortURL, err := h.GetShortURL(ctx, longURL, id)
	if errors.As(err, &e) {
		w.WriteHeader(http.StatusConflict)
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusCreated)
	}
	w.Write([]byte(shortURL))
}
