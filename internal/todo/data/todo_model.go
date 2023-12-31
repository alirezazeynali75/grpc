package data

import "time"

type ToDoModel struct {
	ID          int       `gorm:"primaryKey"`
	Title       string    `gorm:"column:title"`
	Description string    `gorm:"column:description"`
	Status      string    `gorm:"column:status"`
	CreatedAt   time.Time `gorm:"autoCreateTime,column:created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime:nano,column:created_at"`
}
