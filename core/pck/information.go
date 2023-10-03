package pck

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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

func fetchPackageInformation(name string, version string) (*PackageInfo, error) {
	url := fmt.Sprintf("https://registry.npmjs.org/%s/%s", name, version)

	// Make an HTTPS GET request.
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer response.Body.Close()

	// Check the status code of the response.
	if response.StatusCode != http.StatusOK {
		fmt.Println("HTTP Status Code:", response.Status)
		return nil, err
	}

	// Read the response body.
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	// Parse JSON into a struct.
	var packageInfo PackageInfo
	if err := json.Unmarshal(body, &packageInfo); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}

	// // Access the parsed data.
	// fmt.Println("Package Name:", packageInfo.Name)
	// fmt.Println("Package Version:", packageInfo.Version)
	// fmt.Println("Package Description:", packageInfo.Description)

	// Access the dynamic dependencies.
	// fmt.Println("Dependencies:")
	// for key, value := range packageInfo.Dependencies {
	// 	fmt.Printf("%s: %v\n", key, value)
	// }

	fmt.Printf("Package %s@%s fetched.\n", packageInfo.Name, packageInfo.Version)
	return &packageInfo, nil
}
