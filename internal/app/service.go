package app

import (
	"time"

	"logger-app/internal/domain"
)

type LoggerService struct {
	logger domain.Logger
}

func NewLoggerService(logger domain.Logger) *LoggerService {
	return &LoggerService{logger: logger}
}

func (s *LoggerService) Debug(msg string, attrs map[string]interface{}, tx *domain.Transaction) {
	s.log(domain.Debug, msg, attrs, tx)
}

func (s *LoggerService) Info(msg string, attrs map[string]interface{}, tx *domain.Transaction) {
	s.log(domain.Info, msg, attrs, tx)
}

func (s *LoggerService) Warning(msg string, attrs map[string]interface{}, tx *domain.Transaction) {
	s.log(domain.Warning, msg, attrs, tx)
}

func (s *LoggerService) Error(msg string, attrs map[string]interface{}, tx *domain.Transaction) {
	s.log(domain.Error, msg, attrs, tx)
}

func (s *LoggerService) log(level domain.LogLevel, msg string, attrs map[string]interface{}, tx *domain.Transaction) {
	entry := domain.LogEntry{
		Timestamp:   time.Now().UTC(),
		Level:       level,
		Message:     msg,
		Attributes:  attrs,
		Transaction: tx,
	}
	_ = s.logger.Log(entry)
}

func (s *LoggerService) StartTransaction(attrs map[string]interface{}) *domain.Transaction {
	return s.logger.StartTransaction(attrs)
}

func (s *LoggerService) EndTransaction(tx *domain.Transaction) {
	s.logger.EndTransaction(tx)
}
