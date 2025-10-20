package create

import (
	"context"
	"fmt"
	cc "graphdbcli/internal/channels/commons"
	c "graphdbcli/internal/data_objects/graphdb_cluster"
	ini "graphdbcli/internal/tool_configurations/initialization"
	"graphdbcli/internal/tool_configurations/logging"
	tc "graphdbcli/internal/tool_configurations/statics"
	"graphdbcli/internal/tui/common_components"
	sp "graphdbcli/internal/tui/instancetui/create"
	"path"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/enescakir/emoji"
)

var p *tea.Program

// createInstanceStructure configures the instance setup, among which:
// - Creates the instance directory
// - Put the Platform independent distribution file content inside
// - Configures the properties
// - Sets a license
func createInstanceStructure(ctx context.Context, ctxCancel context.CancelFunc) {
	instancePath := path.Join(ini.GetWorkbenchesDirectory(), c.Instance.Name)
	zipfileName := "graphdb-" + c.Instance.Version + ".zip"
	zipFilePath := path.Join(ini.GetDistDirectory(), zipfileName)

	// Creating an instance structure
	p = tea.NewProgram(common_components.InitialModel(ctx, ctxCancel, sp.CreatingInstanceStructure, &cc.Success, &cc.Failure))
	go func() {
		p.Run()
	}()
	setupBaseStructure(instancePath, zipFilePath, &cc.Failure)
	cc.HandleEvent(&cc.Success, p)

	p = tea.NewProgram(common_components.InitialModel(ctx, ctxCancel, sp.SettingUpProperties, &cc.Success, &cc.Failure))
	go func() {
		p.Run()
	}()
	ConfigureProperties(c.Instance.Name, p)

	if c.Instance.StoredLicenseFilename == "" {
		fmt.Printf("%s No license file specified. Skipping...", common_components.PadStatusIndicator(emoji.NextTrackButton.String(), tc.NotTUIStatusIndicatorAdditionalPadding))
		logging.LOGGER.Warn("no license file specified")
		return
	}

	p = tea.NewProgram(common_components.InitialModel(ctx, ctxCancel, sp.SettingUpLicense, &cc.Success, &cc.Failure))
	go func() {
		p.Run()
	}()
	configureInstanceLicense(c.Instance.StoredLicenseFilename, c.Instance.Name)
	cc.HandleEvent(&cc.Success, p)
}
