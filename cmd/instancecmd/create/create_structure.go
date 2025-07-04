package create

import (
	"context"
	cc "graphdbcli/internal/channels/commons"
	c "graphdbcli/internal/data_objects/graphdb_cluster"
	ini "graphdbcli/internal/tool_configurations/initialization"
	"graphdbcli/internal/tui/common_components"
	sp "graphdbcli/internal/tui/instancetui/create"
	"path"

	tea "github.com/charmbracelet/bubbletea"
)

var p *tea.Program

// createInstanceStructure configures the instance setup, among which:
// - Creates the instance directory
// - Put the Platform independent distribution file content inside
// - Configures the properties
// - Sets a license
func createInstanceStructure(ctx context.Context, ctxCancel context.CancelFunc) {
	instancePath := path.Join(ini.GetClustersDirectory(), c.Instance.Name)
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

	p = tea.NewProgram(common_components.InitialModel(ctx, ctxCancel, sp.SettingUpLicense, &cc.Success, &cc.Failure))
	go func() {
		p.Run()
	}()
	configureInstanceLicense(c.Instance.StoredLicenseFilename, c.Instance.Name)
	cc.HandleEvent(&cc.Success, p)
}
