package create

import (
	"fmt"
	cc "graphdbcli/internal/channels/commons"
	c "graphdbcli/internal/data_objects/graphdb_cluster"
	ini "graphdbcli/internal/tool_configurations/initialization"
	"graphdbcli/internal/tool_configurations/logging"
	tc "graphdbcli/internal/tool_configurations/statics"
	"os"
	"path"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/enescakir/emoji"
	"go.uber.org/zap"
)

func ConfigureProperties(instanceName string, p *tea.Program) {
	segments := append([]string{ini.GetWorkbenchesDirectory(), instanceName}, tc.DefaultConfPropertiesPath...)
	propertiesFullPath := path.Join(segments...)

	logging.LOGGER.Debug("Configuring properties for GraphDB instance", zap.String("propertiesFile", propertiesFullPath))
	propertiesFile, err := os.OpenFile(propertiesFullPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		cc.HandleEvent(&cc.Failure, p)
		fmt.Printf("%s Error occured whilst opening properties file", emoji.RedCircle)
		logging.LOGGER.Fatal("Error occured when opening properties file",
			zap.Error(err))
	}
	defer propertiesFile.Close()

	configureInstancePort(c.Instance.Port)

	_, err = propertiesFile.WriteString(c.Instance.PropertyOverrides)
	if err != nil {
		fmt.Println("Error occurred whilst writing properties file")
		logging.LOGGER.Fatal("Error occured when writing properties file", zap.Error(err))
	}

	cc.HandleEvent(&cc.Success, p)
}
