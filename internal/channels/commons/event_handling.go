package commons

import tea "github.com/charmbracelet/bubbletea"

// HandleEvent is used to manipulate tea Programs and channels.
// The channels are closed, so everything waiting for signal is notified.
// The tea programs are stopped gracefully and new channel is initiated.
func HandleEvent(eventChannel *chan bool, p *tea.Program) {
	close(*eventChannel)
	p.Quit()
	p.Wait()
	*eventChannel = make(chan bool, 1)
}
