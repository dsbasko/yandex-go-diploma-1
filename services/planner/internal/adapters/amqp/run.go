package amqp

import (
	"fmt"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/core/rmq"
	"github.com/dsbasko/yandex-go-diploma-1/services/planner/internal/config"
)

type Adapter struct {
	conn *rmq.Connector
}

func RunAdapter(log *logger.Logger) (*Adapter, error) {
	conn, err := rmq.Connect(log, config.GetRmqConnectingString())
	if err != nil {
		return nil, fmt.Errorf("rmq.Connect: %w", err)
	}

	return &Adapter{conn: conn}, nil
}
