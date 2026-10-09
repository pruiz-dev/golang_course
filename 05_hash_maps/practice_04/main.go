package main

import (
	"fmt"
	"time"
)

func main() {
	const elementsCount = 5_000_000

	// ==========================================
	// 1. Медленная мапа (постоянные эвакуации)
	// ==========================================
	slowMap := make(map[int]int)

	startSlow := time.Now() // Засекаем время

	// TODO: Напиши цикл на 5 миллионов итераций,
	// который кладет данные в slowMap.

	for i := range 5000000 {
		slowMap[i] = i * 2
	}

	durationSlow := time.Since(startSlow) // Считаем потраченное время
	fmt.Printf("Медленная мапа отработала за: %v\n", durationSlow)

	// ==========================================
	// 2. Быстрая мапа (без эвакуаций)
	// ==========================================
	fastMap := make(map[int]int, elementsCount)

	startFast := time.Now()

	// TODO: Напиши точно такой же цикл для fastMap

	for i := range 5000000 {
		fastMap[i] = i * 2
	}

	durationFast := time.Since(startFast)
	fmt.Printf("Быстрая мапа отработала за:   %v\n", durationFast)
}
