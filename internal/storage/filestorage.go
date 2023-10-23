package storage

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type fileStorage struct {
	path string
}

func newFileStorage(path string) fileStorage {
	return fileStorage{path}
}

func (s fileStorage) GetLongURL(ctx context.Context, key string) (string, error) {
	file, err := os.OpenFile(s.path, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	val := ""
	for scanner.Scan() {
		val = string(scanner.Bytes())
		splVal := strings.Split(val, "|")
		if splVal[0] == key {
			if len(splVal) == 4 {
				return "", NewDeletedError()
			}
			return splVal[1], nil
		}
	}
	return "", errors.New("not found")
}

func (s fileStorage) CheckExists(ctx context.Context, longURL string) (string, error) {
	file, err := os.OpenFile(s.path, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	val := ""

	for scanner.Scan() {
		val = string(scanner.Bytes())
		splVal := strings.Split(val, "|")
		if len(splVal) < 2 {
			continue
		}
		if splVal[1] == strings.Split(longURL, "|")[0] {
			return splVal[0], nil
		}
	}
	return "", nil
}

func (s fileStorage) PrintStorage() error {
	file, err := os.OpenFile(s.path, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(string(scanner.Bytes()))
	}
	return nil
}

func (s fileStorage) AppendURL(ctx context.Context, shortURL string, longURL string, id string) error {
	file, err := os.OpenFile(s.path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	w.WriteByte('\n')
	w.WriteString(shortURL + "|" + longURL + "|" + id)
	w.Flush()

	return nil
}

func (s fileStorage) AppendBatch(ctx context.Context, data map[string]string, id string) error {
	file, err := os.OpenFile(s.path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer file.Close()
	w := bufio.NewWriter(file)

	for k, v := range data {
		w.WriteByte('\n')
		w.WriteString(k + "|" + v + "|" + id)
	}
	w.Flush()
	return nil
}
func (s fileStorage) GetURLsByID(ctx context.Context, id string, baseURL string) ([]URLsData, error) {
	data := make([]URLsData, 0)
	file, err := os.OpenFile(s.path, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := string(scanner.Bytes())
		splVal := strings.Split(val, "|")
		if len(splVal) < 2 {
			continue
		}
		if splVal[2] == id {
			data = append(data, URLsData{baseURL + "/" + splVal[0], splVal[1]})
		}
	}
	return data, nil
}

func (s fileStorage) DeleteBatch(ctx context.Context, URLs []string) error {
	file, err := os.OpenFile(s.path, os.O_RDONLY, 0777)
	if err != nil {
		return err
	}

	fBytes, err := io.ReadAll(file)
	file.Close()
	if err != nil {
		return err
	}
	f := string(fBytes)

	lines := strings.Split(f, "\n")
	for i := 0; i < len(lines); i++ {
		for _, URL := range URLs {
			if strings.Contains(lines[i], URL) && !strings.Contains(lines[i], "true") {
				lines[i] += "|true"
			}
		}
	}
	if len(lines) > 0 {
		file, _ := os.OpenFile(s.path, os.O_WRONLY, 0777)
		defer file.Close()
		outData := strings.Join(lines, "\n")
		w := bufio.NewWriter(file)
		w.WriteString(outData)
		w.Flush()
	}
	return nil
}
