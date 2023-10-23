package server

import (
	"github.com/rvecwxqz/url-shortener/internal/storage"
	"net/http"
)

func (h Handlers) PingHandler(w http.ResponseWriter, r *http.Request) {
	err := storage.DBPing(h.Config.DatabaseDSN)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
