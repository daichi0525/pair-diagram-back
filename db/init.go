package db

import (
	"fmt"
	"log"
	"time"

	"github.com/daichi0525/pair-diagram-back.git/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB // グローバルなデータベースインスタンス

// Init 関数はデータベース接続とマイグレーションを初期化します
func Init() {
	// データベース接続文字列 (DSN: Data Source Name)
	dsn := "host=localhost port=5432 user=pair-diagram-postgres password=pair-diagram-postgres dbname=pair-diagram-postgres sslmode=disable"
	var err error
	// データベース接続の確立
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Database connected Errro.")
	}

	// Scheduleテーブルのマイグレーション
	err = db.AutoMigrate(&model.Schedule{})
	if err != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", err))
	}

	// Userテーブルのマイグレーション
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", err))
	}
	// 初期データを挿入
	initializeData(db)
	dbInstance = db
}

// initializeData 関数はデータベースに初期データを挿入します
func initializeData(db *gorm.DB) {
	// スケジュールのテストデータを作成
	schedules := []model.Schedule{
		{UserId: "user1", Title: "Meeting with Bob", Priority: "High", StartTime: time.Now().Add(48 * time.Hour), EndTime: time.Now().Add(50 * time.Hour), Detail: "Discuss project details", Completed: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserId: "user2", Title: "Project deadline", Priority: "Medium", StartTime: time.Now().Add(72 * time.Hour), EndTime: time.Now().Add(74 * time.Hour), Detail: "Submit final project report", Completed: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserId: "user3", Title: "Gym workout", Priority: "Low", StartTime: time.Now().Add(24 * time.Hour), EndTime: time.Now().Add(25 * time.Hour), Detail: "Cardio session", Completed: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserId: "user4", Title: "Doctor appointment", Priority: "High", StartTime: time.Now().Add(120 * time.Hour), EndTime: time.Now().Add(122 * time.Hour), Detail: "Routine check-up", Completed: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserId: "user5", Title: "Team meeting", Priority: "Medium", StartTime: time.Now().Add(96 * time.Hour), EndTime: time.Now().Add(98 * time.Hour), Detail: "Project planning", Completed: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserId: "user6", Title: "Call with client", Priority: "High", StartTime: time.Now().Add(36 * time.Hour), EndTime: time.Now().Add(37 * time.Hour), Detail: "Discuss contract terms", Completed: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserId: "user7", Title: "Dentist appointment", Priority: "Low", StartTime: time.Now().Add(60 * time.Hour), EndTime: time.Now().Add(61 * time.Hour), Detail: "Dental cleaning", Completed: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserId: "user8", Title: "Conference", Priority: "High", StartTime: time.Now().Add(144 * time.Hour), EndTime: time.Now().Add(146 * time.Hour), Detail: "Tech conference", Completed: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserId: "user9", Title: "Lunch with friend", Priority: "Low", StartTime: time.Now().Add(168 * time.Hour), EndTime: time.Now().Add(169 * time.Hour), Detail: "Catch up over lunch", Completed: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserId: "user10", Title: "Finish report", Priority: "Medium", StartTime: time.Now().Add(200 * time.Hour), EndTime: time.Now().Add(202 * time.Hour), Detail: "Finalize quarterly report", Completed: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	// スケジュールのテストデータをデータベースに保存
	for _, schedule := range schedules {
		db.Create(&schedule)
	}

	// ユーザーのテストデータを作成
	users := []model.User{
		{UserName: "Alice", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserName: "Bob", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserName: "Charlie", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserName: "David", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserName: "Eve", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserName: "Frank", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserName: "Grace", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserName: "Hank", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserName: "Ivy", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserName: "Jack", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	// ユーザーのテストデータをデータベースに保存
	for _, user := range users {
		db.Create(&user)
	}
}

// GetDbInstance 関数はグローバルなデータベースインスタンスを返します
func GetDbInstance() *gorm.DB {
	return dbInstance
}
