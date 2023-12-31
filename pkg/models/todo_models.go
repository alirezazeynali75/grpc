package models

import "time"

type ToDo struct {
	ID          int
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
