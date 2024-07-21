package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/daichi0525/pair-diagram-back.git/usecase"
	"github.com/gin-gonic/gin"
)

// ScheduleHandler はスケジュールに関連するリクエストをハンドリングするための構造体です。
type ScheduleHandler struct {
	scheduleUsecase *usecase.ScheduleUsecase
}

// NewScheduleHandler は ScheduleHandler の新しいインスタンスを作成して返します。
// これは依存性注入の一環で、ScheduleUsecase を ScheduleHandler に渡します。
func NewScheduleHandler(scheduleUsecase *usecase.ScheduleUsecase) *ScheduleHandler {
	return &ScheduleHandler{
		scheduleUsecase: scheduleUsecase,
	}
}

// GetSchedules はスケジュールを取得するためのハンドラーです。
// HTTP GET リクエストに応答し、スケジュールのリストを返します。
func (th *ScheduleHandler) GetSchedules(c *gin.Context) {
	// スケジュールを取得するために Usecase のメソッドを呼び出します。
	schedules, err := th.scheduleUsecase.GetSchedules()

	// エラーが発生した場合、クライアントに対して HTTP 400 ステータスコードとエラーメッセージを返します。
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "値がおかしいので確認してください", // 日本語のエラーメッセージ
		})
		return // エラーが発生した場合、これ以上の処理を行わないためにリターンします。
	}

	// 正常にスケジュールを取得できた場合、クライアントに対して HTTP 200 ステータスコードとスケジュールのデータを返します。
	c.JSON(http.StatusOK, schedules)
}

// UpdateSchedule メソッドの追加
func (th *ScheduleHandler) UpdateSchedule(c *gin.Context) {
	var json struct {
		Title     string    `json:"title"`
		StartTime time.Time `json:"start_time"`
		EndTime   time.Time `json:"end_time"`
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	scheduleID := c.Param("id") // URL パラメータからスケジュールIDを取得
	id, err := strconv.ParseUint(scheduleID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なID"})
		return
	}

	err = th.scheduleUsecase.UpdateSchedule(uint(id), json.Title, json.StartTime, json.EndTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新に失敗しました"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "更新に成功しました"})
}
