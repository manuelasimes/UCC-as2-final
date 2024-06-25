package service

import (

	log "github.com/sirupsen/logrus"
	"fmt"
	"time"
	"strconv"
	// e "containers-api/utils/errors"

	docker_client "containers-api/client"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
)

type dockerService struct{
	cli *client.Client
}

type dockerServiceInterface interface {
	ListContainers() ([]types.Container, error)
	CreateContainer(imageName string, containerName string, runningContainerID string) (string, error)
	StartContainer(containerID string) error
	StopContainer(containerID string) error
	RemoveContainer(containerID string) error
	AutoScale() string
	GetContainerStats(containerID string) (float64, error)
}

var (
	DockerService dockerServiceInterface
)

// Initialize the DockerService variable with an instance of dockerService
func init() {
	dockerClient, err := docker_client.CreateDockerClient()  // Initialize the docker client
	if err != nil {
		log.Fatalf("Failed to create Docker client: %v", err)
	}
	DockerService = &dockerService{cli: dockerClient}
}

func (s *dockerService) ListContainers() ([]types.Container, error) {

	containers, err := docker_client.ListContainers(s.cli)

	if err != nil {
		return nil, fmt.Errorf("failed to list containers: %w", err)
	}

	return containers, nil

}

func (s *dockerService) CreateContainer(imageName string, containerName string, runningContainerID string) (string, error) {

	id, err := docker_client.CreateContainer(s.cli, imageName, containerName, runningContainerID)

	if err != nil {
		return "", fmt.Errorf("failed to create container: %w", err)
	}

	return id, nil
}

func (s *dockerService) StartContainer(containerID string) error {

	err := docker_client.StartContainer(s.cli, containerID)

	if err != nil {
		return fmt.Errorf("Failed to start container: %w", err)
	}

	return nil

}

func (s *dockerService) StopContainer(containerID string) error {

	err := docker_client.StopContainer(s.cli, containerID)

	if err != nil {
		return fmt.Errorf("Failed to stop container: %w", err)
	}

	return nil

}

func (s *dockerService) RemoveContainer(containerID string) error {

	err := docker_client.RemoveContainer(s.cli, containerID)

	if err != nil {
		return fmt.Errorf("Failed to remove container: %w", err)
	}

	return nil

}

func (s *dockerService) AutoScale() string {

	for {
        containers, err := docker_client.ListContainers(s.cli)

		if err != nil {
            log.Printf("Error listing containers: %s", err)
            continue
        }

		serviceContainerMap := make(map[string][]types.Container)

		for _, container := range containers {

			if (container.Image == "ucc-as2-final-user-res-api" || container.Image == "ucc-as2-final-frontend" || container.Image == "ucc-as2-final-search-api" || container.Image == "ucc-as2-final-hotels-api") {

				service := container.Labels["com.docker.compose.service"]
				serviceContainerMap[service] = append(serviceContainerMap[service], container)

			}

		}

        for _, serviceContainers := range serviceContainerMap {

			for _, container := range serviceContainers {

            cpuUsage, err := docker_client.GetContainerStats(s.cli, container.ID)

            if err != nil {
                log.Printf("Error getting container stats: %s", err)
                continue
            }

			containerNumber, _ := strconv.Atoi(container.Labels["com.docker.compose.container-number"])

            // Example condition: CPU usage threshold
            if cpuUsage > 20 {
                // Logic to scale up
                log.Printf("Scaling up due to high CPU usage for container %s, Id: %s", container.Image, container.ID)

				idScaledContainer, errCreate := s.CreateContainer(container.Image, fmt.Sprintf("%s-%d", container.Image , containerNumber+1), container.ID)

				if errCreate != nil {
					log.Printf("Error scaling container: %s. Error: %s", container.Image, errCreate)
				}

				errStarting := s.StartContainer(idScaledContainer)

				if errStarting != nil {
					log.Printf("Error starting container: %s. Error: %s", container.Image, errStarting)
				}

            } else if (cpuUsage < 5 && containerNumber > 1) {

				log.Printf("Scaling down due to low CPU usage for container %s", container.ID)

				errStop := s.StopContainer(container.ID)
                if errStop != nil {
                    log.Printf("Error stopping container: %s. Error: %s", container.ID, errStop)
                    continue
                }

                errRemove := s.RemoveContainer(container.ID)
                if errRemove != nil {
                    log.Printf("Error removing container: %s. Error: %s", container.ID, errRemove)
                }

			}
       	 }

		}

        time.Sleep(1 * time.Minute)
    }

}

func (s *dockerService) GetContainerStats(containerID string) (float64, error) {

	cpuUsage, err := docker_client.GetContainerStats(s.cli, containerID)

	if err != nil {
		log.Printf("Error getting container stats: %s", err)
		
		return 0.0, err
	}

	return cpuUsage, nil
	
}