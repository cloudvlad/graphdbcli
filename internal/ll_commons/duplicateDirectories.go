// Package ll_commons
// Provides a reusable interfaces for functionally low-level operations.
package ll_commons

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// DuplicateContentInternally handles the data coping recursively from the initial place (a.k.a the source)
// the the internal directory (a.k.a. destination).
func DuplicateContentInternally(src string, dest string) error {
	src = filepath.Clean(src)
	dest = filepath.Clean(dest)

	info, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("source is not a directory: %s", src)
	}

	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		targetPath := filepath.Join(dest, relPath)

		if info.IsDir() {
			if err := os.MkdirAll(targetPath, info.Mode()); err != nil {
				return err
			}
			return nil
		}

		if err := copyFile(path, targetPath, info.Mode()); err != nil {
			return err
		}

		return nil
	})
}

// copyFile handles the coping of files
func copyFile(srcFile string, destFile string, mode os.FileMode) error {
	in, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, mode)
	if err != nil {
		return err
	}
	defer func() {
		_ = out.Close()
	}()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}

	return nil
}
