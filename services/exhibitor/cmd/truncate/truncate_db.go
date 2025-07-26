package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/shivajee98/opexn-exhibitors/internal/config"
	"github.com/shivajee98/opexn-exhibitors/internal/db"
)

func main() {
	fmt.Println("Starting DB Truncate...")

	cfg := config.LoadEnv()
	fmt.Println("Using database URL:", cfg.DatabaseURL)

	dbConn, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Define the tables to be truncated.
	// GORM defaults to snake_case, pluralized names. Adjust if you use a custom naming strategy.
	tables := []string{
		"product_images",
		"products",
		"user_types",
		"revenue_infos",
		"funding_infos",
		"event_intents",
		"spocs",
		"directors",
		"startups",
		"addresses",
	}

	// Construct the TRUNCATE command.
	// RESTART IDENTITY resets auto-incrementing columns (like primary keys).
	// CASCADE will automatically truncate tables with foreign-key relationships.
	// This avoids errors related to table dependencies.
	truncateSQL := fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE;", strings.Join(tables, ", "))

	// Execute the raw SQL command.
	if err := dbConn.Exec(truncateSQL).Error; err != nil {
		log.Fatalf("Failed to truncate tables: %v", err)
	}

	fmt.Println("Tables truncated successfully.")
}
