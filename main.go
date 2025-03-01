package main

import (
	"log"

	"github.com/alexohneander/dnsaurus/database"
	"github.com/alexohneander/dnsaurus/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	database.ConnectDB()
	router.SetupRoutes(app)

	log.Fatal(app.Listen(":4000"))
}
