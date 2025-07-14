package services

import "github.com/shivajee98/opexn-exhibitors/internal/model"

type ExhibitorService interface {
	RegisterExhibitor(exhibitor *model.Exhibitor) error
}