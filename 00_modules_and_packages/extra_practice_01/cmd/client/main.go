package main

import (
	"fmt"
	database "tricky_app/internal/db_connector"
)

func main() {
	database.Connect()
	fmt.Println("Клиент запущен")
}
