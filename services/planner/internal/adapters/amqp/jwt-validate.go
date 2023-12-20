package amqp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dsbasko/yandex-go-diploma-1/core/rmq"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/pkg/api"
	"github.com/rabbitmq/amqp091-go"
)

func (a *Adapter) JWTValidate(ctx context.Context, token string) (*api.JWTValidationResponseV1, error) {
	dtoBytes, err := json.Marshal(api.JWTValidationRequestV1{Token: token})
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}

	body, err := a.conn.SimplePublishAndWaitResponse(ctx, &rmq.SimplePublisherConfig{
		Exchange:  api.AMQPExchange,
		Key:       api.AMQPKeyJWTValidation,
		Mandatory: true,
		Msg: amqp091.Publishing{
			Body: dtoBytes,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("conn.SimplePublishAndWaitResponse: %w", err)
	}

	var response api.JWTValidationResponseV1
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return &response, nil
}
