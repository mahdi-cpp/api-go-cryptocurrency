package amqp

import (
	"fmt"
	"github.com/streadway/amqp"
)

var ch *amqp.Channel

func Init() {
	fmt.Println("Go RabbitMQ Tutorial")
	conn, err := amqp.Dial("amqp://mahdi.cpp:aliali@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		//panic(1)
	}
	defer conn.Close()

	// Let's start by opening a channel to our RabbitMQ instance
	// over the connection we have already established
	ch, _ = conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()
}

func Send() {

	// with this channel open, we can then start to interact
	// with the instance and declare Queues that we can publish and
	// subscribe to
	q, err := ch.QueueDeclare(
		"Mahdi_Queue",
		false,
		false,
		false,
		false,
		nil,
	)
	// We can print out the status of our Queue here
	// this will information like the amount of messages on
	// the queue
	fmt.Println(q)
	// Handle any errors if we were unable to create the queue
	if err != nil {
		fmt.Println(err)
	}

	// attempt to publish a message to the queue!
	err = ch.Publish(
		"",
		"Mahdi_Queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(" Mahdi ??????????"),
		},
	)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Published Message to Queue")

}
