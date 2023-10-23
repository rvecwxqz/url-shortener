package config

import (
	"flag"
	"github.com/caarlos0/env/v9"
)

type Config struct {
	ServerAddress    string `env:"SERVER_ADDRESS"`
	BaseURL          string `env:"BASE_URL"`
	FileStoragePath  string `env:"FILE_STORAGE_PATH"`
	DatabaseDSN      string `env:"DATABASE_DSN"`
	DeleteBufferSize int
}

var cfg *Config

func NewConfig() *Config {
	if cfg == nil {
		cfg = &Config{}
		if err := env.Parse(cfg); err == nil {
			addressFromFlag := flag.String("a", ":8080", "address:port to listen")
			URLFromFlag := flag.String("b", "http://localhost:8080", "base url of shortener")
			fileStorageFromFlag := flag.String("f", "", "file storage path")
			DSNFromFlag := flag.String("d", "user=postgres password=123456 host=localhost"+
				" port=5432 dbname=shortener", "database DSN format: user=user password=pass host=host port=port dbname=name")
			flag.Parse()
			if cfg.ServerAddress == "" {
				cfg.ServerAddress = *addressFromFlag
			}
			if cfg.BaseURL == "" {
				cfg.BaseURL = *URLFromFlag
			}
			if cfg.FileStoragePath == "" {
				cfg.FileStoragePath = *fileStorageFromFlag
			}
			if cfg.DatabaseDSN == "" {
				cfg.DatabaseDSN = *DSNFromFlag
			}
		}
		cfg.DeleteBufferSize = 10
	}
	return cfg
}
