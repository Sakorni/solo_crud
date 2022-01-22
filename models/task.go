package models

import "time"

type Task struct {
	Id        string    `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
