package commons

import (
	"fmt"
	"graphdbcli/internal/data_objects/instance_metadata"
	"graphdbcli/internal/tool_configurations/initialization"
	"graphdbcli/internal/tool_configurations/logging"
	"os"
	"path"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

func CollectInstancesInformation() []instance_metadata.InstanceMetadata {
	instancesPath := initialization.GetWorkbenchesDirectory()
	instances, err := os.ReadDir(instancesPath)
	if err != nil {
		fmt.Printf("Error occured while reading instances directory\n")
		logging.LOGGER.Fatal("error reading instances directory", zap.Error(err))
	}

	var instanceMetadata []instance_metadata.InstanceMetadata

	for _, instance := range instances {
		instancePath := path.Join(instancesPath, instance.Name(), "metadata.yaml")
		meta := GetInstanceInfo(instancePath)
		if meta != nil {
			instanceMetadata = append(instanceMetadata, *meta)
		}
	}

	return instanceMetadata
}

func GetInstanceInfo(metadataFilePath string) *instance_metadata.InstanceMetadata {
	dataBytes, err := os.ReadFile(metadataFilePath)
	if err != nil {
		logging.LOGGER.Error("error reading instance metadata", zap.String("metadataFilePath", metadataFilePath), zap.Error(err))
		return nil
	}

	var meta instance_metadata.InstanceMetadata
	err = yaml.Unmarshal(dataBytes, &meta)
	if err != nil {
		logging.LOGGER.Error("error unmarshalling YAML metadata file", zap.String("metadataFilePath", metadataFilePath), zap.Error(err))
		return nil
	}

	return &meta
}
