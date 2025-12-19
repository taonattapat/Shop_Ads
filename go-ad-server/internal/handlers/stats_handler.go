package handlers

import (
	"net/http"
	"strconv"

	"go-ad-server/internal/repository"

	"github.com/gin-gonic/gin"
)

type StatsHandler struct {
	LogRepo *repository.LogRepository
}

func NewStatsHandler(logRepo *repository.LogRepository) *StatsHandler {
	return &StatsHandler{LogRepo: logRepo}
}

func (h *StatsHandler) GetAdStats(c *gin.Context) {
	id := c.Param("id")
	stats, err := h.LogRepo.GetStats(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stats)
}

func (h *StatsHandler) GetRealTimeLogs(c *gin.Context) {
	// Optional query param: limit
	limitStr := c.DefaultQuery("limit", "50")
	limit, _ := strconv.ParseInt(limitStr, 10, 64)

	logs, err := h.LogRepo.GetLogs(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, logs)
}
