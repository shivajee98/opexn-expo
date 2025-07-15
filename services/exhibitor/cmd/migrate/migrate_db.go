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
	println(cfg)

	dbConn, err := db.Connect("postgresql://neondb_owner:npg_bgh1NmrauZe2@ep-plain-cherry-a1l9b9gk-pooler.ap-southeast-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// AutoMigrate creates/updates tables for all models
	err = dbConn.AutoMigrate(
		&model.Startup{},
		&model.Address{},
		&model.Product{},
		&model.UserType{},
		&model.RevenueInfo{},
		&model.FundingInfo{},
		&model.EventIntent{},
		&model.SPOC{},
		&model.Director{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}


	fmt.Println("Migration completed successfully.")
}
