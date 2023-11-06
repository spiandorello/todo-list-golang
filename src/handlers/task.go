package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/spiandorello/todo-list-golang/src/dtos"
	"github.com/spiandorello/todo-list-golang/src/interfaces"
	"github.com/spiandorello/todo-list-golang/src/services"
)

type Task struct {
	taskService interfaces.TaskService
}

func NewTask(services *services.Container) *Task {
	return &Task{
		taskService: services.TaskService,
	}
}

func (t Task) Routes(app *fiber.App) {
	router := app.Group("/tasks")

	router.Get("/", t.list)
	router.Post("/", t.create)
	router.Put("/:id/complete", t.complete)
}

func (t Task) list(c *fiber.Ctx) error {
	ctx := c.UserContext()

	data, err := t.taskService.List(ctx)
	if err != nil {
		return c.SendStatus(http.StatusBadGateway)
	}

	return c.JSON(data)
}

func (t Task) create(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var body dtos.TaskCreateRequest
	err := json.Unmarshal(c.Request().Body(), &body)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	err = t.taskService.Create(ctx, body)
	if err != nil {
		return c.SendStatus(http.StatusBadGateway)
	}

	return c.SendStatus(http.StatusCreated)
}

func (t Task) complete(c *fiber.Ctx) error {
	ctx := c.UserContext()

	ID := c.Params("id")
	task, err := t.taskService.Get(ctx, ID)
	if err != nil {
		return c.SendStatus(http.StatusBadGateway)
	}

	err = t.taskService.Complete(ctx, task)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.SendStatus(http.StatusNoContent)
}
