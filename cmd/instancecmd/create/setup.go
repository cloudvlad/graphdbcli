package create

import (
	"archive/zip"
	"fmt"
	cc "graphdbcli/internal/channels/commons"
	c "graphdbcli/internal/data_objects/graphdb_cluster"
	ini "graphdbcli/internal/tool_configurations/initialization"
	"graphdbcli/internal/tool_configurations/logging"
	"graphdbcli/internal/tool_configurations/properties_manager"
	tc "graphdbcli/internal/tool_configurations/statics"
	"graphdbcli/internal/tui/common_components"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/enescakir/emoji"
	"go.uber.org/zap"
)

// ConfigureInstance extracts a zip file to the node directory, skipping the
// top-level folder (usually the zip's name).
func ConfigureInstance(zipFilePath, nodeFolder string) error {
	_, err := os.Stat(zipFilePath)
	if os.IsNotExist(err) {
		fmt.Printf("%s The distribution package for GraphDb is missing", emoji.Collision)
		logging.LOGGER.Fatal("missing distribution package", zap.String("version", c.Instance.Version))
	}
	r, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		parts := strings.SplitN(f.Name, "/", 2)
		var relativePath string
		if len(parts) == 2 {
			relativePath = parts[1]
		} else {
			relativePath = parts[0]
		}
		if relativePath == "" {
			continue
		}
		fpath := filepath.Join(nodeFolder, relativePath)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		rc, err := f.Open()
		if err != nil {
			outFile.Close()
			return err
		}
		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func configureInstancePort(port string) {
	propertyName := properties_manager.GDB10_8["instancePort"]
	properties_manager.FindAndReplacePropertie(&c.Instance.PropertyOverrides, propertyName, port)
}

// configureInstanceLicense will add a specified, previously stored, license
// It will store it with the default name that GraphDB expects,
// ignoring the name set when storing it.
func configureInstanceLicense(licenseName, instanceName string) {
	fmt.Printf("%s Using license file %s\n", common_components.PadStatusIndicator(emoji.Information.String(), tc.NotTUIStatusIndicatorAdditionalPadding), c.Instance.StoredLicenseFilename)
	licensePath := path.Join(ini.GetLicensesDirectory(), licenseName)
	license, err := os.ReadFile(licensePath)
	if err != nil {
		fmt.Printf("%s Error occured whilst reading license file", common_components.PadStatusIndicator(emoji.Collision.String(), tc.NotTUIStatusIndicatorAdditionalPadding))
		logging.LOGGER.Fatal("Error occurred when reading license file",
			zap.String("license_name", licenseName),
			zap.Error(err))
	}

	_, err = os.Stat(licensePath)
	if err != nil {
		fmt.Printf("%s There was problem reading license file", emoji.Collision)
		logging.LOGGER.Fatal("Error occured when reading license file",
			zap.String("license_name", licenseName),
			zap.Error(err))
	}

	segments := append([]string{ini.GetClustersDirectory(), instanceName}, tc.DefaultLicensePath...)
	nodeLicenseFullPath := path.Join(segments...)

	logging.LOGGER.Debug("Configuring licesense for instance", zap.String("instanceName", instanceName))

	err = os.WriteFile(nodeLicenseFullPath, license, os.ModePerm)
	if err != nil {
		fmt.Printf("%s Error occured whilst writing license file", emoji.Collision)
		logging.LOGGER.Fatal("Error occured when writing license file",
			zap.Error(err))
	}
}

func setupBaseStructure(instancePath, zipFilePath string, failureChannel *chan bool) {
	if err := os.Mkdir(instancePath, 0700); err != nil {
		cc.HandleEvent(failureChannel, p)
		return
	}

	if err := ConfigureInstance(zipFilePath, instancePath); err != nil {
		cc.HandleEvent(failureChannel, p)
		return
	}
}
