package main

func main() {
	var vInt int
	PrintIntInfo(vInt)

	var v32Int int32 = 32
	PrintInt32Info(v32Int)

	var v64Int int64 = 64
	PrintInt64Info(v64Int)

	var vString string = "Hello, Go! Это очень длинная строка для теста."
	PrintStringInfo(vString)

	var vEmptyString string
	PrintStringInfo(vEmptyString)

	var vBool bool = true
	PrintBoolInfo(vBool)
}
