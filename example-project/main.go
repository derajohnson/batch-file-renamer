package main

import (
	"fmt"

	"github.com/derajohnson/batchrenamer"
)

func main() {
	prefixErr := batchrenamer.AddPrefix("./my-files", "open")
	if prefixErr != nil {
		fmt.Println("Error:", prefixErr)
	} else {
		fmt.Println("Prefix added successfully!")
	}

	suffixErr := batchrenamer.AddSuffix("./my-files", "source")
	if suffixErr != nil {
		fmt.Println("Error:", suffixErr)
	} else {
		fmt.Println("Suffix added successfully!")
	}

	extensionErr := batchrenamer.ChangeExtension("./my-files", ".md")
	if extensionErr != nil {
		fmt.Println("Error:", suffixErr)
	} else {
		fmt.Println("Extension added successfully!")
	}
}
