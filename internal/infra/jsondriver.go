package infra

import (
	"encoding/json"
	"os"
	"sync"

	"logger-app/internal/domain"

	"github.com/google/uuid"
)

type JsonFileLogger struct {
	mu   sync.Mutex
	file string
}

func NewJsonFileLogger(file string) *JsonFileLogger {
	return &JsonFileLogger{file: file}
}

func (j *JsonFileLogger) Log(entry domain.LogEntry) error {
	j.mu.Lock()
	defer j.mu.Unlock()

	var entries []domain.LogEntry

	if data, err := os.ReadFile(j.file); err == nil && len(data) > 0 {
		_ = json.Unmarshal(data, &entries)
	}

	entries = append(entries, entry)

	f, err := os.Create(j.file)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(entries)
}

func (j *JsonFileLogger) StartTransaction(attrs map[string]interface{}) *domain.Transaction {
	return &domain.Transaction{
		ID:         uuid.NewString(),
		Attributes: attrs,
	}
}

func (j *JsonFileLogger) EndTransaction(tx *domain.Transaction) {
	// no-op
}
