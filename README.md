# Batch File Renamer Utility

A reusable Go package for batch renaming files in a directory by appending a specified prefix to their names.

## Features
- Renames all files in a given directory by adding a custom prefix.
- Renames only files and skips directories
- Validates file names (file names cannot begin with special characters)

## Call the `FileRenamer` Function
The FileRenamer function accepts two arguments:

- directoryPath (string): The path to the directory containing files to rename.
- newName (string): The prefix to append to file names.

Example usage:

```go
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

```

Example Output:

For a directory `./my-files` containing files:

- `hello.txt`
- `world.txt`
  
Running the above code will rename the files to:

- `velvety_hello.txt`
- `velvety_world.txt`
