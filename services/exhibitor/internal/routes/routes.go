package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shivajee98/opexn-exhibitors/internal/handlers"
)

// SetupStartupRoutes registers routes for the StartupHandler
func SetupStartupRoutes(app *fiber.App, startupHandler *handlers.StartupHandler) {
	// Group routes under /api/startup
	startup := app.Group("/api/startup")

	// Public routes
	startup.Get("/check/:dpiitCertNumber", startupHandler.CheckStartupByDPIITCertNumber)
	startup.Get("/:id", startupHandler.GetStartupByID)
	startup.Get("/", startupHandler.GetAllStartups)

	// Protected routes (require Clerk authentication)
	startup.Post("/register", startupHandler.RegisterStartup)
	startup.Put("/update", startupHandler.UpdateStartup)
	startup.Delete("/:id", startupHandler.DeleteStartup)

	app.Get("/api/startup-products", startupHandler.GetAllStartupsProducts)

	app.Get("/api/startup/:id/product", startupHandler.GetStartupProductByID)


}
