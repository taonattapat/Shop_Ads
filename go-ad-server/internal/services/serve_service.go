package services

import (
	"errors"
	"math/rand"
	"time"

	"go-ad-server/internal/models"
	"go-ad-server/internal/repository"
)

type ServeService struct {
	AdRepo  *repository.AdRepository
	LogRepo *repository.LogRepository
}

func NewServeService(adRepo *repository.AdRepository, logRepo *repository.LogRepository) *ServeService {
	return &ServeService{
		AdRepo:  adRepo,
		LogRepo: logRepo,
	}
}

func (s *ServeService) ServeAd(ip, userAgent string) (*models.Ad, error) {
	// 1. Get candidate ads (Active & Budget > Spent)
	ads, err := s.AdRepo.GetActiveAds()
	if err != nil {
		return nil, err
	}
	if len(ads) == 0 {
		return nil, errors.New("no active ads available")
	}

	// 2. Select algorithm: Random Weighted by Remaining Budget
	// Higher remaining budget = higher chance
	selectedAd := s.selectWeightedAd(ads)

	// 3. Async Log View (Fire and forget or wait? "Fastest" implies async but reliability implies sync)
	// For "High Performance" usually we push to a queue. Here we just do it.
	go func() {
		s.LogRepo.CreateLog(&models.TrackingLog{
			AdID:      selectedAd.ID,
			Type:      "view",
			IP:        ip,
			UserAgent: userAgent,
			Timestamp: time.Now(),
		})
	}()

	return &selectedAd, nil
}

func (s *ServeService) selectWeightedAd(ads []models.Ad) models.Ad {
	// Simple random for now if total weight is 0
	if len(ads) == 1 {
		return ads[0]
	}

	var totalWeight float64
	weights := make([]float64, len(ads))

	for i, ad := range ads {
		remaining := ad.Budget - ad.Spent
		if remaining < 0 {
			remaining = 0
		}
		weights[i] = remaining
		totalWeight += remaining
	}

	if totalWeight == 0 {
		return ads[rand.Intn(len(ads))]
	}

	r := rand.Float64() * totalWeight
	for i, w := range weights {
		r -= w
		if r <= 0 {
			return ads[i]
		}
	}

	return ads[len(ads)-1]
}
