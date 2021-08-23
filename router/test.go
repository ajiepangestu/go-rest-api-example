package router

import "github.com/gofiber/fiber/v2"

func TestRoutes(app *fiber.App) {
	app.Static("/test", "./static/public")
}
