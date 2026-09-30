package main

import (
	"fmt"
)

func main() {
	hugeFile := make([]byte, 1000000)
	header := hugeFile[:5]
	fmt.Printf(
		"Len of header: %d\n"+
			"Cap of header: %d\n",
		len(header), cap(header),
	)
}
