package batchrenamer

import (
	"errors"
	"fmt"
	"os"
	"regexp"
)

func FileRenamer(directoryPath string, newName string) error {
	// opening the directory
	directory, err := os.Open(directoryPath)
	if err != nil {
		return fmt.Errorf("invalid path to directory: %w", err)
	}

	defer directory.Close()

	// reading the files in the directory
	files, err := directory.ReadDir(-1)
	if err != nil {
		return fmt.Errorf("unable to read directory contents: %w", err)

	}

	for _, file := range files {
		for i := 0; i <= len(files); i++ {
			acceptedCharacters := regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_-]*$`)
			filePath := fmt.Sprintf("%s/%s", directoryPath, file.Name())
			if acceptedCharacters.MatchString(newName) && !file.IsDir() {
				newFilePath := fmt.Sprintf("%s/%s_%s", directoryPath, newName, file.Name())
				os.Rename(filePath, newFilePath)
			} else if !acceptedCharacters.MatchString(newName) {
				return errors.New("new file name must not include special characters")
			}

		}

	}
	fmt.Println("All files renamed successfully")
	return nil

}
