package create

import (
	"context"
	"fmt"
	ic "graphdbcli/cmd/instancecmd/commons"
	channels "graphdbcli/internal/channels/commons"
	c "graphdbcli/internal/data_objects/graphdb_cluster"
	"graphdbcli/internal/data_objects/instance_metadata"
	"graphdbcli/internal/tool_configurations/initialization"
	"graphdbcli/internal/tool_configurations/logging"
	tc "graphdbcli/internal/tool_configurations/statics"
	"graphdbcli/internal/tui/common_components"
	"os"
	"path"
	"time"

	"github.com/enescakir/emoji"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

// createGraphDBInstance mainly do the necessary checks to ensure it is safe to continue
func createGraphDBInstance(ctx context.Context, ctxCancel context.CancelFunc) {
	logging.LOGGER.Info("Creating GraphDB Instance...")

	if c.Instance.Name == "" {
		c.Instance.Name = generateRandomName()
		fmt.Printf("%s Assigning random name - %s\n", common_components.PadStatusIndicator(emoji.Information.String(), 1),
			c.Instance.Name)
		logging.LOGGER.Info("No instance name provided, assigning random name")
	}

	if c.Instance.Version == "" {
		c.Instance.Version = tc.Versions[0].Version
		fmt.Printf("%s Utilizing version %s\n", common_components.PadStatusIndicator(emoji.Information.String(), 1),
			c.Instance.Version)
		logging.LOGGER.Info("No instance version provided, utilizing predefined one")
	}

	logging.LOGGER.Debug("Instance name: " + c.Instance.Name)
	logging.LOGGER.Debug("Instance version: " + c.Instance.Version)
	logging.LOGGER.Debug("Instance Stored license filename: " + c.Instance.StoredLicenseFilename)
	logging.LOGGER.Debug("Property overrides file: " + c.PropertyOverridesFilepath)

	c.Instance.PropertyOverrides = c.ReadPropertyOverrides()

	if !c.CheckDoesInstancesDirExists() {
		fmt.Println("The directory for storing instances data is not created")
		logging.LOGGER.Fatal("The directory for storing instances data is not created")
	}

	if IsClusterPresent() {
		fmt.Printf("%s The cluster %s already exists", common_components.PadStatusIndicator(emoji.StopSign.String(), tc.NotTUIStatusIndicatorAdditionalPadding), c.Instance.Name)
		logging.LOGGER.Fatal("A cluster with the used name already exists",
			zap.String("name", c.Instance.Name))
	}

	setupInstancePort(ctx, ctxCancel)
	createInstanceStructure(ctx, ctxCancel)
	if c.Instance.IsActive {
		ic.StartInstance(ctx, ctxCancel, &channels.Success, &channels.Failure)
		ic.CheckProtocolEndpointAccessible(ctx, ctxCancel, c.Instance.Port)

		fmt.Printf("%s Instance can be accessed at: %s\n",
			common_components.PadStatusIndicator(emoji.Star.String(), 0),
			"http://localhost:"+c.Instance.Port)
	}
	storeMetadata()
}

func storeMetadata() {
	instancePath := path.Join(initialization.GetClustersDirectory(), c.Instance.Name)
	metadataFile := path.Join(instancePath, "metadata.yaml")

	instanceMetadata := instance_metadata.InstanceMetadata{}

	instanceMetadata.Name = c.Instance.Name
	if c.Instance.IsActive {
		instanceMetadata.Status = "Status"
	} else {
		instanceMetadata.Status = "Inactive"
	}

	instanceMetadata.Version = c.Instance.Version
	instanceMetadata.Port = c.Instance.Port
	instanceMetadata.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	instanceMetadata.LicenseName = c.Instance.StoredLicenseFilename

	// Marshal to YAML
	yamlContent, err := yaml.Marshal(instanceMetadata)
	if err != nil {
		logging.LOGGER.Error("Failed to marshal instance metadata to YAML", zap.Error(err))
		return
	}

	// Write the YAML content to the file
	err = os.WriteFile(metadataFile, yamlContent, 0644)
	if err != nil {
		logging.LOGGER.Error("Failed to write metadata file", zap.Error(err))
		return
	}

	logging.LOGGER.Info("Instance metadata stored", zap.String("path", metadataFile))
}
