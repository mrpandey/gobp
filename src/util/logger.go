package util

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type StandardLogger struct {
	// In addition to zap's structured logging,
	// sugared logger allows printf style APIs.
	*zap.SugaredLogger
}

// Creates new StandardLogger.
// dgn: the environment of deployment. Colored logs are printed for non-production.
// disableCaller: if set to true, caller's file and line are not logged.
func NewLogger(logLevel string, dgn DGNType, disableCaller bool) *StandardLogger {
	var cfg zap.Config

	outputLevel, err := zapcore.ParseLevel(logLevel)
	if err != nil {
		outputLevel = zapcore.InfoLevel
	}

	if dgn != Prod {
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		cfg = zap.NewProductionConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	}

	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stdout"}
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.TimeKey = "time"
	cfg.EncoderConfig.FunctionKey = "func"
	cfg.Level = zap.NewAtomicLevelAt(outputLevel)
	cfg.DisableCaller = disableCaller

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return &StandardLogger{SugaredLogger: logger.Sugar()}
}

func (l *StandardLogger) Printf(format string, v ...[]any) {
	if strings.Contains(format, "failed") {
		l.Errorf(format, v)
	} else {
		l.Infof(format, v)
	}
}
