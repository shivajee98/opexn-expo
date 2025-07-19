package main

import (
	"fmt"
	"log"

	"github.com/shivajee98/opexn-exhibitors/internal/config"
	"github.com/shivajee98/opexn-exhibitors/internal/db"
	"github.com/shivajee98/opexn-exhibitors/internal/model"
)

func main() {
	fmt.Println("Starting DB Drop...")

	cfg := config.LoadEnv()

	dbConn, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Drop the tables if they exist
	err = dbConn.Migrator().DropTable(&model.Address{}, &model.Director{}, &model.EventIntent{}, &model.FundingInfo{}, &model.Product{}, &model.RevenueInfo{}, &model.SPOC{}, &model.Startup{}, &model.UserType{})

	fmt.Println("Tables Dropped successfully.")
}
