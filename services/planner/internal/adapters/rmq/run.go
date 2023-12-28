package rmq

import (
	"context"
	"fmt"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/core/rmq"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/config"
)

func RunAdapter(ctx context.Context, log *logger.Logger) (*rmq.Connector, error) {
	conn, err := rmq.Connect(ctx, log, config.GetRmqConnectingString())
	if err != nil {
		return nil, fmt.Errorf("rmq.Connect: %w", err)
	}

	return conn, nil
}
