package app

import (
	"auth-project/internal/transport/rest"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"log"
)

func Run() {

	app := fiber.New()
	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Cookie",
		AllowMethods:     "GET,POST",
		AllowCredentials: true,
	}))

	api := app.Group("/api")
	rest.AuthRouter(api)
	log.Fatal(app.Listen(":3000"))

}
