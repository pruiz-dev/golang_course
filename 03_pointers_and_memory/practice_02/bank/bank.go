package deposit

import (
	"fmt"
)

func Deposit(
	balance *int,
	amount int,
) {
	if balance == nil {
		fmt.Println(
			"Критическая ошибка" +
				": счет не найден!",
		)
		return
	}
	*balance += amount
}
