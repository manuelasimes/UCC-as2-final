package app

import (

	log "github.com/sirupsen/logrus"
	solrController "UCC-as2-final/controller"
	
)

func mapUrls() {

	// Search mappings

	router.GET("/search=:searchQuery", solrController.GetQuery)
	router.GET("/searchAll=:searchQuery", solrController.GetQueryAllFields)

	log.Info("Finishing mappings configurations")
}
