package main

import "fmt"

func TryHeal(hp int) {
	hp += 50
}

func TakeDamage(hpPtr *int) {
	*hpPtr -= 30
	fmt.Printf("HP address "+
		"inside TakeDamage "+
		"function: %p\n", hpPtr)
}

func ToggleCombat(inCombatPtr *bool) {
	*inCombatPtr = !*inCombatPtr
}
