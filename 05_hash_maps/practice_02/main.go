package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"Red":   "Красный",
		"Green": "Зеленый",
		"Blue":  "Синий",
		"Black": "Черный",
		"White": "Белый",
	}

	for key, value := range colors {
		fmt.Printf(
			"Цвет %v переводится "+
				"как %v\n",
			key, value,
		)
	}
}
