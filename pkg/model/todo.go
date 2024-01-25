package model

import "time"

type Todo struct {
	Title       string
	Description string
	DueDate     time.Time
}