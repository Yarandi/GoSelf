package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	unixPath := filepath.Join("folder", "sub", "sub1", "sample.txt")
	windowsPath := filepath.Join("folder", "sub", "sub1", "sample.txt")

	fmt.Println("Unix-like Path:", unixPath)
	// Output on Unix-like systems: folder/sub/sub1/sample.txt
	fmt.Println("Windows Path:", windowsPath)
	// Output on Windows: folder\sub\sub1\sample.txt
	
}