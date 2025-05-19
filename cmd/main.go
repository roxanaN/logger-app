package main

import (
	"fmt"
	"os"

	"logger-app/config"
	"logger-app/internal/app"
	"logger-app/internal/infra"
)

func main() {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		os.Exit(1)
	}

	logger, err := infra.NewLoggerFromConfig(cfg.Logger)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize logger: %v\n", err)
		os.Exit(2)
	}

	telemetry := app.NewLoggerService(logger)

	tx := telemetry.StartTransaction(map[string]interface{}{"customerId": 123, "origin": "http"})
	telemetry.Info("Processing started", map[string]interface{}{"step": 1}, tx)
	telemetry.Warning("Slow DB response", map[string]interface{}{"step": 2, "ms": 234}, tx)
	telemetry.Error("Failed to process payment", map[string]interface{}{"error": "insufficient funds"}, tx)
	telemetry.EndTransaction(tx)
}
