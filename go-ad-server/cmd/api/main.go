package main

import (
	"log"

	"go-ad-server/config"
	"go-ad-server/internal/handlers"
	"go-ad-server/internal/repository"
	"go-ad-server/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Config
	cfg := config.LoadConfig()

	// 2. Database
	db := config.ConnectDB(cfg)

	// 3. Repositories
	adRepo := repository.NewAdRepository(db)
	logRepo := repository.NewLogRepository(db)

	// 4. Services
	adService := services.NewAdService(adRepo)
	serveService := services.NewServeService(adRepo, logRepo)
	trackService := services.NewTrackService(adRepo, logRepo)

	// 5. Handlers
	adHandler := handlers.NewAdHandler(adService)
	serveHandler := handlers.NewServeHandler(serveService, trackService)
	statsHandler := handlers.NewStatsHandler(logRepo)

	// 6. Router
	r := gin.Default()

	api := r.Group("/") // Root group or /api/v1? User prompt URLs are /ads, /ad-serve directly.
	// User requested:
	// POST /ads
	// GET /ads
	// PATCH /ads/:id/status
	// GET /ad-serve
	// GET /track/click/:id
	// GET /ads/:id/stats
	// GET /logs

	// Group A: Ad Management
	api.POST("/ads", adHandler.CreateAd)
	api.GET("/ads", adHandler.GetAds)
	api.PATCH("/ads/:id/status", adHandler.UpdateStatus)

	// Group B: Ad Serving & Tracking
	api.GET("/ad-serve", serveHandler.ServeAd)
	api.GET("/track/click/:id", serveHandler.TrackClick)

	// Group C: Analytics
	api.GET("/ads/:id/stats", statsHandler.GetAdStats)
	api.GET("/logs", statsHandler.GetRealTimeLogs)

	// Start
	log.Printf("Server running on port %s", cfg.Port)
	r.Run(":" + cfg.Port)
}
