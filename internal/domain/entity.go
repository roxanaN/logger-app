package domain

import (
	"time"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warning
	Error
)

type LogEntry struct {
	Timestamp   time.Time
	Level       LogLevel
	Message     string
	Attributes  map[string]interface{}
	Transaction *Transaction
}