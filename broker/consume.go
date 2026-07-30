package broker

import "context"

func (r *RabbitMQ) Consume(ctx context.Context, queue string, handler func([]byte) error) error {
	msgs, err := r.channel.Consume(
		queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for msg := range msgs {
		if err := handler(msg.Body); err != nil {
			msg.Nack(false, true)
			continue
		}

		msg.Ack(false)
	}

	return nil
}