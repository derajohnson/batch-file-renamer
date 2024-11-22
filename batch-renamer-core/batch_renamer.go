package batchrenamer

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"strings"
)

func readDirectory(directoryPath string) ([]fs.DirEntry, error) {
	// opening the directory
	directory, err := os.Open(directoryPath)
	if err != nil {
		return nil, fmt.Errorf("invalid path to directory: %w", err)
	}

	defer directory.Close()

	// reading the files in the directory
	files, err := directory.ReadDir(-1)
	if err != nil {
		return nil, fmt.Errorf("unable to read directory contents: %w", err)

	}
	return files, nil

}

func AddPrefix(directoryPath string, newName string) error {
	files, _ := readDirectory(directoryPath)

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
	fmt.Printf("Prefix %v added to all files\n", newName)
	return nil

}

func AddSuffix(directoryPath string, newName string) error {
	files, _ := readDirectory(directoryPath)

	for _, file := range files {
		for i := 0; i <= len(files); i++ {
			acceptedCharacters := regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_-]*$`)
			matchExtension := regexp.MustCompile(`^[^.]*`)
			extension := matchExtension.ReplaceAllString(file.Name(), "")
			subString := strings.LastIndex(file.Name(), ".")
			fileName := file.Name()[:subString]
			filePath := fmt.Sprintf("%s/%s", directoryPath, file.Name())
			if acceptedCharacters.MatchString(newName) && !file.IsDir() {
				newFilePath := fmt.Sprintf("%s/%s_%s%s", directoryPath, fileName, newName, extension)
				os.Rename(filePath, newFilePath)
			} else if !acceptedCharacters.MatchString(newName) {
				return errors.New("new file name must not include special characters")
			}

		}

	}
	fmt.Printf("Suffix %v added to all files\n", newName)
	return nil

}

func ChangeExtension(directoryPath string, extension string) error {
	files, _ := readDirectory(directoryPath)
	for _, file := range files {
		for i := 0; i <= len(files); i++ {
			acceptedCharacters := regexp.MustCompile(`^\.[a-zA-Z]+$`)
			subString := strings.LastIndex(file.Name(), ".")
			fileName := file.Name()[:subString]
			filePath := fmt.Sprintf("%s/%s", directoryPath, file.Name())
			if acceptedCharacters.MatchString(extension) && !file.IsDir() {
				newFilePath := fmt.Sprintf("%s/%s%s", directoryPath, fileName, extension)
				os.Rename(filePath, newFilePath)
			} else if !acceptedCharacters.MatchString(extension) {
				return errors.New("extension must not include special characters, numbers and must start with a dot(.)")
			}

		}

	}
	fmt.Printf("Extension %v added to all files\n", extension)
	return nil

}
