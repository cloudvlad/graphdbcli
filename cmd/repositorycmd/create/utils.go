package create

import (
	"bytes"
	"context"
	"encoding/json"
	cc "graphdbcli/internal/channels/commons"
	"graphdbcli/internal/data_objects/repository"
	"graphdbcli/internal/tool_configurations/logging"
	"graphdbcli/internal/tui/common_components"
	ss "graphdbcli/internal/tui/repository/create"
	"io"
	"net/http"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"go.uber.org/zap"
)

var p *tea.Program

func createRepository(ctx context.Context, ctxCancel context.CancelFunc, config repository.Config, gdbLocation string) {
	p = tea.NewProgram(common_components.InitialModel(ctx, ctxCancel, ss.PrepareConfigurations, &cc.Success, &cc.Failure))
	go func() {
		p.Run()
	}()

	aaa, _ := config.ToJSON()
	println(string(aaa))

	apiPayload, err := json.MarshalIndent(aaa, "", "  ")
	if err != nil {
		cc.HandleEvent(&cc.Failure, p)
		logging.LOGGER.Fatal("error marshaling API payload:", zap.Error(err))
	}

	cc.HandleEvent(&cc.Success, p)

	p = tea.NewProgram(common_components.InitialModel(ctx, ctxCancel, ss.PrepareRequest, &cc.Success, &cc.Failure))
	go func() {
		p.Run()
	}()
	url := strings.TrimRight(gdbLocation, "/") + "/rest/repositories"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(apiPayload))
	if err != nil {
		cc.HandleEvent(&cc.Failure, p)
		logging.LOGGER.Error("failed to create request", zap.Error(err))
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json, text/plain, */*")

	cc.HandleEvent(&cc.Success, p)
	p = tea.NewProgram(common_components.InitialModel(ctx, ctxCancel, ss.CreateRepository, &cc.Success, &cc.Failure))
	go func() {
		p.Run()
	}()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		cc.HandleEvent(&cc.Failure, p)
		logging.LOGGER.Fatal("request failed", zap.Error(err))
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		cc.HandleEvent(&cc.Failure, p)
		logging.LOGGER.Fatal("request failed", zap.String("response status", resp.Status), zap.String("response body", string(body)))
	}

	cc.HandleEvent(&cc.Success, p)
	logging.LOGGER.Info("repository created successfully", zap.String("repository", config.ID))
}
