package main

import (
	"github.com/henrymoreirasilva/go-expert-eventos/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChanel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "producer message", "amq.direct")
}
