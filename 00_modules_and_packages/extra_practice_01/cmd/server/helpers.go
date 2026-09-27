package main

import (
	database "tricky_app/internal/db_connector"
)

func LaunchServer() {
	database.Connect()
}
