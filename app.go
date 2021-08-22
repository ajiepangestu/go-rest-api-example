package main

import (
	"github.com/ajiepangestu/go-rest-api-example/config"
	"github.com/ajiepangestu/go-rest-api-example/database"
	"github.com/ajiepangestu/go-rest-api-example/handlers"
	"github.com/ajiepangestu/go-rest-api-example/router"

	"flag"
	"fmt"
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
	app.Use(handlers.NotFound)

	// Listen on port 3000
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", config.Config("SERVER_HOST"), config.Config("SERVER_PORT")))) // go run app.go -port=:3000
}
