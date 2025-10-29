package upload

import (
	"bytes"
	"context"
	"fmt"
	cc "graphdbcli/internal/channels/commons"
	pf "graphdbcli/internal/flags/rdf4jcmd"
	"graphdbcli/internal/tool_configurations/logging"
	"graphdbcli/internal/tui/common_components"
	sp "graphdbcli/internal/tui/rdf4j/sparql/upload/spinner"
	"io"
	"net/http"
	"net/url"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/enescakir/emoji"
	"go.uber.org/zap"
)

var p *tea.Program

// uploadData uploads a data file to the GraphDB instance with the given base URI.
func uploadData(ctx context.Context, ctxCancel context.CancelFunc, datafilePath, baseUri string) {
	queryParams := url.Values{}
	if baseUri != "" {
		queryParams.Add("baseURI", baseUri)
	}

	requestUrl := fmt.Sprintf("%s/repositories/%s/statements?%s", pf.GraphdbAddress, pf.Repository, queryParams.Encode())

	logging.LOGGER.Debug("request URL", zap.String("url", requestUrl))

	p = tea.NewProgram(common_components.InitialModel(ctx, ctxCancel, sp.ReadData, &cc.Success, &cc.Failure))
	go func() {
		p.Run()
	}()

	logging.LOGGER.Debug("datafile path", zap.String("path", datafilePath))

	// Read the data file
	data, err := os.ReadFile(datafilePath)
	if err != nil {
		cc.HandleEvent(&cc.Failure, p)
		logging.LOGGER.Fatal("failed to read data file", zap.Error(err))
	}

	cc.HandleEvent(&cc.Success, p)
	p = tea.NewProgram(common_components.InitialModel(ctx, ctxCancel, sp.CreateHTTPRequest, &cc.Success, &cc.Failure))
	go func() {
		p.Run()
	}()

	// Create the HTTP request
	req, err := http.NewRequest("POST", requestUrl, bytes.NewReader(data))
	if err != nil {
		cc.HandleEvent(&cc.Failure, p)
		logging.LOGGER.Fatal("failed to create HTTP request", zap.Error(err))
	}
	req.Header.Set("Content-Type", pf.RdfFormat)

	logging.LOGGER.Debug("rdf format", zap.String("rdf_format", pf.RdfFormat))

	cc.HandleEvent(&cc.Success, p)
	p = tea.NewProgram(common_components.InitialModel(ctx, ctxCancel, sp.SendHTTPRequest, &cc.Success, &cc.Failure))
	go func() {
		p.Run()
	}()

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		cc.HandleEvent(&cc.Failure, p)
		logging.LOGGER.Fatal("failed to send HTTP request", zap.Error(err))
	}
	defer resp.Body.Close()

	cc.HandleEvent(&cc.Success, p)

	body, _ := io.ReadAll(resp.Body)
	logging.LOGGER.Debug("response status", zap.String("status", resp.Status), zap.String("response body", string(body)))

	// Check the response status
	// We check against 204 because it is the status code returned when everything is alright
	if resp.StatusCode != http.StatusNoContent {
		fmt.Printf("%s HTTP request failed with status %d",
			common_components.PadStatusIndicator(emoji.CrossMark.String(), 0), resp.StatusCode)
		logging.LOGGER.Fatal("HTTP request failed", zap.Int("status_code", resp.StatusCode), zap.String("body", string(body)))
	}

	logging.LOGGER.Info("HTTP request succeeded", zap.Int("status_code", resp.StatusCode), zap.String("body", string(body)))
}
