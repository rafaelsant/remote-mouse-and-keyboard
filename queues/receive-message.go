package queues

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func InitReceiverQueue(queueName string) (*amqp.Channel, amqp.Queue) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	return ch, declareQueue(ch, queueName)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func getChannel(ch *amqp.Channel, q *amqp.Queue) <-chan amqp.Delivery {
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to get channel")
	return msgs
}

func declareQueue(ch *amqp.Channel, queueName string) amqp.Queue {
	fmt.Println(queueName)
	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")
	return q
}
