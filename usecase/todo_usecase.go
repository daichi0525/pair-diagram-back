package usecase

import (
	"github.com/daichi0525/pair-diagram-back.git/model"
	"github.com/daichi0525/pair-diagram-back.git/repository"
)

type ScheduleUsecase struct {
	scheduleRepository *repository.ScheduleRepository
}

func NewScheduleUsecase(scheduleRepository *repository.ScheduleRepository) *ScheduleUsecase {
	return &ScheduleUsecase{
		scheduleRepository: scheduleRepository,
	}
}

func (tu *ScheduleUsecase) GetSchedules() ([]model.Schedule, error) {
	result, err := tu.scheduleRepository.GetSchedules()
	if err != nil {
		return result, err
	}
	return result, err
}
