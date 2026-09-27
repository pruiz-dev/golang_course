package main

import "fmt"

func main() {
	salaries := []int{100, 200, 300}
	fmt.Println("Зарплаты до:", salaries)

	// ПЛОХОЙ КОД: Попытка повысить зарплату всем на 50
	// for _, s := range salaries {
	// 	s = s + 50
	// }

	for i := range salaries {
		salaries[i] += 50
	}

	fmt.Println("Зарплаты после:", salaries)

	// Твоя задача сделать так, чтобы массив salaries реально изменился!
}
