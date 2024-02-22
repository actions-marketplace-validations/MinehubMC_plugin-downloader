package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

func PrepareOutputFolder(outputPath string) error {
	// Verify if the output folder exists, otherwise create it
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		fmt.Println("Output folder does not exist. Creating...")
		if err := os.MkdirAll(outputPath, 0755); err != nil {
			return fmt.Errorf("error creating output folder: %v", err)
		}
		fmt.Println("Output folder created successfully.")
	} else if err != nil {
		return fmt.Errorf("error checking output folder: %v", err)
	} else {
		fmt.Println("Output folder already exists.")
	}

	// Check if the current user can read and write to the output folder
	if err := checkPermissions(outputPath); err != nil {
		return fmt.Errorf("error checking permissions: %v", err)
	}

	return nil
}

func checkPermissions(path string) error {
	// Get the absolute path of the output folder
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("error getting absolute path: %v", err)
	}

	// Check if the current user can write to the output folder
	if _, err := os.Stat(absPath); os.IsPermission(err) {
		return fmt.Errorf("current user does not have permission to access %s", absPath)
	}

	return nil
}