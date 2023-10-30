package queue

import (
	"encoding/json"
	"log"
	"hotels-api/dtos"
	"hotels-api/config"


	"github.com/streadway/amqp"
)

var amqpChannel *amqp.Channel

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func QueueConnection() {
	conn, err := amqp.Dial(config.AMPQConnectionURL) // Use the same connection URL as the consumer
	handleError(err, "Can't connect to AMQP")
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create an amqpChannel")
	defer amqpChannel.Close()

}

// Function to create and send a message
func SendMessage( id int, action string) {

	addQueue, err := amqpChannel.QueueDeclare("add", true, false, false, false, nil)
	handleError(err, "Could not declare `add` queue")


	// Prepare a message in the same format as the consumer expects
	queueDto := dto.QueueDto{
		Id:     id,     // Generate a unique ID (modify this based on your use case)
		Action: action,       // Define the action as needed
		// Add any other fields as needed for your specific use case
	}

	body, err := json.Marshal(queueDto)
	if err != nil {
		handleError(err, "Error encoding JSON")
	}

	err = amqpChannel.Publish("", addQueue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json", // Change content type to JSON
		Body:         body,
	})

	if err != nil {
		log.Fatalf("Error publishing message: %s", err)
	}

	log.Printf("Message published: ID %d, Action %s", id, action)
}
