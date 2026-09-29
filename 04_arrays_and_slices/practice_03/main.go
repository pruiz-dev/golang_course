package main

import (
	"fmt"
)

func main() {
	key := [5]int{1, 2, 3, 4, 5}
	fmt.Printf(
		"Key before:\n%v\n",
		key,
	)

	EncryptKey(&key)

	fmt.Printf(
		"Key after:\n%v\n",
		key,
	)
}

func EncryptKey(arrPtr *[5]int) {
	//for i := range *arrPtr {

	// Go сам разыменует arrPtr
	// перед применением индекса [i]
	for i := range arrPtr {
		(*arrPtr)[i] *= 2
	}
}
