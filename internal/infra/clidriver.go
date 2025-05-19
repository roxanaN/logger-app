package infra

import (
	"fmt"
	"time"

	"logger-app/internal/domain"

	"github.com/google/uuid"
)

type CliLogger struct{}

func NewCliLogger() *CliLogger {
	return &CliLogger{}
}

func (c *CliLogger) Log(entry domain.LogEntry) error {
	tx := ""
	if entry.Transaction != nil {
		tx = fmt.Sprintf("[TX:%s] ", entry.Transaction.ID)
	}
	fmt.Printf("%s [%s]%s%s | attrs=%v\n",
		entry.Timestamp.Format(time.RFC3339),
		entry.Level.String(),
		tx,
		entry.Message,
		entry.Attributes)
	return nil
}

func (c *CliLogger) StartTransaction(attrs map[string]interface{}) *domain.Transaction {
	tx := &domain.Transaction{
		ID:         uuid.NewString(),
		Attributes: attrs,
	}
	fmt.Printf("=== Start Transaction: %s %+v ===\n", tx.ID, attrs)
	return tx
}

func (c *CliLogger) EndTransaction(tx *domain.Transaction) {
	fmt.Printf("=== End Transaction: %s ===\n", tx.ID)
}
