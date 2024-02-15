package main

import (
	"github.com/Tracking-Detector/td-backend/go/td_public/handler"
	"github.com/Tracking-Detector/td-backend/go/td_public/resources"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	res := resources.LoadHomeResource()
	homeHandler := handler.NewHomeHandler(res)
	app.Static("/static", "static")
	app.Get("/", homeHandler.HandleHomeShow)

	app.Listen(":8081")
}
