package server

import (
	"encoding/json"
	"errors"
	"github.com/rvecwxqz/url-shortener/internal/middleware"
	"github.com/rvecwxqz/url-shortener/internal/service"
	"github.com/rvecwxqz/url-shortener/internal/storage"
	"io"
	"net/http"
	"net/url"
)

type LongURL struct {
	URL string `json:"url"`
}

type ShortURL struct {
	Result string `json:"result"`
}

func (h Handlers) APIHandler(w http.ResponseWriter, r *http.Request) {
	uIDKey := middleware.ContextKey("user_id")
	ctx := r.Context()
	id := ctx.Value(uIDKey).(string)
	lURL := LongURL{}
	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(b, &lURL); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lURL.URL, _ = url.QueryUnescape(lURL.URL)
	// проверяем url на валидность
	if !service.IsURLValid(lURL.URL) {
		http.Error(w, "url not valid", http.StatusBadRequest)
		return
	}

	var e *storage.NotUniqueError
	sURL := ShortURL{}
	sURL.Result, err = h.GetShortURL(ctx, lURL.URL, id)
	w.Header().Set("Content-Type", "application/json")
	if errors.As(err, &e) {
		w.WriteHeader(http.StatusConflict)
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusCreated)
	}
	jsonVal, err := json.Marshal(sURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonVal)
}
