package pck

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func extractTarball(tarballPath, destination string) error {
	fmt.Println("Extracting tarball:", tarballPath)

	tarballFile, err := os.Open(tarballPath)
	if err != nil {
		return err
	}
	defer tarballFile.Close()

	gzipReader, err := gzip.NewReader(tarballFile)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break // End of tarball
		}
		if err != nil {
			return err
		}

		targetPath := filepath.Join(destination, strings.TrimPrefix(header.Name, "package/"))
		fmt.Println("Extracting file:", targetPath)
		switch header.Typeflag {
		case tar.TypeDir:
			// Create directories.
			if err := os.MkdirAll(targetPath, 0755); err != nil {
				return err
			}
		case tar.TypeReg:

			// Ensure the parent directory exists.
			parentDir := filepath.Dir(targetPath)
			if err := os.MkdirAll(parentDir, 0755); err != nil {
				return err
			}

			// Create regular files and copy data.
			file, err := os.Create(targetPath)
			if err != nil {
				return err
			}
			defer file.Close()

			if _, err := io.Copy(file, tarReader); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported file type in tarball: %c", header.Typeflag)
		}
	}

	return nil
}

func deleteTarball(path string) error {
	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}
