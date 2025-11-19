package show

import (
	"fmt"
	"graphdbcli/cmd/workbenchcmd/commons"
	"graphdbcli/internal/tui/workbenchtui/show"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

func showWorkbenches() {
	workbenchesMetadata := commons.CollectWorkbenchesMetadata()
	var rows []table.Row

	for _, wb := range workbenchesMetadata {
		rows = append(rows, wb.Compact())
	}

	m := show.InitialModel(rows)

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
