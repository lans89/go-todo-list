package handlers

import (
	todoheader "api-todolist/todo"

	"github.com/gofiber/fiber/v2"
)

func NewTodoList(c *fiber.Ctx) error {
	todoHeader := new(todoheader.ToDoHeaderDto)
	if err := c.BodyParser(todoHeader); err != nil {
		return err
	}
	err := todoheader.SaveHeader(todoHeader)
	if err != nil {
		return err
	}
	return c.JSON(todoHeader)
}

func ListTodoList(c *fiber.Ctx) error {
	return c.JSON([]string{"1", "2", "3"})
}
