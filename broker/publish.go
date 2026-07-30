package broker

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (r *RabbitMQ) Publish(ctx context.Context, queue string, body []byte) error {
	return r.channel.Publish(
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: body,
		},
	)
}