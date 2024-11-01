package routes

import (
	"data_app/handlers"

	"github.com/gofiber/fiber/v2"
)

var customerHandler handlers.CustomerHandler

func SetRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Get("/customer", customerHandler.Get)
	api.Get("/customer_transaction_summary", customerHandler.GetTransactionSum)

	api.Post("/generate_customers/:count", customerHandler.CreateRandom)
	api.Post("/customer", customerHandler.Create)
}
