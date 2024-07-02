package main

import (

	"containers-api/app"
	service "containers-api/service"
	// containerService "containers-api/service"
	// "os"
)

func main() {

	// Set the Docker API version
	// os.Setenv("DOCKER_API_VERSION", "1.43")

	go service.DockerService.AutoScale()

	app.StartRoute()
}