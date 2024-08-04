package usecase

import (
	"time"

	"github.com/daichi0525/pair-diagram-back.git/model"
	"github.com/daichi0525/pair-diagram-back.git/repository"
)

// ScheduleUsecase はスケジュールに関連するビジネスロジックを処理するための構造体です。
type ScheduleUsecase struct {
	scheduleRepository *repository.ScheduleRepository
}

// NewScheduleUsecase は ScheduleUsecase の新しいインスタンスを作成して返します。
// これは依存性注入の一環で、ScheduleRepository を ScheduleUsecase に渡します。
func NewScheduleUsecase(scheduleRepository *repository.ScheduleRepository) *ScheduleUsecase {
	return &ScheduleUsecase{
		scheduleRepository: scheduleRepository,
	}
}

// GetSchedules はスケジュールのリストを取得するためのユースケースメソッドです。
// このメソッドはリポジトリを介してスケジュールデータを取得し、エラーが発生した場合はエラーを返します。
func (tu *ScheduleUsecase) GetSchedules() ([]model.Schedule, error) {
	// スケジュールデータを取得するためにリポジトリの GetSchedules メソッドを呼び出します。
	result, err := tu.scheduleRepository.GetSchedules()

	// エラーが発生した場合、結果とエラーを返します。
	if err != nil {
		return result, err
	}

	// エラーが発生しなかった場合、結果を返します。
	return result, err
}

func (tu *ScheduleUsecase) InsertSchedule(schedule model.Schedule) error {
	return tu.scheduleRepository.InsertSchedule(schedule)
}

// UpdateSchedule
func (tu *ScheduleUsecase) UpdateSchedule(scheduleID uint, title string, startTime time.Time, endTime time.Time) error {
	return tu.scheduleRepository.UpdateSchedule(scheduleID, title, startTime, endTime)
}
