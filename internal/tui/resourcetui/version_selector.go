package versiontui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"graphdbcli/internal/tool_configurations/statics"
)

var (
	docStyle = lipgloss.NewStyle().Margin(1, 2)
)
var SelectedVersion string

type Item struct {
	Name, Desc string
}

func (i Item) Title() string       { return i.Name }
func (i Item) Description() string { return i.Desc }
func (i Item) FilterValue() string { return i.Name }

type resourceSelectorModel struct {
	list list.Model
}

func (m resourceSelectorModel) Init() tea.Cmd {
	return nil
}

func (m resourceSelectorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			elementNumber := (m.list.Paginator.Page * m.list.Paginator.PerPage) + m.list.Cursor()
			version := statics.Versions[elementNumber].Version
			SelectedVersion = version
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

func (m resourceSelectorModel) View() string {
	return docStyle.Render(m.list.View())
}

func Initialize(versions []list.Item) *tea.Program {

	m := resourceSelectorModel{list: list.New(versions, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Available resources"
	m.list.ShowPagination()

	p := tea.NewProgram(m, tea.WithAltScreen())

	return p
}
