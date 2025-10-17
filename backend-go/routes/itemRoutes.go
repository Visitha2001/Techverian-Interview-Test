package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/visitha2001/backend-go/handlers"
)

func RegisterItemRoutes(app *fiber.App, handler *handlers.ItemHandler) {
	api := app.Group("/items")
	api.Post("", handler.CreateItem)
	api.Get("", handler.GetItems)
	api.Get("/:id", handler.GetItem)
	api.Get("/all/summary", handler.GetSummary)
}
