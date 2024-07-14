package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/daichi0525/pair-diagram-back.git/handler"
	"github.com/daichi0525/pair-diagram-back.git/repository"
	"github.com/daichi0525/pair-diagram-back.git/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func main() {
	dsn := "host=127.0.0.1 port=5432 user=pair-diagram-postgres password=pair-diagram-postgres dbname=pair-diagram-postgres sslmode=disable"
	dialector := postgres.Open(dsn)
	var err error
	if Db, err = gorm.Open(dialector); err != nil {
		connect(dialector, 100)
	}
	fmt.Println("db connected!!")

	todoHandler := handler.NewTodoHandler(
		usecase.NewTodoUsecase(
			repository.NewTodoRepository()
		)
	)

	r := gin.Default()
	r.GET("/", todoHandler)
	// r.POST("/login", todoHandler)
	// r.GET("/login", sampleApi)
	r.Run(":8080")
}

func sampleApi(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func connect(dialector gorm.Dialector, count uint) {
	var err error
	if Db, err = gorm.Open(dialector); err != nil {
		if count > 1 {
			time.Sleep(time.Second * 2)
			count--
			fmt.Printf("retry... count:%v\n", count)
			connect(dialector, count)
			return
		}
		panic(err.Error())
	}
}
