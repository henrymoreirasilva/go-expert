package main

import (
	"fmt"

	"github.com/henrymoreirasilva/go-expert-eventos/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChanel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgs, "orders")

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}

}
