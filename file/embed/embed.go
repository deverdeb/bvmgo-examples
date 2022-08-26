package main

import (
	"embed"
	"fmt"
	"log"
)

//go:embed file1.txt
var file1TextContent string

//go:embed file1.txt
var file1BinaryContent []byte

//go:embed file1.txt
var file1Fs embed.FS

//go:embed other-files
var otherFilesFs embed.FS

func main() {
	// Display string
	fmt.Println(file1TextContent)

	// Display bytes
	fmt.Println(string(file1BinaryContent))

	// Display from "embed" file system
	content, err := file1Fs.ReadFile("file1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))

	// Display from "embed" file system
	content, err = otherFilesFs.ReadFile("other-files/file2.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))

	// List directory of "embed" file system
	entries, err := otherFilesFs.ReadDir("other-files")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("other-files directory contains:")
	for _, entry := range entries {
		fmt.Println(" -", entry.Name())
	}
}
