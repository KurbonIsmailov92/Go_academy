package models

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Title          string       `gorm:"varchar(12)"`
	Description    string       `gorm:"varchar(30)"`
	TaskPriorityID int          `gorm:"column:priority_id"`
	TaskPriority   TaskPriority `gorm:"foreignKey:TaskPriorityID;references:ID"`
	UserName       string       `gorm:"varchar(12)"`
	IsDone         bool         `gorm:"default:false"`
	DoneAt         time.Time    `gorm:"timestamp with time zone"`
	IsDeleted      bool         `gorm:"default:false"`
}

type TaskPriority struct {
	ID       int    `gorm:"primaryKey"`
	Priority string `gorm:"varchar(20)"`
}
