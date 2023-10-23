package storage

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

type mapStorage map[string]string

func (s mapStorage) GetLongURL(ctx context.Context, key string) (string, error) {
	val, ok := s[key]
	if !ok {
		return "", errors.New("not found")
	}
	if strings.Contains(val, "true") {
		return "", NewDeletedError()
	}
	return strings.Split(val, "|")[0], nil
}

func (s mapStorage) CheckExists(ctx context.Context, longURL string) (string, error) {
	for k, v := range s {
		if strings.Split(v, "|")[0] == longURL {
			return k, nil
		}
	}
	return "", nil
}

func (s mapStorage) PrintStorage() error {
	fmt.Println(s)
	return nil
}

func (s mapStorage) GetURLsByID(ctx context.Context, id string, baseURL string) ([]URLsData, error) {
	data := make([]URLsData, 0)
	for k, v := range s {
		if strings.Split(v, "|")[1] == id {
			data = append(data, URLsData{baseURL + "/" + k, strings.Split(v, "|")[0]})
		}
	}
	return data, nil
}
func (s mapStorage) AppendURL(ctx context.Context, shortURL string, longURL string, id string) error {
	s[shortURL] = longURL + "|" + id
	return nil
}

func (s mapStorage) AppendBatch(ctx context.Context, data map[string]string, id string) error {
	for k, v := range data {
		s.AppendURL(ctx, k, v, id)
	}
	return nil
}

func (s mapStorage) DeleteBatch(ctx context.Context, URLs []string) error {
	for _, short := range URLs {
		for k := range s {
			if k == short && !strings.Contains(k, "true") {
				s[k] += "|true"
			}
		}
	}
	return nil
}
