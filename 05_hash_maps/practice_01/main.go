package main

import (
	"fmt"
)

func main() {
	players := make(map[string]int)

	players["Fighter"] = 10
	players["Mage"] = 5

	delete(players, "Mage")

	rogueLvl := players["Rogue"]
	fmt.Printf(
		"players[Rogue] is: %d\n",
		rogueLvl,
	)

	printMapKey(players, "Fighter")
	printMapKey(players, "Mage")

	// if fighterLvl, ok := players["Fighter"]; ok {
	// 	fmt.Printf(
	// 		"Fighter level is: %d\n",
	// 		fighterLvl,
	// 	)
	// } else {
	// 	fmt.Printf(
	// 		"There is no " +
	// 			"Fighter player\n",
	// 	)
	// }

	// if fighterLvl, ok := players["Mage"]; ok {
	// 	fmt.Printf(
	// 		"Mage level is: %d\n",
	// 		fighterLvl,
	// 	)
	// } else {
	// 	fmt.Printf(
	// 		"There is no " +
	// 			"Mage player\n",
	// 	)
	// }
}

func printMapKey(m map[string]int, key string) {
	if lvl, ok := m[key]; ok {
		fmt.Printf(
			"%v level is: %d\n",
			key, lvl,
		)
	} else {
		fmt.Printf(
			"There is no "+
				"%v player\n",
			key,
		)
	}
}
