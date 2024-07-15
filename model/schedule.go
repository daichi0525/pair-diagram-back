package model

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	ID        uint      `gorm:"primaryKey" json:"schedule_id"`
	UserId    string    `gorm:"not null" json:"user_id"`
	Title     string    `gorm:"type:varchar(100);not null" json:"title"`
	Priority  time.Time `gorm:"type:timestamp;" json:"deadline"`
	Start     string    `gorm:"type:varchar(100);" json:"detail"`
	Required  bool      `gorm:"type:boolean;not null" json:"completed"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
