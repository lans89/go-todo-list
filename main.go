package main

import (
	"api-todolist/db"
	"api-todolist/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db.InitDB()
	defer db.SqlDB.Close()

	app := fiber.New()
	group := app.Group("/todo")
	group.Post("/", handlers.NewTodoList)
	group.Get("/:owner", handlers.ListTodoList)
	app.Listen(":3000")

}
