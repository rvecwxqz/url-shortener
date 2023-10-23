package server

import (
	"context"
	"github.com/rvecwxqz/url-shortener/internal/config"
	"github.com/rvecwxqz/url-shortener/internal/service"
	"github.com/rvecwxqz/url-shortener/internal/storage"
)

type Handlers struct {
	storage      storage.Storage
	Config       *config.Config
	ctx          context.Context
	deleteWorker *service.DeleteWorker
}

func NewHandlers(ctx context.Context) Handlers {
	cfg := config.NewConfig()
	s := storage.NewStorage(cfg.FileStoragePath, cfg.DatabaseDSN)
	return Handlers{
		storage:      s,
		Config:       cfg,
		ctx:          ctx,
		deleteWorker: service.NewDeleteWorker(ctx, s, cfg),
	}
}

func NewHandlersTest(ctx context.Context) Handlers {
	return Handlers{
		storage: storage.NewStorage("", ""),
		Config:  &config.Config{},
		ctx:     ctx,
	}
}

func (h Handlers) GetStorage() storage.Storage {
	return h.storage
}

func (h Handlers) PrintStorage() {
	h.storage.PrintStorage()
}

func (h Handlers) GetShortURL(ctx context.Context, longURL string, id string) (string, error) {
	// проверяем есть ли уже запись в storage для этого url,
	// если нет - добавляем
	shortURL, err := h.storage.CheckExists(ctx, longURL)
	var e error
	if err != nil {
		return "", nil
	}
	if shortURL == "" {
		shortURL = service.GetShortURL()
		err = h.storage.AppendURL(ctx, shortURL, longURL, id)
		if err != nil {
			return "", err
		}
	} else {
		e = storage.NewNotUniqueError()
	}
	return h.Config.BaseURL + "/" + shortURL, e
}
