package main

import (
	"fmt"
)

func main() {
	equipment := [3]string{"Шлем", "Броня", "Оружие"}
	inventory := []string{"Зелье", "Свисток"}

	fmt.Printf("Equipment before: %v\n", equipment)
	fmt.Printf("Inventory before: %v\n", inventory)

	BreakEquipment(equipment)
	UseItem(inventory)

	fmt.Printf("Equipment after: %v\n", equipment)
	fmt.Printf("Inventory after: %v\n", inventory)

}

func BreakEquipment(eq [3]string) {
	eq[0] = "Сломанный шлем"
}

func UseItem(inv []string) {
	inv[0] = "Пустая колба"
}
