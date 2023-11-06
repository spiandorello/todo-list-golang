package structs

import (
	"time"

	"github.com/google/uuid"
)

type Metadata struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
