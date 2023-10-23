package server

import (
	"encoding/json"
	"github.com/rvecwxqz/url-shortener/internal/middleware"
	"github.com/rvecwxqz/url-shortener/internal/service"
	"io"
	"net/http"
)

type BatchRequest struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

type BatchResponse struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}

func (h Handlers) BatchHandler(w http.ResponseWriter, r *http.Request) {
	uIDKey := middleware.ContextKey("user_id")
	ctx := r.Context()
	id := ctx.Value(uIDKey).(string)

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var bRequest []BatchRequest
	if err = json.Unmarshal(b, &bRequest); err != nil {
		http.Error(w, "Can't unmarshal JSON", http.StatusBadRequest)
		return
	}
	bResp := make([]BatchResponse, len(bRequest))
	data := make(map[string]string)
	// checking if URL already in storage
	// if not - append to map for AppendBatсh method
	for i, v := range bRequest {
		if !service.IsURLValid(v.OriginalURL) {
			http.Error(w, "URL not valid", http.StatusBadRequest)
			return
		}
		var resp BatchResponse
		resp.CorrelationID = v.CorrelationID
		sURL, err := h.storage.CheckExists(ctx, v.OriginalURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if sURL == "" {
			sURL = service.GetShortURL()
			data[sURL] = v.OriginalURL
		}
		resp.ShortURL = h.Config.BaseURL + "/" + sURL
		bResp[i] = resp
	}
	// appending batch to storage
	if len(data) > 0 {
		err = h.storage.AppendBatch(ctx, data, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	// generating response
	jsonVal, err := json.Marshal(bResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonVal)
}
