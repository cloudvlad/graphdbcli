package commons

import (
	"context"
	"fmt"
	"graphdbcli/internal/channels/commons"
	"graphdbcli/internal/tui/common_components"
	sp "graphdbcli/internal/tui/instancetui/create"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func CheckProtocolEndpointAccessible(ctx context.Context, ctxCancel context.CancelFunc, port string, successChannel, failureChannel *chan bool, p *tea.Program) {
	p = tea.NewProgram(common_components.InitialModel(ctx, ctxCancel, sp.CheckingIsInstanceAccessible, successChannel, failureChannel))
	go func() {
		p.Run()
	}()

	url := fmt.Sprintf("http://localhost:%s", port)
	maxRetries := 340
	retryDelay := 500 * time.Millisecond

	protocolURL := url + "/protocol"
	for i := 0; i < maxRetries; i++ {
		resp, err := http.Get(protocolURL)
		if err == nil && resp.StatusCode == http.StatusOK {
			resp.Body.Close()
			commons.HandleEvent(successChannel, p)
			return
		}
		if resp != nil {
			resp.Body.Close()
		}

		time.Sleep(retryDelay)
	}

	commons.HandleEvent(failureChannel, p)
}
