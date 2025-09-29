package commons

import (
	"fmt"
	"graphdbcli/internal/data_objects/intance_metadata"
	"graphdbcli/internal/tool_configurations/initialization"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

func CollectInstancesInformation() []intance_metadata.InstanceMetadata {
	instancesPath := initialization.GetClustersDirectory()
	instances, err := os.ReadDir(instancesPath)
	if err != nil {
		fmt.Println("Error reading clusters directory:", err)
		os.Exit(1)
	}

	var instanceMetadata []intance_metadata.InstanceMetadata

	for _, instance := range instances {
		instancePath := path.Join(instancesPath, instance.Name(), "metadata.yaml")
		meta := GetInstanceInfo(instancePath)
		if meta != nil {
			instanceMetadata = append(instanceMetadata, *meta)
		}
	}

	return instanceMetadata
}

func GetInstanceInfo(metadataFilePath string) *intance_metadata.InstanceMetadata {
	dataBytes, err := os.ReadFile(metadataFilePath)
	if err != nil {
		fmt.Println("Error reading instance metadata:", err)
		return nil
	}

	var meta intance_metadata.InstanceMetadata
	err = yaml.Unmarshal(dataBytes, &meta)
	if err != nil {
		fmt.Println("Error unmarshaling YAML:", err)
		return nil
	}

	return &meta
}
