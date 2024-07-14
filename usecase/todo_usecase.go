package usecase

import (
	"github.com/daichi0525/pair-diagram-back.git/model"
	"github.com/daichi0525/pair-diagram-back.git/repository"
)

type TodoUsecase struct {
	todoRepository *repository.TodoRepository
}

func NewTodoUsecase(todoRepository *repository.TodoRepository) *TodoUsecase {
	return &TodoUsecase{
		todoRepository: todoRepository,
	}
}

func (tu *TodoUsecase) GetTodos() ([]model.Todo, error) {
	result, err := tu.todoRepository.GetTodos()
	if err != nil {
		return result, err
	}
	return result, err
}
