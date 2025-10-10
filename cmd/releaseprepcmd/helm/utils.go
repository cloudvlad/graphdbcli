package helm

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

// Parse YAML file
type UpdateConf struct {
	Version       string   `yaml:"version"`
	AppVersion    string   `yaml:"appVersion"`
	Documentation []string `yaml:"documentation"`
}

func prepare(updateConfFilePath, updatedTargetDirPath string, removeConfFileWhenDone bool) {
	readConfFile(updateConfFilePath)
	prepareDocumentation(updateConfFilePath)
}

func readConfFile(confFilePath string) UpdateConf {
	data, err := os.ReadFile(confFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read YAML file: %v\n", err)
		return nil
	}

	var conf UpdateConf
	if err := yaml.Unmarshal(data, &conf); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse YAML: %v\n", err)
		return nil
	}

	return conf
}

// prepareDocumentation parses the YAML config and updates documentation files as required.
func prepareDocumentation(confFilePath string) {
	conf := readConfFile(confFilePath)
	appVersion := conf.AppVersion

	parts := strings.Split(appVersion, ".")
	appVersionShort := parts[0] + "." + parts[1]

	for _, docPath := range conf.Documentation {
		// Try to read the file
		content, err := os.ReadFile(docPath)

		updated := string(content)
		
		// Replace GraphDB documentation URLs
		updated = regexp.MustCompile(`https://graphdb\.ontotext\.com/documentation/[\d\.]+/([\w\-]+\.html)`).ReplaceAllStringFunc(updated, func(url string) string {
			return regexp.MustCompile(`[\d\.]+`).ReplaceAllString(url, appVersionShort)
		})

		// Replace release-notes URLs
		updated = regexp.MustCompile(`(https://graphdb\.ontotext\.com/documentation/[\d\.]+/release-notes\.html)(?!#)`).ReplaceAllString(updated, "$1#"+appVersion)

		// Write back
		err = os.WriteFile(docPath, []byte(updated), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write %s: %v\n", docPath, err)
		}
	}
}
