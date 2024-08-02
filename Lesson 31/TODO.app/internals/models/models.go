package models

import "time"

type User struct {
	ID   int
	Name string
}

type Task struct {
	ID             int
	Title          string
	Description    string
	TaskPriorityID int
	TaskPriority   string
	UserID         int
	UserName       string
	CreatedAt      time.Time
	IsDone         bool
	DoneAt         time.Time
	IsDeleted      bool
	DeletedAt      time.Time
}
