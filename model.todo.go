package model

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID        uint      `gorm:"primaryKey"`
	UserId    string    `gorm:"not null" json:"user_id"`
	Title     string    `gorm:"type:varchar(100);not null" json:"title"`
	Deadline  time.Time `gorm:"type:timestamp;" json:"deadline"`
	Detail    string    `gorm:"type:varchar(100);" json:"detail"`
	Completed bool      `gorm:"type:boolean;not null" json:"completed"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
