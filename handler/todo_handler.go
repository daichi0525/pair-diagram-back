package handler

import (
	"net/http"

	"github.com/daichi0525/pair-diagram-back.git/usecase"
	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	todoUsecase *usecase.TodoUsecase
}

func NewTodoHandler(todoUsecase *usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{
		todoUsecase: todoUsecase,
	}
}

func (th *TodoHandler) GetTodos(c *gin.Context) {
	// userId := c.Param("user_id")
	todos, err := th.todoUsecase.GetTodos()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "値がおかしいので確認してください",
		})
	}
	c.JSON(http.StatusOK, todos)
}
