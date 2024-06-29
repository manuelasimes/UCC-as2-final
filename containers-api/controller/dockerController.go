package controller

import (

	"net/http"
	// "strconv"
	"github.com/gin-gonic/gin"
	// log "github.com/sirupsen/logrus"

	service "containers-api/service"

)

func ListContainers(c *gin.Context) {

	containers, err := service.DockerService.ListContainers()

	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, containers)

}

func CreateContainer(c *gin.Context) {

	imageName := c.Param("image")
	containerName := c.Param("name")
	runningContainerID := c.Param("id")

	// imageUrl := "docker.io/mateonegri"

	// image := fmt.Sprintf("%s%s", imageUrl, imageName)

	createdContainerID, err := service.DockerService.CreateContainer(imageName, containerName, runningContainerID)
	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"container_id": createdContainerID})

}

// Deberia almacenar los ID de los contenedores, para poder tener el id para start, stop y remove. 


func StartContainer(c *gin.Context) {
 
	containerID := c.Param("id")

	err := service.DockerService.StartContainer(containerID)

	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, "Container started")

}

func StopContainer(c *gin.Context) {
 
	containerID := c.Param("id")

	err := service.DockerService.StopContainer(containerID)

	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, "Container stopped")

}

func RemoveContainer(c *gin.Context) {
 
	containerID := c.Param("id")

	err := service.DockerService.RemoveContainer(containerID)

	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, "Container removed")

}

func GetContainerStats(c *gin.Context) {

	containerID := c.Param("id")

	cpuUsage, err := service.DockerService.GetContainerStats(containerID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
	}

	c.JSON(http.StatusOK, gin.H{"CPU_Usage": cpuUsage })

}
