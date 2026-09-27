package main

import (
	"fmt"
)

func main() {
	var vUint8 uint8 = 255
	vUint8 += 1
	fmt.Printf("%d\n", vUint8)

	var vInt8 int8 = 127
	vInt8 += 1
	fmt.Printf("%d\n", vInt8)

	var vvInt16 int16 = 1000
	var vvInt8 = int8(vvInt16)
	fmt.Printf("%d\n", vvInt8)

	var vFl32 float32 = 2.99
	fmt.Printf("%d\n", int(vFl32))

	//var impossible int8 = 1000
}
