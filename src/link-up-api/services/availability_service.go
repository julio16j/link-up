package services

import (
	"context"
	"linkupapi/models"
	"linkupapi/repositories"
	"time"
)

type AvailabilityService struct {
	Repo repositories.AvailabilityRepository
}

func NewAvailabilityService(repo repositories.AvailabilityRepository) *AvailabilityService {
	return &AvailabilityService{Repo: repo}
}

func (s *AvailabilityService) IsAddressAvailable(address models.Address) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.Repo.CheckAvailability(ctx, address.ZipCode)
}
