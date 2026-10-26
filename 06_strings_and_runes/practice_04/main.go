package main

import (
	"fmt"
	"runtime"
	"strings"
	"unicode/utf8"
)

func main() {
	token := "TOKEN_9876543210"
	n := 10
	var naiveTokens []string = make([]string, n)
	for i := range n {
		str := generateLargeString(n, token)
		naiveTokens[i] = extractTokenNaive(str, utf8.RuneCountInString(token))
	}

	printMemUsage("После наивных срезов")

	//naiveTokens = nil

	runtime.GC()

	var clonedTokens []string = make([]string, n)

	for i := range n {
		str := generateLargeString(n, token)
		clonedTokens[i] = extractTokenCloned(str, utf8.RuneCountInString(token))
	}

	fmt.Println(naiveTokens[0] == clonedTokens[0])

	printMemUsage("После strings.Clone")
}

func printMemUsage(label string) {
	var m runtime.MemStats
	runtime.GC() // Принудительно вызываем сборщик мусора
	runtime.ReadMemStats(&m)

	// m.Alloc — количество байт в куче, которые реально заняты живыми объектами
	fmt.Printf("[%s] Занято в куче: %d МБ (%d КБ)\n",
		label, m.Alloc/1024/1024, m.Alloc/1024)

}

func generateLargeString(sizeMB int, token string) string {
	var sb strings.Builder
	sb.Grow(1024 * 1024 * sizeMB)
	sb.WriteString(token)
	for range sb.Cap() - 20 {
		sb.WriteByte('a')
	}

	return sb.String()
}

func extractTokenNaive(data string, end int) string {
	return data[0:end]
}

func extractTokenCloned(data string, end int) string {
	return strings.Clone(data[0:end])
}
