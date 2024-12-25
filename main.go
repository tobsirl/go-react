package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body	  string `json:"body"`
}

func main() {
	fmt.Println("Hello, World!")
	app := fiber.New()

	todos := []Todo{}

	app.Get("/", func(c *fiber.Ctx) error { return  c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Hello World!"}) })

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		if todo.Body == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Body is required"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(fiber.StatusCreated).JSON(todo)
	})

	log.Fatal(app.Listen(":4000"))
}
