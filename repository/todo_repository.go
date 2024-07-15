package repository

import (
	"fmt"

	"github.com/daichi0525/pair-diagram-back.git/db"
	"github.com/daichi0525/pair-diagram-back.git/model"
	"gorm.io/gorm"
)

type ScheduleRepository struct {
	Database *gorm.DB
}

func NewScheduleRepository() *ScheduleRepository {
	return &ScheduleRepository{
		Database: db.GetDbInstance(),
	}
}

func (scheduleRepo *ScheduleRepository) GetSchedules() ([]model.Schedule, error) {
	result := []model.Schedule{}
	err := scheduleRepo.Database.Find(&result)
	if err != nil {
		fmt.Println("Error:", err)
		return result, err.Error
	}
	return result, nil
}
