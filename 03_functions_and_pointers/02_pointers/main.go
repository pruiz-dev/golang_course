package main

import (
	"fmt"
)

func main() {
	score := 100

	fmt.Printf("Score is %d\n", score)

	doubleValue(score)

	fmt.Printf("Score is %d\n", score)

	doublePointer(&score)

	fmt.Printf("Score is %d\n", score)
}

func doubleValue(val int) {
	val *= 2
}

func doublePointer(val *int) {
	*val *= 2
}
