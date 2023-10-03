package pck

import (
	"fmt"
	"os"
)

func createSymlink(source string, nodeModulesDir string, modName string) error {

	// Create the destination directory if it doesn't exist.
	if err := os.MkdirAll(nodeModulesDir, 0755); err != nil {
		fmt.Println("Error creating destination directory:", err)
		return err
	}

	modPath := fmt.Sprintf("%s\\%s", nodeModulesDir, modName)

	err := os.Symlink(source, modPath)
	if err != nil {
		fmt.Println("Error creating symlink:", err)
		return err
	}

	fmt.Println("Symlink created:", modPath)

	return nil
}
