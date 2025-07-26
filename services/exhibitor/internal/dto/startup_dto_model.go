package dto

import (
	"github.com/shivajee98/opexn-exhibitors/internal/model"
)

type StartupRegistrationPayload struct {
	Name            string `json:"name" validate:"required"`
	WebsiteURL      string `json:"websiteURL"`
	DPIITCertNumber string `json:"dpiitCertNumber" validate:"required"`

	PitchDeckURL string  `json:"pitchDeck" validate:"required"`
	LogoURL      string  `json:"logo" validate:"required"`
	BannerURL    *string `json:"banner,omitempty"`

	Address     model.Address     `json:"address" validate:"required"`
	Director    model.Director    `json:"director" validate:"required"`
	EventIntent model.EventIntent `json:"eventIntent" validate:"required"`
	FundingInfo model.FundingInfo `json:"fundingInfo" validate:"required"`
	RevenueInfo model.RevenueInfo `json:"revenueInfo" validate:"required"`
	SPOC        model.SPOC        `json:"spoc" validate:"required"`

	Products []model.Product `json:"products" validate:"required"`
}

func (p *StartupRegistrationPayload) ToModel() *model.Startup {
	return &model.Startup{
		Name:            p.Name,
		WebsiteURL:      p.WebsiteURL,
		DPIITCertNumber: p.DPIITCertNumber,
		PitchDeckURL:    p.PitchDeckURL,
		LogoURL:         p.LogoURL,
		BannerURL:       p.BannerURL,
		Address:         p.Address,
		Director:        p.Director,
		EventIntent:     p.EventIntent,
		FundingInfo:     p.FundingInfo,
		RevenueInfo:     p.RevenueInfo,
		SPOC:            p.SPOC,
		Products:        p.Products,
	}
}
