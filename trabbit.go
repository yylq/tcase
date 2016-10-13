package main

import (
	"log"

	"github.com/streadway/amqp"
)

type MRabbit struct {
	supper Mbase
}

func (m MRabbit) Supper() Mbase {
	return m.supper
}

var mrabbit = MRabbit{supper: NewSupper("rabbit")}

func init() {

	ModuleRegisterCase(mrabbit, "Send")
	ModuleRegisterCase(mrabbit, "Receive")
	RegisterModule(mrabbit)
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (m MRabbit) Send() {
	conn, err := amqp.Dial("amqp://root:root@192.168.176.7:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "hello"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}
func (m MRabbit) Receive() {
	conn, err := amqp.Dial("amqp://root:root@192.168.176.7:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
func (m MRabbit) Clear() {
	conn, err := amqp.Dial("amqp://root:root@192.168.176.7:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	n, err := ch.QueuePurge("hello", true)
	failOnError(err, "Failed to QueuePurge")
	log.Printf("n %d", n)
}
