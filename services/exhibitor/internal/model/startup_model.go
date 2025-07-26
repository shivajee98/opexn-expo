package model

import (
	"gorm.io/gorm"
)

// Note: Use a validation layer or GORM hooks to enforce enum constraints where needed.

type Startup struct {
	gorm.Model
	Name            string `gorm:"not null" json:"name"`
	WebsiteURL      string `json:"websiteURL"`
	DPIITCertNumber string `gorm:"not null;uniqueIndex" json:"dpiitCertNumber"`

	PitchDeckURL string  `gorm:"not null" json:"pitchDeck"`
	LogoURL      string  `gorm:"not null" json:"logo"`
	BannerURL    *string `json:"banner,omitempty"` // optional

	AddressID uint    `json:"-"`
	Address   Address `gorm:"constraint:OnDelete:CASCADE" json:"address"`

	Products []Product `gorm:"constraint:OnDelete:CASCADE" json:"products"`

	RevenueInfoID uint        `json:"-"`
	RevenueInfo   RevenueInfo `gorm:"foreignKey:RevenueInfoID;constraint:OnDelete:CASCADE" json:"revenueInfo"`

	FundingInfoID uint        `json:"-"`
	FundingInfo   FundingInfo `gorm:"foreignKey:FundingInfoID;constraint:OnDelete:CASCADE" json:"fundingInfo"`

	EventIntentID uint        `json:"-"`
	EventIntent   EventIntent `gorm:"foreignKey:EventIntentID;constraint:OnDelete:CASCADE" json:"eventIntent"`

	SPOCID uint `json:"-"`
	SPOC   SPOC `gorm:"foreignKey:SPOCID;constraint:OnDelete:CASCADE" json:"spoc"`

	DirectorID uint     `json:"-"`
	Director   Director `gorm:"foreignKey:DirectorID;constraint:OnDelete:CASCADE" json:"director"`
}

type Address struct {
	gorm.Model
	Street  string `gorm:"not null" json:"street"`
	City    string `gorm:"not null" json:"city"`
	State   string `gorm:"not null" json:"state"`
	Pincode string `gorm:"not null" json:"pincode"`
}

type Product struct {
	gorm.Model
	StartupID   uint           `gorm:"not null" json:"-"`
	Title       string         `gorm:"not null" json:"title"`
	Stage       string         `gorm:"not null" json:"productStage"` // enum in frontend
	Users       []*UserType    `gorm:"many2many:product_users;constraint:OnDelete:CASCADE" json:"userTypes"`
	Price       float64        `gorm:"not null" json:"price"`
	Quantity    int            `gorm:"not null" json:"quantity"`
	Category    string         `gorm:"not null" json:"category"`
	Tags        string         `json:"tags"`                        // comma-separated
	ProductType string         `gorm:"not null" json:"productType"` // Physical, Digital, Service
	Images      []ProductImage `gorm:"foreignKey:ProductID;references:ID;constraint:OnDelete:CASCADE" json:"images"`
}

type ProductImage struct {
	gorm.Model
	ProductID uint   `gorm:"not null;index" json:"-"`
	URL       string `gorm:"not null" json:"url"`
}

type UserType struct {
	gorm.Model
	Label string `gorm:"uniqueIndex;not null" json:"label"` // e.g. "Students", "Teachers"`
}

type RevenueInfo struct {
	gorm.Model
	RevenueBracket string `gorm:"not null" json:"revenueBracket"` // Should be one of revenueBrackets enum from frontend
	UserImpact     int    `gorm:"not null" json:"userImpact"`
}

type FundingInfo struct {
	gorm.Model
	Type string `gorm:"not null" json:"fundingType"` // Should be one of fundingTypes enum from frontend
}

type EventIntent struct {
	gorm.Model
	WhyParticipate string `gorm:"type:text;not null" json:"whyParticipate"`
	Expectation    string `gorm:"type:text;not null" json:"expectation"`
	ConsentToPay   bool   `gorm:"not null" json:"consentToPay"`
}

type SPOC struct {
	gorm.Model
	Name     string `gorm:"not null" json:"Name"` // Note uppercase to match your frontend keys, or you can unify to lowercase
	Email    string `gorm:"not null" json:"Email"`
	Phone    string `gorm:"not null;uniqueIndex" json:"Phone"`
	Position string `gorm:"not null" json:"Position"`
}

type Director struct {
	gorm.Model
	Name  string `gorm:"not null" json:"directorName"`
	Email string `gorm:"not null" json:"directorEmail"`
}
