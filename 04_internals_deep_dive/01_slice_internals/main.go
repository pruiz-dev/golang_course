package main

import (
	"fmt"
)

func main() {
	nums := make([]int, 0, 3)
	for i := range 7 {
		nums = append(nums, i)
		fmt.Printf("i is: %d, len(nums) is: %d, "+
			"cap(nums) is: %d\n", i, len(nums), cap(nums))
	}
}
