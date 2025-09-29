package show

import (
	"fmt"
	"graphdbcli/cmd/instancecmd/commons"
	"graphdbcli/internal/tui/instancetui/show"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

func showInstances() {

	instancesMetadata := commons.CollectInstancesInformation()
	rows := []table.Row{}

	for _, im := range instancesMetadata {
		rows = append(rows, im.Compact())
	}

	m := show.InitialModel(rows)

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
