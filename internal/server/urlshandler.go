package server

import (
	"encoding/json"
	"github.com/rvecwxqz/url-shortener/internal/middleware"
	"net/http"
)

func (h Handlers) URLsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	uIDKey := middleware.ContextKey("user_id")
	id := ctx.Value(uIDKey).(string)
	data, err := h.storage.GetURLsByID(ctx, id, h.Config.BaseURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(data) == 0 {
		w.WriteHeader(http.StatusNoContent)
	}
	resp, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
