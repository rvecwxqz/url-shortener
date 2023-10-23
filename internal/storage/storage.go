package storage

import (
	"context"
)

type URLsData struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type Storage interface {
	GetLongURL(ctx context.Context, key string) (string, error)
	CheckExists(ctx context.Context, longURL string) (string, error)
	AppendURL(ctx context.Context, shortURL string, longURL string, id string) error
	GetURLsByID(ctx context.Context, id string, baseURL string) ([]URLsData, error)
	AppendBatch(ctx context.Context, data map[string]string, id string) error
	DeleteBatch(ctx context.Context, URLs []string) error
	PrintStorage() error
}

func NewStorage(path string, DSN string) Storage {
	if err := DBPing(DSN); err == nil {
		return newDBStorage(DSN)
	}
	if path == "" {
		return mapStorage{}
	} else {
		return newFileStorage(path)
	}
}
