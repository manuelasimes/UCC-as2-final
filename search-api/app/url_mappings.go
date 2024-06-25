package app

import (

	log "github.com/sirupsen/logrus"
	solrController "search-api/controller"
	
)

func mapUrls() {

	// Search mappings

	router.GET("/search-api/search=:searchQuery", solrController.GetQuery)
	router.GET("/search-api/searchAll=:searchQuery", solrController.GetQueryAllFields)

	log.Info("Finishing mappings configurations")
}
