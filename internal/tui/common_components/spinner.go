package common_components

import (
	"context"
	"fmt"
	ss "graphdbcli/internal/data_objects/spinner_status"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/enescakir/emoji"
)

type errMsg error

type model struct {
	spinner        spinner.Model
	quitting       bool
	err            error
	ctx            context.Context
	ctxCancel      context.CancelFunc
	statuses       ss.SpinnerStatuses
	currentStatus  ss.SpinnerStatusMessage
	successChannel *chan bool
	failureChannel *chan bool
}

func InitialModel(ctx context.Context, ctxCancel context.CancelFunc, statuses ss.SpinnerStatuses, successChannel *chan bool, failureChannel *chan bool) model {
	s := spinner.New()
	s.Spinner = spinner.Line
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return model{
		spinner:        s,
		ctx:            ctx,
		ctxCancel:      ctxCancel,
		statuses:       statuses,
		successChannel: successChannel,
		failureChannel: failureChannel,
	}
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.ctxCancel()
			return m, tea.Quit
		default:
			return m, nil
		}

	case errMsg:
		m.err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		select {
		case <-*m.failureChannel:
			return m, tea.Quit
		case <-*m.successChannel:
			return m, tea.Quit
		default:
			return m, cmd
		}
	}
}

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	select {
	case <-m.ctx.Done():
		m.currentStatus = m.statuses.CancelledMessage
	case <-*m.failureChannel:
		m.currentStatus = m.statuses.FailureMessage
	case <-*m.successChannel:
		m.currentStatus = m.statuses.SuccessMessage
	default:
		m.currentStatus = m.statuses.InProgressMessage
		m.currentStatus.Status = emoji.Emoji(m.spinner.View())
	}

	return fmt.Sprintf("%s %s\n", PadStatusIndicator(m.currentStatus.Status.String(), 0), m.currentStatus.Message)
}
