package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/spiandorello/todo-list-golang/src/dtos"
	"github.com/spiandorello/todo-list-golang/src/structs"
)

type TaskService interface {
	Get(ctx context.Context, ID string) (structs.Task, error)
	List(ctx context.Context) ([]structs.Task, error)
	Create(ctx context.Context, param dtos.TaskCreateRequest) error
	Complete(ctx context.Context, task structs.Task) error
}

type TaskRepository interface {
	GetOne(ctx context.Context, ID uuid.UUID) (structs.Task, error)
	Create(ctx context.Context, task structs.Task) error
	GetAll(ctx context.Context) ([]structs.Task, error)
	Save(ctx context.Context, task structs.Task) error
}
