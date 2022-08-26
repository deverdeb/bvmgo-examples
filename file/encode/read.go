package encode

import (
	"bufio"
	"golang.org/x/text/encoding"
	"log"
	"os"
)

// readTextFileAndProcessLineByLine reads a file (UTF-8 encoding) and process file line by line.
func readTextFileAndProcessLineByLine(filename string, encoding encoding.Encoding, lineProcessor func(int, string)) {
	// open file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the function
	defer file.Close()

	// create reader to convert initial encoding to UTF-8
	reader := encoding.NewDecoder().Reader(file)

	// create a scanner to read the file line by line
	scanner := bufio.NewScanner(reader)

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
