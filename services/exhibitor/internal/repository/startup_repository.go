package repository

import (
	"github.com/shivajee98/opexn-exhibitors/internal/model"
	"gorm.io/gorm"
)

// StartupRepository defines the interface for Startup CRUD operations
type StartupRepository interface {
	CreateStartup(startup *model.Startup) error
	GetStartupByID(id uint) (*model.Startup, error)
	GetAllStartups() ([]model.Startup, error)
	UpdateStartup(startup *model.Startup) error
	DeleteStartup(id uint) error
}

type startupRepository struct {
	db *gorm.DB
}

// InitStartupRepository initializes the repository with a GORM DB instance
func InitStartupRepository(db *gorm.DB) StartupRepository {
	return &startupRepository{db: db}
}

// CreateStartup creates a new Startup and its related entities
func (r *startupRepository) CreateStartup(startup *model.Startup) error {
	return r.db.Create(startup).Error
}

// GetStartupByID retrieves a Startup by ID with all related entities
func (r *startupRepository) GetStartupByID(id uint) (*model.Startup, error) {
	var startup model.Startup
	err := r.db.Preload("Address").
		Preload("Product.Users").
		Preload("RevenueInfo").
		Preload("FundingInfo").
		Preload("EventIntent").
		Preload("SPOC").
		Preload("Director").
		First(&startup, id).Error
	if err != nil {
		return nil, err
	}
	return &startup, nil
}

// GetAllStartups retrieves all Startups with related entities
func (r *startupRepository) GetAllStartups() ([]model.Startup, error) {
	var startups []model.Startup
	err := r.db.Preload("Address").
		Preload("Product.Users").
		Preload("RevenueInfo").
		Preload("FundingInfo").
		Preload("EventIntent").
		Preload("SPOC").
		Preload("Director").
		Find(&startups).Error
	if err != nil {
		return nil, err
	}
	return startups, nil
}

// UpdateStartup updates a Startup and its related entities
func (r *startupRepository) UpdateStartup(startup *model.Startup) error {
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(startup).Error
}

// DeleteStartup deletes a Startup by ID (soft delete via gorm.Model)
func (r *startupRepository) DeleteStartup(id uint) error {
	return r.db.Delete(&model.Startup{}, id).Error
}
