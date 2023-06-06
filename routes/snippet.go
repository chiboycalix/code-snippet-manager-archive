package routes

import (
	"github.com/chiboycalix/code-snippet-manager/handlers"

	"github.com/gofiber/fiber/v2"
)

func SnippetRoute(app *fiber.App) {
	app.Post("/api/v1/snippets", handlers.CreateSnippet)
	app.Get("/api/v1/snippets", handlers.GetAllSnippets)
	app.Get("/api/v1/snippets/:snippetId", handlers.GetSnippet)
	app.Put("/api/v1/snippets/:snippetId", handlers.UpdateSnippet)
	app.Delete("/api/v1/snippets/:snippetId", handlers.DeleteSnippet)
}
