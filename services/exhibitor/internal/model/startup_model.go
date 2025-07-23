package model

import "gorm.io/gorm"

type Startup struct {
	gorm.Model
	Name            string `gorm:"not null"`
	WebsiteURL      string
	DPIITCertNumber string `gorm:"not null;uniqueIndex"`

	AddressID     uint
	Address       Address
	ProductID     uint
	Product       Product
	RevenueInfoID uint
	RevenueInfo   RevenueInfo
	FundingInfoID uint
	FundingInfo   FundingInfo
	EventIntentID uint
	EventIntent   EventIntent
	PitchDeckURL  string `gorm:"not null"` // S3/Cloud URL

	LogoURL   string `gorm:"not null"` // Add this
	BannerURL string // Optional

	SPOCID     uint
	SPOC       SPOC
	DirectorID uint
	Director   Director
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
	Title       string      `gorm:"not null"`
	Description string      `gorm:"type:text;not null"`
	Problem     string      `gorm:"type:text;not null"`
	Stage       string      `gorm:"not null"` // Enum-like
	Users       []*UserType `gorm:"many2many:product_users;"`

	Price       float64 `gorm:"not null"`  // INR
	Quantity    int     `gorm:"not null"`  // Stock
	Category    string  `gorm:"not null"`  // Eg: Health, EduTech
	Tags        string  `gorm:"type:text"` // Comma-separated tags: "AI,ML,Health"
	ProductType string  `gorm:"not null"`  // Enum: Physical, Digital, Service

	Images []ProductImage `gorm:"foreignKey:ProductID"`
}

type ProductImage struct {
	gorm.Model
	ProductID uint   `gorm:"not null"`
	URL       string `gorm:"not null"`
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
