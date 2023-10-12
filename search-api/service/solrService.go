package service

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
	"UCC-as2-final/config"
	"UCC-as2-final/dto"
	client "UCC-as2-final/client/solr"
	e "UCC-as2-final/utils/errors"
	"strconv"
	"os"
	"sync"

	"github.com/streadway/amqp"
)

type SolrService struct {
	solr *client.SolrClient
}

func NewSolrServiceImpl(solr *client.SolrClient) *SolrService {
	return &SolrService{
		solr: solr,
	}
}

func (s *SolrService) GetQuery(query string) (dto.HotelsDto, e.ApiError) {
	var hotelsDto dto.HotelsDto
	queryParams := strings.Split(query, "_")

	numParams := len(queryParams)

	field, query := queryParams[0], queryParams[1] 

	log.Printf("%s and %s", field, query)

	hotelsDto, err := s.solr.GetQuery(query, field)
	if err != nil {
		return hotelsDto, e.NewBadRequestApiError("Solr failed")
	}

	if numParams == 4 {

	startdate, enddate := queryParams[2], queryParams[3]

	sDate, _ := strconv.Atoi(startdate)
	eDate, _ := strconv.Atoi(enddate)

	// Create a channel to collect results
	resultsChan := make(chan dto.HotelDto, len(hotelsDto))

	// Create a WaitGroup
	var wg sync.WaitGroup
	var hotel dto.HotelDto

	// Iterate through each hotel and make concurrent API calls
	for _, hotel = range hotelsDto {
		wg.Add(1) // Increment the WaitGroup counter for each Goroutine
		go func(hotel dto.HotelDto) {
			defer wg.Done() // Decrement the WaitGroup counter when Goroutine is done

			log.Debug("hola")

			id, _ := strconv.Atoi(hotel.Id)

			// Make API call for each hotel and send the hotel ID
			result, err := s.GetHotelInfo(id, sDate, eDate) // Assuming you have a function to get hotel info
			if err != nil {
				result = false
			}

			var response dto.HotelDto

			if result == true {
				response = hotel
			}

			resultsChan <- response
		}(hotel)
	}


	log.Debug("hola3")
	// Create a slice to store the results
	var hotelResults dto.HotelsDto

	// Start a Goroutine to close the channel when all Goroutines are done
	go func() {
		wg.Wait()     // Wait for all Goroutines to finish
		close(resultsChan) // Close the channel when all Goroutines are done
	}()

	// Collect results from the channel
	for response := range resultsChan {
			hotelResults = append(hotelResults, response)
	}

	log.Debug("hola2") 

	return hotelResults, nil

	}

	return hotelsDto, nil
}

func (s *SolrService) GetHotelInfo(id int, startdate int, enddate int) (bool, error) {

		resp, err := http.Get(fmt.Sprintf("http:://%s:%d/hotel/availability/%d/%d/%d", config.USERAPIHOST, config.USERAPIPORT, id, startdate, enddate))

		if err != nil {
			return false, e.NewBadRequestApiError("user-res-api failed")
		}

		var body []byte
		body, _ = io.ReadAll(resp.Body)

		var responseDto dto.AvailabilityResponse
		err = json.Unmarshal(body,&responseDto)

		if err != nil {
			log.Debugf("error in unmarshal")
			return false, e.NewBadRequestApiError("getHotelInfo failed")
		}

		status := responseDto.Status

		log.Debug("llego aca", status)

		return status, nil
}
	
func (s *SolrService) GetQueryAllFields(query string) (dto.HotelsDto, e.ApiError) {
	var hotelsDto dto.HotelsDto
	hotelsDto, err := s.solr.GetQueryAllFields(query)
	if err != nil {
		return hotelsDto, e.NewBadRequestApiError("Solr failed")
	}
	return hotelsDto, nil
}

func (s *SolrService) AddFromId(id string) e.ApiError {
	var hotelDto dto.HotelDto
	// resp, err := http.Get(fmt.Sprintf("http://%s:%d/hotels/%s", config.HOTELSHOST, config.HOTELSPORT, id))
	resp, err := http.Get(fmt.Sprintf("http://localhost:8070/hotel/%s", id))
	if err != nil {
		log.Debugf("error getting item %s", id)
		return e.NewBadRequestApiError("error getting hotel " + id)
	}
	var body []byte
	body, _ = io.ReadAll(resp.Body)
	log.Debugf("%s", body)
	err = json.Unmarshal(body, &hotelDto)
	if err != nil {
		log.Debugf("error in unmarshal of hotel %s", id)
		return e.NewBadRequestApiError("error in unmarshal of hotel")
	}
	er := s.solr.Add(hotelDto)
	if er != nil {
		log.Debugf("error adding to solr")
		return e.NewInternalServerApiError("Adding to Solr error", err)
	}
	return nil
}

func (s *SolrService) Delete(id string) e.ApiError {
	err := s.solr.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

var QueueConn *amqp.Connection

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}

}

func (s *SolrService) QueueStart(){

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

			log.Printf("ID %d, Action %s", queueDto.Id, queueDto.Action)

			if err := d.Ack(false); err != nil {
				log.Printf("Error acknowledging message : %s", err)
			} else {
				log.Printf("Acknowledged message")
			}

			if ( queueDto.Action == "INSERT" || queueDto.Action == "UPDATE" ) {
			s.AddFromId(strconv.Itoa(queueDto.Id))
			} else if queueDto.Action == "DELETE" {
				s.Delete(strconv.Itoa(queueDto.Id))
			}

		}
	}()

	// Stop for program termination
	<-stopChan

}
