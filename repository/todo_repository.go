package repository

import (
	"fmt"
	"time"

	"github.com/daichi0525/pair-diagram-back.git/db"
	"github.com/daichi0525/pair-diagram-back.git/model"
	"gorm.io/gorm"
)

// ScheduleRepository はスケジュールデータにアクセスするための構造体です。
// この構造体は、データベースへの参照を保持します。
type ScheduleRepository struct {
	Database *gorm.DB
}

// NewScheduleRepository は ScheduleRepository の新しいインスタンスを作成して返します。
// この関数は、データベースインスタンスを取得し、それを ScheduleRepository に設定します。
func NewScheduleRepository() *ScheduleRepository {
	return &ScheduleRepository{
		Database: db.GetDbInstance(), // データベースインスタンスを取得
	}
}

// GetSchedules はデータベースからすべてのスケジュールを取得するメソッドです。
// このメソッドは、スケジュールのスライスとエラーを返します。
func (scheduleRepo *ScheduleRepository) GetSchedules() ([]model.Schedule, error) {
	// スケジュールのスライスを初期化
	result := []model.Schedule{}

	// データベースからスケジュールを取得し、result に格納
	err := scheduleRepo.Database.Order("user_id asc").Find(&result)

	// エラーが発生した場合
	if err.Error != nil {
		// エラーメッセージをコンソールに出力
		fmt.Println("Error:", err)
		// エラーと取得した結果を返す
		return result, err.Error
	}

	// エラーが発生しなかった場合、結果とnil（エラーなし）を返す
	return result, nil
}

// UpdateSchedule
func (scheduleRepo *ScheduleRepository) UpdateSchedule(scheduleID uint, title string, startTime time.Time, endTime time.Time) error {
	var schedule model.Schedule
	if err := scheduleRepo.Database.First(&schedule, scheduleID).Error; err != nil {
		return err
	}
	schedule.Title = title
	schedule.StartTime = startTime
	schedule.EndTime = endTime
	schedule.UpdatedAt = time.Now()
	return scheduleRepo.Database.Save(&schedule).Error
}
