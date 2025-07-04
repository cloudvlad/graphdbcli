package versiontui

import (
	"context"
	"fmt"
	channels "graphdbcli/internal/channels/commons"
	"graphdbcli/internal/tool_configurations/initialization"
	"graphdbcli/internal/tool_configurations/statics"
	"io"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/enescakir/emoji"
)

func NotTUIinstallVersion(ctx context.Context, version statics.Version) error {
	url := fmt.Sprintf("https://ipfs.io/ipfs/%s", version.IpfsCID)
	fmt.Printf("%s Downloading GraphDB %s... \n", emoji.Watch, version.Version)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		channels.Success <- false
		return fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		channels.Success <- false
		return fmt.Errorf("failed to fetch the distribution file: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		channels.Success <- false
		return fmt.Errorf("failed to fetch resource: status %s", resp.Status)
	}

	zipFullpath := path.Join(initialization.GetDistDirectory(), "graphdb-"+version.Version+".zip")
	out, err := os.Create(zipFullpath)
	if err != nil {
		channels.Success <- false
		return fmt.Errorf("failed to create file: %w", err)
	}

	defer func() {
		out.Close()
		if err != nil {
			os.Remove(zipFullpath)
		}
	}()

	buf := make([]byte, 32*1024)
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s Download cancelled\n", emoji.StopSign)
			err = ctx.Err()
			return err
		case <-ticker.C:
			fmt.Printf("%s Download is still in progress...\n", emoji.HourglassNotDone)
		default:
			n, readErr := resp.Body.Read(buf)
			if n > 0 {
				if _, writeErr := out.Write(buf[:n]); writeErr != nil {
					err = writeErr
					IsFileDownloadedSuccessful = false
					return fmt.Errorf("failed to save file: %w", writeErr)
				}
			}
			if readErr == io.EOF {
				IsFileDownloadedSuccessful = true
				return nil
			}
			if readErr != nil {
				err = readErr
				IsFileDownloadedSuccessful = false
				return fmt.Errorf("failed to read response: %w", readErr)
			}
		}

	}
}
