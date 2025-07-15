package main

import (
	"log"
	"os"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/shivajee98/opexn-exhibitors/internal/config"
	"github.com/shivajee98/opexn-exhibitors/internal/db"
	"github.com/shivajee98/opexn-exhibitors/internal/handlers"
	"github.com/shivajee98/opexn-exhibitors/internal/repository"
	"github.com/shivajee98/opexn-exhibitors/internal/routes"
	"github.com/shivajee98/opexn-exhibitors/internal/services"
	// "github.com/shivajee98/opexn-exhibitors/internal/uploader"
	"github.com/shivajee98/opexn-exhibitors/pkg/utils"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Load environment variables
	cfg := config.LoadEnv()
	utils.CheckError("Failed to load environment variables", nil)

	// Initialize Clerk
	clerk.SetKey(cfg.CLERK_SECRET_KEY)

	// Connect to database
	dbConn, err := db.Connect(cfg.DatabaseURL)
	utils.CheckError("Database connection failed", err)

	// Initialize Cloudinary uploader (for PitchDeckURL)
	// cloudinaryUploader := uploader.NewCloudinaryUploader(cfg)

	// Wire dependencies
	startupRepo := repository.InitStartupRepository(dbConn)
	startupService := services.InitStartupService(startupRepo)
	startupHandler := handlers.InitStartupHandler(startupService)

	// Apply CORS middleware
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Content-Type,Authorization",
		AllowOriginsFunc: func(origin string) bool {
			allowed := map[string]bool{
				"https://www.opexn-exhibitors.com": true,
				"http://localhost:3000":            true,
				"http://localhost:3001":            true,
				"http://localhost:3002":            true,
			}
			return allowed[origin]
		},
	}))

	// Setup routes
	routes.SetupStartupRoutes(app, startupHandler)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
