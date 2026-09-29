package main

import (
	"fmt"
	deposit "practice_02/bank"
)

func main() {
	var balance int = 1000
	deposit.Deposit(
		&balance,
		500,
	)
	fmt.Printf(
		"Deposit is: "+
			"%d\n",
		balance,
	)

	var emptyPtr *int
	deposit.Deposit(
		emptyPtr,
		500,
	)

	fmt.Println(
		"Программа успешно" +
			" завершена",
	)
}
