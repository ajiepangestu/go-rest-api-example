package router

import (
	"github.com/ajiepangestu/go-rest-api-example/handlers"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(api fiber.Router) {
	api.Get("/:id", handlers.GetUser)
	api.Get("/", handlers.GetUserList)
	api.Post("/create", handlers.CreateUser)
}
