package handlers

import (
	"net/http"

	"go-ad-server/internal/models"
	"go-ad-server/internal/services"

	"github.com/gin-gonic/gin"
)

type AdHandler struct {
	Service *services.AdService
}

func NewAdHandler(service *services.AdService) *AdHandler {
	return &AdHandler{Service: service}
}

func (h *AdHandler) CreateAd(c *gin.Context) {
	var ad models.Ad
	if err := c.ShouldBindJSON(&ad); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateAd(&ad); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, ad)
}

func (h *AdHandler) GetAds(c *gin.Context) {
	// Simple implementation: Get all ads.
	// In real world, filter by user_id if Auth was present.
	ads, err := h.Service.GetAllAds()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ads)
}

func (h *AdHandler) UpdateStatus(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.UpdateStatus(id, req.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "status updated"})
}
