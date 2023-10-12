package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"UCC-as2-final/config"
	"UCC-as2-final/dto"
	"UCC-as2-final/service"
	client "UCC-as2-final/client/solr"
	con "UCC-as2-final/db"
	"strconv"
	e "UCC-as2-final/utils/errors"
)

var (
	Solr = service.NewSolrServiceImpl(
		(*client.SolrClient)(con.NewSolrClient(config.SOLRHOST, config.SOLRPORT, config.SOLRCOLLECTION)),
	)
)

func GetQuery(c *gin.Context) {
	var hotelsDto dto.HotelsDto
	query := c.Param("searchQuery")

	hotelsDto, err := Solr.GetQuery(query)
	if err != nil {
		log.Debug(hotelsDto)
		c.JSON(http.StatusBadRequest, hotelsDto)
		return
	}

	c.JSON(http.StatusOK, hotelsDto)
}

func GetQueryAllFields(c *gin.Context) {
	var hotelsDto dto.HotelsDto
	query := c.Param("searchQuery")

	hotelsDto, err := Solr.GetQueryAllFields(query)
	if err != nil {
		log.Debug(hotelsDto)
		c.JSON(http.StatusBadRequest, hotelsDto)
		return
	}

	c.JSON(http.StatusOK, hotelsDto)

}

func AddFromId(id int) error {   // agregar e.NewBadResquest para manejar el error
	err := Solr.AddFromId(strconv.Itoa(id))
	if err != nil {
		e.NewBadRequestApiError("Error adding hotel to Solr")
		return err
	}

	var w http.ResponseWriter
	
	w.WriteHeader(http.StatusOK)
	return nil
}

func Delete(id int) error {
	err := Solr.Delete(strconv.Itoa(id))
	if err != nil {
		e.NewBadRequestApiError("Error deleting hotel from Solr")
		return err
	}

	var w http.ResponseWriter
	
	w.WriteHeader(http.StatusOK)
	return nil
}