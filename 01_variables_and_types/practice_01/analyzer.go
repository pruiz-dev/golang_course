package main

import (
	"fmt"
	"unsafe"
)

func PrintIntInfo(v int) {
	fmt.Printf("%T\n", v)
	fmt.Printf("%d\n", unsafe.Sizeof(v))
}

func PrintInt32Info(v int32) {
	fmt.Printf("%T\n", v)
	fmt.Printf("%d\n", unsafe.Sizeof(v))
}

func PrintInt64Info(v int64) {
	fmt.Printf("%T\n", v)
	fmt.Printf("%d\n", unsafe.Sizeof(v))
}

func PrintStringInfo(v string) {
	fmt.Printf("%T\n", v)
	fmt.Printf("%d\n", unsafe.Sizeof(v))
}

func PrintBoolInfo(v bool) {
	fmt.Printf("%T\n", v)
	fmt.Printf("%d\n", unsafe.Sizeof(v))
}
