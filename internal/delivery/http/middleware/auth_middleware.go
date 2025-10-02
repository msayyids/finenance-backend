package middleware

import (
	"strings"

	"finenance-app/internal/utils"

	"github.com/gofiber/fiber/v3"
)

func NewAuth() fiber.Handler {
	return func(c fiber.Ctx) error {
		// Ambil header Authorization
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return fiber.ErrUnauthorized
		}

		// Pastikan formatnya "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return fiber.ErrUnauthorized
		}

		tokenString := parts[1]

		// Validasi token pakai utils.ValidateToken
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			return fiber.ErrUnauthorized
		}

		// Simpan claims ke context biar bisa dipakai di handler berikutnya
		c.Locals("user", claims)

		// Lanjut ke handler berikutnya
		return c.Next()
	}
}
