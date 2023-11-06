package repositories

import (
	"context"

	"github.com/google/uuid"
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

func (t Task) GetOne(ctx context.Context, ID uuid.UUID) (structs.Task, error) {
	var task = structs.Task{Metadata: structs.Metadata{ID: ID}}
	tx := t.db.WithContext(ctx).First(&task, ID)
	if tx.Error != nil {
		return structs.Task{}, tx.Error
	}

	return task, nil
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
	return tx.Error
}

func (t Task) Save(ctx context.Context, task structs.Task) error {
	tx := t.db.WithContext(ctx).Save(&task)

	return tx.Error
}
