package consumers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dsbasko/yandex-go-diploma-1/core/logger"
	"github.com/dsbasko/yandex-go-diploma-1/core/rmq"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/services/jwt"
	"github.com/dsbasko/yandex-go-diploma-1/services/auth/pkg/api"
)

func Validation(ctx context.Context, log *logger.Logger, jwtService *jwt.Service, conn *rmq.Connector) (func(), error) {
	msgs, consumeClose, err := conn.SimpleConsume(ctx, &rmq.SimpleConsumeConfig{
		Exchange: api.AMQPExchange,
		Queue:    api.AMQPQueueJWTValidation,
		Key:      api.AMQPKeyJWTValidation,
		Consumer: "auth.validation.consumer",
	})
	if err != nil {
		return consumeClose, fmt.Errorf("conn.SimpleConsume: %w", err)
	}

	go func() {
		for msg := range msgs {
			if msg.ReplyTo == "" {
				continue
			}

			var dto api.JWTValidationRequestV1
			var response *api.JWTValidationResponseV1

			if errMsg := msg.Ack(true); errMsg != nil {
				log.Errorf("msg.Ack: %v", errMsg)
			}

			if errMsg := json.Unmarshal(msg.Body, &dto); errMsg != nil {
				if errMsg = reply(ctx, conn, &msg, response); errMsg != nil {
					log.Errorf("replyFn: %v", errMsg)
				}
				log.Errorf("json.Decode: %v", errMsg)
				continue
			}

			response, errMsg := jwtService.Validation(ctx, &dto)
			fmt.Println(response, errMsg)
			if errMsg != nil {
				log.Errorf("jwtService.Validation: %v", errMsg)
				continue
			}

			if errMsg = reply(ctx, conn, &msg, response); errMsg != nil {
				log.Errorf("replyFn: %v", errMsg)
				continue
			}
		}
	}()

	return consumeClose, nil
}
