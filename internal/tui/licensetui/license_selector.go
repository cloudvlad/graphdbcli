package licensetui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"graphdbcli/internal/data_objects/license"
)

var (
	docStyle = lipgloss.NewStyle().Margin(1, 2)
)
var SelectedLicense string

type Item struct {
	Name string
	Note string
}

func (i Item) Title() string       { return i.Name }
func (i Item) Description() string { return i.Note }
func (i Item) FilterValue() string { return i.Name }

type versionSelectorModel struct {
	list list.Model
}

func (m versionSelectorModel) Init() tea.Cmd {
	return nil
}

func (m versionSelectorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			elementNumber := (m.list.Paginator.Page * m.list.Paginator.PerPage) + m.list.Cursor()
			SelectedLicense = license.GetLicensesData()[elementNumber].Name
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m versionSelectorModel) View() string {
	return docStyle.Render(m.list.View())
}

func Initialize(versions []list.Item) *tea.Program {
	m := versionSelectorModel{list: list.New(versions, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Stored licenses"
	m.list.ShowPagination()

	p := tea.NewProgram(m, tea.WithAltScreen())

	return p
}
