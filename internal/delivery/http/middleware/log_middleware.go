package middleware

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

func ZapLogger(logger *zap.Logger) fiber.Handler {
	return func(c fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		// default ambil status dari response
		status := c.Response().StatusCode()

		// kalau ada error dan belum diubah statusnya, cek fiber.Error
		if err != nil {
			if e, ok := err.(*fiber.Error); ok {
				status = e.Code
			} else {
				status = fiber.StatusInternalServerError
			}
		}

		latency := time.Since(start)
		method := c.Method()
		path := c.Path()
		ip := c.IP()

		logger.Info("request log",
			zap.String("method", method),
			zap.String("path", path),
			zap.Int("status", status),
			zap.Duration("latency", latency),
			zap.String("ip", ip),
		)

		return err
	}
}
