package services

import (
	"github.com/shivajee98/opexn-exhibitors/internal/model"
	"github.com/shivajee98/opexn-exhibitors/internal/repository"
)

type StartupService interface {
	CreateStartup(startup *model.Startup) error
	CreateAddress(address *model.Address) error
	CreateDirector(director *model.Director) error
	CreateEventIntent(eventIntent *model.EventIntent) error
	CreateFundingInfo(fundingInfo *model.FundingInfo) error
	CreateRevenueInfo(revenueInfo *model.RevenueInfo) error
	CreateProduct(product *model.Product) error
	CreateSPOC(spoc *model.SPOC) error
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
	return s.startupRepo.CreateStartup(startup)
}

func (s *startupService) CreateAddress(address *model.Address) error {
	return s.startupRepo.CreateAddress(address)
}

func (s *startupService) CreateDirector(director *model.Director) error {
	return s.startupRepo.CreateDirector(director)
}

func (s *startupService) CreateEventIntent(eventIntent *model.EventIntent) error {
	return s.startupRepo.CreateEventIntent(eventIntent)
}

func (s *startupService) CreateFundingInfo(fundingInfo *model.FundingInfo) error {
	return s.startupRepo.CreateFundingInfo(fundingInfo)
}

func (s *startupService) CreateRevenueInfo(revenueInfo *model.RevenueInfo) error {
	return s.startupRepo.CreateRevenueInfo(revenueInfo)
}

func (s *startupService) CreateProduct(product *model.Product) error {
	return s.startupRepo.CreateProduct(product)
}

func (s *startupService) CreateSPOC(spoc *model.SPOC) error {
	return s.startupRepo.CreateSPOC(spoc)
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
