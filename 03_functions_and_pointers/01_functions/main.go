package main

import (
	"fmt"
)

func main() {
	sum, diff := calculate(5, -6)

	fmt.Printf("Summ is: %d, Difference is %d", sum, diff)
}

func calculate(a, b int) (int, int) {
	return a + b, a - b
}
