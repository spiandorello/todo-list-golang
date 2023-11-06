package services

import (
	"github.com/spiandorello/todo-list-golang/src/interfaces"
	"github.com/spiandorello/todo-list-golang/src/repositories"
)

type Container struct {
	TaskService interfaces.TaskService
}

func GetContainer(repository *repositories.Container) *Container {
	taskService := NewTask(repository)

	return &Container{
		TaskService: taskService,
	}
}
