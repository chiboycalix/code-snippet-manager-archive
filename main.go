package main

import (
	"github.com/chiboycalix/code-snippet-manager/configs"
	"github.com/chiboycalix/code-snippet-manager/handlers"
	"github.com/chiboycalix/code-snippet-manager/routes"
	"github.com/chiboycalix/code-snippet-manager/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

type Todo struct {
	Object string
	Finish bool
}

func main() {

	// Views
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
		// EnablePrintRoutes: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	app.Get("/snippets", handlers.GetAllSnippets)

	app.Static("/", "./public", fiber.Static{})
	utils.FormatCode()
	// run database
	configs.ConnectDatabase()

	routes.SnippetRoute(app)

	app.Listen(":3000")
}
