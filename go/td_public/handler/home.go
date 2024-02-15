package handler

import (
	"github.com/Tracking-Detector/td-backend/go/td_public/model"
	"github.com/Tracking-Detector/td-backend/go/td_public/views/home"
	"github.com/gofiber/fiber/v2"
)

type HomeHandler struct {
	Home *model.Home
}

func NewHomeHandler(home *model.Home) *HomeHandler {
	return &HomeHandler{
		Home: home,
	}
}

func (h *HomeHandler) HandleHomeShow(c *fiber.Ctx) error {
	return Render(c, home.Index(h.Home))
}
