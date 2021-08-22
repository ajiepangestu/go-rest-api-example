package main

import (
	"github.com/ajiepangestu/go-rest-api-example/config"
	"github.com/ajiepangestu/go-rest-api-example/database"
	"github.com/ajiepangestu/go-rest-api-example/handler"
	"github.com/ajiepangestu/go-rest-api-example/router"

	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var prod = flag.Bool("prod", false, "Enable prefork in Production")

func main() {
	// Parse command-line flags
	flag.Parse()

	// Connected with database
	database.Connect()

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run app.go -prod
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	router.SetupRoutes(app)

	// Handle not founds
	app.Use(handler.NotFound)

	// Listen on port 3000
	log.Fatal(app.Listen(config.Config("SERVER_PORT"))) // go run app.go -port=:3000
}
