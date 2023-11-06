package connections

import (

	"log" 
	"search-api/config"
	"os"
	"search-api/dto"
	"encoding/json"

	"github.com/streadway/amqp"
	controller "search-api/controller"
)

var QueueConn *amqp.Connection

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}

}

func QueueConnection() {

	QueueConn, err := amqp.Dial(config.AMPQConnectionURL)
	handleError(err, "Can't connect to AMQP")
	defer QueueConn.Close()

	amqpChannel, err := QueueConn.Channel()
	handleError(err, "Can't create a amqpChannel")

	defer amqpChannel.Close()

	queue, err := amqpChannel.QueueDeclare("add", true, false, false, false, nil)
	handleError(err, "Could not declare `add` queue")

	err = amqpChannel.Qos(1, 0, false)
	handleError(err, "Could not configure QoS")

	messageChannel, err := amqpChannel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Could not register consumer")

	stopChan := make(chan bool)

	go func() {
		log.Printf("Consumer ready, PID: %d", os.Getpid())
		for d := range messageChannel {
			log.Printf("Received a message: %s", d.Body)

			var queueDto  dto.QueueDto

			err := json.Unmarshal(d.Body, &queueDto)

			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
			}

			log.Printf("ID %s, Action %s", queueDto.Id, queueDto.Action)

			if err := d.Ack(false); err != nil {
				log.Printf("Error acknowledging message : %s", err)
			} else {
				log.Printf("Acknowledged message")
			}

			if ( queueDto.Action == "INSERT" || queueDto.Action == "UPDATE" ) {

				if (queueDto.Action == "UPDATE"){

					err := controller.Delete(queueDto.Id) // Me evito repetidos en Solr

					if err != nil {
						handleError(err, "Error deleting from Solr")
					}

				}

				err := controller.AddFromId(queueDto.Id)

				if err != nil {
					handleError(err, "Error inserting or deleting from Solr")
				}

			} else if queueDto.Action == "DELETE" {
				err := controller.Delete(queueDto.Id)

				if err != nil {
					handleError(err, "Error inserting or deleting from Solr")
				}
			}

		}
	}()

	// Stop for program termination
	<-stopChan
	
}