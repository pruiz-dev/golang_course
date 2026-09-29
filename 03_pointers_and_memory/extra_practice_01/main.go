package main

import (
	"fmt"
)

func main() {
	x, y := 10, 20
	ptrX, ptrY := &x, &y

	fmt.Printf(
		"Before swap:\n"+
			"x is: %d\n"+
			"y is: %d\n",
		x, y,
	)

	fmt.Printf(
		"Before swap:\n"+
			"ptrX is: %p\n"+
			"ptrY is: %p\n",
		ptrX, ptrY,
	)

	HardcoreSwap(&ptrX, &ptrY)

	fmt.Printf(
		"After swap:\n"+
			"x is: %d\n"+
			"y is: %d\n",
		x, y,
	)

	fmt.Printf(
		"After swap:\n"+
			"ptrX is: %p\n"+
			"ptrY is: %p\n",
		ptrX, ptrY,
	)
}

func HardcoreSwap(a, b **int) {
	*a, *b = *b, *a
}
