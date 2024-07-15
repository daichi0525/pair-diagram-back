package model

import (
	"time"

	"gorm.io/gorm"
)

// Scheduleはスケジュール情報を表します
type Schedule struct {
	ID        uint           `gorm:"primaryKey" json:"schedule_id"`              // 主キーとなるID
	UserId    string         `gorm:"not null" json:"user_id"`                    // スケジュールを作成したユーザーのID
	Title     string         `gorm:"type:varchar(100);not null" json:"title"`    // スケジュールのタイトル
	Priority  string         `gorm:"type:varchar(100);not null" json:"priority"` // スケジュールの優先度
	StartTime time.Time      `gorm:"type:timestamp;not null" json:"start_time"`  // スケジュールの開始時間
	EndTime   time.Time      `gorm:"type:timestamp;not null" json:"end_time"`    // スケジュールの終了時間
	Detail    string         `gorm:"type:varchar(255);" json:"detail"`           // スケジュールの詳細説明
	Completed bool           `gorm:"type:boolean;not null" json:"completed"`     // スケジュールが完了したかどうかのフラグ
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`           // レコードの作成時間（自動生成）
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`           // レコードの更新時間（自動更新）
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`                    // レコードの削除時間（ソフトデリート用）
}

// // 優先度の定数定義
// const (
//     PriorityHigh   = "High"   // 高い優先度
//     PriorityMedium = "Medium" // 中程度の優先度
//     PriorityLow    = "Low"    // 低い優先度
// )
