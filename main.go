package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {

	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title": "அறுவடை ...",
		})
	})

	app.Get("/cart", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title": "அறுவடை ...",
		})
	})

	app.Get("/order", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "அறுவடை ...",
		})
	})

	log.Fatal(app.Listen(":3000"))

}
