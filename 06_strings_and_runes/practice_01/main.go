package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Go: Привет 🚀"
	strDiagnostics(s)
	strRange(s)
	if ch, ok := getStrSymbolByIndex(s, 5); ok {
		fmt.Printf("5-й символ: %s\n", ch)
	} else {
		fmt.Println("Индекс за пределами строки!")
	}
}

func getStrSymbolByIndex(s string, i int) (string, bool) {
	if i < 0 {
		return "", false
	}

	counter := -1
	for _, r := range s {
		counter++
		if counter == i {
			return string(r), true
		}
	}

	return "", false
}

func strRange(s string) {
	for idx, r := range s {
		fmt.Printf(
			"Symbol offset in bytes: %d\n"+
				"Bytes size is: %d\n"+
				"Rune: %c (code: %U)\n\n",
			idx, utf8.RuneLen(r), r, r)
	}
}

func strDiagnostics(s string) {
	bytesCount := len(s)
	symbolsCount := utf8.RuneCountInString(s)

	fmt.Printf(
		"Bytes count in str: %d\n"+
			"Symbols count in str: %d\n\n",
		bytesCount, symbolsCount,
	)

	firstByte := s[0]

	fmt.Printf(
		"first byte is: %v\n"+
			"type of firstByte is: %T\n\n",
		firstByte, firstByte,
	)

	// failed code
	fmt.Println(string(s[4]))

}
