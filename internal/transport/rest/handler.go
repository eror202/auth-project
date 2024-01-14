package rest

import (
	"github.com/gofiber/fiber/v3"
)

func AuthRouter(app fiber.Router) {
	app.Get("/createOtp", CreateOtp())
}
