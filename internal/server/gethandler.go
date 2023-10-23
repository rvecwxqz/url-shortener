package server

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/rvecwxqz/url-shortener/internal/storage"
	"net/http"
)

func (h Handlers) GetHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is missing", http.StatusBadRequest)
		return
	}

	longURL, err := h.storage.GetLongURL(ctx, id)
	var e *storage.DeletedError
	if errors.As(err, &e) {
		w.WriteHeader(http.StatusGone)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Location", longURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
