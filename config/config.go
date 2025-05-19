package config

import (
	"encoding/json"
	"os"
	"logger-app/internal/infra"
)

type AppConfig struct {
	Logger infra.DriverConfig `json:"logger"`
}

func LoadConfig(path string) (*AppConfig, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	decoder := json.NewDecoder(f)
	var cfg AppConfig
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}