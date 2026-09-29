package main

import (
	"fmt"
)

func main() {
	var hp int = 100
	var inCombat bool = false

	fmt.Printf("HP from the "+
		"very beginning: %d\n", hp)

	fmt.Printf("inCombat from "+
		"the very beginning"+
		": %t\n", inCombat)

	TryHeal(hp)
	fmt.Printf("HP after heal: "+
		"%d\n", hp)

	TakeDamage(&hp)
	fmt.Printf("HP after damage:"+
		" %d\n", hp)

	ToggleCombat(&inCombat)
	fmt.Printf("inCombat after "+
		"ToggleCombat: %t\n", inCombat)

	fmt.Printf("HP address "+
		"inside main: %p\n", &hp)

	fmt.Printf("inCombat address"+
		" is: %p\n", &inCombat)
}
