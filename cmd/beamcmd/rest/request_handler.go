package rest

import (
	"graphdbcli/cmd/beamcmd/rest/repositories"
	"graphdbcli/internal/tool_configurations/logging"
	"net/http"
	"regexp"
)

func ChangeRequest(req *http.Request) {
	re := regexp.MustCompile(`^/rest/repositories`)
	if re.MatchString(req.URL.Path) {
		logging.LOGGER.Debug("classified as repositories request")
		repositories.Handle(req)
		return
	}
}
