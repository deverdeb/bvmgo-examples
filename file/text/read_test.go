package text

import (
	"fmt"
	"testing"
)

// Example_readTextFile reads a file, prints file content to output and checks output.
// WARNING: line separators are important...
func Example_readTextFile() {
	fmt.Println(readTextFile("file.txt"))
	// Output:
	// Lorem ipsum dolor sit amet, consectetur adipiscing elit.
	// Suspendisse pellentesque nisl at lacus mollis iaculis.
	// Integer id elit leo. Duis aliquam eleifend lorem in euismod.
	// Quisque lacinia ac ex nec convallis.
	// Mauris elementum justo quis lacus elementum, id molestie arcu tempor.
	// Nullam eleifend eu massa eu varius.
	// Cras finibus ligula non turpis mattis, non sodales odio tincidunt.
	// Ut placerat sapien dolor, nec pharetra dui rutrum ac.
	// Nam ullamcorper lacinia nulla.
	// Vestibulum pellentesque, augue a convallis iaculis, sem augue nunc, vitae hendrerit purus tellus viverra ligula.
	// Cras sem nunc, euismod vitae commodo eu, aliquam id ligula.
	// Quisque molestie erat bibendum sem vestibulum, varius fermentum justo mollis.
}

// Example_readTextFileLineByLine reads a file, prints file lines to output and checks output.
func Example_readTextFileLineByLine() {
	lines := readTextFileLineByLine("file.txt")
	for idx, line := range lines {
		fmt.Println(idx+1, line)
	}
	// Output:
	// 1 Lorem ipsum dolor sit amet, consectetur adipiscing elit.
	// 2 Suspendisse pellentesque nisl at lacus mollis iaculis.
	// 3 Integer id elit leo. Duis aliquam eleifend lorem in euismod.
	// 4 Quisque lacinia ac ex nec convallis.
	// 5 Mauris elementum justo quis lacus elementum, id molestie arcu tempor.
	// 6 Nullam eleifend eu massa eu varius.
	// 7 Cras finibus ligula non turpis mattis, non sodales odio tincidunt.
	// 8 Ut placerat sapien dolor, nec pharetra dui rutrum ac.
	// 9 Nam ullamcorper lacinia nulla.
	// 10 Vestibulum pellentesque, augue a convallis iaculis, sem augue nunc, vitae hendrerit purus tellus viverra ligula.
	// 11 Cras sem nunc, euismod vitae commodo eu, aliquam id ligula.
	// 12 Quisque molestie erat bibendum sem vestibulum, varius fermentum justo mollis.
}

// Example_readTextFileAndProcessLineByLine reads a file, prints file lines to output and checks output.
func Example_readTextFileAndProcessLineByLine() {
	readTextFileAndProcessLineByLine("file.txt",
		func(idx int, line string) {
			fmt.Println(idx, line)
		})
	// Output:
	// 1 Lorem ipsum dolor sit amet, consectetur adipiscing elit.
	// 2 Suspendisse pellentesque nisl at lacus mollis iaculis.
	// 3 Integer id elit leo. Duis aliquam eleifend lorem in euismod.
	// 4 Quisque lacinia ac ex nec convallis.
	// 5 Mauris elementum justo quis lacus elementum, id molestie arcu tempor.
	// 6 Nullam eleifend eu massa eu varius.
	// 7 Cras finibus ligula non turpis mattis, non sodales odio tincidunt.
	// 8 Ut placerat sapien dolor, nec pharetra dui rutrum ac.
	// 9 Nam ullamcorper lacinia nulla.
	// 10 Vestibulum pellentesque, augue a convallis iaculis, sem augue nunc, vitae hendrerit purus tellus viverra ligula.
	// 11 Cras sem nunc, euismod vitae commodo eu, aliquam id ligula.
	// 12 Quisque molestie erat bibendum sem vestibulum, varius fermentum justo mollis.
}

func Benchmark_readTextFile_small_file(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = readTextFile("file.txt")
	}
}

func Benchmark_readTextFileLineByLine_small_file(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = readTextFileLineByLine("file.txt")
	}
}

func Benchmark_readTextFileAndProcessLineByLine_small_file(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readTextFileAndProcessLineByLine("file.txt", func(_ int, _ string) {})
	}
}

func Benchmark_readTextFile_large_file(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = readTextFile("large-file.txt")
	}
}

func Benchmark_readTextFileLineByLine_large_file(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = readTextFileLineByLine("large-file.txt")
	}
}

func Benchmark_readTextFileAndProcessLineByLine_large_file(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readTextFileAndProcessLineByLine("large-file.txt", func(_ int, _ string) {})
	}
}
