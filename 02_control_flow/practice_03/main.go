package main

import (
	"fmt"
)

func main() {
	var mapGrid = [][]string{
		{"O", "O", "O"},
		{"O", "X", "O"},
		{"O", "O", "O"},
	}

SearchLoop:
	for i := range mapGrid {
		for j := range mapGrid[i] {
			if mapGrid[i][j] == "X" {
				fmt.Printf("Сокровище найдено на координатах [%d, %d]\n", i, j)
				break SearchLoop
			}
		}
	}

	fmt.Println("Поиск завершен")
}
