package main

import (
	"fmt"
	"math/rand"
)

// Имитируем запрос к серверу
func fetchStatusCode() int {
	codes := []int{200, 201, 400, 404, 500, 503, 999}
	return codes[rand.Intn(len(codes))]
}

func main() {
	// ПЛОХОЙ КОД: Переменная code торчит снаружи, хотя нужна только для проверок
	//code := fetchStatusCode()

	switch code := fetchStatusCode(); code {
	case 200, 201:
		fmt.Printf("Код %d: Успех\n", code)
	case 400, 404:
		fmt.Printf("Код %d: Ошибка клиента\n", code)
	case 500, 503:
		fmt.Printf("Код %d: Ошибка сервера. Требуется перезапуск.\n", code)
	default:
		fmt.Printf("Код %d: Неизвестный статус\n", code)
	}

	// if code == 200 || code == 201 {
	// 	fmt.Printf("Код %d: Успех\n", code)
	// } else if code == 400 || code == 404 {
	// 	fmt.Printf("Код %d: Ошибка клиента\n", code)
	// } else if code == 500 || code == 503 {
	// 	fmt.Printf("Код %d: Ошибка сервера. Требуется перезапуск.\n", code)
	// } else {
	// 	fmt.Printf("Код %d: Неизвестный статус\n", code)
	// }

	// ПЛОХОЙ КОД: Мы всё еще имеем доступ к code здесь, хотя она нам больше не нужна.
	//fmt.Println("Последний обработанный код был:", code)
}
