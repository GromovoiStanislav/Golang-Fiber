package main

import (
	"fiber-example/internal/transport"
	"fiber-example/internal/database"
)


func main() {
	app := transport.Setup()
	database.InitDatabase("books.db")
	app.Listen(":3000")
}