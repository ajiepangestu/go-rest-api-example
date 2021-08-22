package middleware

import (
	"github.com/ajiepangestu/go-rest-api-example/config"
	"github.com/gofiber/basicauth"
	"github.com/gofiber/fiber/v2"
)

func AuthReq() func(*fiber.Ctx) {
	cfg := basicauth.Config{
		Users: map[string]string{
			config.Config("USERNAME"): config.Config("PASSWORD"),
		},
	}
	err := basicauth.New(cfg)
	return err
}
