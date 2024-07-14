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

	err = db.AutoMigrate(&model.Todo{})
	if err != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", err))
	}
	// 初期データを挿入
	initializeTodos(db)
	dbInstance = db

}

func initializeTodos(db *gorm.DB) {
	todos := []model.Todo{
		{UserId: "user1", Title: "Todo 1", Deadline: time.Now().AddDate(0, 1, 0), Detail: "Detail 1", Completed: false},
		{UserId: "user2", Title: "Todo 2", Deadline: time.Now().AddDate(0, 1, 1), Detail: "Detail 2", Completed: false},
		{UserId: "user3", Title: "Todo 3", Deadline: time.Now().AddDate(0, 1, 2), Detail: "Detail 3", Completed: false},
		{UserId: "user4", Title: "Todo 4", Deadline: time.Now().AddDate(0, 1, 3), Detail: "Detail 4", Completed: false},
		{UserId: "user5", Title: "Todo 5", Deadline: time.Now().AddDate(0, 1, 4), Detail: "Detail 5", Completed: false},
		{UserId: "user6", Title: "Todo 6", Deadline: time.Now().AddDate(0, 1, 5), Detail: "Detail 6", Completed: false},
		{UserId: "user7", Title: "Todo 7", Deadline: time.Now().AddDate(0, 1, 6), Detail: "Detail 7", Completed: false},
		{UserId: "user8", Title: "Todo 8", Deadline: time.Now().AddDate(0, 1, 7), Detail: "Detail 8", Completed: false},
		{UserId: "user9", Title: "Todo 9", Deadline: time.Now().AddDate(0, 1, 8), Detail: "Detail 9", Completed: false},
		{UserId: "user10", Title: "Todo 10", Deadline: time.Now().AddDate(0, 1, 9), Detail: "Detail 10", Completed: false},
	}

	for _, todo := range todos {
		todo.CreatedAt = time.Now()
		todo.UpdatedAt = time.Now()
		db.Create(&todo)
	}

	fmt.Println("Initial todos inserted!")
}

func GetDbInstantce() *gorm.DB {
	return dbInstance
}
