package batchrenamer

import (
	"fmt"
	"os"
	"regexp"
)

func FileRenamer(directoryPath string, newName string) {
	// opening the directory
	directory, err := os.Open(directoryPath)
	if err != nil {
		fmt.Println("Error invalid path to directory")
	}

	defer directory.Close()

	// reading the files in the directory
	files, err := directory.ReadDir(-1)
	if err != nil {
		fmt.Println("Error directory does not contain files")
	}

	for _, file := range files {
		for i := 0; i <= len(files); i++ {
			acceptedCharacters := regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_-]*$`)
			filePath := fmt.Sprintf("%s/%s", directoryPath, file.Name())
			if acceptedCharacters.MatchString(newName) && !file.IsDir() {
				newFilePath := fmt.Sprintf("%s/%s_%s", directoryPath, newName, file.Name())
				os.Rename(filePath, newFilePath)
				fmt.Println("Files renamed successfully")

			} else {
				fmt.Println("File name must not include a special character")
			}

		}

	}

}
