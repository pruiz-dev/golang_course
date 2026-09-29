package main

//import "fmt"

func main() {
	// var v1 int = CreateValue()
	// var v2 *int = CreatePointer()

	// fmt.Printf("v1 is: %d\n", v1)
	// fmt.Printf("value of v2 is: %d\n", *v2)

	_ = CreateValue()
	_ = CreatePointer()
}

func CreateValue() int {
	var v int = 42
	return v
}

func CreatePointer() *int {
	var v int = 99
	return &v
}

/*
# practice_03
./main.go:16:6: can inline CreateValue
./main.go:21:6: can inline CreatePointer
./main.go:5:6: can inline main
./main.go:12:17: inlining call to CreateValue
./main.go:13:19: inlining call to CreatePointer
./main.go:22:6: moved to heap: v
*/
