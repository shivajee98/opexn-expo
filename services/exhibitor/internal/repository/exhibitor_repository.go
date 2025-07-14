package repository

import (
	model "github.com/shivajee98/opexn-exhibitors/internal/model"
	"gorm.io/gorm"
)

type ExhibitorRepository interface {
	RegisterExhibitor(exhibitor *model.Exhibitor) error
}

type exhibitorRepository struct {
	db *gorm.DB
}

func InitExhibitorRepository(db *gorm.DB) ExhibitorRepository {
	return &exhibitorRepository{db: db}
}

func (r *exhibitorRepository) RegisterExhibitor(exhibitor *model.Exhibitor) error {
	return r.db.Create(exhibitor).Error
}
