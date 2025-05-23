package config

import (
	_ "embed"
	"encoding/json"
	"fmt"

	"github.com/caarlos0/env/v11"
)

//go:embed config.json
var configJSON []byte

type Config struct {
	ClientType  string `json:"client_type"`
	RequestType string `json:"request_type"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	GRPC        GRPC
	HTTP        HTTP
	Logger      Logger
}

type GRPC struct {
	PORT int `env:"GRPC_PORT,required"`
}

type HTTP struct {
	PORT int `env:"HTTP_PORT,required"`
}

type Logger struct {
	Level string `env:"LOG_LEVEL,required"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := json.Unmarshal(configJSON, cfg)
	if err != nil {
		return nil, err
	}

	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
