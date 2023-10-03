package pck

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func downloadPackage(url string, name string, version string, destinationDirectory string) (destionationPath string, err error) {
	// destinationDirectory := fmt.Sprintf("D:\\_temp\\go_node_modules_test\\%s@%s", name, version)

	// Make an HTTPS GET request.
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	defer response.Body.Close()

	// Check the status code of the response.
	if response.StatusCode != http.StatusOK {
		fmt.Println("HTTP Status Code:", response.Status)
		return "", fmt.Errorf("httpresponse: response status code is not ok, it is %s", response.Status)
	}

	// Create the destination directory if it doesn't exist.
	if err := os.MkdirAll(destinationDirectory, 0755); err != nil {
		fmt.Println("Error creating destination directory:", err)
		return "", err
	}

	// Create a file to save the downloaded tarball.
	tarballPath := filepath.Join(destinationDirectory, fmt.Sprintf("%s-%s.tgz", name, version))
	tarballFile, err := os.Create(tarballPath)
	if err != nil {
		fmt.Println("Error creating tarball file:", err)
		return "", err
	}
	defer tarballFile.Close()

	// Copy the tarball data to the file.
	_, err = io.Copy(tarballFile, response.Body)
	if err != nil {
		fmt.Println("Error copying tarball data:", err)
		return "", err
	}

	// Extract the tarball to the destination directory.
	// if err := extractTarball(tarballPath, destinationDirectory); err != nil {
	// 	fmt.Println("Error extracting tarball:", err)
	// 	return
	// }

	fmt.Println("Package downloaded to:", tarballPath)
	return tarballPath, nil

	// if err = deleteTarball(tarballPath); err != nil {
	// 	fmt.Println("Error deleting tarball:", err)
	// 	return
	// }

}
