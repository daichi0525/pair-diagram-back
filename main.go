package main

import (
	"time"

	"github.com/daichi0525/pair-diagram-back.git/db"
	"github.com/daichi0525/pair-diagram-back.git/handler"
	"github.com/daichi0525/pair-diagram-back.git/repository"
	"github.com/daichi0525/pair-diagram-back.git/usecase"
	"github.com/gin-contrib/cors"
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
	// CORSの設定
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(config))

	r.GET("/", todoHandler.GetTodos)
	// r.POST("/login", todoHandler)
	// r.GET("/login", sampleApi)
	r.Run(":8080")
}
