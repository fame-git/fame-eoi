package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", viewTemplate)

	app.Listen(":8080")
}

func viewTemplate(c *fiber.Ctx) error {
	return c.Render("template", fiber.Map{
		"Name": "World",
	})
}
