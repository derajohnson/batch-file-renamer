# Batch File Renamer Utility

A reusable Go package for batch renaming files in a directory by appending a specified prefix to their names.

## Features
- `AddPrefix()` renames all files in a given directory by adding a custom prefix.
- `AddSuffix()` renames all files in a given directory by adding a custom suffix.
- `ChangeExtension()` changes the extension of all files in a given directory.
- Renames only files and skips directories
- Validates file names (file names cannot begin with special characters)

## Call the `AddPrefix` Function
The AddPrefix function accepts two arguments:

- directoryPath (string): The path to the directory containing files to rename.
- prefix (string): The prefix to append to file names.

Example usage:

```go
package main

import (
	"fmt"

	"github.com/derajohnson/batchrenamer"
)

func main() {
	err := batchrenamer.AddPrefix("./my-files", "velvety")
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

## Call the `AddSuffix` Function
The AddSuffix function accepts two arguments:

- directoryPath (string): The path to the directory containing files to rename.
- suffix (string): The suffix to append to file names.

Example usage:

```go
package main

import (
	"fmt"

	"github.com/derajohnson/batchrenamer"
)

func main() {
	err := batchrenamer.AddSuffix("./my-files", "there")
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
- `hi.txt`
  
Running the above code will rename the files to:

- `hello_there.txt`
- `hi_there.txt`

## Call the `ChangeExtension` Function
The ChangeExtension function accepts two arguments:

- directoryPath (string): The path to the directory containing files to rename.
- extension (string): The new extension to give the files.

Example usage:

```go
package main

import (
	"fmt"

	"github.com/derajohnson/batchrenamer"
)

func main() {
	err := batchrenamer.ChangeExtension("./my-files", ".md")
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

- `hello.md`
- `world.md`




