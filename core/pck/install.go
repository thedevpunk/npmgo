package pck

import (
	"fmt"
	"regexp"
	"sync"
)

func InstallPackage(wg *sync.WaitGroup, name string, version string, destination string) {

	// Cleanup version
	version = getExplicitVersion(version)

	packageInfo, err := fetchPackageInformation(name, version)
	if err != nil {
		fmt.Println("Error fetching package information:", err)
		return
	}

	// Install package dependencies (recursively)
	for name, version := range packageInfo.Dependencies {
		wg.Add(1)
		go InstallPackage(wg, name, version, destination)
	}

	// Download the package
	globalStore := "D:\\_temp\\go_node_modules_test\\global"
	globalPackageDestination := globalStore + fmt.Sprintf("\\%s@%s", name, version)
	tarballPath, err := downloadPackage(packageInfo.Dist.Tarball, packageInfo.Name, packageInfo.Version, globalPackageDestination)
	if err != nil {
		fmt.Println("Error downloading package:", err)
		return
	}

	extractTarball(tarballPath, globalPackageDestination)

	createLink(globalPackageDestination, destination, name)

	deleteTarball(tarballPath)

	wg.Done()
}

func getExplicitVersion(version string) string {
	pattern := `(\d+\.\d+\.\d+)`
	versionRegex := regexp.MustCompile(pattern)
	return versionRegex.FindString(version)
}
