package models

import (
	"time"
)

type Task struct {
	ID             int          `json:"id" gorm:"primary_key"`
	Title          string       `json:"title" gorm:"size:20"`
	Description    string       `json:"description" gorm:"size:30"`
	TaskPriorityID int          `json:"task_priority_id" gorm:"column:priority_id"`
	TaskPriority   TaskPriority `json:"task_priority" gorm:"foreignKey:TaskPriorityID;references:ID"`
	UserID         int          `json:"user_id" gorm:"column:user_id"`
	User           User         `json:"user" gorm:"foreignKey:UserID;references:ID"`
	CreatedAt      time.Time    `json:"created_at"`
	IsDone         bool         `json:"is_done" gorm:"default:false"`
	DoneAt         time.Time    `json:"done_at" gorm:"timestamp with time zone"`
	IsDeleted      bool         `json:"is_deleted" gorm:"default:false"`
	DeletedAt      time.Time    `json:"deleted_at" gorm:"timestamp with time zone"`
}

type TaskPriority struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Priority string `json:"priority" gorm:"varchar(20)"`
}
