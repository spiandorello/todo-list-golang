package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spiandorello/todo-list-golang/src/services"
)

type Handler interface {
	Routes(app *fiber.App)
}

type Handle struct {
	handles []Handler
}

func New(services *services.Container) *Handle {
	return &Handle{
		handles: []Handler{
			NewTask(services),
		},
	}
}

func (h *Handle) GetRoutes(app *fiber.App) {
	for _, h := range h.handles {
		h.Routes(app)
	}
}
