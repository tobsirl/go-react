package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hello, World!")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error { return c.SendString("Hello, Worldasdfas!") })

	log.Fatal(app.Listen(":4000"))
}
