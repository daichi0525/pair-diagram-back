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
	scheduleHandler := handler.NewScheduleHandler(
		usecase.NewScheduleUsecase(
			repository.NewScheduleRepository(),
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

	r.GET("/", scheduleHandler.GetSchedules)
	r.POST("/", scheduleHandler.InsertSchedule)
	r.PUT("/", scheduleHandler.UpdateSchedule)
	// r.POST("/login", scheduleHandler)
	// r.GET("/login", sampleApi)
	r.Run("localhost:8080")
}
