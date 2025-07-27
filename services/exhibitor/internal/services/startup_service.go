package services

import (
	"github.com/shivajee98/opexn-exhibitors/internal/dto"
	"github.com/shivajee98/opexn-exhibitors/internal/model"
	"github.com/shivajee98/opexn-exhibitors/internal/repository"
)

type StartupService interface {
	GetStartupByID(id uint) (*model.Startup, error)
	GetAllStartups() ([]model.Startup, error)
	GetStartupProductByID(id uint) (*model.Product, error)
	GetAllStartupsProducts() ([]model.Startup, error)
	UpdateStartup(startup *model.Startup) error
	DeleteStartup(id uint) error
	RegisterStartupService(payload *dto.StartupRegistrationPayload) (*model.Startup, error)
}

type startupService struct {
	startupRepo repository.StartupRepository
}

func InitStartupService(startupRepo repository.StartupRepository) StartupService {
	return &startupService{startupRepo: startupRepo}
}

// GetStartupByID implements StartupService.
func (s *startupService) GetStartupByID(id uint) (*model.Startup, error) {
	return s.startupRepo.GetStartupByID(id)
}

// GetAllStartups implements StartupService.
func (s *startupService) GetAllStartups() ([]model.Startup, error) {
	return s.startupRepo.GetAllStartups()
}

func (s *startupService) RegisterStartupService(payload *dto.StartupRegistrationPayload) (*model.Startup, error) {
	return s.startupRepo.RegisterStartup(payload)
}

func (s *startupService) GetStartupProductByID(id uint) (*model.Product, error) {
	return s.startupRepo.GetStartupProductByID(id)
}

func (s *startupService) GetAllStartupsProducts() ([]model.Startup, error) {
	return s.startupRepo.GetAllStartupsProducts()
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
