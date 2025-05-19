package infra

import (
	"fmt"
	"os"
	"sync"
	"time"

	"logger-app/internal/domain"

	"github.com/google/uuid"
)

type TextFileLogger struct {
	mu   sync.Mutex
	file string
}

func NewTextFileLogger(file string) *TextFileLogger {
	return &TextFileLogger{file: file}
}

func (t *TextFileLogger) Log(entry domain.LogEntry) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	f, err := os.OpenFile(t.file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	tx := ""
	if entry.Transaction != nil {
		tx = fmt.Sprintf("[TX:%s] ", entry.Transaction.ID)
	}
	logLine := fmt.Sprintf("%s [%s]%s%s | attrs=%v\n",
		entry.Timestamp.Format(time.RFC3339),
		entry.Level.String(),
		tx,
		entry.Message,
		entry.Attributes,
	)
	_, err = f.WriteString(logLine)
	return err
}

func (t *TextFileLogger) StartTransaction(attrs map[string]interface{}) *domain.Transaction {
	return &domain.Transaction{
		ID:         uuid.NewString(),
		Attributes: attrs,
	}
}

func (t *TextFileLogger) EndTransaction(tx *domain.Transaction) {
	// no-op
}
