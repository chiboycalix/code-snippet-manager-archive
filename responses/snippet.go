package responses

import "github.com/gofiber/fiber/v2"

type SnippetResonse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}

func SuccessResponse(c *fiber.Ctx, status int, entity string, data interface{}) error {
	return c.Status(status).JSON(&SnippetResonse{
		Status:  status,
		Message: "Success",
		Data:    &fiber.Map{entity: data},
	})
}

func ErrorResponse(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(&SnippetResonse{
		Status:  status,
		Message: message,
		Data:    nil,
	})
}
