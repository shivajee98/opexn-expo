package handlers

import (
	"errors"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/shivajee98/opexn-exhibitors/internal/dto"
	"github.com/shivajee98/opexn-exhibitors/internal/services"
	"gorm.io/gorm"
)

type StartupHandler struct {
	startupService services.StartupService
}

func InitStartupHandler(startupService services.StartupService) *StartupHandler {
	return &StartupHandler{startupService: startupService}
}

// RegisterStartup handles POST /api/startup/register
func (h *StartupHandler) RegisterStartup(c *fiber.Ctx) error {
	var payload dto.StartupRegistrationPayload
	if err := c.BodyParser(&payload); err != nil {
		log.Printf("RegisterStartup: body parse error: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	// Check for duplicate DPIITCertNumber
	existingStartups, err := h.startupService.GetAllStartups()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("RegisterStartup: fetch error: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Internal error")
	}
	for _, s := range existingStartups {
		if s.DPIITCertNumber == payload.DPIITCertNumber {
			return fiber.NewError(fiber.StatusConflict, "Startup with this DPIIT certificate number already exists")
		}
	}

	startup, err := h.startupService.RegisterStartupService(&payload)
	if err != nil {
		log.Printf("RegisterStartup: transaction failed: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to register startup")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Startup registered successfully",
		"startup": fiber.Map{
			"id":              startup.ID,
			"name":            startup.Name,
			"dpiitCertNumber": startup.DPIITCertNumber,
		},
	})
}

// CheckStartupByDPIITCertNumber handles GET /api/startup/check/:dpiitCertNumber
func (h *StartupHandler) CheckStartupByDPIITCertNumber(c *fiber.Ctx) error {
	dpiitCertNumber := c.Params("dpiitCertNumber")
	if dpiitCertNumber == "" {
		log.Println("CheckStartupByDPIITCertNumber: missing DPIIT certificate number")
		return fiber.NewError(fiber.StatusBadRequest, "DPIIT certificate number is required")
	}

	// Assume DPIITCertNumber is unique and used to check existence
	startups, err := h.startupService.GetAllStartups() // No direct GetByDPIITCertNumber in repo, so filter manually
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("CheckStartupByDPIITCertNumber: error checking startup existence: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Internal error")
	}

	for _, s := range startups {
		if s.DPIITCertNumber == dpiitCertNumber {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"exists": true,
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"exists": false,
	})
}

// GetStartupByID handles GET /api/startup/:id
func (h *StartupHandler) GetStartupByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("GetStartupByID: invalid ID: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid startup ID")
	}

	startup, err := h.startupService.GetStartupByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Startup not found")
		}
		log.Printf("GetStartupByID: error retrieving startup: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Internal error")
	}

	return c.JSON(startup)
}

// GetAllStartups handles GET /api/startup
func (h *StartupHandler) GetAllStartups(c *fiber.Ctx) error {
	startups, err := h.startupService.GetAllStartups()
	if err != nil {
		log.Printf("GetAllStartups: error retrieving startups: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Internal error")
	}

	return c.JSON(startups)
}

// UpdateStartup handles PUT /api/startup/update
// func (h *StartupHandler) UpdateStartup(c *fiber.Ctx) error {
// 	// Step 1: Auth check
// 	clerkIDValue := c.Locals("clerk_id")
// 	clerkID, ok := clerkIDValue.(string)
// 	if !ok || clerkID == "" {
// 		log.Println("UpdateStartup: missing or invalid Clerk ID")
// 		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
// 	}

// 	// Step 2: Parse update request
// 	var updateData model.Startup
// 	if err := c.BodyParser(&updateData); err != nil {
// 		log.Printf("UpdateStartup: body parse error: %v", err)
// 		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
// 	}

// 	// Step 3: Find existing startup
// 	existingStartup, err := h.startupService.GetStartupByID(updateData.ID)
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			log.Printf("UpdateStartup: startup not found for ID %d", updateData.ID)
// 			return fiber.NewError(fiber.StatusNotFound, "Startup not found")
// 		}
// 		log.Printf("UpdateStartup: error retrieving startup: %v", err)
// 		return fiber.NewError(fiber.StatusInternalServerError, "Internal error")
// 	}

// 	// Step 4: Validate and sanitize updatable fields
// 	updateData.Name = sanitizeString(updateData.Name)
// 	if updateData.Name == "" {
// 		return fiber.NewError(fiber.StatusBadRequest, "Name cannot be empty")
// 	}
// 	updateData.PitchDeckURL = sanitizeString(updateData.PitchDeckURL)
// 	if updateData.PitchDeckURL == "" {
// 		return fiber.NewError(fiber.StatusBadRequest, "Pitch deck URL cannot be empty")
// 	}

// 	// Step 5: Update only allowed fields
// 	existingStartup.Name = updateData.Name
// 	existingStartup.WebsiteURL = sanitizeString(updateData.WebsiteURL)
// 	existingStartup.PitchDeckURL = updateData.PitchDeckURL
// 	existingStartup.Address = updateData.Address
// 	existingStartup.Product = updateData.Product
// 	existingStartup.RevenueInfo = updateData.RevenueInfo
// 	existingStartup.FundingInfo = updateData.FundingInfo
// 	existingStartup.EventIntent = updateData.EventIntent
// 	existingStartup.SPOC = updateData.SPOC
// 	existingStartup.Director = updateData.Director

// 	// Step 6: Save updated startup
// 	if err := h.startupService.UpdateStartup(existingStartup); err != nil {
// 		log.Printf("UpdateStartup: failed to update startup: %v", err)
// 		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update startup")
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"message": "Startup updated successfully",
// 		"startup": fiber.Map{
// 			"id":              existingStartup.ID,
// 			"name":            existingStartup.Name,
// 			"dpiitCertNumber": existingStartup.DPIITCertNumber,
// 		},
// 	})
// }

// DeleteStartup handles DELETE /api/startup/:id
func (h *StartupHandler) DeleteStartup(c *fiber.Ctx) error {
	// Step 1: Auth check
	clerkIDValue := c.Locals("clerk_id")
	clerkID, ok := clerkIDValue.(string)
	if !ok || clerkID == "" {
		log.Println("DeleteStartup: missing or invalid Clerk ID")
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	// Step 2: Get ID from params
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("DeleteStartup: invalid ID: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid startup ID")
	}

	// Step 3: Check if startup exists
	_, err = h.startupService.GetStartupByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Startup not found")
		}
		log.Printf("DeleteStartup: error retrieving startup: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Internal error")
	}

	// Step 4: Delete startup
	if err := h.startupService.DeleteStartup(uint(id)); err != nil {
		log.Printf("DeleteStartup: failed to delete startup: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete startup")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Startup deleted successfully",
	})
}

func (h *StartupHandler) GetAllStartupsProducts(c *fiber.Ctx) error {
	startups, err := h.startupService.GetAllStartupsProducts()
	if err != nil {
		log.Printf("GetAllStartupsProducts: failed to retrieve startups: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve startup products")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"startups": startups,
	})
}

func (h *StartupHandler) GetStartupProductByID(c *fiber.Ctx) error {
	startupIDParam := c.Params("id")
	startupID, err := strconv.ParseUint(startupIDParam, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid startup ID")
	}

	product, err := h.startupService.GetStartupProductByID(uint(startupID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Product not found for given startup")
		}
		log.Printf("GetStartupProductByID: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve product")
	}

	return c.Status(fiber.StatusOK).JSON(product)
}
