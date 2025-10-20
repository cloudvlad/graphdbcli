package commons

import (
	"fmt"
	wbmd "graphdbcli/internal/data_objects/workbench_metadata"
	"graphdbcli/internal/tool_configurations/initialization"
	"graphdbcli/internal/tool_configurations/logging"
	"os"
	"path"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

func CollectWorkbenchesMetadata() []wbmd.Data {
	workbenchesPath := initialization.GetWorkbenchDirectory()
	workbenches, err := os.ReadDir(workbenchesPath)
	if err != nil {
		fmt.Println("Error reading clusters directory:", err)
		os.Exit(1)
	}

	var workbencehsMetadata []wbmd.Data

	for _, workbench := range workbenches {
		workbenchPath := path.Join(workbenchesPath, workbench.Name(), "metadata.yaml")
		meta := GetWorkbenchInfo(workbenchPath)
		if meta != nil {
			workbencehsMetadata = append(workbencehsMetadata, *meta)
		}
	}

	return workbencehsMetadata
}

func GetWorkbenchInfo(metadataFilePath string) *wbmd.Data {
	dataBytes, err := os.ReadFile(metadataFilePath)
	if err != nil {
		logging.LOGGER.Error("error reading instance metadata", zap.Error(err))
		return nil
	}

	var meta wbmd.Data
	err = yaml.Unmarshal(dataBytes, &meta)
	if err != nil {
		logging.LOGGER.Error("error unmarshaling YAML", zap.Error(err))
		return nil
	}

	return &meta
}
