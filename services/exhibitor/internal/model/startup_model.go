package model

import "gorm.io/gorm"

type Startup struct {
	gorm.Model
	Name            string `gorm:"not null"`
	WebsiteURL      string
	DPIITCertNumber string `gorm:"not null;uniqueIndex"`
	AddressID       uint
	Address         Address
	ProductID       uint
	Product         Product
	RevenueInfoID   uint
	RevenueInfo     RevenueInfo
	FundingInfoID   uint
	FundingInfo     FundingInfo
	EventIntentID   uint
	EventIntent     EventIntent
	PitchDeckURL    string `gorm:"not null"` // Assuming it's uploaded to cloud
	SPOCID          uint
	SPOC            SPOC
	DirectorID      uint
	Director        Director
}

type Address struct {
	gorm.Model
	Street  string `gorm:"not null"`
	City    string `gorm:"not null"`
	State   string `gorm:"not null"`
	Pincode string `gorm:"not null"`
}

type Product struct {
	gorm.Model
	Description string      `gorm:"type:text;not null"`
	Problem     string      `gorm:"type:text;not null"`
	Stage       string      `gorm:"not null"` // ENUM-like validation at API layer
	Users       []*UserType `gorm:"many2many:product_users;"`
}

type UserType struct {
	gorm.Model
	Label string `gorm:"uniqueIndex;not null"` // "Students", "Teachers", etc.
}

type RevenueInfo struct {
	gorm.Model
	RevenueBracket string `gorm:"not null"` // ENUM-like: ₹0–₹5L, ₹5–₹25L, etc.
	UserImpact     int    `gorm:"not null"`
}

type FundingInfo struct {
	gorm.Model
	Type string `gorm:"not null"` // ENUM-like: Angel, VC, Govt, None, etc.
}

type EventIntent struct {
	gorm.Model
	WhyParticipate string `gorm:"type:text;not null"`
	Expectation    string `gorm:"type:text;not null"`
	ConsentToPay   bool   `gorm:"not null"`
}

type SPOC struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Phone    string `gorm:"not null;uniqueIndex"`
	Position string `gorm:"not null"`
}

type Director struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Email string `gorm:"not null"`
}
