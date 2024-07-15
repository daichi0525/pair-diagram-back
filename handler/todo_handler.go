package handler

import (
	"net/http"

	"github.com/daichi0525/pair-diagram-back.git/usecase"
	"github.com/gin-gonic/gin"
)

type ScheduleHandler struct {
	scheduleUsecase *usecase.ScheduleUsecase
}

func NewScheduleHandler(scheduleUsecase *usecase.ScheduleUsecase) *ScheduleHandler {
	return &ScheduleHandler{
		scheduleUsecase: scheduleUsecase,
	}
}

func (th *ScheduleHandler) GetSchedules(c *gin.Context) {
	// userId := c.Param("user_id")
	schedules, err := th.scheduleUsecase.GetSchedules()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "値がおかしいので確認してください",
		})
	}
	c.JSON(http.StatusOK, schedules)
}
