package repository

import (
	"fmt"

	"github.com/daichi0525/pair-diagram-back.git/db"
	"github.com/daichi0525/pair-diagram-back.git/model"
	"gorm.io/gorm"
)

type TodoRepository struct {
	Database *gorm.DB
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{
		Database: db.GetDbInstantce(),
	}
}

func (todoRepo *TodoRepository) GetTodos() ([]model.Todo, error) {
	result := []model.Todo{}
	err := todoRepo.Database.Find(&result)
	if err != nil {
		fmt.Println("Error:", err)
		return result, err.Error
	}
	return result, nil
}
