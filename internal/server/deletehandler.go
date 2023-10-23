package server

import (
	"encoding/json"
	"github.com/rvecwxqz/url-shortener/internal/middleware"
	"io"
	"net/http"
)

func (h Handlers) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	uIDKey := middleware.ContextKey("user_id")
	ctx := r.Context()
	id := ctx.Value(uIDKey).(string)

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var urls []string
	if err = json.Unmarshal(b, &urls); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	go h.deleteWorker.AppendURLs(h.ctx, urls, h.Config.BaseURL, id)

	w.WriteHeader(http.StatusAccepted)
}
