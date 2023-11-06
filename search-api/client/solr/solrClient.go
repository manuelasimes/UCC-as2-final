package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"github.com/stevenferrer/solr-go"
	"io"
	"net/http"
	"search-api/config"
	"search-api/dto"
	e "search-api/utils/errors"
	"log"
)

type SolrClient struct {
	Client     *solr.JSONClient
	Collection string
}

func (sc *SolrClient) GetQuery(query string, field string) (dto.HotelsDto, e.ApiError) {
	var response dto.SolrResponseDto
	var hotelsDto dto.HotelsDto
	q, err := http.Get(fmt.Sprintf("http://%s:%d/solr/hotelSearch/select?q=%s%s%s", config.SOLRHOST, config.SOLRPORT, field, "%3A", query))

	if err != nil {
		return hotelsDto, e.NewBadRequestApiError("error getting from solr")
	}

	defer q.Body.Close()
	err = json.NewDecoder(q.Body).Decode(&response)
	if err != nil {
		log.Printf("Response Body: %s", q.Body) // Add this line
		log.Printf("Error: %s", err.Error())
		return hotelsDto, e.NewBadRequestApiError("error in unmarshal")
	}
	hotelsDto = response.Response.Docs

	log.Printf("hotels:", hotelsDto)

	return hotelsDto, nil
}

func (sc *SolrClient) GetQueryAllFields(query string) (dto.HotelsDto, e.ApiError) {
	var response dto.SolrResponseDto
	var hotelsDto dto.HotelsDto

	q, err := http.Get(
		fmt.Sprintf("http://%s:%d/solr/hotelSearch/select?q=*:*", config.SOLRHOST, config.SOLRPORT) )
	if err != nil {
		return hotelsDto, e.NewBadRequestApiError("error getting from solr")
	}

	var body []byte
	body, err = io.ReadAll(q.Body)
	if err != nil {
		return hotelsDto, e.NewBadRequestApiError("error reading body")
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return hotelsDto, e.NewBadRequestApiError("error in unmarshal")
	}

	hotelsDto = response.Response.Docs
	return hotelsDto, nil
}

func (sc *SolrClient) Add(hotelDto dto.HotelDto) e.ApiError {
	var addHotelDto dto.AddDto
	addHotelDto.Add = dto.DocDto{Doc: hotelDto}
	data, err := json.Marshal(addHotelDto)

	reader := bytes.NewReader(data)
	if err != nil {
		return e.NewBadRequestApiError("Error getting json")
	}
	resp, err := sc.Client.Update(context.TODO(), sc.Collection, solr.JSON, reader)
	logger.Debug(resp)
	if err != nil {
		return e.NewBadRequestApiError("Error in solr")
	}

	er := sc.Client.Commit(context.TODO(), sc.Collection)
	if er != nil {
		logger.Debug("Error committing load")
		return e.NewInternalServerApiError("Error committing to solr", er)
	}
	return nil
}

func (sc *SolrClient) Delete(id string) e.ApiError {
	var deleteDto dto.DeleteDto
	deleteDto.Delete = dto.DeleteDoc{Query: fmt.Sprintf("id:%s", id)}
	data, err := json.Marshal(deleteDto)
	reader := bytes.NewReader(data)
	if err != nil {
		return e.NewBadRequestApiError("Error getting json")
	}
	resp, err := sc.Client.Update(context.TODO(), sc.Collection, solr.JSON, reader)
	logger.Debug(resp)
	if err != nil {
		return e.NewBadRequestApiError("Error in solr")
	}

	er := sc.Client.Commit(context.TODO(), sc.Collection)
	if er != nil {
		logger.Debug("Error committing load")
		return e.NewInternalServerApiError("Error committing to solr", er)
	}
	return nil
}