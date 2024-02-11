package util

import (
	"net/http"

	"github.com/Tracking-Detector/td-backend/go/td_common/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func DefaultFiberApp() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	return app
}

func GetHealth(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(response.NewSuccessResponse("System is running correct."))
}
