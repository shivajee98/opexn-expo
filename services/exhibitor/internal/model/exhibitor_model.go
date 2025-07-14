package model

import "gorm.io/gorm"

type Exhibitor struct {
	gorm.Model
	Name    string `gorm:"size:200;not null"`
	ClerkID string `gorm:"not null"`
	Phone   string `gorm:"size:20;uniqueIndex;not null"`
}
