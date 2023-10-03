package main

import (
	"github.com/thedevpunk/npmgo/pck"
)

// func installGlobalModule(moduleName, moduleVersion string, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	globalModulePath := filepath.Join("/usr/local/node_modules", moduleName)
// 	projectModulePath := filepath.Join("path/to/your/project", "node_modules", moduleName)

// 	_, err := os.Stat(globalModulePath)
// 	if err != nil {
// 		// Module doesn't exist globally, install it.
// 		cmd := exec.Command("npm", "install", moduleName+"@"+moduleVersion, "-g")
// 		cmd.Run()
// 	}

// 	// Create a symlink to the global module in the project's node_modules folder.
// 	os.Symlink(globalModulePath, projectModulePath)
// }

// type PackageInfo struct {
// 	Name         string            `json:"name"`
// 	Version      string            `json:"version"`
// 	Description  string            `json:"description"`
// 	Dependencies map[string]string `json:"dependencies"`
// 	Dist         struct {
// 		Tarball string `json:"tarball"`
// 	} `json:"dist"`
// 	// Add other fields from the JSON response as needed.
// }

// func fetchPackageInformation(name string, version string) (*PackageInfo, error) {
// 	url := fmt.Sprintf("https://registry.npmjs.org/%s/%s", name, version)

// 	// Make an HTTPS GET request.
// 	response, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return nil, err
// 	}
// 	defer response.Body.Close()

// 	// Check the status code of the response.
// 	if response.StatusCode != http.StatusOK {
// 		fmt.Println("HTTP Status Code:", response.Status)
// 		return nil, err
// 	}

// 	// Read the response body.
// 	body, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response body:", err)
// 		return nil, err
// 	}

// 	// Parse JSON into a struct.
// 	var packageInfo PackageInfo
// 	if err := json.Unmarshal(body, &packageInfo); err != nil {
// 		fmt.Println("Error decoding JSON:", err)
// 		return nil, err
// 	}

// 	// Access the parsed data.
// 	fmt.Println("Package Name:", packageInfo.Name)
// 	fmt.Println("Package Version:", packageInfo.Version)
// 	fmt.Println("Package Description:", packageInfo.Description)

// 	// Access the dynamic dependencies.
// 	fmt.Println("Dependencies:")
// 	for key, value := range packageInfo.Dependencies {
// 		fmt.Printf("%s: %v\n", key, value)
// 	}

// 	return &packageInfo, nil
// }

// func downloadPackage(url string, name string, version string) {
// 	destinationDirectory := fmt.Sprintf("D:\\_temp\\go_node_modules_test\\%s@%s", name, version)

// 	// Make an HTTPS GET request.
// 	response, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	defer response.Body.Close()

// 	// Check the status code of the response.
// 	if response.StatusCode != http.StatusOK {
// 		fmt.Println("HTTP Status Code:", response.Status)
// 		return
// 	}

// 	// Create the destination directory if it doesn't exist.
// 	if err := os.MkdirAll(destinationDirectory, 0755); err != nil {
// 		fmt.Println("Error creating destination directory:", err)
// 		return
// 	}

// 	// Create a file to save the downloaded tarball.
// 	tarballPath := filepath.Join(destinationDirectory, fmt.Sprintf("%s-%s.tgz", name, version))
// 	tarballFile, err := os.Create(tarballPath)
// 	if err != nil {
// 		fmt.Println("Error creating tarball file:", err)
// 		return
// 	}
// 	defer tarballFile.Close()

// 	// Copy the tarball data to the file.
// 	_, err = io.Copy(tarballFile, response.Body)
// 	if err != nil {
// 		fmt.Println("Error copying tarball data:", err)
// 		return
// 	}

// 	// Extract the tarball to the destination directory.
// 	if err := extractTarball(tarballPath, destinationDirectory); err != nil {
// 		fmt.Println("Error extracting tarball:", err)
// 		return
// 	}

// 	// if err = deleteTarball(tarballPath); err != nil {
// 	// 	fmt.Println("Error deleting tarball:", err)
// 	// 	return
// 	// }

// 	fmt.Println("Package downloaded and extracted to:", destinationDirectory)
// }

// func installPackage(name string, version string) {
// 	packageInfo, err := pck.FetchPackageInformation(name, version)
// 	if err != nil {
// 		fmt.Println("Error fetching package information:", err)
// 		return
// 	}

// 	// Download the package.
// 	destinationDirectory := fmt.Sprintf("D:\\_temp\\go_node_modules_test\\%s@%s", name, version)
// 	tarballPath, err := pck.DownloadPackage(packageInfo.Dist.Tarball, packageInfo.Name, packageInfo.Version, destinationDirectory)
// 	if err != nil {
// 		fmt.Println("Error downloading package:", err)
// 		return
// 	}

// 	pck.ExtractTarball(tarballPath, destinationDirectory)
// 	pck.DeleteTarball(tarballPath)
// }

// func extractTarball(tarballPath, destination string) error {
// 	fmt.Println("Extracting tarball:", tarballPath)

// 	tarballFile, err := os.Open(tarballPath)
// 	if err != nil {
// 		return err
// 	}
// 	defer tarballFile.Close()

// 	gzipReader, err := gzip.NewReader(tarballFile)
// 	if err != nil {
// 		return err
// 	}
// 	defer gzipReader.Close()

// 	tarReader := tar.NewReader(gzipReader)

// 	for {
// 		header, err := tarReader.Next()
// 		if err == io.EOF {
// 			break // End of tarball
// 		}
// 		if err != nil {
// 			return err
// 		}

// 		targetPath := filepath.Join(destination, strings.TrimPrefix(header.Name, "package/"))
// 		fmt.Println("Extracting file:", targetPath)
// 		switch header.Typeflag {
// 		case tar.TypeDir:
// 			// Create directories.
// 			if err := os.MkdirAll(targetPath, 0755); err != nil {
// 				return err
// 			}
// 		case tar.TypeReg:

// 			// Ensure the parent directory exists.
// 			parentDir := filepath.Dir(targetPath)
// 			if err := os.MkdirAll(parentDir, 0755); err != nil {
// 				return err
// 			}

// 			// Create regular files and copy data.
// 			file, err := os.Create(targetPath)
// 			if err != nil {
// 				return err
// 			}
// 			defer file.Close()

// 			if _, err := io.Copy(file, tarReader); err != nil {
// 				return err
// 			}
// 		default:
// 			return fmt.Errorf("unsupported file type in tarball: %c", header.Typeflag)
// 		}
// 	}

// 	return nil
// }

// func deleteTarball(path string) error {
// 	if err := os.Remove(path); err != nil {
// 		return err
// 	}

// 	return nil
// }

func main() {
	// packages := []struct {
	// 	Name    string
	// 	Version string
	// }{
	// 	{"express", "4.17.1"},
	// 	// Add more modules here
	// }

	packageInfo := pck.PackageInfo{
		Name:    "test-app",
		Version: "1.0.0",
		Dependencies: map[string]string{
			"express": "4.17.1",
		},
	}

	pck.Install(packageInfo)

	// var wg sync.WaitGroup

	// for _, pack := range packages {
	// 	wg.Add(1)

	// 	go pck.InstallPackage(&wg, pack.Name, pack.Version, "D:\\_temp\\go_node_modules_test\\node_modules")
	// }

	// wg.Wait()
}
