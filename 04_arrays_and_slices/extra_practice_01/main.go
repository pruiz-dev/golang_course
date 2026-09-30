package main

import (
	"fmt"
)

func main() {
	// Инвентарь на 2 слота, оба заняты.
	inventory := []string{"Меч", "Зелье"}

	fmt.Printf("До: %v\n", inventory)

	inventory = addLoot(inventory)

	// Разработчик ожидает увидеть ["Меч", "Зелье", "Золото"]
	// Но видит только старые предметы!
	fmt.Printf("После: %v\n", inventory)
}

func addLoot(items []string) []string {
	// Функция нашла золото и добавляет его в инвентарь
	items = append(items, "Золото")

	return items
}
