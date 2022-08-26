package text

import (
	"bufio"
	"log"
	"os"
)

// readTextFile reads a file (UTF-8 encoding) and it returns all file in string.
// Warning: it does not supporte large files.
func readTextFile(filename string) string {
	// read file
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	// convert to string
	return string(content)
}

// readTextFileLineByLine reads a file (UTF-8 encoding) and it returns all file in string slice.
// Warning: it does not supporte large files.
func readTextFileLineByLine(filename string) []string {
	// open file
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the function
	defer f.Close()

	// create a scanner to read the file line by line
	scanner := bufio.NewScanner(f)

	// read file line by line
	result := make([]string, 0, 0)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

// readTextFileAndProcessLineByLine reads a file (UTF-8 encoding) and process file line by line.
func readTextFileAndProcessLineByLine(filename string, lineProcessor func(int, string)) {
	// open file
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the function
	defer f.Close()

	// create a scanner to read the file line by line
	scanner := bufio.NewScanner(f)

	// read file line by line
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		lineProcessor(lineNumber, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
