package main

import (
	"fmt"
)

func main() {
	loot := []string{"Золото", "Кинжал", "Броня", "Свиток", "Зелье"}

	var inventory []string

	for _, v := range loot {
		inventory = append(inventory, v)
		fmt.Printf(
			"Appended string is: %v\n"+
				"Current inventory length is: %d\n"+
				"Current capacity is: %d\n\n",
			v, len(inventory), cap(inventory),
		)
	}

}
