package main

import (

	"containers-api/app"
	service "containers-api/service"
	// containerService "containers-api/service"
)

func main() {

	go service.DockerService.AutoScale()

	app.StartRoute()
}