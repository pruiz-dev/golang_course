package main

import (
	"fmt"
)

func main() {
	header := getHeaderFromNetwork()
	fmt.Printf(
		"Len of header: %d\n"+
			"Cap of header: %d\n",
		len(header), cap(header),
	)
}

func getHeaderFromNetwork() []byte {
	hugeFile := make([]byte, 1000000)
	safeHeader := make([]byte, 5)
	copy(safeHeader, hugeFile)
	return safeHeader
}
