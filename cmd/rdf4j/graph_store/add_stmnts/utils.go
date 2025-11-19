package add_stmnts

import (
	"bytes"
	"context"
	"fmt"
	cc "graphdbcli/internal/channels/commons"
	dlo "graphdbcli/internal/flags/rdf4jcmd"
	"graphdbcli/internal/tool_configurations/logging"
	"graphdbcli/internal/tui/common_components"
	sp "graphdbcli/internal/tui/rdf4j/graph_store/add_stmnts/spinner"
	"io"
	"net/http"
	"net/url"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/enescakir/emoji"
	"go.uber.org/zap"
)

var p *tea.Program

// addStatements uploads a data file to the GraphDB instance with the given graph name.
func addStatements(ctx context.Context, ctxCancel context.CancelFunc, datafilePath, graphName string) error {
	queryParams := url.Values{}
	if graphName != "" {
		queryParams.Add("graph", graphName)
	}

	requestUrl := fmt.Sprintf("%s/repositories/%s/rdf-graphs/service?%s", dlo.GraphdbAddress, dlo.Repository, queryParams.Encode())

	p = tea.NewProgram(common_components.InitialModel(ctx, ctxCancel, sp.ReadData, &cc.Success, &cc.Failure))
	go func() {
		p.Run()
	}()

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
	req.Header.Set("Content-Type", dlo.RdfFormat)

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

	// Check the response status
	// We check against 204 because it is the status code returned when everything is alright
	if resp.StatusCode != http.StatusNoContent {
		fmt.Printf("%s HTTP request failed with status %d",
			common_components.PadStatusIndicator(emoji.CrossMark.String(), 0), resp.StatusCode)
		logging.LOGGER.Fatal("HTTP request failed", zap.Int("status_code", resp.StatusCode), zap.String("body", string(body)))
	}

	return nil
}
