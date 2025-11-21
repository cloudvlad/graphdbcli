package create

import (
	"bytes"
	"encoding/json"
	"graphdbcli/internal/data_objects/repository"
	"graphdbcli/internal/tool_configurations/logging"
	"io"
	"net/http"

	"go.uber.org/zap"
)

// HandleCreate handles requests that are supposed to create repositories.
func HandleCreate(req *http.Request) {
	logging.LOGGER.Debug("create repository handler invoked")

	defer req.Body.Close()
	bodyBefore, _ := io.ReadAll(req.Body)

	var repoConfig repository.Config
	if err := json.Unmarshal(bodyBefore, &repoConfig); err != nil {
		logging.LOGGER.Error("unable to unmarshal body", zap.Error(err))
		return
	}

	logging.LOGGER.Debug("body before conversion", zap.String("body", string(bodyBefore)))

	oldFormat := ProcessConfigToExpectedPayload(repoConfig, ExampleExpectedPayload())

	bodyAfter, err := json.Marshal(oldFormat)
	println(string(bodyAfter))
	logging.LOGGER.Debug("body after conversion", zap.String("body", string(bodyAfter)))
	if err == nil {
		req.Body = io.NopCloser(bytes.NewReader(bodyAfter))
		req.ContentLength = int64(len(bodyAfter))
	}

	return
}
