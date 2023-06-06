package main

import (
	"github.com/chiboycalix/code-snippet-manager/configs"
	"github.com/chiboycalix/code-snippet-manager/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {

	// Views
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:             engine,
		ViewsLayout:       "layouts/main",
		EnablePrintRoutes: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})
	app.Get("/snippets", func(c *fiber.Ctx) error {
		return c.Render("snippets", fiber.Map{
			"Snippets": "My snippetss",
		})
	})

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "Hello, World 👋!",
		})
	})

	app.Static("/", "./public", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		MaxAge:    3600,
	})

	// run database
	configs.ConnectDatabase()

	routes.SnippetRoute(app)

	app.Listen(":3000")
}
