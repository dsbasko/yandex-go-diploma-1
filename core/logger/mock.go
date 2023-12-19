package logger

import (
	"fmt"

	"go.uber.org/zap"
)

func NewMock() (*Logger, error) {
	var logger *zap.SugaredLogger
	zapConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DPanicLevel),
		Encoding:         "console",
		OutputPaths:      []string{},
		ErrorOutputPaths: []string{},
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
