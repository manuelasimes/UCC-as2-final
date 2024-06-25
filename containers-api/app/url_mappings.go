package app

import (

	log "github.com/sirupsen/logrus"
	dockerController "containers-api/controller"
	
)

func mapUrls() {

	// Search mappings

	router.GET("/containers", dockerController.ListContainers)
	// router.GET("/containers/:id/stats", dockerController.GetContainerStats)
	router.POST("/containers/:image/:name/:id", dockerController.CreateContainer)
	router.POST("/containers/start/:id", dockerController.StartContainer)
	router.POST("/containers/stop/:id", dockerController.StopContainer)
	router.POST("/containers/remove/:id", dockerController.RemoveContainer)
	router.GET("containers/stats/:id", dockerController.GetContainerStats)

	log.Info("Finishing mappings configurations")
}
