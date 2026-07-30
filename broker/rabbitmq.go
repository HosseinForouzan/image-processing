package broker

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn *amqp.Connection
	channel *amqp.Channel
}

func New(url string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &RabbitMQ{
		conn: conn,
		channel: ch,
	}, nil
}

func (r *RabbitMQ) Close() {
	r.channel.Close()

	r.conn.Close()
}

type Consumer interface {
	Consume(ctx context.Context, queue string, handler func([]byte) error) error
}
