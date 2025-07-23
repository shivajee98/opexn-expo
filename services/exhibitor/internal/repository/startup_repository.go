package repository

import (
	"github.com/shivajee98/opexn-exhibitors/internal/model"
	"gorm.io/gorm"
)

// StartupRepository defines the interface for Startup CRUD operations
type StartupRepository interface {
	CreateStartup(startup *model.Startup) error
	CreateAddress(startup *model.Address) error
	CreateDirector(startup *model.Director) error
	CreateFundingInfo(startup *model.FundingInfo) error
	CreateSPOC(startup *model.SPOC) error
	CreateProduct(startup *model.Product) error
	CreateRevenueInfo(startup *model.RevenueInfo) error
	CreateEventIntent(startup *model.EventIntent) error
	GetStartupByID(id uint) (*model.Startup, error)
	GetStartupProductByID(id uint) (*model.Product, error)
	GetAllStartups() ([]model.Startup, error)
	GetAllStartupsProducts() ([]model.Startup, error)
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

func (r *startupRepository) CreateAddress(address *model.Address) error {
	return r.db.Create(address).Error
}

func (r *startupRepository) CreateDirector(director *model.Director) error {
	return r.db.Create(director).Error
}

func (r *startupRepository) CreateFundingInfo(fundingInfo *model.FundingInfo) error {
	return r.db.Create(fundingInfo).Error
}

func (r *startupRepository) CreateEventIntent(eventIntent *model.EventIntent) error {
	return r.db.Create(eventIntent).Error
}

func (r *startupRepository) CreateProduct(product *model.Product) error {
	return r.db.Create(product).Error
}

func (r *startupRepository) CreateSPOC(spoc *model.SPOC) error {
	return r.db.Create(spoc).Error
}

func (r *startupRepository) CreateRevenueInfo(revenueInfo *model.RevenueInfo) error {
	return r.db.Create(revenueInfo).Error
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

func (r *startupRepository) GetStartupProductByID(id uint) (*model.Product, error) {
	var startup model.Startup
	err := r.db.
		Preload("Product").
		First(&startup, id).Error

	if err != nil {
		return nil, err
	}

	return &startup.Product, nil
}

func (r *startupRepository) GetAllStartupsProducts() ([]model.Startup, error) {
	var startups []model.Startup
	err := r.db.
		Preload("Product").
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
