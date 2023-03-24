package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var conn *amqp.Connection

func Connection() *amqp.Connection {
	var err error

	if conn != nil {
		return conn
	}

	conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}

	return conn
}
