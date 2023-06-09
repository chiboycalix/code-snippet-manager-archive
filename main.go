package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Snippet struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Snippet     string             `json:"snippet" validate:"required"`
	Language    string             `json:"language" validate:"required"`
	Title       string             `json:"title" validate:"required"`
	Description string             `json:"description" validate:"required"`
	Theme       string             `json:"theme" validate:"required"`
}

var (
	db         *mongo.Database
	collection *mongo.Collection
)

func main() {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
		// ViewsLayout: "layouts/main",
		// EnablePrintRoutes: true,
	})
	app.Static("/", "./public", fiber.Static{})
	// Connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/code_snippet_manager"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	db = client.Database("code_snippet_manager")
	collection = db.Collection("snippets")

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	// Define routes
	app.Get("/", indexHandler)
	app.Get("/snippets", snippetsHandler)
	app.Post("/snippets", saveSnippetHandler)

	// Start the server
	fmt.Println("Server started on http://localhost:4000")
	app.Listen(":4000")
}

func indexHandler(c *fiber.Ctx) error {
	snippets := getAllSnippets()
	return c.Render("index", fiber.Map{
		"Snippets": snippets,
		"Theme":    "androidstudio",
	})
}

func snippetsHandler(c *fiber.Ctx) error {
	snippets := getAllSnippets()
	fmt.Println(snippets)
	return c.JSON(snippets)
}

func saveSnippetHandler(c *fiber.Ctx) error {
	snippet := new(Snippet)
	if err := c.BodyParser(snippet); err != nil {
		return fiber.NewError(http.StatusBadRequest, "Invalid request")
	}

	if snippet.Description == "" || snippet.Snippet == "" {
		return fiber.NewError(http.StatusBadRequest, "Name and code are required")
	}

	_, err := collection.InsertOne(context.Background(), snippet)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "Failed to save snippet")
	}

	return c.Redirect("/")
}

func getAllSnippets() []Snippet {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	var snippets []Snippet

	for cursor.Next(context.Background()) {
		var snippet Snippet
		err := cursor.Decode(&snippet)
		if err != nil {
			log.Fatal(err)
		}
		snippets = append(snippets, snippet)
	}

	return snippets
}
