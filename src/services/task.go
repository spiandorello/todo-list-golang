package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/spiandorello/todo-list-golang/src/dtos"
	"github.com/spiandorello/todo-list-golang/src/interfaces"
	"github.com/spiandorello/todo-list-golang/src/repositories"
	"github.com/spiandorello/todo-list-golang/src/structs"
)

type Task struct {
	taskRepository interfaces.TaskRepository
}

func NewTask(container *repositories.Container) *Task {
	return &Task{
		taskRepository: container.TaskRepository,
	}
}

func (t Task) Get(ctx context.Context, ID string) (structs.Task, error) {
	uid, _ := uuid.Parse(ID)

	return t.taskRepository.GetOne(ctx, uid)
}

func (t Task) List(ctx context.Context) ([]structs.Task, error) {
	return t.taskRepository.GetAll(ctx)
}

func (t Task) Create(ctx context.Context, params dtos.TaskCreateRequest) error {
	return t.taskRepository.Create(ctx, structs.Task{
		Metadata: structs.Metadata{
			ID: uuid.New(),
		},
		Title:       params.Title,
		Description: params.Description,
	})
}

func (t Task) Complete(ctx context.Context, task structs.Task) error {
	task.Completed = true
	task.UpdateAt = time.Now()

	return t.taskRepository.Save(ctx, task)
}
