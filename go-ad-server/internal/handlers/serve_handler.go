package handlers

import (
	"net/http"

	"go-ad-server/internal/services"

	"github.com/gin-gonic/gin"
)

type ServeHandler struct {
	ServeService *services.ServeService
	TrackService *services.TrackService
}

func NewServeHandler(serve *services.ServeService, track *services.TrackService) *ServeHandler {
	return &ServeHandler{
		ServeService: serve,
		TrackService: track,
	}
}

func (h *ServeHandler) ServeAd(c *gin.Context) {
	ip := c.ClientIP()
	userAgent := c.Request.UserAgent()

	ad, err := h.ServeService.ServeAd(ip, userAgent)
	if err != nil {
		// If no ad found, return 204 or 404
		c.JSON(http.StatusNotFound, gin.H{"error": "no active ads available"})
		return
	}

	c.JSON(http.StatusOK, ad)
}

func (h *ServeHandler) TrackClick(c *gin.Context) {
	id := c.Param("id")
	ip := c.ClientIP()
	userAgent := c.Request.UserAgent()

	targetURL, err := h.TrackService.TrackClick(id, ip, userAgent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Redirect to target URL
	c.Redirect(http.StatusFound, targetURL)
}
