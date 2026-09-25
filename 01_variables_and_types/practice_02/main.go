package main

import (
	"fmt"
)

// Имитируем запрос к базе данных
func fetchUser() (string, error) {
	return "Dmitry", nil
}

// Имитируем получение возраста и статуса
func getStats() (int, bool) {
	return 25, true
}

func main() {
	var currentUser string
	var err error

	// ЧАСТЬ 1: Баг с затенением (Shadowing)
	isReady := true
	if isReady {
		// ВНИМАНИЕ: Здесь закрался баг с затенением!
		currentUser, err = fetchUser()
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		fmt.Println("Пользователь загружен локально:", currentUser)
	}

	// ЧАСТЬ 2: Множественное присваивание
	// У нас уже есть переменные age и isActive, но мы хотим их обновить.
	age, isActive := 20, false
	fmt.Printf("Возраст: %d, Статус: %t\n", age, isActive)

	// Твоя задача: правильно использовать функцию getStats(), чтобы обновить age,
	// создать НОВУЮ переменную newStatus, и НЕ использовать ключевое слово var.

	// ДОПИШИ КОД ЗДЕСЬ
	age, newStatus := getStats()

	// Выводы для проверки
	fmt.Printf("Финальный результат: пользователь '%s' авторизован.\n", currentUser)
	fmt.Printf("Возраст: %d, Статус: %t\n", age, newStatus)
}
