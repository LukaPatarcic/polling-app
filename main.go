package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Static("/", "public")

	app.Post("/create-poll", func(c *fiber.Ctx) error {
		return c.Render("create-poll", fiber.Map{})
	})

	app.Listen(":3000")
}
