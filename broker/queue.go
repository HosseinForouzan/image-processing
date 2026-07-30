package broker

import amqp "github.com/rabbitmq/amqp091-go"

func (r *RabbitMQ) DeclareQueue(name string) (amqp.Queue, error) {
	return r.channel.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		nil,
	)
}