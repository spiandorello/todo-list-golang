package structs

import "time"

type Task struct {
	Metadata
	Title       string     `json:"title" gorm:"type:varchar(32);not null"`
	Description string     `json:"description" gorm:"type:varchar(255);not null"`
	DateLimit   *time.Time `json:"date_limit"`
	Completed   bool       `json:"completed" gorm:"type:boolean"`
}
