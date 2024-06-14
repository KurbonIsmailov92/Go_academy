package models

import "time"

type Task struct {
	Name      string
	Memo      string
	CreatedAt time.Time
	IsDone    bool
}