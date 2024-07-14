package repository

import (
	"fmt"

	"github.com/daichi0525/pair-diagram-back.git/usecase"
	"gorm.io/gorm"
)

type TodoRepository struct {
	Database *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{
		Database: db,
	}
}

func (todoRepo *TodoRepository) GetTodos() (usecase.Todo, error) {
	result := usecase.Todo{}
	err := todoRepo.Database.Find(&result)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return result, nil
}
