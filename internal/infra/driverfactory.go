package infra

import (
	"fmt"

	"logger-app/internal/domain"
)

// Config struct for driver type
type DriverConfig struct {
	Type string
	Path string // for file loggers
}

// Factory method to instantiate the right logger
func NewLoggerFromConfig(cfg DriverConfig) (domain.Logger, error) {
	switch cfg.Type {
	case "cli":
		return NewCliLogger(), nil
	case "json":
		return NewJsonFileLogger(cfg.Path), nil
	case "txt":
		return NewTextFileLogger(cfg.Path), nil
	default:
		return nil, fmt.Errorf("unknown logger type: %s", cfg.Type)
	}
}
