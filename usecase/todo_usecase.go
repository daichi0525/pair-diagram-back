package usecase

import (
	"time"

	"github.com/daichi0525/pair-diagram-back.git/repository"
	"gorm.io/gorm"
)

type Todo struct {
	ID        uint      `gorm:"primaryKey"`
	UserId    string    `gorm:"not null" json:"user_id"`
	Title     string    `gorm:"type:varchar(100);not null" json:"title"`
	Deadline  time.Time `gorm:"type:timestamp;" json:"deadline"`
	Detail    string    `gorm:"type:varchar(100);" json:"detail"`
	Completed bool      `gorm:"type:boolean;not null" json:"completed"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type TodoUsecase struct {
	todoRepository *repository.TodoRepository
}

func NewTodoUsecase(todoRepository *repository.TodoRepository) *TodoUsecase {
	return &TodoUsecase{
		todoRepository: todoRepository,
	}
}

func (tu *TodoUsecase) GetTodos() (Todo, error) {
	result, err := tu.todoRepository.GetTodos()
	if err != nil {
		// log.Println("failed to GetTodos :", err)
		return nil, err
	}
	return result, err
}
