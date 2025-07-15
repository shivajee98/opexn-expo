package services

import (
	"github.com/shivajee98/opexn-exhibitors/internal/model"
	"github.com/shivajee98/opexn-exhibitors/internal/repository"
)

type StartupService interface {
	CreateStartup(startup *model.Startup) error
	GetStartupByID(id uint) (*model.Startup, error)
	GetAllStartups() ([]model.Startup, error)
	UpdateStartup(startup *model.Startup) error
	DeleteStartup(id uint) error
}

type startupService struct {
	startupRepo repository.StartupRepository
}

func InitStartupService(startupRepo repository.StartupRepository) StartupService {
	return &startupService{startupRepo: startupRepo}
}

// CreateStartup implements StartupService.
func (s *startupService) CreateStartup(startup *model.Startup) error {
	// Add business logic/validation here if needed (e.g., validate DPIITCertNumber format)
	return s.startupRepo.CreateStartup(startup)
}

// GetStartupByID implements StartupService.
func (s *startupService) GetStartupByID(id uint) (*model.Startup, error) {
	return s.startupRepo.GetStartupByID(id)
}

// GetAllStartups implements StartupService.
func (s *startupService) GetAllStartups() ([]model.Startup, error) {
	return s.startupRepo.GetAllStartups()
}

// UpdateStartup implements StartupService.
func (s *startupService) UpdateStartup(startup *model.Startup) error {
	// Add business logic/validation here if needed (e.g., ensure PitchDeckURL is valid)
	return s.startupRepo.UpdateStartup(startup)
}

// DeleteStartup implements StartupService.
func (s *startupService) DeleteStartup(id uint) error {
	return s.startupRepo.DeleteStartup(id)
}
