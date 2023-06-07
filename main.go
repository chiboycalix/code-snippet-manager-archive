package main

import (
	"html/template"
	"io/ioutil"

	"github.com/chiboycalix/code-snippet-manager/configs"
	"github.com/chiboycalix/code-snippet-manager/handlers"
	"github.com/chiboycalix/code-snippet-manager/routes"
	"github.com/chiboycalix/code-snippet-manager/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/russross/blackfriday/v2"
)

func main() {

	// Views
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
		// EnablePrintRoutes: true,
	})
	type TemplateData struct {
		MarkdownHTML template.HTML
		Theme        string
	}
	app.Get("/", func(c *fiber.Ctx) error {
		filePath := "./README.md"
		markdown, err := ioutil.ReadFile(filePath)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		renderer := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{})
		html := blackfriday.Run([]byte(markdown), blackfriday.WithRenderer(renderer))

		data := TemplateData{
			MarkdownHTML: template.HTML(html),
			Theme:        "monokai",
		}

		return c.Render("index", data)
	})

	app.Get("/snippets", handlers.GetAllSnippets)

	app.Static("/", "./public", fiber.Static{})
	utils.FormatCode()
	// run database
	configs.ConnectDatabase()

	routes.SnippetRoute(app)

	app.Listen(":3000")
}
