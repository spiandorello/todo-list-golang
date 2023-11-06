package repositories

import (
	"context"

	"github.com/spiandorello/todo-list-golang/src/structs"
	"gorm.io/gorm"
)

type Task struct {
	db *gorm.DB
}

func NewTask(db *gorm.DB) *Task {
	return &Task{
		db: db,
	}
}

func (t Task) GetAll(ctx context.Context) ([]structs.Task, error) {
	var tasks []structs.Task
	tx := t.db.WithContext(ctx).Find(&tasks)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return tasks, nil
}

func (t Task) Create(ctx context.Context, task structs.Task) error {
	tx := t.db.WithContext(ctx).Create(&task)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
