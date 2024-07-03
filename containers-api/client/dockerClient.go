package client

import (
	"context"
	"encoding/json"
	"io"
	"time"

	// log "github.com/sirupsen/logrus"

	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	// "github.com/opencontainers/image-spec/specs-go/v1"
)

func CreateDockerClient() (*client.Client, error) {

	// Set the Docker API version
	os.Setenv("DOCKER_API_VERSION", "1.41")
	// Create a new Docker client with default configuration
	var err error
	dockerClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	// Ping Docker to verify the connection
	_, err = dockerClient.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return dockerClient, nil
}

func PullImage(dockerClient *client.Client, imageName string) error {
	ctx := context.Background()
	options := types.ImagePullOptions{}

	reader, err := dockerClient.ImagePull(ctx, imageName, options)
	if err != nil {
		return err
	}
	defer reader.Close()

	// Read the pull progress and handle it as needed
	// ...

	return nil
}

func ListImages(dockerClient *client.Client) ([]types.ImageSummary, error) {
	ctx := context.Background()

	images, err := dockerClient.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		return nil, err
	}

	return images, nil
}

func CreateContainer(dockerClient *client.Client, imageName string, containerName string, runningContainerID string) (string, error) {
	ctx := context.Background()

	inspectResp, err := dockerClient.ContainerInspect(ctx, runningContainerID)
	if err != nil {
		panic(err)
	}

	newContainerConfig := inspectResp.Config
	newHostConfig := inspectResp.HostConfig
	newNetworkingConfig := &network.NetworkingConfig{
		EndpointsConfig: inspectResp.NetworkSettings.Networks,
	}

	// Remove MAC address from the network configuration (if present)
	for _, endpoint := range newNetworkingConfig.EndpointsConfig {
		endpoint.MacAddress = ""
	}

	resp, err := dockerClient.ContainerCreate(ctx, newContainerConfig, newHostConfig, newNetworkingConfig, nil, containerName)
	if err != nil {
		return "", err
	}

	containerID := resp.ID
	return containerID, nil
}

func ListContainers(dockerClient *client.Client) ([]types.Container, error) {
	ctx := context.Background()

	containers, err := dockerClient.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return nil, err
	}

	return containers, nil
}

func StartContainer(dockerClient *client.Client, containerID string) error {
	ctx := context.Background()

	err := dockerClient.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		return err
	}

	return nil
}

func StopContainer(dockerClient *client.Client, containerID string) error {
	ctx := context.Background()

	timeoutDuration := time.Second * 10
	timeout := int(timeoutDuration.Seconds()) // Convert to int
	stopOptions := container.StopOptions{
		Timeout: &timeout,
	}
	err := dockerClient.ContainerStop(ctx, containerID, stopOptions)
	if err != nil {
		return err
	}

	return nil
}

func RemoveContainer(dockerClient *client.Client, containerID string) error {
	ctx := context.Background()

	options := types.ContainerRemoveOptions{
		Force: true,
	}
	err := dockerClient.ContainerRemove(ctx, containerID, options)
	if err != nil {
		return err
	}

	return nil
}

func GetContainerStats(dockerClient *client.Client, containerID string) (float64, error) {
	ctx := context.Background()
	stats, err := dockerClient.ContainerStats(ctx, containerID, false)
	if err != nil {
		return 0.0, err
	}
	defer stats.Body.Close()

	var statsJSON types.StatsJSON
	decoder := json.NewDecoder(stats.Body)
	if err := decoder.Decode(&statsJSON); err != nil && err != io.EOF {
		return 0.0, err
	}

	cpuUsage := calculateCPUPercent(&statsJSON)

	// log.Printf("CPU Usage: %.2f%%", cpuUsage)

	return cpuUsage, nil

}

func calculateCPUPercent(stats *types.StatsJSON) float64 {

	cpuDelta := float64(stats.CPUStats.CPUUsage.TotalUsage) - float64(stats.PreCPUStats.CPUUsage.TotalUsage)
	systemDelta := float64(stats.CPUStats.SystemUsage) - float64(stats.PreCPUStats.SystemUsage)

	if systemDelta > 0.0 && cpuDelta > 0.0 {
		// log.Printf("cpuDelta: %.2f, systemDelta: %.2f", cpuDelta, systemDelta)
		// log.Printf("Stats: %.2f", float64(stats.CPUStats.OnlineCPUs))
		return (cpuDelta / systemDelta) * float64(stats.CPUStats.OnlineCPUs) * 100.0
	}
	return 0.0

}
