package router

import (
	"github.com/ajiepangestu/go-rest-api-example/handler"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(api fiber.Router) {
	api.Get("/", handler.UserList)
}
