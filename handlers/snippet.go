package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/chiboycalix/code-snippet-manager/configs"
	"github.com/chiboycalix/code-snippet-manager/models"
	"github.com/chiboycalix/code-snippet-manager/responses"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var snippetCollection *mongo.Collection = configs.GetCollection(configs.DB, "snippets")
var validate = validator.New()

func GetAllSnippets(c *fiber.Ctx) error {

	// Get all snippets
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Defer canceling context
	defer cancel()
	// Find all snippets
	cursor, err := snippetCollection.Find(ctx, bson.M{})

	// Check if there was an error finding snippets
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	// Defer closing cursor
	var snippets []models.Snippet = make([]models.Snippet, 0)

	// Iterate through cursor
	if err := cursor.All(ctx, &snippets); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	fmt.Println(snippets)
	// Return success response
	return responses.SuccessResponse(c, http.StatusOK, "snippets", snippets)
}

func GetSnippet(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Get id from params
	id := c.Params("snippetId")
	objId, _ := primitive.ObjectIDFromHex(id)
	// Defer canceling context
	defer cancel()
	// Find snippet by id
	var snippet models.Snippet
	err := snippetCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&snippet)

	// Check if there was an error finding snippet
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	// Return success response
	return responses.SuccessResponse(c, http.StatusOK, "snippet", snippet)
}

func CreateSnippet(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Create a new instance of Snippet
	var snippet models.Snippet

	// Defer canceling context
	defer cancel()

	// Parse body into struct
	if err := c.BodyParser(&snippet); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	// Validate struct
	if err := validate.Struct(snippet); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	snippet = models.Snippet{
		ID:       primitive.NewObjectID(),
		Snippet:  snippet.Snippet,
		Language: snippet.Language,
	}

	// Insert new snippet into database

	result, err := snippetCollection.InsertOne(ctx, snippet)
	// Check if there was an error inserting into database
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	// Return success response
	return responses.SuccessResponse(c, http.StatusOK, "snippet", result)
}

func UpdateSnippet(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Get id from params
	id := c.Params("snippetId")
	var snippet models.Snippet
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	// Parse body into struct
	if err := c.BodyParser(&snippet); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	// Validate struct
	if err := validate.Struct(snippet); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	// Update snippet in database
	update := bson.M{"snippet": snippet.Snippet, "language": snippet.Language}
	updatedSnippet, err := snippetCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

	// Check if there was an error updating snippet
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	// Return success response
	return responses.SuccessResponse(c, http.StatusOK, "snippet", updatedSnippet)
}

func DeleteSnippet(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Get id from params
	id := c.Params("snippetId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	// Delete snippet from database
	deletedSnippet, err := snippetCollection.DeleteOne(ctx, bson.M{"_id": objId})

	if deletedSnippet.DeletedCount < 1 {
		return responses.ErrorResponse(c, http.StatusNotFound, err.Error())
	}

	// Check if there was an error deleting snippet
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	// Return success response
	return responses.SuccessResponse(c, http.StatusOK, "snippet", deletedSnippet)
}
