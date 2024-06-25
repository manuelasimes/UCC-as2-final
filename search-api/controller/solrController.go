package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"search-api/config"
	"search-api/dto"
	"search-api/service"
	client "search-api/client/solr"
	con "search-api/db"
	// "strconv"
	e "search-api/utils/errors"
	"fmt"
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
		c.JSON(http.StatusBadRequest, hotelsDto)
		return
	}

	log.Debug(hotelsDto)
	log.Debug("HOLA ACA")

	c.JSON(http.StatusOK, hotelsDto)
}

func GetQueryAllFields(c *gin.Context) {
	var hotelsDto dto.HotelsDto
	// query := c.Param("searchQuery")

	log.Debug("ACA ESTOY")

	query := "*:*"

	hotelsDto, err := Solr.GetQueryAllFields(query)
	if err != nil {
		log.Debug(err)
		c.JSON(http.StatusBadRequest, hotelsDto)
		return
	}

	c.JSON(http.StatusOK, hotelsDto)

}

func AddFromId(id string) error {   // agregar e.NewBadResquest para manejar el error
	err := Solr.AddFromId(id)
	if err != nil {
		e.NewBadRequestApiError("Error adding hotel to Solr")
		return err
	}

	fmt.Println(http.StatusOK)

	return nil
}

func Delete(id string) error {
	err := Solr.Delete(id)
	if err != nil {
		e.NewBadRequestApiError("Error deleting hotel from Solr")
		return err
	}

	fmt.Println(http.StatusOK)

	return nil
}