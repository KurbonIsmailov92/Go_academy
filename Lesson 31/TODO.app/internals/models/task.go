package models

import "time"

type Task struct {
	ID             int
	Title          string
	Description    string
	TaskPriorityID int
	TaskPriority   string
	UserName       string
	CreatedAt      time.Time
	IsDone         bool
	DoneAt         time.Time
	IsDeleted      bool
	DeletedAt      time.Time
}
