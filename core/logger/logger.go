package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger = zap.SugaredLogger

func NewLogger(env, serviceName string) (*Logger, error) {
	var logger *zap.SugaredLogger
	var zapConfig zap.Config

	if env == "prod" {
		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.TimeKey = "timestamp"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderCfg.EncodeLevel = zapcore.LowercaseLevelEncoder

		zapConfig = zap.Config{
			Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
			Development:      false,
			Encoding:         "json",
			EncoderConfig:    encoderCfg,
			OutputPaths:      []string{os.Stdout.Name()},
			ErrorOutputPaths: []string{os.Stderr.Name()},
			InitialFields:    map[string]any{"services": serviceName},
		}
	} else {
		encoderCfg := zap.NewDevelopmentEncoderConfig()
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

		zapConfig = zap.Config{
			Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
			Development:      true,
			Encoding:         "console",
			EncoderConfig:    encoderCfg,
			OutputPaths:      []string{os.Stdout.Name()},
			ErrorOutputPaths: []string{os.Stderr.Name()},
		}
	}

	zapLogger, err := zapConfig.Build()
	if err != nil {
		return nil, fmt.Errorf("zapConfig.Build: %w", err)
	}

	logger = zapLogger.Sugar()
	defer func() {
		err = logger.Sync()
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	return logger, nil
}
