package logger

import "go.uber.org/zap"

// struct of Logger to wrap zap.Logger and has its level

type Logger struct {
	Zap   *zap.Logger
	Level string
}

func (log *Logger) Message(message string) {
	logger := log.Zap
	defer logger.Sync()
	sugar := logger.Sugar()
	if log.Level == "Info" {
		sugar.Infof(message)
	} else {
		sugar.Debugf(message)
	}
}
