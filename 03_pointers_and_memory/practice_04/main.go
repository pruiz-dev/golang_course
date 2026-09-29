package main

import "fmt"

func main() {
	x := 10
	y := 20
	FailedSwap(&x, &y)
	fmt.Printf(
		"After failed swap:\n"+
			"x is: %d\n"+
			"y is: %d\n",
		x, y,
	)

	SuccessSwap(&x, &y)
	fmt.Printf(
		"After success swap:\n"+
			"x is: %d\n"+
			"y is: %d\n",
		x, y,
	)
}

func FailedSwap(a *int, b *int) {
	// var c *int
	// c = a
	// a = b
	// b = c
	a, b = b, a
}

func SuccessSwap(a *int, b *int) {
	// var c int
	// c = *a
	// *a = *b
	// *b = c
	*a, *b = *b, *a
}
