package main

import (
	"github.com/daichi0525/pair-diagram-back.git/db"
	"github.com/daichi0525/pair-diagram-back.git/handler"
	"github.com/daichi0525/pair-diagram-back.git/repository"
	"github.com/daichi0525/pair-diagram-back.git/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	todoHandler := handler.NewTodoHandler(
		usecase.NewTodoUsecase(
			repository.NewTodoRepository(),
		),
	)

	r := gin.Default()
	r.GET("/", todoHandler.GetTodos)
	// r.POST("/login", todoHandler)
	// r.GET("/login", sampleApi)
	r.Run(":8080")
}
