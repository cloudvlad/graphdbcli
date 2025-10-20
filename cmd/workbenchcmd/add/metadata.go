package add

import (
	wbmd "graphdbcli/internal/data_objects/workbench_metadata"
	"graphdbcli/internal/tool_configurations/logging"
	"os"
	"path"
	"time"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

// createWorkbenchMetadata creates the initial metadata file and sets the default properties.
func createWorkbenchMetadata(workbenchPath string, workbenchName string) {
	var workbenchMetadata wbmd.Data

	workbenchMetadata.WorkbenchPort = "9000"
	workbenchMetadata.GraphDBPort = "7200"
	workbenchMetadata.GraphDBHost = "localhost"
	workbenchMetadata.Status = "Inactive"
	workbenchMetadata.Name = workbenchName
	workbenchMetadata.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	yamlContent, err := yaml.Marshal(workbenchMetadata)
	if err != nil {
		logging.LOGGER.Error("Failed to marshal instance metadata to YAML", zap.Error(err))
		return
	}

	metadataFile := path.Join(workbenchPath, "metadata.yaml")
	err = os.WriteFile(metadataFile, yamlContent, 0644)
	if err != nil {
		logging.LOGGER.Error("Failed to write metadata file", zap.Error(err))
		return
	}

	logging.LOGGER.Info("Workbench metadata stored", zap.String("path", metadataFile))
}
