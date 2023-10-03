package pck

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sync"
)

var storePath = "D:\\_temp\\go_node_modules_test\\global"
var storeCache = make(map[string]struct{})
var storeCacheMu sync.RWMutex

type PackageInfo struct {
	Name         string            `json:"name"`
	Version      string            `json:"version"`
	Description  string            `json:"description"`
	Dependencies map[string]string `json:"dependencies"`
	Dist         struct {
		Tarball string `json:"tarball"`
	} `json:"dist"`
	// Add other fields from the JSON response as needed.
}

func Install(info PackageInfo) {
	readPackagesIntoCache()

	var wg sync.WaitGroup

	for name, version := range info.Dependencies {
		wg.Add(1)

		go installPackage(&wg, name, version, "D:\\_temp\\go_node_modules_test\\node_modules")
	}

	wg.Wait()
}

func installPackage(wg *sync.WaitGroup, name string, version string, destination string) {

	// Cleanup version
	version = getExplicitVersion(version)

	if isPackageInCache(name, version) {
		wg.Done()
		return
	}

	packageInfo, err := fetchPackageInformation(name, version)
	if err != nil {
		fmt.Println("Error fetching package information:", err)
		return
	}

	// Install package dependencies (recursively)
	for name, version := range packageInfo.Dependencies {
		wg.Add(1)
		go installPackage(wg, name, version, destination)
	}

	// Download the package
	globalPackageDestination := storePath + fmt.Sprintf("\\%s@%s", name, version)
	tarballPath, err := downloadPackage(packageInfo.Dist.Tarball, packageInfo.Name, packageInfo.Version, globalPackageDestination)
	if err != nil {
		fmt.Println("Error downloading package:", err)
		return
	}

	extractTarball(tarballPath, globalPackageDestination)

	createLink(globalPackageDestination, destination, name)

	deleteTarball(tarballPath)

	addPackageToCache(name, version)

	fmt.Printf("Installed: %s@%s\n", name, version)

	wg.Done()
}

func getExplicitVersion(version string) string {
	pattern := `(\d+\.\d+\.\d+)`
	versionRegex := regexp.MustCompile(pattern)
	return versionRegex.FindString(version)
}

// func isPackageInstalled(name string, version string) (bool, error) {
// 	packageDir := fmt.Sprintf("\\%s@%s", name, version)

// 	cmd := exec.Command("npm", "list", "--global")

// 	// Capture the command output
// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		return false, err
// 	}

// 	// Check if the package name is found in the output
// 	return strings.Contains(string(output), packageDir), nil
// }

// func isPackageInstalled(name string, version string) bool {
// 	packageDir := fmt.Sprintf("\\%s@%s", name, version)

// 	// Construct the path where the package should be installed
// 	packagePath := filepath.Join(storePath, packageDir)

// 	// Check if the package directory exists
// 	_, err := os.Stat(packagePath)
// 	if err == nil {
// 		fmt.Printf("Package directory exists: %s\n", packagePath)
// 		return true // Package directory exists
// 	} else if os.IsNotExist(err) {
// 		return false // Package directory does not exist
// 	}

// 	return false // Error occurred
// }

func addPackageToCache(name string, version string) {

	storeCacheMu.Lock()
	storeCache[fmt.Sprintf("%s@%s", name, version)] = struct{}{}
	storeCacheMu.Unlock()
	fmt.Printf("Stored in cache: %s@%s\n", name, version)
}

func isPackageInCache(name string, version string) bool {
	packageDir := fmt.Sprintf("%s@%s", name, version)
	fmt.Println(storeCache)

	storeCacheMu.RLock()
	_, exists := storeCache[packageDir]
	storeCacheMu.RUnlock()

	if exists {
		fmt.Printf("Package exists in cache: %s\n", packageDir)
		return true
	}

	fmt.Printf("Package does not exist in cache: %s\n", packageDir)
	return false
}

func readPackagesIntoCache() {
	if _, err := os.Stat(storePath); os.IsNotExist(err) {
		return
	}

	// Read the contents of the directory
	dirs, err := os.ReadDir(storePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Filter out directories
	for _, file := range dirs {
		storeCacheMu.Lock()
		storeCache[file.Name()] = struct{}{}
		storeCacheMu.Unlock()

		fmt.Printf("store: %s\n", file.Name())
	}
}
