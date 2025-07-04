package fetch

import (
	"errors"
	"fmt"
	ini "graphdbcli/internal/tool_configurations/initialization"
	"graphdbcli/internal/tool_configurations/statics"
	"io"
	"net/http"
	"os"
	"path"
)

func findResourceByName(name string) (*statics.Resource, error) {
	for _, r := range statics.Resources {
		if r.Name == name {
			return &r, nil
		}
	}
	return nil, errors.New("resource not found")
}

func FetchResource(resource statics.Resource) error {
	url := fmt.Sprintf("https://ipfs.io/ipfs/%s", resource.IpfsCID)
	fmt.Printf("Downloading %s from %s...\n", resource.Name, url)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download: status %s", resp.Status)
	}
	resourceFullpath := path.Join(ini.GetResourcesDirectory(), resource.Name)
	out, err := os.Create(resourceFullpath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}

	fmt.Printf("Resource saved as %s\n", resource.Name)
	return nil
}
