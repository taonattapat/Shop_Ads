package services

import (
	"errors"
	"time"

	"go-ad-server/internal/models"
	"go-ad-server/internal/repository"
)

type TrackService struct {
	AdRepo  *repository.AdRepository
	LogRepo *repository.LogRepository
}

func NewTrackService(adRepo *repository.AdRepository, logRepo *repository.LogRepository) *TrackService {
	return &TrackService{
		AdRepo:  adRepo,
		LogRepo: logRepo,
	}
}

func (s *TrackService) TrackClick(adID, ip, userAgent string) (string, error) {
	ad, err := s.AdRepo.GetByID(adID)
	if err != nil {
		return "", err
	}
	if ad == nil {
		return "", errors.New("ad not found")
	}

	// 1. Record Click Log
	err = s.LogRepo.CreateLog(&models.TrackingLog{
		AdID:      ad.ID,
		Type:      "click",
		IP:        ip,
		UserAgent: userAgent,
		Timestamp: time.Now(),
	})
	if err != nil {
		// Log error but proceed? Strict tracking requires this to succeed.
		return "", err
	}

	// 2. Increment Spent (Simulate CPC)
	// Assume 1 unit cost per click for demo
	go s.AdRepo.IncrementSpent(ad.ID, 1.0)

	return ad.TargetURL, nil
}
