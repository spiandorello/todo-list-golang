package repositories

import (
	"github.com/spiandorello/todo-list-golang/src/interfaces"
	"gorm.io/gorm"
)

type Container struct {
	TaskRepository interfaces.TaskRepository
}

func GetContainer(db *gorm.DB) *Container {
	taskRepository := NewTask(db)

	return &Container{
		TaskRepository: taskRepository,
	}
}
