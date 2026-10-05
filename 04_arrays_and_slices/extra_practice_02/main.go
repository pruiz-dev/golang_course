package main

import (
	"fmt"
)

func main() {
	// 1. Слайс: Оружие (урон в единицах)
	weapons := []int{10, 20, 30}

	// 2. Массив: Броня (прочность в процентах)
	var armor [4]int = [4]int{50, 30, 10, 85}

	fmt.Printf("До апгрейда оружия: %v\n", weapons)
	fmt.Printf("До ремонта брони:   %v\n", armor)

	upgradeWeapons(weapons)
	repairArmor(&armor) // Передаем указатель, так как массив копируется целиком!

	fmt.Printf("После апгрейда оружия: %v\n", weapons)
	fmt.Printf("После ремонта брони:   %v\n", armor)
}

func upgradeWeapons(items []int) {
	// ЗАДАЧА 1: Исправь этот цикл.
	// Сейчас урон не увеличивается из-за копирования значений!
	for i := range items {
		items[i] += +5
	}
}

func repairArmor(arrPtr *[4]int) {
	// ЗАДАЧА 2: Напиши цикл for range, который обойдет массив по указателю.
	// Прибавь +10 к прочности каждой детали брони.
	// Вспомни про синтаксический сахар для указателей на массивы!
	for i := range arrPtr {
		arrPtr[i] += 10
	}

}
