package rest

import (
	"auth-project/internal/models/response"
	"auth-project/pkg"
	"github.com/gofiber/fiber/v3"
)

func CreateOtp() fiber.Handler {
	return func(c fiber.Ctx) error {
		secretKey, otpImage := pkg.GenerateOtp()
		response := &response.Otp{
			SecretCode:  secretKey,
			OtpInBase64: otpImage,
		}
		return c.Status(200).JSON(response)
	}
}
