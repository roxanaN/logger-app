package tests

import (
	"testing"

	"logger-app/internal/app"
	"logger-app/internal/domain"
)

type mockLogger struct {
	lastLog domain.LogEntry
	tx      *domain.Transaction
}

func (m *mockLogger) Log(entry domain.LogEntry) error {
	m.lastLog = entry
	return nil
}
func (m *mockLogger) StartTransaction(attrs map[string]interface{}) *domain.Transaction {
	m.tx = &domain.Transaction{ID: "mock-tx", Attributes: attrs}
	return m.tx
}
func (m *mockLogger) EndTransaction(tx *domain.Transaction) {
	m.tx = nil
}

func TestDebugLog(t *testing.T) {
	mock := &mockLogger{}
	svc := app.NewLoggerService(mock)
	svc.Debug("test debug", map[string]interface{}{"foo": "bar"}, nil)
	if mock.lastLog.Level != domain.Debug {
		t.Error("Expected log level Debug")
	}
	if mock.lastLog.Message != "test debug" {
		t.Error("Expected message")
	}
	if mock.lastLog.Attributes["foo"] != "bar" {
		t.Error("Expected attribute foo=bar")
	}
}

func TestTransaction(t *testing.T) {
	mock := &mockLogger{}
	svc := app.NewLoggerService(mock)
	tx := svc.StartTransaction(map[string]interface{}{"customer": 1})
	if tx.ID != "mock-tx" {
		t.Error("Expected tx ID to be mock-tx")
	}
	svc.Info("msg", nil, tx)
	if mock.lastLog.Transaction == nil || mock.lastLog.Transaction.ID != "mock-tx" {
		t.Error("Transaction not attached correctly")
	}
	svc.EndTransaction(tx)
	if mock.tx != nil {
		t.Error("Transaction should be ended")
	}
}
