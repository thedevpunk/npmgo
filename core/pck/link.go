package pck

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func createLink(source string, nodeModulesDir string, modName string) error {

	// Create the destination directory if it doesn't exist.
	if err := os.MkdirAll(nodeModulesDir, 0755); err != nil {
		fmt.Println("Error creating destination directory:", err)
		return err
	}

	modPath := fmt.Sprintf("%s\\%s", nodeModulesDir, modName)

	if runtime.GOOS == "windows" {
		return createJunction(source, modPath)
	} else {
		return createSymlink(source, modPath)
	}
}

func createSymlink(source string, destionation string) error {
	err := os.Symlink(source, destionation)
	if err != nil {
		fmt.Println("Error creating symlink:", err)
		return err
	}

	fmt.Println("Symlink created:", destionation)

	return nil
}

func createJunction(source string, destionation string) error {
	cmd := exec.Command("cmd.exe", "/C", "mklink", "/J", destionation, source)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error creating junction: %v\n", err)
		fmt.Println(string(output))
		return err
	}

	fmt.Println("Junction created:", destionation)

	return nil
}
