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

func AddFromId(c *gin.Context) {
	id := c.Param("id")
	err := Solr.AddFromId(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, err)
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	err := Solr.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, err)
}