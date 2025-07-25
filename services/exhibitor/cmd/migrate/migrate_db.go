package main

import (
	"fmt"
	"log"

	"github.com/shivajee98/opexn-exhibitors/internal/config"
	"github.com/shivajee98/opexn-exhibitors/internal/db"
	"github.com/shivajee98/opexn-exhibitors/internal/model"
)

func main() {
	fmt.Println("Starting migration...")

	cfg := config.LoadEnv()
	fmt.Println("Using database URL:", cfg.DatabaseURL)

	dbConn, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// AutoMigrate creates/updates tables for all models in correct order
	err = dbConn.AutoMigrate(
		&model.Address{},
		&model.UserType{},
		&model.RevenueInfo{},
		&model.FundingInfo{},
		&model.EventIntent{},
		&model.SPOC{},
		&model.Director{},
		&model.Startup{},      // ✅ before Product
		&model.Product{},      // ✅ now safe
		&model.ProductImage{}, // ✅ now safe
	)

	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("Migration completed successfully.")
}
