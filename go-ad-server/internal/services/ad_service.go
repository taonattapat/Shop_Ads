package services

import (
	"go-ad-server/internal/models"
	"go-ad-server/internal/repository"

	"go.mongodb.org/mongo-driver/bson"
)

type AdService struct {
	Repo *repository.AdRepository
}

func NewAdService(repo *repository.AdRepository) *AdService {
	return &AdService{Repo: repo}
}

func (s *AdService) CreateAd(ad *models.Ad) error {
	return s.Repo.Create(ad)
}

func (s *AdService) GetAllAds() ([]models.Ad, error) {
	return s.Repo.GetAds(bson.M{})
}

func (s *AdService) UpdateStatus(id string, status string) error {
	return s.Repo.UpdateStatus(id, status)
}

func (s *AdService) GetAdByID(id string) (*models.Ad, error) {
	return s.Repo.GetByID(id)
}
