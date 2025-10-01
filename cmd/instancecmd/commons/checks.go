package commons

import (
	"context"
	"fmt"
	channels "graphdbcli/internal/channels/commons"
	"graphdbcli/internal/tui/common_components"
	sp "graphdbcli/internal/tui/instancetui/create"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func CheckProtocolEndpointAccessible(ctx context.Context, ctxCancel context.CancelFunc, port string) {
	p = tea.NewProgram(common_components.InitialModel(ctx, ctxCancel, sp.CheckingIsInstanceAccessible, &channels.Success, &channels.Failure))
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
			channels.HandleEvent(&channels.Success, p)
			return
		}
		if resp != nil {
			resp.Body.Close()
		}

		time.Sleep(retryDelay)
	}

	channels.HandleEvent(&channels.Success, p)
}
