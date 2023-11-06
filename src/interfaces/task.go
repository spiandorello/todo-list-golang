package interfaces

import (
	"context"

	"github.com/spiandorello/todo-list-golang/src/dtos"
	"github.com/spiandorello/todo-list-golang/src/structs"
)

type TaskService interface {
	List(ctx context.Context) ([]structs.Task, error)
	Create(ctx context.Context, param dtos.TaskCreateRequest) error
}

type TaskRepository interface {
	Create(ctx context.Context, task structs.Task) error
	GetAll(ctx context.Context) ([]structs.Task, error)
}
