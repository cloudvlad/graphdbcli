package repositories

import (
	"net/http"
	"regexp"
)

func Handle(req *http.Request) {
	path := req.URL.Path
	method := req.Method

	// Regex patterns for endpoints
	reRepositories := regexp.MustCompile(`^/rest/repositories$`)
	reRepositoryID := regexp.MustCompile(`^/rest/repositories/([^/]+)$`)
	reRepositoryValidateText := regexp.MustCompile(`^/rest/repositories/([^/]+)/validate/text$`)
	reRepositoryValidateRepo := regexp.MustCompile(`^/rest/repositories/([^/]+)/validate/repository/([^/]+)$`)
	reRepositoryValidateFile := regexp.MustCompile(`^/rest/repositories/([^/]+)/validate/file$`)
	reRepositoryRestart := regexp.MustCompile(`^/rest/repositories/([^/]+)/restart$`)
	reRepositorySize := regexp.MustCompile(`^/rest/repositories/([^/]+)/size$`)

	switch {
	// /rest/repositories
	case reRepositories.MatchString(path) && method == http.MethodGet:
		HandleGetAll(req)
	case reRepositories.MatchString(path) && method == http.MethodPost:
		HandleCreate(req)

	// /rest/repositories/{repositoryID}
	case reRepositoryID.MatchString(path) && method == http.MethodGet:
		HandleGetConfig(req)
	case reRepositoryID.MatchString(path) && method == http.MethodPut:
		HandleEditConfig(req)
	case reRepositoryID.MatchString(path) && method == http.MethodDelete:
		HandleDelete(req)

	// /rest/repositories/{repositoryID}/validate/text
	case reRepositoryValidateText.MatchString(path) && method == http.MethodPost:
		HandleValidateText(req)

	// /rest/repositories/{repositoryID}/validate/repository/{shapesRepositoryID}
	case reRepositoryValidateRepo.MatchString(path) && method == http.MethodPost:
		HandleValidateRepository(req)

	// /rest/repositories/{repositoryID}/validate/file
	case reRepositoryValidateFile.MatchString(path) && method == http.MethodPost:
		HandleValidateFile(req)

	// /rest/repositories/{repositoryID}/restart
	case reRepositoryRestart.MatchString(path) && method == http.MethodPost:
		HandleRestart(req)

	// /rest/repositories/{repositoryID}/size
	case reRepositorySize.MatchString(path) && method == http.MethodGet:
		HandleGetSize(req)

	default:
		HandleNotFound(req)
	}
}

// Placeholder handler functions
func HandleGetAll(req *http.Request)             {}
func HandleGetConfig(req *http.Request)          {}
func HandleEditConfig(req *http.Request)         {}
func HandleDelete(req *http.Request)             {}
func HandleValidateText(req *http.Request)       {}
func HandleValidateRepository(req *http.Request) {}
func HandleValidateFile(req *http.Request)       {}
func HandleRestart(req *http.Request)            {}
func HandleGetSize(req *http.Request)            {}
func HandleNotFound(req *http.Request)           {}
