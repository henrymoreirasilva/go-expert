package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func OpenChanel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return ch, nil
}

func Consume(ch *amqp.Channel, out chan<- amqp.Delivery, fila string) error {
	msgs, err := ch.Consume(
		fila,
		"go-consumer",
		false, false, false, false, nil,
	)

	if err != nil {
		return err
	}

	for msg := range msgs {
		out <- msg
	}

	return nil
}

func Publish(ch *amqp.Channel, body string, exchangeName string) error {
	err := ch.Publish(
		exchangeName,
		"", false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

	if err != nil {
		return err
	}

	return nil
}
