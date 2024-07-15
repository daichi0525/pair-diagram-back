package db

import (
	"fmt"
	"log"
	"time"

	"github.com/daichi0525/pair-diagram-back.git/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func Init() {
	dsn := "host=localhost port=5432 user=pair-diagram-postgres password=pair-diagram-postgres dbname=pair-diagram-postgres sslmode=disable"
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Database connected Errro.")
	}

	err = db.AutoMigrate(&model.Schedule{})
	if err != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", err))
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", err))
	}
	// 初期データを挿入
	initializeData(db)
	dbInstance = db

}

func initializeData(db *gorm.DB) {
	// テストデータを作成
	schedules := []model.Schedule{
		{UserId: "user1", Title: "Meeting with Bob", Priority: time.Now().Add(48 * time.Hour), Start: "2024-07-16T09:00:00Z", Required: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserId: "user2", Title: "Project deadline", Priority: time.Now().Add(72 * time.Hour), Start: "2024-07-17T12:00:00Z", Required: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserId: "user3", Title: "Gym workout", Priority: time.Now().Add(24 * time.Hour), Start: "2024-07-15T17:00:00Z", Required: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserId: "user4", Title: "Doctor appointment", Priority: time.Now().Add(120 * time.Hour), Start: "2024-07-18T10:00:00Z", Required: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserId: "user5", Title: "Team meeting", Priority: time.Now().Add(96 * time.Hour), Start: "2024-07-19T15:00:00Z", Required: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserId: "user6", Title: "Call with client", Priority: time.Now().Add(36 * time.Hour), Start: "2024-07-15T13:00:00Z", Required: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserId: "user7", Title: "Dentist appointment", Priority: time.Now().Add(60 * time.Hour), Start: "2024-07-16T11:00:00Z", Required: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserId: "user8", Title: "Conference", Priority: time.Now().Add(144 * time.Hour), Start: "2024-07-20T09:00:00Z", Required: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserId: "user9", Title: "Lunch with friend", Priority: time.Now().Add(168 * time.Hour), Start: "2024-07-21T12:00:00Z", Required: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserId: "user10", Title: "Finish report", Priority: time.Now().Add(200 * time.Hour), Start: "2024-07-22T17:00:00Z", Required: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	// テストデータをデータベースに保存
	for _, schedule := range schedules {
		db.Create(&schedule)
	}

	// テストデータを作成
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
	// テストデータをデータベースに保存
	for _, user := range users {
		db.Create(&user)
	}

}

func GetDbInstantce() *gorm.DB {
	return dbInstance
}
