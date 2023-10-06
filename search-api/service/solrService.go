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
)

type SolrService struct {
	solr *client.SolrClient
}

func NewSolrServiceImpl(
	solr *client.SolrClient,
) *SolrService {
	return &SolrService{
		solr: solr,
	}
}

func (s *SolrService) GetQuery(query string) (dto.HotelsDto, e.ApiError) {
	var hotelsDto dto.HotelsDto
	queryParams := strings.Split(query, "_")
	field, query := queryParams[0], queryParams[1]
	hotelsDto, err := s.solr.GetQuery(query, field)
	if err != nil {
		return hotelsDto, e.NewBadRequestApiError("Solr failed")
	}
	return hotelsDto, nil
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
	resp, err := http.Get(fmt.Sprintf("http://%s:%d/hotels/%s", config.HOTELSHOST, config.HOTELSPORT, id))
	if err != nil {
		log.Debugf("error getting item %s", id)
		return e.NewBadRequestApiError("error getting hotel " + id)
	}
	var body []byte
	body, _ = io.ReadAll(resp.Body)
	log.Debugf("%s", body)
	err = json.Unmarshal(body, &hotelDto)
	if err != nil {
		log.Debugf("error in unmarshal of item %s", id)
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