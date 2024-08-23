package models

import "time"

type User struct {
	ID            int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Username      string    `json:"username" gorm:"unique"`
	Password      string    `json:"password" gorm:"not null"`
	FullName      string    `json:"full_name" gorm:"not null"`
	Email         string    `json:"email" gorm:"unique"`
	IsActive      bool      `json:"is_active" gorm:"default:true"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	DeactivatedAt time.Time `json:"deactivated_at" gorm:"default:null"`
}
