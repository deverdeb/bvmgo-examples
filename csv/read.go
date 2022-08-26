package main

import (
	"embed"
	"encoding/csv"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"io"
	"log"
)

//go:embed files
var fileSystem embed.FS

func main() {
	fmt.Println("readCsvWithComma()")
	readCsvWithComma()

	fmt.Println("readCsvWithSemicolon()")
	readCsvWithSemicolon()

	fmt.Println("readCsvUtf8()")
	readCsvUtf8()

	fmt.Println("readCsvWindows1252()")
	readCsvWindows1252()
}

func readCsvWithComma() {
	// file, err := os.Open("csv/files/csv-file-comma.csv")
	file, err := fileSystem.Open("files/csv-file-comma.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	idxLine := 0
	for {
		idxLine++
		record, err := csvReader.Read()
		if err == io.EOF {
			// end of file - stop read
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%3d - len=%3d - data=%v\n", idxLine, len(record), record)
	}
}

func readCsvWithSemicolon() {
	// file, err := os.Open("csv/files/csv-file-semicolon.csv")
	file, err := fileSystem.Open("files/csv-file-semicolon.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	// configure CSV reader
	csvReader.Comma = ';'
	csvReader.Comment = '#'

	idxLine := 0
	for {
		idxLine++
		record, err := csvReader.Read()
		if err == io.EOF {
			// end of file - stop read
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%3d - len=%3d - data=%v\n", idxLine, len(record), record)
	}
}

func readCsvUtf8() {
	// file, err := os.Open("csv/files/csv-file-utf-8.csv")
	file, err := fileSystem.Open("files/csv-file-utf-8.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	idxLine := 0
	for {
		idxLine++
		record, err := csvReader.Read()
		if err == io.EOF {
			// end of file - stop read
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%3d - len=%3d - data=%v\n", idxLine, len(record), record)
	}
}

func readCsvWindows1252() {
	// file, err := os.Open("csv/files/csv-file-windows-1252.csv")
	file, err := fileSystem.Open("files/csv-file-windows-1252.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Windows-1252 to UTF-8 converter
	windows1252Reader := charmap.Windows1252.NewDecoder().Reader(file)

	// Text to CSV parser
	csvReader := csv.NewReader(windows1252Reader)

	idxLine := 0
	for {
		idxLine++
		record, err := csvReader.Read()
		if err == io.EOF {
			// end of file - stop read
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%3d - len=%3d - data=%v\n", idxLine, len(record), record)
	}
}
