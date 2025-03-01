package router

import (
	"github.com/alexohneander/dnsaurus/handler"
	"github.com/alexohneander/dnsaurus/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/register", handler.CreateUser)
	auth.Post("/login", handler.Login)

	// User
	user := api.Group("/user", middleware.Protected())
	user.Get("/:id", handler.GetUser)
	user.Patch("/:id", handler.UpdateUser)
	user.Delete("/:id", handler.DeleteUser)

	// Product
	// product := api.Group("/product")
	// product.Get("/", handler.GetAllProducts)
	// product.Get("/:id", handler.GetProduct)
	// product.Post("/", middleware.Protected(), handler.CreateProduct)
	// product.Delete("/:id", middleware.Protected(), handler.DeleteProduct)
}
