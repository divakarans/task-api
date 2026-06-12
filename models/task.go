package models

import "time"

type Task struct {
	ID        int
	Title     string
	Status    string
	CreatedAt time.Time
}
