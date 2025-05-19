package domain

type Logger interface {
	Log(entry LogEntry) error
	StartTransaction(attrs map[string]interface{}) *Transaction
	EndTransaction(tx *Transaction)
}