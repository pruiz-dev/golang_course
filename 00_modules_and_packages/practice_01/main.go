package main

import (
	"fmt"
	"shop_app/discount"
)

func main() {
	discountPrice := discount.Calculate(100)
	fmt.Printf("Discount price is: %d\n", discountPrice)
}
