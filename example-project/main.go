package main

import (
	"fmt"

	"github.com/derajohnson/batchrenamer"
)

func main() {
	err := batchrenamer.FileRenamer("./my-files", "velvety")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Files renamed successfully!")
	}
}
