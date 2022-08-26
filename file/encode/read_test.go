package encode

import (
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/unicode"
)

func Example_readTextFileAndProcessLineByLine_iso_8859_1() {
	readTextFileAndProcessLineByLine("iso-8859-1.txt", charmap.ISO8859_1,
		func(_ int, line string) {
			fmt.Println(line)
		})
	// Output:
	// Test ISO-8859-1
	// abcdefghijklmnopqrstuvwxyz
	// ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// 0123456789
	// +-*/+²#"'{([|_\@)]}°
	// àâäçéèëêìïîñòôöùûüÿ
	// ÀÂÄÈËÊÌÏÎÑÒÔÖÙÜÛ
	// µ£§
}

func Example_readTextFileAndProcessLineByLine_iso_8859_15() {
	readTextFileAndProcessLineByLine("iso-8859-15.txt", charmap.ISO8859_15,
		func(_ int, line string) {
			fmt.Println(line)
		})
	// Output:
	// Test ISO-8859-15
	// abcdefghijklmnopqrstuvwxyz
	// ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// 0123456789
	// +-*/+²#"'{([|_\@)]}°
	// àâäçéèëêìïîñòôöùûüÿ
	// ÀÂÄÈËÊÌÏÎÑÒÔÖÙÜÛ
	// µ£§€
}

func Example_readTextFileAndProcessLineByLine_UTF_8() {
	readTextFileAndProcessLineByLine("UTF-8.txt", unicode.UTF8,
		func(_ int, line string) {
			fmt.Println(line)
		})
	// Output:
	// Test UTF-8
	// abcdefghijklmnopqrstuvwxyz
	// ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// 0123456789
	// +-*/+²#"'{([|_\@)]}°
	// àâäçéèëêìïîñòôöùûüÿ
	// ÀÂÄÈËÊÌÏÎÑÒÔÖÙÜÛ
	// µ£§€
}
