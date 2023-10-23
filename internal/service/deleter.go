package service

import (
	"context"
	"fmt"
	"github.com/rvecwxqz/url-shortener/internal/config"
	"github.com/rvecwxqz/url-shortener/internal/storage"
	"log"
	"strings"
	"time"
)

type DeleteWorker struct {
	buff        []string
	maxBuffSize int
	ch          chan string
	errCh       chan error
	ticker      *time.Ticker
	storage     storage.Storage
	ctx         context.Context
}

func NewDeleteWorker(ctx context.Context, s storage.Storage, cfg *config.Config) *DeleteWorker {
	w := DeleteWorker{
		buff:        make([]string, 0, cfg.DeleteBufferSize),
		maxBuffSize: cfg.DeleteBufferSize,
		ch:          make(chan string),
		errCh:       make(chan error),
		ticker:      time.NewTicker(3 * time.Second),
		storage:     s,
		ctx:         ctx,
	}
	go func() {
		for {
			select {
			case val := <-w.ch:
				w.buff = append(w.buff, val)
				if len(w.buff) == w.maxBuffSize {
					go w.deleteURLs(ctx)
				}
			case <-w.ticker.C:
				if len(w.buff) > 0 {
					go w.deleteURLs(ctx)
				}
			}
		}
	}()
	go func() {
		for err := range w.errCh {
			fmt.Println(err.Error())
		}
	}()
	return &w
}

func (w *DeleteWorker) AppendURLs(ctx context.Context, URLs []string, baseURL string, id string) {
	// shortURLs in
	userURls, err := w.storage.GetURLsByID(ctx, id, baseURL)
	if err != nil {
		log.Println(err.Error(), "s")
	}

	urlsMap := make(map[string]bool, len(URLs))
	for _, val := range URLs {
		urlsMap[val] = false
	}
	for _, val := range userURls {
		sURL := strings.TrimPrefix(val.ShortURL, baseURL+"/")
		_, ok := urlsMap[sURL]
		if ok {
			w.ch <- sURL
		}
	}
}

func (w *DeleteWorker) deleteURLs(ctx context.Context) {
	// delete URLs from storage and clean buffer
	go func(URLs []string) {
		err := w.storage.DeleteBatch(ctx, URLs)
		if err != nil {
			w.errCh <- err
		}
	}(w.buff)
	w.buff = w.buff[:0]
}
