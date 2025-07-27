package repository

import (
	"fmt"

	"github.com/shivajee98/opexn-exhibitors/internal/dto"
	"github.com/shivajee98/opexn-exhibitors/internal/model"
	"gorm.io/gorm"
)

// StartupRepository defines the interface for Startup CRUD operations
type StartupRepository interface {
	GetStartupByID(id uint) (*model.Startup, error)
	GetStartupProductByID(id uint) (*model.Product, error)
	GetAllStartups() ([]model.Startup, error)
	GetAllStartupsProducts() ([]model.Startup, error)
	UpdateStartup(startup *model.Startup) error
	DeleteStartup(id uint) error
	RegisterStartup(payload *dto.StartupRegistrationPayload) (*model.Startup, error)
}

type startupRepository struct {
	db *gorm.DB
}

// InitStartupRepository initializes the repository with a GORM DB instance
func InitStartupRepository(db *gorm.DB) StartupRepository {
	return &startupRepository{db: db}
}
func (r *startupRepository) RegisterStartup(payload *dto.StartupRegistrationPayload) (*model.Startup, error) {
	startup := payload.ToModel()

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&startup.Address).Error; err != nil {
			return err
		}
		if err := tx.Create(&startup.Director).Error; err != nil {
			return err
		}
		if err := tx.Create(&startup.EventIntent).Error; err != nil {
			return err
		}
		if err := tx.Create(&startup.FundingInfo).Error; err != nil {
			return err
		}
		if err := tx.Create(&startup.RevenueInfo).Error; err != nil {
			return err
		}
		if err := tx.Create(&startup.SPOC).Error; err != nil {
			return err
		}
		if err := tx.Create(&startup).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &model.Startup{}, nil
}

// GetStartupByID retrieves a Startup by ID with all related entities
func (r *startupRepository) GetStartupByID(id uint) (*model.Startup, error) {
	var startup model.Startup
	fmt.Println("Debug...............")
	err := r.db.Preload("Address").
		Preload("Products.Users").
		Preload("RevenueInfo").
		Preload("FundingInfo").
		Preload("EventIntent").
		Preload("Products.Images").
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
		Preload("Products.Users").
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
	var product model.Product
	err := r.db.Preload("Users").Preload("Images").First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *startupRepository) GetAllStartupsProducts() ([]model.Startup, error) {
	var startups []model.Startup
	err := r.db.
		Preload("Products").
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
