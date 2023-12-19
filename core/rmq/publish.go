package rmq

import (
	"context"
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

/*
Отправка сообщения
*/

type PublisherConfig struct {
	Exchange  string
	Key       string
	Mandatory bool
	Msg       amqp091.Publishing
}

func (c *Connector) Publish(
	ctx context.Context,
	cfgPublisher *PublisherConfig,
) error {
	ch, err := c.conn.Channel()
	if err != nil {
		return fmt.Errorf("conn.Channel: %w", err)
	}
	defer ch.Close()

	return ch.PublishWithContext(
		ctx,
		cfgPublisher.Exchange,
		cfgPublisher.Key,
		cfgPublisher.Mandatory,
		false,
		cfgPublisher.Msg,
	)
}

/*
Отправка сообщения и ожидание ответа
*/

type SimplePublisherAndWaitResponseConfig struct {
	Exchange  string
	Key       string
	Mandatory bool
	Msg       amqp091.Publishing
}

func (c *Connector) SimplePublishAndWaitResponse(
	ctx context.Context,
	cfgPublisher *SimplePublisherAndWaitResponseConfig,
) (<-chan amqp091.Delivery, func(), error) {
	replyQueue, err := c.QueueDeclare(&QueueConfig{
		Name:       "",
		Durable:    false,
		AutoDelete: true,
		Exclusive:  true,
		NoWait:     false,
		Args:       nil,
	})
	if err != nil {
		return nil, func() {}, err
	}

	msgs, closeFn, err := c.Consume(ctx, &ConsumeConfig{
		Queue:     replyQueue.Name,
		Consumer:  "",
		AutoAck:   true,
		Exclusive: false,
		NoLocal:   false,
		NoWait:    false,
		Args:      nil,
	})
	if err != nil {
		return nil, func() {}, err
	}

	cfgPublisher.Msg.ReplyTo = replyQueue.Name
	if errPub := c.Publish(ctx, &PublisherConfig{
		cfgPublisher.Exchange,
		cfgPublisher.Key,
		cfgPublisher.Mandatory,
		cfgPublisher.Msg,
	}); errPub != nil {
		return nil, func() {}, errPub
	}

	return msgs, closeFn, nil
}
